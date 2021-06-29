package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
	ui "github.com/sm4rtshr1mp/gocui-controls"
)

func quit(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
	return gocui.ErrQuit
}
func selectClickedView(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
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

	root := ui.NewDockPanel("root", ui.Fill, ui.Fill)

	centerBtn := ui.NewButton("center", ui.Auto, ui.Auto)
	centerBtn.Label = "Center"
	root.Dock(centerBtn, ui.Center)

	leftBtn := ui.NewButton("left", ui.Auto, ui.Auto)
	leftBtn.Label = "Left"
	root.Dock(leftBtn, ui.Left)

	rightBtn := ui.NewButton("right", ui.Auto, ui.Auto)
	rightBtn.Label = "Right"
	root.Dock(rightBtn, ui.Right)

	logBox := ui.NewTextBox("log", ui.Fill, 7)
	logBox.Options.Title = "log"
	root.Dock(logBox, ui.Bottom)

	titleDock := ui.NewDockPanel("titleDock", ui.Fill, 3)
	root.Dock(titleDock, ui.Top)

	closeBtn := ui.NewButton("close", 4, 4)
	closeBtn.Label = "\n  "

	closeBtn.OnClick = quit

	titleDock.Dock(closeBtn, ui.Right)

	g.SetManager(root)

	err = root.SetKeybindings(g)
	if err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, selectClickedView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
