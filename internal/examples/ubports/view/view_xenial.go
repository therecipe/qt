//go:build xenial
// +build xenial

package main

import "github.com/akiyosi/qt/core"

func EnableHighDPI()               { core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true) }
func NewQUrl(in string) *core.QUrl { return core.NewQUrl3("qrc:/qml/view.qml", 0) }
