package theme

import (
	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/internal/examples/showcases/sia/theme/controller"
)

func init() { themeTemplate_QmlRegisterType2("ThemeTemplate", 1, 0, "ThemeTemplate") }

type themeTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"accent"`
	_ string `property:"nextAccent"`

	_ string `property:"background"`
	_ string `property:"darkBackground"`

	_ string `property:"walletTableHeader"`
	_ string `property:"walletTableAlternate"`
	_ string `property:"walletTableHighlight"`

	_ string `property:"inputFieldBackground"`

	_ string `property:"font"`
	_ string `property:"fontHighlight"`
}

func (t *themeTemplate) init() {
	c := controller.NewThemeController(nil)

	t.ConnectAccent(c.Accent)
	c.ConnectAccentChanged(t.AccentChanged)

	t.ConnectNextAccent(c.NextAccent)
	c.ConnectNextAccentChanged(t.NextAccentChanged)

	t.ConnectBackground(c.Background)
	c.ConnectBackgroundChanged(t.BackgroundChanged)

	t.ConnectDarkBackground(c.DarkBackground)
	c.ConnectDarkBackgroundChanged(t.DarkBackgroundChanged)

	t.ConnectWalletTableHeader(c.WalletTableHeader)
	c.ConnectWalletTableHeaderChanged(t.WalletTableHeaderChanged)

	t.ConnectWalletTableAlternate(c.WalletTableAlternate)
	c.ConnectWalletTableAlternateChanged(t.WalletTableAlternateChanged)

	t.ConnectWalletTableHighlight(c.WalletTableHighlight)
	c.ConnectWalletTableHighlightChanged(t.WalletTableHighlightChanged)

	t.ConnectInputFieldBackground(c.InputFieldBackground)
	c.ConnectInputFieldBackgroundChanged(t.InputFieldBackgroundChanged)

	t.ConnectFont(c.Font)
	c.ConnectFontChanged(t.FontChanged)

	t.ConnectFontHighlight(c.FontHighlight)
	c.ConnectFontHighlightChanged(t.FontHighlightChanged)
}
