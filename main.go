package main

import (
	"fmt"

	"github.com/Konstantin8105/file_manipulation_go/fileManipulation"
)

func main() {

	var (
		inputFolders = []Folder{
			"Z:\\Temp",
			"C:\\TEMP",
			"Z:\\GoogleDisk",
			"Z:\\SVNSERVER",
			"X:\\2 Project Execution\\Steel Structure Calculations",
		}
		outputFolder Folder = "E:\\Temp"
	)

	// fmt.Println("Input folders:")
	// for _, inputFolder := range inputFolders {
	// 	fmt.Println(inputFolder)
	// }
	//
	// fmt.Println("Output folder:")
	// fmt.Println(outputFolder)

	for _, inputFolder := range inputFolders {
		err := fileManipulation.Cleaning(inputFolder, outputFolder)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}
