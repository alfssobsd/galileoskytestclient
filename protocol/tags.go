package protocol

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
)

const (
	HWversion = iota
	FWversion
	IMEI
	DeviceID
	ArchiveNumber
	DateTime
	LatLon
)

func TagEncoder(tag int, value string) string {

	switch tag {

	case HWversion:
		return encodeHwVersion(value)
	case IMEI:
		return encodeIMEI(value)
	case FWversion:
		return encodeFwVersion(value)
	case DeviceID:
		return encodeDeviceID(value)
	case ArchiveNumber:
		return encodeArchiveNumber(value)
	case DateTime:
		return encodeDateTime(value)
	case LatLon:
		return encodeLatLon(value)
	default:
		return "unknow type"
	}

}

func encodeHwVersion(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "01" + hex.EncodeToString(bs[:1])
}

func encodeIMEI(value string) string {
	if len(value) != 15 {
		log.Panicln("incorrect IMEI = ", value)
	}
	hx := hex.EncodeToString([]byte(value))

	return "03" + hx
}

func encodeFwVersion(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "02" + hex.EncodeToString(bs[:1])
}

func encodeDeviceID(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "04" + hex.EncodeToString(bs[:2])
}

func encodeArchiveNumber(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "10" + hex.EncodeToString(bs[:2])
}

func encodeDateTime(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return "20" + hex.EncodeToString(bs[:4])
}

func encodeLatLon(value string) string {
	arr := strings.Split(value, ";")

	tmp1, _ := strconv.ParseFloat(arr[2], 32)
	lat := int32(tmp1 * 1000000)
	bslat := make([]byte, 4)
	binary.LittleEndian.PutUint32(bslat, uint32(lat))

	tmp2, _ := strconv.ParseFloat(arr[3], 32)
	lon := int32(tmp2 * 1000000)
	bslon := make([]byte, 4)
	binary.LittleEndian.PutUint32(bslon, uint32(lon))

	return "30" + arr[0] + arr[1] + hex.EncodeToString(bslat[:]) + hex.EncodeToString(bslon[:])
}
