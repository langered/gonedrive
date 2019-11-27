package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//NewVersionCmd returns the version cobra command
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Shows current used version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("1.0.0")
		},
	}
}
