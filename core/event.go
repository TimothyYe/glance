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

func displayHelp() {
	showHelp = !showHelp
	if showHelp {
		p.Text = "[↑]:j [↓]:k [b]:Boss Key [g]:Show/Hide Grid, [q]:Quit"
	} else {
		p.Text = ""
	}
}

func displayBorder() {
	showBorder = !showBorder
	p.Border = showBorder
}

func displayProgress() {
	showProgress = !showProgress
	if showProgress {
		p.Text = "(10/99)"
	} else {
		p.Text = ""
	}
}

func displayBossKey() {
	bossKey = !bossKey
	if bossKey {
		p.Border = false
		p.Text = "[root@localhost]$"
	} else {
		p.Text = ""
	}
}

// HandleEvents handles all the keyboard events
func HandleEvents() {
	uiEvents := ui.PollEvents()
	defer ui.Close()

	for {
		e := <-uiEvents
		switch e.ID {
		case "?":
			// show help menu
			displayHelp()
		case "p":
			// show progress
			displayProgress()
		case "f":
			// show border
			displayBorder()
		case "b":
			// boss key
			displayBossKey()
		case "q", "<C-c>":
			// quit
			return
		case "j", "<C-n>":
			// show next content
			updateParagraph(e.ID)
		case "k", "<C-p>":
			// show previous content
			updateParagraph(e.ID)
		}

		ui.Render(p)
	}
}
