package server

import (
	"github.com/elahe-dastan/interview/config"
	"github.com/elahe-dastan/interview/server"

	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Runs server",
		Run: func(cmd *cobra.Command, args []string) {
			c := config.Read()
			server.Start(c)
		},
	},
	)
}
