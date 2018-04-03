package controller

import "github.com/therecipe/qt/core"

var Controller *themeController

type themeController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"name"`

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

	_ func() `signal:"change"`

	_ func() `signal:"hide"`
}

func (c *themeController) init() {
	Controller = c

	c.ConnectNameChanged(c.nameChanged)

	c.ConnectAccent(c.accent)
	c.ConnectNextAccent(c.nextAccent)

	c.ConnectBackground(c.background)
	c.ConnectDarkBackground(c.darkBackground)

	c.ConnectWalletTableHeader(c.walletTableHeader)
	c.ConnectWalletTableAlternate(c.walletTableAlternate)
	c.ConnectWalletTableHighlight(c.walletTableHighlight)

	c.ConnectInputFieldBackground(c.inputFieldBackground)

	c.ConnectFont(c.font)
	c.ConnectFontHighlight(c.fontHighlight)

	c.SetName("dark")

	c.ConnectChange(c.change)

	initColorDialog()
}

func (c *themeController) nameChanged(string) {
	c.AccentChanged(c.Accent())
	c.NextAccentChanged(c.NextAccent())

	c.BackgroundChanged(c.Background())
	c.DarkBackgroundChanged(c.DarkBackground())

	c.WalletTableHeaderChanged(c.WalletTableHeader())
	c.WalletTableAlternateChanged(c.WalletTableAlternate())
	c.WalletTableHighlightChanged(c.WalletTableHighlight())

	c.InputFieldBackgroundChanged(c.InputFieldBackground())

	c.FontChanged(c.Font())
	c.FontHighlightChanged(c.FontHighlight())
}

func (c *themeController) accent() string {
	if color := c.AccentDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#00CA9F"

	default:
		return "#00CA9F"
	}
}

func (c *themeController) nextAccent() string {
	if color := c.NextAccentDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#00CA9F"

	default:
		return "#00CA9F"
	}
}

func (c *themeController) background() string {
	if color := c.BackgroundDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#343C58"

	default:
		return "#FFFFFF"
	}
}

func (c *themeController) darkBackground() string {
	if color := c.DarkBackgroundDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#272E49"

	default:
		return "#E6E9EE"
	}
}

func (c *themeController) walletTableHeader() string {
	if color := c.WalletTableHeaderDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#4C547A"

	default:
		return "#18CCA1"
	}
}

func (c *themeController) walletTableAlternate() string {
	if color := c.WalletTableAlternateDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#2F3655"

	default:
		return "#3CD6B4"
	}
}

func (c *themeController) walletTableHighlight() string {
	if color := c.WalletTableHighlightDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#FDB06F"

	default:
		return "white"
	}
}

func (c *themeController) inputFieldBackground() string {
	if color := c.InputFieldBackgroundDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#232942"

	default:
		return "#747579"
	}
}

func (c *themeController) font() string {
	if color := c.FontDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "#A9AEBE"

	default:
		return "#B1B2B3"
	}
}

func (c *themeController) fontHighlight() string {
	if color := c.FontHighlightDefault(); color != "" {
		return color
	}

	switch c.Name() {
	case "dark":
		return "white"

	default:
		return "black"
	}
}

func (c *themeController) change() {
	colorDialog.Show()
	//TODO: there is no light theme yet
	/*
		if c.Name() == "dark" {
			c.SetName("light")
		} else {
			c.SetName("dark")
		}
	*/
}
