// +build !minimal

package xmlpatterns

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "xmlpatterns.h"
import "C"
import (
	"github.com/StarAurryon/qt"
	"github.com/StarAurryon/qt/core"
	"github.com/StarAurryon/qt/network"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtXmlPatterns_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtXmlPatterns_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractMessageHandler) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandler_PTR().Pointer()
	}
	return nil
}

func NewQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) (n *QAbstractMessageHandler) {
	n = new(QAbstractMessageHandler)
	n.SetPointer(ptr)
	return
}

//export callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler
func callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractMessageHandler"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DestroyQAbstractMessageHandlerDefault()
	}
}

func (ptr *QAbstractMessageHandler) ConnectDestroyQAbstractMessageHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractMessageHandler"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractMessageHandler", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractMessageHandler", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractMessageHandler) DisconnectDestroyQAbstractMessageHandler() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractMessageHandler")
	}
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandlerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractMessageHandler) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractMessageHandler___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractMessageHandler) __children_newList() unsafe.Pointer {
	return C.QAbstractMessageHandler___children_newList(ptr.Pointer())
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractMessageHandler___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractMessageHandler___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractMessageHandler) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractMessageHandler___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractMessageHandler) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractMessageHandler___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractMessageHandler) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractMessageHandler___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractMessageHandler) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractMessageHandler___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractMessageHandler_ChildEvent
func callbackQAbstractMessageHandler_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractMessageHandler_ConnectNotify
func callbackQAbstractMessageHandler_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_CustomEvent
func callbackQAbstractMessageHandler_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractMessageHandler_DeleteLater
func callbackQAbstractMessageHandler_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractMessageHandler) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractMessageHandler_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractMessageHandler_Destroyed
func callbackQAbstractMessageHandler_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractMessageHandler_DisconnectNotify
func callbackQAbstractMessageHandler_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_Event
func callbackQAbstractMessageHandler_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractMessageHandler) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractMessageHandler_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_EventFilter
func callbackQAbstractMessageHandler_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractMessageHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractMessageHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_MetaObject
func callbackQAbstractMessageHandler_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractMessageHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractMessageHandler) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractMessageHandler_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractMessageHandler_ObjectNameChanged
func callbackQAbstractMessageHandler_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractMessageHandler_TimerEvent
func callbackQAbstractMessageHandler_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractMessageHandler) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractUriResolver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) (n *QAbstractUriResolver) {
	n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	return
}
func NewQAbstractUriResolver(parent core.QObject_ITF) *QAbstractUriResolver {
	tmpValue := NewQAbstractUriResolverFromPointer(C.QAbstractUriResolver_NewQAbstractUriResolver(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractUriResolver_Resolve
func callbackQAbstractUriResolver_Resolve(ptr unsafe.Pointer, relative unsafe.Pointer, baseURI unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "resolve"); signal != nil {
		return core.PointerFromQUrl((*(*func(*core.QUrl, *core.QUrl) *core.QUrl)(signal))(core.NewQUrlFromPointer(relative), core.NewQUrlFromPointer(baseURI)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractUriResolver) ConnectResolve(f func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resolve"); signal != nil {
			f := func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl {
				(*(*func(*core.QUrl, *core.QUrl) *core.QUrl)(signal))(relative, baseURI)
				return f(relative, baseURI)
			}
			qt.ConnectSignal(ptr.Pointer(), "resolve", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resolve", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractUriResolver) DisconnectResolve() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resolve")
	}
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QAbstractUriResolver_Resolve(ptr.Pointer(), core.PointerFromQUrl(relative), core.PointerFromQUrl(baseURI)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractUriResolver_DestroyQAbstractUriResolver
func callbackQAbstractUriResolver_DestroyQAbstractUriResolver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractUriResolver"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DestroyQAbstractUriResolverDefault()
	}
}

func (ptr *QAbstractUriResolver) ConnectDestroyQAbstractUriResolver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractUriResolver"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractUriResolver", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractUriResolver", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractUriResolver) DisconnectDestroyQAbstractUriResolver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractUriResolver")
	}
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolverDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractUriResolver_DestroyQAbstractUriResolverDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractUriResolver___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractUriResolver) __children_newList() unsafe.Pointer {
	return C.QAbstractUriResolver___children_newList(ptr.Pointer())
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractUriResolver___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAbstractUriResolver___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAbstractUriResolver) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractUriResolver___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractUriResolver) __findChildren_newList() unsafe.Pointer {
	return C.QAbstractUriResolver___findChildren_newList(ptr.Pointer())
}

func (ptr *QAbstractUriResolver) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractUriResolver___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractUriResolver) __findChildren_newList3() unsafe.Pointer {
	return C.QAbstractUriResolver___findChildren_newList3(ptr.Pointer())
}

//export callbackQAbstractUriResolver_ChildEvent
func callbackQAbstractUriResolver_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractUriResolver_ConnectNotify
func callbackQAbstractUriResolver_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_CustomEvent
func callbackQAbstractUriResolver_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractUriResolver_DeleteLater
func callbackQAbstractUriResolver_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractUriResolver) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractUriResolver_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQAbstractUriResolver_Destroyed
func callbackQAbstractUriResolver_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractUriResolver_DisconnectNotify
func callbackQAbstractUriResolver_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_Event
func callbackQAbstractUriResolver_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractUriResolver) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractUriResolver_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_EventFilter
func callbackQAbstractUriResolver_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractUriResolver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractUriResolver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_MetaObject
func callbackQAbstractUriResolver_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQAbstractUriResolverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractUriResolver) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQAbstractUriResolver_ObjectNameChanged
func callbackQAbstractUriResolver_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractUriResolver_TimerEvent
func callbackQAbstractUriResolver_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return ptr
}

func (ptr *QAbstractXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedData_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QSharedData_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) (n *QAbstractXmlNodeModel) {
	n = new(QAbstractXmlNodeModel)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAbstractXmlNodeModel__SimpleAxis
//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     QAbstractXmlNodeModel__SimpleAxis = QAbstractXmlNodeModel__SimpleAxis(3)
)

