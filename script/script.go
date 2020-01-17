// +build !minimal

package script

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtScript_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtScript_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QScriptClass struct {
	ptr unsafe.Pointer
}

type QScriptClass_ITF interface {
	QScriptClass_PTR() *QScriptClass
}

func (ptr *QScriptClass) QScriptClass_PTR() *QScriptClass {
	return ptr
}

func (ptr *QScriptClass) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptClass) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptClass(ptr QScriptClass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClass_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassFromPointer(ptr unsafe.Pointer) (n *QScriptClass) {
	n = new(QScriptClass)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptClass__QueryFlag
//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int64

const (
	QScriptClass__HandlesReadAccess  QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x02)
)

//go:generate stringer -type=QScriptClass__Extension
//QScriptClass::Extension
type QScriptClass__Extension int64

const (
	QScriptClass__Callable    QScriptClass__Extension = QScriptClass__Extension(0)
	QScriptClass__HasInstance QScriptClass__Extension = QScriptClass__Extension(1)
)

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {
	return NewQScriptClassFromPointer(C.QScriptClass_NewQScriptClass(PointerFromQScriptEngine(engine)))
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptEngineFromPointer(C.QScriptClass_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_Extension
func callbackQScriptClass_Extension(ptr unsafe.Pointer, extension C.longlong, argument unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQVariant((*(*func(QScriptClass__Extension, *core.QVariant) *core.QVariant)(signal))(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptClassFromPointer(ptr).ExtensionDefault(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptClass) ConnectExtension(f func(extension QScriptClass__Extension, argument *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			f := func(extension QScriptClass__Extension, argument *core.QVariant) *core.QVariant {
				(*(*func(QScriptClass__Extension, *core.QVariant) *core.QVariant)(signal))(extension, argument)
				return f(extension, argument)
			}
			qt.ConnectSignal(ptr.Pointer(), "extension", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "extension")
	}
}

func (ptr *QScriptClass) Extension(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QScriptClass_Extension(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) ExtensionDefault(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QScriptClass_ExtensionDefault(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_Name
func callbackQScriptClass_Name(ptr unsafe.Pointer) C.struct_QtScript_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtScript_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQScriptClassFromPointer(ptr).NameDefault()
	return C.struct_QtScript_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QScriptClass) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QScriptClass) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptClass_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptClass) NameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptClass_NameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQScriptClass_NewIterator
func callbackQScriptClass_NewIterator(ptr unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "newIterator"); signal != nil {
		return PointerFromQScriptClassPropertyIterator((*(*func(*QScriptValue) *QScriptClassPropertyIterator)(signal))(NewQScriptValueFromPointer(object)))
	}

	return PointerFromQScriptClassPropertyIterator(NewQScriptClassFromPointer(ptr).NewIteratorDefault(NewQScriptValueFromPointer(object)))
}

func (ptr *QScriptClass) ConnectNewIterator(f func(object *QScriptValue) *QScriptClassPropertyIterator) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newIterator"); signal != nil {
			f := func(object *QScriptValue) *QScriptClassPropertyIterator {
				(*(*func(*QScriptValue) *QScriptClassPropertyIterator)(signal))(object)
				return f(object)
			}
			qt.ConnectSignal(ptr.Pointer(), "newIterator", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newIterator", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectNewIterator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newIterator")
	}
}

func (ptr *QScriptClass) NewIterator(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIterator(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

func (ptr *QScriptClass) NewIteratorDefault(object QScriptValue_ITF) *QScriptClassPropertyIterator {
	if ptr.Pointer() != nil {
		return NewQScriptClassPropertyIteratorFromPointer(C.QScriptClass_NewIteratorDefault(ptr.Pointer(), PointerFromQScriptValue(object)))
	}
	return nil
}

//export callbackQScriptClass_Property
func callbackQScriptClass_Property(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, id C.uint) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "property"); signal != nil {
		return PointerFromQScriptValue((*(*func(*QScriptValue, *QScriptString, uint) *QScriptValue)(signal))(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
	}

	return PointerFromQScriptValue(NewQScriptClassFromPointer(ptr).PropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectProperty(f func(object *QScriptValue, name *QScriptString, id uint) *QScriptValue) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			f := func(object *QScriptValue, name *QScriptString, id uint) *QScriptValue {
				(*(*func(*QScriptValue, *QScriptString, uint) *QScriptValue)(signal))(object, name, id)
				return f(object, name, id)
			}
			qt.ConnectSignal(ptr.Pointer(), "property", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "property")
	}
}

func (ptr *QScriptClass) Property(object QScriptValue_ITF, name QScriptString_ITF, id uint) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptClass_Property(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) PropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptClass_PropertyDefault(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_PropertyFlags
func callbackQScriptClass_PropertyFlags(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, id C.uint) C.longlong {
	if signal := qt.GetSignal(ptr, "propertyFlags"); signal != nil {
		return C.longlong((*(*func(*QScriptValue, *QScriptString, uint) QScriptValue__PropertyFlag)(signal))(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
	}

	return C.longlong(NewQScriptClassFromPointer(ptr).PropertyFlagsDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectPropertyFlags(f func(object *QScriptValue, name *QScriptString, id uint) QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "propertyFlags"); signal != nil {
			f := func(object *QScriptValue, name *QScriptString, id uint) QScriptValue__PropertyFlag {
				(*(*func(*QScriptValue, *QScriptString, uint) QScriptValue__PropertyFlag)(signal))(object, name, id)
				return f(object, name, id)
			}
			qt.ConnectSignal(ptr.Pointer(), "propertyFlags", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "propertyFlags", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectPropertyFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "propertyFlags")
	}
}

func (ptr *QScriptClass) PropertyFlags(object QScriptValue_ITF, name QScriptString_ITF, id uint) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptClass_PropertyFlags(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
	}
	return 0
}

func (ptr *QScriptClass) PropertyFlagsDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptClass_PropertyFlagsDefault(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
	}
	return 0
}

//export callbackQScriptClass_Prototype
func callbackQScriptClass_Prototype(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "prototype"); signal != nil {
		return PointerFromQScriptValue((*(*func() *QScriptValue)(signal))())
	}

	return PointerFromQScriptValue(NewQScriptClassFromPointer(ptr).PrototypeDefault())
}

func (ptr *QScriptClass) ConnectPrototype(f func() *QScriptValue) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prototype"); signal != nil {
			f := func() *QScriptValue {
				(*(*func() *QScriptValue)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "prototype", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prototype", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectPrototype() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "prototype")
	}
}

func (ptr *QScriptClass) Prototype() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptClass_Prototype(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) PrototypeDefault() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptClass_PrototypeDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_QueryProperty
func callbackQScriptClass_QueryProperty(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, flags C.longlong, id C.uint) C.longlong {
	if signal := qt.GetSignal(ptr, "queryProperty"); signal != nil {
		return C.longlong((*(*func(*QScriptValue, *QScriptString, QScriptClass__QueryFlag, uint) QScriptClass__QueryFlag)(signal))(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), QScriptClass__QueryFlag(flags), uint(uint32(id))))
	}

	return C.longlong(NewQScriptClassFromPointer(ptr).QueryPropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), QScriptClass__QueryFlag(flags), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectQueryProperty(f func(object *QScriptValue, name *QScriptString, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "queryProperty"); signal != nil {
			f := func(object *QScriptValue, name *QScriptString, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {
				(*(*func(*QScriptValue, *QScriptString, QScriptClass__QueryFlag, uint) QScriptClass__QueryFlag)(signal))(object, name, flags, id)
				return f(object, name, flags, id)
			}
			qt.ConnectSignal(ptr.Pointer(), "queryProperty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "queryProperty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectQueryProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "queryProperty")
	}
}

func (ptr *QScriptClass) QueryProperty(object QScriptValue_ITF, name QScriptString_ITF, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {
	if ptr.Pointer() != nil {
		return QScriptClass__QueryFlag(C.QScriptClass_QueryProperty(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.longlong(flags), C.uint(uint32(id))))
	}
	return 0
}

func (ptr *QScriptClass) QueryPropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {
	if ptr.Pointer() != nil {
		return QScriptClass__QueryFlag(C.QScriptClass_QueryPropertyDefault(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.longlong(flags), C.uint(uint32(id))))
	}
	return 0
}

//export callbackQScriptClass_SetProperty
func callbackQScriptClass_SetProperty(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, id C.uint, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		(*(*func(*QScriptValue, *QScriptString, uint, *QScriptValue))(signal))(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id)), NewQScriptValueFromPointer(value))
	} else {
		NewQScriptClassFromPointer(ptr).SetPropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id)), NewQScriptValueFromPointer(value))
	}
}

func (ptr *QScriptClass) ConnectSetProperty(f func(object *QScriptValue, name *QScriptString, id uint, value *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			f := func(object *QScriptValue, name *QScriptString, id uint, value *QScriptValue) {
				(*(*func(*QScriptValue, *QScriptString, uint, *QScriptValue))(signal))(object, name, id, value)
				f(object, name, id, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "setProperty", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QScriptClass) SetProperty(object QScriptValue_ITF, name QScriptString_ITF, id uint, value QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptClass_SetProperty(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id)), PointerFromQScriptValue(value))
	}
}

func (ptr *QScriptClass) SetPropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint, value QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptClass_SetPropertyDefault(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id)), PointerFromQScriptValue(value))
	}
}

