package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
	ui "github.com/sm4rtshr1mp/gocui-controls"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func selectClickedView(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(v.Name())
	return err
}

func main() {
	g, err := ui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Mouse = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorRed

	dockPanel := ui.NewDockPanel("root", ui.Fill, ui.Fill)

	centerBtn := ui.NewButton("center", ui.Auto, ui.Auto)
	centerBtn.Label = "Center"
	dockPanel.DockElem(centerBtn, ui.Center)

	leftBtn := ui.NewButton("left", ui.Auto, ui.Auto)
	leftBtn.Label = "Left"
	dockPanel.DockElem(leftBtn, ui.Left)

	rightBtn := ui.NewButton("right", ui.Auto, ui.Auto)
	rightBtn.Label = "Right"
	dockPanel.DockElem(rightBtn, ui.Right)

	logBox := ui.NewTextBox("log", ui.Fill, 7)
	logBox.Options.Title = "log"
	dockPanel.DockElem(logBox, ui.Bottom)

	titleDock := ui.NewDockPanel("titleDock", ui.Fill, 3)
	dockPanel.DockElem(titleDock, ui.Top)

	closeBtn := ui.NewButton("close", ui.Auto, ui.Auto)
	closeBtn.Label = "X"

	closeBtn.OnClick = quit

	titleDock.DockElem(closeBtn, ui.Right)

	g.SetManager(dockPanel)

	err = dockPanel.SetKeybindings(g)
	if err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.MouseRelease, gocui.ModNone, selectClickedView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