//export callbackQAbstractXmlNodeModel_BaseUri
func callbackQAbstractXmlNodeModel_BaseUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "baseUri"); signal != nil {
		return core.PointerFromQUrl((*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectBaseUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "baseUri"); signal != nil {
			f := func(n *QXmlNodeModelIndex) *core.QUrl {
				(*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(n)
				return f(n)
			}
			qt.ConnectSignal(ptr.Pointer(), "baseUri", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectBaseUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "baseUri")
	}
}

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_CompareOrder
func callbackQAbstractXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "compareOrder"); signal != nil {
		return C.longlong((*(*func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(signal))(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "compareOrder"); signal != nil {
			f := func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder {
				(*(*func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(signal))(ni1, ni2)
				return f(ni1, ni2)
			}
			qt.ConnectSignal(ptr.Pointer(), "compareOrder", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "compareOrder", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectCompareOrder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "compareOrder")
	}
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) CreateIndex(data int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex(ptr.Pointer(), C.longlong(data)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex2(pointer unsafe.Pointer, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex2(ptr.Pointer(), pointer, C.longlong(additionalData)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex3(data int64, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex3(ptr.Pointer(), C.longlong(data), C.longlong(additionalData)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DocumentUri
func callbackQAbstractXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "documentUri"); signal != nil {
		return core.PointerFromQUrl((*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "documentUri"); signal != nil {
			f := func(n *QXmlNodeModelIndex) *core.QUrl {
				(*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(n)
				return f(n)
			}
			qt.ConnectSignal(ptr.Pointer(), "documentUri", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "documentUri", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectDocumentUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "documentUri")
	}
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QAbstractXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_ElementById
func callbackQAbstractXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(*QXmlName) *QXmlNodeModelIndex)(signal))(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementById"); signal != nil {
			f := func(id *QXmlName) *QXmlNodeModelIndex {
				(*(*func(*QXmlName) *QXmlNodeModelIndex)(signal))(id)
				return f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "elementById", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementById", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectElementById() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementById")
	}
}

func (ptr *QAbstractXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Kind
func callbackQAbstractXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "kind"); signal != nil {
		return C.longlong((*(*func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(signal))(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "kind"); signal != nil {
			f := func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind {
				(*(*func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(signal))(ni)
				return f(ni)
			}
			qt.ConnectSignal(ptr.Pointer(), "kind", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "kind", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectKind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "kind")
	}
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

//export callbackQAbstractXmlNodeModel_Name
func callbackQAbstractXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		return PointerFromQXmlName((*(*func(*QXmlNodeModelIndex) *QXmlName)(signal))(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQXmlName())
}

func (ptr *QAbstractXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			f := func(ni *QXmlNodeModelIndex) *QXmlName {
				(*(*func(*QXmlNodeModelIndex) *QXmlName)(signal))(ni)
				return f(ni)
			}
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QAbstractXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QAbstractXmlNodeModel_Name(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NamespaceBindings
func callbackQAbstractXmlNodeModel_NamespaceBindings(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "namespaceBindings"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
			for _, v := range (*(*func(*QXmlNodeModelIndex) []*QXmlName)(signal))(NewQXmlNodeModelIndexFromPointer(n)) {
				tmpList.__namespaceBindings_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
		for _, v := range make([]*QXmlName, 0) {
			tmpList.__namespaceBindings_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QAbstractXmlNodeModel) ConnectNamespaceBindings(f func(n *QXmlNodeModelIndex) []*QXmlName) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "namespaceBindings"); signal != nil {
			f := func(n *QXmlNodeModelIndex) []*QXmlName {
				(*(*func(*QXmlNodeModelIndex) []*QXmlName)(signal))(n)
				return f(n)
			}
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNamespaceBindings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "namespaceBindings")
	}
}

func (ptr *QAbstractXmlNodeModel) NamespaceBindings(n QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			out := make([]*QXmlName, int(l.len))
			tmpList := NewQAbstractXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__namespaceBindings_atList(i)
			}
			return out
		}(C.QAbstractXmlNodeModel_NamespaceBindings(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return make([]*QXmlName, 0)
}

//export callbackQAbstractXmlNodeModel_NextFromSimpleAxis
func callbackQAbstractXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextFromSimpleAxis"); signal != nil {
			f := func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex {
				(*(*func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(axis, origin)
				return f(axis, origin)
			}
			qt.ConnectSignal(ptr.Pointer(), "nextFromSimpleAxis", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextFromSimpleAxis", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNextFromSimpleAxis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nextFromSimpleAxis")
	}
}

func (ptr *QAbstractXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NodesByIdref
func callbackQAbstractXmlNodeModel_NodesByIdref(ptr unsafe.Pointer, idref unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nodesByIdref"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__nodesByIdref_newList())
			for _, v := range (*(*func(*QXmlName) []*QXmlNodeModelIndex)(signal))(NewQXmlNameFromPointer(idref)) {
				tmpList.__nodesByIdref_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__nodesByIdref_newList())
		for _, v := range make([]*QXmlNodeModelIndex, 0) {
			tmpList.__nodesByIdref_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QAbstractXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nodesByIdref"); signal != nil {
			f := func(idref *QXmlName) []*QXmlNodeModelIndex {
				(*(*func(*QXmlName) []*QXmlNodeModelIndex)(signal))(idref)
				return f(idref)
			}
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectNodesByIdref() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nodesByIdref")
	}
}

func (ptr *QAbstractXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			out := make([]*QXmlNodeModelIndex, int(l.len))
			tmpList := NewQAbstractXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__nodesByIdref_atList(i)
			}
			return out
		}(C.QAbstractXmlNodeModel_NodesByIdref(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return make([]*QXmlNodeModelIndex, 0)
}

//export callbackQAbstractXmlNodeModel_Root
func callbackQAbstractXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "root"); signal != nil {
			f := func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex {
				(*(*func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(n)
				return f(n)
			}
			qt.ConnectSignal(ptr.Pointer(), "root", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "root", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectRoot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "root")
	}
}

func (ptr *QAbstractXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) SourceLocation(index QXmlNodeModelIndex_ITF) *QSourceLocation {
	if ptr.Pointer() != nil {
		tmpValue := NewQSourceLocationFromPointer(C.QAbstractXmlNodeModel_SourceLocation(ptr.Pointer(), PointerFromQXmlNodeModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_StringValue
func callbackQAbstractXmlNodeModel_StringValue(ptr unsafe.Pointer, n unsafe.Pointer) C.struct_QtXmlPatterns_PackedString {
	if signal := qt.GetSignal(ptr, "stringValue"); signal != nil {
		tempVal := (*(*func(*QXmlNodeModelIndex) string)(signal))(NewQXmlNodeModelIndexFromPointer(n))
		return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QAbstractXmlNodeModel) ConnectStringValue(f func(n *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stringValue"); signal != nil {
			f := func(n *QXmlNodeModelIndex) string {
				(*(*func(*QXmlNodeModelIndex) string)(signal))(n)
				return f(n)
			}
			qt.ConnectSignal(ptr.Pointer(), "stringValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectStringValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stringValue")
	}
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return ""
}

//export callbackQAbstractXmlNodeModel_TypedValue
func callbackQAbstractXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "typedValue"); signal != nil {
		return core.PointerFromQVariant((*(*func(*QXmlNodeModelIndex) *core.QVariant)(signal))(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QAbstractXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "typedValue"); signal != nil {
			f := func(node *QXmlNodeModelIndex) *core.QVariant {
				(*(*func(*QXmlNodeModelIndex) *core.QVariant)(signal))(node)
				return f(node)
			}
			qt.ConnectSignal(ptr.Pointer(), "typedValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typedValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectTypedValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "typedValue")
	}
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAbstractXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel
func callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractXmlNodeModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractXmlNodeModelFromPointer(ptr).DestroyQAbstractXmlNodeModelDefault()
	}
}

func (ptr *QAbstractXmlNodeModel) ConnectDestroyQAbstractXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractXmlNodeModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlNodeModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlNodeModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlNodeModel) DisconnectDestroyQAbstractXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractXmlNodeModel")
	}
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QAbstractXmlNodeModel___namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_setList(i QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel___namespaceBindings_setList(ptr.Pointer(), PointerFromQXmlName(i))
	}
}

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_newList() unsafe.Pointer {
	return C.QAbstractXmlNodeModel___namespaceBindings_newList(ptr.Pointer())
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_atList(i int) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel___nodesByIdref_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_setList(i QXmlNodeModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel___nodesByIdref_setList(ptr.Pointer(), PointerFromQXmlNodeModelIndex(i))
	}
}

func (ptr *QAbstractXmlNodeModel) __nodesByIdref_newList() unsafe.Pointer {
	return C.QAbstractXmlNodeModel___nodesByIdref_newList(ptr.Pointer())
}

type QAbstractXmlReceiver struct {
	ptr unsafe.Pointer
}

type QAbstractXmlReceiver_ITF interface {
	QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver
}

func (ptr *QAbstractXmlReceiver) QAbstractXmlReceiver_PTR() *QAbstractXmlReceiver {
	return ptr
}

func (ptr *QAbstractXmlReceiver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractXmlReceiver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAbstractXmlReceiver(ptr QAbstractXmlReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlReceiverFromPointer(ptr unsafe.Pointer) (n *QAbstractXmlReceiver) {
	n = new(QAbstractXmlReceiver)
	n.SetPointer(ptr)
	return
}
func NewQAbstractXmlReceiver() *QAbstractXmlReceiver {
	return NewQAbstractXmlReceiverFromPointer(C.QAbstractXmlReceiver_NewQAbstractXmlReceiver())
}

//export callbackQAbstractXmlReceiver_AtomicValue
func callbackQAbstractXmlReceiver_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "atomicValue"); signal != nil {
		(*(*func(*core.QVariant))(signal))(core.NewQVariantFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "atomicValue"); signal != nil {
			f := func(value *core.QVariant) {
				(*(*func(*core.QVariant))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAtomicValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "atomicValue")
	}
}

func (ptr *QAbstractXmlReceiver) AtomicValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQAbstractXmlReceiver_Attribute
func callbackQAbstractXmlReceiver_Attribute(ptr unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "attribute"); signal != nil {
		(*(*func(*QXmlName, *core.QStringRef))(signal))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "attribute"); signal != nil {
			f := func(name *QXmlName, value *core.QStringRef) {
				(*(*func(*QXmlName, *core.QStringRef))(signal))(name, value)
				f(name, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "attribute", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "attribute", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "attribute")
	}
}

func (ptr *QAbstractXmlReceiver) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Characters
func callbackQAbstractXmlReceiver_Characters(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characters"); signal != nil {
		(*(*func(*core.QStringRef))(signal))(core.NewQStringRefFromPointer(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "characters"); signal != nil {
			f := func(value *core.QStringRef) {
				(*(*func(*core.QStringRef))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "characters", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characters", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "characters")
	}
}

func (ptr *QAbstractXmlReceiver) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

//export callbackQAbstractXmlReceiver_Comment
func callbackQAbstractXmlReceiver_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			f := func(value string) {
				(*(*func(string))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "comment", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "comment")
	}
}

func (ptr *QAbstractXmlReceiver) Comment(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QAbstractXmlReceiver_Comment(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackQAbstractXmlReceiver_EndDocument
func callbackQAbstractXmlReceiver_EndDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endDocument"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDocument")
	}
}

func (ptr *QAbstractXmlReceiver) EndDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndElement
func callbackQAbstractXmlReceiver_EndElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endElement"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endElement"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endElement", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endElement", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endElement")
	}
}

func (ptr *QAbstractXmlReceiver) EndElement() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndElement(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_EndOfSequence
func callbackQAbstractXmlReceiver_EndOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endOfSequence"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endOfSequence"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectEndOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_EndOfSequence(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_NamespaceBinding
func callbackQAbstractXmlReceiver_NamespaceBinding(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "namespaceBinding"); signal != nil {
		(*(*func(*QXmlName))(signal))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectNamespaceBinding(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "namespaceBinding"); signal != nil {
			f := func(name *QXmlName) {
				(*(*func(*QXmlName))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectNamespaceBinding() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "namespaceBinding")
	}
}

func (ptr *QAbstractXmlReceiver) NamespaceBinding(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_ProcessingInstruction
func callbackQAbstractXmlReceiver_ProcessingInstruction(ptr unsafe.Pointer, target unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "processingInstruction"); signal != nil {
		(*(*func(*QXmlName, string))(signal))(NewQXmlNameFromPointer(target), cGoUnpackString(value))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectProcessingInstruction(f func(target *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processingInstruction"); signal != nil {
			f := func(target *QXmlName, value string) {
				(*(*func(*QXmlName, string))(signal))(target, value)
				f(target, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processingInstruction")
	}
}

func (ptr *QAbstractXmlReceiver) ProcessingInstruction(target QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QAbstractXmlReceiver_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(target), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackQAbstractXmlReceiver_StartDocument
func callbackQAbstractXmlReceiver_StartDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startDocument"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDocument")
	}
}

func (ptr *QAbstractXmlReceiver) StartDocument() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartDocument(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_StartElement
func callbackQAbstractXmlReceiver_StartElement(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startElement"); signal != nil {
		(*(*func(*QXmlName))(signal))(NewQXmlNameFromPointer(name))
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startElement"); signal != nil {
			f := func(name *QXmlName) {
				(*(*func(*QXmlName))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "startElement", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startElement", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startElement")
	}
}

func (ptr *QAbstractXmlReceiver) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQAbstractXmlReceiver_StartOfSequence
func callbackQAbstractXmlReceiver_StartOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startOfSequence"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QAbstractXmlReceiver) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startOfSequence"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectStartOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startOfSequence")
	}
}

func (ptr *QAbstractXmlReceiver) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_StartOfSequence(ptr.Pointer())
	}
}

//export callbackQAbstractXmlReceiver_DestroyQAbstractXmlReceiver
func callbackQAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractXmlReceiver"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAbstractXmlReceiverFromPointer(ptr).DestroyQAbstractXmlReceiverDefault()
	}
}

func (ptr *QAbstractXmlReceiver) ConnectDestroyQAbstractXmlReceiver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractXmlReceiver"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlReceiver", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlReceiver", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAbstractXmlReceiver) DisconnectDestroyQAbstractXmlReceiver() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractXmlReceiver")
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiver() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractXmlReceiver) DestroyQAbstractXmlReceiverDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlReceiver_DestroyQAbstractXmlReceiverDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlNodeModel_PTR().SetPointer(p)
	}
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) (n *QSimpleXmlNodeModel) {
	n = new(QSimpleXmlNodeModel)
	n.SetPointer(ptr)
	return
}

//export callbackQSimpleXmlNodeModel_BaseUri
func callbackQSimpleXmlNodeModel_BaseUri(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "baseUri"); signal != nil {
		return core.PointerFromQUrl((*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).BaseUriDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectBaseUri(f func(node *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "baseUri"); signal != nil {
			f := func(node *QXmlNodeModelIndex) *core.QUrl {
				(*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(node)
				return f(node)
			}
			qt.ConnectSignal(ptr.Pointer(), "baseUri", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectBaseUri() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "baseUri")
	}
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) BaseUriDefault(node QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_ElementById
func callbackQSimpleXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(*QXmlName) *QXmlNodeModelIndex)(signal))(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).ElementByIdDefault(NewQXmlNameFromPointer(id)))
}

func (ptr *QSimpleXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementById"); signal != nil {
			f := func(id *QXmlName) *QXmlNodeModelIndex {
				(*(*func(*QXmlName) *QXmlNodeModelIndex)(signal))(id)
				return f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "elementById", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementById", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectElementById() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementById")
	}
}

func (ptr *QSimpleXmlNodeModel) ElementById(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementById(ptr.Pointer(), PointerFromQXmlName(id)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ElementByIdDefault(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementByIdDefault(ptr.Pointer(), PointerFromQXmlName(id)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QSimpleXmlNodeModel_NamePool(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NamespaceBindings
func callbackQSimpleXmlNodeModel_NamespaceBindings(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "namespaceBindings"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSimpleXmlNodeModelFromPointer(NewQSimpleXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
			for _, v := range (*(*func(*QXmlNodeModelIndex) []*QXmlName)(signal))(NewQXmlNodeModelIndexFromPointer(node)) {
				tmpList.__namespaceBindings_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSimpleXmlNodeModelFromPointer(NewQSimpleXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
		for _, v := range NewQSimpleXmlNodeModelFromPointer(ptr).NamespaceBindingsDefault(NewQXmlNodeModelIndexFromPointer(node)) {
			tmpList.__namespaceBindings_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSimpleXmlNodeModel) ConnectNamespaceBindings(f func(node *QXmlNodeModelIndex) []*QXmlName) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "namespaceBindings"); signal != nil {
			f := func(node *QXmlNodeModelIndex) []*QXmlName {
				(*(*func(*QXmlNodeModelIndex) []*QXmlName)(signal))(node)
				return f(node)
			}
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNamespaceBindings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "namespaceBindings")
	}
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindings(node QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			out := make([]*QXmlName, int(l.len))
			tmpList := NewQSimpleXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__namespaceBindings_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NamespaceBindings(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return make([]*QXmlName, 0)
}

func (ptr *QSimpleXmlNodeModel) NamespaceBindingsDefault(node QXmlNodeModelIndex_ITF) []*QXmlName {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlName {
			out := make([]*QXmlName, int(l.len))
			tmpList := NewQSimpleXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__namespaceBindings_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NamespaceBindingsDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return make([]*QXmlName, 0)
}

//export callbackQSimpleXmlNodeModel_NodesByIdref
func callbackQSimpleXmlNodeModel_NodesByIdref(ptr unsafe.Pointer, idref unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nodesByIdref"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSimpleXmlNodeModelFromPointer(NewQSimpleXmlNodeModelFromPointer(nil).__nodesByIdref_newList())
			for _, v := range (*(*func(*QXmlName) []*QXmlNodeModelIndex)(signal))(NewQXmlNameFromPointer(idref)) {
				tmpList.__nodesByIdref_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQSimpleXmlNodeModelFromPointer(NewQSimpleXmlNodeModelFromPointer(nil).__nodesByIdref_newList())
		for _, v := range NewQSimpleXmlNodeModelFromPointer(ptr).NodesByIdrefDefault(NewQXmlNameFromPointer(idref)) {
			tmpList.__nodesByIdref_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QSimpleXmlNodeModel) ConnectNodesByIdref(f func(idref *QXmlName) []*QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nodesByIdref"); signal != nil {
			f := func(idref *QXmlName) []*QXmlNodeModelIndex {
				(*(*func(*QXmlName) []*QXmlNodeModelIndex)(signal))(idref)
				return f(idref)
			}
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectNodesByIdref() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nodesByIdref")
	}
}

func (ptr *QSimpleXmlNodeModel) NodesByIdref(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			out := make([]*QXmlNodeModelIndex, int(l.len))
			tmpList := NewQSimpleXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__nodesByIdref_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NodesByIdref(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return make([]*QXmlNodeModelIndex, 0)
}

func (ptr *QSimpleXmlNodeModel) NodesByIdrefDefault(idref QXmlName_ITF) []*QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtXmlPatterns_PackedList) []*QXmlNodeModelIndex {
			out := make([]*QXmlNodeModelIndex, int(l.len))
			tmpList := NewQSimpleXmlNodeModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__nodesByIdref_atList(i)
			}
			return out
		}(C.QSimpleXmlNodeModel_NodesByIdrefDefault(ptr.Pointer(), PointerFromQXmlName(idref)))
	}
	return make([]*QXmlNodeModelIndex, 0)
}

//export callbackQSimpleXmlNodeModel_StringValue
func callbackQSimpleXmlNodeModel_StringValue(ptr unsafe.Pointer, node unsafe.Pointer) C.struct_QtXmlPatterns_PackedString {
	if signal := qt.GetSignal(ptr, "stringValue"); signal != nil {
		tempVal := (*(*func(*QXmlNodeModelIndex) string)(signal))(NewQXmlNodeModelIndexFromPointer(node))
		return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSimpleXmlNodeModelFromPointer(ptr).StringValueDefault(NewQXmlNodeModelIndexFromPointer(node))
	return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSimpleXmlNodeModel) ConnectStringValue(f func(node *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stringValue"); signal != nil {
			f := func(node *QXmlNodeModelIndex) string {
				(*(*func(*QXmlNodeModelIndex) string)(signal))(node)
				return f(node)
			}
			qt.ConnectSignal(ptr.Pointer(), "stringValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectStringValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stringValue")
	}
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) StringValueDefault(node QXmlNodeModelIndex_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QSimpleXmlNodeModel_StringValueDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

//export callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel
func callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSimpleXmlNodeModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQSimpleXmlNodeModelFromPointer(ptr).DestroyQSimpleXmlNodeModelDefault()
	}
}

func (ptr *QSimpleXmlNodeModel) ConnectDestroyQSimpleXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSimpleXmlNodeModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QSimpleXmlNodeModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSimpleXmlNodeModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QSimpleXmlNodeModel) DisconnectDestroyQSimpleXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QSimpleXmlNodeModel")
	}
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSimpleXmlNodeModel_CompareOrder
func callbackQSimpleXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "compareOrder"); signal != nil {
		return C.longlong((*(*func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(signal))(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.longlong(NewQSimpleXmlNodeModelFromPointer(ptr).CompareOrderDefault(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
}

func (ptr *QSimpleXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QSimpleXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QSimpleXmlNodeModel) CompareOrderDefault(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QSimpleXmlNodeModel_CompareOrderDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_DocumentUri
func callbackQSimpleXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "documentUri"); signal != nil {
		return core.PointerFromQUrl((*(*func(*QXmlNodeModelIndex) *core.QUrl)(signal))(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).DocumentUriDefault(NewQXmlNodeModelIndexFromPointer(n)))
}

func (ptr *QSimpleXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) DocumentUriDefault(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Kind
func callbackQSimpleXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "kind"); signal != nil {
		return C.longlong((*(*func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(signal))(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.longlong(NewQSimpleXmlNodeModelFromPointer(ptr).KindDefault(NewQXmlNodeModelIndexFromPointer(ni)))
}

func (ptr *QSimpleXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QSimpleXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

func (ptr *QSimpleXmlNodeModel) KindDefault(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QSimpleXmlNodeModel_KindDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

//export callbackQSimpleXmlNodeModel_Name
func callbackQSimpleXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		return PointerFromQXmlName((*(*func(*QXmlNodeModelIndex) *QXmlName)(signal))(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQSimpleXmlNodeModelFromPointer(ptr).NameDefault(NewQXmlNodeModelIndexFromPointer(ni)))
}

func (ptr *QSimpleXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_Name(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NameDefault(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_NameDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NextFromSimpleAxis
func callbackQSimpleXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).NextFromSimpleAxisDefault(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxisDefault(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxisDefault(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Root
func callbackQSimpleXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex((*(*func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(signal))(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).RootDefault(NewQXmlNodeModelIndexFromPointer(n)))
}

func (ptr *QSimpleXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) RootDefault(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_RootDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_TypedValue
func callbackQSimpleXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "typedValue"); signal != nil {
		return core.PointerFromQVariant((*(*func(*QXmlNodeModelIndex) *core.QVariant)(signal))(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(NewQSimpleXmlNodeModelFromPointer(ptr).TypedValueDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSimpleXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) TypedValueDefault(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSimpleXmlNodeModel_TypedValueDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QSourceLocation struct {
	ptr unsafe.Pointer
}

type QSourceLocation_ITF interface {
	QSourceLocation_PTR() *QSourceLocation
}

func (ptr *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return ptr
}

func (ptr *QSourceLocation) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QSourceLocation) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQSourceLocation(ptr QSourceLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocation_PTR().Pointer()
	}
	return nil
}

func NewQSourceLocationFromPointer(ptr unsafe.Pointer) (n *QSourceLocation) {
	n = new(QSourceLocation)
	n.SetPointer(ptr)
	return
}
func NewQSourceLocation() *QSourceLocation {
	tmpValue := NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation())
	qt.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	tmpValue := NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
	qt.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	tmpValue := NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(int32(l)), C.int(int32(c))))
	qt.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func (ptr *QSourceLocation) Column() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSourceLocation_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSourceLocation) Line() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) SetColumn(newColumn int64) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetColumn(ptr.Pointer(), C.longlong(newColumn))
	}
}

func (ptr *QSourceLocation) SetLine(newLine int64) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetLine(ptr.Pointer(), C.longlong(newLine))
	}
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(ptr.Pointer(), core.PointerFromQUrl(newUri))
	}
}

func (ptr *QSourceLocation) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSourceLocation_Uri(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QSourceLocation_DestroyQSourceLocation(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlFormatter struct {
	QXmlSerializer
}

type QXmlFormatter_ITF interface {
	QXmlSerializer_ITF
	QXmlFormatter_PTR() *QXmlFormatter
}

func (ptr *QXmlFormatter) QXmlFormatter_PTR() *QXmlFormatter {
	return ptr
}

func (ptr *QXmlFormatter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlFormatter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QXmlSerializer_PTR().SetPointer(p)
	}
}

func PointerFromQXmlFormatter(ptr QXmlFormatter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlFormatter_PTR().Pointer()
	}
	return nil
}

func NewQXmlFormatterFromPointer(ptr unsafe.Pointer) (n *QXmlFormatter) {
	n = new(QXmlFormatter)
	n.SetPointer(ptr)
	return
}
func (ptr *QXmlFormatter) DestroyQXmlFormatter() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlFormatter_IndentationDepth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

type QXmlItem struct {
	ptr unsafe.Pointer
}

type QXmlItem_ITF interface {
	QXmlItem_PTR() *QXmlItem
}

func (ptr *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return ptr
}

func (ptr *QXmlItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlItem(ptr QXmlItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItem_PTR().Pointer()
	}
	return nil
}

func NewQXmlItemFromPointer(ptr unsafe.Pointer) (n *QXmlItem) {
	n = new(QXmlItem)
	n.SetPointer(ptr)
	return
}
func NewQXmlItem() *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
	qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
	qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
	qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
	qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func (ptr *QXmlItem) IsAtomicValue() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlItem_IsAtomicValue(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNode() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlItem_IsNode(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlItem_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlItem) ToAtomicValue() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QXmlItem_ToAtomicValue(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) ToNodeModelIndex() *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlItem_ToNodeModelIndex(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlName struct {
	ptr unsafe.Pointer
}

type QXmlName_ITF interface {
	QXmlName_PTR() *QXmlName
}

func (ptr *QXmlName) QXmlName_PTR() *QXmlName {
	return ptr
}

func (ptr *QXmlName) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlName) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlName(ptr QXmlName_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlName_PTR().Pointer()
	}
	return nil
}

func NewQXmlNameFromPointer(ptr unsafe.Pointer) (n *QXmlName) {
	n = new(QXmlName)
	n.SetPointer(ptr)
	return
}
func (ptr *QXmlName) DestroyQXmlName() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQXmlName() *QXmlName {
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_NewQXmlName())
	qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func NewQXmlName2(namePool QXmlNamePool_ITF, localName string, namespaceURI string, prefix string) *QXmlName {
	var localNameC *C.char
	if localName != "" {
		localNameC = C.CString(localName)
		defer C.free(unsafe.Pointer(localNameC))
	}
	var namespaceURIC *C.char
	if namespaceURI != "" {
		namespaceURIC = C.CString(namespaceURI)
		defer C.free(unsafe.Pointer(namespaceURIC))
	}
	var prefixC *C.char
	if prefix != "" {
		prefixC = C.CString(prefix)
		defer C.free(unsafe.Pointer(prefixC))
	}
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_NewQXmlName2(PointerFromQXmlNamePool(namePool), C.struct_QtXmlPatterns_PackedString{data: localNameC, len: C.longlong(len(localName))}, C.struct_QtXmlPatterns_PackedString{data: namespaceURIC, len: C.longlong(len(namespaceURI))}, C.struct_QtXmlPatterns_PackedString{data: prefixC, len: C.longlong(len(prefix))}))
	qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func NewQXmlName3(other QXmlName_ITF) *QXmlName {
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_NewQXmlName3(PointerFromQXmlName(other)))
	qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func QXmlName_FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC *C.char
	if clarkName != "" {
		clarkNameC = C.CString(clarkName)
		defer C.free(unsafe.Pointer(clarkNameC))
	}
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(C.struct_QtXmlPatterns_PackedString{data: clarkNameC, len: C.longlong(len(clarkName))}, PointerFromQXmlNamePool(namePool)))
	qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func (ptr *QXmlName) FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC *C.char
	if clarkName != "" {
		clarkNameC = C.CString(clarkName)
		defer C.free(unsafe.Pointer(clarkNameC))
	}
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(C.struct_QtXmlPatterns_PackedString{data: clarkNameC, len: C.longlong(len(clarkName))}, PointerFromQXmlNamePool(namePool)))
	qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func QXmlName_IsNCName(candidate string) bool {
	var candidateC *C.char
	if candidate != "" {
		candidateC = C.CString(candidate)
		defer C.free(unsafe.Pointer(candidateC))
	}
	return int8(C.QXmlName_QXmlName_IsNCName(C.struct_QtXmlPatterns_PackedString{data: candidateC, len: C.longlong(len(candidate))})) != 0
}

func (ptr *QXmlName) IsNCName(candidate string) bool {
	var candidateC *C.char
	if candidate != "" {
		candidateC = C.CString(candidate)
		defer C.free(unsafe.Pointer(candidateC))
	}
	return int8(C.QXmlName_QXmlName_IsNCName(C.struct_QtXmlPatterns_PackedString{data: candidateC, len: C.longlong(len(candidate))})) != 0
}

func (ptr *QXmlName) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlName_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlName) LocalName(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_LocalName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) NamespaceUri(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_NamespaceUri(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) Prefix(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_Prefix(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

func (ptr *QXmlName) ToClarkName(namePool QXmlNamePool_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QXmlName_ToClarkName(ptr.Pointer(), PointerFromQXmlNamePool(namePool)))
	}
	return ""
}

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (ptr *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return ptr
}

func (ptr *QXmlNamePool) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlNamePool) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlNamePool(ptr QXmlNamePool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePool_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamePoolFromPointer(ptr unsafe.Pointer) (n *QXmlNamePool) {
	n = new(QXmlNamePool)
	n.SetPointer(ptr)
	return
}
func NewQXmlNamePool() *QXmlNamePool {
	tmpValue := NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool())
	qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	tmpValue := NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
	qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return ptr
}

func (ptr *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) (n *QXmlNodeModelIndex) {
	n = new(QXmlNodeModelIndex)
	n.SetPointer(ptr)
	return
}
func (ptr *QXmlNodeModelIndex) DestroyQXmlNodeModelIndex() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QXmlNodeModelIndex__NodeKind
//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  QXmlNodeModelIndex__NodeKind = QXmlNodeModelIndex__NodeKind(64)
)

//go:generate stringer -type=QXmlNodeModelIndex__DocumentOrder
//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(1)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
	qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
	qt.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
}

func (ptr *QXmlNodeModelIndex) AdditionalData() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_AdditionalData(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) Data() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QXmlNodeModelIndex_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_InternalPointer(ptr.Pointer())
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlNodeModelIndex_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	if ptr.Pointer() != nil {
		return NewQAbstractXmlNodeModelFromPointer(C.QXmlNodeModelIndex_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QXmlNodeModelIndex___namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_setList(i QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlNodeModelIndex___namespaceBindings_setList(ptr.Pointer(), PointerFromQXmlName(i))
	}
}

func (ptr *QXmlNodeModelIndex) __namespaceBindings_newList() unsafe.Pointer {
	return C.QXmlNodeModelIndex___namespaceBindings_newList(ptr.Pointer())
}

type QXmlQuery struct {
	ptr unsafe.Pointer
}

type QXmlQuery_ITF interface {
	QXmlQuery_PTR() *QXmlQuery
}

func (ptr *QXmlQuery) QXmlQuery_PTR() *QXmlQuery {
	return ptr
}

func (ptr *QXmlQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlQuery(ptr QXmlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlQuery_PTR().Pointer()
	}
	return nil
}

func NewQXmlQueryFromPointer(ptr unsafe.Pointer) (n *QXmlQuery) {
	n = new(QXmlQuery)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QXmlQuery__QueryLanguage
//QXmlQuery::QueryLanguage
type QXmlQuery__QueryLanguage int64

const (
	QXmlQuery__XQuery10                              QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1)
	QXmlQuery__XSLT20                                QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2)
	QXmlQuery__XmlSchema11IdentityConstraintSelector QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(1024)
	QXmlQuery__XmlSchema11IdentityConstraintField    QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(2048)
	QXmlQuery__XPath20                               QXmlQuery__QueryLanguage = QXmlQuery__QueryLanguage(4096)
)

func NewQXmlQuery() *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery())
	qt.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
	qt.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
	qt.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.longlong(queryLanguage), PointerFromQXmlNamePool(np)))
	qt.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable2(localName string, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		C.QXmlQuery_BindVariable2(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: localNameC, len: C.longlong(len(localName))}, PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable4(localName string, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		C.QXmlQuery_BindVariable4(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: localNameC, len: C.longlong(len(localName))}, core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable5(name QXmlName_ITF, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) BindVariable6(localName string, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		C.QXmlQuery_BindVariable6(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: localNameC, len: C.longlong(len(localName))}, PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) EvaluateTo(result QXmlResultItems_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_EvaluateTo(ptr.Pointer(), PointerFromQXmlResultItems(result))
	}
}

func (ptr *QXmlQuery) EvaluateTo2(callback QAbstractXmlReceiver_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_EvaluateTo2(ptr.Pointer(), PointerFromQAbstractXmlReceiver(callback))) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo3(target []string) bool {
	if ptr.Pointer() != nil {
		targetC := C.CString(strings.Join(target, "¡¦!"))
		defer C.free(unsafe.Pointer(targetC))
		return int8(C.QXmlQuery_EvaluateTo3(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: targetC, len: C.longlong(len(strings.Join(target, "¡¦!")))})) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo4(target core.QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_EvaluateTo4(ptr.Pointer(), core.PointerFromQIODevice(target))) != 0
	}
	return false
}

func (ptr *QXmlQuery) EvaluateTo5(output string) bool {
	if ptr.Pointer() != nil {
		var outputC *C.char
		if output != "" {
			outputC = C.CString(output)
			defer C.free(unsafe.Pointer(outputC))
		}
		return int8(C.QXmlQuery_EvaluateTo5(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: outputC, len: C.longlong(len(output))})) != 0
	}
	return false
}

func (ptr *QXmlQuery) InitialTemplateName() *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QXmlQuery_InitialTemplateName(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlQuery) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractMessageHandlerFromPointer(C.QXmlQuery_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlQuery_NamePool(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkAccessManagerFromPointer(C.QXmlQuery_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) QueryLanguage() QXmlQuery__QueryLanguage {
	if ptr.Pointer() != nil {
		return QXmlQuery__QueryLanguage(C.QXmlQuery_QueryLanguage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
	}
}

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI))) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus3(document core.QIODevice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_SetFocus3(ptr.Pointer(), core.PointerFromQIODevice(document))) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetFocus4(focus string) bool {
	if ptr.Pointer() != nil {
		var focusC *C.char
		if focus != "" {
			focusC = C.CString(focus)
			defer C.free(unsafe.Pointer(focusC))
		}
		return int8(C.QXmlQuery_SetFocus4(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: focusC, len: C.longlong(len(focus))})) != 0
	}
	return false
}

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlQuery) SetInitialTemplateName2(localName string) {
	if ptr.Pointer() != nil {
		var localNameC *C.char
		if localName != "" {
			localNameC = C.CString(localName)
			defer C.free(unsafe.Pointer(localNameC))
		}
		C.QXmlQuery_SetInitialTemplateName2(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: localNameC, len: C.longlong(len(localName))})
	}
}

func (ptr *QXmlQuery) SetMessageHandler(aMessageHandler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(aMessageHandler))
	}
}

func (ptr *QXmlQuery) SetNetworkAccessManager(newManager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(newManager))
	}
}

func (ptr *QXmlQuery) SetQuery(sourceCode core.QIODevice_ITF, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery(ptr.Pointer(), core.PointerFromQIODevice(sourceCode), core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery2(sourceCode string, documentURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		var sourceCodeC *C.char
		if sourceCode != "" {
			sourceCodeC = C.CString(sourceCode)
			defer C.free(unsafe.Pointer(sourceCodeC))
		}
		C.QXmlQuery_SetQuery2(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: sourceCodeC, len: C.longlong(len(sourceCode))}, core.PointerFromQUrl(documentURI))
	}
}

func (ptr *QXmlQuery) SetQuery3(queryURI core.QUrl_ITF, baseURI core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetQuery3(ptr.Pointer(), core.PointerFromQUrl(queryURI), core.PointerFromQUrl(baseURI))
	}
}

func (ptr *QXmlQuery) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlQuery) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractUriResolverFromPointer(C.QXmlQuery_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlQuery_DestroyQXmlQuery(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlResultItems struct {
	ptr unsafe.Pointer
}

type QXmlResultItems_ITF interface {
	QXmlResultItems_PTR() *QXmlResultItems
}

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func (ptr *QXmlResultItems) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlResultItems) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlResultItems(ptr QXmlResultItems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItems_PTR().Pointer()
	}
	return nil
}

func NewQXmlResultItemsFromPointer(ptr unsafe.Pointer) (n *QXmlResultItems) {
	n = new(QXmlResultItems)
	n.SetPointer(ptr)
	return
}
func NewQXmlResultItems() *QXmlResultItems {
	tmpValue := NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
	qt.SetFinalizer(tmpValue, (*QXmlResultItems).DestroyQXmlResultItems)
	return tmpValue
}

func (ptr *QXmlResultItems) Current() *QXmlItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlItemFromPointer(C.QXmlResultItems_Current(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlResultItems) HasError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlResultItems_HasError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlResultItems) Next() *QXmlItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlItemFromPointer(C.QXmlResultItems_Next(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

//export callbackQXmlResultItems_DestroyQXmlResultItems
func callbackQXmlResultItems_DestroyQXmlResultItems(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlResultItems"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlResultItemsFromPointer(ptr).DestroyQXmlResultItemsDefault()
	}
}

func (ptr *QXmlResultItems) ConnectDestroyQXmlResultItems(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlResultItems"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QXmlResultItems", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlResultItems", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlResultItems) DisconnectDestroyQXmlResultItems() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QXmlResultItems")
	}
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QXmlResultItems) DestroyQXmlResultItemsDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlResultItems_DestroyQXmlResultItemsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlSchema struct {
	ptr unsafe.Pointer
}

type QXmlSchema_ITF interface {
	QXmlSchema_PTR() *QXmlSchema
}

func (ptr *QXmlSchema) QXmlSchema_PTR() *QXmlSchema {
	return ptr
}

func (ptr *QXmlSchema) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlSchema) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlSchema(ptr QXmlSchema_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchema_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaFromPointer(ptr unsafe.Pointer) (n *QXmlSchema) {
	n = new(QXmlSchema)
	n.SetPointer(ptr)
	return
}
func NewQXmlSchema() *QXmlSchema {
	tmpValue := NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema())
	qt.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	tmpValue := NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
	qt.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
}

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QXmlSchema_DocumentUri(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchema_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source))) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchema_Load2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri))) != 0
	}
	return false
}

