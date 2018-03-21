package top

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { searchTemplate_QmlRegisterType2("TopTemplate", 1, 0, "SearchTemplate") }

type searchTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(string) `signal:"search"`
}

func (t *searchTemplate) init() {
	c := ctop.NewSearchController(nil)

	t.ConnectSearch(c.Search)
}
