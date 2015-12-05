// +build !android

package androidextras

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

	return nil
}

func NewQAndroidJniObject2(className string) *QAndroidJniObject {

	return nil
}

func NewQAndroidJniObject3(className string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func NewQAndroidJniObject4(clazz unsafe.Pointer) *QAndroidJniObject {

	return nil
}

func NewQAndroidJniObject5(clazz unsafe.Pointer, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func NewQAndroidJniObject6(object unsafe.Pointer) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {

	return 0
}
func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {

	return false
}
func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {

}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, signature string, v ...interface{}) int {

	return 0
}
func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, signature string, v ...interface{}) bool {

	return false
}
func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, signature string, v ...interface{}) {

}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {

	return 0
}
func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {

	return false
}
func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {

}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {

	return 0
}
func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {

	return false
}
func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {

}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {

	return 0
}
func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {

	return false
}
func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {

}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {

	return 0
}
func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {

	return false
}
func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {

}

func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {

	return 0
}
func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {

	return false
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {

	return 0
}
func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {

	return false
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {

	return 0
}
func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {

	return false
}

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) IsValid() bool {

	return false
}

func (ptr *QAndroidJniObject) ToString() string {

	return ""
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {

}
func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {

}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {

}
func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {

}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {

}
