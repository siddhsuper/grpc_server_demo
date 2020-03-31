FROM golang:1.14
RUN mkdir -p /tmp/src/github.com
ADD grpc_server_demo /tmp/src/github.com/grpc_server_demo
ADD grpc_hello_world /tmp/src/github.com/grpc_hello_world
RUN GOPATH=/tmp && \
    cd /tmp/src/github.com/grpc_server_demo && \
    go get &&\
    go build
EXPOSE 8090
WORKDIR /tmp/src/github.com/grpc_server_demo
CMD ["./grpc_server_demo"]

