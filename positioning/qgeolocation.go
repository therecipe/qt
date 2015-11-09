package positioning

//#include "qgeolocation.h"
import "C"
import (
	"unsafe"
)

type QGeoLocation struct {
	ptr unsafe.Pointer
}

type QGeoLocation_ITF interface {
	QGeoLocation_PTR() *QGeoLocation
}

func (p *QGeoLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoLocation(ptr QGeoLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoLocation_PTR().Pointer()
	}
	return nil
}

func NewQGeoLocationFromPointer(ptr unsafe.Pointer) *QGeoLocation {
	var n = new(QGeoLocation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoLocation) QGeoLocation_PTR() *QGeoLocation {
	return ptr
}
