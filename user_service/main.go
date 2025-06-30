package main

import (
	"log"
	"micro_services/user_service/config"
	"micro_services/user_service/handler"
	"micro_services/user_service/proto"
	"micro_services/user_service/utils"
)

func main() {
	db, err := utils.DBConnect(config.DBName)

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}

	gServer, listener, err := utils.GrpcServer(config.Port, nil)

	if err != nil {
		log.Fatalf("Failed to start gPRC server: %v", err)
		return
	}

	// 注册某一个 gRPC 的 server
	proto.RegisterUserServer(gServer, &handler.User{
		DB: db,
	})

	if err := gServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
		return
	}
}
