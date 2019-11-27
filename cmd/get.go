package cmd

import (
	"fmt"
	"net/http"

	"github.com/langered/gonedrive/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [path]",
		Short: "Get the content of a given file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			content, err := service.Get(http.DefaultClient, viper.Get("access_token").(string), args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(content)
		},
	}
}
