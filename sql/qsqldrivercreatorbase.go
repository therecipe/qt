package sql

//#include "qsqldrivercreatorbase.h"
import "C"
import (
	"unsafe"
)

type QSqlDriverCreatorBase struct {
	ptr unsafe.Pointer
}

type QSqlDriverCreatorBaseITF interface {
	QSqlDriverCreatorBasePTR() *QSqlDriverCreatorBase
}

func (p *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlDriverCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBasePTR().Pointer()
	}
	return nil
}

func QSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) *QSqlDriverCreatorBase {
	var n = new(QSqlDriverCreatorBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBasePTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	if ptr.Pointer() != nil {
		return QSqlDriverFromPointer(unsafe.Pointer(C.QSqlDriverCreatorBase_CreateObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(C.QtObjectPtr(ptr.Pointer()))
	}
}
