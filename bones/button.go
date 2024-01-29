package bones

type ButtonData struct {
	Content string
	Form    string
	Type    string
}

func (bd ButtonData) Add() error {
	return nil
}
