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
	"fmt"
	"io"
	"os"
	"strings"
)

type FileExecutableInfo struct {
	wordCount bool
	lineCount bool
	charCount bool
	filePath  []string
}

var fileExecutable FileExecutableInfo

func extractFileData(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if len(fileExecutable.filePath) >= 1 {
		if err != nil {
			return string(dat), err
		}
	}
	if err != nil {
		text, _ := io.ReadAll(os.Stdin)
		return string(text), nil
	}
	return string(dat), nil

}

func formatResponse(data string) string {

	var result string

	if !fileExecutable.lineCount && !fileExecutable.wordCount && !fileExecutable.charCount { //default behaviour when used with standard input
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

	return result
}

func parseArguements() {
	flagsArgs := os.Args

	for i := 1; i < len(flagsArgs); i++ {
		if flagsArgs[i] == "-l" && !fileExecutable.lineCount { //line count
			fileExecutable.lineCount = true
		} else if flagsArgs[i] == "-w" && !fileExecutable.wordCount { //word count
			fileExecutable.wordCount = true

		} else if flagsArgs[i] == "-m" && !fileExecutable.charCount { //char count
			fileExecutable.charCount = true
		} else {

			fileExecutable.filePath = append(fileExecutable.filePath, flagsArgs[i])

		}
	}

}

func main() {

	parseArguements()
	if len(fileExecutable.filePath) == 0 {
		data, _ := extractFileData(" ")
		result := formatResponse(data)
		fmt.Println(result + " ")
	} else {
		for i := 0; i < len(fileExecutable.filePath); i++ {
			data, err := extractFileData(fileExecutable.filePath[i])
			if err != nil {
				fmt.Println("NO FILE FOUND")
				continue
			}
			result := formatResponse(data)

			fmt.Println(result + " " + fileExecutable.filePath[i])
		}
	}

}
