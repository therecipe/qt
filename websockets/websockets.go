// +build !minimal

package websockets

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtWebSockets_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWebSockets_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QMaskGenerator struct {
	core.QObject
}

type QMaskGenerator_ITF interface {
	core.QObject_ITF
	QMaskGenerator_PTR() *QMaskGenerator
}

func (ptr *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QMaskGenerator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) (n *QMaskGenerator) {
	n = new(QMaskGenerator)
	n.SetPointer(ptr)
	return
}
func NewQMaskGenerator2(parent core.QObject_ITF) *QMaskGenerator {
	tmpValue := NewQMaskGeneratorFromPointer(C.QMaskGenerator_NewQMaskGenerator2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMaskGenerator_NextMask
func callbackQMaskGenerator_NextMask(ptr unsafe.Pointer) C.uint {
	if signal := qt.GetSignal(ptr, "nextMask"); signal != nil {
		return C.uint(uint32((*(*func() uint)(signal))()))
	}

	return C.uint(uint32(0))
}

func (ptr *QMaskGenerator) ConnectNextMask(f func() uint) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextMask"); signal != nil {
			f := func() uint {
				(*(*func() uint)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nextMask", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextMask", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMaskGenerator) DisconnectNextMask() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nextMask")
	}
}

func (ptr *QMaskGenerator) NextMask() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QMaskGenerator_NextMask(ptr.Pointer())))
	}
	return 0
}

//export callbackQMaskGenerator_Seed
func callbackQMaskGenerator_Seed(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "seed"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMaskGenerator) ConnectSeed(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "seed"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "seed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMaskGenerator) DisconnectSeed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "seed")
	}
}

func (ptr *QMaskGenerator) Seed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QMaskGenerator_Seed(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQMaskGenerator_DestroyQMaskGenerator
func callbackQMaskGenerator_DestroyQMaskGenerator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMaskGenerator"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DestroyQMaskGeneratorDefault()
	}
}

func (ptr *QMaskGenerator) ConnectDestroyQMaskGenerator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMaskGenerator"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QMaskGenerator", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMaskGenerator", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QMaskGenerator) DisconnectDestroyQMaskGenerator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QMaskGenerator")
	}
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) DestroyQMaskGeneratorDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMaskGenerator_DestroyQMaskGeneratorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMaskGenerator___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMaskGenerator) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMaskGenerator) __children_newList() unsafe.Pointer {
	return C.QMaskGenerator___children_newList(ptr.Pointer())
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QMaskGenerator___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QMaskGenerator___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QMaskGenerator) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMaskGenerator___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMaskGenerator) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMaskGenerator) __findChildren_newList() unsafe.Pointer {
	return C.QMaskGenerator___findChildren_newList(ptr.Pointer())
}

func (ptr *QMaskGenerator) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QMaskGenerator___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMaskGenerator) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMaskGenerator) __findChildren_newList3() unsafe.Pointer {
	return C.QMaskGenerator___findChildren_newList3(ptr.Pointer())
}

//export callbackQMaskGenerator_ChildEvent
func callbackQMaskGenerator_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMaskGenerator_ConnectNotify
func callbackQMaskGenerator_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_CustomEvent
func callbackQMaskGenerator_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMaskGenerator_DeleteLater
func callbackQMaskGenerator_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMaskGenerator) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QMaskGenerator_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQMaskGenerator_Destroyed
func callbackQMaskGenerator_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQMaskGenerator_DisconnectNotify
func callbackQMaskGenerator_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_Event
func callbackQMaskGenerator_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMaskGenerator) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMaskGenerator_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQMaskGenerator_EventFilter
func callbackQMaskGenerator_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMaskGenerator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QMaskGenerator_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQMaskGenerator_MetaObject
func callbackQMaskGenerator_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQMaskGeneratorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMaskGenerator) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMaskGenerator_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQMaskGenerator_ObjectNameChanged
func callbackQMaskGenerator_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQMaskGenerator_TimerEvent
func callbackQMaskGenerator_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlWebSocket struct {
	core.QObject
}

type QQmlWebSocket_ITF interface {
	core.QObject_ITF
	QQmlWebSocket_PTR() *QQmlWebSocket
}

func (ptr *QQmlWebSocket) QQmlWebSocket_PTR() *QQmlWebSocket {
	return ptr
}

func (ptr *QQmlWebSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWebSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWebSocket(ptr QQmlWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWebSocket_PTR().Pointer()
	}
	return nil
}

func NewQQmlWebSocketFromPointer(ptr unsafe.Pointer) (n *QQmlWebSocket) {
	n = new(QQmlWebSocket)
	n.SetPointer(ptr)
	return
}

type QQmlWebSocketServer struct {
	core.QObject
}

type QQmlWebSocketServer_ITF interface {
	core.QObject_ITF
	QQmlWebSocketServer_PTR() *QQmlWebSocketServer
}

func (ptr *QQmlWebSocketServer) QQmlWebSocketServer_PTR() *QQmlWebSocketServer {
	return ptr
}

func (ptr *QQmlWebSocketServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWebSocketServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWebSocketServer(ptr QQmlWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWebSocketServer_PTR().Pointer()
	}
	return nil
}

func NewQQmlWebSocketServerFromPointer(ptr unsafe.Pointer) (n *QQmlWebSocketServer) {
	n = new(QQmlWebSocketServer)
	n.SetPointer(ptr)
	return
}

type QWebSocket struct {
	core.QObject
}

type QWebSocket_ITF interface {
	core.QObject_ITF
	QWebSocket_PTR() *QWebSocket
}

