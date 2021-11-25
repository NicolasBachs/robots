package input

import (
	"github.com/gin-gonic/gin"
)

type httpGinInput struct {
	context *gin.Context
}

func NewHttpGinInput(context *gin.Context) Input {
	return &httpGinInput{
		context: context,
	}
}

func (i *httpGinInput) ExecuteInstructions() (string, error) {
	var content string

	b, err := i.context.GetRawData()
	if err != nil {
		return "", err
	}

	content = string(b)

	textInput := NewTextInput(content)

	return textInput.ExecuteInstructions()
}
