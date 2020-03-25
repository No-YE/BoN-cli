package main

import (
	"github.com/No-YE/BoN-cli/keystore"
	"github.com/spf13/cobra"
)

var (
	tokenCmd = &cobra.Command{
		Use:   "token",
		Short: "register token",
		Long:  `register jwt token for BoN-cli`,
		Run: func(cmd *cobra.Command, args []string) {
			keystore.SaveToken(token)
		},
	}
	token string
)

func init() {
	tokenCmd.Flags().StringVarP(&token, "token", "t", "", "jwt token (required)")
	tokenCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(tokenCmd)
}