//export callbackQScriptClass_SupportsExtension
func callbackQScriptClass_SupportsExtension(ptr unsafe.Pointer, extension C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "supportsExtension"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QScriptClass__Extension) bool)(signal))(QScriptClass__Extension(extension)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptClassFromPointer(ptr).SupportsExtensionDefault(QScriptClass__Extension(extension)))))
}

func (ptr *QScriptClass) ConnectSupportsExtension(f func(extension QScriptClass__Extension) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportsExtension"); signal != nil {
			f := func(extension QScriptClass__Extension) bool {
				(*(*func(QScriptClass__Extension) bool)(signal))(extension)
				return f(extension)
			}
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectSupportsExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportsExtension")
	}
}

func (ptr *QScriptClass) SupportsExtension(extension QScriptClass__Extension) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptClass_SupportsExtension(ptr.Pointer(), C.longlong(extension))) != 0
	}
	return false
}

func (ptr *QScriptClass) SupportsExtensionDefault(extension QScriptClass__Extension) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptClass_SupportsExtensionDefault(ptr.Pointer(), C.longlong(extension))) != 0
	}
	return false
}

//export callbackQScriptClass_DestroyQScriptClass
func callbackQScriptClass_DestroyQScriptClass(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptClass"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptClassFromPointer(ptr).DestroyQScriptClassDefault()
	}
}

func (ptr *QScriptClass) ConnectDestroyQScriptClass(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptClass"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScriptClass", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptClass", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptClass) DisconnectDestroyQScriptClass() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScriptClass")
	}
}

func (ptr *QScriptClass) DestroyQScriptClass() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptClass_DestroyQScriptClass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptClass) DestroyQScriptClassDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptClass_DestroyQScriptClassDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptClassPropertyIterator struct {
	ptr unsafe.Pointer
}

type QScriptClassPropertyIterator_ITF interface {
	QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator
}

func (ptr *QScriptClassPropertyIterator) QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator {
	return ptr
}

func (ptr *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptClassPropertyIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptClassPropertyIterator(ptr QScriptClassPropertyIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPropertyIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) (n *QScriptClassPropertyIterator) {
	n = new(QScriptClassPropertyIterator)
	n.SetPointer(ptr)
	return
}

type QScriptContext struct {
	ptr unsafe.Pointer
}

type QScriptContext_ITF interface {
	QScriptContext_PTR() *QScriptContext
}

func (ptr *QScriptContext) QScriptContext_PTR() *QScriptContext {
	return ptr
}

func (ptr *QScriptContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptContext(ptr QScriptContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContext_PTR().Pointer()
	}
	return nil
}

func NewQScriptContextFromPointer(ptr unsafe.Pointer) (n *QScriptContext) {
	n = new(QScriptContext)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptContext__ExecutionState
//QScriptContext::ExecutionState
type QScriptContext__ExecutionState int64

const (
	QScriptContext__NormalState    QScriptContext__ExecutionState = QScriptContext__ExecutionState(0)
	QScriptContext__ExceptionState QScriptContext__ExecutionState = QScriptContext__ExecutionState(1)
)

//go:generate stringer -type=QScriptContext__Error
//QScriptContext::Error
type QScriptContext__Error int64

const (
	QScriptContext__UnknownError   QScriptContext__Error = QScriptContext__Error(0)
	QScriptContext__ReferenceError QScriptContext__Error = QScriptContext__Error(1)
	QScriptContext__SyntaxError    QScriptContext__Error = QScriptContext__Error(2)
	QScriptContext__TypeError      QScriptContext__Error = QScriptContext__Error(3)
	QScriptContext__RangeError     QScriptContext__Error = QScriptContext__Error(4)
	QScriptContext__URIError       QScriptContext__Error = QScriptContext__Error(5)
)

func (ptr *QScriptContext) ActivationObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ActivationObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) Argument(index int) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_Argument(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ArgumentCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContext_ArgumentCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContext) ArgumentsObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ArgumentsObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) Backtrace() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptContext_Backtrace(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QScriptContext) Callee() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_Callee(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptEngineFromPointer(C.QScriptContext_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) IsCalledAsConstructor() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptContext_IsCalledAsConstructor(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptContext) ParentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) SetActivationObject(activation QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptContext_SetActivationObject(ptr.Pointer(), PointerFromQScriptValue(activation))
	}
}

func (ptr *QScriptContext) SetThisObject(thisObject QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptContext_SetThisObject(ptr.Pointer(), PointerFromQScriptValue(thisObject))
	}
}

