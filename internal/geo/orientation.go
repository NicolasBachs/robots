package geo

import "errors"

type Orientation int8

const (
	North Orientation = iota
	East
	South
	West
)

func (o Orientation) String() string {
	switch o {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	}
	return "unknown"
}

func NewOrientation(orientation string) (Orientation, error) {
	switch orientation {
	case "N":
		return North, nil
	case "S":
		return South, nil
	case "E":
		return East, nil
	case "W":
		return West, nil
	}

	return -1, errors.New("invalid orientation")
}
