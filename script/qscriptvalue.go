package script

//#include "qscriptvalue.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScriptValue struct {
	ptr unsafe.Pointer
}

type QScriptValueITF interface {
	QScriptValuePTR() *QScriptValue
}

func (p *QScriptValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptValue(ptr QScriptValueITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValuePTR().Pointer()
	}
	return nil
}

func QScriptValueFromPointer(ptr unsafe.Pointer) *QScriptValue {
	var n = new(QScriptValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptValue) QScriptValuePTR() *QScriptValue {
	return ptr
}

//QScriptValue::PropertyFlag
type QScriptValue__PropertyFlag int

var (
	QScriptValue__ReadOnly          = QScriptValue__PropertyFlag(0x00000001)
	QScriptValue__Undeletable       = QScriptValue__PropertyFlag(0x00000002)
	QScriptValue__SkipInEnumeration = QScriptValue__PropertyFlag(0x00000004)
	QScriptValue__PropertyGetter    = QScriptValue__PropertyFlag(0x00000008)
	QScriptValue__PropertySetter    = QScriptValue__PropertyFlag(0x00000010)
	QScriptValue__QObjectMember     = QScriptValue__PropertyFlag(0x00000020)
	QScriptValue__KeepExistingFlags = QScriptValue__PropertyFlag(0x00000800)
	QScriptValue__UserRange         = QScriptValue__PropertyFlag(0xff000000)
)

//QScriptValue::ResolveFlag
type QScriptValue__ResolveFlag int

var (
	QScriptValue__ResolveLocal     = QScriptValue__ResolveFlag(0x00)
	QScriptValue__ResolvePrototype = QScriptValue__ResolveFlag(0x01)
	QScriptValue__ResolveScope     = QScriptValue__ResolveFlag(0x02)
	QScriptValue__ResolveFull      = QScriptValue__ResolveFlag(QScriptValue__ResolvePrototype | QScriptValue__ResolveScope)
)

//QScriptValue::SpecialValue
type QScriptValue__SpecialValue int

var (
	QScriptValue__NullValue      = QScriptValue__SpecialValue(0)
	QScriptValue__UndefinedValue = QScriptValue__SpecialValue(1)
)

func NewQScriptValue() *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue()))
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue10(C.int(value))))
}

func NewQScriptValue11(value bool) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue11(C.int(qt.GoBoolToInt(value)))))
}

func NewQScriptValue16(value core.QLatin1StringITF) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue16(C.QtObjectPtr(core.PointerFromQLatin1String(value)))))
}

func NewQScriptValue2(other QScriptValueITF) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue2(C.QtObjectPtr(PointerFromQScriptValue(other)))))
}

func NewQScriptValue15(value string) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue15(C.CString(value))))
}

func NewQScriptValue17(value string) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue17(C.CString(value))))
}

func NewQScriptValue12(value int) *QScriptValue {
	return QScriptValueFromPointer(unsafe.Pointer(C.QScriptValue_NewQScriptValue12(C.int(value))))
}

func (ptr *QScriptValue) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		return QScriptEngineFromPointer(unsafe.Pointer(C.QScriptValue_Engine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptValue) Equals(other QScriptValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_Equals(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) InstanceOf(other QScriptValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_InstanceOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsArray(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsDate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsFunction() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsFunction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNumber(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQMetaObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQMetaObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsRegExp(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsString(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsUndefined(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsVariant(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) LessThan(other QScriptValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_LessThan(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptStringITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptString(name)), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) PropertyFlags(name string, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) ScriptClass() *QScriptClass {
	if ptr.Pointer() != nil {
		return QScriptClassFromPointer(unsafe.Pointer(C.QScriptValue_ScriptClass(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptValue) SetData(data QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(data)))
	}
}

func (ptr *QScriptValue) SetProperty2(name QScriptStringITF, value QScriptValueITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptString(name)), C.QtObjectPtr(PointerFromQScriptValue(value)), C.int(flags))
	}
}

func (ptr *QScriptValue) SetProperty(name string, value QScriptValueITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQScriptValue(value)), C.int(flags))
	}
}

func (ptr *QScriptValue) SetPrototype(prototype QScriptValueITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetPrototype(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(prototype)))
	}
}

func (ptr *QScriptValue) SetScriptClass(scriptClass QScriptClassITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetScriptClass(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptClass(scriptClass)))
	}
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValueITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_StrictlyEquals(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_ToBool(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.QMetaObjectFromPointer(unsafe.Pointer(C.QScriptValue_ToQMetaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QScriptValue_ToQObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScriptValue) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptValue_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptValue) ToVariant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptValue_ToVariant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QScriptValue) DestroyQScriptValue() {
	if ptr.Pointer() != nil {
		C.QScriptValue_DestroyQScriptValue(C.QtObjectPtr(ptr.Pointer()))
	}
}
