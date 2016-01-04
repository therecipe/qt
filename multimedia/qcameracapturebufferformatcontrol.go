package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraCaptureBufferFormatControl struct {
	QMediaControl
}

type QCameraCaptureBufferFormatControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl
}

func PointerFromQCameraCaptureBufferFormatControl(ptr QCameraCaptureBufferFormatControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraCaptureBufferFormatControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraCaptureBufferFormatControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureBufferFormatControl {
	var n = new(QCameraCaptureBufferFormatControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraCaptureBufferFormatControl_") {
		n.SetObjectName("QCameraCaptureBufferFormatControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraCaptureBufferFormatControl) QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl {
	return ptr
}

func (ptr *QCameraCaptureBufferFormatControl) BufferFormat() QVideoFrame__PixelFormat {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::bufferFormat")

	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraCaptureBufferFormatControl_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectBufferFormatChanged() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraCaptureBufferFormatControlBufferFormatChanged
func callbackQCameraCaptureBufferFormatControlBufferFormatChanged(ptr unsafe.Pointer, ptrName *C.char, format C.int) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged"); signal != nil {
		signal.(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
	}

}

func (ptr *QCameraCaptureBufferFormatControl) BufferFormatChanged(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::bufferFormatChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_BufferFormatChanged(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) SetBufferFormat(format QVideoFrame__PixelFormat) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::setBufferFormat")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DestroyQCameraCaptureBufferFormatControl() {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::~QCameraCaptureBufferFormatControl")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlTimerEvent
func callbackQCameraCaptureBufferFormatControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlChildEvent
func callbackQCameraCaptureBufferFormatControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCaptureBufferFormatControlCustomEvent
func callbackQCameraCaptureBufferFormatControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraCaptureBufferFormatControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureBufferFormatControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
