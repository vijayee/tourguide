package tour

import (
	"errors"
	u "github.com/ipfs/go-ipfs/util"
	"github.com/nsf/termbox-go"
	"github.com/vijayee/tourguide/passage"
	"strconv"
	"strings"
)

var log = u.Logger("tour")

// ID is a string identifier for topics
type ID string

// LessThan returns whether this ID is sorted earlier than another.
func (i ID) LessThan(o ID) bool {
	return compareDottedInts(string(i), string(o))
}

// IDSlice implements the sort interface for ID slices.
type IDSlice []ID

func (a IDSlice) Len() int           { return len(a) }
func (a IDSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IDSlice) Less(i, j int) bool { return a[i].LessThan(a[j]) }

// Topic is a type of objects that structures a tour topic.
type Topic struct {
	ID ID
	Content
	hasPassed bool
}

type Content struct {
	Title  string
	Text   string
	verify func([]byte) (bool, error)
}

// Topics is a sorted list of topic IDs
var IDs []ID

// Topics contains a mapping of Tour Topic ID to Topic
var Topics = map[ID]Topic{}

var VerifcationInput []byte

// NextTopic returns the next in-order topic
func NextTopic(topic ID) ID {
	for _, id := range IDs {
		if topic.LessThan(id) {
			return id
		}
	}
	return topic // last one, it seems.
}
func (t *Topic) Mark() {

	t.hasPassed = true
}
func (t *Topic) Text() string {
	return t.Content.Text
}
func (t *Topic) Title() string {
	if t.hasPassed {
		return "\u2713" + string(t.ID) + "  " + t.Content.Title
	} else {
		return string(t.ID) + "  " + t.Content.Title
	}

}
func (t *Topic) Verify() (bool, error) {
	if t.Content.verify == nil {
		return false, errors.New("Content has no verification method")
	}
	return t.Content.verify(VerifcationInput)
}

func (t *Topic) Invoke() error {
	p := passage.NewPassage(t, termbox.ColorWhite, termbox.ColorBlue)
	p.Draw()
	p.ListenToKeys()
	return nil
}

// TopicID returns a valid tour topic ID from given string
func TopicID(t string) ID {
	if t == "" { // if empty, use first ID
		return IDs[0]
	}
	return ID(t)
}

func compareDottedInts(i, o string) bool {
	is := strings.Split(i, ".")
	os := strings.Split(o, ".")

	for n, vis := range is {
		if n >= len(os) {
			return false // os is smaller.
		}

		vos := os[n]
		ivis, err1 := strconv.Atoi(vis)
		ivos, err2 := strconv.Atoi(vos)
		if err1 != nil || err2 != nil {
			log.Debug(err1)
			log.Debug(err2)
			panic("tour ID LessThan: not an int")
		}

		if ivis < ivos {
			return true
		}

		if ivis > ivos {
			return false
		}
	}

	return len(os) > len(is)
}
