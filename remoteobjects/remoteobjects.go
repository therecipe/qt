// +build !minimal

package remoteobjects

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "remoteobjects.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtRemoteObjects_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtRemoteObjects_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QRemoteObjectAbstractPersistedStore struct {
	core.QObject
}

type QRemoteObjectAbstractPersistedStore_ITF interface {
	core.QObject_ITF
	QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore
}

func (ptr *QRemoteObjectAbstractPersistedStore) QRemoteObjectAbstractPersistedStore_PTR() *QRemoteObjectAbstractPersistedStore {
	return ptr
}

func (ptr *QRemoteObjectAbstractPersistedStore) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectAbstractPersistedStore(ptr QRemoteObjectAbstractPersistedStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectAbstractPersistedStore_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectAbstractPersistedStore) {
	n = new(QRemoteObjectAbstractPersistedStore)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectAbstractPersistedStore(parent core.QObject_ITF) *QRemoteObjectAbstractPersistedStore {
	tmpValue := NewQRemoteObjectAbstractPersistedStoreFromPointer(C.QRemoteObjectAbstractPersistedStore_NewQRemoteObjectAbstractPersistedStore(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QRemoteObjectAbstractPersistedStore_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_Tr(sC, cC, C.int(int32(n))))
}

func QRemoteObjectAbstractPersistedStore_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQRemoteObjectAbstractPersistedStore_RestoreProperties
func callbackQRemoteObjectAbstractPersistedStore_RestoreProperties(ptr unsafe.Pointer, repName C.struct_QtRemoteObjects_PackedString, repSig unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "restoreProperties"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__restoreProperties_newList())
			for _, v := range signal.(func(string, *core.QByteArray) []*core.QVariant)(cGoUnpackString(repName), core.NewQByteArrayFromPointer(repSig)) {
				tmpList.__restoreProperties_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__restoreProperties_newList())
		for _, v := range make([]*core.QVariant, 0) {
			tmpList.__restoreProperties_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectRestoreProperties(f func(repName string, repSig *core.QByteArray) []*core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "restoreProperties"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "restoreProperties", func(repName string, repSig *core.QByteArray) []*core.QVariant {
				signal.(func(string, *core.QByteArray) []*core.QVariant)(repName, repSig)
				return f(repName, repSig)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "restoreProperties", f)
		}
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectRestoreProperties() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "restoreProperties")
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) RestoreProperties(repName string, repSig core.QByteArray_ITF) []*core.QVariant {
	if ptr.Pointer() != nil {
		var repNameC *C.char
		if repName != "" {
			repNameC = C.CString(repName)
			defer C.free(unsafe.Pointer(repNameC))
		}
		return func(l C.struct_QtRemoteObjects_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__restoreProperties_atList(i)
			}
			return out
		}(C.QRemoteObjectAbstractPersistedStore_RestoreProperties(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: repNameC, len: C.longlong(len(repName))}, core.PointerFromQByteArray(repSig)))
	}
	return make([]*core.QVariant, 0)
}

//export callbackQRemoteObjectAbstractPersistedStore_SaveProperties
func callbackQRemoteObjectAbstractPersistedStore_SaveProperties(ptr unsafe.Pointer, repName C.struct_QtRemoteObjects_PackedString, repSig unsafe.Pointer, values C.struct_QtRemoteObjects_PackedList) {
	if signal := qt.GetSignal(ptr, "saveProperties"); signal != nil {
		signal.(func(string, *core.QByteArray, []*core.QVariant))(cGoUnpackString(repName), core.NewQByteArrayFromPointer(repSig), func(l C.struct_QtRemoteObjects_PackedList) []*core.QVariant {
			out := make([]*core.QVariant, int(l.len))
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__saveProperties_values_atList(i)
			}
			return out
		}(values))
	}

}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectSaveProperties(f func(repName string, repSig *core.QByteArray, values []*core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveProperties"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "saveProperties", func(repName string, repSig *core.QByteArray, values []*core.QVariant) {
				signal.(func(string, *core.QByteArray, []*core.QVariant))(repName, repSig, values)
				f(repName, repSig, values)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveProperties", f)
		}
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectSaveProperties() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveProperties")
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) SaveProperties(repName string, repSig core.QByteArray_ITF, values []*core.QVariant) {
	if ptr.Pointer() != nil {
		var repNameC *C.char
		if repName != "" {
			repNameC = C.CString(repName)
			defer C.free(unsafe.Pointer(repNameC))
		}
		C.QRemoteObjectAbstractPersistedStore_SaveProperties(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: repNameC, len: C.longlong(len(repName))}, core.PointerFromQByteArray(repSig), func() unsafe.Pointer {
			tmpList := NewQRemoteObjectAbstractPersistedStoreFromPointer(NewQRemoteObjectAbstractPersistedStoreFromPointer(nil).__saveProperties_values_newList())
			for _, v := range values {
				tmpList.__saveProperties_values_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStore
func callbackQRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStore(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectAbstractPersistedStore"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).DestroyQRemoteObjectAbstractPersistedStoreDefault()
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectDestroyQRemoteObjectAbstractPersistedStore(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectAbstractPersistedStore"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectAbstractPersistedStore", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectAbstractPersistedStore", f)
		}
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectDestroyQRemoteObjectAbstractPersistedStore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectAbstractPersistedStore")
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DestroyQRemoteObjectAbstractPersistedStore() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStore(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DestroyQRemoteObjectAbstractPersistedStoreDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStoreDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_MetaObject
func callbackQRemoteObjectAbstractPersistedStore_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectAbstractPersistedStore) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectAbstractPersistedStore_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectAbstractPersistedStore___restoreProperties_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___restoreProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __restoreProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___restoreProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectAbstractPersistedStore___saveProperties_values_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___saveProperties_values_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __saveProperties_values_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___saveProperties_values_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList2() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___findChildren_newList2(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___findChildren_newList3(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectAbstractPersistedStore___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectAbstractPersistedStore___children_newList(ptr.Pointer())
}

//export callbackQRemoteObjectAbstractPersistedStore_Event
func callbackQRemoteObjectAbstractPersistedStore_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectAbstractPersistedStore_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectAbstractPersistedStore_EventFilter
func callbackQRemoteObjectAbstractPersistedStore_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectAbstractPersistedStore) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectAbstractPersistedStore_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectAbstractPersistedStore_ChildEvent
func callbackQRemoteObjectAbstractPersistedStore_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_ConnectNotify
func callbackQRemoteObjectAbstractPersistedStore_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_CustomEvent
func callbackQRemoteObjectAbstractPersistedStore_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_DeleteLater
func callbackQRemoteObjectAbstractPersistedStore_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_Destroyed
func callbackQRemoteObjectAbstractPersistedStore_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify
func callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged
func callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQRemoteObjectAbstractPersistedStore_TimerEvent
func callbackQRemoteObjectAbstractPersistedStore_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRemoteObjectAbstractPersistedStoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectAbstractPersistedStore) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectAbstractPersistedStore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QRemoteObjectDynamicReplica struct {
	QRemoteObjectReplica
}

type QRemoteObjectDynamicReplica_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica
}

func (ptr *QRemoteObjectDynamicReplica) QRemoteObjectDynamicReplica_PTR() *QRemoteObjectDynamicReplica {
	return ptr
}

func (ptr *QRemoteObjectDynamicReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectDynamicReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectDynamicReplica(ptr QRemoteObjectDynamicReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectDynamicReplica_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectDynamicReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectDynamicReplica) {
	n = new(QRemoteObjectDynamicReplica)
	n.SetPointer(ptr)
	return
}

//export callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica
func callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectDynamicReplica"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectDynamicReplicaFromPointer(ptr).DestroyQRemoteObjectDynamicReplicaDefault()
	}
}

func (ptr *QRemoteObjectDynamicReplica) ConnectDestroyQRemoteObjectDynamicReplica(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica", f)
		}
	}
}

