// +build !minimal

package qml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "qml.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtQml_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func (ptr *QJSEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QJSEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) *QJSEngine {
	var n = new(QJSEngine)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QJSEngine__Extension
//QJSEngine::Extension
type QJSEngine__Extension int64

const (
	QJSEngine__TranslationExtension       QJSEngine__Extension = QJSEngine__Extension(0x1)
	QJSEngine__ConsoleExtension           QJSEngine__Extension = QJSEngine__Extension(0x2)
	QJSEngine__GarbageCollectionExtension QJSEngine__Extension = QJSEngine__Extension(0x4)
	QJSEngine__AllExtensions              QJSEngine__Extension = QJSEngine__Extension(0xffffffff)
)

func NewQJSEngine() *QJSEngine {
	var tmpValue = NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	var tmpValue = NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QJSEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	if ptr.Pointer() != nil {
		var programC = C.CString(program)
		defer C.free(unsafe.Pointer(programC))
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_Evaluate(ptr.Pointer(), programC, fileNameC, C.int(int32(lineNumber))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) InstallExtensions(extensions QJSEngine__Extension, object QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_InstallExtensions(ptr.Pointer(), C.longlong(extensions), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) NewArray(length uint) *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewArray(ptr.Pointer(), C.uint(uint32(length))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

//export callbackQJSEngine_DestroyQJSEngine
func callbackQJSEngine_DestroyQJSEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::~QJSEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQJSEngineFromPointer(ptr).DestroyQJSEngineDefault()
	}
}

func (ptr *QJSEngine) ConnectDestroyQJSEngine(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::~QJSEngine", f)
	}
}

func (ptr *QJSEngine) DisconnectDestroyQJSEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::~QJSEngine")
	}
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) DestroyQJSEngineDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngineDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QJSEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QJSEngine___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QJSEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QJSEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QJSEngine___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QJSEngine___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QJSEngine___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QJSEngine___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQJSEngine_TimerEvent
func callbackQJSEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::timerEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::timerEvent")
	}
}

func (ptr *QJSEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQJSEngine_ChildEvent
func callbackQJSEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::childEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::childEvent")
	}
}

func (ptr *QJSEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQJSEngine_ConnectNotify
func callbackQJSEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::connectNotify", f)
	}
}

func (ptr *QJSEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::connectNotify")
	}
}

func (ptr *QJSEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QJSEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_CustomEvent
func callbackQJSEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::customEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::customEvent")
	}
}

func (ptr *QJSEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQJSEngine_DeleteLater
func callbackQJSEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQJSEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QJSEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::deleteLater", f)
	}
}

func (ptr *QJSEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::deleteLater")
	}
}

func (ptr *QJSEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQJSEngine_DisconnectNotify
func callbackQJSEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::disconnectNotify", f)
	}
}

func (ptr *QJSEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::disconnectNotify")
	}
}

func (ptr *QJSEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QJSEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_Event
func callbackQJSEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QJSEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::event", f)
	}
}

func (ptr *QJSEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::event")
	}
}

func (ptr *QJSEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QJSEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQJSEngine_EventFilter
func callbackQJSEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QJSEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::eventFilter", f)
	}
}

func (ptr *QJSEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::eventFilter")
	}
}

func (ptr *QJSEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QJSEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QJSEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQJSEngine_MetaObject
func callbackQJSEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQJSEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QJSEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::metaObject", f)
	}
}

func (ptr *QJSEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::metaObject")
	}
}

