package cmd

import (
	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to OneDrive",
		Run: func(cmd *cobra.Command, args []string) {
			service.Login()
		},
	}
}
