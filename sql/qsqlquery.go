package sql

//#include "qsqlquery.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlQuery struct {
	ptr unsafe.Pointer
}

type QSqlQueryITF interface {
	QSqlQueryPTR() *QSqlQuery
}

func (p *QSqlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlQuery(ptr QSqlQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryPTR().Pointer()
	}
	return nil
}

func QSqlQueryFromPointer(ptr unsafe.Pointer) *QSqlQuery {
	var n = new(QSqlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlQuery) QSqlQueryPTR() *QSqlQuery {
	return ptr
}

//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int

var (
	QSqlQuery__ValuesAsRows    = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns = QSqlQuery__BatchExecutionMode(1)
)

func NewQSqlQuery3(db QSqlDatabaseITF) *QSqlQuery {
	return QSqlQueryFromPointer(unsafe.Pointer(C.QSqlQuery_NewQSqlQuery3(C.QtObjectPtr(PointerFromQSqlDatabase(db)))))
}

func NewQSqlQuery(result QSqlResultITF) *QSqlQuery {
	return QSqlQueryFromPointer(unsafe.Pointer(C.QSqlQuery_NewQSqlQuery(C.QtObjectPtr(PointerFromQSqlResult(result)))))
}

func NewQSqlQuery4(other QSqlQueryITF) *QSqlQuery {
	return QSqlQueryFromPointer(unsafe.Pointer(C.QSqlQuery_NewQSqlQuery4(C.QtObjectPtr(PointerFromQSqlQuery(other)))))
}

func NewQSqlQuery2(query string, db QSqlDatabaseITF) *QSqlQuery {
	return QSqlQueryFromPointer(unsafe.Pointer(C.QSqlQuery_NewQSqlQuery2(C.CString(query), C.QtObjectPtr(PointerFromQSqlDatabase(db)))))
}

func (ptr *QSqlQuery) At() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_At(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) BoundValue(placeholder string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_BoundValue(C.QtObjectPtr(ptr.Pointer()), C.CString(placeholder)))
	}
	return ""
}

func (ptr *QSqlQuery) BoundValue2(pos int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_BoundValue2(C.QtObjectPtr(ptr.Pointer()), C.int(pos)))
	}
	return ""
}

func (ptr *QSqlQuery) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	if ptr.Pointer() != nil {
		return QSqlDriverFromPointer(unsafe.Pointer(C.QSqlQuery_Driver(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSqlQuery) Exec2() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec2(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec(C.QtObjectPtr(ptr.Pointer()), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_ExecBatch(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_ExecutedQuery(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlQuery) Finish() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlQuery) First() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_First(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsForwardOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull(C.QtObjectPtr(ptr.Pointer()), C.int(field)) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsSelect(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Last(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) LastInsertId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastInsertId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlQuery) LastQuery() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastQuery(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlQuery) Next() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Next(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_NextResult(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_NumRowsAffected(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Prepare(C.QtObjectPtr(ptr.Pointer()), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Previous(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	if ptr.Pointer() != nil {
		return QSqlResultFromPointer(unsafe.Pointer(C.QSqlQuery_Result(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQuery_Seek(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(relative))) != 0
	}
	return false
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlQuery) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQuery_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) Value2(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_Value2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QSqlQuery) Value(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_Value(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(C.QtObjectPtr(ptr.Pointer()))
	}
}
