package multimedia

//#include "qabstractvideosurface.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractVideoSurface struct {
	core.QObject
}

type QAbstractVideoSurfaceITF interface {
	core.QObjectITF
	QAbstractVideoSurfacePTR() *QAbstractVideoSurface
}

func PointerFromQAbstractVideoSurface(ptr QAbstractVideoSurfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoSurfacePTR().Pointer()
	}
	return nil
}

func QAbstractVideoSurfaceFromPointer(ptr unsafe.Pointer) *QAbstractVideoSurface {
	var n = new(QAbstractVideoSurface)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractVideoSurface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractVideoSurface) QAbstractVideoSurfacePTR() *QAbstractVideoSurface {
	return ptr
}

//QAbstractVideoSurface::Error
type QAbstractVideoSurface__Error int

var (
	QAbstractVideoSurface__NoError                = QAbstractVideoSurface__Error(0)
	QAbstractVideoSurface__UnsupportedFormatError = QAbstractVideoSurface__Error(1)
	QAbstractVideoSurface__IncorrectFormatError   = QAbstractVideoSurface__Error(2)
	QAbstractVideoSurface__StoppedError           = QAbstractVideoSurface__Error(3)
	QAbstractVideoSurface__ResourceError          = QAbstractVideoSurface__Error(4)
)

func (ptr *QAbstractVideoSurface) ConnectActiveChanged(f func(active bool)) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoSurfaceActiveChanged
func callbackQAbstractVideoSurfaceActiveChanged(ptrName *C.char, active C.int) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func(bool))(int(active) != 0)
}

func (ptr *QAbstractVideoSurface) Error() QAbstractVideoSurface__Error {
	if ptr.Pointer() != nil {
		return QAbstractVideoSurface__Error(C.QAbstractVideoSurface_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractVideoSurface) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) IsFormatSupported(format QVideoSurfaceFormatITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsFormatSupported(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVideoSurfaceFormat(format))) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Present(frame QVideoFrameITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Present(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVideoFrame(frame))) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Start(format QVideoSurfaceFormatITF) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Start(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQVideoSurfaceFormat(format))) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Stop() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractVideoSurface) ConnectSupportedFormatsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectSupportedFormatsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "supportedFormatsChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectSupportedFormatsChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectSupportedFormatsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "supportedFormatsChanged")
	}
}

//export callbackQAbstractVideoSurfaceSupportedFormatsChanged
func callbackQAbstractVideoSurfaceSupportedFormatsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "supportedFormatsChanged").(func())()
}

func (ptr *QAbstractVideoSurface) DestroyQAbstractVideoSurface() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DestroyQAbstractVideoSurface(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
