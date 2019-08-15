package main

import (
	"fmt"
	"log"

	"github.com/gizak/termui/v3"
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

	type token struct {
		color termui.Color
		text  string
	}

	/*
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
	*/

	decodeInit(xpub)

	var top wrapper
	tabpane := widgets.NewTabPane("version", "depth", "fingerprint", "index", "chaincode", "key", "private")
	tabpane.SetRect(5, 0, 20, 3)
	tabpane.Border = true
	top = wrapper{tabpane}
	top = wrapper{decodeTop}

	var bottom wrapper
	bottom = wrapper{decodeBottom}

	//var test2 wrapper
	tabpane2 := widgets.NewTabPane("davel", "tim", "mom", "maya", "chaincode", "key", "private")
	tabpane2.SetRect(5, 0, 20, 3)
	tabpane2.Border = true
	//test2 = tabpane2

	xpubPar := widgets.NewParagraph()
	xpubPar.Text = "[xpub](fg:green)[123ABC](fg:yellow)[fbeef](fg:red)[dave](mod:bold,fg:cyan,bg:white)"
	xpubPar.SetRect(0, 0, 30, 3)

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = []string{
		"[0] Decode",
		"[1] Deserialize",
	}
	l.SelectedRowStyle = ui.NewStyle(ui.ColorYellow, ui.ColorBlack, ui.ModifierBold)
	//l.TextStyle = ui.NewStyle(ui.ColorGrey)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	grid := ui.NewGrid()
	termwidth, termheight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termwidth, termheight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(.4/2, l),
			ui.NewCol(1.6/2,
				ui.NewRow(.3/2, &top),
				ui.NewRow(1.7/2, &bottom),
			),
		),
	)
	ui.Render(grid)

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			//ui.Render(p2)
			xpubPar.Text = "Potato"
		case 1:
			xpubPar.Text = "[RED DRAGON](fg:red) Potato"
			//ui.Render(bc)
		case 2:
			//test.Drawable = decodeTop
			//xpubPar.Text = fmt.Sprintf("%+v", test)
		default:
			xpubPar.Text = "Javascript is cool"
		}
	}

	//renderRHS := func() {
	//switch l.Ac
	//}

	// colored paragraph up top
	// each tab is associated with 1 chunk

	//ui.Render(xpubPar, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			renderTab()
			ui.Render(grid)
			//ui.Render(xpubPar, tabpane)
			//ui.Render(header, tabpane)
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			renderTab()
			ui.Render(grid)
			//ui.Render(xpubPar, tabpane)
			//ui.Render(header, tabpane)
		case "j", "<Down>":
			l.ScrollDown()
			ui.Render(l)
		case "k", "<Up>":
			l.ScrollUp()
			ui.Render(l)

		}

	}
}

type wrapper struct {
	ui.Drawable
}

/*
type wrapper struct {
	drawable ui.Drawable
}

func (w *wrapper) GetRect() image.Rectangle {
	return w.drawable.GetRect()
}

func (w *wrapper) SetRect(a, b, c, d int) {
	w.drawable.SetRect(a, b, c, d)
}

func (w *wrapper) Draw(b *ui.Buffer) {
	w.drawable.Draw(b)
}

func (w *wrapper) Lock() {
	w.drawable.Lock()
}

func (w *wrapper) Unlock() {
	w.drawable.Unlock()
}
*/
