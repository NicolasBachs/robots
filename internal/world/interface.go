package world

import "robots/internal/geo"

type World interface {
	IsOutOfBounds(geo.Position) bool
	AddScentPosition(geo.Position)
	IsScentPosition(geo.Position) bool
}
