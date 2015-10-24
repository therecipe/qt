package macextras

//#include "qmactoolbar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacToolBar struct {
	core.QObject
}

type QMacToolBarITF interface {
	core.QObjectITF
	QMacToolBarPTR() *QMacToolBar
}

func PointerFromQMacToolBar(ptr QMacToolBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarPTR().Pointer()
	}
	return nil
}

func QMacToolBarFromPointer(ptr unsafe.Pointer) *QMacToolBar {
	var n = new(QMacToolBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMacToolBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacToolBar) QMacToolBarPTR() *QMacToolBar {
	return ptr
}

func NewQMacToolBar(parent core.QObjectITF) *QMacToolBar {
	return QMacToolBarFromPointer(unsafe.Pointer(C.QMacToolBar_NewQMacToolBar(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQMacToolBar2(identifier string, parent core.QObjectITF) *QMacToolBar {
	return QMacToolBarFromPointer(unsafe.Pointer(C.QMacToolBar_NewQMacToolBar2(C.CString(identifier), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIconITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		return QMacToolBarItemFromPointer(unsafe.Pointer(C.QMacToolBar_AddAllowedItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIconITF, text string) *QMacToolBarItem {
	if ptr.Pointer() != nil {
		return QMacToolBarItemFromPointer(unsafe.Pointer(C.QMacToolBar_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text))))
	}
	return nil
}

func (ptr *QMacToolBar) AddSeparator() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindowITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQWindow(window)))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
