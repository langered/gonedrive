package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/langered/gonedrive/service/azure"
	"github.com/langered/gonedrive/service/secret"
)

//NewSecretCmd returns the secret cobra command
func NewSecretCmd() *cobra.Command {
	secretCMD := &cobra.Command{
		Use:   "secret",
		Short: "Work with secrets",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	secretCMD.AddCommand(newInitSubCmd())
	secretCMD.AddCommand(newPushSubCmd())
	secretCMD.AddCommand(newGetSubCmd())
	return secretCMD
}

func newPushSubCmd() *cobra.Command {
	var credenitalName string
	var credenitalValue string

	pushCMD := &cobra.Command{
		Use:   "push",
		Short: "Push a secret",
		Run: func(cmd *cobra.Command, args []string) {
			if !viper.IsSet("secret_path") {
				fmt.Println("Please execute 'secret init' first")
				return
			}
			password := promptForPassword()
			client := azure.AzureClient{}
			err := secret.Push(client, viper.Get("access_token").(string), password, credenitalName, credenitalValue, viper.Get("secret_path").(string))
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}
	pushCMD.Flags().StringVarP(&credenitalName, "name", "n", "", "name of the credential")
	pushCMD.MarkFlagRequired("name")
	pushCMD.Flags().StringVarP(&credenitalValue, "value", "v", "", "value of the credential")
	pushCMD.MarkFlagRequired("value")
	return pushCMD
}

func newGetSubCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [secret-name]",
		Short: "Get a secret by a given name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if !viper.IsSet("secret_path") {
				fmt.Println("Please execute 'secret init' first")
				return
			}
			password := promptForPassword()
			client := azure.AzureClient{}
			secret, err := secret.Get(client, viper.Get("access_token").(string), password, args[0], viper.Get("secret_path").(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(secret)
		},
	}
}

func newInitSubCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init [path]",
		Short: "Set the .gdsecret path to store and write secrets",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if filepath.Ext(args[0]) != ".gdsecret" {
				fmt.Println("The stated file is not a secret file for gonedrive:", args[0])
				return
			}
			viper.Set("secret_path", args[0])
			viper.WriteConfig()
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}
}
