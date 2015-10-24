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

type QJsonValueITF interface {
	QJsonValuePTR() *QJsonValue
}

func (p *QJsonValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonValue(ptr QJsonValueITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonValuePTR().Pointer()
	}
	return nil
}

func QJsonValueFromPointer(ptr unsafe.Pointer) *QJsonValue {
	var n = new(QJsonValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonValue) QJsonValuePTR() *QJsonValue {
	return ptr
}

//QJsonValue::Type
type QJsonValue__Type int

var (
	QJsonValue__Null      = QJsonValue__Type(0x0)
	QJsonValue__Bool      = QJsonValue__Type(0x1)
	QJsonValue__Double    = QJsonValue__Type(0x2)
	QJsonValue__String    = QJsonValue__Type(0x3)
	QJsonValue__Array     = QJsonValue__Type(0x4)
	QJsonValue__Object    = QJsonValue__Type(0x5)
	QJsonValue__Undefined = QJsonValue__Type(0x80)
)

func NewQJsonValue5(s QLatin1StringITF) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue5(C.QtObjectPtr(PointerFromQLatin1String(s)))))
}

func NewQJsonValue(ty QJsonValue__Type) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue(C.int(ty))))
}

func NewQJsonValue2(b bool) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue2(C.int(qt.GoBoolToInt(b)))))
}

func NewQJsonValue7(a QJsonArrayITF) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue7(C.QtObjectPtr(PointerFromQJsonArray(a)))))
}

func NewQJsonValue8(o QJsonObjectITF) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue8(C.QtObjectPtr(PointerFromQJsonObject(o)))))
}

func NewQJsonValue9(other QJsonValueITF) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue9(C.QtObjectPtr(PointerFromQJsonValue(other)))))
}

func NewQJsonValue4(s string) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue4(C.CString(s))))
}

func NewQJsonValue6(s string) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue6(C.CString(s))))
}

func NewQJsonValue12(n int) *QJsonValue {
	return QJsonValueFromPointer(unsafe.Pointer(C.QJsonValue_NewQJsonValue12(C.int(n))))
}

func (ptr *QJsonValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsArray(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsDouble() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsDouble(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsString(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_IsUndefined(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonValue) ToBool(defaultValue bool) bool {
	if ptr.Pointer() != nil {
		return C.QJsonValue_ToBool(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(defaultValue))) != 0
	}
	return false
}

func (ptr *QJsonValue) ToInt(defaultValue int) int {
	if ptr.Pointer() != nil {
		return int(C.QJsonValue_ToInt(C.QtObjectPtr(ptr.Pointer()), C.int(defaultValue)))
	}
	return 0
}

func (ptr *QJsonValue) ToString(defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJsonValue_ToString(C.QtObjectPtr(ptr.Pointer()), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QJsonValue) ToVariant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJsonValue_ToVariant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QJsonValue) Type() QJsonValue__Type {
	if ptr.Pointer() != nil {
		return QJsonValue__Type(C.QJsonValue_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJsonValue) DestroyQJsonValue() {
	if ptr.Pointer() != nil {
		C.QJsonValue_DestroyQJsonValue(C.QtObjectPtr(ptr.Pointer()))
	}
}
