// +build android android_emulator

package androidextras

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "androidextras_android.h"
import "C"
import (
	"errors"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtAndroidExtras_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtAndroidExtras_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func QAndroidJniEnvironment_ExceptionCatch() error {
	var err error
	if QAndroidJniEnvironment_ExceptionCheck() {
		tmpExcPtr := QAndroidJniEnvironment_ExceptionOccurred()
		QAndroidJniEnvironment_ExceptionClear()
		tmpExc := NewQAndroidJniObject6(tmpExcPtr)
		err = errors.New(tmpExc.CallMethodString2("toString", "()Ljava/lang/String;"))
		tmpExc.DestroyQAndroidJniObject()
	}
	return err
}

func (e *QAndroidJniEnvironment) ExceptionCatch() error {
	return QAndroidJniEnvironment_ExceptionCatch()
}

type QAndroidActivityResultReceiver struct {
	ptr unsafe.Pointer
}

type QAndroidActivityResultReceiver_ITF interface {
	QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver
}

func (ptr *QAndroidActivityResultReceiver) QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver {
	return ptr
}

func (ptr *QAndroidActivityResultReceiver) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidActivityResultReceiver) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidActivityResultReceiver(ptr QAndroidActivityResultReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidActivityResultReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAndroidActivityResultReceiverFromPointer(ptr unsafe.Pointer) (n *QAndroidActivityResultReceiver) {
	n = new(QAndroidActivityResultReceiver)
	n.SetPointer(ptr)
	return
}

func (ptr *QAndroidActivityResultReceiver) DestroyQAndroidActivityResultReceiver() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAndroidActivityResultReceiver_HandleActivityResult
func callbackQAndroidActivityResultReceiver_HandleActivityResult(ptr unsafe.Pointer, receiverRequestCode C.int, resultCode C.int, data unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "handleActivityResult"); signal != nil {
		signal.(func(int, int, *QAndroidJniObject))(int(int32(receiverRequestCode)), int(int32(resultCode)), NewQAndroidJniObjectFromPointer(data))
	}

}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "handleActivityResult"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "handleActivityResult", func(receiverRequestCode int, resultCode int, data *QAndroidJniObject) {
				signal.(func(int, int, *QAndroidJniObject))(receiverRequestCode, resultCode, data)
				f(receiverRequestCode, resultCode, data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "handleActivityResult", f)
		}
	}
}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "handleActivityResult")
	}
}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAndroidActivityResultReceiver_HandleActivityResult(ptr.Pointer(), C.int(int32(receiverRequestCode)), C.int(int32(resultCode)), PointerFromQAndroidJniObject(data))
	}
}

type QAndroidBinder struct {
	ptr unsafe.Pointer
}

type QAndroidBinder_ITF interface {
	QAndroidBinder_PTR() *QAndroidBinder
}

func (ptr *QAndroidBinder) QAndroidBinder_PTR() *QAndroidBinder {
	return ptr
}

func (ptr *QAndroidBinder) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidBinder) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidBinder(ptr QAndroidBinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidBinder_PTR().Pointer()
	}
	return nil
}

func NewQAndroidBinderFromPointer(ptr unsafe.Pointer) (n *QAndroidBinder) {
	n = new(QAndroidBinder)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAndroidBinder__CallType
//QAndroidBinder::CallType
type QAndroidBinder__CallType int64

const (
	QAndroidBinder__Normal QAndroidBinder__CallType = QAndroidBinder__CallType(0)
	QAndroidBinder__OneWay QAndroidBinder__CallType = QAndroidBinder__CallType(1)
)

func NewQAndroidBinder() *QAndroidBinder {
	return NewQAndroidBinderFromPointer(C.QAndroidBinder_NewQAndroidBinder())
}

func NewQAndroidBinder2(binder QAndroidJniObject_ITF) *QAndroidBinder {
	return NewQAndroidBinderFromPointer(C.QAndroidBinder_NewQAndroidBinder2(PointerFromQAndroidJniObject(binder)))
}

//export callbackQAndroidBinder_DestroyQAndroidBinder
func callbackQAndroidBinder_DestroyQAndroidBinder(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAndroidBinder"); signal != nil {
		signal.(func())()
	} else {
		NewQAndroidBinderFromPointer(ptr).DestroyQAndroidBinderDefault()
	}
}

func (ptr *QAndroidBinder) ConnectDestroyQAndroidBinder(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAndroidBinder"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidBinder", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidBinder", f)
		}
	}
}

func (ptr *QAndroidBinder) DisconnectDestroyQAndroidBinder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAndroidBinder")
	}
}

func (ptr *QAndroidBinder) DestroyQAndroidBinder() {
	if ptr.Pointer() != nil {
		C.QAndroidBinder_DestroyQAndroidBinder(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidBinder) DestroyQAndroidBinderDefault() {
	if ptr.Pointer() != nil {
		C.QAndroidBinder_DestroyQAndroidBinderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidBinder) Handle() *QAndroidJniObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidBinder_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidBinder) Transact(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF, flags QAndroidBinder__CallType) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAndroidBinder_Transact(ptr.Pointer(), C.int(int32(code)), PointerFromQAndroidParcel(data), PointerFromQAndroidParcel(reply), C.longlong(flags))) != 0
	}
	return false
}

type QAndroidIntent struct {
	ptr unsafe.Pointer
}

type QAndroidIntent_ITF interface {
	QAndroidIntent_PTR() *QAndroidIntent
}

func (ptr *QAndroidIntent) QAndroidIntent_PTR() *QAndroidIntent {
	return ptr
}

func (ptr *QAndroidIntent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidIntent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidIntent(ptr QAndroidIntent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidIntent_PTR().Pointer()
	}
	return nil
}