func (ptr *QScriptContext) State() QScriptContext__ExecutionState {
	if ptr.Pointer() != nil {
		return QScriptContext__ExecutionState(C.QScriptContext_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContext) ThisObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ThisObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ThrowError(error QScriptContext__Error, text string) *QScriptValue {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ThrowError(ptr.Pointer(), C.longlong(error), C.struct_QtScript_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ThrowError2(text string) *QScriptValue {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ThrowError2(ptr.Pointer(), C.struct_QtScript_PackedString{data: textC, len: C.longlong(len(text))}))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ThrowValue(value QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptContext_ThrowValue(ptr.Pointer(), PointerFromQScriptValue(value)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptContext_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContext) DestroyQScriptContext() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptContext_DestroyQScriptContext(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptContextInfo struct {
	ptr unsafe.Pointer
}

type QScriptContextInfo_ITF interface {
	QScriptContextInfo_PTR() *QScriptContextInfo
}

func (ptr *QScriptContextInfo) QScriptContextInfo_PTR() *QScriptContextInfo {
	return ptr
}

func (ptr *QScriptContextInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptContextInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptContextInfo(ptr QScriptContextInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptContextInfo_PTR().Pointer()
	}
	return nil
}

func NewQScriptContextInfoFromPointer(ptr unsafe.Pointer) (n *QScriptContextInfo) {
	n = new(QScriptContextInfo)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptContextInfo__FunctionType
//QScriptContextInfo::FunctionType
type QScriptContextInfo__FunctionType int64

const (
	QScriptContextInfo__ScriptFunction     QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(0)
	QScriptContextInfo__QtFunction         QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(1)
	QScriptContextInfo__QtPropertyFunction QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(2)
	QScriptContextInfo__NativeFunction     QScriptContextInfo__FunctionType = QScriptContextInfo__FunctionType(3)
)

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {
	tmpValue := NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo(PointerFromQScriptContext(context)))
	qt.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {
	tmpValue := NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo2(PointerFromQScriptContextInfo(other)))
	qt.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func NewQScriptContextInfo3() *QScriptContextInfo {
	tmpValue := NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo3())
	qt.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func (ptr *QScriptContextInfo) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptContextInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionEndLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContextInfo_FunctionEndLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionMetaIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContextInfo_FunctionMetaIndex(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptContextInfo_FunctionName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptContextInfo_FunctionParameterNames(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContextInfo_FunctionStartLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {
	if ptr.Pointer() != nil {
		return QScriptContextInfo__FunctionType(C.QScriptContextInfo_FunctionType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptContextInfo_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptContextInfo) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContextInfo_LineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptContextInfo) ScriptId() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QScriptContextInfo_ScriptId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptContextInfo_DestroyQScriptContextInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptEngine struct {
	core.QObject
}

type QScriptEngine_ITF interface {
	core.QObject_ITF
	QScriptEngine_PTR() *QScriptEngine
}

func (ptr *QScriptEngine) QScriptEngine_PTR() *QScriptEngine {
	return ptr
}

func (ptr *QScriptEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScriptEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScriptEngine(ptr QScriptEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngine_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineFromPointer(ptr unsafe.Pointer) (n *QScriptEngine) {
	n = new(QScriptEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptEngine__ValueOwnership
//QScriptEngine::ValueOwnership
type QScriptEngine__ValueOwnership int64

const (
	QScriptEngine__QtOwnership     QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(0)
	QScriptEngine__ScriptOwnership QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(1)
	QScriptEngine__AutoOwnership   QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(2)
)

//go:generate stringer -type=QScriptEngine__QObjectWrapOption
//QScriptEngine::QObjectWrapOption
type QScriptEngine__QObjectWrapOption int64

const (
	QScriptEngine__ExcludeChildObjects         QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0001)
	QScriptEngine__ExcludeSuperClassMethods    QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0002)
	QScriptEngine__ExcludeSuperClassProperties QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0004)
	QScriptEngine__ExcludeSuperClassContents   QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0006)
	QScriptEngine__SkipMethodsInEnumeration    QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0008)
	QScriptEngine__ExcludeDeleteLater          QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0010)
	QScriptEngine__ExcludeSlots                QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0020)
	QScriptEngine__AutoCreateDynamicProperties QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0100)
	QScriptEngine__PreferExistingWrapperObject QScriptEngine__QObjectWrapOption = QScriptEngine__QObjectWrapOption(0x0200)
)

func NewQScriptEngine() *QScriptEngine {
	tmpValue := NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {
	tmpValue := NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_AbortEvaluation(ptr.Pointer(), PointerFromQScriptValue(result))
	}
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {
	if ptr.Pointer() != nil {
		return NewQScriptEngineAgentFromPointer(C.QScriptEngine_Agent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) AvailableExtensions() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptEngine_AvailableExtensions(ptr.Pointer())))
	}
	return make([]string, 0)
}

func QScriptEngine_CheckSyntax(program string) *QScriptSyntaxCheckResult {
	var programC *C.char
	if program != "" {
		programC = C.CString(program)
		defer C.free(unsafe.Pointer(programC))
	}
	tmpValue := NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}))
	qt.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptEngine) CheckSyntax(program string) *QScriptSyntaxCheckResult {
	var programC *C.char
	if program != "" {
		programC = C.CString(program)
		defer C.free(unsafe.Pointer(programC))
	}
	tmpValue := NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}))
	qt.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptEngine) ClearExceptions() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ClearExceptions(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_CurrentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DefaultPrototype(metaTypeId int) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_DefaultPrototype(ptr.Pointer(), C.int(int32(metaTypeId))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate(program string, fileName string, lineNumber int) *QScriptValue {
	if ptr.Pointer() != nil {
		var programC *C.char
		if program != "" {
			programC = C.CString(program)
			defer C.free(unsafe.Pointer(programC))
		}
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_Evaluate(ptr.Pointer(), C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(lineNumber))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) Evaluate2(program QScriptProgram_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_Evaluate2(ptr.Pointer(), PointerFromQScriptProgram(program)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) GlobalObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_GlobalObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) HasUncaughtException() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngine_HasUncaughtException(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptEngine) ImportExtension(extension string) *QScriptValue {
	if ptr.Pointer() != nil {
		var extensionC *C.char
		if extension != "" {
			extensionC = C.CString(extension)
			defer C.free(unsafe.Pointer(extensionC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_ImportExtension(ptr.Pointer(), C.struct_QtScript_PackedString{data: extensionC, len: C.longlong(len(extension))}))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) ImportedExtensions() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptEngine_ImportedExtensions(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) IsEvaluating() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngine_IsEvaluating(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptEngine) NewArray(length uint) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewArray(ptr.Pointer(), C.uint(uint32(length))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewDate2(value core.QDateTime_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewDate2(ptr.Pointer(), core.PointerFromQDateTime(value)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewObject2(scriptClass QScriptClass_ITF, data QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewObject2(ptr.Pointer(), PointerFromQScriptClass(scriptClass), PointerFromQScriptValue(data)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQMetaObject(metaObject core.QMetaObject_ITF, ctor QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewQMetaObject(ptr.Pointer(), core.PointerFromQMetaObject(metaObject), PointerFromQScriptValue(ctor)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject(object core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object), C.longlong(ownership), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject2(scriptObject QScriptValue_ITF, qtObject core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewQObject2(ptr.Pointer(), PointerFromQScriptValue(scriptObject), core.PointerFromQObject(qtObject), C.longlong(ownership), C.longlong(options)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp(regexp core.QRegExp_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp(ptr.Pointer(), core.PointerFromQRegExp(regexp)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp2(pattern string, flags string) *QScriptValue {
	if ptr.Pointer() != nil {
		var patternC *C.char
		if pattern != "" {
			patternC = C.CString(pattern)
			defer C.free(unsafe.Pointer(patternC))
		}
		var flagsC *C.char
		if flags != "" {
			flagsC = C.CString(flags)
			defer C.free(unsafe.Pointer(flagsC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp2(ptr.Pointer(), C.struct_QtScript_PackedString{data: patternC, len: C.longlong(len(pattern))}, C.struct_QtScript_PackedString{data: flagsC, len: C.longlong(len(flags))}))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant(value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewVariant(ptr.Pointer(), core.PointerFromQVariant(value)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant2(object QScriptValue_ITF, value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NewVariant2(ptr.Pointer(), PointerFromQScriptValue(object), core.PointerFromQVariant(value)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NullValue() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_NullValue(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) PopContext() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_PopContext(ptr.Pointer())
	}
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptEngine_ProcessEventsInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) PushContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_PushContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) ReportAdditionalMemoryCost(size int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ReportAdditionalMemoryCost(ptr.Pointer(), C.int(int32(size)))
	}
}

func (ptr *QScriptEngine) SetAgent(agent QScriptEngineAgent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetAgent(ptr.Pointer(), PointerFromQScriptEngineAgent(agent))
	}
}

func (ptr *QScriptEngine) SetDefaultPrototype(metaTypeId int, prototype QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetDefaultPrototype(ptr.Pointer(), C.int(int32(metaTypeId)), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptEngine) SetGlobalObject(object QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetGlobalObject(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) SetProcessEventsInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SetProcessEventsInterval(ptr.Pointer(), C.int(int32(interval)))
	}
}

//export callbackQScriptEngine_SignalHandlerException
func callbackQScriptEngine_SignalHandlerException(ptr unsafe.Pointer, exception unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "signalHandlerException"); signal != nil {
		(*(*func(*QScriptValue))(signal))(NewQScriptValueFromPointer(exception))
	}

}

func (ptr *QScriptEngine) ConnectSignalHandlerException(f func(exception *QScriptValue)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "signalHandlerException") {
			C.QScriptEngine_ConnectSignalHandlerException(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "signalHandlerException")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "signalHandlerException"); signal != nil {
			f := func(exception *QScriptValue) {
				(*(*func(*QScriptValue))(signal))(exception)
				f(exception)
			}
			qt.ConnectSignal(ptr.Pointer(), "signalHandlerException", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "signalHandlerException", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngine) DisconnectSignalHandlerException() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectSignalHandlerException(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "signalHandlerException")
	}
}

func (ptr *QScriptEngine) SignalHandlerException(exception QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_SignalHandlerException(ptr.Pointer(), PointerFromQScriptValue(exception))
	}
}

func (ptr *QScriptEngine) ToObject(value QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_ToObject(ptr.Pointer(), PointerFromQScriptValue(value)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) ToStringHandle(str string) *QScriptString {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		tmpValue := NewQScriptStringFromPointer(C.QScriptEngine_ToStringHandle(ptr.Pointer(), C.struct_QtScript_PackedString{data: strC, len: C.longlong(len(str))}))
		qt.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtException() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_UncaughtException(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptEngine_UncaughtExceptionBacktrace(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptEngine_UncaughtExceptionLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) UndefinedValue() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptEngine_UndefinedValue(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngine_DestroyQScriptEngine
func callbackQScriptEngine_DestroyQScriptEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineFromPointer(ptr).DestroyQScriptEngineDefault()
	}
}

func (ptr *QScriptEngine) ConnectDestroyQScriptEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngine) DisconnectDestroyQScriptEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScriptEngine")
	}
}

func (ptr *QScriptEngine) DestroyQScriptEngine() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngine_DestroyQScriptEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngine) DestroyQScriptEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngine_DestroyQScriptEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngine) __children_newList() unsafe.Pointer {
	return C.QScriptEngine___children_newList(ptr.Pointer())
}

func (ptr *QScriptEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScriptEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScriptEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QScriptEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QScriptEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngine) __findChildren_newList() unsafe.Pointer {
	return C.QScriptEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QScriptEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QScriptEngine___findChildren_newList3(ptr.Pointer())
}

//export callbackQScriptEngine_ChildEvent
func callbackQScriptEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptEngine_ConnectNotify
func callbackQScriptEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngine_CustomEvent
func callbackQScriptEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptEngine_DeleteLater
func callbackQScriptEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngine_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQScriptEngine_Destroyed
func callbackQScriptEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptEngine_DisconnectNotify
func callbackQScriptEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngine_Event
func callbackQScriptEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQScriptEngine_EventFilter
func callbackQScriptEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQScriptEngine_MetaObject
func callbackQScriptEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptEngine_ObjectNameChanged
func callbackQScriptEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScript_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptEngine_TimerEvent
func callbackQScriptEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QScriptEngineAgent struct {
	ptr unsafe.Pointer
}

type QScriptEngineAgent_ITF interface {
	QScriptEngineAgent_PTR() *QScriptEngineAgent
}

func (ptr *QScriptEngineAgent) QScriptEngineAgent_PTR() *QScriptEngineAgent {
	return ptr
}

func (ptr *QScriptEngineAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptEngineAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptEngineAgent(ptr QScriptEngineAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptEngineAgent_PTR().Pointer()
	}
	return nil
}

func NewQScriptEngineAgentFromPointer(ptr unsafe.Pointer) (n *QScriptEngineAgent) {
	n = new(QScriptEngineAgent)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptEngineAgent__Extension
//QScriptEngineAgent::Extension
type QScriptEngineAgent__Extension int64

const (
	QScriptEngineAgent__DebuggerInvocationRequest QScriptEngineAgent__Extension = QScriptEngineAgent__Extension(0)
)

func NewQScriptEngineAgent(engine QScriptEngine_ITF) *QScriptEngineAgent {
	return NewQScriptEngineAgentFromPointer(C.QScriptEngineAgent_NewQScriptEngineAgent(PointerFromQScriptEngine(engine)))
}

//export callbackQScriptEngineAgent_ContextPop
func callbackQScriptEngineAgent_ContextPop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextPop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPopDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextPop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "contextPop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextPop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectContextPop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contextPop")
	}
}

func (ptr *QScriptEngineAgent) ContextPop() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPop(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) ContextPopDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPopDefault(ptr.Pointer())
	}
}

//export callbackQScriptEngineAgent_ContextPush
func callbackQScriptEngineAgent_ContextPush(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextPush"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPushDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPush(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextPush"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "contextPush", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextPush", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectContextPush() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contextPush")
	}
}

func (ptr *QScriptEngineAgent) ContextPush() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPush(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) ContextPushDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ContextPushDefault(ptr.Pointer())
	}
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptEngineFromPointer(C.QScriptEngineAgent_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngineAgent_ExceptionCatch
func callbackQScriptEngineAgent_ExceptionCatch(ptr unsafe.Pointer, scriptId C.longlong, exception unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exceptionCatch"); signal != nil {
		(*(*func(int64, *QScriptValue))(signal))(int64(scriptId), NewQScriptValueFromPointer(exception))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionCatchDefault(int64(scriptId), NewQScriptValueFromPointer(exception))
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionCatch(f func(scriptId int64, exception *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exceptionCatch"); signal != nil {
			f := func(scriptId int64, exception *QScriptValue) {
				(*(*func(int64, *QScriptValue))(signal))(scriptId, exception)
				f(scriptId, exception)
			}
			qt.ConnectSignal(ptr.Pointer(), "exceptionCatch", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exceptionCatch", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectExceptionCatch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exceptionCatch")
	}
}

func (ptr *QScriptEngineAgent) ExceptionCatch(scriptId int64, exception QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionCatch(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception))
	}
}

func (ptr *QScriptEngineAgent) ExceptionCatchDefault(scriptId int64, exception QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionCatchDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception))
	}
}

//export callbackQScriptEngineAgent_ExceptionThrow
func callbackQScriptEngineAgent_ExceptionThrow(ptr unsafe.Pointer, scriptId C.longlong, exception unsafe.Pointer, hasHandler C.char) {
	if signal := qt.GetSignal(ptr, "exceptionThrow"); signal != nil {
		(*(*func(int64, *QScriptValue, bool))(signal))(int64(scriptId), NewQScriptValueFromPointer(exception), int8(hasHandler) != 0)
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionThrowDefault(int64(scriptId), NewQScriptValueFromPointer(exception), int8(hasHandler) != 0)
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionThrow(f func(scriptId int64, exception *QScriptValue, hasHandler bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exceptionThrow"); signal != nil {
			f := func(scriptId int64, exception *QScriptValue, hasHandler bool) {
				(*(*func(int64, *QScriptValue, bool))(signal))(scriptId, exception, hasHandler)
				f(scriptId, exception, hasHandler)
			}
			qt.ConnectSignal(ptr.Pointer(), "exceptionThrow", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exceptionThrow", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectExceptionThrow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exceptionThrow")
	}
}

func (ptr *QScriptEngineAgent) ExceptionThrow(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionThrow(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception), C.char(int8(qt.GoBoolToInt(hasHandler))))
	}
}

func (ptr *QScriptEngineAgent) ExceptionThrowDefault(scriptId int64, exception QScriptValue_ITF, hasHandler bool) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ExceptionThrowDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(exception), C.char(int8(qt.GoBoolToInt(hasHandler))))
	}
}

//export callbackQScriptEngineAgent_Extension
func callbackQScriptEngineAgent_Extension(ptr unsafe.Pointer, extension C.longlong, argument unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQVariant((*(*func(QScriptEngineAgent__Extension, *core.QVariant) *core.QVariant)(signal))(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptEngineAgentFromPointer(ptr).ExtensionDefault(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptEngineAgent) ConnectExtension(f func(extension QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			f := func(extension QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant {
				(*(*func(QScriptEngineAgent__Extension, *core.QVariant) *core.QVariant)(signal))(extension, argument)
				return f(extension, argument)
			}
			qt.ConnectSignal(ptr.Pointer(), "extension", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "extension")
	}
}

func (ptr *QScriptEngineAgent) Extension(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineAgent) ExtensionDefault(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QScriptEngineAgent_ExtensionDefault(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngineAgent_FunctionEntry
func callbackQScriptEngineAgent_FunctionEntry(ptr unsafe.Pointer, scriptId C.longlong) {
	if signal := qt.GetSignal(ptr, "functionEntry"); signal != nil {
		(*(*func(int64))(signal))(int64(scriptId))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionEntryDefault(int64(scriptId))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionEntry(f func(scriptId int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "functionEntry"); signal != nil {
			f := func(scriptId int64) {
				(*(*func(int64))(signal))(scriptId)
				f(scriptId)
			}
			qt.ConnectSignal(ptr.Pointer(), "functionEntry", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "functionEntry", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectFunctionEntry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "functionEntry")
	}
}

func (ptr *QScriptEngineAgent) FunctionEntry(scriptId int64) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionEntry(ptr.Pointer(), C.longlong(scriptId))
	}
}

func (ptr *QScriptEngineAgent) FunctionEntryDefault(scriptId int64) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionEntryDefault(ptr.Pointer(), C.longlong(scriptId))
	}
}

//export callbackQScriptEngineAgent_FunctionExit
func callbackQScriptEngineAgent_FunctionExit(ptr unsafe.Pointer, scriptId C.longlong, returnValue unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "functionExit"); signal != nil {
		(*(*func(int64, *QScriptValue))(signal))(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionExitDefault(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionExit(f func(scriptId int64, returnValue *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "functionExit"); signal != nil {
			f := func(scriptId int64, returnValue *QScriptValue) {
				(*(*func(int64, *QScriptValue))(signal))(scriptId, returnValue)
				f(scriptId, returnValue)
			}
			qt.ConnectSignal(ptr.Pointer(), "functionExit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "functionExit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectFunctionExit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "functionExit")
	}
}

func (ptr *QScriptEngineAgent) FunctionExit(scriptId int64, returnValue QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionExit(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(returnValue))
	}
}

func (ptr *QScriptEngineAgent) FunctionExitDefault(scriptId int64, returnValue QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_FunctionExitDefault(ptr.Pointer(), C.longlong(scriptId), PointerFromQScriptValue(returnValue))
	}
}

//export callbackQScriptEngineAgent_PositionChange
func callbackQScriptEngineAgent_PositionChange(ptr unsafe.Pointer, scriptId C.longlong, lineNumber C.int, columnNumber C.int) {
	if signal := qt.GetSignal(ptr, "positionChange"); signal != nil {
		(*(*func(int64, int, int))(signal))(int64(scriptId), int(int32(lineNumber)), int(int32(columnNumber)))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).PositionChangeDefault(int64(scriptId), int(int32(lineNumber)), int(int32(columnNumber)))
	}
}

func (ptr *QScriptEngineAgent) ConnectPositionChange(f func(scriptId int64, lineNumber int, columnNumber int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionChange"); signal != nil {
			f := func(scriptId int64, lineNumber int, columnNumber int) {
				(*(*func(int64, int, int))(signal))(scriptId, lineNumber, columnNumber)
				f(scriptId, lineNumber, columnNumber)
			}
			qt.ConnectSignal(ptr.Pointer(), "positionChange", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionChange", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectPositionChange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "positionChange")
	}
}

func (ptr *QScriptEngineAgent) PositionChange(scriptId int64, lineNumber int, columnNumber int) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_PositionChange(ptr.Pointer(), C.longlong(scriptId), C.int(int32(lineNumber)), C.int(int32(columnNumber)))
	}
}

func (ptr *QScriptEngineAgent) PositionChangeDefault(scriptId int64, lineNumber int, columnNumber int) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_PositionChangeDefault(ptr.Pointer(), C.longlong(scriptId), C.int(int32(lineNumber)), C.int(int32(columnNumber)))
	}
}

//export callbackQScriptEngineAgent_ScriptLoad
func callbackQScriptEngineAgent_ScriptLoad(ptr unsafe.Pointer, id C.longlong, program C.struct_QtScript_PackedString, fileName C.struct_QtScript_PackedString, baseLineNumber C.int) {
	if signal := qt.GetSignal(ptr, "scriptLoad"); signal != nil {
		(*(*func(int64, string, string, int))(signal))(int64(id), cGoUnpackString(program), cGoUnpackString(fileName), int(int32(baseLineNumber)))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptLoadDefault(int64(id), cGoUnpackString(program), cGoUnpackString(fileName), int(int32(baseLineNumber)))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptLoad(f func(id int64, program string, fileName string, baseLineNumber int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scriptLoad"); signal != nil {
			f := func(id int64, program string, fileName string, baseLineNumber int) {
				(*(*func(int64, string, string, int))(signal))(id, program, fileName, baseLineNumber)
				f(id, program, fileName, baseLineNumber)
			}
			qt.ConnectSignal(ptr.Pointer(), "scriptLoad", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scriptLoad", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectScriptLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scriptLoad")
	}
}

func (ptr *QScriptEngineAgent) ScriptLoad(id int64, program string, fileName string, baseLineNumber int) {
	if ptr.Pointer() != nil {
		var programC *C.char
		if program != "" {
			programC = C.CString(program)
			defer C.free(unsafe.Pointer(programC))
		}
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QScriptEngineAgent_ScriptLoad(ptr.Pointer(), C.longlong(id), C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(baseLineNumber)))
	}
}

func (ptr *QScriptEngineAgent) ScriptLoadDefault(id int64, program string, fileName string, baseLineNumber int) {
	if ptr.Pointer() != nil {
		var programC *C.char
		if program != "" {
			programC = C.CString(program)
			defer C.free(unsafe.Pointer(programC))
		}
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QScriptEngineAgent_ScriptLoadDefault(ptr.Pointer(), C.longlong(id), C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(baseLineNumber)))
	}
}

//export callbackQScriptEngineAgent_ScriptUnload
func callbackQScriptEngineAgent_ScriptUnload(ptr unsafe.Pointer, id C.longlong) {
	if signal := qt.GetSignal(ptr, "scriptUnload"); signal != nil {
		(*(*func(int64))(signal))(int64(id))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptUnloadDefault(int64(id))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptUnload(f func(id int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scriptUnload"); signal != nil {
			f := func(id int64) {
				(*(*func(int64))(signal))(id)
				f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "scriptUnload", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scriptUnload", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectScriptUnload() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scriptUnload")
	}
}

func (ptr *QScriptEngineAgent) ScriptUnload(id int64) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptUnload(ptr.Pointer(), C.longlong(id))
	}
}

func (ptr *QScriptEngineAgent) ScriptUnloadDefault(id int64) {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_ScriptUnloadDefault(ptr.Pointer(), C.longlong(id))
	}
}

//export callbackQScriptEngineAgent_SupportsExtension
func callbackQScriptEngineAgent_SupportsExtension(ptr unsafe.Pointer, extension C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "supportsExtension"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(QScriptEngineAgent__Extension) bool)(signal))(QScriptEngineAgent__Extension(extension)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineAgentFromPointer(ptr).SupportsExtensionDefault(QScriptEngineAgent__Extension(extension)))))
}

