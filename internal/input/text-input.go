package input

import (
	"errors"
	"robots/internal/util"
	"strings"
)

type textInput struct {
	text string
}

func NewTextInput(text string) Input {
	return &textInput{
		text: text,
	}
}

func (i *textInput) ExecuteInstructions() (string, error) {
	lines := strings.Split(i.text, "\n")

	if len(lines) < 3 {
		return "", errors.New("invalid text file, number of lines are not enough")
	}

	worldDimension := lines[0]

	world, err := util.InitWorldFromString(worldDimension)

	if err != nil {
		return "", errors.New("invalid text file, " + err.Error())
	}

	outputs := ""

	for i := 1; i < len(lines)-1; i += 3 {
		robotPosition := lines[i]
		robotInstructions := lines[i+1]

		robot, err := util.InitRobotFromString(robotPosition, robotInstructions)
		if err != nil {
			return "", errors.New("invalid text file, " + err.Error())
		}

		output, err := robot.RunInstrunctions(world)

		if err != nil {
			return "", errors.New("error running instructions")
		}

		outputs += output + "\n"
	}

	return outputs[:len(outputs)-1], nil
}
