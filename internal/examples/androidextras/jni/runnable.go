package main

import "C"
import (
	"unsafe"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/androidextras"
)

type Runnable struct {
	ptr *androidextras.QAndroidJniObject
}

type Runnable_ITF interface {
	Runnable_PTR() *Runnable
}

func (p *Runnable) Pointer() *androidextras.QAndroidJniObject {
	return p.ptr
}

func (p *Runnable) SetPointer(ptr *androidextras.QAndroidJniObject) {
	p.ptr = ptr
}

func (ptr *Runnable) ConnectRun(f func()) {
	if ptr.Pointer() != nil && ptr.Pointer().IsValid() {

		qt.ConnectSignal(ptr.Pointer().ToString(), "Runnable::run", f)
	}
}

func (ptr *Runnable) DisconnectRun() {
	if ptr.Pointer() != nil && ptr.Pointer().IsValid() {

		qt.DisconnectSignal(ptr.Pointer().ToString(), "Runnable::run")
	}
}

//export Java_qt_java_lang_Runnable_qtrun
func Java_qt_java_lang_Runnable_qtrun(_, _, this unsafe.Pointer) {
	if signal := qt.GetSignal(androidextras.NewQAndroidJniObject6(this).ToString(), "Runnable::run"); signal != nil {
		signal.(func())()
	}

}

func (ptr *Runnable) Run() {
	if ptr.Pointer() != nil && ptr.Pointer().IsValid() {
		ptr.Pointer().CallMethodVoid2("run", "()V")
	}
}

func PointerFromRunnable(ptr Runnable_ITF) *androidextras.QAndroidJniObject {
	if ptr != nil {
		return ptr.Runnable_PTR().Pointer()
	}
	return nil
}

func NewRunnableFromPointer(ptr androidextras.QAndroidJniObject_ITF) *Runnable {
	var n = new(Runnable)
	n.SetPointer(ptr.QAndroidJniObject_PTR())
	return n
}

func (ptr *Runnable) Runnable_PTR() *Runnable {
	return ptr
}

func Runnable__IsClassAvailable() bool {
	return androidextras.QAndroidJniObject_IsClassAvailable("java/lang/Runnable")
}

func Runnable__IsQtClassAvailable() bool {
	return androidextras.QAndroidJniObject_IsClassAvailable("qt/java/lang/Runnable")
}

func (ptr *Runnable) SetObjectName(n string) {
	if ptr.Pointer() != nil && ptr.Pointer().IsValid() {
		ptr.Pointer().SetField2("ObjectName", "Ljava/lang/String;", n)
	}
}

func (ptr *Runnable) ObjectName() string {
	if ptr.Pointer() != nil && ptr.Pointer().IsValid() {
		return ptr.Pointer().GetObjectField2("ObjectName", "Ljava/lang/String;").ToString()
	}
	return ""
}

func NewQtRunnable() *Runnable {
	return NewRunnableFromPointer(androidextras.NewQAndroidJniObject3("qt/java/lang/Runnable", "()V"))
}
func NewRunnable() *Runnable {
	return NewRunnableFromPointer(androidextras.NewQAndroidJniObject3("java/lang/Runnable", "()V"))
}
