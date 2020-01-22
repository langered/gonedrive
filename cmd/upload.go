package cmd

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/crypto"
	"github.com/langered/gonedrive/service/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//NewUploadCmd returns the upload cobra command
func NewUploadCmd() *cobra.Command {
	var encryption bool

	uploadCMD := &cobra.Command{
		Use:   "upload [remote-filepath] [content as string]",
		Short: "Upload a stdin to onedrive by into the given file",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			client := azure.AzureClient{}
			remoteFilePath := args[0]
			content := args[1]

			if encryption {
				var err error
				content, err = encryptContent(content)
				if err != nil {
					fmt.Println("Encryption failed: ", err)
					return
				}
			}

			success, err := client.Upload(http.DefaultClient, viper.Get("access_token").(string), remoteFilePath, content)
			if err != nil {
				fmt.Println(err)
				return
			}

			if success {
				fmt.Println("content got uploaded to:", remoteFilePath)
			} else {
				fmt.Println("upload failed")
			}
		},
	}
	uploadCMD.PersistentFlags().BoolVarP(&encryption, "encrypt", "e", false, "encrypt the content before uploading")

	return uploadCMD
}

func encryptContent(content string) (string, error) {
	fmt.Println("Content will be encrypted before uploading")
	password := promptForPassword()
	uploadContent, err := crypto.Encrypt(content, password)
	if err != nil {
		return "", err
	}
	return uploadContent, nil
}
