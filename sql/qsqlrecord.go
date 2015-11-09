package sql

//#include "qsqlrecord.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlRecord struct {
	ptr unsafe.Pointer
}

type QSqlRecord_ITF interface {
	QSqlRecord_PTR() *QSqlRecord
}

func (p *QSqlRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlRecord(ptr QSqlRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func NewQSqlRecordFromPointer(ptr unsafe.Pointer) *QSqlRecord {
	var n = new(QSqlRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlRecord) QSqlRecord_PTR() *QSqlRecord {
	return ptr
}

func NewQSqlRecord() *QSqlRecord {
	return NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord())
}

func NewQSqlRecord2(other QSqlRecord_ITF) *QSqlRecord {
	return NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord2(PointerFromQSqlRecord(other)))
}

func (ptr *QSqlRecord) Append(field QSqlField_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) ClearValues() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_ClearValues(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) Contains(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_Contains(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlRecord) FieldName(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_FieldName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QSqlRecord) IndexOf(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlRecord_IndexOf(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QSqlRecord) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated2(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull(name string) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull2(index int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull2(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated(ptr.Pointer(), C.CString(name), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated2(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(generated)))
	}
}

func (ptr *QSqlRecord) SetNull2(name string) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull2(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlRecord) SetNull(index int) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSqlRecord) SetValue2(name string, val core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue2(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) SetValue(index int, val core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue(ptr.Pointer(), C.int(index), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) Value2(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRecord_Value2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QSqlRecord) Value(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRecord_Value(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {
	if ptr.Pointer() != nil {
		C.QSqlRecord_DestroyQSqlRecord(ptr.Pointer())
	}
}
