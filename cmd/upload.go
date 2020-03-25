package main

import (
	"github.com/No-YE/BoN-cli/client"
	"github.com/No-YE/BoN-cli/keystore"
	"github.com/spf13/cobra"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "upload markdown file",
		Long:  `upload markdown file to BoN with token`,
		Run: func(cmd *cobra.Command, args []string) {
			post := client.Post{title, categories, "test"}
			token := keystore.GetToken()

			client.Upload(post, token)
		},
	}
	title      string
	categories []string
)

func init() {
	uploadCmd.Flags().StringVarP(&title, "title", "t", "", "title (required)")
	uploadCmd.Flags().StringSliceVarP(&categories, "categories", "c", []string{}, "categories")
	uploadCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(uploadCmd)
}