func (ptr *QJSEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

func (ptr *QJSValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJSValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
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

//go:generate stringer -type=QJSValue__SpecialValue
//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      QJSValue__SpecialValue = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue QJSValue__SpecialValue = QJSValue__SpecialValue(1)
)

func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue(C.longlong(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue4(value bool) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.char(int8(qt.GoBoolToInt(value)))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue8(value string) *QJSValue {
	var valueC = C.CString(value)
	defer C.free(unsafe.Pointer(valueC))
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue8(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue10(value string) *QJSValue {
	var valueC = C.CString(value)
	defer C.free(unsafe.Pointer(valueC))
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue10(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue7(value float64) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue7(C.double(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue5(value int) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(int32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue6(value uint) *QJSValue {
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue6(C.uint(uint32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_DeleteProperty(ptr.Pointer(), nameC) != 0
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
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_HasOwnProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_HasProperty(ptr.Pointer(), nameC) != 0
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
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = NewQJSValueFromPointer(C.QJSValue_Property(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Property2(arrayIndex uint) *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Prototype() *QJSValue {
	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSValue_Prototype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QJSValue_SetProperty(ptr.Pointer(), nameC, PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetProperty2(arrayIndex uint, value QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty2(ptr.Pointer(), C.uint(uint32(arrayIndex)), PointerFromQJSValue(value))
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
		var tmpValue = core.NewQDateTimeFromPointer(C.QJSValue_ToDateTime(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ToInt() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJSValue_ToInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToNumber() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QJSValue_ToNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QJSValue_ToQObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJSValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QJSValue) ToUInt() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QJSValue_ToUInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) DestroyQJSValue() {
	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QJSValueIterator struct {
	ptr unsafe.Pointer
}

type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return ptr
}

func (ptr *QJSValueIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJSValueIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJSValueIterator(ptr QJSValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQJSValueIteratorFromPointer(ptr unsafe.Pointer) *QJSValueIterator {
	var n = new(QJSValueIterator)
	n.SetPointer(ptr)
	return n
}

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

func (ptr *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlAbstractUrlInterceptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptor_PTR().Pointer()
	}
	return nil
}

func NewQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) *QQmlAbstractUrlInterceptor {
	var n = new(QQmlAbstractUrlInterceptor)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QQmlAbstractUrlInterceptor__DataType
//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {
	return NewQQmlAbstractUrlInterceptorFromPointer(C.QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor())
}

//export callbackQQmlAbstractUrlInterceptor_Intercept
func callbackQQmlAbstractUrlInterceptor_Intercept(ptr unsafe.Pointer, url unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlAbstractUrlInterceptor::intercept"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, QQmlAbstractUrlInterceptor__DataType) *core.QUrl)(core.NewQUrlFromPointer(url), QQmlAbstractUrlInterceptor__DataType(ty)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectIntercept(f func(url *core.QUrl, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::intercept", f)
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectIntercept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::intercept")
	}
}

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlAbstractUrlInterceptor_Intercept(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor
func callbackQQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlAbstractUrlInterceptor::~QQmlAbstractUrlInterceptor"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlAbstractUrlInterceptorFromPointer(ptr).DestroyQQmlAbstractUrlInterceptorDefault()
	}
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectDestroyQQmlAbstractUrlInterceptor(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::~QQmlAbstractUrlInterceptor", f)
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectDestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::~QQmlAbstractUrlInterceptor")
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptorDefault() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return ptr
}

func (ptr *QQmlApplicationEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlApplicationEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQmlEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) *QQmlApplicationEngine {
	var n = new(QQmlApplicationEngine)
	n.SetPointer(ptr)
	return n
}
func NewQQmlApplicationEngine(parent core.QObject_ITF) *QQmlApplicationEngine {
	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {
	var filePathC = C.CString(filePath)
	defer C.free(unsafe.Pointer(filePathC))
	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(filePathC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {
	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlApplicationEngine_Load2
func callbackQQmlApplicationEngine_Load2(ptr unsafe.Pointer, filePath C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::load2"); signal != nil {
		signal.(func(string))(cGoUnpackString(filePath))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).Load2Default(cGoUnpackString(filePath))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoad2(f func(filePath string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load2", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load2")
	}
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	if ptr.Pointer() != nil {
		var filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
		C.QQmlApplicationEngine_Load2(ptr.Pointer(), filePathC)
	}
}

func (ptr *QQmlApplicationEngine) Load2Default(filePath string) {
	if ptr.Pointer() != nil {
		var filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
		C.QQmlApplicationEngine_Load2Default(ptr.Pointer(), filePathC)
	}
}

//export callbackQQmlApplicationEngine_Load
func callbackQQmlApplicationEngine_Load(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::load"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).LoadDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoad(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load")
	}
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_LoadData
func callbackQQmlApplicationEngine_LoadData(ptr unsafe.Pointer, data unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::loadData"); signal != nil {
		signal.(func(*core.QByteArray, *core.QUrl))(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).LoadDataDefault(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoadData(f func(data *core.QByteArray, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::loadData", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::loadData")
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadDataDefault(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_ObjectCreated
func callbackQQmlApplicationEngine_ObjectCreated(ptr unsafe.Pointer, object unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::objectCreated"); signal != nil {
		signal.(func(*core.QObject, *core.QUrl))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectObjectCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::objectCreated", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::objectCreated")
	}
}

func (ptr *QQmlApplicationEngine) ObjectCreated(object core.QObject_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ObjectCreated(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) RootObjects() []*core.QObject {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*core.QObject {
			var out = make([]*core.QObject, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQmlApplicationEngineFromPointer(l.data).__rootObjects_atList(i)
			}
			return out
		}(C.QQmlApplicationEngine_RootObjects(ptr.Pointer()))
	}
	return make([]*core.QObject, 0)
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlApplicationEngine) __rootObjects_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlApplicationEngine___rootObjects_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __rootObjects_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___rootObjects_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __rootObjects_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___rootObjects_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __importPlugin_errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlApplicationEngine___importPlugin_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __importPlugin_errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___importPlugin_errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlApplicationEngine) __importPlugin_errors_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___importPlugin_errors_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __warnings_warnings_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlApplicationEngine___warnings_warnings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __warnings_warnings_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___warnings_warnings_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlApplicationEngine) __warnings_warnings_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___warnings_warnings_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlApplicationEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlApplicationEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlApplicationEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlApplicationEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlApplicationEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlApplicationEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlApplicationEngine___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlApplicationEngine_Event
func callbackQQmlApplicationEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlApplicationEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlApplicationEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::event", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::event")
	}
}

func (ptr *QQmlApplicationEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlApplicationEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlApplicationEngine_TimerEvent
func callbackQQmlApplicationEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::timerEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::timerEvent")
	}
}

func (ptr *QQmlApplicationEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlApplicationEngine_ChildEvent
func callbackQQmlApplicationEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::childEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::childEvent")
	}
}

func (ptr *QQmlApplicationEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlApplicationEngine_ConnectNotify
func callbackQQmlApplicationEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::connectNotify", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::connectNotify")
	}
}

func (ptr *QQmlApplicationEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlApplicationEngine_CustomEvent
func callbackQQmlApplicationEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::customEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::customEvent")
	}
}

func (ptr *QQmlApplicationEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlApplicationEngine_DeleteLater
func callbackQQmlApplicationEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlApplicationEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::deleteLater", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::deleteLater")
	}
}

func (ptr *QQmlApplicationEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlApplicationEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlApplicationEngine_DisconnectNotify
func callbackQQmlApplicationEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::disconnectNotify", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::disconnectNotify")
	}
}

func (ptr *QQmlApplicationEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlApplicationEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlApplicationEngine_EventFilter
func callbackQQmlApplicationEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlApplicationEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlApplicationEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::eventFilter", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::eventFilter")
	}
}

func (ptr *QQmlApplicationEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlApplicationEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlApplicationEngine_MetaObject
func callbackQQmlApplicationEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlApplicationEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlApplicationEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::metaObject", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::metaObject")
	}
}

func (ptr *QQmlApplicationEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlApplicationEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlApplicationEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlComponent struct {
	core.QObject
}

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return ptr
}

func (ptr *QQmlComponent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlComponent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlComponent(ptr QQmlComponent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponent_PTR().Pointer()
	}
	return nil
}

func NewQQmlComponentFromPointer(ptr unsafe.Pointer) *QQmlComponent {
	var n = new(QQmlComponent)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QQmlComponent__CompilationMode
//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(1)
)

//go:generate stringer -type=QQmlComponent__Status
//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    QQmlComponent__Status = QQmlComponent__Status(0)
	QQmlComponent__Ready   QQmlComponent__Status = QQmlComponent__Status(1)
	QQmlComponent__Loading QQmlComponent__Status = QQmlComponent__Status(2)
	QQmlComponent__Error   QQmlComponent__Status = QQmlComponent__Status(3)
)

func (ptr *QQmlComponent) Progress() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlComponent_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), fileNameC, C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), fileNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlComponent_BeginCreate
func callbackQQmlComponent_BeginCreate(ptr unsafe.Pointer, publicContext unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::beginCreate"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(publicContext)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).BeginCreateDefault(NewQQmlContextFromPointer(publicContext)))
}

func (ptr *QQmlComponent) ConnectBeginCreate(f func(publicContext *QQmlContext) *core.QObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::beginCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectBeginCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::beginCreate")
	}
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent_BeginCreate(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) BeginCreateDefault(publicContext QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent_BeginCreateDefault(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlComponent_CompleteCreate
func callbackQQmlComponent_CompleteCreate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::completeCreate"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).CompleteCreateDefault()
	}
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::completeCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::completeCreate")
	}
}

func (ptr *QQmlComponent) CompleteCreate() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) CompleteCreateDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreateDefault(ptr.Pointer())
	}
}

//export callbackQQmlComponent_Create
func callbackQQmlComponent_Create(ptr unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::create"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(context)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).CreateDefault(NewQQmlContextFromPointer(context)))
}

func (ptr *QQmlComponent) ConnectCreate(f func(context *QQmlContext) *core.QObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::create", f)
	}
}

func (ptr *QQmlComponent) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::create")
	}
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent_Create(ptr.Pointer(), PointerFromQQmlContext(context)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) CreateDefault(context QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent_CreateDefault(ptr.Pointer(), PointerFromQQmlContext(context)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) Errors() []*QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*QQmlError {
			var out = make([]*QQmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQmlComponentFromPointer(l.data).__errors_atList(i)
			}
			return out
		}(C.QQmlComponent_Errors(ptr.Pointer()))
	}
	return make([]*QQmlError, 0)
}