func (ptr *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebSocket(ptr QWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocket_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketFromPointer(ptr unsafe.Pointer) (n *QWebSocket) {
	n = new(QWebSocket)
	n.SetPointer(ptr)
	return
}
func NewQWebSocket2(origin string, version QWebSocketProtocol__Version, parent core.QObject_ITF) *QWebSocket {
	var originC *C.char
	if origin != "" {
		originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
	}
	tmpValue := NewQWebSocketFromPointer(C.QWebSocket_NewQWebSocket2(C.struct_QtWebSockets_PackedString{data: originC, len: C.longlong(len(origin))}, C.longlong(version), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

//export callbackQWebSocket_AboutToClose
func callbackQWebSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "aboutToClose") {
			C.QWebSocket_ConnectAboutToClose(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "aboutToClose")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "aboutToClose"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "aboutToClose", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "aboutToClose", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "aboutToClose")
	}
}

func (ptr *QWebSocket) AboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_AboutToClose(ptr.Pointer())
	}
}

//export callbackQWebSocket_BinaryFrameReceived
func callbackQWebSocket_BinaryFrameReceived(ptr unsafe.Pointer, frame unsafe.Pointer, isLastFrame C.char) {
	if signal := qt.GetSignal(ptr, "binaryFrameReceived"); signal != nil {
		(*(*func(*core.QByteArray, bool))(signal))(core.NewQByteArrayFromPointer(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "binaryFrameReceived") {
			C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "binaryFrameReceived")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "binaryFrameReceived"); signal != nil {
			f := func(frame *core.QByteArray, isLastFrame bool) {
				(*(*func(*core.QByteArray, bool))(signal))(frame, isLastFrame)
				f(frame, isLastFrame)
			}
			qt.ConnectSignal(ptr.Pointer(), "binaryFrameReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryFrameReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "binaryFrameReceived")
	}
}

func (ptr *QWebSocket) BinaryFrameReceived(frame core.QByteArray_ITF, isLastFrame bool) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryFrameReceived(ptr.Pointer(), core.PointerFromQByteArray(frame), C.char(int8(qt.GoBoolToInt(isLastFrame))))
	}
}

//export callbackQWebSocket_BinaryMessageReceived
func callbackQWebSocket_BinaryMessageReceived(ptr unsafe.Pointer, message unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "binaryMessageReceived"); signal != nil {
		(*(*func(*core.QByteArray))(signal))(core.NewQByteArrayFromPointer(message))
	}

}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "binaryMessageReceived") {
			C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "binaryMessageReceived")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "binaryMessageReceived"); signal != nil {
			f := func(message *core.QByteArray) {
				(*(*func(*core.QByteArray))(signal))(message)
				f(message)
			}
			qt.ConnectSignal(ptr.Pointer(), "binaryMessageReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryMessageReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "binaryMessageReceived")
	}
}

func (ptr *QWebSocket) BinaryMessageReceived(message core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryMessageReceived(ptr.Pointer(), core.PointerFromQByteArray(message))
	}
}

func (ptr *QWebSocket) BytesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_BytesWritten
func callbackQWebSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "bytesWritten") {
			C.QWebSocket_ConnectBytesWritten(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "bytesWritten")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "bytesWritten"); signal != nil {
			f := func(bytes int64) {
				(*(*func(int64))(signal))(bytes)
				f(bytes)
			}
			qt.ConnectSignal(ptr.Pointer(), "bytesWritten", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bytesWritten", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectBytesWritten() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "bytesWritten")
	}
}

func (ptr *QWebSocket) BytesWritten(bytes int64) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BytesWritten(ptr.Pointer(), C.longlong(bytes))
	}
}

//export callbackQWebSocket_Close
func callbackQWebSocket_Close(ptr unsafe.Pointer, closeCode C.longlong, reason C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func(QWebSocketProtocol__CloseCode, string))(signal))(QWebSocketProtocol__CloseCode(closeCode), cGoUnpackString(reason))
	} else {
		NewQWebSocketFromPointer(ptr).CloseDefault(QWebSocketProtocol__CloseCode(closeCode), cGoUnpackString(reason))
	}
}

func (ptr *QWebSocket) ConnectClose(f func(closeCode QWebSocketProtocol__CloseCode, reason string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			f := func(closeCode QWebSocketProtocol__CloseCode, reason string) {
				(*(*func(QWebSocketProtocol__CloseCode, string))(signal))(closeCode, reason)
				f(closeCode, reason)
			}
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QWebSocket) Close(closeCode QWebSocketProtocol__CloseCode, reason string) {
	if ptr.Pointer() != nil {
		var reasonC *C.char
		if reason != "" {
			reasonC = C.CString(reason)
			defer C.free(unsafe.Pointer(reasonC))
		}
		C.QWebSocket_Close(ptr.Pointer(), C.longlong(closeCode), C.struct_QtWebSockets_PackedString{data: reasonC, len: C.longlong(len(reason))})
	}
}

func (ptr *QWebSocket) CloseDefault(closeCode QWebSocketProtocol__CloseCode, reason string) {
	if ptr.Pointer() != nil {
		var reasonC *C.char
		if reason != "" {
			reasonC = C.CString(reason)
			defer C.free(unsafe.Pointer(reasonC))
		}
		C.QWebSocket_CloseDefault(ptr.Pointer(), C.longlong(closeCode), C.struct_QtWebSockets_PackedString{data: reasonC, len: C.longlong(len(reason))})
	}
}

func (ptr *QWebSocket) CloseCode() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocket_CloseCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) CloseReason() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebSocket_Connected
func callbackQWebSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QWebSocket_ConnectConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connected")
	}
}

