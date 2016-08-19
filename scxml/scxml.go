// +build !minimal

package scxml

//#include <stdlib.h>
//#include "scxml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"runtime"
	"strings"
	"unsafe"
)

type QScxmlCppDataModel struct {
	QScxmlDataModel
}

type QScxmlCppDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlCppDataModel_PTR() *QScxmlCppDataModel
}

func (p *QScxmlCppDataModel) QScxmlCppDataModel_PTR() *QScxmlCppDataModel {
	return p
}

func (p *QScxmlCppDataModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlCppDataModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QScxmlDataModel_PTR().SetPointer(ptr)
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

func newQScxmlCppDataModelFromPointer(ptr unsafe.Pointer) *QScxmlCppDataModel {
	var n = NewQScxmlCppDataModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QScxmlCppDataModel_") {
		n.SetObjectName("QScxmlCppDataModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QScxmlCppDataModel) DestroyQScxmlCppDataModel() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QScxmlCppDataModel) In(stateName string) bool {
	defer qt.Recovering("QScxmlCppDataModel::In")

	if ptr.Pointer() != nil {
		var stateNameC = C.CString(stateName)
		defer C.free(unsafe.Pointer(stateNameC))
		return C.QScxmlCppDataModel_In(ptr.Pointer(), stateNameC) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_HasScxmlProperty
func callbackQScxmlCppDataModel_HasScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QScxmlCppDataModel::hasScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).HasScxmlPropertyDefault(C.GoString(name))))
}