func (ptr *QQmlComponent) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQmlComponent_LoadUrl
func callbackQQmlComponent_LoadUrl(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::loadUrl"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQmlComponentFromPointer(ptr).LoadUrlDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlComponent) ConnectLoadUrl(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl", f)
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl")
	}
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) LoadUrlDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrlDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_LoadUrl2
func callbackQQmlComponent_LoadUrl2(ptr unsafe.Pointer, url unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::loadUrl2"); signal != nil {
		signal.(func(*core.QUrl, QQmlComponent__CompilationMode))(core.NewQUrlFromPointer(url), QQmlComponent__CompilationMode(mode))
	} else {
		NewQQmlComponentFromPointer(ptr).LoadUrl2Default(core.NewQUrlFromPointer(url), QQmlComponent__CompilationMode(mode))
	}
}

func (ptr *QQmlComponent) ConnectLoadUrl2(f func(url *core.QUrl, mode QQmlComponent__CompilationMode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl2", f)
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl2")
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(mode))
	}
}

func (ptr *QQmlComponent) LoadUrl2Default(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2Default(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(mode))
	}
}

//export callbackQQmlComponent_ProgressChanged
func callbackQQmlComponent_ProgressChanged(ptr unsafe.Pointer, progress C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::progressChanged"); signal != nil {
		signal.(func(float64))(float64(progress))
	}

}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectProgressChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::progressChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectProgressChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::progressChanged")
	}
}

func (ptr *QQmlComponent) ProgressChanged(progress float64) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ProgressChanged(ptr.Pointer(), C.double(progress))
	}
}

//export callbackQQmlComponent_SetData
func callbackQQmlComponent_SetData(ptr unsafe.Pointer, data unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::setData"); signal != nil {
		signal.(func(*core.QByteArray, *core.QUrl))(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	} else {
		NewQQmlComponentFromPointer(ptr).SetDataDefault(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlComponent) ConnectSetData(f func(data *core.QByteArray, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::setData", f)
	}
}

func (ptr *QQmlComponent) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::setData")
	}
}

func (ptr *QQmlComponent) SetData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) SetDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetDataDefault(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_StatusChanged
func callbackQQmlComponent_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::statusChanged"); signal != nil {
		signal.(func(QQmlComponent__Status))(QQmlComponent__Status(status))
	}

}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::statusChanged")
	}
}

func (ptr *QQmlComponent) StatusChanged(status QQmlComponent__Status) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

//export callbackQQmlComponent_DestroyQQmlComponent
func callbackQQmlComponent_DestroyQQmlComponent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::~QQmlComponent"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).DestroyQQmlComponentDefault()
	}
}

func (ptr *QQmlComponent) ConnectDestroyQQmlComponent(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::~QQmlComponent", f)
	}
}

func (ptr *QQmlComponent) DisconnectDestroyQQmlComponent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::~QQmlComponent")
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponentDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponentDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) __errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlComponent___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlComponent) __errors_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___errors_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlComponent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlComponent) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlComponent___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlComponent_TimerEvent
func callbackQQmlComponent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::timerEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::timerEvent")
	}
}

func (ptr *QQmlComponent) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlComponent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlComponent_ChildEvent
func callbackQQmlComponent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::childEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::childEvent")
	}
}

func (ptr *QQmlComponent) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlComponent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlComponent_ConnectNotify
func callbackQQmlComponent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::connectNotify", f)
	}
}

func (ptr *QQmlComponent) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::connectNotify")
	}
}

func (ptr *QQmlComponent) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlComponent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_CustomEvent
func callbackQQmlComponent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::customEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::customEvent")
	}
}

func (ptr *QQmlComponent) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlComponent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlComponent_DeleteLater
func callbackQQmlComponent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlComponent) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::deleteLater", f)
	}
}

func (ptr *QQmlComponent) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::deleteLater")
	}
}

func (ptr *QQmlComponent) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlComponent_DisconnectNotify
func callbackQQmlComponent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::disconnectNotify", f)
	}
}

func (ptr *QQmlComponent) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::disconnectNotify")
	}
}

func (ptr *QQmlComponent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlComponent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_Event
func callbackQQmlComponent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlComponent) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::event", f)
	}
}

func (ptr *QQmlComponent) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::event")
	}
}

func (ptr *QQmlComponent) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlComponent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlComponent_EventFilter
func callbackQQmlComponent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlComponent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::eventFilter", f)
	}
}

func (ptr *QQmlComponent) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::eventFilter")
	}
}

func (ptr *QQmlComponent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlComponent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlComponent_MetaObject
func callbackQQmlComponent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlComponentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlComponent) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::metaObject", f)
	}
}

func (ptr *QQmlComponent) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::metaObject")
	}
}

func (ptr *QQmlComponent) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlComponent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlComponent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func (ptr *QQmlContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) *QQmlContext {
	var n = new(QQmlContext)
	n.SetPointer(ptr)
	return n
}
func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {
	var tmpValue = NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	var tmpValue = NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlContext_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlContext_ContextObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QQmlContext_ContextProperty(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlEngineFromPointer(C.QQmlContext_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlContextFromPointer(C.QQmlContext_ParentContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ResolvedUrl(src core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlContext_ResolvedUrl(ptr.Pointer(), core.PointerFromQUrl(src)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QQmlContext_SetContextProperty(ptr.Pointer(), nameC, core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), nameC, core.PointerFromQVariant(value))
	}
}

//export callbackQQmlContext_DestroyQQmlContext
func callbackQQmlContext_DestroyQQmlContext(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::~QQmlContext"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlContextFromPointer(ptr).DestroyQQmlContextDefault()
	}
}

func (ptr *QQmlContext) ConnectDestroyQQmlContext(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::~QQmlContext", f)
	}
}

func (ptr *QQmlContext) DisconnectDestroyQQmlContext() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::~QQmlContext")
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) DestroyQQmlContextDefault() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContextDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlContext___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlContext___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlContext___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlContext) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlContext___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlContext___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlContext___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlContext___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlContext_TimerEvent
func callbackQQmlContext_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::timerEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::timerEvent")
	}
}

func (ptr *QQmlContext) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlContext) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlContext_ChildEvent
func callbackQQmlContext_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::childEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::childEvent")
	}
}

func (ptr *QQmlContext) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlContext) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlContext_ConnectNotify
func callbackQQmlContext_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::connectNotify", f)
	}
}

func (ptr *QQmlContext) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::connectNotify")
	}
}

func (ptr *QQmlContext) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_CustomEvent
func callbackQQmlContext_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::customEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::customEvent")
	}
}

func (ptr *QQmlContext) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlContext) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlContext_DeleteLater
func callbackQQmlContext_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlContextFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlContext) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::deleteLater", f)
	}
}

func (ptr *QQmlContext) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::deleteLater")
	}
}

func (ptr *QQmlContext) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlContext_DisconnectNotify
func callbackQQmlContext_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::disconnectNotify", f)
	}
}

func (ptr *QQmlContext) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::disconnectNotify")
	}
}

func (ptr *QQmlContext) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_Event
func callbackQQmlContext_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlContext) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::event", f)
	}
}

func (ptr *QQmlContext) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::event")
	}
}

