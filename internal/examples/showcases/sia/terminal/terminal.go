package terminal

import (
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/showcases/sia/terminal/controller"
)

func init() { terminalTemplate_QmlRegisterType2("TerminalTemplate", 1, 0, "TerminalTemplate") }

type terminalTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(cmd string) string `slot:"command"`
}

func (t *terminalTemplate) init() {
	c := cterminal.NewTerminalController(nil)

	t.ConnectCommand(c.Command)
}
