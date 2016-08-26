// +build !minimal

package qml

//#include <stdint.h>
//#include <stdlib.h>
//#include "qml.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

//QJSEngine::Extension
type QJSEngine__Extension int64

const (
	QJSEngine__TranslationExtension       = QJSEngine__Extension(0x1)
	QJSEngine__ConsoleExtension           = QJSEngine__Extension(0x2)
	QJSEngine__GarbageCollectionExtension = QJSEngine__Extension(0x4)
	QJSEngine__AllExtensions              = QJSEngine__Extension(0xffffffff)
)

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (p *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return p
}

func (p *QJSEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QJSEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
func NewQJSEngine() *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	var tmpValue = NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	defer qt.Recovering("QJSEngine::QJSEngine")

	var tmpValue = NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QJSEngine) CollectGarbage() {
	defer qt.Recovering("QJSEngine::collectGarbage")

	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	defer qt.Recovering("QJSEngine::evaluate")

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
	defer qt.Recovering("QJSEngine::globalObject")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) InstallExtensions(extensions QJSEngine__Extension, object QJSValue_ITF) {
	defer qt.Recovering("QJSEngine::installExtensions")

	if ptr.Pointer() != nil {
		C.QJSEngine_InstallExtensions(ptr.Pointer(), C.longlong(extensions), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) NewArray(length uint) *QJSValue {
	defer qt.Recovering("QJSEngine::newArray")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewArray(ptr.Pointer(), C.uint(uint32(length))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	defer qt.Recovering("QJSEngine::newObject")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	defer qt.Recovering("QJSEngine::newQObject")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	defer qt.Recovering("QJSEngine::~QJSEngine")

	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQJSEngine_TimerEvent
func callbackQJSEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::timerEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QJSEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::timerEvent")
	}
}

func (ptr *QJSEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QJSEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQJSEngine_ChildEvent
func callbackQJSEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::childEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QJSEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::childEvent")
	}
}

func (ptr *QJSEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QJSEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQJSEngine_ConnectNotify
func callbackQJSEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QJSEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::connectNotify", f)
	}
}

func (ptr *QJSEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QJSEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::connectNotify")
	}
}

func (ptr *QJSEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QJSEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QJSEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QJSEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_CustomEvent
func callbackQJSEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::customEvent", f)
	}
}

func (ptr *QJSEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QJSEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::customEvent")
	}
}

func (ptr *QJSEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QJSEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQJSEngine_DeleteLater
func callbackQJSEngine_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQJSEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QJSEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QJSEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::deleteLater", f)
	}
}

func (ptr *QJSEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QJSEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::deleteLater")
	}
}

func (ptr *QJSEngine) DeleteLater() {
	defer qt.Recovering("QJSEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QJSEngine) DeleteLaterDefault() {
	defer qt.Recovering("QJSEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQJSEngine_DisconnectNotify
func callbackQJSEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QJSEngine::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QJSEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::disconnectNotify", f)
	}
}

func (ptr *QJSEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QJSEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::disconnectNotify")
	}
}

func (ptr *QJSEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QJSEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QJSEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QJSEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_Event
func callbackQJSEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QJSEngine::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QJSEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QJSEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::event", f)
	}
}

func (ptr *QJSEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QJSEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::event")
	}
}

func (ptr *QJSEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QJSEngine::event")

	if ptr.Pointer() != nil {
		return C.QJSEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QJSEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QJSEngine::event")

	if ptr.Pointer() != nil {
		return C.QJSEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQJSEngine_EventFilter
func callbackQJSEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QJSEngine::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QJSEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QJSEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::eventFilter", f)
	}
}

func (ptr *QJSEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QJSEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::eventFilter")
	}
}

func (ptr *QJSEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QJSEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QJSEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QJSEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QJSEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QJSEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQJSEngine_MetaObject
func callbackQJSEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QJSEngine::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QJSEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQJSEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QJSEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QJSEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::metaObject", f)
	}
}

func (ptr *QJSEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QJSEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QJSEngine::metaObject")
	}
}

func (ptr *QJSEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QJSEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QJSEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue = QJSValue__SpecialValue(1)
)

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (p *QJSValue) QJSValue_PTR() *QJSValue {
	return p
}

