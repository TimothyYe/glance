package core

import (
	ui "github.com/gizak/termui/v3"
)

var (
	showBorder   = false
	showHelp     = false
	showProgress = false
	bossKey      = false
)

func updateParagraph(key string) {
	p.Text = key
}

func setBorder() {
	showBorder = !showBorder
	p.Border = showBorder
}

func HandleEvents() {
	uiEvents := ui.PollEvents()
	defer ui.Close()

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
		case "j", "<C-n>":
			// show next content
			updateParagraph(e.ID)
			ui.Render(p)
		case "k", "<C-p":
			// show previous content
			updateParagraph(e.ID)
			ui.Render(p)
			return
		}
	}
}
