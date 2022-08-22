/*
Copyright © 2022 karampok

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/karampok/keepme/xurl"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keepme",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	 Run: func(cmd *cobra.Command, args []string) { 
    if err:=xurl.URLToText(); err!=nil{
      fmt.Println(err)
  }
  },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
