package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QAbstractVideoSurface_" + qt.RandomIdentifier())
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

func (ptr *QAbstractVideoSurface) ConnectActiveChanged(f func(active bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectActiveChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoSurfaceActiveChanged
func callbackQAbstractVideoSurfaceActiveChanged(ptrName *C.char, active C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::activeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func(bool))(int(active) != 0)
}

func (ptr *QAbstractVideoSurface) Error() QAbstractVideoSurface__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractVideoSurface__Error(C.QAbstractVideoSurface_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractVideoSurface) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) IsFormatSupported(format QVideoSurfaceFormat_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::isFormatSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_IsFormatSupported(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Present(frame QVideoFrame_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::present")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Present(ptr.Pointer(), PointerFromQVideoFrame(frame)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Start(format QVideoSurfaceFormat_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::start")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractVideoSurface_Start(ptr.Pointer(), PointerFromQVideoSurfaceFormat(format)) != 0
	}
	return false
}

func (ptr *QAbstractVideoSurface) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_Stop(ptr.Pointer())
	}
}

func (ptr *QAbstractVideoSurface) ConnectSupportedFormatsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::supportedFormatsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_ConnectSupportedFormatsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "supportedFormatsChanged", f)
	}
}

func (ptr *QAbstractVideoSurface) DisconnectSupportedFormatsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::supportedFormatsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DisconnectSupportedFormatsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "supportedFormatsChanged")
	}
}

//export callbackQAbstractVideoSurfaceSupportedFormatsChanged
func callbackQAbstractVideoSurfaceSupportedFormatsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::supportedFormatsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "supportedFormatsChanged").(func())()
}

func (ptr *QAbstractVideoSurface) DestroyQAbstractVideoSurface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoSurface::~QAbstractVideoSurface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoSurface_DestroyQAbstractVideoSurface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
