package main

import (
	"fmt"
	"io/ioutil"
)

var excludeFolders = []folder{
	".git",
}

func search(inputFolder folder, internalPath folder, toFilter chan fileParam) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s\\%s", inputFolder, internalPath))
	if err != nil {
		return
	}
	folderLists := make([]folder, 0)
	for _, file := range files {
		if file.IsDir() {
			isEclude := false
			for _, excludeFolder := range excludeFolders {
				if file.Name() == string(excludeFolder) {
					isEclude = true
				}
			}
			if !isEclude {
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
	for _, fold := range folderLists {
		search(inputFolder, fold, toFilter)
	}
}
