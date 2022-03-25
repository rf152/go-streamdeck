package devices

import (
	"image"

	streamdeck "github.com/rf152/go-streamdeck"
)

var (
	oName                     string
	oButtonWidth              uint
	oButtonHeight             uint
	oImageReportPayloadLength uint
)

// GetImageHeaderO returns the USB comms header for a button image for the XL
func GetImageHeaderO(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	thisLength := uint(0)
	if oImageReportPayloadLength < bytesRemaining {
		thisLength = oImageReportPayloadLength
	} else {
		thisLength = bytesRemaining
	}
	header := []byte{'\x02', '\x07', byte(btnIndex)}
	if thisLength == bytesRemaining {
		header = append(header, '\x01')
	} else {
		header = append(header, '\x00')
	}

	header = append(header, byte(thisLength&0xff))
	header = append(header, byte(thisLength>>8))

	header = append(header, byte(pageNumber&0xff))
	header = append(header, byte(pageNumber>>8))

	return header
}

func init() {
	oName = "Streamdeck (original)"
	oButtonWidth = 72
	oButtonHeight = 72
	oImageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		oName, // Name
		image.Point{X: int(oButtonWidth), Y: int(oButtonHeight)}, // Width/height of a button
		0x80,                        // USB productID
		resetPacket32(),             // Reset packet
		15,                          // Number of buttons
		3,                           // Number of rows
		5,                           // Number of columns
		brightnessPacket32(),        // Set brightness packet preamble
		4,                           // Button read offset
		"JPEG",                      // Image format
		oImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderO,           // Function to get the comms image header
	)
}
