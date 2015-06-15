package main

import (
	"github.com/nsf/termbox-go"
	"github.com/vijayee/termbox-menu"
	"github.com/vijayee/tourguide/tour"
)

func main() {
	tour.Init()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	topicItems := make([]menu.Item, len(tour.Topics))
	index := 0
	for _, key := range tour.IDs {
		c := tour.Topics[key]
		topicItems[index] = &c
		index++
	}

	mainMenu := menu.NewMenu("Tour of IPFS", topicItems, termbox.ColorWhite, termbox.ColorBlue)
	go menu.ListenToKeys()
	mainMenu.Invoke()
}
