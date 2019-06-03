package protocol

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/alfssobsd/galileoskytestclient/utils"
	"strings"
)

type DataPackEntity struct {
	HwVersion   string
	FwVersion   string
	IMEI        string
	DeviceModel string
	Timestamp   string
	Lat         string
	Lon         string
	HDOP        string
	SpeedKm     string
	Hight       string
	HwStatus    string
}

func MakePack(input DataPackEntity, isHeadPack bool) (string, string) {

	var tags []string
	tags = append(tags, TagEncoder(HWversion, input.HwVersion))
	tags = append(tags, TagEncoder(FWversion, input.FwVersion))
	tags = append(tags, TagEncoder(IMEI, input.IMEI))
	tags = append(tags, TagEncoder(DeviceID, input.DeviceModel))
	if isHeadPack == false {
		tags = append(tags, TagEncoder(DateTime, input.Timestamp))
		tags = append(tags, TagEncoder(LatLon, "0;7;"+input.Lat+";"+input.Lon))
		tags = append(tags, TagEncoder(SpeedNDirection, input.SpeedKm+";180"))
		tags = append(tags, TagEncoder(HeightMeters, input.Hight))
		tags = append(tags, TagEncoder(HDOP, input.HDOP))
		tags = append(tags, TagEncoder(HWStatus, input.HwStatus))
	}

	dataPack := strings.Join(tags, "")
	dataBinPack, _ := hex.DecodeString(dataPack)

	dataLenB := make([]byte, 2)
	binary.LittleEndian.PutUint16(dataLenB, uint16(len(dataBinPack)))
	hexLenDataBinPack := hex.EncodeToString(dataLenB)

	hexPack := "01" + hexLenDataBinPack + dataPack

	hexBinPack, _ := hex.DecodeString(hexPack)
	checkSumUint16 := utils.Crc16CheckSum(hexBinPack)
	checkBinSum := make([]byte, 2)
	binary.LittleEndian.PutUint16(checkBinSum, checkSumUint16)

	return hexPack + hex.EncodeToString(checkBinSum), "02" + hex.EncodeToString(checkBinSum)
}
