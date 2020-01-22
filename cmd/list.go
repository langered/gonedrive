package cmd

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//NewListCmd returns the list cobra command
func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list [path]",
		Short: "List items under given path",
		Long:  "List items under given path, when no path given it will list items on the root level",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := azure.AzureClient{}
			path := ""
			if len(args) == 1 {
				path = args[0]
			}
			items, err := client.List(http.DefaultClient, viper.Get("access_token").(string), path)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(items)
		},
	}
}
