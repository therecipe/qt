// +build !minimal

package scxml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "scxml.h"
import "C"
import (
	"fmt"
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
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCompiler) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QScxmlCompiler_SetFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QScxmlCompiler) DestroyQScxmlCompiler() {
	if ptr.Pointer() != nil {
		C.QScxmlCompiler_DestroyQScxmlCompiler(ptr.Pointer())
		ptr.SetPointer(nil)
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlCompiler___errors_newList(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlCppDataModel) DestroyQScxmlCppDataModel() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlCppDataModel_SetScxmlProperty
func callbackQScxmlCppDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlCppDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::setScxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::setScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlCppDataModel_SetScxmlProperty(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlCppDataModel_SetScxmlPropertyDefault(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_Setup
func callbackQScxmlCppDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::setup"); signal != nil {
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

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::setup", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::setup")
	}
}

func (ptr *QScxmlCppDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlCppDataModelFromPointer(NewQScxmlCppDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
			var tmpList = NewQScxmlCppDataModelFromPointer(NewQScxmlCppDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlCppDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlCppDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::scxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::scxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlCppDataModel_ScxmlProperty(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCppDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlCppDataModel_ScxmlPropertyDefault(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlCppDataModel_HasScxmlProperty
func callbackQScxmlCppDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlCppDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::hasScxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::hasScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlCppDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlCppDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) InState(stateName string) bool {
	if ptr.Pointer() != nil {
		var stateNameC = C.CString(stateName)
		defer C.free(unsafe.Pointer(stateNameC))
		return C.QScxmlCppDataModel_InState(ptr.Pointer(), stateNameC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) ScxmlEvent() *QScxmlEvent {
	if ptr.Pointer() != nil {
		return NewQScxmlEventFromPointer(C.QScxmlCppDataModel_ScxmlEvent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlCppDataModel) __setup_initialDataValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlCppDataModel___setup_initialDataValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlCppDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlCppDataModel___setup_initialDataValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlCppDataModel) __setup_initialDataValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlCppDataModel___setup_initialDataValues_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlCppDataModel) __setup_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlCppDataModelFromPointer(l.data).____setup_keyList_atList(i)
			}
			return out
		}(C.QScxmlCppDataModel___setup_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlCppDataModel) ____setup_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlCppDataModel_____setup_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlCppDataModel) ____setup_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlCppDataModel_____setup_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlCppDataModel) ____setup_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlCppDataModel_____setup_keyList_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlCppDataModel_Event
func callbackQScxmlCppDataModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlCppDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::event", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::event")
	}
}

func (ptr *QScxmlCppDataModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_EventFilter
func callbackQScxmlCppDataModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlCppDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::eventFilter", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::eventFilter")
	}
}

func (ptr *QScxmlCppDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_ChildEvent
func callbackQScxmlCppDataModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::childEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::childEvent")
	}
}

func (ptr *QScxmlCppDataModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlCppDataModel_ConnectNotify
func callbackQScxmlCppDataModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::connectNotify", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::connectNotify")
	}
}

func (ptr *QScxmlCppDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlCppDataModel_CustomEvent
func callbackQScxmlCppDataModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::customEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::customEvent")
	}
}

func (ptr *QScxmlCppDataModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlCppDataModel_DeleteLater
func callbackQScxmlCppDataModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlCppDataModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::deleteLater", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::deleteLater")
	}
}

func (ptr *QScxmlCppDataModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlCppDataModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlCppDataModel_DisconnectNotify
func callbackQScxmlCppDataModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::disconnectNotify", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::disconnectNotify")
	}
}

func (ptr *QScxmlCppDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlCppDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlCppDataModel_TimerEvent
func callbackQScxmlCppDataModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::timerEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::timerEvent")
	}
}

func (ptr *QScxmlCppDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlCppDataModel_MetaObject
func callbackQScxmlCppDataModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlCppDataModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlCppDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlCppDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::metaObject", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlCppDataModel::metaObject")
	}
}

func (ptr *QScxmlCppDataModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlCppDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlCppDataModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlCppDataModel_MetaObjectDefault(ptr.Pointer()))
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

func (ptr *QScxmlDataModel) DestroyQScxmlDataModel() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlDataModel_SetScxmlProperty
func callbackQScxmlDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setScxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlDataModel_SetScxmlProperty(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

//export callbackQScxmlDataModel_Setup
func callbackQScxmlDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::setup"); signal != nil {
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

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setup", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setup")
	}
}

