package location

//#include "qgeoserviceproviderfactory.h"
import "C"
import (
	"unsafe"
)

type QGeoServiceProviderFactory struct {
	ptr unsafe.Pointer
}

type QGeoServiceProviderFactory_ITF interface {
	QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory
}

func (p *QGeoServiceProviderFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoServiceProviderFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoServiceProviderFactory(ptr QGeoServiceProviderFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactory_PTR().Pointer()
	}
	return nil
}

func NewQGeoServiceProviderFactoryFromPointer(ptr unsafe.Pointer) *QGeoServiceProviderFactory {
	var n = new(QGeoServiceProviderFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoServiceProviderFactory) QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory {
	return ptr
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactory() {
	if ptr.Pointer() != nil {
		C.QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(ptr.Pointer())
	}
}
