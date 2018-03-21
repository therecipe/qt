package ctop

import (
	"fmt"
	"math/big"

	"github.com/therecipe/qt/core"

	"github.com/NebulousLabs/Sia/node/api"
	"github.com/NebulousLabs/Sia/types"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
)

type statusController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"balance"`
	_ string `property:"delta"`
}

func (c *statusController) init() {
	controller.Controller.ConnectWalletChanged(c.wallet)
}

func (c *statusController) wallet(wgI interface{}) {
	wg := wgI.(api.WalletGET)
	if wg.Unlocked {
		confirmedFloat, _ := new(big.Rat).SetFrac(wg.ConfirmedSiacoinBalance.Big(), types.SiacoinPrecision.Big()).Float64()
		c.SetBalance(fmt.Sprintf("%1.5f SC", confirmedFloat))

		unconfirmedBalance := wg.ConfirmedSiacoinBalance.Add(wg.UnconfirmedIncomingSiacoins).Sub(wg.UnconfirmedOutgoingSiacoins)
		if unconfirmedBalance.Cmp(wg.ConfirmedSiacoinBalance) >= 0 {
			deltaFloat, _ := new(big.Rat).SetFrac(unconfirmedBalance.Sub(wg.ConfirmedSiacoinBalance).Big(), types.SiacoinPrecision.Big()).Float64()
			c.SetDelta(fmt.Sprintf("+%1.5f SC", deltaFloat))
		} else {
			deltaFloat, _ := new(big.Rat).SetFrac(wg.ConfirmedSiacoinBalance.Sub(unconfirmedBalance).Big(), types.SiacoinPrecision.Big()).Float64()
			c.SetDelta(fmt.Sprintf("-%1.5f SC", deltaFloat))
		}
	} else {
		c.SetBalance("0")
		c.SetDelta("0")
	}
}
