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
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebSockets_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	return n
}
func NewQMaskGenerator(parent core.QObject_ITF) *QMaskGenerator {
	var tmpValue = NewQMaskGeneratorFromPointer(C.QMaskGenerator_NewQMaskGenerator(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMaskGenerator_Seed
func callbackQMaskGenerator_Seed(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "seed"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMaskGenerator) ConnectSeed(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "seed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "seed", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "seed", f)
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
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQMaskGenerator_NextMask
func callbackQMaskGenerator_NextMask(ptr unsafe.Pointer) C.uint {
	if signal := qt.GetSignal(ptr, "nextMask"); signal != nil {
		return C.uint(uint32(signal.(func() uint)()))
	}

	return C.uint(uint32(0))
}

func (ptr *QMaskGenerator) ConnectNextMask(f func() uint) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextMask"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nextMask", func() uint {
				signal.(func() uint)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextMask", f)
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

//export callbackQMaskGenerator_DestroyQMaskGenerator
func callbackQMaskGenerator_DestroyQMaskGenerator(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QMaskGenerator"); signal != nil {
		signal.(func())()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DestroyQMaskGeneratorDefault()
	}
}

func (ptr *QMaskGenerator) ConnectDestroyQMaskGenerator(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QMaskGenerator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QMaskGenerator", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QMaskGenerator", f)
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
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMaskGenerator) DestroyQMaskGeneratorDefault() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGeneratorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QMaskGenerator___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QMaskGenerator___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QMaskGenerator) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QMaskGenerator___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QMaskGenerator) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QMaskGenerator) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QMaskGenerator___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QMaskGenerator) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QMaskGenerator___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QMaskGenerator___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QMaskGenerator) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QMaskGenerator___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QMaskGenerator___findChildren_newList(ptr.Pointer()))
}

func (ptr *QMaskGenerator) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QMaskGenerator___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QMaskGenerator___children_newList(ptr.Pointer()))
}

//export callbackQMaskGenerator_Event
func callbackQMaskGenerator_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMaskGenerator) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMaskGenerator_EventFilter
func callbackQMaskGenerator_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMaskGenerator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMaskGenerator_ChildEvent
func callbackQMaskGenerator_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMaskGenerator) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQMaskGenerator_Destroyed
func callbackQMaskGenerator_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQMaskGenerator_DisconnectNotify
func callbackQMaskGenerator_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_ObjectNameChanged
func callbackQMaskGenerator_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQMaskGenerator_TimerEvent
func callbackQMaskGenerator_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMaskGenerator_MetaObject
func callbackQMaskGenerator_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMaskGeneratorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMaskGenerator) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMaskGenerator_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = new(QWebSocket)
	n.SetPointer(ptr)
	return n
}
func NewQWebSocket(origin string, version QWebSocketProtocol__Version, parent core.QObject_ITF) *QWebSocket {
	var originC *C.char
	if origin != "" {
		originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
	}
	var tmpValue = NewQWebSocketFromPointer(C.QWebSocket_NewQWebSocket(C.struct_QtWebSockets_PackedString{data: originC, len: C.longlong(len(origin))}, C.longlong(version), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QWebSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

//export callbackQWebSocket_AboutToClose
func callbackQWebSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "aboutToClose") {
			C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "aboutToClose"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "aboutToClose", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "aboutToClose", f)
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
		signal.(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "binaryFrameReceived") {
			C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "binaryFrameReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "binaryFrameReceived", func(frame *core.QByteArray, isLastFrame bool) {
				signal.(func(*core.QByteArray, bool))(frame, isLastFrame)
				f(frame, isLastFrame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryFrameReceived", f)
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
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
	}

}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "binaryMessageReceived") {
			C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "binaryMessageReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "binaryMessageReceived", func(message *core.QByteArray) {
				signal.(func(*core.QByteArray))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "binaryMessageReceived", f)
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

//export callbackQWebSocket_BytesWritten
func callbackQWebSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "bytesWritten") {
			C.QWebSocket_ConnectBytesWritten(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "bytesWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "bytesWritten", func(bytes int64) {
				signal.(func(int64))(bytes)
				f(bytes)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "bytesWritten", f)
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
		signal.(func(QWebSocketProtocol__CloseCode, string))(QWebSocketProtocol__CloseCode(closeCode), cGoUnpackString(reason))
	} else {
		NewQWebSocketFromPointer(ptr).CloseDefault(QWebSocketProtocol__CloseCode(closeCode), cGoUnpackString(reason))
	}
}

