package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoProbe struct {
	core.QObject
}

type QVideoProbe_ITF interface {
	core.QObject_ITF
	QVideoProbe_PTR() *QVideoProbe
}

func PointerFromQVideoProbe(ptr QVideoProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoProbe_PTR().Pointer()
	}
	return nil
}

func NewQVideoProbeFromPointer(ptr unsafe.Pointer) *QVideoProbe {
	var n = new(QVideoProbe)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoProbe_") {
		n.SetObjectName("QVideoProbe_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoProbe) QVideoProbe_PTR() *QVideoProbe {
	return ptr
}

func NewQVideoProbe(parent core.QObject_ITF) *QVideoProbe {
	defer qt.Recovering("QVideoProbe::QVideoProbe")

	return NewQVideoProbeFromPointer(C.QVideoProbe_NewQVideoProbe(core.PointerFromQObject(parent)))
}

func (ptr *QVideoProbe) ConnectFlush(f func()) {
	defer qt.Recovering("connect QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QVideoProbe) DisconnectFlush() {
	defer qt.Recovering("disconnect QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQVideoProbeFlush
func callbackQVideoProbeFlush(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoProbe::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoProbe) Flush() {
	defer qt.Recovering("QVideoProbe::flush")

	if ptr.Pointer() != nil {
		C.QVideoProbe_Flush(ptr.Pointer())
	}
}

func (ptr *QVideoProbe) IsActive() bool {
	defer qt.Recovering("QVideoProbe::isActive")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource(source QMediaObject_ITF) bool {
	defer qt.Recovering("QVideoProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QVideoProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	defer qt.Recovering("QVideoProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QVideoProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QVideoProbe) DestroyQVideoProbe() {
	defer qt.Recovering("QVideoProbe::~QVideoProbe")

	if ptr.Pointer() != nil {
		C.QVideoProbe_DestroyQVideoProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoProbe) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoProbeTimerEvent
func callbackQVideoProbeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoProbe) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoProbe::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoProbe) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoProbeChildEvent
func callbackQVideoProbeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoProbe) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoProbe::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoProbe) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoProbe) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoProbeCustomEvent
func callbackQVideoProbeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoProbe::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoProbeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoProbe) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoProbe) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoProbe::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoProbe_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
