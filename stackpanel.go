package ui

import "math"

type StackOrientation int

const (
	Horizontal StackOrientation = iota
	Vertical
)

type StackPanel struct {
	container
	orientation StackOrientation
}

func NewStackPanel(id string, width Length, height Length, orientation StackOrientation) *StackPanel {
	return &StackPanel{
		container:   newContainer(id, width, height, Options{}),
		orientation: orientation,
	}
}
func (s *StackPanel) Append(elem Element) {
	previous := s.Last()
	s.container.Append(elem)

	switch s.orientation {
	case Horizontal:
		elem.SetY(0)
		if previous == nil {
			elem.SetX(0)
		} else {
			elem.SetX(previous.MaxAbsX())
		}
	case Vertical:
		elem.SetX(0)
		if previous == nil {
			elem.SetY(0)
		} else {
			elem.SetY(previous.MaxAbsY())
		}
	}
}

// func (s *StackPanel) Layout(g *gocui.Gui) error {
// 	v, err := s.Create2(g, s.AbsX(), s.AbsY(), s.MaxAbsX(), s.MaxAbsY(), byte(0))
// 	if err == gocui.ErrUnknownView {
// 		fmt.Fprintln(v, b.Label)
// 		return nil
// 	}

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
func (s *StackPanel) ActualWidth() int {
	if s.width == Auto {
		switch s.orientation {
		case Horizontal:
			combinedWidth := 0
			for _, child := range s.children {
				combinedWidth += child.ActualWidth()
			}
			return combinedWidth
		case Vertical:
			maxWidth := -1
			for _, child := range s.children {
				maxWidth = int(math.Max(float64(maxWidth), float64(child.ActualWidth())))
			}
			return maxWidth
		}
	}
	return s.element.ActualWidth()
}
func (s *StackPanel) ActualHeight() int {
	if s.width == Auto {
		switch s.orientation {
		case Horizontal:
			maxHeight := -1
			for _, child := range s.children {
				maxHeight = int(math.Max(float64(maxHeight), float64(child.ActualHeight())))
			}
			return maxHeight
		case Vertical:
			combinedHeight := 0
			for _, child := range s.children {
				combinedHeight += child.ActualHeight()
			}
			return combinedHeight
		}
	}
	return s.element.ActualHeight()
}
func (s *StackPanel) MaxAbsX() int {
	return s.AbsX() + s.ActualWidth()
}
func (s *StackPanel) MaxAbsY() int {
	return s.AbsY() + s.ActualHeight()
}
