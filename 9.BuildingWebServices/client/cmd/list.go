/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		URL := "http://" + SERVER + ":" + PORT + "/list"

		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("List of entries:")
		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
