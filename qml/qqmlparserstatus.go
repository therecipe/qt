package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (p *QQmlParserStatus) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlParserStatus) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) *QQmlParserStatus {
	var n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) ClassBegin() {
	defer qt.Recovering("QQmlParserStatus::classBegin")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	defer qt.Recovering("QQmlParserStatus::componentComplete")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}