func (ptr *QScriptEngineAgent) ConnectSupportsExtension(f func(extension QScriptEngineAgent__Extension) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportsExtension"); signal != nil {
			f := func(extension QScriptEngineAgent__Extension) bool {
				(*(*func(QScriptEngineAgent__Extension) bool)(signal))(extension)
				return f(extension)
			}
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectSupportsExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "supportsExtension")
	}
}

func (ptr *QScriptEngineAgent) SupportsExtension(extension QScriptEngineAgent__Extension) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineAgent_SupportsExtension(ptr.Pointer(), C.longlong(extension))) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) SupportsExtensionDefault(extension QScriptEngineAgent__Extension) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptEngineAgent_SupportsExtensionDefault(ptr.Pointer(), C.longlong(extension))) != 0
	}
	return false
}

//export callbackQScriptEngineAgent_DestroyQScriptEngineAgent
func callbackQScriptEngineAgent_DestroyQScriptEngineAgent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngineAgent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).DestroyQScriptEngineAgentDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectDestroyQScriptEngineAgent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngineAgent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineAgent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineAgent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptEngineAgent) DisconnectDestroyQScriptEngineAgent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScriptEngineAgent")
	}
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgentDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptEngineAgent_DestroyQScriptEngineAgentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptExtensionInterface struct {
	ptr unsafe.Pointer
}

