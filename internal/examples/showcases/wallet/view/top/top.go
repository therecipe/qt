package top

import "github.com/bluszcz/cutego/quick"

func init() { topTemplate_QmlRegisterType2("TopTemplate", 1, 0, "TopTemplate") }

type topTemplate struct {
	quick.QQuickItem
}
