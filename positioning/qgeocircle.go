package positioning

//#include "qgeocircle.h"
import "C"
import (
	"unsafe"
)

type QGeoCircle struct {
	QGeoShape
}

type QGeoCircleITF interface {
	QGeoShapeITF
	QGeoCirclePTR() *QGeoCircle
}

func PointerFromQGeoCircle(ptr QGeoCircleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCirclePTR().Pointer()
	}
	return nil
}

func QGeoCircleFromPointer(ptr unsafe.Pointer) *QGeoCircle {
	var n = new(QGeoCircle)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoCircle) QGeoCirclePTR() *QGeoCircle {
	return ptr
}

func NewQGeoCircle() *QGeoCircle {
	return QGeoCircleFromPointer(unsafe.Pointer(C.QGeoCircle_NewQGeoCircle()))
}

func NewQGeoCircle3(other QGeoCircleITF) *QGeoCircle {
	return QGeoCircleFromPointer(unsafe.Pointer(C.QGeoCircle_NewQGeoCircle3(C.QtObjectPtr(PointerFromQGeoCircle(other)))))
}

func NewQGeoCircle4(other QGeoShapeITF) *QGeoCircle {
	return QGeoCircleFromPointer(unsafe.Pointer(C.QGeoCircle_NewQGeoCircle4(C.QtObjectPtr(PointerFromQGeoShape(other)))))
}

func (ptr *QGeoCircle) SetCenter(center QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoCircle_SetCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(center)))
	}
}

func (ptr *QGeoCircle) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCircle_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	if ptr.Pointer() != nil {
		C.QGeoCircle_DestroyQGeoCircle(C.QtObjectPtr(ptr.Pointer()))
	}
}