type QScriptExtensionInterface_ITF interface {
	QScriptExtensionInterface_PTR() *QScriptExtensionInterface
}

func (ptr *QScriptExtensionInterface) QScriptExtensionInterface_PTR() *QScriptExtensionInterface {
	return ptr
}

func (ptr *QScriptExtensionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptExtensionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptExtensionInterface(ptr QScriptExtensionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionInterface_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionInterfaceFromPointer(ptr unsafe.Pointer) (n *QScriptExtensionInterface) {
	n = new(QScriptExtensionInterface)
	n.SetPointer(ptr)
	return
}
func (ptr *QScriptExtensionInterface) DestroyQScriptExtensionInterface() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptExtensionPlugin struct {
	core.QObject
	QScriptExtensionInterface
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionInterface_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QScriptExtensionInterface_PTR().SetPointer(p)
	}
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QScriptExtensionPlugin) {
	n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	return
}
func NewQScriptExtensionPlugin(parent core.QObject_ITF) *QScriptExtensionPlugin {
	tmpValue := NewQScriptExtensionPluginFromPointer(C.QScriptExtensionPlugin_NewQScriptExtensionPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScriptExtensionPlugin_Initialize
func callbackQScriptExtensionPlugin_Initialize(ptr unsafe.Pointer, key C.struct_QtScript_PackedString, engine unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialize"); signal != nil {
		(*(*func(string, *QScriptEngine))(signal))(cGoUnpackString(key), NewQScriptEngineFromPointer(engine))
	}

}

func (ptr *QScriptExtensionPlugin) ConnectInitialize(f func(key string, engine *QScriptEngine)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initialize"); signal != nil {
			f := func(key string, engine *QScriptEngine) {
				(*(*func(string, *QScriptEngine))(signal))(key, engine)
				f(key, engine)
			}
			qt.ConnectSignal(ptr.Pointer(), "initialize", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialize", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectInitialize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "initialize")
	}
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.struct_QtScript_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQScriptEngine(engine))
	}
}

