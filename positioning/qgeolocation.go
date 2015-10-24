package positioning

//#include "qgeolocation.h"
import "C"
import (
	"unsafe"
)

type QGeoLocation struct {
	ptr unsafe.Pointer
}

type QGeoLocationITF interface {
	QGeoLocationPTR() *QGeoLocation
}

func (p *QGeoLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoLocation(ptr QGeoLocationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoLocationPTR().Pointer()
	}
	return nil
}

func QGeoLocationFromPointer(ptr unsafe.Pointer) *QGeoLocation {
	var n = new(QGeoLocation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoLocation) QGeoLocationPTR() *QGeoLocation {
	return ptr
}
