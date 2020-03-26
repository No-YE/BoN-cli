package main

import (
	"fmt"
	"os"
	"path"

	"github.com/No-YE/BoN-cli/client"
	localFile "github.com/No-YE/BoN-cli/file"
	"github.com/No-YE/BoN-cli/keystore"
	"github.com/spf13/cobra"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "upload markdown file",
		Long:  `upload markdown file to BoN with token`,
		Run: func(cmd *cobra.Command, args []string) {
			mdFile := string(localFile.Read(getMdPath(file)))

			post := client.Post{
				Title:      title,
				Categories: categories,
				Content:    mdFile,
			}
			token := keystore.GetToken()

			client.Upload(post, token)
		},
	}
	title      string
	categories []string
	file       string
)

func init() {
	uploadCmd.Flags().StringVarP(&title, "title", "t", "", "title (required)")
	uploadCmd.Flags().StringSliceVarP(&categories, "categories", "c", []string{}, "categories")
	uploadCmd.Flags().StringVarP(&file, "file", "f", "", "file name (required)")

	uploadCmd.MarkFlagRequired("title")
	uploadCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(uploadCmd)
}

func getMdPath(name string) string {
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path.Join(dirPath, name)
}
