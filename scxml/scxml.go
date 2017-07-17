// +build !minimal

package scxml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "scxml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtScxml_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QScxmlCompiler struct {
	ptr unsafe.Pointer
}

type QScxmlCompiler_ITF interface {
	QScxmlCompiler_PTR() *QScxmlCompiler
}

func (ptr *QScxmlCompiler) QScxmlCompiler_PTR() *QScxmlCompiler {
	return ptr
}

func (ptr *QScxmlCompiler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlCompiler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlCompiler(ptr QScxmlCompiler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlCompiler_PTR().Pointer()
	}
	return nil
}

func NewQScxmlCompilerFromPointer(ptr unsafe.Pointer) *QScxmlCompiler {
	var n = new(QScxmlCompiler)
	n.SetPointer(ptr)
	return n
}
func NewQScxmlCompiler(reader core.QXmlStreamReader_ITF) *QScxmlCompiler {
	var tmpValue = NewQScxmlCompilerFromPointer(C.QScxmlCompiler_NewQScxmlCompiler(core.PointerFromQXmlStreamReader(reader)))
	runtime.SetFinalizer(tmpValue, (*QScxmlCompiler).DestroyQScxmlCompiler)
	return tmpValue
}

func (ptr *QScxmlCompiler) Compile() *QScxmlStateMachine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlCompiler_Compile(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCompiler) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QScxmlCompiler_SetFileName(ptr.Pointer(), C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QScxmlCompiler) DestroyQScxmlCompiler() {
	if ptr.Pointer() != nil {
		C.QScxmlCompiler_DestroyQScxmlCompiler(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScxmlCompiler) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlCompiler_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlCompiler) Errors() []*QScxmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []*QScxmlError {
			var out = make([]*QScxmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlCompilerFromPointer(l.data).__errors_atList(i)
			}
			return out
		}(C.QScxmlCompiler_Errors(ptr.Pointer()))
	}
	return make([]*QScxmlError, 0)
}

func (ptr *QScxmlCompiler) __errors_atList(i int) *QScxmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlCompiler___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCompiler) __errors_setList(i QScxmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCompiler___errors_setList(ptr.Pointer(), PointerFromQScxmlError(i))
	}
}

func (ptr *QScxmlCompiler) __errors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlCompiler___errors_newList(ptr.Pointer()))
}

type QScxmlCppDataModel struct {
	QScxmlDataModel
}

type QScxmlCppDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlCppDataModel_PTR() *QScxmlCppDataModel
}

func (ptr *QScxmlCppDataModel) QScxmlCppDataModel_PTR() *QScxmlCppDataModel {
	return ptr
}

func (ptr *QScxmlCppDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlCppDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlCppDataModel(ptr QScxmlCppDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlCppDataModel_PTR().Pointer()
	}
	return nil
}

func NewQScxmlCppDataModelFromPointer(ptr unsafe.Pointer) *QScxmlCppDataModel {
	var n = new(QScxmlCppDataModel)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlCppDataModel_SetScxmlProperty
func callbackQScxmlCppDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlCppDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", func(name string, value *core.QVariant, context string) bool {
				signal.(func(string, *core.QVariant, string) bool)(name, value, context)
				return f(name, value, context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlCppDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlCppDataModel_SetScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlCppDataModel_SetScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_Setup
func callbackQScxmlCppDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setup"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(map[string]*core.QVariant) bool)(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQScxmlCppDataModelFromPointer(l.data).__setup_keyList() {
				out[i] = NewQScxmlCppDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
			}
			return out
		}(initialDataValues)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).SetupDefault(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
		var out = make(map[string]*core.QVariant, int(l.len))
		for _, i := range NewQScxmlCppDataModelFromPointer(l.data).__setup_keyList() {
			out[i] = NewQScxmlCppDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
		}
		return out
	}(initialDataValues)))))
}

func (ptr *QScxmlCppDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setup", func(initialDataValues map[string]*core.QVariant) bool {
				signal.(func(map[string]*core.QVariant) bool)(initialDataValues)
				return f(initialDataValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setup", f)
		}
	}
}

func (ptr *QScxmlCppDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setup")
	}
}

func (ptr *QScxmlCppDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlCppDataModelFromPointer(NewQScxmlCppDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_SetupDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlCppDataModelFromPointer(NewQScxmlCppDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlCppDataModel_ScxmlProperty
func callbackQScxmlCppDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlCppDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlCppDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", func(name string) *core.QVariant {
				signal.(func(string) *core.QVariant)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", f)
		}
	}
}

func (ptr *QScxmlCppDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlCppDataModel_ScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCppDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlCppDataModel_ScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlCppDataModel_HasScxmlProperty
func callbackQScxmlCppDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlCppDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlCppDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlCppDataModel_HasScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlCppDataModel_HasScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) InState(stateName string) bool {
	if ptr.Pointer() != nil {
		var stateNameC *C.char
		if stateName != "" {
			stateNameC = C.CString(stateName)
			defer C.free(unsafe.Pointer(stateNameC))
		}
		return C.QScxmlCppDataModel_InState(ptr.Pointer(), C.struct_QtScxml_PackedString{data: stateNameC, len: C.longlong(len(stateName))}) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) ScxmlEvent() *QScxmlEvent {
	if ptr.Pointer() != nil {
		return NewQScxmlEventFromPointer(C.QScxmlCppDataModel_ScxmlEvent(ptr.Pointer()))
	}
	return nil
}

type QScxmlDataModel struct {
	core.QObject
}

type QScxmlDataModel_ITF interface {
	core.QObject_ITF
	QScxmlDataModel_PTR() *QScxmlDataModel
}

func (ptr *QScxmlDataModel) QScxmlDataModel_PTR() *QScxmlDataModel {
	return ptr
}

func (ptr *QScxmlDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlDataModel(ptr QScxmlDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func NewQScxmlDataModelFromPointer(ptr unsafe.Pointer) *QScxmlDataModel {
	var n = new(QScxmlDataModel)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlDataModel_SetScxmlProperty
func callbackQScxmlDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", func(name string, value *core.QVariant, context string) bool {
				signal.(func(string, *core.QVariant, string) bool)(name, value, context)
				return f(name, value, context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlDataModel_SetScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

//export callbackQScxmlDataModel_Setup
func callbackQScxmlDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setup"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(map[string]*core.QVariant) bool)(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQScxmlDataModelFromPointer(l.data).__setup_keyList() {
				out[i] = NewQScxmlDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
			}
			return out
		}(initialDataValues)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setup", func(initialDataValues map[string]*core.QVariant) bool {
				signal.(func(map[string]*core.QVariant) bool)(initialDataValues)
				return f(initialDataValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setup", f)
		}
	}
}

func (ptr *QScxmlDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setup")
	}
}

