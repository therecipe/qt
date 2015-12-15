package sql

//#include "sql.h"
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
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery3(PointerFromQSqlDatabase(db)))
}

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery(PointerFromQSqlResult(result)))
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery4(PointerFromQSqlQuery(other)))
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	return NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery2(C.CString(query), PointerFromQSqlDatabase(db)))
}

func (ptr *QSqlQuery) At() int {
	defer qt.Recovering("QSqlQuery::at")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_At(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue(ptr.Pointer(), C.CString(placeholder)))
	}
	return nil
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_BoundValue2(ptr.Pointer(), C.int(pos)))
	}
	return nil
}

func (ptr *QSqlQuery) Clear() {
	defer qt.Recovering("QSqlQuery::clear")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlQuery::driver")

	if ptr.Pointer() != nil {
		return NewQSqlDriverFromPointer(C.QSqlQuery_Driver(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Exec2() bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	defer qt.Recovering("QSqlQuery::execBatch")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_ExecBatch(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	defer qt.Recovering("QSqlQuery::executedQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Finish() {
	defer qt.Recovering("QSqlQuery::finish")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) First() bool {
	defer qt.Recovering("QSqlQuery::first")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_First(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsActive() bool {
	defer qt.Recovering("QSqlQuery::isActive")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	defer qt.Recovering("QSqlQuery::isForwardOnly")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsForwardOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull2(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull(ptr.Pointer(), C.int(field)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	defer qt.Recovering("QSqlQuery::isSelect")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	defer qt.Recovering("QSqlQuery::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	defer qt.Recovering("QSqlQuery::last")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Last(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlQuery::lastInsertId")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_LastInsertId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) LastQuery() string {
	defer qt.Recovering("QSqlQuery::lastQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Next() bool {
	defer qt.Recovering("QSqlQuery::next")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Next(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	defer qt.Recovering("QSqlQuery::nextResult")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_NextResult(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	defer qt.Recovering("QSqlQuery::numRowsAffected")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_NumRowsAffected(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	defer qt.Recovering("QSqlQuery::prepare")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Prepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	defer qt.Recovering("QSqlQuery::previous")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Previous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	defer qt.Recovering("QSqlQuery::result")

	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlQuery_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	defer qt.Recovering("QSqlQuery::seek")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Seek(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(relative))) != 0
	}
	return false
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlQuery::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlQuery) Size() int {
	defer qt.Recovering("QSqlQuery::size")

	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQuery_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	defer qt.Recovering("QSqlQuery::~QSqlQuery")

	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(ptr.Pointer())
	}
}
