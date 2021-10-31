package file

import (
	"document-measuring-api/pkg/error"

	"log"
	"os"
)

func Read(path string) []byte {
	
	_, err := os.Open(path)
    error.Check("file", err)

	log.Panicln("READING FILE")

	return nil
}
