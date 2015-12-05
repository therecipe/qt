package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	return n
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return ptr
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkProxyFactory::setApplicationProxyFactory")
		}
	}()

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkProxyFactory::setUseSystemConfiguration")
		}
	}()

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.int(qt.GoBoolToInt(enable)))
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkProxyFactory::~QNetworkProxyFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
	}
}