func (ptr *QScxmlDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlDataModelFromPointer(NewQScxmlDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
			for k, v := range initialDataValues {
				tmpList.__setup_initialDataValues_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackQScxmlDataModel_SetScxmlEvent
func callbackQScxmlDataModel_SetScxmlEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setScxmlEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::setScxmlEvent")
	}
}

func (ptr *QScxmlDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlDataModel) SetStateMachine(stateMachine QScxmlStateMachine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetStateMachine(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

//export callbackQScxmlDataModel_StateMachineChanged
func callbackQScxmlDataModel_StateMachineChanged(ptr unsafe.Pointer, stateMachine unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::stateMachineChanged"); signal != nil {
		signal.(func(*QScxmlStateMachine))(NewQScxmlStateMachineFromPointer(stateMachine))
	}

}

func (ptr *QScxmlDataModel) ConnectStateMachineChanged(f func(stateMachine *QScxmlStateMachine)) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectStateMachineChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::stateMachineChanged", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectStateMachineChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectStateMachineChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::stateMachineChanged")
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
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlDataModel_ScxmlProperty
func callbackQScxmlDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QScxmlDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::scxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::scxmlProperty")
	}
}

func (ptr *QScxmlDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlDataModel_ScxmlProperty(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlDataModel_HasScxmlProperty
func callbackQScxmlDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::hasScxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::hasScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlDataModel___setup_initialDataValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlDataModel___setup_initialDataValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlDataModel) __setup_initialDataValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___setup_initialDataValues_newList(ptr.Pointer()))
	}
	return nil
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
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlDataModel_____setup_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlDataModel) ____setup_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel_____setup_keyList_newList(ptr.Pointer()))
	}
	return nil
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDataModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDataModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlDataModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDataModel___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlDataModel_Event
func callbackQScxmlDataModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::event", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::event")
	}
}

func (ptr *QScxmlDataModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_EventFilter
func callbackQScxmlDataModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::eventFilter", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::eventFilter")
	}
}

func (ptr *QScxmlDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_ChildEvent
func callbackQScxmlDataModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::childEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::childEvent")
	}
}

func (ptr *QScxmlDataModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlDataModel_ConnectNotify
func callbackQScxmlDataModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::connectNotify", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::connectNotify")
	}
}

func (ptr *QScxmlDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_CustomEvent
func callbackQScxmlDataModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::customEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::customEvent")
	}
}

func (ptr *QScxmlDataModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlDataModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlDataModel_DeleteLater
func callbackQScxmlDataModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlDataModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::deleteLater", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::deleteLater")
	}
}

func (ptr *QScxmlDataModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlDataModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlDataModel_DisconnectNotify
func callbackQScxmlDataModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::disconnectNotify", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::disconnectNotify")
	}
}

func (ptr *QScxmlDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_TimerEvent
func callbackQScxmlDataModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::timerEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::timerEvent")
	}
}

func (ptr *QScxmlDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlDataModel_MetaObject
func callbackQScxmlDataModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDataModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::metaObject", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDataModel::metaObject")
	}
}

func (ptr *QScxmlDataModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlDynamicScxmlServiceFactory) DestroyQScxmlDynamicScxmlServiceFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_Invoke
func callbackQScxmlDynamicScxmlServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).InvokeDefault(NewQScxmlStateMachineFromPointer(parentStateMachine)))
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::invoke", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::invoke")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlDynamicScxmlServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlDynamicScxmlServiceFactory_InvokeDefault(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_names_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlDynamicScxmlServiceFactory_parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlDynamicScxmlServiceFactory_parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlInvokableServiceFactory_names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlInvokableServiceFactory_names_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __QScxmlInvokableServiceFactory_parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___QScxmlInvokableServiceFactory_parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) __names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlDynamicScxmlServiceFactory___names_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlDynamicScxmlServiceFactory_Event
func callbackQScxmlDynamicScxmlServiceFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::event", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::event")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDynamicScxmlServiceFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlDynamicScxmlServiceFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDynamicScxmlServiceFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlDynamicScxmlServiceFactory_EventFilter
func callbackQScxmlDynamicScxmlServiceFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::eventFilter", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::eventFilter")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDynamicScxmlServiceFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlDynamicScxmlServiceFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlDynamicScxmlServiceFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlDynamicScxmlServiceFactory_ChildEvent
func callbackQScxmlDynamicScxmlServiceFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::childEvent", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::childEvent")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_ConnectNotify
func callbackQScxmlDynamicScxmlServiceFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::connectNotify", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::connectNotify")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_CustomEvent
func callbackQScxmlDynamicScxmlServiceFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::customEvent", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::customEvent")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_DeleteLater
func callbackQScxmlDynamicScxmlServiceFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::deleteLater", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::deleteLater")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_DisconnectNotify
func callbackQScxmlDynamicScxmlServiceFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::disconnectNotify", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::disconnectNotify")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_TimerEvent
func callbackQScxmlDynamicScxmlServiceFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::timerEvent", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::timerEvent")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDynamicScxmlServiceFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlDynamicScxmlServiceFactory_MetaObject
func callbackQScxmlDynamicScxmlServiceFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlDynamicScxmlServiceFactory::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlDynamicScxmlServiceFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlDynamicScxmlServiceFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::metaObject", f)
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlDynamicScxmlServiceFactory::metaObject")
	}
}

