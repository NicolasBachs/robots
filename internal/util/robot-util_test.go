package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRobotFromString_Success(t *testing.T) {
	robot, err := InitRobotFromString("1 1 N", "LRF")

	assert.Nil(t, err)
	assert.NotNil(t, robot)
}

func TestInitRobotFromString_InvalidPosition_X(t *testing.T) {
	robot, err := InitRobotFromString("A 5 N", "LRF")

	assert.NotNil(t, err)
	assert.Nil(t, robot)
}

func TestInitRobotFromString_InvalidPosition_Y(t *testing.T) {
	robot, err := InitRobotFromString("5 A N", "LRF")

	assert.NotNil(t, err)
	assert.Nil(t, robot)
}

func TestInitRobotFromString_InvalidPosition_Orientation(t *testing.T) {
	robot, err := InitRobotFromString("5 5 A", "LRF")

	assert.NotNil(t, err)
	assert.Nil(t, robot)
}

func TestInitRobotFromString_InvalidPosition_length(t *testing.T) {
	robot, err := InitRobotFromString("5 5", "LRF")

	assert.NotNil(t, err)
	assert.Nil(t, robot)
}
