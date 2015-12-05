package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
	"unsafe"
)

type QQmlNetworkAccessManagerFactory struct {
	ptr unsafe.Pointer
}

type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (p *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlNetworkAccessManagerFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactory_PTR().Pointer()
	}
	return nil
}

func NewQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	var n = new(QQmlNetworkAccessManagerFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlNetworkAccessManagerFactory::create")
		}
	}()

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
	}
}
