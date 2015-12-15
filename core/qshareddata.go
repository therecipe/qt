package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSharedData struct {
	ptr unsafe.Pointer
}

type QSharedData_ITF interface {
	QSharedData_PTR() *QSharedData
}

func (p *QSharedData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedData(ptr QSharedData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedData_PTR().Pointer()
	}
	return nil
}

func NewQSharedDataFromPointer(ptr unsafe.Pointer) *QSharedData {
	var n = new(QSharedData)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedData) QSharedData_PTR() *QSharedData {
	return ptr
}

func NewQSharedData() *QSharedData {
	defer qt.Recovering("QSharedData::QSharedData")

	return NewQSharedDataFromPointer(C.QSharedData_NewQSharedData())
}

func NewQSharedData2(other QSharedData_ITF) *QSharedData {
	defer qt.Recovering("QSharedData::QSharedData")

	return NewQSharedDataFromPointer(C.QSharedData_NewQSharedData2(PointerFromQSharedData(other)))
}
