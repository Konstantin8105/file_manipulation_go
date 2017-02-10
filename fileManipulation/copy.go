package fileManipulation 

import(
	"io"
	"os"
)

// Copy - copy files
func Copy(inputFileName, outputFileName string) error {

 	inputFile, err := os.Open(inputFileName)
 	defer inputFile.Close()
 	if err != nil {
 		return err
 	}

 	outputFile, err := os.Create(outputFileName)
 	defer outputFile.Close()
 	if err != nil {
 		return err
 	}

	_, err = io.Copy(outputFile, inputFile)
 	if err != nil {
 		return err
 	}

	err = outputFile.Sync()
 	if err != nil {
 		return err
 	}

 	return nil
 }
