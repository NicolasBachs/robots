package util

import (
	"errors"
	"robots/internal/geo"
	"robots/internal/robot"
	"strconv"
	"strings"
)

func InitRobotFromString(robotPosition string, robotInstructions string) (robot.Robot, error) {
	position := strings.Split(robotPosition, " ")

	if len(position) < 3 {
		return nil, errors.New("invalid robot position")
	}

	posXstr, posYstr, orientationStr := position[0], position[1], position[2]

	posX, err := strconv.Atoi(posXstr)
	if err != nil {
		return nil, errors.New("invalid robot position X '" + posXstr + "'")
	}

	posY, err := strconv.Atoi(posYstr)
	if err != nil {
		return nil, errors.New("invalid robot position Y '" + posYstr + "'")
	}

	orientation, err := geo.NewOrientation(orientationStr)
	if err != nil {
		return nil, err
	}

	return robot.NewRobot(geo.Position{X: int16(posX), Y: int16(posY)}, orientation, robotInstructions)
}
