package server

import (
	"log"

	"github.com/elahe-dastan/gossip/config"
	"github.com/elahe-dastan/gossip/server"

	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Runs server",
		Run: func(cmd *cobra.Command, args []string) {
			c := config.Read()
			s := server.NewChatServer()
			if err := s.Start(c); err != nil {
				log.Fatal(err)
			}
		},
	},
	)
}
