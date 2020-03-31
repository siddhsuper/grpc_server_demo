# grpc_hello_world
GRPC Hello World

#Install grpc
go get -u google.golang.org/grpc

#Install the protoc Go plugin
go get -u github.com/golang/protobuf/protoc-gen-go

#To generate the GRPC Proto File
protoc -I . helloworld.proto --go_out=plugins=grpc:.

#Server
go run main.go

#Client
go run main.go

Create a folder github.com inside github.com copy paste grpc_hello_world/ grpc_server_demo/