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
		Use:   "upload [remote-path] [remote-filename] [content as string]",
		Short: "Upload a string to onedrive by the given path and into the given file",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			remotePath := args[0]
			remoteFilename := args[1]
			content := args[2]

			success, err := service.Upload(http.DefaultClient, viper.Get("access_token").(string), remotePath, remoteFilename, content)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(success)
		},
	}
}
