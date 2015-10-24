package positioning

//#include "qgeopositioninfosourcefactory.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoPositionInfoSourceFactory struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfoSourceFactoryITF interface {
	QGeoPositionInfoSourceFactoryPTR() *QGeoPositionInfoSourceFactory
}

func (p *QGeoPositionInfoSourceFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfoSourceFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfoSourceFactory(ptr QGeoPositionInfoSourceFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfoSourceFactoryPTR().Pointer()
	}
	return nil
}

func QGeoPositionInfoSourceFactoryFromPointer(ptr unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	var n = new(QGeoPositionInfoSourceFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactoryPTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (ptr *QGeoPositionInfoSourceFactory) AreaMonitor(parent core.QObjectITF) *QGeoAreaMonitorSource {
	if ptr.Pointer() != nil {
		return QGeoAreaMonitorSourceFromPointer(unsafe.Pointer(C.QGeoPositionInfoSourceFactory_AreaMonitor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)))))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) PositionInfoSource(parent core.QObjectITF) *QGeoPositionInfoSource {
	if ptr.Pointer() != nil {
		return QGeoPositionInfoSourceFromPointer(unsafe.Pointer(C.QGeoPositionInfoSourceFactory_PositionInfoSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)))))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent core.QObjectITF) *QGeoSatelliteInfoSource {
	if ptr.Pointer() != nil {
		return QGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(C.QGeoPositionInfoSourceFactory_SatelliteInfoSource(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)))))
	}
	return nil
}

func (ptr *QGeoPositionInfoSourceFactory) DestroyQGeoPositionInfoSourceFactory() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(C.QtObjectPtr(ptr.Pointer()))
	}
}
