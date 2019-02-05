package main

import (
	"fmt"
	"time"

	"github.com/Konstantin8105/file_manipulation_go/fileManipulation"
)

func main() {
	cleaning()
	backup()
	// pause //
	fmt.Println("===\nPAUSE")
	time.Sleep(time.Second * 8)
}

func cleaning() {
	var (
		inputFolders = []fileManipulation.Folder{
			"Z:\\Temp",
			"C:\\TEMP",
			"Z:\\GoogleDisk\\Fired Heaters\\Projects",
			"Z:\\SVNSERVER",
			"X:\\2 Project Execution\\Steel Structure Calculations",
		}
		outputFolder fileManipulation.Folder = "E:\\Temp"
	)

	for _, inputFolder := range inputFolders {
		fmt.Println("Cleaning")
		fmt.Println("input  folder = ", string(inputFolder))
		fmt.Println("output folder = ", string(outputFolder))
		err := fileManipulation.Cleaning(inputFolder, outputFolder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

func backup() {
	for _, inputFolder := range []fileManipulation.Folder{
		"Z:\\SVNSERVER",    // old location
		"Z:\\git-projects", // new location
	} {
		for _, outputFolder := range []fileManipulation.Folder{
			"X:\\2 Project Execution\\Steel Structure Calculations",
		} {
			fmt.Println("BackUp")
			fmt.Println("input  folder = ", string(inputFolder))
			fmt.Println("output folder = ", string(outputFolder))
			err := fileManipulation.BackUp(inputFolder, outputFolder)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
		}
	}
}