func (ptr *QXmlSchema) Load3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchema_Load3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri))) != 0
	}
	return false
}

func (ptr *QXmlSchema) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractMessageHandlerFromPointer(C.QXmlSchema_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlSchema_NamePool(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkAccessManagerFromPointer(C.QXmlSchema_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchema) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchema) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchema_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchema) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractUriResolverFromPointer(C.QXmlSchema_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlSchema_DestroyQXmlSchema(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlSchemaValidator struct {
	ptr unsafe.Pointer
}

type QXmlSchemaValidator_ITF interface {
	QXmlSchemaValidator_PTR() *QXmlSchemaValidator
}

func (ptr *QXmlSchemaValidator) QXmlSchemaValidator_PTR() *QXmlSchemaValidator {
	return ptr
}

func (ptr *QXmlSchemaValidator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQXmlSchemaValidator(ptr QXmlSchemaValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSchemaValidator_PTR().Pointer()
	}
	return nil
}

func NewQXmlSchemaValidatorFromPointer(ptr unsafe.Pointer) (n *QXmlSchemaValidator) {
	n = new(QXmlSchemaValidator)
	n.SetPointer(ptr)
	return
}
func NewQXmlSchemaValidator() *QXmlSchemaValidator {
	tmpValue := NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator())
	qt.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	tmpValue := NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
	qt.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
}

func (ptr *QXmlSchemaValidator) MessageHandler() *QAbstractMessageHandler {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractMessageHandlerFromPointer(C.QXmlSchemaValidator_MessageHandler(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlSchemaValidator_NamePool(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkAccessManagerFromPointer(C.QXmlSchemaValidator_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Schema() *QXmlSchema {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlSchemaFromPointer(C.QXmlSchemaValidator_Schema(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) SetMessageHandler(handler QAbstractMessageHandler_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetMessageHandler(ptr.Pointer(), PointerFromQAbstractMessageHandler(handler))
	}
}

func (ptr *QXmlSchemaValidator) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QXmlSchemaValidator) SetSchema(schema QXmlSchema_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetSchema(ptr.Pointer(), PointerFromQXmlSchema(schema))
	}
}

func (ptr *QXmlSchemaValidator) SetUriResolver(resolver QAbstractUriResolver_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_SetUriResolver(ptr.Pointer(), PointerFromQAbstractUriResolver(resolver))
	}
}

func (ptr *QXmlSchemaValidator) UriResolver() *QAbstractUriResolver {
	if ptr.Pointer() != nil {
		tmpValue := NewQAbstractUriResolverFromPointer(C.QXmlSchemaValidator_UriResolver(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source))) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate2(source core.QIODevice_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchemaValidator_Validate2(ptr.Pointer(), core.PointerFromQIODevice(source), core.PointerFromQUrl(documentUri))) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) Validate3(data core.QByteArray_ITF, documentUri core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchemaValidator_Validate3(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(documentUri))) != 0
	}
	return false
}

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QXmlSerializer struct {
	QAbstractXmlReceiver
}

type QXmlSerializer_ITF interface {
	QAbstractXmlReceiver_ITF
	QXmlSerializer_PTR() *QXmlSerializer
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func (ptr *QXmlSerializer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlReceiver_PTR().Pointer()
	}
	return nil
}

func (ptr *QXmlSerializer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractXmlReceiver_PTR().SetPointer(p)
	}
}

func PointerFromQXmlSerializer(ptr QXmlSerializer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlSerializer_PTR().Pointer()
	}
	return nil
}

