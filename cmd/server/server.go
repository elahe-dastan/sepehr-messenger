package server

import (
	"log"
	"sync"

	"github.com/elahe-dastan/interview/config"
	"github.com/elahe-dastan/interview/protocol"
	"github.com/elahe-dastan/interview/server"

	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Runs server",
		Run: func(cmd *cobra.Command, args []string) {
			c := config.Read()
			s := server.ChatServer {
				Seq:    0,
				Mutex:  sync.Mutex{},
				Queues: make(map[int32]chan protocol.Data),
			}
			if err := s.Start(c); err != nil {
				log.Fatal(err)
			}
		},
	},
	)
}
