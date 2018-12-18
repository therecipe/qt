// +build !minimal

package qml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtQml_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtQml_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func (ptr *QJSEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QJSEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) (n *QJSEngine) {
	n = new(QJSEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QJSEngine__Extension
//QJSEngine::Extension
type QJSEngine__Extension int64

const (
	QJSEngine__TranslationExtension       QJSEngine__Extension = QJSEngine__Extension(0x1)
	QJSEngine__ConsoleExtension           QJSEngine__Extension = QJSEngine__Extension(0x2)
	QJSEngine__GarbageCollectionExtension QJSEngine__Extension = QJSEngine__Extension(0x4)
	QJSEngine__AllExtensions              QJSEngine__Extension = QJSEngine__Extension(0xffffffff)
)

func NewQJSEngine() *QJSEngine {
	tmpValue := NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	tmpValue := NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QJSEngine) Evaluate(program string, fileName string, lineNumber int) *QJSValue {
	if ptr.Pointer() != nil {
		var programC *C.char
		if program != "" {
			programC = C.CString(program)
			defer C.free(unsafe.Pointer(programC))
		}
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_Evaluate(ptr.Pointer(), C.struct_QtQml_PackedString{data: programC, len: C.longlong(len(program))}, C.struct_QtQml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.int(int32(lineNumber))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) ImportModule(fileName string) *QJSValue {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_ImportModule(ptr.Pointer(), C.struct_QtQml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewArray(length uint) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewArray(ptr.Pointer(), C.uint(uint32(length))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewErrorObject(errorType QJSValue__ErrorType, message string) *QJSValue {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewErrorObject(ptr.Pointer(), C.longlong(errorType), C.struct_QtQml_PackedString{data: messageC, len: C.longlong(len(message))}))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewObject() *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewQMetaObject(metaObject core.QMetaObject_ITF) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewQMetaObject(ptr.Pointer(), core.PointerFromQMetaObject(metaObject)))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func QJSEngine_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QJSEngine_QJSEngine_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QJSEngine) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QJSEngine_QJSEngine_Tr(sC, cC, C.int(int32(n))))
}

func QJSEngine_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QJSEngine_QJSEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QJSEngine) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QJSEngine_QJSEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QJSEngine) CollectGarbage() {
	if ptr.Pointer() != nil {
		C.QJSEngine_CollectGarbage(ptr.Pointer())
	}
}

func (ptr *QJSEngine) InstallExtensions(extensions QJSEngine__Extension, object QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_InstallExtensions(ptr.Pointer(), C.longlong(extensions), PointerFromQJSValue(object))
	}
}

func (ptr *QJSEngine) ThrowError2(errorType QJSValue__ErrorType, message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QJSEngine_ThrowError2(ptr.Pointer(), C.longlong(errorType), C.struct_QtQml_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

func (ptr *QJSEngine) ThrowError(message string) {
	if ptr.Pointer() != nil {
		var messageC *C.char
		if message != "" {
			messageC = C.CString(message)
			defer C.free(unsafe.Pointer(messageC))
		}
		C.QJSEngine_ThrowError(ptr.Pointer(), C.struct_QtQml_PackedString{data: messageC, len: C.longlong(len(message))})
	}
}

//export callbackQJSEngine_DestroyQJSEngine
func callbackQJSEngine_DestroyQJSEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QJSEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQJSEngineFromPointer(ptr).DestroyQJSEngineDefault()
	}
}

func (ptr *QJSEngine) ConnectDestroyQJSEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QJSEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QJSEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QJSEngine", f)
		}
	}
}

