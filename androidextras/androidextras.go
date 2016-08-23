// +build !android

package androidextras

//#include <stdint.h>
//#include <stdlib.h>
//#include "androidextras_android.h"
import "C"
import (
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

func (ptr *QAndroidActivityResultReceiver) DestroyQAndroidActivityResultReceiver() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQAndroidActivityResultReceiver_HandleActivityResult
func callbackQAndroidActivityResultReceiver_HandleActivityResult(ptr unsafe.Pointer, receiverRequestCode C.int, resultCode C.int, data unsafe.Pointer) {

}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {

}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {

}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {

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
func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {

	return nil
}

func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniEnvironment) JavaVM() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {

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

func (ptr *QAndroidJniObject) CallMethodString(methodName string) string {

	return ""
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) CallMethodString2(methodName string, signature string, v ...interface{}) string {

	return ""
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

func QAndroidJniObject_CallStaticMethodString(className string, methodName string) string {

	return ""
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticMethodString2(className string, methodName string, signature string, v ...interface{}) string {

	return ""
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticMethodString3(clazz unsafe.Pointer, methodName string) string {

	return ""
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticMethodString4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) string {

	return ""
}

func QAndroidJniObject_FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) FromLocalRef(localRef unsafe.Pointer) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_FromString(stri string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) FromString(stri string) *QAndroidJniObject {

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

func (ptr *QAndroidJniObject) GetFieldString(fieldName string) string {

	return ""
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) GetFieldString2(fieldName string, signature string) string {

	return ""
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

func QAndroidJniObject_GetStaticFieldString(className string, fieldName string) string {

	return ""
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticFieldString2(className string, fieldName string, signature string) string {

	return ""
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticFieldString3(clazz unsafe.Pointer, fieldName string) string {

	return ""
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticFieldString4(clazz unsafe.Pointer, fieldName string, signature string) string {

	return ""
}

func QAndroidJniObject_IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {

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

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {

}

func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {

}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {

}
