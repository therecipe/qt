package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QVideoRendererControl struct {
	QMediaControl
}

type QVideoRendererControl_ITF interface {
	QMediaControl_ITF
	QVideoRendererControl_PTR() *QVideoRendererControl
}

func PointerFromQVideoRendererControl(ptr QVideoRendererControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoRendererControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoRendererControlFromPointer(ptr unsafe.Pointer) *QVideoRendererControl {
	var n = new(QVideoRendererControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoRendererControl_") {
		n.SetObjectName("QVideoRendererControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoRendererControl) QVideoRendererControl_PTR() *QVideoRendererControl {
	return ptr
}

func (ptr *QVideoRendererControl) SetSurface(surface QAbstractVideoSurface_ITF) {
	defer qt.Recovering("QVideoRendererControl::setSurface")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_SetSurface(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QVideoRendererControl) Surface() *QAbstractVideoSurface {
	defer qt.Recovering("QVideoRendererControl::surface")

	if ptr.Pointer() != nil {
		return NewQAbstractVideoSurfaceFromPointer(C.QVideoRendererControl_Surface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoRendererControl) DestroyQVideoRendererControl() {
	defer qt.Recovering("QVideoRendererControl::~QVideoRendererControl")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_DestroyQVideoRendererControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoRendererControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoRendererControlTimerEvent
func callbackQVideoRendererControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoRendererControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoRendererControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoRendererControlChildEvent
func callbackQVideoRendererControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoRendererControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoRendererControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoRendererControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoRendererControlCustomEvent
func callbackQVideoRendererControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoRendererControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoRendererControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoRendererControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoRendererControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoRendererControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoRendererControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
