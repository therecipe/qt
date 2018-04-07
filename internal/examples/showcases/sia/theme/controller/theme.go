package controller

import "github.com/therecipe/qt/core"

var Controller *ThemeController

type ThemeController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"name,auto,changed"`

	_ string `property:"accent,auto,get"`
	_ string `property:"nextAccent,auto,get"`

	_ string `property:"background,auto,get"`
	_ string `property:"darkBackground,auto,get"`

	_ string `property:"walletTableHeader,auto,get"`
	_ string `property:"walletTableAlternate,auto,get"`
	_ string `property:"walletTableHighlight,auto,get"`

	_ string `property:"inputFieldBackground,auto,get"`

	_ string `property:"font,auto,get"`
	_ string `property:"fontHighlight,auto,get"`

	_ func() `signal:"change,auto"`

	_ func() `signal:"hide"`
}

func (c *ThemeController) init() {
	Controller = c

	c.SetName("dark")

	initColorDialog()
}

func (c *ThemeController) nameChanged(string) {
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

func (c *ThemeController) accent() string {
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

func (c *ThemeController) nextAccent() string {
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

func (c *ThemeController) background() string {
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

func (c *ThemeController) darkBackground() string {
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

func (c *ThemeController) walletTableHeader() string {
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

func (c *ThemeController) walletTableAlternate() string {
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

func (c *ThemeController) walletTableHighlight() string {
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

func (c *ThemeController) inputFieldBackground() string {
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

func (c *ThemeController) font() string {
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

func (c *ThemeController) fontHighlight() string {
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

func (c *ThemeController) change() {
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
