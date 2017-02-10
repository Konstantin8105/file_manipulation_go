package fileManipulation

//
// func movingFiles(outputFolder folder, fileFromChannel chan fileParam) chan bool {
// 	success := make(chan bool)
// 	go func() {
// 		var counter uint64
// 		for file := range fileFromChannel {
// 			counter++
// 			fmt.Println("counter", counter)
// 			fmt.Println("TAKE IN MOVING", file.Sting())
//
// 			// create general folder
// 			generalFolderName := strings.Replace(string(file.basePath), "\\", "_", -1)
// 			generalFolderName = strings.Replace(generalFolderName, ":", "_", -1)
// 			generalFolderName = strings.Replace(generalFolderName, " ", "_", -1)
// 			generalFolderName = fmt.Sprintf("%s\\%s\\%s", outputFolder, generalFolderName, string(file.path)[1:])
//
// 			// create internal folder
// 			if _, err := os.Stat(generalFolderName); os.IsNotExist(err) {
// 				os.MkdirAll(generalFolderName, os.ModePerm)
// 				fmt.Println("counter", counter)
// 				fmt.Println("create folder -->", generalFolderName)
// 			} else {
// 				// fmt.Println("counter", counter)
// 				// fmt.Println("folder is exist -->", generalFolderName)
// 			}
//
// 			inputFileName := fmt.Sprintf("%s\\%s\\%s", string(file.basePath), string(file.path)[1:], file.fileInfo.Name())
// 			outputFileName := fmt.Sprintf("%s\\%s", generalFolderName, file.fileInfo.Name())
// 			err := fileManipulation.copy(inputFileName, outputFileName)
// 			if err != nil {
// 				return
// 			}
//
// 			err = outputFile.Sync()
// 			if err != nil {
// 				fmt.Println("counter", counter)
// 				fmt.Println("Can not flush the file ", outputFileName)
// 				return
// 			}
// 			inputFile.Close()
//
// 			// check md5
//
// 			// remove resource
//
// 			fmt.Println("Remove file ->", inputFileName)
// 			err = os.Remove(inputFileName)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 				return
// 			}
// 		}
// 		success <- true
// 		close(success)
// 	}()
// 	return success
// }

