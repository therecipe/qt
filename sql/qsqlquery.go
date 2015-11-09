package sql

//#include "qsqlquery.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlQuery struct {
	ptr unsafe.Pointer
}

type QSqlQuery_ITF interface {
	QSqlQuery_PTR() *QSqlQuery
}

func (p *QSqlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlQuery(ptr QSqlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQuery_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryFromPointer(ptr unsafe.Pointer) *QSqlQuery {
	var n = new(QSqlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlQuery) QSqlQuery_PTR() *QSqlQuery {
	return ptr
}

//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int64

const (
	QSqlQuery__ValuesAsRows    = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns = QSqlQuery__BatchExecutionMode(1)
)

func NewQSqlQuery3(db QSqlDatabase_ITF) *QSqlQuery {
	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery3(PointerFromQSqlDatabase(db)))
}

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {
	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery(PointerFromQSqlResult(result)))
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {
	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery4(PointerFromQSqlQuery(other)))
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {
	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery2(C.CString(query), PointerFromQSqlDatabase(db)))
}

func (ptr *QSqlQuery) At() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_At(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue(ptr.Pointer(), C.CString(placeholder)))
	}
	return nil
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue2(ptr.Pointer(), C.int(pos)))
	}
	return nil
}

func (ptr *QSqlQuery) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlQuery_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Exec2() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_ExecBatch(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Finish() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) First() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_First(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsForwardOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull2(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull(ptr.Pointer(), C.int(field)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Last(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_LastInsertId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) LastQuery() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Next() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Next(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_NextResult(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_NumRowsAffected(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Prepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Previous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlQuery_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Seek(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(relative))) != 0
	}
	return false
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlQuery) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(ptr.Pointer())
	}
}
