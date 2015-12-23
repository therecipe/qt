package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextList struct {
	QTextBlockGroup
}

type QTextList_ITF interface {
	QTextBlockGroup_ITF
	QTextList_PTR() *QTextList
}

func PointerFromQTextList(ptr QTextList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextList_PTR().Pointer()
	}
	return nil
}

func NewQTextListFromPointer(ptr unsafe.Pointer) *QTextList {
	var n = new(QTextList)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextList_") {
		n.SetObjectName("QTextList_" + qt.Identifier())
	}
	return n
}

func (ptr *QTextList) QTextList_PTR() *QTextList {
	return ptr
}

func (ptr *QTextList) ItemNumber(block QTextBlock_ITF) int {
	defer qt.Recovering("QTextList::itemNumber")

	if ptr.Pointer() != nil {
		return int(C.QTextList_ItemNumber(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return 0
}

func (ptr *QTextList) ItemText(block QTextBlock_ITF) string {
	defer qt.Recovering("QTextList::itemText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextList_ItemText(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return ""
}

func (ptr *QTextList) Add(block QTextBlock_ITF) {
	defer qt.Recovering("QTextList::add")

	if ptr.Pointer() != nil {
		C.QTextList_Add(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QTextList) Count() int {
	defer qt.Recovering("QTextList::count")

	if ptr.Pointer() != nil {
		return int(C.QTextList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextList) RemoveItem(i int) {
	defer qt.Recovering("QTextList::removeItem")

	if ptr.Pointer() != nil {
		C.QTextList_RemoveItem(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTextList) SetFormat(format QTextListFormat_ITF) {
	defer qt.Recovering("QTextList::setFormat")

	if ptr.Pointer() != nil {
		C.QTextList_SetFormat(ptr.Pointer(), PointerFromQTextListFormat(format))
	}
}

func (ptr *QTextList) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTextList::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTextList) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTextList::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTextListTimerEvent
func callbackQTextListTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextList::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextList) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTextList::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTextList) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTextList::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTextListChildEvent
func callbackQTextListChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextList::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTextList) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTextList::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTextList) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTextList::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTextListCustomEvent
func callbackQTextListCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTextList::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
