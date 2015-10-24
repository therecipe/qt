package location

//#include "qgeoserviceproviderfactory.h"
import "C"
import (
	"unsafe"
)

type QGeoServiceProviderFactory struct {
	ptr unsafe.Pointer
}

type QGeoServiceProviderFactoryITF interface {
	QGeoServiceProviderFactoryPTR() *QGeoServiceProviderFactory
}

func (p *QGeoServiceProviderFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoServiceProviderFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoServiceProviderFactory(ptr QGeoServiceProviderFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoServiceProviderFactoryPTR().Pointer()
	}
	return nil
}

func QGeoServiceProviderFactoryFromPointer(ptr unsafe.Pointer) *QGeoServiceProviderFactory {
	var n = new(QGeoServiceProviderFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoServiceProviderFactory) QGeoServiceProviderFactoryPTR() *QGeoServiceProviderFactory {
	return ptr
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactory() {
	if ptr.Pointer() != nil {
		C.QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(C.QtObjectPtr(ptr.Pointer()))
	}
}
