package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
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
	for len(n.ObjectNameAbs()) < len("QQmlNetworkAccessManagerFactory_") {
		n.SetObjectNameAbs("QQmlNetworkAccessManagerFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::create")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) ObjectNameAbs() string {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlNetworkAccessManagerFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlNetworkAccessManagerFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
