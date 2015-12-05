package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSwipeGesture struct {
	QGesture
}

type QSwipeGesture_ITF interface {
	QGesture_ITF
	QSwipeGesture_PTR() *QSwipeGesture
}

func PointerFromQSwipeGesture(ptr QSwipeGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSwipeGesture_PTR().Pointer()
	}
	return nil
}

func NewQSwipeGestureFromPointer(ptr unsafe.Pointer) *QSwipeGesture {
	var n = new(QSwipeGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSwipeGesture_") {
		n.SetObjectName("QSwipeGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSwipeGesture) QSwipeGesture_PTR() *QSwipeGesture {
	return ptr
}

//QSwipeGesture::SwipeDirection
type QSwipeGesture__SwipeDirection int64

const (
	QSwipeGesture__NoDirection = QSwipeGesture__SwipeDirection(0)
	QSwipeGesture__Left        = QSwipeGesture__SwipeDirection(1)
	QSwipeGesture__Right       = QSwipeGesture__SwipeDirection(2)
	QSwipeGesture__Up          = QSwipeGesture__SwipeDirection(3)
	QSwipeGesture__Down        = QSwipeGesture__SwipeDirection(4)
)

func (ptr *QSwipeGesture) HorizontalDirection() QSwipeGesture__SwipeDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSwipeGesture::horizontalDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return QSwipeGesture__SwipeDirection(C.QSwipeGesture_HorizontalDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSwipeGesture) SetSwipeAngle(value float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSwipeGesture::setSwipeAngle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSwipeGesture_SetSwipeAngle(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QSwipeGesture) SwipeAngle() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSwipeGesture::swipeAngle")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSwipeGesture_SwipeAngle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSwipeGesture) VerticalDirection() QSwipeGesture__SwipeDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSwipeGesture::verticalDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return QSwipeGesture__SwipeDirection(C.QSwipeGesture_VerticalDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSwipeGesture) DestroyQSwipeGesture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSwipeGesture::~QSwipeGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSwipeGesture_DestroyQSwipeGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