func (ptr *QScxmlDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlDataModelFromPointer(NewQScxmlDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) SetStateMachine(stateMachine QScxmlStateMachine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetStateMachine(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

//export callbackQScxmlDataModel_StateMachineChanged
func callbackQScxmlDataModel_StateMachineChanged(ptr unsafe.Pointer, stateMachine unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stateMachineChanged"); signal != nil {
		signal.(func(*QScxmlStateMachine))(NewQScxmlStateMachineFromPointer(stateMachine))
	}

}

func (ptr *QScxmlDataModel) ConnectStateMachineChanged(f func(stateMachine *QScxmlStateMachine)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateMachineChanged") {
			C.QScxmlDataModel_ConnectStateMachineChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateMachineChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateMachineChanged", func(stateMachine *QScxmlStateMachine) {
				signal.(func(*QScxmlStateMachine))(stateMachine)
				f(stateMachine)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateMachineChanged", f)
		}
	}
}

func (ptr *QScxmlDataModel) DisconnectStateMachineChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectStateMachineChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateMachineChanged")
	}
}

func (ptr *QScxmlDataModel) StateMachineChanged(stateMachine QScxmlStateMachine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_StateMachineChanged(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

func (ptr *QScxmlDataModel) StateMachine() *QScxmlStateMachine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlDataModel_StateMachine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlDataModel_ScxmlProperty
func callbackQScxmlDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QScxmlDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", func(name string) *core.QVariant {
				signal.(func(string) *core.QVariant)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", f)
		}
	}
}

func (ptr *QScxmlDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scxmlProperty")
	}
}

func (ptr *QScxmlDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlDataModel_ScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlDataModel_HasScxmlProperty
func callbackQScxmlDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlDataModel_HasScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlDataModel___setup_initialDataValues_atList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QScxmlDataModel___setup_initialDataValues_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___setup_initialDataValues_newList(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __setup_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlDataModelFromPointer(l.data).____setup_keyList_atList(i)
			}
			return out
		}(C.QScxmlDataModel___setup_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlDataModel) ____setup_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlDataModel_____setup_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlDataModel) ____setup_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QScxmlDataModel_____setup_keyList_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QScxmlDataModel) ____setup_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel_____setup_keyList_newList(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlDataModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlDataModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlDataModel) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlDataModel) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlDataModel) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScxmlDataModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlDataModel) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDataModel___children_newList(ptr.Pointer()))
}

//export callbackQScxmlDataModel_Event
func callbackQScxmlDataModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlDataModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_EventFilter
func callbackQScxmlDataModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_ChildEvent
func callbackQScxmlDataModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlDataModel_ConnectNotify
func callbackQScxmlDataModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_CustomEvent
func callbackQScxmlDataModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlDataModel_DeleteLater
func callbackQScxmlDataModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlDataModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScxmlDataModel_Destroyed
func callbackQScxmlDataModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScxmlDataModel_DisconnectNotify
func callbackQScxmlDataModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_ObjectNameChanged
func callbackQScxmlDataModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScxmlDataModel_TimerEvent
func callbackQScxmlDataModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlDataModel_MetaObject
func callbackQScxmlDataModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlDataModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlDataModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QScxmlDynamicScxmlServiceFactory struct {
	QScxmlInvokableServiceFactory
}

type QScxmlDynamicScxmlServiceFactory_ITF interface {
	QScxmlInvokableServiceFactory_ITF
	QScxmlDynamicScxmlServiceFactory_PTR() *QScxmlDynamicScxmlServiceFactory
}

func (ptr *QScxmlDynamicScxmlServiceFactory) QScxmlDynamicScxmlServiceFactory_PTR() *QScxmlDynamicScxmlServiceFactory {
	return ptr
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlInvokableServiceFactory_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlDynamicScxmlServiceFactory(ptr QScxmlDynamicScxmlServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDynamicScxmlServiceFactory_PTR().Pointer()
	}
	return nil
}

func NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr unsafe.Pointer) *QScxmlDynamicScxmlServiceFactory {
	var n = new(QScxmlDynamicScxmlServiceFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlDynamicScxmlServiceFactory_Invoke
func callbackQScxmlDynamicScxmlServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).InvokeDefault(NewQScxmlStateMachineFromPointer(parentStateMachine)))
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "invoke"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "invoke", func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService {
				signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(parentStateMachine)
				return f(parentStateMachine)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invoke", f)
		}
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "invoke")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlDynamicScxmlServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlDynamicScxmlServiceFactory_InvokeDefault(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_names_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_names_newList(ptr.Pointer()))
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_parameters_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_parameters_newList(ptr.Pointer()))
}

type QScxmlEcmaScriptDataModel struct {
	QScxmlDataModel
}

type QScxmlEcmaScriptDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlEcmaScriptDataModel_PTR() *QScxmlEcmaScriptDataModel
}

func (ptr *QScxmlEcmaScriptDataModel) QScxmlEcmaScriptDataModel_PTR() *QScxmlEcmaScriptDataModel {
	return ptr
}

func (ptr *QScxmlEcmaScriptDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlEcmaScriptDataModel(ptr QScxmlEcmaScriptDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlEcmaScriptDataModel_PTR().Pointer()
	}
	return nil
}

