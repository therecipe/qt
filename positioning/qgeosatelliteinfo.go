package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGeoSatelliteInfo struct {
	ptr unsafe.Pointer
}

type QGeoSatelliteInfo_ITF interface {
	QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo
}

func (p *QGeoSatelliteInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoSatelliteInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoSatelliteInfo(ptr QGeoSatelliteInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoSatelliteInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoSatelliteInfoFromPointer(ptr unsafe.Pointer) *QGeoSatelliteInfo {
	var n = new(QGeoSatelliteInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoSatelliteInfo) QGeoSatelliteInfo_PTR() *QGeoSatelliteInfo {
	return ptr
}

//QGeoSatelliteInfo::Attribute
type QGeoSatelliteInfo__Attribute int64

const (
	QGeoSatelliteInfo__Elevation = QGeoSatelliteInfo__Attribute(0)
	QGeoSatelliteInfo__Azimuth   = QGeoSatelliteInfo__Attribute(1)
)

//QGeoSatelliteInfo::SatelliteSystem
type QGeoSatelliteInfo__SatelliteSystem int64

const (
	QGeoSatelliteInfo__Undefined = QGeoSatelliteInfo__SatelliteSystem(0x00)
	QGeoSatelliteInfo__GPS       = QGeoSatelliteInfo__SatelliteSystem(0x01)
	QGeoSatelliteInfo__GLONASS   = QGeoSatelliteInfo__SatelliteSystem(0x02)
)

func NewQGeoSatelliteInfo() *QGeoSatelliteInfo {
	defer qt.Recovering("QGeoSatelliteInfo::QGeoSatelliteInfo")

	return NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo())
}

func NewQGeoSatelliteInfo2(other QGeoSatelliteInfo_ITF) *QGeoSatelliteInfo {
	defer qt.Recovering("QGeoSatelliteInfo::QGeoSatelliteInfo")

	return NewQGeoSatelliteInfoFromPointer(C.QGeoSatelliteInfo_NewQGeoSatelliteInfo2(PointerFromQGeoSatelliteInfo(other)))
}

func (ptr *QGeoSatelliteInfo) Attribute(attribute QGeoSatelliteInfo__Attribute) float64 {
	defer qt.Recovering("QGeoSatelliteInfo::attribute")

	if ptr.Pointer() != nil {
		return float64(C.QGeoSatelliteInfo_Attribute(ptr.Pointer(), C.int(attribute)))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) HasAttribute(attribute QGeoSatelliteInfo__Attribute) bool {
	defer qt.Recovering("QGeoSatelliteInfo::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QGeoSatelliteInfo_HasAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoSatelliteInfo) RemoveAttribute(attribute QGeoSatelliteInfo__Attribute) {
	defer qt.Recovering("QGeoSatelliteInfo::removeAttribute")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_RemoveAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QGeoSatelliteInfo) SatelliteIdentifier() int {
	defer qt.Recovering("QGeoSatelliteInfo::satelliteIdentifier")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SatelliteIdentifier(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SatelliteSystem() QGeoSatelliteInfo__SatelliteSystem {
	defer qt.Recovering("QGeoSatelliteInfo::satelliteSystem")

	if ptr.Pointer() != nil {
		return QGeoSatelliteInfo__SatelliteSystem(C.QGeoSatelliteInfo_SatelliteSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) SetAttribute(attribute QGeoSatelliteInfo__Attribute, value float64) {
	defer qt.Recovering("QGeoSatelliteInfo::setAttribute")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetAttribute(ptr.Pointer(), C.int(attribute), C.double(value))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteIdentifier(satId int) {
	defer qt.Recovering("QGeoSatelliteInfo::setSatelliteIdentifier")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteIdentifier(ptr.Pointer(), C.int(satId))
	}
}

func (ptr *QGeoSatelliteInfo) SetSatelliteSystem(system QGeoSatelliteInfo__SatelliteSystem) {
	defer qt.Recovering("QGeoSatelliteInfo::setSatelliteSystem")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSatelliteSystem(ptr.Pointer(), C.int(system))
	}
}

func (ptr *QGeoSatelliteInfo) SetSignalStrength(signalStrength int) {
	defer qt.Recovering("QGeoSatelliteInfo::setSignalStrength")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_SetSignalStrength(ptr.Pointer(), C.int(signalStrength))
	}
}

func (ptr *QGeoSatelliteInfo) SignalStrength() int {
	defer qt.Recovering("QGeoSatelliteInfo::signalStrength")

	if ptr.Pointer() != nil {
		return int(C.QGeoSatelliteInfo_SignalStrength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoSatelliteInfo) DestroyQGeoSatelliteInfo() {
	defer qt.Recovering("QGeoSatelliteInfo::~QGeoSatelliteInfo")

	if ptr.Pointer() != nil {
		C.QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(ptr.Pointer())
	}
}
