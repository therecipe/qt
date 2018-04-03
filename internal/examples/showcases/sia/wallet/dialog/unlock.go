package dialog

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { unlockTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "UnlockTemplate") }

type unlockTemplate struct {
	dialogTemplate

	_ func() `constructor:"init"`

	_ func(string) *core.QVariant `slot:"unlock"`
}

func (t *unlockTemplate) init() {
	t.ConnectUnlock(controller.Controller.Unlock)
}
