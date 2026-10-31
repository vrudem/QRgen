package main

import (
	"io"
	"log"
	"os"

	"CLI_QRgen/pkg/argparser"
	"CLI_QRgen/pkg/displayers"

	"github.com/floatpane/termimage"
)

func main() {
	// для termimage, инициализация
	termimage.MaybeRunWorker()
	sourceFile, finalFile := argparser.ParseArgs()

	file, err := os.Open(sourceFile)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("failed to seek: %v", err)
	}

	buffer := make([]byte, argparser.MAX_SIZE)
	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatalf("failed to read bytes: %v", err)
	}

	content := string(buffer[:bytesRead]) //не включая крайний элемент, от [0 до bytesRead-1] или [0 до bytesRead)

	if finalFile != "" {
		displayers.WriteQrToFile(content, finalFile)
	} else {
		displayers.WriteQrToTerminal(content)
	}
}
