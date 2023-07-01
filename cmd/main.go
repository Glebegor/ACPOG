package main

import (
	"fmt"
	"os"

	"github.com/Glebegor/ACPOG/pkg/creating"
)

func main() {
	mainDirsArray := []string{"/", "/cmd", "/pkg", "/schema", "/configs", "/interfaces", "/pkg/repository", "/pkg/service", "/pkg/handler"}
	mainFilesArray := []string{"/.env", "/.gitignore", "/server.go", "/README.md", "/DockerFile", "/MakeFile", "/cmd/main.go", "/schema/000001_init.up.sql", "/schema/000001_init.down.sql", "/configs/config.yml", "/pkg/repository/repository.go", "/pkg/service/service.go", "/pkg/handler/handler.go", "/pkg/handler/cors.go"}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	myAppName := os.Args
	workWithElems := creating.NewWorkWithElements(myAppName[1], path, mainDirsArray, mainFilesArray, getDataToWrite(myAppName[1]))
	fmt.Printf("Started Creating App - %s - by clear arch\n", myAppName[1])

	workWithElems.CreateArch(*workWithElems, myAppName[1], path)
	workWithElems.CreateFiles(*workWithElems, myAppName[1], path)
	workWithElems.WriteInfoForFiles(*workWithElems, myAppName[1], path)
}

func getDataToWrite(appName string) []string {
	return []string{
		"Secret_key='Secret_key'",
		"venv\n.env\nvenv/\n/venv\n",
		"",
		fmt.Sprintf("# Project %s", appName),
		"",
		"run:\n	go run cmd/main.go",
		"package main\n\nfunc main(){\n	return\n}",
		"",
		"",
		"",
		"package repository",
		"package service",
		"package handler",
		"package handler",
	}
}
