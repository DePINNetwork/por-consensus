package types

import (
	"github.com/depinnetwork/por-consensus/crypto/bls12381"
	"github.com/depinnetwork/por-consensus/crypto/ed25519"
	"github.com/depinnetwork/por-consensus/crypto/secp256k1eth"
	cmtmath "github.com/depinnetwork/por-consensus/libs/math"
)

// MaxSignatureSize is a maximum allowed signature size for the Proposal
// and Vote.
// XXX: secp256k1 does not have max signature size defined.
var MaxSignatureSize = cmtmath.MaxInt(
	ed25519.SignatureSize,
	cmtmath.MaxInt(secp256k1eth.SignatureLength, bls12381.SignatureLength))

// Signable is an interface for all signable things.
// It typically removes signatures before serializing.
// SignBytes returns the bytes to be signed
// NOTE: chainIDs are part of the SignBytes but not
// necessarily the object themselves.
// NOTE: Expected to panic if there is an error marshaling.
type Signable interface {
	SignBytes(chainID string) []byte
}
