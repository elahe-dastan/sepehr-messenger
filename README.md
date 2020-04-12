# Simple chat server 

In this project I implemented a vary simple chat server using golang in which each client writes a message and the <br/>
server broadcasts the message to all other clients. <br/>

#### How was it implemented? (gRPC)
I have used gRPC to write less, simpler and more efficient code. <br/>
gRPC is an open source remote procedure call (RPC) system, using this framework I no longer had to keep the TCP <br/>
connection open :smiley: and also wrote much less code.<br/>
:speaking_head:Let's get deeper, I had implemented this server char without gRPC and what I had to do was initiating a TCP <br/>
connection and keep it open so the server and client can talk to each other using this connection but as you know it wastes our <br/>
resources, there may be connections which no one writes to for a long time. gRPC works in a request/response manner so the<br>
client initiates a connection and sends its request to server then the server responds and the connection will be closed <br/>
everything was looking good till I thought about realtime chat :thinking: if a client wants to read its messages so <br/>
many times in a short time then the previous open TCP connection manner is better than opening a lot of new connections,<br/>
it led me to use gRPC stream keyword which keeps the connection open.To sum it up when a user wants to read its messages <br/>
it initiates a TCP connection which will be closed after there is no write action for 10 seconds.
 


#### How to use gRPC?
First of all I had to write the functions I needed in a .proto file and then I could compile the file to go file using

````
$ protoc -I go_compiled_file_directory proto_file_directory --go_out=plugins=grpc:.
```` 


