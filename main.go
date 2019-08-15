package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {

	xpub := "xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz"
	fmt.Println(xpub)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	/*
		header := widgets.NewParagraph()
		header.Text = "Press q to quit, Press h or l to switch tabs"
		header.SetRect(0, 0, 50, 1)
		header.Border = false
		header.TextStyle.Bg = ui.ColorBlue
	*/

	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress h or l to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 35, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	tabpane := widgets.NewTabPane("pierwszy", "[drugi](fg:green)", "trzeci", "żółw", "four", "five")
	tabpane.SetRect(0, 3, 50, 7)
	tabpane.Border = true

	xpubPar := widgets.NewParagraph()
	xpubPar.Text = "[xpub](fg:green)[123ABC](fg:yellow)[fbeef](fg:red)[dave](mod:bold,fg:cyan,bg:white)"
	xpubPar.SetRect(0, 0, 30, 3)

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
		case 1:
			ui.Render(bc)
		}
	}

	// colored paragraph up top
	// each tab is associated with 1 chunk

	ui.Render(xpubPar, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(xpubPar, tabpane)
			//ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(xpubPar, tabpane)
			//ui.Render(header, tabpane)
			renderTab()
		}
	}
}
