package cmd

import (
	"fmt"
	"net/http"
	"syscall"

	"github.com/langered/gonedrive/crypto"
	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

//NewGetCmd returns the get cobra command
func NewGetCmd() *cobra.Command {
	var encryption bool

	getCMD := &cobra.Command{
		Use:   "get [path]",
		Short: "Get the content of a given file as stdout",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			content, err := service.Get(http.DefaultClient, viper.Get("access_token").(string), args[0])

			if encryption {
				content, err = decryptContent(content)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(content)
		},
	}
	getCMD.PersistentFlags().BoolVarP(&encryption, "encrypt", "e", false, "decrypt encrypted content")

	return getCMD
}

func decryptContent(content string) (string, error) {
	fmt.Println("Content is encrypted. Enter the password to decrypt it:")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	uploadContent, err := crypto.Decrypt(content, password)
	if err != nil {
		return "", err
	}
	return uploadContent, nil
}
