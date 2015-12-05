package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAbstractVideoFilter struct {
	core.QObject
}

type QAbstractVideoFilter_ITF interface {
	core.QObject_ITF
	QAbstractVideoFilter_PTR() *QAbstractVideoFilter
}

func PointerFromQAbstractVideoFilter(ptr QAbstractVideoFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoFilter_PTR().Pointer()
	}
	return nil
}

func NewQAbstractVideoFilterFromPointer(ptr unsafe.Pointer) *QAbstractVideoFilter {
	var n = new(QAbstractVideoFilter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractVideoFilter_") {
		n.SetObjectName("QAbstractVideoFilter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilter_PTR() *QAbstractVideoFilter {
	return ptr
}

func (ptr *QAbstractVideoFilter) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAbstractVideoFilter_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoFilter) SetActive(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::setActive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractVideoFilter) ConnectActiveChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectActiveChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::activeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoFilterActiveChanged
func callbackQAbstractVideoFilterActiveChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::activeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractVideoFilter::createFilterRunnable")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVideoFilterRunnableFromPointer(C.QAbstractVideoFilter_CreateFilterRunnable(ptr.Pointer()))
	}
	return nil
}
