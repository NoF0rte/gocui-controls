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

	root := ui.NewStackPanel("root", ui.Fill, ui.Fill, ui.Vertical)

	btn1 := ui.NewButton("btn1", ui.Auto, ui.Auto)
	btn1.Label = "Button 1"
	root.Append(btn1)

	btn2 := ui.NewButton("btn2", ui.Auto, ui.Auto)
	btn2.Label = "Button 2"
	root.Append(btn2)

	btn3 := ui.NewButton("btn3", ui.Auto, ui.Auto)
	btn3.Label = "Button 3"
	root.Append(btn3)

	horizontalStack := ui.NewStackPanel("horizontalStack", ui.Fill, ui.Auto, ui.Horizontal)
	horizontalStack.Options.Frame = true
	btn4 := ui.NewButton("btn4", ui.Auto, ui.Auto)
	btn4.Label = "Button 4"
	btn4.Options.Frame = false
	horizontalStack.Append(btn4)

	btn5 := ui.NewButton("btn5", ui.Auto, ui.Auto)
	btn5.Label = "Button 5"
	btn5.Options.Frame = false
	horizontalStack.Append(btn5)

	btn6 := ui.NewButton("btn6", ui.Auto, ui.Auto)
	btn6.Label = "Button 6"
	btn6.Options.Frame = false
	horizontalStack.Append(btn6)

	root.Append(horizontalStack)

	g.SetManager(root)

	err = root.SetKeybindings(g)
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
