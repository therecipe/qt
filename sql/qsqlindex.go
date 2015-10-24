package sql

//#include "qsqlindex.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndexITF interface {
	QSqlRecordITF
	QSqlIndexPTR() *QSqlIndex
}

func PointerFromQSqlIndex(ptr QSqlIndexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndexPTR().Pointer()
	}
	return nil
}

func QSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = new(QSqlIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlIndex) QSqlIndexPTR() *QSqlIndex {
	return ptr
}

func NewQSqlIndex2(other QSqlIndexITF) *QSqlIndex {
	return QSqlIndexFromPointer(unsafe.Pointer(C.QSqlIndex_NewQSqlIndex2(C.QtObjectPtr(PointerFromQSqlIndex(other)))))
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	return QSqlIndexFromPointer(unsafe.Pointer(C.QSqlIndex_NewQSqlIndex(C.CString(cursorname), C.CString(name))))
}

func (ptr *QSqlIndex) Append(field QSqlFieldITF) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_Append(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSqlField(field)))
	}
}

func (ptr *QSqlIndex) Append2(field QSqlFieldITF, desc bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSqlField(field)), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) CursorName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_CursorName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlIndex_IsDescending(C.QtObjectPtr(ptr.Pointer()), C.int(i)) != 0
	}
	return false
}

func (ptr *QSqlIndex) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_SetCursorName(C.QtObjectPtr(ptr.Pointer()), C.CString(cursorName))
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(C.QtObjectPtr(ptr.Pointer()), C.int(i), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlIndex_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(C.QtObjectPtr(ptr.Pointer()))
	}
}
