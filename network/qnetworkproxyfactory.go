package network

//#include "qnetworkproxyfactory.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QNetworkProxyFactory struct {
	ptr unsafe.Pointer
}

type QNetworkProxyFactoryITF interface {
	QNetworkProxyFactoryPTR() *QNetworkProxyFactory
}

func (p *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactoryPTR().Pointer()
	}
	return nil
}

func QNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactoryPTR() *QNetworkProxyFactory {
	return ptr
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactoryITF) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(C.QtObjectPtr(PointerFromQNetworkProxyFactory(factory)))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.int(qt.GoBoolToInt(enable)))
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(C.QtObjectPtr(ptr.Pointer()))
	}
}
