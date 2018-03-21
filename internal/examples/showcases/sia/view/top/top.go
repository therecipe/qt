package top

import "github.com/therecipe/qt/quick"

func init() { topTemplate_QmlRegisterType2("TopTemplate", 1, 0, "TopTemplate") }

type topTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`
}

func (t *topTemplate) init() {

}
