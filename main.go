package main

import (
	"fmt"
	"os"
)

type folder string
type extension string
type fileParam struct {
	fileInfo os.FileInfo
	basePath folder
	path     folder
}

func (file fileParam) Sting() string {
	if !file.fileInfo.IsDir() {
		return fmt.Sprintln("File:", string(file.basePath), string(file.path), file.fileInfo.Name())
	}
	return fmt.Sprintln("Directory:", string(file.basePath), string(file.path), file.fileInfo.Name())
}

var exts = []extension{
	"bmd", "CFR", "cod",
	"cut", "day", "dbi",
	"dbs", "dgn", "dsp",
	"slg", "ecf", "ejt",
	"EQL", "est", "EU2",
	"emf", "NLD", "num",
	"rea", "REI_SPRO_Auxilary_Data",
	"sbk", "scn", "slv",
	"ANL", "u01", "u02",
	"u03", "u04", "u05",
	"u06", "u07", "u08",
	"UID", "REI_Saved_Picture",
	"err", "bsh", "cfc",
	"ben", "MMTX", "dsn",
	"jst", "REI_STEELDESIGNER_DATA",
	"str",
}

var suffixs = []extension{
	"_MASS.TXT", "ed.Backup",
}

// General*.msh
// Preliminary*.msh
// General*.log
// Preliminary*.log

func main() {

	var (
		inputFolders = []folder{
			"Z:\\Temp",
			"C:\\TEMP",
			"Z:\\GoogleDisk",
			"Z:\\SVNSERVER",
			"X:\\2 Project Execution\\Steel Structure Calculations",
		}
		outputFolder folder = "E:\\Temp"
	)

	fmt.Println("Input folders:")
	for _, inputFolder := range inputFolders {
		fmt.Println(inputFolder)
	}

	fmt.Println("Search extentions:")
	for _, ext := range exts {
		fmt.Println(ext)
	}

	fmt.Println("Output folder:")
	fmt.Println(outputFolder)

	for _, inputFolder := range inputFolders {
		toFilter := search(inputFolder, []folder{""})
		toMoving := filter(toFilter)
		success := movingFiles(outputFolder, toMoving)
		if <-success {
			fmt.Println("Done...", inputFolder)
		}
	}

	// toFilter := make(chan fileParam)
	// toMovingFile := make(chan fileParam)
	//
	// go movingFiles(outputFolder, toMovingFile)
	// go filter(exts, toFilter, toMovingFile)
	//
	// for _, inputFolder := range inputFolders {
	// 	search(inputFolder, "", toFilter)
	// }
	//
	// time.Sleep(1000)
}
