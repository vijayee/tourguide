package main

import (
	"github.com/nsf/termbox-go"
	"github.com/robfig/config"
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
	first := ""
	for _, key := range tour.IDs {
		c := tour.Topics[key]
		if first == "" {
			first = string(c.ID)
		}
		topicItems[index] = &c

		index++
	}

	c, err := config.ReadDefault("tour/config.cfg")
	if err != nil {
		c = config.NewDefault()
		c.AddSection("Topics")
		c.AddOption("Topics", "Current", first)
		c.WriteFile("tour/config.cfg", 0644, "Tour Guide Configuration")
	}
	mainMenu := menu.NewMenu("Tour of IPFS", topicItems, termbox.ColorWhite, termbox.ColorBlue)
	go menu.ListenToKeys()
	mainMenu.Invoke()
}