func (ptr *QQmlContext) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlContext) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlContext_EventFilter
func callbackQQmlContext_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlContext) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::eventFilter", f)
	}
}

func (ptr *QQmlContext) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::eventFilter")
	}
}

func (ptr *QQmlContext) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlContext_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlContext_MetaObject
func callbackQQmlContext_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlContextFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlContext) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::metaObject", f)
	}
}

func (ptr *QQmlContext) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::metaObject")
	}
}

func (ptr *QQmlContext) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlContext_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlContext_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

func (ptr *QQmlEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QJSEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) *QQmlEngine {
	var n = new(QQmlEngine)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QQmlEngine__ObjectOwnership
//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(1)
)

func (ptr *QQmlEngine) OfflineStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	if ptr.Pointer() != nil {
		var dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), dirC)
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	var tmpValue = NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), providerIdC, PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQmlEngine_AddImportPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQmlEngine) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlEngine_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) ClearComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	var tmpValue = NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) ContextForObject(object core.QObject_ITF) *QQmlContext {
	var tmpValue = NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlEngine_Event
func callbackQQmlEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::event", f)
	}
}

func (ptr *QQmlEngine) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::event")
	}
}

func (ptr *QQmlEngine) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), providerIdC))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlEngine_ImportPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) ImportPlugin(filePath string, uri string, errors []*QQmlError) bool {
	if ptr.Pointer() != nil {
		var filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		return C.QQmlEngine_ImportPlugin(ptr.Pointer(), filePathC, uriC, func() unsafe.Pointer {
			var tmpList = NewQQmlEngineFromPointer(NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(1))).__importPlugin_errors_newList())
			for _, v := range errors {
				tmpList.__importPlugin_errors_setList(v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlEngine_PluginPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQQmlEngine_Quit
func callbackQQmlEngine_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::quit"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::quit")
	}
}

func (ptr *QQmlEngine) Quit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Quit(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), providerIdC)
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), pathsC)
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), pathsC)
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

//export callbackQQmlEngine_Warnings
func callbackQQmlEngine_Warnings(ptr unsafe.Pointer, warnings C.struct_QtQml_PackedList) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::warnings"); signal != nil {
		signal.(func([]*QQmlError))(func(l C.struct_QtQml_PackedList) []*QQmlError {
			var out = make([]*QQmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQmlEngineFromPointer(l.data).__warnings_warnings_atList(i)
			}
			return out
		}(warnings))
	}

}

func (ptr *QQmlEngine) ConnectWarnings(f func(warnings []*QQmlError)) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectWarnings(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::warnings", f)
	}
}

func (ptr *QQmlEngine) DisconnectWarnings() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectWarnings(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::warnings")
	}
}

func (ptr *QQmlEngine) Warnings(warnings []*QQmlError) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Warnings(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQQmlEngineFromPointer(NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(1))).__warnings_warnings_newList())
			for _, v := range warnings {
				tmpList.__warnings_warnings_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQQmlEngine_DestroyQQmlEngine
func callbackQQmlEngine_DestroyQQmlEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::~QQmlEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlEngineFromPointer(ptr).DestroyQQmlEngineDefault()
	}
}

func (ptr *QQmlEngine) ConnectDestroyQQmlEngine(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::~QQmlEngine", f)
	}
}

func (ptr *QQmlEngine) DisconnectDestroyQQmlEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::~QQmlEngine")
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngineDefault() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngineDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QQmlEngine_QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC = C.CString(uri)
	defer C.free(unsafe.Pointer(uriC))
	var qmlNameC = C.CString(qmlName)
	defer C.free(unsafe.Pointer(qmlNameC))
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterSingletonType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC = C.CString(uri)
	defer C.free(unsafe.Pointer(uriC))
	var qmlNameC = C.CString(qmlName)
	defer C.free(unsafe.Pointer(qmlNameC))
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterSingletonType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) __importPlugin_errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlEngine___importPlugin_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __importPlugin_errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___importPlugin_errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __importPlugin_errors_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___importPlugin_errors_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __warnings_warnings_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlEngine___warnings_warnings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __warnings_warnings_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___warnings_warnings_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __warnings_warnings_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___warnings_warnings_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlEngine) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlEngine) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlEngine) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlEngine) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlEngine___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlEngine_TimerEvent
func callbackQQmlEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::timerEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::timerEvent")
	}
}

func (ptr *QQmlEngine) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlEngine_ChildEvent
func callbackQQmlEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::childEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::childEvent")
	}
}

func (ptr *QQmlEngine) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlEngine_ConnectNotify
func callbackQQmlEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::connectNotify", f)
	}
}

func (ptr *QQmlEngine) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::connectNotify")
	}
}

func (ptr *QQmlEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlEngine_CustomEvent
func callbackQQmlEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::customEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::customEvent")
	}
}

func (ptr *QQmlEngine) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlEngine_DeleteLater
func callbackQQmlEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlEngine) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::deleteLater", f)
	}
}

func (ptr *QQmlEngine) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::deleteLater")
	}
}

func (ptr *QQmlEngine) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlEngine_DisconnectNotify
func callbackQQmlEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::disconnectNotify", f)
	}
}

func (ptr *QQmlEngine) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::disconnectNotify")
	}
}

func (ptr *QQmlEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlEngine_EventFilter
func callbackQQmlEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::eventFilter", f)
	}
}

func (ptr *QQmlEngine) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::eventFilter")
	}
}

func (ptr *QQmlEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlEngine_MetaObject
func callbackQQmlEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::metaObject", f)
	}
}

func (ptr *QQmlEngine) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::metaObject")
	}
}

func (ptr *QQmlEngine) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func (ptr *QQmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = new(QQmlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlError) DestroyQQmlError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQQmlError() *QQmlError {
	var tmpValue = NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	var tmpValue = NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func (ptr *QQmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Line(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlError) SetColumn(column int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QQmlError_SetDescription(ptr.Pointer(), descriptionC)
	}
}

func (ptr *QQmlError) SetLine(line int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(int32(line)))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlError_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QQmlExpression struct {
	core.QObject
}

type QQmlExpression_ITF interface {
	core.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return ptr
}

func (ptr *QQmlExpression) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExpression) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExpression(ptr QQmlExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpression_PTR().Pointer()
	}
	return nil
}

func NewQQmlExpressionFromPointer(ptr unsafe.Pointer) *QQmlExpression {
	var n = new(QQmlExpression)
	n.SetPointer(ptr)
	return n
}
func NewQQmlExpression() *QQmlExpression {
	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {
	var expressionC = C.CString(expression)
	defer C.free(unsafe.Pointer(expressionC))
	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression2(PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), expressionC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression3(script QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {
	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression3(PointerFromQQmlScriptString(script), PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlExpression) ClearError() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlContextFromPointer(C.QQmlExpression_Context(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlEngineFromPointer(C.QQmlExpression_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Error() *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlExpression_Error(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined bool) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQmlExpression_Evaluate(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(valueIsUndefined)))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Expression() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlExpression_Expression(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_LineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_NotifyOnValueChanged(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExpression_ScopeObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) SetExpression(expression string) {
	if ptr.Pointer() != nil {
		var expressionC = C.CString(expression)
		defer C.free(unsafe.Pointer(expressionC))
		C.QQmlExpression_SetExpression(ptr.Pointer(), expressionC)
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(notifyOnChange))))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	if ptr.Pointer() != nil {
		var urlC = C.CString(url)
		defer C.free(unsafe.Pointer(urlC))
		C.QQmlExpression_SetSourceLocation(ptr.Pointer(), urlC, C.int(int32(line)), C.int(int32(column)))
	}
}

func (ptr *QQmlExpression) SourceFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlExpression_SourceFile(ptr.Pointer()))
	}
	return ""
}

//export callbackQQmlExpression_ValueChanged
func callbackQQmlExpression_ValueChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::valueChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::valueChanged", f)
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::valueChanged")
	}
}

