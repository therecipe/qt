// +build !minimal

package dbus

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtDBus_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtDBus_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QDBus struct {
	ptr unsafe.Pointer
}

type QDBus_ITF interface {
	QDBus_PTR() *QDBus
}

func (ptr *QDBus) QDBus_PTR() *QDBus {
	return ptr
}

func (ptr *QDBus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBus(ptr QDBus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBus_PTR().Pointer()
	}
	return nil
}

func NewQDBusFromPointer(ptr unsafe.Pointer) (n *QDBus) {
	n = new(QDBus)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBus) DestroyQDBus() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QDBus__CallMode
//QDBus::CallMode
type QDBus__CallMode int64

const (
	QDBus__NoBlock      QDBus__CallMode = QDBus__CallMode(0)
	QDBus__Block        QDBus__CallMode = QDBus__CallMode(1)
	QDBus__BlockWithGui QDBus__CallMode = QDBus__CallMode(2)
	QDBus__AutoDetect   QDBus__CallMode = QDBus__CallMode(3)
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptor_ITF interface {
	core.QObject_ITF
	QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor
}

func (ptr *QDBusAbstractAdaptor) QDBusAbstractAdaptor_PTR() *QDBusAbstractAdaptor {
	return ptr
}

func (ptr *QDBusAbstractAdaptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) (n *QDBusAbstractAdaptor) {
	n = new(QDBusAbstractAdaptor)
	n.SetPointer(ptr)
	return
}
func NewQDBusAbstractAdaptor(obj core.QObject_ITF) *QDBusAbstractAdaptor {
	tmpValue := NewQDBusAbstractAdaptorFromPointer(C.QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(core.PointerFromQObject(obj)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDBusAbstractAdaptor_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractAdaptor_QDBusAbstractAdaptor_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractAdaptor) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractAdaptor_QDBusAbstractAdaptor_Tr(sC, cC, C.int(int32(n))))
}

func QDBusAbstractAdaptor_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractAdaptor_QDBusAbstractAdaptor_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractAdaptor) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractAdaptor_QDBusAbstractAdaptor_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractAdaptor) SetAutoRelaySignals(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_SetAutoRelaySignals(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

//export callbackQDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor
func callbackQDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusAbstractAdaptor"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DestroyQDBusAbstractAdaptorDefault()
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectDestroyQDBusAbstractAdaptor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusAbstractAdaptor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusAbstractAdaptor", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusAbstractAdaptor", f)
		}
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectDestroyQDBusAbstractAdaptor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusAbstractAdaptor")
	}
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptorDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusAbstractAdaptor) AutoRelaySignals() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractAdaptor_AutoRelaySignals(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_MetaObject
func callbackQDBusAbstractAdaptor_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractAdaptorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractAdaptor) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractAdaptor_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusAbstractAdaptor___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusAbstractAdaptor) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusAbstractAdaptor___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractAdaptor) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractAdaptor___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractAdaptor) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusAbstractAdaptor___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusAbstractAdaptor) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractAdaptor___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractAdaptor) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusAbstractAdaptor___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusAbstractAdaptor) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractAdaptor___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractAdaptor) __findChildren_newList() unsafe.Pointer {
	return C.QDBusAbstractAdaptor___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractAdaptor) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractAdaptor___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractAdaptor) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractAdaptor) __children_newList() unsafe.Pointer {
	return C.QDBusAbstractAdaptor___children_newList(ptr.Pointer())
}

