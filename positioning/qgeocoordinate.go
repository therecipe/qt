package positioning

//#include "qgeocoordinate.h"
import "C"
import (
	"unsafe"
)

type QGeoCoordinate struct {
	ptr unsafe.Pointer
}

type QGeoCoordinateITF interface {
	QGeoCoordinatePTR() *QGeoCoordinate
}

func (p *QGeoCoordinate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoCoordinate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoCoordinate(ptr QGeoCoordinateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCoordinatePTR().Pointer()
	}
	return nil
}

func QGeoCoordinateFromPointer(ptr unsafe.Pointer) *QGeoCoordinate {
	var n = new(QGeoCoordinate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoCoordinate) QGeoCoordinatePTR() *QGeoCoordinate {
	return ptr
}

//QGeoCoordinate::CoordinateFormat
type QGeoCoordinate__CoordinateFormat int

var (
	QGeoCoordinate__Degrees                             = QGeoCoordinate__CoordinateFormat(0)
	QGeoCoordinate__DegreesWithHemisphere               = QGeoCoordinate__CoordinateFormat(1)
	QGeoCoordinate__DegreesMinutes                      = QGeoCoordinate__CoordinateFormat(2)
	QGeoCoordinate__DegreesMinutesWithHemisphere        = QGeoCoordinate__CoordinateFormat(3)
	QGeoCoordinate__DegreesMinutesSeconds               = QGeoCoordinate__CoordinateFormat(4)
	QGeoCoordinate__DegreesMinutesSecondsWithHemisphere = QGeoCoordinate__CoordinateFormat(5)
)

//QGeoCoordinate::CoordinateType
type QGeoCoordinate__CoordinateType int

var (
	QGeoCoordinate__InvalidCoordinate = QGeoCoordinate__CoordinateType(0)
	QGeoCoordinate__Coordinate2D      = QGeoCoordinate__CoordinateType(1)
	QGeoCoordinate__Coordinate3D      = QGeoCoordinate__CoordinateType(2)
)

func NewQGeoCoordinate() *QGeoCoordinate {
	return QGeoCoordinateFromPointer(unsafe.Pointer(C.QGeoCoordinate_NewQGeoCoordinate()))
}

func NewQGeoCoordinate4(other QGeoCoordinateITF) *QGeoCoordinate {
	return QGeoCoordinateFromPointer(unsafe.Pointer(C.QGeoCoordinate_NewQGeoCoordinate4(C.QtObjectPtr(PointerFromQGeoCoordinate(other)))))
}

func (ptr *QGeoCoordinate) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoCoordinate_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoCoordinate) ToString(format QGeoCoordinate__CoordinateFormat) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCoordinate_ToString(C.QtObjectPtr(ptr.Pointer()), C.int(format)))
	}
	return ""
}

func (ptr *QGeoCoordinate) Type() QGeoCoordinate__CoordinateType {
	if ptr.Pointer() != nil {
		return QGeoCoordinate__CoordinateType(C.QGeoCoordinate_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoCoordinate) DestroyQGeoCoordinate() {
	if ptr.Pointer() != nil {
		C.QGeoCoordinate_DestroyQGeoCoordinate(C.QtObjectPtr(ptr.Pointer()))
	}
}
