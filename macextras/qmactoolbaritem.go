package macextras

//#include "qmactoolbaritem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacToolBarItem struct {
	core.QObject
}

type QMacToolBarItemITF interface {
	core.QObjectITF
	QMacToolBarItemPTR() *QMacToolBarItem
}

func PointerFromQMacToolBarItem(ptr QMacToolBarItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarItemPTR().Pointer()
	}
	return nil
}

func QMacToolBarItemFromPointer(ptr unsafe.Pointer) *QMacToolBarItem {
	var n = new(QMacToolBarItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMacToolBarItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacToolBarItem) QMacToolBarItemPTR() *QMacToolBarItem {
	return ptr
}

//QMacToolBarItem::StandardItem
type QMacToolBarItem__StandardItem int

var (
	QMacToolBarItem__NoStandardItem = QMacToolBarItem__StandardItem(0)
	QMacToolBarItem__Space          = QMacToolBarItem__StandardItem(1)
	QMacToolBarItem__FlexibleSpace  = QMacToolBarItem__StandardItem(2)
)

func NewQMacToolBarItem(parent core.QObjectITF) *QMacToolBarItem {
	return QMacToolBarItemFromPointer(unsafe.Pointer(C.QMacToolBarItem_NewQMacToolBarItem(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQMacToolBarItemActivated
func callbackQMacToolBarItemActivated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func())()
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) Selectable() bool {
	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Selectable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(C.QtObjectPtr(ptr.Pointer()), C.int(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMacToolBarItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