func (ptr *QScxmlDynamicScxmlServiceFactory) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlDynamicScxmlServiceFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDynamicScxmlServiceFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlDynamicScxmlServiceFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlEcmaScriptDataModel) DestroyQScxmlEcmaScriptDataModel() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQScxmlEcmaScriptDataModel(parent core.QObject_ITF) *QScxmlEcmaScriptDataModel {
	var tmpValue = NewQScxmlEcmaScriptDataModelFromPointer(C.QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlEcmaScriptDataModel_SetScxmlProperty
func callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setScxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlEcmaScriptDataModel_SetScxmlProperty(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlEcmaScriptDataModel_SetScxmlPropertyDefault(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_Setup
func callbackQScxmlEcmaScriptDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::setup"); signal != nil {
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

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setup", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setup")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlEcmaScriptDataModelFromPointer(NewQScxmlEcmaScriptDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
			var tmpList = NewQScxmlEcmaScriptDataModelFromPointer(NewQScxmlEcmaScriptDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setScxmlEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::setScxmlEvent")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlEcmaScriptDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::scxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::scxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEcmaScriptDataModel_ScxmlProperty(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEcmaScriptDataModel_ScxmlPropertyDefault(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlEcmaScriptDataModel_HasScxmlProperty
func callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::hasScxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::hasScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlEcmaScriptDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) __setup_initialDataValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEcmaScriptDataModel___setup_initialDataValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlEcmaScriptDataModel___setup_initialDataValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __setup_initialDataValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___setup_initialDataValues_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __setup_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlEcmaScriptDataModelFromPointer(l.data).____setup_keyList_atList(i)
			}
			return out
		}(C.QScxmlEcmaScriptDataModel___setup_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlEcmaScriptDataModel) ____setup_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEcmaScriptDataModel_____setup_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlEcmaScriptDataModel) ____setup_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlEcmaScriptDataModel_____setup_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ____setup_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel_____setup_keyList_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlEcmaScriptDataModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlEcmaScriptDataModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlEcmaScriptDataModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlEcmaScriptDataModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlEcmaScriptDataModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlEcmaScriptDataModel___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlEcmaScriptDataModel_Event
func callbackQScxmlEcmaScriptDataModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::event", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::event")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_EventFilter
func callbackQScxmlEcmaScriptDataModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::eventFilter", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::eventFilter")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_ChildEvent
func callbackQScxmlEcmaScriptDataModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::childEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::childEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_ConnectNotify
func callbackQScxmlEcmaScriptDataModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::connectNotify", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::connectNotify")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlEcmaScriptDataModel_CustomEvent
func callbackQScxmlEcmaScriptDataModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::customEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::customEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_DeleteLater
func callbackQScxmlEcmaScriptDataModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::deleteLater", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::deleteLater")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlEcmaScriptDataModel_DisconnectNotify
func callbackQScxmlEcmaScriptDataModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::disconnectNotify", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::disconnectNotify")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlEcmaScriptDataModel_TimerEvent
func callbackQScxmlEcmaScriptDataModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::timerEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::timerEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_MetaObject
func callbackQScxmlEcmaScriptDataModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEcmaScriptDataModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlEcmaScriptDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::metaObject", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEcmaScriptDataModel::metaObject")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlEcmaScriptDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlEcmaScriptDataModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var descriptionC = C.CString(description)
	defer C.free(unsafe.Pointer(descriptionC))
	var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError2(fileNameC, C.int(int32(line)), C.int(int32(column)), descriptionC))
	runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
	return tmpValue
}

func (ptr *QScxmlError) DestroyQScxmlError() {
	if ptr.Pointer() != nil {
		C.QScxmlError_DestroyQScxmlError(ptr.Pointer())
		ptr.SetPointer(nil)
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
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		C.QScxmlEvent_SetErrorMessage(ptr.Pointer(), messageC)
	}
}

func (ptr *QScxmlEvent) SetEventType(ty QScxmlEvent__EventType) {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetEventType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QScxmlEvent) SetInvokeId(invokeid string) {
	if ptr.Pointer() != nil {
		var invokeidC = C.CString(invokeid)
		defer C.free(unsafe.Pointer(invokeidC))
		C.QScxmlEvent_SetInvokeId(ptr.Pointer(), invokeidC)
	}
}

func (ptr *QScxmlEvent) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QScxmlEvent_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QScxmlEvent) SetOrigin(origin string) {
	if ptr.Pointer() != nil {
		var originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
		C.QScxmlEvent_SetOrigin(ptr.Pointer(), originC)
	}
}

func (ptr *QScxmlEvent) SetOriginType(origintype string) {
	if ptr.Pointer() != nil {
		var origintypeC = C.CString(origintype)
		defer C.free(unsafe.Pointer(origintypeC))
		C.QScxmlEvent_SetOriginType(ptr.Pointer(), origintypeC)
	}
}

func (ptr *QScxmlEvent) SetSendId(sendid string) {
	if ptr.Pointer() != nil {
		var sendidC = C.CString(sendid)
		defer C.free(unsafe.Pointer(sendidC))
		C.QScxmlEvent_SetSendId(ptr.Pointer(), sendidC)
	}
}

