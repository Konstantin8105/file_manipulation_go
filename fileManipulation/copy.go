package fileManipulation

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyWithCheckingMd5 - copy file with chacking md5
func CopyWithCheckingMd5(inputFileName, outputFileName string) error {
	err := copyFile(inputFileName, outputFileName)
	if err != nil {
		return err
	}

	// copy of time
	inputStat, err := os.Stat(inputFileName)
	if err != nil {
		return err
	}
	if err = os.Chtimes(outputFileName, inputStat.ModTime(), inputStat.ModTime()); err != nil {
		return err
	}

	// check hash md5
	in, err := hashFileMd5(inputFileName)
	if err != nil {
		return err
	}

	out, err := hashFileMd5(outputFileName)
	if err != nil {
		return err
	}

	if in != out {
		return fmt.Errorf("hash md5 files is different")
	}
	return nil
}

// Copy - copy files
func copyFile(inputFileName, outputFileName string) (err error) {

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
	defer func() {
		errFile := inputFile.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer func() {
		errFile := outputFile.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

func hashFileMd5(fileName string) (returnMD5String string, err error) {

	//Open the passed argument and check for any error
	file, err := os.Open(fileName)
	if err != nil {
		return returnMD5String, err
	}
	defer func() {
		errFile := file.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

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

	// Check input data
	if len(string(inputFolder)) == 0 {
		return "", "", "", fmt.Errorf("(convert): inputFolder is empty")
	}

	if len(string(outputFolder)) == 0 {
		return "", "", "", fmt.Errorf("(convert): outputFolder is empty")
	}

	folder = Folder(inputFolder)
	folder = Folder(strings.Replace(string(folder), "\\", "_", -1))
	folder = Folder(strings.Replace(string(folder), ":", "_", -1))
	folder = Folder(strings.Replace(string(folder), " ", "_", -1))

	if len(string(inputFolder)) == len(string(file.path)) {
		folder = Folder(folder)
	} else {
		folder = Folder(filepath.Join(string(folder), string(file.path)[(len(inputFolder)+1):]))
		//fmt.Sprintf("%s\\%s", string(folder), string(file.path)[(len(inputFolder)+1):]))
	}

	// in = fmt.Sprintf("%s\\%s", file.path, file.fileInfo.Name())
	in = filepath.Join(string(file.path), file.fileInfo.Name())
	// out = fmt.Sprintf("%s\\%s\\%s", outputFolder, folder, file.fileInfo.Name())
	out = filepath.Join(string(outputFolder), string(folder), file.fileInfo.Name())

	folder = Folder(fmt.Sprintf("%s\\%s", outputFolder, string(folder)))

	return in, out, folder, nil
}

func createDirectory(folder Folder) error {
	// Check input data
	if len(string(folder)) == 0 {
		return fmt.Errorf("folder for creating is empty")
	}

	if _, err := os.Stat(string(folder)); os.IsNotExist(err) {
		err := os.MkdirAll(string(folder), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("removeFile error : %v\n", err)
		return nil
	}
	return nil
}
