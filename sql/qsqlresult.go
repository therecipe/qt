package sql

//#include "qsqlresult.h"
import "C"
import (
	"unsafe"
)

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResultITF interface {
	QSqlResultPTR() *QSqlResult
}

func (p *QSqlResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlResult(ptr QSqlResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResultPTR().Pointer()
	}
	return nil
}

func QSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = new(QSqlResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlResult) QSqlResultPTR() *QSqlResult {
	return ptr
}

func (ptr *QSqlResult) Handle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_Handle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlResult) DestroyQSqlResult() {
	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(C.QtObjectPtr(ptr.Pointer()))
	}
}
