package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	myProgramName := os.Args
	fmt.Printf("Started Creating App - %s - by clear arch", myProgramName)

	creating.createFolder(myProgramName, path)
}
