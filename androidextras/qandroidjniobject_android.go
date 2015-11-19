package androidextras

//#include "qandroidjniobject_android.h"
import "C"
import (
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
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject())
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject2(C.CString(className)))
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject4(clazz))
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_NewQAndroidJniObject6(object))
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {
	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodInt(ptr.Pointer(), C.CString(methodName)))
	}
	return 0
}
func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {
	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_CallMethodBoolean(ptr.Pointer(), C.CString(methodName))) != 0
	}
	return false
}
func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {
	if ptr.Pointer() != nil {
		C.QAndroidJniObject_CallMethodVoid(ptr.Pointer(), C.CString(methodName))
	}
}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_CallObjectMethod(ptr.Pointer(), C.CString(methodName)))
	}
	return nil
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt(C.CString(className), C.CString(methodName)))
}
func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean(C.CString(className), C.CString(methodName))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid(C.CString(className), C.CString(methodName))
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodInt3(clazz, C.CString(methodName)))
}
func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {
	return int(C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodBoolean3(clazz, C.CString(methodName))) != 0
}
func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {
	C.QAndroidJniObject_QAndroidJniObject_CallStaticMethodVoid3(clazz, C.CString(methodName))
}

func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod(C.CString(className), C.CString(methodName)))
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_CallStaticObjectMethod3(clazz, C.CString(methodName)))
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_FromString(C.CString(stri)))
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {
	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldInt(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}
func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {
	if ptr.Pointer() != nil {
		return int(C.QAndroidJniObject_GetFieldBoolean(ptr.Pointer(), C.CString(fieldName))) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField(ptr.Pointer(), C.CString(fieldName)))
	}
	return nil
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {
	if ptr.Pointer() != nil {
		return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_GetObjectField2(ptr.Pointer(), C.CString(fieldName), C.CString(signature)))
	}
	return nil
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {
	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt(C.CString(className), C.CString(fieldName)))
}
func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {
	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean(C.CString(className), C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {
	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldInt2(clazz, C.CString(fieldName)))
}
func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {
	return int(C.QAndroidJniObject_QAndroidJniObject_GetStaticFieldBoolean2(clazz, C.CString(fieldName))) != 0
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField(C.CString(className), C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField2(C.CString(className), C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField3(clazz, C.CString(fieldName)))
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(C.QAndroidJniObject_QAndroidJniObject_GetStaticObjectField4(clazz, C.CString(fieldName), C.CString(signature)))
}

func QAndroidJniObject_IsClassAvailable(className string) bool {
	return C.QAndroidJniObject_QAndroidJniObject_IsClassAvailable(C.CString(className)) != 0
}

func (ptr *QAndroidJniObject) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAndroidJniObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAndroidJniObject) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAndroidJniObject_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QAndroidJniObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {
	if ptr.Pointer() != nil {
		C.QAndroidJniObject_DestroyQAndroidJniObject(ptr.Pointer())
	}
}
