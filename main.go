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
			//	"Y:\\",
		}
		outputFolder fileManipulation.Folder = "E:\\Temp"
	)

	for _, inputFolder := range inputFolders {
		fmt.Println("\nCleaning")
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
	for _, f := range []struct {
		in, out fileManipulation.Folder
	}{
		{
			in:  "\\\\192.168.5.195\\web",
			out: "Z:\\git-projects\\Articles",
		},
		{
			in:  "Y:\\",
			out: "Z:\\git-projects",
		},
		{
			in:  "Z:\\git-projects",
			out: "X:\\2 Project Execution\\Steel Structure Calculations",
		},
		{
			in:  "E:\\Outlook data",
			out: "G:\\",
		},
	} {
		fmt.Println("\nBackup")
		fmt.Println(f.in, " --> ", f.out)
		err := fileManipulation.BackUp(f.in, f.out)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
