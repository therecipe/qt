// +build cmd

package cmd

import (
	_ "github.com/therecipe/qt/internal"

	_ "github.com/therecipe/qt/cmd/qtdeploy"
	_ "github.com/therecipe/qt/cmd/qtminimal"
	_ "github.com/therecipe/qt/cmd/qtmoc"
	_ "github.com/therecipe/qt/cmd/qtrcc"
	_ "github.com/therecipe/qt/cmd/qtsetup"

	_ "github.com/gopherjs/gopherjs"
)
