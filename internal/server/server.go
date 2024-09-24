package server

import (
	"log"
	"net"
	"user-service/internal/config"
)

func ConnGrpc(cfg config.Config) {
	lis, err := net.Listen("tcp", cfg.USER_SERVER)
	if err != nil {
		log.Fatal("Unable to listen: ", err)
	}
	defer lis.Close()

}
