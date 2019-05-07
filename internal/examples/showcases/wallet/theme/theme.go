package theme

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/wallet/theme/controller"
)

func init() { themeTemplate_QmlRegisterType2("ThemeTemplate", 1, 0, "ThemeTemplate") }

type themeTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"accent,<-(this.c)"`
	_ string `property:"nextAccent,<-(this.c)"`

	_ string `property:"background,<-(this.c)"`
	_ string `property:"darkBackground,<-(this.c)"`

	_ string `property:"walletTableHeader,<-(this.c)"`
	_ string `property:"walletTableAlternate,<-(this.c)"`
	_ string `property:"walletTableHighlight,<-(this.c)"`

	_ string `property:"inputFieldBackground,<-(this.c)"`

	_ string `property:"font,<-(this.c)"`
	_ string `property:"fontHighlight,<-(this.c)"`

	c *controller.ThemeController
}

func (t *themeTemplate) init() {
	t.c = controller.NewThemeController(nil)
}
