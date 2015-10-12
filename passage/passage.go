package passage

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/vijayee/termbox-menu"
	"strings"
	"unicode/utf8"
)

type Article interface {
	Title() string
	Text() string
	Mark()
	Verify() (bool, error)
}

type Passage struct {
	article         Article
	foreground      termbox.Attribute
	background      termbox.Attribute
	keyEventService chan termbox.Event
	isFocused       bool
	hasFailed       bool
	failMessages    string
	displayStart    int
	displayEnd      int
	lines           []string
}

func NewPassage(article Article, foreground termbox.Attribute, background termbox.Attribute) Passage {
	text := fmt.Sprintf("%+q", article.Text())
	text = strings.Replace(text, "\\t", "  ", -1)
	if text != "\"\"" && len(text) > 1 {
		//trim weird quotes
		if text[0] == '\u0022' {
			text = text[1 : len(text)-1]
		}
		if text[len(text)-1] == '\u0022' {
			text = text[0 : len(text)-2]
		}
	} else {
		text = ""
	}
	lines := strings.Split(text, "\\n")

	return Passage{article, foreground, background, nil, false, false, "", 0, 0, lines}

}
func (p *Passage) drawTitle() {
	w, _ := termbox.Size()
	title := p.article.Title()

	titleStart := (w / 2) - (len(title) / 2)
	titleRow := 2
	titleIndex := 0
	instruction1 := "ESC: Back"
	instruction2 := "Enter: Verify"
	spaces := w - (len(instruction1) + len(instruction2))
	space := ""
	for s := 0; s < spaces; s++ {
		space += " "
	}
	instructions := ""
	instructions += instruction1
	instructions += space
	instructions += instruction2
	for i, c := range instructions {
		termbox.SetCell(i, 0, c, p.foreground, p.background)
	}
	for x := 0; x < w; x++ {
		if x >= titleStart && titleIndex < len(title) {
			c, rw := utf8.DecodeRuneInString(title[titleIndex:])
			titleIndex += rw
			titleStart += rw
			termbox.SetCell(x, titleRow, c, p.foreground, p.background)
		}
		termbox.SetCell(x, titleRow+1, '_', p.foreground, p.background)
	}

}
func (p *Passage) drawContent() {
	w, h := termbox.Size()
	currrentRow := 5
	var lines []string
	if p.hasFailed {
		text := p.failMessages
		text = strings.Replace(text, "\\t", "  ", -1)
		if text != "\"\"" && len(text) > 1 {
			//trim weird quotes
			if text[0] == '\u0022' {
				text = text[1 : len(text)-1]
			}
			if text[len(text)-1] == '\u0022' {
				text = text[0 : len(text)-2]
			}
		} else {
			text = ""
		}
		lines = strings.Split(text, "\\n")
	} else {
		lines = p.lines
	}

	if p.displayEnd == 0 && len(lines) != 0 {
		p.displayEnd = len(lines)
	}
	for _, line := range lines[p.displayStart:] {
		lineIndex := 0
		lineStart := 3
		lineEnd := w - lineStart
		if len(line) == 0 {
			continue
		}
		for lineIndex < p.displayEnd {
			if currrentRow > h {
				break
			}
			for x := 0; x < w; x++ {
				var c rune
				var rw int
				if x > lineStart && x < lineEnd && lineIndex < len(line) {
					c, rw = utf8.DecodeRuneInString(line[lineIndex:])
					lineIndex += rw
					lineStart += rw
				} else {
					c = ' '
				}
				termbox.SetCell(x, currrentRow, c, p.foreground, p.background)

			}
			currrentRow++
			lineStart = 3
		}
		currrentRow++
	}
	if (p.displayEnd - p.displayStart) > (h - 3) {
		a, _ := utf8.DecodeRuneInString("\u25BC")
		termbox.SetCell(w-1, h-1, a, p.background, p.foreground)

	}
	if p.displayStart > 0 {
		a, _ := utf8.DecodeRuneInString("\u25B2")
		termbox.SetCell(w-1, 4, a, p.background, p.foreground)

	}

}
func (p *Passage) Draw() error {
	termbox.Clear(p.background, p.background)
	p.drawTitle()
	p.drawContent()
	termbox.Flush()
	return nil
}

func (p *Passage) verify() {
	result, err := p.article.Verify()
	if result {
		p.hasFailed = false
		p.article.Mark()
	} else {
		p.hasFailed = true
		p.failMessages = err.Error()
	}
}

func (p *Passage) Up() {
	if p.displayStart != 0 {
		p.displayStart--
	}
}

func (p *Passage) Down() {
	_, h := termbox.Size()
	if (p.displayEnd - p.displayStart) > (h - 3) {
		p.displayStart++
	}
}

func (p *Passage) ListenToKeys() {
	p.keyEventService = make(chan termbox.Event)
	menu.Subscribe(p.keyEventService)
	p.isFocused = true
	for {
		select {
		case keyEvent := <-p.keyEventService:
			switch keyEvent.Type {
			case termbox.EventKey:
				switch keyEvent.Key {
				case termbox.KeyEsc:
					if p.isFocused == true {
						return
					}
				case termbox.KeyArrowUp:
					if p.isFocused == true {
						go func() {
							p.Up()
							p.Draw()
						}()
					}
				case termbox.KeyArrowDown:
					if p.isFocused == true {
						go func() {
							p.Down()
							p.Draw()
						}()
					}
				case termbox.KeyEnter:
					if p.isFocused == true {
						go func() {
							p.verify()
							p.Draw()
						}()
					}
				}
			case termbox.EventError:
				panic(keyEvent.Err)
			}
		}

		if p.isFocused == true {
			p.Draw()
		}
	}

}
