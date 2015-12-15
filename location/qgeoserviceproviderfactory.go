package location

//#include "location.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QGeoServiceProviderFactory_") {
		n.SetObjectNameAbs("QGeoServiceProviderFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QGeoServiceProviderFactory) QGeoServiceProviderFactory_PTR() *QGeoServiceProviderFactory {
	return ptr
}

func (ptr *QGeoServiceProviderFactory) DestroyQGeoServiceProviderFactory() {
	defer qt.Recovering("QGeoServiceProviderFactory::~QGeoServiceProviderFactory")

	if ptr.Pointer() != nil {
		C.QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(ptr.Pointer())
	}
}

func (ptr *QGeoServiceProviderFactory) ObjectNameAbs() string {
	defer qt.Recovering("QGeoServiceProviderFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoServiceProviderFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoServiceProviderFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QGeoServiceProviderFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QGeoServiceProviderFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
