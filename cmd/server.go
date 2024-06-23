package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"phonebook/pkg/bootstrap"
)

func Execute() {
	root := &cobra.Command{
		Use:     "Phonebook",
		Short:   "serve the Phone Book",
		Version: "0.1",
	}

	bootstrap.Server()

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
