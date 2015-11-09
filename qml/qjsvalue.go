package qml

//#include "qjsvalue.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (p *QJSValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValue(ptr QJSValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValue_PTR().Pointer()
	}
	return nil
}

func NewQJSValueFromPointer(ptr unsafe.Pointer) *QJSValue {
	var n = new(QJSValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue = QJSValue__SpecialValue(1)
)

func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue(C.int(value)))
}

func NewQJSValue4(value bool) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.int(qt.GoBoolToInt(value))))
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
}

func NewQJSValue8(value string) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue8(C.CString(value)))
}

func NewQJSValue10(value string) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue10(C.CString(value)))
}

func NewQJSValue5(value int) *QJSValue {
	return NewQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(value)))
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_DeleteProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) Equals(other QJSValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_Equals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_HasOwnProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_HasProperty(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsCallable() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsCallable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) Property(name string) *QJSValue {
	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSValue_Property(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QJSValue) Prototype() *QJSValue {
	if ptr.Pointer() != nil {
		return NewQJSValueFromPointer(C.QJSValue_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetPrototype(ptr.Pointer(), PointerFromQJSValue(prototype))
	}
}

func (ptr *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_StrictlyEquals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) ToDateTime() *core.QDateTime {
	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QJSValue_ToDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QJSValue_ToQObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJSValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) DestroyQJSValue() {
	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(ptr.Pointer())
	}
}
