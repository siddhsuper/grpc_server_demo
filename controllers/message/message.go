package currency

import (
	"fmt"
	"net/http"

	grpc "google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/grpc_hello_world/helloworld"
	"github.com/grpc_server_demo/config"
	"github.com/grpc_server_demo/types"
)

func GetMessages(c *gin.Context) {

	gctx := types.ConvertContext(c)
	gctx.INFO("--GetMessages Function Initiated --")

	// Set up a connection to the server.
	conn, err := grpc.Dial(config.Config.GRPC_Port, grpc.WithInsecure())
	if err != nil {
		gctx.FATAL(err, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	cHello := helloworld.NewGreeterClient(conn)

	message := gctx.GetUrlParam("message")
	if message == "" || message == "null" {
		gctx.CUSTOMERROR("Invalid Request ", http.StatusBadRequest)
		return
	}

	r, err := cHello.SayHello(c, &helloworld.HelloRequest{Name: message})
	if err != nil {
		gctx.FATAL(err, http.StatusInternalServerError)
		return
	}

	gctx.INFO("--GetMessages Function Executed --")
	gctx.OK(r.GetMessage(), "Successfully fetched data")
	return
}
