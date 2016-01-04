package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractVideoSurface struct {
	core.QObject
}

type QAbstractVideoSurface_ITF interface {
	core.QObject_ITF
	QAbstractVideoSurface_PTR() *QAbstractVideoSurface
}

func PointerFromQAbstractVideoSurface(ptr QAbstractVideoSurface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoSurface_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoSurfaceFromPointer(ptr unsafe.Pointer) *QAbstractVideoSurface {
	var n = new(QAbstractVideoSurface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractVideoSurface_") {
		n.SetObjectName("QAbstractVideoSurface_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractVideoSurface) QAbstractVideoSurface_PTR() *QAbstractVideoSurface {
	return ptr
}

//QAbstractVideoSurface::Error
type QAbstractVideoSurface__Error int64

const (
	QAbstractVideoSurface__NoError                = QAbstractVideoSurface__Error(0)
	QAbstractVideoSurface__UnsupportedFormatError = QAbstractVideoSurface__Error(1)
	QAbstractVideoSurface__IncorrectFormatError   = QAbstractVideoSurface__Error(2)
	QAbstractVideoSurface__StoppedError           = QAbstractVideoSurface__Error(3)
	QAbstractVideoSurface__ResourceError          = QAbstractVideoSurface__Error(4)
)

func (ptr *QAbstractVideoSurface) NativeResolution() *core.QSize {
	defer qt.Recovering("QAbstractVideoSurface::nativeResolution")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractVideoSurface_NativeResolution(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractVideoSurface) ConnectActiveChanged(f func(active bool)) {
	defer qt.Recovering("connect QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoSurfaceActiveChanged
func callbackQAbstractVideoSurfaceActiveChanged(ptr unsafe.Pointer, ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QAbstractVideoSurface::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func(bool))(int(active) != 0)
	}

}

func (ptr *QAbstractVideoSurface) ActiveChanged(active bool) {
	defer qt.Recovering("QAbstractVideoSurface::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ActiveChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QAbstractVideoSurface) Error() QAbstractVideoSurface__Error {
	defer qt.Recovering("QAbstractVideoSurface::error")

	if ptr.Pointer() != nil {
		return QAbstractVideoSurface__Error(C.QAbstractVideoSurface_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoSurface) IsActive() bool {
	defer qt.Recovering("QAbstractVideoSurface::isActive")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) IsFormatSupported(format QVideoSurfaceFormat_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::isFormatSupported")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsFormatSupported(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) ConnectNativeResolutionChanged(f func(resolution *core.QSize)) {
	defer qt.Recovering("connect QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectNativeResolutionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nativeResolutionChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectNativeResolutionChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectNativeResolutionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nativeResolutionChanged")
	}
}

//export callbackQAbstractVideoSurfaceNativeResolutionChanged
func callbackQAbstractVideoSurfaceNativeResolutionChanged(ptr unsafe.Pointer, ptrName *C.char, resolution unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::nativeResolutionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeResolutionChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(resolution))
	}

}

func (ptr *QAbstractVideoSurface) NativeResolutionChanged(resolution core.QSize_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::nativeResolutionChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_NativeResolutionChanged(ptr.Pointer(), core.PointerFromQSize(resolution))
	}
}

func (ptr *QAbstractVideoSurface) Present(frame QVideoFrame_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::present")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Present(ptr.Pointer(), PointerFromQVideoFrame(frame)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Start(format QVideoSurfaceFormat_ITF) bool {
	defer qt.Recovering("QAbstractVideoSurface::start")

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Start(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) ConnectStop(f func()) {
	defer qt.Recovering("connect QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectStop() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

//export callbackQAbstractVideoSurfaceStop
func callbackQAbstractVideoSurfaceStop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoSurface::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).StopDefault()
	}
}

func (ptr *QAbstractVideoSurface) Stop() {
	defer qt.Recovering("QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_Stop(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) StopDefault() {
	defer qt.Recovering("QAbstractVideoSurface::stop")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_StopDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) ConnectSupportedFormatsChanged(f func()) {
	defer qt.Recovering("connect QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectSupportedFormatsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "supportedFormatsChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectSupportedFormatsChanged() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectSupportedFormatsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "supportedFormatsChanged")
	}
}

//export callbackQAbstractVideoSurfaceSupportedFormatsChanged
func callbackQAbstractVideoSurfaceSupportedFormatsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractVideoSurface::supportedFormatsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "supportedFormatsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractVideoSurface) SupportedFormatsChanged() {
	defer qt.Recovering("QAbstractVideoSurface::supportedFormatsChanged")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_SupportedFormatsChanged(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) DestroyQAbstractVideoSurface() {
	defer qt.Recovering("QAbstractVideoSurface::~QAbstractVideoSurface")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DestroyQAbstractVideoSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractVideoSurface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractVideoSurfaceTimerEvent
func callbackQAbstractVideoSurfaceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractVideoSurfaceChildEvent
func callbackQAbstractVideoSurfaceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractVideoSurfaceCustomEvent
func callbackQAbstractVideoSurfaceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractVideoSurface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractVideoSurfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractVideoSurface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractVideoSurface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractVideoSurface::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
