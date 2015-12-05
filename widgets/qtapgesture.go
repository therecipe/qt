package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTapGesture struct {
	QGesture
}

type QTapGesture_ITF interface {
	QGesture_ITF
	QTapGesture_PTR() *QTapGesture
}

func PointerFromQTapGesture(ptr QTapGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapGesture_PTR().Pointer()
	}
	return nil
}

func NewQTapGestureFromPointer(ptr unsafe.Pointer) *QTapGesture {
	var n = new(QTapGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapGesture_") {
		n.SetObjectName("QTapGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapGesture) QTapGesture_PTR() *QTapGesture {
	return ptr
}

func (ptr *QTapGesture) SetPosition(pos core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapGesture::setPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapGesture_SetPosition(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QTapGesture) DestroyQTapGesture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapGesture::~QTapGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapGesture_DestroyQTapGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
