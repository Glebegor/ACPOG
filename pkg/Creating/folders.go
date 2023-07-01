package creating

import (
	"log"
	"os"
)

func createFolder(name string, path string) {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