func (p *QJSValue) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QJSValue) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue(C.longlong(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue4(value bool) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.char(int8(qt.GoBoolToInt(value)))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue8(value string) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var valueC = C.CString(value)
	defer C.free(unsafe.Pointer(valueC))
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue8(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue10(value string) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var valueC = C.CString(value)
	defer C.free(unsafe.Pointer(valueC))
	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue10(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue7(value float64) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue7(C.double(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue5(value int) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(int32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue6(value uint) *QJSValue {
	defer qt.Recovering("QJSValue::QJSValue")

	var tmpValue = NewQJSValueFromPointer(C.QJSValue_NewQJSValue6(C.uint(uint32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	defer qt.Recovering("QJSValue::deleteProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_DeleteProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QJSValue) Equals(other QJSValue_ITF) bool {
	defer qt.Recovering("QJSValue::equals")

	if ptr.Pointer() != nil {
		return C.QJSValue_Equals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {
	defer qt.Recovering("QJSValue::hasOwnProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_HasOwnProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	defer qt.Recovering("QJSValue::hasProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QJSValue_HasProperty(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QJSValue) IsArray() bool {
	defer qt.Recovering("QJSValue::isArray")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsBool() bool {
	defer qt.Recovering("QJSValue::isBool")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsCallable() bool {
	defer qt.Recovering("QJSValue::isCallable")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsCallable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsDate() bool {
	defer qt.Recovering("QJSValue::isDate")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsError() bool {
	defer qt.Recovering("QJSValue::isError")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNull() bool {
	defer qt.Recovering("QJSValue::isNull")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsNumber() bool {
	defer qt.Recovering("QJSValue::isNumber")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsObject() bool {
	defer qt.Recovering("QJSValue::isObject")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsQObject() bool {
	defer qt.Recovering("QJSValue::isQObject")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsRegExp() bool {
	defer qt.Recovering("QJSValue::isRegExp")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsString() bool {
	defer qt.Recovering("QJSValue::isString")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsUndefined() bool {
	defer qt.Recovering("QJSValue::isUndefined")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) IsVariant() bool {
	defer qt.Recovering("QJSValue::isVariant")

	if ptr.Pointer() != nil {
		return C.QJSValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) Property(name string) *QJSValue {
	defer qt.Recovering("QJSValue::property")

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
	defer qt.Recovering("QJSValue::property")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Prototype() *QJSValue {
	defer qt.Recovering("QJSValue::prototype")

	if ptr.Pointer() != nil {
		var tmpValue = NewQJSValueFromPointer(C.QJSValue_Prototype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	defer qt.Recovering("QJSValue::setProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QJSValue_SetProperty(ptr.Pointer(), nameC, PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetProperty2(arrayIndex uint, value QJSValue_ITF) {
	defer qt.Recovering("QJSValue::setProperty")

	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty2(ptr.Pointer(), C.uint(uint32(arrayIndex)), PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	defer qt.Recovering("QJSValue::setPrototype")

	if ptr.Pointer() != nil {
		C.QJSValue_SetPrototype(ptr.Pointer(), PointerFromQJSValue(prototype))
	}
}

func (ptr *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	defer qt.Recovering("QJSValue::strictlyEquals")

	if ptr.Pointer() != nil {
		return C.QJSValue_StrictlyEquals(ptr.Pointer(), PointerFromQJSValue(other)) != 0
	}
	return false
}

func (ptr *QJSValue) ToBool() bool {
	defer qt.Recovering("QJSValue::toBool")

	if ptr.Pointer() != nil {
		return C.QJSValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJSValue) ToDateTime() *core.QDateTime {
	defer qt.Recovering("QJSValue::toDateTime")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QJSValue_ToDateTime(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ToInt() int {
	defer qt.Recovering("QJSValue::toInt")

	if ptr.Pointer() != nil {
		return int(int32(C.QJSValue_ToInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToNumber() float64 {
	defer qt.Recovering("QJSValue::toNumber")

	if ptr.Pointer() != nil {
		return float64(C.QJSValue_ToNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	defer qt.Recovering("QJSValue::toQObject")

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
	defer qt.Recovering("QJSValue::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QJSValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QJSValue) ToUInt() uint {
	defer qt.Recovering("QJSValue::toUInt")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QJSValue_ToUInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	defer qt.Recovering("QJSValue::toVariant")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) DestroyQJSValue() {
	defer qt.Recovering("QJSValue::~QJSValue")

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

func (p *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return p
}

func (p *QJSValueIterator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QJSValueIterator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (p *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return p
}

func (p *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlAbstractUrlInterceptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::QQmlAbstractUrlInterceptor")

	return NewQQmlAbstractUrlInterceptorFromPointer(C.QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor())
}

//export callbackQQmlAbstractUrlInterceptor_Intercept
func callbackQQmlAbstractUrlInterceptor_Intercept(ptr unsafe.Pointer, url unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QQmlAbstractUrlInterceptor::intercept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlAbstractUrlInterceptor::intercept"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, QQmlAbstractUrlInterceptor__DataType) *core.QUrl)(core.NewQUrlFromPointer(url), QQmlAbstractUrlInterceptor__DataType(ty)))
	}

	return core.PointerFromQUrl(nil)
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectIntercept(f func(url *core.QUrl, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl) {
	defer qt.Recovering("connect QQmlAbstractUrlInterceptor::intercept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::intercept", f)
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectIntercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) {
	defer qt.Recovering("disconnect QQmlAbstractUrlInterceptor::intercept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlAbstractUrlInterceptor::intercept")
	}
}

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::intercept")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlAbstractUrlInterceptor_Intercept(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	defer qt.Recovering("QQmlAbstractUrlInterceptor::~QQmlAbstractUrlInterceptor")

	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr.Pointer())
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

func (p *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return p
}

func (p *QQmlApplicationEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func (p *QQmlApplicationEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQmlEngine_PTR().SetPointer(ptr)
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
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	var filePathC = C.CString(filePath)
	defer C.free(unsafe.Pointer(filePathC))
	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(filePathC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	var tmpValue = NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlApplicationEngine_Load2
func callbackQQmlApplicationEngine_Load2(ptr unsafe.Pointer, filePath *C.char) {
	defer qt.Recovering("callback QQmlApplicationEngine::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::load2"); signal != nil {
		signal.(func(string))(C.GoString(filePath))
	}

}

func (ptr *QQmlApplicationEngine) ConnectLoad2(f func(filePath string)) {
	defer qt.Recovering("connect QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load2", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad2(filePath string) {
	defer qt.Recovering("disconnect QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load2")
	}
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		var filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
		C.QQmlApplicationEngine_Load2(ptr.Pointer(), filePathC)
	}
}

//export callbackQQmlApplicationEngine_Load
func callbackQQmlApplicationEngine_Load(ptr unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::load"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ConnectLoad(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::load")
	}
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_LoadData
func callbackQQmlApplicationEngine_LoadData(ptr unsafe.Pointer, data *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::loadData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::loadData"); signal != nil {
		signal.(func(string, *core.QUrl))(qt.HexDecodeToString(C.GoString(data)), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ConnectLoadData(f func(data string, url *core.QUrl)) {
	defer qt.Recovering("connect QQmlApplicationEngine::loadData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::loadData", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoadData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQmlApplicationEngine::loadData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::loadData")
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::loadData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		C.QQmlApplicationEngine_LoadData(ptr.Pointer(), dataC, core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_ObjectCreated
func callbackQQmlApplicationEngine_ObjectCreated(ptr unsafe.Pointer, object unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::objectCreated")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::objectCreated"); signal != nil {
		signal.(func(*core.QObject, *core.QUrl))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {
	defer qt.Recovering("connect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectObjectCreated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::objectCreated", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::objectCreated")
	}
}

func (ptr *QQmlApplicationEngine) ObjectCreated(object core.QObject_ITF, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ObjectCreated(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	defer qt.Recovering("QQmlApplicationEngine::~QQmlApplicationEngine")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlApplicationEngine_Event
func callbackQQmlApplicationEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlApplicationEngine::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlApplicationEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlApplicationEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlApplicationEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::event", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::event")
	}
}

func (ptr *QQmlApplicationEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlApplicationEngine::event")

	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlApplicationEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlApplicationEngine::event")

	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlApplicationEngine_TimerEvent
func callbackQQmlApplicationEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::timerEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::timerEvent")
	}
}

func (ptr *QQmlApplicationEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlApplicationEngine_ChildEvent
func callbackQQmlApplicationEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::childEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::childEvent")
	}
}

func (ptr *QQmlApplicationEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlApplicationEngine_ConnectNotify
func callbackQQmlApplicationEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlApplicationEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::connectNotify", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::connectNotify")
	}
}

func (ptr *QQmlApplicationEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlApplicationEngine_CustomEvent
func callbackQQmlApplicationEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlApplicationEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::customEvent", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::customEvent")
	}
}

func (ptr *QQmlApplicationEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlApplicationEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlApplicationEngine_DeleteLater
func callbackQQmlApplicationEngine_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlApplicationEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlApplicationEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::deleteLater", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::deleteLater")
	}
}

func (ptr *QQmlApplicationEngine) DeleteLater() {
	defer qt.Recovering("QQmlApplicationEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlApplicationEngine) DeleteLaterDefault() {
	defer qt.Recovering("QQmlApplicationEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlApplicationEngine_DisconnectNotify
func callbackQQmlApplicationEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlApplicationEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlApplicationEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::disconnectNotify", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::disconnectNotify")
	}
}

func (ptr *QQmlApplicationEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlApplicationEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlApplicationEngine_EventFilter
func callbackQQmlApplicationEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlApplicationEngine::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlApplicationEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlApplicationEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlApplicationEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::eventFilter", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::eventFilter")
	}
}

func (ptr *QQmlApplicationEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlApplicationEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlApplicationEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlApplicationEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlApplicationEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlApplicationEngine_MetaObject
func callbackQQmlApplicationEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlApplicationEngine::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlApplicationEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlApplicationEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlApplicationEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlApplicationEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::metaObject", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlApplicationEngine::metaObject")
	}
}

func (ptr *QQmlApplicationEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlApplicationEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlApplicationEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlApplicationEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlApplicationEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlApplicationEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      = QQmlComponent__CompilationMode(1)
)

//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    = QQmlComponent__Status(0)
	QQmlComponent__Ready   = QQmlComponent__Status(1)
	QQmlComponent__Loading = QQmlComponent__Status(2)
	QQmlComponent__Error   = QQmlComponent__Status(3)
)

type QQmlComponent struct {
	core.QObject
}

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func (p *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return p
}

func (p *QQmlComponent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlComponent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
func (ptr *QQmlComponent) Progress() float64 {
	defer qt.Recovering("QQmlComponent::progress")

	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	defer qt.Recovering("QQmlComponent::status")

	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) Url() *core.QUrl {
	defer qt.Recovering("QQmlComponent::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlComponent_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), fileNameC, C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), fileNameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	defer qt.Recovering("QQmlComponent::QQmlComponent")

	var tmpValue = NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlComponent_BeginCreate
func callbackQQmlComponent_BeginCreate(ptr unsafe.Pointer, publicContext unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlComponent::beginCreate")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::beginCreate"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(publicContext)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).BeginCreateDefault(NewQQmlContextFromPointer(publicContext)))
}

func (ptr *QQmlComponent) ConnectBeginCreate(f func(publicContext *QQmlContext) *core.QObject) {
	defer qt.Recovering("connect QQmlComponent::beginCreate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::beginCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectBeginCreate() {
	defer qt.Recovering("disconnect QQmlComponent::beginCreate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::beginCreate")
	}
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::beginCreate")

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
	defer qt.Recovering("QQmlComponent::beginCreate")

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
	defer qt.Recovering("callback QQmlComponent::completeCreate")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::completeCreate"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).CompleteCreateDefault()
	}
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {
	defer qt.Recovering("connect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::completeCreate", f)
	}
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {
	defer qt.Recovering("disconnect QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::completeCreate")
	}
}

func (ptr *QQmlComponent) CompleteCreate() {
	defer qt.Recovering("QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) CompleteCreateDefault() {
	defer qt.Recovering("QQmlComponent::completeCreate")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreateDefault(ptr.Pointer())
	}
}

//export callbackQQmlComponent_Create
func callbackQQmlComponent_Create(ptr unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlComponent::create")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::create"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(context)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).CreateDefault(NewQQmlContextFromPointer(context)))
}

func (ptr *QQmlComponent) ConnectCreate(f func(context *QQmlContext) *core.QObject) {
	defer qt.Recovering("connect QQmlComponent::create")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::create", f)
	}
}

func (ptr *QQmlComponent) DisconnectCreate() {
	defer qt.Recovering("disconnect QQmlComponent::create")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::create")
	}
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	defer qt.Recovering("QQmlComponent::create")

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
	defer qt.Recovering("QQmlComponent::create")

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
	defer qt.Recovering("QQmlComponent::create")

	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	defer qt.Recovering("QQmlComponent::creationContext")

	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	defer qt.Recovering("QQmlComponent::isError")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	defer qt.Recovering("QQmlComponent::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	defer qt.Recovering("QQmlComponent::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	defer qt.Recovering("QQmlComponent::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_IsReady(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQmlComponent_LoadUrl
func callbackQQmlComponent_LoadUrl(ptr unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::loadUrl")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::loadUrl"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlComponent) ConnectLoadUrl(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl", f)
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl")
	}
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_LoadUrl2
func callbackQQmlComponent_LoadUrl2(ptr unsafe.Pointer, url unsafe.Pointer, mode C.longlong) {
	defer qt.Recovering("callback QQmlComponent::loadUrl")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::loadUrl2"); signal != nil {
		signal.(func(*core.QUrl, QQmlComponent__CompilationMode))(core.NewQUrlFromPointer(url), QQmlComponent__CompilationMode(mode))
	}

}

func (ptr *QQmlComponent) ConnectLoadUrl2(f func(url *core.QUrl, mode QQmlComponent__CompilationMode)) {
	defer qt.Recovering("connect QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl2", f)
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	defer qt.Recovering("disconnect QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::loadUrl2")
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	defer qt.Recovering("QQmlComponent::loadUrl")

	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(mode))
	}
}

//export callbackQQmlComponent_ProgressChanged
func callbackQQmlComponent_ProgressChanged(ptr unsafe.Pointer, progress C.double) {
	defer qt.Recovering("callback QQmlComponent::progressChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::progressChanged"); signal != nil {
		signal.(func(float64))(float64(progress))
	}

}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {
	defer qt.Recovering("connect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectProgressChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::progressChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {
	defer qt.Recovering("disconnect QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectProgressChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::progressChanged")
	}
}

func (ptr *QQmlComponent) ProgressChanged(progress float64) {
	defer qt.Recovering("QQmlComponent::progressChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ProgressChanged(ptr.Pointer(), C.double(progress))
	}
}

//export callbackQQmlComponent_SetData
func callbackQQmlComponent_SetData(ptr unsafe.Pointer, data *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::setData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::setData"); signal != nil {
		signal.(func(string, *core.QUrl))(qt.HexDecodeToString(C.GoString(data)), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlComponent) ConnectSetData(f func(data string, url *core.QUrl)) {
	defer qt.Recovering("connect QQmlComponent::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::setData", f)
	}
}

func (ptr *QQmlComponent) DisconnectSetData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQmlComponent::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::setData")
	}
}

func (ptr *QQmlComponent) SetData(data string, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlComponent::setData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		C.QQmlComponent_SetData(ptr.Pointer(), dataC, core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_StatusChanged
func callbackQQmlComponent_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	defer qt.Recovering("callback QQmlComponent::statusChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::statusChanged"); signal != nil {
		signal.(func(QQmlComponent__Status))(QQmlComponent__Status(status))
	}

}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	defer qt.Recovering("connect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::statusChanged", f)
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::statusChanged")
	}
}

func (ptr *QQmlComponent) StatusChanged(status QQmlComponent__Status) {
	defer qt.Recovering("QQmlComponent::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlComponent_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	defer qt.Recovering("QQmlComponent::~QQmlComponent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlComponent_TimerEvent
func callbackQQmlComponent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::timerEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::timerEvent")
	}
}

func (ptr *QQmlComponent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlComponent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlComponent::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlComponent_ChildEvent
func callbackQQmlComponent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::childEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlComponent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::childEvent")
	}
}

func (ptr *QQmlComponent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlComponent::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlComponent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlComponent::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlComponent_ConnectNotify
func callbackQQmlComponent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlComponent::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::connectNotify", f)
	}
}

func (ptr *QQmlComponent) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlComponent::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::connectNotify")
	}
}

func (ptr *QQmlComponent) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlComponent::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlComponent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlComponent::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_CustomEvent
func callbackQQmlComponent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::customEvent", f)
	}
}

func (ptr *QQmlComponent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlComponent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::customEvent")
	}
}

func (ptr *QQmlComponent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlComponent::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlComponent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlComponent::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlComponent_DeleteLater
func callbackQQmlComponent_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlComponent) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlComponent::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::deleteLater", f)
	}
}

func (ptr *QQmlComponent) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlComponent::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::deleteLater")
	}
}

func (ptr *QQmlComponent) DeleteLater() {
	defer qt.Recovering("QQmlComponent::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlComponent) DeleteLaterDefault() {
	defer qt.Recovering("QQmlComponent::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlComponent_DisconnectNotify
func callbackQQmlComponent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlComponent::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlComponent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::disconnectNotify", f)
	}
}

func (ptr *QQmlComponent) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlComponent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::disconnectNotify")
	}
}

func (ptr *QQmlComponent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlComponent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlComponent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlComponent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_Event
func callbackQQmlComponent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlComponent::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlComponent) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlComponent::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::event", f)
	}
}

func (ptr *QQmlComponent) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlComponent::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::event")
	}
}

func (ptr *QQmlComponent) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlComponent::event")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlComponent) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlComponent::event")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlComponent_EventFilter
func callbackQQmlComponent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlComponent::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlComponent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlComponent::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::eventFilter", f)
	}
}

func (ptr *QQmlComponent) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlComponent::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::eventFilter")
	}
}

func (ptr *QQmlComponent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlComponent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlComponent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlComponent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlComponent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlComponent_MetaObject
func callbackQQmlComponent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlComponent::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlComponent::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlComponentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlComponent) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlComponent::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::metaObject", f)
	}
}

func (ptr *QQmlComponent) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlComponent::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlComponent::metaObject")
	}
}

func (ptr *QQmlComponent) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlComponent::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlComponent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlComponent::metaObject")

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

func (p *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return p
}

func (p *QQmlContext) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlContext) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
	defer qt.Recovering("QQmlContext::QQmlContext")

	var tmpValue = NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlContext::QQmlContext")

	var tmpValue = NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlContext::baseUrl")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlContext_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	defer qt.Recovering("QQmlContext::contextObject")

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
	defer qt.Recovering("QQmlContext::contextProperty")

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
	defer qt.Recovering("QQmlContext::engine")

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
	defer qt.Recovering("QQmlContext::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlContext_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	defer qt.Recovering("QQmlContext::nameForObject")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	defer qt.Recovering("QQmlContext::parentContext")

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
	defer qt.Recovering("QQmlContext::resolvedUrl")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlContext_ResolvedUrl(ptr.Pointer(), core.PointerFromQUrl(src)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QQmlContext::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextObject")

	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QQmlContext_SetContextProperty(ptr.Pointer(), nameC, core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlContext::setContextProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), nameC, core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	defer qt.Recovering("QQmlContext::~QQmlContext")

	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlContext_TimerEvent
func callbackQQmlContext_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::timerEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlContext::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::timerEvent")
	}
}

func (ptr *QQmlContext) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlContext::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlContext) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlContext::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlContext_ChildEvent
func callbackQQmlContext_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::childEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlContext::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::childEvent")
	}
}

func (ptr *QQmlContext) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlContext::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlContext) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlContext::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlContext_ConnectNotify
func callbackQQmlContext_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlContext::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::connectNotify", f)
	}
}

