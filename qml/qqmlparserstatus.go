package qml

//#include "qqmlparserstatus.h"
import "C"
import (
	"unsafe"
)

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatusITF interface {
	QQmlParserStatusPTR() *QQmlParserStatus
}

func (p *QQmlParserStatus) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlParserStatus) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatusITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatusPTR().Pointer()
	}
	return nil
}

func QQmlParserStatusFromPointer(ptr unsafe.Pointer) *QQmlParserStatus {
	var n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlParserStatus) QQmlParserStatusPTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(C.QtObjectPtr(ptr.Pointer()))
	}
}
