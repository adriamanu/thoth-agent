package main

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SplitFilesByLine(fileContent string) []string {
	splittedFile := strings.Split(fileContent, "\n")
	return splittedFile
}