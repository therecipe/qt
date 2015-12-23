package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioProbe struct {
	core.QObject
}

type QAudioProbe_ITF interface {
	core.QObject_ITF
	QAudioProbe_PTR() *QAudioProbe
}

func PointerFromQAudioProbe(ptr QAudioProbe_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioProbe_PTR().Pointer()
	}
	return nil
}

func NewQAudioProbeFromPointer(ptr unsafe.Pointer) *QAudioProbe {
	var n = new(QAudioProbe)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioProbe_") {
		n.SetObjectName("QAudioProbe_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioProbe) QAudioProbe_PTR() *QAudioProbe {
	return ptr
}

func NewQAudioProbe(parent core.QObject_ITF) *QAudioProbe {
	defer qt.Recovering("QAudioProbe::QAudioProbe")

	return NewQAudioProbeFromPointer(C.QAudioProbe_NewQAudioProbe(core.PointerFromQObject(parent)))
}

func (ptr *QAudioProbe) ConnectFlush(f func()) {
	defer qt.Recovering("connect QAudioProbe::flush")

	if ptr.Pointer() != nil {
		C.QAudioProbe_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QAudioProbe) DisconnectFlush() {
	defer qt.Recovering("disconnect QAudioProbe::flush")

	if ptr.Pointer() != nil {
		C.QAudioProbe_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQAudioProbeFlush
func callbackQAudioProbeFlush(ptrName *C.char) {
	defer qt.Recovering("callback QAudioProbe::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioProbe) IsActive() bool {
	defer qt.Recovering("QAudioProbe::isActive")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource(source QMediaObject_ITF) bool {
	defer qt.Recovering("QAudioProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource(ptr.Pointer(), PointerFromQMediaObject(source)) != 0
	}
	return false
}

func (ptr *QAudioProbe) SetSource2(mediaRecorder QMediaRecorder_ITF) bool {
	defer qt.Recovering("QAudioProbe::setSource")

	if ptr.Pointer() != nil {
		return C.QAudioProbe_SetSource2(ptr.Pointer(), PointerFromQMediaRecorder(mediaRecorder)) != 0
	}
	return false
}

func (ptr *QAudioProbe) DestroyQAudioProbe() {
	defer qt.Recovering("QAudioProbe::~QAudioProbe")

	if ptr.Pointer() != nil {
		C.QAudioProbe_DestroyQAudioProbe(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioProbe) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioProbe::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioProbeTimerEvent
func callbackQAudioProbeTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioProbe::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAudioProbe) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioProbe::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioProbeChildEvent
func callbackQAudioProbeChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioProbe::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAudioProbe) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioProbe) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioProbe::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioProbeCustomEvent
func callbackQAudioProbeCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioProbe::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