func NewQAndroidIntentFromPointer(ptr unsafe.Pointer) (n *QAndroidIntent) {
	n = new(QAndroidIntent)
	n.SetPointer(ptr)
	return
}
func NewQAndroidIntent() *QAndroidIntent {
	return NewQAndroidIntentFromPointer(C.QAndroidIntent_NewQAndroidIntent())
}

func NewQAndroidIntent2(intent QAndroidJniObject_ITF) *QAndroidIntent {
	return NewQAndroidIntentFromPointer(C.QAndroidIntent_NewQAndroidIntent2(PointerFromQAndroidJniObject(intent)))
}

func NewQAndroidIntent4(packageContext QAndroidJniObject_ITF, className string) *QAndroidIntent {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	return NewQAndroidIntentFromPointer(C.QAndroidIntent_NewQAndroidIntent4(PointerFromQAndroidJniObject(packageContext), classNameC))
}

func NewQAndroidIntent3(action string) *QAndroidIntent {
	var actionC *C.char
	if action != "" {
		actionC = C.CString(action)
		defer C.free(unsafe.Pointer(actionC))
	}
	return NewQAndroidIntentFromPointer(C.QAndroidIntent_NewQAndroidIntent3(C.struct_QtAndroidExtras_PackedString{data: actionC, len: C.longlong(len(action))}))
}

func (ptr *QAndroidIntent) ExtraBytes(key string) *core.QByteArray {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQByteArrayFromPointer(C.QAndroidIntent_ExtraBytes(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: keyC, len: C.longlong(len(key))}))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidIntent) ExtraVariant(key string) *core.QVariant {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QAndroidIntent_ExtraVariant(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: keyC, len: C.longlong(len(key))}))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidIntent) PutExtra(key string, data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QAndroidIntent_PutExtra(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQByteArray(data))
	}
}

func (ptr *QAndroidIntent) PutExtra2(key string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QAndroidIntent_PutExtra2(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(value))
	}
}

//export callbackQAndroidIntent_DestroyQAndroidIntent
func callbackQAndroidIntent_DestroyQAndroidIntent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAndroidIntent"); signal != nil {
		signal.(func())()
	} else {
		NewQAndroidIntentFromPointer(ptr).DestroyQAndroidIntentDefault()
	}
}

func (ptr *QAndroidIntent) ConnectDestroyQAndroidIntent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAndroidIntent"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidIntent", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidIntent", f)
		}
	}
}

func (ptr *QAndroidIntent) DisconnectDestroyQAndroidIntent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAndroidIntent")
	}
}

func (ptr *QAndroidIntent) DestroyQAndroidIntent() {
	if ptr.Pointer() != nil {
		C.QAndroidIntent_DestroyQAndroidIntent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidIntent) DestroyQAndroidIntentDefault() {
	if ptr.Pointer() != nil {
		C.QAndroidIntent_DestroyQAndroidIntentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidIntent) Handle() *QAndroidJniObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidIntent_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

type QAndroidJniEnvironment struct {
	ptr unsafe.Pointer
}

type QAndroidJniEnvironment_ITF interface {
	QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment
}

func (ptr *QAndroidJniEnvironment) QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment {
	return ptr
}

func (ptr *QAndroidJniEnvironment) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidJniEnvironment) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidJniEnvironment(ptr QAndroidJniEnvironment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidJniEnvironment_PTR().Pointer()
	}
	return nil
}

func NewQAndroidJniEnvironmentFromPointer(ptr unsafe.Pointer) (n *QAndroidJniEnvironment) {
	n = new(QAndroidJniEnvironment)
	n.SetPointer(ptr)
	return
}
func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	tmpValue := NewQAndroidJniEnvironmentFromPointer(C.QAndroidJniEnvironment_NewQAndroidJniEnvironment())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniEnvironment).DestroyQAndroidJniEnvironment)
	return tmpValue
}

func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {
	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) JavaVM() unsafe.Pointer {
	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) FindClass(className string) unsafe.Pointer {
	if ptr.Pointer() != nil {
		var classNameC *C.char
		if className != "" {
			classNameC = C.CString(className)
			defer C.free(unsafe.Pointer(classNameC))
		}
		return unsafe.Pointer(C.QAndroidJniEnvironment_FindClass(ptr.Pointer(), classNameC))
	}
	return nil
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {
	if ptr.Pointer() != nil {
		C.QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QAndroidJniEnvironment_ExceptionCheck() bool {
	return int8(C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionCheck()) != 0
}

func (ptr *QAndroidJniEnvironment) ExceptionCheck() bool {
	return int8(C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionCheck()) != 0
}

func QAndroidJniEnvironment_ExceptionDescribe() {
	C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionDescribe()
}

func (ptr *QAndroidJniEnvironment) ExceptionDescribe() {
	C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionDescribe()
}

func QAndroidJniEnvironment_ExceptionClear() {
	C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionClear()
}

func (ptr *QAndroidJniEnvironment) ExceptionClear() {
	C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionClear()
}

func QAndroidJniEnvironment_ExceptionOccurred() unsafe.Pointer {
	return C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionOccurred()
}

func (ptr *QAndroidJniEnvironment) ExceptionOccurred() unsafe.Pointer {
	return C.QAndroidJniEnvironment_QAndroidJniEnvironment_ExceptionOccurred()
}

type QAndroidJniExceptionCleaner struct {
	ptr unsafe.Pointer
}

type QAndroidJniExceptionCleaner_ITF interface {
	QAndroidJniExceptionCleaner_PTR() *QAndroidJniExceptionCleaner
}

func (ptr *QAndroidJniExceptionCleaner) QAndroidJniExceptionCleaner_PTR() *QAndroidJniExceptionCleaner {
	return ptr
}

func (ptr *QAndroidJniExceptionCleaner) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidJniExceptionCleaner) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidJniExceptionCleaner(ptr QAndroidJniExceptionCleaner_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidJniExceptionCleaner_PTR().Pointer()
	}
	return nil
}

func NewQAndroidJniExceptionCleanerFromPointer(ptr unsafe.Pointer) (n *QAndroidJniExceptionCleaner) {
	n = new(QAndroidJniExceptionCleaner)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QAndroidJniExceptionCleaner__OutputMode
//QAndroidJniExceptionCleaner::OutputMode
type QAndroidJniExceptionCleaner__OutputMode int64

const (
	QAndroidJniExceptionCleaner__Silent  QAndroidJniExceptionCleaner__OutputMode = QAndroidJniExceptionCleaner__OutputMode(0)
	QAndroidJniExceptionCleaner__Verbose QAndroidJniExceptionCleaner__OutputMode = QAndroidJniExceptionCleaner__OutputMode(1)
)

func NewQAndroidJniExceptionCleaner(outputMode QAndroidJniExceptionCleaner__OutputMode) *QAndroidJniExceptionCleaner {
	tmpValue := NewQAndroidJniExceptionCleanerFromPointer(C.QAndroidJniExceptionCleaner_NewQAndroidJniExceptionCleaner(C.longlong(outputMode)))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniExceptionCleaner).DestroyQAndroidJniExceptionCleaner)
	return tmpValue
}

