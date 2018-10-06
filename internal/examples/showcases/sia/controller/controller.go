package controller

import (
	"time"

	"github.com/therecipe/qt/core"
)

var DEMO bool

var Controller *controller

type controller struct {
	core.QObject

	_ func() `constructor:"init"`

	_ bool   `property:"synced"`
	_ uint64 `property:"height"`

	_ bool        `property:"locked"`
	_ bool        `property:"encrypted"`
	_ interface{} `property:"wallet"`
}

func (c *controller) init() {
	Controller = c

	c.SetSynced(false)
	c.SetLocked(!DEMO)

	go c.loop()
}

func (c *controller) loop() {
	for range time.NewTicker(1 * time.Second).C {

	}
}
