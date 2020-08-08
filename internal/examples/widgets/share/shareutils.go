package main

import (
	"runtime"

	"github.com/StarAurryon/qt/androidextras"
	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/quick"
)

type ShareUtils struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ func(text string, url *core.QUrl) `slot:"share"`
}

func (u *ShareUtils) init() {

	switch {
	case runtime.GOOS == "android" && runtime.GOARCH == "arm":
		u.ConnectShare(func(text string, url *core.QUrl) {
			androidextras.QAndroidJniObject_CallStaticMethodVoid2Caught("com/lasconic/QShareUtils", "share", "(Ljava/lang/String;Ljava/lang/String;)V", text, url.ToString(0))
		})

	case runtime.GOOS == "darwin" && runtime.GOARCH == "arm64":
		//TODO:

	default:
		u.ConnectShare(func(text string, url *core.QUrl) {
			println(text, url.ToString(0))
		})
	}

}
