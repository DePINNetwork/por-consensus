package config

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/depinnetwork/por-consensus/cmd/cometbft/commands"
	cfg "github.com/depinnetwork/por-consensus/config"
)

func defaultConfigPath(cmd *cobra.Command) string {
	home, err := commands.ConfigHome(cmd)
	if err != nil {
		return ""
	}
	return filepath.Join(home, cfg.DefaultConfigDir, cfg.DefaultConfigFileName)
}
