package main

import (
	"gobuyright/pkg/mongo"
	"gobuyright/pkg/mongo/service"
	"gobuyright/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatal("Unable to connect to mongo")
	}
	defer ms.Close()

	u := service.NewIUserService(ms.Copy(), "buyright", "iUser")
	s := server.NewServer(u)

	s.Start()
}
