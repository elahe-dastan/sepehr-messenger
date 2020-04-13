[![Drone (cloud)](https://img.shields.io/drone/build/elahe-dastan/gossip.svg?style=flat-square)](https://cloud.drone.io/elahe-dastan/gossip)
# Simple chat server 

In this project I implemented a vary simple chat server using golang in which each client writes a message and the
server broadcasts the message to all other clients.

#### How was it implemented? (gRPC)
I have used gRPC to write less, simpler and more efficient code.

gRPC is an open source remote procedure call (RPC) system, using this framework I no longer had to keep the TCP
connection open :smiley: and also wrote much less code.

:speaking_head:Let's get deeper, I had implemented this server char without gRPC and what I had to do was initiating a TCP
connection and keep it open so the server and client can talk to each other using this connection but as you know it wastes our
resources, there may be connections which no one writes to for a long time.

gRPC works in a request/response manner so the
client initiates a connection and sends its request to server then the server responds and the connection will be closed
everything was looking good till I thought about realtime chat :thinking: if a client wants to read its messages so
many times in a short time then the previous open TCP connection manner is better than opening a lot of new connections,
it led me to use gRPC stream keyword which keeps the connection open.To sum it up when a user wants to read its messages
it initiates a TCP connection which will be closed after there is no write action for 10 seconds.
 


#### How to use gRPC?
First of all I had to write the functions I needed in a .proto file and then I could compile the file to go file using

```sh
protoc -I go_compiled_file_directory proto_file_directory --go_out=plugins=grpc:.
``` 


