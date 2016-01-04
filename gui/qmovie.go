package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMovie struct {
	core.QObject
}

type QMovie_ITF interface {
	core.QObject_ITF
	QMovie_PTR() *QMovie
}

func PointerFromQMovie(ptr QMovie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMovie_PTR().Pointer()
	}
	return nil
}

func NewQMovieFromPointer(ptr unsafe.Pointer) *QMovie {
	var n = new(QMovie)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMovie_") {
		n.SetObjectName("QMovie_" + qt.Identifier())
	}
	return n
}

func (ptr *QMovie) QMovie_PTR() *QMovie {
	return ptr
}

//QMovie::CacheMode
type QMovie__CacheMode int64

const (
	QMovie__CacheNone = QMovie__CacheMode(0)
	QMovie__CacheAll  = QMovie__CacheMode(1)
)

//QMovie::MovieState
type QMovie__MovieState int64

const (
	QMovie__NotRunning = QMovie__MovieState(0)
	QMovie__Paused     = QMovie__MovieState(1)
	QMovie__Running    = QMovie__MovieState(2)
)

func (ptr *QMovie) CacheMode() QMovie__CacheMode {
	defer qt.Recovering("QMovie::cacheMode")

	if ptr.Pointer() != nil {
		return QMovie__CacheMode(C.QMovie_CacheMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) SetCacheMode(mode QMovie__CacheMode) {
	defer qt.Recovering("QMovie::setCacheMode")

	if ptr.Pointer() != nil {
		C.QMovie_SetCacheMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMovie) SetSpeed(percentSpeed int) {
	defer qt.Recovering("QMovie::setSpeed")

	if ptr.Pointer() != nil {
		C.QMovie_SetSpeed(ptr.Pointer(), C.int(percentSpeed))
	}
}

func (ptr *QMovie) Speed() int {
	defer qt.Recovering("QMovie::speed")

	if ptr.Pointer() != nil {
		return int(C.QMovie_Speed(ptr.Pointer()))
	}
	return 0
}

func NewQMovie2(device core.QIODevice_ITF, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	defer qt.Recovering("QMovie::QMovie")

	return NewQMovieFromPointer(C.QMovie_NewQMovie2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func NewQMovie(parent core.QObject_ITF) *QMovie {
	defer qt.Recovering("QMovie::QMovie")

	return NewQMovieFromPointer(C.QMovie_NewQMovie(core.PointerFromQObject(parent)))
}

func NewQMovie3(fileName string, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	defer qt.Recovering("QMovie::QMovie")

	return NewQMovieFromPointer(C.QMovie_NewQMovie3(C.CString(fileName), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func (ptr *QMovie) BackgroundColor() *QColor {
	defer qt.Recovering("QMovie::backgroundColor")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QMovie_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) CurrentFrameNumber() int {
	defer qt.Recovering("QMovie::currentFrameNumber")

	if ptr.Pointer() != nil {
		return int(C.QMovie_CurrentFrameNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) Device() *core.QIODevice {
	defer qt.Recovering("QMovie::device")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMovie_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectError(f func(error QImageReader__ImageReaderError)) {
	defer qt.Recovering("connect QMovie::error")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMovie) DisconnectError() {
	defer qt.Recovering("disconnect QMovie::error")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMovieError
func callbackQMovieError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QMovie::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QImageReader__ImageReaderError))(QImageReader__ImageReaderError(error))
	}

}

func (ptr *QMovie) Error(error QImageReader__ImageReaderError) {
	defer qt.Recovering("QMovie::error")

	if ptr.Pointer() != nil {
		C.QMovie_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QMovie) FileName() string {
	defer qt.Recovering("QMovie::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMovie_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMovie) ConnectFinished(f func()) {
	defer qt.Recovering("connect QMovie::finished")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QMovie) DisconnectFinished() {
	defer qt.Recovering("disconnect QMovie::finished")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQMovieFinished
func callbackQMovieFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMovie::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMovie) Finished() {
	defer qt.Recovering("QMovie::finished")

	if ptr.Pointer() != nil {
		C.QMovie_Finished(ptr.Pointer())
	}
}

func (ptr *QMovie) Format() *core.QByteArray {
	defer qt.Recovering("QMovie::format")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QMovie_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectFrameChanged(f func(frameNumber int)) {
	defer qt.Recovering("connect QMovie::frameChanged")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectFrameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QMovie) DisconnectFrameChanged() {
	defer qt.Recovering("disconnect QMovie::frameChanged")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFrameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQMovieFrameChanged
func callbackQMovieFrameChanged(ptr unsafe.Pointer, ptrName *C.char, frameNumber C.int) {
	defer qt.Recovering("callback QMovie::frameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "frameChanged"); signal != nil {
		signal.(func(int))(int(frameNumber))
	}

}

func (ptr *QMovie) FrameChanged(frameNumber int) {
	defer qt.Recovering("QMovie::frameChanged")

	if ptr.Pointer() != nil {
		C.QMovie_FrameChanged(ptr.Pointer(), C.int(frameNumber))
	}
}

func (ptr *QMovie) FrameCount() int {
	defer qt.Recovering("QMovie::frameCount")

	if ptr.Pointer() != nil {
		return int(C.QMovie_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) FrameRect() *core.QRect {
	defer qt.Recovering("QMovie::frameRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QMovie_FrameRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) IsValid() bool {
	defer qt.Recovering("QMovie::isValid")

	if ptr.Pointer() != nil {
		return C.QMovie_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) JumpToFrame(frameNumber int) bool {
	defer qt.Recovering("QMovie::jumpToFrame")

	if ptr.Pointer() != nil {
		return C.QMovie_JumpToFrame(ptr.Pointer(), C.int(frameNumber)) != 0
	}
	return false
}

func (ptr *QMovie) JumpToNextFrame() bool {
	defer qt.Recovering("QMovie::jumpToNextFrame")

	if ptr.Pointer() != nil {
		return C.QMovie_JumpToNextFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) LoopCount() int {
	defer qt.Recovering("QMovie::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QMovie_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) NextFrameDelay() int {
	defer qt.Recovering("QMovie::nextFrameDelay")

	if ptr.Pointer() != nil {
		return int(C.QMovie_NextFrameDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) ConnectResized(f func(size *core.QSize)) {
	defer qt.Recovering("connect QMovie::resized")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "resized", f)
	}
}

func (ptr *QMovie) DisconnectResized() {
	defer qt.Recovering("disconnect QMovie::resized")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "resized")
	}
}

//export callbackQMovieResized
func callbackQMovieResized(ptr unsafe.Pointer, ptrName *C.char, size unsafe.Pointer) {
	defer qt.Recovering("callback QMovie::resized")

	if signal := qt.GetSignal(C.GoString(ptrName), "resized"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

func (ptr *QMovie) Resized(size core.QSize_ITF) {
	defer qt.Recovering("QMovie::resized")

	if ptr.Pointer() != nil {
		C.QMovie_Resized(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QMovie) ScaledSize() *core.QSize {
	defer qt.Recovering("QMovie::scaledSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMovie_ScaledSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) SetBackgroundColor(color QColor_ITF) {
	defer qt.Recovering("QMovie::setBackgroundColor")

	if ptr.Pointer() != nil {
		C.QMovie_SetBackgroundColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QMovie) SetDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QMovie::setDevice")

	if ptr.Pointer() != nil {
		C.QMovie_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QMovie) SetFileName(fileName string) {
	defer qt.Recovering("QMovie::setFileName")

	if ptr.Pointer() != nil {
		C.QMovie_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QMovie) SetFormat(format core.QByteArray_ITF) {
	defer qt.Recovering("QMovie::setFormat")

	if ptr.Pointer() != nil {
		C.QMovie_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QMovie) SetPaused(paused bool) {
	defer qt.Recovering("QMovie::setPaused")

	if ptr.Pointer() != nil {
		C.QMovie_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QMovie) SetScaledSize(size core.QSize_ITF) {
	defer qt.Recovering("QMovie::setScaledSize")

	if ptr.Pointer() != nil {
		C.QMovie_SetScaledSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QMovie) Start() {
	defer qt.Recovering("QMovie::start")

	if ptr.Pointer() != nil {
		C.QMovie_Start(ptr.Pointer())
	}
}

func (ptr *QMovie) ConnectStarted(f func()) {
	defer qt.Recovering("connect QMovie::started")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QMovie) DisconnectStarted() {
	defer qt.Recovering("disconnect QMovie::started")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQMovieStarted
func callbackQMovieStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMovie::started")

	if signal := qt.GetSignal(C.GoString(ptrName), "started"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMovie) Started() {
	defer qt.Recovering("QMovie::started")

	if ptr.Pointer() != nil {
		C.QMovie_Started(ptr.Pointer())
	}
}

func (ptr *QMovie) State() QMovie__MovieState {
	defer qt.Recovering("QMovie::state")

	if ptr.Pointer() != nil {
		return QMovie__MovieState(C.QMovie_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) ConnectStateChanged(f func(state QMovie__MovieState)) {
	defer qt.Recovering("connect QMovie::stateChanged")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMovie) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMovie::stateChanged")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMovieStateChanged
func callbackQMovieStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMovie::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QMovie__MovieState))(QMovie__MovieState(state))
	}

}

func (ptr *QMovie) StateChanged(state QMovie__MovieState) {
	defer qt.Recovering("QMovie::stateChanged")

	if ptr.Pointer() != nil {
		C.QMovie_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QMovie) Stop() {
	defer qt.Recovering("QMovie::stop")

	if ptr.Pointer() != nil {
		C.QMovie_Stop(ptr.Pointer())
	}
}

func (ptr *QMovie) ConnectUpdated(f func(rect *core.QRect)) {
	defer qt.Recovering("connect QMovie::updated")

	if ptr.Pointer() != nil {
		C.QMovie_ConnectUpdated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updated", f)
	}
}

func (ptr *QMovie) DisconnectUpdated() {
	defer qt.Recovering("disconnect QMovie::updated")

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updated")
	}
}

//export callbackQMovieUpdated
func callbackQMovieUpdated(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer) {
	defer qt.Recovering("callback QMovie::updated")

	if signal := qt.GetSignal(C.GoString(ptrName), "updated"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(rect))
	}

}

func (ptr *QMovie) Updated(rect core.QRect_ITF) {
	defer qt.Recovering("QMovie::updated")

	if ptr.Pointer() != nil {
		C.QMovie_Updated(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QMovie) DestroyQMovie() {
	defer qt.Recovering("QMovie::~QMovie")

	if ptr.Pointer() != nil {
		C.QMovie_DestroyQMovie(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMovie) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMovie::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMovie) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMovie::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMovieTimerEvent
func callbackQMovieTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMovie::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMovieFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMovie) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMovie::timerEvent")

	if ptr.Pointer() != nil {
		C.QMovie_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMovie) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMovie::timerEvent")

	if ptr.Pointer() != nil {
		C.QMovie_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMovie) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMovie::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMovie) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMovie::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMovieChildEvent
func callbackQMovieChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMovie::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMovieFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMovie) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMovie::childEvent")

	if ptr.Pointer() != nil {
		C.QMovie_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMovie) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMovie::childEvent")

	if ptr.Pointer() != nil {
		C.QMovie_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMovie) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMovie::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMovie) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMovie::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMovieCustomEvent
func callbackQMovieCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMovie::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMovieFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMovie) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMovie::customEvent")

	if ptr.Pointer() != nil {
		C.QMovie_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMovie) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMovie::customEvent")

	if ptr.Pointer() != nil {
		C.QMovie_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
