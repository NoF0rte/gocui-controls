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
	g.Cursor = true
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
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorRed

	root := ui.NewDockPanel("root", ui.Fill, ui.Fill)

	menu := ui.NewMenu("mainMenu", ui.Length(15), ui.Length(10))
	menu.Items = []*ui.MenuItem{
		ui.NewMenuItem("Menu Item 1", func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
			ui.Logln("Menu Item 1 Clicked/Executed")
			return nil
		}),
		ui.NewMenuItem("Menu Item 2", func(g *gocui.Gui, v *gocui.View, keyEv *gocui.KeyEvent) error {
			ui.Logln("Menu Item 2 Clicked/Executed")
			return nil
		}),
	}
	menu.Options.Title = "Main Menu"

	root.Dock(menu, ui.Left)

	logBox := ui.NewTextBox("log", ui.Fill, 7)
	logBox.Options.Title = "log"
	root.Dock(logBox, ui.Bottom)

	g.SetManager(root)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, selectClickedView); err != nil {
		log.Panicln(err)
	}

	err = root.SetKeybindings(g)
	if err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
