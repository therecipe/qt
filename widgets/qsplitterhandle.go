package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSplitterHandle struct {
	QWidget
}

type QSplitterHandle_ITF interface {
	QWidget_ITF
	QSplitterHandle_PTR() *QSplitterHandle
}

func PointerFromQSplitterHandle(ptr QSplitterHandle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitterHandle_PTR().Pointer()
	}
	return nil
}

func NewQSplitterHandleFromPointer(ptr unsafe.Pointer) *QSplitterHandle {
	var n = new(QSplitterHandle)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitterHandle_") {
		n.SetObjectName("QSplitterHandle_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSplitterHandle) QSplitterHandle_PTR() *QSplitterHandle {
	return ptr
}

func NewQSplitterHandle(orientation core.Qt__Orientation, parent QSplitter_ITF) *QSplitterHandle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::QSplitterHandle")
		}
	}()

	return NewQSplitterHandleFromPointer(C.QSplitterHandle_NewQSplitterHandle(C.int(orientation), PointerFromQSplitter(parent)))
}

func (ptr *QSplitterHandle) OpaqueResize() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::opaqueResize")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSplitterHandle_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitterHandle) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitterHandle_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitterHandle) SetOrientation(orientation core.Qt__Orientation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::setOrientation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QSplitterHandle) Splitter() *QSplitter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::splitter")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSplitterFromPointer(C.QSplitterHandle_Splitter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitterHandle) DestroyQSplitterHandle() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSplitterHandle::~QSplitterHandle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DestroyQSplitterHandle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