func NewQXmlSerializerFromPointer(ptr unsafe.Pointer) (n *QXmlSerializer) {
	n = new(QXmlSerializer)
	n.SetPointer(ptr)
	return
}
func (ptr *QXmlSerializer) DestroyQXmlSerializer() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlSerializer_AtomicValue
func callbackQXmlSerializer_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "atomicValue"); signal != nil {
		(*(*func(*core.QVariant))(signal))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "atomicValue"); signal != nil {
			f := func(value *core.QVariant) {
				(*(*func(*core.QVariant))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectAtomicValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "atomicValue")
	}
}

func (ptr *QXmlSerializer) AtomicValue(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QXmlSerializer) AtomicValueDefault(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AtomicValueDefault(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

//export callbackQXmlSerializer_Attribute
func callbackQXmlSerializer_Attribute(ptr unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "attribute"); signal != nil {
		(*(*func(*QXmlName, *core.QStringRef))(signal))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "attribute"); signal != nil {
			f := func(name *QXmlName, value *core.QStringRef) {
				(*(*func(*QXmlName, *core.QStringRef))(signal))(name, value)
				f(name, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "attribute", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "attribute", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "attribute")
	}
}

func (ptr *QXmlSerializer) Attribute(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Attribute(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) AttributeDefault(name QXmlName_ITF, value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_AttributeDefault(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQStringRef(value))
	}
}

//export callbackQXmlSerializer_Characters
func callbackQXmlSerializer_Characters(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characters"); signal != nil {
		(*(*func(*core.QStringRef))(signal))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "characters"); signal != nil {
			f := func(value *core.QStringRef) {
				(*(*func(*core.QStringRef))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "characters", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characters", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectCharacters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "characters")
	}
}

func (ptr *QXmlSerializer) Characters(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_Characters(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) CharactersDefault(value core.QStringRef_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_CharactersDefault(ptr.Pointer(), core.PointerFromQStringRef(value))
	}
}

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}

//export callbackQXmlSerializer_Comment
func callbackQXmlSerializer_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CommentDefault(cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			f := func(value string) {
				(*(*func(string))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "comment", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "comment")
	}
}

func (ptr *QXmlSerializer) Comment(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QXmlSerializer_Comment(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QXmlSerializer) CommentDefault(value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QXmlSerializer_CommentDefault(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

//export callbackQXmlSerializer_EndDocument
func callbackQXmlSerializer_EndDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endDocument"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectEndDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endDocument")
	}
}

func (ptr *QXmlSerializer) EndDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_EndElement
func callbackQXmlSerializer_EndElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endElement"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endElement"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endElement", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endElement", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectEndElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endElement")
	}
}

