package cmd

import (
	"fmt"

	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to OneDrive",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Login to:\n\n\n", service.AuthURL(), "\n\n")
			fmt.Print("Enter the token: ")
			var input string
			fmt.Scanln(&input)

			accessToken := input
			err := service.ValidateToken(accessToken)
			if err != nil {
				fmt.Println("Invalid token")
				fmt.Println(err)
			}

			fmt.Println(accessToken)
		},
	}
}