//export callbackQDBusAbstractAdaptor_Event
func callbackQDBusAbstractAdaptor_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractAdaptor) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractAdaptor_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_EventFilter
func callbackQDBusAbstractAdaptor_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractAdaptorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractAdaptor) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractAdaptor_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusAbstractAdaptor_ChildEvent
func callbackQDBusAbstractAdaptor_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_ConnectNotify
func callbackQDBusAbstractAdaptor_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_CustomEvent
func callbackQDBusAbstractAdaptor_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractAdaptor_DeleteLater
func callbackQDBusAbstractAdaptor_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractAdaptor) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusAbstractAdaptor_Destroyed
func callbackQDBusAbstractAdaptor_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusAbstractAdaptor_DisconnectNotify
func callbackQDBusAbstractAdaptor_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractAdaptor) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractAdaptor_ObjectNameChanged
func callbackQDBusAbstractAdaptor_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusAbstractAdaptor_TimerEvent
func callbackQDBusAbstractAdaptor_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractAdaptorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractAdaptor) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusAbstractInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusAbstractInterface) {
	n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	return
}
func (ptr *QDBusAbstractInterface) Call2(mode QDBus__CallMode, method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusMessage {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		tmpValue := NewQDBusMessageFromPointer(C.QDBusAbstractInterface_Call2(ptr.Pointer(), C.longlong(mode), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, core.PointerFromQVariant(arg1), core.PointerFromQVariant(arg2), core.PointerFromQVariant(arg3), core.PointerFromQVariant(arg4), core.PointerFromQVariant(arg5), core.PointerFromQVariant(arg6), core.PointerFromQVariant(arg7), core.PointerFromQVariant(arg8)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Call(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusMessage {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		tmpValue := NewQDBusMessageFromPointer(C.QDBusAbstractInterface_Call(ptr.Pointer(), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, core.PointerFromQVariant(arg1), core.PointerFromQVariant(arg2), core.PointerFromQVariant(arg3), core.PointerFromQVariant(arg4), core.PointerFromQVariant(arg5), core.PointerFromQVariant(arg6), core.PointerFromQVariant(arg7), core.PointerFromQVariant(arg8)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) CallWithArgumentList(mode QDBus__CallMode, method string, args []*core.QVariant) *QDBusMessage {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		tmpValue := NewQDBusMessageFromPointer(C.QDBusAbstractInterface_CallWithArgumentList(ptr.Pointer(), C.longlong(mode), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, func() unsafe.Pointer {
			tmpList := NewQDBusAbstractInterfaceFromPointer(NewQDBusAbstractInterfaceFromPointer(nil).__callWithArgumentList_args_newList())
			for _, v := range args {
				tmpList.__callWithArgumentList_args_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) AsyncCall(method string, arg1 core.QVariant_ITF, arg2 core.QVariant_ITF, arg3 core.QVariant_ITF, arg4 core.QVariant_ITF, arg5 core.QVariant_ITF, arg6 core.QVariant_ITF, arg7 core.QVariant_ITF, arg8 core.QVariant_ITF) *QDBusPendingCall {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		tmpValue := NewQDBusPendingCallFromPointer(C.QDBusAbstractInterface_AsyncCall(ptr.Pointer(), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, core.PointerFromQVariant(arg1), core.PointerFromQVariant(arg2), core.PointerFromQVariant(arg3), core.PointerFromQVariant(arg4), core.PointerFromQVariant(arg5), core.PointerFromQVariant(arg6), core.PointerFromQVariant(arg7), core.PointerFromQVariant(arg8)))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) AsyncCallWithArgumentList(method string, args []*core.QVariant) *QDBusPendingCall {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		tmpValue := NewQDBusPendingCallFromPointer(C.QDBusAbstractInterface_AsyncCallWithArgumentList(ptr.Pointer(), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, func() unsafe.Pointer {
			tmpList := NewQDBusAbstractInterfaceFromPointer(NewQDBusAbstractInterfaceFromPointer(nil).__asyncCallWithArgumentList_args_newList())
			for _, v := range args {
				tmpList.__asyncCallWithArgumentList_args_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func QDBusAbstractInterface_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractInterface_QDBusAbstractInterface_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractInterface) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractInterface_QDBusAbstractInterface_Tr(sC, cC, C.int(int32(n))))
}

func QDBusAbstractInterface_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractInterface_QDBusAbstractInterface_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractInterface) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusAbstractInterface_QDBusAbstractInterface_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusAbstractInterface) CallWithCallback(method string, args []*core.QVariant, receiver core.QObject_ITF, returnMethod string, errorMethod string) bool {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		var returnMethodC *C.char
		if returnMethod != "" {
			returnMethodC = C.CString(returnMethod)
			defer C.free(unsafe.Pointer(returnMethodC))
		}
		var errorMethodC *C.char
		if errorMethod != "" {
			errorMethodC = C.CString(errorMethod)
			defer C.free(unsafe.Pointer(errorMethodC))
		}
		return int8(C.QDBusAbstractInterface_CallWithCallback(ptr.Pointer(), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, func() unsafe.Pointer {
			tmpList := NewQDBusAbstractInterfaceFromPointer(NewQDBusAbstractInterfaceFromPointer(nil).__callWithCallback_args_newList())
			for _, v := range args {
				tmpList.__callWithCallback_args_setList(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQObject(receiver), returnMethodC, errorMethodC)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) CallWithCallback2(method string, args []*core.QVariant, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var methodC *C.char
		if method != "" {
			methodC = C.CString(method)
			defer C.free(unsafe.Pointer(methodC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusAbstractInterface_CallWithCallback2(ptr.Pointer(), C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}, func() unsafe.Pointer {
			tmpList := NewQDBusAbstractInterfaceFromPointer(NewQDBusAbstractInterfaceFromPointer(nil).__callWithCallback_args_newList2())
			for _, v := range args {
				tmpList.__callWithCallback_args_setList2(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

//export callbackQDBusAbstractInterface_DestroyQDBusAbstractInterface
func callbackQDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusAbstractInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DestroyQDBusAbstractInterfaceDefault()
	}
}

func (ptr *QDBusAbstractInterface) ConnectDestroyQDBusAbstractInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusAbstractInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusAbstractInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusAbstractInterface", f)
		}
	}
}

func (ptr *QDBusAbstractInterface) DisconnectDestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusAbstractInterface")
	}
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusAbstractInterface) Connection() *QDBusConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusConnectionFromPointer(C.QDBusAbstractInterface_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusErrorFromPointer(C.QDBusAbstractInterface_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Interface() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractInterface_IsValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_MetaObject
func callbackQDBusAbstractInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusAbstractInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusAbstractInterface) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusAbstractInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDBusAbstractInterface_Timeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusAbstractInterface___callWithArgumentList_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___callWithArgumentList_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusAbstractInterface) __callWithArgumentList_args_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___callWithArgumentList_args_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusAbstractInterface___asyncCallWithArgumentList_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___asyncCallWithArgumentList_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusAbstractInterface) __asyncCallWithArgumentList_args_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___asyncCallWithArgumentList_args_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusAbstractInterface___callWithCallback_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___callWithCallback_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___callWithCallback_args_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_atList2(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusAbstractInterface___callWithCallback_args_atList2(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_setList2(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___callWithCallback_args_setList2(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusAbstractInterface) __callWithCallback_args_newList2() unsafe.Pointer {
	return C.QDBusAbstractInterface___callWithCallback_args_newList2(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusAbstractInterface___internalConstCall_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___internalConstCall_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusAbstractInterface) __internalConstCall_args_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___internalConstCall_args_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusAbstractInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusAbstractInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractInterface) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusAbstractInterface___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractInterface) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusAbstractInterface___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractInterface) __findChildren_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusAbstractInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusAbstractInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusAbstractInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusAbstractInterface) __children_newList() unsafe.Pointer {
	return C.QDBusAbstractInterface___children_newList(ptr.Pointer())
}

//export callbackQDBusAbstractInterface_Event
func callbackQDBusAbstractInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusAbstractInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_EventFilter
func callbackQDBusAbstractInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusAbstractInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusAbstractInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusAbstractInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusAbstractInterface_ChildEvent
func callbackQDBusAbstractInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusAbstractInterface_ConnectNotify
func callbackQDBusAbstractInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_CustomEvent
func callbackQDBusAbstractInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusAbstractInterface_DeleteLater
func callbackQDBusAbstractInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusAbstractInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusAbstractInterface_Destroyed
func callbackQDBusAbstractInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusAbstractInterface_DisconnectNotify
func callbackQDBusAbstractInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusAbstractInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusAbstractInterface_ObjectNameChanged
func callbackQDBusAbstractInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusAbstractInterface_TimerEvent
func callbackQDBusAbstractInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusAbstractInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusAbstractInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDBusArgument struct {
	ptr unsafe.Pointer
}

type QDBusArgument_ITF interface {
	QDBusArgument_PTR() *QDBusArgument
}

func (ptr *QDBusArgument) QDBusArgument_PTR() *QDBusArgument {
	return ptr
}

func (ptr *QDBusArgument) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusArgument) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusArgument(ptr QDBusArgument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusArgument_PTR().Pointer()
	}
	return nil
}

func NewQDBusArgumentFromPointer(ptr unsafe.Pointer) (n *QDBusArgument) {
	n = new(QDBusArgument)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDBusArgument__ElementType
//QDBusArgument::ElementType
type QDBusArgument__ElementType int64

const (
	QDBusArgument__BasicType     QDBusArgument__ElementType = QDBusArgument__ElementType(0)
	QDBusArgument__VariantType   QDBusArgument__ElementType = QDBusArgument__ElementType(1)
	QDBusArgument__ArrayType     QDBusArgument__ElementType = QDBusArgument__ElementType(2)
	QDBusArgument__StructureType QDBusArgument__ElementType = QDBusArgument__ElementType(3)
	QDBusArgument__MapType       QDBusArgument__ElementType = QDBusArgument__ElementType(4)
	QDBusArgument__MapEntryType  QDBusArgument__ElementType = QDBusArgument__ElementType(5)
	QDBusArgument__UnknownType   QDBusArgument__ElementType = QDBusArgument__ElementType(-1)
)

func NewQDBusArgument() *QDBusArgument {
	tmpValue := NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument())
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument3(other QDBusArgument_ITF) *QDBusArgument {
	tmpValue := NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument3(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func NewQDBusArgument2(other QDBusArgument_ITF) *QDBusArgument {
	tmpValue := NewQDBusArgumentFromPointer(C.QDBusArgument_NewQDBusArgument2(PointerFromQDBusArgument(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusArgument).DestroyQDBusArgument)
	return tmpValue
}

func (ptr *QDBusArgument) BeginArray(id int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QDBusArgument) BeginMap(kid int, vid int) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap(ptr.Pointer(), C.int(int32(kid)), C.int(int32(vid)))
	}
}

func (ptr *QDBusArgument) BeginMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) Swap(other QDBusArgument_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusArgument_Swap(ptr.Pointer(), PointerFromQDBusArgument(other))
	}
}

func (ptr *QDBusArgument) DestroyQDBusArgument() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_DestroyQDBusArgument(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusArgument) CurrentType() QDBusArgument__ElementType {
	if ptr.Pointer() != nil {
		return QDBusArgument__ElementType(C.QDBusArgument_CurrentType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusArgument) AsVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusArgument_AsVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusArgument) AtEnd() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusArgument_AtEnd(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusArgument) BeginArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) BeginStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_BeginStructure2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndArray2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndArray2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMap2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMap2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndMapEntry2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndMapEntry2(ptr.Pointer())
	}
}

func (ptr *QDBusArgument) EndStructure2() {
	if ptr.Pointer() != nil {
		C.QDBusArgument_EndStructure2(ptr.Pointer())
	}
}

type QDBusConnection struct {
	ptr unsafe.Pointer
}

type QDBusConnection_ITF interface {
	QDBusConnection_PTR() *QDBusConnection
}

func (ptr *QDBusConnection) QDBusConnection_PTR() *QDBusConnection {
	return ptr
}

func (ptr *QDBusConnection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusConnection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusConnection(ptr QDBusConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnection_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionFromPointer(ptr unsafe.Pointer) (n *QDBusConnection) {
	n = new(QDBusConnection)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDBusConnection__BusType
//QDBusConnection::BusType
type QDBusConnection__BusType int64

const (
	QDBusConnection__SessionBus    QDBusConnection__BusType = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     QDBusConnection__BusType = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus QDBusConnection__BusType = QDBusConnection__BusType(2)
)

//go:generate stringer -type=QDBusConnection__ConnectionCapability
//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int64

const (
	QDBusConnection__UnixFileDescriptorPassing QDBusConnection__ConnectionCapability = QDBusConnection__ConnectionCapability(0x0001)
)

//go:generate stringer -type=QDBusConnection__RegisterOption
//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int64

const (
	QDBusConnection__ExportAdaptors                QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportChildObjects            QDBusConnection__RegisterOption = QDBusConnection__RegisterOption(0x1000)
)

//go:generate stringer -type=QDBusConnection__UnregisterMode
//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int64

const (
	QDBusConnection__UnregisterNode QDBusConnection__UnregisterMode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree QDBusConnection__UnregisterMode = QDBusConnection__UnregisterMode(1)
)

func QDBusConnection_LocalMachineId() *core.QByteArray {
	tmpValue := core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *QDBusConnection) LocalMachineId() *core.QByteArray {
	tmpValue := core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
	runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func QDBusConnection_ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus(ty QDBusConnection__BusType, name string) *QDBusConnection {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus(C.longlong(ty), C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToBus2(address string, name string) *QDBusConnection {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(C.struct_QtDBus_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToBus2(address string, name string) *QDBusConnection {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToBus2(C.struct_QtDBus_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_ConnectToPeer(address string, name string) *QDBusConnection {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(C.struct_QtDBus_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) ConnectToPeer(address string, name string) *QDBusConnection {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_ConnectToPeer(C.struct_QtDBus_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_SessionBus() *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SessionBus() *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SessionBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func QDBusConnection_SystemBus() *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) SystemBus() *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_QDBusConnection_SystemBus())
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection3(other QDBusConnection_ITF) *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection3(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection2(PointerFromQDBusConnection(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func NewQDBusConnection(name string) *QDBusConnection {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
	return tmpValue
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Connect(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Connect2(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		argumentMatchC := C.CString(strings.Join(argumentMatch, "|"))
		defer C.free(unsafe.Pointer(argumentMatchC))
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Connect3(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: argumentMatchC, len: C.longlong(len(strings.Join(argumentMatch, "|")))}, C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Disconnect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Disconnect(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Disconnect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Disconnect2(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Disconnect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		argumentMatchC := C.CString(strings.Join(argumentMatch, "|"))
		defer C.free(unsafe.Pointer(argumentMatchC))
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QDBusConnection_Disconnect3(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: argumentMatchC, len: C.longlong(len(strings.Join(argumentMatch, "|")))}, C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))}, core.PointerFromQObject(receiver), slotC)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		return int8(C.QDBusConnection_RegisterObject(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, core.PointerFromQObject(object), C.longlong(options))) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		var interfaC *C.char
		if interfa != "" {
			interfaC = C.CString(interfa)
			defer C.free(unsafe.Pointer(interfaC))
		}
		return int8(C.QDBusConnection_RegisterObject2(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, core.PointerFromQObject(object), C.longlong(options))) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		return int8(C.QDBusConnection_RegisterService(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))})) != 0
	}
	return false
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		return int8(C.QDBusConnection_UnregisterService(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))})) != 0
	}
	return false
}

func QDBusConnection_DisconnectFromBus(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))})
}

func (ptr *QDBusConnection) DisconnectFromBus(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))})
}

func QDBusConnection_DisconnectFromPeer(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))})
}

func (ptr *QDBusConnection) DisconnectFromPeer(name string) {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))})
}

