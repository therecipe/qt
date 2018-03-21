package left

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { progressBarTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ProgressBarTemplate") }

type progressBarTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string  `property:"text"`
	_ float64 `property:"value"`

	_ func() `signal:"clicked"`
}

func (t *progressBarTemplate) init() {
	c := cleft.NewProgressBarController(nil)

	t.ConnectText(c.Text)
	c.ConnectTextChanged(t.TextChanged)

	t.ConnectValue(c.Value)
	c.ConnectValueChanged(t.ValueChanged)

	t.ConnectClicked(c.Clicked)
}
