package macextras

//#include "macextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QMacToolBarItem struct {
	core.QObject
}

type QMacToolBarItem_ITF interface {
	core.QObject_ITF
	QMacToolBarItem_PTR() *QMacToolBarItem
}

func PointerFromQMacToolBarItem(ptr QMacToolBarItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBarItem_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarItemFromPointer(ptr unsafe.Pointer) *QMacToolBarItem {
	var n = new(QMacToolBarItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacToolBarItem_") {
		n.SetObjectName("QMacToolBarItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacToolBarItem) QMacToolBarItem_PTR() *QMacToolBarItem {
	return ptr
}

//QMacToolBarItem::StandardItem
type QMacToolBarItem__StandardItem int64

const (
	QMacToolBarItem__NoStandardItem = QMacToolBarItem__StandardItem(0)
	QMacToolBarItem__Space          = QMacToolBarItem__StandardItem(1)
	QMacToolBarItem__FlexibleSpace  = QMacToolBarItem__StandardItem(2)
)

func NewQMacToolBarItem(parent core.QObject_ITF) *QMacToolBarItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::QMacToolBarItem")
		}
	}()

	return NewQMacToolBarItemFromPointer(C.QMacToolBarItem_NewQMacToolBarItem(core.PointerFromQObject(parent)))
}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQMacToolBarItemActivated
func callbackQMacToolBarItemActivated(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::activated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activated").(func())()
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::~QMacToolBarItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) Selectable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::selectable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Selectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::setSelectable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(ptr.Pointer(), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::setStandardItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(ptr.Pointer(), C.int(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::standardItem")
		}
	}()

	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMacToolBarItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacToolBarItem_Text(ptr.Pointer()))
	}
	return ""
}
