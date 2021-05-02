package ui

type TextBox struct {
	element
}

func NewTextBox(id string, width Length, height Length) *TextBox {
	return &TextBox{
		element: newElement(id, width, height, Options{
			Wrap:       true,
			Autoscroll: true,
			Frame:      true,
		}),
	}
}