func (ptr *QQmlExpression) ValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ValueChanged(ptr.Pointer())
	}
}

//export callbackQQmlExpression_DestroyQQmlExpression
func callbackQQmlExpression_DestroyQQmlExpression(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::~QQmlExpression"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExpressionFromPointer(ptr).DestroyQQmlExpressionDefault()
	}
}

func (ptr *QQmlExpression) ConnectDestroyQQmlExpression(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::~QQmlExpression", f)
	}
}

func (ptr *QQmlExpression) DisconnectDestroyQQmlExpression() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::~QQmlExpression")
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpressionDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpressionDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExpression) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExpression___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExpression___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlExpression___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlExpression) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExpression___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExpression___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExpression___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExpression___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlExpression_TimerEvent
func callbackQQmlExpression_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::timerEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::timerEvent")
	}
}

func (ptr *QQmlExpression) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExpression) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlExpression_ChildEvent
func callbackQQmlExpression_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::childEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::childEvent")
	}
}

func (ptr *QQmlExpression) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExpression) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExpression_ConnectNotify
func callbackQQmlExpression_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::connectNotify", f)
	}
}

func (ptr *QQmlExpression) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::connectNotify")
	}
}

func (ptr *QQmlExpression) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExpression) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_CustomEvent
func callbackQQmlExpression_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::customEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::customEvent")
	}
}

func (ptr *QQmlExpression) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExpression) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExpression_DeleteLater
func callbackQQmlExpression_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExpressionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExpression) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::deleteLater", f)
	}
}

func (ptr *QQmlExpression) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::deleteLater")
	}
}

func (ptr *QQmlExpression) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExpression) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExpression_DisconnectNotify
func callbackQQmlExpression_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::disconnectNotify", f)
	}
}

func (ptr *QQmlExpression) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::disconnectNotify")
	}
}

func (ptr *QQmlExpression) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExpression) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_Event
func callbackQQmlExpression_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExpression) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::event", f)
	}
}

func (ptr *QQmlExpression) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::event")
	}
}

func (ptr *QQmlExpression) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlExpression) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlExpression_EventFilter
func callbackQQmlExpression_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExpression) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::eventFilter", f)
	}
}

func (ptr *QQmlExpression) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::eventFilter")
	}
}

func (ptr *QQmlExpression) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlExpression) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlExpression_MetaObject
func callbackQQmlExpression_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExpressionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExpression) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::metaObject", f)
	}
}

func (ptr *QQmlExpression) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::metaObject")
	}
}

func (ptr *QQmlExpression) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExpression_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExpression_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlExtensionPlugin) DestroyQQmlExtensionPlugin() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExtensionPlugin_InitializeEngine
func callbackQQmlExtensionPlugin_InitializeEngine(ptr unsafe.Pointer, engine unsafe.Pointer, uri C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::initializeEngine"); signal != nil {
		signal.(func(*QQmlEngine, string))(NewQQmlEngineFromPointer(engine), cGoUnpackString(uri))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).InitializeEngineDefault(NewQQmlEngineFromPointer(engine), cGoUnpackString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectInitializeEngine(f func(engine *QQmlEngine, uri string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::initializeEngine", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectInitializeEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::initializeEngine")
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_InitializeEngine(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngineDefault(engine QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_InitializeEngineDefault(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

func NewQQmlExtensionPlugin(parent core.QObject_ITF) *QQmlExtensionPlugin {
	var tmpValue = NewQQmlExtensionPluginFromPointer(C.QQmlExtensionPlugin_NewQQmlExtensionPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlExtensionPlugin_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQmlExtensionPlugin_RegisterTypes
func callbackQQmlExtensionPlugin_RegisterTypes(ptr unsafe.Pointer, uri C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::registerTypes"); signal != nil {
		signal.(func(string))(cGoUnpackString(uri))
	}

}

func (ptr *QQmlExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::registerTypes", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectRegisterTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::registerTypes")
	}
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), uriC)
	}
}

func (ptr *QQmlExtensionPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExtensionPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExtensionPlugin___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlExtensionPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExtensionPlugin___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExtensionPlugin___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExtensionPlugin___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlExtensionPlugin___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlExtensionPlugin_TimerEvent
func callbackQQmlExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::timerEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::timerEvent")
	}
}

func (ptr *QQmlExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_ChildEvent
func callbackQQmlExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::childEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::childEvent")
	}
}

func (ptr *QQmlExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_ConnectNotify
func callbackQQmlExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::connectNotify", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::connectNotify")
	}
}

func (ptr *QQmlExtensionPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_CustomEvent
func callbackQQmlExtensionPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::customEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::customEvent")
	}
}

func (ptr *QQmlExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_DeleteLater
func callbackQQmlExtensionPlugin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExtensionPlugin) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::deleteLater", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::deleteLater")
	}
}

func (ptr *QQmlExtensionPlugin) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExtensionPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExtensionPlugin_DisconnectNotify
func callbackQQmlExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::disconnectNotify", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::disconnectNotify")
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_Event
func callbackQQmlExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExtensionPlugin) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::event", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::event")
	}
}

func (ptr *QQmlExtensionPlugin) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_EventFilter
func callbackQQmlExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExtensionPlugin) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::eventFilter", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::eventFilter")
	}
}

func (ptr *QQmlExtensionPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_MetaObject
func callbackQQmlExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExtensionPlugin) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::metaObject", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::metaObject")
	}
}

func (ptr *QQmlExtensionPlugin) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExtensionPlugin_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func (ptr *QQmlFileSelector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlFileSelector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	return n
}
func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {
	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlFileSelector) Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlFileSelector) Selector() *core.QFileSelector {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQFileSelectorFromPointer(C.QQmlFileSelector_Selector(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	if ptr.Pointer() != nil {
		var strinC = C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), strinC)
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	if ptr.Pointer() != nil {
		var strinC = C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), strinC)
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlFileSelector) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlFileSelector___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlFileSelector___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlFileSelector___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlFileSelector___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlFileSelector___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlFileSelector___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlFileSelector___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlFileSelector_TimerEvent
func callbackQQmlFileSelector_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::timerEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::timerEvent")
	}
}

func (ptr *QQmlFileSelector) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlFileSelector) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlFileSelector_ChildEvent
func callbackQQmlFileSelector_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::childEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::childEvent")
	}
}

