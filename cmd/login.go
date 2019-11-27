package cmd

import (
	"fmt"

	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//NewLoginCmd returns the login cobra command
func NewLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to OneDrive",
		Run: func(cmd *cobra.Command, args []string) {
			accessToken := enterAccessToken()
			valid := validateEnteredToken(accessToken)
			if !valid {
				return
			}
			writeAccessTokenToConfig(accessToken)
		},
	}
}

func enterAccessToken() string {
	var input string
	fmt.Println("Login to:\n\n\n", service.AuthURL())
	fmt.Print("\n\nEnter the token: ")
	fmt.Scanln(&input)
	return input
}

func validateEnteredToken(token string) bool {
	err := service.ValidateToken(token)
	if err != nil {
		fmt.Println("Invalid token")
		fmt.Println(err)
		return false
	}
	return true
}

func writeAccessTokenToConfig(accessToken string) {
	viper.Set("access_token", accessToken)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Access token stored in config file")
}
