package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "gonedrive",
		Short: "OneDrive CLI",
		Long:  `A CLI to interact with items stored in OneDrive`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gonedrive.yaml)")

	rootCmd.AddCommand(
		NewVersionCmd(),
		NewLoginCmd(),
		NewListCmd(),
		NewGetCmd(),
		NewUploadCmd(),
	)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			return
		}
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(workingDir)
		viper.SetConfigName(".gonedrive")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//Execute the root cmd of gonedrive
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