func (ptr *QXmlSerializer) EndElement() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElement(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndElementDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndElementDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_EndOfSequence
func callbackQXmlSerializer_EndOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "endOfSequence"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endOfSequence"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectEndOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "endOfSequence")
	}
}

func (ptr *QXmlSerializer) EndOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) EndOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_EndOfSequenceDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_NamespaceBinding
func callbackQXmlSerializer_NamespaceBinding(ptr unsafe.Pointer, nb unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "namespaceBinding"); signal != nil {
		(*(*func(*QXmlName))(signal))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlSerializerFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlSerializer) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "namespaceBinding"); signal != nil {
			f := func(nb *QXmlName) {
				(*(*func(*QXmlName))(signal))(nb)
				f(nb)
			}
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectNamespaceBinding() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "namespaceBinding")
	}
}

func (ptr *QXmlSerializer) NamespaceBinding(nb QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBinding(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) NamespaceBindingDefault(nb QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_NamespaceBindingDefault(ptr.Pointer(), PointerFromQXmlName(nb))
	}
}

func (ptr *QXmlSerializer) OutputDevice() *core.QIODevice {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQIODeviceFromPointer(C.QXmlSerializer_OutputDevice(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQXmlSerializer_ProcessingInstruction
func callbackQXmlSerializer_ProcessingInstruction(ptr unsafe.Pointer, name unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "processingInstruction"); signal != nil {
		(*(*func(*QXmlName, string))(signal))(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processingInstruction"); signal != nil {
			f := func(name *QXmlName, value string) {
				(*(*func(*QXmlName, string))(signal))(name, value)
				f(name, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectProcessingInstruction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processingInstruction")
	}
}

func (ptr *QXmlSerializer) ProcessingInstruction(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QXmlSerializer_ProcessingInstruction(ptr.Pointer(), PointerFromQXmlName(name), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QXmlSerializer) ProcessingInstructionDefault(name QXmlName_ITF, value string) {
	if ptr.Pointer() != nil {
		var valueC *C.char
		if value != "" {
			valueC = C.CString(value)
			defer C.free(unsafe.Pointer(valueC))
		}
		C.QXmlSerializer_ProcessingInstructionDefault(ptr.Pointer(), PointerFromQXmlName(name), C.struct_QtXmlPatterns_PackedString{data: valueC, len: C.longlong(len(value))})
	}
}

func (ptr *QXmlSerializer) SetCodec(outputCodec core.QTextCodec_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_SetCodec(ptr.Pointer(), core.PointerFromQTextCodec(outputCodec))
	}
}

//export callbackQXmlSerializer_StartDocument
func callbackQXmlSerializer_StartDocument(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startDocument"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDocument"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startDocument", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectStartDocument() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startDocument")
	}
}

func (ptr *QXmlSerializer) StartDocument() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocument(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartDocumentDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartDocumentDefault(ptr.Pointer())
	}
}

//export callbackQXmlSerializer_StartElement
func callbackQXmlSerializer_StartElement(ptr unsafe.Pointer, name unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startElement"); signal != nil {
		(*(*func(*QXmlName))(signal))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlSerializerFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlSerializer) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startElement"); signal != nil {
			f := func(name *QXmlName) {
				(*(*func(*QXmlName))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "startElement", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startElement", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectStartElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startElement")
	}
}

func (ptr *QXmlSerializer) StartElement(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElement(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

func (ptr *QXmlSerializer) StartElementDefault(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartElementDefault(ptr.Pointer(), PointerFromQXmlName(name))
	}
}

//export callbackQXmlSerializer_StartOfSequence
func callbackQXmlSerializer_StartOfSequence(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "startOfSequence"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startOfSequence"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QXmlSerializer) DisconnectStartOfSequence() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "startOfSequence")
	}
}

func (ptr *QXmlSerializer) StartOfSequence() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequence(ptr.Pointer())
	}
}

