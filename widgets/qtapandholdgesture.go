package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapAndHoldGesture struct {
	QGesture
}

type QTapAndHoldGesture_ITF interface {
	QGesture_ITF
	QTapAndHoldGesture_PTR() *QTapAndHoldGesture
}

func PointerFromQTapAndHoldGesture(ptr QTapAndHoldGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapAndHoldGesture_PTR().Pointer()
	}
	return nil
}

func NewQTapAndHoldGestureFromPointer(ptr unsafe.Pointer) *QTapAndHoldGesture {
	var n = new(QTapAndHoldGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapAndHoldGesture_") {
		n.SetObjectName("QTapAndHoldGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapAndHoldGesture) QTapAndHoldGesture_PTR() *QTapAndHoldGesture {
	return ptr
}

func (ptr *QTapAndHoldGesture) SetPosition(pos core.QPointF_ITF) {
	defer qt.Recovering("QTapAndHoldGesture::setPosition")

	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_SetPosition(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	defer qt.Recovering("QTapAndHoldGesture::setTimeout")

	C.QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(C.int(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	defer qt.Recovering("QTapAndHoldGesture::timeout")

	return int(C.QTapAndHoldGesture_QTapAndHoldGesture_Timeout())
}

func (ptr *QTapAndHoldGesture) DestroyQTapAndHoldGesture() {
	defer qt.Recovering("QTapAndHoldGesture::~QTapAndHoldGesture")

	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_DestroyQTapAndHoldGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