func NewQScxmlEcmaScriptDataModelFromPointer(ptr unsafe.Pointer) *QScxmlEcmaScriptDataModel {
	var n = new(QScxmlEcmaScriptDataModel)
	n.SetPointer(ptr)
	return n
}
func NewQScxmlEcmaScriptDataModel(parent core.QObject_ITF) *QScxmlEcmaScriptDataModel {
	var tmpValue = NewQScxmlEcmaScriptDataModelFromPointer(C.QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlEcmaScriptDataModel_SetScxmlProperty
func callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", func(name string, value *core.QVariant, context string) bool {
				signal.(func(string, *core.QVariant, string) bool)(name, value, context)
				return f(name, value, context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlEcmaScriptDataModel_SetScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlEcmaScriptDataModel_SetScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_Setup
func callbackQScxmlEcmaScriptDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setup"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(map[string]*core.QVariant) bool)(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQScxmlEcmaScriptDataModelFromPointer(l.data).__setup_keyList() {
				out[i] = NewQScxmlEcmaScriptDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
			}
			return out
		}(initialDataValues)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetupDefault(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
		var out = make(map[string]*core.QVariant, int(l.len))
		for _, i := range NewQScxmlEcmaScriptDataModelFromPointer(l.data).__setup_keyList() {
			out[i] = NewQScxmlEcmaScriptDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
		}
		return out
	}(initialDataValues)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setup", func(initialDataValues map[string]*core.QVariant) bool {
				signal.(func(map[string]*core.QVariant) bool)(initialDataValues)
				return f(initialDataValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setup", f)
		}
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setup")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlEcmaScriptDataModelFromPointer(NewQScxmlEcmaScriptDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_SetupDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlEcmaScriptDataModelFromPointer(NewQScxmlEcmaScriptDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_SetScxmlEvent
func callbackQScxmlEcmaScriptDataModel_SetScxmlEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlEvent", func(event *QScxmlEvent) {
				signal.(func(*QScxmlEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlEvent", f)
		}
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetScxmlEventDefault(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_ScxmlProperty
func callbackQScxmlEcmaScriptDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlEcmaScriptDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", func(name string) *core.QVariant {
				signal.(func(string) *core.QVariant)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", f)
		}
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEcmaScriptDataModel_ScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEcmaScriptDataModel_ScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlEcmaScriptDataModel_HasScxmlProperty
func callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlEcmaScriptDataModel_HasScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

type QScxmlError struct {
	ptr unsafe.Pointer
}

type QScxmlError_ITF interface {
	QScxmlError_PTR() *QScxmlError
}

func (ptr *QScxmlError) QScxmlError_PTR() *QScxmlError {
	return ptr
}

func (ptr *QScxmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlError(ptr QScxmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlError_PTR().Pointer()
	}
	return nil
}

func NewQScxmlErrorFromPointer(ptr unsafe.Pointer) *QScxmlError {
	var n = new(QScxmlError)
	n.SetPointer(ptr)
	return n
}
func NewQScxmlError() *QScxmlError {
	var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError())
	runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
	return tmpValue
}

func NewQScxmlError3(other QScxmlError_ITF) *QScxmlError {
	var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError3(PointerFromQScxmlError(other)))
	runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
	return tmpValue
}

func NewQScxmlError2(fileName string, line int, column int, description string) *QScxmlError {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var descriptionC *C.char
	if description != "" {
		descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
	}
	var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError2(C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(line)), C.int(int32(column)), C.struct_QtScxml_PackedString{data: descriptionC, len: C.longlong(len(description))}))
	runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
	return tmpValue
}

func (ptr *QScxmlError) DestroyQScxmlError() {
	if ptr.Pointer() != nil {
		C.QScxmlError_DestroyQScxmlError(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScxmlError) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlError_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlError_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScxmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlError_Line(ptr.Pointer())))
	}
	return 0
}

type QScxmlEvent struct {
	ptr unsafe.Pointer
}

type QScxmlEvent_ITF interface {
	QScxmlEvent_PTR() *QScxmlEvent
}

func (ptr *QScxmlEvent) QScxmlEvent_PTR() *QScxmlEvent {
	return ptr
}

func (ptr *QScxmlEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlEvent(ptr QScxmlEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlEvent_PTR().Pointer()
	}
	return nil
}

func NewQScxmlEventFromPointer(ptr unsafe.Pointer) *QScxmlEvent {
	var n = new(QScxmlEvent)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QScxmlEvent__EventType
//QScxmlEvent::EventType
type QScxmlEvent__EventType int64

const (
	QScxmlEvent__PlatformEvent QScxmlEvent__EventType = QScxmlEvent__EventType(0)
	QScxmlEvent__InternalEvent QScxmlEvent__EventType = QScxmlEvent__EventType(1)
	QScxmlEvent__ExternalEvent QScxmlEvent__EventType = QScxmlEvent__EventType(2)
)

func NewQScxmlEvent() *QScxmlEvent {
	var tmpValue = NewQScxmlEventFromPointer(C.QScxmlEvent_NewQScxmlEvent())
	runtime.SetFinalizer(tmpValue, (*QScxmlEvent).DestroyQScxmlEvent)
	return tmpValue
}

func NewQScxmlEvent2(other QScxmlEvent_ITF) *QScxmlEvent {
	var tmpValue = NewQScxmlEventFromPointer(C.QScxmlEvent_NewQScxmlEvent2(PointerFromQScxmlEvent(other)))
	runtime.SetFinalizer(tmpValue, (*QScxmlEvent).DestroyQScxmlEvent)
	return tmpValue
}

func (ptr *QScxmlEvent) Clear() {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_Clear(ptr.Pointer())
	}
}

func (ptr *QScxmlEvent) SetData(data core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetData(ptr.Pointer(), core.PointerFromQVariant(data))
	}
}

func (ptr *QScxmlEvent) SetDelay(delayInMiliSecs int) {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetDelay(ptr.Pointer(), C.int(int32(delayInMiliSecs)))
	}
}

func (ptr *QScxmlEvent) SetErrorMessage(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QScxmlEvent_SetErrorMessage(ptr.Pointer(), C.struct_QtScxml_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *QScxmlEvent) SetEventType(ty QScxmlEvent__EventType) {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetEventType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QScxmlEvent) SetInvokeId(invokeid string) {
	if ptr.Pointer() != nil {
		var invokeidC *C.char
		if invokeid != "" {
			invokeidC = C.CString(invokeid)
			defer C.free(unsafe.Pointer(invokeidC))
		}
		C.QScxmlEvent_SetInvokeId(ptr.Pointer(), C.struct_QtScxml_PackedString{data: invokeidC, len: C.longlong(len(invokeid))})
	}
}

func (ptr *QScxmlEvent) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QScxmlEvent_SetName(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QScxmlEvent) SetOrigin(origin string) {
	if ptr.Pointer() != nil {
		var originC *C.char
		if origin != "" {
			originC = C.CString(origin)
			defer C.free(unsafe.Pointer(originC))
		}
		C.QScxmlEvent_SetOrigin(ptr.Pointer(), C.struct_QtScxml_PackedString{data: originC, len: C.longlong(len(origin))})
	}
}

func (ptr *QScxmlEvent) SetOriginType(origintype string) {
	if ptr.Pointer() != nil {
		var origintypeC *C.char
		if origintype != "" {
			origintypeC = C.CString(origintype)
			defer C.free(unsafe.Pointer(origintypeC))
		}
		C.QScxmlEvent_SetOriginType(ptr.Pointer(), C.struct_QtScxml_PackedString{data: origintypeC, len: C.longlong(len(origintype))})
	}
}

func (ptr *QScxmlEvent) SetSendId(sendid string) {
	if ptr.Pointer() != nil {
		var sendidC *C.char
		if sendid != "" {
			sendidC = C.CString(sendid)
			defer C.free(unsafe.Pointer(sendidC))
		}
		C.QScxmlEvent_SetSendId(ptr.Pointer(), C.struct_QtScxml_PackedString{data: sendidC, len: C.longlong(len(sendid))})
	}
}

func (ptr *QScxmlEvent) DestroyQScxmlEvent() {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_DestroyQScxmlEvent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QScxmlEvent) EventType() QScxmlEvent__EventType {
	if ptr.Pointer() != nil {
		return QScxmlEvent__EventType(C.QScxmlEvent_EventType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlEvent) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) InvokeId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_InvokeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) OriginType() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_OriginType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) ScxmlType() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_ScxmlType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) SendId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_SendId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) Data() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEvent_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEvent) IsErrorEvent() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEvent_IsErrorEvent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlEvent) Delay() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlEvent_Delay(ptr.Pointer())))
	}
	return 0
}

