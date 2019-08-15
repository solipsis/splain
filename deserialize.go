package main

import (
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var deserializeTop ui.Drawable
var deserializeBottom ui.Drawable

var deserializeBuf []byte
var activeTab int

func deserializeInit(xpub string) {
	deserializeBuf = base58.Decode(xpub)

	top := widgets.NewTabPane(strings.Repeat(" ", 70)+"version", "depth", "fingerprint", "index", "chaincode", "key", "checksum")
	top.SetRect(5, 0, 20, 3)
	top.Border = true
	top.ActiveTabStyle = ui.NewStyle(ui.ColorYellow, ui.ColorBlack, ui.ModifierBold)
	//top = wrapper{tabpane}
	//top = wrapper{decodeTop}
	deserializeTop = top

	bottom := widgets.NewParagraph()
	bottom.Text = printXpub()
	deserializeBottom = bottom

}

func redrawContent() {
	par := (deserializeBottom).(*widgets.Paragraph)
	//par.Text = printXpub()

	var b strings.Builder
	pre := strings.Repeat(" ", 60)

	b.WriteString("\n\n")
	if activeTab == 0 {
		b.WriteString(pre + "The [version](fg:green,mod:bold) gives information into what kind of key is encoded\n")
		b.WriteString(pre + "This is also what gives an XPUB its distinct form (XPUB, LTUB, ZPUB)\n")
	}
	if activeTab == 1 {
		b.WriteString(pre + "The [Depth](fg:red,mod:bold) byte tells you have what generation key this is\n")
		b.WriteString(pre + "In other words it tells you how many parent keys or ancestors lead up to this key\n")
	}
	if activeTab == 2 {
		b.WriteString(pre + "The [Fingerprint](fg:yellow,mod:bold) is used to verify the parent key\n")
	}
	if activeTab == 3 {
		b.WriteString(pre + "The [Index](fg:cyan,mod:bold) tells you what child of the parent key this is\n")
		b.WriteString(pre + "Each parent can support up to 2^32 child keys")
	}
	if activeTab == 4 {
		b.WriteString(pre + "The [Chaincode](fg:blue,mod:bold) is used to deterministically derive child keys of this key\n")
	}
	if activeTab == 5 {
		b.WriteString(pre + "The [Keydata](fg:magenta,mod:bold) is the actual bytes of this extended key\n")
		b.WriteString(pre + "If the first byte is 0x00 you know that this is a public child key\n")
		b.WriteString(pre + "Otherwise this is a private child\n")
	}
	if activeTab == 6 {
		b.WriteString(pre + "The [Checksum](fg:white,mod:bold) is used to verify that the other data was encoded and transmitted properly\n")
	}

	par.Text = printXpub() + "\n" + b.String()
}

func printXpub() string {
	idx := activeTab
	var b strings.Builder
	buf := deserializeBuf
	bold := ",bg:black,mod:bold"

	// HOT garbage hackathon life
	var version, depth, fingerprint, index, chaincode, keydata, checksum string
	if idx == 0 {
		version = "                          [" + hex.EncodeToString(buf[0:4]) + "](fg:green" + bold + ") " // version (4)
	} else {
		version = "                          [" + hex.EncodeToString(buf[0:4]) + "](fg:green) " // version (4)
	}

	if idx == 1 {
		depth = "[" + hex.EncodeToString(buf[4:5]) + "](fg:red" + bold + ") "
	} else {
		depth = "[" + hex.EncodeToString(buf[4:5]) + "](fg:red) "
	}

	if idx == 2 {
		fingerprint = "[" + hex.EncodeToString(buf[5:9]) + "](fg:yellow" + bold + ") "
	} else {
		fingerprint = "[" + hex.EncodeToString(buf[5:9]) + "](fg:yellow) "
	}

	if idx == 3 {
		index = "[" + hex.EncodeToString(buf[9:13]) + "](fg:cyan" + bold + ") "
	} else {
		index = "[" + hex.EncodeToString(buf[9:13]) + "](fg:cyan) "
	}

	if idx == 4 {
		chaincode = "[" + hex.EncodeToString(buf[13:45]) + "](fg:blue" + bold + ") "
	} else {
		chaincode = "[" + hex.EncodeToString(buf[13:45]) + "](fg:blue) "
	}

	if idx == 5 {
		keydata = "[" + hex.EncodeToString(buf[45:78]) + "](fg:magenta" + bold + ") "
	} else {
		keydata = "[" + hex.EncodeToString(buf[45:78]) + "](fg:magenta) "
	}

	if idx == 6 {
		checksum = "[" + hex.EncodeToString(buf[78:82]) + "](fg:white" + bold + ") "

	} else {
		checksum = "[" + hex.EncodeToString(buf[78:82]) + "](fg:white) "
	}

	//b.WriteString("                          [" + hex.EncodeToString(buf[0:4]) + "](fg:green" + bold + ") ") // version (4)
	//b.WriteString("[" + hex.EncodeToString(buf[4:5]) + "](fg:red" + bold + ") ")       // depth (1)
	//b.WriteString("[" + hex.EncodeToString(buf[5:9]) + "](fg:yellow" + bold + ") ")    // fingerprint (4)
	//b.WriteString("[" + hex.EncodeToString(buf[9:13]) + "](fg:cyan" + bold + ") ")     // Index (4)
	//b.WriteString("[" + hex.EncodeToString(buf[13:45]) + "](fg:blue" + bold + ") ")    // Chaincode (32)
	//b.WriteString("[" + hex.EncodeToString(buf[45:78]) + "](fg:magenta" + bold + ") ") // Keydata (33)
	//b.WriteString("[" + hex.EncodeToString(buf[78:82]) + "](fg:white" + bold + ") ") // checksum (4)
	b.WriteString(version)
	b.WriteString(depth)
	b.WriteString(fingerprint)
	b.WriteString(index)
	b.WriteString(chaincode)
	b.WriteString(keydata)
	b.WriteString(checksum)

	return b.String()
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