func (ptr *QWebSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Connected(ptr.Pointer())
	}
}

//export callbackQWebSocket_Disconnected
func callbackQWebSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QWebSocket_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QWebSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_Error2
func callbackQWebSocket_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(network.QAbstractSocket__SocketError))(signal))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QWebSocket_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error network.QAbstractSocket__SocketError) {
				(*(*func(network.QAbstractSocket__SocketError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QWebSocket) Error2(error network.QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Error2(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QWebSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocket_Flush(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQWebSocket_IgnoreSslErrors
func callbackQWebSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ignoreSslErrors"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebSocketFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QWebSocket) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignoreSslErrors"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ignoreSslErrors")
	}
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IgnoreSslErrorsDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrorsDefault(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IgnoreSslErrors2(errors []*network.QSslError) {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWebSocketFromPointer(NewQWebSocketFromPointer(nil).__ignoreSslErrors_errors_newList2())
			for _, v := range errors {
				tmpList.__ignoreSslErrors_errors_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocket_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocket) LocalAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQHostAddressFromPointer(C.QWebSocket_LocalAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	if ptr.Pointer() != nil {
		tmpValue := NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebSocket_Open
func callbackQWebSocket_Open(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		(*(*func(*core.QUrl))(signal))(core.NewQUrlFromPointer(url))
	} else {
		NewQWebSocketFromPointer(ptr).OpenDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QWebSocket) ConnectOpen(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			f := func(url *core.QUrl) {
				(*(*func(*core.QUrl))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) OpenDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_OpenDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebSocket_Open2
func callbackQWebSocket_Open2(ptr unsafe.Pointer, request unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open2"); signal != nil {
		(*(*func(*network.QNetworkRequest))(signal))(network.NewQNetworkRequestFromPointer(request))
	} else {
		NewQWebSocketFromPointer(ptr).Open2Default(network.NewQNetworkRequestFromPointer(request))
	}
}

func (ptr *QWebSocket) ConnectOpen2(f func(request *network.QNetworkRequest)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open2"); signal != nil {
			f := func(request *network.QNetworkRequest) {
				(*(*func(*network.QNetworkRequest))(signal))(request)
				f(request)
			}
			qt.ConnectSignal(ptr.Pointer(), "open2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open2", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectOpen2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open2")
	}
}

func (ptr *QWebSocket) Open2(request network.QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open2(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

func (ptr *QWebSocket) Open2Default(request network.QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open2Default(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

func (ptr *QWebSocket) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQHostAddressFromPointer(C.QWebSocket_PeerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_Ping
func callbackQWebSocket_Ping(ptr unsafe.Pointer, payload unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ping"); signal != nil {
		(*(*func(*core.QByteArray))(signal))(core.NewQByteArrayFromPointer(payload))
	} else {
		NewQWebSocketFromPointer(ptr).PingDefault(core.NewQByteArrayFromPointer(payload))
	}
}

func (ptr *QWebSocket) ConnectPing(f func(payload *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ping"); signal != nil {
			f := func(payload *core.QByteArray) {
				(*(*func(*core.QByteArray))(signal))(payload)
				f(payload)
			}
			qt.ConnectSignal(ptr.Pointer(), "ping", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ping", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectPing() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "ping")
	}
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QWebSocket) PingDefault(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_PingDefault(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

//export callbackQWebSocket_Pong
func callbackQWebSocket_Pong(ptr unsafe.Pointer, elapsedTime C.ulonglong, payload unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pong"); signal != nil {
		(*(*func(uint64, *core.QByteArray))(signal))(uint64(elapsedTime), core.NewQByteArrayFromPointer(payload))
	}

}

func (ptr *QWebSocket) ConnectPong(f func(elapsedTime uint64, payload *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pong") {
			C.QWebSocket_ConnectPong(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pong")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pong"); signal != nil {
			f := func(elapsedTime uint64, payload *core.QByteArray) {
				(*(*func(uint64, *core.QByteArray))(signal))(elapsedTime, payload)
				f(elapsedTime, payload)
			}
			qt.ConnectSignal(ptr.Pointer(), "pong", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pong", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectPong() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectPong(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pong")
	}
}

func (ptr *QWebSocket) Pong(elapsedTime uint64, payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Pong(ptr.Pointer(), C.ulonglong(elapsedTime), core.PointerFromQByteArray(payload))
	}
}

//export callbackQWebSocket_PreSharedKeyAuthenticationRequired
func callbackQWebSocket_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preSharedKeyAuthenticationRequired"); signal != nil {
		(*(*func(*network.QSslPreSharedKeyAuthenticator))(signal))(network.NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QWebSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preSharedKeyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			f := func(authenticator *network.QSslPreSharedKeyAuthenticator) {
				(*(*func(*network.QSslPreSharedKeyAuthenticator))(signal))(authenticator)
				f(authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QWebSocket) PreSharedKeyAuthenticationRequired(authenticator network.QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_PreSharedKeyAuthenticationRequired(ptr.Pointer(), network.PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QWebSocket) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkProxyFromPointer(C.QWebSocket_Proxy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocket_ProxyAuthenticationRequired
func callbackQWebSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "proxyAuthenticationRequired"); signal != nil {
		(*(*func(*network.QNetworkProxy, *network.QAuthenticator))(signal))(network.NewQNetworkProxyFromPointer(proxy), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectProxyAuthenticationRequired(f func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "proxyAuthenticationRequired") {
			C.QWebSocket_ConnectProxyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "proxyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "proxyAuthenticationRequired"); signal != nil {
			f := func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator) {
				(*(*func(*network.QNetworkProxy, *network.QAuthenticator))(signal))(proxy, authenticator)
				f(proxy, authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "proxyAuthenticationRequired")
	}
}

func (ptr *QWebSocket) ProxyAuthenticationRequired(proxy network.QNetworkProxy_ITF, authenticator network.QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ProxyAuthenticationRequired(ptr.Pointer(), network.PointerFromQNetworkProxy(proxy), network.PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QWebSocket) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_ReadChannelFinished
func callbackQWebSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "readChannelFinished") {
			C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "readChannelFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "readChannelFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "readChannelFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readChannelFinished", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "readChannelFinished")
	}
}

func (ptr *QWebSocket) ReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_ReadChannelFinished(ptr.Pointer())
	}
}

func (ptr *QWebSocket) Request() *network.QNetworkRequest {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkRequestFromPointer(C.QWebSocket_Request(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QNetworkRequest).DestroyQNetworkRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebSocket_RequestUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) ResourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SendBinaryMessage(data core.QByteArray_ITF) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendBinaryMessage(ptr.Pointer(), core.PointerFromQByteArray(data)))
	}
	return 0
}

func (ptr *QWebSocket) SendTextMessage(message string) int64 {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		return int64(C.QWebSocket_SendTextMessage(ptr.Pointer(), C.struct_QtWebSockets_PackedString{data: messageC, len: C.longlong(len(message))}))
	}
	return 0
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.longlong(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslConfigurationFromPointer(C.QWebSocket_SslConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocket_SslErrors
func callbackQWebSocket_SslErrors(ptr unsafe.Pointer, errors C.struct_QtWebSockets_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		(*(*func([]*network.QSslError))(signal))(func(l C.struct_QtWebSockets_PackedList) []*network.QSslError {
			out := make([]*network.QSslError, int(l.len))
			tmpList := NewQWebSocketFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QWebSocket) ConnectSslErrors(f func(errors []*network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QWebSocket_ConnectSslErrors(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sslErrors")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			f := func(errors []*network.QSslError) {
				(*(*func([]*network.QSslError))(signal))(errors)
				f(errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectSslErrors(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sslErrors")
	}
}

func (ptr *QWebSocket) SslErrors(errors []*network.QSslError) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SslErrors(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWebSocketFromPointer(NewQWebSocketFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketState(C.QWebSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_StateChanged
func callbackQWebSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(network.QAbstractSocket__SocketState))(signal))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QWebSocket_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state network.QAbstractSocket__SocketState) {
				(*(*func(network.QAbstractSocket__SocketState))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QWebSocket) StateChanged(state network.QAbstractSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QWebSocket_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQWebSocket_TextFrameReceived
func callbackQWebSocket_TextFrameReceived(ptr unsafe.Pointer, frame C.struct_QtWebSockets_PackedString, isLastFrame C.char) {
	if signal := qt.GetSignal(ptr, "textFrameReceived"); signal != nil {
		(*(*func(string, bool))(signal))(cGoUnpackString(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textFrameReceived") {
			C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "textFrameReceived")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textFrameReceived"); signal != nil {
			f := func(frame string, isLastFrame bool) {
				(*(*func(string, bool))(signal))(frame, isLastFrame)
				f(frame, isLastFrame)
			}
			qt.ConnectSignal(ptr.Pointer(), "textFrameReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textFrameReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textFrameReceived")
	}
}

func (ptr *QWebSocket) TextFrameReceived(frame string, isLastFrame bool) {
	if ptr.Pointer() != nil {
		var frameC *C.char
		if frame != "" {
			frameC = C.CString(frame)
			defer C.free(unsafe.Pointer(frameC))
		}
		C.QWebSocket_TextFrameReceived(ptr.Pointer(), C.struct_QtWebSockets_PackedString{data: frameC, len: C.longlong(len(frame))}, C.char(int8(qt.GoBoolToInt(isLastFrame))))
	}
}

//export callbackQWebSocket_TextMessageReceived
func callbackQWebSocket_TextMessageReceived(ptr unsafe.Pointer, message C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "textMessageReceived"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(message))
	}

}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textMessageReceived") {
			C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "textMessageReceived")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textMessageReceived"); signal != nil {
			f := func(message string) {
				(*(*func(string))(signal))(message)
				f(message)
			}
			qt.ConnectSignal(ptr.Pointer(), "textMessageReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textMessageReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "textMessageReceived")
	}
}

func (ptr *QWebSocket) TextMessageReceived(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QWebSocket_TextMessageReceived(ptr.Pointer(), C.struct_QtWebSockets_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *QWebSocket) Version() QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__Version(C.QWebSocket_Version(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_DestroyQWebSocket
func callbackQWebSocket_DestroyQWebSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebSocketFromPointer(ptr).DestroyQWebSocketDefault()
	}
}

func (ptr *QWebSocket) ConnectDestroyQWebSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocket", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocket) DisconnectDestroyQWebSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebSocket")
	}
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) DestroyQWebSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocket_DestroyQWebSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_atList2(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslErrorFromPointer(C.QWebSocket___ignoreSslErrors_errors_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_setList2(i network.QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___ignoreSslErrors_errors_setList2(ptr.Pointer(), network.PointerFromQSslError(i))
	}
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_newList2() unsafe.Pointer {
	return C.QWebSocket___ignoreSslErrors_errors_newList2(ptr.Pointer())
}

func (ptr *QWebSocket) __sslErrors_errors_atList(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslErrorFromPointer(C.QWebSocket___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __sslErrors_errors_setList(i network.QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___sslErrors_errors_setList(ptr.Pointer(), network.PointerFromQSslError(i))
	}
}

func (ptr *QWebSocket) __sslErrors_errors_newList() unsafe.Pointer {
	return C.QWebSocket___sslErrors_errors_newList(ptr.Pointer())
}

func (ptr *QWebSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocket) __children_newList() unsafe.Pointer {
	return C.QWebSocket___children_newList(ptr.Pointer())
}

func (ptr *QWebSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebSocket) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebSocket___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocket) __findChildren_newList() unsafe.Pointer {
	return C.QWebSocket___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocket) __findChildren_newList3() unsafe.Pointer {
	return C.QWebSocket___findChildren_newList3(ptr.Pointer())
}

//export callbackQWebSocket_ChildEvent
func callbackQWebSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocket_ConnectNotify
func callbackQWebSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_CustomEvent
func callbackQWebSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocket_DeleteLater
func callbackQWebSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocket_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWebSocket_Destroyed
func callbackQWebSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebSocket_DisconnectNotify
func callbackQWebSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_Event
func callbackQWebSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebSocket_EventFilter
func callbackQWebSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebSocket_MetaObject
func callbackQWebSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebSocket_ObjectNameChanged
func callbackQWebSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebSocket_TimerEvent
func callbackQWebSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWebSocketCorsAuthenticator struct {
	ptr unsafe.Pointer
}

type QWebSocketCorsAuthenticator_ITF interface {
	QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator
}

func (ptr *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator {
	return ptr
}

func (ptr *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebSocketCorsAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QWebSocketCorsAuthenticator) {
	n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return
}
func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	var originC *C.char
	if origin != "" {
		originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
	}
	tmpValue := NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.struct_QtWebSockets_PackedString{data: originC, len: C.longlong(len(origin))}))
	qt.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	tmpValue := NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
	qt.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	tmpValue := NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
	qt.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allowed))))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(other))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebSocketProtocol struct {
	ptr unsafe.Pointer
}

type QWebSocketProtocol_ITF interface {
	QWebSocketProtocol_PTR() *QWebSocketProtocol
}

func (ptr *QWebSocketProtocol) QWebSocketProtocol_PTR() *QWebSocketProtocol {
	return ptr
}

func (ptr *QWebSocketProtocol) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWebSocketProtocol) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWebSocketProtocol(ptr QWebSocketProtocol_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketProtocol_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketProtocolFromPointer(ptr unsafe.Pointer) (n *QWebSocketProtocol) {
	n = new(QWebSocketProtocol)
	n.SetPointer(ptr)
	return
}
func (ptr *QWebSocketProtocol) DestroyQWebSocketProtocol() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QWebSocketProtocol__CloseCode
//QWebSocketProtocol::CloseCode
type QWebSocketProtocol__CloseCode int64

const (
	QWebSocketProtocol__CloseCodeNormal                QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1000)
	QWebSocketProtocol__CloseCodeGoingAway             QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1001)
	QWebSocketProtocol__CloseCodeProtocolError         QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1002)
	QWebSocketProtocol__CloseCodeDatatypeNotSupported  QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1003)
	QWebSocketProtocol__CloseCodeReserved1004          QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1004)
	QWebSocketProtocol__CloseCodeMissingStatusCode     QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1005)
	QWebSocketProtocol__CloseCodeAbnormalDisconnection QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1006)
	QWebSocketProtocol__CloseCodeWrongDatatype         QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1007)
	QWebSocketProtocol__CloseCodePolicyViolated        QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1008)
	QWebSocketProtocol__CloseCodeTooMuchData           QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1009)
	QWebSocketProtocol__CloseCodeMissingExtension      QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1010)
	QWebSocketProtocol__CloseCodeBadOperation          QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1011)
	QWebSocketProtocol__CloseCodeTlsHandshakeFailed    QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1015)
)

//go:generate stringer -type=QWebSocketProtocol__Version
//QWebSocketProtocol::Version
type QWebSocketProtocol__Version int64

const (
	QWebSocketProtocol__VersionUnknown QWebSocketProtocol__Version = QWebSocketProtocol__Version(-1)
	QWebSocketProtocol__Version0       QWebSocketProtocol__Version = QWebSocketProtocol__Version(0)
	QWebSocketProtocol__Version4       QWebSocketProtocol__Version = QWebSocketProtocol__Version(4)
	QWebSocketProtocol__Version5       QWebSocketProtocol__Version = QWebSocketProtocol__Version(5)
	QWebSocketProtocol__Version6       QWebSocketProtocol__Version = QWebSocketProtocol__Version(6)
	QWebSocketProtocol__Version7       QWebSocketProtocol__Version = QWebSocketProtocol__Version(7)
	QWebSocketProtocol__Version8       QWebSocketProtocol__Version = QWebSocketProtocol__Version(8)
	QWebSocketProtocol__Version13      QWebSocketProtocol__Version = QWebSocketProtocol__Version(13)
	QWebSocketProtocol__VersionLatest  QWebSocketProtocol__Version = QWebSocketProtocol__Version(QWebSocketProtocol__Version13)
)

type QWebSocketServer struct {
	core.QObject
}

type QWebSocketServer_ITF interface {
	core.QObject_ITF
	QWebSocketServer_PTR() *QWebSocketServer
}

func (ptr *QWebSocketServer) QWebSocketServer_PTR() *QWebSocketServer {
	return ptr
}

func (ptr *QWebSocketServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebSocketServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebSocketServer(ptr QWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketServer_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketServerFromPointer(ptr unsafe.Pointer) (n *QWebSocketServer) {
	n = new(QWebSocketServer)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QWebSocketServer__SslMode
//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int64

const (
	QWebSocketServer__SecureMode    QWebSocketServer__SslMode = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode QWebSocketServer__SslMode = QWebSocketServer__SslMode(1)
)

func NewQWebSocketServer2(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {
	var serverNameC *C.char
	if serverName != "" {
		serverNameC = C.CString(serverName)
		defer C.free(unsafe.Pointer(serverNameC))
	}
	tmpValue := NewQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer2(C.struct_QtWebSockets_PackedString{data: serverNameC, len: C.longlong(len(serverName))}, C.longlong(secureMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebSocketServer_AcceptError
func callbackQWebSocketServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(ptr, "acceptError"); signal != nil {
		(*(*func(network.QAbstractSocket__SocketError))(signal))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "acceptError") {
			C.QWebSocketServer_ConnectAcceptError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "acceptError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "acceptError"); signal != nil {
			f := func(socketError network.QAbstractSocket__SocketError) {
				(*(*func(network.QAbstractSocket__SocketError))(signal))(socketError)
				f(socketError)
			}
			qt.ConnectSignal(ptr.Pointer(), "acceptError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "acceptError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "acceptError")
	}
}

func (ptr *QWebSocketServer) AcceptError(socketError network.QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_AcceptError(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QWebSocketServer) Close() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_Close(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_Closed
func callbackQWebSocketServer_Closed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closed"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closed") {
			C.QWebSocketServer_ConnectClosed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "closed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closed"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "closed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "closed")
	}
}

func (ptr *QWebSocketServer) Closed() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_Closed(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) Error() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocketServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) HandleConnection(socket network.QTcpSocket_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_HandleConnection(ptr.Pointer(), network.PointerFromQTcpSocket(socket))
	}
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketServer_HasPendingConnections(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketServer_IsListening(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketServer) Listen(address network.QHostAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketServer_Listen(ptr.Pointer(), network.PointerFromQHostAddress(address), C.ushort(port))) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSocketServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQWebSocketServer_NewConnection
func callbackQWebSocketServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QWebSocketServer_ConnectNewConnection(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newConnection")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConnection")
	}
}

func (ptr *QWebSocketServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_NextPendingConnection
func callbackQWebSocketServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextPendingConnection"); signal != nil {
		return PointerFromQWebSocket((*(*func() *QWebSocket)(signal))())
	}

	return PointerFromQWebSocket(NewQWebSocketServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QWebSocketServer) ConnectNextPendingConnection(f func() *QWebSocket) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextPendingConnection"); signal != nil {
			f := func() *QWebSocket {
				(*(*func() *QWebSocket)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nextPendingConnection")
	}
}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) NextPendingConnectionDefault() *QWebSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_OriginAuthenticationRequired
func callbackQWebSocketServer_OriginAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "originAuthenticationRequired"); signal != nil {
		(*(*func(*QWebSocketCorsAuthenticator))(signal))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "originAuthenticationRequired") {
			C.QWebSocketServer_ConnectOriginAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "originAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "originAuthenticationRequired"); signal != nil {
			f := func(authenticator *QWebSocketCorsAuthenticator) {
				(*(*func(*QWebSocketCorsAuthenticator))(signal))(authenticator)
				f(authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "originAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "originAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectOriginAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "originAuthenticationRequired")
	}
}

func (ptr *QWebSocketServer) OriginAuthenticationRequired(authenticator QWebSocketCorsAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_OriginAuthenticationRequired(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(authenticator))
	}
}

func (ptr *QWebSocketServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PauseAccepting(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_PeerVerifyError
func callbackQWebSocketServer_PeerVerifyError(ptr unsafe.Pointer, error unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "peerVerifyError"); signal != nil {
		(*(*func(*network.QSslError))(signal))(network.NewQSslErrorFromPointer(error))
	}

}

func (ptr *QWebSocketServer) ConnectPeerVerifyError(f func(error *network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peerVerifyError") {
			C.QWebSocketServer_ConnectPeerVerifyError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "peerVerifyError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peerVerifyError"); signal != nil {
			f := func(error *network.QSslError) {
				(*(*func(*network.QSslError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectPeerVerifyError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "peerVerifyError")
	}
}

func (ptr *QWebSocketServer) PeerVerifyError(error network.QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PeerVerifyError(ptr.Pointer(), network.PointerFromQSslError(error))
	}
}

//export callbackQWebSocketServer_PreSharedKeyAuthenticationRequired
func callbackQWebSocketServer_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "preSharedKeyAuthenticationRequired"); signal != nil {
		(*(*func(*network.QSslPreSharedKeyAuthenticator))(signal))(network.NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QWebSocketServer_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "preSharedKeyAuthenticationRequired")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			f := func(authenticator *network.QSslPreSharedKeyAuthenticator) {
				(*(*func(*network.QSslPreSharedKeyAuthenticator))(signal))(authenticator)
				f(authenticator)
			}
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QWebSocketServer) PreSharedKeyAuthenticationRequired(authenticator network.QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PreSharedKeyAuthenticationRequired(ptr.Pointer(), network.PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QWebSocketServer) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkProxyFromPointer(C.QWebSocketServer_Proxy(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQHostAddressFromPointer(C.QWebSocketServer_ServerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_ServerError
func callbackQWebSocketServer_ServerError(ptr unsafe.Pointer, closeCode C.longlong) {
	if signal := qt.GetSignal(ptr, "serverError"); signal != nil {
		(*(*func(QWebSocketProtocol__CloseCode))(signal))(QWebSocketProtocol__CloseCode(closeCode))
	}

}

func (ptr *QWebSocketServer) ConnectServerError(f func(closeCode QWebSocketProtocol__CloseCode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serverError") {
			C.QWebSocketServer_ConnectServerError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "serverError")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serverError"); signal != nil {
			f := func(closeCode QWebSocketProtocol__CloseCode) {
				(*(*func(QWebSocketProtocol__CloseCode))(signal))(closeCode)
				f(closeCode)
			}
			qt.ConnectSignal(ptr.Pointer(), "serverError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serverError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectServerError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectServerError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serverError")
	}
}

func (ptr *QWebSocketServer) ServerError(closeCode QWebSocketProtocol__CloseCode) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ServerError(ptr.Pointer(), C.longlong(closeCode))
	}
}

func (ptr *QWebSocketServer) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocketServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QWebSocketServer_ServerUrl(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {
	if ptr.Pointer() != nil {
		var serverNameC *C.char
		if serverName != "" {
			serverNameC = C.CString(serverName)
			defer C.free(unsafe.Pointer(serverNameC))
		}
		C.QWebSocketServer_SetServerName(ptr.Pointer(), C.struct_QtWebSockets_PackedString{data: serverNameC, len: C.longlong(len(serverName))})
	}
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocketServer) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslConfigurationFromPointer(C.QWebSocketServer_SslConfiguration(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_SslErrors
func callbackQWebSocketServer_SslErrors(ptr unsafe.Pointer, errors C.struct_QtWebSockets_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		(*(*func([]*network.QSslError))(signal))(func(l C.struct_QtWebSockets_PackedList) []*network.QSslError {
			out := make([]*network.QSslError, int(l.len))
			tmpList := NewQWebSocketServerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QWebSocketServer) ConnectSslErrors(f func(errors []*network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QWebSocketServer_ConnectSslErrors(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sslErrors")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			f := func(errors []*network.QSslError) {
				(*(*func([]*network.QSslError))(signal))(errors)
				f(errors)
			}
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectSslErrors(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sslErrors")
	}
}

func (ptr *QWebSocketServer) SslErrors(errors []*network.QSslError) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SslErrors(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWebSocketServerFromPointer(NewQWebSocketServerFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QWebSocketServer) SupportedVersions() []QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebSockets_PackedList) []QWebSocketProtocol__Version {
			out := make([]QWebSocketProtocol__Version, int(l.len))
			tmpList := NewQWebSocketServerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__supportedVersions_atList(i)
			}
			return out
		}(C.QWebSocketServer_SupportedVersions(ptr.Pointer()))
	}
	return make([]QWebSocketProtocol__Version, 0)
}

//export callbackQWebSocketServer_DestroyQWebSocketServer
func callbackQWebSocketServer_DestroyQWebSocketServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebSocketServer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebSocketServerFromPointer(ptr).DestroyQWebSocketServerDefault()
	}
}

func (ptr *QWebSocketServer) ConnectDestroyQWebSocketServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebSocketServer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocketServer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocketServer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWebSocketServer) DisconnectDestroyQWebSocketServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWebSocketServer")
	}
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocketServer_DestroyQWebSocketServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocketServer) DestroyQWebSocketServerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocketServer_DestroyQWebSocketServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocketServer) __sslErrors_errors_atList(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQSslErrorFromPointer(C.QWebSocketServer___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __sslErrors_errors_setList(i network.QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___sslErrors_errors_setList(ptr.Pointer(), network.PointerFromQSslError(i))
	}
}

func (ptr *QWebSocketServer) __sslErrors_errors_newList() unsafe.Pointer {
	return C.QWebSocketServer___sslErrors_errors_newList(ptr.Pointer())
}

func (ptr *QWebSocketServer) __supportedVersions_atList(i int) QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__Version(C.QWebSocketServer___supportedVersions_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QWebSocketServer) __supportedVersions_setList(i QWebSocketProtocol__Version) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___supportedVersions_setList(ptr.Pointer(), C.longlong(i))
	}
}

func (ptr *QWebSocketServer) __supportedVersions_newList() unsafe.Pointer {
	return C.QWebSocketServer___supportedVersions_newList(ptr.Pointer())
}

func (ptr *QWebSocketServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocketServer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocketServer) __children_newList() unsafe.Pointer {
	return C.QWebSocketServer___children_newList(ptr.Pointer())
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWebSocketServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWebSocketServer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWebSocketServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocketServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocketServer) __findChildren_newList() unsafe.Pointer {
	return C.QWebSocketServer___findChildren_newList(ptr.Pointer())
}

func (ptr *QWebSocketServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWebSocketServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocketServer) __findChildren_newList3() unsafe.Pointer {
	return C.QWebSocketServer___findChildren_newList3(ptr.Pointer())
}

//export callbackQWebSocketServer_ChildEvent
func callbackQWebSocketServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocketServer_ConnectNotify
func callbackQWebSocketServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_CustomEvent
func callbackQWebSocketServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocketServer_DeleteLater
func callbackQWebSocketServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWebSocketServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocketServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWebSocketServer_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_Destroyed
func callbackQWebSocketServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebSocketServer_DisconnectNotify
func callbackQWebSocketServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_Event
func callbackQWebSocketServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocketServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWebSocketServer_EventFilter
func callbackQWebSocketServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocketServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWebSocketServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWebSocketServer_MetaObject
func callbackQWebSocketServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWebSocketServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocketServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebSocketServer_ObjectNameChanged
func callbackQWebSocketServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWebSocketServer_TimerEvent
func callbackQWebSocketServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QtWebSocketsDeclarativeModule struct {
	ptr unsafe.Pointer
}

type QtWebSocketsDeclarativeModule_ITF interface {
	QtWebSocketsDeclarativeModule_PTR() *QtWebSocketsDeclarativeModule
}

func (ptr *QtWebSocketsDeclarativeModule) QtWebSocketsDeclarativeModule_PTR() *QtWebSocketsDeclarativeModule {
	return ptr
}

func (ptr *QtWebSocketsDeclarativeModule) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtWebSocketsDeclarativeModule) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtWebSocketsDeclarativeModule(ptr QtWebSocketsDeclarativeModule_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebSocketsDeclarativeModule_PTR().Pointer()
	}
	return nil
}

