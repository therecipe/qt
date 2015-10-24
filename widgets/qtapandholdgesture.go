package widgets

//#include "qtapandholdgesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapAndHoldGesture struct {
	QGesture
}

type QTapAndHoldGestureITF interface {
	QGestureITF
	QTapAndHoldGesturePTR() *QTapAndHoldGesture
}

func PointerFromQTapAndHoldGesture(ptr QTapAndHoldGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapAndHoldGesturePTR().Pointer()
	}
	return nil
}

func QTapAndHoldGestureFromPointer(ptr unsafe.Pointer) *QTapAndHoldGesture {
	var n = new(QTapAndHoldGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTapAndHoldGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapAndHoldGesture) QTapAndHoldGesturePTR() *QTapAndHoldGesture {
	return ptr
}

func (ptr *QTapAndHoldGesture) SetPosition(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	C.QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(C.int(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	return int(C.QTapAndHoldGesture_QTapAndHoldGesture_Timeout())
}

func (ptr *QTapAndHoldGesture) DestroyQTapAndHoldGesture() {
	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_DestroyQTapAndHoldGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
