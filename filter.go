package main

import (
	"path"
	"strings"
	"time"
)

func filter(exts []extension, fileFromChannel, movingFile chan fileParam) {
	for {
		file := <-fileFromChannel
		//fmt.Println("TAKE TO FILTER: ", file)
		fileExension := path.Ext(file.fileInfo.Name())
		//fmt.Println("fileExension = ", fileExension)
		//if len(ext) > 0 {
		for _, ext := range exts {
			if len(ext) > 0 && len(fileExension) > 0 {
				//if strings.HasSuffix(file.fileInfo.Name(), string(ext)) && len(ext) > 0 {
				if strings.Compare(string(ext), fileExension[1:]) == 0 {
					if time.Since(file.fileInfo.ModTime()).Hours() > 24.0 {
						//fmt.Println("SEND FROM FILTER", file.Sting())
						movingFile <- file
					}
				}
			}
		}
		//}
	}
}
