package fileManipulation

import (
	"fmt"
	"io/ioutil"
)

//
// var excludeFolders = []folder{
// 	".git",
// }

func search(inputFolder Folder, internalPaths []Folder) chan fileParam {
	toFilter := make(chan fileParam, 1000)
	go func() {
	begin:
		folderLists := make([]Folder, 0)
		for _, internalPath := range internalPaths {

			// printAll := false

			//X:\2 Project Execution\Steel Structure Calculations\EF16003\STAAD model
			// if string(internalPath) == "\\EF16003\\STAAD model" {
			// 	fmt.Println(fmt.Sprintf("Search in folder: %s%s", inputFolder, internalPath))
			// 	printAll = true
			// }

			//TODO: Mark folders with std files

			files, err := ioutil.ReadDir(fmt.Sprintf("%s\\%s", inputFolder, internalPath))
			if err != nil {
				fmt.Println("Error: Search function - cannot read dir")
				return
			}
			// if len(files) == 0 {
			// 	fmt.Println("Empty folder:", inputFolder, internalPath)
			// }
			for _, file := range files {
				// if printAll {
				// 	fmt.Println(file.Name())
				// }
				if file.IsDir() {
					// isExclude := false
					// for _, excludeFolder := range excludeFolders {
					// 	if file.Name() == string(excludeFolder) {
					// 		isExclude = true
					// 	}
					// }
					// if !isExclude {
					folderLists = append(folderLists, folder(fmt.Sprintf("%s\\%s", internalPath, file.Name())))
					// }
				} else {
					p := new(fileParam)
					p.fileInfo = file
					p.basePath = inputFolder
					p.path = internalPath
					//fmt.Println("SEND FROM SEARCH:", p.Sting())
					toFilter <- *p
				}
			}
		}
		if len(folderLists) > 0 {
			internalPaths = folderLists
			fmt.Print(".")
			goto begin
		}
		close(toFilter)
	}()
	return toFilter
}
