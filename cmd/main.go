package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var (
	border = false
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	width, height := ui.TerminalDimensions()

	p := widgets.NewParagraph()
	// p.Title = "Text Box"
	p.Text = fmt.Sprintf("%d,%d", width, height)
	p.SetRect(0, 0, width, 3)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.Border = border
	// p.WrapText = true

	updateParagraph := func(key string) {
		p.Text = key
	}

	setBorder := func() {
		border = !border
		p.Border = border
	}

	ui.Render(p)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "?":
			// show help menu
			updateParagraph("[↑]:j [↓]:k [b]:Boss Key [g]:Show/Hide Grid, [q]:Quit")
			ui.Render(p)
		case "p":
			// show progress
		case "f":
			// show border
			setBorder()
			ui.Render(p)
		case "b":
			// boss key
			p.Border = false
			p.Text = "[root@localhost]$"
			ui.Render(p)
		case "q", "<C-c>":
			// quit
			return
		case "j", "k", "<C-n>", "<C-p>":
			// switch content
			updateParagraph(e.ID)
			ui.Render(p)
		}
	}
}
