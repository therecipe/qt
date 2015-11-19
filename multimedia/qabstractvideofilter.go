package multimedia

//#include "qabstractvideofilter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractVideoFilter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilter_PTR() *QAbstractVideoFilter {
	return ptr
}

func (ptr *QAbstractVideoFilter) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoFilter_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractVideoFilter) SetActive(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractVideoFilter) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoFilterActiveChanged
func callbackQAbstractVideoFilterActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable {
	if ptr.Pointer() != nil {
		return NewQVideoFilterRunnableFromPointer(C.QAbstractVideoFilter_CreateFilterRunnable(ptr.Pointer()))
	}
	return nil
}
