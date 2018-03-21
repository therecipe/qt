package cleft

import (
	"fmt"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/NebulousLabs/Sia/types"

	"github.com/therecipe/qt/internal/examples/showcases/sia/controller"
)

type progressBarController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string  `property:"text"`
	_ float64 `property:"value"`

	_ func() `signal:"clicked"`
}

func (c *progressBarController) init() {

	controller.Controller.ConnectHeightChanged(c.heightChanged)

	c.ConnectClicked(c.clicked)
}

func (c *progressBarController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://explore.sia.tech", 0))
}

func (c *progressBarController) heightChanged(height types.BlockHeight) {
	if controller.Controller.IsSynced() {
		c.SetText(fmt.Sprintf("BH: %v", controller.Controller.Height()))
		c.SetValue(100)
	} else {
		estimatedHeight := estimatedHeightAt(time.Now())
		estimatedProgress := float64(height) / float64(estimatedHeight) * 100
		if estimatedProgress > 100 {
			estimatedProgress = 100
		}
		c.SetText(fmt.Sprintf("%.1f%%", estimatedProgress))
		c.SetValue(estimatedProgress)
	}
}

// estimatedHeightAt returns the estimated block height for the given time.
// Block height is estimated by calculating the minutes since a known block in
// the past and dividing by 10 minutes (the block time).
func estimatedHeightAt(t time.Time) types.BlockHeight {
	block100kTimestamp := time.Date(2017, time.April, 13, 23, 29, 49, 0, time.UTC)
	blockTime := float64(9) // overestimate block time for better UX
	diff := t.Sub(block100kTimestamp)
	estimatedHeight := 100e3 + (diff.Minutes() / blockTime)
	return types.BlockHeight(estimatedHeight + 0.5) // round to the nearest block
}
