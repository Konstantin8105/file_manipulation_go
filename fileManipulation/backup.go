package fileManipulation

import "runtime"

//BackUp - copy files from inputFolder to outputFolder
func BackUp(inputFolder, outputFolder Folder) error {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	inputFileFlow, errChannel := getInputFilesFlow(inputFolder)
	defer close(*errChannel)
	copyFileFlow := filterInputFiles(inputFileFlow, outputFolder, errChannel)
	success := copyFiles(copyFileFlow, errChannel)
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
		// TODO
	}()
	return inputFiles, &errFunc
}

func filterInputFiles(inputFiles <-chan fileParam, outputFolder Folder, errChannel *(chan error)) <-chan fileParam {
	copyFiles := make(chan fileParam)
	go func() {
		defer close(copyFiles)
		// TODO
	}()
	return copyFiles
}

func copyFiles(copyFiles <-chan fileParam, errChannel *(chan error)) chan bool {
	success := make(chan bool)
	go func() {
		// TODO
		success <- true
	}()
	return success
}
