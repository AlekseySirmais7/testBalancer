package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const subStrForFind = `class="firstHeading" >`

func readFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	pageHtml := string(b)
	titleIndex := strings.Index(pageHtml, subStrForFind)
	titleIndex += len(subStrForFind)
	pageHtmlWithServerNumber := pageHtml[:titleIndex] + " SERVER №" + serverNumberStr + " " + pageHtml[titleIndex:]
	return pageHtmlWithServerNumber, err
}