func (ptr *QQmlContext) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlContext::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::connectNotify")
	}
}

func (ptr *QQmlContext) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlContext::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlContext_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlContext::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlContext_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_CustomEvent
func callbackQQmlContext_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::customEvent", f)
	}
}

func (ptr *QQmlContext) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlContext::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::customEvent")
	}
}

func (ptr *QQmlContext) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlContext::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlContext) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlContext::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlContext_DeleteLater
func callbackQQmlContext_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlContextFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlContext) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlContext::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::deleteLater", f)
	}
}

func (ptr *QQmlContext) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlContext::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::deleteLater")
	}
}

func (ptr *QQmlContext) DeleteLater() {
	defer qt.Recovering("QQmlContext::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlContext_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlContext) DeleteLaterDefault() {
	defer qt.Recovering("QQmlContext::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlContext_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlContext_DisconnectNotify
func callbackQQmlContext_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlContext::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlContext::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::disconnectNotify", f)
	}
}

func (ptr *QQmlContext) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlContext::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::disconnectNotify")
	}
}

func (ptr *QQmlContext) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlContext::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlContext_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlContext::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlContext_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_Event
func callbackQQmlContext_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlContext::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlContext) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlContext::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::event", f)
	}
}

func (ptr *QQmlContext) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlContext::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::event")
	}
}

