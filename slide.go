package main

type slide struct {
	id                int
	idItem            int
	num               int
	name              string
	content           string
	contentType       string
	direct            string
	contentProportion int
	pageProportion    int
	comment           string
	tabs              []Tab
	tabNum            int
}
