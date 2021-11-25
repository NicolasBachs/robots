package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrientation_North(t *testing.T) {
	//Given
	orientationString := "N"

	//When
	orientation, err := NewOrientation(orientationString)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, North, orientation)
}

func TestNewOrientation_South(t *testing.T) {
	//Given
	orientationString := "S"

	//When
	orientation, err := NewOrientation(orientationString)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, South, orientation)
}

func TestNewOrientation_East(t *testing.T) {
	//Given
	orientationString := "E"

	//When
	orientation, err := NewOrientation(orientationString)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, East, orientation)
}

func TestNewOrientation_West(t *testing.T) {
	//Given
	orientationString := "W"

	//When
	orientation, err := NewOrientation(orientationString)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, West, orientation)
}

func TestNewOrientation_Invalid(t *testing.T) {
	//Given
	orientationString := "INVALID"

	//When
	_, err := NewOrientation(orientationString)

	//Then
	assert.NotNil(t, err)
}
