package main

import (
	"os"

	"github.com/lalo64/SmartEnv-api/src/kafka"
	"github.com/lalo64/SmartEnv-api/src/server"
)

var (
	HOST = os.Getenv("HOST_SERVER")
	PORT = os.Getenv("PORT_SERVER")
)

func main(){
	go kafka.Consumer()
	srv := server.NewServer(HOST, PORT)
	srv.Run()
}