package controller

import (
	"time"

	"github.com/therecipe/qt/core"

	"github.com/NebulousLabs/Sia/node/api/client"
	"github.com/NebulousLabs/Sia/types"
)

var Client = client.New("127.0.0.1:9980")

var Controller *controller

type controller struct {
	core.QObject

	_ func() `constructor:"init"`

	_ bool              `property:"synced"`
	_ types.BlockHeight `property:"height"`

	_ bool        `property:"locked"`
	_ bool        `property:"encrypted"`
	_ interface{} `property:"wallet"`
}

func (c *controller) init() {
	Controller = c

	c.SetLocked(true)

	go c.loop()
}

func (c *controller) loop() {
	for range time.NewTicker(1 * time.Second).C {

		cg, errC := Client.ConsensusGet()
		if errC != nil {
			println(errC.Error())
		} else {
			c.SetSynced(cg.Synced)
			c.SetHeight(cg.Height)
		}

		wg, errW := Client.WalletGet()
		if errW != nil {
			println(errW.Error())
		} else {
			c.SetLocked(!wg.Unlocked)
			c.SetEncrypted(wg.Encrypted)
			c.SetWallet(wg)
		}
	}
}
