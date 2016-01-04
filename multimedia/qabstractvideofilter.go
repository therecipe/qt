package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractVideoFilter struct {
	core.QObject
}

type QAbstractVideoFilter_ITF interface {
	core.QObject_ITF
	QAbstractVideoFilter_PTR() *QAbstractVideoFilter
}

func PointerFromQAbstractVideoFilter(ptr QAbstractVideoFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoFilter_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoFilterFromPointer(ptr unsafe.Pointer) *QAbstractVideoFilter {
	var n = new(QAbstractVideoFilter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractVideoFilter_") {
		n.SetObjectName("QAbstractVideoFilter_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilter_PTR() *QAbstractVideoFilter {
	return ptr
}

func (ptr *QAbstractVideoFilter) IsActive() bool {
	defer qt.Recovering("QAbstractVideoFilter::isActive")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoFilter_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoFilter) SetActive(v bool) {
	defer qt.Recovering("QAbstractVideoFilter::setActive")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractVideoFilter) ConnectActiveChanged(f func()) {
	defer qt.Recovering("connect QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoFilterActiveChanged
func callbackQAbstractVideoFilterActiveChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoFilter::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractVideoFilter) ActiveChanged() {
	defer qt.Recovering("QAbstractVideoFilter::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ActiveChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable {
	defer qt.Recovering("QAbstractVideoFilter::createFilterRunnable")

	if ptr.Pointer() != nil {
		return NewQVideoFilterRunnableFromPointer(C.QAbstractVideoFilter_CreateFilterRunnable(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoFilter) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractVideoFilterTimerEvent
func callbackQAbstractVideoFilterTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractVideoFilterChildEvent
func callbackQAbstractVideoFilterChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractVideoFilterCustomEvent
func callbackQAbstractVideoFilterCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoFilter::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractVideoFilterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoFilter) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractVideoFilter) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoFilter::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