func (ptr *QDBusConnection) Swap(other QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnection_Swap(ptr.Pointer(), PointerFromQDBusConnection(other))
	}
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QDBusConnection_UnregisterObject(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.longlong(mode))
	}
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
	if ptr.Pointer() != nil {
		C.QDBusConnection_DestroyQDBusConnection(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusConnectionInterfaceFromPointer(C.QDBusConnection_Interface(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusErrorFromPointer(C.QDBusConnection_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) Call(message QDBusMessage_ITF, mode QDBus__CallMode, timeout int) *QDBusMessage {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusMessageFromPointer(C.QDBusConnection_Call(ptr.Pointer(), PointerFromQDBusMessage(message), C.longlong(mode), C.int(int32(timeout))))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) AsyncCall(message QDBusMessage_ITF, timeout int) *QDBusPendingCall {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusPendingCallFromPointer(C.QDBusConnection_AsyncCall(ptr.Pointer(), PointerFromQDBusMessage(message), C.int(int32(timeout))))
		runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		tmpValue := core.NewQObjectFromPointer(C.QDBusConnection_ObjectRegisteredAt(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusConnection) BaseService() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusConnection_BaseService(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusConnection_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {
	if ptr.Pointer() != nil {
		var returnMethodC *C.char
		if returnMethod != "" {
			returnMethodC = C.CString(returnMethod)
			defer C.free(unsafe.Pointer(returnMethodC))
		}
		var errorMethodC *C.char
		if errorMethod != "" {
			errorMethodC = C.CString(errorMethod)
			defer C.free(unsafe.Pointer(errorMethodC))
		}
		return int8(C.QDBusConnection_CallWithCallback(ptr.Pointer(), PointerFromQDBusMessage(message), core.PointerFromQObject(receiver), returnMethodC, errorMethodC, C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QDBusConnection) IsConnected() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusConnection_IsConnected(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusConnection_Send(ptr.Pointer(), PointerFromQDBusMessage(message))) != 0
	}
	return false
}

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func (ptr *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return ptr
}

func (ptr *QDBusConnectionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusConnectionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusAbstractInterface_PTR().SetPointer(p)
	}
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusConnectionInterface) {
	n = new(QDBusConnectionInterface)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDBusConnectionInterface__RegisterServiceReply
//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        QDBusConnectionInterface__RegisterServiceReply = QDBusConnectionInterface__RegisterServiceReply(2)
)

//go:generate stringer -type=QDBusConnectionInterface__ServiceQueueOptions
//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService QDBusConnectionInterface__ServiceQueueOptions = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//go:generate stringer -type=QDBusConnectionInterface__ServiceReplacementOptions
//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement QDBusConnectionInterface__ServiceReplacementOptions = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     QDBusConnectionInterface__ServiceReplacementOptions = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

//export callbackQDBusConnectionInterface_CallWithCallbackFailed
func callbackQDBusConnectionInterface_CallWithCallbackFailed(ptr unsafe.Pointer, error unsafe.Pointer, call unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "callWithCallbackFailed"); signal != nil {
		signal.(func(*QDBusError, *QDBusMessage))(NewQDBusErrorFromPointer(error), NewQDBusMessageFromPointer(call))
	}

}

func (ptr *QDBusConnectionInterface) ConnectCallWithCallbackFailed(f func(error *QDBusError, call *QDBusMessage)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "callWithCallbackFailed") {
			C.QDBusConnectionInterface_ConnectCallWithCallbackFailed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "callWithCallbackFailed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "callWithCallbackFailed", func(error *QDBusError, call *QDBusMessage) {
				signal.(func(*QDBusError, *QDBusMessage))(error, call)
				f(error, call)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "callWithCallbackFailed", f)
		}
	}
}

func (ptr *QDBusConnectionInterface) DisconnectCallWithCallbackFailed() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectCallWithCallbackFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "callWithCallbackFailed")
	}
}

func (ptr *QDBusConnectionInterface) CallWithCallbackFailed(error QDBusError_ITF, call QDBusMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_CallWithCallbackFailed(ptr.Pointer(), PointerFromQDBusError(error), PointerFromQDBusMessage(call))
	}
}

//export callbackQDBusConnectionInterface_ServiceRegistered
func callbackQDBusConnectionInterface_ServiceRegistered(ptr unsafe.Pointer, service C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "serviceRegistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(service))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(service string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceRegistered") {
			C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceRegistered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceRegistered", func(service string) {
				signal.(func(string))(service)
				f(service)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceRegistered", f)
		}
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceRegistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceRegistered(service string) {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		C.QDBusConnectionInterface_ServiceRegistered(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))})
	}
}

//export callbackQDBusConnectionInterface_ServiceUnregistered
func callbackQDBusConnectionInterface_ServiceUnregistered(ptr unsafe.Pointer, service C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "serviceUnregistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(service))
	}

}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(service string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceUnregistered") {
			C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceUnregistered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceUnregistered", func(service string) {
				signal.(func(string))(service)
				f(service)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceUnregistered", f)
		}
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceUnregistered")
	}
}

func (ptr *QDBusConnectionInterface) ServiceUnregistered(service string) {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		C.QDBusConnectionInterface_ServiceUnregistered(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))})
	}
}

type QDBusContext struct {
	ptr unsafe.Pointer
}

type QDBusContext_ITF interface {
	QDBusContext_PTR() *QDBusContext
}

func (ptr *QDBusContext) QDBusContext_PTR() *QDBusContext {
	return ptr
}

func (ptr *QDBusContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusContext(ptr QDBusContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContext_PTR().Pointer()
	}
	return nil
}

func NewQDBusContextFromPointer(ptr unsafe.Pointer) (n *QDBusContext) {
	n = new(QDBusContext)
	n.SetPointer(ptr)
	return
}
func NewQDBusContext() *QDBusContext {
	tmpValue := NewQDBusContextFromPointer(C.QDBusContext_NewQDBusContext())
	runtime.SetFinalizer(tmpValue, (*QDBusContext).DestroyQDBusContext)
	return tmpValue
}