func (ptr *QRemoteObjectDynamicReplica) DisconnectDestroyQRemoteObjectDynamicReplica() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectDynamicReplica")
	}
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplica() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectDynamicReplica) DestroyQRemoteObjectDynamicReplicaDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplicaDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QRemoteObjectHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectHost_PTR() *QRemoteObjectHost
}

func (ptr *QRemoteObjectHost) QRemoteObjectHost_PTR() *QRemoteObjectHost {
	return ptr
}

func (ptr *QRemoteObjectHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHost(ptr QRemoteObjectHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHost_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHost) {
	n = new(QRemoteObjectHost)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectHost(parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectHost3(address core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost3(core.PointerFromQUrl(address), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectHost2(address core.QUrl_ITF, registryAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas, parent core.QObject_ITF) *QRemoteObjectHost {
	tmpValue := NewQRemoteObjectHostFromPointer(C.QRemoteObjectHost_NewQRemoteObjectHost2(core.PointerFromQUrl(address), core.PointerFromQUrl(registryAddress), C.longlong(allowedSchemas), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQRemoteObjectHost_SetHostUrl
func callbackQRemoteObjectHost_SetHostUrl(ptr unsafe.Pointer, hostAddress unsafe.Pointer, allowedSchemas C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "setHostUrl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl, QRemoteObjectHostBase__AllowedSchemas) bool)(core.NewQUrlFromPointer(hostAddress), QRemoteObjectHostBase__AllowedSchemas(allowedSchemas)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectHostFromPointer(ptr).SetHostUrlDefault(core.NewQUrlFromPointer(hostAddress), QRemoteObjectHostBase__AllowedSchemas(allowedSchemas)))))
}

func (ptr *QRemoteObjectHost) ConnectSetHostUrl(f func(hostAddress *core.QUrl, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHostUrl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setHostUrl", func(hostAddress *core.QUrl, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
				signal.(func(*core.QUrl, QRemoteObjectHostBase__AllowedSchemas) bool)(hostAddress, allowedSchemas)
				return f(hostAddress, allowedSchemas)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHostUrl", f)
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectSetHostUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHostUrl")
	}
}

func (ptr *QRemoteObjectHost) SetHostUrl(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHost_SetHostUrl(ptr.Pointer(), core.PointerFromQUrl(hostAddress), C.longlong(allowedSchemas))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHost) SetHostUrlDefault(hostAddress core.QUrl_ITF, allowedSchemas QRemoteObjectHostBase__AllowedSchemas) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHost_SetHostUrlDefault(ptr.Pointer(), core.PointerFromQUrl(hostAddress), C.longlong(allowedSchemas))) != 0
	}
	return false
}

//export callbackQRemoteObjectHost_DestroyQRemoteObjectHost
func callbackQRemoteObjectHost_DestroyQRemoteObjectHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectHost"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectHostFromPointer(ptr).DestroyQRemoteObjectHostDefault()
	}
}

func (ptr *QRemoteObjectHost) ConnectDestroyQRemoteObjectHost(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectHost"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectHost", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectHost", f)
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectDestroyQRemoteObjectHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectHost")
	}
}

func (ptr *QRemoteObjectHost) DestroyQRemoteObjectHost() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHost_DestroyQRemoteObjectHost(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectHost) DestroyQRemoteObjectHostDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHost_DestroyQRemoteObjectHostDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRemoteObjectHost_HostUrl
func callbackQRemoteObjectHost_HostUrl(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "hostUrl"); signal != nil {
		return core.PointerFromQUrl(signal.(func() *core.QUrl)())
	}

	return core.PointerFromQUrl(NewQRemoteObjectHostFromPointer(ptr).HostUrlDefault())
}

func (ptr *QRemoteObjectHost) ConnectHostUrl(f func() *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hostUrl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hostUrl", func() *core.QUrl {
				signal.(func() *core.QUrl)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostUrl", f)
		}
	}
}

