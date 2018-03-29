// +build vivid

package main

import (
	"os"

	"github.com/therecipe/qt/core"
)

func EnableHighDPI()               { os.Setenv("QT_DEVICE_PIXEL_RATIO", "auto") }
func NewQUrl(in string) *core.QUrl { return core.NewQUrl4("qrc:/qml/view.qml", 0) }
