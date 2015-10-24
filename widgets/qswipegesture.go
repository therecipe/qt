package widgets

//#include "qswipegesture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSwipeGesture struct {
	QGesture
}

type QSwipeGestureITF interface {
	QGestureITF
	QSwipeGesturePTR() *QSwipeGesture
}

func PointerFromQSwipeGesture(ptr QSwipeGestureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSwipeGesturePTR().Pointer()
	}
	return nil
}

func QSwipeGestureFromPointer(ptr unsafe.Pointer) *QSwipeGesture {
	var n = new(QSwipeGesture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSwipeGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSwipeGesture) QSwipeGesturePTR() *QSwipeGesture {
	return ptr
}

//QSwipeGesture::SwipeDirection
type QSwipeGesture__SwipeDirection int

var (
	QSwipeGesture__NoDirection = QSwipeGesture__SwipeDirection(0)
	QSwipeGesture__Left        = QSwipeGesture__SwipeDirection(1)
	QSwipeGesture__Right       = QSwipeGesture__SwipeDirection(2)
	QSwipeGesture__Up          = QSwipeGesture__SwipeDirection(3)
	QSwipeGesture__Down        = QSwipeGesture__SwipeDirection(4)
)

func (ptr *QSwipeGesture) HorizontalDirection() QSwipeGesture__SwipeDirection {
	if ptr.Pointer() != nil {
		return QSwipeGesture__SwipeDirection(C.QSwipeGesture_HorizontalDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSwipeGesture) VerticalDirection() QSwipeGesture__SwipeDirection {
	if ptr.Pointer() != nil {
		return QSwipeGesture__SwipeDirection(C.QSwipeGesture_VerticalDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSwipeGesture) DestroyQSwipeGesture() {
	if ptr.Pointer() != nil {
		C.QSwipeGesture_DestroyQSwipeGesture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
