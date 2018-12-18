// +build !minimal

package webchannel

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebChannel_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWebChannel_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QWebChannel struct {
	core.QObject
}

type QWebChannel_ITF interface {
	core.QObject_ITF
	QWebChannel_PTR() *QWebChannel
}

func (ptr *QWebChannel) QWebChannel_PTR() *QWebChannel {
	return ptr
}

func (ptr *QWebChannel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebChannel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebChannel(ptr QWebChannel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannel_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelFromPointer(ptr unsafe.Pointer) (n *QWebChannel) {
	n = new(QWebChannel)
	n.SetPointer(ptr)
	return
}
func QWebChannel_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannel_QWebChannel_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebChannel) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannel_QWebChannel_Tr(sC, cC, C.int(int32(n))))
}

func QWebChannel_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannel_QWebChannel_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QWebChannel) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannel_QWebChannel_TrUtf8(sC, cC, C.int(int32(n))))
}

func NewQWebChannel(parent core.QObject_ITF) *QWebChannel {
	tmpValue := NewQWebChannelFromPointer(C.QWebChannel_NewQWebChannel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebChannel_BlockUpdatesChanged
func callbackQWebChannel_BlockUpdatesChanged(ptr unsafe.Pointer, block C.char) {
	if signal := qt.GetSignal(ptr, "blockUpdatesChanged"); signal != nil {
		signal.(func(bool))(int8(block) != 0)
	}

}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "blockUpdatesChanged") {
			C.QWebChannel_ConnectBlockUpdatesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "blockUpdatesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "blockUpdatesChanged", func(block bool) {
				signal.(func(bool))(block)
				f(block)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "blockUpdatesChanged", f)
		}
	}
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectBlockUpdatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "blockUpdatesChanged")
	}
}

func (ptr *QWebChannel) BlockUpdatesChanged(block bool) {
	if ptr.Pointer() != nil {
		C.QWebChannel_BlockUpdatesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(block))))
	}
}

//export callbackQWebChannel_ConnectTo
func callbackQWebChannel_ConnectTo(ptr unsafe.Pointer, transport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectTo"); signal != nil {
		signal.(func(*QWebChannelAbstractTransport))(NewQWebChannelAbstractTransportFromPointer(transport))
	} else {
		NewQWebChannelFromPointer(ptr).ConnectToDefault(NewQWebChannelAbstractTransportFromPointer(transport))
	}
}

func (ptr *QWebChannel) ConnectConnectTo(f func(transport *QWebChannelAbstractTransport)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connectTo"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connectTo", func(transport *QWebChannelAbstractTransport) {
				signal.(func(*QWebChannelAbstractTransport))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectTo", f)
		}
	}
}

func (ptr *QWebChannel) DisconnectConnectTo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connectTo")
	}
}

func (ptr *QWebChannel) ConnectTo(transport QWebChannelAbstractTransport_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectTo(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) ConnectToDefault(transport QWebChannelAbstractTransport_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectToDefault(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) DeregisterObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DeregisterObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQWebChannel_DisconnectFrom
func callbackQWebChannel_DisconnectFrom(ptr unsafe.Pointer, transport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectFrom"); signal != nil {
		signal.(func(*QWebChannelAbstractTransport))(NewQWebChannelAbstractTransportFromPointer(transport))
	} else {
		NewQWebChannelFromPointer(ptr).DisconnectFromDefault(NewQWebChannelAbstractTransportFromPointer(transport))
	}
}

func (ptr *QWebChannel) ConnectDisconnectFrom(f func(transport *QWebChannelAbstractTransport)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "disconnectFrom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "disconnectFrom", func(transport *QWebChannelAbstractTransport) {
				signal.(func(*QWebChannelAbstractTransport))(transport)
				f(transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnectFrom", f)
		}
	}
}

func (ptr *QWebChannel) DisconnectDisconnectFrom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "disconnectFrom")
	}
}

func (ptr *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransport_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectFrom(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) DisconnectFromDefault(transport QWebChannelAbstractTransport_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectFromDefault(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) RegisterObject(id string, object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.QWebChannel_RegisterObject(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: idC, len: C.longlong(len(id))}, core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) RegisterObjects(objects map[string]*core.QObject) {
	if ptr.Pointer() != nil {
		C.QWebChannel_RegisterObjects(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWebChannelFromPointer(NewQWebChannelFromPointer(nil).__registerObjects_objects_newList())
			for k, v := range objects {
				tmpList.__registerObjects_objects_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QWebChannel) SetBlockUpdates(block bool) {
	if ptr.Pointer() != nil {
		C.QWebChannel_SetBlockUpdates(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(block))))
	}
}

//export callbackQWebChannel_DestroyQWebChannel
func callbackQWebChannel_DestroyQWebChannel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebChannel"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelFromPointer(ptr).DestroyQWebChannelDefault()
	}
}

func (ptr *QWebChannel) ConnectDestroyQWebChannel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebChannel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QWebChannel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebChannel", f)
		}
	}
}

