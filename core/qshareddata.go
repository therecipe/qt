package core

//#include "qshareddata.h"
import "C"
import (
	"unsafe"
)

type QSharedData struct {
	ptr unsafe.Pointer
}

type QSharedDataITF interface {
	QSharedDataPTR() *QSharedData
}

func (p *QSharedData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedData(ptr QSharedDataITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedDataPTR().Pointer()
	}
	return nil
}

func QSharedDataFromPointer(ptr unsafe.Pointer) *QSharedData {
	var n = new(QSharedData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedData) QSharedDataPTR() *QSharedData {
	return ptr
}

func NewQSharedData() *QSharedData {
	return QSharedDataFromPointer(unsafe.Pointer(C.QSharedData_NewQSharedData()))
}

func NewQSharedData2(other QSharedDataITF) *QSharedData {
	return QSharedDataFromPointer(unsafe.Pointer(C.QSharedData_NewQSharedData2(C.QtObjectPtr(PointerFromQSharedData(other)))))
}
