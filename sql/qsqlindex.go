package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndex_ITF interface {
	QSqlRecord_ITF
	QSqlIndex_PTR() *QSqlIndex
}

func PointerFromQSqlIndex(ptr QSqlIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndex_PTR().Pointer()
	}
	return nil
}

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = new(QSqlIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlIndex) QSqlIndex_PTR() *QSqlIndex {
	return ptr
}

func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	return NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex2(PointerFromQSqlIndex(other)))
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	return NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex(C.CString(cursorname), C.CString(name)))
}

func (ptr *QSqlIndex) Append(field QSqlField_ITF) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(ptr.Pointer(), PointerFromQSqlField(field), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) CursorName() string {
	defer qt.Recovering("QSqlIndex::cursorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_CursorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	defer qt.Recovering("QSqlIndex::isDescending")

	if ptr.Pointer() != nil {
		return C.QSqlIndex_IsDescending(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QSqlIndex) Name() string {
	defer qt.Recovering("QSqlIndex::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	defer qt.Recovering("QSqlIndex::setCursorName")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetCursorName(ptr.Pointer(), C.CString(cursorName))
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	defer qt.Recovering("QSqlIndex::setDescending")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(ptr.Pointer(), C.int(i), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	defer qt.Recovering("QSqlIndex::setName")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	defer qt.Recovering("QSqlIndex::~QSqlIndex")

	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(ptr.Pointer())
	}
}
