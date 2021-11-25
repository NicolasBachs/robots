package test

import (
	"io/ioutil"
	"log"
	"os"
	"robots/internal/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteInstructions_Cases(t *testing.T) {
	pwd, _ := os.Getwd()

	files, err := ioutil.ReadDir(pwd + "/input")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		inputContent, err := ioutil.ReadFile(pwd + "/input/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		outputContent, err := ioutil.ReadFile(pwd + "/output/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		inputText := string(inputContent)
		outputText := string(outputContent)

		textInput := input.NewTextInput(inputText)

		result, err := textInput.ExecuteInstructions()

		assert.Nil(t, err)
		assert.Equal(t, outputText, result)
	}
}