func (ptr *QJSEngine) DisconnectDestroyQJSEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QJSEngine")
	}
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSEngine) DestroyQJSEngineDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSEngine) GlobalObject() *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_GlobalObject(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

//export callbackQJSEngine_MetaObject
func callbackQJSEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQJSEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QJSEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QJSEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QJSEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QJSEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QJSEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList2() unsafe.Pointer {
	return C.QJSEngine___findChildren_newList2(ptr.Pointer())
}

func (ptr *QJSEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QJSEngine___findChildren_newList3(ptr.Pointer())
}

func (ptr *QJSEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList() unsafe.Pointer {
	return C.QJSEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QJSEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __children_newList() unsafe.Pointer {
	return C.QJSEngine___children_newList(ptr.Pointer())
}

//export callbackQJSEngine_Event
func callbackQJSEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QJSEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQJSEngine_EventFilter
func callbackQJSEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QJSEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQJSEngine_ChildEvent
func callbackQJSEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQJSEngine_ConnectNotify
func callbackQJSEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_CustomEvent
func callbackQJSEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQJSEngine_DeleteLater
func callbackQJSEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQJSEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QJSEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQJSEngine_Destroyed
func callbackQJSEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQJSEngine_DisconnectNotify
func callbackQJSEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_ObjectNameChanged
func callbackQJSEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQJSEngine_TimerEvent
func callbackQJSEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

func (ptr *QJSValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJSValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJSValue(ptr QJSValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValue_PTR().Pointer()
	}
	return nil
}

func NewQJSValueFromPointer(ptr unsafe.Pointer) (n *QJSValue) {
	n = new(QJSValue)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QJSValue__ErrorType
//QJSValue::ErrorType
type QJSValue__ErrorType int64

const (
	QJSValue__NoError        QJSValue__ErrorType = QJSValue__ErrorType(0)
	QJSValue__GenericError   QJSValue__ErrorType = QJSValue__ErrorType(1)
	QJSValue__EvalError      QJSValue__ErrorType = QJSValue__ErrorType(2)
	QJSValue__RangeError     QJSValue__ErrorType = QJSValue__ErrorType(3)
	QJSValue__ReferenceError QJSValue__ErrorType = QJSValue__ErrorType(4)
	QJSValue__SyntaxError    QJSValue__ErrorType = QJSValue__ErrorType(5)
	QJSValue__TypeError      QJSValue__ErrorType = QJSValue__ErrorType(6)
	QJSValue__URIError       QJSValue__ErrorType = QJSValue__ErrorType(7)
)

//go:generate stringer -type=QJSValue__SpecialValue
//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      QJSValue__SpecialValue = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue QJSValue__SpecialValue = QJSValue__SpecialValue(1)
)

func (ptr *QJSValue) Call(args []*QJSValue) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Call(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQJSValueFromPointer(NewQJSValueFromPointer(nil).__call_args_newList())
			for _, v := range args {
				tmpList.__call_args_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) CallAsConstructor(args []*QJSValue) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_CallAsConstructor(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQJSValueFromPointer(NewQJSValueFromPointer(nil).__callAsConstructor_args_newList())
			for _, v := range args {
				tmpList.__callAsConstructor_args_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) CallWithInstance(instance QJSValue_ITF, args []*QJSValue) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_CallWithInstance(ptr.Pointer(), PointerFromQJSValue(instance), func() unsafe.Pointer {
			tmpList := NewQJSValueFromPointer(NewQJSValueFromPointer(nil).__callWithInstance_args_newList())
			for _, v := range args {
				tmpList.__callWithInstance_args_setList(v)
			}
			return tmpList.Pointer()
		}()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue(C.longlong(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue4(value bool) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.char(int8(qt.GoBoolToInt(value)))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue8(value string) *QJSValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue8(C.struct_QtQml_PackedString{data: valueC, len: C.longlong(len(value))}))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue10(value string) *QJSValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue10(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue7(value float64) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue7(C.double(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue5(value int) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(int32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue6(value uint) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue6(C.uint(uint32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func (ptr *QJSValue) DeleteProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QJSValue_DeleteProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QJSValue) SetProperty(name string, value QJSValue_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QJSValue_SetProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetProperty2(arrayIndex uint, value QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetProperty2(ptr.Pointer(), C.uint(uint32(arrayIndex)), PointerFromQJSValue(value))
	}
}

func (ptr *QJSValue) SetPrototype(prototype QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue_SetPrototype(ptr.Pointer(), PointerFromQJSValue(prototype))
	}
}

func (ptr *QJSValue) DestroyQJSValue() {
	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSValue) ToDateTime() *core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDateTimeFromPointer(C.QJSValue_ToDateTime(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Property(name string) *QJSValue {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Property(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Property2(arrayIndex uint) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Prototype() *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Prototype(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ErrorType() QJSValue__ErrorType {
	if ptr.Pointer() != nil {
		return QJSValue__ErrorType(C.QJSValue_ErrorType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJSValue) ToQObject() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSValue_ToQObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QJSValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Equals(other QJSValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_Equals(ptr.Pointer(), PointerFromQJSValue(other))) != 0
	}
	return false
}

func (ptr *QJSValue) HasOwnProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QJSValue_HasOwnProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QJSValue) HasProperty(name string) bool {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int8(C.QJSValue_HasProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))})) != 0
	}
	return false
}

func (ptr *QJSValue) IsArray() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsArray(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsBool() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsBool(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsCallable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsCallable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsDate() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsDate(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsNumber() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsNumber(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsQMetaObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsQMetaObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsQObject() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsQObject(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsRegExp() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsRegExp(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsString() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsString(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsUndefined() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsUndefined(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) IsVariant() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_IsVariant(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) StrictlyEquals(other QJSValue_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_StrictlyEquals(ptr.Pointer(), PointerFromQJSValue(other))) != 0
	}
	return false
}

func (ptr *QJSValue) ToBool() bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSValue_ToBool(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJSValue) ToQMetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QJSValue_ToQMetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJSValue) ToNumber() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QJSValue_ToNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QJSValue) ToInt() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJSValue_ToInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToUInt() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QJSValue_ToUInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) __call_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___call_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __call_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___call_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __call_args_newList() unsafe.Pointer {
	return C.QJSValue___call_args_newList(ptr.Pointer())
}

func (ptr *QJSValue) __callAsConstructor_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___callAsConstructor_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __callAsConstructor_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___callAsConstructor_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __callAsConstructor_args_newList() unsafe.Pointer {
	return C.QJSValue___callAsConstructor_args_newList(ptr.Pointer())
}

func (ptr *QJSValue) __callWithInstance_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___callWithInstance_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __callWithInstance_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___callWithInstance_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __callWithInstance_args_newList() unsafe.Pointer {
	return C.QJSValue___callWithInstance_args_newList(ptr.Pointer())
}

type QJSValueIterator struct {
	ptr unsafe.Pointer
}

type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return ptr
}

func (ptr *QJSValueIterator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJSValueIterator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJSValueIterator(ptr QJSValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQJSValueIteratorFromPointer(ptr unsafe.Pointer) (n *QJSValueIterator) {
	n = new(QJSValueIterator)
	n.SetPointer(ptr)
	return
}

type QQmlAbstractUrlInterceptor struct {
	ptr unsafe.Pointer
}

type QQmlAbstractUrlInterceptor_ITF interface {
	QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor
}

func (ptr *QQmlAbstractUrlInterceptor) QQmlAbstractUrlInterceptor_PTR() *QQmlAbstractUrlInterceptor {
	return ptr
}

func (ptr *QQmlAbstractUrlInterceptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlAbstractUrlInterceptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlAbstractUrlInterceptor(ptr QQmlAbstractUrlInterceptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlAbstractUrlInterceptor_PTR().Pointer()
	}
	return nil
}

func NewQQmlAbstractUrlInterceptorFromPointer(ptr unsafe.Pointer) (n *QQmlAbstractUrlInterceptor) {
	n = new(QQmlAbstractUrlInterceptor)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQmlAbstractUrlInterceptor__DataType
//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

func NewQQmlAbstractUrlInterceptor() *QQmlAbstractUrlInterceptor {
	return NewQQmlAbstractUrlInterceptorFromPointer(C.QQmlAbstractUrlInterceptor_NewQQmlAbstractUrlInterceptor())
}

//export callbackQQmlAbstractUrlInterceptor_Intercept
func callbackQQmlAbstractUrlInterceptor_Intercept(ptr unsafe.Pointer, url unsafe.Pointer, ty C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "intercept"); signal != nil {
		return core.PointerFromQUrl(signal.(func(*core.QUrl, QQmlAbstractUrlInterceptor__DataType) *core.QUrl)(core.NewQUrlFromPointer(url), QQmlAbstractUrlInterceptor__DataType(ty)))
	}

	return core.PointerFromQUrl(core.NewQUrl())
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectIntercept(f func(url *core.QUrl, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "intercept"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "intercept", func(url *core.QUrl, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {
				signal.(func(*core.QUrl, QQmlAbstractUrlInterceptor__DataType) *core.QUrl)(url, ty)
				return f(url, ty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "intercept", f)
		}
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectIntercept() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "intercept")
	}
}

func (ptr *QQmlAbstractUrlInterceptor) Intercept(url core.QUrl_ITF, ty QQmlAbstractUrlInterceptor__DataType) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlAbstractUrlInterceptor_Intercept(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(ty)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor
func callbackQQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlAbstractUrlInterceptor"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlAbstractUrlInterceptorFromPointer(ptr).DestroyQQmlAbstractUrlInterceptorDefault()
	}
}

func (ptr *QQmlAbstractUrlInterceptor) ConnectDestroyQQmlAbstractUrlInterceptor(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlAbstractUrlInterceptor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlAbstractUrlInterceptor", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlAbstractUrlInterceptor", f)
		}
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DisconnectDestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlAbstractUrlInterceptor")
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptor() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlAbstractUrlInterceptor) DestroyQQmlAbstractUrlInterceptorDefault() {
	if ptr.Pointer() != nil {
		C.QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return ptr
}

func (ptr *QQmlApplicationEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlApplicationEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQmlEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) (n *QQmlApplicationEngine) {
	n = new(QQmlApplicationEngine)
	n.SetPointer(ptr)
	return
}
func NewQQmlApplicationEngine(parent core.QObject_ITF) *QQmlApplicationEngine {
	tmpValue := NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {
	var filePathC *C.char
	if filePath != "" {
		filePathC = C.CString(filePath)
		defer C.free(unsafe.Pointer(filePathC))
	}
	tmpValue := NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(C.struct_QtQml_PackedString{data: filePathC, len: C.longlong(len(filePath))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {
	tmpValue := NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQQmlApplicationEngine_Load2
func callbackQQmlApplicationEngine_Load2(ptr unsafe.Pointer, filePath C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "load2"); signal != nil {
		signal.(func(string))(cGoUnpackString(filePath))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).Load2Default(cGoUnpackString(filePath))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoad2(f func(filePath string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load2", func(filePath string) {
				signal.(func(string))(filePath)
				f(filePath)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load2", f)
		}
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load2")
	}
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		C.QQmlApplicationEngine_Load2(ptr.Pointer(), C.struct_QtQml_PackedString{data: filePathC, len: C.longlong(len(filePath))})
	}
}

func (ptr *QQmlApplicationEngine) Load2Default(filePath string) {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		C.QQmlApplicationEngine_Load2Default(ptr.Pointer(), C.struct_QtQml_PackedString{data: filePathC, len: C.longlong(len(filePath))})
	}
}

//export callbackQQmlApplicationEngine_Load
func callbackQQmlApplicationEngine_Load(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).LoadDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoad(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", f)
		}
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_LoadData
func callbackQQmlApplicationEngine_LoadData(ptr unsafe.Pointer, data unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadData"); signal != nil {
		signal.(func(*core.QByteArray, *core.QUrl))(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).LoadDataDefault(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectLoadData(f func(data *core.QByteArray, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadData", func(data *core.QByteArray, url *core.QUrl) {
				signal.(func(*core.QByteArray, *core.QUrl))(data, url)
				f(data, url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadData", f)
		}
	}
}

func (ptr *QQmlApplicationEngine) DisconnectLoadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadData")
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadDataDefault(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_ObjectCreated
func callbackQQmlApplicationEngine_ObjectCreated(ptr unsafe.Pointer, object unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "objectCreated"); signal != nil {
		signal.(func(*core.QObject, *core.QUrl))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "objectCreated") {
			C.QQmlApplicationEngine_ConnectObjectCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "objectCreated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "objectCreated", func(object *core.QObject, url *core.QUrl) {
				signal.(func(*core.QObject, *core.QUrl))(object, url)
				f(object, url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "objectCreated", f)
		}
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "objectCreated")
	}
}

func (ptr *QQmlApplicationEngine) ObjectCreated(object core.QObject_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ObjectCreated(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlApplicationEngine_DestroyQQmlApplicationEngine
func callbackQQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlApplicationEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlApplicationEngineFromPointer(ptr).DestroyQQmlApplicationEngineDefault()
	}
}

func (ptr *QQmlApplicationEngine) ConnectDestroyQQmlApplicationEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlApplicationEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlApplicationEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlApplicationEngine", f)
		}
	}
}

func (ptr *QQmlApplicationEngine) DisconnectDestroyQQmlApplicationEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlApplicationEngine")
	}
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngineDefault() {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlApplicationEngine) RootObjects() []*core.QObject {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*core.QObject {
			out := make([]*core.QObject, int(l.len))
			tmpList := NewQQmlApplicationEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__rootObjects_atList(i)
			}
			return out
		}(C.QQmlApplicationEngine_RootObjects(ptr.Pointer()))
	}
	return make([]*core.QObject, 0)
}

func (ptr *QQmlApplicationEngine) __rootObjects_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlApplicationEngine___rootObjects_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __rootObjects_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___rootObjects_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __rootObjects_newList2() unsafe.Pointer {
	return C.QQmlApplicationEngine___rootObjects_newList2(ptr.Pointer())
}

func (ptr *QQmlApplicationEngine) __rootObjects_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlApplicationEngine___rootObjects_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlApplicationEngine) __rootObjects_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine___rootObjects_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlApplicationEngine) __rootObjects_newList() unsafe.Pointer {
	return C.QQmlApplicationEngine___rootObjects_newList(ptr.Pointer())
}

type QQmlComponent struct {
	core.QObject
}

type QQmlComponent_ITF interface {
	core.QObject_ITF
	QQmlComponent_PTR() *QQmlComponent
}

func (ptr *QQmlComponent) QQmlComponent_PTR() *QQmlComponent {
	return ptr
}

func (ptr *QQmlComponent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlComponent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlComponent(ptr QQmlComponent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlComponent_PTR().Pointer()
	}
	return nil
}

func NewQQmlComponentFromPointer(ptr unsafe.Pointer) (n *QQmlComponent) {
	n = new(QQmlComponent)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQmlComponent__CompilationMode
//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(1)
)

//go:generate stringer -type=QQmlComponent__Status
//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    QQmlComponent__Status = QQmlComponent__Status(0)
	QQmlComponent__Ready   QQmlComponent__Status = QQmlComponent__Status(1)
	QQmlComponent__Loading QQmlComponent__Status = QQmlComponent__Status(2)
	QQmlComponent__Error   QQmlComponent__Status = QQmlComponent__Status(3)
)

//export callbackQQmlComponent_BeginCreate
func callbackQQmlComponent_BeginCreate(ptr unsafe.Pointer, publicContext unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "beginCreate"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(publicContext)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).BeginCreateDefault(NewQQmlContextFromPointer(publicContext)))
}

func (ptr *QQmlComponent) ConnectBeginCreate(f func(publicContext *QQmlContext) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "beginCreate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "beginCreate", func(publicContext *QQmlContext) *core.QObject {
				signal.(func(*QQmlContext) *core.QObject)(publicContext)
				return f(publicContext)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "beginCreate", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectBeginCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "beginCreate")
	}
}

func (ptr *QQmlComponent) BeginCreate(publicContext QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent_BeginCreate(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) BeginCreateDefault(publicContext QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent_BeginCreateDefault(ptr.Pointer(), PointerFromQQmlContext(publicContext)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlComponent_Create
func callbackQQmlComponent_Create(ptr unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "create"); signal != nil {
		return core.PointerFromQObject(signal.(func(*QQmlContext) *core.QObject)(NewQQmlContextFromPointer(context)))
	}

	return core.PointerFromQObject(NewQQmlComponentFromPointer(ptr).CreateDefault(NewQQmlContextFromPointer(context)))
}

func (ptr *QQmlComponent) ConnectCreate(f func(context *QQmlContext) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "create"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "create", func(context *QQmlContext) *core.QObject {
				signal.(func(*QQmlContext) *core.QObject)(context)
				return f(context)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "create", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "create")
	}
}

func (ptr *QQmlComponent) Create(context QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent_Create(ptr.Pointer(), PointerFromQQmlContext(context)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) CreateDefault(context QQmlContext_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent_CreateDefault(ptr.Pointer(), PointerFromQQmlContext(context)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func NewQQmlComponent(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlComponent {
	tmpValue := NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent3(engine QQmlEngine_ITF, fileName string, parent core.QObject_ITF) *QQmlComponent {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent3(PointerFromQQmlEngine(engine), C.struct_QtQml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent4(engine QQmlEngine_ITF, fileName string, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent4(PointerFromQQmlEngine(engine), C.struct_QtQml_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent5(engine QQmlEngine_ITF, url core.QUrl_ITF, parent core.QObject_ITF) *QQmlComponent {
	tmpValue := NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent5(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlComponent6(engine QQmlEngine_ITF, url core.QUrl_ITF, mode QQmlComponent__CompilationMode, parent core.QObject_ITF) *QQmlComponent {
	tmpValue := NewQQmlComponentFromPointer(C.QQmlComponent_NewQQmlComponent6(PointerFromQQmlEngine(engine), core.PointerFromQUrl(url), C.longlong(mode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlComponent_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlComponent_QQmlComponent_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlComponent) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlComponent_QQmlComponent_Tr(sC, cC, C.int(int32(n))))
}

func QQmlComponent_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlComponent_QQmlComponent_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlComponent) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlComponent_QQmlComponent_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQmlComponent_CompleteCreate
func callbackQQmlComponent_CompleteCreate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "completeCreate"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).CompleteCreateDefault()
	}
}

func (ptr *QQmlComponent) ConnectCompleteCreate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "completeCreate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "completeCreate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "completeCreate", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectCompleteCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "completeCreate")
	}
}

func (ptr *QQmlComponent) CompleteCreate() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreate(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) CompleteCreateDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CompleteCreateDefault(ptr.Pointer())
	}
}

func (ptr *QQmlComponent) Create2(incubator QQmlIncubator_ITF, context QQmlContext_ITF, forContext QQmlContext_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_Create2(ptr.Pointer(), PointerFromQQmlIncubator(incubator), PointerFromQQmlContext(context), PointerFromQQmlContext(forContext))
	}
}

//export callbackQQmlComponent_LoadUrl
func callbackQQmlComponent_LoadUrl(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadUrl"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	} else {
		NewQQmlComponentFromPointer(ptr).LoadUrlDefault(core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlComponent) ConnectLoadUrl(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadUrl"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadUrl", func(url *core.QUrl) {
				signal.(func(*core.QUrl))(url)
				f(url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadUrl", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadUrl")
	}
}

func (ptr *QQmlComponent) LoadUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) LoadUrlDefault(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrlDefault(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_LoadUrl2
func callbackQQmlComponent_LoadUrl2(ptr unsafe.Pointer, url unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "loadUrl2"); signal != nil {
		signal.(func(*core.QUrl, QQmlComponent__CompilationMode))(core.NewQUrlFromPointer(url), QQmlComponent__CompilationMode(mode))
	} else {
		NewQQmlComponentFromPointer(ptr).LoadUrl2Default(core.NewQUrlFromPointer(url), QQmlComponent__CompilationMode(mode))
	}
}

func (ptr *QQmlComponent) ConnectLoadUrl2(f func(url *core.QUrl, mode QQmlComponent__CompilationMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadUrl2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "loadUrl2", func(url *core.QUrl, mode QQmlComponent__CompilationMode) {
				signal.(func(*core.QUrl, QQmlComponent__CompilationMode))(url, mode)
				f(url, mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadUrl2", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectLoadUrl2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadUrl2")
	}
}

func (ptr *QQmlComponent) LoadUrl2(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(mode))
	}
}

func (ptr *QQmlComponent) LoadUrl2Default(url core.QUrl_ITF, mode QQmlComponent__CompilationMode) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_LoadUrl2Default(ptr.Pointer(), core.PointerFromQUrl(url), C.longlong(mode))
	}
}

//export callbackQQmlComponent_ProgressChanged
func callbackQQmlComponent_ProgressChanged(ptr unsafe.Pointer, progress C.double) {
	if signal := qt.GetSignal(ptr, "progressChanged"); signal != nil {
		signal.(func(float64))(float64(progress))
	}

}

func (ptr *QQmlComponent) ConnectProgressChanged(f func(progress float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "progressChanged") {
			C.QQmlComponent_ConnectProgressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "progressChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "progressChanged", func(progress float64) {
				signal.(func(float64))(progress)
				f(progress)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "progressChanged", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectProgressChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectProgressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "progressChanged")
	}
}

func (ptr *QQmlComponent) ProgressChanged(progress float64) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ProgressChanged(ptr.Pointer(), C.double(progress))
	}
}

