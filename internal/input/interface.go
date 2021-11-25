package input

type Input interface {
	ExecuteInstructions() (string, error)
}
