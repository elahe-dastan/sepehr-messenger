package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/gossip/cmd/client"
	"github.com/elahe-dastan/gossip/cmd/server"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "gossip",
		Short: "Chat wit gRPC",
	}

	server.Register(rootCmd)
	client.Register(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