func (ptr *QDBusContext) DestroyQDBusContext() {
	if ptr.Pointer() != nil {
		C.QDBusContext_DestroyQDBusContext(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusContext) Connection() *QDBusConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusConnectionFromPointer(C.QDBusContext_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusContext_CalledFromDBus(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusContext_IsDelayedReply(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusContext) Message() *QDBusMessage {
	if ptr.Pointer() != nil {
		return NewQDBusMessageFromPointer(C.QDBusContext_Message(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QDBusContext_SendErrorReply2(ptr.Pointer(), C.longlong(ty), C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QDBusContext_SendErrorReply(ptr.Pointer(), C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

type QDBusError struct {
	ptr unsafe.Pointer
}

type QDBusError_ITF interface {
	QDBusError_PTR() *QDBusError
}

func (ptr *QDBusError) QDBusError_PTR() *QDBusError {
	return ptr
}

func (ptr *QDBusError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusError(ptr QDBusError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusError_PTR().Pointer()
	}
	return nil
}

func NewQDBusErrorFromPointer(ptr unsafe.Pointer) (n *QDBusError) {
	n = new(QDBusError)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusError) DestroyQDBusError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QDBusError__ErrorType
//QDBusError::ErrorType
type QDBusError__ErrorType int64

const (
	QDBusError__NoError           QDBusError__ErrorType = QDBusError__ErrorType(0)
	QDBusError__Other             QDBusError__ErrorType = QDBusError__ErrorType(1)
	QDBusError__Failed            QDBusError__ErrorType = QDBusError__ErrorType(2)
	QDBusError__NoMemory          QDBusError__ErrorType = QDBusError__ErrorType(3)
	QDBusError__ServiceUnknown    QDBusError__ErrorType = QDBusError__ErrorType(4)
	QDBusError__NoReply           QDBusError__ErrorType = QDBusError__ErrorType(5)
	QDBusError__BadAddress        QDBusError__ErrorType = QDBusError__ErrorType(6)
	QDBusError__NotSupported      QDBusError__ErrorType = QDBusError__ErrorType(7)
	QDBusError__LimitsExceeded    QDBusError__ErrorType = QDBusError__ErrorType(8)
	QDBusError__AccessDenied      QDBusError__ErrorType = QDBusError__ErrorType(9)
	QDBusError__NoServer          QDBusError__ErrorType = QDBusError__ErrorType(10)
	QDBusError__Timeout           QDBusError__ErrorType = QDBusError__ErrorType(11)
	QDBusError__NoNetwork         QDBusError__ErrorType = QDBusError__ErrorType(12)
	QDBusError__AddressInUse      QDBusError__ErrorType = QDBusError__ErrorType(13)
	QDBusError__Disconnected      QDBusError__ErrorType = QDBusError__ErrorType(14)
	QDBusError__InvalidArgs       QDBusError__ErrorType = QDBusError__ErrorType(15)
	QDBusError__UnknownMethod     QDBusError__ErrorType = QDBusError__ErrorType(16)
	QDBusError__TimedOut          QDBusError__ErrorType = QDBusError__ErrorType(17)
	QDBusError__InvalidSignature  QDBusError__ErrorType = QDBusError__ErrorType(18)
	QDBusError__UnknownInterface  QDBusError__ErrorType = QDBusError__ErrorType(19)
	QDBusError__UnknownObject     QDBusError__ErrorType = QDBusError__ErrorType(20)
	QDBusError__UnknownProperty   QDBusError__ErrorType = QDBusError__ErrorType(21)
	QDBusError__PropertyReadOnly  QDBusError__ErrorType = QDBusError__ErrorType(22)
	QDBusError__InternalError     QDBusError__ErrorType = QDBusError__ErrorType(23)
	QDBusError__InvalidService    QDBusError__ErrorType = QDBusError__ErrorType(24)
	QDBusError__InvalidObjectPath QDBusError__ErrorType = QDBusError__ErrorType(25)
	QDBusError__InvalidInterface  QDBusError__ErrorType = QDBusError__ErrorType(26)
	QDBusError__InvalidMember     QDBusError__ErrorType = QDBusError__ErrorType(27)
)

func NewQDBusError(other QDBusError_ITF) *QDBusError {
	tmpValue := NewQDBusErrorFromPointer(C.QDBusError_NewQDBusError(PointerFromQDBusError(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
	return tmpValue
}

func QDBusError_ErrorString(error QDBusError__ErrorType) string {
	return cGoUnpackString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) ErrorString(error QDBusError__ErrorType) string {
	return cGoUnpackString(C.QDBusError_QDBusError_ErrorString(C.longlong(error)))
}

func (ptr *QDBusError) Swap(other QDBusError_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusError_Swap(ptr.Pointer(), PointerFromQDBusError(other))
	}
}

func (ptr *QDBusError) Type() QDBusError__ErrorType {
	if ptr.Pointer() != nil {
		return QDBusError__ErrorType(C.QDBusError_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusError) Message() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusError_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusError_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusError) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusError_IsValid(ptr.Pointer())) != 0
	}
	return false
}

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func (ptr *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return ptr
}

func (ptr *QDBusInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusAbstractInterface_PTR().SetPointer(p)
	}
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) (n *QDBusInterface) {
	n = new(QDBusInterface)
	n.SetPointer(ptr)
	return
}
func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	tmpValue := NewQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDBusInterface_DestroyQDBusInterface
func callbackQDBusInterface_DestroyQDBusInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusInterfaceFromPointer(ptr).DestroyQDBusInterfaceDefault()
	}
}

func (ptr *QDBusInterface) ConnectDestroyQDBusInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusInterface", f)
		}
	}
}

func (ptr *QDBusInterface) DisconnectDestroyQDBusInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusInterface")
	}
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusInterface) DestroyQDBusInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QDBusMessage struct {
	ptr unsafe.Pointer
}

type QDBusMessage_ITF interface {
	QDBusMessage_PTR() *QDBusMessage
}

func (ptr *QDBusMessage) QDBusMessage_PTR() *QDBusMessage {
	return ptr
}

func (ptr *QDBusMessage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusMessage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusMessage(ptr QDBusMessage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusMessage_PTR().Pointer()
	}
	return nil
}

func NewQDBusMessageFromPointer(ptr unsafe.Pointer) (n *QDBusMessage) {
	n = new(QDBusMessage)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDBusMessage__MessageType
//QDBusMessage::MessageType
type QDBusMessage__MessageType int64

const (
	QDBusMessage__InvalidMessage    QDBusMessage__MessageType = QDBusMessage__MessageType(0)
	QDBusMessage__MethodCallMessage QDBusMessage__MessageType = QDBusMessage__MessageType(1)
	QDBusMessage__ReplyMessage      QDBusMessage__MessageType = QDBusMessage__MessageType(2)
	QDBusMessage__ErrorMessage      QDBusMessage__MessageType = QDBusMessage__MessageType(3)
	QDBusMessage__SignalMessage     QDBusMessage__MessageType = QDBusMessage__MessageType(4)
)

func QDBusMessage_CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	var msgC *C.char
	if msg != "" {
		msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	var msgC *C.char
	if msg != "" {
		msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError3(C.longlong(ty), C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError2(error QDBusError_ITF) *QDBusMessage {
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError2(error QDBusError_ITF) *QDBusMessage {
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError2(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateError(name string, msg string) *QDBusMessage {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var msgC *C.char
	if msg != "" {
		msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateError(name string, msg string) *QDBusMessage {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var msgC *C.char
	if msg != "" {
		msgC = C.CString(msg)
		defer C.free(unsafe.Pointer(msgC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateError(C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var methodC *C.char
	if method != "" {
		methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateMethodCall(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateMethodCall(service string, path string, interfa string, method string) *QDBusMessage {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var methodC *C.char
	if method != "" {
		methodC = C.CString(method)
		defer C.free(unsafe.Pointer(methodC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateMethodCall(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: methodC, len: C.longlong(len(method))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateSignal(path string, interfa string, name string) *QDBusMessage {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateSignal(C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateSignal(path string, interfa string, name string) *QDBusMessage {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateSignal(C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func QDBusMessage_CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateTargetedSignal(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) CreateTargetedSignal(service string, path string, interfa string, name string) *QDBusMessage {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	var interfaC *C.char
	if interfa != "" {
		interfaC = C.CString(interfa)
		defer C.free(unsafe.Pointer(interfaC))
	}
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_QDBusMessage_CreateTargetedSignal(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}, C.struct_QtDBus_PackedString{data: interfaC, len: C.longlong(len(interfa))}, C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func NewQDBusMessage() *QDBusMessage {
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage())
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func NewQDBusMessage2(other QDBusMessage_ITF) *QDBusMessage {
	tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_NewQDBusMessage2(PointerFromQDBusMessage(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
	return tmpValue
}

func (ptr *QDBusMessage) SetArguments(arguments []*core.QVariant) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetArguments(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQDBusMessageFromPointer(NewQDBusMessageFromPointer(nil).__setArguments_arguments_newList())
			for _, v := range arguments {
				tmpList.__setArguments_arguments_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QDBusMessage) SetAutoStartService(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetAutoStartService(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) SetInteractiveAuthorizationAllowed(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetInteractiveAuthorizationAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) Swap(other QDBusMessage_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_Swap(ptr.Pointer(), PointerFromQDBusMessage(other))
	}
}

func (ptr *QDBusMessage) DestroyQDBusMessage() {
	if ptr.Pointer() != nil {
		C.QDBusMessage_DestroyQDBusMessage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusMessage) CreateErrorReply3(ty QDBusError__ErrorType, msg string) *QDBusMessage {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply3(ptr.Pointer(), C.longlong(ty), C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateErrorReply2(error QDBusError_ITF) *QDBusMessage {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply2(ptr.Pointer(), PointerFromQDBusError(error)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateErrorReply(name string, msg string) *QDBusMessage {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_CreateErrorReply(ptr.Pointer(), C.struct_QtDBus_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_QtDBus_PackedString{data: msgC, len: C.longlong(len(msg))}))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateReply(arguments []*core.QVariant) *QDBusMessage {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_CreateReply(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQDBusMessageFromPointer(NewQDBusMessageFromPointer(nil).__createReply_arguments_newList())
			for _, v := range arguments {
				tmpList.__createReply_arguments_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) CreateReply2(argument core.QVariant_ITF) *QDBusMessage {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusMessageFromPointer(C.QDBusMessage_CreateReply2(ptr.Pointer(), core.PointerFromQVariant(argument)))
		runtime.SetFinalizer(tmpValue, (*QDBusMessage).DestroyQDBusMessage)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) Type() QDBusMessage__MessageType {
	if ptr.Pointer() != nil {
		return QDBusMessage__MessageType(C.QDBusMessage_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusMessage) Arguments() []*core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDBus_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQDBusMessageFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__arguments_atList(i)
			}
			return out
		}(C.QDBusMessage_Arguments(ptr.Pointer()))
	}
	return make([]*core.QVariant, 0)
}

func (ptr *QDBusMessage) ErrorMessage() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) ErrorName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_ErrorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Interface() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Member() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Member(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Service() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) Signature() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusMessage_Signature(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusMessage) AutoStartService() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusMessage_AutoStartService(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusMessage_IsDelayedReply(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsInteractiveAuthorizationAllowed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusMessage_IsInteractiveAuthorizationAllowed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) IsReplyRequired() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusMessage_IsReplyRequired(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusMessage) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusMessage_SetDelayedReply(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QDBusMessage) __setArguments_arguments_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusMessage___setArguments_arguments_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) __setArguments_arguments_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusMessage___setArguments_arguments_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusMessage) __setArguments_arguments_newList() unsafe.Pointer {
	return C.QDBusMessage___setArguments_arguments_newList(ptr.Pointer())
}

func (ptr *QDBusMessage) __createReply_arguments_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusMessage___createReply_arguments_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) __createReply_arguments_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusMessage___createReply_arguments_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusMessage) __createReply_arguments_newList() unsafe.Pointer {
	return C.QDBusMessage___createReply_arguments_newList(ptr.Pointer())
}

func (ptr *QDBusMessage) __arguments_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusMessage___arguments_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusMessage) __arguments_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusMessage___arguments_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QDBusMessage) __arguments_newList() unsafe.Pointer {
	return C.QDBusMessage___arguments_newList(ptr.Pointer())
}

type QDBusObjectPath struct {
	ptr unsafe.Pointer
}

type QDBusObjectPath_ITF interface {
	QDBusObjectPath_PTR() *QDBusObjectPath
}

func (ptr *QDBusObjectPath) QDBusObjectPath_PTR() *QDBusObjectPath {
	return ptr
}

func (ptr *QDBusObjectPath) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusObjectPath) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPath_PTR().Pointer()
	}
	return nil
}

func NewQDBusObjectPathFromPointer(ptr unsafe.Pointer) (n *QDBusObjectPath) {
	n = new(QDBusObjectPath)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusObjectPath) DestroyQDBusObjectPath() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDBusObjectPath() *QDBusObjectPath {
	tmpValue := NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath())
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {
	tmpValue := NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath3(core.PointerFromQLatin1String(path)))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath5(p string) *QDBusObjectPath {
	var pC *C.char
	if p != "" {
		pC = C.CString(p)
		defer C.free(unsafe.Pointer(pC))
	}
	tmpValue := NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath5(C.struct_QtDBus_PackedString{data: pC, len: C.longlong(len(p))}))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	tmpValue := NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath4(C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	var pathC *C.char
	if path != "" {
		pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
	}
	tmpValue := NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath2(pathC))
	runtime.SetFinalizer(tmpValue, (*QDBusObjectPath).DestroyQDBusObjectPath)
	return tmpValue
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QDBusObjectPath_SetPath(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QDBusObjectPath) Swap(other QDBusObjectPath_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusObjectPath_Swap(ptr.Pointer(), PointerFromQDBusObjectPath(other))
	}
}

func (ptr *QDBusObjectPath) Path() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusObjectPath_Path(ptr.Pointer()))
	}
	return ""
}

type QDBusPendingCall struct {
	ptr unsafe.Pointer
}

type QDBusPendingCall_ITF interface {
	QDBusPendingCall_PTR() *QDBusPendingCall
}

func (ptr *QDBusPendingCall) QDBusPendingCall_PTR() *QDBusPendingCall {
	return ptr
}

func (ptr *QDBusPendingCall) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusPendingCall) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallFromPointer(ptr unsafe.Pointer) (n *QDBusPendingCall) {
	n = new(QDBusPendingCall)
	n.SetPointer(ptr)
	return
}
func QDBusPendingCall_FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	tmpValue := NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromCompletedCall(msg QDBusMessage_ITF) *QDBusPendingCall {
	tmpValue := NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromCompletedCall(PointerFromQDBusMessage(msg)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func QDBusPendingCall_FromError(error QDBusError_ITF) *QDBusPendingCall {
	tmpValue := NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) FromError(error QDBusError_ITF) *QDBusPendingCall {
	tmpValue := NewQDBusPendingCallFromPointer(C.QDBusPendingCall_QDBusPendingCall_FromError(PointerFromQDBusError(error)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func NewQDBusPendingCall(other QDBusPendingCall_ITF) *QDBusPendingCall {
	tmpValue := NewQDBusPendingCallFromPointer(C.QDBusPendingCall_NewQDBusPendingCall(PointerFromQDBusPendingCall(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusPendingCall).DestroyQDBusPendingCall)
	return tmpValue
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(ptr.Pointer(), PointerFromQDBusPendingCall(other))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCall_DestroyQDBusPendingCall(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QDBusPendingCallWatcher struct {
	core.QObject
	QDBusPendingCall
}

type QDBusPendingCallWatcher_ITF interface {
	core.QObject_ITF
	QDBusPendingCall_ITF
	QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher
}

func (ptr *QDBusPendingCallWatcher) QDBusPendingCallWatcher_PTR() *QDBusPendingCallWatcher {
	return ptr
}

func (ptr *QDBusPendingCallWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QDBusPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQDBusPendingCallWatcher(ptr QDBusPendingCallWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCallWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallWatcherFromPointer(ptr unsafe.Pointer) (n *QDBusPendingCallWatcher) {
	n = new(QDBusPendingCallWatcher)
	n.SetPointer(ptr)
	return
}
func NewQDBusPendingCallWatcher(call QDBusPendingCall_ITF, parent core.QObject_ITF) *QDBusPendingCallWatcher {
	tmpValue := NewQDBusPendingCallWatcherFromPointer(C.QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(PointerFromQDBusPendingCall(call), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDBusPendingCallWatcher_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusPendingCallWatcher_QDBusPendingCallWatcher_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusPendingCallWatcher) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusPendingCallWatcher_QDBusPendingCallWatcher_Tr(sC, cC, C.int(int32(n))))
}

func QDBusPendingCallWatcher_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusPendingCallWatcher_QDBusPendingCallWatcher_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusPendingCallWatcher) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusPendingCallWatcher_QDBusPendingCallWatcher_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQDBusPendingCallWatcher_Finished
func callbackQDBusPendingCallWatcher_Finished(ptr unsafe.Pointer, self unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(*QDBusPendingCallWatcher))(NewQDBusPendingCallWatcherFromPointer(self))
	}

}

func (ptr *QDBusPendingCallWatcher) ConnectFinished(f func(self *QDBusPendingCallWatcher)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QDBusPendingCallWatcher_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func(self *QDBusPendingCallWatcher) {
				signal.(func(*QDBusPendingCallWatcher))(self)
				f(self)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QDBusPendingCallWatcher) Finished(self QDBusPendingCallWatcher_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_Finished(ptr.Pointer(), PointerFromQDBusPendingCallWatcher(self))
	}
}

func (ptr *QDBusPendingCallWatcher) WaitForFinished() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_WaitForFinished(ptr.Pointer())
	}
}

//export callbackQDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher
func callbackQDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusPendingCallWatcher"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DestroyQDBusPendingCallWatcherDefault()
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectDestroyQDBusPendingCallWatcher(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusPendingCallWatcher"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusPendingCallWatcher", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusPendingCallWatcher", f)
		}
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectDestroyQDBusPendingCallWatcher() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusPendingCallWatcher")
	}
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusPendingCallWatcher) DestroyQDBusPendingCallWatcherDefault() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcherDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusPendingCallWatcher) IsFinished() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusPendingCallWatcher_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_MetaObject
func callbackQDBusPendingCallWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusPendingCallWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusPendingCallWatcher) ConnectMetaObject(f func() *core.QMetaObject) {
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

func (ptr *QDBusPendingCallWatcher) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "metaObject")
	}
}

func (ptr *QDBusPendingCallWatcher) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusPendingCallWatcher_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusPendingCallWatcher_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusPendingCallWatcher___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusPendingCallWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusPendingCallWatcher___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusPendingCallWatcher) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusPendingCallWatcher___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusPendingCallWatcher) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusPendingCallWatcher___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusPendingCallWatcher) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusPendingCallWatcher___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusPendingCallWatcher) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusPendingCallWatcher___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusPendingCallWatcher) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusPendingCallWatcher___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusPendingCallWatcher) __findChildren_newList() unsafe.Pointer {
	return C.QDBusPendingCallWatcher___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusPendingCallWatcher) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusPendingCallWatcher___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusPendingCallWatcher) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusPendingCallWatcher) __children_newList() unsafe.Pointer {
	return C.QDBusPendingCallWatcher___children_newList(ptr.Pointer())
}

//export callbackQDBusPendingCallWatcher_Event
func callbackQDBusPendingCallWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusPendingCallWatcher) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusPendingCallWatcher_Event(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusPendingCallWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_EventFilter
func callbackQDBusPendingCallWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusPendingCallWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusPendingCallWatcher) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusPendingCallWatcher_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QDBusPendingCallWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusPendingCallWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusPendingCallWatcher_ChildEvent
func callbackQDBusPendingCallWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_ConnectNotify
func callbackQDBusPendingCallWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_CustomEvent
func callbackQDBusPendingCallWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusPendingCallWatcher_DeleteLater
func callbackQDBusPendingCallWatcher_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusPendingCallWatcher) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusPendingCallWatcher_Destroyed
func callbackQDBusPendingCallWatcher_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusPendingCallWatcher_DisconnectNotify
func callbackQDBusPendingCallWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDBusPendingCallWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusPendingCallWatcher_ObjectNameChanged
func callbackQDBusPendingCallWatcher_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusPendingCallWatcher_TimerEvent
func callbackQDBusPendingCallWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusPendingCallWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusPendingCallWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusPendingCallWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDBusPendingReply struct {
	QDBusPendingCall
}

type QDBusPendingReply_ITF interface {
	QDBusPendingCall_ITF
	QDBusPendingReply_PTR() *QDBusPendingReply
}

func (ptr *QDBusPendingReply) QDBusPendingReply_PTR() *QDBusPendingReply {
	return ptr
}

func (ptr *QDBusPendingReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusPendingReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDBusPendingCall_PTR().SetPointer(p)
	}
}

func PointerFromQDBusPendingReply(ptr QDBusPendingReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingReplyFromPointer(ptr unsafe.Pointer) (n *QDBusPendingReply) {
	n = new(QDBusPendingReply)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusPendingReply) DestroyQDBusPendingReply() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QDBusPendingReply__anonymous
//QDBusPendingReply::anonymous
type QDBusPendingReply__anonymous int64

const (
	QDBusPendingReply__Count QDBusPendingReply__anonymous = QDBusPendingReply__anonymous(0)
)

type QDBusPendingReplyTypes struct {
	ptr unsafe.Pointer
}

type QDBusPendingReplyTypes_ITF interface {
	QDBusPendingReplyTypes_PTR() *QDBusPendingReplyTypes
}

func (ptr *QDBusPendingReplyTypes) QDBusPendingReplyTypes_PTR() *QDBusPendingReplyTypes {
	return ptr
}

func (ptr *QDBusPendingReplyTypes) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusPendingReplyTypes) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusPendingReplyTypes(ptr QDBusPendingReplyTypes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingReplyTypes_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingReplyTypesFromPointer(ptr unsafe.Pointer) (n *QDBusPendingReplyTypes) {
	n = new(QDBusPendingReplyTypes)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusPendingReplyTypes) DestroyQDBusPendingReplyTypes() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QDBusPendingReplyTypes_MetaTypeFor(vqv core.QVariant_ITF) int {
	return int(int32(C.QDBusPendingReplyTypes_QDBusPendingReplyTypes_MetaTypeFor(core.PointerFromQVariant(vqv))))
}

func (ptr *QDBusPendingReplyTypes) MetaTypeFor(vqv core.QVariant_ITF) int {
	return int(int32(C.QDBusPendingReplyTypes_QDBusPendingReplyTypes_MetaTypeFor(core.PointerFromQVariant(vqv))))
}

type QDBusReply struct {
	ptr unsafe.Pointer
}

type QDBusReply_ITF interface {
	QDBusReply_PTR() *QDBusReply
}

func (ptr *QDBusReply) QDBusReply_PTR() *QDBusReply {
	return ptr
}

func (ptr *QDBusReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusReply(ptr QDBusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusReply_PTR().Pointer()
	}
	return nil
}

func NewQDBusReplyFromPointer(ptr unsafe.Pointer) (n *QDBusReply) {
	n = new(QDBusReply)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusReply) DestroyQDBusReply() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func (ptr *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return ptr
}

func (ptr *QDBusServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) (n *QDBusServer) {
	n = new(QDBusServer)
	n.SetPointer(ptr)
	return
}
func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {
	tmpValue := NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	tmpValue := NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(C.struct_QtDBus_PackedString{data: addressC, len: C.longlong(len(address))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDBusServer_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServer_QDBusServer_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusServer) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServer_QDBusServer_Tr(sC, cC, C.int(int32(n))))
}

func QDBusServer_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServer_QDBusServer_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusServer) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServer_QDBusServer_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQDBusServer_NewConnection
func callbackQDBusServer_NewConnection(ptr unsafe.Pointer, connection unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		signal.(func(*QDBusConnection))(NewQDBusConnectionFromPointer(connection))
	}

}

func (ptr *QDBusServer) ConnectNewConnection(f func(connection *QDBusConnection)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QDBusServer_ConnectNewConnection(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", func(connection *QDBusConnection) {
				signal.(func(*QDBusConnection))(connection)
				f(connection)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", f)
		}
	}
}

func (ptr *QDBusServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConnection")
	}
}

func (ptr *QDBusServer) NewConnection(connection QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_NewConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(value))))
	}
}

//export callbackQDBusServer_DestroyQDBusServer
func callbackQDBusServer_DestroyQDBusServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusServer"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServerFromPointer(ptr).DestroyQDBusServerDefault()
	}
}

func (ptr *QDBusServer) ConnectDestroyQDBusServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusServer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusServer", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusServer", f)
		}
	}
}

