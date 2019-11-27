package cmd

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list [path]",
		Short: "List items under given path",
		Long:  "List items under given path, when no path given it will list items on the root level",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := ""
			if len(args) == 1 {
				path = args[0]
			}
			items, err := service.ListItems(http.DefaultClient, path, viper.Get("access_token").(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(items)
		},
	}
}
