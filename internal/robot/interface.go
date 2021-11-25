package robot

import (
	"robots/internal/geo"
	"robots/internal/world"
)

type Robot interface {
	TurnLeft()
	TurnRight()
	MoveForward(world.World)
	GetPosition() geo.Position
	GetOrientation() geo.Orientation
	AddInstructions(string) error
	RunInstrunctions(world.World) (string, error)
	Output() string
}
