/*
A presentation of the tview package, implemented with tview.

# Navigation

The presentation will advance to the next slide when the primitive demonstrated
in the current slide is left (usually by hitting Enter or Escape). Additionally,
the following shortcuts can be used:

  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

var curPageNum int

// Starting point for the presentation.
func main() {
	file, err := os.OpenFile("app.log", os.O_TRUNC|os.O_CREATE, 0666)
	check(err)
	log.SetOutput(file)

	// The presentation slides.
	slides := []Slide{
		//Cover,
		//Introduction,
		//HelloWorld,
		//InputField,
		//Form,
		//TextView1,
		//TextView2,
		Table,
		//TreeView,
		//Flex,
		//Grid,
		//Colors,
		//End,
	}

	err = database.Connect()
	if err != nil {
		panic(err)
	}

	pages := tview.NewPages()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			if len(added) == 0 {
				return
			}
			curPageNum, err = strconv.Atoi(added[0])
			check(err)
			pages.SwitchToPage(added[0])
		})

	// Create the pages for all slides.
	previousSlide := func() {
		log.Println("previousSlide")
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		newPageNum := slide - 1
		log.Println("len(slides", len(slides))
		log.Println("newPageNum", newPageNum)
		if newPageNum >= 0 {
			log.Println("there")
			//slide = (slide + 1) % len(slides)
			//log.Println(slide)
			//curPageNum = newPageNum
			info.Highlight(strconv.Itoa(newPageNum)).
				ScrollToHighlight()
		} else {
			info.Highlight(strconv.Itoa(pages.GetPageCount() - 1)).
				ScrollToHighlight()
		}
	}
	nextSlide := func() {
		log.Println("nextSlide")
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		log.Println(slide)
		log.Println("info.GetHighlights()")
		log.Println(info.GetHighlights())
		newPageNum := slide + 1
		log.Println("len(slides", len(slides))
		log.Println("newPageNum", newPageNum)
		if newPageNum < pages.GetPageCount() {
			log.Println("there")
			//slide = (slide + 1) % len(slides)
			//log.Println(slide)
			//curPageNum = newPageNum
			info.Highlight(strconv.Itoa(newPageNum)).
				ScrollToHighlight()
		} else {
			info.Highlight(strconv.Itoa(0)).
				ScrollToHighlight()
		}
	}

	contArr := getContent(1)

	for index, cont := range contArr {
		primitive := demo(nextSlide, cont)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, cont.name)
	}
	//for index, slide := range slides {
	//	title, primitive := slide(nextSlide)
	//	pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
	//	fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	//}

	info.Highlight("0")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