func (ptr *QRemoteObjectHost) DisconnectHostUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hostUrl")
	}
}

func (ptr *QRemoteObjectHost) HostUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectHost_HostUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectHost) HostUrlDefault() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectHost_HostUrlDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

type QRemoteObjectHostBase struct {
	QRemoteObjectNode
}

type QRemoteObjectHostBase_ITF interface {
	QRemoteObjectNode_ITF
	QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase
}

func (ptr *QRemoteObjectHostBase) QRemoteObjectHostBase_PTR() *QRemoteObjectHostBase {
	return ptr
}

func (ptr *QRemoteObjectHostBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectHostBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectNode_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectHostBase(ptr QRemoteObjectHostBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectHostBaseFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectHostBase) {
	n = new(QRemoteObjectHostBase)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectHostBase__AllowedSchemas
//QRemoteObjectHostBase::AllowedSchemas
type QRemoteObjectHostBase__AllowedSchemas int64

const (
	QRemoteObjectHostBase__BuiltInSchemasOnly        QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(0)
	QRemoteObjectHostBase__AllowExternalRegistration QRemoteObjectHostBase__AllowedSchemas = QRemoteObjectHostBase__AllowedSchemas(1)
)

func (ptr *QRemoteObjectHostBase) DisableRemoting(remoteObject core.QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectHostBase_DisableRemoting(ptr.Pointer(), core.PointerFromQObject(remoteObject))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) EnableRemoting3(model core.QAbstractItemModel_ITF, name string, roles []int, selectionModel core.QItemSelectionModel_ITF) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QRemoteObjectHostBase_EnableRemoting3(ptr.Pointer(), core.PointerFromQAbstractItemModel(model), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))}, func() unsafe.Pointer {
			tmpList := NewQRemoteObjectHostBaseFromPointer(NewQRemoteObjectHostBaseFromPointer(nil).__enableRemoting_roles_newList3())
			for _, v := range roles {
				tmpList.__enableRemoting_roles_setList3(v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQItemSelectionModel(selectionModel))) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) EnableRemoting2(object core.QObject_ITF, name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QRemoteObjectHostBase_EnableRemoting2(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QRemoteObjectHostBase) AddHostSideConnection(ioDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase_AddHostSideConnection(ptr.Pointer(), core.PointerFromQIODevice(ioDevice))
	}
}

//export callbackQRemoteObjectHostBase_DestroyQRemoteObjectHostBase
func callbackQRemoteObjectHostBase_DestroyQRemoteObjectHostBase(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectHostBase"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectHostBaseFromPointer(ptr).DestroyQRemoteObjectHostBaseDefault()
	}
}

func (ptr *QRemoteObjectHostBase) ConnectDestroyQRemoteObjectHostBase(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectHostBase"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectHostBase", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectHostBase", f)
		}
	}
}

func (ptr *QRemoteObjectHostBase) DisconnectDestroyQRemoteObjectHostBase() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectHostBase")
	}
}

