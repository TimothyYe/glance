package core

import (
	"github.com/TimothyYe/glance/lib"
	ui "github.com/gizak/termui/v3"
)


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
		case "<C-n>":
			// show next content
			updateParagraph(r.Next())
		case "<C-p>":
			// show previous content
			updateParagraph(r.Prev())
		case "j":
			if rowNumber == "" {
				// show next content
				updateParagraph(r.Next())
			} else {
				// parse the row number
				if num, err := lib.ParseRowNum(rowNumber); err != nil {
					updateParagraph(err.Error())
				} else {
					updateParagraph(r.Goto(r.CurrentPos() + num))
				}
				rowNumber = ""
			}
		case "k":
			if rowNumber == "" {
				// show previous content
				updateParagraph(r.Prev())
			} else {
				// parse the row number
				if num, err := lib.ParseRowNum(rowNumber); err != nil {
					updateParagraph(err.Error())
				} else {
					updateParagraph(r.Goto(r.CurrentPos() - num))
				}
				rowNumber = ""
			}
		case "G":
			// jump to the specified row
			 if rowNumber == "" {
				 updateParagraph("Invalid row number!")
			 } else {
				// parse the row number
				if num, err := lib.ParseRowNum(rowNumber); err != nil {
					updateParagraph(err.Error())
				} else {
					updateParagraph(r.Goto(num))
				}
				rowNumber = ""
			}
		case "c":
			// change color
		case "0","1","2","3","4","5","6","7","8","9":
			// jump to rows
			rowNumber += e.ID
			updateParagraph(rowNumber)
		}

		ui.Render(p)
	}
}
