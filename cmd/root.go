package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "onedrive",
		Short: "OneDrive CLI",
		Long:  `OneDrive CLI long`,
	}

	rootCmd.AddCommand(
		NewVersionCmd(),
	)

	return rootCmd
}

func Execute() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
