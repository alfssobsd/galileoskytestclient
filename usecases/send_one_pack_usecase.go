package usecases

import (
	"encoding/hex"
	gatewaynet "github.com/alfssobsd/galileoskytestclient/gateways/net"
	"log"
	"os"
)

func SendOnePackUseCase(host string, port int, hexMessage string, hexExpectedResponse string) {
	log.Println("prepare message for send")
	message, _ := hex.DecodeString(hexMessage)
	expectedResponse, _ := hex.DecodeString(hexExpectedResponse)

	err := gatewaynet.SendOnePack(host, port, message, expectedResponse)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	log.Println("Successful send")
}
