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

func displayHelp(current string) {
	showHelp = !showHelp
	if showHelp {
		p.Text = "[↑]:j [↓]:k [b]:Boss Key [g]:Show/Hide Grid, [q]:Quit"
	} else {
		p.Text = current
	}
}

func displayBorder() {
	showBorder = !showBorder
	p.Border = showBorder
}

func displayProgress(current, progress string) {
	showProgress = !showProgress
	if showProgress {
		p.Text = progress
	} else {
		p.Text = current
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

// handleEvents handles all the keyboard events
func handleEvents() {
	uiEvents := ui.PollEvents()
	defer ui.Close()

	for {
		e := <-uiEvents
		switch e.ID {
		case "?":
			// show help menu
			displayHelp(r.Current())
		case "p":
			// show progress
			displayProgress(r.Current(), r.GetProgress())
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
			updateParagraph(r.Next())
		case "k", "<C-p>":
			// show previous content
			updateParagraph(r.Prev())
		}

		ui.Render(p)
	}
}
