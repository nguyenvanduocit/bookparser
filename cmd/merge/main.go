package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	var txtlines string
	for i := 1; i <= 185 ; i++ {
		fileContent, err := processFile(fmt.Sprintf("text/page-%d.txt", i))
		if err != nil {
			log.Fatal(err)
		}

		fileContent = strings.TrimSpace(fileContent)

		if len(fileContent) == 0 {
			continue
		}
		lastChar := fileContent[len(fileContent) - 1]

		if lastChar == '.' || lastChar == ':' || lastChar == ';' || lastChar == '!' || unicode.IsUpper(rune(lastChar)) {
			txtlines += "\n\n" + fileContent
		} else {
			txtlines += " " + fileContent
		}
	}

	ioutil.WriteFile("chapters/01.md", []byte(txtlines), fs.ModePerm)
}

func processFile (fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines string
	currentLine := ""
	for scanner.Scan() {
		currentLine += " " + strings.TrimSpace(scanner.Text())

		lastChar := currentLine[len(currentLine) - 1]

		if lastChar == '.' || lastChar == ':' || lastChar == ';' || lastChar == '!' || unicode.IsUpper(rune(lastChar)) {
			txtlines += "\n\n" + currentLine
			currentLine = ""
		}
	}

	if currentLine != "" {
		txtlines += "\n\n" + currentLine
	}

	return txtlines, nil
}