func (ptr *QQmlFileSelector) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlFileSelector) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlFileSelector_ConnectNotify
func callbackQQmlFileSelector_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::connectNotify", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::connectNotify")
	}
}

func (ptr *QQmlFileSelector) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_CustomEvent
func callbackQQmlFileSelector_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::customEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::customEvent")
	}
}

func (ptr *QQmlFileSelector) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlFileSelector) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlFileSelector_DeleteLater
func callbackQQmlFileSelector_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlFileSelector) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::deleteLater", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::deleteLater")
	}
}

func (ptr *QQmlFileSelector) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlFileSelector) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlFileSelector_DisconnectNotify
func callbackQQmlFileSelector_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::disconnectNotify", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::disconnectNotify")
	}
}

func (ptr *QQmlFileSelector) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlFileSelector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_Event
func callbackQQmlFileSelector_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlFileSelector) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::event", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::event")
	}
}

func (ptr *QQmlFileSelector) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlFileSelector) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlFileSelector_EventFilter
func callbackQQmlFileSelector_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlFileSelector) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::eventFilter", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::eventFilter")
	}
}

func (ptr *QQmlFileSelector) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlFileSelector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlFileSelector_MetaObject
func callbackQQmlFileSelector_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlFileSelectorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlFileSelector) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::metaObject", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::metaObject")
	}
}

func (ptr *QQmlFileSelector) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlFileSelector_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlFileSelector_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return ptr
}

func (ptr *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlImageProviderBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func NewQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) *QQmlImageProviderBase {
	var n = new(QQmlImageProviderBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlImageProviderBase) DestroyQQmlImageProviderBase() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QQmlImageProviderBase__Flag
//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading QQmlImageProviderBase__Flag = QQmlImageProviderBase__Flag(0x01)
)

//go:generate stringer -type=QQmlImageProviderBase__ImageType
//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image         QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap        QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(3)
	QQmlImageProviderBase__ImageResponse QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(4)
)

//export callbackQQmlImageProviderBase_Flags
func callbackQQmlImageProviderBase_Flags(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlImageProviderBase::flags"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__Flag)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectFlags(f func() QQmlImageProviderBase__Flag) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::flags", f)
	}
}

func (ptr *QQmlImageProviderBase) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::flags")
	}
}

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlImageProviderBase_ImageType
func callbackQQmlImageProviderBase_ImageType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlImageProviderBase::imageType"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectImageType(f func() QQmlImageProviderBase__ImageType) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::imageType", f)
	}
}

func (ptr *QQmlImageProviderBase) DisconnectImageType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::imageType")
	}
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__ImageType(C.QQmlImageProviderBase_ImageType(ptr.Pointer()))
	}
	return 0
}

type QQmlIncubationController struct {
	ptr unsafe.Pointer
}

type QQmlIncubationController_ITF interface {
	QQmlIncubationController_PTR() *QQmlIncubationController
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func (ptr *QQmlIncubationController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlIncubationController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationController_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) *QQmlIncubationController {
	var n = new(QQmlIncubationController)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlIncubationController) DestroyQQmlIncubationController() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQQmlIncubationController() *QQmlIncubationController {
	return NewQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlEngineFromPointer(C.QQmlIncubationController_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQQmlIncubationController_IncubatingObjectCountChanged
func callbackQQmlIncubationController_IncubatingObjectCountChanged(ptr unsafe.Pointer, incubatingObjectCount C.int) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubationController::incubatingObjectCountChanged"); signal != nil {
		signal.(func(int))(int(int32(incubatingObjectCount)))
	} else {
		NewQQmlIncubationControllerFromPointer(ptr).IncubatingObjectCountChangedDefault(int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubationController::incubatingObjectCountChanged", f)
	}
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubationController::incubatingObjectCountChanged")
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChanged(incubatingObjectCount int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChanged(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChangedDefault(incubatingObjectCount int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChangedDefault(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return ptr
}

func (ptr *QQmlIncubator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlIncubator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlIncubator(ptr QQmlIncubator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubator_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubatorFromPointer(ptr unsafe.Pointer) *QQmlIncubator {
	var n = new(QQmlIncubator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlIncubator) DestroyQQmlIncubator() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QQmlIncubator__IncubationMode
//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(2)
)

//go:generate stringer -type=QQmlIncubator__Status
//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    QQmlIncubator__Status = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   QQmlIncubator__Status = QQmlIncubator__Status(1)
	QQmlIncubator__Loading QQmlIncubator__Status = QQmlIncubator__Status(2)
	QQmlIncubator__Error   QQmlIncubator__Status = QQmlIncubator__Status(3)
)

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	return NewQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.longlong(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) Errors() []*QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*QQmlError {
			var out = make([]*QQmlError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQQmlIncubatorFromPointer(l.data).__errors_atList(i)
			}
			return out
		}(C.QQmlIncubator_Errors(ptr.Pointer()))
	}
	return make([]*QQmlError, 0)
}

func (ptr *QQmlIncubator) ForceCompletion() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlIncubator_Object(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlIncubator_SetInitialState
func callbackQQmlIncubator_SetInitialState(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubator::setInitialState"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQQmlIncubatorFromPointer(ptr).SetInitialStateDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::setInitialState", f)
	}
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::setInitialState")
	}
}

func (ptr *QQmlIncubator) SetInitialState(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialState(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) SetInitialStateDefault(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialStateDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlIncubator_StatusChanged
func callbackQQmlIncubator_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubator::statusChanged"); signal != nil {
		signal.(func(QQmlIncubator__Status))(QQmlIncubator__Status(status))
	} else {
		NewQQmlIncubatorFromPointer(ptr).StatusChangedDefault(QQmlIncubator__Status(status))
	}
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::statusChanged", f)
	}
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::statusChanged")
	}
}

func (ptr *QQmlIncubator) StatusChanged(status QQmlIncubator__Status) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlIncubator) StatusChangedDefault(status QQmlIncubator__Status) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChangedDefault(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlIncubator) __errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlIncubator___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlIncubator) __errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator___errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlIncubator) __errors_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlIncubator___errors_newList(ptr.Pointer()))
	}
	return nil
}

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (ptr *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return ptr
}

