package tour

import (
	"github.com/robfig/config"
	"os"
	"testing"
)

func TestCurrentTopic(t *testing.T) {
	Init()
	var currentTest Topic
	for _, key := range IDs {
		c := Topics[key]

		c.Content.Verify = func(t *testing.T) {
			t.Error("This should always fail")
		}
		Topics[key] = c
	}
	c, err := config.ReadDefault("config.cfg")
	if err != nil {
		os.Exit(1)
	} else {
		id, err2 := c.String("Topics", "Current")
		if err2 != nil {
			os.Exit(1)
		}
		currentTest = Topics[ID(id)]
		currentTest.Content.Verify(t)
	}

}