func (ptr *QWebChannel) DisconnectDestroyQWebChannel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebChannel")
	}
}

func (ptr *QWebChannel) DestroyQWebChannel() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebChannel) DestroyQWebChannelDefault() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebChannel) RegisteredObjects() map[string]*core.QObject {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebChannel_PackedList) map[string]*core.QObject {
			out := make(map[string]*core.QObject, int(l.len))
			tmpList := NewQWebChannelFromPointer(l.data)
			for i, v := range tmpList.__registeredObjects_keyList() {
				out[v] = tmpList.__registeredObjects_atList(v, i)
			}
			return out
		}(C.QWebChannel_RegisteredObjects(ptr.Pointer()))
	}
	return make(map[string]*core.QObject, 0)
}

func (ptr *QWebChannel) BlockUpdates() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebChannel_BlockUpdates(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebChannel_MetaObject
func callbackQWebChannel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebChannelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebChannel) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebChannel) __registerObjects_objects_atList(v string, i int) *core.QObject {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___registerObjects_objects_atList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __registerObjects_objects_setList(key string, i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebChannel___registerObjects_objects_setList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __registerObjects_objects_newList() unsafe.Pointer {
	return C.QWebChannel___registerObjects_objects_newList(ptr.Pointer())
}

func (ptr *QWebChannel) __registerObjects_objects_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebChannel_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQWebChannelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____registerObjects_objects_keyList_atList(i)
			}
			return out
		}(C.QWebChannel___registerObjects_objects_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebChannel) __registeredObjects_atList(v string, i int) *core.QObject {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___registeredObjects_atList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __registeredObjects_setList(key string, i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QWebChannel___registeredObjects_setList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __registeredObjects_newList() unsafe.Pointer {
	return C.QWebChannel___registeredObjects_newList(ptr.Pointer())
}

func (ptr *QWebChannel) __registeredObjects_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebChannel_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQWebChannelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____registeredObjects_keyList_atList(i)
			}
			return out
		}(C.QWebChannel___registeredObjects_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebChannel_____registerObjects_objects_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebChannel_____registerObjects_objects_keyList_setList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebChannel) ____registerObjects_objects_keyList_newList() unsafe.Pointer {
	return C.QWebChannel_____registerObjects_objects_keyList_newList(ptr.Pointer())
}

func (ptr *QWebChannel) ____registeredObjects_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebChannel_____registeredObjects_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QWebChannel) ____registeredObjects_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QWebChannel_____registeredObjects_keyList_setList(ptr.Pointer(), C.struct_QtWebChannel_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QWebChannel) ____registeredObjects_keyList_newList() unsafe.Pointer {
	return C.QWebChannel_____registeredObjects_keyList_newList(ptr.Pointer())
}

func (ptr *QWebChannel) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebChannel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebChannel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebChannel___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebChannel) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __findChildren_newList2() unsafe.Pointer {
	return C.QWebChannel___findChildren_newList2(ptr.Pointer())
}

func (ptr *QWebChannel) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __findChildren_newList3() unsafe.Pointer {
	return C.QWebChannel___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebChannel) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __findChildren_newList() unsafe.Pointer {
	return C.QWebChannel___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebChannel) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannel) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannel) __children_newList() unsafe.Pointer {
	return C.QWebChannel___children_newList(ptr.Pointer())
}

//export callbackQWebChannel_Event
func callbackQWebChannel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebChannel) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebChannel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebChannel_EventFilter
func callbackQWebChannel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebChannel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebChannel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebChannel_ChildEvent
func callbackQWebChannel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannel) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebChannel_ConnectNotify
func callbackQWebChannel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannel_CustomEvent
func callbackQWebChannel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannel) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebChannel_DeleteLater
func callbackQWebChannel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebChannel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebChannel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebChannel_Destroyed
func callbackQWebChannel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebChannel_DisconnectNotify
func callbackQWebChannel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannel_ObjectNameChanged
func callbackQWebChannel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebChannel_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebChannel_TimerEvent
func callbackQWebChannel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannel) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransport_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) (n *QWebChannelAbstractTransport) {
	n = new(QWebChannelAbstractTransport)
	n.SetPointer(ptr)
	return
}
func QWebChannelAbstractTransport_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannelAbstractTransport_QWebChannelAbstractTransport_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QWebChannelAbstractTransport) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannelAbstractTransport_QWebChannelAbstractTransport_Tr(sC, cC, C.int(int32(n))))
}