func (ptr *QAndroidJniExceptionCleaner) Clean() {
	if ptr.Pointer() != nil {
		C.QAndroidJniExceptionCleaner_Clean(ptr.Pointer())
	}
}

func (ptr *QAndroidJniExceptionCleaner) DestroyQAndroidJniExceptionCleaner() {
	if ptr.Pointer() != nil {
		C.QAndroidJniExceptionCleaner_DestroyQAndroidJniExceptionCleaner(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

type QAndroidJniObject struct {
	ptr unsafe.Pointer
}

type QAndroidJniObject_ITF interface {
	QAndroidJniObject_PTR() *QAndroidJniObject
}

func (ptr *QAndroidJniObject) QAndroidJniObject_PTR() *QAndroidJniObject {
	return ptr
}

func (ptr *QAndroidJniObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidJniObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidJniObject(ptr QAndroidJniObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidJniObject_PTR().Pointer()
	}
	return nil
}

func NewQAndroidJniObjectFromPointer(ptr unsafe.Pointer) (n *QAndroidJniObject) {
	n = new(QAndroidJniObject)
	n.SetPointer(ptr)
	return
}
func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(classNameC, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticObjectMethodCaught(className string, methodName string) (*QAndroidJniObject, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethodCaught(classNameC, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodString(className string, methodName string) string {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(classNameC, methodNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_CallStaticMethodStringCaught(className string, methodName string) (string, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodStringCaught(classNameC, methodNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticObjectMethod2Caught(className string, methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2Caught(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodString2(className string, methodName string, signature string, v ...interface{}) string {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_CallStaticMethodString2Caught(className string, methodName string, signature string, v ...interface{}) (string, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2Caught(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(clazz, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticObjectMethod3Caught(clazz unsafe.Pointer, methodName string) (*QAndroidJniObject, error) {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3Caught(clazz, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodString3(clazz unsafe.Pointer, methodName string) string {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(clazz, methodNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_CallStaticMethodString3Caught(clazz unsafe.Pointer, methodName string) (string, error) {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3Caught(clazz, methodNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticObjectMethod4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4Caught(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodString4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) string {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_CallStaticMethodString4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (string, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4Caught(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromLocalRef(localRef))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromLocalRef(localRef))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.struct_QtAndroidExtras_PackedString{data: striC, len: C.longlong(len(stri))}))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) FromString(stri string) *QAndroidJniObject {
	var striC *C.char
	if stri != "" {
		striC = C.CString(stri)
		defer C.free(unsafe.Pointer(striC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.struct_QtAndroidExtras_PackedString{data: striC, len: C.longlong(len(stri))}))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(classNameC, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticObjectFieldCaught(className string, fieldName string) (*QAndroidJniObject, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectFieldCaught(classNameC, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldString(className string, fieldName string) string {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(classNameC, fieldNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_GetStaticFieldStringCaught(className string, fieldName string) (string, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldStringCaught(classNameC, fieldNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(classNameC, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticObjectField2Caught(className string, fieldName string, signature string) (*QAndroidJniObject, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2Caught(classNameC, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldString2(className string, fieldName string, signature string) string {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(classNameC, fieldNameC, signatureC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_GetStaticFieldString2Caught(className string, fieldName string, signature string) (string, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2Caught(classNameC, fieldNameC, signatureC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(clazz, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticObjectField3Caught(clazz unsafe.Pointer, fieldName string) (*QAndroidJniObject, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3Caught(clazz, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldString3(clazz unsafe.Pointer, fieldName string) string {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(clazz, fieldNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_GetStaticFieldString3Caught(clazz unsafe.Pointer, fieldName string) (string, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3Caught(clazz, fieldNameC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(clazz, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticObjectField4Caught(clazz unsafe.Pointer, fieldName string, signature string) (*QAndroidJniObject, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4Caught(clazz, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldString4(clazz unsafe.Pointer, fieldName string, signature string) string {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(clazz, fieldNameC, signatureC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString
}

func QAndroidJniObject_GetStaticFieldString4Caught(clazz unsafe.Pointer, fieldName string, signature string) (string, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4Caught(clazz, fieldNameC, signatureC))
	tmpValueToString := tmpValue.ToString()
	tmpValue.DestroyQAndroidJniObject()
	return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
}

func NewQAndroidJniObject() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject2(classNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject3(className string, signature string, v ...interface{}) *QAndroidJniObject {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject3(classNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject6(object))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject4(clazz))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject5(clazz unsafe.Pointer, signature string, v ...interface{}) *QAndroidJniObject {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject5(clazz, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(classNameC, methodNameC)))
}

func QAndroidJniObject_CallStaticMethodIntCaught(className string, methodName string) (int, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodIntCaught(classNameC, methodNameC))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(classNameC, methodNameC)) != 0
}

func QAndroidJniObject_CallStaticMethodBooleanCaught(className string, methodName string) (bool, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBooleanCaught(classNameC, methodNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(classNameC, methodNameC)
}

func QAndroidJniObject_CallStaticMethodVoidCaught(className string, methodName string) error {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoidCaught(classNameC, methodNameC)
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
}

func QAndroidJniObject_CallStaticMethodInt2Caught(className string, methodName string, signature string, v ...interface{}) (int, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2Caught(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodBoolean2Caught(className string, methodName string, signature string, v ...interface{}) (bool, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2Caught(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticMethodVoid2Caught(className string, methodName string, signature string, v ...interface{}) error {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2Caught(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(clazz, methodNameC)))
}

func QAndroidJniObject_CallStaticMethodInt3Caught(clazz unsafe.Pointer, methodName string) (int, error) {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3Caught(clazz, methodNameC))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(clazz, methodNameC)) != 0
}

func QAndroidJniObject_CallStaticMethodBoolean3Caught(clazz unsafe.Pointer, methodName string) (bool, error) {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3Caught(clazz, methodNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(clazz, methodNameC)
}

func QAndroidJniObject_CallStaticMethodVoid3Caught(clazz unsafe.Pointer, methodName string) error {
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3Caught(clazz, methodNameC)
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
}

func QAndroidJniObject_CallStaticMethodInt4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (int, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4Caught(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodBoolean4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (bool, error) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4Caught(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticMethodVoid4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) error {
	p0, d0 := assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	p1, d1 := assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	p2, d2 := assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	p3, d3 := assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	p4, d4 := assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	p5, d5 := assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	p6, d6 := assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	p7, d7 := assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	p8, d8 := assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	p9, d9 := assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC *C.char
	if methodName != "" {
		methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4Caught(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(classNameC, fieldNameC)))
}

func QAndroidJniObject_GetStaticFieldIntCaught(className string, fieldName string) (int, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldIntCaught(classNameC, fieldNameC))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(classNameC, fieldNameC)) != 0
}

func QAndroidJniObject_GetStaticFieldBooleanCaught(className string, fieldName string) (bool, error) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBooleanCaught(classNameC, fieldNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(clazz, fieldNameC)))
}

func QAndroidJniObject_GetStaticFieldInt2Caught(clazz unsafe.Pointer, fieldName string) (int, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2Caught(clazz, fieldNameC))), QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(clazz, fieldNameC)) != 0
}

func QAndroidJniObject_GetStaticFieldBoolean2Caught(clazz unsafe.Pointer, fieldName string) (bool, error) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2Caught(clazz, fieldNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_IsClassAvailable(className string) bool {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(classNameC)) != 0
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	return int8(C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(classNameC)) != 0
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		C.QAndroidJniObject_SetField(ptr.Pointer(), fieldNameC, p0)
	}
}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		C.QAndroidJniObject_SetField2(ptr.Pointer(), fieldNameC, signatureC, p0)
	}
}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(classNameC, fieldNameC, C.int(int32(value)))
}

func QAndroidJniObject_SetStaticFieldInt2Caught(className string, fieldName string, value int) error {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2Caught(classNameC, fieldNameC, C.int(int32(value)))
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(classNameC, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
}

func QAndroidJniObject_SetStaticFieldBoolean2Caught(className string, fieldName string, value bool) error {
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2Caught(classNameC, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {
	p0, d0 := assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(classNameC, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {
	p0, d0 := assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var classNameC *C.char
	if className != "" {
		classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(classNameC, fieldNameC, signatureC, p0)
}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(clazz, fieldNameC, C.int(int32(value)))
}

func QAndroidJniObject_SetStaticFieldInt4Caught(clazz unsafe.Pointer, fieldName string, value int) error {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4Caught(clazz, fieldNameC, C.int(int32(value)))
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(clazz, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
}

func QAndroidJniObject_SetStaticFieldBoolean4Caught(clazz unsafe.Pointer, fieldName string, value bool) error {
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4Caught(clazz, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
	return QAndroidJniEnvironment_ExceptionCatch()
}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	p0, d0 := assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	p0, d0 := assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var fieldNameC *C.char
	if fieldName != "" {
		fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
	}
	var signatureC *C.char
	if signature != "" {
		signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {
	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.QAndroidJniObject_DestroyQAndroidJniObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod(ptr.Pointer(), methodNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethodCaught(methodName string) (*QAndroidJniObject, error) {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethodCaught(ptr.Pointer(), methodNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
	}
	return nil, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodString(methodName string) string {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString(ptr.Pointer(), methodNameC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString
	}
	return ""
}

func (ptr *QAndroidJniObject) CallMethodStringCaught(methodName string) (string, error) {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodStringCaught(ptr.Pointer(), methodNameC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
	}
	return "", errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethod2Caught(methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod2Caught(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
	}
	return nil, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodString2(methodName string, signature string, v ...interface{}) string {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString
	}
	return ""
}

func (ptr *QAndroidJniObject) CallMethodString2Caught(methodName string, signature string, v ...interface{}) (string, error) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString2Caught(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
	}
	return "", errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField(ptr.Pointer(), fieldNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) GetObjectFieldCaught(fieldName string) (*QAndroidJniObject, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectFieldCaught(ptr.Pointer(), fieldNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
	}
	return nil, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetFieldString(fieldName string) string {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString(ptr.Pointer(), fieldNameC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString
	}
	return ""
}

func (ptr *QAndroidJniObject) GetFieldStringCaught(fieldName string) (string, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldStringCaught(ptr.Pointer(), fieldNameC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
	}
	return "", errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2(ptr.Pointer(), fieldNameC, signatureC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) GetObjectField2Caught(fieldName string, signature string) (*QAndroidJniObject, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2Caught(ptr.Pointer(), fieldNameC, signatureC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue, QAndroidJniEnvironment_ExceptionCatch()
	}
	return nil, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetFieldString2(fieldName string, signature string) string {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString2(ptr.Pointer(), fieldNameC, signatureC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString
	}
	return ""
}

func (ptr *QAndroidJniObject) GetFieldString2Caught(fieldName string, signature string) (string, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		var signatureC *C.char
		if signature != "" {
			signatureC = C.CString(signature)
			defer C.free(unsafe.Pointer(signatureC))
		}
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString2Caught(ptr.Pointer(), fieldNameC, signatureC))
		tmpValueToString := tmpValue.ToString()
		tmpValue.DestroyQAndroidJniObject()
		return tmpValueToString, QAndroidJniEnvironment_ExceptionCatch()
	}
	return "", errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAndroidJniObject_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		return int(int32(C.QAndroidJniObject_CallMethodInt(ptr.Pointer(), methodNameC)))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodIntCaught(methodName string) (int, error) {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		return int(int32(C.QAndroidJniObject_CallMethodIntCaught(ptr.Pointer(), methodNameC))), QAndroidJniEnvironment_ExceptionCatch()
	}
	return 0, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		return int8(C.QAndroidJniObject_CallMethodBoolean(ptr.Pointer(), methodNameC)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodBooleanCaught(methodName string) (bool, error) {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		return int8(C.QAndroidJniObject_CallMethodBooleanCaught(ptr.Pointer(), methodNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
	}
	return false, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		C.QAndroidJniObject_CallMethodVoid(ptr.Pointer(), methodNameC)
	}
}

func (ptr *QAndroidJniObject) CallMethodVoidCaught(methodName string) error {
	if ptr.Pointer() != nil {
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		C.QAndroidJniObject_CallMethodVoidCaught(ptr.Pointer(), methodNameC)
		return QAndroidJniEnvironment_ExceptionCatch()
	}
	return errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, sig string, v ...interface{}) int {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		return int(int32(C.QAndroidJniObject_CallMethodInt2(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodInt2Caught(methodName string, sig string, v ...interface{}) (int, error) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		return int(int32(C.QAndroidJniObject_CallMethodInt2Caught(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))), QAndroidJniEnvironment_ExceptionCatch()
	}
	return 0, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, sig string, v ...interface{}) bool {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		return int8(C.QAndroidJniObject_CallMethodBoolean2(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodBoolean2Caught(methodName string, sig string, v ...interface{}) (bool, error) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		return int8(C.QAndroidJniObject_CallMethodBoolean2Caught(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0, QAndroidJniEnvironment_ExceptionCatch()
	}
	return false, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, sig string, v ...interface{}) {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		C.QAndroidJniObject_CallMethodVoid2(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (ptr *QAndroidJniObject) CallMethodVoid2Caught(methodName string, sig string, v ...interface{}) error {
	if ptr.Pointer() != nil {
		p0, d0 := assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		p1, d1 := assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		p2, d2 := assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		p3, d3 := assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		p4, d4 := assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		p5, d5 := assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		p6, d6 := assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		p7, d7 := assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		p8, d8 := assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		p9, d9 := assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC *C.char
		if methodName != "" {
			methodNameC = C.CString(methodName)
			defer C.free(unsafe.Pointer(methodNameC))
		}
		var sigC *C.char
		if sig != "" {
			sigC = C.CString(sig)
			defer C.free(unsafe.Pointer(sigC))
		}
		C.QAndroidJniObject_CallMethodVoid2Caught(ptr.Pointer(), methodNameC, sigC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		return QAndroidJniEnvironment_ExceptionCatch()
	}
	return errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		return int(int32(C.QAndroidJniObject_GetFieldInt(ptr.Pointer(), fieldNameC)))
	}
	return 0
}

func (ptr *QAndroidJniObject) GetFieldIntCaught(fieldName string) (int, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		return int(int32(C.QAndroidJniObject_GetFieldIntCaught(ptr.Pointer(), fieldNameC))), QAndroidJniEnvironment_ExceptionCatch()
	}
	return 0, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		return int8(C.QAndroidJniObject_GetFieldBoolean(ptr.Pointer(), fieldNameC)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) GetFieldBooleanCaught(fieldName string) (bool, error) {
	if ptr.Pointer() != nil {
		var fieldNameC *C.char
		if fieldName != "" {
			fieldNameC = C.CString(fieldName)
			defer C.free(unsafe.Pointer(fieldNameC))
		}
		return int8(C.QAndroidJniObject_GetFieldBooleanCaught(ptr.Pointer(), fieldNameC)) != 0, QAndroidJniEnvironment_ExceptionCatch()
	}
	return false, errors.New("*.Pointer() == nil")
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAndroidJniObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAndroidJniObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QAndroidJniObject_IsValid(ptr.Pointer())) != 0
	}
	return false
}

type QAndroidParcel struct {
	ptr unsafe.Pointer
}

type QAndroidParcel_ITF interface {
	QAndroidParcel_PTR() *QAndroidParcel
}

func (ptr *QAndroidParcel) QAndroidParcel_PTR() *QAndroidParcel {
	return ptr
}

func (ptr *QAndroidParcel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidParcel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidParcel(ptr QAndroidParcel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidParcel_PTR().Pointer()
	}
	return nil
}

func NewQAndroidParcelFromPointer(ptr unsafe.Pointer) (n *QAndroidParcel) {
	n = new(QAndroidParcel)
	n.SetPointer(ptr)
	return
}
func NewQAndroidParcel() *QAndroidParcel {
	return NewQAndroidParcelFromPointer(C.QAndroidParcel_NewQAndroidParcel())
}

func NewQAndroidParcel2(parcel QAndroidJniObject_ITF) *QAndroidParcel {
	return NewQAndroidParcelFromPointer(C.QAndroidParcel_NewQAndroidParcel2(PointerFromQAndroidJniObject(parcel)))
}

//export callbackQAndroidParcel_DestroyQAndroidParcel
func callbackQAndroidParcel_DestroyQAndroidParcel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAndroidParcel"); signal != nil {
		signal.(func())()
	} else {
		NewQAndroidParcelFromPointer(ptr).DestroyQAndroidParcelDefault()
	}
}

func (ptr *QAndroidParcel) ConnectDestroyQAndroidParcel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAndroidParcel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidParcel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidParcel", f)
		}
	}
}

func (ptr *QAndroidParcel) DisconnectDestroyQAndroidParcel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAndroidParcel")
	}
}

func (ptr *QAndroidParcel) DestroyQAndroidParcel() {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_DestroyQAndroidParcel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidParcel) DestroyQAndroidParcelDefault() {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_DestroyQAndroidParcelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidParcel) ReadBinder() *QAndroidBinder {
	if ptr.Pointer() != nil {
		tmpValue := NewQAndroidBinderFromPointer(C.QAndroidParcel_ReadBinder(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QAndroidBinder).DestroyQAndroidBinder)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidParcel) Handle() *QAndroidJniObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidParcel_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidParcel) ReadData() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QAndroidParcel_ReadData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidParcel) ReadVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QAndroidParcel_ReadVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidParcel) ReadFileDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAndroidParcel_ReadFileDescriptor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAndroidParcel) WriteBinder(binder QAndroidBinder_ITF) {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_WriteBinder(ptr.Pointer(), PointerFromQAndroidBinder(binder))
	}
}

func (ptr *QAndroidParcel) WriteData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_WriteData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

func (ptr *QAndroidParcel) WriteFileDescriptor(fd int) {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_WriteFileDescriptor(ptr.Pointer(), C.int(int32(fd)))
	}
}

func (ptr *QAndroidParcel) WriteVariant(value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAndroidParcel_WriteVariant(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

type QAndroidService struct {
	ptr unsafe.Pointer
}

type QAndroidService_ITF interface {
	QAndroidService_PTR() *QAndroidService
}

func (ptr *QAndroidService) QAndroidService_PTR() *QAndroidService {
	return ptr
}

func (ptr *QAndroidService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidService(ptr QAndroidService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidService_PTR().Pointer()
	}
	return nil
}

func NewQAndroidServiceFromPointer(ptr unsafe.Pointer) (n *QAndroidService) {
	n = new(QAndroidService)
	n.SetPointer(ptr)
	return
}

//export callbackQAndroidService_OnBind
func callbackQAndroidService_OnBind(ptr unsafe.Pointer, intent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "onBind"); signal != nil {
		return PointerFromQAndroidBinder(signal.(func(*QAndroidIntent) *QAndroidBinder)(NewQAndroidIntentFromPointer(intent)))
	}

	return PointerFromQAndroidBinder(NewQAndroidServiceFromPointer(ptr).OnBindDefault(NewQAndroidIntentFromPointer(intent)))
}

func (ptr *QAndroidService) ConnectOnBind(f func(intent *QAndroidIntent) *QAndroidBinder) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "onBind"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onBind", func(intent *QAndroidIntent) *QAndroidBinder {
				signal.(func(*QAndroidIntent) *QAndroidBinder)(intent)
				return f(intent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onBind", f)
		}
	}
}

func (ptr *QAndroidService) DisconnectOnBind() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "onBind")
	}
}

func (ptr *QAndroidService) OnBind(intent QAndroidIntent_ITF) *QAndroidBinder {
	if ptr.Pointer() != nil {
		return NewQAndroidBinderFromPointer(C.QAndroidService_OnBind(ptr.Pointer(), PointerFromQAndroidIntent(intent)))
	}
	return nil
}

func (ptr *QAndroidService) OnBindDefault(intent QAndroidIntent_ITF) *QAndroidBinder {
	if ptr.Pointer() != nil {
		return NewQAndroidBinderFromPointer(C.QAndroidService_OnBindDefault(ptr.Pointer(), PointerFromQAndroidIntent(intent)))
	}
	return nil
}

func NewQAndroidService(argc int, argv []string) *QAndroidService {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return NewQAndroidServiceFromPointer(C.QAndroidService_NewQAndroidService(C.int(int32(argc)), argvC))
}

//export callbackQAndroidService_DestroyQAndroidService
func callbackQAndroidService_DestroyQAndroidService(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAndroidService"); signal != nil {
		signal.(func())()
	} else {
		NewQAndroidServiceFromPointer(ptr).DestroyQAndroidServiceDefault()
	}
}

func (ptr *QAndroidService) ConnectDestroyQAndroidService(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAndroidService"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidService", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidService", f)
		}
	}
}

func (ptr *QAndroidService) DisconnectDestroyQAndroidService() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAndroidService")
	}
}

func (ptr *QAndroidService) DestroyQAndroidService() {
	if ptr.Pointer() != nil {
		C.QAndroidService_DestroyQAndroidService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAndroidService) DestroyQAndroidServiceDefault() {
	if ptr.Pointer() != nil {
		C.QAndroidService_DestroyQAndroidServiceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QAndroidServiceConnection struct {
	ptr unsafe.Pointer
}

type QAndroidServiceConnection_ITF interface {
	QAndroidServiceConnection_PTR() *QAndroidServiceConnection
}

func (ptr *QAndroidServiceConnection) QAndroidServiceConnection_PTR() *QAndroidServiceConnection {
	return ptr
}

func (ptr *QAndroidServiceConnection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAndroidServiceConnection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQAndroidServiceConnection(ptr QAndroidServiceConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidServiceConnection_PTR().Pointer()
	}
	return nil
}

func NewQAndroidServiceConnectionFromPointer(ptr unsafe.Pointer) (n *QAndroidServiceConnection) {
	n = new(QAndroidServiceConnection)
	n.SetPointer(ptr)
	return
}
func NewQAndroidServiceConnection() *QAndroidServiceConnection {
	return NewQAndroidServiceConnectionFromPointer(C.QAndroidServiceConnection_NewQAndroidServiceConnection())
}

func NewQAndroidServiceConnection2(serviceConnection QAndroidJniObject_ITF) *QAndroidServiceConnection {
	return NewQAndroidServiceConnectionFromPointer(C.QAndroidServiceConnection_NewQAndroidServiceConnection2(PointerFromQAndroidJniObject(serviceConnection)))
}

//export callbackQAndroidServiceConnection_OnServiceConnected
func callbackQAndroidServiceConnection_OnServiceConnected(ptr unsafe.Pointer, name C.struct_QtAndroidExtras_PackedString, serviceBinder unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "onServiceConnected"); signal != nil {
		signal.(func(string, *QAndroidBinder))(cGoUnpackString(name), NewQAndroidBinderFromPointer(serviceBinder))
	}

}

func (ptr *QAndroidServiceConnection) ConnectOnServiceConnected(f func(name string, serviceBinder *QAndroidBinder)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "onServiceConnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onServiceConnected", func(name string, serviceBinder *QAndroidBinder) {
				signal.(func(string, *QAndroidBinder))(name, serviceBinder)
				f(name, serviceBinder)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onServiceConnected", f)
		}
	}
}

func (ptr *QAndroidServiceConnection) DisconnectOnServiceConnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "onServiceConnected")
	}
}

func (ptr *QAndroidServiceConnection) OnServiceConnected(name string, serviceBinder QAndroidBinder_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAndroidServiceConnection_OnServiceConnected(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: nameC, len: C.longlong(len(name))}, PointerFromQAndroidBinder(serviceBinder))
	}
}

//export callbackQAndroidServiceConnection_OnServiceDisconnected
func callbackQAndroidServiceConnection_OnServiceDisconnected(ptr unsafe.Pointer, name C.struct_QtAndroidExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "onServiceDisconnected"); signal != nil {
		signal.(func(string))(cGoUnpackString(name))
	}

}

func (ptr *QAndroidServiceConnection) ConnectOnServiceDisconnected(f func(name string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "onServiceDisconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "onServiceDisconnected", func(name string) {
				signal.(func(string))(name)
				f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "onServiceDisconnected", f)
		}
	}
}

func (ptr *QAndroidServiceConnection) DisconnectOnServiceDisconnected() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "onServiceDisconnected")
	}
}

func (ptr *QAndroidServiceConnection) OnServiceDisconnected(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QAndroidServiceConnection_OnServiceDisconnected(ptr.Pointer(), C.struct_QtAndroidExtras_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQAndroidServiceConnection_DestroyQAndroidServiceConnection
func callbackQAndroidServiceConnection_DestroyQAndroidServiceConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAndroidServiceConnection"); signal != nil {
		signal.(func())()
	} else {
		NewQAndroidServiceConnectionFromPointer(ptr).DestroyQAndroidServiceConnectionDefault()
	}
}

func (ptr *QAndroidServiceConnection) ConnectDestroyQAndroidServiceConnection(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAndroidServiceConnection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidServiceConnection", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAndroidServiceConnection", f)
		}
	}
}

func (ptr *QAndroidServiceConnection) DisconnectDestroyQAndroidServiceConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAndroidServiceConnection")
	}
}

func (ptr *QAndroidServiceConnection) DestroyQAndroidServiceConnection() {
	if ptr.Pointer() != nil {
		C.QAndroidServiceConnection_DestroyQAndroidServiceConnection(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidServiceConnection) DestroyQAndroidServiceConnectionDefault() {
	if ptr.Pointer() != nil {
		C.QAndroidServiceConnection_DestroyQAndroidServiceConnectionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAndroidServiceConnection) Handle() *QAndroidJniObject {
	if ptr.Pointer() != nil {
		tmpValue := NewQAndroidJniObjectFromPointer(C.QAndroidServiceConnection_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

type QtAndroid struct {
	ptr unsafe.Pointer
}

type QtAndroid_ITF interface {
	QtAndroid_PTR() *QtAndroid
}

func (ptr *QtAndroid) QtAndroid_PTR() *QtAndroid {
	return ptr
}

func (ptr *QtAndroid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtAndroid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtAndroid(ptr QtAndroid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtAndroid_PTR().Pointer()
	}
	return nil
}

func NewQtAndroidFromPointer(ptr unsafe.Pointer) (n *QtAndroid) {
	n = new(QtAndroid)
	n.SetPointer(ptr)
	return
}

func (ptr *QtAndroid) DestroyQtAndroid() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QtAndroid__BindFlag
//QtAndroid::BindFlag
type QtAndroid__BindFlag int64

const (
	QtAndroid__None               QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000000)
	QtAndroid__AutoCreate         QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000001)
	QtAndroid__DebugUnbind        QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000002)
	QtAndroid__NotForeground      QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000004)
	QtAndroid__AboveClient        QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000008)
	QtAndroid__AllowOomManagement QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000010)
	QtAndroid__WaivePriority      QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000020)
	QtAndroid__Important          QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000040)
	QtAndroid__AdjustWithActivity QtAndroid__BindFlag = QtAndroid__BindFlag(0x00000080)
	QtAndroid__ExternalService    QtAndroid__BindFlag = QtAndroid__BindFlag(-2147483648)
)

//go:generate stringer -type=QtAndroid__PermissionResult
//QtAndroid::PermissionResult
type QtAndroid__PermissionResult int64

const (
	QtAndroid__Granted QtAndroid__PermissionResult = QtAndroid__PermissionResult(0)
	QtAndroid__Denied  QtAndroid__PermissionResult = QtAndroid__PermissionResult(1)
)

func QtAndroid_AndroidActivity() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidActivity())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QtAndroid) AndroidActivity() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidActivity())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QtAndroid_AndroidContext() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidContext())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QtAndroid) AndroidContext() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidContext())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QtAndroid_AndroidService() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidService())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QtAndroid) AndroidService() *QAndroidJniObject {
	tmpValue := NewQAndroidJniObjectFromPointer(C.QtAndroid_QtAndroid_AndroidService())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QtAndroid_BindService(serviceIntent QAndroidIntent_ITF, serviceConnection QAndroidServiceConnection_ITF, flags QtAndroid__BindFlag) bool {
	return int8(C.QtAndroid_QtAndroid_BindService(PointerFromQAndroidIntent(serviceIntent), PointerFromQAndroidServiceConnection(serviceConnection), C.longlong(flags))) != 0
}

func (ptr *QtAndroid) BindService(serviceIntent QAndroidIntent_ITF, serviceConnection QAndroidServiceConnection_ITF, flags QtAndroid__BindFlag) bool {
	return int8(C.QtAndroid_QtAndroid_BindService(PointerFromQAndroidIntent(serviceIntent), PointerFromQAndroidServiceConnection(serviceConnection), C.longlong(flags))) != 0
}

func QtAndroid_ShouldShowRequestPermissionRationale(permission string) bool {
	var permissionC *C.char
	if permission != "" {
		permissionC = C.CString(permission)
		defer C.free(unsafe.Pointer(permissionC))
	}
	return int8(C.QtAndroid_QtAndroid_ShouldShowRequestPermissionRationale(C.struct_QtAndroidExtras_PackedString{data: permissionC, len: C.longlong(len(permission))})) != 0
}

func (ptr *QtAndroid) ShouldShowRequestPermissionRationale(permission string) bool {
	var permissionC *C.char
	if permission != "" {
		permissionC = C.CString(permission)
		defer C.free(unsafe.Pointer(permissionC))
	}
	return int8(C.QtAndroid_QtAndroid_ShouldShowRequestPermissionRationale(C.struct_QtAndroidExtras_PackedString{data: permissionC, len: C.longlong(len(permission))})) != 0
}

func QtAndroid_AndroidSdkVersion() int {
	return int(int32(C.QtAndroid_QtAndroid_AndroidSdkVersion()))
}

func (ptr *QtAndroid) AndroidSdkVersion() int {
	return int(int32(C.QtAndroid_QtAndroid_AndroidSdkVersion()))
}

func QtAndroid_HideSplashScreen() {
	C.QtAndroid_QtAndroid_HideSplashScreen()
}

func (ptr *QtAndroid) HideSplashScreen() {
	C.QtAndroid_QtAndroid_HideSplashScreen()
}

func QtAndroid_HideSplashScreen2(duration int) {
	C.QtAndroid_QtAndroid_HideSplashScreen2(C.int(int32(duration)))
}

func (ptr *QtAndroid) HideSplashScreen2(duration int) {
	C.QtAndroid_QtAndroid_HideSplashScreen2(C.int(int32(duration)))
}

func QtAndroid_StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {
	C.QtAndroid_QtAndroid_StartActivity(PointerFromQAndroidJniObject(intent), C.int(int32(receiverRequestCode)), PointerFromQAndroidActivityResultReceiver(resultReceiver))
}

func (ptr *QtAndroid) StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {
	C.QtAndroid_QtAndroid_StartActivity(PointerFromQAndroidJniObject(intent), C.int(int32(receiverRequestCode)), PointerFromQAndroidActivityResultReceiver(resultReceiver))
}

func QtAndroid_StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {
	C.QtAndroid_QtAndroid_StartIntentSender(PointerFromQAndroidJniObject(intentSender), C.int(int32(receiverRequestCode)), PointerFromQAndroidActivityResultReceiver(resultReceiver))
}

func (ptr *QtAndroid) StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {
	C.QtAndroid_QtAndroid_StartIntentSender(PointerFromQAndroidJniObject(intentSender), C.int(int32(receiverRequestCode)), PointerFromQAndroidActivityResultReceiver(resultReceiver))
}
