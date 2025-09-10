package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

var application applicationType

var app = tview.NewApplication()

var ItemID int

func main() {

	err := database.Connect()
	check(err)

	if len(os.Args) != 2 {
		fmt.Println("No item name sent")
		os.Exit(1)
	} else {
		count := countItems(os.Args[1])
		if count > 1 {
			getItemId(os.Args[1], false)
		} else {
			ItemID = getItemId(os.Args[1], true)
			application.init()
		}
	}
}
