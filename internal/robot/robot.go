package robot

import (
	"errors"
	"fmt"
	"robots/internal/geo"
	"robots/internal/world"
)

type robot struct {
	position     geo.Position
	orientation  geo.Orientation
	isLost       bool
	instructions string
}

func NewRobot(position geo.Position, orientation geo.Orientation, instructions string) (Robot, error) {
	robot := &robot{
		position:    position,
		orientation: orientation,
	}

	err := robot.AddInstructions(instructions)

	return robot, err
}

func (r *robot) TurnLeft() {
	switch r.orientation {
	case geo.North:
		r.orientation = geo.West
	case geo.South:
		r.orientation = geo.East
	case geo.East:
		r.orientation = geo.North
	case geo.West:
		r.orientation = geo.South
	}
	fmt.Println("Turning left")
}

func (r *robot) TurnRight() {
	switch r.orientation {
	case geo.North:
		r.orientation = geo.East
	case geo.South:
		r.orientation = geo.West
	case geo.East:
		r.orientation = geo.South
	case geo.West:
		r.orientation = geo.North
	}
	fmt.Println("Turning right")
}

func (r *robot) MoveForward(w world.World) {
	newPosition := r.position

	switch r.orientation {
	case geo.North:
		newPosition.Y += 1
	case geo.South:
		newPosition.Y -= 1
	case geo.East:
		newPosition.X += 1
	case geo.West:
		newPosition.X -= 1
	}

	if w.IsOutOfBounds(newPosition) {
		if w.IsScentPosition(r.position) {
			fmt.Println("Move forward ignored to avoid falling, this position has scent")
			return
		} else {
			w.AddScentPosition(r.position)
			r.isLost = true
			fmt.Println("Robot is lost (the robot left a scent before getting lost)")
			return
		}
	}

	fmt.Println("Moving forward")
	r.position = newPosition
}

func (r *robot) GetPosition() geo.Position {
	return r.position
}

func (r *robot) GetOrientation() geo.Orientation {
	return r.orientation
}

func (r *robot) AddInstructions(newInstructions string) error {
	instructions := r.instructions + newInstructions

	if len(instructions) > 100 {
		return errors.New("instructions too long")
	}

	r.instructions = instructions

	return nil
}

func (r *robot) RunInstrunctions(world world.World) (string, error) {
	for _, char := range r.instructions {
		if r.isLost {
			break
		}

		if char == 'F' {
			r.MoveForward(world)
			continue
		}

		if char == 'L' {
			r.TurnLeft()
			continue
		}

		if char == 'R' {
			r.TurnRight()
			continue
		}

		fmt.Printf("Invalid instruction %c\n", char)
	}

	r.instructions = ""

	return r.Output(), nil
}

func (r *robot) Output() string {
	output := fmt.Sprintf("%d %d %s", r.position.X, r.position.Y, r.orientation)

	if r.isLost {
		output += " LOST"
	}

	fmt.Println(output)

	return output
}
