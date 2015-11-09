package qml

//#include "qqmlabstracturlinterceptor.h"
import "C"
import (
	"unsafe"
)

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (p *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlAbstractUrlInterceptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptor_PTR().Pointer()
	}
	return nil
}

func NewQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	var n = new(QQmlAbstractUrlInterceptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr.Pointer())
	}
}
