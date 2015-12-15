package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGeoShape struct {
	ptr unsafe.Pointer
}

type QGeoShape_ITF interface {
	QGeoShape_PTR() *QGeoShape
}

func (p *QGeoShape) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoShape) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoShape(ptr QGeoShape_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoShape_PTR().Pointer()
	}
	return nil
}

func NewQGeoShapeFromPointer(ptr unsafe.Pointer) *QGeoShape {
	var n = new(QGeoShape)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoShape) QGeoShape_PTR() *QGeoShape {
	return ptr
}

//QGeoShape::ShapeType
type QGeoShape__ShapeType int64

const (
	QGeoShape__UnknownType   = QGeoShape__ShapeType(0)
	QGeoShape__RectangleType = QGeoShape__ShapeType(1)
	QGeoShape__CircleType    = QGeoShape__ShapeType(2)
)

func NewQGeoShape() *QGeoShape {
	defer qt.Recovering("QGeoShape::QGeoShape")

	return NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape())
}

func NewQGeoShape2(other QGeoShape_ITF) *QGeoShape {
	defer qt.Recovering("QGeoShape::QGeoShape")

	return NewQGeoShapeFromPointer(C.QGeoShape_NewQGeoShape2(PointerFromQGeoShape(other)))
}

func (ptr *QGeoShape) Contains(coordinate QGeoCoordinate_ITF) bool {
	defer qt.Recovering("QGeoShape::contains")

	if ptr.Pointer() != nil {
		return C.QGeoShape_Contains(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate)) != 0
	}
	return false
}

func (ptr *QGeoShape) ExtendShape(coordinate QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoShape::extendShape")

	if ptr.Pointer() != nil {
		C.QGeoShape_ExtendShape(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoShape) IsEmpty() bool {
	defer qt.Recovering("QGeoShape::isEmpty")

	if ptr.Pointer() != nil {
		return C.QGeoShape_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoShape) IsValid() bool {
	defer qt.Recovering("QGeoShape::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoShape_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoShape) ToString() string {
	defer qt.Recovering("QGeoShape::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoShape_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoShape) Type() QGeoShape__ShapeType {
	defer qt.Recovering("QGeoShape::type")

	if ptr.Pointer() != nil {
		return QGeoShape__ShapeType(C.QGeoShape_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoShape) DestroyQGeoShape() {
	defer qt.Recovering("QGeoShape::~QGeoShape")

	if ptr.Pointer() != nil {
		C.QGeoShape_DestroyQGeoShape(ptr.Pointer())
	}
}
