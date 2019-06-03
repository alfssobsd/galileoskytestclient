package main

import (
	"github.com/alfssobsd/galileoskytestclient/usecases"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Test Client GalileoSky 7.0")

	//usecases.SendOneMessageUseCase("127.0.0.1", 9998,
	//	"011780011a02e703383634343935303330383631333033043200a3af", "02a3af")
	usecases.EmulateMovmentUseCase("127.0.0.1", 9998, "client_v7_test_movement.yml")
}
