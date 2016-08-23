// +build !minimal

package testlib

//#include <stdint.h>
//#include <stdlib.h>
//#include "testlib.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

type QSignalSpy struct {
	core.QObject
	core.QList
}

type QSignalSpy_ITF interface {
	core.QObject_ITF
	core.QList_ITF
	QSignalSpy_PTR() *QSignalSpy
}

func (p *QSignalSpy) QSignalSpy_PTR() *QSignalSpy {
	return p
}

func (p *QSignalSpy) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSignalSpy) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QList_PTR().SetPointer(ptr)
	}
}

func PointerFromQSignalSpy(ptr QSignalSpy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalSpy_PTR().Pointer()
	}
	return nil
}

func NewQSignalSpyFromPointer(ptr unsafe.Pointer) *QSignalSpy {
	var n = new(QSignalSpy)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSignalSpy) DestroyQSignalSpy() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSignalSpy(object core.QObject_ITF, sign string) *QSignalSpy {
	defer qt.Recovering("QSignalSpy::QSignalSpy")

	var signC = C.CString(sign)
	defer C.free(unsafe.Pointer(signC))
	return NewQSignalSpyFromPointer(C.QSignalSpy_NewQSignalSpy(core.PointerFromQObject(object), signC))
}

func (ptr *QSignalSpy) IsValid() bool {
	defer qt.Recovering("QSignalSpy::isValid")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSignalSpy) Signal() string {
	defer qt.Recovering("QSignalSpy::signal")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSignalSpy_Signal(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSignalSpy) Wait(timeout int) bool {
	defer qt.Recovering("QSignalSpy::wait")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_Wait(ptr.Pointer(), C.int(int32(timeout))) != 0
	}
	return false
}

//export callbackQSignalSpy_TimerEvent
func callbackQSignalSpy_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::timerEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSignalSpy::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "timerEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSignalSpy::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "timerEvent")
	}
}

func (ptr *QSignalSpy) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSignalSpy::timerEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSignalSpy) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSignalSpy::timerEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSignalSpy_ChildEvent
func callbackQSignalSpy_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::childEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSignalSpy::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "childEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSignalSpy::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "childEvent")
	}
}

func (ptr *QSignalSpy) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSignalSpy::childEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSignalSpy) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSignalSpy::childEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSignalSpy_ConnectNotify
func callbackQSignalSpy_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::connectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSignalSpy::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "connectNotify", f)
	}
}

func (ptr *QSignalSpy) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSignalSpy::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "connectNotify")
	}
}

func (ptr *QSignalSpy) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSignalSpy::connectNotify")

	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSignalSpy) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSignalSpy::connectNotify")

	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_CustomEvent
func callbackQSignalSpy_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::customEvent")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSignalSpy::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "customEvent", f)
	}
}

func (ptr *QSignalSpy) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSignalSpy::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "customEvent")
	}
}

func (ptr *QSignalSpy) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSignalSpy::customEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSignalSpy) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSignalSpy::customEvent")

	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSignalSpy_DeleteLater
func callbackQSignalSpy_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::deleteLater")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSignalSpyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSignalSpy) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSignalSpy::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "deleteLater", f)
	}
}

func (ptr *QSignalSpy) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSignalSpy::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "deleteLater")
	}
}

func (ptr *QSignalSpy) DeleteLater() {
	defer qt.Recovering("QSignalSpy::deleteLater")

	if ptr.Pointer() != nil {
		C.QSignalSpy_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSignalSpy) DeleteLaterDefault() {
	defer qt.Recovering("QSignalSpy::deleteLater")

	if ptr.Pointer() != nil {
		C.QSignalSpy_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSignalSpy_DisconnectNotify
func callbackQSignalSpy_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSignalSpy::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSignalSpy::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "disconnectNotify", f)
	}
}

func (ptr *QSignalSpy) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSignalSpy::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "disconnectNotify")
	}
}

func (ptr *QSignalSpy) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSignalSpy::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSignalSpy) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSignalSpy::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_Event
func callbackQSignalSpy_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSignalSpy::event")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSignalSpy) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSignalSpy::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "event", f)
	}
}

func (ptr *QSignalSpy) DisconnectEvent() {
	defer qt.Recovering("disconnect QSignalSpy::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "event")
	}
}

