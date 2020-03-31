package main

import (
	"os"

	"github.com/antigloss/go/logger"
	"github.com/gin-gonic/gin"
	"github.com/grpc_server_demo/config"
	"github.com/grpc_server_demo/router"
)

func main() {
	if _, err := os.Stat("./log_files"); os.IsNotExist(err) {
		os.Mkdir("./log_files", 0755)
	}
	// logger.Init must be called first to setup logger
	logger.Init("./log_files", // specify the directory to save the logfiles
		100,   // maximum logfiles allowed under the specified log directory
		50,    // number of logfiles to delete when number of logfiles exceeds the configured limit
		5,     // maximum size of a logfile in MB
		false) // whether logs with Trace level are written down
	log_msg := ""
	log_msg += "Initiated Main function....\n"
	r := gin.Default()

	router.GRPCServiceRouter(r)
	log_msg += "Server Running on port :" + config.Config.Server_Port
	logger.Info(log_msg)
	err := r.Run(":" + config.Config.Server_Port)
	if err != nil {
		logger.Error("Error while starting the server with Port No : " + config.Config.Server_Port + " Error Msg : " + err.Error())
		panic(err)
	}

}
