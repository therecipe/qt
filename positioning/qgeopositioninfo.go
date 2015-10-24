package positioning

//#include "qgeopositioninfo.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoPositionInfo struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfoITF interface {
	QGeoPositionInfoPTR() *QGeoPositionInfo
}

func (p *QGeoPositionInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfo(ptr QGeoPositionInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoPTR().Pointer()
	}
	return nil
}

func QGeoPositionInfoFromPointer(ptr unsafe.Pointer) *QGeoPositionInfo {
	var n = new(QGeoPositionInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoPositionInfo) QGeoPositionInfoPTR() *QGeoPositionInfo {
	return ptr
}

//QGeoPositionInfo::Attribute
type QGeoPositionInfo__Attribute int

var (
	QGeoPositionInfo__Direction          = QGeoPositionInfo__Attribute(0)
	QGeoPositionInfo__GroundSpeed        = QGeoPositionInfo__Attribute(1)
	QGeoPositionInfo__VerticalSpeed      = QGeoPositionInfo__Attribute(2)
	QGeoPositionInfo__MagneticVariation  = QGeoPositionInfo__Attribute(3)
	QGeoPositionInfo__HorizontalAccuracy = QGeoPositionInfo__Attribute(4)
	QGeoPositionInfo__VerticalAccuracy   = QGeoPositionInfo__Attribute(5)
)

func NewQGeoPositionInfo() *QGeoPositionInfo {
	return QGeoPositionInfoFromPointer(unsafe.Pointer(C.QGeoPositionInfo_NewQGeoPositionInfo()))
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinateITF, timestamp core.QDateTimeITF) *QGeoPositionInfo {
	return QGeoPositionInfoFromPointer(unsafe.Pointer(C.QGeoPositionInfo_NewQGeoPositionInfo2(C.QtObjectPtr(PointerFromQGeoCoordinate(coordinate)), C.QtObjectPtr(core.PointerFromQDateTime(timestamp)))))
}

func NewQGeoPositionInfo3(other QGeoPositionInfoITF) *QGeoPositionInfo {
	return QGeoPositionInfoFromPointer(unsafe.Pointer(C.QGeoPositionInfo_NewQGeoPositionInfo3(C.QtObjectPtr(PointerFromQGeoPositionInfo(other)))))
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_HasAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) RemoveAttribute(attribute QGeoPositionInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_RemoveAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute))
	}
}

func (ptr *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinateITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetCoordinate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGeoCoordinate(coordinate)))
	}
}

func (ptr *QGeoPositionInfo) SetTimestamp(timestamp core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetTimestamp(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(timestamp)))
	}
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
