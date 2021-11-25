package util

import (
	"fmt"
	"robots/cmd/robots/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitWorldFromString_Success(t *testing.T) {
	world, err := InitWorldFromString("5 5")

	assert.Nil(t, err)
	assert.NotNil(t, world)
}

func TestInitWorldFromString_InvalidDimension_width(t *testing.T) {
	world, err := InitWorldFromString("A 5")

	assert.NotNil(t, err)
	assert.Nil(t, world)
}

func TestInitWorldFromString_InvalidDimension_height(t *testing.T) {
	world, err := InitWorldFromString("5 A")

	assert.NotNil(t, err)
	assert.Nil(t, world)
}

func TestInitWorldFromString_InvalidDimension_length(t *testing.T) {
	world, err := InitWorldFromString("1")

	assert.NotNil(t, err)
	assert.Nil(t, world)
}

func TestInitWorldFromString_TooBig(t *testing.T) {
	dimensions := fmt.Sprintf(
		"%d %d",
		config.MAX_WIDTH_RECTANGULAR_WORLD+1,
		config.MAX_HEIGHT_RECTANGULAR_WORLD+1,
	)

	world, err := InitWorldFromString(dimensions)

	assert.NotNil(t, err)
	assert.Nil(t, world)
}