func (ptr *QRemoteObjectHostBase) DestroyQRemoteObjectHostBase() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase_DestroyQRemoteObjectHostBase(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectHostBase) DestroyQRemoteObjectHostBaseDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase_DestroyQRemoteObjectHostBaseDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_atList3(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectHostBase___enableRemoting_roles_atList3(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_setList3(i int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectHostBase___enableRemoting_roles_setList3(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QRemoteObjectHostBase) __enableRemoting_roles_newList3() unsafe.Pointer {
	return C.QRemoteObjectHostBase___enableRemoting_roles_newList3(ptr.Pointer())
}

type QRemoteObjectNode struct {
	core.QObject
}

type QRemoteObjectNode_ITF interface {
	core.QObject_ITF
	QRemoteObjectNode_PTR() *QRemoteObjectNode
}

func (ptr *QRemoteObjectNode) QRemoteObjectNode_PTR() *QRemoteObjectNode {
	return ptr
}

func (ptr *QRemoteObjectNode) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectNode) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectNode(ptr QRemoteObjectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectNode_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectNodeFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectNode) {
	n = new(QRemoteObjectNode)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectNode__ErrorCode
//QRemoteObjectNode::ErrorCode
type QRemoteObjectNode__ErrorCode int64

const (
	QRemoteObjectNode__NoError                       QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(0)
	QRemoteObjectNode__RegistryNotAcquired           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(1)
	QRemoteObjectNode__RegistryAlreadyHosted         QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(2)
	QRemoteObjectNode__NodeIsNoServer                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(3)
	QRemoteObjectNode__ServerAlreadyCreated          QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(4)
	QRemoteObjectNode__UnintendedRegistryHosting     QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(5)
	QRemoteObjectNode__OperationNotValidOnClientNode QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(6)
	QRemoteObjectNode__SourceNotRegistered           QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(7)
	QRemoteObjectNode__MissingObjectName             QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(8)
	QRemoteObjectNode__HostUrlInvalid                QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(9)
	QRemoteObjectNode__ProtocolMismatch              QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(10)
	QRemoteObjectNode__ListenFailed                  QRemoteObjectNode__ErrorCode = QRemoteObjectNode__ErrorCode(11)
)

func (ptr *QRemoteObjectNode) AcquireDynamic(name string) *QRemoteObjectDynamicReplica {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQRemoteObjectDynamicReplicaFromPointer(C.QRemoteObjectNode_AcquireDynamic(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQRemoteObjectNode(parent core.QObject_ITF) *QRemoteObjectNode {
	tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectNode_NewQRemoteObjectNode(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRemoteObjectNode2(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectNode {
	tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectNode_NewQRemoteObjectNode2(core.PointerFromQUrl(registryAddress), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QRemoteObjectNode_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectNode_QRemoteObjectNode_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectNode) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectNode_QRemoteObjectNode_Tr(sC, cC, C.int(int32(n))))
}

func QRemoteObjectNode_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectNode_QRemoteObjectNode_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectNode) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectNode_QRemoteObjectNode_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectNode) ConnectToNode(address core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_ConnectToNode(ptr.Pointer(), core.PointerFromQUrl(address))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_SetRegistryUrl
func callbackQRemoteObjectNode_SetRegistryUrl(ptr unsafe.Pointer, registryAddress unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setRegistryUrl"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl) bool)(core.NewQUrlFromPointer(registryAddress)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).SetRegistryUrlDefault(core.NewQUrlFromPointer(registryAddress)))))
}

func (ptr *QRemoteObjectNode) ConnectSetRegistryUrl(f func(registryAddress *core.QUrl) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRegistryUrl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRegistryUrl", func(registryAddress *core.QUrl) bool {
				signal.(func(*core.QUrl) bool)(registryAddress)
				return f(registryAddress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRegistryUrl", f)
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectSetRegistryUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRegistryUrl")
	}
}

func (ptr *QRemoteObjectNode) SetRegistryUrl(registryAddress core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_SetRegistryUrl(ptr.Pointer(), core.PointerFromQUrl(registryAddress))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) SetRegistryUrlDefault(registryAddress core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_SetRegistryUrlDefault(ptr.Pointer(), core.PointerFromQUrl(registryAddress))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) WaitForRegistry(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_WaitForRegistry(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

func (ptr *QRemoteObjectNode) AddClientSideConnection(ioDevice core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_AddClientSideConnection(ptr.Pointer(), core.PointerFromQIODevice(ioDevice))
	}
}

//export callbackQRemoteObjectNode_Error
func callbackQRemoteObjectNode_Error(ptr unsafe.Pointer, errorCode C.longlong) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		signal.(func(QRemoteObjectNode__ErrorCode))(QRemoteObjectNode__ErrorCode(errorCode))
	}

}

func (ptr *QRemoteObjectNode) ConnectError(f func(errorCode QRemoteObjectNode__ErrorCode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QRemoteObjectNode_ConnectError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func(errorCode QRemoteObjectNode__ErrorCode) {
				signal.(func(QRemoteObjectNode__ErrorCode))(errorCode)
				f(errorCode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QRemoteObjectNode) Error(errorCode QRemoteObjectNode__ErrorCode) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_Error(ptr.Pointer(), C.longlong(errorCode))
	}
}

//export callbackQRemoteObjectNode_HeartbeatIntervalChanged
func callbackQRemoteObjectNode_HeartbeatIntervalChanged(ptr unsafe.Pointer, heartbeatInterval C.int) {
	if signal := qt.GetSignal(ptr, "heartbeatIntervalChanged"); signal != nil {
		signal.(func(int))(int(int32(heartbeatInterval)))
	}

}

func (ptr *QRemoteObjectNode) ConnectHeartbeatIntervalChanged(f func(heartbeatInterval int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "heartbeatIntervalChanged") {
			C.QRemoteObjectNode_ConnectHeartbeatIntervalChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "heartbeatIntervalChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "heartbeatIntervalChanged", func(heartbeatInterval int) {
				signal.(func(int))(heartbeatInterval)
				f(heartbeatInterval)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "heartbeatIntervalChanged", f)
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectHeartbeatIntervalChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DisconnectHeartbeatIntervalChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "heartbeatIntervalChanged")
	}
}

func (ptr *QRemoteObjectNode) HeartbeatIntervalChanged(heartbeatInterval int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_HeartbeatIntervalChanged(ptr.Pointer(), C.int(int32(heartbeatInterval)))
	}
}

func (ptr *QRemoteObjectNode) SetHeartbeatInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_SetHeartbeatInterval(ptr.Pointer(), C.int(int32(interval)))
	}
}

//export callbackQRemoteObjectNode_SetName
func callbackQRemoteObjectNode_SetName(ptr unsafe.Pointer, name C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "setName"); signal != nil {
		signal.(func(string))(cGoUnpackString(name))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).SetNameDefault(cGoUnpackString(name))
	}
}

func (ptr *QRemoteObjectNode) ConnectSetName(f func(name string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setName", func(name string) {
				signal.(func(string))(name)
				f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setName", f)
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectSetName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setName")
	}
}

func (ptr *QRemoteObjectNode) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QRemoteObjectNode_SetName(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QRemoteObjectNode) SetNameDefault(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QRemoteObjectNode_SetNameDefault(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QRemoteObjectNode) SetPersistedStore(persistedStore QRemoteObjectAbstractPersistedStore_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_SetPersistedStore(ptr.Pointer(), PointerFromQRemoteObjectAbstractPersistedStore(persistedStore))
	}
}

//export callbackQRemoteObjectNode_TimerEvent
func callbackQRemoteObjectNode_TimerEvent(ptr unsafe.Pointer, vqt unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(vqt))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(vqt))
	}
}