func QWebChannelAbstractTransport_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannelAbstractTransport_QWebChannelAbstractTransport_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QWebChannelAbstractTransport) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QWebChannelAbstractTransport_QWebChannelAbstractTransport_TrUtf8(sC, cC, C.int(int32(n))))
}

func NewQWebChannelAbstractTransport(parent core.QObject_ITF) *QWebChannelAbstractTransport {
	tmpValue := NewQWebChannelAbstractTransportFromPointer(C.QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebChannelAbstractTransport_MessageReceived
func callbackQWebChannelAbstractTransport_MessageReceived(ptr unsafe.Pointer, message unsafe.Pointer, transport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "messageReceived"); signal != nil {
		signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "messageReceived") {
			C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "messageReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "messageReceived", func(message *core.QJsonObject, transport *QWebChannelAbstractTransport) {
				signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(message, transport)
				f(message, transport)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "messageReceived", f)
		}
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "messageReceived")
	}
}

func (ptr *QWebChannelAbstractTransport) MessageReceived(message core.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_MessageReceived(ptr.Pointer(), core.PointerFromQJsonObject(message), PointerFromQWebChannelAbstractTransport(transport))
	}
}

//export callbackQWebChannelAbstractTransport_SendMessage
func callbackQWebChannelAbstractTransport_SendMessage(ptr unsafe.Pointer, message unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sendMessage"); signal != nil {
		signal.(func(*core.QJsonObject))(core.NewQJsonObjectFromPointer(message))
	}

}

func (ptr *QWebChannelAbstractTransport) ConnectSendMessage(f func(message *core.QJsonObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sendMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendMessage", func(message *core.QJsonObject) {
				signal.(func(*core.QJsonObject))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendMessage", f)
		}
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectSendMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sendMessage")
	}
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

//export callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport
func callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebChannelAbstractTransport"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DestroyQWebChannelAbstractTransportDefault()
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectDestroyQWebChannelAbstractTransport(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebChannelAbstractTransport"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QWebChannelAbstractTransport", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebChannelAbstractTransport", f)
		}
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectDestroyQWebChannelAbstractTransport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebChannelAbstractTransport")
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransportDefault() {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransportDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebChannelAbstractTransport_MetaObject
func callbackQWebChannelAbstractTransport_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebChannelAbstractTransportFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebChannelAbstractTransport) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannelAbstractTransport_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebChannelAbstractTransport___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebChannelAbstractTransport) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebChannelAbstractTransport___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebChannelAbstractTransport) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannelAbstractTransport___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannelAbstractTransport) __findChildren_newList2() unsafe.Pointer {
	return C.QWebChannelAbstractTransport___findChildren_newList2(ptr.Pointer())
}

func (ptr *QWebChannelAbstractTransport) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannelAbstractTransport___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannelAbstractTransport) __findChildren_newList3() unsafe.Pointer {
	return C.QWebChannelAbstractTransport___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWebChannelAbstractTransport) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannelAbstractTransport___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannelAbstractTransport) __findChildren_newList() unsafe.Pointer {
	return C.QWebChannelAbstractTransport___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebChannelAbstractTransport) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebChannelAbstractTransport___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebChannelAbstractTransport) __children_newList() unsafe.Pointer {
	return C.QWebChannelAbstractTransport___children_newList(ptr.Pointer())
}

//export callbackQWebChannelAbstractTransport_Event
func callbackQWebChannelAbstractTransport_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelAbstractTransportFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebChannelAbstractTransport) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebChannelAbstractTransport_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebChannelAbstractTransport_EventFilter
func callbackQWebChannelAbstractTransport_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelAbstractTransportFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebChannelAbstractTransport) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebChannelAbstractTransport_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebChannelAbstractTransport_ChildEvent
func callbackQWebChannelAbstractTransport_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebChannelAbstractTransport_ConnectNotify
func callbackQWebChannelAbstractTransport_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannelAbstractTransport_CustomEvent
func callbackQWebChannelAbstractTransport_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebChannelAbstractTransport_DeleteLater
func callbackQWebChannelAbstractTransport_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebChannelAbstractTransport) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebChannelAbstractTransport_Destroyed
func callbackQWebChannelAbstractTransport_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebChannelAbstractTransport_DisconnectNotify
func callbackQWebChannelAbstractTransport_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannelAbstractTransport_ObjectNameChanged
func callbackQWebChannelAbstractTransport_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebChannel_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebChannelAbstractTransport_TimerEvent
func callbackQWebChannelAbstractTransport_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}
