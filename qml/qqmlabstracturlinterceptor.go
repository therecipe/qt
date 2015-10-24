package qml

//#include "qqmlabstracturlinterceptor.h"
import "C"
import (
	"unsafe"
)

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptorITF interface {
	QQmlAbstractUrlInterceptorPTR() *QQmlAbstractUrlInterceptor
}

func (p *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlAbstractUrlInterceptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptorPTR().Pointer()
	}
	return nil
}

func QQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	var n = new(QQmlAbstractUrlInterceptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptorPTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int

var (
	QQmlAbstractUrlInterceptor__QmlFile        = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url string, ty QQmlAbstractUrlInterceptor__DataType) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlAbstractUrlInterceptor_Intercept(C.QtObjectPtr(ptr.Pointer()), C.CString(url), C.int(ty)))
	}
	return ""
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(C.QtObjectPtr(ptr.Pointer()))
	}
}
