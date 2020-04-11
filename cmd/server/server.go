package server

import (
	"alibaba/server"

	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "A brief description of your application",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			//s := server.New()
			server.Start()
		},
	},
	)
}
