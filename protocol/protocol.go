package protocol

import (
	"github.com/ottojo/blnk2/client"
)

type BlnkPacket []byte

func SetPixelsMessage(element *client.LedListElement, length int) BlnkPacket {
	var message = make([]byte, 3*length)
	for i := 0; i < length; i++ {
		r, b, g := element.Data.Color.LinearRgb()
		message[3*i] = byte(r * 255)
		message[3*i+1] = byte(g * 255)
		message[3*i+2] = byte(b * 255)
		element = element.Next

	}
	message = append([]byte{1}, message...)
	return newPacket(message)
}

func newPacket(payload []byte) BlnkPacket {
	return append([]byte{0xaf, byte(len(payload) >> 8), byte(len(payload))}, payload...)
}

func SetIntervalMessage(element *client.LedListElement, length, startIndex int) BlnkPacket {
	var message = make([]byte, 3*length+2)
	message[0] = byte(startIndex >> 8)
	message[1] = byte(startIndex)
	for i := 0; i < length; i++ {
		r, b, g := element.Data.Color.LinearRgb()
		message[3*i+2] = byte(r * 255)
		message[3*i+3] = byte(g * 255)
		message[3*i+4] = byte(b * 255)
		element = element.Next

	}
	message = append([]byte{2}, message...)
	return newPacket(message)
}
