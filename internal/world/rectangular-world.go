package world

import (
	"errors"
	"robots/cmd/robots/config"
	"robots/internal/geo"
)

type rectangularWorld struct {
	sizeX          int16
	sizeY          int16
	scentPositions map[int16]map[int16]bool //{posX1: {posY1: bool, posY2: bool}, posX2: {posY1}}
}

func NewRectangularWorld(sizeX, sizeY int16) (World, error) {
	if sizeX > config.MAX_WIDTH_RECTANGULAR_WORLD || sizeY > config.MAX_HEIGHT_RECTANGULAR_WORLD {
		return nil, errors.New("the world is too big")
	}
	return &rectangularWorld{
		sizeX: sizeX,
		sizeY: sizeY,
	}, nil
}

func (r *rectangularWorld) IsOutOfBounds(pos geo.Position) bool {
	if pos.X < 0 || pos.X > r.sizeX {
		return true
	}
	if pos.Y < 0 || pos.Y > r.sizeY {
		return true
	}
	return false
}

func (r *rectangularWorld) AddScentPosition(pos geo.Position) {
	if yMap, containsX := r.scentPositions[pos.X]; containsX {
		yMap[pos.Y] = true
		r.scentPositions[pos.X] = yMap
	} else {
		r.scentPositions = map[int16]map[int16]bool{pos.X: {pos.Y: true}}
	}
}

func (r *rectangularWorld) IsScentPosition(pos geo.Position) bool {
	if yMap, containsX := r.scentPositions[pos.X]; containsX {
		if value, containsY := yMap[pos.Y]; containsY {
			return value
		}
	}
	return false
}
