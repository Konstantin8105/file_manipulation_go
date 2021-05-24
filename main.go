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
			"X:\\2 Project Execution\\Steel Structure Calculations",
			"Z:\\git-projects",
			"M:\\",
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
	{
		var in,out fileManipulation.Folder = "M:\\", "Z:\\git-projects"
		fmt.Println(in," --> ", out)
		err := fileManipulation.BackUp(in, out)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	{
		var in,out fileManipulation.Folder = "Z:\\git-projects", "X:\\2 Project Execution\\Steel Structure Calculations"
		fmt.Println(in," --> ", out)
		err := fileManipulation.BackUp(in, out)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
