package widgets

//#include "qpangesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPanGesture struct {
	QGesture
}

type QPanGestureITF interface {
	QGestureITF
	QPanGesturePTR() *QPanGesture
}

func PointerFromQPanGesture(ptr QPanGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPanGesturePTR().Pointer()
	}
	return nil
}

func QPanGestureFromPointer(ptr unsafe.Pointer) *QPanGesture {
	var n = new(QPanGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPanGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPanGesture) QPanGesturePTR() *QPanGesture {
	return ptr
}

func (ptr *QPanGesture) SetLastOffset(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPanGesture_SetLastOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QPanGesture) SetOffset(value core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPanGesture_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(value)))
	}
}

func (ptr *QPanGesture) DestroyQPanGesture() {
	if ptr.Pointer() != nil {
		C.QPanGesture_DestroyQPanGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