func (ptr *QQmlContext) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlContext::event")

	if ptr.Pointer() != nil {
		return C.QQmlContext_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlContext) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlContext::event")

	if ptr.Pointer() != nil {
		return C.QQmlContext_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlContext_EventFilter
func callbackQQmlContext_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlContext::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlContext) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlContext::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::eventFilter", f)
	}
}

func (ptr *QQmlContext) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlContext::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::eventFilter")
	}
}

func (ptr *QQmlContext) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlContext::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlContext_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlContext::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlContext_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlContext_MetaObject
func callbackQQmlContext_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlContext::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlContext::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlContextFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlContext) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlContext::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::metaObject", f)
	}
}

func (ptr *QQmlContext) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlContext::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlContext::metaObject")
	}
}

func (ptr *QQmlContext) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlContext::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlContext_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlContext::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlContext_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership = QQmlEngine__ObjectOwnership(1)
)

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (p *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return p
}

func (p *QQmlEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (p *QQmlEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QJSEngine_PTR().SetPointer(ptr)
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
func (ptr *QQmlEngine) OfflineStoragePath() string {
	defer qt.Recovering("QQmlEngine::offlineStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	defer qt.Recovering("QQmlEngine::setOfflineStoragePath")

	if ptr.Pointer() != nil {
		var dirC = C.CString(dir)
		defer C.free(unsafe.Pointer(dirC))
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), dirC)
	}
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	defer qt.Recovering("QQmlEngine::QQmlEngine")

	var tmpValue = NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	defer qt.Recovering("QQmlEngine::addImageProvider")

	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), providerIdC, PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	defer qt.Recovering("QQmlEngine::addImportPath")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQmlEngine_AddImportPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	defer qt.Recovering("QQmlEngine::addPluginPath")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QQmlEngine) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlEngine::baseUrl")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlEngine_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) ClearComponentCache() {
	defer qt.Recovering("QQmlEngine::clearComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlEngine::contextForObject")

	var tmpValue = NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) ContextForObject(object core.QObject_ITF) *QQmlContext {
	defer qt.Recovering("QQmlEngine::contextForObject")

	var tmpValue = NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlEngine_Event
func callbackQQmlEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlEngine::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::event", f)
	}
}

func (ptr *QQmlEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::event")
	}
}

func (ptr *QQmlEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlEngine::event")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlEngine::event")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	defer qt.Recovering("QQmlEngine::imageProvider")

	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), providerIdC))
	}
	return nil
}

func (ptr *QQmlEngine) ImportPathList() []string {
	defer qt.Recovering("QQmlEngine::importPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_ImportPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlEngine::incubationController")

	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlEngine::networkAccessManager")

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
	defer qt.Recovering("QQmlEngine::networkAccessManagerFactory")

	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	defer qt.Recovering("QQmlEngine::objectOwnership")

	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	defer qt.Recovering("QQmlEngine::objectOwnership")

	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	defer qt.Recovering("QQmlEngine::outputWarningsToStandardError")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlEngine) PluginPathList() []string {
	defer qt.Recovering("QQmlEngine::pluginPathList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlEngine_PluginPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQQmlEngine_Quit
func callbackQQmlEngine_Quit(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::quit")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::quit"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	defer qt.Recovering("connect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectQuit(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::quit", f)
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	defer qt.Recovering("disconnect QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::quit")
	}
}

func (ptr *QQmlEngine) Quit() {
	defer qt.Recovering("QQmlEngine::quit")

	if ptr.Pointer() != nil {
		C.QQmlEngine_Quit(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	defer qt.Recovering("QQmlEngine::removeImageProvider")

	if ptr.Pointer() != nil {
		var providerIdC = C.CString(providerId)
		defer C.free(unsafe.Pointer(providerIdC))
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), providerIdC)
	}
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	defer qt.Recovering("QQmlEngine::rootContext")

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
	defer qt.Recovering("QQmlEngine::setBaseUrl")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	defer qt.Recovering("QQmlEngine::setContextForObject")

	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	defer qt.Recovering("QQmlEngine::setContextForObject")

	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setImportPathList")

	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), pathsC)
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	defer qt.Recovering("QQmlEngine::setIncubationController")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	defer qt.Recovering("QQmlEngine::setNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	defer qt.Recovering("QQmlEngine::setObjectOwnership")

	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	defer qt.Recovering("QQmlEngine::setObjectOwnership")

	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	defer qt.Recovering("QQmlEngine::setOutputWarningsToStandardError")

	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	defer qt.Recovering("QQmlEngine::setPluginPathList")

	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), pathsC)
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	defer qt.Recovering("QQmlEngine::trimComponentCache")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	defer qt.Recovering("QQmlEngine::~QQmlEngine")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlEngine_TimerEvent
func callbackQQmlEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::timerEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::timerEvent")
	}
}

func (ptr *QQmlEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlEngine_ChildEvent
func callbackQQmlEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::childEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::childEvent")
	}
}

func (ptr *QQmlEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlEngine_ConnectNotify
func callbackQQmlEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::connectNotify", f)
	}
}

func (ptr *QQmlEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::connectNotify")
	}
}

func (ptr *QQmlEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlEngine_CustomEvent
func callbackQQmlEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::customEvent", f)
	}
}

func (ptr *QQmlEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::customEvent")
	}
}

func (ptr *QQmlEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlEngine_DeleteLater
func callbackQQmlEngine_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::deleteLater", f)
	}
}

func (ptr *QQmlEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::deleteLater")
	}
}

func (ptr *QQmlEngine) DeleteLater() {
	defer qt.Recovering("QQmlEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlEngine) DeleteLaterDefault() {
	defer qt.Recovering("QQmlEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlEngine_DisconnectNotify
func callbackQQmlEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlEngine::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::disconnectNotify", f)
	}
}

func (ptr *QQmlEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::disconnectNotify")
	}
}

func (ptr *QQmlEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlEngine_EventFilter
func callbackQQmlEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlEngine::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::eventFilter", f)
	}
}

func (ptr *QQmlEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::eventFilter")
	}
}

func (ptr *QQmlEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlEngine_MetaObject
func callbackQQmlEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlEngine::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlEngine::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::metaObject", f)
	}
}

func (ptr *QQmlEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlEngine::metaObject")
	}
}