func (ptr *QRemoteObjectNode) TimerEventDefault(vqt core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(vqt))
	}
}

//export callbackQRemoteObjectNode_DestroyQRemoteObjectNode
func callbackQRemoteObjectNode_DestroyQRemoteObjectNode(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectNode"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).DestroyQRemoteObjectNodeDefault()
	}
}

func (ptr *QRemoteObjectNode) ConnectDestroyQRemoteObjectNode(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectNode", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectNode", f)
		}
	}
}

func (ptr *QRemoteObjectNode) DisconnectDestroyQRemoteObjectNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectNode")
	}
}

func (ptr *QRemoteObjectNode) DestroyQRemoteObjectNode() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DestroyQRemoteObjectNode(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectNode) DestroyQRemoteObjectNodeDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DestroyQRemoteObjectNodeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectNode) PersistedStore() *QRemoteObjectAbstractPersistedStore {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectAbstractPersistedStoreFromPointer(C.QRemoteObjectNode_PersistedStore(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) LastError() QRemoteObjectNode__ErrorCode {
	if ptr.Pointer() != nil {
		return QRemoteObjectNode__ErrorCode(C.QRemoteObjectNode_LastError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRemoteObjectNode) Instances2(typeName string) []string {
	if ptr.Pointer() != nil {
		var typeNameC *C.char
		if typeName != "" {
			typeNameC = C.CString(typeName)
			defer C.free(unsafe.Pointer(typeNameC))
		}
		return strings.Split(cGoUnpackString(C.QRemoteObjectNode_Instances2(ptr.Pointer(), C.struct_QtRemoteObjects_PackedString{data: typeNameC, len: C.longlong(len(typeName))})), "|")
	}
	return make([]string, 0)
}

func (ptr *QRemoteObjectNode) RegistryUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QRemoteObjectNode_RegistryUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQRemoteObjectNode_MetaObject
func callbackQRemoteObjectNode_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectNodeFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectNode) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectNode_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRemoteObjectNode) Registry() *QRemoteObjectRegistry {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectRegistryFromPointer(C.QRemoteObjectNode_Registry(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) HeartbeatInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectNode_HeartbeatInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRemoteObjectNode___acquireModel_rolesHint_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_setList(i int) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___acquireModel_rolesHint_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QRemoteObjectNode) __acquireModel_rolesHint_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___acquireModel_rolesHint_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __retrieveProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectNode___retrieveProperties_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __retrieveProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___retrieveProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectNode) __retrieveProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___retrieveProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __persistProperties_props_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectNode___persistProperties_props_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __persistProperties_props_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___persistProperties_props_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectNode) __persistProperties_props_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___persistProperties_props_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectNode___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectNode) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __findChildren_newList2() unsafe.Pointer {
	return C.QRemoteObjectNode___findChildren_newList2(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectNode___findChildren_newList3(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectNode) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectNode___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectNode) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectNode) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectNode___children_newList(ptr.Pointer())
}

//export callbackQRemoteObjectNode_Event
func callbackQRemoteObjectNode_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectNode) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_EventFilter
func callbackQRemoteObjectNode_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectNodeFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectNode) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectNode_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectNode_ChildEvent
func callbackQRemoteObjectNode_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectNode) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectNode_ConnectNotify
func callbackQRemoteObjectNode_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectNode) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectNode_CustomEvent
func callbackQRemoteObjectNode_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectNode) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectNode_DeleteLater
func callbackQRemoteObjectNode_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectNode) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRemoteObjectNode_Destroyed
func callbackQRemoteObjectNode_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectNode_DisconnectNotify
func callbackQRemoteObjectNode_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectNodeFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectNode) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectNode_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectNode_ObjectNameChanged
func callbackQRemoteObjectNode_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

type QRemoteObjectPackets struct {
	ptr unsafe.Pointer
}

type QRemoteObjectPackets_ITF interface {
	QRemoteObjectPackets_PTR() *QRemoteObjectPackets
}

func (ptr *QRemoteObjectPackets) QRemoteObjectPackets_PTR() *QRemoteObjectPackets {
	return ptr
}

func (ptr *QRemoteObjectPackets) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRemoteObjectPackets) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRemoteObjectPackets(ptr QRemoteObjectPackets_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectPackets_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectPacketsFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectPackets) {
	n = new(QRemoteObjectPackets)
	n.SetPointer(ptr)
	return
}

func (ptr *QRemoteObjectPackets) DestroyQRemoteObjectPackets() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QRemoteObjectPackets__ObjectType
//QRemoteObjectPackets::ObjectType
type QRemoteObjectPackets__ObjectType int64

const (
	QRemoteObjectPackets__CLASS QRemoteObjectPackets__ObjectType = QRemoteObjectPackets__ObjectType(0)
	QRemoteObjectPackets__MODEL QRemoteObjectPackets__ObjectType = QRemoteObjectPackets__ObjectType(1)
)

type QRemoteObjectRegistry struct {
	QRemoteObjectReplica
}

type QRemoteObjectRegistry_ITF interface {
	QRemoteObjectReplica_ITF
	QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry
}

func (ptr *QRemoteObjectRegistry) QRemoteObjectRegistry_PTR() *QRemoteObjectRegistry {
	return ptr
}

