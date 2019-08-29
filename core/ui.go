package core

import (
	"fmt"
	"log"

	"github.com/TimothyYe/glance/reader"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var (
	p *widgets.Paragraph
	r reader.GeneralReader

	showBorder   = false
	showHelp     = false
	showProgress = false
	bossKey      = false
	rowNumber = ""
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

// Init ui & components
func Init(gr reader.GeneralReader) {
	r = gr

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	//_, height := ui.TerminalDimensions()
	width := 83

	p = widgets.NewParagraph()
	p.Text = fmt.Sprintf(r.Current())
	p.SetRect(0, 0, width, 3)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.Border = showBorder

	ui.Render(p)
	handleEvents()
}
