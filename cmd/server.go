package cmd

import (
	"github.com/spf13/cobra"
	"phonebook/pkg/bootstrap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.Server()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
