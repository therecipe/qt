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

type QMovieITF interface {
	core.QObjectITF
	QMoviePTR() *QMovie
}

func PointerFromQMovie(ptr QMovieITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMoviePTR().Pointer()
	}
	return nil
}

func QMovieFromPointer(ptr unsafe.Pointer) *QMovie {
	var n = new(QMovie)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMovie_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMovie) QMoviePTR() *QMovie {
	return ptr
}

//QMovie::CacheMode
type QMovie__CacheMode int

var (
	QMovie__CacheNone = QMovie__CacheMode(0)
	QMovie__CacheAll  = QMovie__CacheMode(1)
)

//QMovie::MovieState
type QMovie__MovieState int

var (
	QMovie__NotRunning = QMovie__MovieState(0)
	QMovie__Paused     = QMovie__MovieState(1)
	QMovie__Running    = QMovie__MovieState(2)
)

func (ptr *QMovie) CacheMode() QMovie__CacheMode {
	if ptr.Pointer() != nil {
		return QMovie__CacheMode(C.QMovie_CacheMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMovie) SetCacheMode(mode QMovie__CacheMode) {
	if ptr.Pointer() != nil {
		C.QMovie_SetCacheMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QMovie) SetSpeed(percentSpeed int) {
	if ptr.Pointer() != nil {
		C.QMovie_SetSpeed(C.QtObjectPtr(ptr.Pointer()), C.int(percentSpeed))
	}
}

func (ptr *QMovie) Speed() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_Speed(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQMovie2(device core.QIODeviceITF, format core.QByteArrayITF, parent core.QObjectITF) *QMovie {
	return QMovieFromPointer(unsafe.Pointer(C.QMovie_NewQMovie2(C.QtObjectPtr(core.PointerFromQIODevice(device)), C.QtObjectPtr(core.PointerFromQByteArray(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQMovie(parent core.QObjectITF) *QMovie {
	return QMovieFromPointer(unsafe.Pointer(C.QMovie_NewQMovie(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQMovie3(fileName string, format core.QByteArrayITF, parent core.QObjectITF) *QMovie {
	return QMovieFromPointer(unsafe.Pointer(C.QMovie_NewQMovie3(C.CString(fileName), C.QtObjectPtr(core.PointerFromQByteArray(format)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMovie) CurrentFrameNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_CurrentFrameNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMovie) Device() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QMovie_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMovie) ConnectError(f func(error QImageReader__ImageReaderError)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMovie) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMovieError
func callbackQMovieError(ptrName *C.char, error C.int) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(QImageReader__ImageReaderError))(QImageReader__ImageReaderError(error))
}

func (ptr *QMovie) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMovie_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMovie) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QMovie) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQMovieFinished
func callbackQMovieFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QMovie) ConnectFrameChanged(f func(frameNumber int)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectFrameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QMovie) DisconnectFrameChanged() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFrameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQMovieFrameChanged
func callbackQMovieFrameChanged(ptrName *C.char, frameNumber C.int) {
	qt.GetSignal(C.GoString(ptrName), "frameChanged").(func(int))(int(frameNumber))
}

func (ptr *QMovie) FrameCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_FrameCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMovie) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMovie_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMovie) JumpToFrame(frameNumber int) bool {
	if ptr.Pointer() != nil {
		return C.QMovie_JumpToFrame(C.QtObjectPtr(ptr.Pointer()), C.int(frameNumber)) != 0
	}
	return false
}

func (ptr *QMovie) JumpToNextFrame() bool {
	if ptr.Pointer() != nil {
		return C.QMovie_JumpToNextFrame(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMovie) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_LoopCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMovie) NextFrameDelay() int {
	if ptr.Pointer() != nil {
		return int(C.QMovie_NextFrameDelay(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMovie) SetBackgroundColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetBackgroundColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QMovie) SetDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QMovie) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		C.QMovie_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QMovie) SetFormat(format core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(format)))
	}
}

func (ptr *QMovie) SetPaused(paused bool) {
	if ptr.Pointer() != nil {
		C.QMovie_SetPaused(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QMovie) SetScaledSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QMovie_SetScaledSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QMovie) Start() {
	if ptr.Pointer() != nil {
		C.QMovie_Start(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMovie) ConnectStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QMovie) DisconnectStarted() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQMovieStarted
func callbackQMovieStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QMovie) ConnectStateChanged(f func(state QMovie__MovieState)) {
	if ptr.Pointer() != nil {
		C.QMovie_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMovie) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMovieStateChanged
func callbackQMovieStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMovie__MovieState))(QMovie__MovieState(state))
}

func (ptr *QMovie) Stop() {
	if ptr.Pointer() != nil {
		C.QMovie_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMovie) DestroyQMovie() {
	if ptr.Pointer() != nil {
		C.QMovie_DestroyQMovie(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
