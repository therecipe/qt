package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QNetworkProxyFactory struct {
	ptr unsafe.Pointer
}

type QNetworkProxyFactory_ITF interface {
	QNetworkProxyFactory_PTR() *QNetworkProxyFactory
}

func (p *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactory_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QNetworkProxyFactory_") {
		n.SetObjectNameAbs("QNetworkProxyFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return ptr
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkProxyFactory::setApplicationProxyFactory")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	defer qt.Recovering("QNetworkProxyFactory::setUseSystemConfiguration")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.int(qt.GoBoolToInt(enable)))
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	defer qt.Recovering("QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
	}
}

func (ptr *QNetworkProxyFactory) ObjectNameAbs() string {
	defer qt.Recovering("QNetworkProxyFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QNetworkProxyFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
