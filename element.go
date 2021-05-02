package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

var gui *gocui.Gui

type Keybinding struct {
	Key     interface{}
	Mod     gocui.Modifier
	Handler func(*gocui.Gui, *gocui.View) error
}

func NewGui(mode gocui.OutputMode, supportOverlaps bool) (*gocui.Gui, error) {
	var err error
	gui, err = gocui.NewGui(mode, supportOverlaps)
	return gui, err
}

func Log(a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprint(logView, a...)
	}
}
func Logf(format string, a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintf(logView, format, a...)
	}
}
func Logfln(format string, a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintf(logView, format, a...)
		fmt.Fprintln(logView)
	}
}
func Logln(a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintln(logView, a...)
	}
}

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

type Container interface {
	Element
	Children() []Element
	GetChild(string) Element
	Append(Element)
}

type container struct {
	element
	children []Element
}

func newContainer(id string, width Length, height Length, options Options) container {
	return container{
		element: newElement(id, width, height, options),
	}
}

func (c *container) Layout(g *gocui.Gui) error {
	err := c.element.Layout(g)
	if err == gocui.ErrUnknownView {
		for _, child := range c.children {
			err = child.Layout(g)
			if err != nil && err != gocui.ErrUnknownView {
				return err
			}
		}
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
func (c *container) Children() []Element {
	return c.children
}
func (c *container) GetChild(id string) Element {
	for _, child := range c.children {
		if child.Id() == id {
			return child
		}
	}
	return nil
}
func (c *container) Append(elem Element) {
	elem.SetParent(c)
	c.children = append(c.children, elem)
}
func (c *container) SetKeybindings(g *gocui.Gui) error {
	err := c.element.SetKeybindings(g)
	if err != nil {
		return err
	}

	for _, child := range c.children {
		err = child.SetKeybindings(g)
		if err != nil {
			return err
		}
	}

	return nil
}

type DockType int

const (
	Left DockType = iota
	Bottom
	Right
	Top
	Center
)

type DockPanel struct {
	container
	left   Element
	bottom Element
	right  Element
	top    Element
	center Element
}

func NewDockPanel(id string, width Length, height Length) *DockPanel {
	return &DockPanel{
		container: newContainer(id, width, height, Options{}),
	}
}

func (d *DockPanel) DockElem(elem Element, dockType DockType) {
	d.Append(elem)

	maxX := d.ActualWidth()
	maxY := d.ActualHeight()

	actualWidth := elem.ActualWidth()
	actualHeight := elem.ActualHeight()

	switch dockType {
	case Center:
		elem.SetX((maxX - actualWidth) / 2)
		elem.SetY((maxY - actualHeight) / 2)
		d.center = elem
	case Left:
		elem.SetX(0)
		elem.SetY((maxY - actualHeight) / 2)
		d.left = elem
	case Right:
		elem.SetX(maxX - actualWidth)
		elem.SetY((maxY - actualHeight) / 2)
		d.right = elem
	case Top:
		elem.SetX((maxX - actualWidth) / 2)
		elem.SetY(0)
		d.top = elem
	case Bottom:
		elem.SetX((maxX - actualWidth) / 2)
		elem.SetY(maxY - actualHeight)
		d.bottom = elem
	}
}

type Button struct {
	element
	mouseDown       bool
	processingClick bool
	Label           string
	OnClick         func(g *gocui.Gui, v *gocui.View) error
}

func NewButton(id string, width Length, height Length) *Button {
	options := Options{
		Title:     "",
		Frame:     true,
		Wrap:      true,
		Highlight: true,
	}
	btn := &Button{
		element: newElement(id, width, height, options),
	}
	btn.AddKeybinding(&Keybinding{
		Key: gocui.MouseLeft,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View) error {
			btn.mouseDown = true
			return nil
		},
	})
	btn.AddKeybinding(&Keybinding{
		Key: gocui.MouseRelease,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View) error {
			Logfln("New View: %s", v.Name())
			g.SetCurrentView(btn.id)

			if btn.mouseDown {
				btn.mouseDown = false
				if !btn.processingClick && btn.OnClick != nil {
					btn.processingClick = true
					err := btn.OnClick(g, v)
					btn.processingClick = false
					return err
				}
			}
			return nil
		},
	})
	return btn
}

func (b *Button) Layout(g *gocui.Gui) error {
	v, err := b.Create2(g, b.AbsX(), b.AbsY(), b.MaxAbsX(), b.MaxAbsY(), byte(0))
	if err == gocui.ErrUnknownView {
		fmt.Fprintln(v, b.Label)
		return nil
	}

	if err != nil {
		return err
	}
	return nil
}
func (b *Button) ActualWidth() int {
	if b.width == Auto {
		return len(b.Label) + 1
	}
	return b.element.ActualWidth()
}
func (b *Button) ActualHeight() int {
	if b.height == Auto {
		return 2
	}
	return b.element.ActualHeight()
}
func (b *Button) MaxAbsX() int {
	return b.AbsX() + b.ActualWidth()
}
func (b *Button) MaxAbsY() int {
	return b.AbsY() + b.ActualHeight()
}

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
