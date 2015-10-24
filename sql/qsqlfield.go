package sql

//#include "qsqlfield.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlField struct {
	ptr unsafe.Pointer
}

type QSqlFieldITF interface {
	QSqlFieldPTR() *QSqlField
}

func (p *QSqlField) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlField) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlField(ptr QSqlFieldITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlFieldPTR().Pointer()
	}
	return nil
}

func QSqlFieldFromPointer(ptr unsafe.Pointer) *QSqlField {
	var n = new(QSqlField)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlField) QSqlFieldPTR() *QSqlField {
	return ptr
}

//QSqlField::RequiredStatus
type QSqlField__RequiredStatus int

var (
	QSqlField__Unknown  = QSqlField__RequiredStatus(-1)
	QSqlField__Optional = QSqlField__RequiredStatus(0)
	QSqlField__Required = QSqlField__RequiredStatus(1)
)

func NewQSqlField2(other QSqlFieldITF) *QSqlField {
	return QSqlFieldFromPointer(unsafe.Pointer(C.QSqlField_NewQSqlField2(C.QtObjectPtr(PointerFromQSqlField(other)))))
}

func (ptr *QSqlField) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlField_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlField) DefaultValue() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_DefaultValue(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlField) IsAutoValue() bool {
	if ptr.Pointer() != nil {
		return C.QSqlField_IsAutoValue(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsGenerated() bool {
	if ptr.Pointer() != nil {
		return C.QSqlField_IsGenerated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSqlField_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QSqlField_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlField_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlField) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlField_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlField) Precision() int {
	if ptr.Pointer() != nil {
		return int(C.QSqlField_Precision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {
	if ptr.Pointer() != nil {
		return QSqlField__RequiredStatus(C.QSqlField_RequiredStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetAutoValue(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(autoVal)))
	}
}

func (ptr *QSqlField) SetDefaultValue(value string) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetDefaultValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QSqlField) SetGenerated(gen bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetGenerated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(gen)))
	}
}

func (ptr *QSqlField) SetLength(fieldLength int) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetLength(C.QtObjectPtr(ptr.Pointer()), C.int(fieldLength))
	}
}

func (ptr *QSqlField) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSqlField) SetPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetPrecision(C.QtObjectPtr(ptr.Pointer()), C.int(precision))
	}
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(readOnly)))
	}
}

func (ptr *QSqlField) SetRequired(required bool) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetRequired(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(required)))
	}
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetRequiredStatus(C.QtObjectPtr(ptr.Pointer()), C.int(required))
	}
}

func (ptr *QSqlField) SetValue(value string) {
	if ptr.Pointer() != nil {
		C.QSqlField_SetValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QSqlField) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlField) DestroyQSqlField() {
	if ptr.Pointer() != nil {
		C.QSqlField_DestroyQSqlField(C.QtObjectPtr(ptr.Pointer()))
	}
}
