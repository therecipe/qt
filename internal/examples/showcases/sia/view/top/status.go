package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { statusTemplate_QmlRegisterType2("TopTemplate", 1, 0, "StatusTemplate") }

type statusTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string `property:"balance"`
	_ string `property:"delta"`
}

func (t *statusTemplate) init() {
	c := ctop.NewStatusController(nil)

	t.ConnectBalance(c.Balance)
	c.ConnectBalanceChanged(t.BalanceChanged)

	t.ConnectDelta(c.Delta)
	c.ConnectDeltaChanged(t.DeltaChanged)
}
