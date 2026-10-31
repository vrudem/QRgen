package displayers

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/floatpane/termimage"
	"github.com/skip2/go-qrcode"
)

func WriteQrToFile(content string, outputFile string) {
	err := qrcode.WriteFile(content, qrcode.Medium, 256, outputFile)
	if err != nil {
		log.Fatalf("failed to generate QR code: %v", err)
	}
}

func WriteQrToTerminal(content string) {
	png, err := qrcode.Encode(content, qrcode.Highest, 256)
	if err != nil {
		log.Fatalf("failed to generate QR code: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(png)
	dataURI := "data:image/png;base64," + b64

	err = termimage.Display(os.Stdout, dataURI, termimage.Options{
		Protocol: termimage.Auto,
	})
	if err != nil {
		log.Fatalf("failed displaying QR to terminal: %v", err)
	}
}
