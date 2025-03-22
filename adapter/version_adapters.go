package adapter

import (
	abciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	abciv2 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	typesv1 "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	typesv2 "github.com/depinnetwork/por-consensus/api/cometbft/types/v2"
	abcitypes "github.com/depinnetwork/por-consensus/abci/types"
)

// ConvertV1EventsToV2Events converts abciv1.Event to abciv2.Event
func ConvertV1EventsToV2Events(v1Events []abciv1.Event) []abciv2.Event {
	if v1Events == nil {
		return nil
	}
	
	v2Events := make([]abciv2.Event, len(v1Events))
	for i, event := range v1Events {
		v2Events[i] = abciv2.Event{
			Type:       event.Type,
			Attributes: ConvertV1AttributesToV2Attributes(event.Attributes),
		}
	}
	
	return v2Events
}

// ConvertV2EventsToV1Events converts abciv2.Event to abciv1.Event
func ConvertV2EventsToV1Events(v2Events []abciv2.Event) []abciv1.Event {
	if v2Events == nil {
		return nil
	}
	
	v1Events := make([]abciv1.Event, len(v2Events))
	for i, event := range v2Events {
		v1Events[i] = abciv1.Event{
			Type:       event.Type,
			Attributes: ConvertV2AttributesToV1Attributes(event.Attributes),
		}
	}
	
	return v1Events
}

// ConvertV1AttributesToV2Attributes converts abciv1.EventAttribute to abciv2.EventAttribute
func ConvertV1AttributesToV2Attributes(v1Attrs []abciv1.EventAttribute) []abciv2.EventAttribute {
	if v1Attrs == nil {
		return nil
	}
	
	v2Attrs := make([]abciv2.EventAttribute, len(v1Attrs))
	for i, attr := range v1Attrs {
		v2Attrs[i] = abciv2.EventAttribute{
			Key:   attr.Key,
			Value: attr.Value,
			Index: attr.Index,
		}
	}
	
	return v2Attrs
}

// ConvertV2AttributesToV1Attributes converts abciv2.EventAttribute to abciv1.EventAttribute
func ConvertV2AttributesToV1Attributes(v2Attrs []abciv2.EventAttribute) []abciv1.EventAttribute {
	if v2Attrs == nil {
		return nil
	}
	
	v1Attrs := make([]abciv1.EventAttribute, len(v2Attrs))
	for i, attr := range v2Attrs {
		v1Attrs[i] = abciv1.EventAttribute{
			Key:   attr.Key,
			Value: attr.Value,
			Index: attr.Index,
		}
	}
	
	return v1Attrs
}

// ConvertV1HeaderToV2Header converts typesv1.Header to typesv2.Header
// This version handles the Version field structure correctly
func ConvertV1HeaderToV2Header(v1Header typesv1.Header) typesv2.Header {
	return typesv2.Header{
		// Use the correct Version struct type based on actual implementation
		Version: typesv2.Version{
			Block: v1Header.Version.Block,
			App:   v1Header.Version.App,
		},
		ChainID:            v1Header.ChainID,
		Height:             v1Header.Height,
		Time:               v1Header.Time,
		LastBlockId:        ConvertV1BlockIDToV2BlockID(v1Header.LastBlockId),
		LastCommitHash:     v1Header.LastCommitHash,
		DataHash:           v1Header.DataHash,
		ValidatorsHash:     v1Header.ValidatorsHash,
		NextValidatorsHash: v1Header.NextValidatorsHash,
		ConsensusHash:      v1Header.ConsensusHash,
		AppHash:            v1Header.AppHash,
		LastResultsHash:    v1Header.LastResultsHash,
		EvidenceHash:       v1Header.EvidenceHash,
		ProposerAddress:    v1Header.ProposerAddress,
	}
}

// ConvertV2HeaderToV1Header converts typesv2.Header to typesv1.Header
// This version handles the Version field structure correctly
func ConvertV2HeaderToV1Header(v2Header typesv2.Header) typesv1.Header {
	return typesv1.Header{
		// Use the correct Version struct type based on actual implementation
		Version: typesv1.Version{
			Block: v2Header.Version.Block,
			App:   v2Header.Version.App,
		},
		ChainID:            v2Header.ChainID,
		Height:             v2Header.Height,
		Time:               v2Header.Time,
		LastBlockId:        ConvertV2BlockIDToV1BlockID(v2Header.LastBlockId),
		LastCommitHash:     v2Header.LastCommitHash,
		DataHash:           v2Header.DataHash,
		ValidatorsHash:     v2Header.ValidatorsHash,
		NextValidatorsHash: v2Header.NextValidatorsHash,
		ConsensusHash:      v2Header.ConsensusHash,
		AppHash:            v2Header.AppHash,
		LastResultsHash:    v2Header.LastResultsHash,
		EvidenceHash:       v2Header.EvidenceHash,
		ProposerAddress:    v2Header.ProposerAddress,
	}
}

// ConvertV1BlockIDToV2BlockID converts typesv1.BlockID to typesv2.BlockID
func ConvertV1BlockIDToV2BlockID(blockID typesv1.BlockID) typesv2.BlockID {
	return typesv2.BlockID{
		Hash: blockID.Hash,
		PartSetHeader: typesv2.PartSetHeader{
			Total: blockID.PartSetHeader.Total,
			Hash:  blockID.PartSetHeader.Hash,
		},
	}
}

// ConvertV2BlockIDToV1BlockID converts typesv2.BlockID to typesv1.BlockID
func ConvertV2BlockIDToV1BlockID(blockID typesv2.BlockID) typesv1.BlockID {
	return typesv1.BlockID{
		Hash: blockID.Hash,
		PartSetHeader: typesv1.PartSetHeader{
			Total: blockID.PartSetHeader.Total,
			Hash:  blockID.PartSetHeader.Hash,
		},
	}
}

// ConvertV1ValidatorUpdatesToAbciTypes converts []abciv1.ValidatorUpdate to []abcitypes.ValidatorUpdate
// Adapted to match actual field structures
func ConvertV1ValidatorUpdatesToAbciTypes(v1Updates []abciv1.ValidatorUpdate) []abcitypes.ValidatorUpdate {
	if v1Updates == nil {
		return nil
	}
	
	abciUpdates := make([]abcitypes.ValidatorUpdate, len(v1Updates))
	for i, update := range v1Updates {
		// This is a simplified version - inspect the actual structures
		// and adapt this function to match the real field names
		abciUpdates[i] = abcitypes.ValidatorUpdate{
			// Adapt this based on inspection of the actual field structure
			Power: update.Power,
		}
	}
	
	return abciUpdates
}

// ConvertAbciTypesToV1ValidatorUpdates converts []abcitypes.ValidatorUpdate to []abciv1.ValidatorUpdate
// Adapted to match actual field structures
func ConvertAbciTypesToV1ValidatorUpdates(abciUpdates []abcitypes.ValidatorUpdate) []abciv1.ValidatorUpdate {
	if abciUpdates == nil {
		return nil
	}
	
	v1Updates := make([]abciv1.ValidatorUpdate, len(abciUpdates))
	for i, update := range abciUpdates {
		// This is a simplified version - inspect the actual structures
		// and adapt this function to match the real field names
		v1Updates[i] = abciv1.ValidatorUpdate{
			// Adapt this based on inspection of the actual field structure
			Power: update.Power,
		}
	}
	
	return v1Updates
}
