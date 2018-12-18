// +build !minimal

package xmlpatterns

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtXmlPatterns_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtXmlPatterns_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
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
func QAbstractMessageHandler_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractMessageHandler_QAbstractMessageHandler_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractMessageHandler) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractMessageHandler_QAbstractMessageHandler_Tr(sC, cC, C.int(int32(n))))
}

func QAbstractMessageHandler_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractMessageHandler_QAbstractMessageHandler_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractMessageHandler) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractMessageHandler_QAbstractMessageHandler_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler
func callbackQAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractMessageHandler"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DestroyQAbstractMessageHandlerDefault()
	}
}

func (ptr *QAbstractMessageHandler) ConnectDestroyQAbstractMessageHandler(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractMessageHandler"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractMessageHandler", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractMessageHandler", f)
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
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandlerDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandlerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAbstractMessageHandler_MetaObject
func callbackQAbstractMessageHandler_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractMessageHandlerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractMessageHandler) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractMessageHandler_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractMessageHandler___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QAbstractMessageHandler) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractMessageHandler___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractMessageHandler) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractMessageHandler) __findChildren_newList2() unsafe.Pointer {
	return C.QAbstractMessageHandler___findChildren_newList2(ptr.Pointer())
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

//export callbackQAbstractMessageHandler_Event
func callbackQAbstractMessageHandler_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractMessageHandlerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractMessageHandler) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractMessageHandler_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractMessageHandler_ChildEvent
func callbackQAbstractMessageHandler_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractMessageHandler) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAbstractMessageHandler_Destroyed
func callbackQAbstractMessageHandler_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractMessageHandler_DisconnectNotify
func callbackQAbstractMessageHandler_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractMessageHandlerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractMessageHandler) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractMessageHandler_ObjectNameChanged
func callbackQAbstractMessageHandler_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractMessageHandler_TimerEvent
func callbackQAbstractMessageHandler_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
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

func QAbstractUriResolver_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractUriResolver_QAbstractUriResolver_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractUriResolver) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractUriResolver_QAbstractUriResolver_Tr(sC, cC, C.int(int32(n))))
}

func QAbstractUriResolver_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractUriResolver_QAbstractUriResolver_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QAbstractUriResolver) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QAbstractUriResolver_QAbstractUriResolver_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQAbstractUriResolver_DestroyQAbstractUriResolver
func callbackQAbstractUriResolver_DestroyQAbstractUriResolver(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractUriResolver"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DestroyQAbstractUriResolverDefault()
	}
}

func (ptr *QAbstractUriResolver) ConnectDestroyQAbstractUriResolver(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractUriResolver"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractUriResolver", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractUriResolver", f)
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
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolverDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolverDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAbstractUriResolver_Resolve
func callbackQAbstractUriResolver_Resolve(ptr unsafe.Pointer, relative unsafe.Pointer, baseURI unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "resolve"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, *core.QUrl) *core.QUrl)(core.NewQUrlFromPointer(relative), core.NewQUrlFromPointer(baseURI)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractUriResolver) ConnectResolve(f func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resolve"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resolve", func(relative *core.QUrl, baseURI *core.QUrl) *core.QUrl {
				signal.(func(*core.QUrl, *core.QUrl) *core.QUrl)(relative, baseURI)
				return f(relative, baseURI)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resolve", f)
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
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractUriResolver_MetaObject
func callbackQAbstractUriResolver_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractUriResolverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractUriResolver) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractUriResolver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractUriResolver) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAbstractUriResolver___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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

func (ptr *QAbstractUriResolver) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QAbstractUriResolver___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractUriResolver) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QAbstractUriResolver) __findChildren_newList2() unsafe.Pointer {
	return C.QAbstractUriResolver___findChildren_newList2(ptr.Pointer())
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

//export callbackQAbstractUriResolver_Event
func callbackQAbstractUriResolver_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
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
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractUriResolverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractUriResolver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAbstractUriResolver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAbstractUriResolver_ChildEvent
func callbackQAbstractUriResolver_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractUriResolver) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAbstractUriResolver_Destroyed
func callbackQAbstractUriResolver_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQAbstractUriResolver_DisconnectNotify
func callbackQAbstractUriResolver_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractUriResolver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractUriResolver_ObjectNameChanged
func callbackQAbstractUriResolver_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQAbstractUriResolver_TimerEvent
func callbackQAbstractUriResolver_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
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

