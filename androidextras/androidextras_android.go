package androidextras

//#include "androidextras_android.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAndroidActivityResultReceiver struct {
	ptr unsafe.Pointer
}

type QAndroidActivityResultReceiver_ITF interface {
	QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver
}

func (p *QAndroidActivityResultReceiver) QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver {
	return p
}

func (p *QAndroidActivityResultReceiver) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAndroidActivityResultReceiver) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQAndroidActivityResultReceiverFromPointer(ptr unsafe.Pointer) *QAndroidActivityResultReceiver {
	var n = NewQAndroidActivityResultReceiverFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAndroidActivityResultReceiver_") {
		n.SetObjectNameAbs("QAndroidActivityResultReceiver_" + qt.Identifier())
	}
	return n
}

//export callbackQAndroidActivityResultReceiver_HandleActivityResult
func callbackQAndroidActivityResultReceiver_HandleActivityResult(ptr unsafe.Pointer, ptrName *C.char, receiverRequestCode C.int, resultCode C.int, data unsafe.Pointer) {
	defer qt.Recovering("callback QAndroidActivityResultReceiver::handleActivityResult")

	if signal := qt.GetSignal(C.GoString(ptrName), "handleActivityResult"); signal != nil {
		signal.(func(int, int, *QAndroidJniObject))(int(receiverRequestCode), int(resultCode), NewQAndroidJniObjectFromPointer(data))
	}

}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {
	defer qt.Recovering("connect QAndroidActivityResultReceiver::handleActivityResult")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "handleActivityResult", f)
	}
}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	defer qt.Recovering("disconnect QAndroidActivityResultReceiver::handleActivityResult")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "handleActivityResult")
	}
}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {
	defer qt.Recovering("QAndroidActivityResultReceiver::handleActivityResult")

	if ptr.Pointer() != nil {
		C.QAndroidActivityResultReceiver_HandleActivityResult(ptr.Pointer(), C.int(receiverRequestCode), C.int(resultCode), PointerFromQAndroidJniObject(data))
	}
}

func (ptr *QAndroidActivityResultReceiver) ObjectNameAbs() string {
	defer qt.Recovering("QAndroidActivityResultReceiver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAndroidActivityResultReceiver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidActivityResultReceiver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAndroidActivityResultReceiver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAndroidActivityResultReceiver_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QAndroidJniEnvironment struct {
	ptr unsafe.Pointer
}

type QAndroidJniEnvironment_ITF interface {
	QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment
}

func (p *QAndroidJniEnvironment) QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment {
	return p
}

func (p *QAndroidJniEnvironment) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAndroidJniEnvironment) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQAndroidJniEnvironmentFromPointer(ptr unsafe.Pointer) *QAndroidJniEnvironment {
	var n = NewQAndroidJniEnvironmentFromPointer(ptr)
	return n
}

func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	defer qt.Recovering("QAndroidJniEnvironment::QAndroidJniEnvironment")

	return newQAndroidJniEnvironmentFromPointer(C.QAndroidJniEnvironment_NewQAndroidJniEnvironment())
}

func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {
	defer qt.Recovering("QAndroidJniEnvironment::javaVM")

	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) JavaVM() unsafe.Pointer {
	defer qt.Recovering("QAndroidJniEnvironment::javaVM")

	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {
	defer qt.Recovering("QAndroidJniEnvironment::~QAndroidJniEnvironment")

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

func (p *QAndroidJniObject) QAndroidJniObject_PTR() *QAndroidJniObject {
	return p
}

func (p *QAndroidJniObject) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAndroidJniObject) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
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

func newQAndroidJniObjectFromPointer(ptr unsafe.Pointer) *QAndroidJniObject {
	var n = NewQAndroidJniObjectFromPointer(ptr)
	return n
}

func NewQAndroidJniObject() *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject())
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject2(C.CString(className)))
}