//export callbackQQmlComponent_SetData
func callbackQQmlComponent_SetData(ptr unsafe.Pointer, data unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		signal.(func(*core.QByteArray, *core.QUrl))(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	} else {
		NewQQmlComponentFromPointer(ptr).SetDataDefault(core.NewQByteArrayFromPointer(data), core.NewQUrlFromPointer(url))
	}
}

func (ptr *QQmlComponent) ConnectSetData(f func(data *core.QByteArray, url *core.QUrl)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData", func(data *core.QByteArray, url *core.QUrl) {
				signal.(func(*core.QByteArray, *core.QUrl))(data, url)
				f(data, url)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *QQmlComponent) SetData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlComponent) SetDataDefault(data core.QByteArray_ITF, url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_SetDataDefault(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

//export callbackQQmlComponent_StatusChanged
func callbackQQmlComponent_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		signal.(func(QQmlComponent__Status))(QQmlComponent__Status(status))
	}

}

func (ptr *QQmlComponent) ConnectStatusChanged(f func(status QQmlComponent__Status)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.QQmlComponent_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", func(status QQmlComponent__Status) {
				signal.(func(QQmlComponent__Status))(status)
				f(status)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *QQmlComponent) StatusChanged(status QQmlComponent__Status) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

//export callbackQQmlComponent_DestroyQQmlComponent
func callbackQQmlComponent_DestroyQQmlComponent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlComponent"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).DestroyQQmlComponentDefault()
	}
}

func (ptr *QQmlComponent) ConnectDestroyQQmlComponent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlComponent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlComponent", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlComponent", f)
		}
	}
}

func (ptr *QQmlComponent) DisconnectDestroyQQmlComponent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlComponent")
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponent() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlComponent) DestroyQQmlComponentDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DestroyQQmlComponentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlComponent) Errors() []*QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*QQmlError {
			out := make([]*QQmlError, int(l.len))
			tmpList := NewQQmlComponentFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__errors_atList(i)
			}
			return out
		}(C.QQmlComponent_Errors(ptr.Pointer()))
	}
	return make([]*QQmlError, 0)
}

func (ptr *QQmlComponent) Status() QQmlComponent__Status {
	if ptr.Pointer() != nil {
		return QQmlComponent__Status(C.QQmlComponent_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) CreationContext() *QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlContextFromPointer(C.QQmlComponent_CreationContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlEngineFromPointer(C.QQmlComponent_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlComponent_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) IsError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_IsError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsLoading() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_IsLoading(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlComponent) IsReady() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_IsReady(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQmlComponent_MetaObject
func callbackQQmlComponent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlComponentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlComponent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlComponent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlComponent) Progress() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QQmlComponent_Progress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlComponent) __errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlComponent___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlComponent) __errors_newList() unsafe.Pointer {
	return C.QQmlComponent___errors_newList(ptr.Pointer())
}

func (ptr *QQmlComponent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlComponent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlComponent) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlComponent___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlComponent) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlComponent___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlComponent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlComponent___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlComponent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __findChildren_newList() unsafe.Pointer {
	return C.QQmlComponent___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlComponent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlComponent___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlComponent) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlComponent) __children_newList() unsafe.Pointer {
	return C.QQmlComponent___children_newList(ptr.Pointer())
}

//export callbackQQmlComponent_Event
func callbackQQmlComponent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlComponent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlComponent_EventFilter
func callbackQQmlComponent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlComponentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlComponent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlComponent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlComponent_ChildEvent
func callbackQQmlComponent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlComponent_ConnectNotify
func callbackQQmlComponent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_CustomEvent
func callbackQQmlComponent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlComponent_DeleteLater
func callbackQQmlComponent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlComponentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlComponent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlComponent_Destroyed
func callbackQQmlComponent_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlComponent_DisconnectNotify
func callbackQQmlComponent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlComponentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlComponent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlComponent_ObjectNameChanged
func callbackQQmlComponent_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlComponent_TimerEvent
func callbackQQmlComponent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlComponentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlComponent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlComponent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlContext struct {
	core.QObject
}

type QQmlContext_ITF interface {
	core.QObject_ITF
	QQmlContext_PTR() *QQmlContext
}

func (ptr *QQmlContext) QQmlContext_PTR() *QQmlContext {
	return ptr
}

func (ptr *QQmlContext) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlContext) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlContext(ptr QQmlContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlContext_PTR().Pointer()
	}
	return nil
}

func NewQQmlContextFromPointer(ptr unsafe.Pointer) (n *QQmlContext) {
	n = new(QQmlContext)
	n.SetPointer(ptr)
	return
}
func NewQQmlContext2(parentContext QQmlContext_ITF, parent core.QObject_ITF) *QQmlContext {
	tmpValue := NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext2(PointerFromQQmlContext(parentContext), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlContext(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlContext {
	tmpValue := NewQQmlContextFromPointer(C.QQmlContext_NewQQmlContext(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlContext_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlContext_QQmlContext_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlContext) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlContext_QQmlContext_Tr(sC, cC, C.int(int32(n))))
}

func QQmlContext_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlContext_QQmlContext_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlContext) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlContext_QQmlContext_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlContext) ResolvedUrl(src core.QUrl_ITF) *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlContext_ResolvedUrl(ptr.Pointer(), core.PointerFromQUrl(src)))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) SetBaseUrl(baseUrl core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QQmlContext) SetContextObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_SetContextObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlContext) SetContextProperty(name string, value core.QObject_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QQmlContext_SetContextProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQObject(value))
	}
}

func (ptr *QQmlContext) SetContextProperty2(name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QQmlContext_SetContextProperty2(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))
	}
}

//export callbackQQmlContext_DestroyQQmlContext
func callbackQQmlContext_DestroyQQmlContext(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlContext"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlContextFromPointer(ptr).DestroyQQmlContextDefault()
	}
}

func (ptr *QQmlContext) ConnectDestroyQQmlContext(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlContext"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlContext", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlContext", f)
		}
	}
}

func (ptr *QQmlContext) DisconnectDestroyQQmlContext() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlContext")
	}
}

func (ptr *QQmlContext) DestroyQQmlContext() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContext(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlContext) DestroyQQmlContextDefault() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DestroyQQmlContextDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlContext) ContextObject() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlContext_ContextObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ParentContext() *QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlContextFromPointer(C.QQmlContext_ParentContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlEngineFromPointer(C.QQmlContext_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) NameForObject(object core.QObject_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlContext_NameForObject(ptr.Pointer(), core.PointerFromQObject(object)))
	}
	return ""
}

func (ptr *QQmlContext) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlContext_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) ContextProperty(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QQmlContext_ContextProperty(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlContext_IsValid(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQmlContext_MetaObject
func callbackQQmlContext_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlContextFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlContext) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlContext_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlContext) __setContextProperties_properties_newList() unsafe.Pointer {
	return C.QQmlContext___setContextProperties_properties_newList(ptr.Pointer())
}

func (ptr *QQmlContext) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlContext___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlContext) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlContext___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlContext) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlContext___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlContext) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlContext___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlContext) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlContext___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __findChildren_newList() unsafe.Pointer {
	return C.QQmlContext___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlContext) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlContext___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlContext) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlContext) __children_newList() unsafe.Pointer {
	return C.QQmlContext___children_newList(ptr.Pointer())
}

//export callbackQQmlContext_Event
func callbackQQmlContext_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlContext) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlContext_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlContext_EventFilter
func callbackQQmlContext_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlContextFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlContext) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlContext_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlContext_ChildEvent
func callbackQQmlContext_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlContext) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlContext_ConnectNotify
func callbackQQmlContext_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_CustomEvent
func callbackQQmlContext_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlContext) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlContext_DeleteLater
func callbackQQmlContext_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlContextFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlContext) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlContext_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlContext_Destroyed
func callbackQQmlContext_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlContext_DisconnectNotify
func callbackQQmlContext_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlContextFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlContext) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlContext_ObjectNameChanged
func callbackQQmlContext_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlContext_TimerEvent
func callbackQQmlContext_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlContextFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlContext) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlContext_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

