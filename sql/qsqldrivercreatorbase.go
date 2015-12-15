package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QSqlDriverCreatorBase_") {
		n.SetObjectNameAbs("QSqlDriverCreatorBase_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return ptr
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	defer qt.Recovering("QSqlDriverCreatorBase::createObject")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlDriverCreatorBase_CreateObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	defer qt.Recovering("QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr.Pointer())
	}
}

func (ptr *QSqlDriverCreatorBase) ObjectNameAbs() string {
	defer qt.Recovering("QSqlDriverCreatorBase::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriverCreatorBase_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDriverCreatorBase) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSqlDriverCreatorBase::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
