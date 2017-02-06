package main

import (
	"fmt"
	"path"
	"strings"
	"time"
)

func filter(fromSearch chan fileParam) chan fileParam {
	movingFile := make(chan fileParam, 1000)
	go func() {
		for file := range fromSearch {

			fileExension := path.Ext(file.fileInfo.Name())
			isFound := false

			if time.Since(file.fileInfo.ModTime()).Hours() < 24.0 {
				continue
			}
			for _, ext := range exts {
				if len(ext) > 0 && len(fileExension) > 0 {
					if strings.Compare(string(ext), fileExension[1:]) == 0 {
						fmt.Println("SEND FROM FILTER", file.Sting())
						movingFile <- file
						isFound = true
						break
					}
				}
			}
			if !isFound {
				for _, suffix := range suffixs {
					if len(suffix) > 0 && len(fileExension) > 0 {
						if strings.HasSuffix(file.fileInfo.Name(), string(suffix)) {
							fmt.Println("SEND FROM FILTER BY SUFFIX --", file.Sting())
							movingFile <- file
							break
						}
					}
				}
			}
		}
		close(movingFile)
	}()
	return movingFile
}
