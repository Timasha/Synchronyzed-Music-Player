package tokens

import (
	"math/rand"
	"strings"
	"time"
)

var bodyChars []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func byte8bitToASCII(char byte) byte {
	if char < 32 {
		char += 64
	}
	return char
}

func fromASCIIToByte(char byte) byte {
	if char > 63 {
		char -= 64
	}
	return byte(char)
}

func timeToASCII(resultChan chan string, t time.Time) {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()

	date := uint32(year<<20) | uint32(month<<16) | uint32(day<<11) | uint32(hour<<6) | uint32(minute)

	var (
		sBuilder strings.Builder
		b6Bit    byte
		i        int8 = 25
	)
	for n := 0; n < 6; n++ {
		b6Bit = byte(date>>i) & 0x3F
		sBuilder.WriteByte(byte8bitToASCII(b6Bit))

		i -= 6
		if i < 0 {
			i = 0
		}
	}

	resultChan <- sBuilder.String()
}

func fromASCIIToTime(resultChan chan time.Time, str string) {
	defer func() {
		if recover() != nil {
			return
		}
	}()

	var (
		timeBitArray uint32
		b8Bit        byte
		i            int8 = 25
	)
	for n := 0; n < 6; n++ {
		b8Bit = fromASCIIToByte(byte(str[n]))
		timeBitArray |= uint32(b8Bit) << i

		i -= 6
		if i < 0 {
			i = 0
		}
	}

	year := int((timeBitArray >> 20) & 0x0fff)
	month := int((timeBitArray >> 16) & 0x0f)
	day := int((timeBitArray >> 11) & 0x1f)
	hour := int((timeBitArray >> 6) & 0x1f)
	minute := int((timeBitArray) & 0x3f)

	resultChan <- time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Now().Location())

}
func generateRefreshBody(resultChan chan string, bodyLen int) {
	src := rand.NewSource(time.Now().Unix())
	newRand := rand.New(src)
	refreshBody := make([]byte, bodyLen)
	for i := 0; i < 8; i++ {
		refreshBody[i] = bodyChars[newRand.Intn(62)]
	}
	resultChan <- string(refreshBody)

}
