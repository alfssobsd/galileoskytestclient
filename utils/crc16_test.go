package utils

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrc16CheckSumMainPack(t *testing.T) {
	b, _ := hex.DecodeString("0133000111021403393930303030383632343731383534046e0020000000003007382c920300000000332f010807341400350040310a")

	checkSumUint16 := Crc16CheckSum(b)
	checkBinSum := make([]byte, 2)
	binary.LittleEndian.PutUint16(checkBinSum, checkSumUint16)

	assert.Equal(t, "3cd7", hex.EncodeToString(checkBinSum))

}

func TestCrc16CheckSumHeadPack(t *testing.T) {
	b, _ := hex.DecodeString("0117000111021403393930303030383632343731383534046e00")

	checkSumUint16 := Crc16CheckSum(b)
	checkBinSum := make([]byte, 2)
	binary.LittleEndian.PutUint16(checkBinSum, checkSumUint16)

	assert.Equal(t, "79fd", hex.EncodeToString(checkBinSum))
}
