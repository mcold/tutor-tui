package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

const tableData = `sale_dt|product_id|status|amount|pcs|_part
2025-06-01|1|successed|100.00|1|all_1_2_2
2025-06-01|2|successed|75.00|3|all_1_2_2`

func demo(nextSlide func(), cont content) (frame tview.Primitive) {

	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

	_, err := fmt.Fprint(code, cont.code)
	check(err)

	// TODO: change to Table
	//tbl := tview.NewTextView().
	//	SetWrap(false).
	//	SetDynamicColors(true)

	//tbl := tview.NewTable().
	//	SetFixed(1, 1)
	//tbl.SetBorderPadding(1, 1, 2, 0)
	//
	//fmt.Fprint(tbl, cont.output)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(code, 0, cont.codeProportion, false)

	if cont.direct == "row" {
		flex.SetDirection(tview.FlexRow)
	} else {
		flex.SetDirection(tview.FlexColumn)
	}

	if cont.outType == "table" {
		table := tview.NewTable().
			SetFixed(1, 1).
			SetSelectable(true, true).
			SetSeparator(' ').
			SetBorders(false)

		// TODO: clear separate ----- if it exists in output
		for row, line := range strings.Split(cont.output, "\n") {
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
				} else if column == 0 || column >= 3 { // Align right for first column and numeric columns
					align = tview.AlignRight
				}

				tableCell := tview.NewTableCell(cell).
					SetTextColor(color).
					SetAlign(align).
					SetSelectable(true)

				// Set expansion for middle columns if needed
				if column >= 1 && column <= 2 {
					tableCell.SetExpansion(1)
				}

				table.SetCell(row, column, tableCell)
			}

		}
		flex.AddItem(table, 0, cont.outProportion, false)
	} else {
		out := tview.NewTextView().
			SetWrap(false).
			SetDynamicColors(true)
		out.SetBorderPadding(1, 1, 2, 0)

		_, err := fmt.Fprint(out, cont.output)
		check(err)

		flex.AddItem(out, 0, cont.outProportion, false)
	}

	return flex
}