func (ptr *QSignalSpy) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSignalSpy::event")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSignalSpy) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSignalSpy::event")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSignalSpy_EventFilter
func callbackQSignalSpy_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSignalSpy::eventFilter")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSignalSpy) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSignalSpy::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "eventFilter", f)
	}
}

func (ptr *QSignalSpy) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSignalSpy::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "eventFilter")
	}
}

func (ptr *QSignalSpy) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSignalSpy::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSignalSpy) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSignalSpy::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSignalSpy_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSignalSpy_MetaObject
func callbackQSignalSpy_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSignalSpy::metaObject")

	if signal := qt.GetSignal(fmt.Sprintf("QSignalSpy(%v)", ptr), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSignalSpyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSignalSpy) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSignalSpy::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "metaObject", f)
	}
}

func (ptr *QSignalSpy) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSignalSpy::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprintf("QSignalSpy(%v)", ptr.Pointer()), "metaObject")
	}
}

func (ptr *QSignalSpy) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSignalSpy::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSignalSpy) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSignalSpy::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTestEventList struct {
	core.QList
}

type QTestEventList_ITF interface {
	core.QList_ITF
	QTestEventList_PTR() *QTestEventList
}

func (p *QTestEventList) QTestEventList_PTR() *QTestEventList {
	return p
}

func (p *QTestEventList) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QList_PTR().Pointer()
	}
	return nil
}

func (p *QTestEventList) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QList_PTR().SetPointer(ptr)
	}
}

func PointerFromQTestEventList(ptr QTestEventList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventList_PTR().Pointer()
	}
	return nil
}

func NewQTestEventListFromPointer(ptr unsafe.Pointer) *QTestEventList {
	var n = new(QTestEventList)
	n.SetPointer(ptr)
	return n
}
func NewQTestEventList() *QTestEventList {
	defer qt.Recovering("QTestEventList::QTestEventList")

	var tmpValue = NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList())
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func NewQTestEventList2(other QTestEventList_ITF) *QTestEventList {
	defer qt.Recovering("QTestEventList::QTestEventList")

	var tmpValue = NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList2(PointerFromQTestEventList(other)))
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func (ptr *QTestEventList) AddDelay(msecs int) {
	defer qt.Recovering("QTestEventList::addDelay")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddDelay(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyClick")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyClick(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyClick")

	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyClick2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClicks(keys string, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyClicks")

	if ptr.Pointer() != nil {
		var keysC = C.CString(keys)
		defer C.free(unsafe.Pointer(keysC))
		C.QTestEventList_AddKeyClicks(ptr.Pointer(), keysC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyPress")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyPress(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyPress")

	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyPress2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyRelease")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyRelease(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	defer qt.Recovering("QTestEventList::addKeyRelease")

	if ptr.Pointer() != nil {
		var asciiC = C.CString(ascii)
		defer C.free(unsafe.Pointer(asciiC))
		C.QTestEventList_AddKeyRelease2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddMouseClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	defer qt.Recovering("QTestEventList::addMouseClick")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseDClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	defer qt.Recovering("QTestEventList::addMouseDClick")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseDClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseMove(pos core.QPoint_ITF, delay int) {
	defer qt.Recovering("QTestEventList::addMouseMove")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseMove(ptr.Pointer(), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMousePress(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	defer qt.Recovering("QTestEventList::addMousePress")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddMousePress(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseRelease(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	defer qt.Recovering("QTestEventList::addMouseRelease")

	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseRelease(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) Clear() {
	defer qt.Recovering("QTestEventList::clear")

	if ptr.Pointer() != nil {
		C.QTestEventList_Clear(ptr.Pointer())
	}
}

func (ptr *QTestEventList) Simulate(w widgets.QWidget_ITF) {
	defer qt.Recovering("QTestEventList::simulate")

	if ptr.Pointer() != nil {
		C.QTestEventList_Simulate(ptr.Pointer(), widgets.PointerFromQWidget(w))
	}
}

func (ptr *QTestEventList) DestroyQTestEventList() {
	defer qt.Recovering("QTestEventList::~QTestEventList")

	if ptr.Pointer() != nil {
		C.QTestEventList_DestroyQTestEventList(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