func (ptr *QDBusServer) DisconnectDestroyQDBusServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusServer")
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusServer) DestroyQDBusServerDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusServer) LastError() *QDBusError {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusErrorFromPointer(C.QDBusServer_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusError).DestroyQDBusError)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) Address() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServer_IsConnected(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQDBusServer_MetaObject
func callbackQDBusServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusServer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusServer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusServer) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServer___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServer) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusServer___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServer) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusServer___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServer) __findChildren_newList() unsafe.Pointer {
	return C.QDBusServer___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServer) __children_newList() unsafe.Pointer {
	return C.QDBusServer___children_newList(ptr.Pointer())
}

//export callbackQDBusServer_Event
func callbackQDBusServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusServer_EventFilter
func callbackQDBusServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusServer_ChildEvent
func callbackQDBusServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServer_ConnectNotify
func callbackQDBusServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_CustomEvent
func callbackQDBusServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServer_DeleteLater
func callbackQDBusServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusServer_Destroyed
func callbackQDBusServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusServer_DisconnectNotify
func callbackQDBusServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServer_ObjectNameChanged
func callbackQDBusServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusServer_TimerEvent
func callbackQDBusServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDBusServiceWatcher struct {
	core.QObject
}

type QDBusServiceWatcher_ITF interface {
	core.QObject_ITF
	QDBusServiceWatcher_PTR() *QDBusServiceWatcher
}

