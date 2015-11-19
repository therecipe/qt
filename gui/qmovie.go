package gui

//#include "qmovie.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QMovie_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return QMovie__CacheMode(C.QMovie_CacheMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) SetCacheMode(mode QMovie__CacheMode) {
	if ptr.Pointer() != nil {
		C.QMovie_SetCacheMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMovie) SetSpeed(percentSpeed int) {
	if ptr.Pointer() != nil {
		C.QMovie_SetSpeed(ptr.Pointer(), C.int(percentSpeed))
	}
}

func (ptr *QMovie) Speed() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_Speed(ptr.Pointer()))
	}
	return 0
}

func NewQMovie2(device core.QIODevice_ITF, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	return NewQMovieFromPointer(C.QMovie_NewQMovie2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func NewQMovie(parent core.QObject_ITF) *QMovie {
	return NewQMovieFromPointer(C.QMovie_NewQMovie(core.PointerFromQObject(parent)))
}

func NewQMovie3(fileName string, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	return NewQMovieFromPointer(C.QMovie_NewQMovie3(C.CString(fileName), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func (ptr *QMovie) BackgroundColor() *QColor {
	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QMovie_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) CurrentFrameNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_CurrentFrameNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMovie_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectError(f func(error QImageReader__ImageReaderError)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMovie) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMovieError
func callbackQMovieError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QImageReader__ImageReaderError))(QImageReader__ImageReaderError(error))
}

func (ptr *QMovie) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMovie_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMovie) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QMovie) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQMovieFinished
func callbackQMovieFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QMovie) Format() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QMovie_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectFrameChanged(f func(frameNumber int)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectFrameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QMovie) DisconnectFrameChanged() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFrameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQMovieFrameChanged
func callbackQMovieFrameChanged(ptrName *C.char, frameNumber C.int) {
	qt.GetSignal(C.GoString(ptrName), "frameChanged").(func(int))(int(frameNumber))
}

func (ptr *QMovie) FrameCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMovie_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) JumpToFrame(frameNumber int) bool {
	if ptr.Pointer() != nil {
		return C.QMovie_JumpToFrame(ptr.Pointer(), C.int(frameNumber)) != 0
	}
	return false
}

func (ptr *QMovie) JumpToNextFrame() bool {
	if ptr.Pointer() != nil {
		return C.QMovie_JumpToNextFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) NextFrameDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_NextFrameDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) SetBackgroundColor(color QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetBackgroundColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QMovie) SetDevice(device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QMovie) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QMovie_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QMovie) SetFormat(format core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QMovie) SetPaused(paused bool) {
	if ptr.Pointer() != nil {
		C.QMovie_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QMovie) SetScaledSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetScaledSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QMovie) Start() {
	if ptr.Pointer() != nil {
		C.QMovie_Start(ptr.Pointer())
	}
}

func (ptr *QMovie) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QMovie) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQMovieStarted
func callbackQMovieStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QMovie) ConnectStateChanged(f func(state QMovie__MovieState)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMovie) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMovieStateChanged
func callbackQMovieStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMovie__MovieState))(QMovie__MovieState(state))
}

func (ptr *QMovie) Stop() {
	if ptr.Pointer() != nil {
		C.QMovie_Stop(ptr.Pointer())
	}
}

func (ptr *QMovie) DestroyQMovie() {
	if ptr.Pointer() != nil {
		C.QMovie_DestroyQMovie(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
