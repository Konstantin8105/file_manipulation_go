package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func movingFiles(outputFolder folder, fileFromChannel chan fileParam) {
	var counter uint64
	for {
		file := <-fileFromChannel
		counter++
		fmt.Println("counter", counter)
		fmt.Println("TAKE IN MOVING", file.Sting())

		// create general folder
		generalFolderName := strings.Replace(string(file.basePath), "\\", "_", -1)
		generalFolderName = strings.Replace(generalFolderName, ":", "_", -1)
		generalFolderName = strings.Replace(generalFolderName, " ", "_", -1)
		generalFolderName = fmt.Sprintf("%s\\%s\\%s", outputFolder, generalFolderName, string(file.path)[1:])

		// create internal folder
		if _, err := os.Stat(generalFolderName); os.IsNotExist(err) {
			os.MkdirAll(generalFolderName, os.ModePerm)
			fmt.Println("counter", counter)
			fmt.Println("create folder -->", generalFolderName)
		} else {
			// fmt.Println("counter", counter)
			// fmt.Println("folder is exist -->", generalFolderName)
		}

		// copy of file
		inputFileName := fmt.Sprintf("%s\\%s\\%s", string(file.basePath), string(file.path)[1:], file.fileInfo.Name())
		inputFile, err := os.Open(inputFileName)
		if err != nil {
			fmt.Println("counter", counter)
			fmt.Println("Can not open file", inputFileName)
			return
		}

		outputFileName := fmt.Sprintf("%s\\%s", generalFolderName, file.fileInfo.Name())
		outputFile, err := os.Create(outputFileName)
		defer outputFile.Close()
		if err != nil {
			fmt.Println("counter", counter)
			fmt.Println("Can not create a new file", outputFileName)
			return
		}

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			fmt.Println("counter", counter)
			fmt.Println("Can not copy file", outputFileName)
			return
		}

		err = outputFile.Sync()
		if err != nil {
			fmt.Println("counter", counter)
			fmt.Println("Can not flush the file ", outputFileName)
			return
		}
		inputFile.Close()

		// check md5

		// remove resource

		err = os.Remove(inputFileName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
