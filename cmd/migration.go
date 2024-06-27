package cmd

import (
	"github.com/spf13/cobra"
	"phonebook/internal/database/migration"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Table migration",
	Long:  "All tables will be migrated on username, password, host and port defined in config.yml ",
	Run: func(cmd *cobra.Command, args []string) {
		migration.Migration()
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