func (ptr *QQmlListProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlListProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlListProperty(ptr QQmlListProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlListPropertyFromPointer(ptr unsafe.Pointer) *QQmlListProperty {
	var n = new(QQmlListProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListProperty) DestroyQQmlListProperty() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (ptr *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return ptr
}

func (ptr *QQmlListReference) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlListReference) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlListReference(ptr QQmlListReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReference_PTR().Pointer()
	}
	return nil
}

func NewQQmlListReferenceFromPointer(ptr unsafe.Pointer) *QQmlListReference {
	var n = new(QQmlListReference)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListReference) DestroyQQmlListReference() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQQmlListReference() *QQmlListReference {
	var tmpValue = NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference())
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {
	var propertyC = C.CString(property)
	defer C.free(unsafe.Pointer(propertyC))
	var tmpValue = NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference2(core.PointerFromQObject(object), propertyC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_Append(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlListReference_At(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlListReference) CanAppend() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAppend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAt(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanClear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanCount(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlListReference_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlListReference) IsManipulable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsManipulable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlListReference_ListElementType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlListReference_Object(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QQmlNetworkAccessManagerFactory struct {
	ptr unsafe.Pointer
}

type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactory_PTR().Pointer()
	}
	return nil
}

func NewQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	var n = new(QQmlNetworkAccessManagerFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQQmlNetworkAccessManagerFactory_Create
func callbackQQmlNetworkAccessManagerFactory_Create(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNetworkAccessManagerFactory::create"); signal != nil {
		return network.PointerFromQNetworkAccessManager(signal.(func(*core.QObject) *network.QNetworkAccessManager)(core.NewQObjectFromPointer(parent)))
	}

	return network.PointerFromQNetworkAccessManager(network.NewQNetworkAccessManager(nil))
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectCreate(f func(parent *core.QObject) *network.QNetworkAccessManager) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::create", f)
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::create")
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory
func callbackQQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlNetworkAccessManagerFactoryFromPointer(ptr).DestroyQQmlNetworkAccessManagerFactoryDefault()
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectDestroyQQmlNetworkAccessManagerFactory(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory", f)
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectDestroyQQmlNetworkAccessManagerFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory")
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactoryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQQmlNetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlNetworkAccessManagerFactory_NewQQmlNetworkAccessManagerFactory())
}

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlParserStatus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) *QQmlParserStatus {
	var n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlParserStatus) DestroyQQmlParserStatus() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlParserStatus_ClassBegin
func callbackQQmlParserStatus_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlParserStatus::classBegin"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::classBegin", f)
	}
}

func (ptr *QQmlParserStatus) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::classBegin")
	}
}

func (ptr *QQmlParserStatus) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

//export callbackQQmlParserStatus_ComponentComplete
func callbackQQmlParserStatus_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlParserStatus::componentComplete"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::componentComplete", f)
	}
}

func (ptr *QQmlParserStatus) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::componentComplete")
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}

type QQmlPrivate struct {
	ptr unsafe.Pointer
}

type QQmlPrivate_ITF interface {
	QQmlPrivate_PTR() *QQmlPrivate
}

func (ptr *QQmlPrivate) QQmlPrivate_PTR() *QQmlPrivate {
	return ptr
}

func (ptr *QQmlPrivate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlPrivate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlPrivate(ptr QQmlPrivate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPrivate_PTR().Pointer()
	}
	return nil
}

func NewQQmlPrivateFromPointer(ptr unsafe.Pointer) *QQmlPrivate {
	var n = new(QQmlPrivate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlPrivate) DestroyQQmlPrivate() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QQmlPrivate__AutoParentResult
//QQmlPrivate::AutoParentResult
type QQmlPrivate__AutoParentResult int64

const (
	QQmlPrivate__Parented           QQmlPrivate__AutoParentResult = QQmlPrivate__AutoParentResult(0)
	QQmlPrivate__IncompatibleObject QQmlPrivate__AutoParentResult = QQmlPrivate__AutoParentResult(1)
	QQmlPrivate__IncompatibleParent QQmlPrivate__AutoParentResult = QQmlPrivate__AutoParentResult(2)
)

//go:generate stringer -type=QQmlPrivate__RegistrationType
//QQmlPrivate::RegistrationType
type QQmlPrivate__RegistrationType int64

const (
	QQmlPrivate__TypeRegistration               QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(0)
	QQmlPrivate__InterfaceRegistration          QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(1)
	QQmlPrivate__AutoParentRegistration         QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(2)
	QQmlPrivate__SingletonRegistration          QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(3)
	QQmlPrivate__CompositeRegistration          QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(4)
	QQmlPrivate__CompositeSingletonRegistration QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(5)
	QQmlPrivate__QmlUnitCacheHookRegistration   QQmlPrivate__RegistrationType = QQmlPrivate__RegistrationType(6)
)

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return ptr
}

func (ptr *QQmlProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlProperty(ptr QQmlProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyFromPointer(ptr unsafe.Pointer) *QQmlProperty {
	var n = new(QQmlProperty)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlProperty) DestroyQQmlProperty() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QQmlProperty__PropertyTypeCategory
//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(3)
)

//go:generate stringer -type=QQmlProperty__Type
//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        QQmlProperty__Type = QQmlProperty__Type(0)
	QQmlProperty__Property       QQmlProperty__Type = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty QQmlProperty__Type = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty())
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty2(core.PointerFromQObject(obj)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty3(core.PointerFromQObject(obj), PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty4(core.PointerFromQObject(obj), PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty5(core.PointerFromQObject(obj), nameC))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty6(core.PointerFromQObject(obj), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty7(core.PointerFromQObject(obj), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty8(PointerFromQQmlProperty(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QQmlProperty_ConnectNotifySignal(ptr.Pointer(), core.PointerFromQObject(dest), slotC) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal2(ptr.Pointer(), core.PointerFromQObject(dest), C.int(int32(method))) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_HasNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Index() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_Index(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) IsDesignable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsDesignable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsResettable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsSignalProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Method() *core.QMetaMethod {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQMetaMethodFromPointer(C.QQmlProperty_Method(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QMetaMethod).DestroyQMetaMethod)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_NeedsNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Object() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlProperty_Object(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) PropertyType() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_PropertyType(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlProperty_PropertyTypeName(ptr.Pointer()))
	}
	return ""
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read2(object core.QObject_ITF, name string) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read() *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_Read(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) Reset() bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value)) != 0
}

func (ptr *QQmlProperty) Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value)) != 0
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func (ptr *QQmlProperty) Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlProperty_Write(ptr.Pointer(), core.PointerFromQVariant(value)) != 0
	}
	return false
}

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func (ptr *QQmlPropertyMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlPropertyMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) *QQmlPropertyMap {
	var n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	return n
}
func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {
	var tmpValue = NewQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_Clear(ptr.Pointer(), keyC)
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QQmlPropertyMap_Contains(ptr.Pointer(), keyC) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) Insert(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_Insert(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlPropertyMap_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Size(ptr.Pointer())))
	}
	return 0
}

//export callbackQQmlPropertyMap_UpdateValue
func callbackQQmlPropertyMap_UpdateValue(ptr unsafe.Pointer, key C.struct_QtQml_PackedString, input unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::updateValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string, *core.QVariant) *core.QVariant)(cGoUnpackString(key), core.NewQVariantFromPointer(input)))
	}

	return core.PointerFromQVariant(NewQQmlPropertyMapFromPointer(ptr).UpdateValueDefault(cGoUnpackString(key), core.NewQVariantFromPointer(input)))
}

