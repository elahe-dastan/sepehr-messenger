package client

import (
	"alibaba/client"
	"alibaba/protocol"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "client",
		Short: "A brief description of your application",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			cli := client.New()
			id, err := cli.Who(context.Background(), &empty.Empty{})

			if err != nil {
				log.Fatal(err)
			}

			consoleReader := bufio.NewReader(os.Stdin)

			for {
				fmt.Print(" >> ")
				text, err := consoleReader.ReadString('\n')
				if err != nil {
					log.Println(err)
				}

				if text == "show\n" {
					res, err := cli.Receive(context.Background(), id)

					if err != nil {
						log.Println(err)
					}

					for  {
						m, err := res.Recv()

						if err != nil {
							break
						}

						fmt.Printf("id:%d > %s",m.Id.Id, m.Text)
					}

				}else {
					cli.Send(context.Background(), &protocol.Data{
						Id:                   id,
						Text:                 text,
					})
				}
			}
		},
	},
	)
}