type QScxmlInvokableService struct {
	core.QObject
}

type QScxmlInvokableService_ITF interface {
	core.QObject_ITF
	QScxmlInvokableService_PTR() *QScxmlInvokableService
}

func (ptr *QScxmlInvokableService) QScxmlInvokableService_PTR() *QScxmlInvokableService {
	return ptr
}

func (ptr *QScxmlInvokableService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlInvokableService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlInvokableService(ptr QScxmlInvokableService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableService_PTR().Pointer()
	}
	return nil
}

func NewQScxmlInvokableServiceFromPointer(ptr unsafe.Pointer) *QScxmlInvokableService {
	var n = new(QScxmlInvokableService)
	n.SetPointer(ptr)
	return n
}
func NewQScxmlInvokableService(parentStateMachine QScxmlStateMachine_ITF, parent QScxmlInvokableServiceFactory_ITF) *QScxmlInvokableService {
	var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlInvokableService_NewQScxmlInvokableService(PointerFromQScxmlStateMachine(parentStateMachine), PointerFromQScxmlInvokableServiceFactory(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlInvokableService_Start
func callbackQScxmlInvokableService_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlInvokableService) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "start", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", f)
		}
	}
}

func (ptr *QScxmlInvokableService) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QScxmlInvokableService) Start() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_Start(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQScxmlInvokableService_PostEvent
func callbackQScxmlInvokableService_PostEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "postEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlInvokableService) ConnectPostEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "postEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "postEvent", func(event *QScxmlEvent) {
				signal.(func(*QScxmlEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "postEvent", f)
		}
	}
}

func (ptr *QScxmlInvokableService) DisconnectPostEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "postEvent")
	}
}

func (ptr *QScxmlInvokableService) PostEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_PostEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlInvokableService) ParentStateMachine() *QScxmlStateMachine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlInvokableService_ParentStateMachine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlInvokableService_Id
func callbackQScxmlInvokableService_Id(ptr unsafe.Pointer) C.struct_QtScxml_PackedString {
	if signal := qt.GetSignal(ptr, "id"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QScxmlInvokableService) ConnectId(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "id"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "id", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "id", f)
		}
	}
}

func (ptr *QScxmlInvokableService) DisconnectId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "id")
	}
}

func (ptr *QScxmlInvokableService) Id() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlInvokableService_Id(ptr.Pointer()))
	}
	return ""
}

//export callbackQScxmlInvokableService_Name
func callbackQScxmlInvokableService_Name(ptr unsafe.Pointer) C.struct_QtScxml_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QScxmlInvokableService) ConnectName(f func() string) {
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

func (ptr *QScxmlInvokableService) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QScxmlInvokableService) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlInvokableService_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlInvokableService___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlInvokableService) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableService___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableService) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableService) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScxmlInvokableService) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableService) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScxmlInvokableService) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableService) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableService) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableService) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableService) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableService___children_newList(ptr.Pointer()))
}

//export callbackQScxmlInvokableService_Event
func callbackQScxmlInvokableService_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlInvokableService) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlInvokableService_EventFilter
func callbackQScxmlInvokableService_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlInvokableService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlInvokableService_ChildEvent
func callbackQScxmlInvokableService_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlInvokableService_ConnectNotify
func callbackQScxmlInvokableService_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableService_CustomEvent
func callbackQScxmlInvokableService_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlInvokableService_DeleteLater
func callbackQScxmlInvokableService_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlInvokableService) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScxmlInvokableService_Destroyed
func callbackQScxmlInvokableService_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScxmlInvokableService_DisconnectNotify
func callbackQScxmlInvokableService_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableService_ObjectNameChanged
func callbackQScxmlInvokableService_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScxmlInvokableService_TimerEvent
func callbackQScxmlInvokableService_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlInvokableService_MetaObject
func callbackQScxmlInvokableService_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlInvokableServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlInvokableService) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlInvokableService_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QScxmlInvokableServiceFactory struct {
	core.QObject
}

type QScxmlInvokableServiceFactory_ITF interface {
	core.QObject_ITF
	QScxmlInvokableServiceFactory_PTR() *QScxmlInvokableServiceFactory
}

func (ptr *QScxmlInvokableServiceFactory) QScxmlInvokableServiceFactory_PTR() *QScxmlInvokableServiceFactory {
	return ptr
}

func (ptr *QScxmlInvokableServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlInvokableServiceFactory(ptr QScxmlInvokableServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func NewQScxmlInvokableServiceFactoryFromPointer(ptr unsafe.Pointer) *QScxmlInvokableServiceFactory {
	var n = new(QScxmlInvokableServiceFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlInvokableServiceFactory_Invoke
func callbackQScxmlInvokableServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(nil)
}

func (ptr *QScxmlInvokableServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "invoke"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "invoke", func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService {
				signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(parentStateMachine)
				return f(parentStateMachine)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invoke", f)
		}
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "invoke")
	}
}

func (ptr *QScxmlInvokableServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlInvokableServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_names_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_names_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_parameters_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_parameters_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __parameters_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___parameters_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __names_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___names_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlInvokableServiceFactory___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlInvokableServiceFactory) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScxmlInvokableServiceFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlInvokableServiceFactory) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlInvokableServiceFactory___children_newList(ptr.Pointer()))
}