func (ptr *QQmlEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QJSEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) (n *QQmlEngine) {
	n = new(QQmlEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQmlEngine__ObjectOwnership
//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(1)
)

func QQmlEngine_ContextForObject(object core.QObject_ITF) *QQmlContext {
	tmpValue := NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) ContextForObject(object core.QObject_ITF) *QQmlContext {
	tmpValue := NewQQmlContextFromPointer(C.QQmlEngine_QQmlEngine_ContextForObject(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	tmpValue := NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlEngine_ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ObjectOwnership(object core.QObject_ITF) QQmlEngine__ObjectOwnership {
	return QQmlEngine__ObjectOwnership(C.QQmlEngine_QQmlEngine_ObjectOwnership(core.PointerFromQObject(object)))
}

func (ptr *QQmlEngine) ImportPlugin(filePath string, uri string, errors []*QQmlError) bool {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		return int8(C.QQmlEngine_ImportPlugin(ptr.Pointer(), C.struct_QtQml_PackedString{data: filePathC, len: C.longlong(len(filePath))}, C.struct_QtQml_PackedString{data: uriC, len: C.longlong(len(uri))}, func() unsafe.Pointer {
			tmpList := NewQQmlEngineFromPointer(NewQQmlEngineFromPointer(nil).__importPlugin_errors_newList())
			for _, v := range errors {
				tmpList.__importPlugin_errors_setList(v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

func (ptr *QQmlEngine) AddImageProvider(providerId string, provider QQmlImageProviderBase_ITF) {
	if ptr.Pointer() != nil {
		var providerIdC *C.char
		if providerId != "" {
			providerIdC = C.CString(providerId)
			defer C.free(unsafe.Pointer(providerIdC))
		}
		C.QQmlEngine_AddImageProvider(ptr.Pointer(), C.struct_QtQml_PackedString{data: providerIdC, len: C.longlong(len(providerId))}, PointerFromQQmlImageProviderBase(provider))
	}
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.struct_QtQml_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QQmlEngine) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQmlEngine_AddPluginPath(ptr.Pointer(), C.struct_QtQml_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QQmlEngine) ClearComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_ClearComponentCache(ptr.Pointer())
	}
}

//export callbackQQmlEngine_Exit
func callbackQQmlEngine_Exit(ptr unsafe.Pointer, retCode C.int) {
	if signal := qt.GetSignal(ptr, "exit"); signal != nil {
		signal.(func(int))(int(int32(retCode)))
	}

}

func (ptr *QQmlEngine) ConnectExit(f func(retCode int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "exit") {
			C.QQmlEngine_ConnectExit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "exit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exit", func(retCode int) {
				signal.(func(int))(retCode)
				f(retCode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exit", f)
		}
	}
}

func (ptr *QQmlEngine) DisconnectExit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectExit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "exit")
	}
}

func (ptr *QQmlEngine) Exit(retCode int) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Exit(ptr.Pointer(), C.int(int32(retCode)))
	}
}

//export callbackQQmlEngine_Quit
func callbackQQmlEngine_Quit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "quit"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlEngine) ConnectQuit(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "quit") {
			C.QQmlEngine_ConnectQuit(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "quit"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "quit", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "quit", f)
		}
	}
}

func (ptr *QQmlEngine) DisconnectQuit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "quit")
	}
}

func (ptr *QQmlEngine) Quit() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Quit(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) RemoveImageProvider(providerId string) {
	if ptr.Pointer() != nil {
		var providerIdC *C.char
		if providerId != "" {
			providerIdC = C.CString(providerId)
			defer C.free(unsafe.Pointer(providerIdC))
		}
		C.QQmlEngine_RemoveImageProvider(ptr.Pointer(), C.struct_QtQml_PackedString{data: providerIdC, len: C.longlong(len(providerId))})
	}
}

//export callbackQQmlEngine_Retranslate
func callbackQQmlEngine_Retranslate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "retranslate"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlEngineFromPointer(ptr).RetranslateDefault()
	}
}

func (ptr *QQmlEngine) ConnectRetranslate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "retranslate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "retranslate", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "retranslate", f)
		}
	}
}

func (ptr *QQmlEngine) DisconnectRetranslate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "retranslate")
	}
}

func (ptr *QQmlEngine) Retranslate() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Retranslate(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) RetranslateDefault() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_RetranslateDefault(ptr.Pointer())
	}
}

func (ptr *QQmlEngine) SetBaseUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetBaseUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func QQmlEngine_SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetContextForObject(object core.QObject_ITF, context QQmlContext_ITF) {
	C.QQmlEngine_QQmlEngine_SetContextForObject(core.PointerFromQObject(object), PointerFromQQmlContext(context))
}

func (ptr *QQmlEngine) SetImportPathList(paths []string) {
	if ptr.Pointer() != nil {
		pathsC := C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetImportPathList(ptr.Pointer(), C.struct_QtQml_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))})
	}
}

func (ptr *QQmlEngine) SetIncubationController(controller QQmlIncubationController_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetIncubationController(ptr.Pointer(), PointerFromQQmlIncubationController(controller))
	}
}

func (ptr *QQmlEngine) SetNetworkAccessManagerFactory(factory QQmlNetworkAccessManagerFactory_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetNetworkAccessManagerFactory(ptr.Pointer(), PointerFromQQmlNetworkAccessManagerFactory(factory))
	}
}

func QQmlEngine_SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetObjectOwnership(object core.QObject_ITF, ownership QQmlEngine__ObjectOwnership) {
	C.QQmlEngine_QQmlEngine_SetObjectOwnership(core.PointerFromQObject(object), C.longlong(ownership))
}

func (ptr *QQmlEngine) SetOfflineStoragePath(dir string) {
	if ptr.Pointer() != nil {
		var dirC *C.char
		if dir != "" {
			dirC = C.CString(dir)
			defer C.free(unsafe.Pointer(dirC))
		}
		C.QQmlEngine_SetOfflineStoragePath(ptr.Pointer(), C.struct_QtQml_PackedString{data: dirC, len: C.longlong(len(dir))})
	}
}

func (ptr *QQmlEngine) SetOutputWarningsToStandardError(enabled bool) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_SetOutputWarningsToStandardError(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QQmlEngine) SetPluginPathList(paths []string) {
	if ptr.Pointer() != nil {
		pathsC := C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		C.QQmlEngine_SetPluginPathList(ptr.Pointer(), C.struct_QtQml_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))})
	}
}

func (ptr *QQmlEngine) TrimComponentCache() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_TrimComponentCache(ptr.Pointer())
	}
}

//export callbackQQmlEngine_Warnings
func callbackQQmlEngine_Warnings(ptr unsafe.Pointer, warnings C.struct_QtQml_PackedList) {
	if signal := qt.GetSignal(ptr, "warnings"); signal != nil {
		signal.(func([]*QQmlError))(func(l C.struct_QtQml_PackedList) []*QQmlError {
			out := make([]*QQmlError, int(l.len))
			tmpList := NewQQmlEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__warnings_warnings_atList(i)
			}
			return out
		}(warnings))
	}

}

func (ptr *QQmlEngine) ConnectWarnings(f func(warnings []*QQmlError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "warnings") {
			C.QQmlEngine_ConnectWarnings(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "warnings"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "warnings", func(warnings []*QQmlError) {
				signal.(func([]*QQmlError))(warnings)
				f(warnings)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "warnings", f)
		}
	}
}

func (ptr *QQmlEngine) DisconnectWarnings() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DisconnectWarnings(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "warnings")
	}
}

func (ptr *QQmlEngine) Warnings(warnings []*QQmlError) {
	if ptr.Pointer() != nil {
		C.QQmlEngine_Warnings(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQQmlEngineFromPointer(NewQQmlEngineFromPointer(nil).__warnings_warnings_newList())
			for _, v := range warnings {
				tmpList.__warnings_warnings_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQQmlEngine_DestroyQQmlEngine
func callbackQQmlEngine_DestroyQQmlEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlEngineFromPointer(ptr).DestroyQQmlEngineDefault()
	}
}

func (ptr *QQmlEngine) ConnectDestroyQQmlEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlEngine", f)
		}
	}
}

func (ptr *QQmlEngine) DisconnectDestroyQQmlEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlEngine")
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngineDefault() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlEngine) NetworkAccessManager() *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkAccessManagerFromPointer(C.QQmlEngine_NetworkAccessManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) RootContext() *QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlContextFromPointer(C.QQmlEngine_RootContext(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) ImageProvider(providerId string) *QQmlImageProviderBase {
	if ptr.Pointer() != nil {
		var providerIdC *C.char
		if providerId != "" {
			providerIdC = C.CString(providerId)
			defer C.free(unsafe.Pointer(providerIdC))
		}
		return NewQQmlImageProviderBaseFromPointer(C.QQmlEngine_ImageProvider(ptr.Pointer(), C.struct_QtQml_PackedString{data: providerIdC, len: C.longlong(len(providerId))}))
	}
	return nil
}

func (ptr *QQmlEngine) IncubationController() *QQmlIncubationController {
	if ptr.Pointer() != nil {
		return NewQQmlIncubationControllerFromPointer(C.QQmlEngine_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) NetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	if ptr.Pointer() != nil {
		return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlEngine_NetworkAccessManagerFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlEngine) OfflineStorageDatabaseFilePath(databaseName string) string {
	if ptr.Pointer() != nil {
		var databaseNameC *C.char
		if databaseName != "" {
			databaseNameC = C.CString(databaseName)
			defer C.free(unsafe.Pointer(databaseNameC))
		}
		return cGoUnpackString(C.QQmlEngine_OfflineStorageDatabaseFilePath(ptr.Pointer(), C.struct_QtQml_PackedString{data: databaseNameC, len: C.longlong(len(databaseName))}))
	}
	return ""
}

func (ptr *QQmlEngine) OfflineStoragePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlEngine_OfflineStoragePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlEngine) ImportPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlEngine_ImportPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) PluginPathList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlEngine_PluginPathList(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlEngine) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlEngine_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) OutputWarningsToStandardError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlEngine_OutputWarningsToStandardError(ptr.Pointer())) != 0
	}
	return false
}

func QQmlEngine_QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterSingletonType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) QmlRegisterSingletonType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterSingletonType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func QQmlEngine_QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) __importPlugin_errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___importPlugin_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __importPlugin_errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___importPlugin_errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __importPlugin_errors_newList() unsafe.Pointer {
	return C.QQmlEngine___importPlugin_errors_newList(ptr.Pointer())
}

func (ptr *QQmlEngine) __warnings_warnings_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___warnings_warnings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __warnings_warnings_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___warnings_warnings_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __warnings_warnings_newList() unsafe.Pointer {
	return C.QQmlEngine___warnings_warnings_newList(ptr.Pointer())
}

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func (ptr *QQmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) (n *QQmlError) {
	n = new(QQmlError)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlError) DestroyQQmlError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQQmlError() *QQmlError {
	tmpValue := NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	tmpValue := NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func (ptr *QQmlError) SetColumn(column int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(int32(column)))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QQmlError_SetDescription(ptr.Pointer(), C.struct_QtQml_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func (ptr *QQmlError) SetLine(line int) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(int32(line)))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) Object() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlError) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlError_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlError_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Line() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Line(ptr.Pointer())))
	}
	return 0
}

type QQmlExpression struct {
	core.QObject
}

type QQmlExpression_ITF interface {
	core.QObject_ITF
	QQmlExpression_PTR() *QQmlExpression
}

func (ptr *QQmlExpression) QQmlExpression_PTR() *QQmlExpression {
	return ptr
}

func (ptr *QQmlExpression) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExpression) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExpression(ptr QQmlExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExpression_PTR().Pointer()
	}
	return nil
}

func NewQQmlExpressionFromPointer(ptr unsafe.Pointer) (n *QQmlExpression) {
	n = new(QQmlExpression)
	n.SetPointer(ptr)
	return
}
func NewQQmlExpression() *QQmlExpression {
	tmpValue := NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression2(ctxt QQmlContext_ITF, scope core.QObject_ITF, expression string, parent core.QObject_ITF) *QQmlExpression {
	var expressionC *C.char
	if expression != "" {
		expressionC = C.CString(expression)
		defer C.free(unsafe.Pointer(expressionC))
	}
	tmpValue := NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression2(PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), C.struct_QtQml_PackedString{data: expressionC, len: C.longlong(len(expression))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlExpression3(scri QQmlScriptString_ITF, ctxt QQmlContext_ITF, scope core.QObject_ITF, parent core.QObject_ITF) *QQmlExpression {
	tmpValue := NewQQmlExpressionFromPointer(C.QQmlExpression_NewQQmlExpression3(PointerFromQQmlScriptString(scri), PointerFromQQmlContext(ctxt), core.PointerFromQObject(scope), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlExpression_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExpression_QQmlExpression_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlExpression) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExpression_QQmlExpression_Tr(sC, cC, C.int(int32(n))))
}

func QQmlExpression_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExpression_QQmlExpression_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlExpression) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExpression_QQmlExpression_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlExpression) Evaluate(valueIsUndefined *bool) *core.QVariant {
	if ptr.Pointer() != nil {
		valueIsUndefinedC := C.char(int8(qt.GoBoolToInt(*valueIsUndefined)))
		defer func() { *valueIsUndefined = int8(valueIsUndefinedC) != 0 }()
		tmpValue := core.NewQVariantFromPointer(C.QQmlExpression_Evaluate(ptr.Pointer(), &valueIsUndefinedC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) ClearError() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ClearError(ptr.Pointer())
	}
}

func (ptr *QQmlExpression) SetExpression(expression string) {
	if ptr.Pointer() != nil {
		var expressionC *C.char
		if expression != "" {
			expressionC = C.CString(expression)
			defer C.free(unsafe.Pointer(expressionC))
		}
		C.QQmlExpression_SetExpression(ptr.Pointer(), C.struct_QtQml_PackedString{data: expressionC, len: C.longlong(len(expression))})
	}
}

func (ptr *QQmlExpression) SetNotifyOnValueChanged(notifyOnChange bool) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_SetNotifyOnValueChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(notifyOnChange))))
	}
}

