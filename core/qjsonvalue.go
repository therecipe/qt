package core

//#include "qjsonvalue.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QJsonValue struct {
	ptr unsafe.Pointer
}

type QJsonValue_ITF interface {
	QJsonValue_PTR() *QJsonValue
}

func (p *QJsonValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonValue(ptr QJsonValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonValue_PTR().Pointer()
	}
	return nil
}

func NewQJsonValueFromPointer(ptr unsafe.Pointer) *QJsonValue {
	var n = new(QJsonValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonValue) QJsonValue_PTR() *QJsonValue {
	return ptr
}

//QJsonValue::Type
type QJsonValue__Type int64

const (
	QJsonValue__Null      = QJsonValue__Type(0x0)
	QJsonValue__Bool      = QJsonValue__Type(0x1)
	QJsonValue__Double    = QJsonValue__Type(0x2)
	QJsonValue__String    = QJsonValue__Type(0x3)
	QJsonValue__Array     = QJsonValue__Type(0x4)
	QJsonValue__Object    = QJsonValue__Type(0x5)
	QJsonValue__Undefined = QJsonValue__Type(0x80)
)

func NewQJsonValue5(s QLatin1String_ITF) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue5(PointerFromQLatin1String(s)))
}

func NewQJsonValue(ty QJsonValue__Type) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue(C.int(ty)))
}

func NewQJsonValue2(b bool) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue2(C.int(qt.GoBoolToInt(b))))
}

func NewQJsonValue7(a QJsonArray_ITF) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue7(PointerFromQJsonArray(a)))
}

func NewQJsonValue8(o QJsonObject_ITF) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue8(PointerFromQJsonObject(o)))
}

func NewQJsonValue9(other QJsonValue_ITF) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue9(PointerFromQJsonValue(other)))
}

func NewQJsonValue4(s string) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue4(C.CString(s)))
}

func NewQJsonValue6(s string) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue6(C.CString(s)))
}

func NewQJsonValue12(n int) *QJsonValue {
	return NewQJsonValueFromPointer(C.QJsonValue_NewQJsonValue12(C.int(n)))
}

func (ptr *QJsonValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsDouble() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsDouble(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonValue) ToArray2() *QJsonArray {
	if ptr.Pointer() != nil {
		return NewQJsonArrayFromPointer(C.QJsonValue_ToArray2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonValue) ToArray(defaultValue QJsonArray_ITF) *QJsonArray {
	if ptr.Pointer() != nil {
		return NewQJsonArrayFromPointer(C.QJsonValue_ToArray(ptr.Pointer(), PointerFromQJsonArray(defaultValue)))
	}
	return nil
}

func (ptr *QJsonValue) ToBool(defaultValue bool) bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_ToBool(ptr.Pointer(), C.int(qt.GoBoolToInt(defaultValue))) != 0
	}
	return false
}

func (ptr *QJsonValue) ToInt(defaultValue int) int {
	if ptr.Pointer() != nil {
		return int(C.QJsonValue_ToInt(ptr.Pointer(), C.int(defaultValue)))
	}
	return 0
}

func (ptr *QJsonValue) ToObject2() *QJsonObject {
	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QJsonValue_ToObject2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonValue) ToObject(defaultValue QJsonObject_ITF) *QJsonObject {
	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QJsonValue_ToObject(ptr.Pointer(), PointerFromQJsonObject(defaultValue)))
	}
	return nil
}

func (ptr *QJsonValue) ToString(defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJsonValue_ToString(ptr.Pointer(), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QJsonValue) ToVariant() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QJsonValue_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonValue) Type() QJsonValue__Type {
	if ptr.Pointer() != nil {
		return QJsonValue__Type(C.QJsonValue_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJsonValue) DestroyQJsonValue() {
	if ptr.Pointer() != nil {
		C.QJsonValue_DestroyQJsonValue(ptr.Pointer())
	}
}
