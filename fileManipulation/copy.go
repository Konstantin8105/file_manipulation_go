package fileManipulation 

import(
	"io"
	"fmt"
	"os"
	"strings"
)

// Copy - copy files
func Copy(inputFileName, outputFileName string) error {

	if len(inputFileName) == 0{
		return fmt.Errorf("inputFileName is zero: %s",inputFileName)
	}


	if len(outputFileName) == 0{
		return fmt.Errorf("inputFileName is zero: %s",outputFileName)
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

			//err := md5checkFileCompare(inputFileName, outputFileName)
			//if err != nil{
			//	errChannel <- err
			//	return
			//}

	err = outputFile.Sync()
 	if err != nil {
 		return err
 	}

 	return nil
 }

 func convert(file fileParam, inputFolder, outputFolder Folder) (in string,out string, folder Folder,err error) {

	 if file == nil{
	 	return nil,nil,nil,fmt.Errorf("(convert): file is empty")
	 }

	 if len(string(inputFolder))==0 {
	 	return nil,nil,nil,fmt.Errorf("(convert): inputFolder is empty")
	 }

	 if len(string(outputFolder))==0 {
	 	return nil,nil,nil,fmt.Errorf("(convert): outputFolder is empty")
	 }

	 fmt.Println("--------")
	 fmt.Println("file.fileInfo = ",file.fileInfo)
	 fmt.Println("file.path     = ",file.path)
	 fmt.Println("input         = ",inputFolder)
	 fmt.Println("output        = ",outputFolder)
	 fmt.Println("--------")

	 folder = strings.Replace(string(file.path),"\\","-",-1)
	 folder =  strings.Replace(folder,":","_",-1)
	 folder =  strings.Replace(folder," ","_",-1)

	 in = fmt.Sprintf("%s%s",file.path,file.fileInfo.Name())
	 out = fmt.Sprintf("%s\\%s\\%s",outputFolder,folder,file.fileInfo.Name())

	 fmt.Println("========")
	 fmt.Println("in  = ",in)
	 fmt.Println("out = ",out)
	 fmt.Println("========")


	 return in,out,nil,nil
 }
