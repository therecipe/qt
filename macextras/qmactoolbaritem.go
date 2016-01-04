package macextras

//#include "macextras.h"
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
		n.SetObjectName("QMacToolBarItem_" + qt.Identifier())
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
	defer qt.Recovering("QMacToolBarItem::QMacToolBarItem")

	return NewQMacToolBarItemFromPointer(C.QMacToolBarItem_NewQMacToolBarItem(core.PointerFromQObject(parent)))
}

func (ptr *QMacToolBarItem) ConnectActivated(f func()) {
	defer qt.Recovering("connect QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectActivated() {
	defer qt.Recovering("disconnect QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQMacToolBarItemActivated
func callbackQMacToolBarItemActivated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMacToolBarItem::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMacToolBarItem) Activated() {
	defer qt.Recovering("QMacToolBarItem::activated")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_Activated(ptr.Pointer())
	}
}

func (ptr *QMacToolBarItem) DestroyQMacToolBarItem() {
	defer qt.Recovering("QMacToolBarItem::~QMacToolBarItem")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_DestroyQMacToolBarItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMacToolBarItem) Icon() *gui.QIcon {
	defer qt.Recovering("QMacToolBarItem::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QMacToolBarItem_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMacToolBarItem) Selectable() bool {
	defer qt.Recovering("QMacToolBarItem::selectable")

	if ptr.Pointer() != nil {
		return C.QMacToolBarItem_Selectable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMacToolBarItem) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QMacToolBarItem::setIcon")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMacToolBarItem) SetSelectable(selectable bool) {
	defer qt.Recovering("QMacToolBarItem::setSelectable")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetSelectable(ptr.Pointer(), C.int(qt.GoBoolToInt(selectable)))
	}
}

func (ptr *QMacToolBarItem) SetStandardItem(standardItem QMacToolBarItem__StandardItem) {
	defer qt.Recovering("QMacToolBarItem::setStandardItem")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetStandardItem(ptr.Pointer(), C.int(standardItem))
	}
}

func (ptr *QMacToolBarItem) SetText(text string) {
	defer qt.Recovering("QMacToolBarItem::setText")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMacToolBarItem) StandardItem() QMacToolBarItem__StandardItem {
	defer qt.Recovering("QMacToolBarItem::standardItem")

	if ptr.Pointer() != nil {
		return QMacToolBarItem__StandardItem(C.QMacToolBarItem_StandardItem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMacToolBarItem) Text() string {
	defer qt.Recovering("QMacToolBarItem::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMacToolBarItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMacToolBarItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMacToolBarItemTimerEvent
func callbackQMacToolBarItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBarItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMacToolBarItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMacToolBarItemChildEvent
func callbackQMacToolBarItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBarItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::childEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMacToolBarItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMacToolBarItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMacToolBarItemCustomEvent
func callbackQMacToolBarItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMacToolBarItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMacToolBarItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMacToolBarItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMacToolBarItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMacToolBarItem::customEvent")

	if ptr.Pointer() != nil {
		C.QMacToolBarItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
