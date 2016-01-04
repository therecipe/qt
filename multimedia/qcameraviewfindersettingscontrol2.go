package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraViewfinderSettingsControl2 struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl2_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2
}

func PointerFromQCameraViewfinderSettingsControl2(ptr QCameraViewfinderSettingsControl2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl2_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControl2FromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	var n = new(QCameraViewfinderSettingsControl2)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl2_") {
		n.SetObjectName("QCameraViewfinderSettingsControl2_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl2) QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2 {
	return ptr
}

func (ptr *QCameraViewfinderSettingsControl2) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::setViewfinderSettings")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_SetViewfinderSettings(ptr.Pointer(), PointerFromQCameraViewfinderSettings(settings))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DestroyQCameraViewfinderSettingsControl2() {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::~QCameraViewfinderSettingsControl2")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2TimerEvent
func callbackQCameraViewfinderSettingsControl2TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2ChildEvent
func callbackQCameraViewfinderSettingsControl2ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraViewfinderSettingsControl2) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraViewfinderSettingsControl2CustomEvent
func callbackQCameraViewfinderSettingsControl2CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraViewfinderSettingsControl2::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraViewfinderSettingsControl2FromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraViewfinderSettingsControl2) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraViewfinderSettingsControl2::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl2_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