func (ptr *QQmlPropertyMap) ConnectUpdateValue(f func(key string, input *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::updateValue", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectUpdateValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::updateValue")
	}
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValue(ptr.Pointer(), keyC, core.PointerFromQVariant(input)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) UpdateValueDefault(key string, input core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValueDefault(ptr.Pointer(), keyC, core.PointerFromQVariant(input)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QQmlPropertyMap_Value(ptr.Pointer(), keyC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQQmlPropertyMap_ValueChanged
func callbackQQmlPropertyMap_ValueChanged(ptr unsafe.Pointer, key C.struct_QtQml_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::valueChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(cGoUnpackString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::valueChanged")
	}
}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_ValueChanged(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

//export callbackQQmlPropertyMap_DestroyQQmlPropertyMap
func callbackQQmlPropertyMap_DestroyQQmlPropertyMap(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::~QQmlPropertyMap"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DestroyQQmlPropertyMapDefault()
	}
}

func (ptr *QQmlPropertyMap) ConnectDestroyQQmlPropertyMap(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::~QQmlPropertyMap", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectDestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::~QQmlPropertyMap")
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMapDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMapDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlPropertyMap___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __children_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlPropertyMap___children_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QQmlPropertyMap___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlPropertyMap___dynamicPropertyNames_newList(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlPropertyMap___findChildren_newList2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList3() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlPropertyMap___findChildren_newList3(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QQmlPropertyMap___findChildren_newList(ptr.Pointer()))
	}
	return nil
}

//export callbackQQmlPropertyMap_TimerEvent
func callbackQQmlPropertyMap_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::timerEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::timerEvent")
	}
}

func (ptr *QQmlPropertyMap) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlPropertyMap_ChildEvent
func callbackQQmlPropertyMap_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::childEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::childEvent")
	}
}

func (ptr *QQmlPropertyMap) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlPropertyMap_ConnectNotify
func callbackQQmlPropertyMap_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::connectNotify", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::connectNotify")
	}
}

func (ptr *QQmlPropertyMap) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_CustomEvent
func callbackQQmlPropertyMap_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::customEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::customEvent")
	}
}

func (ptr *QQmlPropertyMap) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlPropertyMap_DeleteLater
func callbackQQmlPropertyMap_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlPropertyMap) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::deleteLater", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::deleteLater")
	}
}

func (ptr *QQmlPropertyMap) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlPropertyMap_DisconnectNotify
func callbackQQmlPropertyMap_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::disconnectNotify", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::disconnectNotify")
	}
}

func (ptr *QQmlPropertyMap) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlPropertyMap) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_Event
func callbackQQmlPropertyMap_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlPropertyMap) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::event", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::event")
	}
}

func (ptr *QQmlPropertyMap) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_EventFilter
func callbackQQmlPropertyMap_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlPropertyMap) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::eventFilter", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::eventFilter")
	}
}

func (ptr *QQmlPropertyMap) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_MetaObject
func callbackQQmlPropertyMap_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlPropertyMapFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlPropertyMap) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::metaObject", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::metaObject")
	}
}

func (ptr *QQmlPropertyMap) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlPropertyMap_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlPropertyMap_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQmlPropertyValueSource struct {
	ptr unsafe.Pointer
}

type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlPropertyValueSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSource_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) *QQmlPropertyValueSource {
	var n = new(QQmlPropertyValueSource)
	n.SetPointer(ptr)
	return n
}
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {
	return NewQQmlPropertyValueSourceFromPointer(C.QQmlPropertyValueSource_NewQQmlPropertyValueSource())
}

//export callbackQQmlPropertyValueSource_SetTarget
func callbackQQmlPropertyValueSource_SetTarget(ptr unsafe.Pointer, property unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyValueSource::setTarget"); signal != nil {
		signal.(func(*QQmlProperty))(NewQQmlPropertyFromPointer(property))
	}

}

func (ptr *QQmlPropertyValueSource) ConnectSetTarget(f func(property *QQmlProperty)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::setTarget", f)
	}
}

func (ptr *QQmlPropertyValueSource) DisconnectSetTarget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::setTarget")
	}
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

//export callbackQQmlPropertyValueSource_DestroyQQmlPropertyValueSource
func callbackQQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyValueSource::~QQmlPropertyValueSource"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyValueSourceFromPointer(ptr).DestroyQQmlPropertyValueSourceDefault()
	}
}

func (ptr *QQmlPropertyValueSource) ConnectDestroyQQmlPropertyValueSource(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::~QQmlPropertyValueSource", f)
	}
}

func (ptr *QQmlPropertyValueSource) DisconnectDestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::~QQmlPropertyValueSource")
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSourceDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSourceDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QQmlScriptString struct {
	ptr unsafe.Pointer
}

type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return ptr
}

func (ptr *QQmlScriptString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlScriptString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlScriptString(ptr QQmlScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptString_PTR().Pointer()
	}
	return nil
}

func NewQQmlScriptStringFromPointer(ptr unsafe.Pointer) *QQmlScriptString {
	var n = new(QQmlScriptString)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlScriptString) DestroyQQmlScriptString() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQQmlScriptString() *QQmlScriptString {
	var tmpValue = NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString())
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {
	var tmpValue = NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString2(PointerFromQQmlScriptString(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func (ptr *QQmlScriptString) BooleanLiteral(ok bool) bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_BooleanLiteral(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsNullLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsUndefinedLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) NumberLiteral(ok bool) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQmlScriptString_NumberLiteral(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))))
	}
	return 0
}

func (ptr *QQmlScriptString) StringLiteral() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlScriptString_StringLiteral(ptr.Pointer()))
	}
	return ""
}

type QV4 struct {
	ptr unsafe.Pointer
}

type QV4_ITF interface {
	QV4_PTR() *QV4
}

func (ptr *QV4) QV4_PTR() *QV4 {
	return ptr
}

func (ptr *QV4) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QV4) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQV4(ptr QV4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4_PTR().Pointer()
	}
	return nil
}

func NewQV4FromPointer(ptr unsafe.Pointer) *QV4 {
	var n = new(QV4)
	n.SetPointer(ptr)
	return n
}

func (ptr *QV4) DestroyQV4() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QV4__PropertyFlag
//QV4::PropertyFlag
type QV4__PropertyFlag int64

const (
	QV4__Attr_Data            QV4__PropertyFlag = QV4__PropertyFlag(0)
	QV4__Attr_Accessor        QV4__PropertyFlag = QV4__PropertyFlag(0x1)
	QV4__Attr_NotWritable     QV4__PropertyFlag = QV4__PropertyFlag(0x2)
	QV4__Attr_NotEnumerable   QV4__PropertyFlag = QV4__PropertyFlag(0x4)
	QV4__Attr_NotConfigurable QV4__PropertyFlag = QV4__PropertyFlag(0x8)
	QV4__Attr_ReadOnly        QV4__PropertyFlag = QV4__PropertyFlag(QV4__Attr_NotWritable | QV4__Attr_NotEnumerable | QV4__Attr_NotConfigurable)
	QV4__Attr_Invalid         QV4__PropertyFlag = QV4__PropertyFlag(0xff)
)

//go:generate stringer -type=QV4__TypeHint
//QV4::TypeHint
type QV4__TypeHint int64

const (
	QV4__PREFERREDTYPE_HINT QV4__TypeHint = QV4__TypeHint(0)
	QV4__NUMBER_HINT        QV4__TypeHint = QV4__TypeHint(1)
	QV4__STRING_HINT        QV4__TypeHint = QV4__TypeHint(2)
)
