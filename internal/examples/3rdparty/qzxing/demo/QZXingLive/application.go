package main

import (
	"runtime"

	"github.com/therecipe/qt/androidextras"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type Application struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func() `slot:"initializeQML,auto"`

	_ func() `signal:"onPermissionsGranted,auto(this.Application_PTR().initializeQML)"`
	_ func() `signal:"onPermissionsDenied,auto(this.Application_PTR().initializeQML)"`

	engine *qml.QQmlApplicationEngine
}

func (a *Application) init() {
	NativeHelper.registerApplicationInstance(a.QObject_PTR())
	a.engine = qml.NewQQmlApplicationEngine(nil)
}

func (a *Application) initializeQML() {
	a.engine.Load(core.NewQUrl3("qrc:/main.qml", 0))
}

func (a *Application) checkPermissions() {
	if runtime.GOOS == "android" {
		//intentionally called in the C++ thread since it is blocking and will continue after the check
		println("About to request permissions")
		androidextras.QAndroidJniObject_CallStaticMethodVoid2("org/ftylitak/qzxing/Utilities", "requestQZXingPermissions", "(Landroid/app/Activity;)V", androidextras.QtAndroid_AndroidActivity().Object())
		println("Permissions granted")
	} else {
		a.OnPermissionsGranted()
	}
}
