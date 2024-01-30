package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	entries, err := os.ReadDir("./tests")
	if err != nil {
		log.Fatal(err)
	}
	var inputs []string
	var outputs []string
	for _, e := range entries {
		if strings.Contains(e.Name(), `.a`) {
			outputs = append(outputs, e.Name())
		} else {
			inputs = append(inputs, e.Name())
		}

	}

	for i, name := range inputs {

		t.Run(name, func(t *testing.T) {
			file, err := os.Open("./tests/" + name)
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}

			reader := bufio.NewReader(file)
			var buffer bytes.Buffer
			writer := bufio.NewWriter(&buffer)
			Solution(reader, writer)

			content, err := os.ReadFile("./tests/" + outputs[i])
			if err != nil {
				log.Fatal(err)
			}
			testContent := replaceReplacer(string(content))

			if err != nil {
				log.Fatal(err)
			}
			resultContent := buffer.String()
			if resultContent != testContent {
				t.Errorf("Solution() = \n%v, want \n%v", resultContent, testContent)
			}
		})
	}
}

var replacer = strings.NewReplacer(
	"\r\n", "\n",
)

func replaceReplacer(s string) string {
	return replacer.Replace(s)
}
