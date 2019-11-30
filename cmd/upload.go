package cmd

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//NewUploadCmd returns the upload cobra command
func NewUploadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "upload [remote-filepath] [content as string]",
		Short: "Upload a string to onedrive by the given path and into the given file",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			remoteFilePath := args[0]
			content := args[1]

			success, err := service.Upload(http.DefaultClient, viper.Get("access_token").(string), remoteFilePath, content)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(success)
		},
	}
}
