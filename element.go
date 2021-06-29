package ui

import (
	"github.com/awesome-gocui/gocui"
)

var gui *gocui.Gui

type Keybinding struct {
	Key     interface{}
	Mod     gocui.Modifier
	Handler func(*gocui.Gui, *gocui.View, *gocui.KeyEvent) error
}

func NewGui(mode gocui.OutputMode, supportOverlaps bool) (*gocui.Gui, error) {
	var err error
	gui, err = gocui.NewGui(mode, supportOverlaps)
	return gui, err
}

type element struct {
	id            string
	elemType      string
	x, y          int
	width, height Length
	parent        Container
	Options       Options
	keybindings   []*Keybinding
}

type Options struct {
	Title                  string
	Editable               bool
	Wrap                   bool
	Autoscroll             bool
	Highlight              bool
	Frame                  bool
	FrameColor             gocui.Attribute
	BgColor, FgColor       gocui.Attribute
	SelBgColor, SelFgColor gocui.Attribute
}

func newElement(id string, width Length, height Length, options Options) element {
	return element{
		id:      id,
		width:   width,
		height:  height,
		Options: options,
	}
}

func (e *element) Create(g *gocui.Gui) (*gocui.View, error) {
	return e.Create2(g, e.AbsX(), e.AbsY(), e.MaxAbsX(), e.MaxAbsY(), byte(0))
}
func (e *element) Create2(g *gocui.Gui, x0 int, y0 int, x1 int, y1 int, overlaps byte) (*gocui.View, error) {
	v, err := g.View(e.id)
	if err == nil {
		return v, nil
	}

	if err != gocui.ErrUnknownView {
		return nil, err
	}

	v, err = g.SetView(e.id, x0, y0, x1, y1, overlaps)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	v.Title = e.Options.Title
	v.Editable = e.Options.Editable
	v.Wrap = e.Options.Wrap
	v.Autoscroll = e.Options.Autoscroll
	v.Highlight = e.Options.Highlight
	v.Frame = e.Options.Frame
	v.BgColor = e.Options.BgColor
	v.FgColor = e.Options.FgColor
	v.SelBgColor = e.Options.SelBgColor
	v.SelFgColor = e.Options.SelFgColor

	return v, err
}
func (e *element) Layout(g *gocui.Gui) error {
	_, err := e.Create(g)
	return err
}
func (e *element) X() int {
	return e.x
}
func (e *element) SetX(x int) {
	e.x = x
}
func (e *element) Y() int {
	return e.y
}
func (e *element) SetY(y int) {
	e.y = y
}
func (e *element) AbsX() int {
	x := e.x
	if e.parent != nil {
		x += e.parent.AbsX()
	}
	return x
}
func (e *element) AbsY() int {
	y := e.y
	if e.parent != nil {
		y += e.parent.AbsY()
	}
	return y
}
func (e *element) Width() Length {
	return e.width
}
func (e *element) SetWidth(width Length) {
	e.width = width
}
func (e *element) ActualWidth() int {
	switch e.width {
	case Fill:
		if e.parent == nil {
			width, _ := gui.Size()
			return width - 1
		}
		return e.parent.ActualWidth()
	}

	return e.width.toInt()
}
func (e *element) Height() Length {
	return e.height
}
func (e *element) SetHeight(height Length) {
	e.height = height
}
func (e *element) ActualHeight() int {
	switch e.height {
	case Fill:
		if e.parent == nil {
			_, height := gui.Size()
			return height - 1
		}
		return e.parent.ActualHeight()
	}

	return e.height.toInt()
}
func (e *element) MaxX() int {
	return e.x + e.ActualWidth()
}
func (e *element) MaxY() int {
	return e.y + e.ActualHeight()
}
func (e *element) MaxAbsX() int {
	return e.AbsX() + e.ActualWidth()
}
func (e *element) MaxAbsY() int {
	return e.AbsY() + e.ActualHeight()
}
func (e *element) Parent() Container {
	return e.parent
}
func (e *element) SetParent(c Container) {
	e.parent = c
}
func (e *element) Type() string {
	return e.elemType
}
func (e *element) Id() string {
	return e.id
}
func (e *element) SetKeybindings(g *gocui.Gui) error {
	for _, keybinding := range e.keybindings {
		err := g.SetKeybinding(e.id, keybinding.Key, keybinding.Mod, keybinding.Handler)
		if err != nil {
			return err
		}
	}
	return nil
}
func (e *element) AddKeybinding(keybinding *Keybinding) {
	e.keybindings = append(e.keybindings, keybinding)
}
