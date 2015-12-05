package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsAnchor struct {
	core.QObject
}

type QGraphicsAnchor_ITF interface {
	core.QObject_ITF
	QGraphicsAnchor_PTR() *QGraphicsAnchor
}

func PointerFromQGraphicsAnchor(ptr QGraphicsAnchor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchor_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsAnchorFromPointer(ptr unsafe.Pointer) *QGraphicsAnchor {
	var n = new(QGraphicsAnchor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsAnchor_") {
		n.SetObjectName("QGraphicsAnchor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsAnchor) QGraphicsAnchor_PTR() *QGraphicsAnchor {
	return ptr
}

func (ptr *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::setSizePolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_SetSizePolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGraphicsAnchor) SetSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::sizePolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QSizePolicy__Policy(C.QGraphicsAnchor_SizePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchor) Spacing() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::spacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsAnchor_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsAnchor) UnsetSpacing() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::unsetSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_UnsetSpacing(ptr.Pointer())
	}
}

func (ptr *QGraphicsAnchor) DestroyQGraphicsAnchor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsAnchor::~QGraphicsAnchor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_DestroyQGraphicsAnchor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
