package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGeoCoordinate struct {
	ptr unsafe.Pointer
}

type QGeoCoordinate_ITF interface {
	QGeoCoordinate_PTR() *QGeoCoordinate
}

func (p *QGeoCoordinate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoCoordinate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoCoordinate(ptr QGeoCoordinate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCoordinate_PTR().Pointer()
	}
	return nil
}

func NewQGeoCoordinateFromPointer(ptr unsafe.Pointer) *QGeoCoordinate {
	var n = new(QGeoCoordinate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoCoordinate) QGeoCoordinate_PTR() *QGeoCoordinate {
	return ptr
}

//QGeoCoordinate::CoordinateFormat
type QGeoCoordinate__CoordinateFormat int64

const (
	QGeoCoordinate__Degrees                             = QGeoCoordinate__CoordinateFormat(0)
	QGeoCoordinate__DegreesWithHemisphere               = QGeoCoordinate__CoordinateFormat(1)
	QGeoCoordinate__DegreesMinutes                      = QGeoCoordinate__CoordinateFormat(2)
	QGeoCoordinate__DegreesMinutesWithHemisphere        = QGeoCoordinate__CoordinateFormat(3)
	QGeoCoordinate__DegreesMinutesSeconds               = QGeoCoordinate__CoordinateFormat(4)
	QGeoCoordinate__DegreesMinutesSecondsWithHemisphere = QGeoCoordinate__CoordinateFormat(5)
)

//QGeoCoordinate::CoordinateType
type QGeoCoordinate__CoordinateType int64

const (
	QGeoCoordinate__InvalidCoordinate = QGeoCoordinate__CoordinateType(0)
	QGeoCoordinate__Coordinate2D      = QGeoCoordinate__CoordinateType(1)
	QGeoCoordinate__Coordinate3D      = QGeoCoordinate__CoordinateType(2)
)

func NewQGeoCoordinate() *QGeoCoordinate {
	defer qt.Recovering("QGeoCoordinate::QGeoCoordinate")

	return NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate())
}

func NewQGeoCoordinate4(other QGeoCoordinate_ITF) *QGeoCoordinate {
	defer qt.Recovering("QGeoCoordinate::QGeoCoordinate")

	return NewQGeoCoordinateFromPointer(C.QGeoCoordinate_NewQGeoCoordinate4(PointerFromQGeoCoordinate(other)))
}

func (ptr *QGeoCoordinate) AzimuthTo(other QGeoCoordinate_ITF) float64 {
	defer qt.Recovering("QGeoCoordinate::azimuthTo")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_AzimuthTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) DistanceTo(other QGeoCoordinate_ITF) float64 {
	defer qt.Recovering("QGeoCoordinate::distanceTo")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCoordinate_DistanceTo(ptr.Pointer(), PointerFromQGeoCoordinate(other)))
	}
	return 0
}

func (ptr *QGeoCoordinate) IsValid() bool {
	defer qt.Recovering("QGeoCoordinate::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoCoordinate_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {
	defer qt.Recovering("QGeoCoordinate::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCoordinate_ToString(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {
	defer qt.Recovering("QGeoCoordinate::type")

	if ptr.Pointer() != nil {
		return QGeoCoordinate__CoordinateType(C.QGeoCoordinate_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {
	defer qt.Recovering("QGeoCoordinate::~QGeoCoordinate")

	if ptr.Pointer() != nil {
		C.QGeoCoordinate_DestroyQGeoCoordinate(ptr.Pointer())
	}
}
