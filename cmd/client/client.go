package client

import (
	"alibaba/client"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "client",
		Short: "A brief description of your application",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			c, err := net.Dial("tcp", "localhost:1373")

			if err != nil {
				log.Println(err)
			}

			cli := client.New(c)

			go cli.Show()

			consoleReader := bufio.NewReader(os.Stdin)

			for {
				fmt.Print(" >> ")
				text, err := consoleReader.ReadString('\n')
				if err != nil {
					log.Println(err)
				}

				cli.Send(text)
			}
		},
	},
	)
}