func (ptr *QQmlExpression) SetSourceLocation(url string, line int, column int) {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		C.QQmlExpression_SetSourceLocation(ptr.Pointer(), C.struct_QtQml_PackedString{data: urlC, len: C.longlong(len(url))}, C.int(int32(line)), C.int(int32(column)))
	}
}

//export callbackQQmlExpression_ValueChanged
func callbackQQmlExpression_ValueChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlExpression) ConnectValueChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QQmlExpression_ConnectValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", f)
		}
	}
}

func (ptr *QQmlExpression) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QQmlExpression) ValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ValueChanged(ptr.Pointer())
	}
}

//export callbackQQmlExpression_DestroyQQmlExpression
func callbackQQmlExpression_DestroyQQmlExpression(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlExpression"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExpressionFromPointer(ptr).DestroyQQmlExpressionDefault()
	}
}

func (ptr *QQmlExpression) ConnectDestroyQQmlExpression(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlExpression"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlExpression", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlExpression", f)
		}
	}
}

func (ptr *QQmlExpression) DisconnectDestroyQQmlExpression() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlExpression")
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpression() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpression(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlExpression) DestroyQQmlExpressionDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DestroyQQmlExpressionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlExpression) ScopeObject() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExpression_ScopeObject(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Context() *QQmlContext {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlContextFromPointer(C.QQmlExpression_Context(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlEngineFromPointer(C.QQmlExpression_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Error() *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlExpression_Error(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) Expression() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlExpression_Expression(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) SourceFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlExpression_SourceFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlExpression) HasError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExpression_HasError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlExpression) NotifyOnValueChanged() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExpression_NotifyOnValueChanged(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQmlExpression_MetaObject
func callbackQQmlExpression_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExpressionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExpression) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExpression_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExpression) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_ColumnNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlExpression_LineNumber(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlExpression) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlExpression___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlExpression) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlExpression___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlExpression) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlExpression___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlExpression) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlExpression___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlExpression) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExpression___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __findChildren_newList() unsafe.Pointer {
	return C.QQmlExpression___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlExpression) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExpression___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExpression) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExpression) __children_newList() unsafe.Pointer {
	return C.QQmlExpression___children_newList(ptr.Pointer())
}

//export callbackQQmlExpression_Event
func callbackQQmlExpression_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExpression) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExpression_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlExpression_EventFilter
func callbackQQmlExpression_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExpressionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExpression) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExpression_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlExpression_ChildEvent
func callbackQQmlExpression_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExpression_ConnectNotify
func callbackQQmlExpression_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_CustomEvent
func callbackQQmlExpression_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExpression_DeleteLater
func callbackQQmlExpression_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExpressionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExpression) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlExpression_Destroyed
func callbackQQmlExpression_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlExpression_DisconnectNotify
func callbackQQmlExpression_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExpressionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExpression) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExpression_ObjectNameChanged
func callbackQQmlExpression_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlExpression_TimerEvent
func callbackQQmlExpression_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExpressionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExpression) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExpression_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) (n *QQmlExtensionPlugin) {
	n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	return
}
func NewQQmlExtensionPlugin(parent core.QObject_ITF) *QQmlExtensionPlugin {
	tmpValue := NewQQmlExtensionPluginFromPointer(C.QQmlExtensionPlugin_NewQQmlExtensionPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlExtensionPlugin_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExtensionPlugin_QQmlExtensionPlugin_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlExtensionPlugin) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExtensionPlugin_QQmlExtensionPlugin_Tr(sC, cC, C.int(int32(n))))
}

func QQmlExtensionPlugin_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExtensionPlugin_QQmlExtensionPlugin_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlExtensionPlugin) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlExtensionPlugin_QQmlExtensionPlugin_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQmlExtensionPlugin_InitializeEngine
func callbackQQmlExtensionPlugin_InitializeEngine(ptr unsafe.Pointer, engine unsafe.Pointer, uri C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "initializeEngine"); signal != nil {
		signal.(func(*QQmlEngine, string))(NewQQmlEngineFromPointer(engine), cGoUnpackString(uri))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).InitializeEngineDefault(NewQQmlEngineFromPointer(engine), cGoUnpackString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectInitializeEngine(f func(engine *QQmlEngine, uri string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initializeEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initializeEngine", func(engine *QQmlEngine, uri string) {
				signal.(func(*QQmlEngine, string))(engine, uri)
				f(engine, uri)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initializeEngine", f)
		}
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectInitializeEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "initializeEngine")
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QQmlExtensionPlugin_InitializeEngine(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

func (ptr *QQmlExtensionPlugin) InitializeEngineDefault(engine QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QQmlExtensionPlugin_InitializeEngineDefault(ptr.Pointer(), PointerFromQQmlEngine(engine), uriC)
	}
}

//export callbackQQmlExtensionPlugin_RegisterTypes
func callbackQQmlExtensionPlugin_RegisterTypes(ptr unsafe.Pointer, uri C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "registerTypes"); signal != nil {
		signal.(func(string))(cGoUnpackString(uri))
	}

}

func (ptr *QQmlExtensionPlugin) ConnectRegisterTypes(f func(uri string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerTypes"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "registerTypes", func(uri string) {
				signal.(func(string))(uri)
				f(uri)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerTypes", f)
		}
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectRegisterTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerTypes")
	}
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		var uriC *C.char
		if uri != "" {
			uriC = C.CString(uri)
			defer C.free(unsafe.Pointer(uriC))
		}
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), uriC)
	}
}

func (ptr *QQmlExtensionPlugin) BaseUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlExtensionPlugin_BaseUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//export callbackQQmlExtensionPlugin_MetaObject
func callbackQQmlExtensionPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlExtensionPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlExtensionPlugin) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlExtensionPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlExtensionPlugin___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlExtensionPlugin) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlExtensionPlugin___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlExtensionPlugin___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlExtensionPlugin___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlExtensionPlugin) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExtensionPlugin___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __findChildren_newList() unsafe.Pointer {
	return C.QQmlExtensionPlugin___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlExtensionPlugin) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlExtensionPlugin___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlExtensionPlugin) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlExtensionPlugin) __children_newList() unsafe.Pointer {
	return C.QQmlExtensionPlugin___children_newList(ptr.Pointer())
}

//export callbackQQmlExtensionPlugin_Event
func callbackQQmlExtensionPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlExtensionPlugin) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExtensionPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_EventFilter
func callbackQQmlExtensionPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlExtensionPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlExtensionPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlExtensionPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlExtensionPlugin_ChildEvent
func callbackQQmlExtensionPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_ConnectNotify
func callbackQQmlExtensionPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_CustomEvent
func callbackQQmlExtensionPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlExtensionPlugin_DeleteLater
func callbackQQmlExtensionPlugin_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlExtensionPlugin) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlExtensionPlugin_Destroyed
func callbackQQmlExtensionPlugin_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlExtensionPlugin_DisconnectNotify
func callbackQQmlExtensionPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlExtensionPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlExtensionPlugin_ObjectNameChanged
func callbackQQmlExtensionPlugin_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlExtensionPlugin_TimerEvent
func callbackQQmlExtensionPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlExtensionPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlExtensionPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func (ptr *QQmlFileSelector) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlFileSelector) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) (n *QQmlFileSelector) {
	n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	return
}
func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	tmpValue := NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlFileSelector) Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	tmpValue := NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {
	tmpValue := NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlFileSelector_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlFileSelector_QQmlFileSelector_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlFileSelector) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlFileSelector_QQmlFileSelector_Tr(sC, cC, C.int(int32(n))))
}

func QQmlFileSelector_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlFileSelector_QQmlFileSelector_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlFileSelector) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlFileSelector_QQmlFileSelector_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	if ptr.Pointer() != nil {
		strinC := C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), C.struct_QtQml_PackedString{data: strinC, len: C.longlong(len(strings.Join(strin, "|")))})
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	if ptr.Pointer() != nil {
		strinC := C.CString(strings.Join(strin, "|"))
		defer C.free(unsafe.Pointer(strinC))
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), C.struct_QtQml_PackedString{data: strinC, len: C.longlong(len(strings.Join(strin, "|")))})
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

//export callbackQQmlFileSelector_DestroyQQmlFileSelector
func callbackQQmlFileSelector_DestroyQQmlFileSelector(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlFileSelector"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DestroyQQmlFileSelectorDefault()
	}
}

func (ptr *QQmlFileSelector) ConnectDestroyQQmlFileSelector(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlFileSelector"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlFileSelector", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlFileSelector", f)
		}
	}
}

func (ptr *QQmlFileSelector) DisconnectDestroyQQmlFileSelector() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlFileSelector")
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelectorDefault() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelectorDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlFileSelector) Selector() *core.QFileSelector {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQFileSelectorFromPointer(C.QQmlFileSelector_Selector(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlFileSelector_MetaObject
func callbackQQmlFileSelector_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlFileSelectorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlFileSelector) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlFileSelector_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlFileSelector___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlFileSelector) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlFileSelector___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlFileSelector) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlFileSelector___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlFileSelector) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlFileSelector___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlFileSelector) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlFileSelector___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __findChildren_newList() unsafe.Pointer {
	return C.QQmlFileSelector___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlFileSelector) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlFileSelector___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlFileSelector) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlFileSelector) __children_newList() unsafe.Pointer {
	return C.QQmlFileSelector___children_newList(ptr.Pointer())
}

//export callbackQQmlFileSelector_Event
func callbackQQmlFileSelector_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlFileSelector) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlFileSelector_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlFileSelector_EventFilter
func callbackQQmlFileSelector_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlFileSelectorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlFileSelector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlFileSelector_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlFileSelector_ChildEvent
func callbackQQmlFileSelector_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlFileSelector_ConnectNotify
func callbackQQmlFileSelector_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_CustomEvent
func callbackQQmlFileSelector_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlFileSelector_DeleteLater
func callbackQQmlFileSelector_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlFileSelector) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlFileSelector_Destroyed
func callbackQQmlFileSelector_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlFileSelector_DisconnectNotify
func callbackQQmlFileSelector_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlFileSelector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlFileSelector_ObjectNameChanged
func callbackQQmlFileSelector_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlFileSelector_TimerEvent
func callbackQQmlFileSelector_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlFileSelectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlFileSelector) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlFileSelector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlImageProviderBase struct {
	ptr unsafe.Pointer
}

type QQmlImageProviderBase_ITF interface {
	QQmlImageProviderBase_PTR() *QQmlImageProviderBase
}

func (ptr *QQmlImageProviderBase) QQmlImageProviderBase_PTR() *QQmlImageProviderBase {
	return ptr
}

