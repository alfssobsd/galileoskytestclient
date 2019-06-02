package usecases

import (
	"encoding/hex"
	gatewaynet "github.com/alfssobsd/galileoskytestclient/gateways/net"
	"log"
)

func SendOneMessageUseCase(host string, port int, hexMessage string, hexExpectedResponse string) {
	log.Println("prepare message for send")
	message, _ := hex.DecodeString(hexMessage)
	expectedResponse, _ := hex.DecodeString(hexExpectedResponse)

	err := gatewaynet.SendOneMessage(host, port, message, expectedResponse)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successful send")
}
