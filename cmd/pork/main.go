package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nap/git"
	"nap/search"
	"os"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "project Forking Tool for GitHub",
	}

	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(git.CloneCmd)
	//rootCmd.AddCommand(git.DocsCmd)
	//rootCmd.AddCommand(git.ForkCmd)

	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("pork")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