func (ptr *QQmlImageProviderBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlImageProviderBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlImageProviderBase(ptr QQmlImageProviderBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func NewQQmlImageProviderBaseFromPointer(ptr unsafe.Pointer) (n *QQmlImageProviderBase) {
	n = new(QQmlImageProviderBase)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlImageProviderBase) DestroyQQmlImageProviderBase() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QQmlImageProviderBase__Flag
//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading QQmlImageProviderBase__Flag = QQmlImageProviderBase__Flag(0x01)
)

//go:generate stringer -type=QQmlImageProviderBase__ImageType
//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image         QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap        QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(3)
	QQmlImageProviderBase__ImageResponse QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(4)
)

//export callbackQQmlImageProviderBase_Flags
func callbackQQmlImageProviderBase_Flags(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__Flag)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectFlags(f func() QQmlImageProviderBase__Flag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "flags"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "flags", func() QQmlImageProviderBase__Flag {
				signal.(func() QQmlImageProviderBase__Flag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "flags", f)
		}
	}
}

func (ptr *QQmlImageProviderBase) DisconnectFlags() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "flags")
	}
}

func (ptr *QQmlImageProviderBase) Flags() QQmlImageProviderBase__Flag {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__Flag(C.QQmlImageProviderBase_Flags(ptr.Pointer()))
	}
	return 0
}

//export callbackQQmlImageProviderBase_ImageType
func callbackQQmlImageProviderBase_ImageType(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "imageType"); signal != nil {
		return C.longlong(signal.(func() QQmlImageProviderBase__ImageType)())
	}

	return C.longlong(0)
}

func (ptr *QQmlImageProviderBase) ConnectImageType(f func() QQmlImageProviderBase__ImageType) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "imageType"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "imageType", func() QQmlImageProviderBase__ImageType {
				signal.(func() QQmlImageProviderBase__ImageType)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "imageType", f)
		}
	}
}

func (ptr *QQmlImageProviderBase) DisconnectImageType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "imageType")
	}
}

func (ptr *QQmlImageProviderBase) ImageType() QQmlImageProviderBase__ImageType {
	if ptr.Pointer() != nil {
		return QQmlImageProviderBase__ImageType(C.QQmlImageProviderBase_ImageType(ptr.Pointer()))
	}
	return 0
}

type QQmlImport struct {
	ptr unsafe.Pointer
}

type QQmlImport_ITF interface {
	QQmlImport_PTR() *QQmlImport
}

func (ptr *QQmlImport) QQmlImport_PTR() *QQmlImport {
	return ptr
}

func (ptr *QQmlImport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlImport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlImport(ptr QQmlImport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlImport_PTR().Pointer()
	}
	return nil
}

func NewQQmlImportFromPointer(ptr unsafe.Pointer) (n *QQmlImport) {
	n = new(QQmlImport)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlImport) DestroyQQmlImport() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QQmlImport__RecursionRestriction
//QQmlImport::RecursionRestriction
type QQmlImport__RecursionRestriction int64

const (
	QQmlImport__PreventRecursion QQmlImport__RecursionRestriction = QQmlImport__RecursionRestriction(0)
	QQmlImport__AllowRecursion   QQmlImport__RecursionRestriction = QQmlImport__RecursionRestriction(1)
)

type QQmlIncubationController struct {
	ptr unsafe.Pointer
}

type QQmlIncubationController_ITF interface {
	QQmlIncubationController_PTR() *QQmlIncubationController
}

func (ptr *QQmlIncubationController) QQmlIncubationController_PTR() *QQmlIncubationController {
	return ptr
}

func (ptr *QQmlIncubationController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlIncubationController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlIncubationController(ptr QQmlIncubationController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubationController_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubationControllerFromPointer(ptr unsafe.Pointer) (n *QQmlIncubationController) {
	n = new(QQmlIncubationController)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlIncubationController) DestroyQQmlIncubationController() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQQmlIncubationController() *QQmlIncubationController {
	return NewQQmlIncubationControllerFromPointer(C.QQmlIncubationController_NewQQmlIncubationController())
}

func (ptr *QQmlIncubationController) IncubateFor(msecs int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubateFor(ptr.Pointer(), C.int(int32(msecs)))
	}
}

//export callbackQQmlIncubationController_IncubatingObjectCountChanged
func callbackQQmlIncubationController_IncubatingObjectCountChanged(ptr unsafe.Pointer, incubatingObjectCount C.int) {
	if signal := qt.GetSignal(ptr, "incubatingObjectCountChanged"); signal != nil {
		signal.(func(int))(int(int32(incubatingObjectCount)))
	} else {
		NewQQmlIncubationControllerFromPointer(ptr).IncubatingObjectCountChangedDefault(int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) ConnectIncubatingObjectCountChanged(f func(incubatingObjectCount int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "incubatingObjectCountChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "incubatingObjectCountChanged", func(incubatingObjectCount int) {
				signal.(func(int))(incubatingObjectCount)
				f(incubatingObjectCount)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "incubatingObjectCountChanged", f)
		}
	}
}

func (ptr *QQmlIncubationController) DisconnectIncubatingObjectCountChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "incubatingObjectCountChanged")
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChanged(incubatingObjectCount int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChanged(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) IncubatingObjectCountChangedDefault(incubatingObjectCount int) {
	if ptr.Pointer() != nil {
		C.QQmlIncubationController_IncubatingObjectCountChangedDefault(ptr.Pointer(), C.int(int32(incubatingObjectCount)))
	}
}

func (ptr *QQmlIncubationController) Engine() *QQmlEngine {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlEngineFromPointer(C.QQmlIncubationController_Engine(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlIncubationController) IncubatingObjectCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlIncubationController_IncubatingObjectCount(ptr.Pointer())))
	}
	return 0
}

type QQmlIncubator struct {
	ptr unsafe.Pointer
}

type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator {
	return ptr
}

func (ptr *QQmlIncubator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlIncubator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlIncubator(ptr QQmlIncubator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlIncubator_PTR().Pointer()
	}
	return nil
}

func NewQQmlIncubatorFromPointer(ptr unsafe.Pointer) (n *QQmlIncubator) {
	n = new(QQmlIncubator)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlIncubator) DestroyQQmlIncubator() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QQmlIncubator__IncubationMode
//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(2)
)

//go:generate stringer -type=QQmlIncubator__Status
//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    QQmlIncubator__Status = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   QQmlIncubator__Status = QQmlIncubator__Status(1)
	QQmlIncubator__Loading QQmlIncubator__Status = QQmlIncubator__Status(2)
	QQmlIncubator__Error   QQmlIncubator__Status = QQmlIncubator__Status(3)
)

func NewQQmlIncubator(mode QQmlIncubator__IncubationMode) *QQmlIncubator {
	return NewQQmlIncubatorFromPointer(C.QQmlIncubator_NewQQmlIncubator(C.longlong(mode)))
}

func (ptr *QQmlIncubator) Clear() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_Clear(ptr.Pointer())
	}
}

func (ptr *QQmlIncubator) ForceCompletion() {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_ForceCompletion(ptr.Pointer())
	}
}

//export callbackQQmlIncubator_SetInitialState
func callbackQQmlIncubator_SetInitialState(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setInitialState"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	} else {
		NewQQmlIncubatorFromPointer(ptr).SetInitialStateDefault(core.NewQObjectFromPointer(object))
	}
}

func (ptr *QQmlIncubator) ConnectSetInitialState(f func(object *core.QObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInitialState"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setInitialState", func(object *core.QObject) {
				signal.(func(*core.QObject))(object)
				f(object)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setInitialState", f)
		}
	}
}

func (ptr *QQmlIncubator) DisconnectSetInitialState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setInitialState")
	}
}

func (ptr *QQmlIncubator) SetInitialState(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialState(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlIncubator) SetInitialStateDefault(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_SetInitialStateDefault(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQQmlIncubator_StatusChanged
func callbackQQmlIncubator_StatusChanged(ptr unsafe.Pointer, status C.longlong) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		signal.(func(QQmlIncubator__Status))(QQmlIncubator__Status(status))
	} else {
		NewQQmlIncubatorFromPointer(ptr).StatusChangedDefault(QQmlIncubator__Status(status))
	}
}

func (ptr *QQmlIncubator) ConnectStatusChanged(f func(status QQmlIncubator__Status)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", func(status QQmlIncubator__Status) {
				signal.(func(QQmlIncubator__Status))(status)
				f(status)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", f)
		}
	}
}

func (ptr *QQmlIncubator) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *QQmlIncubator) StatusChanged(status QQmlIncubator__Status) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChanged(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlIncubator) StatusChangedDefault(status QQmlIncubator__Status) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator_StatusChangedDefault(ptr.Pointer(), C.longlong(status))
	}
}

func (ptr *QQmlIncubator) Errors() []*QQmlError {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtQml_PackedList) []*QQmlError {
			out := make([]*QQmlError, int(l.len))
			tmpList := NewQQmlIncubatorFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__errors_atList(i)
			}
			return out
		}(C.QQmlIncubator_Errors(ptr.Pointer()))
	}
	return make([]*QQmlError, 0)
}

