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
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtScript_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QSOperator struct {
	ptr unsafe.Pointer
}

type QSOperator_ITF interface {
	QSOperator_PTR() *QSOperator
}

func (ptr *QSOperator) QSOperator_PTR() *QSOperator {
	return ptr
}

func (ptr *QSOperator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSOperator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSOperator(ptr QSOperator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSOperator_PTR().Pointer()
	}
	return nil
}

func NewQSOperatorFromPointer(ptr unsafe.Pointer) *QSOperator {
	var n = new(QSOperator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSOperator) DestroyQSOperator() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QSOperator__Op
//QSOperator::Op
type QSOperator__Op int64

const (
	QSOperator__Add                QSOperator__Op = QSOperator__Op(0)
	QSOperator__And                QSOperator__Op = QSOperator__Op(1)
	QSOperator__InplaceAnd         QSOperator__Op = QSOperator__Op(2)
	QSOperator__Assign             QSOperator__Op = QSOperator__Op(3)
	QSOperator__BitAnd             QSOperator__Op = QSOperator__Op(4)
	QSOperator__BitOr              QSOperator__Op = QSOperator__Op(5)
	QSOperator__BitXor             QSOperator__Op = QSOperator__Op(6)
	QSOperator__InplaceSub         QSOperator__Op = QSOperator__Op(7)
	QSOperator__Div                QSOperator__Op = QSOperator__Op(8)
	QSOperator__InplaceDiv         QSOperator__Op = QSOperator__Op(9)
	QSOperator__Equal              QSOperator__Op = QSOperator__Op(10)
	QSOperator__Ge                 QSOperator__Op = QSOperator__Op(11)
	QSOperator__Gt                 QSOperator__Op = QSOperator__Op(12)
	QSOperator__In                 QSOperator__Op = QSOperator__Op(13)
	QSOperator__InplaceAdd         QSOperator__Op = QSOperator__Op(14)
	QSOperator__InstanceOf         QSOperator__Op = QSOperator__Op(15)
	QSOperator__Le                 QSOperator__Op = QSOperator__Op(16)
	QSOperator__LShift             QSOperator__Op = QSOperator__Op(17)
	QSOperator__InplaceLeftShift   QSOperator__Op = QSOperator__Op(18)
	QSOperator__Lt                 QSOperator__Op = QSOperator__Op(19)
	QSOperator__Mod                QSOperator__Op = QSOperator__Op(20)
	QSOperator__InplaceMod         QSOperator__Op = QSOperator__Op(21)
	QSOperator__Mul                QSOperator__Op = QSOperator__Op(22)
	QSOperator__InplaceMul         QSOperator__Op = QSOperator__Op(23)
	QSOperator__NotEqual           QSOperator__Op = QSOperator__Op(24)
	QSOperator__Or                 QSOperator__Op = QSOperator__Op(25)
	QSOperator__InplaceOr          QSOperator__Op = QSOperator__Op(26)
	QSOperator__RShift             QSOperator__Op = QSOperator__Op(27)
	QSOperator__InplaceRightShift  QSOperator__Op = QSOperator__Op(28)
	QSOperator__StrictEqual        QSOperator__Op = QSOperator__Op(29)
	QSOperator__StrictNotEqual     QSOperator__Op = QSOperator__Op(30)
	QSOperator__Sub                QSOperator__Op = QSOperator__Op(31)
	QSOperator__URShift            QSOperator__Op = QSOperator__Op(32)
	QSOperator__InplaceURightShift QSOperator__Op = QSOperator__Op(33)
	QSOperator__InplaceXor         QSOperator__Op = QSOperator__Op(34)
)

type QScript struct {
	ptr unsafe.Pointer
}

type QScript_ITF interface {
	QScript_PTR() *QScript
}

func (ptr *QScript) QScript_PTR() *QScript {
	return ptr
}

func (ptr *QScript) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScript) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScript(ptr QScript_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScript_PTR().Pointer()
	}
	return nil
}

func NewQScriptFromPointer(ptr unsafe.Pointer) *QScript {
	var n = new(QScript)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScript) DestroyQScript() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQScriptClassFromPointer(ptr unsafe.Pointer) *QScriptClass {
	var n = new(QScriptClass)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QScriptClass__Extension
//QScriptClass::Extension
type QScriptClass__Extension int64

const (
	QScriptClass__Callable    QScriptClass__Extension = QScriptClass__Extension(0)
	QScriptClass__HasInstance QScriptClass__Extension = QScriptClass__Extension(1)
)

//go:generate stringer -type=QScriptClass__QueryFlag
//QScriptClass::QueryFlag
type QScriptClass__QueryFlag int64

const (
	QScriptClass__HandlesReadAccess  QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x01)
	QScriptClass__HandlesWriteAccess QScriptClass__QueryFlag = QScriptClass__QueryFlag(0x02)
)

func NewQScriptClass(engine QScriptEngine_ITF) *QScriptClass {
	return NewQScriptClassFromPointer(C.QScriptClass_NewQScriptClass(PointerFromQScriptEngine(engine)))
}

