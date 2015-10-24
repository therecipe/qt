package positioning

//#include "qgeosatelliteinfo.h"
import "C"
import (
	"unsafe"
)

type QGeoSatelliteInfo struct {
	ptr unsafe.Pointer
}

type QGeoSatelliteInfoITF interface {
	QGeoSatelliteInfoPTR() *QGeoSatelliteInfo
}

func (p *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoSatelliteInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoSatelliteInfo(ptr QGeoSatelliteInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfoPTR().Pointer()
	}
	return nil
}

func QGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfo {
	var n = new(QGeoSatelliteInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfoPTR() *QGeoSatelliteInfo {
	return ptr
}

//QGeoSatelliteInfo::Attribute
type QGeoSatelliteInfo__Attribute int

var (
	QGeoSatelliteInfo__Elevation = QGeoSatelliteInfo__Attribute(0)
	QGeoSatelliteInfo__Azimuth   = QGeoSatelliteInfo__Attribute(1)
)

//QGeoSatelliteInfo::SatelliteSystem
type QGeoSatelliteInfo__SatelliteSystem int

var (
	QGeoSatelliteInfo__Undefined = QGeoSatelliteInfo__SatelliteSystem(0x00)
	QGeoSatelliteInfo__GPS       = QGeoSatelliteInfo__SatelliteSystem(0x01)
	QGeoSatelliteInfo__GLONASS   = QGeoSatelliteInfo__SatelliteSystem(0x02)
)

func NewQGeoSatelliteInfo() *QGeoSatelliteInfo {
	return QGeoSatelliteInfoFromPointer(unsafe.Pointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo()))
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfoITF) *QGeoSatelliteInfo {
	return QGeoSatelliteInfoFromPointer(unsafe.Pointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo2(C.QtObjectPtr(PointerFromQGeoSatelliteInfo(other)))))
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfo_HasAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_RemoveAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute))
	}
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SatelliteIdentifier(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfo__SatelliteSystem(C.QGeoSatelliteInfo_SatelliteSystem(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteIdentifier(C.QtObjectPtr(ptr.Pointer()), C.int(satId))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteSystem(system QGeoSatelliteInfo__SatelliteSystem) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteSystem(C.QtObjectPtr(ptr.Pointer()), C.int(system))
	}
}

func (ptr *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSignalStrength(C.QtObjectPtr(ptr.Pointer()), C.int(signalStrength))
	}
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {
	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SignalStrength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {
	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
