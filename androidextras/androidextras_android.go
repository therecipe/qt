// +build android

package androidextras

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "androidextras_android.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtAndroidExtras_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQAndroidActivityResultReceiverFromPointer(ptr unsafe.Pointer) *QAndroidActivityResultReceiver {
	var n = new(QAndroidActivityResultReceiver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAndroidActivityResultReceiver) DestroyQAndroidActivityResultReceiver() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAndroidActivityResultReceiver_HandleActivityResult
func callbackQAndroidActivityResultReceiver_HandleActivityResult(ptr unsafe.Pointer, receiverRequestCode C.int, resultCode C.int, data unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAndroidActivityResultReceiver::handleActivityResult"); signal != nil {
		signal.(func(int, int, *QAndroidJniObject))(int(int32(receiverRequestCode)), int(int32(resultCode)), NewQAndroidJniObjectFromPointer(data))
	}

}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAndroidActivityResultReceiver::handleActivityResult", f)
	}
}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAndroidActivityResultReceiver::handleActivityResult")
	}
}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAndroidActivityResultReceiver_HandleActivityResult(ptr.Pointer(), C.int(int32(receiverRequestCode)), C.int(int32(resultCode)), PointerFromQAndroidJniObject(data))
	}
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

func NewQAndroidJniEnvironmentFromPointer(ptr unsafe.Pointer) *QAndroidJniEnvironment {
	var n = new(QAndroidJniEnvironment)
	n.SetPointer(ptr)
	return n
}
func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	var tmpValue = NewQAndroidJniEnvironmentFromPointer(C.QAndroidJniEnvironment_NewQAndroidJniEnvironment())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniEnvironment).DestroyQAndroidJniEnvironment)
	return tmpValue
}

