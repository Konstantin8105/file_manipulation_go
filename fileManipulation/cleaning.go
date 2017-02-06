package fileManipulation

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
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
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	staadFolders, errChannel := getStaadFolders(inputFolder)
	defer close(*errChannel)
	tempFiles := filterTempStaadFiles(staadFolders, errChannel)
	success := moveTempStaadFiles(tempFiles, inputFolder, outputFolder, errChannel)

	select {
	case <-success:
		return nil
	case err := <-*errChannel:
		return err
	}
}

func moveTempStaadFiles(tempFiles <-chan fileParam, inputFolder, outputFolder Folder, errChannel *chan error) <-chan bool {
	success := make(chan bool)
	go func() {
		defer close(success)
		for tempFile := range tempFiles {
			// TODO moving
			fmt.Println("temp file ", tempFile.path, tempFile.fileInfo.Name())
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
							fmt.Println("Maybe add for delete list : ", file.Name())
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
		folders := make([]Folder, 1)
		folders = append(folders, inputFolder)
		err := getInternalDirectory(inputFolder, folders)
		if err != nil {
			errFunc <- err
			return
		}
		for _, folder := range folders {
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

func getInternalDirectory(folder Folder, internalDir []Folder) error {
	files, err := ioutil.ReadDir(string(folder))
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			var intDir Folder = Folder(string(folder) + "//" + file.Name())
			internalDir = append(internalDir, intDir)
			err := getInternalDirectory(intDir, internalDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
