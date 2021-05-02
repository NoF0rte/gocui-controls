package ui

import "github.com/awesome-gocui/gocui"

type Length int

const (
	Fill Length = iota - 1
	Auto
)

func (l Length) toInt() int {
	return int(l)
}

type Element interface {
	gocui.Manager
	Parent() Container
	SetParent(Container)
	X() int
	Y() int
	SetX(int)
	SetY(int)
	AbsX() int
	AbsY() int
	Width() Length
	Height() Length
	SetWidth(Length)
	SetHeight(Length)
	ActualWidth() int
	ActualHeight() int
	MaxX() int
	MaxY() int
	MaxAbsX() int
	MaxAbsY() int
	Type() string
	Id() string
	SetKeybindings(*gocui.Gui) error
}

type Container interface {
	Element
	Children() []Element
	GetChild(string) Element
	Append(Element)
}
