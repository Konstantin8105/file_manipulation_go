package fileManipulation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
)

//BackUp - copy files from inputFolder to outputFolder
func BackUp(inputFolder, outputFolder Folder) error {

	// Check input data
	if string(inputFolder)[len(inputFolder)-1] == '\\' {
		inputFolder = Folder(string(inputFolder)[:(len(inputFolder) - 1)])
	}

	if string(outputFolder)[len(outputFolder)-1] == '\\' {
		outputFolder = Folder(string(outputFolder)[:(len(outputFolder) - 1)])
	}

	if string(inputFolder) == string(outputFolder) {
		return fmt.Errorf("Input and output folder cannot be same")
	}

	// use all allowable proccesors
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	// working
	files, errChannel := getInputFilesFlow(inputFolder)
	defer close(*errChannel)
	success := copyFiles(files, inputFolder, outputFolder, errChannel)
	defer close(success)

	select {
	case <-success:
		return nil
	case err := <-*errChannel:
		return err
	}
}

func getInputFilesFlow(inputFolder Folder) (<-chan fileParam, *(chan error)) {
	inputFiles := make(chan fileParam)
	errFunc := make(chan error)
	go func() {
		defer close(inputFiles)
		folders := getInternalDirectory(inputFolder, &errFunc)
		for folder := range folders {
			isStaadFolder, err := folder.withStaadFiles()
			if err != nil {
				errFunc <- err
				return
			}
			files, err := ioutil.ReadDir(string(folder))
			if err != nil {
				errFunc <- err
				return
			}
			f := folder
			for _, file := range files {
				if !file.IsDir() {
					if (isStaadFolder && !isStaadTempFile(file.Name())) || !isStaadFolder {

						inputFiles <- fileParam{fileInfo: file, path: f}
					}
				}
			}
		}

	}()
	return inputFiles, &errFunc
}

// files come with format:
// file.info .... Name = "kernel.dll" ...
// file.path = "C://Windows//"
// inputFolder = "C:"
// outputFolder = "X:"
// Transformations of file name:
// inFileName  = C://Windows//kernel.dll
// outFileName = X://Windows//kernel.dll
func copyFiles(files <-chan fileParam, inputFolder Folder, outputFolder Folder, errChannel *(chan error)) chan bool {
	success := make(chan bool)
	go func() {
		for file := range files {
			fmt.Println("Income file  = ", file)
			fmt.Println("inputFolder  = ", inputFolder)
			fmt.Println("outputFolder = ", outputFolder)

			inFileName := fmt.Sprintf("%s\\%s", string(inputFolder), file.fileInfo.Name())
			outFileName := fmt.Sprintf("%s\\%s\\%s", string(outputFolder), string(file.path)[(len(string(inputFolder))+1):], file.fileInfo.Name())

			fmt.Println("inFile ======> ", inFileName)
			fmt.Println("outFile =====> ", outFileName)

			// optimization of copy time:
			// - if files with same time, then no copy
			// - if files with same size, then no copy
			var copy bool
			copy, err := isNeedCopy(inFileName, outFileName)
			if err != nil {
				*errChannel <- err
				return
			}

			fmt.Println("isNeedCopy == ", copy)
			fmt.Println("\n\n\n\n")
			/*
				if copy {
					err := CopyWithCheckingMd5(inFileName, outFileName)
					if err != nil {
						*errChannel <- err
						return
					}
				}
			*/
		}
		success <- true
	}()
	return success
}

func isNeedCopy(in, out string) (b bool, err error) {
	b = true
	var inFileInfo, outFileInfo os.FileInfo
	// check out file is exist, if not exist, then need copy
	if outFileInfo, err = os.Stat(out); os.IsNotExist(err) {
		// out does not exist
		return b, nil
	}
	// check in file is exist, if not exist, then error
	if inFileInfo, err = os.Stat(out); os.IsNotExist(err) {
		// out does not exist
		return b, err
	}

	// if diff time or size, then need copy
	if inFileInfo.Size() != outFileInfo.Size() {
		return b, nil
	}
	if inFileInfo.ModTime() != outFileInfo.ModTime() {
		return b, nil
	}

	// in 0.1% cases just copy //
	var (
		amount      = 1000
		luckyNumber = 222
	)
	if rand.Intn(amount) == luckyNumber {
		return b, nil
	}

	return false, nil
}