func NewQAndroidJniObject3(className string, signature string, v ...interface{}) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

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
	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject3(C.CString(className), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject4(clazz))
}

func NewQAndroidJniObject5(clazz unsafe.Pointer, signature string, v ...interface{}) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

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
	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject5(clazz, C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::QAndroidJniObject")

	return newQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject6(object))
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {
	defer qt.Recovering("QAndroidJniObject::callMethod")

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodInt(ptr.Pointer(), C.CString(methodName)))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {
	defer qt.Recovering("QAndroidJniObject::callMethod")

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodBoolean(ptr.Pointer(), C.CString(methodName))) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {
	defer qt.Recovering("QAndroidJniObject::callMethod")

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_CallMethodVoid(ptr.Pointer(), C.CString(methodName))
	}
}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, signature string, v ...interface{}) int {
	defer qt.Recovering("QAndroidJniObject::callMethod")

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
		return int(C.QAndroidJniObject_CallMethodInt2(ptr.Pointer(), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	}
	return 0
}

func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, signature string, v ...interface{}) bool {
	defer qt.Recovering("QAndroidJniObject::callMethod")

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
		return int(C.QAndroidJniObject_CallMethodBoolean2(ptr.Pointer(), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, signature string, v ...interface{}) {
	defer qt.Recovering("QAndroidJniObject::callMethod")

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
		C.QAndroidJniObject_CallMethodVoid2(ptr.Pointer(), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callObjectMethod")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod(ptr.Pointer(), C.CString(methodName)))
	}
	return nil
}

func (ptr *QAndroidJniObject) CallMethodString(methodName string) string {
	defer qt.Recovering("QAndroidJniObject::callObjectMethod")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString(ptr.Pointer(), C.CString(methodName))).ToString()
	}
	return ""
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callObjectMethod")

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
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod2(ptr.Pointer(), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
	}
	return nil
}

func (ptr *QAndroidJniObject) CallMethodString2(methodName string, signature string, v ...interface{}) string {
	defer qt.Recovering("QAndroidJniObject::callObjectMethod")

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
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallMethodString2(ptr.Pointer(), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)).ToString()
	}
	return ""
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(C.CString(className), C.CString(methodName)))
}

func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(C.CString(className), C.CString(methodName))) != 0
}

func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(C.CString(className), C.CString(methodName))
}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(C.CString(className), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(C.CString(className), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(C.CString(className), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(clazz, C.CString(methodName)))
}

func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(clazz, C.CString(methodName))) != 0
}

func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(clazz, C.CString(methodName))
}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(clazz, C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(clazz, C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)) != 0
}

func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {
	defer qt.Recovering("QAndroidJniObject::callStaticMethod")

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
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(clazz, C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(C.CString(className), C.CString(methodName)))
}

func QAndroidJniObject_CallStaticMethodString(className string, methodName string) string {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString(C.CString(className), C.CString(methodName))).ToString()
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

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
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(C.CString(className), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func QAndroidJniObject_CallStaticMethodString2(className string, methodName string, signature string, v ...interface{}) string {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

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
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString2(C.CString(className), C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)).ToString()
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(clazz, C.CString(methodName)))
}

func QAndroidJniObject_CallStaticMethodString3(clazz unsafe.Pointer, methodName string) string {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString3(clazz, C.CString(methodName))).ToString()
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

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
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(clazz, C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9))
}

func QAndroidJniObject_CallStaticMethodString4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) string {
	defer qt.Recovering("QAndroidJniObject::callStaticObjectMethod")

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
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodString4(clazz, C.CString(methodName), C.CString(signature), p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)).ToString()
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::fromString")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.CString(stri)))
}

func (ptr *QAndroidJniObject) FromString(stri string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::fromString")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.CString(stri)))
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {
	defer qt.Recovering("QAndroidJniObject::getField")

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldInt(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}

func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {
	defer qt.Recovering("QAndroidJniObject::getField")

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldBoolean(ptr.Pointer(), C.CString(fieldName))) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getObjectField")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField(ptr.Pointer(), C.CString(fieldName)))
	}
	return nil
}

