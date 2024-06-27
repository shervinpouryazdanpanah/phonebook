package cmd

import (
	"github.com/spf13/cobra"
	"phonebook/internal/database/seeder"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Generate admin seed",
	Long:  `Generate admin seed with username:admin and password:admin`,
	Run: func(cmd *cobra.Command, args []string) {
		seeder.Seed()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
