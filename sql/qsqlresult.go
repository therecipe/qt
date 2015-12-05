package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (p *QSqlResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = new(QSqlResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return ptr
}

func (ptr *QSqlResult) Handle() *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlResult::handle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) DestroyQSqlResult() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlResult::~QSqlResult")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(ptr.Pointer())
	}
}
