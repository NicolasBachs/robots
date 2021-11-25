package world

import (
	"robots/cmd/robots/config"
	"robots/internal/geo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangularWorld_Success(t *testing.T) {
	world, err := NewRectangularWorld(1, 1)

	assert.Nil(t, err)
	assert.NotNil(t, world)
}

func TestNewRectangularWorld_TooBig_Width(t *testing.T) {
	world, err := NewRectangularWorld(config.MAX_WIDTH_RECTANGULAR_WORLD+1, 1)

	assert.NotNil(t, err)
	assert.Nil(t, world)
}

func TestNewRectangularWorld_TooBig_Height(t *testing.T) {
	world, err := NewRectangularWorld(1, config.MAX_HEIGHT_RECTANGULAR_WORLD+1)

	assert.NotNil(t, err)
	assert.Nil(t, world)
}

func TestIsOutOfBounds_false(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	position := geo.Position{
		X: 0,
		Y: 0,
	}

	//When
	result := world.IsOutOfBounds(position)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func TestIsOutOfBounds_true_X(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	position := geo.Position{
		X: 10,
		Y: 0,
	}

	//When
	result := world.IsOutOfBounds(position)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestIsOutOfBounds_true_Y(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	position := geo.Position{
		X: 0,
		Y: 10,
	}

	//When
	result := world.IsOutOfBounds(position)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestAddScentPosition_IsScentPosition_true(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	scentPosition := geo.Position{
		X: 0,
		Y: 0,
	}

	world.AddScentPosition(scentPosition)

	//When
	result := world.IsScentPosition(scentPosition)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestAddScentPosition_IsScentPosition_false(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	scentPosition := geo.Position{
		X: 0,
		Y: 0,
	}

	noScentPosition := geo.Position{
		X: 1,
		Y: 1,
	}

	world.AddScentPosition(scentPosition)

	//When
	result := world.IsScentPosition(noScentPosition)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func TestAddScentPosition_IsScentPosition_sameXCoordinate(t *testing.T) {
	//Given
	world, err := NewRectangularWorld(1, 1)

	scentPosition1 := geo.Position{
		X: 0,
		Y: 0,
	}

	scentPosition2 := geo.Position{
		X: 0,
		Y: 1,
	}

	world.AddScentPosition(scentPosition1)
	world.AddScentPosition(scentPosition2)

	//When
	result1 := world.IsScentPosition(scentPosition1)
	result2 := world.IsScentPosition(scentPosition2)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, true, result1)
	assert.Equal(t, true, result2)
}