func (ptr *QRemoteObjectRegistry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectReplica_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistry(ptr QRemoteObjectRegistry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistry_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectRegistryFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistry) {
	n = new(QRemoteObjectRegistry)
	n.SetPointer(ptr)
	return
}
func QRemoteObjectRegistry_RegisterMetatypes() {
	C.QRemoteObjectRegistry_QRemoteObjectRegistry_RegisterMetatypes()
}

func (ptr *QRemoteObjectRegistry) RegisterMetatypes() {
	C.QRemoteObjectRegistry_QRemoteObjectRegistry_RegisterMetatypes()
}

//export callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry
func callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectRegistry"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectRegistryFromPointer(ptr).DestroyQRemoteObjectRegistryDefault()
	}
}

func (ptr *QRemoteObjectRegistry) ConnectDestroyQRemoteObjectRegistry(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectRegistry"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry", f)
		}
	}
}

func (ptr *QRemoteObjectRegistry) DisconnectDestroyQRemoteObjectRegistry() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectRegistry")
	}
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistry() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectRegistry) DestroyQRemoteObjectRegistryDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectRegistry_DestroyQRemoteObjectRegistryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QRemoteObjectRegistryHost struct {
	QRemoteObjectHostBase
}

type QRemoteObjectRegistryHost_ITF interface {
	QRemoteObjectHostBase_ITF
	QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost
}

func (ptr *QRemoteObjectRegistryHost) QRemoteObjectRegistryHost_PTR() *QRemoteObjectRegistryHost {
	return ptr
}

func (ptr *QRemoteObjectRegistryHost) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectHostBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectRegistryHost) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QRemoteObjectHostBase_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectRegistryHost(ptr QRemoteObjectRegistryHost_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectRegistryHost_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectRegistryHostFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectRegistryHost) {
	n = new(QRemoteObjectRegistryHost)
	n.SetPointer(ptr)
	return
}
func NewQRemoteObjectRegistryHost(registryAddress core.QUrl_ITF, parent core.QObject_ITF) *QRemoteObjectRegistryHost {
	tmpValue := NewQRemoteObjectRegistryHostFromPointer(C.QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(core.PointerFromQUrl(registryAddress), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost
func callbackQRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QRemoteObjectRegistryHost"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectRegistryHostFromPointer(ptr).DestroyQRemoteObjectRegistryHostDefault()
	}
}

func (ptr *QRemoteObjectRegistryHost) ConnectDestroyQRemoteObjectRegistryHost(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QRemoteObjectRegistryHost"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistryHost", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QRemoteObjectRegistryHost", f)
		}
	}
}

func (ptr *QRemoteObjectRegistryHost) DisconnectDestroyQRemoteObjectRegistryHost() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QRemoteObjectRegistryHost")
	}
}

func (ptr *QRemoteObjectRegistryHost) DestroyQRemoteObjectRegistryHost() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QRemoteObjectRegistryHost) DestroyQRemoteObjectRegistryHostDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHostDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QRemoteObjectReplica struct {
	core.QObject
}

type QRemoteObjectReplica_ITF interface {
	core.QObject_ITF
	QRemoteObjectReplica_PTR() *QRemoteObjectReplica
}

func (ptr *QRemoteObjectReplica) QRemoteObjectReplica_PTR() *QRemoteObjectReplica {
	return ptr
}

func (ptr *QRemoteObjectReplica) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QRemoteObjectReplica) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQRemoteObjectReplica(ptr QRemoteObjectReplica_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectReplica_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectReplicaFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectReplica) {
	n = new(QRemoteObjectReplica)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QRemoteObjectReplica__State
//QRemoteObjectReplica::State
type QRemoteObjectReplica__State int64

const (
	QRemoteObjectReplica__Uninitialized     QRemoteObjectReplica__State = QRemoteObjectReplica__State(0)
	QRemoteObjectReplica__Default           QRemoteObjectReplica__State = QRemoteObjectReplica__State(1)
	QRemoteObjectReplica__Valid             QRemoteObjectReplica__State = QRemoteObjectReplica__State(2)
	QRemoteObjectReplica__Suspect           QRemoteObjectReplica__State = QRemoteObjectReplica__State(3)
	QRemoteObjectReplica__SignatureMismatch QRemoteObjectReplica__State = QRemoteObjectReplica__State(4)
)

func QRemoteObjectReplica_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectReplica_QRemoteObjectReplica_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectReplica) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectReplica_QRemoteObjectReplica_Tr(sC, cC, C.int(int32(n))))
}

func QRemoteObjectReplica_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectReplica_QRemoteObjectReplica_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectReplica) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QRemoteObjectReplica_QRemoteObjectReplica_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QRemoteObjectReplica) WaitForSource(timeout int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_WaitForSource(ptr.Pointer(), C.int(int32(timeout)))) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_Initialized
func callbackQRemoteObjectReplica_Initialized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialized"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QRemoteObjectReplica) ConnectInitialized(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "initialized") {
			C.QRemoteObjectReplica_ConnectInitialized(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "initialized"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initialized", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialized", f)
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectInitialized() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "initialized")
	}
}

func (ptr *QRemoteObjectReplica) Initialized() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_Initialized(ptr.Pointer())
	}
}

//export callbackQRemoteObjectReplica_SetNode
func callbackQRemoteObjectReplica_SetNode(ptr unsafe.Pointer, node unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setNode"); signal != nil {
		signal.(func(*QRemoteObjectNode))(NewQRemoteObjectNodeFromPointer(node))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).SetNodeDefault(NewQRemoteObjectNodeFromPointer(node))
	}
}

func (ptr *QRemoteObjectReplica) ConnectSetNode(f func(node *QRemoteObjectNode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setNode", func(node *QRemoteObjectNode) {
				signal.(func(*QRemoteObjectNode))(node)
				f(node)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNode", f)
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectSetNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNode")
	}
}