//export callbackQScriptClass_NewIterator
func callbackQScriptClass_NewIterator(ptr unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "newIterator"); signal != nil {
		return PointerFromQScriptClassPropertyIterator(signal.(func(*QScriptValue) *QScriptClassPropertyIterator)(NewQScriptValueFromPointer(object)))
	}

	return PointerFromQScriptClassPropertyIterator(NewQScriptClassFromPointer(ptr).NewIteratorDefault(NewQScriptValueFromPointer(object)))
}

func (ptr *QScriptClass) ConnectNewIterator(f func(object *QScriptValue) *QScriptClassPropertyIterator) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newIterator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newIterator", func(object *QScriptValue) *QScriptClassPropertyIterator {
				signal.(func(*QScriptValue) *QScriptClassPropertyIterator)(object)
				return f(object)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newIterator", f)
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
		return PointerFromQScriptValue(signal.(func(*QScriptValue, *QScriptString, uint) *QScriptValue)(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
	}

	return PointerFromQScriptValue(NewQScriptClassFromPointer(ptr).PropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectProperty(f func(object *QScriptValue, name *QScriptString, id uint) *QScriptValue) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "property", func(object *QScriptValue, name *QScriptString, id uint) *QScriptValue {
				signal.(func(*QScriptValue, *QScriptString, uint) *QScriptValue)(object, name, id)
				return f(object, name, id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", f)
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
		var tmpValue = NewQScriptValueFromPointer(C.QScriptClass_Property(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) PropertyDefault(object QScriptValue_ITF, name QScriptString_ITF, id uint) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptClass_PropertyDefault(ptr.Pointer(), PointerFromQScriptValue(object), PointerFromQScriptString(name), C.uint(uint32(id))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_PropertyFlags
func callbackQScriptClass_PropertyFlags(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, id C.uint) C.longlong {
	if signal := qt.GetSignal(ptr, "propertyFlags"); signal != nil {
		return C.longlong(signal.(func(*QScriptValue, *QScriptString, uint) QScriptValue__PropertyFlag)(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
	}

	return C.longlong(NewQScriptClassFromPointer(ptr).PropertyFlagsDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectPropertyFlags(f func(object *QScriptValue, name *QScriptString, id uint) QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "propertyFlags"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "propertyFlags", func(object *QScriptValue, name *QScriptString, id uint) QScriptValue__PropertyFlag {
				signal.(func(*QScriptValue, *QScriptString, uint) QScriptValue__PropertyFlag)(object, name, id)
				return f(object, name, id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "propertyFlags", f)
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

//export callbackQScriptClass_Extension
func callbackQScriptClass_Extension(ptr unsafe.Pointer, extension C.longlong, argument unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QScriptClass__Extension, *core.QVariant) *core.QVariant)(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptClassFromPointer(ptr).ExtensionDefault(QScriptClass__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptClass) ConnectExtension(f func(extension QScriptClass__Extension, argument *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "extension", func(extension QScriptClass__Extension, argument *core.QVariant) *core.QVariant {
				signal.(func(QScriptClass__Extension, *core.QVariant) *core.QVariant)(extension, argument)
				return f(extension, argument)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", f)
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
		var tmpValue = core.NewQVariantFromPointer(C.QScriptClass_Extension(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) ExtensionDefault(extension QScriptClass__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScriptClass_ExtensionDefault(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_QueryProperty
func callbackQScriptClass_QueryProperty(ptr unsafe.Pointer, object unsafe.Pointer, name unsafe.Pointer, flags C.longlong, id C.uint) C.longlong {
	if signal := qt.GetSignal(ptr, "queryProperty"); signal != nil {
		return C.longlong(signal.(func(*QScriptValue, *QScriptString, QScriptClass__QueryFlag, uint) QScriptClass__QueryFlag)(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), QScriptClass__QueryFlag(flags), uint(uint32(id))))
	}

	return C.longlong(NewQScriptClassFromPointer(ptr).QueryPropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), QScriptClass__QueryFlag(flags), uint(uint32(id))))
}

func (ptr *QScriptClass) ConnectQueryProperty(f func(object *QScriptValue, name *QScriptString, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "queryProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "queryProperty", func(object *QScriptValue, name *QScriptString, flags QScriptClass__QueryFlag, id uint) QScriptClass__QueryFlag {
				signal.(func(*QScriptValue, *QScriptString, QScriptClass__QueryFlag, uint) QScriptClass__QueryFlag)(object, name, flags, id)
				return f(object, name, flags, id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "queryProperty", f)
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
		signal.(func(*QScriptValue, *QScriptString, uint, *QScriptValue))(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id)), NewQScriptValueFromPointer(value))
	} else {
		NewQScriptClassFromPointer(ptr).SetPropertyDefault(NewQScriptValueFromPointer(object), NewQScriptStringFromPointer(name), uint(uint32(id)), NewQScriptValueFromPointer(value))
	}
}

func (ptr *QScriptClass) ConnectSetProperty(f func(object *QScriptValue, name *QScriptString, id uint, value *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(object *QScriptValue, name *QScriptString, id uint, value *QScriptValue) {
				signal.(func(*QScriptValue, *QScriptString, uint, *QScriptValue))(object, name, id, value)
				f(object, name, id, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
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

//export callbackQScriptClass_DestroyQScriptClass
func callbackQScriptClass_DestroyQScriptClass(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptClass"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptClassFromPointer(ptr).DestroyQScriptClassDefault()
	}
}

func (ptr *QScriptClass) ConnectDestroyQScriptClass(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptClass"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptClass", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptClass", f)
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
		C.QScriptClass_DestroyQScriptClass(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptClass) DestroyQScriptClassDefault() {
	if ptr.Pointer() != nil {
		C.QScriptClass_DestroyQScriptClassDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptClass) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptEngineFromPointer(C.QScriptClass_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_Prototype
func callbackQScriptClass_Prototype(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "prototype"); signal != nil {
		return PointerFromQScriptValue(signal.(func() *QScriptValue)())
	}

	return PointerFromQScriptValue(NewQScriptClassFromPointer(ptr).PrototypeDefault())
}

func (ptr *QScriptClass) ConnectPrototype(f func() *QScriptValue) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "prototype"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "prototype", func() *QScriptValue {
				signal.(func() *QScriptValue)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "prototype", f)
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
		var tmpValue = NewQScriptValueFromPointer(C.QScriptClass_Prototype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptClass) PrototypeDefault() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptClass_PrototypeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptClass_Name
func callbackQScriptClass_Name(ptr unsafe.Pointer) C.struct_QtScript_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtScript_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQScriptClassFromPointer(ptr).NameDefault()
	return C.struct_QtScript_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QScriptClass) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "name", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", f)
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

//export callbackQScriptClass_SupportsExtension
func callbackQScriptClass_SupportsExtension(ptr unsafe.Pointer, extension C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "supportsExtension"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QScriptClass__Extension) bool)(QScriptClass__Extension(extension)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptClassFromPointer(ptr).SupportsExtensionDefault(QScriptClass__Extension(extension)))))
}

func (ptr *QScriptClass) ConnectSupportsExtension(f func(extension QScriptClass__Extension) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportsExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", func(extension QScriptClass__Extension) bool {
				signal.(func(QScriptClass__Extension) bool)(extension)
				return f(extension)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", f)
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
		return C.QScriptClass_SupportsExtension(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
}

func (ptr *QScriptClass) SupportsExtensionDefault(extension QScriptClass__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QScriptClass_SupportsExtensionDefault(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
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

func NewQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) *QScriptClassPropertyIterator {
	var n = new(QScriptClassPropertyIterator)
	n.SetPointer(ptr)
	return n
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

func NewQScriptContextFromPointer(ptr unsafe.Pointer) *QScriptContext {
	var n = new(QScriptContext)
	n.SetPointer(ptr)
	return n
}

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

//go:generate stringer -type=QScriptContext__ExecutionState
//QScriptContext::ExecutionState
type QScriptContext__ExecutionState int64

const (
	QScriptContext__NormalState    QScriptContext__ExecutionState = QScriptContext__ExecutionState(0)
	QScriptContext__ExceptionState QScriptContext__ExecutionState = QScriptContext__ExecutionState(1)
)

func (ptr *QScriptContext) ThrowError(error QScriptContext__Error, text string) *QScriptValue {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ThrowError(ptr.Pointer(), C.longlong(error), C.struct_QtScript_PackedString{data: textC, len: C.longlong(len(text))}))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
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
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ThrowError2(ptr.Pointer(), C.struct_QtScript_PackedString{data: textC, len: C.longlong(len(text))}))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ThrowValue(value QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ThrowValue(ptr.Pointer(), PointerFromQScriptValue(value)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
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

func (ptr *QScriptContext) DestroyQScriptContext() {
	if ptr.Pointer() != nil {
		C.QScriptContext_DestroyQScriptContext(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptContext) State() QScriptContext__ExecutionState {
	if ptr.Pointer() != nil {
		return QScriptContext__ExecutionState(C.QScriptContext_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContext) ParentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptContext_ParentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptContext) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptEngineFromPointer(C.QScriptContext_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ActivationObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ActivationObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) Argument(index int) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_Argument(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ArgumentsObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ArgumentsObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) Callee() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_Callee(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptContext) ThisObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptContext_ThisObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
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

func (ptr *QScriptContext) Backtrace() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScriptContext_Backtrace(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContext) IsCalledAsConstructor() bool {
	if ptr.Pointer() != nil {
		return C.QScriptContext_IsCalledAsConstructor(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptContext) ArgumentCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContext_ArgumentCount(ptr.Pointer())))
	}
	return 0
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

func NewQScriptContextInfoFromPointer(ptr unsafe.Pointer) *QScriptContextInfo {
	var n = new(QScriptContextInfo)
	n.SetPointer(ptr)
	return n
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

func NewQScriptContextInfo3() *QScriptContextInfo {
	var tmpValue = NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo3())
	runtime.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func NewQScriptContextInfo(context QScriptContext_ITF) *QScriptContextInfo {
	var tmpValue = NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo(PointerFromQScriptContext(context)))
	runtime.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func NewQScriptContextInfo2(other QScriptContextInfo_ITF) *QScriptContextInfo {
	var tmpValue = NewQScriptContextInfoFromPointer(C.QScriptContextInfo_NewQScriptContextInfo2(PointerFromQScriptContextInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QScriptContextInfo).DestroyQScriptContextInfo)
	return tmpValue
}

func (ptr *QScriptContextInfo) DestroyQScriptContextInfo() {
	if ptr.Pointer() != nil {
		C.QScriptContextInfo_DestroyQScriptContextInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptContextInfo) FunctionType() QScriptContextInfo__FunctionType {
	if ptr.Pointer() != nil {
		return QScriptContextInfo__FunctionType(C.QScriptContextInfo_FunctionType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptContextInfo) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptContextInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptContextInfo_FunctionName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptContextInfo) FunctionParameterNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScriptContextInfo_FunctionParameterNames(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptContextInfo) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptContextInfo_IsNull(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QScriptContextInfo) FunctionStartLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptContextInfo_FunctionStartLineNumber(ptr.Pointer())))
	}
	return 0
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

func NewQScriptEngineFromPointer(ptr unsafe.Pointer) *QScriptEngine {
	var n = new(QScriptEngine)
	n.SetPointer(ptr)
	return n
}

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

//go:generate stringer -type=QScriptEngine__ValueOwnership
//QScriptEngine::ValueOwnership
type QScriptEngine__ValueOwnership int64

const (
	QScriptEngine__QtOwnership     QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(0)
	QScriptEngine__ScriptOwnership QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(1)
	QScriptEngine__AutoOwnership   QScriptEngine__ValueOwnership = QScriptEngine__ValueOwnership(2)
)

func (ptr *QScriptEngine) PushContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_PushContext(ptr.Pointer()))
	}
	return nil
}

func NewQScriptEngine() *QScriptEngine {
	var tmpValue = NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQScriptEngine2(parent core.QObject_ITF) *QScriptEngine {
	var tmpValue = NewQScriptEngineFromPointer(C.QScriptEngine_NewQScriptEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScriptEngine) ToStringHandle(str string) *QScriptString {
	if ptr.Pointer() != nil {
		var strC *C.char
		if str != "" {
			strC = C.CString(str)
			defer C.free(unsafe.Pointer(strC))
		}
		var tmpValue = NewQScriptStringFromPointer(C.QScriptEngine_ToStringHandle(ptr.Pointer(), C.struct_QtScript_PackedString{data: strC, len: C.longlong(len(str))}))
		runtime.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
		return tmpValue
	}
	return nil
}

func QScriptEngine_CheckSyntax(program string) *QScriptSyntaxCheckResult {
	var programC *C.char
	if program != "" {
		programC = C.CString(program)
		defer C.free(unsafe.Pointer(programC))
	}
	var tmpValue = NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}))
	runtime.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptEngine) CheckSyntax(program string) *QScriptSyntaxCheckResult {
	var programC *C.char
	if program != "" {
		programC = C.CString(program)
		defer C.free(unsafe.Pointer(programC))
	}
	var tmpValue = NewQScriptSyntaxCheckResultFromPointer(C.QScriptEngine_QScriptEngine_CheckSyntax(C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}))
	runtime.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptEngine) Evaluate2(program QScriptProgram_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_Evaluate2(ptr.Pointer(), PointerFromQScriptProgram(program)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
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
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_Evaluate(ptr.Pointer(), C.struct_QtScript_PackedString{data: programC, len: C.longlong(len(program))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(lineNumber))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) ImportExtension(extension string) *QScriptValue {
	if ptr.Pointer() != nil {
		var extensionC *C.char
		if extension != "" {
			extensionC = C.CString(extension)
			defer C.free(unsafe.Pointer(extensionC))
		}
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_ImportExtension(ptr.Pointer(), C.struct_QtScript_PackedString{data: extensionC, len: C.longlong(len(extension))}))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewArray(length uint) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewArray(ptr.Pointer(), C.uint(uint32(length))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewDate2(value core.QDateTime_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewDate2(ptr.Pointer(), core.PointerFromQDateTime(value)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewObject2(scriptClass QScriptClass_ITF, data QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewObject2(ptr.Pointer(), PointerFromQScriptClass(scriptClass), PointerFromQScriptValue(data)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQMetaObject(metaObject core.QMetaObject_ITF, ctor QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewQMetaObject(ptr.Pointer(), core.PointerFromQMetaObject(metaObject), PointerFromQScriptValue(ctor)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject(object core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object), C.longlong(ownership), C.longlong(options)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewQObject2(scriptObject QScriptValue_ITF, qtObject core.QObject_ITF, ownership QScriptEngine__ValueOwnership, options QScriptEngine__QObjectWrapOption) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewQObject2(ptr.Pointer(), PointerFromQScriptValue(scriptObject), core.PointerFromQObject(qtObject), C.longlong(ownership), C.longlong(options)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewRegExp(regexp core.QRegExp_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp(ptr.Pointer(), core.PointerFromQRegExp(regexp)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
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
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewRegExp2(ptr.Pointer(), C.struct_QtScript_PackedString{data: patternC, len: C.longlong(len(pattern))}, C.struct_QtScript_PackedString{data: flagsC, len: C.longlong(len(flags))}))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant2(object QScriptValue_ITF, value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewVariant2(ptr.Pointer(), PointerFromQScriptValue(object), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NewVariant(value core.QVariant_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NewVariant(ptr.Pointer(), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) NullValue() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_NullValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) ToObject(value QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_ToObject(ptr.Pointer(), PointerFromQScriptValue(value)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) UndefinedValue() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_UndefinedValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) AbortEvaluation(result QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_AbortEvaluation(ptr.Pointer(), PointerFromQScriptValue(result))
	}
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

func (ptr *QScriptEngine) InstallTranslatorFunctions(object QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_InstallTranslatorFunctions(ptr.Pointer(), PointerFromQScriptValue(object))
	}
}

func (ptr *QScriptEngine) PopContext() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_PopContext(ptr.Pointer())
	}
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
		signal.(func(*QScriptValue))(NewQScriptValueFromPointer(exception))
	}

}

func (ptr *QScriptEngine) ConnectSignalHandlerException(f func(exception *QScriptValue)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "signalHandlerException") {
			C.QScriptEngine_ConnectSignalHandlerException(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "signalHandlerException"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "signalHandlerException", func(exception *QScriptValue) {
				signal.(func(*QScriptValue))(exception)
				f(exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "signalHandlerException", f)
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

//export callbackQScriptEngine_DestroyQScriptEngine
func callbackQScriptEngine_DestroyQScriptEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineFromPointer(ptr).DestroyQScriptEngineDefault()
	}
}

func (ptr *QScriptEngine) ConnectDestroyQScriptEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngine", f)
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
		C.QScriptEngine_DestroyQScriptEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptEngine) DestroyQScriptEngineDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DestroyQScriptEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptEngine) CurrentContext() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptEngine_CurrentContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) Agent() *QScriptEngineAgent {
	if ptr.Pointer() != nil {
		return NewQScriptEngineAgentFromPointer(C.QScriptEngine_Agent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptEngine) DefaultPrototype(metaTypeId int) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_DefaultPrototype(ptr.Pointer(), C.int(int32(metaTypeId))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) GlobalObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_GlobalObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) UncaughtException() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptEngine_UncaughtException(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) AvailableExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScriptEngine_AvailableExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) ImportedExtensions() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScriptEngine_ImportedExtensions(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) UncaughtExceptionBacktrace() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScriptEngine_UncaughtExceptionBacktrace(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptEngine) HasUncaughtException() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_HasUncaughtException(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) IsEvaluating() bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_IsEvaluating(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptEngine) ProcessEventsInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptEngine_ProcessEventsInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) UncaughtExceptionLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptEngine_UncaughtExceptionLineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScriptEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScriptEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QScriptEngine___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScriptEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptEngine) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScriptEngine___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScriptEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptEngine___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScriptEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptEngine___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScriptEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptEngine___children_newList(ptr.Pointer()))
}

//export callbackQScriptEngine_Event
func callbackQScriptEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScriptEngine_EventFilter
func callbackQScriptEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScriptEngine_ChildEvent
func callbackQScriptEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQScriptEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScriptEngine_Destroyed
func callbackQScriptEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptEngine_DisconnectNotify
func callbackQScriptEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptEngine_ObjectNameChanged
func callbackQScriptEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScript_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptEngine_TimerEvent
func callbackQScriptEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScriptEngine_MetaObject
func callbackQScriptEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQScriptEngineAgentFromPointer(ptr unsafe.Pointer) *QScriptEngineAgent {
	var n = new(QScriptEngineAgent)
	n.SetPointer(ptr)
	return n
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

//export callbackQScriptEngineAgent_Extension
func callbackQScriptEngineAgent_Extension(ptr unsafe.Pointer, extension C.longlong, argument unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QScriptEngineAgent__Extension, *core.QVariant) *core.QVariant)(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
	}

	return core.PointerFromQVariant(NewQScriptEngineAgentFromPointer(ptr).ExtensionDefault(QScriptEngineAgent__Extension(extension), core.NewQVariantFromPointer(argument)))
}

func (ptr *QScriptEngineAgent) ConnectExtension(f func(extension QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "extension", func(extension QScriptEngineAgent__Extension, argument *core.QVariant) *core.QVariant {
				signal.(func(QScriptEngineAgent__Extension, *core.QVariant) *core.QVariant)(extension, argument)
				return f(extension, argument)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", f)
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
		var tmpValue = core.NewQVariantFromPointer(C.QScriptEngineAgent_Extension(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptEngineAgent) ExtensionDefault(extension QScriptEngineAgent__Extension, argument core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScriptEngineAgent_ExtensionDefault(ptr.Pointer(), C.longlong(extension), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngineAgent_ContextPop
func callbackQScriptEngineAgent_ContextPop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextPop"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPopDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextPop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contextPop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextPop", f)
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
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ContextPushDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectContextPush(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contextPush"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contextPush", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contextPush", f)
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

//export callbackQScriptEngineAgent_ExceptionCatch
func callbackQScriptEngineAgent_ExceptionCatch(ptr unsafe.Pointer, scriptId C.longlong, exception unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exceptionCatch"); signal != nil {
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(exception))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionCatchDefault(int64(scriptId), NewQScriptValueFromPointer(exception))
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionCatch(f func(scriptId int64, exception *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exceptionCatch"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exceptionCatch", func(scriptId int64, exception *QScriptValue) {
				signal.(func(int64, *QScriptValue))(scriptId, exception)
				f(scriptId, exception)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exceptionCatch", f)
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
		signal.(func(int64, *QScriptValue, bool))(int64(scriptId), NewQScriptValueFromPointer(exception), int8(hasHandler) != 0)
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ExceptionThrowDefault(int64(scriptId), NewQScriptValueFromPointer(exception), int8(hasHandler) != 0)
	}
}

func (ptr *QScriptEngineAgent) ConnectExceptionThrow(f func(scriptId int64, exception *QScriptValue, hasHandler bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exceptionThrow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exceptionThrow", func(scriptId int64, exception *QScriptValue, hasHandler bool) {
				signal.(func(int64, *QScriptValue, bool))(scriptId, exception, hasHandler)
				f(scriptId, exception, hasHandler)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exceptionThrow", f)
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

//export callbackQScriptEngineAgent_FunctionEntry
func callbackQScriptEngineAgent_FunctionEntry(ptr unsafe.Pointer, scriptId C.longlong) {
	if signal := qt.GetSignal(ptr, "functionEntry"); signal != nil {
		signal.(func(int64))(int64(scriptId))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionEntryDefault(int64(scriptId))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionEntry(f func(scriptId int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "functionEntry"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "functionEntry", func(scriptId int64) {
				signal.(func(int64))(scriptId)
				f(scriptId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "functionEntry", f)
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
		signal.(func(int64, *QScriptValue))(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).FunctionExitDefault(int64(scriptId), NewQScriptValueFromPointer(returnValue))
	}
}

func (ptr *QScriptEngineAgent) ConnectFunctionExit(f func(scriptId int64, returnValue *QScriptValue)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "functionExit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "functionExit", func(scriptId int64, returnValue *QScriptValue) {
				signal.(func(int64, *QScriptValue))(scriptId, returnValue)
				f(scriptId, returnValue)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "functionExit", f)
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
		signal.(func(int64, int, int))(int64(scriptId), int(int32(lineNumber)), int(int32(columnNumber)))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).PositionChangeDefault(int64(scriptId), int(int32(lineNumber)), int(int32(columnNumber)))
	}
}

func (ptr *QScriptEngineAgent) ConnectPositionChange(f func(scriptId int64, lineNumber int, columnNumber int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "positionChange"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "positionChange", func(scriptId int64, lineNumber int, columnNumber int) {
				signal.(func(int64, int, int))(scriptId, lineNumber, columnNumber)
				f(scriptId, lineNumber, columnNumber)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "positionChange", f)
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
		signal.(func(int64, string, string, int))(int64(id), cGoUnpackString(program), cGoUnpackString(fileName), int(int32(baseLineNumber)))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptLoadDefault(int64(id), cGoUnpackString(program), cGoUnpackString(fileName), int(int32(baseLineNumber)))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptLoad(f func(id int64, program string, fileName string, baseLineNumber int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scriptLoad"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scriptLoad", func(id int64, program string, fileName string, baseLineNumber int) {
				signal.(func(int64, string, string, int))(id, program, fileName, baseLineNumber)
				f(id, program, fileName, baseLineNumber)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scriptLoad", f)
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
		signal.(func(int64))(int64(id))
	} else {
		NewQScriptEngineAgentFromPointer(ptr).ScriptUnloadDefault(int64(id))
	}
}

func (ptr *QScriptEngineAgent) ConnectScriptUnload(f func(id int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scriptUnload"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scriptUnload", func(id int64) {
				signal.(func(int64))(id)
				f(id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scriptUnload", f)
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

//export callbackQScriptEngineAgent_DestroyQScriptEngineAgent
func callbackQScriptEngineAgent_DestroyQScriptEngineAgent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScriptEngineAgent"); signal != nil {
		signal.(func())()
	} else {
		NewQScriptEngineAgentFromPointer(ptr).DestroyQScriptEngineAgentDefault()
	}
}

func (ptr *QScriptEngineAgent) ConnectDestroyQScriptEngineAgent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScriptEngineAgent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineAgent", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScriptEngineAgent", f)
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
		C.QScriptEngineAgent_DestroyQScriptEngineAgent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptEngineAgent) DestroyQScriptEngineAgentDefault() {
	if ptr.Pointer() != nil {
		C.QScriptEngineAgent_DestroyQScriptEngineAgentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptEngineAgent) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptEngineFromPointer(C.QScriptEngineAgent_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScriptEngineAgent_SupportsExtension
func callbackQScriptEngineAgent_SupportsExtension(ptr unsafe.Pointer, extension C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "supportsExtension"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QScriptEngineAgent__Extension) bool)(QScriptEngineAgent__Extension(extension)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptEngineAgentFromPointer(ptr).SupportsExtensionDefault(QScriptEngineAgent__Extension(extension)))))
}

func (ptr *QScriptEngineAgent) ConnectSupportsExtension(f func(extension QScriptEngineAgent__Extension) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "supportsExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", func(extension QScriptEngineAgent__Extension) bool {
				signal.(func(QScriptEngineAgent__Extension) bool)(extension)
				return f(extension)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "supportsExtension", f)
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
		return C.QScriptEngineAgent_SupportsExtension(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
}

func (ptr *QScriptEngineAgent) SupportsExtensionDefault(extension QScriptEngineAgent__Extension) bool {
	if ptr.Pointer() != nil {
		return C.QScriptEngineAgent_SupportsExtensionDefault(ptr.Pointer(), C.longlong(extension)) != 0
	}
	return false
}

type QScriptExtensionPlugin struct {
	core.QObject
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
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
	}
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	return n
}
func NewQScriptExtensionPlugin(parent core.QObject_ITF) *QScriptExtensionPlugin {
	var tmpValue = NewQScriptExtensionPluginFromPointer(C.QScriptExtensionPlugin_NewQScriptExtensionPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScriptExtensionPlugin_Initialize
func callbackQScriptExtensionPlugin_Initialize(ptr unsafe.Pointer, key C.struct_QtScript_PackedString, engine unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialize"); signal != nil {
		signal.(func(string, *QScriptEngine))(cGoUnpackString(key), NewQScriptEngineFromPointer(engine))
	}

}

func (ptr *QScriptExtensionPlugin) ConnectInitialize(f func(key string, engine *QScriptEngine)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initialize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initialize", func(key string, engine *QScriptEngine) {
				signal.(func(string, *QScriptEngine))(key, engine)
				f(key, engine)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialize", f)
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

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		var tmpValue = NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.struct_QtScript_PackedString{data: keyC, len: C.longlong(len(key))}, PointerFromQScriptEngine(engine)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

//export callbackQScriptExtensionPlugin_Keys
func callbackQScriptExtensionPlugin_Keys(ptr unsafe.Pointer) C.struct_QtScript_PackedString {
	if signal := qt.GetSignal(ptr, "keys"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtScript_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtScript_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QScriptExtensionPlugin) ConnectKeys(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "keys"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "keys", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "keys", f)
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
		return strings.Split(cGoUnpackString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScriptExtensionPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QScriptExtensionPlugin___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptExtensionPlugin___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScriptExtensionPlugin) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScriptExtensionPlugin___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptExtensionPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptExtensionPlugin___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScriptExtensionPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptExtensionPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptExtensionPlugin___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScriptExtensionPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptExtensionPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QScriptExtensionPlugin___children_newList(ptr.Pointer()))
}

//export callbackQScriptExtensionPlugin_Event
func callbackQScriptExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScriptExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_EventFilter
func callbackQScriptExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScriptExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScriptExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScriptExtensionPlugin_ChildEvent
func callbackQScriptExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScriptExtensionPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScriptExtensionPlugin_Destroyed
func callbackQScriptExtensionPlugin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScriptExtensionPlugin_DisconnectNotify
func callbackQScriptExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScriptExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScriptExtensionPlugin_ObjectNameChanged
func callbackQScriptExtensionPlugin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScript_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScriptExtensionPlugin_TimerEvent
func callbackQScriptExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScriptExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScriptExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScriptExtensionPlugin_MetaObject
func callbackQScriptExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScriptExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScriptExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQScriptProgramFromPointer(ptr unsafe.Pointer) *QScriptProgram {
	var n = new(QScriptProgram)
	n.SetPointer(ptr)
	return n
}
func NewQScriptProgram() *QScriptProgram {
	var tmpValue = NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram())
	runtime.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
	return tmpValue
}

func NewQScriptProgram3(other QScriptProgram_ITF) *QScriptProgram {
	var tmpValue = NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram3(PointerFromQScriptProgram(other)))
	runtime.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
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
	var tmpValue = NewQScriptProgramFromPointer(C.QScriptProgram_NewQScriptProgram2(C.struct_QtScript_PackedString{data: sourceCodeC, len: C.longlong(len(sourceCode))}, C.struct_QtScript_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(firstLineNumber))))
	runtime.SetFinalizer(tmpValue, (*QScriptProgram).DestroyQScriptProgram)
	return tmpValue
}

func (ptr *QScriptProgram) DestroyQScriptProgram() {
	if ptr.Pointer() != nil {
		C.QScriptProgram_DestroyQScriptProgram(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptProgram) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptProgram_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) SourceCode() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptProgram_SourceCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptProgram) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptProgram_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptProgram) FirstLineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptProgram_FirstLineNumber(ptr.Pointer())))
	}
	return 0
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

func NewQScriptStringFromPointer(ptr unsafe.Pointer) *QScriptString {
	var n = new(QScriptString)
	n.SetPointer(ptr)
	return n
}
func NewQScriptString() *QScriptString {
	var tmpValue = NewQScriptStringFromPointer(C.QScriptString_NewQScriptString())
	runtime.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
	return tmpValue
}

func NewQScriptString2(other QScriptString_ITF) *QScriptString {
	var tmpValue = NewQScriptStringFromPointer(C.QScriptString_NewQScriptString2(PointerFromQScriptString(other)))
	runtime.SetFinalizer(tmpValue, (*QScriptString).DestroyQScriptString)
	return tmpValue
}

func (ptr *QScriptString) DestroyQScriptString() {
	if ptr.Pointer() != nil {
		C.QScriptString_DestroyQScriptString(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptString) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptString_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptString) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScriptString_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptString) ToArrayIndex(ok bool) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QScriptString_ToArrayIndex(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))))
	}
	return 0
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

func NewQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) *QScriptSyntaxCheckResult {
	var n = new(QScriptSyntaxCheckResult)
	n.SetPointer(ptr)
	return n
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
	var tmpValue = NewQScriptSyntaxCheckResultFromPointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(PointerFromQScriptSyntaxCheckResult(other)))
	runtime.SetFinalizer(tmpValue, (*QScriptSyntaxCheckResult).DestroyQScriptSyntaxCheckResult)
	return tmpValue
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {
	if ptr.Pointer() != nil {
		C.QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func NewQScriptValueFromPointer(ptr unsafe.Pointer) *QScriptValue {
	var n = new(QScriptValue)
	n.SetPointer(ptr)
	return n
}

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

func (ptr *QScriptValue) Construct2(arguments QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Construct2(ptr.Pointer(), PointerFromQScriptValue(arguments)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func NewQScriptValue() *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue())
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func (ptr *QScriptValue) Call2(thisObject QScriptValue_ITF, arguments QScriptValue_ITF) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Call2(ptr.Pointer(), PointerFromQScriptValue(thisObject), PointerFromQScriptValue(arguments)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue10(C.longlong(value)))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue11(value bool) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue11(C.char(int8(qt.GoBoolToInt(value)))))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue16(core.PointerFromQLatin1String(value)))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue2(PointerFromQScriptValue(other)))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue15(value string) *QScriptValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue15(C.struct_QtScript_PackedString{data: valueC, len: C.longlong(len(value))}))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue17(value string) *QScriptValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue17(valueC))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue12(value int) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue12(C.int(int32(value))))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func NewQScriptValue13(value uint) *QScriptValue {
	var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue13(C.uint(uint32(value))))
	runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
	return tmpValue
}

func (ptr *QScriptValue) SetData(data QScriptValue_ITF) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetData(ptr.Pointer(), PointerFromQScriptValue(data))
	}
}

func (ptr *QScriptValue) SetProperty3(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty3(ptr.Pointer(), PointerFromQScriptString(name), PointerFromQScriptValue(value), C.longlong(flags))
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

func (ptr *QScriptValue) DestroyQScriptValue() {
	if ptr.Pointer() != nil {
		C.QScriptValue_DestroyQScriptValue(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptValue) ToDateTime() *core.QDateTime {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QScriptValue_ToDateTime(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScriptValue_ToQObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) ToRegExp() *core.QRegExp {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRegExpFromPointer(C.QScriptValue_ToRegExp(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRegExp).DestroyQRegExp)
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

func (ptr *QScriptValue) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptEngineFromPointer(C.QScriptValue_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Data() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Property3(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Property3(ptr.Pointer(), PointerFromQScriptString(name), C.longlong(mode)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Property(name string, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Property(ptr.Pointer(), C.struct_QtScript_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(mode)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Property2(arrayIndex uint, mode QScriptValue__ResolveFlag) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex)), C.longlong(mode)))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Prototype() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptValue_Prototype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags2(ptr.Pointer(), PointerFromQScriptString(name), C.longlong(mode)))
	}
	return 0
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

func (ptr *QScriptValue) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScriptValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScriptValue_ToVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptValue) Equals(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_Equals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) InstanceOf(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_InstanceOf(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsFunction() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsFunction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQMetaObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQMetaObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsString() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) LessThan(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_LessThan(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValue_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_StrictlyEquals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return C.QScriptValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptValue_ToQMetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToInt32() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScriptValue_ToInt32(ptr.Pointer())))
	}
	return 0
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

func NewQScriptValueIteratorFromPointer(ptr unsafe.Pointer) *QScriptValueIterator {
	var n = new(QScriptValueIterator)
	n.SetPointer(ptr)
	return n
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

func NewQScriptableFromPointer(ptr unsafe.Pointer) *QScriptable {
	var n = new(QScriptable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptable) DestroyQScriptable() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScriptable) Context() *QScriptContext {
	if ptr.Pointer() != nil {
		return NewQScriptContextFromPointer(C.QScriptable_Context(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptable) Engine() *QScriptEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptEngineFromPointer(C.QScriptable_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScriptable) Argument(index int) *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptable_Argument(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
		return tmpValue
	}
	return nil
}

func (ptr *QScriptable) ThisObject() *QScriptValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScriptValueFromPointer(C.QScriptable_ThisObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QScriptValue).DestroyQScriptValue)
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
