package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Konstantin8105/file_manipulation_go/fileManipulation"
)

func main() {
	cleaning()
	backup()
	// pause //
	fmt.Fprintf(os.Stdout, "===\nPAUSE\n")
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
		fmt.Fprintf(os.Stdout, "\nCleaning\n")
		fmt.Fprintf(os.Stdout, "input  folder = %s\n", inputFolder)
		fmt.Fprintf(os.Stdout, "output folder = %s\n", outputFolder)
		err := fileManipulation.Cleaning(inputFolder, outputFolder)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error: %v\n", err)
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
			in:  "\\\\192.168.5.195\\web",
			out: "O:\\Archive\\Articles",
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
			in:  "Z:\\git-projects",
			out: "O:\\Work\\git-projects",
		},
		{
			in:  "E:\\Outlook data",
			out: "H:\\",
		},
		{
			in:  "X:\\2 Project Execution\\Steel Structure Calculations\\git-projects",
			out: "O:\\Backup\\git-projects",
		},
	} {
		fmt.Fprintf(os.Stdout, "\nBackup\n")
		fmt.Fprintf(os.Stdout, "%s --> %s\n", f.in, f.out)
		err := fileManipulation.BackUp(f.in, f.out)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error: %v\n", err)
		}
	}
}
