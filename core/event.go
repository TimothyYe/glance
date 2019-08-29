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
		p.Text = menuText
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

func displayBossKey(current string) {
	bossKey = !bossKey
	if bossKey {
		p.Border = false
		p.Text = fakeShell
	} else {
		p.Text = current
	}
}

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
			displayBossKey(r.Current())
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
