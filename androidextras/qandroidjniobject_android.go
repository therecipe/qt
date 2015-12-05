package androidextras

//#include "androidextras_android.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QAndroidJniObject struct {
	ptr unsafe.Pointer
}

type QAndroidJniObject_ITF interface {
	QAndroidJniObject_PTR() *QAndroidJniObject
}

func (p *QAndroidJniObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAndroidJniObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QAndroidJniObject) QAndroidJniObject_PTR() *QAndroidJniObject {
	return ptr
}

func NewQAndroidJniObject() *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject())
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject2(C.CString(className)))
}

func NewQAndroidJniObject3(className string, signature string, v ...interface{}) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject3(C.CString(className), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject4(clazz))
}

func NewQAndroidJniObject5(clazz unsafe.Pointer, signature string, v ...interface{}) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject5(clazz, C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::QAndroidJniObject")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject6(object))
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodInt(ptr.Pointer(), C.CString(methodName)))
	}
	return 0
}
func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodBoolean(ptr.Pointer(), C.CString(methodName))) != 0
	}
	return false
}
func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_CallMethodVoid(ptr.Pointer(), C.CString(methodName))
	}
}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, signature string, v ...interface{}) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodInt2(ptr.Pointer(), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
	}
	return 0
}
func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, signature string, v ...interface{}) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodBoolean2(ptr.Pointer(), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))) != 0
	}
	return false
}
func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, signature string, v ...interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callMethod")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_CallMethodVoid2(ptr.Pointer(), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))
	}
}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callObjectMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod(ptr.Pointer(), C.CString(methodName)))
	}
	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callObjectMethod")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod2(ptr.Pointer(), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
	}
	return nil
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(C.CString(className), C.CString(methodName)))
}
func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(C.CString(className), C.CString(methodName))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(C.CString(className), C.CString(methodName))
}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt2(C.CString(className), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}
func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean2(C.CString(className), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid2(C.CString(className), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(clazz, C.CString(methodName)))
}
func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(clazz, C.CString(methodName))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(clazz, C.CString(methodName))
}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt4(clazz, C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}
func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean4(clazz, C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticMethod")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid4(clazz, C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...))
}

func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticObjectMethod")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(C.CString(className), C.CString(methodName)))
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticObjectMethod")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod2(C.CString(className), C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticObjectMethod")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(clazz, C.CString(methodName)))
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::callStaticObjectMethod")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod4(clazz, C.CString(methodName), C.CString(signature), assertion(0, v...), assertion(1, v...), assertion(2, v...), assertion(3, v...), assertion(4, v...), assertion(5, v...), assertion(6, v...), assertion(7, v...), assertion(8, v...), assertion(9, v...)))
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::fromString")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.CString(stri)))
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getField")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldInt(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}
func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getField")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldBoolean(ptr.Pointer(), C.CString(fieldName))) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getObjectField")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField(ptr.Pointer(), C.CString(fieldName)))
	}
	return nil
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getObjectField")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2(ptr.Pointer(), C.CString(fieldName), C.CString(signature)))
	}
	return nil
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticField")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(C.CString(className), C.CString(fieldName)))
}
func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticField")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(C.CString(className), C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticField")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(clazz, C.CString(fieldName)))
}
func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticField")
		}
	}()

	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(clazz, C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticObjectField")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(C.CString(className), C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticObjectField")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(C.CString(className), C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticObjectField")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(clazz, C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::getStaticObjectField")
		}
	}()

	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(clazz, C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_IsClassAvailable(className string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::isClassAvailable")
		}
	}()

	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(C.CString(className)) != 0
}

func (ptr *QAndroidJniObject) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAndroidJniObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAndroidJniObject_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::object")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAndroidJniObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setField")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_SetField(ptr.Pointer(), C.CString(fieldName), assertion(0, value))
	}
}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setField")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_SetField2(ptr.Pointer(), C.CString(fieldName), C.CString(signature), assertion(0, value))
	}
}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt2(C.CString(className), C.CString(fieldName), C.int(value))
}
func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean2(C.CString(className), C.CString(fieldName), C.int(qt.GoBoolToInt(value)))
}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticField(C.CString(className), C.CString(fieldName), C.CString(signature), assertion(0, value))
}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldInt4(clazz, C.CString(fieldName), C.int(value))
}
func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticFieldBoolean4(clazz, C.CString(fieldName), C.int(qt.GoBoolToInt(value)))
}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::setStaticField")
		}
	}()

	C.QAndroidJniObject_QAndroidJniObject_SetStaticField3(clazz, C.CString(fieldName), C.CString(signature), assertion(0, value))
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniObject::~QAndroidJniObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniObject_DestroyQAndroidJniObject(ptr.Pointer())
	}
}