func (ptr *QRemoteObjectReplica) SetNode(node QRemoteObjectNode_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_SetNode(ptr.Pointer(), PointerFromQRemoteObjectNode(node))
	}
}

func (ptr *QRemoteObjectReplica) SetNodeDefault(node QRemoteObjectNode_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_SetNodeDefault(ptr.Pointer(), PointerFromQRemoteObjectNode(node))
	}
}

//export callbackQRemoteObjectReplica_StateChanged
func callbackQRemoteObjectReplica_StateChanged(ptr unsafe.Pointer, state C.longlong, oldState C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QRemoteObjectReplica__State, QRemoteObjectReplica__State))(QRemoteObjectReplica__State(state), QRemoteObjectReplica__State(oldState))
	}

}

func (ptr *QRemoteObjectReplica) ConnectStateChanged(f func(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QRemoteObjectReplica_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State) {
				signal.(func(QRemoteObjectReplica__State, QRemoteObjectReplica__State))(state, oldState)
				f(state, oldState)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QRemoteObjectReplica) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QRemoteObjectReplica) StateChanged(state QRemoteObjectReplica__State, oldState QRemoteObjectReplica__State) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_StateChanged(ptr.Pointer(), C.longlong(state), C.longlong(oldState))
	}
}

func (ptr *QRemoteObjectReplica) Node() *QRemoteObjectNode {
	if ptr.Pointer() != nil {
		tmpValue := NewQRemoteObjectNodeFromPointer(C.QRemoteObjectReplica_Node(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) State() QRemoteObjectReplica__State {
	if ptr.Pointer() != nil {
		return QRemoteObjectReplica__State(C.QRemoteObjectReplica_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRemoteObjectReplica) IsInitialized() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_IsInitialized(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRemoteObjectReplica) IsReplicaValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_IsReplicaValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_MetaObject
func callbackQRemoteObjectReplica_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQRemoteObjectReplicaFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QRemoteObjectReplica) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QRemoteObjectReplica_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___sendWithReply_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___sendWithReply_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __sendWithReply_args_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___sendWithReply_args_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __send_args_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___send_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __send_args_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___send_args_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __send_args_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___send_args_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___setProperties_properties_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___setProperties_properties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __setProperties_properties_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___setProperties_properties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___retrieveProperties_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___retrieveProperties_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __retrieveProperties_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___retrieveProperties_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_atList(i int) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QRemoteObjectReplica___persistProperties_props_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_setList(i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___persistProperties_props_setList(ptr.Pointer(), core.PointerFromQVariant(i))
	}
}

func (ptr *QRemoteObjectReplica) __persistProperties_props_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___persistProperties_props_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QRemoteObjectReplica___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QRemoteObjectReplica) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __findChildren_newList2() unsafe.Pointer {
	return C.QRemoteObjectReplica___findChildren_newList2(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __findChildren_newList3() unsafe.Pointer {
	return C.QRemoteObjectReplica___findChildren_newList3(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __findChildren_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___findChildren_newList(ptr.Pointer())
}

func (ptr *QRemoteObjectReplica) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QRemoteObjectReplica___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRemoteObjectReplica) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QRemoteObjectReplica) __children_newList() unsafe.Pointer {
	return C.QRemoteObjectReplica___children_newList(ptr.Pointer())
}

//export callbackQRemoteObjectReplica_Event
func callbackQRemoteObjectReplica_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectReplicaFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QRemoteObjectReplica) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_EventFilter
func callbackQRemoteObjectReplica_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRemoteObjectReplicaFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QRemoteObjectReplica) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRemoteObjectReplica_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQRemoteObjectReplica_ChildEvent
func callbackQRemoteObjectReplica_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQRemoteObjectReplica_ConnectNotify
func callbackQRemoteObjectReplica_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectReplica) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectReplica_CustomEvent
func callbackQRemoteObjectReplica_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQRemoteObjectReplica_DeleteLater
func callbackQRemoteObjectReplica_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRemoteObjectReplica) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRemoteObjectReplica_Destroyed
func callbackQRemoteObjectReplica_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRemoteObjectReplica_DisconnectNotify
func callbackQRemoteObjectReplica_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRemoteObjectReplica) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRemoteObjectReplica_ObjectNameChanged
func callbackQRemoteObjectReplica_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtRemoteObjects_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQRemoteObjectReplica_TimerEvent
func callbackQRemoteObjectReplica_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQRemoteObjectReplicaFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRemoteObjectReplica) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRemoteObjectReplica_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QRemoteObjectStringLiterals struct {
	ptr unsafe.Pointer
}

type QRemoteObjectStringLiterals_ITF interface {
	QRemoteObjectStringLiterals_PTR() *QRemoteObjectStringLiterals
}

func (ptr *QRemoteObjectStringLiterals) QRemoteObjectStringLiterals_PTR() *QRemoteObjectStringLiterals {
	return ptr
}

func (ptr *QRemoteObjectStringLiterals) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRemoteObjectStringLiterals) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRemoteObjectStringLiterals(ptr QRemoteObjectStringLiterals_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRemoteObjectStringLiterals_PTR().Pointer()
	}
	return nil
}

func NewQRemoteObjectStringLiteralsFromPointer(ptr unsafe.Pointer) (n *QRemoteObjectStringLiterals) {
	n = new(QRemoteObjectStringLiterals)
	n.SetPointer(ptr)
	return
}

func (ptr *QRemoteObjectStringLiterals) DestroyQRemoteObjectStringLiterals() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QRemoteObjectStringLiterals_CLASS() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_CLASS())
}

func (ptr *QRemoteObjectStringLiterals) CLASS() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_CLASS())
}

