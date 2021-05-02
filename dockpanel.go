package ui

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

func (d *DockPanel) Dock(elem Element, dockType DockType) {
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
