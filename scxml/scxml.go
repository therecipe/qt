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
	"github.com/therecipe/qt/qml"
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

func (ptr *QScxmlCppDataModel) In(stateName string) bool {
	if ptr.Pointer() != nil {
		var stateNameC = C.CString(stateName)
		defer C.free(unsafe.Pointer(stateNameC))
		return C.QScxmlCppDataModel_In(ptr.Pointer(), stateNameC) != 0
	}
	return false
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

func (ptr *QScxmlCppDataModel) ScxmlEvent() *QScxmlEvent {
	if ptr.Pointer() != nil {
		return NewQScxmlEventFromPointer(C.QScxmlCppDataModel_ScxmlEvent(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlCppDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
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

func (ptr *QScxmlDataModel) SetStateMachine(stateMachine QScxmlStateMachine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetStateMachine(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
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

func (ptr *QScxmlEcmaScriptDataModel) Engine() *qml.QJSEngine {
	if ptr.Pointer() != nil {
		var tmpValue = qml.NewQJSEngineFromPointer(C.QScxmlEcmaScriptDataModel_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
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

func (ptr *QScxmlEcmaScriptDataModel) SetEngine(engine qml.QJSEngine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetEngine(ptr.Pointer(), qml.PointerFromQJSEngine(engine))
	}
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

func (ptr *QScxmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlError_Column(ptr.Pointer())))
	}
	return 0
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

func (ptr *QScxmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlError_Line(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScxmlError) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) DestroyQScxmlError() {
	if ptr.Pointer() != nil {
		C.QScxmlError_DestroyQScxmlError(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QScxmlEvent struct {
	core.QEvent
}

type QScxmlEvent_ITF interface {
	core.QEvent_ITF
	QScxmlEvent_PTR() *QScxmlEvent
}

func (ptr *QScxmlEvent) QScxmlEvent_PTR() *QScxmlEvent {
	return ptr
}

func (ptr *QScxmlEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QScxmlEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
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

func (ptr *QScxmlEvent) Data() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEvent_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEvent) Delay() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScxmlEvent_Delay(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScxmlEvent) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) EventType() QScxmlEvent__EventType {
	if ptr.Pointer() != nil {
		return QScxmlEvent__EventType(C.QScxmlEvent_EventType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlEvent) InvokeId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlEvent_InvokeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) IsErrorEvent() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEvent_IsErrorEvent(ptr.Pointer()) != 0
	}
	return false
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

type QScxmlEventFilter struct {
	ptr unsafe.Pointer
}

type QScxmlEventFilter_ITF interface {
	QScxmlEventFilter_PTR() *QScxmlEventFilter
}

func (ptr *QScxmlEventFilter) QScxmlEventFilter_PTR() *QScxmlEventFilter {
	return ptr
}

func (ptr *QScxmlEventFilter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlEventFilter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlEventFilter(ptr QScxmlEventFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlEventFilter_PTR().Pointer()
	}
	return nil
}

func NewQScxmlEventFilterFromPointer(ptr unsafe.Pointer) *QScxmlEventFilter {
	var n = new(QScxmlEventFilter)
	n.SetPointer(ptr)
	return n
}

//export callbackQScxmlEventFilter_Handle
func callbackQScxmlEventFilter_Handle(ptr unsafe.Pointer, event unsafe.Pointer, stateMachine unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEventFilter::handle"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QScxmlEvent, *QScxmlStateMachine) bool)(NewQScxmlEventFromPointer(event), NewQScxmlStateMachineFromPointer(stateMachine)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QScxmlEventFilter) ConnectHandle(f func(event *QScxmlEvent, stateMachine *QScxmlStateMachine) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEventFilter::handle", f)
	}
}

func (ptr *QScxmlEventFilter) DisconnectHandle() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEventFilter::handle")
	}
}

func (ptr *QScxmlEventFilter) Handle(event QScxmlEvent_ITF, stateMachine QScxmlStateMachine_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScxmlEventFilter_Handle(ptr.Pointer(), PointerFromQScxmlEvent(event), PointerFromQScxmlStateMachine(stateMachine)) != 0
	}
	return false
}

//export callbackQScxmlEventFilter_DestroyQScxmlEventFilter
func callbackQScxmlEventFilter_DestroyQScxmlEventFilter(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlEventFilter::~QScxmlEventFilter"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlEventFilterFromPointer(ptr).DestroyQScxmlEventFilterDefault()
	}
}

func (ptr *QScxmlEventFilter) ConnectDestroyQScxmlEventFilter(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEventFilter::~QScxmlEventFilter", f)
	}
}

func (ptr *QScxmlEventFilter) DisconnectDestroyQScxmlEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlEventFilter::~QScxmlEventFilter")
	}
}

func (ptr *QScxmlEventFilter) DestroyQScxmlEventFilter() {
	if ptr.Pointer() != nil {
		C.QScxmlEventFilter_DestroyQScxmlEventFilter(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlEventFilter) DestroyQScxmlEventFilterDefault() {
	if ptr.Pointer() != nil {
		C.QScxmlEventFilter_DestroyQScxmlEventFilterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
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

func (ptr *QScxmlNullDataModel) DestroyQScxmlNullDataModel() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQScxmlNullDataModel(parent core.QObject_ITF) *QScxmlNullDataModel {
	var tmpValue = NewQScxmlNullDataModelFromPointer(C.QScxmlNullDataModel_NewQScxmlNullDataModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
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

type QScxmlParser struct {
	ptr unsafe.Pointer
}

type QScxmlParser_ITF interface {
	QScxmlParser_PTR() *QScxmlParser
}

func (ptr *QScxmlParser) QScxmlParser_PTR() *QScxmlParser {
	return ptr
}

func (ptr *QScxmlParser) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QScxmlParser) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQScxmlParser(ptr QScxmlParser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScxmlParser_PTR().Pointer()
	}
	return nil
}

func NewQScxmlParserFromPointer(ptr unsafe.Pointer) *QScxmlParser {
	var n = new(QScxmlParser)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QScxmlParser__QtMode
//QScxmlParser::QtMode
type QScxmlParser__QtMode int64

const (
	QScxmlParser__QtModeDisabled      QScxmlParser__QtMode = QScxmlParser__QtMode(0)
	QScxmlParser__QtModeEnabled       QScxmlParser__QtMode = QScxmlParser__QtMode(1)
	QScxmlParser__QtModeFromInputFile QScxmlParser__QtMode = QScxmlParser__QtMode(2)
)

func NewQScxmlParser(reader core.QXmlStreamReader_ITF) *QScxmlParser {
	var tmpValue = NewQScxmlParserFromPointer(C.QScxmlParser_NewQScxmlParser(core.PointerFromQXmlStreamReader(reader)))
	runtime.SetFinalizer(tmpValue, (*QScxmlParser).DestroyQScxmlParser)
	return tmpValue
}

func (ptr *QScxmlParser) AddError(msg string) {
	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QScxmlParser_AddError(ptr.Pointer(), msgC)
	}
}

func (ptr *QScxmlParser) Errors() []*QScxmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []*QScxmlError {
			var out = make([]*QScxmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlParserFromPointer(l.data).errors_atList(i)
			}
			return out
		}(C.QScxmlParser_Errors(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlParser) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlParser_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlParser) InstantiateDataModel(stateMachine QScxmlStateMachine_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlParser_InstantiateDataModel(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

func (ptr *QScxmlParser) InstantiateStateMachine() *QScxmlStateMachine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlStateMachineFromPointer(C.QScxmlParser_InstantiateStateMachine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlParser) Parse() {
	if ptr.Pointer() != nil {
		C.QScxmlParser_Parse(ptr.Pointer())
	}
}

func (ptr *QScxmlParser) QtMode() QScxmlParser__QtMode {
	if ptr.Pointer() != nil {
		return QScxmlParser__QtMode(C.QScxmlParser_QtMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlParser) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QScxmlParser_SetFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QScxmlParser) SetQtMode(mode QScxmlParser__QtMode) {
	if ptr.Pointer() != nil {
		C.QScxmlParser_SetQtMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QScxmlParser) DestroyQScxmlParser() {
	if ptr.Pointer() != nil {
		C.QScxmlParser_DestroyQScxmlParser(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlParser) errors_atList(i int) *QScxmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlParser_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
		return tmpValue
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

//go:generate stringer -type=QScxmlStateMachine__BindingMethod
//QScxmlStateMachine::BindingMethod
type QScxmlStateMachine__BindingMethod int64

const (
	QScxmlStateMachine__EarlyBinding QScxmlStateMachine__BindingMethod = QScxmlStateMachine__BindingMethod(0)
	QScxmlStateMachine__LateBinding  QScxmlStateMachine__BindingMethod = QScxmlStateMachine__BindingMethod(1)
)

func (ptr *QScxmlStateMachine) IsInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) ActiveStateNames(compress bool) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScxmlStateMachine_ActiveStateNames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(compress))))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) CancelDelayedEvent(sendId string) {
	if ptr.Pointer() != nil {
		var sendIdC = C.CString(sendId)
		defer C.free(unsafe.Pointer(sendIdC))
		C.QScxmlStateMachine_CancelDelayedEvent(ptr.Pointer(), sendIdC)
	}
}

func (ptr *QScxmlStateMachine) DataBinding() QScxmlStateMachine__BindingMethod {
	if ptr.Pointer() != nil {
		return QScxmlStateMachine__BindingMethod(C.QScxmlStateMachine_DataBinding(ptr.Pointer()))
	}
	return 0
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

//export callbackQScxmlStateMachine_EventOccurred
func callbackQScxmlStateMachine_EventOccurred(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::eventOccurred"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlStateMachine) ConnectEventOccurred(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectEventOccurred(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::eventOccurred", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEventOccurred() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectEventOccurred(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::eventOccurred")
	}
}

func (ptr *QScxmlStateMachine) EventOccurred(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_EventOccurred(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlStateMachine_ExternalEventOccurred
func callbackQScxmlStateMachine_ExternalEventOccurred(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QScxmlStateMachine::externalEventOccurred"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlStateMachine) ConnectExternalEventOccurred(f func(event *QScxmlEvent)) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectExternalEventOccurred(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::externalEventOccurred", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectExternalEventOccurred() {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectExternalEventOccurred(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QScxmlStateMachine::externalEventOccurred")
	}
}

func (ptr *QScxmlStateMachine) ExternalEventOccurred(event QScxmlEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ExternalEventOccurred(ptr.Pointer(), PointerFromQScxmlEvent(event))
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

func QScxmlStateMachine_GenerateSessionId(prefix string) string {
	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	return cGoUnpackString(C.QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(prefixC))
}

func (ptr *QScxmlStateMachine) GenerateSessionId(prefix string) string {
	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	return cGoUnpackString(C.QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(prefixC))
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

func (ptr *QScxmlStateMachine) IsActive(scxmlStateName string) bool {
	if ptr.Pointer() != nil {
		var scxmlStateNameC = C.CString(scxmlStateName)
		defer C.free(unsafe.Pointer(scxmlStateNameC))
		return C.QScxmlStateMachine_IsActive(ptr.Pointer(), scxmlStateNameC) != 0
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

func (ptr *QScxmlStateMachine) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlStateMachine) ParseErrors() []*QScxmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtScxml_PackedList) []*QScxmlError {
			var out = make([]*QScxmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQScxmlStateMachineFromPointer(l.data).parseErrors_atList(i)
			}
			return out
		}(C.QScxmlStateMachine_ParseErrors(ptr.Pointer()))
	}
	return nil
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

func (ptr *QScxmlStateMachine) ScxmlEventFilter() *QScxmlEventFilter {
	if ptr.Pointer() != nil {
		return NewQScxmlEventFilterFromPointer(C.QScxmlStateMachine_ScxmlEventFilter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) SessionId() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScxmlStateMachine_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlStateMachine) SetDataModel(model QScxmlDataModel_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetDataModel(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

func (ptr *QScxmlStateMachine) SetRunning(running bool) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetRunning(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(running))))
	}
}

func (ptr *QScxmlStateMachine) SetScxmlEventFilter(newFilter QScxmlEventFilter_ITF) {
	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetScxmlEventFilter(ptr.Pointer(), PointerFromQScxmlEventFilter(newFilter))
	}
}

func (ptr *QScxmlStateMachine) SetSessionId(id string) {
	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QScxmlStateMachine_SetSessionId(ptr.Pointer(), idC)
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

func (ptr *QScxmlStateMachine) StateNames(compress bool) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QScxmlStateMachine_StateNames(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(compress))))), "|")
	}
	return make([]string, 0)
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

func (ptr *QScxmlStateMachine) parseErrors_atList(i int) *QScxmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQScxmlErrorFromPointer(C.QScxmlStateMachine_parseErrors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QScxmlError).DestroyQScxmlError)
		return tmpValue
	}
	return nil
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
