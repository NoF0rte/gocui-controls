package ui

import "github.com/awesome-gocui/gocui"

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
