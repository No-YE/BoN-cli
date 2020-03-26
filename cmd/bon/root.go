package main

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "bon",
		Short: "A CRUD client for BoN",
		Long:  `BoN-cli is a client for BoN CRUD`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
