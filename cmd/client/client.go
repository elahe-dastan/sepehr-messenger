package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elahe-dastan/interview/client"
	"github.com/elahe-dastan/interview/protocol"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	c := cobra.Command{
		Use:   "client",
		Short: "Runs client",
		Run: func(cmd *cobra.Command, args []string) {
			cli := client.New()
			id, err := cli.Who(context.Background(), &empty.Empty{})

			if err != nil {
				log.Fatal(err)
			}

			consoleReader := bufio.NewReader(os.Stdin)

			go client.Show(cli, id)

			for {
				fmt.Print(" >> ")
				text, err := consoleReader.ReadString('\n')
				if err != nil {
					log.Println(err)
				}

				cli.Send(context.Background(), &protocol.Data{
					Id:                   id,
					Text:                 text,
				})
			}
		},
	}

	c.Flags().StringP("server", "s", "127.0.0.1:1373","server address")
	rootCmd.AddCommand(&c,)
}
