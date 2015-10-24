package sql

//#include "qsqldrivercreator.h"
import "C"
import (
	"unsafe"
)

type QSqlDriverCreator struct {
	QSqlDriverCreatorBase
}

type QSqlDriverCreatorITF interface {
	QSqlDriverCreatorBaseITF
	QSqlDriverCreatorPTR() *QSqlDriverCreator
}

func PointerFromQSqlDriverCreator(ptr QSqlDriverCreatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorPTR().Pointer()
	}
	return nil
}

func QSqlDriverCreatorFromPointer(ptr unsafe.Pointer) *QSqlDriverCreator {
	var n = new(QSqlDriverCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlDriverCreator) QSqlDriverCreatorPTR() *QSqlDriverCreator {
	return ptr
}
