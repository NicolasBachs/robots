package robot

import (
	"robots/internal/geo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRobot(t *testing.T) {
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.North, "LFR")

	assert.Nil(t, err)
	assert.NotNil(t, robot)
}

func TestTurnRight_fromNorth(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.North, "LFR")

	//When
	robot.TurnRight()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.East, robot.GetOrientation())
}

func TestTurnRight_fromEast(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.East, "LFR")

	//When
	robot.TurnRight()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.South, robot.GetOrientation())
}

func TestTurnRight_fromSouth(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.South, "LFR")

	//When
	robot.TurnRight()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.West, robot.GetOrientation())
}

func TestTurnRight_fromWest(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.West, "LFR")

	//When
	robot.TurnRight()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.North, robot.GetOrientation())
}

func TestTurnLeft_fromNorth(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.North, "LFR")

	//When
	robot.TurnLeft()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.West, robot.GetOrientation())
}

func TestTurnLeft_fromEast(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.East, "LFR")

	//When
	robot.TurnLeft()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.North, robot.GetOrientation())
}

func TestTurnLeft_fromSouth(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.South, "LFR")

	//When
	robot.TurnLeft()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.East, robot.GetOrientation())
}

func TestTurnLeft_fromWest(t *testing.T) {
	//Given
	robot, err := NewRobot(geo.Position{X: 0, Y: 0}, geo.West, "LFR")

	//When
	robot.TurnLeft()

	//Then
	assert.Nil(t, err)
	assert.Equal(t, geo.South, robot.GetOrientation())
}
