package sql

//#include "qsqlrelationaldelegate.h"
import "C"
import (
	"unsafe"
)

type QSqlRelationalDelegate struct {
	ptr unsafe.Pointer
}

type QSqlRelationalDelegate_ITF interface {
	QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate
}

func (p *QSqlRelationalDelegate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRelationalDelegate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRelationalDelegate(ptr QSqlRelationalDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalDelegate_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) *QSqlRelationalDelegate {
	var n = new(QSqlRelationalDelegate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlRelationalDelegate) QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate {
	return ptr
}