func (ptr *QQmlEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlEngine::metaObject")

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

func (p *QQmlError) QQmlError_PTR() *QQmlError {
	return p
}

func (p *QQmlError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQQmlError() *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	var tmpValue = NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	var tmpValue = NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func (ptr *QQmlError) Column() int {
	defer qt.Recovering("QQmlError::column")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	defer qt.Recovering("QQmlError::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	defer qt.Recovering("QQmlError::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	defer qt.Recovering("QQmlError::line")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Line(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	defer qt.Recovering("QQmlError::object")

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
	defer qt.Recovering("QQmlError::setColumn")

	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	defer qt.Recovering("QQmlError::setDescription")

	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QQmlError_SetDescription(ptr.Pointer(), descriptionC)
	}
}

func (ptr *QQmlError) SetLine(line int) {
	defer qt.Recovering("QQmlError::setLine")

	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(int32(line)))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlError::setObject")

	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlError::setUrl")

	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) ToString() string {
	defer qt.Recovering("QQmlError::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) Url() *core.QUrl {
	defer qt.Recovering("QQmlError::url")

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

func (p *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return p
}

func (p *QQmlExpression) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlExpression) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	var expressionC = C.CString(expression)
	defer C.free(unsafe.Pointer(expressionC))
	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression2(PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), expressionC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression3(script QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {
	defer qt.Recovering("QQmlExpression::QQmlExpression")

	var tmpValue = NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression3(PointerFromQQmlScriptString(script), PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlExpression) ClearError() {
	defer qt.Recovering("QQmlExpression::clearError")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) ColumnNumber() int {
	defer qt.Recovering("QQmlExpression::columnNumber")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	defer qt.Recovering("QQmlExpression::context")

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
	defer qt.Recovering("QQmlExpression::engine")

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
	defer qt.Recovering("QQmlExpression::error")

	if ptr.Pointer() != nil {
		var tmpValue = NewQQmlErrorFromPointer(C.QQmlExpression_Error(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined bool) *core.QVariant {
	defer qt.Recovering("QQmlExpression::evaluate")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQmlExpression_Evaluate(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(valueIsUndefined)))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Expression() string {
	defer qt.Recovering("QQmlExpression::expression")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_Expression(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	defer qt.Recovering("QQmlExpression::hasError")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) LineNumber() int {
	defer qt.Recovering("QQmlExpression::lineNumber")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_LineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	defer qt.Recovering("QQmlExpression::notifyOnValueChanged")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_NotifyOnValueChanged(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	defer qt.Recovering("QQmlExpression::scopeObject")

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
	defer qt.Recovering("QQmlExpression::setExpression")

	if ptr.Pointer() != nil {
		var expressionC = C.CString(expression)
		defer C.free(unsafe.Pointer(expressionC))
		C.QQmlExpression_SetExpression(ptr.Pointer(), expressionC)
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	defer qt.Recovering("QQmlExpression::setNotifyOnValueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(notifyOnChange))))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	defer qt.Recovering("QQmlExpression::setSourceLocation")

	if ptr.Pointer() != nil {
		var urlC = C.CString(url)
		defer C.free(unsafe.Pointer(urlC))
		C.QQmlExpression_SetSourceLocation(ptr.Pointer(), urlC, C.int(int32(line)), C.int(int32(column)))
	}
}

func (ptr *QQmlExpression) SourceFile() string {
	defer qt.Recovering("QQmlExpression::sourceFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExpression_SourceFile(ptr.Pointer()))
	}
	return ""
}

//export callbackQQmlExpression_ValueChanged
func callbackQQmlExpression_ValueChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::valueChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::valueChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	defer qt.Recovering("connect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::valueChanged", f)
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::valueChanged")
	}
}

func (ptr *QQmlExpression) ValueChanged() {
	defer qt.Recovering("QQmlExpression::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ValueChanged(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	defer qt.Recovering("QQmlExpression::~QQmlExpression")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExpression_TimerEvent
func callbackQQmlExpression_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::timerEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::timerEvent")
	}
}

func (ptr *QQmlExpression) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExpression) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExpression::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlExpression_ChildEvent
func callbackQQmlExpression_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlExpression::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::childEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlExpression::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::childEvent")
	}
}

func (ptr *QQmlExpression) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExpression::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExpression) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExpression::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExpression_ConnectNotify
func callbackQQmlExpression_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlExpression::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::connectNotify", f)
	}
}

func (ptr *QQmlExpression) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlExpression::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::connectNotify")
	}
}

func (ptr *QQmlExpression) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExpression::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExpression) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExpression::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_CustomEvent
func callbackQQmlExpression_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlExpression::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::customEvent", f)
	}
}

func (ptr *QQmlExpression) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlExpression::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::customEvent")
	}
}

func (ptr *QQmlExpression) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExpression::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExpression) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExpression::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExpression_DeleteLater
func callbackQQmlExpression_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExpressionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExpression) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlExpression::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::deleteLater", f)
	}
}

func (ptr *QQmlExpression) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlExpression::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::deleteLater")
	}
}

func (ptr *QQmlExpression) DeleteLater() {
	defer qt.Recovering("QQmlExpression::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExpression) DeleteLaterDefault() {
	defer qt.Recovering("QQmlExpression::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExpression_DisconnectNotify
func callbackQQmlExpression_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExpression::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlExpression::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::disconnectNotify", f)
	}
}

func (ptr *QQmlExpression) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlExpression::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::disconnectNotify")
	}
}

func (ptr *QQmlExpression) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExpression::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExpression) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExpression::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_Event
func callbackQQmlExpression_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlExpression::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExpression) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlExpression::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::event", f)
	}
}

func (ptr *QQmlExpression) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlExpression::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::event")
	}
}

func (ptr *QQmlExpression) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExpression::event")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlExpression) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExpression::event")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlExpression_EventFilter
func callbackQQmlExpression_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlExpression::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExpression) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlExpression::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::eventFilter", f)
	}
}

func (ptr *QQmlExpression) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlExpression::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::eventFilter")
	}
}

func (ptr *QQmlExpression) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExpression::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlExpression) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExpression::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlExpression_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlExpression_MetaObject
func callbackQQmlExpression_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlExpression::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExpression::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExpressionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExpression) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlExpression::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::metaObject", f)
	}
}

func (ptr *QQmlExpression) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlExpression::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExpression::metaObject")
	}
}

func (ptr *QQmlExpression) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlExpression::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExpression_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlExpression::metaObject")

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

func (p *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return p
}

func (p *QQmlExtensionPlugin) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlExtensionPlugin) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQQmlExtensionPlugin_InitializeEngine
func callbackQQmlExtensionPlugin_InitializeEngine(ptr unsafe.Pointer, engine unsafe.Pointer, uri *C.char) {
	defer qt.Recovering("callback QQmlExtensionPlugin::initializeEngine")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::initializeEngine"); signal != nil {
		signal.(func(*QQmlEngine, string))(NewQQmlEngineFromPointer(engine), C.GoString(uri))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).InitializeEngineDefault(NewQQmlEngineFromPointer(engine), C.GoString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectInitializeEngine(f func(engine *QQmlEngine, uri string)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::initializeEngine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::initializeEngine", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectInitializeEngine() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::initializeEngine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::initializeEngine")
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {
	defer qt.Recovering("QQmlExtensionPlugin::initializeEngine")

	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_InitializeEngine(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngineDefault(engine QQmlEngine_ITF, uri string) {
	defer qt.Recovering("QQmlExtensionPlugin::initializeEngine")

	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_InitializeEngineDefault(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

func NewQQmlExtensionPlugin(parent core.QObject_ITF) *QQmlExtensionPlugin {
	defer qt.Recovering("QQmlExtensionPlugin::QQmlExtensionPlugin")

	var tmpValue = NewQQmlExtensionPluginFromPointer(C.QQmlExtensionPlugin_NewQQmlExtensionPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {
	defer qt.Recovering("QQmlExtensionPlugin::baseUrl")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QQmlExtensionPlugin_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQmlExtensionPlugin_RegisterTypes
func callbackQQmlExtensionPlugin_RegisterTypes(ptr unsafe.Pointer, uri *C.char) {
	defer qt.Recovering("callback QQmlExtensionPlugin::registerTypes")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::registerTypes"); signal != nil {
		signal.(func(string))(C.GoString(uri))
	}

}

func (ptr *QQmlExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::registerTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::registerTypes", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectRegisterTypes(uri string) {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::registerTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::registerTypes")
	}
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	defer qt.Recovering("QQmlExtensionPlugin::registerTypes")

	if ptr.Pointer() != nil {
		var uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), uriC)
	}
}

//export callbackQQmlExtensionPlugin_TimerEvent
func callbackQQmlExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::timerEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::timerEvent")
	}
}

func (ptr *QQmlExtensionPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_ChildEvent
func callbackQQmlExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::childEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::childEvent")
	}
}

func (ptr *QQmlExtensionPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_ConnectNotify
func callbackQQmlExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::connectNotify", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::connectNotify")
	}
}

func (ptr *QQmlExtensionPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_CustomEvent
func callbackQQmlExtensionPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::customEvent", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::customEvent")
	}
}

