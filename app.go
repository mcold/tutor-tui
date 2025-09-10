package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"os"
	"strconv"
)

type applicationType struct {
	pages *tview.Pages
}

func (application *applicationType) init() {
	file, err := os.OpenFile("app.log", os.O_TRUNC|os.O_CREATE, 0666)
	check(err)
	log.SetOutput(file)

	app = tview.NewApplication()

	application.pages = tview.NewPages()
	pageMain.build()

	application.registerGlobalShortcuts()
	application.pages.SwitchToPage("main")

	if err := app.SetRoot(application.pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func (application *applicationType) registerGlobalShortcuts() {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			pageMain.demos[pageMain.curSlideNum].pages.SwitchToPage(strconv.Itoa(pageMain.demos[pageMain.curSlideNum].curTabNum))
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			pageMain.demos[pageMain.curSlideNum].pages.SwitchToPage(strconv.Itoa(pageMain.demos[pageMain.curSlideNum].curTabNum))
			return nil
		}

		if event.Key() == tcell.KeyRight && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(&pageMain.demos[pageMain.curSlideNum].pages)
		}
		if event.Key() == tcell.KeyLeft && event.Modifiers() == tcell.ModAlt {
			app.SetFocus(&pageMain.demos[pageMain.curSlideNum].slideContent)
		}

		return event
	})
}