func (ptr *QQmlIncubator) Object() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlIncubator_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlIncubator) IncubationMode() QQmlIncubator__IncubationMode {
	if ptr.Pointer() != nil {
		return QQmlIncubator__IncubationMode(C.QQmlIncubator_IncubationMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) Status() QQmlIncubator__Status {
	if ptr.Pointer() != nil {
		return QQmlIncubator__Status(C.QQmlIncubator_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlIncubator) IsError() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlIncubator_IsError(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsLoading() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlIncubator_IsLoading(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlIncubator_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) IsReady() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlIncubator_IsReady(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlIncubator) __errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlIncubator___errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlIncubator) __errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlIncubator___errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlIncubator) __errors_newList() unsafe.Pointer {
	return C.QQmlIncubator___errors_newList(ptr.Pointer())
}

type QQmlListProperty struct {
	ptr unsafe.Pointer
}

type QQmlListProperty_ITF interface {
	QQmlListProperty_PTR() *QQmlListProperty
}

func (ptr *QQmlListProperty) QQmlListProperty_PTR() *QQmlListProperty {
	return ptr
}

func (ptr *QQmlListProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlListProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlListProperty(ptr QQmlListProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlListPropertyFromPointer(ptr unsafe.Pointer) (n *QQmlListProperty) {
	n = new(QQmlListProperty)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlListProperty) DestroyQQmlListProperty() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (ptr *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return ptr
}

func (ptr *QQmlListReference) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlListReference) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlListReference(ptr QQmlListReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReference_PTR().Pointer()
	}
	return nil
}

func NewQQmlListReferenceFromPointer(ptr unsafe.Pointer) (n *QQmlListReference) {
	n = new(QQmlListReference)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlListReference) DestroyQQmlListReference() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQQmlListReference() *QQmlListReference {
	tmpValue := NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference())
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {
	var propertyC *C.char
	if property != "" {
		propertyC = C.CString(property)
		defer C.free(unsafe.Pointer(propertyC))
	}
	tmpValue := NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference2(core.PointerFromQObject(object), propertyC, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlListReference).DestroyQQmlListReference)
	return tmpValue
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlListReference_At(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlListReference_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_Append(ptr.Pointer(), core.PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAppend() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_CanAppend(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_CanAt(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_CanClear(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_CanCount(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_Clear(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsManipulable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_IsManipulable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_IsReadable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlListReference_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlListReference_ListElementType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlListReference) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlListReference_Count(ptr.Pointer())))
	}
	return 0
}

type QQmlNetworkAccessManagerFactory struct {
	ptr unsafe.Pointer
}

type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (ptr *QQmlNetworkAccessManagerFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlNetworkAccessManagerFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlNetworkAccessManagerFactory(ptr QQmlNetworkAccessManagerFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlNetworkAccessManagerFactory_PTR().Pointer()
	}
	return nil
}

func NewQQmlNetworkAccessManagerFactoryFromPointer(ptr unsafe.Pointer) (n *QQmlNetworkAccessManagerFactory) {
	n = new(QQmlNetworkAccessManagerFactory)
	n.SetPointer(ptr)
	return
}

//export callbackQQmlNetworkAccessManagerFactory_Create
func callbackQQmlNetworkAccessManagerFactory_Create(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "create"); signal != nil {
		return network.PointerFromQNetworkAccessManager(signal.(func(*core.QObject) *network.QNetworkAccessManager)(core.NewQObjectFromPointer(parent)))
	}

	return network.PointerFromQNetworkAccessManager(network.NewQNetworkAccessManager(nil))
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectCreate(f func(parent *core.QObject) *network.QNetworkAccessManager) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "create"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "create", func(parent *core.QObject) *network.QNetworkAccessManager {
				signal.(func(*core.QObject) *network.QNetworkAccessManager)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "create", f)
		}
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "create")
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) Create(parent core.QObject_ITF) *network.QNetworkAccessManager {
	if ptr.Pointer() != nil {
		tmpValue := network.NewQNetworkAccessManagerFromPointer(C.QQmlNetworkAccessManagerFactory_Create(ptr.Pointer(), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory
func callbackQQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlNetworkAccessManagerFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlNetworkAccessManagerFactoryFromPointer(ptr).DestroyQQmlNetworkAccessManagerFactoryDefault()
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) ConnectDestroyQQmlNetworkAccessManagerFactory(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlNetworkAccessManagerFactory"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlNetworkAccessManagerFactory", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlNetworkAccessManagerFactory", f)
		}
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DisconnectDestroyQQmlNetworkAccessManagerFactory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlNetworkAccessManagerFactory")
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactory() {
	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlNetworkAccessManagerFactory) DestroyQQmlNetworkAccessManagerFactoryDefault() {
	if ptr.Pointer() != nil {
		C.QQmlNetworkAccessManagerFactory_DestroyQQmlNetworkAccessManagerFactoryDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQQmlNetworkAccessManagerFactory() *QQmlNetworkAccessManagerFactory {
	return NewQQmlNetworkAccessManagerFactoryFromPointer(C.QQmlNetworkAccessManagerFactory_NewQQmlNetworkAccessManagerFactory())
}

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlParserStatus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) (n *QQmlParserStatus) {
	n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlParserStatus) DestroyQQmlParserStatus() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlParserStatus_ClassBegin
func callbackQQmlParserStatus_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "classBegin"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "classBegin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "classBegin", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "classBegin", f)
		}
	}
}

func (ptr *QQmlParserStatus) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "classBegin")
	}
}

func (ptr *QQmlParserStatus) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

//export callbackQQmlParserStatus_ComponentComplete
func callbackQQmlParserStatus_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "componentComplete"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQmlParserStatus) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "componentComplete"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", f)
		}
	}
}

func (ptr *QQmlParserStatus) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "componentComplete")
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}

type QQmlProperty struct {
	ptr unsafe.Pointer
}

type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty {
	return ptr
}

func (ptr *QQmlProperty) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlProperty) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlProperty(ptr QQmlProperty_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlProperty_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyFromPointer(ptr unsafe.Pointer) (n *QQmlProperty) {
	n = new(QQmlProperty)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlProperty) DestroyQQmlProperty() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QQmlProperty__PropertyTypeCategory
//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(3)
)

//go:generate stringer -type=QQmlProperty__Type
//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        QQmlProperty__Type = QQmlProperty__Type(0)
	QQmlProperty__Property       QQmlProperty__Type = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty QQmlProperty__Type = QQmlProperty__Type(2)
)

func NewQQmlProperty() *QQmlProperty {
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty())
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty2(obj core.QObject_ITF) *QQmlProperty {
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty2(core.PointerFromQObject(obj)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty3(obj core.QObject_ITF, ctxt QQmlContext_ITF) *QQmlProperty {
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty3(core.PointerFromQObject(obj), PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty4(obj core.QObject_ITF, engine QQmlEngine_ITF) *QQmlProperty {
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty4(core.PointerFromQObject(obj), PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty5(obj core.QObject_ITF, name string) *QQmlProperty {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty5(core.PointerFromQObject(obj), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty6(obj core.QObject_ITF, name string, ctxt QQmlContext_ITF) *QQmlProperty {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty6(core.PointerFromQObject(obj), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty7(obj core.QObject_ITF, name string, engine QQmlEngine_ITF) *QQmlProperty {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty7(core.PointerFromQObject(obj), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func NewQQmlProperty8(other QQmlProperty_ITF) *QQmlProperty {
	tmpValue := NewQQmlPropertyFromPointer(C.QQmlProperty_NewQQmlProperty8(PointerFromQQmlProperty(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlProperty).DestroyQQmlProperty)
	return tmpValue
}

func QQmlProperty_Read2(object core.QObject_ITF, name string) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read2(object core.QObject_ITF, name string) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read2(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read3(object core.QObject_ITF, name string, ctxt QQmlContext_ITF) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read3(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlContext(ctxt)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QQmlProperty) Read4(object core.QObject_ITF, name string, engine QQmlEngine_ITF) *core.QVariant {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_QQmlProperty_Read4(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQQmlEngine(engine)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QQmlProperty_Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))) != 0
}

func (ptr *QQmlProperty) Write2(object core.QObject_ITF, name string, value core.QVariant_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write2(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))) != 0
}

func QQmlProperty_Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt))) != 0
}

func (ptr *QQmlProperty) Write3(object core.QObject_ITF, name string, value core.QVariant_ITF, ctxt QQmlContext_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write3(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), PointerFromQQmlContext(ctxt))) != 0
}

func QQmlProperty_Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine))) != 0
}

func (ptr *QQmlProperty) Write4(object core.QObject_ITF, name string, value core.QVariant_ITF, engine QQmlEngine_ITF) bool {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	return int8(C.QQmlProperty_QQmlProperty_Write4(core.PointerFromQObject(object), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), PointerFromQQmlEngine(engine))) != 0
}

func (ptr *QQmlProperty) Method() *core.QMetaMethod {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQMetaMethodFromPointer(C.QQmlProperty_Method(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QMetaMethod).DestroyQMetaMethod)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) Object() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlProperty_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) PropertyTypeCategory() QQmlProperty__PropertyTypeCategory {
	if ptr.Pointer() != nil {
		return QQmlProperty__PropertyTypeCategory(C.QQmlProperty_PropertyTypeCategory(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) Type() QQmlProperty__Type {
	if ptr.Pointer() != nil {
		return QQmlProperty__Type(C.QQmlProperty_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlProperty) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlProperty_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) Read() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QQmlProperty_Read(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlProperty) ConnectNotifySignal(dest core.QObject_ITF, slot string) bool {
	if ptr.Pointer() != nil {
		var slotC *C.char
		if slot != "" {
			slotC = C.CString(slot)
			defer C.free(unsafe.Pointer(slotC))
		}
		return int8(C.QQmlProperty_ConnectNotifySignal(ptr.Pointer(), core.PointerFromQObject(dest), slotC)) != 0
	}
	return false
}

func (ptr *QQmlProperty) ConnectNotifySignal2(dest core.QObject_ITF, method int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_ConnectNotifySignal2(ptr.Pointer(), core.PointerFromQObject(dest), C.int(int32(method)))) != 0
	}
	return false
}

func (ptr *QQmlProperty) HasNotifySignal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_HasNotifySignal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsDesignable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsDesignable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsProperty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsProperty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsResettable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsResettable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsSignalProperty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsSignalProperty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) IsWritable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_IsWritable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) NeedsNotifySignal() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_NeedsNotifySignal(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Reset() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_Reset(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlProperty) Write(value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlProperty_Write(ptr.Pointer(), core.PointerFromQVariant(value))) != 0
	}
	return false
}

func (ptr *QQmlProperty) PropertyTypeName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlProperty_PropertyTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlProperty) Index() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_Index(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlProperty) PropertyType() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlProperty_PropertyType(ptr.Pointer())))
	}
	return 0
}

type QQmlPropertyMap struct {
	core.QObject
}

type QQmlPropertyMap_ITF interface {
	core.QObject_ITF
	QQmlPropertyMap_PTR() *QQmlPropertyMap
}

func (ptr *QQmlPropertyMap) QQmlPropertyMap_PTR() *QQmlPropertyMap {
	return ptr
}

func (ptr *QQmlPropertyMap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlPropertyMap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlPropertyMap(ptr QQmlPropertyMap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyMap_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyMapFromPointer(ptr unsafe.Pointer) (n *QQmlPropertyMap) {
	n = new(QQmlPropertyMap)
	n.SetPointer(ptr)
	return
}
func NewQQmlPropertyMap(parent core.QObject_ITF) *QQmlPropertyMap {
	tmpValue := NewQQmlPropertyMapFromPointer(C.QQmlPropertyMap_NewQQmlPropertyMap(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QQmlPropertyMap_Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlPropertyMap_QQmlPropertyMap_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlPropertyMap) Tr(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlPropertyMap_QQmlPropertyMap_Tr(sC, cC, C.int(int32(n))))
}

func QQmlPropertyMap_TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlPropertyMap_QQmlPropertyMap_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QQmlPropertyMap) TrUtf8(s string, c string, n int) string {
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
	return cGoUnpackString(C.QQmlPropertyMap_QQmlPropertyMap_TrUtf8(sC, cC, C.int(int32(n))))
}

//export callbackQQmlPropertyMap_UpdateValue
func callbackQQmlPropertyMap_UpdateValue(ptr unsafe.Pointer, key C.struct_QtQml_PackedString, input unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "updateValue"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string, *core.QVariant) *core.QVariant)(cGoUnpackString(key), core.NewQVariantFromPointer(input)))
	}

	return core.PointerFromQVariant(NewQQmlPropertyMapFromPointer(ptr).UpdateValueDefault(cGoUnpackString(key), core.NewQVariantFromPointer(input)))
}

func (ptr *QQmlPropertyMap) ConnectUpdateValue(f func(key string, input *core.QVariant) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateValue", func(key string, input *core.QVariant) *core.QVariant {
				signal.(func(string, *core.QVariant) *core.QVariant)(key, input)
				return f(key, input)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateValue", f)
		}
	}
}

func (ptr *QQmlPropertyMap) DisconnectUpdateValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateValue")
	}
}

func (ptr *QQmlPropertyMap) UpdateValue(key string, input core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValue(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(input)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) UpdateValueDefault(key string, input core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QQmlPropertyMap_UpdateValueDefault(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(input)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) Clear(key string) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QQmlPropertyMap_Clear(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))})
	}
}

func (ptr *QQmlPropertyMap) Insert(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QQmlPropertyMap_Insert(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(value))
	}
}

//export callbackQQmlPropertyMap_ValueChanged
func callbackQQmlPropertyMap_ValueChanged(ptr unsafe.Pointer, key C.struct_QtQml_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(cGoUnpackString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QQmlPropertyMap) ConnectValueChanged(f func(key string, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QQmlPropertyMap_ConnectValueChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", func(key string, value *core.QVariant) {
				signal.(func(string, *core.QVariant))(key, value)
				f(key, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", f)
		}
	}
}

func (ptr *QQmlPropertyMap) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QQmlPropertyMap) ValueChanged(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QQmlPropertyMap_ValueChanged(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(value))
	}
}

