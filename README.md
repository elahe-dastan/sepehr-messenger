# Simple chat server 

In this project I implemented a vary simple chat server using golang in which each client writes a message and the <br/>
server broadcasts the message to all other clients. <br/>

#### How was it implemented? (gRPC)
I have used gRPC to write less, simpler and more efficient code. <br/>
gRPC is an open source remote procedure call (RPC) system, using this framework I no longer had to keep the TCP <br/>
connection open :smiley: and also wrote much less code.

#### How to use gRPC?
First of all I had to write the functions I needed in a .proto file and then I could compile the file to go file using

````
$ protoc -I go_compiled_file_directory proto_file_directory --go_out=plugins=grpc:.
```` 