func (ptr *QScxmlEvent) DestroyQScxmlEvent() {
	if ptr.Pointer() != nil {
		C.QScxmlEvent_DestroyQScxmlEvent(ptr.Pointer())
		ptr.SetPointer(nil)
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

func (ptr *QScxmlInvokableService) DestroyQScxmlInvokableService() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQScxmlInvokableService(parentStateMachine QScxmlStateMachine_ITF, parent QScxmlInvokableServiceFactory_ITF) *QScxmlInvokableService {
	var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlInvokableService_NewQScxmlInvokableService(PointerFromQScxmlStateMachine(parentStateMachine), PointerFromQScxmlInvokableServiceFactory(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlInvokableService_Start
func callbackQScxmlInvokableService_Start(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::start"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlInvokableService) ConnectStart(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::start", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::start")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::postEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlInvokableService) ConnectPostEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::postEvent", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectPostEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::postEvent")
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
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlInvokableService_Id
func callbackQScxmlInvokableService_Id(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::id"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QScxmlInvokableService) ConnectId(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::id", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::id")
	}
}

func (ptr *QScxmlInvokableService) Id() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlInvokableService_Id(ptr.Pointer()))
	}
	return ""
}

//export callbackQScxmlInvokableService_Name
func callbackQScxmlInvokableService_Name(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::name"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QScxmlInvokableService) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::name", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::name")
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableService___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableService) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableService___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableService) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableService___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableService___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlInvokableService_Event
func callbackQScxmlInvokableService_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlInvokableService) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::event", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::event")
	}
}

func (ptr *QScxmlInvokableService) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlInvokableService) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlInvokableService_EventFilter
func callbackQScxmlInvokableService_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlInvokableService) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::eventFilter", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::eventFilter")
	}
}

func (ptr *QScxmlInvokableService) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlInvokableService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlInvokableService_ChildEvent
func callbackQScxmlInvokableService_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::childEvent", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::childEvent")
	}
}

func (ptr *QScxmlInvokableService) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlInvokableService) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlInvokableService_ConnectNotify
func callbackQScxmlInvokableService_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableService) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::connectNotify", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::connectNotify")
	}
}

func (ptr *QScxmlInvokableService) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlInvokableService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableService_CustomEvent
func callbackQScxmlInvokableService_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::customEvent", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::customEvent")
	}
}

func (ptr *QScxmlInvokableService) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlInvokableService) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlInvokableService_DeleteLater
func callbackQScxmlInvokableService_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlInvokableService) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::deleteLater", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::deleteLater")
	}
}

func (ptr *QScxmlInvokableService) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlInvokableService) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlInvokableService_DisconnectNotify
func callbackQScxmlInvokableService_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableService) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::disconnectNotify", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::disconnectNotify")
	}
}

func (ptr *QScxmlInvokableService) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlInvokableService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableService_TimerEvent
func callbackQScxmlInvokableService_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::timerEvent", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::timerEvent")
	}
}

func (ptr *QScxmlInvokableService) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlInvokableService) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlInvokableService_MetaObject
func callbackQScxmlInvokableService_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableService::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlInvokableServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlInvokableService) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::metaObject", f)
	}
}

func (ptr *QScxmlInvokableService) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableService::metaObject")
	}
}

func (ptr *QScxmlInvokableService) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlInvokableService_MetaObject(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlInvokableServiceFactory) DestroyQScxmlInvokableServiceFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlInvokableServiceFactory_Invoke
func callbackQScxmlInvokableServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(nil)
}

func (ptr *QScxmlInvokableServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::invoke", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::invoke")
	}
}

func (ptr *QScxmlInvokableServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlInvokableServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_names_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __QScxmlInvokableServiceFactory_parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___QScxmlInvokableServiceFactory_parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___names_newList(ptr.Pointer()))
	}
	return nil
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlInvokableServiceFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlInvokableServiceFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlInvokableServiceFactory___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlInvokableServiceFactory_Event
func callbackQScxmlInvokableServiceFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlInvokableServiceFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::event", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::event")
	}
}

func (ptr *QScxmlInvokableServiceFactory) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlInvokableServiceFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlInvokableServiceFactory_EventFilter
func callbackQScxmlInvokableServiceFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlInvokableServiceFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlInvokableServiceFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::eventFilter", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::eventFilter")
	}
}

func (ptr *QScxmlInvokableServiceFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlInvokableServiceFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlInvokableServiceFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlInvokableServiceFactory_ChildEvent
func callbackQScxmlInvokableServiceFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::childEvent", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::childEvent")
	}
}

func (ptr *QScxmlInvokableServiceFactory) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_ConnectNotify
func callbackQScxmlInvokableServiceFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::connectNotify", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::connectNotify")
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableServiceFactory_CustomEvent
func callbackQScxmlInvokableServiceFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::customEvent", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::customEvent")
	}
}

func (ptr *QScxmlInvokableServiceFactory) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_DeleteLater
func callbackQScxmlInvokableServiceFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::deleteLater", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::deleteLater")
	}
}

