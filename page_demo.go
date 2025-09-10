package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Demo struct {
	frame        tview.Primitive
	pages        tview.Pages
	slideContent tview.TextView
	curTabNum    int
	idSlide      int
}

func demo(sl slide) (frame tview.Primitive, pages *tview.Pages, slideContent *tview.TextView) {

	slideContent = tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	slideContent.SetBorderPadding(1, 1, 2, 0)

	_, err := fmt.Fprint(slideContent, sl.content)
	check(err)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(slideContent, 0, sl.contentProportion, false)

	flex.SetBackgroundColor(tcell.ColorWhite)

	if sl.direct == "row" {
		flex.SetDirection(tview.FlexRow)
	} else {
		flex.SetDirection(tview.FlexColumn)
	}

	pages = tview.NewPages()
	for index, tabSlide := range sl.tabs {
		primitive := getTabPage(tabSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
	}
	if pages.GetPageCount() > 0 {
		flex.AddItem(pages, 0, sl.pageProportion, true)
	}

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'n' && event.Modifiers() == tcell.ModAlt {
			newNum := sl.tabNum + 1
			if newNum < pages.GetPageCount() {
				sl.tabNum = newNum
				pageMain.demos[pageMain.curSlideNum].curTabNum = sl.tabNum
			}
			pages.SwitchToPage(strconv.Itoa(sl.tabNum))
			return nil
		} else if event.Rune() == 'p' && event.Modifiers() == tcell.ModAlt {
			newNum := sl.tabNum - 1
			if newNum >= 0 {
				sl.tabNum = newNum
				pageMain.demos[pageMain.curSlideNum].curTabNum = sl.tabNum
			}
			pages.SwitchToPage(strconv.Itoa(sl.tabNum))
			return nil
		}
		return event

	})

	return flex, pages, slideContent
}

func getTabPage(tabSlide Tab) (frame tview.Primitive) {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBackgroundColor(tcell.ColorWhite)

	if tabSlide.contentType == "table" {
		table := tview.NewTable().
			SetFixed(1, 1).
			SetSelectable(true, true).
			SetSeparator(' ').
			SetBorders(false)

		for row, line := range strings.Split(tabSlide.content, "\n") {
			for column, cell := range strings.Split(line, "|") {
				color := tcell.ColorWhite
				if row == 0 {
					color = tcell.ColorYellow
				} else if column == 0 {
					color = tcell.ColorDarkCyan
				}

				align := tview.AlignLeft
				if row == 0 {
					align = tview.AlignCenter
				} else if column == 0 || column >= 3 {
					align = tview.AlignRight
				}

				tableCell := tview.NewTableCell(cell).
					SetTextColor(color).
					SetAlign(align).
					SetSelectable(true)

				if column >= 1 && column <= 2 {
					tableCell.SetExpansion(1)
				}

				table.SetCell(row, column, tableCell)
			}

		}

		flex.AddItem(table, 0, 5, false)

		flex.SetFocusFunc(func() {
			app.SetFocus(table)
		})
	} else {
		out := tview.NewTextView().
			SetWrap(false).
			SetDynamicColors(true)
		out.SetBorderPadding(1, 1, 2, 0)

		_, err := fmt.Fprint(out, "---- ["+strconv.Itoa(tabSlide.num)+"]\n\n"+tabSlide.content)
		check(err)
		
		flex.AddItem(out, 0, 5, false)

		flex.SetFocusFunc(func() {
			app.SetFocus(out)
		})
	}

	if len(tabSlide.comment) > 0 {
		com := tview.NewTextView().
			SetWrap(false).
			SetDynamicColors(true)

		com.SetText(strings.TrimSpace(tabSlide.comment))
		flex.AddItem(com, 0, 1, false)
	}

	if len(tabSlide.name) > 0 {
		flex.SetTitle(tabSlide.name).SetTitleAlign(tview.AlignLeft)
	}

	flex.SetTitle(strconv.Itoa(tabSlide.num))
	flex.SetBorder(true)
	flex.SetBorderColor(tcell.ColorWhite)

	return flex
}
