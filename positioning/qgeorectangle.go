package positioning

//#include "positioning.h"
import "C"
import (
	"log"
	"unsafe"
)

type QGeoRectangle struct {
	QGeoShape
}

type QGeoRectangle_ITF interface {
	QGeoShape_ITF
	QGeoRectangle_PTR() *QGeoRectangle
}

func PointerFromQGeoRectangle(ptr QGeoRectangle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRectangle_PTR().Pointer()
	}
	return nil
}

func NewQGeoRectangleFromPointer(ptr unsafe.Pointer) *QGeoRectangle {
	var n = new(QGeoRectangle)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRectangle) QGeoRectangle_PTR() *QGeoRectangle {
	return ptr
}

func NewQGeoRectangle() *QGeoRectangle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::QGeoRectangle")
		}
	}()

	return NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle())
}

func NewQGeoRectangle3(topLeft QGeoCoordinate_ITF, bottomRight QGeoCoordinate_ITF) *QGeoRectangle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::QGeoRectangle")
		}
	}()

	return NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle3(PointerFromQGeoCoordinate(topLeft), PointerFromQGeoCoordinate(bottomRight)))
}

func NewQGeoRectangle5(other QGeoRectangle_ITF) *QGeoRectangle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::QGeoRectangle")
		}
	}()

	return NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle5(PointerFromQGeoRectangle(other)))
}

func NewQGeoRectangle6(other QGeoShape_ITF) *QGeoRectangle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::QGeoRectangle")
		}
	}()

	return NewQGeoRectangleFromPointer(C.QGeoRectangle_NewQGeoRectangle6(PointerFromQGeoShape(other)))
}

func (ptr *QGeoRectangle) Contains(rectangle QGeoRectangle_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Contains(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) Intersects(rectangle QGeoRectangle_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::intersects")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoRectangle_Intersects(ptr.Pointer(), PointerFromQGeoRectangle(rectangle)) != 0
	}
	return false
}

func (ptr *QGeoRectangle) SetBottomLeft(bottomLeft QGeoCoordinate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::setBottomLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomLeft(ptr.Pointer(), PointerFromQGeoCoordinate(bottomLeft))
	}
}

func (ptr *QGeoRectangle) SetBottomRight(bottomRight QGeoCoordinate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::setBottomRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetBottomRight(ptr.Pointer(), PointerFromQGeoCoordinate(bottomRight))
	}
}

func (ptr *QGeoRectangle) SetCenter(center QGeoCoordinate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::setCenter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoRectangle) SetTopLeft(topLeft QGeoCoordinate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::setTopLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopLeft(ptr.Pointer(), PointerFromQGeoCoordinate(topLeft))
	}
}

func (ptr *QGeoRectangle) SetTopRight(topRight QGeoCoordinate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::setTopRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_SetTopRight(ptr.Pointer(), PointerFromQGeoCoordinate(topRight))
	}
}

func (ptr *QGeoRectangle) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoRectangle_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoRectangle) DestroyQGeoRectangle() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRectangle::~QGeoRectangle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRectangle_DestroyQGeoRectangle(ptr.Pointer())
	}
}
