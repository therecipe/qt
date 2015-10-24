package positioning

//#include "qgeorectangle.h"
import "C"
import (
	"unsafe"
)

type QGeoRectangle struct {
	QGeoShape
}

type QGeoRectangleITF interface {
	QGeoShapeITF
	QGeoRectanglePTR() *QGeoRectangle
}

func PointerFromQGeoRectangle(ptr QGeoRectangleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRectanglePTR().Pointer()
	}
	return nil
}

func QGeoRectangleFromPointer(ptr unsafe.Pointer) *QGeoRectangle {
	var n = new(QGeoRectangle)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRectangle) QGeoRectanglePTR() *QGeoRectangle {
	return ptr
}

func NewQGeoRectangle() *QGeoRectangle {
	return QGeoRectangleFromPointer(unsafe.Pointer(C.QGeoRectangle_NewQGeoRectangle()))
}

func NewQGeoRectangle3(topLeft QGeoCoordinateITF, bottomRight QGeoCoordinateITF) *QGeoRectangle {
	return QGeoRectangleFromPointer(unsafe.Pointer(C.QGeoRectangle_NewQGeoRectangle3(C.QtObjectPtr(PointerFromQGeoCoordinate(topLeft)), C.QtObjectPtr(PointerFromQGeoCoordinate(bottomRight)))))
}

func NewQGeoRectangle5(other QGeoRectangleITF) *QGeoRectangle {
	return QGeoRectangleFromPointer(unsafe.Pointer(C.QGeoRectangle_NewQGeoRectangle5(C.QtObjectPtr(PointerFromQGeoRectangle(other)))))
}

func NewQGeoRectangle6(other QGeoShapeITF) *QGeoRectangle {
	return QGeoRectangleFromPointer(unsafe.Pointer(C.QGeoRectangle_NewQGeoRectangle6(C.QtObjectPtr(PointerFromQGeoShape(other)))))
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangleITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRectangle(rectangle))) != 0
	}
	return false
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangleITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoRectangle(rectangle))) != 0
	}
	return false
}

func (ptr *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(bottomLeft)))
	}
}

func (ptr *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(bottomRight)))
	}
}

func (ptr *QGeoRectangle) SetCenter(center QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(center)))
	}
}

func (ptr *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(topLeft)))
	}
}

func (ptr *QGeoRectangle) SetTopRight(topRight QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(topRight)))
	}
}

func (ptr *QGeoRectangle) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRectangle_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	if ptr.Pointer() != nil {
		C.QGeoRectangle_DestroyQGeoRectangle(C.QtObjectPtr(ptr.Pointer()))
	}
}
