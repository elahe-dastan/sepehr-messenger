package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/interview/cmd/client"
	"github.com/elahe-dastan/interview/cmd/server"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "alibaba",
		Short: "A brief description of your application",
	}

	server.Register(rootCmd)
	client.Register(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
