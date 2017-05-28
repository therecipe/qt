package main

import "github.com/therecipe/qt/core"

var (
	materialRed        = []string{"#FFCDD2", "#F44336", "#D32F2F", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialPink       = []string{"#F8BBD0", "#E91E63", "#C2185B", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialPurple     = []string{"#E1BEE7", "#9C27B0", "#7B1FA2", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialDeepPurple = []string{"#D1C4E9", "#673AB7", "#512DA8", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialIndigo     = []string{"#C5CAE9", "#3F51B5", "#303F9F", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialBlue       = []string{"#BBDEFB", "#2196F3", "#1976D2", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialLightBlue  = []string{"#B3E5FC", "#03A9F4", "#0288D1", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialCyan       = []string{"#B2EBF2", "#00BCD4", "#0097A7", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialTeal       = []string{"#B2DFDB", "#009688", "#00796B", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialGreen      = []string{"#C8E6C9", "#4CAF50", "#388E3C", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialLightGreen = []string{"#DCEDC8", "#8BC34A", "#689F38", "#000000", "#000000", "#000000", "black", "black", "black"}
	materialLime       = []string{"#F0F4C3", "#CDDC39", "#AFB42B", "#000000", "#000000", "#000000", "black", "black", "black"}
	materialYellow     = []string{"#FFF9C4", "#FFEB3B", "#FBC02D", "#000000", "#000000", "#000000", "black", "black", "black"}
	materialAmber      = []string{"#FFECB3", "#FFC107", "#FFA000", "#000000", "#000000", "#000000", "black", "black", "black"}
	materialOrange     = []string{"#FFE0B2", "#FF9800", "#F57C00", "#000000", "#000000", "#000000", "black", "black", "black"}
	materialDeepOrange = []string{"#FFCCBC", "#FF5722", "#E64A19", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialBrown      = []string{"#D7CCC8", "#795548", "#5D4037", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}
	materialGrey       = []string{"#F5F5F5", "#9E9E9E", "#616161", "#000000", "#000000", "#FFFFFF", "black", "black", "white"}
	materialBlueGrey   = []string{"#CFD8DC", "#607D8B", "#455A64", "#000000", "#FFFFFF", "#FFFFFF", "black", "white", "white"}

	darkPalette  = []string{"#FFFFFF", "#424242", "1.0", "0.70", "0.12", "1.0", "0.3", "white", "1", "#FFFFFF", "#FFFFFF", "1.0", "0.7", "Darkgrey", "0.9"}
	lightPalette = []string{"#000000", "#FFFFFF", "0.87", "0.54", "0.12", "0.54", "0.26", "black", "0", "#424242", "#424242", "1.0", "0.7", "#323232", "0.75"}
)

type ApplicationUI struct {
	core.QObject

	_ func() []string                 `slot:"swapThemePalette"`
	_ func() []string                 `slot:"defaultThemePalette"`
	_ func(paletteIndex int) []string `slot:"primaryPalette"`
	_ func(paletteIndex int) []string `slot:"accentPalette"`
	_ func() []string                 `slot:"defaultPrimaryPalette"`
	_ func() []string                 `slot:"defaultAccentPalette"`

	mIsDarkTheme    bool
	mPrimaryPalette int
	mAccentPalette  int
}

func initApplicationUI() *ApplicationUI {
	var this = NewApplicationUI(nil)

	// default theme is light
	this.mIsDarkTheme = false
	// default primary color is Teal
	this.mPrimaryPalette = 8
	// default accent color is DeepOrange
	this.mAccentPalette = 15

	this.ConnectSwapThemePalette(this.swapThemePalette)
	this.ConnectDefaultThemePalette(this.defaultThemePalette)
	this.ConnectPrimaryPalette(this.primaryPalette)
	this.ConnectAccentPalette(this.accentPalette)
	this.ConnectDefaultPrimaryPalette(this.defaultPrimaryPalette)
	this.ConnectDefaultAccentPalette(this.defaultAccentPalette)

	return this
}

/* Change Theme Palette */
func (this *ApplicationUI) swapThemePalette() []string {
	this.mIsDarkTheme = !this.mIsDarkTheme
	if this.mIsDarkTheme {
		return darkPalette
	}
	return lightPalette
}

/* Get current default Theme Palette */
func (this *ApplicationUI) defaultThemePalette() []string {
	if this.mIsDarkTheme {
		return darkPalette
	}
	return lightPalette
}

/* Get one of the Primary Palettes */
func (this *ApplicationUI) primaryPalette(paletteIndex int) []string {
	this.mPrimaryPalette = paletteIndex
	switch paletteIndex {
	case 0:
		return materialRed

	case 1:
		return materialPink

	case 2:
		return materialPurple

	case 3:
		return materialDeepPurple

	case 4:
		return materialIndigo

	case 5:
		return materialBlue

	case 6:
		return materialLightBlue

	case 7:
		return materialCyan

	case 8:
		return materialTeal

	case 9:
		return materialGreen

	case 10:
		return materialLightGreen

	case 11:
		return materialLime

	case 12:
		return materialYellow

	case 13:
		return materialAmber

	case 14:
		return materialOrange

	case 15:
		return materialDeepOrange

	case 16:
		return materialBrown

	case 17:
		return materialGrey

	default:
		return materialBlueGrey
	}
}

/* Get one of the Accent Palettes */
func (this *ApplicationUI) accentPalette(paletteIndex int) []string {
	this.mAccentPalette = paletteIndex
	var currentPrimary = this.mPrimaryPalette
	var thePalette = this.primaryPalette(paletteIndex)
	this.mPrimaryPalette = currentPrimary
	// we need: primaryColor, textOnPrimary, iconOnPrimaryFolder
	return []string{thePalette[1], thePalette[4], thePalette[7]}
}

/* Get Default Primary Palette */
func (this *ApplicationUI) defaultPrimaryPalette() []string {
	return this.primaryPalette(this.mPrimaryPalette)
}

/* Get Default Accent Palette */
func (this *ApplicationUI) defaultAccentPalette() []string {
	return this.accentPalette(this.mAccentPalette)
}