func (ptr *QScxmlInvokableServiceFactory) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlInvokableServiceFactory_DisconnectNotify
func callbackQScxmlInvokableServiceFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::disconnectNotify", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::disconnectNotify")
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlInvokableServiceFactory_TimerEvent
func callbackQScxmlInvokableServiceFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlInvokableServiceFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::timerEvent", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::timerEvent")
	}
}

func (ptr *QScxmlInvokableServiceFactory) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlInvokableServiceFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlInvokableServiceFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlInvokableServiceFactory_MetaObject
func callbackQScxmlInvokableServiceFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlInvokableServiceFactory::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlInvokableServiceFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlInvokableServiceFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::metaObject", f)
	}
}

func (ptr *QScxmlInvokableServiceFactory) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlInvokableServiceFactory::metaObject")
	}
}

func (ptr *QScxmlInvokableServiceFactory) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlInvokableServiceFactory_MetaObject(ptr.Pointer()))
	}
	return nil
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
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQScxmlNullDataModel_SetScxmlProperty
func callbackQScxmlNullDataModel_SetScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString, value unsafe.Pointer, context C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::setScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).SetScxmlPropertyDefault(cGoUnpackString(name), core.NewQVariantFromPointer(value), cGoUnpackString(context)))))
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setScxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlNullDataModel_SetScxmlProperty(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) SetScxmlPropertyDefault(name string, value core.QVariant_ITF, context string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var contextC = C.CString(context)
		defer C.free(unsafe.Pointer(contextC))
		return C.QScxmlNullDataModel_SetScxmlPropertyDefault(ptr.Pointer(), nameC, core.PointerFromQVariant(value), contextC) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_Setup
func callbackQScxmlNullDataModel_Setup(ptr unsafe.Pointer, initialDataValues C.struct_QtScxml_PackedList) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::setup"); signal != nil {
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

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setup", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setup")
	}
}

func (ptr *QScxmlNullDataModel) Setup(initialDataValues map[string]*core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_Setup(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlNullDataModelFromPointer(NewQScxmlNullDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
			var tmpList = NewQScxmlNullDataModelFromPointer(NewQScxmlNullDataModelFromPointer(unsafe.Pointer(uintptr(1))).__setup_initialDataValues_newList())
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setScxmlEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::setScxmlEvent")
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
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlNullDataModel_ScxmlProperty
func callbackQScxmlNullDataModel_ScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(cGoUnpackString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlNullDataModelFromPointer(ptr).ScxmlPropertyDefault(cGoUnpackString(name)))
}

func (ptr *QScxmlNullDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::scxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::scxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) ScxmlProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlNullDataModel_ScxmlProperty(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) ScxmlPropertyDefault(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlNullDataModel_ScxmlPropertyDefault(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQScxmlNullDataModel_HasScxmlProperty
func callbackQScxmlNullDataModel_HasScxmlProperty(ptr unsafe.Pointer, name C.struct_QtScxml_PackedString) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::hasScxmlProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).HasScxmlPropertyDefault(cGoUnpackString(name)))))
}

func (ptr *QScxmlNullDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::hasScxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectHasScxmlProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::hasScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) HasScxmlProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlNullDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) HasScxmlPropertyDefault(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlNullDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) __setup_initialDataValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlNullDataModel___setup_initialDataValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __setup_initialDataValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlNullDataModel___setup_initialDataValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlNullDataModel) __setup_initialDataValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___setup_initialDataValues_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __setup_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlNullDataModelFromPointer(l.data).____setup_keyList_atList(i)
			}
			return out
		}(C.QScxmlNullDataModel___setup_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QScxmlNullDataModel) ____setup_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlNullDataModel_____setup_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlNullDataModel) ____setup_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlNullDataModel_____setup_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlNullDataModel) ____setup_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel_____setup_keyList_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlNullDataModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlNullDataModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlNullDataModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlNullDataModel) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlNullDataModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlNullDataModel) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlNullDataModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlNullDataModel) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlNullDataModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlNullDataModel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlNullDataModel) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlNullDataModel___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlNullDataModel_Event
func callbackQScxmlNullDataModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlNullDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::event", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::event")
	}
}

func (ptr *QScxmlNullDataModel) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_EventFilter
func callbackQScxmlNullDataModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlNullDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::eventFilter", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::eventFilter")
	}
}

func (ptr *QScxmlNullDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_ChildEvent
func callbackQScxmlNullDataModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::childEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::childEvent")
	}
}

func (ptr *QScxmlNullDataModel) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlNullDataModel_ConnectNotify
func callbackQScxmlNullDataModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::connectNotify", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::connectNotify")
	}
}

func (ptr *QScxmlNullDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlNullDataModel_CustomEvent
func callbackQScxmlNullDataModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::customEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::customEvent")
	}
}

func (ptr *QScxmlNullDataModel) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlNullDataModel_DeleteLater
func callbackQScxmlNullDataModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlNullDataModel) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::deleteLater", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::deleteLater")
	}
}