//export callbackQScriptExtensionPlugin_Keys
func callbackQScriptExtensionPlugin_Keys(ptr unsafe.Pointer) C.struct_QtScript_PackedString {
	if signal := qt.GetSignal(ptr, "keys"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_QtScript_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtScript_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *QScriptExtensionPlugin) ConnectKeys(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keys"); signal != nil {
			f := func() []string {
				(*(*func() []string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "keys", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keys", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectKeys() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "keys")
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.struct_QtScript_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQScriptEngine(engine)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptExtensionPlugin_DestroyQScriptExtensionPlugin
func callbackQScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptExtensionPlugin"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DestroyQScriptExtensionPluginDefault()
	}
}

func (ptr *QScriptExtensionPlugin) ConnectDestroyQScriptExtensionPlugin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptExtensionPlugin"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QScriptExtensionPlugin", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptExtensionPlugin", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectDestroyQScriptExtensionPlugin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScriptExtensionPlugin")
	}
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPluginDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPluginDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScriptExtensionPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptExtensionPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptExtensionPlugin) __children_newList() unsafe.Pointer {
	return C.QScriptExtensionPlugin___children_newList(ptr.Pointer())
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScriptExtensionPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QScriptExtensionPlugin___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptExtensionPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptExtensionPlugin) __findChildren_newList() unsafe.Pointer {
	return C.QScriptExtensionPlugin___findChildren_newList(ptr.Pointer())
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptExtensionPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptExtensionPlugin) __findChildren_newList3() unsafe.Pointer {
	return C.QScriptExtensionPlugin___findChildren_newList3(ptr.Pointer())
}

//export callbackQScriptExtensionPlugin_ChildEvent
func callbackQScriptExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_ConnectNotify
func callbackQScriptExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptExtensionPlugin_CustomEvent
func callbackQScriptExtensionPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_DeleteLater
func callbackQScriptExtensionPlugin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptExtensionPlugin) DeleteLater() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptExtensionPlugin_DeleteLater(ptr.Pointer())
	}
}

func (ptr *QScriptExtensionPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQScriptExtensionPlugin_Destroyed
func callbackQScriptExtensionPlugin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptExtensionPlugin_DisconnectNotify
func callbackQScriptExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptExtensionPlugin_Event
func callbackQScriptExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptExtensionPlugin) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptExtensionPlugin_Event(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QScriptExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_EventFilter
func callbackQScriptExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptExtensionPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptExtensionPlugin_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QScriptExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_MetaObject
func callbackQScriptExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQScriptExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptExtensionPlugin) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptExtensionPlugin_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQScriptExtensionPlugin_ObjectNameChanged
func callbackQScriptExtensionPlugin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScript_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptExtensionPlugin_TimerEvent
func callbackQScriptExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScriptExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QScriptProgram struct {
	ptr unsafe.Pointer
}

type QScriptProgram_ITF interface {
	QScriptProgram_PTR() *QScriptProgram
}

func (ptr *QScriptProgram) QScriptProgram_PTR() *QScriptProgram {
	return ptr
}

func (ptr *QScriptProgram) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptProgram) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptProgram(ptr QScriptProgram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptProgram_PTR().Pointer()
	}
	return nil
}

func NewQScriptProgramFromPointer(ptr unsafe.Pointer) (n *QScriptProgram) {
	n = new(QScriptProgram)
	n.SetPointer(ptr)
	return
}
func NewQScriptProgram() *QScriptProgram {
	tmpValue := NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram())
	qt.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
	return tmpValue
}

func NewQScriptProgram2(sourceCode string, fileName string, firstLineNumber int) *QScriptProgram {
	var sourceCodeC *C.char
	if sourceCode != "" {
		sourceCodeC = C.CString(sourceCode)
		defer C.free(unsafe.Pointer(sourceCodeC))
	}
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram2(C.struct_QtScript_PackedString{data: sourceCodeC, len: C.longlong(len(sourceCode))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(firstLineNumber))))
	qt.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
	return tmpValue
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {
	tmpValue := NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram3(PointerFromQScriptProgram(other)))
	qt.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
	return tmpValue
}

func (ptr *QScriptProgram) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptProgram_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) FirstLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptProgram_FirstLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptProgram) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptProgram_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptProgram) SourceCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptProgram_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptProgram_DestroyQScriptProgram(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptString struct {
	ptr unsafe.Pointer
}

type QScriptString_ITF interface {
	QScriptString_PTR() *QScriptString
}

func (ptr *QScriptString) QScriptString_PTR() *QScriptString {
	return ptr
}

func (ptr *QScriptString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptString(ptr QScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptString_PTR().Pointer()
	}
	return nil
}

func NewQScriptStringFromPointer(ptr unsafe.Pointer) (n *QScriptString) {
	n = new(QScriptString)
	n.SetPointer(ptr)
	return
}
func NewQScriptString() *QScriptString {
	tmpValue := NewQScriptStringFromPointer(C.QScriptString_NewQScriptString())
	qt.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
	return tmpValue
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {
	tmpValue := NewQScriptStringFromPointer(C.QScriptString_NewQScriptString2(PointerFromQScriptString(other)))
	qt.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
	return tmpValue
}

func (ptr *QScriptString) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptString_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptString) ToArrayIndex(ok *bool) uint {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint(uint32(C.QScriptString_ToArrayIndex(ptr.Pointer(), &okC)))
	}
	return 0
}

func (ptr *QScriptString) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptString_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptString) DestroyQScriptString() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptString_DestroyQScriptString(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptSyntaxCheckResult struct {
	ptr unsafe.Pointer
}

type QScriptSyntaxCheckResult_ITF interface {
	QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult
}

func (ptr *QScriptSyntaxCheckResult) QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult {
	return ptr
}

func (ptr *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptSyntaxCheckResult) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptSyntaxCheckResult(ptr QScriptSyntaxCheckResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptSyntaxCheckResult_PTR().Pointer()
	}
	return nil
}

func NewQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) (n *QScriptSyntaxCheckResult) {
	n = new(QScriptSyntaxCheckResult)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptSyntaxCheckResult__State
//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int64

const (
	QScriptSyntaxCheckResult__Error        QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        QScriptSyntaxCheckResult__State = QScriptSyntaxCheckResult__State(2)
)

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResult_ITF) *QScriptSyntaxCheckResult {
	tmpValue := NewQScriptSyntaxCheckResultFromPointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(PointerFromQScriptSyntaxCheckResult(other)))
	qt.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptSyntaxCheckResult) ErrorColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptSyntaxCheckResult_ErrorColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptSyntaxCheckResult_ErrorLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptSyntaxCheckResult_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptSyntaxCheckResult) State() QScriptSyntaxCheckResult__State {
	if ptr.Pointer() != nil {
		return QScriptSyntaxCheckResult__State(C.QScriptSyntaxCheckResult_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptValue struct {
	ptr unsafe.Pointer
}

type QScriptValue_ITF interface {
	QScriptValue_PTR() *QScriptValue
}

func (ptr *QScriptValue) QScriptValue_PTR() *QScriptValue {
	return ptr
}

func (ptr *QScriptValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptValue(ptr QScriptValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValue_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueFromPointer(ptr unsafe.Pointer) (n *QScriptValue) {
	n = new(QScriptValue)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QScriptValue__ResolveFlag
//QScriptValue::ResolveFlag
type QScriptValue__ResolveFlag int64

const (
	QScriptValue__ResolveLocal     QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x00)
	QScriptValue__ResolvePrototype QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x01)
	QScriptValue__ResolveScope     QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(0x02)
	QScriptValue__ResolveFull      QScriptValue__ResolveFlag = QScriptValue__ResolveFlag(QScriptValue__ResolvePrototype | QScriptValue__ResolveScope)
)

//go:generate stringer -type=QScriptValue__SpecialValue
//QScriptValue::SpecialValue
type QScriptValue__SpecialValue int64

const (
	QScriptValue__NullValue      QScriptValue__SpecialValue = QScriptValue__SpecialValue(0)
	QScriptValue__UndefinedValue QScriptValue__SpecialValue = QScriptValue__SpecialValue(1)
)

//go:generate stringer -type=QScriptValue__PropertyFlag
//QScriptValue::PropertyFlag
type QScriptValue__PropertyFlag int64

const (
	QScriptValue__ReadOnly          QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000001)
	QScriptValue__Undeletable       QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000002)
	QScriptValue__SkipInEnumeration QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000004)
	QScriptValue__PropertyGetter    QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000008)
	QScriptValue__PropertySetter    QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000010)
	QScriptValue__QObjectMember     QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000020)
	QScriptValue__KeepExistingFlags QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0x00000800)
	QScriptValue__UserRange         QScriptValue__PropertyFlag = QScriptValue__PropertyFlag(0xff000000)
)

