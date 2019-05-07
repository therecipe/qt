package controller

import (
	"fmt"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/controller"
)

type ProgressBarController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string  `property:"text"`
	_ float64 `property:"value"`

	_ func() `signal:"clicked,auto"`

	_ func(uint64) `signal:"heightChanged,<-(controller.Controller)"`
}

func (c *ProgressBarController) init() {
	if controller.DEMO {
		c.heightChanged(uint64(0))
	}
}

func (c *ProgressBarController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com/", 0))
}

func (c *ProgressBarController) heightChanged(height uint64) {
	if controller.Controller.IsSynced() {
		c.SetText(fmt.Sprintf("BH: %v", controller.Controller.Height()))
		c.SetValue(100)
	} else if controller.DEMO {
		c.SetText("BH: DEMO")
		c.SetValue(80)
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
func estimatedHeightAt(t time.Time) uint64 {
	block100kTimestamp := time.Date(2017, time.April, 13, 23, 29, 49, 0, time.UTC)
	blockTime := float64(9) // overestimate block time for better UX
	diff := t.Sub(block100kTimestamp)
	estimatedHeight := 100e3 + (diff.Minutes() / blockTime)
	return uint64(estimatedHeight + 0.5) // round to the nearest block
}
