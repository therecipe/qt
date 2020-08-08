package terminal

import (
	"github.com/StarAurryon/qt/quick"

	_ "github.com/StarAurryon/qt/internal/examples/showcases/wallet/terminal/controller"
)

func init() { terminalTemplate_QmlRegisterType2("TerminalTemplate", 1, 0, "TerminalTemplate") }

type terminalTemplate struct {
	quick.QQuickItem

	_ func(cmd string) string `slot:"command,->(controller.NewTerminalController(nil))"`
}
