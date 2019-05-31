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
	SpeedNDirection
	HeightMeters
	HDOP
	HWStatus
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
	case SpeedNDirection:
		return encodeSpeedNDirection(value)
	case HeightMeters:
		return encodeHeightMeters(value)
	case HDOP:
		return encodeHdop(value)
	case HWStatus:
		return encodeHWstatus(value)
	default:
		return "unknow type"
	}

}

// Encode tag 0x01
// value = unit, hardware version of device
// example:
//    17 - Galileosky GPS/GLONASS 5.0.11
// return encode tag as hex string
func encodeHwVersion(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "01" + hex.EncodeToString(bs[:1])
}

// Encode tag 0x03
// value = string, is IMEI
//
// return encode tag as hex string
func encodeIMEI(value string) string {
	if len(value) != 15 {
		log.Panicln("incorrect IMEI = ", value)
	}
	hx := hex.EncodeToString([]byte(value))

	return "03" + hx
}

// Encode tag 0x02
// value = uint, version of firmware
//
// return encode tag as hex string
func encodeFwVersion(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "02" + hex.EncodeToString(bs[:1])
}

// Encode tag 0x04
// value = uint, device model identifier
// example: 130 - terminal version 7, 110  it is BaseBlock terminal
//
// return encode tag as hex string
func encodeDeviceID(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "04" + hex.EncodeToString(bs[:2])
}

// Encode tag 0x10
// value = uint, number of record in archive
// return encode tag as hex string
func encodeArchiveNumber(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return "10" + hex.EncodeToString(bs[:2])
}

// Encode tag 0x20
// value = timestamp as string
// return encode tag as hex string
func encodeDateTime(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return "20" + hex.EncodeToString(bs[:4])
}

// Encode tag 0x30
// value = string, like: "0;7;53.612224;86.890426"
//
// 0 - how are the coordinates obtained 0 = GPS/GLONASS, 2 = Mobile network
// 7 - number of satellites
// 53.612224 - Latitude
// 86.890426 - Longitude
//
// return encode tag as hex string
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

// Encode tag 0x33
// value = string, speed and direction
//
// example: "9.2;212"
// 9.2 - speed in km/h
// 212 - direction in degrees
//
// return encode tag as hex string
func encodeSpeedNDirection(value string) string {
	arr := strings.Split(value, ";")

	speed, _ := strconv.ParseFloat(arr[0], 16)
	speed = speed * 10
	direction, _ := strconv.ParseFloat(arr[1], 16)
	direction = direction * 10

	bss := make([]byte, 2)
	binary.LittleEndian.PutUint16(bss, uint16(speed))

	bsd := make([]byte, 2)
	binary.LittleEndian.PutUint16(bsd, uint16(direction))

	return "33" + hex.EncodeToString(bss[:]) + hex.EncodeToString(bsd[:])
}

// Encode tag 0x34
// value = unit, height in meters
//
// return encode tag as hex string
func encodeHeightMeters(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))

	return "34" + hex.EncodeToString(bs[:])
}

// Encode tag 0x35
// value = unit, hdop or error in meters
// Only 1 bite
//
// return ecnode tag as hex string
func encodeHdop(value string) string {
	i, _ := strconv.Atoi(value)
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))

	return "35" + hex.EncodeToString(bs[:1])
}

// Encode tag 0x40
// value = string (bits), status of device
//
// example: "1;0;0;0;1;1;0;0;0;1;0;1;0;0;0;0"
// bits values:
// 0: 0 - vibration level eq parking, 1 - vibration level eq movement
// 1: 0 - tilt angle ok, 1 - tilt angle exceeded
// 2: 0 - disable iButton key, 1 - enable iButton key
// 3: 0 - SIM Card enable, 1 - SIM Card disable
// 4: 0 - terminal out of GEO zone, 1 - terminal into GEO zone
// 5: 0 - voltage is problem, 1 - voltage is ok
// 6: 0 - enable GPS antenna, 1 - disable GPS antenna
// 7: 0 - voltage on internal bus is ok, 1 - voltage on internal bus is problem
// 8: 0 - external voltage is ok, 1 - external voltage  is problem
// 9: 0 - car is off, 1 - car is on
// 10: 0 - vibration level eq normal move, 1 - vibration level eq hit
// 11: GPS: 0 - geo coords from internal module, 1 - geo coords from external module
//     GLONASS: 1 - geo coords from internal module, 0 - geo coords from external module
// 12 - 13: quality of signal 0 to 3, less is good
// 14: 0 - alarm on, 1 - alarm off
// 15: 0 - all fine, 1 - alarm triggered
func encodeHWstatus(value string) string {
	arr := strings.Split(value, ";")
	tmp := uint16(0)

	for i := range arr {
		if arr[i] == "1" {
			tmp |= 1 << uint(i)
		}
	}

	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(tmp))

	return "40" + hex.EncodeToString(bs[:])
}
