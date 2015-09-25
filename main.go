package main

import (
	"code.google.com/p/go.crypto/ssh/terminal"
	//"fmt"
	"github.com/nsf/termbox-go"
	"github.com/robfig/config"
	"github.com/vijayee/termbox-menu"
	"github.com/vijayee/tourguide/tour"
	"io/ioutil"
	"os"
)

func main() {
	if !terminal.IsTerminal(0) {
		tour.VerifcationInput, _ = ioutil.ReadAll(os.Stdin)
	}
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
	id := ""
	currentIndex := 0
	cfg, err := config.ReadDefault("tour/config.cfg")

	if err == nil {
		id, err = cfg.String("Topics", "Current")
	}

	for _, key := range tour.IDs {
		c := tour.Topics[key]
		if err != nil && first == "" {
			first = string(c.ID)
		}
		if err == nil {
			if id == string(c.ID) {
				currentIndex = index
			}

		}
		topicItems[index] = &c

		index++
	}
	if err != nil {
		cfg = config.NewDefault()
		cfg.AddSection("Topics")
		cfg.AddOption("Topics", "Current", first)
		cfg.WriteFile("tour/config.cfg", 0644, "Tour Guide Configuration")
	}
	if currentIndex == 1 {
	}
	mainMenu := menu.NewMenu("Tour of IPFS", topicItems, termbox.ColorWhite, termbox.ColorBlue)
	go menu.ListenToKeys()
	mainMenu.Invoke()

}
