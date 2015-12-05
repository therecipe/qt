package location

//#include "location.h"
import "C"
import (
	"log"
	"unsafe"
)

type QGeoRouteSegment struct {
	ptr unsafe.Pointer
}

type QGeoRouteSegment_ITF interface {
	QGeoRouteSegment_PTR() *QGeoRouteSegment
}

func (p *QGeoRouteSegment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoRouteSegment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoRouteSegment(ptr QGeoRouteSegment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoRouteSegment_PTR().Pointer()
	}
	return nil
}

func NewQGeoRouteSegmentFromPointer(ptr unsafe.Pointer) *QGeoRouteSegment {
	var n = new(QGeoRouteSegment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoRouteSegment) QGeoRouteSegment_PTR() *QGeoRouteSegment {
	return ptr
}

func NewQGeoRouteSegment() *QGeoRouteSegment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::QGeoRouteSegment")
		}
	}()

	return NewQGeoRouteSegmentFromPointer(C.QGeoRouteSegment_NewQGeoRouteSegment())
}

func NewQGeoRouteSegment2(other QGeoRouteSegment_ITF) *QGeoRouteSegment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::QGeoRouteSegment")
		}
	}()

	return NewQGeoRouteSegmentFromPointer(C.QGeoRouteSegment_NewQGeoRouteSegment2(PointerFromQGeoRouteSegment(other)))
}

func (ptr *QGeoRouteSegment) Distance() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::distance")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGeoRouteSegment_Distance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteSegment) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGeoRouteSegment_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoRouteSegment) SetDistance(distance float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::setDistance")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetDistance(ptr.Pointer(), C.double(distance))
	}
}

func (ptr *QGeoRouteSegment) SetManeuver(maneuver QGeoManeuver_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::setManeuver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetManeuver(ptr.Pointer(), PointerFromQGeoManeuver(maneuver))
	}
}

func (ptr *QGeoRouteSegment) SetNextRouteSegment(routeSegment QGeoRouteSegment_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::setNextRouteSegment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetNextRouteSegment(ptr.Pointer(), PointerFromQGeoRouteSegment(routeSegment))
	}
}

func (ptr *QGeoRouteSegment) SetTravelTime(secs int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::setTravelTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_SetTravelTime(ptr.Pointer(), C.int(secs))
	}
}

func (ptr *QGeoRouteSegment) TravelTime() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::travelTime")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGeoRouteSegment_TravelTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoRouteSegment) DestroyQGeoRouteSegment() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoRouteSegment::~QGeoRouteSegment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoRouteSegment_DestroyQGeoRouteSegment(ptr.Pointer())
	}
}
