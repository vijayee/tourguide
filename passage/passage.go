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
}

func NewPassage(article Article, foreground termbox.Attribute, background termbox.Attribute) Passage {
	return Passage{article, foreground, background, nil, false, false, ""}

}
func (p *Passage) drawTitle() {
	w, _ := termbox.Size()
	title := p.article.Title()

	titleStart := (w / 2) - (len(title) / 2)
	titleRow := 2
	titleIndex := 0
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
	var text string

	if p.hasFailed {
		text = p.failMessages
	} else {
		text = fmt.Sprintf("%+q", p.article.Text())
	}

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
	for _, line := range lines {
		lineIndex := 0
		lineStart := 3
		lineEnd := w - lineStart
		if len(line) == 0 {
			continue
		}
		for lineIndex < len(line) {
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