//export callbackQScxmlInvokableServiceFactory_Event
func callbackQScxmlInvokableServiceFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlInvokableServiceFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlInvokableServiceFactory_EventFilter
func callbackQScxmlInvokableServiceFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlInvokableServiceFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlInvokableServiceFactory_ChildEvent
func callbackQScxmlInvokableServiceFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_ConnectNotify
func callbackQScxmlInvokableServiceFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableServiceFactory_CustomEvent
func callbackQScxmlInvokableServiceFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_DeleteLater
func callbackQScxmlInvokableServiceFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlInvokableServiceFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScxmlInvokableServiceFactory_Destroyed
func callbackQScxmlInvokableServiceFactory_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScxmlInvokableServiceFactory_DisconnectNotify
func callbackQScxmlInvokableServiceFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableServiceFactory_ObjectNameChanged
func callbackQScxmlInvokableServiceFactory_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScxmlInvokableServiceFactory_TimerEvent
func callbackQScxmlInvokableServiceFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_MetaObject
func callbackQScxmlInvokableServiceFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlInvokableServiceFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlInvokableServiceFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlInvokableServiceFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QScxmlNullDataModel struct {
	QScxmlDataModel
}

type QScxmlNullDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlNullDataModel_PTR() *QScxmlNullDataModel
}

func (ptr *QScxmlNullDataModel) QScxmlNullDataModel_PTR() *QScxmlNullDataModel {
	return ptr
}

func (ptr *QScxmlNullDataModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlNullDataModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlDataModel_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlNullDataModel(ptr QScxmlNullDataModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlNullDataModel_PTR().Pointer()
	}
	return nil
}

func NewQScxmlNullDataModelFromPointer(ptr unsafe.Pointer) *QScxmlNullDataModel {
	var n = new(QScxmlNullDataModel)
	n.SetPointer(ptr)
	return n
}
func NewQScxmlNullDataModel(parent core.QObject_ITF) *QScxmlNullDataModel {
	var tmpValue = NewQScxmlNullDataModelFromPointer(C.QScxmlNullDataModel_NewQScxmlNullDataModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlNullDataModel_SetScxmlProperty
func callbackQScxmlNullDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", func(name string, value *core.QVariant, context string) bool {
				signal.(func(string, *core.QVariant, string) bool)(name, value, context)
				return f(name, value, context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlNullDataModel_SetScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var contextC *C.char
		if context != "" {
			contextC = C.CString(context)
			defer C.free(unsafe.Pointer(contextC))
		}
		return C.QScxmlNullDataModel_SetScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.struct_QtScxml_PackedString{data: contextC, len: C.longlong(len(context))}) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_Setup
func callbackQScxmlNullDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setup"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(map[string]*core.QVariant) bool)(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQScxmlNullDataModelFromPointer(l.data).__setup_keyList() {
				out[i] = NewQScxmlNullDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
			}
			return out
		}(initialDataValues)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).SetupDefault(func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
		var out = make(map[string]*core.QVariant, int(l.len))
		for _, i := range NewQScxmlNullDataModelFromPointer(l.data).__setup_keyList() {
			out[i] = NewQScxmlNullDataModelFromPointer(l.data).__setup_initialDataValues_atList(i)
		}
		return out
	}(initialDataValues)))))
}

func (ptr *QScxmlNullDataModel) ConnectSetup(f func(initialDataValues map[string]*core.QVariant) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setup", func(initialDataValues map[string]*core.QVariant) bool {
				signal.(func(map[string]*core.QVariant) bool)(initialDataValues)
				return f(initialDataValues)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setup", f)
		}
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setup")
	}
}

func (ptr *QScxmlNullDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlNullDataModelFromPointer(NewQScxmlNullDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) SetupDefault(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_SetupDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlNullDataModelFromPointer(NewQScxmlNullDataModelFromPointer(nil).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_SetScxmlEvent
func callbackQScxmlNullDataModel_SetScxmlEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setScxmlEvent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlEvent", func(event *QScxmlEvent) {
				signal.(func(*QScxmlEvent))(event)
				f(event)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setScxmlEvent", f)
		}
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setScxmlEvent")
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_SetScxmlEventDefault(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) DestroyQScxmlNullDataModel() {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DestroyQScxmlNullDataModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScxmlNullDataModel_ScxmlProperty
func callbackQScxmlNullDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlNullDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlNullDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "scxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", func(name string) *core.QVariant {
				signal.(func(string) *core.QVariant)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "scxmlProperty", f)
		}
	}
}

func (ptr *QScxmlNullDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "scxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlNullDataModel_ScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlNullDataModel_ScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlNullDataModel_HasScxmlProperty
func callbackQScxmlNullDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlNullDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasScxmlProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", func(name string) bool {
				signal.(func(string) bool)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasScxmlProperty", f)
		}
	}
}

func (ptr *QScxmlNullDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlNullDataModel_HasScxmlProperty(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return C.QScxmlNullDataModel_HasScxmlPropertyDefault(ptr.Pointer(), C.struct_QtScxml_PackedString{data: nameC, len: C.longlong(len(name))}) != 0
	}
	return false
}

type QScxmlStateMachine struct {
	core.QObject
}

type QScxmlStateMachine_ITF interface {
	core.QObject_ITF
	QScxmlStateMachine_PTR() *QScxmlStateMachine
}

func (ptr *QScxmlStateMachine) QScxmlStateMachine_PTR() *QScxmlStateMachine {
	return ptr
}

func (ptr *QScxmlStateMachine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlStateMachine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlStateMachine(ptr QScxmlStateMachine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlStateMachine_PTR().Pointer()
	}
	return nil
}

func NewQScxmlStateMachineFromPointer(ptr unsafe.Pointer) *QScxmlStateMachine {
	var n = new(QScxmlStateMachine)
	n.SetPointer(ptr)
	return n
}
func QScxmlStateMachine_FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScxmlStateMachine) FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QScxmlStateMachine_FromFile(fileName string) *QScxmlStateMachine {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScxmlStateMachine) FromFile(fileName string) *QScxmlStateMachine {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(C.struct_QtScxml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQScxmlStateMachine(metaObject core.QMetaObject_ITF, parent core.QObject_ITF) *QScxmlStateMachine {
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_NewQScxmlStateMachine(core.PointerFromQMetaObject(metaObject), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScxmlStateMachine) InitialValues() map[string]*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) map[string]*core.QVariant {
			var out = make(map[string]*core.QVariant, int(l.len))
			for _, i := range NewQScxmlStateMachineFromPointer(l.data).__initialValues_keyList() {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).__initialValues_atList(i)
			}
			return out
		}(C.QScxmlStateMachine_InitialValues(ptr.Pointer()))
	}
	return make(map[string]*core.QVariant, 0)
}