//export callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel
func callbackQAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractXmlNodeModel"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractXmlNodeModelFromPointer(ptr).DestroyQAbstractXmlNodeModelDefault()
	}
}

func (ptr *QAbstractXmlNodeModel) ConnectDestroyQAbstractXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractXmlNodeModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlNodeModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractXmlNodeModel", f)
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
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModelDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractXmlNodeModel) SourceLocation(index QXmlNodeModelIndex_ITF) *QSourceLocation {
	if ptr.Pointer() != nil {
		tmpValue := NewQSourceLocationFromPointer(C.QAbstractXmlNodeModel_SourceLocation(ptr.Pointer(), PointerFromQXmlNodeModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_StringValue
func callbackQAbstractXmlNodeModel_StringValue(ptr unsafe.Pointer, n unsafe.Pointer) C.struct_QtXmlPatterns_PackedString {
	if signal := qt.GetSignal(ptr, "stringValue"); signal != nil {
		tempVal := signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(n))
		return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QAbstractXmlNodeModel) ConnectStringValue(f func(n *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stringValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", func(n *QXmlNodeModelIndex) string {
				signal.(func(*QXmlNodeModelIndex) string)(n)
				return f(n)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", f)
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

//export callbackQAbstractXmlNodeModel_BaseUri
func callbackQAbstractXmlNodeModel_BaseUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectBaseUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "baseUri"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", func(n *QXmlNodeModelIndex) *core.QUrl {
				signal.(func(*QXmlNodeModelIndex) *core.QUrl)(n)
				return f(n)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", f)
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
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_DocumentUri
func callbackQAbstractXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QAbstractXmlNodeModel) ConnectDocumentUri(f func(n *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "documentUri"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "documentUri", func(n *QXmlNodeModelIndex) *core.QUrl {
				signal.(func(*QXmlNodeModelIndex) *core.QUrl)(n)
				return f(n)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "documentUri", f)
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
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_TypedValue
func callbackQAbstractXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QAbstractXmlNodeModel) ConnectTypedValue(f func(node *QXmlNodeModelIndex) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "typedValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "typedValue", func(node *QXmlNodeModelIndex) *core.QVariant {
				signal.(func(*QXmlNodeModelIndex) *core.QVariant)(node)
				return f(node)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typedValue", f)
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
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NamespaceBindings
func callbackQAbstractXmlNodeModel_NamespaceBindings(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "namespaceBindings"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
			for _, v := range signal.(func(*QXmlNodeModelIndex) []*QXmlName)(NewQXmlNodeModelIndexFromPointer(n)) {
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
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", func(n *QXmlNodeModelIndex) []*QXmlName {
				signal.(func(*QXmlNodeModelIndex) []*QXmlName)(n)
				return f(n)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", f)
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

//export callbackQAbstractXmlNodeModel_NodesByIdref
func callbackQAbstractXmlNodeModel_NodesByIdref(ptr unsafe.Pointer, idref unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nodesByIdref"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQAbstractXmlNodeModelFromPointer(NewQAbstractXmlNodeModelFromPointer(nil).__nodesByIdref_newList())
			for _, v := range signal.(func(*QXmlName) []*QXmlNodeModelIndex)(NewQXmlNameFromPointer(idref)) {
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
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", func(idref *QXmlName) []*QXmlNodeModelIndex {
				signal.(func(*QXmlName) []*QXmlNodeModelIndex)(idref)
				return f(idref)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", f)
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

//export callbackQAbstractXmlNodeModel_Name
func callbackQAbstractXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQXmlName())
}

func (ptr *QAbstractXmlNodeModel) ConnectName(f func(ni *QXmlNodeModelIndex) *QXmlName) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "name", func(ni *QXmlNodeModelIndex) *QXmlName {
				signal.(func(*QXmlNodeModelIndex) *QXmlName)(ni)
				return f(ni)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", f)
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
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex(data int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex(ptr.Pointer(), C.longlong(data)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex3(data int64, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex3(ptr.Pointer(), C.longlong(data), C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) CreateIndex2(pointer unsafe.Pointer, additionalData int64) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QAbstractXmlNodeModel_CreateIndex2(ptr.Pointer(), pointer, C.longlong(additionalData)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_ElementById
func callbackQAbstractXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementById"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementById", func(id *QXmlName) *QXmlNodeModelIndex {
				signal.(func(*QXmlName) *QXmlNodeModelIndex)(id)
				return f(id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementById", f)
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
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_NextFromSimpleAxis
func callbackQAbstractXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectNextFromSimpleAxis(f func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextFromSimpleAxis"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nextFromSimpleAxis", func(axis QAbstractXmlNodeModel__SimpleAxis, origin *QXmlNodeModelIndex) *QXmlNodeModelIndex {
				signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(axis, origin)
				return f(axis, origin)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextFromSimpleAxis", f)
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
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_Root
func callbackQAbstractXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQXmlNodeModelIndex())
}

func (ptr *QAbstractXmlNodeModel) ConnectRoot(f func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "root"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "root", func(n *QXmlNodeModelIndex) *QXmlNodeModelIndex {
				signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(n)
				return f(n)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "root", f)
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
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractXmlNodeModel_CompareOrder
func callbackQAbstractXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "compareOrder"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectCompareOrder(f func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "compareOrder"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "compareOrder", func(ni1 *QXmlNodeModelIndex, ni2 *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder {
				signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(ni1, ni2)
				return f(ni1, ni2)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "compareOrder", f)
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

//export callbackQAbstractXmlNodeModel_Kind
func callbackQAbstractXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "kind"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return C.longlong(0)
}

func (ptr *QAbstractXmlNodeModel) ConnectKind(f func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "kind"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "kind", func(ni *QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind {
				signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(ni)
				return f(ni)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "kind", f)
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

func (ptr *QAbstractXmlNodeModel) __namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QAbstractXmlNodeModel___namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
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
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
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

type QPatternist struct {
	ptr unsafe.Pointer
}

type QPatternist_ITF interface {
	QPatternist_PTR() *QPatternist
}

func (ptr *QPatternist) QPatternist_PTR() *QPatternist {
	return ptr
}

func (ptr *QPatternist) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPatternist) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPatternist(ptr QPatternist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPatternist_PTR().Pointer()
	}
	return nil
}

func NewQPatternistFromPointer(ptr unsafe.Pointer) (n *QPatternist) {
	n = new(QPatternist)
	n.SetPointer(ptr)
	return
}

func (ptr *QPatternist) DestroyQPatternist() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QPatternist__DerivedIntegerLimitsUsage
//QPatternist::DerivedIntegerLimitsUsage
type QPatternist__DerivedIntegerLimitsUsage int64

const (
	QPatternist__None           QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(1)
	QPatternist__LimitUpwards   QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(2)
	QPatternist__LimitDownwards QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(4)
	QPatternist__LimitBoth      QPatternist__DerivedIntegerLimitsUsage = QPatternist__DerivedIntegerLimitsUsage(QPatternist__LimitUpwards | QPatternist__LimitDownwards)
)

//go:generate stringer -type=QPatternist__TypeOfDerivedInteger
//QPatternist::TypeOfDerivedInteger
type QPatternist__TypeOfDerivedInteger int64

const (
	QPatternist__TypeByte               QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(0)
	QPatternist__TypeInt                QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(1)
	QPatternist__TypeLong               QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(2)
	QPatternist__TypeNegativeInteger    QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(3)
	QPatternist__TypeNonNegativeInteger QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(4)
	QPatternist__TypeNonPositiveInteger QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(5)
	QPatternist__TypePositiveInteger    QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(6)
	QPatternist__TypeShort              QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(7)
	QPatternist__TypeUnsignedByte       QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(8)
	QPatternist__TypeUnsignedInt        QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(9)
	QPatternist__TypeUnsignedLong       QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(10)
	QPatternist__TypeUnsignedShort      QPatternist__TypeOfDerivedInteger = QPatternist__TypeOfDerivedInteger(11)
)

//go:generate stringer -type=QPatternist__TypeOfDerivedString
//QPatternist::TypeOfDerivedString
type QPatternist__TypeOfDerivedString int64

const (
	QPatternist__TypeString           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(0)
	QPatternist__TypeNormalizedString QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(1)
	QPatternist__TypeToken            QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(2)
	QPatternist__TypeLanguage         QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(3)
	QPatternist__TypeNMTOKEN          QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(4)
	QPatternist__TypeName             QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(5)
	QPatternist__TypeNCName           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(6)
	QPatternist__TypeID               QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(7)
	QPatternist__TypeIDREF            QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(8)
	QPatternist__TypeENTITY           QPatternist__TypeOfDerivedString = QPatternist__TypeOfDerivedString(9)
)

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

//export callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel
func callbackQSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QSimpleXmlNodeModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSimpleXmlNodeModelFromPointer(ptr).DestroyQSimpleXmlNodeModelDefault()
	}
}

func (ptr *QSimpleXmlNodeModel) ConnectDestroyQSimpleXmlNodeModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QSimpleXmlNodeModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QSimpleXmlNodeModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QSimpleXmlNodeModel", f)
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
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModelDefault() {
	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQSimpleXmlNodeModel_StringValue
func callbackQSimpleXmlNodeModel_StringValue(ptr unsafe.Pointer, node unsafe.Pointer) C.struct_QtXmlPatterns_PackedString {
	if signal := qt.GetSignal(ptr, "stringValue"); signal != nil {
		tempVal := signal.(func(*QXmlNodeModelIndex) string)(NewQXmlNodeModelIndexFromPointer(node))
		return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQSimpleXmlNodeModelFromPointer(ptr).StringValueDefault(NewQXmlNodeModelIndexFromPointer(node))
	return C.struct_QtXmlPatterns_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QSimpleXmlNodeModel) ConnectStringValue(f func(node *QXmlNodeModelIndex) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stringValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", func(node *QXmlNodeModelIndex) string {
				signal.(func(*QXmlNodeModelIndex) string)(node)
				return f(node)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stringValue", f)
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

//export callbackQSimpleXmlNodeModel_BaseUri
func callbackQSimpleXmlNodeModel_BaseUri(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "baseUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).BaseUriDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) ConnectBaseUri(f func(node *QXmlNodeModelIndex) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "baseUri"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", func(node *QXmlNodeModelIndex) *core.QUrl {
				signal.(func(*QXmlNodeModelIndex) *core.QUrl)(node)
				return f(node)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "baseUri", f)
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
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) BaseUriDefault(node QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NamespaceBindings
func callbackQSimpleXmlNodeModel_NamespaceBindings(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "namespaceBindings"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQSimpleXmlNodeModelFromPointer(NewQSimpleXmlNodeModelFromPointer(nil).__namespaceBindings_newList())
			for _, v := range signal.(func(*QXmlNodeModelIndex) []*QXmlName)(NewQXmlNodeModelIndexFromPointer(node)) {
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
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", func(node *QXmlNodeModelIndex) []*QXmlName {
				signal.(func(*QXmlNodeModelIndex) []*QXmlName)(node)
				return f(node)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBindings", f)
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
			for _, v := range signal.(func(*QXmlName) []*QXmlNodeModelIndex)(NewQXmlNameFromPointer(idref)) {
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
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", func(idref *QXmlName) []*QXmlNodeModelIndex {
				signal.(func(*QXmlName) []*QXmlNodeModelIndex)(idref)
				return f(idref)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodesByIdref", f)
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

func (ptr *QSimpleXmlNodeModel) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QSimpleXmlNodeModel_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_ElementById
func callbackQSimpleXmlNodeModel_ElementById(ptr unsafe.Pointer, id unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "elementById"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlName) *QXmlNodeModelIndex)(NewQXmlNameFromPointer(id)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).ElementByIdDefault(NewQXmlNameFromPointer(id)))
}

func (ptr *QSimpleXmlNodeModel) ConnectElementById(f func(id *QXmlName) *QXmlNodeModelIndex) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementById"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementById", func(id *QXmlName) *QXmlNodeModelIndex {
				signal.(func(*QXmlName) *QXmlNodeModelIndex)(id)
				return f(id)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementById", f)
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
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) ElementByIdDefault(id QXmlName_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_ElementByIdDefault(ptr.Pointer(), PointerFromQXmlName(id)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_DocumentUri
func callbackQSimpleXmlNodeModel_DocumentUri(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "documentUri"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*QXmlNodeModelIndex) *core.QUrl)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return core.PointerFromQUrl(NewQSimpleXmlNodeModelFromPointer(ptr).DocumentUriDefault(NewQXmlNodeModelIndexFromPointer(n)))
}

func (ptr *QSimpleXmlNodeModel) DocumentUri(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) DocumentUriDefault(n QXmlNodeModelIndex_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_DocumentUriDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_TypedValue
func callbackQSimpleXmlNodeModel_TypedValue(ptr unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "typedValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(*QXmlNodeModelIndex) *core.QVariant)(NewQXmlNodeModelIndexFromPointer(node)))
	}

	return core.PointerFromQVariant(NewQSimpleXmlNodeModelFromPointer(ptr).TypedValueDefault(NewQXmlNodeModelIndexFromPointer(node)))
}

func (ptr *QSimpleXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSimpleXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) TypedValueDefault(node QXmlNodeModelIndex_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QSimpleXmlNodeModel_TypedValueDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Name
func callbackQSimpleXmlNodeModel_Name(ptr unsafe.Pointer, ni unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		return PointerFromQXmlName(signal.(func(*QXmlNodeModelIndex) *QXmlName)(NewQXmlNodeModelIndexFromPointer(ni)))
	}

	return PointerFromQXmlName(NewQSimpleXmlNodeModelFromPointer(ptr).NameDefault(NewQXmlNodeModelIndexFromPointer(ni)))
}

func (ptr *QSimpleXmlNodeModel) Name(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_Name(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NameDefault(ni QXmlNodeModelIndex_ITF) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QSimpleXmlNodeModel_NameDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_NextFromSimpleAxis
func callbackQSimpleXmlNodeModel_NextFromSimpleAxis(ptr unsafe.Pointer, axis C.longlong, origin unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextFromSimpleAxis"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(QAbstractXmlNodeModel__SimpleAxis, *QXmlNodeModelIndex) *QXmlNodeModelIndex)(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).NextFromSimpleAxisDefault(QAbstractXmlNodeModel__SimpleAxis(axis), NewQXmlNodeModelIndexFromPointer(origin)))
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxis(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxis(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) NextFromSimpleAxisDefault(axis QAbstractXmlNodeModel__SimpleAxis, origin QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_NextFromSimpleAxisDefault(ptr.Pointer(), C.longlong(axis), PointerFromQXmlNodeModelIndex(origin)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_Root
func callbackQSimpleXmlNodeModel_Root(ptr unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "root"); signal != nil {
		return PointerFromQXmlNodeModelIndex(signal.(func(*QXmlNodeModelIndex) *QXmlNodeModelIndex)(NewQXmlNodeModelIndexFromPointer(n)))
	}

	return PointerFromQXmlNodeModelIndex(NewQSimpleXmlNodeModelFromPointer(ptr).RootDefault(NewQXmlNodeModelIndexFromPointer(n)))
}

func (ptr *QSimpleXmlNodeModel) Root(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_Root(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) RootDefault(n QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QSimpleXmlNodeModel_RootDefault(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSimpleXmlNodeModel_CompareOrder
func callbackQSimpleXmlNodeModel_CompareOrder(ptr unsafe.Pointer, ni1 unsafe.Pointer, ni2 unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "compareOrder"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex, *QXmlNodeModelIndex) QXmlNodeModelIndex__DocumentOrder)(NewQXmlNodeModelIndexFromPointer(ni1), NewQXmlNodeModelIndexFromPointer(ni2)))
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

//export callbackQSimpleXmlNodeModel_Kind
func callbackQSimpleXmlNodeModel_Kind(ptr unsafe.Pointer, ni unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "kind"); signal != nil {
		return C.longlong(signal.(func(*QXmlNodeModelIndex) QXmlNodeModelIndex__NodeKind)(NewQXmlNodeModelIndexFromPointer(ni)))
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
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	tmpValue := NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	tmpValue := NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(int32(l)), C.int(int32(c))))
	runtime.SetFinalizer(tmpValue, (*QSourceLocation).DestroyQSourceLocation)
	return tmpValue
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

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	if ptr.Pointer() != nil {
		C.QSourceLocation_DestroyQSourceLocation(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QSourceLocation) Uri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QSourceLocation_Uri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QSourceLocation) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QSourceLocation_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSourceLocation) Column() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSourceLocation) Line() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QSourceLocation_Line(ptr.Pointer()))
	}
	return 0
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
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQXmlFormatter(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlFormatter {
	return NewQXmlFormatterFromPointer(C.QXmlFormatter_NewQXmlFormatter(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

func (ptr *QXmlFormatter) SetIndentationDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QXmlFormatter_SetIndentationDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QXmlFormatter) IndentationDepth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QXmlFormatter_IndentationDepth(ptr.Pointer())))
	}
	return 0
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
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	tmpValue := NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
	runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
	return tmpValue
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlItem) ToAtomicValue() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QXmlItem_ToAtomicValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlItem) ToNodeModelIndex() *QXmlNodeModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlItem_ToNodeModelIndex(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
		return tmpValue
	}
	return nil
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
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QXmlName_FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC *C.char
	if clarkName != "" {
		clarkNameC = C.CString(clarkName)
		defer C.free(unsafe.Pointer(clarkNameC))
	}
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(C.struct_QtXmlPatterns_PackedString{data: clarkNameC, len: C.longlong(len(clarkName))}, PointerFromQXmlNamePool(namePool)))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func (ptr *QXmlName) FromClarkName(clarkName string, namePool QXmlNamePool_ITF) *QXmlName {
	var clarkNameC *C.char
	if clarkName != "" {
		clarkNameC = C.CString(clarkName)
		defer C.free(unsafe.Pointer(clarkNameC))
	}
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_QXmlName_FromClarkName(C.struct_QtXmlPatterns_PackedString{data: clarkNameC, len: C.longlong(len(clarkName))}, PointerFromQXmlNamePool(namePool)))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func NewQXmlName() *QXmlName {
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_NewQXmlName())
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
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
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
	return tmpValue
}

func NewQXmlName3(other QXmlName_ITF) *QXmlName {
	tmpValue := NewQXmlNameFromPointer(C.QXmlName_NewQXmlName3(PointerFromQXmlName(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
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

func (ptr *QXmlName) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlName_IsNull(ptr.Pointer())) != 0
	}
	return false
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
	runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	tmpValue := NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
	return tmpValue
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QXmlNodeModelIndex__DocumentOrder
//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  QXmlNodeModelIndex__DocumentOrder = QXmlNodeModelIndex__DocumentOrder(1)
)

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

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
	runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	tmpValue := NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlNodeModelIndex).DestroyQXmlNodeModelIndex)
	return tmpValue
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

func (ptr *QXmlNodeModelIndex) __namespaceBindings_atList(i int) *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QXmlNodeModelIndex___namespaceBindings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
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
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery4(queryLanguage QXmlQuery__QueryLanguage, np QXmlNamePool_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery4(C.longlong(queryLanguage), PointerFromQXmlNamePool(np)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery3(np QXmlNamePool_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery3(PointerFromQXmlNamePool(np)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
}

func NewQXmlQuery2(other QXmlQuery_ITF) *QXmlQuery {
	tmpValue := NewQXmlQueryFromPointer(C.QXmlQuery_NewQXmlQuery2(PointerFromQXmlQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlQuery).DestroyQXmlQuery)
	return tmpValue
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

func (ptr *QXmlQuery) SetFocus2(documentURI core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_SetFocus2(ptr.Pointer(), core.PointerFromQUrl(documentURI))) != 0
	}
	return false
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

func (ptr *QXmlQuery) BindVariable3(name QXmlName_ITF, device core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable3(ptr.Pointer(), PointerFromQXmlName(name), core.PointerFromQIODevice(device))
	}
}

func (ptr *QXmlQuery) BindVariable(name QXmlName_ITF, value QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlItem(value))
	}
}

func (ptr *QXmlQuery) BindVariable5(name QXmlName_ITF, query QXmlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_BindVariable5(ptr.Pointer(), PointerFromQXmlName(name), PointerFromQXmlQuery(query))
	}
}

func (ptr *QXmlQuery) SetFocus(item QXmlItem_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetFocus(ptr.Pointer(), PointerFromQXmlItem(item))
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

func (ptr *QXmlQuery) SetInitialTemplateName(name QXmlName_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_SetInitialTemplateName(ptr.Pointer(), PointerFromQXmlName(name))
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

func (ptr *QXmlQuery) DestroyQXmlQuery() {
	if ptr.Pointer() != nil {
		C.QXmlQuery_DestroyQXmlQuery(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func (ptr *QXmlQuery) InitialTemplateName() *QXmlName {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNameFromPointer(C.QXmlQuery_InitialTemplateName(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlName).DestroyQXmlName)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlQuery) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlQuery_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
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

func (ptr *QXmlQuery) EvaluateTo3(target []string) bool {
	if ptr.Pointer() != nil {
		targetC := C.CString(strings.Join(target, "|"))
		defer C.free(unsafe.Pointer(targetC))
		return int8(C.QXmlQuery_EvaluateTo3(ptr.Pointer(), C.struct_QtXmlPatterns_PackedString{data: targetC, len: C.longlong(len(strings.Join(target, "|")))})) != 0
	}
	return false
}

func (ptr *QXmlQuery) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlQuery_IsValid(ptr.Pointer())) != 0
	}
	return false
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

func (ptr *QXmlQuery) EvaluateTo(result QXmlResultItems_ITF) {
	if ptr.Pointer() != nil {
		C.QXmlQuery_EvaluateTo(ptr.Pointer(), PointerFromQXmlResultItems(result))
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
func (ptr *QXmlResultItems) Next() *QXmlItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlItemFromPointer(C.QXmlResultItems_Next(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
		return tmpValue
	}
	return nil
}

func NewQXmlResultItems() *QXmlResultItems {
	return NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

//export callbackQXmlResultItems_DestroyQXmlResultItems
func callbackQXmlResultItems_DestroyQXmlResultItems(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QXmlResultItems"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlResultItemsFromPointer(ptr).DestroyQXmlResultItemsDefault()
	}
}

func (ptr *QXmlResultItems) ConnectDestroyQXmlResultItems(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QXmlResultItems"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlResultItems", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QXmlResultItems", f)
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
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlResultItems) DestroyQXmlResultItemsDefault() {
	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItemsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QXmlResultItems) Current() *QXmlItem {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlItemFromPointer(C.QXmlResultItems_Current(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlItem).DestroyQXmlItem)
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
	runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
}

func NewQXmlSchema2(other QXmlSchema_ITF) *QXmlSchema {
	tmpValue := NewQXmlSchemaFromPointer(C.QXmlSchema_NewQXmlSchema2(PointerFromQXmlSchema(other)))
	runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
	return tmpValue
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

func (ptr *QXmlSchema) Load(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchema_Load(ptr.Pointer(), core.PointerFromQUrl(source))) != 0
	}
	return false
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

func (ptr *QXmlSchema) DestroyQXmlSchema() {
	if ptr.Pointer() != nil {
		C.QXmlSchema_DestroyQXmlSchema(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func (ptr *QXmlSchema) DocumentUri() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QXmlSchema_DocumentUri(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchema) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlSchema_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
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
	runtime.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
}

func NewQXmlSchemaValidator2(schema QXmlSchema_ITF) *QXmlSchemaValidator {
	tmpValue := NewQXmlSchemaValidatorFromPointer(C.QXmlSchemaValidator_NewQXmlSchemaValidator2(PointerFromQXmlSchema(schema)))
	runtime.SetFinalizer(tmpValue, (*QXmlSchemaValidator).DestroyQXmlSchemaValidator)
	return tmpValue
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

func (ptr *QXmlSchemaValidator) DestroyQXmlSchemaValidator() {
	if ptr.Pointer() != nil {
		C.QXmlSchemaValidator_DestroyQXmlSchemaValidator(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func (ptr *QXmlSchemaValidator) NamePool() *QXmlNamePool {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlNamePoolFromPointer(C.QXmlSchemaValidator_NamePool(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlNamePool).DestroyQXmlNamePool)
		return tmpValue
	}
	return nil
}

func (ptr *QXmlSchemaValidator) Schema() *QXmlSchema {
	if ptr.Pointer() != nil {
		tmpValue := NewQXmlSchemaFromPointer(C.QXmlSchemaValidator_Schema(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QXmlSchema).DestroyQXmlSchema)
		return tmpValue
	}
	return nil
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

func (ptr *QXmlSchemaValidator) Validate(source core.QUrl_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QXmlSchemaValidator_Validate(ptr.Pointer(), core.PointerFromQUrl(source))) != 0
	}
	return false
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

type QXmlSerializer struct {
	ptr unsafe.Pointer
}

type QXmlSerializer_ITF interface {
	QXmlSerializer_PTR() *QXmlSerializer
}

func (ptr *QXmlSerializer) QXmlSerializer_PTR() *QXmlSerializer {
	return ptr
}

func (ptr *QXmlSerializer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QXmlSerializer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQXmlSerializer(query QXmlQuery_ITF, outputDevice core.QIODevice_ITF) *QXmlSerializer {
	return NewQXmlSerializerFromPointer(C.QXmlSerializer_NewQXmlSerializer(PointerFromQXmlQuery(query), core.PointerFromQIODevice(outputDevice)))
}

//export callbackQXmlSerializer_AtomicValue
func callbackQXmlSerializer_AtomicValue(ptr unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "atomicValue"); signal != nil {
		signal.(func(*core.QVariant))(core.NewQVariantFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AtomicValueDefault(core.NewQVariantFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAtomicValue(f func(value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "atomicValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", func(value *core.QVariant) {
				signal.(func(*core.QVariant))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "atomicValue", f)
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
		signal.(func(*QXmlName, *core.QStringRef))(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).AttributeDefault(NewQXmlNameFromPointer(name), core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectAttribute(f func(name *QXmlName, value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "attribute"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "attribute", func(name *QXmlName, value *core.QStringRef) {
				signal.(func(*QXmlName, *core.QStringRef))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "attribute", f)
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
		signal.(func(*core.QStringRef))(core.NewQStringRefFromPointer(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CharactersDefault(core.NewQStringRefFromPointer(value))
	}
}

func (ptr *QXmlSerializer) ConnectCharacters(f func(value *core.QStringRef)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "characters"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characters", func(value *core.QStringRef) {
				signal.(func(*core.QStringRef))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characters", f)
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

//export callbackQXmlSerializer_Comment
func callbackQXmlSerializer_Comment(ptr unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		signal.(func(string))(cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).CommentDefault(cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectComment(f func(value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "comment", func(value string) {
				signal.(func(string))(value)
				f(value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", f)
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
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endDocument", f)
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
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndElementDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndElement(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endElement", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endElement", f)
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
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).EndOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectEndOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "endOfSequence"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "endOfSequence", f)
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
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(nb))
	} else {
		NewQXmlSerializerFromPointer(ptr).NamespaceBindingDefault(NewQXmlNameFromPointer(nb))
	}
}

func (ptr *QXmlSerializer) ConnectNamespaceBinding(f func(nb *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "namespaceBinding"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", func(nb *QXmlName) {
				signal.(func(*QXmlName))(nb)
				f(nb)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "namespaceBinding", f)
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

//export callbackQXmlSerializer_ProcessingInstruction
func callbackQXmlSerializer_ProcessingInstruction(ptr unsafe.Pointer, name unsafe.Pointer, value C.struct_QtXmlPatterns_PackedString) {
	if signal := qt.GetSignal(ptr, "processingInstruction"); signal != nil {
		signal.(func(*QXmlName, string))(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	} else {
		NewQXmlSerializerFromPointer(ptr).ProcessingInstructionDefault(NewQXmlNameFromPointer(name), cGoUnpackString(value))
	}
}

func (ptr *QXmlSerializer) ConnectProcessingInstruction(f func(name *QXmlName, value string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processingInstruction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", func(name *QXmlName, value string) {
				signal.(func(*QXmlName, string))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processingInstruction", f)
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
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartDocumentDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartDocument(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startDocument"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startDocument", f)
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
		signal.(func(*QXmlName))(NewQXmlNameFromPointer(name))
	} else {
		NewQXmlSerializerFromPointer(ptr).StartElementDefault(NewQXmlNameFromPointer(name))
	}
}

func (ptr *QXmlSerializer) ConnectStartElement(f func(name *QXmlName)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startElement", func(name *QXmlName) {
				signal.(func(*QXmlName))(name)
				f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startElement", f)
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
		signal.(func())()
	} else {
		NewQXmlSerializerFromPointer(ptr).StartOfSequenceDefault()
	}
}

func (ptr *QXmlSerializer) ConnectStartOfSequence(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "startOfSequence"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "startOfSequence", f)
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

func (ptr *QXmlSerializer) Codec() *core.QTextCodec {
	if ptr.Pointer() != nil {
		return core.NewQTextCodecFromPointer(C.QXmlSerializer_Codec(ptr.Pointer()))
	}
	return nil
}
