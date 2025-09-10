package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

type pageMainType struct {
	demos       []Demo
	flex        *tview.Flex
	info        *tview.TextView
	pages       *tview.Pages
	curSlideNum int
}

var pageMain pageMainType

func (pageMain *pageMainType) build() {

	pageMain.curSlideNum = 0

	contArr := getContent(ItemID)

	pageMain.pages = tview.NewPages()
	pageMain.info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			if len(added) == 0 {
				return
			}
			pageMain.pages.SwitchToPage(added[0])
		})

	for index, cont := range contArr {
		primitive, pgs, slideContent := demo(cont)
		pageMain.demos = append(pageMain.demos, Demo{primitive, *pgs, *slideContent, 0, cont.id})
		pageMain.pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(pageMain.info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, cont.name)
	}

	pageMain.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pageMain.pages, 0, 1, true).
		AddItem(pageMain.info, 1, 1, false)

	pageMain.info.Highlight("0")

	application.pages.AddPage("main", pageMain.flex, true, true)
}

func nextSlide() {

	slide, _ := strconv.Atoi(pageMain.info.GetHighlights()[0])
	newPageNum := slide + 1
	if newPageNum < pageMain.pages.GetPageCount() {
		pageMain.info.Highlight(strconv.Itoa(newPageNum)).
			ScrollToHighlight()
		pageMain.curSlideNum = newPageNum
	} else {
		pageMain.info.Highlight(strconv.Itoa(0)).
			ScrollToHighlight()

		pageMain.curSlideNum = 0
	}
}

func previousSlide() {
	slide, _ := strconv.Atoi(pageMain.info.GetHighlights()[0])
	newPageNum := slide - 1
	if newPageNum >= 0 {
		pageMain.info.Highlight(strconv.Itoa(newPageNum)).
			ScrollToHighlight()
		pageMain.curSlideNum = newPageNum
	} else {
		pageMain.info.Highlight(strconv.Itoa(pageMain.pages.GetPageCount() - 1)).
			ScrollToHighlight()

		pageMain.curSlideNum = pageMain.pages.GetPageCount() - 1
	}
}
