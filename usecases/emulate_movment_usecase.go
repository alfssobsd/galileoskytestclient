package usecases

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/alfssobsd/galileoskytestclient/gateways/file"
	gatewaynet "github.com/alfssobsd/galileoskytestclient/gateways/net"
	"github.com/alfssobsd/galileoskytestclient/utils"
	"github.com/alfssobsd/galileoskytestclient/utils/protocol"
	"log"
	"os"
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

	timestamp := int(time.Now().Unix())
	for _, element := range config.Route {
		arr := strings.Split(element, ",")
		hexMessage, hexExpectedResponse := makeMessage(arr[0]+";"+arr[1], config.IMEI,
			config.HwVersion, config.FwVersion, string(timestamp), config.DeviceModel, config.HDOP,
			config.SpeedKm, config.HightMeters, config.HWStatus)
		message, _ := hex.DecodeString(hexMessage)
		expectedResponse, _ := hex.DecodeString(hexExpectedResponse)

		log.Printf("Send Lat,Lon = %v, timestamp = %v\n", arr, timestamp)
		err := gatewaynet.SendMessage(conn, message, expectedResponse)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		log.Println("Successful send Lat, Lon = ", arr)

		timestamp += config.IntervalInSeconds
	}
	log.Println("Successful send all data")
}

func makeMessage(latLon string, imei string, hwVersion string, fwVersion string, timestamp string,
	deviceModel string, hdop string, speedKm string, hight string, hwStatus string) (string, string) {

	var message []string
	message = append(message, protocol.TagEncoder(protocol.HWversion, hwVersion))
	message = append(message, protocol.TagEncoder(protocol.FWversion, fwVersion))
	message = append(message, protocol.TagEncoder(protocol.IMEI, imei))
	message = append(message, protocol.TagEncoder(protocol.DeviceID, deviceModel))
	message = append(message, protocol.TagEncoder(protocol.DateTime, timestamp))
	message = append(message, protocol.TagEncoder(protocol.LatLon, "0;7;"+latLon))
	message = append(message, protocol.TagEncoder(protocol.SpeedNDirection, speedKm+";180"))
	message = append(message, protocol.TagEncoder(protocol.HeightMeters, hight))
	message = append(message, protocol.TagEncoder(protocol.HDOP, hdop))
	message = append(message, protocol.TagEncoder(protocol.HWStatus, hwStatus))

	data := strings.Join(message, "")
	dataB, _ := hex.DecodeString(data)

	dataLenB := make([]byte, 2)
	binary.LittleEndian.PutUint16(dataLenB, uint16(len(dataB)))
	dataLen := hex.EncodeToString(dataLenB)

	hexMessage := "01" + dataLen + data

	hexMessageB, _ := hex.DecodeString(hexMessage)
	checkSumUint16 := utils.Crc16CheckSum(hexMessageB)
	checkSubB := make([]byte, 2)
	binary.LittleEndian.PutUint16(checkSubB, checkSumUint16)

	return hexMessage + hex.EncodeToString(checkSubB), "02" + hex.EncodeToString(checkSubB)
}
