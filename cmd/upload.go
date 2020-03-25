package main

import (
	"github.com/spf13/cobra"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "upload markdown file",
		Long:  `upload markdown file to BoN with token`,
		Run: func(cmd *cobra.Command, args []string) {
			println(categories[0])
		},
	}
	title      string
	categories []string
)

func init() {
	uploadCmd.Flags().StringVarP(&title, "title", "t", "", "title (required)")
	uploadCmd.Flags().StringArrayVarP(&categories, "categories", "c", []string{"abc"}, "categories")
	uploadCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(uploadCmd)
}