func (ptr *QWebSocket) ConnectClose(f func(closeCode QWebSocketProtocol__CloseCode, reason string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func(closeCode QWebSocketProtocol__CloseCode, reason string) {
				signal.(func(QWebSocketProtocol__CloseCode, string))(closeCode, reason)
				f(closeCode, reason)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
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

//export callbackQWebSocket_Connected
func callbackQWebSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QWebSocket_ConnectConnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", f)
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
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", f)
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

//export callbackQWebSocket_Error2
func callbackQWebSocket_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QWebSocket_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(error network.QAbstractSocket__SocketError) {
				signal.(func(network.QAbstractSocket__SocketError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
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

//export callbackQWebSocket_IgnoreSslErrors
func callbackQWebSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ignoreSslErrors"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QWebSocket) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ignoreSslErrors"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ignoreSslErrors", f)
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
			var tmpList = NewQWebSocketFromPointer(NewQWebSocketFromPointer(nil).__ignoreSslErrors_errors_newList2())
			for _, v := range errors {
				tmpList.__ignoreSslErrors_errors_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQWebSocket_Open2
func callbackQWebSocket_Open2(ptr unsafe.Pointer, request unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open2"); signal != nil {
		signal.(func(*network.QNetworkRequest))(network.NewQNetworkRequestFromPointer(request))
	} else {
		NewQWebSocketFromPointer(ptr).Open2Default(network.NewQNetworkRequestFromPointer(request))
	}
}

func (ptr *QWebSocket) ConnectOpen2(f func(request *network.QNetworkRequest)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open2", func(request *network.QNetworkRequest) {
				signal.(func(*network.QNetworkRequest))(request)
				f(request)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open2", f)
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

//export callbackQWebSocket_Open
func callbackQWebSocket_Open(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQWebSocketFromPointer(ptr).OpenDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QWebSocket) ConnectOpen(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
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

//export callbackQWebSocket_Ping
func callbackQWebSocket_Ping(ptr unsafe.Pointer, payload unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "ping"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(payload))
	} else {
		NewQWebSocketFromPointer(ptr).PingDefault(core.NewQByteArrayFromPointer(payload))
	}
}

func (ptr *QWebSocket) ConnectPing(f func(payload *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "ping"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ping", func(payload *core.QByteArray) {
				signal.(func(*core.QByteArray))(payload)
				f(payload)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ping", f)
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
		signal.(func(uint64, *core.QByteArray))(uint64(elapsedTime), core.NewQByteArrayFromPointer(payload))
	}

}

func (ptr *QWebSocket) ConnectPong(f func(elapsedTime uint64, payload *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pong") {
			C.QWebSocket_ConnectPong(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pong"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pong", func(elapsedTime uint64, payload *core.QByteArray) {
				signal.(func(uint64, *core.QByteArray))(elapsedTime, payload)
				f(elapsedTime, payload)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pong", f)
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
		signal.(func(*network.QSslPreSharedKeyAuthenticator))(network.NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QWebSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", func(authenticator *network.QSslPreSharedKeyAuthenticator) {
				signal.(func(*network.QSslPreSharedKeyAuthenticator))(authenticator)
				f(authenticator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", f)
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

//export callbackQWebSocket_ProxyAuthenticationRequired
func callbackQWebSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "proxyAuthenticationRequired"); signal != nil {
		signal.(func(*network.QNetworkProxy, *network.QAuthenticator))(network.NewQNetworkProxyFromPointer(proxy), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectProxyAuthenticationRequired(f func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "proxyAuthenticationRequired") {
			C.QWebSocket_ConnectProxyAuthenticationRequired(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "proxyAuthenticationRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator) {
				signal.(func(*network.QNetworkProxy, *network.QAuthenticator))(proxy, authenticator)
				f(proxy, authenticator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "proxyAuthenticationRequired", f)
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

//export callbackQWebSocket_ReadChannelFinished
func callbackQWebSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "readChannelFinished") {
			C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "readChannelFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readChannelFinished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readChannelFinished", f)
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

func (ptr *QWebSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
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

//export callbackQWebSocket_SslErrors
func callbackQWebSocket_SslErrors(ptr unsafe.Pointer, errors C.struct_QtWebSockets_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		signal.(func([]*network.QSslError))(func(l C.struct_QtWebSockets_PackedList) []*network.QSslError {
			var out = make([]*network.QSslError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebSocketFromPointer(l.data).__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QWebSocket) ConnectSslErrors(f func(errors []*network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QWebSocket_ConnectSslErrors(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", func(errors []*network.QSslError) {
				signal.(func([]*network.QSslError))(errors)
				f(errors)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", f)
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
			var tmpList = NewQWebSocketFromPointer(NewQWebSocketFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQWebSocket_StateChanged
func callbackQWebSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state network.QAbstractSocket__SocketState) {
				signal.(func(network.QAbstractSocket__SocketState))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
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
		signal.(func(string, bool))(cGoUnpackString(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textFrameReceived") {
			C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textFrameReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textFrameReceived", func(frame string, isLastFrame bool) {
				signal.(func(string, bool))(frame, isLastFrame)
				f(frame, isLastFrame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textFrameReceived", f)
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
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "textMessageReceived") {
			C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "textMessageReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "textMessageReceived", func(message string) {
				signal.(func(string))(message)
				f(message)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "textMessageReceived", f)
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

//export callbackQWebSocket_DestroyQWebSocket
func callbackQWebSocket_DestroyQWebSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).DestroyQWebSocketDefault()
	}
}

func (ptr *QWebSocket) ConnectDestroyQWebSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebSocket"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocket", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocket", f)
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
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSocket) DestroyQWebSocketDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketState(C.QWebSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) LocalAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) PeerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkProxyFromPointer(C.QWebSocket_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) Request() *network.QNetworkRequest {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkRequestFromPointer(C.QWebSocket_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkRequest).DestroyQNetworkRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslConfigurationFromPointer(C.QWebSocket_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) CloseReason() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) ResourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebSocket_RequestUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) CloseCode() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocket_CloseCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) Version() QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__Version(C.QWebSocket_Version(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	if ptr.Pointer() != nil {
		var tmpValue = NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_atList2(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslErrorFromPointer(C.QWebSocket___ignoreSslErrors_errors_atList2(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
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
	return unsafe.Pointer(C.QWebSocket___ignoreSslErrors_errors_newList2(ptr.Pointer()))
}

func (ptr *QWebSocket) __sslErrors_errors_atList(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslErrorFromPointer(C.QWebSocket___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
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
	return unsafe.Pointer(C.QWebSocket___sslErrors_errors_newList(ptr.Pointer()))
}

func (ptr *QWebSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QWebSocket___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebSocket) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocket___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocket) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebSocket___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocket___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocket___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocket___children_newList(ptr.Pointer()))
}

//export callbackQWebSocket_Event
func callbackQWebSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocket_EventFilter
func callbackQWebSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocket_ChildEvent
func callbackQWebSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebSocket_Destroyed
func callbackQWebSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebSocket_DisconnectNotify
func callbackQWebSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_ObjectNameChanged
func callbackQWebSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebSocket_TimerEvent
func callbackQWebSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocket_MetaObject
func callbackQWebSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return n
}
func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	var originC *C.char
	if origin != "" {
		originC = C.CString(origin)
		defer C.free(unsafe.Pointer(originC))
	}
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.struct_QtWebSockets_PackedString{data: originC, len: C.longlong(len(origin))}))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
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
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer()) != 0
	}
	return false
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

func NewQWebSocketProtocolFromPointer(ptr unsafe.Pointer) *QWebSocketProtocol {
	var n = new(QWebSocketProtocol)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebSocketProtocol) DestroyQWebSocketProtocol() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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

func NewQWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = new(QWebSocketServer)
	n.SetPointer(ptr)
	return n
}

//go:generate stringer -type=QWebSocketServer__SslMode
//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int64

const (
	QWebSocketServer__SecureMode    QWebSocketServer__SslMode = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode QWebSocketServer__SslMode = QWebSocketServer__SslMode(1)
)

//export callbackQWebSocketServer_NextPendingConnection
func callbackQWebSocketServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "nextPendingConnection"); signal != nil {
		return PointerFromQWebSocket(signal.(func() *QWebSocket)())
	}

	return PointerFromQWebSocket(NewQWebSocketServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QWebSocketServer) ConnectNextPendingConnection(f func() *QWebSocket) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nextPendingConnection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", func() *QWebSocket {
				signal.(func() *QWebSocket)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nextPendingConnection", f)
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
		var tmpValue = NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) NextPendingConnectionDefault() *QWebSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQWebSocketServer(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {
	var serverNameC *C.char
	if serverName != "" {
		serverNameC = C.CString(serverName)
		defer C.free(unsafe.Pointer(serverNameC))
	}
	var tmpValue = NewQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer(C.struct_QtWebSockets_PackedString{data: serverNameC, len: C.longlong(len(serverName))}, C.longlong(secureMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebSocketServer) Listen(address network.QHostAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_Listen(ptr.Pointer(), network.PointerFromQHostAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) SetSocketDescriptor(socketDescriptor int) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_SetSocketDescriptor(ptr.Pointer(), C.int(int32(socketDescriptor))) != 0
	}
	return false
}

//export callbackQWebSocketServer_AcceptError
func callbackQWebSocketServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {
	if signal := qt.GetSignal(ptr, "acceptError"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "acceptError") {
			C.QWebSocketServer_ConnectAcceptError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "acceptError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "acceptError", func(socketError network.QAbstractSocket__SocketError) {
				signal.(func(network.QAbstractSocket__SocketError))(socketError)
				f(socketError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "acceptError", f)
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
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "closed") {
			C.QWebSocketServer_ConnectClosed(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "closed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "closed", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closed", f)
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

//export callbackQWebSocketServer_NewConnection
func callbackQWebSocketServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QWebSocketServer_ConnectNewConnection(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", f)
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

//export callbackQWebSocketServer_OriginAuthenticationRequired
func callbackQWebSocketServer_OriginAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "originAuthenticationRequired"); signal != nil {
		signal.(func(*QWebSocketCorsAuthenticator))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "originAuthenticationRequired") {
			C.QWebSocketServer_ConnectOriginAuthenticationRequired(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "originAuthenticationRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "originAuthenticationRequired", func(authenticator *QWebSocketCorsAuthenticator) {
				signal.(func(*QWebSocketCorsAuthenticator))(authenticator)
				f(authenticator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "originAuthenticationRequired", f)
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
		signal.(func(*network.QSslError))(network.NewQSslErrorFromPointer(error))
	}

}

func (ptr *QWebSocketServer) ConnectPeerVerifyError(f func(error *network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peerVerifyError") {
			C.QWebSocketServer_ConnectPeerVerifyError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peerVerifyError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", func(error *network.QSslError) {
				signal.(func(*network.QSslError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peerVerifyError", f)
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
		signal.(func(*network.QSslPreSharedKeyAuthenticator))(network.NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired") {
			C.QWebSocketServer_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", func(authenticator *network.QSslPreSharedKeyAuthenticator) {
				signal.(func(*network.QSslPreSharedKeyAuthenticator))(authenticator)
				f(authenticator)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preSharedKeyAuthenticationRequired", f)
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

func (ptr *QWebSocketServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_ServerError
func callbackQWebSocketServer_ServerError(ptr unsafe.Pointer, closeCode C.longlong) {
	if signal := qt.GetSignal(ptr, "serverError"); signal != nil {
		signal.(func(QWebSocketProtocol__CloseCode))(QWebSocketProtocol__CloseCode(closeCode))
	}

}

func (ptr *QWebSocketServer) ConnectServerError(f func(closeCode QWebSocketProtocol__CloseCode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serverError") {
			C.QWebSocketServer_ConnectServerError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serverError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serverError", func(closeCode QWebSocketProtocol__CloseCode) {
				signal.(func(QWebSocketProtocol__CloseCode))(closeCode)
				f(closeCode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serverError", f)
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

//export callbackQWebSocketServer_SslErrors
func callbackQWebSocketServer_SslErrors(ptr unsafe.Pointer, errors C.struct_QtWebSockets_PackedList) {
	if signal := qt.GetSignal(ptr, "sslErrors"); signal != nil {
		signal.(func([]*network.QSslError))(func(l C.struct_QtWebSockets_PackedList) []*network.QSslError {
			var out = make([]*network.QSslError, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebSocketServerFromPointer(l.data).__sslErrors_errors_atList(i)
			}
			return out
		}(errors))
	}

}

func (ptr *QWebSocketServer) ConnectSslErrors(f func(errors []*network.QSslError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sslErrors") {
			C.QWebSocketServer_ConnectSslErrors(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sslErrors"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", func(errors []*network.QSslError) {
				signal.(func([]*network.QSslError))(errors)
				f(errors)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sslErrors", f)
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
			var tmpList = NewQWebSocketServerFromPointer(NewQWebSocketServerFromPointer(nil).__sslErrors_errors_newList())
			for _, v := range errors {
				tmpList.__sslErrors_errors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQWebSocketServer_DestroyQWebSocketServer
func callbackQWebSocketServer_DestroyQWebSocketServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWebSocketServer"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketServerFromPointer(ptr).DestroyQWebSocketServerDefault()
	}
}

func (ptr *QWebSocketServer) ConnectDestroyQWebSocketServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWebSocketServer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocketServer", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWebSocketServer", f)
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
		C.QWebSocketServer_DestroyQWebSocketServer(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSocketServer) DestroyQWebSocketServerDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DestroyQWebSocketServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWebSocketServer) ServerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocketServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) SupportedVersions() []QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWebSockets_PackedList) []QWebSocketProtocol__Version {
			var out = make([]QWebSocketProtocol__Version, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQWebSocketServerFromPointer(l.data).__supportedVersions_atList(i)
			}
			return out
		}(C.QWebSocketServer_SupportedVersions(ptr.Pointer()))
	}
	return make([]QWebSocketProtocol__Version, 0)
}

func (ptr *QWebSocketServer) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkProxyFromPointer(C.QWebSocketServer_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslConfigurationFromPointer(C.QWebSocketServer_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebSocketServer_ServerUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) Error() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocketServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSocketServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSocketServer_SocketDescriptor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocketServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) HandleConnection(socket network.QTcpSocket_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_HandleConnection(ptr.Pointer(), network.PointerFromQTcpSocket(socket))
	}
}

func (ptr *QWebSocketServer) __sslErrors_errors_atList(i int) *network.QSslError {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslErrorFromPointer(C.QWebSocketServer___sslErrors_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*network.QSslError).DestroyQSslError)
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
	return unsafe.Pointer(C.QWebSocketServer___sslErrors_errors_newList(ptr.Pointer()))
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
	return unsafe.Pointer(C.QWebSocketServer___supportedVersions_newList(ptr.Pointer()))
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QWebSocketServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return unsafe.Pointer(C.QWebSocketServer___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QWebSocketServer) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocketServer___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWebSocketServer) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QWebSocketServer___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QWebSocketServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocketServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocketServer___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QWebSocketServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocketServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocketServer___findChildren_newList(ptr.Pointer()))
}

func (ptr *QWebSocketServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QWebSocketServer___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QWebSocketServer___children_newList(ptr.Pointer()))
}

//export callbackQWebSocketServer_Event
func callbackQWebSocketServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocketServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocketServer_EventFilter
func callbackQWebSocketServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocketServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocketServer_ChildEvent
func callbackQWebSocketServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
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
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
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
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
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
		signal.(func())()
	} else {
		NewQWebSocketServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocketServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWebSocketServer_Destroyed
func callbackQWebSocketServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWebSocketServer_DisconnectNotify
func callbackQWebSocketServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_ObjectNameChanged
func callbackQWebSocketServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQWebSocketServer_TimerEvent
func callbackQWebSocketServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocketServer_MetaObject
func callbackQWebSocketServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocketServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