func (ptr *QAndroidJniObject) GetFieldString(fieldName string) string {
	defer qt.Recovering("QAndroidJniObject::getObjectField")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString(ptr.Pointer(), C.CString(fieldName))).ToString()
	}
	return ""
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getObjectField")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2(ptr.Pointer(), C.CString(fieldName), C.CString(signature)))
	}
	return nil
}

func (ptr *QAndroidJniObject) GetFieldString2(fieldName string, signature string) string {
	defer qt.Recovering("QAndroidJniObject::getObjectField")

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetFieldString2(ptr.Pointer(), C.CString(fieldName), C.CString(signature))).ToString()
	}
	return ""
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {
	defer qt.Recovering("QAndroidJniObject::getStaticField")

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(C.CString(className), C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {
	defer qt.Recovering("QAndroidJniObject::getStaticField")

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(C.CString(className), C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {
	defer qt.Recovering("QAndroidJniObject::getStaticField")

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(clazz, C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {
	defer qt.Recovering("QAndroidJniObject::getStaticField")

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(clazz, C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(C.CString(className), C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticFieldString(className string, fieldName string) string {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString(C.CString(className), C.CString(fieldName))).ToString()
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(C.CString(className), C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_GetStaticFieldString2(className string, fieldName string, signature string) string {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString2(C.CString(className), C.CString(fieldName), C.CString(signature))).ToString()
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(clazz, C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticFieldString3(clazz unsafe.Pointer, fieldName string) string {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString3(clazz, C.CString(fieldName))).ToString()
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(clazz, C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_GetStaticFieldString4(clazz unsafe.Pointer, fieldName string, signature string) string {
	defer qt.Recovering("QAndroidJniObject::getStaticObjectField")

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldString4(clazz, C.CString(fieldName), C.CString(signature))).ToString()
}

func QAndroidJniObject_IsClassAvailable(className string) bool {
	defer qt.Recovering("QAndroidJniObject::isClassAvailable")

	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(C.CString(className)) != 0
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {
	defer qt.Recovering("QAndroidJniObject::isClassAvailable")

	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(C.CString(className)) != 0
}

func (ptr *QAndroidJniObject) IsValid() bool {
	defer qt.Recovering("QAndroidJniObject::isValid")

	if ptr.Pointer() != nil {
		return C.QAndroidJniObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) ToString() string {
	defer qt.Recovering("QAndroidJniObject::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAndroidJniObject_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {
	defer qt.Recovering("QAndroidJniObject::object")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAndroidJniObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setField")

	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		C.QAndroidJniObject_SetField(ptr.Pointer(), C.CString(fieldName), p0)
	}
}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setField")

	if ptr.Pointer() != nil {
		var p0, d0 = assertion(0, value)
		if d0 != nil {
			defer d0()
		}
		C.QAndroidJniObject_SetField2(ptr.Pointer(), C.CString(fieldName), C.CString(signature), p0)
	}
}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(C.CString(className), C.CString(fieldName), C.int(value))
}

func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(C.CString(className), C.CString(fieldName), C.int(qt.GoBoolToInt(value)))
}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(C.CString(className), C.CString(fieldName), C.CString(signature), p0)
}

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(C.CString(className), C.CString(fieldName), C.CString(signature), p0)
}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(clazz, C.CString(fieldName), C.int(value))
}

func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(clazz, C.CString(fieldName), C.int(qt.GoBoolToInt(value)))
}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, C.CString(fieldName), C.CString(signature), p0)
}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	defer qt.Recovering("QAndroidJniObject::setStaticField")

	var p0, d0 = assertion(0, value)
	if d0 != nil {
		defer d0()
	}
	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, C.CString(fieldName), C.CString(signature), p0)
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {
	defer qt.Recovering("QAndroidJniObject::~QAndroidJniObject")

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_DestroyQAndroidJniObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
