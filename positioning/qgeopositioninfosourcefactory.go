package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoPositionInfoSourceFactory struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfoSourceFactory_ITF interface {
	QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory
}

func (p *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfoSourceFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfoSourceFactory(ptr QGeoPositionInfoSourceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourceFactory_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	var n = new(QGeoPositionInfoSourceFactory)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QGeoPositionInfoSourceFactory_") {
		n.SetObjectNameAbs("QGeoPositionInfoSourceFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::areaMonitor")

	if ptr.Pointer() != nil {
		return NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::positionInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::satelliteInfoSource")

	if ptr.Pointer() != nil {
		return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr.Pointer())
	}
}

func (ptr *QGeoPositionInfoSourceFactory) ObjectNameAbs() string {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoPositionInfoSourceFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoPositionInfoSourceFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGeoPositionInfoSourceFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