func (ptr *QQmlExtensionPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_DeleteLater
func callbackQQmlExtensionPlugin_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExtensionPlugin) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::deleteLater", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::deleteLater")
	}
}

func (ptr *QQmlExtensionPlugin) DeleteLater() {
	defer qt.Recovering("QQmlExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlExtensionPlugin) DeleteLaterDefault() {
	defer qt.Recovering("QQmlExtensionPlugin::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlExtensionPlugin_DisconnectNotify
func callbackQQmlExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlExtensionPlugin::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::disconnectNotify", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::disconnectNotify")
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlExtensionPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_Event
func callbackQQmlExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlExtensionPlugin::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExtensionPlugin) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlExtensionPlugin::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::event", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::event")
	}
}

func (ptr *QQmlExtensionPlugin) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExtensionPlugin::event")

	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExtensionPlugin::event")

	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_EventFilter
func callbackQQmlExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlExtensionPlugin::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExtensionPlugin) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::eventFilter", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::eventFilter")
	}
}

func (ptr *QQmlExtensionPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlExtensionPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_MetaObject
func callbackQQmlExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlExtensionPlugin::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlExtensionPlugin::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExtensionPlugin) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::metaObject", f)
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlExtensionPlugin::metaObject")
	}
}

func (ptr *QQmlExtensionPlugin) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlExtensionPlugin::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExtensionPlugin_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlExtensionPlugin::metaObject")

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

func (p *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return p
}

func (p *QQmlFileSelector) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlFileSelector) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
	defer qt.Recovering("QQmlFileSelector::QQmlFileSelector")

	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::get")

	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlFileSelector) Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::get")

	var tmpValue = NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlFileSelector) Selector() *core.QFileSelector {
	defer qt.Recovering("QQmlFileSelector::selector")

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
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		var strinC = C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), strinC)
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		var strinC = C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), strinC)
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	defer qt.Recovering("QQmlFileSelector::setSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	defer qt.Recovering("QQmlFileSelector::~QQmlFileSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlFileSelector_TimerEvent
func callbackQQmlFileSelector_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::timerEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::timerEvent")
	}
}

func (ptr *QQmlFileSelector) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlFileSelector) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlFileSelector_ChildEvent
func callbackQQmlFileSelector_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::childEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::childEvent")
	}
}

func (ptr *QQmlFileSelector) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlFileSelector) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlFileSelector_ConnectNotify
func callbackQQmlFileSelector_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlFileSelector::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::connectNotify", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlFileSelector::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::connectNotify")
	}
}

func (ptr *QQmlFileSelector) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlFileSelector::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlFileSelector::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_CustomEvent
func callbackQQmlFileSelector_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::customEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::customEvent")
	}
}

func (ptr *QQmlFileSelector) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlFileSelector) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlFileSelector_DeleteLater
func callbackQQmlFileSelector_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlFileSelector) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlFileSelector::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::deleteLater", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlFileSelector::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::deleteLater")
	}
}

func (ptr *QQmlFileSelector) DeleteLater() {
	defer qt.Recovering("QQmlFileSelector::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlFileSelector) DeleteLaterDefault() {
	defer qt.Recovering("QQmlFileSelector::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlFileSelector_DisconnectNotify
func callbackQQmlFileSelector_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlFileSelector::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlFileSelector::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::disconnectNotify", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlFileSelector::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::disconnectNotify")
	}
}

func (ptr *QQmlFileSelector) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlFileSelector::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlFileSelector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlFileSelector::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_Event
func callbackQQmlFileSelector_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlFileSelector::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlFileSelector) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlFileSelector::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::event", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::event")
	}
}

func (ptr *QQmlFileSelector) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlFileSelector::event")

	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlFileSelector) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlFileSelector::event")

	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlFileSelector_EventFilter
func callbackQQmlFileSelector_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlFileSelector::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlFileSelector) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlFileSelector::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::eventFilter", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlFileSelector::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::eventFilter")
	}
}

func (ptr *QQmlFileSelector) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlFileSelector::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlFileSelector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlFileSelector::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlFileSelector_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlFileSelector_MetaObject
func callbackQQmlFileSelector_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlFileSelector::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlFileSelector::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlFileSelectorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlFileSelector) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlFileSelector::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::metaObject", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlFileSelector::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlFileSelector::metaObject")
	}
}

func (ptr *QQmlFileSelector) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlFileSelector::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlFileSelector_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlFileSelector::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlFileSelector_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading = QQmlImageProviderBase__Flag(0x01)
)

//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image         = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap        = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture       = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid       = QQmlImageProviderBase__ImageType(3)
	QQmlImageProviderBase__ImageResponse = QQmlImageProviderBase__ImageType(4)
)

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (p *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return p
}

func (p *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlImageProviderBase) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQQmlImageProviderBase_Flags
func callbackQQmlImageProviderBase_Flags(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QQmlImageProviderBase::flags")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlImageProviderBase::flags"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__Flag)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectFlags(f func() QQmlImageProviderBase__Flag) {
	defer qt.Recovering("connect QQmlImageProviderBase::flags")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::flags", f)
	}
}

func (ptr *QQmlImageProviderBase) DisconnectFlags() {
	defer qt.Recovering("disconnect QQmlImageProviderBase::flags")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::flags")
	}
}

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	defer qt.Recovering("QQmlImageProviderBase::flags")

	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlImageProviderBase_ImageType
func callbackQQmlImageProviderBase_ImageType(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QQmlImageProviderBase::imageType")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlImageProviderBase::imageType"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectImageType(f func() QQmlImageProviderBase__ImageType) {
	defer qt.Recovering("connect QQmlImageProviderBase::imageType")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::imageType", f)
	}
}

func (ptr *QQmlImageProviderBase) DisconnectImageType() {
	defer qt.Recovering("disconnect QQmlImageProviderBase::imageType")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlImageProviderBase::imageType")
	}
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	defer qt.Recovering("QQmlImageProviderBase::imageType")

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

func (p *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return p
}

