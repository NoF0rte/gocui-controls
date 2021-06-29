package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type Menu struct {
	element
	Items []*MenuItem
}

func NewMenu(id string, width Length, height Length) *Menu {
	menu := &Menu{
		element: newElement(id, width, height, Options{
			Frame:      true,
			Highlight:  true,
			SelBgColor: gocui.ColorWhite,
			SelFgColor: gocui.ColorRed,
		}),
	}
	menu.elemType = "Menu"

	menu.AddKeybinding(&Keybinding{
		Key: gocui.KeyArrowDown,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View, ke *gocui.KeyEvent) error {
			v.MoveCursor(0, 1)
			return nil
		},
	})
	menu.AddKeybinding(&Keybinding{
		Key: gocui.KeyArrowUp,
		Mod: gocui.ModNone,
		Handler: func(g *gocui.Gui, v *gocui.View, ke *gocui.KeyEvent) error {
			v.MoveCursor(0, -1)
			return nil
		},
	})

	menu.AddKeybinding(&Keybinding{
		Key:     gocui.KeyEnter,
		Mod:     gocui.ModNone,
		Handler: menu.execSelectedItem,
	})
	menu.AddKeybinding(&Keybinding{
		Key:     gocui.MouseLeft,
		Mod:     gocui.ModNone,
		Handler: menu.onClick,
	})
	return menu
}
func (m *Menu) execSelectedItem(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
	selectedItem := m.getSelectedItem(g, v)
	if selectedItem != nil {
		return selectedItem.OnClick(g, v, keyEv)
	}
	return nil
}
func (m *Menu) getSelectedIndex(g *gocui.Gui, v *gocui.View) int {
	var err error

	_, cy := v.Cursor()
	lineNum := cy
	if _, err = v.Line(cy); err != nil {
		lineNum = -1
	}
	return lineNum
}
func (m *Menu) getSelectedItem(g *gocui.Gui, v *gocui.View) *MenuItem {
	index := m.getSelectedIndex(g, v)
	if index == -1 || index > len(m.Items)-1 {
		return nil
	}
	return m.Items[index]
}
func (m *Menu) isMouseClickOnItem(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) bool {
	var line string
	var err error

	my := keyEv.MouseY - m.y
	mx := keyEv.MouseX - m.x
	if line, err = v.Line(my - 1); err != nil {
		return false
	}

	return mx <= len(line)-1
}
func (m *Menu) onClick(g *gocui.Gui, v *gocui.View, ke *gocui.KeyEvent) error {
	g.SetCurrentView(v.Name())

	g.Cursor = false

	selectedItem := m.getSelectedItem(g, v)
	if selectedItem == nil {
		return v.SetCursor(0, len(m.Items)-1)
	}

	// Check if an item was clicked directly
	if m.isMouseClickOnItem(g, v, ke) {
		return selectedItem.OnClick(g, v, ke)
	}

	return nil
}
func (m *Menu) Layout(g *gocui.Gui) error {
	v, err := m.Create2(g, m.AbsX(), m.AbsY(), m.MaxAbsX(), m.MaxAbsY(), byte(0))
	if err == gocui.ErrUnknownView {
		for i, item := range m.Items {
			if i != len(m.Items)-1 {
				fmt.Fprintln(v, item.Label)
			} else {
				fmt.Fprint(v, item.Label)
			}
		}
		return nil
	}

	if err != nil {
		return err
	}
	return nil
}

type MenuItem struct {
	Label   string
	OnClick func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error
}

func NewMenuItem(label string, onClick func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error) *MenuItem {
	return &MenuItem{
		Label:   label,
		OnClick: onClick,
	}
}