func NewQScriptValue() *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue())
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue2(PointerFromQScriptValue(other)))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue10(C.longlong(value)))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue11(value bool) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue11(C.char(int8(qt.GoBoolToInt(value)))))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue12(value int) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue12(C.int(int32(value))))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue13(value uint) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue13(C.uint(uint32(value))))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue15(value string) *QScriptValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue15(C.struct_QtScript_PackedString{data: valueC, len: C.longlong(len(value))}))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue16(core.PointerFromQLatin1String(value)))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue17(value string) *QScriptValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue17(valueC))
	qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func (ptr *QScriptValue) Call2(thisObject QScriptValue_ITF, arguments QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Call2(ptr.Pointer(), PointerFromQScriptValue(thisObject), PointerFromQScriptValue(arguments)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Construct2(arguments QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Construct2(ptr.Pointer(), PointerFromQScriptValue(arguments)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Data() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Data(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptEngineFromPointer(C.QScriptValue_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Equals(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_Equals(ptr.Pointer(), PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) InstanceOf(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_InstanceOf(ptr.Pointer(), PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsArray(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsBool(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsDate(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsFunction() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsFunction(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsNumber(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQMetaObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsQMetaObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsQObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsRegExp(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsString() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsString(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsUndefined(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_IsVariant(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) LessThan(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_LessThan(ptr.Pointer(), PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) Property(name string, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Property(ptr.Pointer(), C.struct_QtScript_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Property2(arrayIndex uint, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex)), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Property3(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Property3(ptr.Pointer(), PointerFromQScriptString(name), C.longlong(mode)))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) PropertyFlags(name string, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags(ptr.Pointer(), C.struct_QtScript_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(mode)))
	}
	return 0
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags2(ptr.Pointer(), PointerFromQScriptString(name), C.longlong(mode)))
	}
	return 0
}

func (ptr *QScriptValue) Prototype() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptValue_Prototype(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ScriptClass() *QScriptClass {
	if ptr.Pointer() != nil {
		return NewQScriptClassFromPointer(C.QScriptValue_ScriptClass(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) SetData(data QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetData(ptr.Pointer(), PointerFromQScriptValue(data))
	}
}

func (ptr *QScriptValue) SetProperty(name string, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QScriptValue_SetProperty(ptr.Pointer(), C.struct_QtScript_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQScriptValue(value), C.longlong(flags))
	}
}

func (ptr *QScriptValue) SetProperty2(arrayIndex uint, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty2(ptr.Pointer(), C.uint(uint32(arrayIndex)), PointerFromQScriptValue(value), C.longlong(flags))
	}
}

func (ptr *QScriptValue) SetProperty3(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty3(ptr.Pointer(), PointerFromQScriptString(name), PointerFromQScriptValue(value), C.longlong(flags))
	}
}

func (ptr *QScriptValue) SetPrototype(prototype QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetPrototype(ptr.Pointer(), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptValue) SetScriptClass(scriptClass QScriptClass_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetScriptClass(ptr.Pointer(), PointerFromQScriptClass(scriptClass))
	}
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_StrictlyEquals(ptr.Pointer(), PointerFromQScriptValue(other))) != 0
	}
	return false
}

func (ptr *QScriptValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return int8(C.QScriptValue_ToBool(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QScriptValue) ToDateTime() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QScriptValue_ToDateTime(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ToInt32() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptValue_ToInt32(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptValue_ToQMetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScriptValue_ToQObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ToRegExp() *core.QRegExp {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRegExpFromPointer(C.QScriptValue_ToRegExp(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptValue) ToUInt16() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QScriptValue_ToUInt16(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptValue) ToUInt32() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QScriptValue_ToUInt32(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QScriptValue_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) DestroyQScriptValue() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScriptValue_DestroyQScriptValue(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScriptValueIterator struct {
	ptr unsafe.Pointer
}

type QScriptValueIterator_ITF interface {
	QScriptValueIterator_PTR() *QScriptValueIterator
}

func (ptr *QScriptValueIterator) QScriptValueIterator_PTR() *QScriptValueIterator {
	return ptr
}

func (ptr *QScriptValueIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptValueIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptValueIterator(ptr QScriptValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueIteratorFromPointer(ptr unsafe.Pointer) (n *QScriptValueIterator) {
	n = new(QScriptValueIterator)
	n.SetPointer(ptr)
	return
}

type QScriptable struct {
	ptr unsafe.Pointer
}

type QScriptable_ITF interface {
	QScriptable_PTR() *QScriptable
}

func (ptr *QScriptable) QScriptable_PTR() *QScriptable {
	return ptr
}

func (ptr *QScriptable) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScriptable) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScriptable(ptr QScriptable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptable_PTR().Pointer()
	}
	return nil
}

func NewQScriptableFromPointer(ptr unsafe.Pointer) (n *QScriptable) {
	n = new(QScriptable)
	n.SetPointer(ptr)
	return
}
func (ptr *QScriptable) DestroyQScriptable() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QScriptable) Argument(index int) *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptable_Argument(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptable) ArgumentCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptable_ArgumentCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptable) Context() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptable_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptEngineFromPointer(C.QScriptable_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptable) ThisObject() *QScriptValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQScriptValueFromPointer(C.QScriptable_ThisObject(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func init() {
	qt.ItfMap["script.QScriptClass_ITF"] = QScriptClass{}
	qt.FuncMap["script.NewQScriptClass"] = NewQScriptClass
	qt.EnumMap["script.QScriptClass__HandlesReadAccess"] = int64(QScriptClass__HandlesReadAccess)
	qt.EnumMap["script.QScriptClass__HandlesWriteAccess"] = int64(QScriptClass__HandlesWriteAccess)
	qt.EnumMap["script.QScriptClass__Callable"] = int64(QScriptClass__Callable)
	qt.EnumMap["script.QScriptClass__HasInstance"] = int64(QScriptClass__HasInstance)
	qt.ItfMap["script.QScriptContext_ITF"] = QScriptContext{}
	qt.EnumMap["script.QScriptContext__NormalState"] = int64(QScriptContext__NormalState)
	qt.EnumMap["script.QScriptContext__ExceptionState"] = int64(QScriptContext__ExceptionState)
	qt.EnumMap["script.QScriptContext__UnknownError"] = int64(QScriptContext__UnknownError)
	qt.EnumMap["script.QScriptContext__ReferenceError"] = int64(QScriptContext__ReferenceError)
	qt.EnumMap["script.QScriptContext__SyntaxError"] = int64(QScriptContext__SyntaxError)
	qt.EnumMap["script.QScriptContext__TypeError"] = int64(QScriptContext__TypeError)
	qt.EnumMap["script.QScriptContext__RangeError"] = int64(QScriptContext__RangeError)
	qt.EnumMap["script.QScriptContext__URIError"] = int64(QScriptContext__URIError)
	qt.ItfMap["script.QScriptContextInfo_ITF"] = QScriptContextInfo{}
	qt.FuncMap["script.NewQScriptContextInfo"] = NewQScriptContextInfo
	qt.FuncMap["script.NewQScriptContextInfo2"] = NewQScriptContextInfo2
	qt.FuncMap["script.NewQScriptContextInfo3"] = NewQScriptContextInfo3
	qt.EnumMap["script.QScriptContextInfo__ScriptFunction"] = int64(QScriptContextInfo__ScriptFunction)
	qt.EnumMap["script.QScriptContextInfo__QtFunction"] = int64(QScriptContextInfo__QtFunction)
	qt.EnumMap["script.QScriptContextInfo__QtPropertyFunction"] = int64(QScriptContextInfo__QtPropertyFunction)
	qt.EnumMap["script.QScriptContextInfo__NativeFunction"] = int64(QScriptContextInfo__NativeFunction)
	qt.ItfMap["script.QScriptEngine_ITF"] = QScriptEngine{}
	qt.FuncMap["script.NewQScriptEngine"] = NewQScriptEngine
	qt.FuncMap["script.NewQScriptEngine2"] = NewQScriptEngine2
	qt.FuncMap["script.QScriptEngine_CheckSyntax"] = QScriptEngine_CheckSyntax
	qt.EnumMap["script.QScriptEngine__QtOwnership"] = int64(QScriptEngine__QtOwnership)
	qt.EnumMap["script.QScriptEngine__ScriptOwnership"] = int64(QScriptEngine__ScriptOwnership)
	qt.EnumMap["script.QScriptEngine__AutoOwnership"] = int64(QScriptEngine__AutoOwnership)
	qt.EnumMap["script.QScriptEngine__ExcludeChildObjects"] = int64(QScriptEngine__ExcludeChildObjects)
	qt.EnumMap["script.QScriptEngine__ExcludeSuperClassMethods"] = int64(QScriptEngine__ExcludeSuperClassMethods)
	qt.EnumMap["script.QScriptEngine__ExcludeSuperClassProperties"] = int64(QScriptEngine__ExcludeSuperClassProperties)
	qt.EnumMap["script.QScriptEngine__ExcludeSuperClassContents"] = int64(QScriptEngine__ExcludeSuperClassContents)
	qt.EnumMap["script.QScriptEngine__SkipMethodsInEnumeration"] = int64(QScriptEngine__SkipMethodsInEnumeration)
	qt.EnumMap["script.QScriptEngine__ExcludeDeleteLater"] = int64(QScriptEngine__ExcludeDeleteLater)
	qt.EnumMap["script.QScriptEngine__ExcludeSlots"] = int64(QScriptEngine__ExcludeSlots)
	qt.EnumMap["script.QScriptEngine__AutoCreateDynamicProperties"] = int64(QScriptEngine__AutoCreateDynamicProperties)
	qt.EnumMap["script.QScriptEngine__PreferExistingWrapperObject"] = int64(QScriptEngine__PreferExistingWrapperObject)
	qt.ItfMap["script.QScriptEngineAgent_ITF"] = QScriptEngineAgent{}
	qt.FuncMap["script.NewQScriptEngineAgent"] = NewQScriptEngineAgent
	qt.EnumMap["script.QScriptEngineAgent__DebuggerInvocationRequest"] = int64(QScriptEngineAgent__DebuggerInvocationRequest)
	qt.ItfMap["script.QScriptExtensionInterface_ITF"] = QScriptExtensionInterface{}
	qt.ItfMap["script.QScriptExtensionPlugin_ITF"] = QScriptExtensionPlugin{}
	qt.FuncMap["script.NewQScriptExtensionPlugin"] = NewQScriptExtensionPlugin
	qt.ItfMap["script.QScriptProgram_ITF"] = QScriptProgram{}
	qt.FuncMap["script.NewQScriptProgram"] = NewQScriptProgram
	qt.FuncMap["script.NewQScriptProgram2"] = NewQScriptProgram2
	qt.FuncMap["script.NewQScriptProgram3"] = NewQScriptProgram3
	qt.ItfMap["script.QScriptString_ITF"] = QScriptString{}
	qt.FuncMap["script.NewQScriptString"] = NewQScriptString
	qt.FuncMap["script.NewQScriptString2"] = NewQScriptString2
	qt.ItfMap["script.QScriptSyntaxCheckResult_ITF"] = QScriptSyntaxCheckResult{}
	qt.FuncMap["script.NewQScriptSyntaxCheckResult"] = NewQScriptSyntaxCheckResult
	qt.EnumMap["script.QScriptSyntaxCheckResult__Error"] = int64(QScriptSyntaxCheckResult__Error)
	qt.EnumMap["script.QScriptSyntaxCheckResult__Intermediate"] = int64(QScriptSyntaxCheckResult__Intermediate)
	qt.EnumMap["script.QScriptSyntaxCheckResult__Valid"] = int64(QScriptSyntaxCheckResult__Valid)
	qt.ItfMap["script.QScriptValue_ITF"] = QScriptValue{}
	qt.FuncMap["script.NewQScriptValue"] = NewQScriptValue
	qt.FuncMap["script.NewQScriptValue2"] = NewQScriptValue2
	qt.FuncMap["script.NewQScriptValue10"] = NewQScriptValue10
	qt.FuncMap["script.NewQScriptValue11"] = NewQScriptValue11
	qt.FuncMap["script.NewQScriptValue12"] = NewQScriptValue12
	qt.FuncMap["script.NewQScriptValue13"] = NewQScriptValue13
	qt.FuncMap["script.NewQScriptValue15"] = NewQScriptValue15
	qt.FuncMap["script.NewQScriptValue16"] = NewQScriptValue16
	qt.FuncMap["script.NewQScriptValue17"] = NewQScriptValue17
	qt.EnumMap["script.QScriptValue__ResolveLocal"] = int64(QScriptValue__ResolveLocal)
	qt.EnumMap["script.QScriptValue__ResolvePrototype"] = int64(QScriptValue__ResolvePrototype)
	qt.EnumMap["script.QScriptValue__ResolveScope"] = int64(QScriptValue__ResolveScope)
	qt.EnumMap["script.QScriptValue__ResolveFull"] = int64(QScriptValue__ResolveFull)
	qt.EnumMap["script.QScriptValue__NullValue"] = int64(QScriptValue__NullValue)
	qt.EnumMap["script.QScriptValue__UndefinedValue"] = int64(QScriptValue__UndefinedValue)
	qt.EnumMap["script.QScriptValue__ReadOnly"] = int64(QScriptValue__ReadOnly)
	qt.EnumMap["script.QScriptValue__Undeletable"] = int64(QScriptValue__Undeletable)
	qt.EnumMap["script.QScriptValue__SkipInEnumeration"] = int64(QScriptValue__SkipInEnumeration)
	qt.EnumMap["script.QScriptValue__PropertyGetter"] = int64(QScriptValue__PropertyGetter)
	qt.EnumMap["script.QScriptValue__PropertySetter"] = int64(QScriptValue__PropertySetter)
	qt.EnumMap["script.QScriptValue__QObjectMember"] = int64(QScriptValue__QObjectMember)
	qt.EnumMap["script.QScriptValue__KeepExistingFlags"] = int64(QScriptValue__KeepExistingFlags)
	qt.EnumMap["script.QScriptValue__UserRange"] = int64(QScriptValue__UserRange)
	qt.ItfMap["script.QScriptable_ITF"] = QScriptable{}
}
