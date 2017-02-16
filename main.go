package main

import (
	"fmt"

	"github.com/Konstantin8105/file_manipulation_go/fileManipulation"
)

func main() {
	cleaning()
	//backup()
}

func cleaning() {
	var (
		inputFolders = []fileManipulation.Folder{
			"Z:\\Temp",
			"C:\\TEMP",
			"Z:\\GoogleDisk",
			"Z:\\SVNSERVER",
			"X:\\2 Project Execution\\Steel Structure Calculations",
		}
		outputFolder fileManipulation.Folder = "E:\\Temp"
	)

	for _, inputFolder := range inputFolders {
		err := fileManipulation.Cleaning(inputFolder, outputFolder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

func backup() {
	var (
		inputFolder   fileManipulation.Folder = "Z:\\SVNSERVER"
		outputFolders                         = []fileManipulation.Folder{
			"X:\\2 Project Execution\\Steel Structure Calculations",
			"Z:\\GoogleDisk",
		}
	)
	for _, outputFolder := range outputFolders {
		err := fileManipulation.BackUp(inputFolder, outputFolder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}
