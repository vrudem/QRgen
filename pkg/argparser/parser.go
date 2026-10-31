package argparser

import (
	"log"
	"os"
)

const MAX_SIZE = 2953

var ProgName string = "qrgen"

func PrintHelp() {
	log.Fatalf(`# incorrect usage.
		Please
		%s <SOURCE_FILE> <FINAL_FILE>
		or
		%s <SOURCE_FILE>
		# note that the max size u can encode with QR is %d bytes`,
		ProgName, ProgName, MAX_SIZE,
	)
}

func ParseArgs() (sourceFile string, finalFile string) {
	if len(os.Args) < 2 {
		PrintHelp()
		os.Exit(1)
	} else if len(os.Args) == 2 {
		return os.Args[1], ""
	}
	return os.Args[1], os.Args[2]
}
