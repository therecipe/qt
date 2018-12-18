// +build !minimal

package testlib

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "testlib.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtTestLib_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtTestLib_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QAbstractItemModelTester struct {
	ptr unsafe.Pointer
}

type QAbstractItemModelTester_ITF interface {
	QAbstractItemModelTester_PTR() *QAbstractItemModelTester
}

func (ptr *QAbstractItemModelTester) QAbstractItemModelTester_PTR() *QAbstractItemModelTester {
	return ptr
}

func (ptr *QAbstractItemModelTester) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractItemModelTester) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractItemModelTester(ptr QAbstractItemModelTester_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemModelTester_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemModelTesterFromPointer(ptr unsafe.Pointer) (n *QAbstractItemModelTester) {
	n = new(QAbstractItemModelTester)
	n.SetPointer(ptr)
	return
}

func (ptr *QAbstractItemModelTester) DestroyQAbstractItemModelTester() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QAbstractItemModelTester__FailureReportingMode
//QAbstractItemModelTester::FailureReportingMode
type QAbstractItemModelTester__FailureReportingMode int64

const (
	QAbstractItemModelTester__QtTest  QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(0)
	QAbstractItemModelTester__Warning QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(1)
	QAbstractItemModelTester__Fatal   QAbstractItemModelTester__FailureReportingMode = QAbstractItemModelTester__FailureReportingMode(2)
)

func NewQAbstractItemModelTester2(model core.QAbstractItemModel_ITF, mode QAbstractItemModelTester__FailureReportingMode, parent core.QObject_ITF) *QAbstractItemModelTester {
	return NewQAbstractItemModelTesterFromPointer(C.QAbstractItemModelTester_NewQAbstractItemModelTester2(core.PointerFromQAbstractItemModel(model), C.longlong(mode), core.PointerFromQObject(parent)))
}

func NewQAbstractItemModelTester(model core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QAbstractItemModelTester {
	return NewQAbstractItemModelTesterFromPointer(C.QAbstractItemModelTester_NewQAbstractItemModelTester(core.PointerFromQAbstractItemModel(model), core.PointerFromQObject(parent)))
}

func QAbstractItemModelTester_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractItemModelTester_QAbstractItemModelTester_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractItemModelTester) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractItemModelTester_QAbstractItemModelTester_Tr(sC, cC, C.int(int32(n))))
}

func QAbstractItemModelTester_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractItemModelTester_QAbstractItemModelTester_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractItemModelTester) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QAbstractItemModelTester_QAbstractItemModelTester_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractItemModelTester) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQAbstractItemModelFromPointer(C.QAbstractItemModelTester_Model(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractItemModelTester_MetaObject
func callbackQAbstractItemModelTester_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractItemModelTesterFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractItemModelTester) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "metaObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "metaObject", func() *core.QMetaObject {
				signal.(func() *core.QMetaObject)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "metaObject", f)
		}
	}
}

func (ptr *QAbstractItemModelTester) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metaObject")
	}
}

func (ptr *QAbstractItemModelTester) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractItemModelTester_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemModelTester) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractItemModelTester_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSignalSpy struct {
	core.QObject
}

type QSignalSpy_ITF interface {
	core.QObject_ITF
	QSignalSpy_PTR() *QSignalSpy
}

func (ptr *QSignalSpy) QSignalSpy_PTR() *QSignalSpy {
	return ptr
}

func (ptr *QSignalSpy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QSignalSpy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQSignalSpy(ptr QSignalSpy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSignalSpy_PTR().Pointer()
	}
	return nil
}

func NewQSignalSpyFromPointer(ptr unsafe.Pointer) (n *QSignalSpy) {
	n = new(QSignalSpy)
	n.SetPointer(ptr)
	return
}
func NewQSignalSpy(object core.QObject_ITF, sign string) *QSignalSpy {
	var signC *C.char
	if sign != "" {
		signC = C.CString(sign)
		defer C.free(unsafe.Pointer(signC))
	}
	tmpValue := NewQSignalSpyFromPointer(C.QSignalSpy_NewQSignalSpy(core.PointerFromQObject(object), signC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSignalSpy) Wait(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_Wait(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QSignalSpy) Signal() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSignalSpy_Signal(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSignalSpy) __args_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSignalSpy___args_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSignalSpy) __args_setList(i int) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___args_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSignalSpy) __args_newList() unsafe.Pointer {
	return C.QSignalSpy___args_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __setArgs__atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QSignalSpy___setArgs__atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QSignalSpy) __setArgs__setList(i int) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___setArgs__setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QSignalSpy) __setArgs__newList() unsafe.Pointer {
	return C.QSignalSpy___setArgs__newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QSignalSpy___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QSignalSpy) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QSignalSpy___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __findChildren_newList2() unsafe.Pointer {
	return C.QSignalSpy___findChildren_newList2(ptr.Pointer())
}

