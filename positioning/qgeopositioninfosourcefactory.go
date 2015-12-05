package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
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
	return n
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObject_ITF) *QGeoAreaMonitorSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSourceFactory::areaMonitor")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoAreaMonitorSourceFromPointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObject_ITF) *QGeoPositionInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSourceFactory::positionInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoPositionInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObject_ITF) *QGeoSatelliteInfoSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSourceFactory::satelliteInfoSource")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGeoSatelliteInfoSourceFromPointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGeoPositionInfoSourceFactory::~QGeoPositionInfoSourceFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(ptr.Pointer())
	}
}