func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {
	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) JavaVM() unsafe.Pointer {
	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {
	if ptr.Pointer() != nil {
		C.QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(ptr.Pointer())
		ptr.SetPointer(nil)
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

func NewQAndroidJniObjectFromPointer(ptr unsafe.Pointer) *QAndroidJniObject {
	var n = new(QAndroidJniObject)
	n.SetPointer(ptr)
	return n
}
func NewQAndroidJniObject() *QAndroidJniObject {
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject())
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject2(classNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject3(className string, signature string, v ...interface{}) *QAndroidJniObject {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject3(classNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject4(clazz))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject5(clazz unsafe.Pointer, signature string, v ...interface{}) *QAndroidJniObject {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject5(clazz, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject6(object))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {
	if ptr.Pointer() != nil {
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		return int(int32(C.QAndroidJniObject_CallMethodInt(ptr.Pointer(), methodNameC)))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {
	if ptr.Pointer() != nil {
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		return int8(C.QAndroidJniObject_CallMethodBoolean(ptr.Pointer(), methodNameC)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {
	if ptr.Pointer() != nil {
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		C.QAndroidJniObject_CallMethodVoid(ptr.Pointer(), methodNameC)
	}
}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, signature string, v ...interface{}) int {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		var p1, d1 = assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		var p2, d2 = assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		var p3, d3 = assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		var p4, d4 = assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		var p5, d5 = assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		var p6, d6 = assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		var p7, d7 = assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		var p8, d8 = assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		var p9, d9 = assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		return int(int32(C.QAndroidJniObject_CallMethodInt2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, signature string, v ...interface{}) bool {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		var p1, d1 = assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		var p2, d2 = assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		var p3, d3 = assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		var p4, d4 = assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		var p5, d5 = assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		var p6, d6 = assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		var p7, d7 = assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		var p8, d8 = assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		var p9, d9 = assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		return int8(C.QAndroidJniObject_CallMethodBoolean2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, signature string, v ...interface{}) {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		var p1, d1 = assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		var p2, d2 = assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		var p3, d3 = assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		var p4, d4 = assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		var p5, d5 = assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		var p6, d6 = assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		var p7, d7 = assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		var p8, d8 = assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		var p9, d9 = assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		C.QAndroidJniObject_CallMethodVoid2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod(ptr.Pointer(), methodNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) CallMethodString(methodName string) string {
	if ptr.Pointer() != nil {
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString(ptr.Pointer(), methodNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue.ToString()
	}
	return ""
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		var p1, d1 = assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		var p2, d2 = assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		var p3, d3 = assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		var p4, d4 = assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		var p5, d5 = assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		var p6, d6 = assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		var p7, d7 = assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		var p8, d8 = assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		var p9, d9 = assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) CallMethodString2(methodName string, signature string, v ...interface{}) string {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, v...)
		if d0 != nil {
			defer d0()
		}
		var p1, d1 = assertion(1, v...)
		if d1 != nil {
			defer d1()
		}
		var p2, d2 = assertion(2, v...)
		if d2 != nil {
			defer d2()
		}
		var p3, d3 = assertion(3, v...)
		if d3 != nil {
			defer d3()
		}
		var p4, d4 = assertion(4, v...)
		if d4 != nil {
			defer d4()
		}
		var p5, d5 = assertion(5, v...)
		if d5 != nil {
			defer d5()
		}
		var p6, d6 = assertion(6, v...)
		if d6 != nil {
			defer d6()
		}
		var p7, d7 = assertion(7, v...)
		if d7 != nil {
			defer d7()
		}
		var p8, d8 = assertion(8, v...)
		if d8 != nil {
			defer d8()
		}
		var p9, d9 = assertion(9, v...)
		if d9 != nil {
			defer d9()
		}
		var methodNameC = C.CString(methodName)
		defer C.free(unsafe.Pointer(methodNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString2(ptr.Pointer(), methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue.ToString()
	}
	return ""
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(classNameC, methodNameC)))
}

func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(classNameC, methodNameC)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(classNameC, methodNameC)
}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
}

func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(clazz, methodNameC)))
}

func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(clazz, methodNameC)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(clazz, methodNameC)
}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)))
}

func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(classNameC, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticMethodString(className string, methodName string) string {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(classNameC, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticMethodString2(className string, methodName string, signature string, v ...interface{}) string {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(classNameC, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(clazz, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticMethodString3(clazz unsafe.Pointer, methodName string) string {
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(clazz, methodNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_CallStaticMethodString4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) string {
	var p0, d0 = assertion(0, v...)
	if d0 != nil {
		defer d0()
	}
	var p1, d1 = assertion(1, v...)
	if d1 != nil {
		defer d1()
	}
	var p2, d2 = assertion(2, v...)
	if d2 != nil {
		defer d2()
	}
	var p3, d3 = assertion(3, v...)
	if d3 != nil {
		defer d3()
	}
	var p4, d4 = assertion(4, v...)
	if d4 != nil {
		defer d4()
	}
	var p5, d5 = assertion(5, v...)
	if d5 != nil {
		defer d5()
	}
	var p6, d6 = assertion(6, v...)
	if d6 != nil {
		defer d6()
	}
	var p7, d7 = assertion(7, v...)
	if d7 != nil {
		defer d7()
	}
	var p8, d8 = assertion(8, v...)
	if d8 != nil {
		defer d8()
	}
	var p9, d9 = assertion(9, v...)
	if d9 != nil {
		defer d9()
	}
	var methodNameC = C.CString(methodName)
	defer C.free(unsafe.Pointer(methodNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(clazz, methodNameC, signatureC, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromLocalRef(localRef))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromLocalRef(localRef))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {
	var striC = C.CString(stri)
	defer C.free(unsafe.Pointer(striC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(striC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) FromString(stri string) *QAndroidJniObject {
	var striC = C.CString(stri)
	defer C.free(unsafe.Pointer(striC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(striC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		return int(int32(C.QAndroidJniObject_GetFieldInt(ptr.Pointer(), fieldNameC)))
	}
	return 0
}

func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		return int8(C.QAndroidJniObject_GetFieldBoolean(ptr.Pointer(), fieldNameC)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField(ptr.Pointer(), fieldNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) GetFieldString(fieldName string) string {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString(ptr.Pointer(), fieldNameC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue.ToString()
	}
	return ""
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2(ptr.Pointer(), fieldNameC, signatureC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue
	}
	return nil
}

func (ptr *QAndroidJniObject) GetFieldString2(fieldName string, signature string) string {
	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString2(ptr.Pointer(), fieldNameC, signatureC))
		runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
		return tmpValue.ToString()
	}
	return ""
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(classNameC, fieldNameC)))
}

func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(classNameC, fieldNameC)) != 0
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	return int(int32(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(clazz, fieldNameC)))
}

func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	return int8(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(clazz, fieldNameC)) != 0
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(classNameC, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticFieldString(className string, fieldName string) string {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(classNameC, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(classNameC, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticFieldString2(className string, fieldName string, signature string) string {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(classNameC, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(clazz, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticFieldString3(clazz unsafe.Pointer, fieldName string) string {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(clazz, fieldNameC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(clazz, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue
}

func QAndroidJniObject_GetStaticFieldString4(clazz unsafe.Pointer, fieldName string, signature string) string {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	var tmpValue = NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(clazz, fieldNameC, signatureC))
	runtime.SetFinalizer(tmpValue, (*QAndroidJniObject).DestroyQAndroidJniObject)
	return tmpValue.ToString()
}

func QAndroidJniObject_IsClassAvailable(className string) bool {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(classNameC) != 0
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(classNameC) != 0
}

func (ptr *QAndroidJniObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAndroidJniObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAndroidJniObject_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAndroidJniObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		C.QAndroidJniObject_SetField(ptr.Pointer(), fieldNameC, p0)
	}
}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {
	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		var signatureC = C.CString(signature)
		defer C.free(unsafe.Pointer(signatureC))
		C.QAndroidJniObject_SetField2(ptr.Pointer(), fieldNameC, signatureC, p0)
	}
}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(classNameC, fieldNameC, C.int(int32(value)))
}

func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(classNameC, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {
	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(classNameC, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {
	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var classNameC = C.CString(className)
	defer C.free(unsafe.Pointer(classNameC))
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(classNameC, fieldNameC, signatureC, p0)
}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(clazz, fieldNameC, C.int(int32(value)))
}

func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(clazz, fieldNameC, C.char(int8(qt.GoBoolToInt(value))))
}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	var fieldNameC = C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameC))
	var signatureC = C.CString(signature)
	defer C.free(unsafe.Pointer(signatureC))
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, fieldNameC, signatureC, p0)
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {
	if ptr.Pointer() != nil {
		C.QAndroidJniObject_DestroyQAndroidJniObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
