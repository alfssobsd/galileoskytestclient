package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeTagHWVersion(t *testing.T) {
	result := TagEncoder(HWversion, "17")
	assert.Equal(t, "0111", result)

	result1 := TagEncoder(HWversion, "18")
	assert.Equal(t, "0112", result1)

	result2 := TagEncoder(HWversion, "26")
	assert.Equal(t, "011a", result2)

	result3 := TagEncoder(HWversion, "7")
	assert.Equal(t, "0107", result3)

	result4 := TagEncoder(HWversion, "256")
	assert.Equal(t, "0100", result4)
}

func TestEncodeTagIMEI(t *testing.T) {
	result := TagEncoder(IMEI, "868204005647838")
	assert.Equal(t, "03383638323034303035363437383338", result)

	assert.Panics(t, func() { TagEncoder(IMEI, "4005647838") }, "The code did not panic")
}

func TestEncodeTagFWVersion(t *testing.T) {
	result := TagEncoder(FWversion, "223")
	assert.Equal(t, "02df", result)
}

func TestEncodeTagDeviceID(t *testing.T) {
	result := TagEncoder(DeviceID, "50")
	assert.Equal(t, "043200", result)

	result2 := TagEncoder(DeviceID, "256")
	assert.Equal(t, "040001", result2)
}

func TestEncodeTagArchiveNumber(t *testing.T) {
	result := TagEncoder(ArchiveNumber, "10")
	assert.Equal(t, "100a00", result)

}

func TestDateTime(t *testing.T) {
	result := TagEncoder(DateTime, "1558916283")
	assert.Equal(t, "20bb2ceb5c", result)

	result2 := TagEncoder(DateTime, "1552831531")
	assert.Equal(t, "202b548e5c", result2)
}

func TestLatLon(t *testing.T) {
	result := TagEncoder(LatLon, "0;7;53.612224;86.890426")
	assert.Equal(t, "3007c00e3203bad72d05", result)
}

func TestEncodeSpeedNDirection(t *testing.T) {
	result := TagEncoder(SpeedNDirection, "9.2;212")
	assert.Equal(t, "335c004808", result)

	result2 := TagEncoder(SpeedNDirection, "45.2;90")
	assert.Equal(t, "33c4018403", result2)
}

func TestEncodeHeightMeters(t *testing.T) {
	result := TagEncoder(HeightMeters, "100")
	assert.Equal(t, "346400", result)

	result2 := TagEncoder(HeightMeters, "1000")
	assert.Equal(t, "34e803", result2)
}

func TestEncodeHdop(t *testing.T) {
	result := TagEncoder(HDOP, "0")
	assert.Equal(t, "3500", result)

	result2 := TagEncoder(HDOP, "10")
	assert.Equal(t, "350a", result2)
}

func TestEncodeHWstatus(t *testing.T) {
	result := TagEncoder(HWStatus, "1;0;0;1;0;0;0;1;1;0;0;1;1;1;1;1")
	assert.Equal(t, "4089f9", result)

	result1 := TagEncoder(HWStatus, "1;0;0;1;0;0;0;1;1;0;0;1;1;1;0;0")
	assert.Equal(t, "408939", result1)
}

func TestEncodeSupplyVoltage(t *testing.T) {
	result := TagEncoder(SupplyVoltage, "24")
	assert.Equal(t, "411800", result)

	result2 := TagEncoder(SupplyVoltage, "300")
	assert.Equal(t, "412c01", result2)
}

func TestEncodeBatteryVoltage(t *testing.T) {
	result := TagEncoder(BetteryVoltage, "24")
	assert.Equal(t, "421800", result)

	result2 := TagEncoder(BetteryVoltage, "300")
	assert.Equal(t, "422c01", result2)
}

func TestEncodeTemperatureOfTerminal(t *testing.T) {
	result := TagEncoder(TemperatureOfTerminal, "-5")
	assert.Equal(t, "43fb", result)

	result2 := TagEncoder(TemperatureOfTerminal, "-25")
	assert.Equal(t, "43e7", result2)

	result3 := TagEncoder(TemperatureOfTerminal, "0")
	assert.Equal(t, "4300", result3)

	result4 := TagEncoder(TemperatureOfTerminal, "10")
	assert.Equal(t, "430a", result4)
}
