package main

import (
	"runtime"

	"github.com/therecipe/qt/androidextras"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

var Application *application

type application struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func() `signal:"onPermissionsGranted"`
	_ func() `signal:"onPermissionsDenied"`

	engine *qml.QQmlApplicationEngine
}

func (a *application) init() {
	Application = a

	a.engine = qml.NewQQmlApplicationEngine(nil)
	a.ConnectOnPermissionsGranted(a.initializeQML)
	a.ConnectOnPermissionsDenied(a.initializeQML)
}

func (a *application) initializeQML() {
	a.engine.Load(core.NewQUrl3("qrc:/main.qml", 0))
}

func (a *application) checkPermissions() {
	if runtime.GOOS == "android" {
		//intentionally called in the C++ thread since it is blocking and will continue after the check
		println("About to request permissions")
		androidextras.QAndroidJniObject_CallStaticMethodVoid2("org/ftylitak/qzxing/Utilities", "requestQZXingPermissions", "(Landroid/app/Activity;)V", androidextras.QtAndroid_AndroidActivity().Object())
		println("Permissions granted")
	} else {
		a.OnPermissionsGranted()
	}
}
