package merkle

import (
	"bytes"
	"errors"
	"fmt"

	cmtcrypto "github.com/depinnetwork/por-consensus/api/cometbft/crypto/v1"
)

var ErrKeyPathNotConsumed = errors.New("merkle: keypath not consumed")

// ----------------------------------------
// ProofOp gets converted to an instance of ProofOperator:

// ProofOperator is a layer for calculating intermediate Merkle roots
// when a series of Merkle trees are chained together.
// Run() takes leaf values from a tree and returns the Merkle
// root for the corresponding tree. It takes and returns a list of bytes
// to allow multiple leaves to be part of a single proof, for instance in a range proof.
// ProofOp() encodes the ProofOperator in a generic way so it can later be
// decoded with OpDecoder.
type ProofOperator interface {
	Run(leaves [][]byte) ([][]byte, error)
	GetKey() []byte
	ProofOp() cmtcrypto.ProofOp
}

// ----------------------------------------
// Operations on a list of ProofOperators

// ProofOperators is a slice of ProofOperator(s).
// Each operator will be applied to the input value sequentially
// and the last Merkle root will be verified with already known data.
type ProofOperators []ProofOperator

func (poz ProofOperators) VerifyValue(root []byte, keypath string, value []byte) (err error) {
	return poz.Verify(root, keypath, [][]byte{value})
}

// Verify applies a series of ProofOperators to verify the provided args (byte slices).
func (poz ProofOperators) Verify(root []byte, keypath string, args [][]byte) error {
	keys, err := KeyPathToKeys(keypath)
	if err != nil {
		return err
	}

	for i, op := range poz {
		key := op.GetKey()
		if len(key) != 0 {
			if len(keys) == 0 {
				return ErrInvalidKey{
					Err: fmt.Errorf("key path has insufficient # of parts: expected no more keys but got %+v", string(key)),
				}
			}
			lastKey := keys[len(keys)-1]
			if !bytes.Equal(lastKey, key) {
				return ErrInvalidKey{
					Err: fmt.Errorf("key mismatch on operation #%d: expected %+v but got %+v", i, string(lastKey), string(key)),
				}
			}
			keys = keys[:len(keys)-1]
		}
		args, err = op.Run(args)
		if err != nil {
			return err
		}
	}
	if !bytes.Equal(root, args[0]) {
		return ErrInvalidHash{
			Err: fmt.Errorf("root %x, want %x", args[0], root),
		}
	}
	if len(keys) != 0 {
		return ErrKeyPathNotConsumed
	}
	return nil
}

// ----------------------------------------
// ProofRuntime - main entrypoint

type OpDecoder func(cmtcrypto.ProofOp) (ProofOperator, error)

type ProofRuntime struct {
	decoders map[string]OpDecoder
}

func NewProofRuntime() *ProofRuntime {
	return &ProofRuntime{
		decoders: make(map[string]OpDecoder),
	}
}

// RegisterOpDecoder registers a new OpDecoder for a specific proof operation type.
func (prt *ProofRuntime) RegisterOpDecoder(typ string, dec OpDecoder) {
	_, ok := prt.decoders[typ]
	if ok {
		panic("already registered for type " + typ)
	}
	prt.decoders[typ] = dec
}

// Decode decodes a cmtcrypto.ProofOp into a ProofOperator using the appropriate decoder.
func (prt *ProofRuntime) Decode(pop cmtcrypto.ProofOp) (ProofOperator, error) {
	decoder := prt.decoders[pop.Type]
	if decoder == nil {
		return nil, ErrInvalidProof{
			Err: fmt.Errorf("unrecognized proof type %v", pop.Type),
		}
	}
	return decoder(pop)
}

// DecodeProof decodes a list of cmtcrypto.ProofOps into a ProofOperators slice.
func (prt *ProofRuntime) DecodeProof(proof *cmtcrypto.ProofOps) (ProofOperators, error) {
	poz := make(ProofOperators, 0, len(proof.Ops))
	for _, pop := range proof.Ops {
		operator, err := prt.Decode(pop)
		if err != nil {
			return nil, ErrInvalidProof{
				Err: fmt.Errorf("decoding a proof operator: %w", err),
			}
		}
		poz = append(poz, operator)
	}
	return poz, nil
}

func (prt *ProofRuntime) VerifyValue(proof *cmtcrypto.ProofOps, root []byte, keypath string, value []byte) (err error) {
	return prt.Verify(proof, root, keypath, [][]byte{value})
}

// TODO In the long run we'll need a method of classification of ops,
// whether existence or absence or perhaps a third?
func (prt *ProofRuntime) VerifyAbsence(proof *cmtcrypto.ProofOps, root []byte, keypath string) (err error) {
	return prt.Verify(proof, root, keypath, nil)
}

// Verify verifies a proof by decoding it into ProofOperators.
func (prt *ProofRuntime) Verify(proof *cmtcrypto.ProofOps, root []byte, keypath string, args [][]byte) (err error) {
	poz, err := prt.DecodeProof(proof)
	if err != nil {
		return ErrInvalidProof{
			Err: fmt.Errorf("decoding proof: %w", err),
		}
	}
	return poz.Verify(root, keypath, args)
}

// DefaultProofRuntime only knows about value proofs.
// To use e.g. IAVL proofs, register op-decoders as
// defined in the IAVL package.
func DefaultProofRuntime() (prt *ProofRuntime) {
	prt = NewProofRuntime()
	prt.RegisterOpDecoder(ProofOpValue, ValueOpDecoder)
	return prt
}
