package sql

//#include "qsqlrelationaldelegate.h"
import "C"
import (
	"unsafe"
)

type QSqlRelationalDelegate struct {
	ptr unsafe.Pointer
}

type QSqlRelationalDelegateITF interface {
	QSqlRelationalDelegatePTR() *QSqlRelationalDelegate
}

func (p *QSqlRelationalDelegate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRelationalDelegate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRelationalDelegate(ptr QSqlRelationalDelegateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalDelegatePTR().Pointer()
	}
	return nil
}

func QSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) *QSqlRelationalDelegate {
	var n = new(QSqlRelationalDelegate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlRelationalDelegate) QSqlRelationalDelegatePTR() *QSqlRelationalDelegate {
	return ptr
}
