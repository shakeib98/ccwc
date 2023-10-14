package main

import (
	"fmt"
	"io/fs"
	"os"
)

func openFile() (fs.FileInfo, error) {
	fi, err := os.Stat("./text.txt")
	return fi, err
}

func checkFileWordCount() {
	fi, err := openFile()

	if err != nil {
		panic("File name does not exists")
	}

	// get the size
	size := fi.Size()

	fmt.Println(size)
}

func main() {
	checkFileWordCount()
}
