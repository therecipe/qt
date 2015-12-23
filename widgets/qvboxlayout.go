package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVBoxLayout struct {
	QBoxLayout
}

type QVBoxLayout_ITF interface {
	QBoxLayout_ITF
	QVBoxLayout_PTR() *QVBoxLayout
}

func PointerFromQVBoxLayout(ptr QVBoxLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func NewQVBoxLayoutFromPointer(ptr unsafe.Pointer) *QVBoxLayout {
	var n = new(QVBoxLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVBoxLayout_") {
		n.SetObjectName("QVBoxLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QVBoxLayout) QVBoxLayout_PTR() *QVBoxLayout {
	return ptr
}

func NewQVBoxLayout() *QVBoxLayout {
	defer qt.Recovering("QVBoxLayout::QVBoxLayout")

	return NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout())
}

func NewQVBoxLayout2(parent QWidget_ITF) *QVBoxLayout {
	defer qt.Recovering("QVBoxLayout::QVBoxLayout")

	return NewQVBoxLayoutFromPointer(C.QVBoxLayout_NewQVBoxLayout2(PointerFromQWidget(parent)))
}

func (ptr *QVBoxLayout) DestroyQVBoxLayout() {
	defer qt.Recovering("QVBoxLayout::~QVBoxLayout")

	if ptr.Pointer() != nil {
		C.QVBoxLayout_DestroyQVBoxLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVBoxLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	defer qt.Recovering("connect QVBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addItem", f)
	}
}

func (ptr *QVBoxLayout) DisconnectAddItem() {
	defer qt.Recovering("disconnect QVBoxLayout::addItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addItem")
	}
}

//export callbackQVBoxLayoutAddItem
func callbackQVBoxLayoutAddItem(ptrName *C.char, item unsafe.Pointer) bool {
	defer qt.Recovering("callback QVBoxLayout::addItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "addItem"); signal != nil {
		signal.(func(*QLayoutItem))(NewQLayoutItemFromPointer(item))
		return true
	}
	return false

}

func (ptr *QVBoxLayout) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QVBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "invalidate", f)
	}
}

func (ptr *QVBoxLayout) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QVBoxLayout::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "invalidate")
	}
}

//export callbackQVBoxLayoutInvalidate
func callbackQVBoxLayoutInvalidate(ptrName *C.char) bool {
	defer qt.Recovering("callback QVBoxLayout::invalidate")

	if signal := qt.GetSignal(C.GoString(ptrName), "invalidate"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QVBoxLayout) ConnectSetGeometry(f func(r *core.QRect)) {
	defer qt.Recovering("connect QVBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QVBoxLayout) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QVBoxLayout::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

//export callbackQVBoxLayoutSetGeometry
func callbackQVBoxLayoutSetGeometry(ptrName *C.char, r unsafe.Pointer) bool {
	defer qt.Recovering("callback QVBoxLayout::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(r))
		return true
	}
	return false

}

func (ptr *QVBoxLayout) ConnectChildEvent(f func(e *core.QChildEvent)) {
	defer qt.Recovering("connect QVBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVBoxLayout) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVBoxLayout::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVBoxLayoutChildEvent
func callbackQVBoxLayoutChildEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QVBoxLayout::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QVBoxLayout) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVBoxLayout) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVBoxLayout::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVBoxLayoutTimerEvent
func callbackQVBoxLayoutTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVBoxLayout::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QVBoxLayout) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVBoxLayout) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVBoxLayout::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVBoxLayoutCustomEvent
func callbackQVBoxLayoutCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVBoxLayout::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
