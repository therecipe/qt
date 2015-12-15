package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
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
func callbackQCameraCaptureBufferFormatControlBufferFormatChanged(ptrName *C.char, format C.int) {
	defer qt.Recovering("callback QCameraCaptureBufferFormatControl::bufferFormatChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged")
	if signal != nil {
		signal.(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
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
