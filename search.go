package main

import (
	"fmt"
	"io/ioutil"
)

var excludeFolders = []folder{
	".git",
}

func search(inputFolder folder, internalPaths []folder) chan fileParam {
	toFilter := make(chan fileParam)
	go func() {
	begin:
		folderLists := make([]folder, 0)
		for _, internalPath := range internalPaths {
			files, err := ioutil.ReadDir(fmt.Sprintf("%s\\%s", inputFolder, internalPath))
			if err != nil {
				return
			}
			for _, file := range files {
				if file.IsDir() {
					isExclude := false
					for _, excludeFolder := range excludeFolders {
						if file.Name() == string(excludeFolder) {
							isExclude = true
						}
					}
					if !isExclude {
						folderLists = append(folderLists, folder(fmt.Sprintf("%s\\%s", internalPath, file.Name())))
					}
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