func (ptr *QScxmlNullDataModel) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlNullDataModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlNullDataModel_DisconnectNotify
func callbackQScxmlNullDataModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::disconnectNotify", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::disconnectNotify")
	}
}

func (ptr *QScxmlNullDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlNullDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlNullDataModel_TimerEvent
func callbackQScxmlNullDataModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::timerEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::timerEvent")
	}
}

func (ptr *QScxmlNullDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlNullDataModel_MetaObject
func callbackQScxmlNullDataModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlNullDataModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlNullDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlNullDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::metaObject", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlNullDataModel::metaObject")
	}
}

func (ptr *QScxmlNullDataModel) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlNullDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlNullDataModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlStateMachine) DestroyQScxmlStateMachine() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QScxmlStateMachine_FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), fileNameC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScxmlStateMachine) FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), fileNameC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QScxmlStateMachine_FromFile(fileName string) *QScxmlStateMachine {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(fileNameC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QScxmlStateMachine) FromFile(fileName string) *QScxmlStateMachine {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(fileNameC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQScxmlStateMachine(metaObject core.QMetaObject_ITF, parent core.QObject_ITF) *QScxmlStateMachine {
	var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_NewQScxmlStateMachine(core.PointerFromQMetaObject(metaObject), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::init"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).InitDefault())))
}

func (ptr *QScxmlStateMachine) ConnectInit(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::init", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectInit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::init")
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
		var sendIdC = C.CString(sendId)
		defer C.free(unsafe.Pointer(sendIdC))
		C.QScxmlStateMachine_CancelDelayedEvent(ptr.Pointer(), sendIdC)
	}
}

//export callbackQScxmlStateMachine_DataModelChanged
func callbackQScxmlStateMachine_DataModelChanged(ptr unsafe.Pointer, model unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::dataModelChanged"); signal != nil {
		signal.(func(*QScxmlDataModel))(NewQScxmlDataModelFromPointer(model))
	}

}

func (ptr *QScxmlStateMachine) ConnectDataModelChanged(f func(model *QScxmlDataModel)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectDataModelChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::dataModelChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDataModelChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectDataModelChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::dataModelChanged")
	}
}

func (ptr *QScxmlStateMachine) DataModelChanged(model QScxmlDataModel_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DataModelChanged(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

//export callbackQScxmlStateMachine_Finished
func callbackQScxmlStateMachine_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::finished", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::finished")
	}
}

func (ptr *QScxmlStateMachine) Finished() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Finished(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_InitializedChanged
func callbackQScxmlStateMachine_InitializedChanged(ptr unsafe.Pointer, initialized C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::initializedChanged"); signal != nil {
		signal.(func(bool))(int8(initialized) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectInitializedChanged(f func(initialized bool)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectInitializedChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::initializedChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectInitializedChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectInitializedChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::initializedChanged")
	}
}

func (ptr *QScxmlStateMachine) InitializedChanged(initialized bool) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_InitializedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(initialized))))
	}
}

//export callbackQScxmlStateMachine_InvokedServicesChanged
func callbackQScxmlStateMachine_InvokedServicesChanged(ptr unsafe.Pointer, invokedServices C.struct_QtScxml_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::invokedServicesChanged"); signal != nil {
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
		C.QScxmlStateMachine_ConnectInvokedServicesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::invokedServicesChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectInvokedServicesChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectInvokedServicesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::invokedServicesChanged")
	}
}

func (ptr *QScxmlStateMachine) InvokedServicesChanged(invokedServices []*QScxmlInvokableService) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_InvokedServicesChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQScxmlStateMachineFromPointer(NewQScxmlStateMachineFromPointer(unsafe.Pointer(uintptr(1))).__invokedServicesChanged_invokedServices_newList())
			for _, v := range invokedServices {
				tmpList.__invokedServicesChanged_invokedServices_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQScxmlStateMachine_Log
func callbackQScxmlStateMachine_Log(ptr unsafe.Pointer, label C.struct_QtScxml_PackedString, msg C.struct_QtScxml_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::log"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(label), cGoUnpackString(msg))
	}

}

func (ptr *QScxmlStateMachine) ConnectLog(f func(label string, msg string)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectLog(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::log", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectLog() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectLog(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::log")
	}
}

func (ptr *QScxmlStateMachine) Log(label string, msg string) {
	if ptr.Pointer() != nil {
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QScxmlStateMachine_Log(ptr.Pointer(), labelC, msgC)
	}
}

//export callbackQScxmlStateMachine_ReachedStableState
func callbackQScxmlStateMachine_ReachedStableState(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::reachedStableState"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectReachedStableState(f func()) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectReachedStableState(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::reachedStableState", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectReachedStableState() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectReachedStableState(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::reachedStableState")
	}
}

