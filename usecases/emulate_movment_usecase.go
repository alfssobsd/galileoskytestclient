package usecases

import (
	"encoding/hex"
	"github.com/alfssobsd/galileoskytestclient/gateways/file"
	gatewaynet "github.com/alfssobsd/galileoskytestclient/gateways/net"
	"github.com/alfssobsd/galileoskytestclient/utils/protocol"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func EmulateMovmentUseCase(host string, port int, pathToConfig string) {

	config, err := file.ReadConfigEmulateMovment(pathToConfig)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	conn, err := gatewaynet.Connect(host, port)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	hexPack, hexExpectedResponse := protocol.MakePack(protocol.DataPackEntity{
		HwVersion:   config.HwVersion,
		FwVersion:   config.FwVersion,
		IMEI:        config.IMEI,
		DeviceModel: config.DeviceModel,
	}, true)

	//make bin message for send
	message, _ := hex.DecodeString(hexPack)
	expectedResponse, _ := hex.DecodeString(hexExpectedResponse)

	log.Printf("Send HeadPack\n")
	err = gatewaynet.SendPack(conn, message, expectedResponse)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println("Successful send head message")

	log.Printf("Prepare send MainPacks\n")
	timestamp := int(time.Now().Unix())

	for _, element := range config.Route {
		//split lat, lon
		arr := strings.Split(element, ",")
		//make hex pack
		hexPack, hexExpectedResponse := protocol.MakePack(protocol.DataPackEntity{
			HwVersion:   config.HwVersion,
			FwVersion:   config.FwVersion,
			IMEI:        config.IMEI,
			DeviceModel: config.DeviceModel,
			Timestamp:   strconv.Itoa(timestamp),
			Lat:         arr[0],
			Lon:         arr[1],
			HDOP:        config.HDOP,
			SpeedKm:     config.SpeedKm,
			Hight:       config.HightMeters,
			HwStatus:    config.HWStatus,
		}, false)

		//make bin pack for send
		message, _ := hex.DecodeString(hexPack)
		expectedResponse, _ := hex.DecodeString(hexExpectedResponse)

		log.Printf("Send Lat,Lon = %v, timestamp = %v\n", arr, timestamp)
		err := gatewaynet.SendPack(conn, message, expectedResponse)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		log.Println("Successful send Lat, Lon = ", arr)

		timestamp += config.IntervalInSeconds
	}
	log.Println("Successful send all data")
}
