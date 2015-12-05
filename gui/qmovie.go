package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::cacheMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QMovie__CacheMode(C.QMovie_CacheMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) SetCacheMode(mode QMovie__CacheMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setCacheMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetCacheMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMovie) SetSpeed(percentSpeed int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setSpeed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetSpeed(ptr.Pointer(), C.int(percentSpeed))
	}
}

func (ptr *QMovie) Speed() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::speed")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMovie_Speed(ptr.Pointer()))
	}
	return 0
}

func NewQMovie2(device core.QIODevice_ITF, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::QMovie")
		}
	}()

	return NewQMovieFromPointer(C.QMovie_NewQMovie2(core.PointerFromQIODevice(device), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func NewQMovie(parent core.QObject_ITF) *QMovie {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::QMovie")
		}
	}()

	return NewQMovieFromPointer(C.QMovie_NewQMovie(core.PointerFromQObject(parent)))
}

func NewQMovie3(fileName string, format core.QByteArray_ITF, parent core.QObject_ITF) *QMovie {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::QMovie")
		}
	}()

	return NewQMovieFromPointer(C.QMovie_NewQMovie3(C.CString(fileName), core.PointerFromQByteArray(format), core.PointerFromQObject(parent)))
}

func (ptr *QMovie) BackgroundColor() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::backgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QMovie_BackgroundColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) CurrentFrameNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::currentFrameNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMovie_CurrentFrameNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) Device() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::device")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMovie_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectError(f func(error QImageReader__ImageReaderError)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMovie) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMovieError
func callbackQMovieError(ptrName *C.char, error C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(QImageReader__ImageReaderError))(QImageReader__ImageReaderError(error))
}

func (ptr *QMovie) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMovie_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMovie) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QMovie) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQMovieFinished
func callbackQMovieFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QMovie) Format() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::format")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QMovie_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMovie) ConnectFrameChanged(f func(frameNumber int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::frameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_ConnectFrameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameChanged", f)
	}
}

func (ptr *QMovie) DisconnectFrameChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::frameChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectFrameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameChanged")
	}
}

//export callbackQMovieFrameChanged
func callbackQMovieFrameChanged(ptrName *C.char, frameNumber C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::frameChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "frameChanged").(func(int))(int(frameNumber))
}

func (ptr *QMovie) FrameCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::frameCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMovie_FrameCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMovie_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) JumpToFrame(frameNumber int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::jumpToFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMovie_JumpToFrame(ptr.Pointer(), C.int(frameNumber)) != 0
	}
	return false
}

func (ptr *QMovie) JumpToNextFrame() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::jumpToNextFrame")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMovie_JumpToNextFrame(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMovie) LoopCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::loopCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMovie_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) NextFrameDelay() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::nextFrameDelay")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMovie_NextFrameDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMovie) SetBackgroundColor(color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setBackgroundColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetBackgroundColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QMovie) SetDevice(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QMovie) SetFileName(fileName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setFileName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QMovie) SetFormat(format core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetFormat(ptr.Pointer(), core.PointerFromQByteArray(format))
	}
}

func (ptr *QMovie) SetPaused(paused bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setPaused")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetPaused(ptr.Pointer(), C.int(qt.GoBoolToInt(paused)))
	}
}

func (ptr *QMovie) SetScaledSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::setScaledSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_SetScaledSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QMovie) Start() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_Start(ptr.Pointer())
	}
}

func (ptr *QMovie) ConnectStarted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_ConnectStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "started", f)
	}
}

func (ptr *QMovie) DisconnectStarted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::started")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "started")
	}
}

//export callbackQMovieStarted
func callbackQMovieStarted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::started")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "started").(func())()
}

func (ptr *QMovie) ConnectStateChanged(f func(state QMovie__MovieState)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMovie) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMovieStateChanged
func callbackQMovieStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMovie__MovieState))(QMovie__MovieState(state))
}

func (ptr *QMovie) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_Stop(ptr.Pointer())
	}
}

func (ptr *QMovie) DestroyQMovie() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMovie::~QMovie")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMovie_DestroyQMovie(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