func (ptr *QScxmlCppDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	defer qt.Recovering("connect QScxmlCppDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasScxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectHasScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) HasScxmlProperty(name string) bool {
	defer qt.Recovering("QScxmlCppDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlCppDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) HasScxmlPropertyDefault(name string) bool {
	defer qt.Recovering("QScxmlCppDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlCppDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) ScxmlEvent() *QScxmlEvent {
	defer qt.Recovering("QScxmlCppDataModel::scxmlEvent")

	if ptr.Pointer() != nil {
		return NewQScxmlEventFromPointer(C.QScxmlCppDataModel_ScxmlEvent(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlCppDataModel_ScxmlProperty
func callbackQScxmlCppDataModel_ScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlCppDataModel::scxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(C.GoString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlCppDataModelFromPointer(ptr).ScxmlPropertyDefault(C.GoString(name)))
}

func (ptr *QScxmlCppDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	defer qt.Recovering("connect QScxmlCppDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) ScxmlProperty(name string) *core.QVariant {
	defer qt.Recovering("QScxmlCppDataModel::scxmlProperty")

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
	defer qt.Recovering("QScxmlCppDataModel::scxmlProperty")

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
	defer qt.Recovering("QScxmlCppDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlCppDataModel_SetScxmlProperty
func callbackQScxmlCppDataModel_SetScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer, context *C.char) C.int {
	defer qt.Recovering("callback QScxmlCppDataModel::setScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).SetScxmlPropertyDefault(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
}

func (ptr *QScxmlCppDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	defer qt.Recovering("connect QScxmlCppDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlProperty", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectSetScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlProperty")
	}
}

func (ptr *QScxmlCppDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	defer qt.Recovering("QScxmlCppDataModel::setScxmlProperty")

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
	defer qt.Recovering("QScxmlCppDataModel::setScxmlProperty")

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
func callbackQScxmlCppDataModel_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlCppDataModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScxmlCppDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScxmlCppDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlCppDataModel_ChildEvent
func callbackQScxmlCppDataModel_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlCppDataModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScxmlCppDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScxmlCppDataModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlCppDataModel_ConnectNotify
func callbackQScxmlCppDataModel_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlCppDataModel::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlCppDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScxmlCppDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlCppDataModel_CustomEvent
func callbackQScxmlCppDataModel_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlCppDataModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlCppDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScxmlCppDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScxmlCppDataModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlCppDataModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlCppDataModel_DeleteLater
func callbackQScxmlCppDataModel_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlCppDataModel::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlCppDataModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScxmlCppDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScxmlCppDataModel) DeleteLater() {
	defer qt.Recovering("QScxmlCppDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlCppDataModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlCppDataModel) DeleteLaterDefault() {
	defer qt.Recovering("QScxmlCppDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlCppDataModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlCppDataModel_DisconnectNotify
func callbackQScxmlCppDataModel_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlCppDataModel::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlCppDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlCppDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlCppDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScxmlCppDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlCppDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlCppDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlCppDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlCppDataModel_Event
func callbackQScxmlCppDataModel_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlCppDataModel::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScxmlCppDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlCppDataModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScxmlCppDataModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlCppDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlCppDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_EventFilter
func callbackQScxmlCppDataModel_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlCppDataModel::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlCppDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScxmlCppDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlCppDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScxmlCppDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlCppDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlCppDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlCppDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlCppDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlCppDataModel_MetaObject
func callbackQScxmlCppDataModel_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlCppDataModel::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlCppDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlCppDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScxmlCppDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScxmlCppDataModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScxmlCppDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScxmlCppDataModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScxmlCppDataModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlCppDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlCppDataModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScxmlCppDataModel::metaObject")

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

func (p *QScxmlDataModel) QScxmlDataModel_PTR() *QScxmlDataModel {
	return p
}

func (p *QScxmlDataModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlDataModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQScxmlDataModelFromPointer(ptr unsafe.Pointer) *QScxmlDataModel {
	var n = NewQScxmlDataModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QScxmlDataModel_") {
		n.SetObjectName("QScxmlDataModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QScxmlDataModel) DestroyQScxmlDataModel() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQScxmlDataModel_HasScxmlProperty
func callbackQScxmlDataModel_HasScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QScxmlDataModel::hasScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QScxmlDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	defer qt.Recovering("connect QScxmlDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasScxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectHasScxmlProperty(name string) {
	defer qt.Recovering("disconnect QScxmlDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) HasScxmlProperty(name string) bool {
	defer qt.Recovering("QScxmlDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

//export callbackQScxmlDataModel_ScxmlProperty
func callbackQScxmlDataModel_ScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlDataModel::scxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(C.GoString(name)))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QScxmlDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	defer qt.Recovering("connect QScxmlDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectScxmlProperty(name string) {
	defer qt.Recovering("disconnect QScxmlDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scxmlProperty")
	}
}

func (ptr *QScxmlDataModel) ScxmlProperty(name string) *core.QVariant {
	defer qt.Recovering("QScxmlDataModel::scxmlProperty")

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
func callbackQScxmlDataModel_SetScxmlEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::setScxmlEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	defer qt.Recovering("connect QScxmlDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlEvent(event QScxmlEvent_ITF) {
	defer qt.Recovering("disconnect QScxmlDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlEvent")
	}
}

func (ptr *QScxmlDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlDataModel_SetScxmlProperty
func callbackQScxmlDataModel_SetScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer, context *C.char) C.int {
	defer qt.Recovering("callback QScxmlDataModel::setScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QScxmlDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	defer qt.Recovering("connect QScxmlDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlProperty", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectSetScxmlProperty(name string, value core.QVariant_ITF, context string) {
	defer qt.Recovering("disconnect QScxmlDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlProperty")
	}
}

func (ptr *QScxmlDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	defer qt.Recovering("QScxmlDataModel::setScxmlProperty")

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
	defer qt.Recovering("QScxmlDataModel::setStateMachine")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_SetStateMachine(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

func (ptr *QScxmlDataModel) StateMachine() *QScxmlStateMachine {
	defer qt.Recovering("QScxmlDataModel::stateMachine")

	if ptr.Pointer() != nil {
		return NewQScxmlStateMachineFromPointer(C.QScxmlDataModel_StateMachine(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlDataModel_StateMachineChanged
func callbackQScxmlDataModel_StateMachineChanged(ptr unsafe.Pointer, ptrName *C.char, stateMachine unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::stateMachineChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateMachineChanged"); signal != nil {
		signal.(func(*QScxmlStateMachine))(NewQScxmlStateMachineFromPointer(stateMachine))
	}

}

func (ptr *QScxmlDataModel) ConnectStateMachineChanged(f func(stateMachine *QScxmlStateMachine)) {
	defer qt.Recovering("connect QScxmlDataModel::stateMachineChanged")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectStateMachineChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateMachineChanged", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectStateMachineChanged() {
	defer qt.Recovering("disconnect QScxmlDataModel::stateMachineChanged")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectStateMachineChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateMachineChanged")
	}
}

func (ptr *QScxmlDataModel) StateMachineChanged(stateMachine QScxmlStateMachine_ITF) {
	defer qt.Recovering("QScxmlDataModel::stateMachineChanged")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_StateMachineChanged(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

//export callbackQScxmlDataModel_TimerEvent
func callbackQScxmlDataModel_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScxmlDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScxmlDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScxmlDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlDataModel_ChildEvent
func callbackQScxmlDataModel_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScxmlDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScxmlDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScxmlDataModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlDataModel_ConnectNotify
func callbackQScxmlDataModel_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScxmlDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScxmlDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_CustomEvent
func callbackQScxmlDataModel_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScxmlDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScxmlDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScxmlDataModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlDataModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlDataModel_DeleteLater
func callbackQScxmlDataModel_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlDataModel::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlDataModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScxmlDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScxmlDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScxmlDataModel) DeleteLater() {
	defer qt.Recovering("QScxmlDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlDataModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlDataModel) DeleteLaterDefault() {
	defer qt.Recovering("QScxmlDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlDataModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlDataModel_DisconnectNotify
func callbackQScxmlDataModel_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlDataModel::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScxmlDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScxmlDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlDataModel_Event
func callbackQScxmlDataModel_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlDataModel::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScxmlDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlDataModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QScxmlDataModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScxmlDataModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_EventFilter
func callbackQScxmlDataModel_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlDataModel::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScxmlDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScxmlDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScxmlDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlDataModel_MetaObject
func callbackQScxmlDataModel_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlDataModel::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScxmlDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScxmlDataModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScxmlDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScxmlDataModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScxmlDataModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlDataModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScxmlDataModel::metaObject")

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

func (p *QScxmlEcmaScriptDataModel) QScxmlEcmaScriptDataModel_PTR() *QScxmlEcmaScriptDataModel {
	return p
}

func (p *QScxmlEcmaScriptDataModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlEcmaScriptDataModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QScxmlDataModel_PTR().SetPointer(ptr)
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

func newQScxmlEcmaScriptDataModelFromPointer(ptr unsafe.Pointer) *QScxmlEcmaScriptDataModel {
	var n = NewQScxmlEcmaScriptDataModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QScxmlEcmaScriptDataModel_") {
		n.SetObjectName("QScxmlEcmaScriptDataModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QScxmlEcmaScriptDataModel) DestroyQScxmlEcmaScriptDataModel() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQScxmlEcmaScriptDataModel(parent core.QObject_ITF) *QScxmlEcmaScriptDataModel {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::QScxmlEcmaScriptDataModel")

	return newQScxmlEcmaScriptDataModelFromPointer(C.QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(core.PointerFromQObject(parent)))
}

func (ptr *QScxmlEcmaScriptDataModel) Engine() *qml.QJSEngine {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::engine")

	if ptr.Pointer() != nil {
		return qml.NewQJSEngineFromPointer(C.QScxmlEcmaScriptDataModel_Engine(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlEcmaScriptDataModel_HasScxmlProperty
func callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::hasScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).HasScxmlPropertyDefault(C.GoString(name))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasScxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectHasScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlProperty(name string) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlEcmaScriptDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) HasScxmlPropertyDefault(name string) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_ScxmlProperty
func callbackQScxmlEcmaScriptDataModel_ScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::scxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(C.GoString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlEcmaScriptDataModelFromPointer(ptr).ScxmlPropertyDefault(C.GoString(name)))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ScxmlProperty(name string) *core.QVariant {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::scxmlProperty")

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
	defer qt.Recovering("QScxmlEcmaScriptDataModel::scxmlProperty")

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
	defer qt.Recovering("QScxmlEcmaScriptDataModel::setEngine")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetEngine(ptr.Pointer(), qml.PointerFromQJSEngine(engine))
	}
}

//export callbackQScxmlEcmaScriptDataModel_SetScxmlEvent
func callbackQScxmlEcmaScriptDataModel_SetScxmlEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::setScxmlEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlEvent() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_SetScxmlEventDefault(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_SetScxmlProperty
func callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer, context *C.char) C.int {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::setScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).SetScxmlPropertyDefault(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlProperty", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectSetScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlProperty")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::setScxmlProperty")

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
	defer qt.Recovering("QScxmlEcmaScriptDataModel::setScxmlProperty")

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
func callbackQScxmlEcmaScriptDataModel_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_ChildEvent
func callbackQScxmlEcmaScriptDataModel_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_ConnectNotify
func callbackQScxmlEcmaScriptDataModel_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlEcmaScriptDataModel_CustomEvent
func callbackQScxmlEcmaScriptDataModel_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlEcmaScriptDataModel_DeleteLater
func callbackQScxmlEcmaScriptDataModel_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DeleteLater() {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlEcmaScriptDataModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DeleteLaterDefault() {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlEcmaScriptDataModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlEcmaScriptDataModel_DisconnectNotify
func callbackQScxmlEcmaScriptDataModel_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlEcmaScriptDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlEcmaScriptDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlEcmaScriptDataModel_Event
func callbackQScxmlEcmaScriptDataModel_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_EventFilter
func callbackQScxmlEcmaScriptDataModel_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlEcmaScriptDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlEcmaScriptDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlEcmaScriptDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlEcmaScriptDataModel_MetaObject
func callbackQScxmlEcmaScriptDataModel_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlEcmaScriptDataModel::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlEcmaScriptDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlEcmaScriptDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScxmlEcmaScriptDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScxmlEcmaScriptDataModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScxmlEcmaScriptDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScxmlEcmaScriptDataModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlEcmaScriptDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlEcmaScriptDataModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScxmlEcmaScriptDataModel::metaObject")

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

func (p *QScxmlError) QScxmlError_PTR() *QScxmlError {
	return p
}

func (p *QScxmlError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScxmlError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQScxmlErrorFromPointer(ptr unsafe.Pointer) *QScxmlError {
	var n = NewQScxmlErrorFromPointer(ptr)
	return n
}

func NewQScxmlError() *QScxmlError {
	defer qt.Recovering("QScxmlError::QScxmlError")

	return newQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError())
}

func NewQScxmlError3(other QScxmlError_ITF) *QScxmlError {
	defer qt.Recovering("QScxmlError::QScxmlError")

	return newQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError3(PointerFromQScxmlError(other)))
}

func NewQScxmlError2(fileName string, line int, column int, description string) *QScxmlError {
	defer qt.Recovering("QScxmlError::QScxmlError")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var descriptionC = C.CString(description)
	defer C.free(unsafe.Pointer(descriptionC))
	return newQScxmlErrorFromPointer(C.QScxmlError_NewQScxmlError2(fileNameC, C.int(line), C.int(column), descriptionC))
}

func (ptr *QScxmlError) Column() int {
	defer qt.Recovering("QScxmlError::column")

	if ptr.Pointer() != nil {
		return int(C.QScxmlError_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlError) Description() string {
	defer qt.Recovering("QScxmlError::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) FileName() string {
	defer qt.Recovering("QScxmlError::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlError_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) IsValid() bool {
	defer qt.Recovering("QScxmlError::isValid")

	if ptr.Pointer() != nil {
		return C.QScxmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlError) Line() int {
	defer qt.Recovering("QScxmlError::line")

	if ptr.Pointer() != nil {
		return int(C.QScxmlError_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlError) ToString() string {
	defer qt.Recovering("QScxmlError::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlError) DestroyQScxmlError() {
	defer qt.Recovering("QScxmlError::~QScxmlError")

	if ptr.Pointer() != nil {
		C.QScxmlError_DestroyQScxmlError(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScxmlEvent::EventType
type QScxmlEvent__EventType int64

const (
	QScxmlEvent__PlatformEvent = QScxmlEvent__EventType(0)
	QScxmlEvent__InternalEvent = QScxmlEvent__EventType(1)
	QScxmlEvent__ExternalEvent = QScxmlEvent__EventType(2)
)

type QScxmlEvent struct {
	core.QEvent
}

type QScxmlEvent_ITF interface {
	core.QEvent_ITF
	QScxmlEvent_PTR() *QScxmlEvent
}

func (p *QScxmlEvent) QScxmlEvent_PTR() *QScxmlEvent {
	return p
}

func (p *QScxmlEvent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QEvent_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlEvent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QEvent_PTR().SetPointer(ptr)
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

func newQScxmlEventFromPointer(ptr unsafe.Pointer) *QScxmlEvent {
	var n = NewQScxmlEventFromPointer(ptr)
	return n
}

func NewQScxmlEvent() *QScxmlEvent {
	defer qt.Recovering("QScxmlEvent::QScxmlEvent")

	return newQScxmlEventFromPointer(C.QScxmlEvent_NewQScxmlEvent())
}

func NewQScxmlEvent2(other QScxmlEvent_ITF) *QScxmlEvent {
	defer qt.Recovering("QScxmlEvent::QScxmlEvent")

	return newQScxmlEventFromPointer(C.QScxmlEvent_NewQScxmlEvent2(PointerFromQScxmlEvent(other)))
}

func (ptr *QScxmlEvent) Clear() {
	defer qt.Recovering("QScxmlEvent::clear")

	if ptr.Pointer() != nil {
		C.QScxmlEvent_Clear(ptr.Pointer())
	}
}

func (ptr *QScxmlEvent) Data() *core.QVariant {
	defer qt.Recovering("QScxmlEvent::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QScxmlEvent_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QScxmlEvent) Delay() int {
	defer qt.Recovering("QScxmlEvent::delay")

	if ptr.Pointer() != nil {
		return int(C.QScxmlEvent_Delay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlEvent) ErrorMessage() string {
	defer qt.Recovering("QScxmlEvent::errorMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) EventType() QScxmlEvent__EventType {
	defer qt.Recovering("QScxmlEvent::eventType")

	if ptr.Pointer() != nil {
		return QScxmlEvent__EventType(C.QScxmlEvent_EventType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlEvent) InvokeId() string {
	defer qt.Recovering("QScxmlEvent::invokeId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_InvokeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) IsErrorEvent() bool {
	defer qt.Recovering("QScxmlEvent::isErrorEvent")

	if ptr.Pointer() != nil {
		return C.QScxmlEvent_IsErrorEvent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlEvent) Name() string {
	defer qt.Recovering("QScxmlEvent::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) Origin() string {
	defer qt.Recovering("QScxmlEvent::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) OriginType() string {
	defer qt.Recovering("QScxmlEvent::originType")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_OriginType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) ScxmlType() string {
	defer qt.Recovering("QScxmlEvent::scxmlType")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_ScxmlType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) SendId() string {
	defer qt.Recovering("QScxmlEvent::sendId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEvent_SendId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEvent) SetData(data core.QVariant_ITF) {
	defer qt.Recovering("QScxmlEvent::setData")

	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetData(ptr.Pointer(), core.PointerFromQVariant(data))
	}
}

func (ptr *QScxmlEvent) SetDelay(delayInMiliSecs int) {
	defer qt.Recovering("QScxmlEvent::setDelay")

	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetDelay(ptr.Pointer(), C.int(delayInMiliSecs))
	}
}

func (ptr *QScxmlEvent) SetErrorMessage(message string) {
	defer qt.Recovering("QScxmlEvent::setErrorMessage")

	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		C.QScxmlEvent_SetErrorMessage(ptr.Pointer(), messageC)
	}
}

func (ptr *QScxmlEvent) SetEventType(ty QScxmlEvent__EventType) {
	defer qt.Recovering("QScxmlEvent::setEventType")

	if ptr.Pointer() != nil {
		C.QScxmlEvent_SetEventType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QScxmlEvent) SetInvokeId(invokeid string) {
	defer qt.Recovering("QScxmlEvent::setInvokeId")

	if ptr.Pointer() != nil {
		var invokeidC = C.CString(invokeid)
		defer C.free(unsafe.Pointer(invokeidC))
		C.QScxmlEvent_SetInvokeId(ptr.Pointer(), invokeidC)
	}
}

func (ptr *QScxmlEvent) SetName(name string) {
	defer qt.Recovering("QScxmlEvent::setName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QScxmlEvent_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QScxmlEvent) SetOrigin(origin string) {
	defer qt.Recovering("QScxmlEvent::setOrigin")

	if ptr.Pointer() != nil {
		var originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
		C.QScxmlEvent_SetOrigin(ptr.Pointer(), originC)
	}
}

func (ptr *QScxmlEvent) SetOriginType(origintype string) {
	defer qt.Recovering("QScxmlEvent::setOriginType")

	if ptr.Pointer() != nil {
		var origintypeC = C.CString(origintype)
		defer C.free(unsafe.Pointer(origintypeC))
		C.QScxmlEvent_SetOriginType(ptr.Pointer(), origintypeC)
	}
}

func (ptr *QScxmlEvent) SetSendId(sendid string) {
	defer qt.Recovering("QScxmlEvent::setSendId")

	if ptr.Pointer() != nil {
		var sendidC = C.CString(sendid)
		defer C.free(unsafe.Pointer(sendidC))
		C.QScxmlEvent_SetSendId(ptr.Pointer(), sendidC)
	}
}

func (ptr *QScxmlEvent) DestroyQScxmlEvent() {
	defer qt.Recovering("QScxmlEvent::~QScxmlEvent")

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

func (p *QScxmlEventFilter) QScxmlEventFilter_PTR() *QScxmlEventFilter {
	return p
}

func (p *QScxmlEventFilter) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScxmlEventFilter) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQScxmlEventFilterFromPointer(ptr unsafe.Pointer) *QScxmlEventFilter {
	var n = NewQScxmlEventFilterFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QScxmlEventFilter_") {
		n.SetObjectNameAbs("QScxmlEventFilter_" + qt.Identifier())
	}
	return n
}

//export callbackQScxmlEventFilter_Handle
func callbackQScxmlEventFilter_Handle(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer, stateMachine unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlEventFilter::handle")

	if signal := qt.GetSignal(C.GoString(ptrName), "handle"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QScxmlEvent, *QScxmlStateMachine) bool)(NewQScxmlEventFromPointer(event), NewQScxmlStateMachineFromPointer(stateMachine))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QScxmlEventFilter) ConnectHandle(f func(event *QScxmlEvent, stateMachine *QScxmlStateMachine) bool) {
	defer qt.Recovering("connect QScxmlEventFilter::handle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "handle", f)
	}
}

func (ptr *QScxmlEventFilter) DisconnectHandle(event QScxmlEvent_ITF, stateMachine QScxmlStateMachine_ITF) {
	defer qt.Recovering("disconnect QScxmlEventFilter::handle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "handle")
	}
}

func (ptr *QScxmlEventFilter) Handle(event QScxmlEvent_ITF, stateMachine QScxmlStateMachine_ITF) bool {
	defer qt.Recovering("QScxmlEventFilter::handle")

	if ptr.Pointer() != nil {
		return C.QScxmlEventFilter_Handle(ptr.Pointer(), PointerFromQScxmlEvent(event), PointerFromQScxmlStateMachine(stateMachine)) != 0
	}
	return false
}

func (ptr *QScxmlEventFilter) DestroyQScxmlEventFilter() {
	defer qt.Recovering("QScxmlEventFilter::~QScxmlEventFilter")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QScxmlEventFilter_DestroyQScxmlEventFilter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlEventFilter) ObjectNameAbs() string {
	defer qt.Recovering("QScxmlEventFilter::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlEventFilter_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlEventFilter) SetObjectNameAbs(name string) {
	defer qt.Recovering("QScxmlEventFilter::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QScxmlEventFilter_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

type QScxmlNullDataModel struct {
	QScxmlDataModel
}

type QScxmlNullDataModel_ITF interface {
	QScxmlDataModel_ITF
	QScxmlNullDataModel_PTR() *QScxmlNullDataModel
}

func (p *QScxmlNullDataModel) QScxmlNullDataModel_PTR() *QScxmlNullDataModel {
	return p
}

func (p *QScxmlNullDataModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QScxmlDataModel_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlNullDataModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QScxmlDataModel_PTR().SetPointer(ptr)
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

func newQScxmlNullDataModelFromPointer(ptr unsafe.Pointer) *QScxmlNullDataModel {
	var n = NewQScxmlNullDataModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QScxmlNullDataModel_") {
		n.SetObjectName("QScxmlNullDataModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QScxmlNullDataModel) DestroyQScxmlNullDataModel() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQScxmlNullDataModel(parent core.QObject_ITF) *QScxmlNullDataModel {
	defer qt.Recovering("QScxmlNullDataModel::QScxmlNullDataModel")

	return newQScxmlNullDataModelFromPointer(C.QScxmlNullDataModel_NewQScxmlNullDataModel(core.PointerFromQObject(parent)))
}

//export callbackQScxmlNullDataModel_HasScxmlProperty
func callbackQScxmlNullDataModel_HasScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) C.int {
	defer qt.Recovering("callback QScxmlNullDataModel::hasScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).HasScxmlPropertyDefault(C.GoString(name))))
}

func (ptr *QScxmlNullDataModel) ConnectHasScxmlProperty(f func(name string) bool) {
	defer qt.Recovering("connect QScxmlNullDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasScxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectHasScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) HasScxmlProperty(name string) bool {
	defer qt.Recovering("QScxmlNullDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlNullDataModel_HasScxmlProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) HasScxmlPropertyDefault(name string) bool {
	defer qt.Recovering("QScxmlNullDataModel::hasScxmlProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QScxmlNullDataModel_HasScxmlPropertyDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_ScxmlProperty
func callbackQScxmlNullDataModel_ScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlNullDataModel::scxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "scxmlProperty"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(C.GoString(name)))
	}

	return core.PointerFromQVariant(NewQScxmlNullDataModelFromPointer(ptr).ScxmlPropertyDefault(C.GoString(name)))
}

func (ptr *QScxmlNullDataModel) ConnectScxmlProperty(f func(name string) *core.QVariant) {
	defer qt.Recovering("connect QScxmlNullDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::scxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) ScxmlProperty(name string) *core.QVariant {
	defer qt.Recovering("QScxmlNullDataModel::scxmlProperty")

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
	defer qt.Recovering("QScxmlNullDataModel::scxmlProperty")

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
func callbackQScxmlNullDataModel_SetScxmlEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::setScxmlEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlEvent"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).SetScxmlEventDefault(NewQScxmlEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlEvent(f func(event *QScxmlEvent)) {
	defer qt.Recovering("connect QScxmlNullDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlEvent() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlEvent")
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlEvent(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_SetScxmlEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlEventDefault(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::setScxmlEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_SetScxmlEventDefault(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlNullDataModel_SetScxmlProperty
func callbackQScxmlNullDataModel_SetScxmlProperty(ptr unsafe.Pointer, ptrName *C.char, name *C.char, value unsafe.Pointer, context *C.char) C.int {
	defer qt.Recovering("callback QScxmlNullDataModel::setScxmlProperty")

	if signal := qt.GetSignal(C.GoString(ptrName), "setScxmlProperty"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, *core.QVariant, string) bool)(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).SetScxmlPropertyDefault(C.GoString(name), core.NewQVariantFromPointer(value), C.GoString(context))))
}

func (ptr *QScxmlNullDataModel) ConnectSetScxmlProperty(f func(name string, value *core.QVariant, context string) bool) {
	defer qt.Recovering("connect QScxmlNullDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setScxmlProperty", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectSetScxmlProperty() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::setScxmlProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setScxmlProperty")
	}
}

func (ptr *QScxmlNullDataModel) SetScxmlProperty(name string, value core.QVariant_ITF, context string) bool {
	defer qt.Recovering("QScxmlNullDataModel::setScxmlProperty")

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
	defer qt.Recovering("QScxmlNullDataModel::setScxmlProperty")

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
func callbackQScxmlNullDataModel_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScxmlNullDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScxmlNullDataModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlNullDataModel_ChildEvent
func callbackQScxmlNullDataModel_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScxmlNullDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScxmlNullDataModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlNullDataModel_ConnectNotify
func callbackQScxmlNullDataModel_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlNullDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScxmlNullDataModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlNullDataModel_CustomEvent
func callbackQScxmlNullDataModel_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlNullDataModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScxmlNullDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScxmlNullDataModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlNullDataModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlNullDataModel_DeleteLater
func callbackQScxmlNullDataModel_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlNullDataModel::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlNullDataModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScxmlNullDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScxmlNullDataModel) DeleteLater() {
	defer qt.Recovering("QScxmlNullDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlNullDataModel_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlNullDataModel) DeleteLaterDefault() {
	defer qt.Recovering("QScxmlNullDataModel::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlNullDataModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlNullDataModel_DisconnectNotify
func callbackQScxmlNullDataModel_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlNullDataModel::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlNullDataModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlNullDataModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlNullDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScxmlNullDataModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlNullDataModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlNullDataModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlNullDataModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlNullDataModel_Event
func callbackQScxmlNullDataModel_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlNullDataModel::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScxmlNullDataModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlNullDataModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScxmlNullDataModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlNullDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlNullDataModel::event")

	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_EventFilter
func callbackQScxmlNullDataModel_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlNullDataModel::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlNullDataModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScxmlNullDataModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlNullDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScxmlNullDataModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlNullDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlNullDataModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlNullDataModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlNullDataModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlNullDataModel_MetaObject
func callbackQScxmlNullDataModel_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlNullDataModel::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlNullDataModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlNullDataModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScxmlNullDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScxmlNullDataModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScxmlNullDataModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScxmlNullDataModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScxmlNullDataModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlNullDataModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlNullDataModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScxmlNullDataModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlNullDataModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QScxmlParser::QtMode
type QScxmlParser__QtMode int64

const (
	QScxmlParser__QtModeDisabled      = QScxmlParser__QtMode(0)
	QScxmlParser__QtModeEnabled       = QScxmlParser__QtMode(1)
	QScxmlParser__QtModeFromInputFile = QScxmlParser__QtMode(2)
)

type QScxmlParser struct {
	ptr unsafe.Pointer
}

type QScxmlParser_ITF interface {
	QScxmlParser_PTR() *QScxmlParser
}

func (p *QScxmlParser) QScxmlParser_PTR() *QScxmlParser {
	return p
}

func (p *QScxmlParser) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QScxmlParser) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQScxmlParserFromPointer(ptr unsafe.Pointer) *QScxmlParser {
	var n = NewQScxmlParserFromPointer(ptr)
	return n
}

func NewQScxmlParser(reader core.QXmlStreamReader_ITF) *QScxmlParser {
	defer qt.Recovering("QScxmlParser::QScxmlParser")

	return newQScxmlParserFromPointer(C.QScxmlParser_NewQScxmlParser(core.PointerFromQXmlStreamReader(reader)))
}

func (ptr *QScxmlParser) AddError(msg string) {
	defer qt.Recovering("QScxmlParser::addError")

	if ptr.Pointer() != nil {
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QScxmlParser_AddError(ptr.Pointer(), msgC)
	}
}

func (ptr *QScxmlParser) FileName() string {
	defer qt.Recovering("QScxmlParser::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlParser_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlParser) InstantiateDataModel(stateMachine QScxmlStateMachine_ITF) {
	defer qt.Recovering("QScxmlParser::instantiateDataModel")

	if ptr.Pointer() != nil {
		C.QScxmlParser_InstantiateDataModel(ptr.Pointer(), PointerFromQScxmlStateMachine(stateMachine))
	}
}

func (ptr *QScxmlParser) InstantiateStateMachine() *QScxmlStateMachine {
	defer qt.Recovering("QScxmlParser::instantiateStateMachine")

	if ptr.Pointer() != nil {
		return NewQScxmlStateMachineFromPointer(C.QScxmlParser_InstantiateStateMachine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlParser) Parse() {
	defer qt.Recovering("QScxmlParser::parse")

	if ptr.Pointer() != nil {
		C.QScxmlParser_Parse(ptr.Pointer())
	}
}

func (ptr *QScxmlParser) QtMode() QScxmlParser__QtMode {
	defer qt.Recovering("QScxmlParser::qtMode")

	if ptr.Pointer() != nil {
		return QScxmlParser__QtMode(C.QScxmlParser_QtMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlParser) SetFileName(fileName string) {
	defer qt.Recovering("QScxmlParser::setFileName")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QScxmlParser_SetFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QScxmlParser) SetQtMode(mode QScxmlParser__QtMode) {
	defer qt.Recovering("QScxmlParser::setQtMode")

	if ptr.Pointer() != nil {
		C.QScxmlParser_SetQtMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QScxmlParser) DestroyQScxmlParser() {
	defer qt.Recovering("QScxmlParser::~QScxmlParser")

	if ptr.Pointer() != nil {
		C.QScxmlParser_DestroyQScxmlParser(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QScxmlStateMachine::BindingMethod
type QScxmlStateMachine__BindingMethod int64

const (
	QScxmlStateMachine__EarlyBinding = QScxmlStateMachine__BindingMethod(0)
	QScxmlStateMachine__LateBinding  = QScxmlStateMachine__BindingMethod(1)
)

type QScxmlStateMachine struct {
	core.QObject
}

type QScxmlStateMachine_ITF interface {
	core.QObject_ITF
	QScxmlStateMachine_PTR() *QScxmlStateMachine
}

func (p *QScxmlStateMachine) QScxmlStateMachine_PTR() *QScxmlStateMachine {
	return p
}

func (p *QScxmlStateMachine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QScxmlStateMachine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQScxmlStateMachineFromPointer(ptr unsafe.Pointer) *QScxmlStateMachine {
	var n = NewQScxmlStateMachineFromPointer(ptr)
	for len(n.ObjectName()) < len("QScxmlStateMachine_") {
		n.SetObjectName("QScxmlStateMachine_" + qt.Identifier())
	}
	return n
}

func (ptr *QScxmlStateMachine) DestroyQScxmlStateMachine() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QScxmlStateMachine) IsInitialized() bool {
	defer qt.Recovering("QScxmlStateMachine::isInitialized")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) ActiveStateNames(compress bool) []string {
	defer qt.Recovering("QScxmlStateMachine::activeStateNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScxmlStateMachine_ActiveStateNames(ptr.Pointer(), C.int(qt.GoBoolToInt(compress)))), "|")
	}
	return make([]string, 0)
}

func (ptr *QScxmlStateMachine) CancelDelayedEvent(sendId string) {
	defer qt.Recovering("QScxmlStateMachine::cancelDelayedEvent")

	if ptr.Pointer() != nil {
		var sendIdC = C.CString(sendId)
		defer C.free(unsafe.Pointer(sendIdC))
		C.QScxmlStateMachine_CancelDelayedEvent(ptr.Pointer(), sendIdC)
	}
}

func (ptr *QScxmlStateMachine) DataBinding() QScxmlStateMachine__BindingMethod {
	defer qt.Recovering("QScxmlStateMachine::dataBinding")

	if ptr.Pointer() != nil {
		return QScxmlStateMachine__BindingMethod(C.QScxmlStateMachine_DataBinding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScxmlStateMachine) DataModel() *QScxmlDataModel {
	defer qt.Recovering("QScxmlStateMachine::dataModel")

	if ptr.Pointer() != nil {
		return NewQScxmlDataModelFromPointer(C.QScxmlStateMachine_DataModel(ptr.Pointer()))
	}
	return nil
}

//export callbackQScxmlStateMachine_DataModelChanged
func callbackQScxmlStateMachine_DataModelChanged(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::dataModelChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "dataModelChanged"); signal != nil {
		signal.(func(*QScxmlDataModel))(NewQScxmlDataModelFromPointer(model))
	}

}

func (ptr *QScxmlStateMachine) ConnectDataModelChanged(f func(model *QScxmlDataModel)) {
	defer qt.Recovering("connect QScxmlStateMachine::dataModelChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectDataModelChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataModelChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDataModelChanged() {
	defer qt.Recovering("disconnect QScxmlStateMachine::dataModelChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectDataModelChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataModelChanged")
	}
}

func (ptr *QScxmlStateMachine) DataModelChanged(model QScxmlDataModel_ITF) {
	defer qt.Recovering("QScxmlStateMachine::dataModelChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DataModelChanged(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

//export callbackQScxmlStateMachine_EventOccurred
func callbackQScxmlStateMachine_EventOccurred(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::eventOccurred")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventOccurred"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlStateMachine) ConnectEventOccurred(f func(event *QScxmlEvent)) {
	defer qt.Recovering("connect QScxmlStateMachine::eventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectEventOccurred(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "eventOccurred", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEventOccurred() {
	defer qt.Recovering("disconnect QScxmlStateMachine::eventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectEventOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "eventOccurred")
	}
}

func (ptr *QScxmlStateMachine) EventOccurred(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::eventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_EventOccurred(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlStateMachine_ExternalEventOccurred
func callbackQScxmlStateMachine_ExternalEventOccurred(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::externalEventOccurred")

	if signal := qt.GetSignal(C.GoString(ptrName), "externalEventOccurred"); signal != nil {
		signal.(func(*QScxmlEvent))(NewQScxmlEventFromPointer(event))
	}

}

func (ptr *QScxmlStateMachine) ConnectExternalEventOccurred(f func(event *QScxmlEvent)) {
	defer qt.Recovering("connect QScxmlStateMachine::externalEventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectExternalEventOccurred(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "externalEventOccurred", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectExternalEventOccurred() {
	defer qt.Recovering("disconnect QScxmlStateMachine::externalEventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectExternalEventOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "externalEventOccurred")
	}
}

func (ptr *QScxmlStateMachine) ExternalEventOccurred(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::externalEventOccurred")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ExternalEventOccurred(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

//export callbackQScxmlStateMachine_Finished
func callbackQScxmlStateMachine_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectFinished(f func()) {
	defer qt.Recovering("connect QScxmlStateMachine::finished")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectFinished() {
	defer qt.Recovering("disconnect QScxmlStateMachine::finished")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QScxmlStateMachine) Finished() {
	defer qt.Recovering("QScxmlStateMachine::finished")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Finished(ptr.Pointer())
	}
}

func QScxmlStateMachine_FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	defer qt.Recovering("QScxmlStateMachine::fromData")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	return NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), fileNameC))
}

func (ptr *QScxmlStateMachine) FromData(data core.QIODevice_ITF, fileName string) *QScxmlStateMachine {
	defer qt.Recovering("QScxmlStateMachine::fromData")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	return NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromData(core.PointerFromQIODevice(data), fileNameC))
}

func QScxmlStateMachine_FromFile(fileName string) *QScxmlStateMachine {
	defer qt.Recovering("QScxmlStateMachine::fromFile")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	return NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(fileNameC))
}

func (ptr *QScxmlStateMachine) FromFile(fileName string) *QScxmlStateMachine {
	defer qt.Recovering("QScxmlStateMachine::fromFile")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	return NewQScxmlStateMachineFromPointer(C.QScxmlStateMachine_QScxmlStateMachine_FromFile(fileNameC))
}

func QScxmlStateMachine_GenerateSessionId(prefix string) string {
	defer qt.Recovering("QScxmlStateMachine::generateSessionId")

	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	return C.GoString(C.QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(prefixC))
}

func (ptr *QScxmlStateMachine) GenerateSessionId(prefix string) string {
	defer qt.Recovering("QScxmlStateMachine::generateSessionId")

	var prefixC = C.CString(prefix)
	defer C.free(unsafe.Pointer(prefixC))
	return C.GoString(C.QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(prefixC))
}

//export callbackQScxmlStateMachine_Init
func callbackQScxmlStateMachine_Init(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QScxmlStateMachine::init")

	if signal := qt.GetSignal(C.GoString(ptrName), "init"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QScxmlStateMachine) ConnectInit(f func() bool) {
	defer qt.Recovering("connect QScxmlStateMachine::init")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "init", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectInit() {
	defer qt.Recovering("disconnect QScxmlStateMachine::init")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "init")
	}
}

func (ptr *QScxmlStateMachine) Init() bool {
	defer qt.Recovering("QScxmlStateMachine::init")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_Init(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_InitializedChanged
func callbackQScxmlStateMachine_InitializedChanged(ptr unsafe.Pointer, ptrName *C.char, initialized C.int) {
	defer qt.Recovering("callback QScxmlStateMachine::initializedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "initializedChanged"); signal != nil {
		signal.(func(bool))(int(initialized) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectInitializedChanged(f func(initialized bool)) {
	defer qt.Recovering("connect QScxmlStateMachine::initializedChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectInitializedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "initializedChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectInitializedChanged() {
	defer qt.Recovering("disconnect QScxmlStateMachine::initializedChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectInitializedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "initializedChanged")
	}
}

func (ptr *QScxmlStateMachine) InitializedChanged(initialized bool) {
	defer qt.Recovering("QScxmlStateMachine::initializedChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_InitializedChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(initialized)))
	}
}

func (ptr *QScxmlStateMachine) IsActive(scxmlStateName string) bool {
	defer qt.Recovering("QScxmlStateMachine::isActive")

	if ptr.Pointer() != nil {
		var scxmlStateNameC = C.CString(scxmlStateName)
		defer C.free(unsafe.Pointer(scxmlStateNameC))
		return C.QScxmlStateMachine_IsActive(ptr.Pointer(), scxmlStateNameC) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsDispatchableTarget(target string) bool {
	defer qt.Recovering("QScxmlStateMachine::isDispatchableTarget")

	if ptr.Pointer() != nil {
		var targetC = C.CString(target)
		defer C.free(unsafe.Pointer(targetC))
		return C.QScxmlStateMachine_IsDispatchableTarget(ptr.Pointer(), targetC) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsInvoked() bool {
	defer qt.Recovering("QScxmlStateMachine::isInvoked")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsInvoked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) IsRunning() bool {
	defer qt.Recovering("QScxmlStateMachine::isRunning")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_Log
func callbackQScxmlStateMachine_Log(ptr unsafe.Pointer, ptrName *C.char, label *C.char, msg *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::log")

	if signal := qt.GetSignal(C.GoString(ptrName), "log"); signal != nil {
		signal.(func(string, string))(C.GoString(label), C.GoString(msg))
	}

}

func (ptr *QScxmlStateMachine) ConnectLog(f func(label string, msg string)) {
	defer qt.Recovering("connect QScxmlStateMachine::log")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectLog(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "log", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectLog() {
	defer qt.Recovering("disconnect QScxmlStateMachine::log")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectLog(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "log")
	}
}

func (ptr *QScxmlStateMachine) Log(label string, msg string) {
	defer qt.Recovering("QScxmlStateMachine::log")

	if ptr.Pointer() != nil {
		var labelC = C.CString(label)
		defer C.free(unsafe.Pointer(labelC))
		var msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
		C.QScxmlStateMachine_Log(ptr.Pointer(), labelC, msgC)
	}
}

func (ptr *QScxmlStateMachine) Name() string {
	defer qt.Recovering("QScxmlStateMachine::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlStateMachine_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQScxmlStateMachine_ReachedStableState
func callbackQScxmlStateMachine_ReachedStableState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::reachedStableState")

	if signal := qt.GetSignal(C.GoString(ptrName), "reachedStableState"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectReachedStableState(f func()) {
	defer qt.Recovering("connect QScxmlStateMachine::reachedStableState")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectReachedStableState(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "reachedStableState", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectReachedStableState() {
	defer qt.Recovering("disconnect QScxmlStateMachine::reachedStableState")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectReachedStableState(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "reachedStableState")
	}
}

func (ptr *QScxmlStateMachine) ReachedStableState() {
	defer qt.Recovering("QScxmlStateMachine::reachedStableState")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ReachedStableState(ptr.Pointer())
	}
}

//export callbackQScxmlStateMachine_RunningChanged
func callbackQScxmlStateMachine_RunningChanged(ptr unsafe.Pointer, ptrName *C.char, running C.int) {
	defer qt.Recovering("callback QScxmlStateMachine::runningChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "runningChanged"); signal != nil {
		signal.(func(bool))(int(running) != 0)
	}

}

func (ptr *QScxmlStateMachine) ConnectRunningChanged(f func(running bool)) {
	defer qt.Recovering("connect QScxmlStateMachine::runningChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectRunningChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "runningChanged", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectRunningChanged() {
	defer qt.Recovering("disconnect QScxmlStateMachine::runningChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectRunningChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "runningChanged")
	}
}

func (ptr *QScxmlStateMachine) RunningChanged(running bool) {
	defer qt.Recovering("QScxmlStateMachine::runningChanged")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_RunningChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QScxmlStateMachine) ScxmlEventFilter() *QScxmlEventFilter {
	defer qt.Recovering("QScxmlStateMachine::scxmlEventFilter")

	if ptr.Pointer() != nil {
		return NewQScxmlEventFilterFromPointer(C.QScxmlStateMachine_ScxmlEventFilter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) SessionId() string {
	defer qt.Recovering("QScxmlStateMachine::sessionId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScxmlStateMachine_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScxmlStateMachine) SetDataModel(model QScxmlDataModel_ITF) {
	defer qt.Recovering("QScxmlStateMachine::setDataModel")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetDataModel(ptr.Pointer(), PointerFromQScxmlDataModel(model))
	}
}

func (ptr *QScxmlStateMachine) SetRunning(running bool) {
	defer qt.Recovering("QScxmlStateMachine::setRunning")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetRunning(ptr.Pointer(), C.int(qt.GoBoolToInt(running)))
	}
}

func (ptr *QScxmlStateMachine) SetScxmlEventFilter(newFilter QScxmlEventFilter_ITF) {
	defer qt.Recovering("QScxmlStateMachine::setScxmlEventFilter")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SetScxmlEventFilter(ptr.Pointer(), PointerFromQScxmlEventFilter(newFilter))
	}
}

func (ptr *QScxmlStateMachine) SetSessionId(id string) {
	defer qt.Recovering("QScxmlStateMachine::setSessionId")

	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QScxmlStateMachine_SetSessionId(ptr.Pointer(), idC)
	}
}

//export callbackQScxmlStateMachine_Start
func callbackQScxmlStateMachine_Start(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::start")

	if signal := qt.GetSignal(C.GoString(ptrName), "start"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectStart(f func()) {
	defer qt.Recovering("connect QScxmlStateMachine::start")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "start", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectStart() {
	defer qt.Recovering("disconnect QScxmlStateMachine::start")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "start")
	}
}

func (ptr *QScxmlStateMachine) Start() {
	defer qt.Recovering("QScxmlStateMachine::start")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Start(ptr.Pointer())
	}
}

func (ptr *QScxmlStateMachine) StateNames(compress bool) []string {
	defer qt.Recovering("QScxmlStateMachine::stateNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScxmlStateMachine_StateNames(ptr.Pointer(), C.int(qt.GoBoolToInt(compress)))), "|")
	}
	return make([]string, 0)
}

//export callbackQScxmlStateMachine_Stop
func callbackQScxmlStateMachine_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QScxmlStateMachine) ConnectStop(f func()) {
	defer qt.Recovering("connect QScxmlStateMachine::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectStop() {
	defer qt.Recovering("disconnect QScxmlStateMachine::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QScxmlStateMachine) Stop() {
	defer qt.Recovering("QScxmlStateMachine::stop")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_Stop(ptr.Pointer())
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent(event QScxmlEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::submitEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_SubmitEvent(ptr.Pointer(), PointerFromQScxmlEvent(event))
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent2(eventName string) {
	defer qt.Recovering("QScxmlStateMachine::submitEvent")

	if ptr.Pointer() != nil {
		var eventNameC = C.CString(eventName)
		defer C.free(unsafe.Pointer(eventNameC))
		C.QScxmlStateMachine_SubmitEvent2(ptr.Pointer(), eventNameC)
	}
}

func (ptr *QScxmlStateMachine) SubmitEvent3(eventName string, data core.QVariant_ITF) {
	defer qt.Recovering("QScxmlStateMachine::submitEvent")

	if ptr.Pointer() != nil {
		var eventNameC = C.CString(eventName)
		defer C.free(unsafe.Pointer(eventNameC))
		C.QScxmlStateMachine_SubmitEvent3(ptr.Pointer(), eventNameC, core.PointerFromQVariant(data))
	}
}

//export callbackQScxmlStateMachine_TimerEvent
func callbackQScxmlStateMachine_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QScxmlStateMachine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScxmlStateMachine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QScxmlStateMachine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QScxmlStateMachine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::timerEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQScxmlStateMachine_ChildEvent
func callbackQScxmlStateMachine_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScxmlStateMachine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScxmlStateMachine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QScxmlStateMachine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QScxmlStateMachine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::childEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQScxmlStateMachine_ConnectNotify
func callbackQScxmlStateMachine_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlStateMachine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QScxmlStateMachine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QScxmlStateMachine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlStateMachine::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlStateMachine::connectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_CustomEvent
func callbackQScxmlStateMachine_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QScxmlStateMachine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScxmlStateMachine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScxmlStateMachine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QScxmlStateMachine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QScxmlStateMachine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QScxmlStateMachine::customEvent")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQScxmlStateMachine_DeleteLater
func callbackQScxmlStateMachine_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QScxmlStateMachine::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QScxmlStateMachine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QScxmlStateMachine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QScxmlStateMachine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QScxmlStateMachine) DeleteLater() {
	defer qt.Recovering("QScxmlStateMachine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlStateMachine_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScxmlStateMachine) DeleteLaterDefault() {
	defer qt.Recovering("QScxmlStateMachine::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QScxmlStateMachine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQScxmlStateMachine_DisconnectNotify
func callbackQScxmlStateMachine_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QScxmlStateMachine::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQScxmlStateMachineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QScxmlStateMachine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QScxmlStateMachine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QScxmlStateMachine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QScxmlStateMachine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlStateMachine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QScxmlStateMachine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QScxmlStateMachine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QScxmlStateMachine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQScxmlStateMachine_Event
func callbackQScxmlStateMachine_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlStateMachine::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QScxmlStateMachine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlStateMachine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEvent() {
	defer qt.Recovering("disconnect QScxmlStateMachine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QScxmlStateMachine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlStateMachine::event")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlStateMachine::event")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_EventFilter
func callbackQScxmlStateMachine_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QScxmlStateMachine::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQScxmlStateMachineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QScxmlStateMachine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QScxmlStateMachine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QScxmlStateMachine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QScxmlStateMachine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlStateMachine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScxmlStateMachine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QScxmlStateMachine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QScxmlStateMachine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQScxmlStateMachine_MetaObject
func callbackQScxmlStateMachine_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QScxmlStateMachine::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQScxmlStateMachineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QScxmlStateMachine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QScxmlStateMachine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QScxmlStateMachine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QScxmlStateMachine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QScxmlStateMachine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QScxmlStateMachine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStateMachine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScxmlStateMachine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QScxmlStateMachine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScxmlStateMachine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