func (ptr *QScxmlStateMachine) ReachedStableState() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ReachedStableState(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_RunningChanged
func callbackQScxmlStateMachine_RunningChanged(ptr unsafe.Pointer, running C.char) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::runningChanged"); signal != nil {
		signal.(func(bool))(int8(running) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectRunningChanged(f func(running bool)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectRunningChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::runningChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectRunningChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::runningChanged")
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
			var tmpList = NewQScxmlStateMachineFromPointer(NewQScxmlStateMachineFromPointer(unsafe.Pointer(uintptr(1))).__setInitialValues_initialValues_newList())
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::start"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).StartDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::start", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::start")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::stop"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).StopDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::stop", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::stop")
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
		var eventNameC = C.CString(eventName)
		defer C.free(unsafe.Pointer(eventNameC))
		C.QScxmlStateMachine_SubmitEvent2(ptr.Pointer(), eventNameC)
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent3(eventName string, data core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var eventNameC = C.CString(eventName)
		defer C.free(unsafe.Pointer(eventNameC))
		C.QScxmlStateMachine_SubmitEvent3(ptr.Pointer(), eventNameC, core.PointerFromQVariant(data))
	}
}

//export callbackQScxmlStateMachine_TableDataChanged
func callbackQScxmlStateMachine_TableDataChanged(ptr unsafe.Pointer, tableData unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::tableDataChanged"); signal != nil {
		signal.(func(*QScxmlTableData))(NewQScxmlTableDataFromPointer(tableData))
	}

}

func (ptr *QScxmlStateMachine) ConnectTableDataChanged(f func(tableData *QScxmlTableData)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectTableDataChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::tableDataChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectTableDataChanged() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectTableDataChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::tableDataChanged")
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
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
		var scxmlStateNameC = C.CString(scxmlStateName)
		defer C.free(unsafe.Pointer(scxmlStateNameC))
		return C.QScxmlStateMachine_IsActive(ptr.Pointer(), scxmlStateNameC) != 0
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
		var targetC = C.CString(target)
		defer C.free(unsafe.Pointer(targetC))
		return C.QScxmlStateMachine_IsDispatchableTarget(ptr.Pointer(), targetC) != 0
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
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___initialValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlStateMachine___initialValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __initialValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___initialValues_newList(ptr.Pointer()))
	}
	return nil
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
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___initialValuesChanged_initialValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlStateMachine___initialValuesChanged_initialValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __initialValuesChanged_initialValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___initialValuesChanged_initialValues_newList(ptr.Pointer()))
	}
	return nil
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
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___invokedServicesChanged_invokedServices_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlStateMachine___setInitialValues_initialValues_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QScxmlStateMachine___setInitialValues_initialValues_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
	}
}

func (ptr *QScxmlStateMachine) __setInitialValues_initialValues_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___setInitialValues_initialValues_newList(ptr.Pointer()))
	}
	return nil
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___parseErrors_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __invokedServices_atList(i int) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStateMachine___invokedServices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___invokedServices_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____initialValues_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlStateMachine_____initialValues_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlStateMachine) ____initialValues_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine_____initialValues_keyList_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____initialValuesChanged_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlStateMachine_____initialValuesChanged_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlStateMachine) ____initialValuesChanged_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine_____initialValuesChanged_keyList_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_____setInitialValues_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QScxmlStateMachine_____setInitialValues_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QScxmlStateMachine) ____setInitialValues_keyList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine_____setInitialValues_keyList_newList(ptr.Pointer()))
	}
	return nil
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStateMachine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStateMachine___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlStateMachine_Event
func callbackQScxmlStateMachine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlStateMachine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::event", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::event")
	}
}

func (ptr *QScxmlStateMachine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_EventFilter
func callbackQScxmlStateMachine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlStateMachine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::eventFilter", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::eventFilter")
	}
}

func (ptr *QScxmlStateMachine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_ChildEvent
func callbackQScxmlStateMachine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::childEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::childEvent")
	}
}

func (ptr *QScxmlStateMachine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlStateMachine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlStateMachine_ConnectNotify
func callbackQScxmlStateMachine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::connectNotify", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::connectNotify")
	}
}

func (ptr *QScxmlStateMachine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_CustomEvent
func callbackQScxmlStateMachine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::customEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::customEvent")
	}
}

func (ptr *QScxmlStateMachine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlStateMachine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlStateMachine_DeleteLater
func callbackQScxmlStateMachine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::deleteLater", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::deleteLater")
	}
}

func (ptr *QScxmlStateMachine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlStateMachine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlStateMachine_DisconnectNotify
func callbackQScxmlStateMachine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::disconnectNotify", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::disconnectNotify")
	}
}

func (ptr *QScxmlStateMachine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStateMachine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_TimerEvent
func callbackQScxmlStateMachine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::timerEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::timerEvent")
	}
}

func (ptr *QScxmlStateMachine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlStateMachine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlStateMachine_MetaObject
func callbackQScxmlStateMachine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlStateMachineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlStateMachine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::metaObject", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::metaObject")
	}
}

