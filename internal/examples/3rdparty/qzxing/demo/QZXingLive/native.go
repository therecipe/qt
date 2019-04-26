package main

import "github.com/therecipe/qt/core"

var NativeHelper = new(nativeHelpers)

type nativeHelpers struct{ application_p_ *core.QObject }

func (nh *nativeHelpers) registerApplicationInstance(app_p *core.QObject) {
	nh.application_p_ = app_p
}

func (nh *nativeHelpers) getApplicationInstance() *core.QObject {
	return nh.application_p_
}