func (ptr *QSignalSpy) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __findChildren_newList3() unsafe.Pointer {
	return C.QSignalSpy___findChildren_newList3(ptr.Pointer())
}

func (ptr *QSignalSpy) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __findChildren_newList() unsafe.Pointer {
	return C.QSignalSpy___findChildren_newList(ptr.Pointer())
}

func (ptr *QSignalSpy) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QSignalSpy___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSignalSpy) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QSignalSpy) __children_newList() unsafe.Pointer {
	return C.QSignalSpy___children_newList(ptr.Pointer())
}

//export callbackQSignalSpy_Event
func callbackQSignalSpy_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSignalSpy) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQSignalSpy_EventFilter
func callbackQSignalSpy_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSignalSpyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSignalSpy) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QSignalSpy_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQSignalSpy_ChildEvent
func callbackQSignalSpy_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSignalSpy_ConnectNotify
func callbackQSignalSpy_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_CustomEvent
func callbackQSignalSpy_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSignalSpy_DeleteLater
func callbackQSignalSpy_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSignalSpyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSignalSpy) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSignalSpy_Destroyed
func callbackQSignalSpy_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQSignalSpy_DisconnectNotify
func callbackQSignalSpy_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSignalSpyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSignalSpy) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSignalSpy_ObjectNameChanged
func callbackQSignalSpy_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtTestLib_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQSignalSpy_TimerEvent
func callbackQSignalSpy_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSignalSpyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSignalSpy) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QSignalSpy_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSignalSpy_MetaObject
func callbackQSignalSpy_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSignalSpyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSignalSpy) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSignalSpy_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTestEventList struct {
	ptr unsafe.Pointer
}

type QTestEventList_ITF interface {
	QTestEventList_PTR() *QTestEventList
}

func (ptr *QTestEventList) QTestEventList_PTR() *QTestEventList {
	return ptr
}

func (ptr *QTestEventList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTestEventList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTestEventList(ptr QTestEventList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTestEventList_PTR().Pointer()
	}
	return nil
}

func NewQTestEventListFromPointer(ptr unsafe.Pointer) (n *QTestEventList) {
	n = new(QTestEventList)
	n.SetPointer(ptr)
	return
}
func NewQTestEventList() *QTestEventList {
	tmpValue := NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList())
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func NewQTestEventList2(other QTestEventList_ITF) *QTestEventList {
	tmpValue := NewQTestEventListFromPointer(C.QTestEventList_NewQTestEventList2(PointerFromQTestEventList(other)))
	runtime.SetFinalizer(tmpValue, (*QTestEventList).DestroyQTestEventList)
	return tmpValue
}

func (ptr *QTestEventList) AddDelay(msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddDelay(ptr.Pointer(), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyClick(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClick2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
		C.QTestEventList_AddKeyClick2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyClicks(keys string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var keysC *C.char
		if keys != "" {
			keysC = C.CString(keys)
			defer C.free(unsafe.Pointer(keysC))
		}
		C.QTestEventList_AddKeyClicks(ptr.Pointer(), C.struct_QtTestLib_PackedString{data: keysC, len: C.longlong(len(keys))}, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyPress(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyPress2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
		C.QTestEventList_AddKeyPress2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease(qtKey core.Qt__Key, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddKeyRelease(ptr.Pointer(), C.longlong(qtKey), C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddKeyRelease2(ascii string, modifiers core.Qt__KeyboardModifier, msecs int) {
	if ptr.Pointer() != nil {
		var asciiC *C.char
		if ascii != "" {
			asciiC = C.CString(ascii)
			defer C.free(unsafe.Pointer(asciiC))
		}
		C.QTestEventList_AddKeyRelease2(ptr.Pointer(), asciiC, C.longlong(modifiers), C.int(int32(msecs)))
	}
}

func (ptr *QTestEventList) AddMouseClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseDClick(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseDClick(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseMove(pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseMove(ptr.Pointer(), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMousePress(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMousePress(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) AddMouseRelease(button core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, pos core.QPoint_ITF, delay int) {
	if ptr.Pointer() != nil {
		C.QTestEventList_AddMouseRelease(ptr.Pointer(), C.longlong(button), C.longlong(modifiers), core.PointerFromQPoint(pos), C.int(int32(delay)))
	}
}

func (ptr *QTestEventList) Clear() {
	if ptr.Pointer() != nil {
		C.QTestEventList_Clear(ptr.Pointer())
	}
}

func (ptr *QTestEventList) Simulate(w widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QTestEventList_Simulate(ptr.Pointer(), widgets.PointerFromQWidget(w))
	}
}

func (ptr *QTestEventList) DestroyQTestEventList() {
	if ptr.Pointer() != nil {
		C.QTestEventList_DestroyQTestEventList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}