//export callbackQQmlPropertyMap_DestroyQQmlPropertyMap
func callbackQQmlPropertyMap_DestroyQQmlPropertyMap(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlPropertyMap"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DestroyQQmlPropertyMapDefault()
	}
}

func (ptr *QQmlPropertyMap) ConnectDestroyQQmlPropertyMap(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlPropertyMap"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlPropertyMap", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlPropertyMap", f)
		}
	}
}

func (ptr *QQmlPropertyMap) DisconnectDestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlPropertyMap")
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMap() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMap(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlPropertyMap) DestroyQQmlPropertyMapDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DestroyQQmlPropertyMapDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlPropertyMap) Keys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QQmlPropertyMap_Keys(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QQmlPropertyMap) Value(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QQmlPropertyMap_Value(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) Contains(key string) bool {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		return int8(C.QQmlPropertyMap_Contains(ptr.Pointer(), C.struct_QtQml_PackedString{data: keyC, len: C.longlong(len(key))})) != 0
	}
	return false
}

func (ptr *QQmlPropertyMap) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlPropertyMap_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_MetaObject
func callbackQQmlPropertyMap_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQmlPropertyMapFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQmlPropertyMap) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlPropertyMap_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlPropertyMap) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) Size() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlPropertyMap_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QQmlPropertyMap___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QQmlPropertyMap) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QQmlPropertyMap___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QQmlPropertyMap) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList2() unsafe.Pointer {
	return C.QQmlPropertyMap___findChildren_newList2(ptr.Pointer())
}

func (ptr *QQmlPropertyMap) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList3() unsafe.Pointer {
	return C.QQmlPropertyMap___findChildren_newList3(ptr.Pointer())
}

func (ptr *QQmlPropertyMap) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlPropertyMap___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __findChildren_newList() unsafe.Pointer {
	return C.QQmlPropertyMap___findChildren_newList(ptr.Pointer())
}

func (ptr *QQmlPropertyMap) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlPropertyMap___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlPropertyMap) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QQmlPropertyMap) __children_newList() unsafe.Pointer {
	return C.QQmlPropertyMap___children_newList(ptr.Pointer())
}

//export callbackQQmlPropertyMap_Event
func callbackQQmlPropertyMap_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QQmlPropertyMap) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlPropertyMap_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_EventFilter
func callbackQQmlPropertyMap_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQQmlPropertyMapFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QQmlPropertyMap) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlPropertyMap_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQQmlPropertyMap_ChildEvent
func callbackQQmlPropertyMap_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQmlPropertyMap_ConnectNotify
func callbackQQmlPropertyMap_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_CustomEvent
func callbackQQmlPropertyMap_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQmlPropertyMap_DeleteLater
func callbackQQmlPropertyMap_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQmlPropertyMap) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlPropertyMap_Destroyed
func callbackQQmlPropertyMap_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQQmlPropertyMap_DisconnectNotify
func callbackQQmlPropertyMap_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQmlPropertyMap) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQmlPropertyMap_ObjectNameChanged
func callbackQQmlPropertyMap_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQQmlPropertyMap_TimerEvent
func callbackQQmlPropertyMap_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQmlPropertyMapFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQmlPropertyMap) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyMap_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QQmlPropertyValueSource struct {
	ptr unsafe.Pointer
}

type QQmlPropertyValueSource_ITF interface {
	QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource
}

func (ptr *QQmlPropertyValueSource) QQmlPropertyValueSource_PTR() *QQmlPropertyValueSource {
	return ptr
}

func (ptr *QQmlPropertyValueSource) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlPropertyValueSource) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlPropertyValueSource(ptr QQmlPropertyValueSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlPropertyValueSource_PTR().Pointer()
	}
	return nil
}

func NewQQmlPropertyValueSourceFromPointer(ptr unsafe.Pointer) (n *QQmlPropertyValueSource) {
	n = new(QQmlPropertyValueSource)
	n.SetPointer(ptr)
	return
}
func NewQQmlPropertyValueSource() *QQmlPropertyValueSource {
	return NewQQmlPropertyValueSourceFromPointer(C.QQmlPropertyValueSource_NewQQmlPropertyValueSource())
}

//export callbackQQmlPropertyValueSource_SetTarget
func callbackQQmlPropertyValueSource_SetTarget(ptr unsafe.Pointer, property unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setTarget"); signal != nil {
		signal.(func(*QQmlProperty))(NewQQmlPropertyFromPointer(property))
	}

}

func (ptr *QQmlPropertyValueSource) ConnectSetTarget(f func(property *QQmlProperty)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTarget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setTarget", func(property *QQmlProperty) {
				signal.(func(*QQmlProperty))(property)
				f(property)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTarget", f)
		}
	}
}

func (ptr *QQmlPropertyValueSource) DisconnectSetTarget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTarget")
	}
}

func (ptr *QQmlPropertyValueSource) SetTarget(property QQmlProperty_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_SetTarget(ptr.Pointer(), PointerFromQQmlProperty(property))
	}
}

//export callbackQQmlPropertyValueSource_DestroyQQmlPropertyValueSource
func callbackQQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlPropertyValueSource"); signal != nil {
		signal.(func())()
	} else {
		NewQQmlPropertyValueSourceFromPointer(ptr).DestroyQQmlPropertyValueSourceDefault()
	}
}

func (ptr *QQmlPropertyValueSource) ConnectDestroyQQmlPropertyValueSource(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlPropertyValueSource"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlPropertyValueSource", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlPropertyValueSource", f)
		}
	}
}

func (ptr *QQmlPropertyValueSource) DisconnectDestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlPropertyValueSource")
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSource() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlPropertyValueSource) DestroyQQmlPropertyValueSourceDefault() {
	if ptr.Pointer() != nil {
		C.QQmlPropertyValueSource_DestroyQQmlPropertyValueSourceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QQmlScriptString struct {
	ptr unsafe.Pointer
}

type QQmlScriptString_ITF interface {
	QQmlScriptString_PTR() *QQmlScriptString
}

func (ptr *QQmlScriptString) QQmlScriptString_PTR() *QQmlScriptString {
	return ptr
}

func (ptr *QQmlScriptString) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlScriptString) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlScriptString(ptr QQmlScriptString_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlScriptString_PTR().Pointer()
	}
	return nil
}

func NewQQmlScriptStringFromPointer(ptr unsafe.Pointer) (n *QQmlScriptString) {
	n = new(QQmlScriptString)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlScriptString) DestroyQQmlScriptString() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQQmlScriptString() *QQmlScriptString {
	tmpValue := NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString())
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func NewQQmlScriptString2(other QQmlScriptString_ITF) *QQmlScriptString {
	tmpValue := NewQQmlScriptStringFromPointer(C.QQmlScriptString_NewQQmlScriptString2(PointerFromQQmlScriptString(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlScriptString).DestroyQQmlScriptString)
	return tmpValue
}

func (ptr *QQmlScriptString) StringLiteral() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QQmlScriptString_StringLiteral(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlScriptString) BooleanLiteral(ok *bool) bool {
	if ptr.Pointer() != nil {
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return int8(C.QQmlScriptString_BooleanLiteral(ptr.Pointer(), &okC)) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlScriptString_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsNullLiteral() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlScriptString_IsNullLiteral(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) IsUndefinedLiteral() bool {
	if ptr.Pointer() != nil {
		return int8(C.QQmlScriptString_IsUndefinedLiteral(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlScriptString) NumberLiteral(ok *bool) float64 {
	if ptr.Pointer() != nil {
		okC := C.char(int8(qt.GoBoolToInt(*ok)))
		defer func() { *ok = int8(okC) != 0 }()
		return float64(C.QQmlScriptString_NumberLiteral(ptr.Pointer(), &okC))
	}
	return 0
}

type QV4 struct {
	ptr unsafe.Pointer
}

type QV4_ITF interface {
	QV4_PTR() *QV4
}

func (ptr *QV4) QV4_PTR() *QV4 {
	return ptr
}

func (ptr *QV4) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QV4) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQV4(ptr QV4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QV4_PTR().Pointer()
	}
	return nil
}

func NewQV4FromPointer(ptr unsafe.Pointer) (n *QV4) {
	n = new(QV4)
	n.SetPointer(ptr)
	return
}

func (ptr *QV4) DestroyQV4() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QV4__GeneratorState
//QV4::GeneratorState
type QV4__GeneratorState int64

const (
	QV4__Undefined      QV4__GeneratorState = QV4__GeneratorState(0)
	QV4__SuspendedStart QV4__GeneratorState = QV4__GeneratorState(1)
	QV4__SuspendedYield QV4__GeneratorState = QV4__GeneratorState(2)
	QV4__Executing      QV4__GeneratorState = QV4__GeneratorState(3)
	QV4__Completed      QV4__GeneratorState = QV4__GeneratorState(4)
)

//go:generate stringer -type=QV4__IteratorKind
//QV4::IteratorKind
type QV4__IteratorKind int64

const (
	QV4__KeyIteratorKind      QV4__IteratorKind = QV4__IteratorKind(0)
	QV4__ValueIteratorKind    QV4__IteratorKind = QV4__IteratorKind(1)
	QV4__KeyValueIteratorKind QV4__IteratorKind = QV4__IteratorKind(2)
)

//go:generate stringer -type=QV4__AtomicModifyOps
//QV4::AtomicModifyOps
type QV4__AtomicModifyOps int64

const (
	QV4__AtomicAdd        QV4__AtomicModifyOps = QV4__AtomicModifyOps(0)
	QV4__AtomicAnd        QV4__AtomicModifyOps = QV4__AtomicModifyOps(1)
	QV4__AtomicExchange   QV4__AtomicModifyOps = QV4__AtomicModifyOps(2)
	QV4__AtomicOr         QV4__AtomicModifyOps = QV4__AtomicModifyOps(3)
	QV4__AtomicSub        QV4__AtomicModifyOps = QV4__AtomicModifyOps(4)
	QV4__AtomicXor        QV4__AtomicModifyOps = QV4__AtomicModifyOps(5)
	QV4__NAtomicModifyOps QV4__AtomicModifyOps = QV4__AtomicModifyOps(6)
)

//go:generate stringer -type=QV4__TypedArrayType
//QV4::TypedArrayType
type QV4__TypedArrayType int64

const (
	QV4__Int8Array         QV4__TypedArrayType = QV4__TypedArrayType(0)
	QV4__UInt8Array        QV4__TypedArrayType = QV4__TypedArrayType(1)
	QV4__Int16Array        QV4__TypedArrayType = QV4__TypedArrayType(2)
	QV4__UInt16Array       QV4__TypedArrayType = QV4__TypedArrayType(3)
	QV4__Int32Array        QV4__TypedArrayType = QV4__TypedArrayType(4)
	QV4__UInt32Array       QV4__TypedArrayType = QV4__TypedArrayType(5)
	QV4__UInt8ClampedArray QV4__TypedArrayType = QV4__TypedArrayType(6)
	QV4__Float32Array      QV4__TypedArrayType = QV4__TypedArrayType(7)
	QV4__Float64Array      QV4__TypedArrayType = QV4__TypedArrayType(8)
	QV4__NTypedArrayTypes  QV4__TypedArrayType = QV4__TypedArrayType(9)
)

//go:generate stringer -type=QV4__TypeHint
//QV4::TypeHint
type QV4__TypeHint int64

const (
	QV4__PREFERREDTYPE_HINT QV4__TypeHint = QV4__TypeHint(0)
	QV4__NUMBER_HINT        QV4__TypeHint = QV4__TypeHint(1)
	QV4__STRING_HINT        QV4__TypeHint = QV4__TypeHint(2)
)