func (ptr *QScxmlStateMachine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStateMachine_MetaObject(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlStaticScxmlServiceFactory) DestroyQScxmlStaticScxmlServiceFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_Invoke
func callbackQScxmlStaticScxmlServiceFactory_Invoke(ptr unsafe.Pointer, parentStateMachine unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::invoke"); signal != nil {
		return PointerFromQScxmlInvokableService(signal.(func(*QScxmlStateMachine) *QScxmlInvokableService)(NewQScxmlStateMachineFromPointer(parentStateMachine)))
	}

	return PointerFromQScxmlInvokableService(NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).InvokeDefault(NewQScxmlStateMachineFromPointer(parentStateMachine)))
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectInvoke(f func(parentStateMachine *QScxmlStateMachine) *QScxmlInvokableService) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::invoke", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectInvoke() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::invoke")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) Invoke(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStaticScxmlServiceFactory_Invoke(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) InvokeDefault(parentStateMachine QScxmlStateMachine_ITF) *QScxmlInvokableService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFromPointer(C.QScxmlStaticScxmlServiceFactory_InvokeDefault(ptr.Pointer(), PointerFromQScxmlStateMachine(parentStateMachine)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_nameList_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_nameList_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlStaticScxmlServiceFactory_parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlStaticScxmlServiceFactory_parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlInvokableServiceFactory_names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlInvokableServiceFactory_names_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __QScxmlInvokableServiceFactory_parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___QScxmlInvokableServiceFactory_parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __parameters_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___parameters_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __names_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___names_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QScxmlStaticScxmlServiceFactory___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStaticScxmlServiceFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStaticScxmlServiceFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStaticScxmlServiceFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QScxmlStaticScxmlServiceFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QScxmlStaticScxmlServiceFactory___children_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlStaticScxmlServiceFactory_Event
func callbackQScxmlStaticScxmlServiceFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::event", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::event")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStaticScxmlServiceFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlStaticScxmlServiceFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStaticScxmlServiceFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlStaticScxmlServiceFactory_EventFilter
func callbackQScxmlStaticScxmlServiceFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::eventFilter", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::eventFilter")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStaticScxmlServiceFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlStaticScxmlServiceFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStaticScxmlServiceFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlStaticScxmlServiceFactory_ChildEvent
func callbackQScxmlStaticScxmlServiceFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::childEvent", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::childEvent")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_ConnectNotify
func callbackQScxmlStaticScxmlServiceFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::connectNotify", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::connectNotify")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_CustomEvent
func callbackQScxmlStaticScxmlServiceFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::customEvent", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::customEvent")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_DeleteLater
func callbackQScxmlStaticScxmlServiceFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::deleteLater", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::deleteLater")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_DisconnectNotify
func callbackQScxmlStaticScxmlServiceFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::disconnectNotify", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::disconnectNotify")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_TimerEvent
func callbackQScxmlStaticScxmlServiceFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::timerEvent", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::timerEvent")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStaticScxmlServiceFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlStaticScxmlServiceFactory_MetaObject
func callbackQScxmlStaticScxmlServiceFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStaticScxmlServiceFactory::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlStaticScxmlServiceFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlStaticScxmlServiceFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::metaObject", f)
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStaticScxmlServiceFactory::metaObject")
	}
}

func (ptr *QScxmlStaticScxmlServiceFactory) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStaticScxmlServiceFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStaticScxmlServiceFactory) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStaticScxmlServiceFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlTableData::~QScxmlTableData"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlTableDataFromPointer(ptr).DestroyQScxmlTableDataDefault()
	}
}

func (ptr *QScxmlTableData) ConnectDestroyQScxmlTableData(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::~QScxmlTableData", f)
	}
}

func (ptr *QScxmlTableData) DisconnectDestroyQScxmlTableData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::~QScxmlTableData")
	}
}

func (ptr *QScxmlTableData) DestroyQScxmlTableData() {
	if ptr.Pointer() != nil {
		C.QScxmlTableData_DestroyQScxmlTableData(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlTableData) DestroyQScxmlTableDataDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlTableData_DestroyQScxmlTableDataDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlTableData_ServiceFactory
func callbackQScxmlTableData_ServiceFactory(ptr unsafe.Pointer, id C.int) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlTableData::serviceFactory"); signal != nil {
		return PointerFromQScxmlInvokableServiceFactory(signal.(func(int) *QScxmlInvokableServiceFactory)(int(int32(id))))
	}

	return PointerFromQScxmlInvokableServiceFactory(nil)
}

func (ptr *QScxmlTableData) ConnectServiceFactory(f func(id int) *QScxmlInvokableServiceFactory) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::serviceFactory", f)
	}
}

func (ptr *QScxmlTableData) DisconnectServiceFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::serviceFactory")
	}
}

func (ptr *QScxmlTableData) ServiceFactory(id int) *QScxmlInvokableServiceFactory {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlInvokableServiceFactoryFromPointer(C.QScxmlTableData_ServiceFactory(ptr.Pointer(), C.int(int32(id))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQScxmlTableData_Name
func callbackQScxmlTableData_Name(ptr unsafe.Pointer) *C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlTableData::name"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QScxmlTableData) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::name", f)
	}
}

func (ptr *QScxmlTableData) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlTableData::name")
	}
}

func (ptr *QScxmlTableData) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlTableData_Name(ptr.Pointer()))
	}
	return ""
}