func (ptr *QDBusServiceWatcher) QDBusServiceWatcher_PTR() *QDBusServiceWatcher {
	return ptr
}

func (ptr *QDBusServiceWatcher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusServiceWatcher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusServiceWatcher(ptr QDBusServiceWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServiceWatcher_PTR().Pointer()
	}
	return nil
}

func NewQDBusServiceWatcherFromPointer(ptr unsafe.Pointer) (n *QDBusServiceWatcher) {
	n = new(QDBusServiceWatcher)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QDBusServiceWatcher__WatchModeFlag
//QDBusServiceWatcher::WatchModeFlag
type QDBusServiceWatcher__WatchModeFlag int64

const (
	QDBusServiceWatcher__WatchForRegistration   QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x01)
	QDBusServiceWatcher__WatchForUnregistration QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x02)
	QDBusServiceWatcher__WatchForOwnerChange    QDBusServiceWatcher__WatchModeFlag = QDBusServiceWatcher__WatchModeFlag(0x03)
)

func NewQDBusServiceWatcher(parent core.QObject_ITF) *QDBusServiceWatcher {
	tmpValue := NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDBusServiceWatcher2(service string, connection QDBusConnection_ITF, watchMode QDBusServiceWatcher__WatchModeFlag, parent core.QObject_ITF) *QDBusServiceWatcher {
	var serviceC *C.char
	if service != "" {
		serviceC = C.CString(service)
		defer C.free(unsafe.Pointer(serviceC))
	}
	tmpValue := NewQDBusServiceWatcherFromPointer(C.QDBusServiceWatcher_NewQDBusServiceWatcher2(C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))}, PointerFromQDBusConnection(connection), C.longlong(watchMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDBusServiceWatcher_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServiceWatcher_QDBusServiceWatcher_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusServiceWatcher) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServiceWatcher_QDBusServiceWatcher_Tr(sC, cC, C.int(int32(n))))
}