//export callbackQScxmlStateMachine_Init
func callbackQScxmlStateMachine_Init(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "init"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).InitDefault())))
}

func (ptr *QScxmlStateMachine) ConnectInit(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "init"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "init", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "init", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectInit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "init")
	}
}

func (ptr *QScxmlStateMachine) Init() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_Init(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) InitDefault() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_InitDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) CancelDelayedEvent(sendId string) {
	if ptr.Pointer() != nil {
		var sendIdC *C.char
		if sendId != "" {
			sendIdC = C.CString(sendId)
			defer C.free(unsafe.Pointer(sendIdC))
		}
		C.QScxmlStateMachine_CancelDelayedEvent(ptr.Pointer(), C.struct_QtScxml_PackedString{data: sendIdC, len: C.longlong(len(sendId))})
	}
}

//export callbackQScxmlStateMachine_DataModelChanged
func callbackQScxmlStateMachine_DataModelChanged(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dataModelChanged"); signal != nil {
		signal.(func(*QScxmlDataModel))(NewQScxmlDataModelFromPointer(model))
	}

}

func (ptr *QScxmlStateMachine) ConnectDataModelChanged(f func(model *QScxmlDataModel)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataModelChanged") {
			C.QScxmlStateMachine_ConnectDataModelChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataModelChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataModelChanged", func(model *QScxmlDataModel) {
				signal.(func(*QScxmlDataModel))(model)
				f(model)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataModelChanged", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectDataModelChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectDataModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataModelChanged")
	}
}

func (ptr *QScxmlStateMachine) DataModelChanged(model QScxmlDataModel_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DataModelChanged(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

//export callbackQScxmlStateMachine_Finished
func callbackQScxmlStateMachine_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QScxmlStateMachine_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QScxmlStateMachine) Finished() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Finished(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_InitializedChanged
func callbackQScxmlStateMachine_InitializedChanged(ptr unsafe.Pointer, initialized C.char) {
	if signal := qt.GetSignal(ptr, "initializedChanged"); signal != nil {
		signal.(func(bool))(int8(initialized) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectInitializedChanged(f func(initialized bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "initializedChanged") {
			C.QScxmlStateMachine_ConnectInitializedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "initializedChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initializedChanged", func(initialized bool) {
				signal.(func(bool))(initialized)
				f(initialized)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initializedChanged", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectInitializedChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectInitializedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "initializedChanged")
	}
}

func (ptr *QScxmlStateMachine) InitializedChanged(initialized bool) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_InitializedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(initialized))))
	}
}

//export callbackQScxmlStateMachine_InvokedServicesChanged
func callbackQScxmlStateMachine_InvokedServicesChanged(ptr unsafe.Pointer, invokedServices C.struct_QtScxml_PackedList) {
	if signal := qt.GetSignal(ptr, "invokedServicesChanged"); signal != nil {
		signal.(func([]*QScxmlInvokableService))(func(l C.struct_QtScxml_PackedList) []*QScxmlInvokableService {
			var out = make([]*QScxmlInvokableService, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).__invokedServicesChanged_invokedServices_atList(i)
			}
			return out
		}(invokedServices))
	}

}

func (ptr *QScxmlStateMachine) ConnectInvokedServicesChanged(f func(invokedServices []*QScxmlInvokableService)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "invokedServicesChanged") {
			C.QScxmlStateMachine_ConnectInvokedServicesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "invokedServicesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "invokedServicesChanged", func(invokedServices []*QScxmlInvokableService) {
				signal.(func([]*QScxmlInvokableService))(invokedServices)
				f(invokedServices)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invokedServicesChanged", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectInvokedServicesChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectInvokedServicesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "invokedServicesChanged")
	}
}

func (ptr *QScxmlStateMachine) InvokedServicesChanged(invokedServices []*QScxmlInvokableService) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_InvokedServicesChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlStateMachineFromPointer(NewQScxmlStateMachineFromPointer(nil).__invokedServicesChanged_invokedServices_newList())
			for _, v := range invokedServices {
				tmpList.__invokedServicesChanged_invokedServices_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQScxmlStateMachine_Log
func callbackQScxmlStateMachine_Log(ptr unsafe.Pointer, label C.struct_QtScxml_PackedString, msg C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(ptr, "log"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(label), cGoUnpackString(msg))
	}

}

func (ptr *QScxmlStateMachine) ConnectLog(f func(label string, msg string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "log") {
			C.QScxmlStateMachine_ConnectLog(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "log"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "log", func(label string, msg string) {
				signal.(func(string, string))(label, msg)
				f(label, msg)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "log", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectLog() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectLog(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "log")
	}
}

func (ptr *QScxmlStateMachine) Log(label string, msg string) {
	if ptr.Pointer() != nil {
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QScxmlStateMachine_Log(ptr.Pointer(), C.struct_QtScxml_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_QtScxml_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQScxmlStateMachine_ReachedStableState
func callbackQScxmlStateMachine_ReachedStableState(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reachedStableState"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectReachedStableState(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "reachedStableState") {
			C.QScxmlStateMachine_ConnectReachedStableState(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "reachedStableState"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reachedStableState", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reachedStableState", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectReachedStableState() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectReachedStableState(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "reachedStableState")
	}
}

func (ptr *QScxmlStateMachine) ReachedStableState() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ReachedStableState(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_RunningChanged
func callbackQScxmlStateMachine_RunningChanged(ptr unsafe.Pointer, running C.char) {
	if signal := qt.GetSignal(ptr, "runningChanged"); signal != nil {
		signal.(func(bool))(int8(running) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectRunningChanged(f func(running bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "runningChanged") {
			C.QScxmlStateMachine_ConnectRunningChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "runningChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "runningChanged", func(running bool) {
				signal.(func(bool))(running)
				f(running)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "runningChanged", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectRunningChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "runningChanged")
	}
}

func (ptr *QScxmlStateMachine) RunningChanged(running bool) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_RunningChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(running))))
	}
}

func (ptr *QScxmlStateMachine) SetDataModel(model QScxmlDataModel_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetDataModel(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

func (ptr *QScxmlStateMachine) SetInitialValues(initialValues map[string]*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetInitialValues(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlStateMachineFromPointer(NewQScxmlStateMachineFromPointer(nil).__setInitialValues_initialValues_newList())
			for k, v := range initialValues {
				tmpList.__setInitialValues_initialValues_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QScxmlStateMachine) SetRunning(running bool) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetRunning(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(running))))
	}
}

func (ptr *QScxmlStateMachine) SetTableData(tableData QScxmlTableData_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetTableData(ptr.Pointer(), PointerFromQScxmlTableData(tableData))
	}
}

//export callbackQScxmlStateMachine_Start
func callbackQScxmlStateMachine_Start(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).StartDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "start", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QScxmlStateMachine) Start() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Start(ptr.Pointer())
	}
}

