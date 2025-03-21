package api

// Import all API packages to ensure they're built correctly
import (
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/crypto/v1"
	_ "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	// Add other imports as needed
)

// Placeholder function to avoid unused import errors
func UseAllAPIPackages() {}
