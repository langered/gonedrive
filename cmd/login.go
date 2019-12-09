package cmd

import (
	"fmt"

	"net/http"

	gonedriveBrowser "github.com/langered/gonedrive/browser"
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
			accessToken, err := service.Login(http.DefaultClient, gonedriveBrowser.GonedriveBrowser{})
			if err != nil {
				fmt.Println(err)
				return
			}
			writeAccessTokenToConfig(accessToken)
		},
	}
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
