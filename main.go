/*
 -l to count number of lines
 -w to count number of words
 -m to count number of chars
 -c to count number of bytes

 If command is called on itself then after flag, filepath should be provided

 If command is called by a pipeline command then filepath is not needed.
 But if filepath is also provided with piping, then standard input's priority
 will become secondary. And filepath will be executed.

 No flag can be run more than once

 Command should be in this pattern "command flags filepath"

 In standard input, the file will handle the standard input if no path is provided
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FileExecutableInfo struct {
	wordCount  bool
	lineCount  bool
	charCount  bool
	bytesCount bool
	filePath   string
}

var fileExecutable FileExecutableInfo

func extractFileData() string {
	filePath := os.Args[len(os.Args)-1]
	dat, err := os.ReadFile(filePath)

	if err != nil {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString(io.SeekEnd)
		return text
	} else {
		fileExecutable.filePath = filePath
		return string(dat)
	}
}

func formatResponse(data string) {

	flagsArgs := os.Args

	var result string

	for i := 1; i < len(flagsArgs); i++ {
		if flagsArgs[i] == "-l" { //line count
			fileExecutable.lineCount = true
		} else if flagsArgs[i] == "-c" { //no of bytes
			fileExecutable.bytesCount = true
		} else if flagsArgs[i] == "-w" { //word count
			fileExecutable.wordCount = true

		} else if flagsArgs[i] == "-m" { //char count
			fileExecutable.charCount = true
		} else {
			if i == 1 {
				fileExecutable.lineCount = true
				fileExecutable.wordCount = true
				fileExecutable.charCount = true
			}
		}
	}

	if !fileExecutable.lineCount && !fileExecutable.bytesCount && !fileExecutable.wordCount && !fileExecutable.charCount { //default behaviour when used with standard input
		fileExecutable.lineCount = true
		fileExecutable.wordCount = true
		fileExecutable.charCount = true
	}

	if fileExecutable.lineCount {
		result = result + " " + fmt.Sprint((len(strings.Split(data, "\n"))))
	}

	if fileExecutable.wordCount {
		words := strings.Fields(data)
		counter := 0
		for counter != len(words) {
			counter++
		}
		result = result + " " + fmt.Sprint(counter)
	}

	if fileExecutable.charCount {
		result = result + " " + fmt.Sprint(len([]rune(data)))
	}

	if fileExecutable.bytesCount {
		result = result + " " + fmt.Sprint(len(data))
	}

	result = result + " " + fileExecutable.filePath

	fmt.Println(result)
}

func main() {

	formatResponse(extractFileData())

}
