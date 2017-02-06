package fileManipulation

import (
	"fmt"
	"os"
	"runtime"
)

// Folder - folder like string
type Folder string

type fileParam struct {
	fileInfo os.FileInfo
	basePath Folder
	path     Folder
}

// Cleaning function for cleaning folder from temp file of STAAD
func Cleaning(inputFolder, outputFolder Folder) error {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	toFilter := search(inputFolder, []folder{""})
	toMoving := filter(toFilter)
	success := movingFiles(outputFolder, toMoving)
	if <-success {
		fmt.Println("Done...", inputFolder)
	}

	return nil
}