func (p *QQmlIncubationController) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlIncubationController) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func NewQQmlIncubationController() *QQmlIncubationController {
	defer qt.Recovering("QQmlIncubationController::QQmlIncubationController")

	return NewQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	defer qt.Recovering("QQmlIncubationController::engine")

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
	defer qt.Recovering("QQmlIncubationController::incubateFor")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQQmlIncubationController_IncubatingObjectCountChanged
func callbackQQmlIncubationController_IncubatingObjectCountChanged(ptr unsafe.Pointer, incubatingObjectCount C.int) {
	defer qt.Recovering("callback QQmlIncubationController::incubatingObjectCountChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubationController::incubatingObjectCountChanged"); signal != nil {
		signal.(func(int))(int(int32(incubatingObjectCount)))
	} else {
		NewQQmlIncubationControllerFromPointer(ptr).IncubatingObjectCountChangedDefault(int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {
	defer qt.Recovering("connect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubationController::incubatingObjectCountChanged", f)
	}
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {
	defer qt.Recovering("disconnect QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubationController::incubatingObjectCountChanged")
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChanged(incubatingObjectCount int) {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChanged(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChangedDefault(incubatingObjectCount int) {
	defer qt.Recovering("QQmlIncubationController::incubatingObjectCountChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChangedDefault(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          = QQmlIncubator__IncubationMode(2)
)

//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   = QQmlIncubator__Status(1)
	QQmlIncubator__Loading = QQmlIncubator__Status(2)
	QQmlIncubator__Error   = QQmlIncubator__Status(3)
)

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (p *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return p
}

func (p *QQmlIncubator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlIncubator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	defer qt.Recovering("QQmlIncubator::QQmlIncubator")

	return NewQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.longlong(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	defer qt.Recovering("QQmlIncubator::clear")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	defer qt.Recovering("QQmlIncubator::forceCompletion")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	defer qt.Recovering("QQmlIncubator::incubationMode")

	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	defer qt.Recovering("QQmlIncubator::isError")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	defer qt.Recovering("QQmlIncubator::isLoading")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsLoading(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	defer qt.Recovering("QQmlIncubator::isNull")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	defer qt.Recovering("QQmlIncubator::isReady")

	if ptr.Pointer() != nil {
		return C.QQmlIncubator_IsReady(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	defer qt.Recovering("QQmlIncubator::object")

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
	defer qt.Recovering("callback QQmlIncubator::setInitialState")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubator::setInitialState"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQQmlIncubatorFromPointer(ptr).SetInitialStateDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {
	defer qt.Recovering("connect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::setInitialState", f)
	}
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {
	defer qt.Recovering("disconnect QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::setInitialState")
	}
}

func (ptr *QQmlIncubator) SetInitialState(object core.QObject_ITF) {
	defer qt.Recovering("QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialState(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) SetInitialStateDefault(object core.QObject_ITF) {
	defer qt.Recovering("QQmlIncubator::setInitialState")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialStateDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	defer qt.Recovering("QQmlIncubator::status")

	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlIncubator_StatusChanged
func callbackQQmlIncubator_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	defer qt.Recovering("callback QQmlIncubator::statusChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlIncubator::statusChanged"); signal != nil {
		signal.(func(QQmlIncubator__Status))(QQmlIncubator__Status(status))
	} else {
		NewQQmlIncubatorFromPointer(ptr).StatusChangedDefault(QQmlIncubator__Status(status))
	}
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {
	defer qt.Recovering("connect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::statusChanged", f)
	}
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlIncubator::statusChanged")
	}
}

func (ptr *QQmlIncubator) StatusChanged(status QQmlIncubator__Status) {
	defer qt.Recovering("QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlIncubator) StatusChangedDefault(status QQmlIncubator__Status) {
	defer qt.Recovering("QQmlIncubator::statusChanged")

	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChangedDefault(ptr.Pointer(), C.longlong(status))
	}
}

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (p *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return p
}

func (p *QQmlListProperty) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlListProperty) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (p *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return p
}

func (p *QQmlListReference) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlListReference) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQQmlListReference() *QQmlListReference {
	defer qt.Recovering("QQmlListReference::QQmlListReference")

	var tmpValue = NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference())
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {
	defer qt.Recovering("QQmlListReference::QQmlListReference")

	var propertyC = C.CString(property)
	defer C.free(unsafe.Pointer(propertyC))
	var tmpValue = NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference2(core.PointerFromQObject(object), propertyC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {
	defer qt.Recovering("QQmlListReference::append")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Append(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	defer qt.Recovering("QQmlListReference::at")

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
	defer qt.Recovering("QQmlListReference::canAppend")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAppend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	defer qt.Recovering("QQmlListReference::canAt")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAt(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	defer qt.Recovering("QQmlListReference::canClear")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanClear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	defer qt.Recovering("QQmlListReference::canCount")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanCount(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	defer qt.Recovering("QQmlListReference::clear")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Count() int {
	defer qt.Recovering("QQmlListReference::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlListReference_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlListReference) IsManipulable() bool {
	defer qt.Recovering("QQmlListReference::isManipulable")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsManipulable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	defer qt.Recovering("QQmlListReference::isReadable")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	defer qt.Recovering("QQmlListReference::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	defer qt.Recovering("QQmlListReference::listElementType")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlListReference_ListElementType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	defer qt.Recovering("QQmlListReference::object")

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

func (p *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return p
}

func (p *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlNetworkAccessManagerFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	defer qt.Recovering("callback QQmlNetworkAccessManagerFactory::create")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlNetworkAccessManagerFactory::create"); signal != nil {
		return network.PointerFromQNetworkAccessManager(signal.(func(*core.QObject) *network.QNetworkAccessManager)(core.NewQObjectFromPointer(parent)))
	}

	return network.PointerFromQNetworkAccessManager(nil)
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectCreate(f func(parent *core.QObject) *network.QNetworkAccessManager) {
	defer qt.Recovering("connect QQmlNetworkAccessManagerFactory::create")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::create", f)
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectCreate(parent core.QObject_ITF) {
	defer qt.Recovering("disconnect QQmlNetworkAccessManagerFactory::create")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlNetworkAccessManagerFactory::create")
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::create")

	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	defer qt.Recovering("QQmlNetworkAccessManagerFactory::~QQmlNetworkAccessManagerFactory")

	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (p *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return p
}

func (p *QQmlParserStatus) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlParserStatus) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

//export callbackQQmlParserStatus_ClassBegin
func callbackQQmlParserStatus_ClassBegin(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlParserStatus::classBegin")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlParserStatus::classBegin"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectClassBegin(f func()) {
	defer qt.Recovering("connect QQmlParserStatus::classBegin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::classBegin", f)
	}
}

func (ptr *QQmlParserStatus) DisconnectClassBegin() {
	defer qt.Recovering("disconnect QQmlParserStatus::classBegin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::classBegin")
	}
}

func (ptr *QQmlParserStatus) ClassBegin() {
	defer qt.Recovering("QQmlParserStatus::classBegin")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

//export callbackQQmlParserStatus_ComponentComplete
func callbackQQmlParserStatus_ComponentComplete(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlParserStatus::componentComplete")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlParserStatus::componentComplete"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectComponentComplete(f func()) {
	defer qt.Recovering("connect QQmlParserStatus::componentComplete")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::componentComplete", f)
	}
}

func (ptr *QQmlParserStatus) DisconnectComponentComplete() {
	defer qt.Recovering("disconnect QQmlParserStatus::componentComplete")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlParserStatus::componentComplete")
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	defer qt.Recovering("QQmlParserStatus::componentComplete")

	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}

//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          = QQmlProperty__PropertyTypeCategory(3)
)

//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        = QQmlProperty__Type(0)
	QQmlProperty__Property       = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty = QQmlProperty__Type(2)
)

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (p *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return p
}

func (p *QQmlProperty) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlProperty) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQQmlProperty() *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty())
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty2(core.PointerFromQObject(obj)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty3(core.PointerFromQObject(obj), PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty4(core.PointerFromQObject(obj), PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty5(core.PointerFromQObject(obj), nameC))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty6(core.PointerFromQObject(obj), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty7(core.PointerFromQObject(obj), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {
	defer qt.Recovering("QQmlProperty::QQmlProperty")

	var tmpValue = NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty8(PointerFromQQmlProperty(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QQmlProperty::connectNotifySignal")

	if ptr.Pointer() != nil {
		var slotC = C.CString(slot)
		defer C.free(unsafe.Pointer(slotC))
		return C.QQmlProperty_ConnectNotifySignal(ptr.Pointer(), core.PointerFromQObject(dest), slotC) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {
	defer qt.Recovering("QQmlProperty::connectNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_ConnectNotifySignal2(ptr.Pointer(), core.PointerFromQObject(dest), C.int(int32(method))) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	defer qt.Recovering("QQmlProperty::hasNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_HasNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Index() int {
	defer qt.Recovering("QQmlProperty::index")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_Index(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) IsDesignable() bool {
	defer qt.Recovering("QQmlProperty::isDesignable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsDesignable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	defer qt.Recovering("QQmlProperty::isProperty")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	defer qt.Recovering("QQmlProperty::isResettable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsResettable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	defer qt.Recovering("QQmlProperty::isSignalProperty")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsSignalProperty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	defer qt.Recovering("QQmlProperty::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	defer qt.Recovering("QQmlProperty::isWritable")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Method() *core.QMetaMethod {
	defer qt.Recovering("QQmlProperty::method")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQMetaMethodFromPointer(C.QQmlProperty_Method(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QMetaMethod).DestroyQMetaMethod)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) Name() string {
	defer qt.Recovering("QQmlProperty::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	defer qt.Recovering("QQmlProperty::needsNotifySignal")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_NeedsNotifySignal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Object() *core.QObject {
	defer qt.Recovering("QQmlProperty::object")

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
	defer qt.Recovering("QQmlProperty::propertyType")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_PropertyType(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	defer qt.Recovering("QQmlProperty::propertyTypeCategory")

	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyTypeName() string {
	defer qt.Recovering("QQmlProperty::propertyTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlProperty_PropertyTypeName(ptr.Pointer()))
	}
	return ""
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read2(object core.QObject_ITF, name string) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), nameC))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), nameC, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), nameC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read() *core.QVariant {
	defer qt.Recovering("QQmlProperty::read")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QQmlProperty_Read(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) Reset() bool {
	defer qt.Recovering("QQmlProperty::reset")

	if ptr.Pointer() != nil {
		return C.QQmlProperty_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	defer qt.Recovering("QQmlProperty::type")

	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(ptr.Pointer()))
	}
	return 0
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value)) != 0
}

func (ptr *QQmlProperty) Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value)) != 0
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func (ptr *QQmlProperty) Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt)) != 0
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), nameC, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine)) != 0
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {
	defer qt.Recovering("QQmlProperty::write")

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

func (p *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return p
}

func (p *QQmlPropertyMap) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQmlPropertyMap) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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
	defer qt.Recovering("QQmlPropertyMap::QQmlPropertyMap")

	var tmpValue = NewQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	defer qt.Recovering("QQmlPropertyMap::clear")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_Clear(ptr.Pointer(), keyC)
	}
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	defer qt.Recovering("QQmlPropertyMap::contains")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		return C.QQmlPropertyMap_Contains(ptr.Pointer(), keyC) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Count() int {
	defer qt.Recovering("QQmlPropertyMap::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) Insert(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlPropertyMap::insert")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_Insert(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	defer qt.Recovering("QQmlPropertyMap::isEmpty")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) Keys() []string {
	defer qt.Recovering("QQmlPropertyMap::keys")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QQmlPropertyMap_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Size() int {
	defer qt.Recovering("QQmlPropertyMap::size")

	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Size(ptr.Pointer())))
	}
	return 0
}

//export callbackQQmlPropertyMap_UpdateValue
func callbackQQmlPropertyMap_UpdateValue(ptr unsafe.Pointer, key *C.char, input unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlPropertyMap::updateValue")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::updateValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string, *core.QVariant) *core.QVariant)(C.GoString(key), core.NewQVariantFromPointer(input)))
	}

	return core.PointerFromQVariant(NewQQmlPropertyMapFromPointer(ptr).UpdateValueDefault(C.GoString(key), core.NewQVariantFromPointer(input)))
}

func (ptr *QQmlPropertyMap) ConnectUpdateValue(f func(key string, input *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QQmlPropertyMap::updateValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::updateValue", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectUpdateValue() {
	defer qt.Recovering("disconnect QQmlPropertyMap::updateValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::updateValue")
	}
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QQmlPropertyMap::updateValue")

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
	defer qt.Recovering("QQmlPropertyMap::updateValue")

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
	defer qt.Recovering("QQmlPropertyMap::value")

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
func callbackQQmlPropertyMap_ValueChanged(ptr unsafe.Pointer, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::valueChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::valueChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::valueChanged", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::valueChanged")
	}
}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QQmlPropertyMap::valueChanged")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QQmlPropertyMap_ValueChanged(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	defer qt.Recovering("QQmlPropertyMap::~QQmlPropertyMap")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlPropertyMap_TimerEvent
func callbackQQmlPropertyMap_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::timerEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::timerEvent")
	}
}

func (ptr *QQmlPropertyMap) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::timerEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQmlPropertyMap_ChildEvent
func callbackQQmlPropertyMap_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::childEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::childEvent")
	}
}

func (ptr *QQmlPropertyMap) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::childEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlPropertyMap_ConnectNotify
func callbackQQmlPropertyMap_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlPropertyMap::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::connectNotify", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQmlPropertyMap::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::connectNotify")
	}
}

func (ptr *QQmlPropertyMap) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlPropertyMap::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlPropertyMap::connectNotify")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_CustomEvent
func callbackQQmlPropertyMap_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::customEvent", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::customEvent")
	}
}

func (ptr *QQmlPropertyMap) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQmlPropertyMap::customEvent")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlPropertyMap_DeleteLater
func callbackQQmlPropertyMap_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlPropertyMap) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQmlPropertyMap::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::deleteLater", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQmlPropertyMap::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::deleteLater")
	}
}

