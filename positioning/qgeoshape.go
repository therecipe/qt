package positioning

//#include "qgeoshape.h"
import "C"
import (
	"unsafe"
)

type QGeoShape struct {
	ptr unsafe.Pointer
}

type QGeoShapeITF interface {
	QGeoShapePTR() *QGeoShape
}

func (p *QGeoShape) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoShape) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoShape(ptr QGeoShapeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShapePTR().Pointer()
	}
	return nil
}

func QGeoShapeFromPointer(ptr unsafe.Pointer) *QGeoShape {
	var n = new(QGeoShape)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoShape) QGeoShapePTR() *QGeoShape {
	return ptr
}

//QGeoShape::ShapeType
type QGeoShape__ShapeType int

var (
	QGeoShape__UnknownType   = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    = QGeoShape__ShapeType(2)
)

func NewQGeoShape() *QGeoShape {
	return QGeoShapeFromPointer(unsafe.Pointer(C.QGeoShape_NewQGeoShape()))
}

func NewQGeoShape2(other QGeoShapeITF) *QGeoShape {
	return QGeoShapeFromPointer(unsafe.Pointer(C.QGeoShape_NewQGeoShape2(C.QtObjectPtr(PointerFromQGeoShape(other)))))
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinateITF) bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(coordinate))) != 0
	}
	return false
}

func (ptr *QGeoShape) ExtendShape(coordinate QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoShape_ExtendShape(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(coordinate)))
	}
}

func (ptr *QGeoShape) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoShape) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoShape_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoShape) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoShape_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGeoShape) Type() QGeoShape__ShapeType {
	if ptr.Pointer() != nil {
		return QGeoShape__ShapeType(C.QGeoShape_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoShape) DestroyQGeoShape() {
	if ptr.Pointer() != nil {
		C.QGeoShape_DestroyQGeoShape(C.QtObjectPtr(ptr.Pointer()))
	}
}
