package widgets

//#include "qgraphicsanchor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsAnchor struct {
	core.QObject
}

type QGraphicsAnchorITF interface {
	core.QObjectITF
	QGraphicsAnchorPTR() *QGraphicsAnchor
}

func PointerFromQGraphicsAnchor(ptr QGraphicsAnchorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsAnchorPTR().Pointer()
	}
	return nil
}

func QGraphicsAnchorFromPointer(ptr unsafe.Pointer) *QGraphicsAnchor {
	var n = new(QGraphicsAnchor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsAnchor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsAnchor) QGraphicsAnchorPTR() *QGraphicsAnchor {
	return ptr
}

func (ptr *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_SetSizePolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	if ptr.Pointer() != nil {
		return QSizePolicy__Policy(C.QGraphicsAnchor_SizePolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsAnchor) UnsetSpacing() {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_UnsetSpacing(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsAnchor) DestroyQGraphicsAnchor() {
	if ptr.Pointer() != nil {
		C.QGraphicsAnchor_DestroyQGraphicsAnchor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
