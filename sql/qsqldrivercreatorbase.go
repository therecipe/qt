package sql

//#include "qsqldrivercreatorbase.h"
import "C"
import (
	"unsafe"
)

type QSqlDriverCreatorBase struct {
	ptr unsafe.Pointer
}

type QSqlDriverCreatorBase_ITF interface {
	QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase
}

func (p *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlDriverCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) *QSqlDriverCreatorBase {
	var n = new(QSqlDriverCreatorBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverCreatorBase_CreateObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr.Pointer())
	}
}