func NewQtWebSocketsDeclarativeModuleFromPointer(ptr unsafe.Pointer) (n *QtWebSocketsDeclarativeModule) {
	n = new(QtWebSocketsDeclarativeModule)
	n.SetPointer(ptr)
	return
}
func (ptr *QtWebSocketsDeclarativeModule) DestroyQtWebSocketsDeclarativeModule() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func init() {
	qt.ItfMap["websockets.QMaskGenerator_ITF"] = QMaskGenerator{}
	qt.FuncMap["websockets.NewQMaskGenerator2"] = NewQMaskGenerator2
	qt.ItfMap["websockets.QWebSocket_ITF"] = QWebSocket{}
	qt.FuncMap["websockets.NewQWebSocket2"] = NewQWebSocket2
	qt.ItfMap["websockets.QWebSocketCorsAuthenticator_ITF"] = QWebSocketCorsAuthenticator{}
	qt.FuncMap["websockets.NewQWebSocketCorsAuthenticator"] = NewQWebSocketCorsAuthenticator
	qt.FuncMap["websockets.NewQWebSocketCorsAuthenticator2"] = NewQWebSocketCorsAuthenticator2
	qt.FuncMap["websockets.NewQWebSocketCorsAuthenticator3"] = NewQWebSocketCorsAuthenticator3
	qt.ItfMap["websockets.QWebSocketProtocol_ITF"] = QWebSocketProtocol{}
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeNormal"] = int64(QWebSocketProtocol__CloseCodeNormal)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeGoingAway"] = int64(QWebSocketProtocol__CloseCodeGoingAway)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeProtocolError"] = int64(QWebSocketProtocol__CloseCodeProtocolError)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeDatatypeNotSupported"] = int64(QWebSocketProtocol__CloseCodeDatatypeNotSupported)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeReserved1004"] = int64(QWebSocketProtocol__CloseCodeReserved1004)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeMissingStatusCode"] = int64(QWebSocketProtocol__CloseCodeMissingStatusCode)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeAbnormalDisconnection"] = int64(QWebSocketProtocol__CloseCodeAbnormalDisconnection)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeWrongDatatype"] = int64(QWebSocketProtocol__CloseCodeWrongDatatype)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodePolicyViolated"] = int64(QWebSocketProtocol__CloseCodePolicyViolated)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeTooMuchData"] = int64(QWebSocketProtocol__CloseCodeTooMuchData)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeMissingExtension"] = int64(QWebSocketProtocol__CloseCodeMissingExtension)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeBadOperation"] = int64(QWebSocketProtocol__CloseCodeBadOperation)
	qt.EnumMap["websockets.QWebSocketProtocol__CloseCodeTlsHandshakeFailed"] = int64(QWebSocketProtocol__CloseCodeTlsHandshakeFailed)
	qt.EnumMap["websockets.QWebSocketProtocol__VersionUnknown"] = int64(QWebSocketProtocol__VersionUnknown)
	qt.EnumMap["websockets.QWebSocketProtocol__Version0"] = int64(QWebSocketProtocol__Version0)
	qt.EnumMap["websockets.QWebSocketProtocol__Version4"] = int64(QWebSocketProtocol__Version4)
	qt.EnumMap["websockets.QWebSocketProtocol__Version5"] = int64(QWebSocketProtocol__Version5)
	qt.EnumMap["websockets.QWebSocketProtocol__Version6"] = int64(QWebSocketProtocol__Version6)
	qt.EnumMap["websockets.QWebSocketProtocol__Version7"] = int64(QWebSocketProtocol__Version7)
	qt.EnumMap["websockets.QWebSocketProtocol__Version8"] = int64(QWebSocketProtocol__Version8)
	qt.EnumMap["websockets.QWebSocketProtocol__Version13"] = int64(QWebSocketProtocol__Version13)
	qt.EnumMap["websockets.QWebSocketProtocol__VersionLatest"] = int64(QWebSocketProtocol__VersionLatest)
	qt.ItfMap["websockets.QWebSocketServer_ITF"] = QWebSocketServer{}
	qt.FuncMap["websockets.NewQWebSocketServer2"] = NewQWebSocketServer2
	qt.EnumMap["websockets.QWebSocketServer__SecureMode"] = int64(QWebSocketServer__SecureMode)
	qt.EnumMap["websockets.QWebSocketServer__NonSecureMode"] = int64(QWebSocketServer__NonSecureMode)
	qt.ItfMap["websockets.QtWebSocketsDeclarativeModule_ITF"] = QtWebSocketsDeclarativeModule{}
}