func (ptr *QScxmlStateMachine) StartDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_StartDefault(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_Stop
func callbackQScxmlStateMachine_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).StopDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QScxmlStateMachine) Stop() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Stop(ptr.Pointer())
	}
}

func (ptr *QScxmlStateMachine) StopDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_StopDefault(ptr.Pointer())
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SubmitEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent2(eventName string) {
	if ptr.Pointer() != nil {
		var eventNameC *C.char
		if eventName != "" {
			eventNameC = C.CString(eventName)
			defer C.free(unsafe.Pointer(eventNameC))
		}
		C.QScxmlStateMachine_SubmitEvent2(ptr.Pointer(), C.struct_QtScxml_PackedString{data: eventNameC, len: C.longlong(len(eventName))})
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent3(eventName string, data core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var eventNameC *C.char
		if eventName != "" {
			eventNameC = C.CString(eventName)
			defer C.free(unsafe.Pointer(eventNameC))
		}
		C.QScxmlStateMachine_SubmitEvent3(ptr.Pointer(), C.struct_QtScxml_PackedString{data: eventNameC, len: C.longlong(len(eventName))}, core.PointerFromQVariant(data))
	}
}

//export callbackQScxmlStateMachine_TableDataChanged
func callbackQScxmlStateMachine_TableDataChanged(ptr unsafe.Pointer, tableData unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tableDataChanged"); signal != nil {
		signal.(func(*QScxmlTableData))(NewQScxmlTableDataFromPointer(tableData))
	}

}

func (ptr *QScxmlStateMachine) ConnectTableDataChanged(f func(tableData *QScxmlTableData)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "tableDataChanged") {
			C.QScxmlStateMachine_ConnectTableDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "tableDataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "tableDataChanged", func(tableData *QScxmlTableData) {
				signal.(func(*QScxmlTableData))(tableData)
				f(tableData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "tableDataChanged", f)
		}
	}
}

func (ptr *QScxmlStateMachine) DisconnectTableDataChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectTableDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "tableDataChanged")
	}
}

func (ptr *QScxmlStateMachine) TableDataChanged(tableData QScxmlTableData_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TableDataChanged(ptr.Pointer(), PointerFromQScxmlTableData(tableData))
	}
}

func (ptr *QScxmlStateMachine) DataModel() *QScxmlDataModel {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlDataModelFromPointer(C.QScxmlStateMachine_DataModel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) TableData() *QScxmlTableData {
	if ptr.Pointer() != nil {
		return NewQScxmlTableDataFromPointer(C.QScxmlStateMachine_TableData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlStateMachine) SessionId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ActiveStateNames(compress bool) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScxmlStateMachine_ActiveStateNames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(compress))))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) StateNames(compress bool) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScxmlStateMachine_StateNames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(compress))))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) ParseErrors() []*QScxmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []*QScxmlError {
			var out = make([]*QScxmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).__parseErrors_atList(i)
			}
			return out
		}(C.QScxmlStateMachine_ParseErrors(ptr.Pointer()))
	}
	return make([]*QScxmlError, 0)
}

func (ptr *QScxmlStateMachine) InvokedServices() []*QScxmlInvokableService {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []*QScxmlInvokableService {
			var out = make([]*QScxmlInvokableService, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).__invokedServices_atList(i)
			}
			return out
		}(C.QScxmlStateMachine_InvokedServices(ptr.Pointer()))
	}
	return make([]*QScxmlInvokableService, 0)
}

func (ptr *QScxmlStateMachine) IsActive(scxmlStateName string) bool {
	if ptr.Pointer() != nil {
		var scxmlStateNameC *C.char
		if scxmlStateName != "" {
			scxmlStateNameC = C.CString(scxmlStateName)
			defer C.free(unsafe.Pointer(scxmlStateNameC))
		}
		return C.QScxmlStateMachine_IsActive(ptr.Pointer(), C.struct_QtScxml_PackedString{data: scxmlStateNameC, len: C.longlong(len(scxmlStateName))}) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsActive2(stateIndex int) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsActive2(ptr.Pointer(), C.int(int32(stateIndex))) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsDispatchableTarget(target string) bool {
	if ptr.Pointer() != nil {
		var targetC *C.char
		if target != "" {
			targetC = C.CString(target)
			defer C.free(unsafe.Pointer(targetC))
		}
		return C.QScxmlStateMachine_IsDispatchableTarget(ptr.Pointer(), C.struct_QtScxml_PackedString{data: targetC, len: C.longlong(len(target))}) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsInvoked() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsInvoked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) __initialValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___initialValues_atList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QScxmlStateMachine___initialValues_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __initialValues_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___initialValues_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __initialValues_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).____initialValues_keyList_atList(i)
			}
			return out
		}(C.QScxmlStateMachine___initialValues_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___initialValuesChanged_initialValues_atList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QScxmlStateMachine___initialValuesChanged_initialValues_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___initialValuesChanged_initialValues_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).____initialValuesChanged_keyList_atList(i)
			}
			return out
		}(C.QScxmlStateMachine___initialValuesChanged_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_atList(i int) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStateMachine___invokedServicesChanged_invokedServices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_setList(i QScxmlInvokableService_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___invokedServicesChanged_invokedServices_setList(ptr.Pointer(), PointerFromQScxmlInvokableService(i))
	}
}

func (ptr *QScxmlStateMachine) __invokedServicesChanged_invokedServices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___invokedServicesChanged_invokedServices_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___setInitialValues_initialValues_atList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QScxmlStateMachine___setInitialValues_initialValues_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___setInitialValues_initialValues_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __setInitialValues_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).____setInitialValues_keyList_atList(i)
			}
			return out
		}(C.QScxmlStateMachine___setInitialValues_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) __parseErrors_atList(i int) *QScxmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlStateMachine___parseErrors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __parseErrors_setList(i QScxmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___parseErrors_setList(ptr.Pointer(), PointerFromQScxmlError(i))
	}
}

