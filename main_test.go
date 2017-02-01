package main

import (
	"strings"
	"testing"
)

func TestExtensionsByUnique(t *testing.T) {
	for index, ext := range exts {
		for i2, ext2 := range exts {
			if index != i2 && strings.Compare(string(ext), string(ext2)) == 0 {
				t.Fatal("Extention is not unique")
			}
		}
	}
}

func TestExtensionsByNotZeo(t *testing.T) {
	for _, ext := range exts {
		if len(ext) < 1 {
			t.Fatal("Extention is not zero")
		}
	}
}
