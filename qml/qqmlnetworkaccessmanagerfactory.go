package qml

//#include "qqmlnetworkaccessmanagerfactory.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QQmlNetworkAccessManagerFactory struct {
	ptr unsafe.Pointer
}

type QQmlNetworkAccessManagerFactoryITF interface {
	QQmlNetworkAccessManagerFactoryPTR() *QQmlNetworkAccessManagerFactory
}

func (p *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlNetworkAccessManagerFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactoryPTR().Pointer()
	}
	return nil
}

func QQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	var n = new(QQmlNetworkAccessManagerFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactoryPTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObjectITF) *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return network.QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QQmlNetworkAccessManagerFactory_Create(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)))))
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(C.QtObjectPtr(ptr.Pointer()))
	}
}
