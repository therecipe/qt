package widgets

//#include "qsplitterhandle.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSplitterHandle struct {
	QWidget
}

type QSplitterHandleITF interface {
	QWidgetITF
	QSplitterHandlePTR() *QSplitterHandle
}

func PointerFromQSplitterHandle(ptr QSplitterHandleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitterHandlePTR().Pointer()
	}
	return nil
}

func QSplitterHandleFromPointer(ptr unsafe.Pointer) *QSplitterHandle {
	var n = new(QSplitterHandle)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSplitterHandle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplitterHandle) QSplitterHandlePTR() *QSplitterHandle {
	return ptr
}

func NewQSplitterHandle(orientation core.Qt__Orientation, parent QSplitterITF) *QSplitterHandle {
	return QSplitterHandleFromPointer(unsafe.Pointer(C.QSplitterHandle_NewQSplitterHandle(C.int(orientation), C.QtObjectPtr(PointerFromQSplitter(parent)))))
}

func (ptr *QSplitterHandle) OpaqueResize() bool {
	if ptr.Pointer() != nil {
		return C.QSplitterHandle_OpaqueResize(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSplitterHandle) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitterHandle_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSplitterHandle) SetOrientation(orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QSplitterHandle) Splitter() *QSplitter {
	if ptr.Pointer() != nil {
		return QSplitterFromPointer(unsafe.Pointer(C.QSplitterHandle_Splitter(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSplitterHandle) DestroyQSplitterHandle() {
	if ptr.Pointer() != nil {
		C.QSplitterHandle_DestroyQSplitterHandle(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
