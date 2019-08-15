package main

import (
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var deserializeTop ui.Drawable
var deserializeBottom ui.Drawable

func deserializeInit(xpub string) {

	top := widgets.NewTabPane(strings.Repeat(" ", 30)+"version", "depth", "fingerprint", "index", "chaincode", "key", "private")
	top.SetRect(5, 0, 20, 3)
	top.Border = true
	//top = wrapper{tabpane}
	//top = wrapper{decodeTop}
	deserializeTop = top

	bottom := widgets.NewParagraph()
	bottom.Text = "Placeholder"
	deserializeBottom = bottom
}

//w, _ := ui.TerminalDimensions()
//pad := strings.Repeat(" ", w/3)

/*
	top := widgets.NewParagraph()
	top.Text = "[" + center("Base 58 Decode") + "](fg:yellow,mod:bold)"
	top.SetRect(3, 3, 30, 6)
	decodeTop = top

	bottom := widgets.NewParagraph()
	var b strings.Builder
	b.WriteString("[" + center("XPUBs are encoded with the base58 encoding") + "](fg:green,mod:bold)\n")
	b.WriteString(center("\n"))
	b.WriteString(center(xpub + "\n"))
	b.WriteString(center("\n"))
	b.WriteString(center("|\n"))
	b.WriteString(center("|\n"))
	b.WriteString(center("|\n"))
	b.WriteString(center("v\n"))
	b.WriteString(center("\n"))
	b.WriteString(center("0x" + hex.EncodeToString(decodeXPUB(xpub))))
	bottom.Text = b.String()
	decodeBottom = bottom
}

func decodeXPUB(xpub string) []byte {
	return base58.Decode(xpub)
}

func center(text string) string {
	w, _ := ui.TerminalDimensions()

	//lenWithoutFormatting := 0
	//temp := strings.Replace(text, "[", "", -1)
	//temp := strings.Replace(text, "[", "", -1)

	buf := int(0.7*float64(w)) - len(text)
	return strings.Repeat(" ", buf/2) + text
}
*/
