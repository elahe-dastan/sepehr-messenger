package client

import (
	"alibaba/protocol"
	"log"

	"google.golang.org/grpc"
)

//type Client struct {
//	//Con      net.Conn
//	//Writer   *bufio.Writer
//	//Reader   *bufio.Reader
//	//Incoming chan string
//}

func New() protocol.SimpleChatClient {
	conn, err := grpc.Dial("localhost:1373", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	return protocol.NewSimpleChatClient(conn)

}

//func (cli *Client) Send(text string)  {
//	//cli.Writer.WriteString(text)
//	//cli.Writer.Flush()
//}


//
//func (cli *Client) Show()  {
//	for {
//		text, err := cli.Reader.ReadString('\n')
//
//		if err != nil {
//			log.Println(err)
//		}
//
//		cli.Incoming<- text
//
//		fmt.Println(text)
//	}
//}

// API GRPC
// 2 days