func QDBusServiceWatcher_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServiceWatcher_QDBusServiceWatcher_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusServiceWatcher) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusServiceWatcher_QDBusServiceWatcher_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusServiceWatcher) RemoveWatchedService(service string) bool {
	if ptr.Pointer() != nil {
		var serviceC *C.char
		if service != "" {
			serviceC = C.CString(service)
			defer C.free(unsafe.Pointer(serviceC))
		}
		return int8(C.QDBusServiceWatcher_RemoveWatchedService(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceC, len: C.longlong(len(service))})) != 0
	}
	return false
}

func (ptr *QDBusServiceWatcher) AddWatchedService(newService string) {
	if ptr.Pointer() != nil {
		var newServiceC *C.char
		if newService != "" {
			newServiceC = C.CString(newService)
			defer C.free(unsafe.Pointer(newServiceC))
		}
		C.QDBusServiceWatcher_AddWatchedService(ptr.Pointer(), C.struct_QtDBus_PackedString{data: newServiceC, len: C.longlong(len(newService))})
	}
}

//export callbackQDBusServiceWatcher_ServiceOwnerChanged
func callbackQDBusServiceWatcher_ServiceOwnerChanged(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString, oldOwner C.struct_QtDBus_PackedString, newOwner C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "serviceOwnerChanged"); signal != nil {
		signal.(func(string, string, string))(cGoUnpackString(serviceName), cGoUnpackString(oldOwner), cGoUnpackString(newOwner))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceOwnerChanged(f func(serviceName string, oldOwner string, newOwner string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceOwnerChanged") {
			C.QDBusServiceWatcher_ConnectServiceOwnerChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceOwnerChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceOwnerChanged", func(serviceName string, oldOwner string, newOwner string) {
				signal.(func(string, string, string))(serviceName, oldOwner, newOwner)
				f(serviceName, oldOwner, newOwner)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceOwnerChanged", f)
		}
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceOwnerChanged() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceOwnerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceOwnerChanged")
	}
}

func (ptr *QDBusServiceWatcher) ServiceOwnerChanged(serviceName string, oldOwner string, newOwner string) {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		var oldOwnerC *C.char
		if oldOwner != "" {
			oldOwnerC = C.CString(oldOwner)
			defer C.free(unsafe.Pointer(oldOwnerC))
		}
		var newOwnerC *C.char
		if newOwner != "" {
			newOwnerC = C.CString(newOwner)
			defer C.free(unsafe.Pointer(newOwnerC))
		}
		C.QDBusServiceWatcher_ServiceOwnerChanged(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))}, C.struct_QtDBus_PackedString{data: oldOwnerC, len: C.longlong(len(oldOwner))}, C.struct_QtDBus_PackedString{data: newOwnerC, len: C.longlong(len(newOwner))})
	}
}

//export callbackQDBusServiceWatcher_ServiceRegistered
func callbackQDBusServiceWatcher_ServiceRegistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "serviceRegistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceRegistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceRegistered") {
			C.QDBusServiceWatcher_ConnectServiceRegistered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceRegistered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceRegistered", func(serviceName string) {
				signal.(func(string))(serviceName)
				f(serviceName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceRegistered", f)
		}
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceRegistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceRegistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceRegistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		C.QDBusServiceWatcher_ServiceRegistered(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))})
	}
}

//export callbackQDBusServiceWatcher_ServiceUnregistered
func callbackQDBusServiceWatcher_ServiceUnregistered(ptr unsafe.Pointer, serviceName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "serviceUnregistered"); signal != nil {
		signal.(func(string))(cGoUnpackString(serviceName))
	}

}

func (ptr *QDBusServiceWatcher) ConnectServiceUnregistered(f func(serviceName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceUnregistered") {
			C.QDBusServiceWatcher_ConnectServiceUnregistered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceUnregistered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceUnregistered", func(serviceName string) {
				signal.(func(string))(serviceName)
				f(serviceName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceUnregistered", f)
		}
	}
}

func (ptr *QDBusServiceWatcher) DisconnectServiceUnregistered() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceUnregistered")
	}
}

func (ptr *QDBusServiceWatcher) ServiceUnregistered(serviceName string) {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		C.QDBusServiceWatcher_ServiceUnregistered(ptr.Pointer(), C.struct_QtDBus_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))})
	}
}

func (ptr *QDBusServiceWatcher) SetConnection(connection QDBusConnection_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetConnection(ptr.Pointer(), PointerFromQDBusConnection(connection))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchMode(mode QDBusServiceWatcher__WatchModeFlag) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_SetWatchMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QDBusServiceWatcher) SetWatchedServices(services []string) {
	if ptr.Pointer() != nil {
		servicesC := C.CString(strings.Join(services, "|"))
		defer C.free(unsafe.Pointer(servicesC))
		C.QDBusServiceWatcher_SetWatchedServices(ptr.Pointer(), C.struct_QtDBus_PackedString{data: servicesC, len: C.longlong(len(strings.Join(services, "|")))})
	}
}

//export callbackQDBusServiceWatcher_DestroyQDBusServiceWatcher
func callbackQDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusServiceWatcher"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DestroyQDBusServiceWatcherDefault()
	}
}

func (ptr *QDBusServiceWatcher) ConnectDestroyQDBusServiceWatcher(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusServiceWatcher"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusServiceWatcher", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusServiceWatcher", f)
		}
	}
}

func (ptr *QDBusServiceWatcher) DisconnectDestroyQDBusServiceWatcher() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusServiceWatcher")
	}
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcher() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusServiceWatcher) DestroyQDBusServiceWatcherDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DestroyQDBusServiceWatcherDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusServiceWatcher) Connection() *QDBusConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQDBusConnectionFromPointer(C.QDBusServiceWatcher_Connection(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QDBusConnection).DestroyQDBusConnection)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) WatchMode() QDBusServiceWatcher__WatchModeFlag {
	if ptr.Pointer() != nil {
		return QDBusServiceWatcher__WatchModeFlag(C.QDBusServiceWatcher_WatchMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusServiceWatcher) WatchedServices() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDBusServiceWatcher_WatchedServices(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDBusServiceWatcher_MetaObject
func callbackQDBusServiceWatcher_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusServiceWatcherFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusServiceWatcher) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusServiceWatcher_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusServiceWatcher___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusServiceWatcher) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusServiceWatcher___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusServiceWatcher) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServiceWatcher___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServiceWatcher) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusServiceWatcher___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusServiceWatcher) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServiceWatcher___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServiceWatcher) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusServiceWatcher___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusServiceWatcher) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServiceWatcher___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServiceWatcher) __findChildren_newList() unsafe.Pointer {
	return C.QDBusServiceWatcher___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusServiceWatcher) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusServiceWatcher___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusServiceWatcher) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusServiceWatcher) __children_newList() unsafe.Pointer {
	return C.QDBusServiceWatcher___children_newList(ptr.Pointer())
}

//export callbackQDBusServiceWatcher_Event
func callbackQDBusServiceWatcher_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusServiceWatcher) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServiceWatcher_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_EventFilter
func callbackQDBusServiceWatcher_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusServiceWatcherFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusServiceWatcher) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusServiceWatcher_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusServiceWatcher_ChildEvent
func callbackQDBusServiceWatcher_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusServiceWatcher_ConnectNotify
func callbackQDBusServiceWatcher_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_CustomEvent
func callbackQDBusServiceWatcher_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusServiceWatcher_DeleteLater
func callbackQDBusServiceWatcher_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusServiceWatcher) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusServiceWatcher_Destroyed
func callbackQDBusServiceWatcher_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusServiceWatcher_DisconnectNotify
func callbackQDBusServiceWatcher_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusServiceWatcher) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusServiceWatcher_ObjectNameChanged
func callbackQDBusServiceWatcher_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusServiceWatcher_TimerEvent
func callbackQDBusServiceWatcher_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServiceWatcherFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServiceWatcher) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusServiceWatcher_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QDBusSignature struct {
	ptr unsafe.Pointer
}

type QDBusSignature_ITF interface {
	QDBusSignature_PTR() *QDBusSignature
}

func (ptr *QDBusSignature) QDBusSignature_PTR() *QDBusSignature {
	return ptr
}

func (ptr *QDBusSignature) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusSignature) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusSignature(ptr QDBusSignature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignature_PTR().Pointer()
	}
	return nil
}

