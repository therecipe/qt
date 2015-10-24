package widgets

//#include "qtapgesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapGesture struct {
	QGesture
}

type QTapGestureITF interface {
	QGestureITF
	QTapGesturePTR() *QTapGesture
}

func PointerFromQTapGesture(ptr QTapGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapGesturePTR().Pointer()
	}
	return nil
}

func QTapGestureFromPointer(ptr unsafe.Pointer) *QTapGesture {
	var n = new(QTapGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTapGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapGesture) QTapGesturePTR() *QTapGesture {
	return ptr
}

func (ptr *QTapGesture) SetPosition(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QTapGesture_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func (ptr *QTapGesture) DestroyQTapGesture() {
	if ptr.Pointer() != nil {
		C.QTapGesture_DestroyQTapGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
