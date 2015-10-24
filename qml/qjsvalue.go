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

type QJSValueITF interface {
	QJSValuePTR() *QJSValue
}

func (p *QJSValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValue(ptr QJSValueITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValuePTR().Pointer()
	}
	return nil
}

func QJSValueFromPointer(ptr unsafe.Pointer) *QJSValue {
	var n = new(QJSValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJSValue) QJSValuePTR() *QJSValue {
	return ptr
}

//QJSValue::SpecialValue
type QJSValue__SpecialValue int

var (
	QJSValue__NullValue      = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue = QJSValue__SpecialValue(1)
)

func NewQJSValue3(other QJSValueITF) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue3(C.QtObjectPtr(PointerFromQJSValue(other)))))
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue(C.int(value))))
}

func NewQJSValue4(value bool) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue4(C.int(qt.GoBoolToInt(value)))))
}

func NewQJSValue2(other QJSValueITF) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue2(C.QtObjectPtr(PointerFromQJSValue(other)))))
}

func NewQJSValue9(value core.QLatin1StringITF) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue9(C.QtObjectPtr(core.PointerFromQLatin1String(value)))))
}

func NewQJSValue8(value string) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue8(C.CString(value))))
}

func NewQJSValue10(value string) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue10(C.CString(value))))
}

func NewQJSValue5(value int) *QJSValue {
	return QJSValueFromPointer(unsafe.Pointer(C.QJSValue_NewQJSValue5(C.int(value))))
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_DeleteProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) Equals(other QJSValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_Equals(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJSValue(other))) != 0
	}
	return false
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_HasOwnProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_HasProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QJSValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsArray(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsCallable() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsCallable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsDate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsNumber(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsQObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsRegExp(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsString(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsUndefined(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_IsVariant(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) SetProperty(name string, value QJSValueITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQJSValue(value)))
	}
}

func (ptr *QJSValue) SetPrototype(prototype QJSValueITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetPrototype(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJSValue(prototype)))
	}
}

func (ptr *QJSValue) StrictlyEquals(other QJSValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_StrictlyEquals(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJSValue(other))) != 0
	}
	return false
}

func (ptr *QJSValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return C.QJSValue_ToBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QJSValue_ToQObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QJSValue) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJSValue_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QJSValue) ToVariant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJSValue_ToVariant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QJSValue) DestroyQJSValue() {
	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(C.QtObjectPtr(ptr.Pointer()))
	}
}