func NewQDBusSignatureFromPointer(ptr unsafe.Pointer) (n *QDBusSignature) {
	n = new(QDBusSignature)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusSignature) DestroyQDBusSignature() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDBusSignature() *QDBusSignature {
	tmpValue := NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature())
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {
	tmpValue := NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature3(core.PointerFromQLatin1String(signature)))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature5(sig string) *QDBusSignature {
	var sigC *C.char
	if sig != "" {
		sigC = C.CString(sig)
		defer C.free(unsafe.Pointer(sigC))
	}
	tmpValue := NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature5(C.struct_QtDBus_PackedString{data: sigC, len: C.longlong(len(sig))}))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature4(C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))}))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature2(signatureC))
	runtime.SetFinalizer(tmpValue, (*QDBusSignature).DestroyQDBusSignature)
	return tmpValue
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	if ptr.Pointer() != nil {
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		C.QDBusSignature_SetSignature(ptr.Pointer(), C.struct_QtDBus_PackedString{data: signatureC, len: C.longlong(len(signature))})
	}
}

func (ptr *QDBusSignature) Swap(other QDBusSignature_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusSignature_Swap(ptr.Pointer(), PointerFromQDBusSignature(other))
	}
}

func (ptr *QDBusSignature) Signature() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDBusSignature_Signature(ptr.Pointer()))
	}
	return ""
}

type QDBusUnixFileDescriptor struct {
	ptr unsafe.Pointer
}

type QDBusUnixFileDescriptor_ITF interface {
	QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor
}

func (ptr *QDBusUnixFileDescriptor) QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor {
	return ptr
}

func (ptr *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusUnixFileDescriptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) (n *QDBusUnixFileDescriptor) {
	n = new(QDBusUnixFileDescriptor)
	n.SetPointer(ptr)
	return
}
func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {
	tmpValue := NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	tmpValue := NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	tmpValue := NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(int32(fileDescriptor))))
	runtime.SetFinalizer(tmpValue, (*QDBusUnixFileDescriptor).DestroyQDBusUnixFileDescriptor)
	return tmpValue
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	return int8(C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported()) != 0
}

func (ptr *QDBusUnixFileDescriptor) IsSupported() bool {
	return int8(C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported()) != 0
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(int32(fileDescriptor)))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer())))
	}
	return 0
}

type QDBusVariant struct {
	ptr unsafe.Pointer
}

type QDBusVariant_ITF interface {
	QDBusVariant_PTR() *QDBusVariant
}

func (ptr *QDBusVariant) QDBusVariant_PTR() *QDBusVariant {
	return ptr
}

func (ptr *QDBusVariant) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDBusVariant) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQDBusVariant(ptr QDBusVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVariant_PTR().Pointer()
	}
	return nil
}

func NewQDBusVariantFromPointer(ptr unsafe.Pointer) (n *QDBusVariant) {
	n = new(QDBusVariant)
	n.SetPointer(ptr)
	return
}

func (ptr *QDBusVariant) DestroyQDBusVariant() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQDBusVariant() *QDBusVariant {
	tmpValue := NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant())
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant3(v core.QVariant_ITF) *QDBusVariant {
	tmpValue := NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant3(core.PointerFromQVariant(v)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func NewQDBusVariant2(variant core.QVariant_ITF) *QDBusVariant {
	tmpValue := NewQDBusVariantFromPointer(C.QDBusVariant_NewQDBusVariant2(core.PointerFromQVariant(variant)))
	runtime.SetFinalizer(tmpValue, (*QDBusVariant).DestroyQDBusVariant)
	return tmpValue
}

func (ptr *QDBusVariant) SetVariant(variant core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_SetVariant(ptr.Pointer(), core.PointerFromQVariant(variant))
	}
}

func (ptr *QDBusVariant) Swap(other QDBusVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVariant_Swap(ptr.Pointer(), PointerFromQDBusVariant(other))
	}
}

func (ptr *QDBusVariant) Variant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QDBusVariant_Variant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func (ptr *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDBusVirtualObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) (n *QDBusVirtualObject) {
	n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	return
}
func NewQDBusVirtualObject(parent core.QObject_ITF) *QDBusVirtualObject {
	tmpValue := NewQDBusVirtualObjectFromPointer(C.QDBusVirtualObject_NewQDBusVirtualObject(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDBusVirtualObject_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusVirtualObject_QDBusVirtualObject_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusVirtualObject) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusVirtualObject_QDBusVirtualObject_Tr(sC, cC, C.int(int32(n))))
}

func QDBusVirtualObject_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusVirtualObject_QDBusVirtualObject_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QDBusVirtualObject) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QDBusVirtualObject_QDBusVirtualObject_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQDBusVirtualObject_HandleMessage
func callbackQDBusVirtualObject_HandleMessage(ptr unsafe.Pointer, message unsafe.Pointer, connection unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "handleMessage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QDBusMessage, *QDBusConnection) bool)(NewQDBusMessageFromPointer(message), NewQDBusConnectionFromPointer(connection)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDBusVirtualObject) ConnectHandleMessage(f func(message *QDBusMessage, connection *QDBusConnection) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "handleMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "handleMessage", func(message *QDBusMessage, connection *QDBusConnection) bool {
				signal.(func(*QDBusMessage, *QDBusConnection) bool)(message, connection)
				return f(message, connection)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "handleMessage", f)
		}
	}
}

func (ptr *QDBusVirtualObject) DisconnectHandleMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "handleMessage")
	}
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection))) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_DestroyQDBusVirtualObject
func callbackQDBusVirtualObject_DestroyQDBusVirtualObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDBusVirtualObject"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DestroyQDBusVirtualObjectDefault()
	}
}

func (ptr *QDBusVirtualObject) ConnectDestroyQDBusVirtualObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDBusVirtualObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusVirtualObject", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDBusVirtualObject", f)
		}
	}
}

func (ptr *QDBusVirtualObject) DisconnectDestroyQDBusVirtualObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDBusVirtualObject")
	}
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObjectDefault() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusVirtualObject_Introspect
func callbackQDBusVirtualObject_Introspect(ptr unsafe.Pointer, path C.struct_QtDBus_PackedString) C.struct_QtDBus_PackedString {
	if signal := qt.GetSignal(ptr, "introspect"); signal != nil {
		tempVal := signal.(func(string) string)(cGoUnpackString(path))
		return C.struct_QtDBus_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDBus_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDBusVirtualObject) ConnectIntrospect(f func(path string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "introspect"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "introspect", func(path string) string {
				signal.(func(string) string)(path)
				return f(path)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "introspect", f)
		}
	}
}

func (ptr *QDBusVirtualObject) DisconnectIntrospect() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "introspect")
	}
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		return cGoUnpackString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), C.struct_QtDBus_PackedString{data: pathC, len: C.longlong(len(path))}))
	}
	return ""
}

//export callbackQDBusVirtualObject_MetaObject
func callbackQDBusVirtualObject_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDBusVirtualObjectFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDBusVirtualObject) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDBusVirtualObject_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QDBusVirtualObject___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDBusVirtualObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QDBusVirtualObject___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QDBusVirtualObject) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusVirtualObject___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusVirtualObject) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusVirtualObject) __findChildren_newList2() unsafe.Pointer {
	return C.QDBusVirtualObject___findChildren_newList2(ptr.Pointer())
}

func (ptr *QDBusVirtualObject) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusVirtualObject___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusVirtualObject) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusVirtualObject) __findChildren_newList3() unsafe.Pointer {
	return C.QDBusVirtualObject___findChildren_newList3(ptr.Pointer())
}

func (ptr *QDBusVirtualObject) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusVirtualObject___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusVirtualObject) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusVirtualObject) __findChildren_newList() unsafe.Pointer {
	return C.QDBusVirtualObject___findChildren_newList(ptr.Pointer())
}

func (ptr *QDBusVirtualObject) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QDBusVirtualObject___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDBusVirtualObject) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDBusVirtualObject) __children_newList() unsafe.Pointer {
	return C.QDBusVirtualObject___children_newList(ptr.Pointer())
}

//export callbackQDBusVirtualObject_Event
func callbackQDBusVirtualObject_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDBusVirtualObject) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusVirtualObject_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_EventFilter
func callbackQDBusVirtualObject_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDBusVirtualObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDBusVirtualObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QDBusVirtualObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQDBusVirtualObject_ChildEvent
func callbackQDBusVirtualObject_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDBusVirtualObject_ConnectNotify
func callbackQDBusVirtualObject_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_CustomEvent
func callbackQDBusVirtualObject_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDBusVirtualObject_DeleteLater
func callbackQDBusVirtualObject_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDBusVirtualObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDBusVirtualObject_Destroyed
func callbackQDBusVirtualObject_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDBusVirtualObject_DisconnectNotify
func callbackQDBusVirtualObject_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDBusVirtualObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDBusVirtualObject_ObjectNameChanged
func callbackQDBusVirtualObject_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDBusVirtualObject_TimerEvent
func callbackQDBusVirtualObject_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusVirtualObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusVirtualObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