func (ptr *QScxmlStateMachine) __parseErrors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___parseErrors_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __invokedServices_atList(i int) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStateMachine___invokedServices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __invokedServices_setList(i QScxmlInvokableService_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___invokedServices_setList(ptr.Pointer(), PointerFromQScxmlInvokableService(i))
	}
}

func (ptr *QScxmlStateMachine) __invokedServices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___invokedServices_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____initialValues_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QScxmlStateMachine_____initialValues_keyList_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine_____initialValues_keyList_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____initialValuesChanged_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QScxmlStateMachine_____initialValuesChanged_keyList_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine_____initialValuesChanged_keyList_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____setInitialValues_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QScxmlStateMachine_____setInitialValues_keyList_setList(ptr.Pointer(), C.struct_QtScxml_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine_____setInitialValues_keyList_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlStateMachine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlStateMachine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStateMachine) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStateMachine) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStateMachine) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList(ptr.Pointer()))
}

func (ptr *QScxmlStateMachine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStateMachine) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStateMachine___children_newList(ptr.Pointer()))
}

//export callbackQScxmlStateMachine_Event
func callbackQScxmlStateMachine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlStateMachine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_EventFilter
func callbackQScxmlStateMachine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlStateMachine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_ChildEvent
func callbackQScxmlStateMachine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlStateMachine_ConnectNotify
func callbackQScxmlStateMachine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_CustomEvent
func callbackQScxmlStateMachine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlStateMachine_DeleteLater
func callbackQScxmlStateMachine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlStateMachine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQScxmlStateMachine_Destroyed
func callbackQScxmlStateMachine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQScxmlStateMachine_DisconnectNotify
func callbackQScxmlStateMachine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_ObjectNameChanged
func callbackQScxmlStateMachine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQScxmlStateMachine_TimerEvent
func callbackQScxmlStateMachine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlStateMachine_MetaObject
func callbackQScxmlStateMachine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlStateMachineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlStateMachine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStateMachine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QScxmlStaticScxmlServiceFactory struct {
	QScxmlInvokableServiceFactory
}

type QScxmlStaticScxmlServiceFactory_ITF interface {
	QScxmlInvokableServiceFactory_ITF
	QScxmlStaticScxmlServiceFactory_PTR() *QScxmlStaticScxmlServiceFactory
}

func (ptr *QScxmlStaticScxmlServiceFactory) QScxmlStaticScxmlServiceFactory_PTR() *QScxmlStaticScxmlServiceFactory {
	return ptr
}

func (ptr *QScxmlStaticScxmlServiceFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlInvokableServiceFactory_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QScxmlInvokableServiceFactory_PTR().SetPointer(p)
	}
}

func PointerFromQScxmlStaticScxmlServiceFactory(ptr QScxmlStaticScxmlServiceFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlStaticScxmlServiceFactory_PTR().Pointer()
	}
	return nil
}

func NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr unsafe.Pointer) *QScxmlStaticScxmlServiceFactory {
	var n = new(QScxmlStaticScxmlServiceFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlStaticScxmlServiceFactory_Invoke
func callbackQScxmlStaticScxmlServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).InvokeDefault(NewQScxmlStateMachineFromPointer(parentStateMachine)))
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "invoke"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "invoke", func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService {
				signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(parentStateMachine)
				return f(parentStateMachine)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "invoke", f)
		}
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "invoke")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStaticScxmlServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStaticScxmlServiceFactory_InvokeDefault(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_nameList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_nameList_newList(ptr.Pointer()))
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_parameters_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_parameters_newList(ptr.Pointer()))
}

type QScxmlTableData struct {
	ptr unsafe.Pointer
}

type QScxmlTableData_ITF interface {
	QScxmlTableData_PTR() *QScxmlTableData
}

func (ptr *QScxmlTableData) QScxmlTableData_PTR() *QScxmlTableData {
	return ptr
}

func (ptr *QScxmlTableData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlTableData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlTableData(ptr QScxmlTableData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlTableData_PTR().Pointer()
	}
	return nil
}

func NewQScxmlTableDataFromPointer(ptr unsafe.Pointer) *QScxmlTableData {
	var n = new(QScxmlTableData)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlTableData_DestroyQScxmlTableData
func callbackQScxmlTableData_DestroyQScxmlTableData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QScxmlTableData"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlTableDataFromPointer(ptr).DestroyQScxmlTableDataDefault()
	}
}

func (ptr *QScxmlTableData) ConnectDestroyQScxmlTableData(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QScxmlTableData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QScxmlTableData", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QScxmlTableData", f)
		}
	}
}

func (ptr *QScxmlTableData) DisconnectDestroyQScxmlTableData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QScxmlTableData")
	}
}

func (ptr *QScxmlTableData) DestroyQScxmlTableData() {
	if ptr.Pointer() != nil {
		C.QScxmlTableData_DestroyQScxmlTableData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlTableData) DestroyQScxmlTableDataDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlTableData_DestroyQScxmlTableDataDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlTableData_ServiceFactory
func callbackQScxmlTableData_ServiceFactory(ptr unsafe.Pointer, id C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "serviceFactory"); signal != nil {
		return PointerFromQScxmlInvokableServiceFactory(signal.(func(int) *QScxmlInvokableServiceFactory)(int(int32(id))))
	}

	return PointerFromQScxmlInvokableServiceFactory(nil)
}

func (ptr *QScxmlTableData) ConnectServiceFactory(f func(id int) *QScxmlInvokableServiceFactory) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "serviceFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceFactory", func(id int) *QScxmlInvokableServiceFactory {
				signal.(func(int) *QScxmlInvokableServiceFactory)(id)
				return f(id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceFactory", f)
		}
	}
}

func (ptr *QScxmlTableData) DisconnectServiceFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "serviceFactory")
	}
}

func (ptr *QScxmlTableData) ServiceFactory(id int) *QScxmlInvokableServiceFactory {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFactoryFromPointer(C.QScxmlTableData_ServiceFactory(ptr.Pointer(), C.int(int32(id))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlTableData_Name
func callbackQScxmlTableData_Name(ptr unsafe.Pointer) C.struct_QtScxml_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtScxml_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QScxmlTableData) ConnectName(f func() string) {
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

func (ptr *QScxmlTableData) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QScxmlTableData) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlTableData_Name(ptr.Pointer()))
	}
	return ""
}
