package sql

//#include "qsqlrecord.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlRecord struct {
	ptr unsafe.Pointer
}

type QSqlRecordITF interface {
	QSqlRecordPTR() *QSqlRecord
}

func (p *QSqlRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRecord(ptr QSqlRecordITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecordPTR().Pointer()
	}
	return nil
}

func QSqlRecordFromPointer(ptr unsafe.Pointer) *QSqlRecord {
	var n = new(QSqlRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlRecord) QSqlRecordPTR() *QSqlRecord {
	return ptr
}

func NewQSqlRecord() *QSqlRecord {
	return QSqlRecordFromPointer(unsafe.Pointer(C.QSqlRecord_NewQSqlRecord()))
}

func NewQSqlRecord2(other QSqlRecordITF) *QSqlRecord {
	return QSqlRecordFromPointer(unsafe.Pointer(C.QSqlRecord_NewQSqlRecord2(C.QtObjectPtr(PointerFromQSqlRecord(other)))))
}

func (ptr *QSqlRecord) Append(field QSqlFieldITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Append(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSqlField(field)))
	}
}

func (ptr *QSqlRecord) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlRecord) ClearValues() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_ClearValues(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlRecord) Contains(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlRecord) FieldName(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_FieldName(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QSqlRecord) IndexOf(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QSqlRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated2(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull2(index int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull2(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetNull2(name string) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull2(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSqlRecord) SetNull(index int) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QSqlRecord) SetValue2(name string, val string) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue2(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(val))
	}
}

func (ptr *QSqlRecord) SetValue(index int, val string) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(val))
	}
}

func (ptr *QSqlRecord) Value2(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_Value2(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QSqlRecord) Value(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_Value(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_DestroyQSqlRecord(C.QtObjectPtr(ptr.Pointer()))
	}
}
