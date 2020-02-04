package cmd

import (
	"net/http"

	"github.com/langered/gonedrive/service/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//NewDeleteCmd returns the delete cobra commannd
func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [remote-filepath]",
		Short: "Delete a remote file by the given path",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := azure.AzureClient{}
			client.Delete(http.DefaultClient, viper.Get("access_token").(string), args[0])
		},
	}
}
