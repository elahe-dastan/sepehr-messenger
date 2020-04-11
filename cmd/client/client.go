package client

import (
	"alibaba/client"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"

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

			time.Sleep(time.Second * 2)

			for {
				fmt.Println("send")
				consoleReader := bufio.NewReader(os.Stdin)
				fmt.Print(" >> ")
				text, er := consoleReader.ReadString('\n')
				if er != nil {
					log.Println(er)
				}

				if text == "show\n" {
					fmt.Println(<- cli.Incoming)
				}else {
					fmt.Println(text)
					cli.Writer.WriteString(text)
					cli.Writer.Flush()
				}
			}

		},
	},
	)
}
