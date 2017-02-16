package fileManipulation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	//"strings"
	"time"
)

// Folder - folder like string
type Folder string

type fileParam struct {
	fileInfo os.FileInfo
	path     Folder
}

// Cleaning function for cleaning folder from temp file of STAAD
func Cleaning(inputFolder, outputFolder Folder) error {

	// Check input data
	//if string(inputFolder)[len(inputFolder)-1] == '\\' {
	//	inputFolder = Folder(string(inputFolder)[:(len(inputFolder) - 1)])
	//}

	//if string(outputFolder)[len(outputFolder)-1] == '\\' {
	//	outputFolder = Folder(string(outputFolder)[:(len(outputFolder) - 1)])
	//}

	//if string(inputFolder) == string(outputFolder) {
	//	return fmt.Errorf("Input and output folder cannot be same")
	//}
	//if strings.HasSuffix(string(outputFolder), string(inputFolder)) {
	//	return fmt.Errorf("Output folder cannot be inside input folder")
	//}

	fmt.Println("inputFolder = ",inputFolder)
	fmt.Println("outputFolder = ",outputFolder)
	
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())
fmt.Println("#1")
	staadFolders, errChannel := getStaadFolders(inputFolder)
	defer close(*errChannel)
fmt.Println("#2")
	tempFiles := filterTempStaadFiles(staadFolders, errChannel)
fmt.Println("#3")
	success := moveTempStaadFiles(tempFiles, inputFolder, outputFolder, errChannel)
fmt.Println("#4")
	defer close(success)

fmt.Println("#5")
	select {
	case <-success:
		return nil
	case err := <-*errChannel:
		return err
	}
}

func moveTempStaadFiles(tempFiles <-chan fileParam, inputFolder, outputFolder Folder, errChannel *chan error) chan bool {
	success := make(chan bool)
	go func() {
		defer close(success)
		for tempFile := range tempFiles {
			inputFileName, outputFileName, folder, err := convert(tempFile, inputFolder, outputFolder)
			if err != nil {
				*errChannel <- err
				return
			}

			err = createDirectory(folder)
			if err != nil {
				*errChannel <- err
				return
			}

			err = CopyWithCheckingMd5(inputFileName, outputFileName)
			if err != nil {
				*errChannel <- err
				return
			}

			err = removeFile(inputFileName)
			if err != nil {
				*errChannel <- err
				return
			}
		}
		success <- true
	}()
	return success
}

func filterTempStaadFiles(staadFolders <-chan Folder, errChannel *chan error) <-chan fileParam {
	tempFiles := make(chan fileParam)
	go func() {
		defer close(tempFiles)
		for folder := range staadFolders {

			fmt.Printf("F")

			files, err := ioutil.ReadDir(string(folder))
			if err != nil {
				*errChannel <- err
				return
			}

			for _, file := range files {
				// filter by last 24 hours files
				if time.Since(file.ModTime()).Hours() < 24.0 {
					continue
				}
				// filter by temp staad files
				if !file.IsDir() {
					if !isStaadFile(file.Name()) {
						if isStaadTempFile(file.Name()) {
							folder := folder
							tempFiles <- fileParam{fileInfo: file, path: folder}
						} else {
							//fmt.Println("Maybe add for delete list : ", file.Name())
						}
					}
				}
			}
		}
	}()
	return tempFiles
}

func getStaadFolders(inputFolder Folder) (<-chan Folder, *(chan error)) {
	staadFolders := make(chan Folder)
	errFunc := make(chan error)
	go func() {
		defer close(staadFolders)
			fmt.Printf("S")
		folders := getInternalDirectory(inputFolder, &errFunc)
		for folder := range folders {
			ok, err := folder.withStaadFiles()
			if err != nil {
				errFunc <- err
				return
			}
			if ok {
				staadFolders <- folder
			}
		}
	}()
	return staadFolders, &errFunc
}

func (folder Folder) withStaadFiles() (bool, error) {
	if len(string(folder)) == 0 {
		return false, errors.New("Null size of folder")
	}
	files, err := ioutil.ReadDir(string(folder))
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if isStaadFile(file.Name()) {
				return true, nil
			}
		}
	}
	return false, nil
}

var t int = 1

func getInternalDirectory(folder Folder, errChannel *chan error) chan Folder {
	channel := make(chan Folder)
	go func() {
		defer close(channel)
		t=t+1
			fmt.Println("I1 + ",t)
		channel <- folder
		files, err := ioutil.ReadDir(string(folder))
		if err != nil {
			*errChannel <- err
		}
		// TODO somethink wrong with channels
			fmt.Println("I2 + ",t)
		for _, file := range files {
			fmt.Println("I2.5 + ",t,"\t",file)
			if file.IsDir() {
			fmt.Println("I3 + ",t)
				if isIgnoreFolder(file.Name()) {
					continue
				}
			fmt.Println("I4 + ",t)
				in := Folder(string(folder) + "\\" + file.Name())
				fs := getInternalDirectory(in, errChannel)
				for f := range fs {
					channel <- f
				}
			}
		}
	}()
	return channel
}
