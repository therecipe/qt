package multimedia

//#include "qcameracapturebufferformatcontrol.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QCameraCaptureBufferFormatControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraCaptureBufferFormatControl) QCameraCaptureBufferFormatControl_PTR() *QCameraCaptureBufferFormatControl {
	return ptr
}

func (ptr *QCameraCaptureBufferFormatControl) BufferFormat() QVideoFrame__PixelFormat {
	if ptr.Pointer() != nil {
		return QVideoFrame__PixelFormat(C.QCameraCaptureBufferFormatControl_BufferFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureBufferFormatControl) ConnectBufferFormatChanged(f func(format QVideoFrame__PixelFormat)) {
	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferFormatChanged", f)
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DisconnectBufferFormatChanged() {
	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferFormatChanged")
	}
}

//export callbackQCameraCaptureBufferFormatControlBufferFormatChanged
func callbackQCameraCaptureBufferFormatControlBufferFormatChanged(ptrName *C.char, format C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferFormatChanged").(func(QVideoFrame__PixelFormat))(QVideoFrame__PixelFormat(format))
}

func (ptr *QCameraCaptureBufferFormatControl) SetBufferFormat(format QVideoFrame__PixelFormat) {
	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_SetBufferFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QCameraCaptureBufferFormatControl) DestroyQCameraCaptureBufferFormatControl() {
	if ptr.Pointer() != nil {
		C.QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
