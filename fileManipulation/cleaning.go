package fileManipulation

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"strings"
	"time"
)

// Folder - folder like string
type Folder string

type fileParam struct {
	fileInfo fs.FileInfo
	path     Folder
}

// Cleaning function for cleaning folder from temp file of STAAD
func Cleaning(inputFolder, outputFolder Folder) (err error) {

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

	if strings.HasSuffix(string(outputFolder), string(inputFolder)) {
		return fmt.Errorf("Output folder cannot be inside input folder")
	}

	// use all allowable proccesors
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	// working
	errChannel := make(chan error)
	staadFolders := getStaadFolders(inputFolder, errChannel)
	tempFiles := filterTempStaadFiles(staadFolders, errChannel)
	success := moveTempStaadFiles(tempFiles, inputFolder, outputFolder, errChannel)

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

func moveTempStaadFiles(
	tempFiles <-chan fileParam,
	inputFolder, outputFolder Folder,
	errChannel chan<- error,
) chan struct{} {
	success := make(chan struct{})
	go func() {
		defer func() {
			success <- struct{}{}
		}()
		for tempFile := range tempFiles {
			inputFileName, outputFileName, folder, err := convert(tempFile, inputFolder, outputFolder)
			if err != nil {
				errChannel <- err
				continue
			}
			if err = createDirectory(folder); err != nil {
				errChannel <- err
				continue
			}
			if err = CopyWithCheckingMd5(inputFileName, outputFileName); err != nil {
				errChannel <- err
				continue
			}
			if err = removeFile(inputFileName); err != nil {
				errChannel <- err
				continue
			}
		}
	}()
	return success
}

func filterTempStaadFiles(staadFolders <-chan Folder, errChannel chan<- error) <-chan fileParam {
	tempFiles := make(chan fileParam)
	go func() {
		defer close(tempFiles)
		for folder := range staadFolders {
			files, err := os.ReadDir(string(folder))
			if err != nil {
				errChannel <- err
				return
			}

			for _, file := range files {
				fi, err := file.Info()
				if err != nil {
					errChannel <- err
					continue
				}
				// filter by last 24*3 hours files
				if time.Since(fi.ModTime()).Hours() < 24.0*3 {
					continue
				}
				// filter by temp staad files
				if !file.IsDir() {
					if !isStaadFile(file.Name()) {
						if isStaadTempFile(file.Name()) {
							folder := folder
							tempFiles <- fileParam{fileInfo: fi, path: folder}
							//} else {
							//fmt.Println("Maybe add for delete list : ", file.Name())
						}
					}
				}
			}
		}
	}()
	return tempFiles
}

func getStaadFolders(inputFolder Folder, errFunc chan<- error) <-chan Folder {
	staadFolders := make(chan Folder)
	go func() {
		defer close(staadFolders)
		folders := getInternalDirectory(inputFolder, errFunc)
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
	return staadFolders
}

func (folder Folder) withStaadFiles() (bool, error) {
	if len(string(folder)) == 0 {
		return false, errors.New("Null size of folder")
	}
	files, err := os.ReadDir(string(folder))
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

func getInternalDirectory(folder Folder, errChannel chan<- error) chan Folder {
	channel := make(chan Folder)
	go func() {
		defer close(channel)
		channel <- folder
		files, err := os.ReadDir(string(folder))
		if err != nil {
			errChannel <- err
		}
		for _, file := range files {
			if file.IsDir() && !isIgnoreFolder(file.Name()) {
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