func QRemoteObjectStringLiterals_MODEL() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_MODEL())
}

func (ptr *QRemoteObjectStringLiterals) MODEL() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_MODEL())
}

func QRemoteObjectStringLiterals_QAIMADAPTER() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_QAIMADAPTER())
}

func (ptr *QRemoteObjectStringLiterals) QAIMADAPTER() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_QAIMADAPTER())
}

func QRemoteObjectStringLiterals_Local() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Local())
}

func (ptr *QRemoteObjectStringLiterals) Local() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Local())
}

func QRemoteObjectStringLiterals_Tcp() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Tcp())
}

func (ptr *QRemoteObjectStringLiterals) Tcp() string {
	return cGoUnpackString(C.QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Tcp())
}

type QtROClientFactory struct {
	ptr unsafe.Pointer
}

type QtROClientFactory_ITF interface {
	QtROClientFactory_PTR() *QtROClientFactory
}

func (ptr *QtROClientFactory) QtROClientFactory_PTR() *QtROClientFactory {
	return ptr
}

func (ptr *QtROClientFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtROClientFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtROClientFactory(ptr QtROClientFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROClientFactory_PTR().Pointer()
	}
	return nil
}

func NewQtROClientFactoryFromPointer(ptr unsafe.Pointer) (n *QtROClientFactory) {
	n = new(QtROClientFactory)
	n.SetPointer(ptr)
	return
}

func (ptr *QtROClientFactory) DestroyQtROClientFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QtROServerFactory struct {
	ptr unsafe.Pointer
}

type QtROServerFactory_ITF interface {
	QtROServerFactory_PTR() *QtROServerFactory
}

func (ptr *QtROServerFactory) QtROServerFactory_PTR() *QtROServerFactory {
	return ptr
}

func (ptr *QtROServerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtROServerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtROServerFactory(ptr QtROServerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtROServerFactory_PTR().Pointer()
	}
	return nil
}

func NewQtROServerFactoryFromPointer(ptr unsafe.Pointer) (n *QtROServerFactory) {
	n = new(QtROServerFactory)
	n.SetPointer(ptr)
	return
}

func (ptr *QtROServerFactory) DestroyQtROServerFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QtRemoteObjects struct {
	ptr unsafe.Pointer
}

type QtRemoteObjects_ITF interface {
	QtRemoteObjects_PTR() *QtRemoteObjects
}

func (ptr *QtRemoteObjects) QtRemoteObjects_PTR() *QtRemoteObjects {
	return ptr
}

func (ptr *QtRemoteObjects) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtRemoteObjects) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtRemoteObjects(ptr QtRemoteObjects_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtRemoteObjects_PTR().Pointer()
	}
	return nil
}

func NewQtRemoteObjectsFromPointer(ptr unsafe.Pointer) (n *QtRemoteObjects) {
	n = new(QtRemoteObjects)
	n.SetPointer(ptr)
	return
}

func (ptr *QtRemoteObjects) DestroyQtRemoteObjects() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QtRemoteObjects__InitialAction
//QtRemoteObjects::InitialAction
type QtRemoteObjects__InitialAction int64

const (
	QtRemoteObjects__FetchRootSize QtRemoteObjects__InitialAction = QtRemoteObjects__InitialAction(0)
	QtRemoteObjects__PrefetchData  QtRemoteObjects__InitialAction = QtRemoteObjects__InitialAction(1)
)

//go:generate stringer -type=QtRemoteObjects__QRemoteObjectPacketTypeEnum
//QtRemoteObjects::QRemoteObjectPacketTypeEnum
type QtRemoteObjects__QRemoteObjectPacketTypeEnum int64

const (
	QtRemoteObjects__Invalid              QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(0)
	QtRemoteObjects__Handshake            QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(1)
	QtRemoteObjects__InitPacket           QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(2)
	QtRemoteObjects__InitDynamicPacket    QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(3)
	QtRemoteObjects__AddObject            QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(4)
	QtRemoteObjects__RemoveObject         QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(5)
	QtRemoteObjects__InvokePacket         QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(6)
	QtRemoteObjects__InvokeReplyPacket    QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(7)
	QtRemoteObjects__PropertyChangePacket QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(8)
	QtRemoteObjects__ObjectList           QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(9)
	QtRemoteObjects__Ping                 QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(10)
	QtRemoteObjects__Pong                 QtRemoteObjects__QRemoteObjectPacketTypeEnum = QtRemoteObjects__QRemoteObjectPacketTypeEnum(11)
)

func QtRemoteObjects_Qt_getEnumName(vqt QtRemoteObjects__InitialAction) string {
	return cGoUnpackString(C.QtRemoteObjects_QtRemoteObjects_Qt_getEnumName(C.longlong(vqt)))
}

func (ptr *QtRemoteObjects) Qt_getEnumName(vqt QtRemoteObjects__InitialAction) string {
	return cGoUnpackString(C.QtRemoteObjects_QtRemoteObjects_Qt_getEnumName(C.longlong(vqt)))
}

func QtRemoteObjects_CopyStoredProperties(mo core.QMetaObject_ITF, src core.QDataStream_ITF, dst unsafe.Pointer) {
	C.QtRemoteObjects_QtRemoteObjects_CopyStoredProperties(core.PointerFromQMetaObject(mo), core.PointerFromQDataStream(src), dst)
}

func (ptr *QtRemoteObjects) CopyStoredProperties(mo core.QMetaObject_ITF, src core.QDataStream_ITF, dst unsafe.Pointer) {
	C.QtRemoteObjects_QtRemoteObjects_CopyStoredProperties(core.PointerFromQMetaObject(mo), core.PointerFromQDataStream(src), dst)
}
