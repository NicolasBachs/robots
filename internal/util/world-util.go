package util

import (
	"errors"
	"robots/internal/world"
	"strconv"
	"strings"
)

func InitWorldFromString(worldDimension string) (world.World, error) {
	dimension := strings.Split(worldDimension, " ")

	if len(dimension) < 2 {
		return nil, errors.New("invalid world dimension")
	}

	worldWidthStr, worldHeightStr := dimension[0], dimension[1]

	worldWidth, err := strconv.Atoi(worldWidthStr)
	if err != nil {
		return nil, errors.New("invalid world width")
	}

	worldHeight, err := strconv.Atoi(worldHeightStr)
	if err != nil {
		return nil, errors.New("invalid world height")
	}

	return world.NewRectangularWorld(int16(worldWidth), int16(worldHeight))
}
