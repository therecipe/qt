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

type QAbstractVideoFilterITF interface {
	core.QObjectITF
	QAbstractVideoFilterPTR() *QAbstractVideoFilter
}

func PointerFromQAbstractVideoFilter(ptr QAbstractVideoFilterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractVideoFilterPTR().Pointer()
	}
	return nil
}

func QAbstractVideoFilterFromPointer(ptr unsafe.Pointer) *QAbstractVideoFilter {
	var n = new(QAbstractVideoFilter)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractVideoFilter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractVideoFilter) QAbstractVideoFilterPTR() *QAbstractVideoFilter {
	return ptr
}

func (ptr *QAbstractVideoFilter) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractVideoFilter_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractVideoFilter) SetActive(v bool) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_SetActive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractVideoFilter) ConnectActiveChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_ConnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractVideoFilter) DisconnectActiveChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractVideoFilter_DisconnectActiveChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractVideoFilterActiveChanged
func callbackQAbstractVideoFilterActiveChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeChanged").(func())()
}

func (ptr *QAbstractVideoFilter) CreateFilterRunnable() *QVideoFilterRunnable {
	if ptr.Pointer() != nil {
		return QVideoFilterRunnableFromPointer(unsafe.Pointer(C.QAbstractVideoFilter_CreateFilterRunnable(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