func (ptr *QQmlPropertyMap) DeleteLater() {
	defer qt.Recovering("QQmlPropertyMap::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyMap) DeleteLaterDefault() {
	defer qt.Recovering("QQmlPropertyMap::deleteLater")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQQmlPropertyMap_DisconnectNotify
func callbackQQmlPropertyMap_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyMap::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQmlPropertyMap::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::disconnectNotify", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQmlPropertyMap::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::disconnectNotify")
	}
}

func (ptr *QQmlPropertyMap) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlPropertyMap::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQmlPropertyMap) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQmlPropertyMap::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_Event
func callbackQQmlPropertyMap_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlPropertyMap::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlPropertyMap) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlPropertyMap::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::event", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectEvent() {
	defer qt.Recovering("disconnect QQmlPropertyMap::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::event")
	}
}

func (ptr *QQmlPropertyMap) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlPropertyMap::event")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlPropertyMap::event")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_EventFilter
func callbackQQmlPropertyMap_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QQmlPropertyMap::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlPropertyMap) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQmlPropertyMap::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::eventFilter", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQmlPropertyMap::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::eventFilter")
	}
}

func (ptr *QQmlPropertyMap) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlPropertyMap::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQmlPropertyMap::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQmlPropertyMap_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_MetaObject
func callbackQQmlPropertyMap_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQmlPropertyMap::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyMap::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlPropertyMapFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlPropertyMap) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQmlPropertyMap::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::metaObject", f)
	}
}

func (ptr *QQmlPropertyMap) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQmlPropertyMap::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyMap::metaObject")
	}
}

func (ptr *QQmlPropertyMap) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQmlPropertyMap::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlPropertyMap_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQmlPropertyMap::metaObject")

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

func (p *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return p
}

func (p *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlPropertyValueSource) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	defer qt.Recovering("QQmlPropertyValueSource::QQmlPropertyValueSource")

	return NewQQmlPropertyValueSourceFromPointer(C.QQmlPropertyValueSource_NewQQmlPropertyValueSource())
}

//export callbackQQmlPropertyValueSource_SetTarget
func callbackQQmlPropertyValueSource_SetTarget(ptr unsafe.Pointer, property unsafe.Pointer) {
	defer qt.Recovering("callback QQmlPropertyValueSource::setTarget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QQmlPropertyValueSource::setTarget"); signal != nil {
		signal.(func(*QQmlProperty))(NewQQmlPropertyFromPointer(property))
	}

}

func (ptr *QQmlPropertyValueSource) ConnectSetTarget(f func(property *QQmlProperty)) {
	defer qt.Recovering("connect QQmlPropertyValueSource::setTarget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::setTarget", f)
	}
}

func (ptr *QQmlPropertyValueSource) DisconnectSetTarget(property QQmlProperty_ITF) {
	defer qt.Recovering("disconnect QQmlPropertyValueSource::setTarget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QQmlPropertyValueSource::setTarget")
	}
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	defer qt.Recovering("QQmlPropertyValueSource::setTarget")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	defer qt.Recovering("QQmlPropertyValueSource::~QQmlPropertyValueSource")

	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
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

func (p *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return p
}

func (p *QQmlScriptString) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QQmlScriptString) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQQmlScriptString() *QQmlScriptString {
	defer qt.Recovering("QQmlScriptString::QQmlScriptString")

	var tmpValue = NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString())
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {
	defer qt.Recovering("QQmlScriptString::QQmlScriptString")

	var tmpValue = NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString2(PointerFromQQmlScriptString(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func (ptr *QQmlScriptString) BooleanLiteral(ok bool) bool {
	defer qt.Recovering("QQmlScriptString::booleanLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_BooleanLiteral(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	defer qt.Recovering("QQmlScriptString::isEmpty")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	defer qt.Recovering("QQmlScriptString::isNullLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsNullLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	defer qt.Recovering("QQmlScriptString::isUndefinedLiteral")

	if ptr.Pointer() != nil {
		return C.QQmlScriptString_IsUndefinedLiteral(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlScriptString) NumberLiteral(ok bool) float64 {
	defer qt.Recovering("QQmlScriptString::numberLiteral")

	if ptr.Pointer() != nil {
		return float64(C.QQmlScriptString_NumberLiteral(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))))
	}
	return 0
}

func (ptr *QQmlScriptString) StringLiteral() string {
	defer qt.Recovering("QQmlScriptString::stringLiteral")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlScriptString_StringLiteral(ptr.Pointer()))
	}
	return ""
}
