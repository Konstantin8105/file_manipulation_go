package fileManipulation

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
)

// BackUp - copy files from inputFolder to outputFolder
func BackUp(inputFolder, outputFolder Folder) (err error) {

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
	errChannel := make(chan error)
	files := getInputFilesFlow(inputFolder, errChannel)
	success := copyFiles(files, inputFolder, outputFolder, errChannel)

	end := make(chan struct{})
	go func() {
		for e := range errChannel {
			err = errors.Join(err, e)
		}
		end <- struct{}{}
	}()

	<-success
	close(errChannel)
	<-end
	return err
}

func getInputFilesFlow(inputFolder Folder, errChannel chan<- error) <-chan fileParam {
	inputFiles := make(chan fileParam)
	go func() {
		defer close(inputFiles)
		folders := getInternalDirectory(inputFolder, errChannel)
		for folder := range folders {
			isStaadFolder, err := folder.withStaadFiles()
			if err != nil {
				errChannel <- err
				continue
			}
			files, err := os.ReadDir(string(folder))
			if err != nil {
				errChannel <- err
				continue
			}
			f := folder
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if isStaadFolder && isStaadTempFile(file.Name()) {
					continue
				}
				if ignoreFile(file.Name()) {
					continue
				}
				fi, err := file.Info()
				if err != nil {
					errChannel <- err
					continue
				}
				inputFiles <- fileParam{fileInfo: fi, path: f}
			}
		}
	}()
	return inputFiles
}

// files come with format:
// file.info .... Name = "kernel.dll" ...
// file.path = "C://Windows//"
// inputFolder = "C:"
// outputFolder = "X:"
// Transformations of file name:
// inFileName  = C://Windows//kernel.dll
// outFileName = X://Windows//kernel.dll
func copyFiles(
	files <-chan fileParam,
	inputFolder Folder,
	outputFolder Folder,
	errChannel chan<- error) chan struct{} {

	success := make(chan struct{})
	go func() {
		defer func() {
			success <- struct{}{}
		}()
		for file := range files {
			inFileName := fmt.Sprintf("%s\\%s", file.path, file.fileInfo.Name())

			var outFileName string
			var outputFullFolder Folder
			if len(string(inputFolder)) == len(string(file.path)) {
				outputFullFolder = outputFolder
			} else {
				outputFullFolder = Folder(fmt.Sprintf(
					"%s\\%s", string(outputFolder), string(file.path)[(len(string(inputFolder))+1):]))
			}
			outFileName = fmt.Sprintf("%s\\%s", outputFullFolder, file.fileInfo.Name())

			// create a output folder if not exist
			err := createDirectory(outputFullFolder)
			if err != nil {
				errChannel <- err
				continue
			}

			// optimization of copy time:
			// - if files with same time, then no copy
			// - if files with same size, then no copy
			var copy bool
			copy, err = isNeedCopy(inFileName, outFileName)
			if err != nil {
				errChannel <- err
				continue
			}

			if copy {
				err := CopyWithCheckingMd5(inFileName, outFileName)
				if err != nil {
					errChannel <- err
					continue
				}
			}
		}
	}()
	return success
}

func isNeedCopy(in, out string) (_ bool, err error) {
	var inFileInfo, outFileInfo os.FileInfo
	// check out file is exist, if not exist, then need copy
	if outFileInfo, err = os.Stat(out); os.IsNotExist(err) {
		// out file does not exist
		return true, nil
	}
	// check in file is exist, if not exist, then error
	if inFileInfo, err = os.Stat(in); os.IsNotExist(err) {
		// in file does not exist
		return true, err
	}

	// if diff time or size, then need copy
	if inFileInfo.Size() != outFileInfo.Size() {
		return true, nil
	}
	if inFileInfo.ModTime() != outFileInfo.ModTime() {
		return true, nil
	}

	// in 0.1% cases just copy //
	var (
		amount      = 1000
		luckyNumber = 222
	)
	if rand.Intn(amount) == luckyNumber {
		return true, nil
	}

	return false, nil
}

func ignoreFile(filename string) bool {
	for _, ext := range []string{
		".dwl",
		".dwl2",
		"Thumbs.db",
	} {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	for _, pre := range []string{
		"~$",
	} {
		if strings.HasPrefix(filename, pre) {
			return true
		}
	}
	return false
}
