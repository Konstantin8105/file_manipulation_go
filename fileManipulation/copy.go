package fileManipulation

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// CopyWithCheckingMd5 - copy file with chacking md5
func CopyWithCheckingMd5(inputFileName, outputFileName string) error {
	// fmt.Println("outputFileName = ", outputFileName)
	// fmt.Println("Copy #1")
	err := Copy(inputFileName, outputFileName)
	if err != nil {
		return err
	}
	// fmt.Println("Copy #2")
	in, err := hashFileMd5(inputFileName)
	if err != nil {
		return err
	}

	// fmt.Println("Copy #3")
	out, err := hashFileMd5(outputFileName)
	if err != nil {
		return err
	}

	// fmt.Println("Copy #4")
	if in != out {
		return fmt.Errorf("hash md5 files is different")
	}
	// fmt.Println("Copy #5")
	return nil
}

// Copy - copy files
func Copy(inputFileName, outputFileName string) error {

	if len(inputFileName) == 0 {
		return fmt.Errorf("inputFileName is zero: %s", inputFileName)
	}

	if len(outputFileName) == 0 {
		return fmt.Errorf("inputFileName is zero: %s", outputFileName)
	}

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

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

func hashFileMd5(fileName string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(fileName)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func convert(file fileParam, inputFolder, outputFolder Folder) (in string, out string, folder Folder, err error) {

	// if file == nil {
	// 	return "", "", "", fmt.Errorf("(convert): file is empty")
	// }

	if len(string(inputFolder)) == 0 {
		return "", "", "", fmt.Errorf("(convert): inputFolder is empty")
	}

	if len(string(outputFolder)) == 0 {
		return "", "", "", fmt.Errorf("(convert): outputFolder is empty")
	}

	// fmt.Println("--------")
	// fmt.Println("file.fileInfo = ", file.fileInfo)
	// fmt.Println("file.path     = ", file.path)
	// fmt.Println("input         = ", inputFolder)
	// fmt.Println("output        = ", outputFolder)
	// fmt.Println("--------")

	folder = Folder(inputFolder)
	folder = Folder(strings.Replace(string(folder), "\\", "_", -1))
	folder = Folder(strings.Replace(string(folder), ":", "_", -1))
	folder = Folder(strings.Replace(string(folder), " ", "_", -1))
	folder = Folder(fmt.Sprintf("%s\\%s", string(folder), string(file.path)[(len(inputFolder)+1):]))

	in = fmt.Sprintf("%s\\%s", file.path, file.fileInfo.Name())
	out = fmt.Sprintf("%s\\%s\\%s", outputFolder, folder, file.fileInfo.Name())

	folder = Folder(fmt.Sprintf("%s\\%s", outputFolder, string(folder)))
	// fmt.Println("========")
	// fmt.Println("in  = ", in)
	// fmt.Println("out = ", out)
	// fmt.Println("folder = ", folder)
	// fmt.Println("========")

	return in, out, folder, nil
}

func createDirectory(folder Folder) error {
	if len(string(folder)) == 0 {
		return fmt.Errorf("folder for creating is empty")
	}
	if _, err := os.Stat(string(folder)); os.IsNotExist(err) {
		err := os.MkdirAll(string(folder), os.ModePerm)
		if err != nil {
			return err
		}
		// fmt.Println("Create directory: ", string(folder))
	}
	return nil
}

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	// fmt.Println("Remove file : ", fileName)
	return nil
}
