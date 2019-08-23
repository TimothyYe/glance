package core

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var (
	p *widgets.Paragraph
)

func Init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	width, height := ui.TerminalDimensions()

	p = widgets.NewParagraph()
	p.Text = fmt.Sprintf("%d,%d", width, height)
	p.SetRect(0, 0, width, 3)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	p.Border = showBorder

	ui.Render(p)
}
