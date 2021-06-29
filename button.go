package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type Button struct {
	element
	mouseDown       bool
	processingClick bool
	Label           string
	OnClick         func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error
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
		Key: gocui.MouseRelease,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
			Logln("MouseLeft")
			btn.mouseDown = true
			return nil
		},
	})
	btn.AddKeybinding(&Keybinding{
		Key: gocui.MouseLeft,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
			Logfln("New View: %s", v.Name())
			g.SetCurrentView(btn.id)

			if btn.mouseDown {
				btn.mouseDown = false
				if !btn.processingClick && btn.OnClick != nil {
					btn.processingClick = true
					err := btn.OnClick(g, v, keyEv)
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
