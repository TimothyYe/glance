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
)

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