func (ptr *QXmlSerializer) StartOfSequenceDefault() {
	if ptr.Pointer() != nil {
		C.QXmlSerializer_StartOfSequenceDefault(ptr.Pointer())
	}
}

func init() {
	qt.ItfMap["xmlpatterns.QAbstractMessageHandler_ITF"] = QAbstractMessageHandler{}
	qt.ItfMap["xmlpatterns.QAbstractUriResolver_ITF"] = QAbstractUriResolver{}
	qt.FuncMap["xmlpatterns.NewQAbstractUriResolver"] = NewQAbstractUriResolver
	qt.ItfMap["xmlpatterns.QAbstractXmlNodeModel_ITF"] = QAbstractXmlNodeModel{}
	qt.EnumMap["xmlpatterns.QAbstractXmlNodeModel__Parent"] = int64(QAbstractXmlNodeModel__Parent)
	qt.EnumMap["xmlpatterns.QAbstractXmlNodeModel__FirstChild"] = int64(QAbstractXmlNodeModel__FirstChild)
	qt.EnumMap["xmlpatterns.QAbstractXmlNodeModel__PreviousSibling"] = int64(QAbstractXmlNodeModel__PreviousSibling)
	qt.EnumMap["xmlpatterns.QAbstractXmlNodeModel__NextSibling"] = int64(QAbstractXmlNodeModel__NextSibling)
	qt.ItfMap["xmlpatterns.QAbstractXmlReceiver_ITF"] = QAbstractXmlReceiver{}
	qt.FuncMap["xmlpatterns.NewQAbstractXmlReceiver"] = NewQAbstractXmlReceiver
	qt.ItfMap["xmlpatterns.QSimpleXmlNodeModel_ITF"] = QSimpleXmlNodeModel{}
	qt.ItfMap["xmlpatterns.QSourceLocation_ITF"] = QSourceLocation{}
	qt.FuncMap["xmlpatterns.NewQSourceLocation"] = NewQSourceLocation
	qt.FuncMap["xmlpatterns.NewQSourceLocation2"] = NewQSourceLocation2
	qt.FuncMap["xmlpatterns.NewQSourceLocation3"] = NewQSourceLocation3
	qt.ItfMap["xmlpatterns.QXmlFormatter_ITF"] = QXmlFormatter{}
	qt.FuncMap["xmlpatterns.NewQXmlFormatter"] = NewQXmlFormatter
	qt.ItfMap["xmlpatterns.QXmlItem_ITF"] = QXmlItem{}
	qt.FuncMap["xmlpatterns.NewQXmlItem"] = NewQXmlItem
	qt.FuncMap["xmlpatterns.NewQXmlItem2"] = NewQXmlItem2
	qt.FuncMap["xmlpatterns.NewQXmlItem3"] = NewQXmlItem3
	qt.FuncMap["xmlpatterns.NewQXmlItem4"] = NewQXmlItem4
	qt.ItfMap["xmlpatterns.QXmlName_ITF"] = QXmlName{}
	qt.FuncMap["xmlpatterns.NewQXmlName"] = NewQXmlName
	qt.FuncMap["xmlpatterns.NewQXmlName2"] = NewQXmlName2
	qt.FuncMap["xmlpatterns.NewQXmlName3"] = NewQXmlName3
	qt.FuncMap["xmlpatterns.QXmlName_FromClarkName"] = QXmlName_FromClarkName
	qt.FuncMap["xmlpatterns.QXmlName_IsNCName"] = QXmlName_IsNCName
	qt.ItfMap["xmlpatterns.QXmlNamePool_ITF"] = QXmlNamePool{}
	qt.FuncMap["xmlpatterns.NewQXmlNamePool"] = NewQXmlNamePool
	qt.FuncMap["xmlpatterns.NewQXmlNamePool2"] = NewQXmlNamePool2
	qt.ItfMap["xmlpatterns.QXmlNodeModelIndex_ITF"] = QXmlNodeModelIndex{}
	qt.FuncMap["xmlpatterns.NewQXmlNodeModelIndex"] = NewQXmlNodeModelIndex
	qt.FuncMap["xmlpatterns.NewQXmlNodeModelIndex2"] = NewQXmlNodeModelIndex2
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Attribute"] = int64(QXmlNodeModelIndex__Attribute)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Comment"] = int64(QXmlNodeModelIndex__Comment)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Document"] = int64(QXmlNodeModelIndex__Document)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Element"] = int64(QXmlNodeModelIndex__Element)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Namespace"] = int64(QXmlNodeModelIndex__Namespace)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__ProcessingInstruction"] = int64(QXmlNodeModelIndex__ProcessingInstruction)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Text"] = int64(QXmlNodeModelIndex__Text)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Precedes"] = int64(QXmlNodeModelIndex__Precedes)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Is"] = int64(QXmlNodeModelIndex__Is)
	qt.EnumMap["xmlpatterns.QXmlNodeModelIndex__Follows"] = int64(QXmlNodeModelIndex__Follows)
	qt.ItfMap["xmlpatterns.QXmlQuery_ITF"] = QXmlQuery{}
	qt.FuncMap["xmlpatterns.NewQXmlQuery"] = NewQXmlQuery
	qt.FuncMap["xmlpatterns.NewQXmlQuery2"] = NewQXmlQuery2
	qt.FuncMap["xmlpatterns.NewQXmlQuery3"] = NewQXmlQuery3
	qt.FuncMap["xmlpatterns.NewQXmlQuery4"] = NewQXmlQuery4
	qt.EnumMap["xmlpatterns.QXmlQuery__XQuery10"] = int64(QXmlQuery__XQuery10)
	qt.EnumMap["xmlpatterns.QXmlQuery__XSLT20"] = int64(QXmlQuery__XSLT20)
	qt.EnumMap["xmlpatterns.QXmlQuery__XmlSchema11IdentityConstraintSelector"] = int64(QXmlQuery__XmlSchema11IdentityConstraintSelector)
	qt.EnumMap["xmlpatterns.QXmlQuery__XmlSchema11IdentityConstraintField"] = int64(QXmlQuery__XmlSchema11IdentityConstraintField)
	qt.EnumMap["xmlpatterns.QXmlQuery__XPath20"] = int64(QXmlQuery__XPath20)
	qt.ItfMap["xmlpatterns.QXmlResultItems_ITF"] = QXmlResultItems{}
	qt.FuncMap["xmlpatterns.NewQXmlResultItems"] = NewQXmlResultItems
	qt.ItfMap["xmlpatterns.QXmlSchema_ITF"] = QXmlSchema{}
	qt.FuncMap["xmlpatterns.NewQXmlSchema"] = NewQXmlSchema
	qt.FuncMap["xmlpatterns.NewQXmlSchema2"] = NewQXmlSchema2
	qt.ItfMap["xmlpatterns.QXmlSchemaValidator_ITF"] = QXmlSchemaValidator{}
	qt.FuncMap["xmlpatterns.NewQXmlSchemaValidator"] = NewQXmlSchemaValidator
	qt.FuncMap["xmlpatterns.NewQXmlSchemaValidator2"] = NewQXmlSchemaValidator2
	qt.ItfMap["xmlpatterns.QXmlSerializer_ITF"] = QXmlSerializer{}
	qt.FuncMap["xmlpatterns.NewQXmlSerializer"] = NewQXmlSerializer
}
