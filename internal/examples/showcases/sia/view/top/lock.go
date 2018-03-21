package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { lockTemplate_QmlRegisterType2("TopTemplate", 1, 0, "LockTemplate") }

type lockTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ bool `property:"locked"`

	_ func() `signal:"change"`
}

func (t *lockTemplate) init() {
	c := ctop.NewLockController(nil)

	t.ConnectIsLocked(c.IsLocked)
	c.ConnectLockedChanged(t.LockedChanged)

	t.ConnectChange(c.Change)
}
