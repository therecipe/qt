// +build !android,!android_emulator

package androidextras

import (
	"unsafe"
)

func QAndroidJniEnvironment_ExceptionCatch() error {
	return nil
}

func (e *QAndroidJniEnvironment) ExceptionCatch() error { return nil }

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

func (ptr *QAndroidActivityResultReceiver) DestroyQAndroidActivityResultReceiver() {}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {

}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult() {

}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {

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
func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniEnvironment) JavaVM() unsafe.Pointer {

	return nil
}

func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {

	return nil
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {

}

func QAndroidJniEnvironment_ExceptionCheck() bool {

	return false
}

func (ptr *QAndroidJniEnvironment) ExceptionCheck() bool {

	return false
}

func QAndroidJniEnvironment_ExceptionDescribe() {

}

func (ptr *QAndroidJniEnvironment) ExceptionDescribe() {

}

func QAndroidJniEnvironment_ExceptionClear() {

}

func (ptr *QAndroidJniEnvironment) ExceptionClear() {

}

func QAndroidJniEnvironment_ExceptionOccurred() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniEnvironment) ExceptionOccurred() unsafe.Pointer {

	return nil
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
func QAndroidJniObject_CallStaticObjectMethod(className string, methodName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethodCaught(className string, methodName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_CallStaticMethodString(className string, methodName string) string {

	return ""
}

func QAndroidJniObject_CallStaticMethodStringCaught(className string, methodName string) (string, error) {

	return "", nil
}

func QAndroidJniObject_CallStaticObjectMethod2(className string, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod2Caught(className string, methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_CallStaticMethodString2(className string, methodName string, signature string, v ...interface{}) string {

	return ""
}

func QAndroidJniObject_CallStaticMethodString2Caught(className string, methodName string, signature string, v ...interface{}) (string, error) {

	return "", nil
}

func QAndroidJniObject_CallStaticObjectMethod3(clazz unsafe.Pointer, methodName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod3Caught(clazz unsafe.Pointer, methodName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_CallStaticMethodString3(clazz unsafe.Pointer, methodName string) string {

	return ""
}

func QAndroidJniObject_CallStaticMethodString3Caught(clazz unsafe.Pointer, methodName string) (string, error) {

	return "", nil
}

func QAndroidJniObject_CallStaticObjectMethod4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_CallStaticObjectMethod4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_CallStaticMethodString4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) string {

	return ""
}

func QAndroidJniObject_CallStaticMethodString4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (string, error) {

	return "", nil
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

func QAndroidJniObject_GetStaticObjectField(className string, fieldName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectFieldCaught(className string, fieldName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_GetStaticFieldString(className string, fieldName string) string {

	return ""
}

func QAndroidJniObject_GetStaticFieldStringCaught(className string, fieldName string) (string, error) {

	return "", nil
}

func QAndroidJniObject_GetStaticObjectField2(className string, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField2Caught(className string, fieldName string, signature string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_GetStaticFieldString2(className string, fieldName string, signature string) string {

	return ""
}

func QAndroidJniObject_GetStaticFieldString2Caught(className string, fieldName string, signature string) (string, error) {

	return "", nil
}

func QAndroidJniObject_GetStaticObjectField3(clazz unsafe.Pointer, fieldName string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField3Caught(clazz unsafe.Pointer, fieldName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_GetStaticFieldString3(clazz unsafe.Pointer, fieldName string) string {

	return ""
}

func QAndroidJniObject_GetStaticFieldString3Caught(clazz unsafe.Pointer, fieldName string) (string, error) {

	return "", nil
}

func QAndroidJniObject_GetStaticObjectField4(clazz unsafe.Pointer, fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func QAndroidJniObject_GetStaticObjectField4Caught(clazz unsafe.Pointer, fieldName string, signature string) (*QAndroidJniObject, error) {

	return nil, nil
}

func QAndroidJniObject_GetStaticFieldString4(clazz unsafe.Pointer, fieldName string, signature string) string {

	return ""
}

func QAndroidJniObject_GetStaticFieldString4Caught(clazz unsafe.Pointer, fieldName string, signature string) (string, error) {

	return "", nil
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

func QAndroidJniObject_CallStaticMethodInt(className string, methodName string) int {

	return 0
}

func QAndroidJniObject_CallStaticMethodIntCaught(className string, methodName string) (int, error) {

	return 0, nil
}

func QAndroidJniObject_CallStaticMethodBoolean(className string, methodName string) bool {

	return false
}

func QAndroidJniObject_CallStaticMethodBooleanCaught(className string, methodName string) (bool, error) {

	return false, nil
}

func QAndroidJniObject_CallStaticMethodVoid(className string, methodName string) {

}

func QAndroidJniObject_CallStaticMethodVoidCaught(className string, methodName string) error {

	return nil
}

func QAndroidJniObject_CallStaticMethodInt2(className string, methodName string, signature string, v ...interface{}) int {

	return 0
}

func QAndroidJniObject_CallStaticMethodInt2Caught(className string, methodName string, signature string, v ...interface{}) (int, error) {

	return 0, nil
}

func QAndroidJniObject_CallStaticMethodBoolean2(className string, methodName string, signature string, v ...interface{}) bool {

	return false
}

func QAndroidJniObject_CallStaticMethodBoolean2Caught(className string, methodName string, signature string, v ...interface{}) (bool, error) {

	return false, nil
}

func QAndroidJniObject_CallStaticMethodVoid2(className string, methodName string, signature string, v ...interface{}) {

}

func QAndroidJniObject_CallStaticMethodVoid2Caught(className string, methodName string, signature string, v ...interface{}) error {

	return nil
}

func QAndroidJniObject_CallStaticMethodInt3(clazz unsafe.Pointer, methodName string) int {

	return 0
}

func QAndroidJniObject_CallStaticMethodInt3Caught(clazz unsafe.Pointer, methodName string) (int, error) {

	return 0, nil
}

func QAndroidJniObject_CallStaticMethodBoolean3(clazz unsafe.Pointer, methodName string) bool {

	return false
}

func QAndroidJniObject_CallStaticMethodBoolean3Caught(clazz unsafe.Pointer, methodName string) (bool, error) {

	return false, nil
}

func QAndroidJniObject_CallStaticMethodVoid3(clazz unsafe.Pointer, methodName string) {

}

func QAndroidJniObject_CallStaticMethodVoid3Caught(clazz unsafe.Pointer, methodName string) error {

	return nil
}

func QAndroidJniObject_CallStaticMethodInt4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) int {

	return 0
}

func QAndroidJniObject_CallStaticMethodInt4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (int, error) {

	return 0, nil
}

func QAndroidJniObject_CallStaticMethodBoolean4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) bool {

	return false
}

func QAndroidJniObject_CallStaticMethodBoolean4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) (bool, error) {

	return false, nil
}

func QAndroidJniObject_CallStaticMethodVoid4(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) {

}

func QAndroidJniObject_CallStaticMethodVoid4Caught(clazz unsafe.Pointer, methodName string, signature string, v ...interface{}) error {

	return nil
}

func QAndroidJniObject_GetStaticFieldInt(className string, fieldName string) int {

	return 0
}

func QAndroidJniObject_GetStaticFieldIntCaught(className string, fieldName string) (int, error) {

	return 0, nil
}

func QAndroidJniObject_GetStaticFieldBoolean(className string, fieldName string) bool {

	return false
}

func QAndroidJniObject_GetStaticFieldBooleanCaught(className string, fieldName string) (bool, error) {

	return false, nil
}

func QAndroidJniObject_GetStaticFieldInt2(clazz unsafe.Pointer, fieldName string) int {

	return 0
}

func QAndroidJniObject_GetStaticFieldInt2Caught(clazz unsafe.Pointer, fieldName string) (int, error) {

	return 0, nil
}

func QAndroidJniObject_GetStaticFieldBoolean2(clazz unsafe.Pointer, fieldName string) bool {

	return false
}

func QAndroidJniObject_GetStaticFieldBoolean2Caught(clazz unsafe.Pointer, fieldName string) (bool, error) {

	return false, nil
}

func QAndroidJniObject_IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticFieldInt2(className string, fieldName string, value int) {

}

func QAndroidJniObject_SetStaticFieldInt2Caught(className string, fieldName string, value int) error {

	return nil
}

func QAndroidJniObject_SetStaticFieldBoolean2(className string, fieldName string, value bool) {

}

func QAndroidJniObject_SetStaticFieldBoolean2Caught(className string, fieldName string, value bool) error {

	return nil
}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticFieldInt4(clazz unsafe.Pointer, fieldName string, value int) {

}

func QAndroidJniObject_SetStaticFieldInt4Caught(clazz unsafe.Pointer, fieldName string, value int) error {

	return nil
}

func QAndroidJniObject_SetStaticFieldBoolean4(clazz unsafe.Pointer, fieldName string, value bool) {

}

func QAndroidJniObject_SetStaticFieldBoolean4Caught(clazz unsafe.Pointer, fieldName string, value bool) error {

	return nil
}

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {

}

func (ptr *QAndroidJniObject) CallObjectMethod(methodName string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethodCaught(methodName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func (ptr *QAndroidJniObject) CallMethodString(methodName string) string {

	return ""
}

func (ptr *QAndroidJniObject) CallMethodStringCaught(methodName string) (string, error) {

	return "", nil
}

func (ptr *QAndroidJniObject) CallObjectMethod2(methodName string, signature string, v ...interface{}) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) CallObjectMethod2Caught(methodName string, signature string, v ...interface{}) (*QAndroidJniObject, error) {

	return nil, nil
}

func (ptr *QAndroidJniObject) CallMethodString2(methodName string, signature string, v ...interface{}) string {

	return ""
}

func (ptr *QAndroidJniObject) CallMethodString2Caught(methodName string, signature string, v ...interface{}) (string, error) {

	return "", nil
}

func (ptr *QAndroidJniObject) GetObjectField(fieldName string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) GetObjectFieldCaught(fieldName string) (*QAndroidJniObject, error) {

	return nil, nil
}

func (ptr *QAndroidJniObject) GetFieldString(fieldName string) string {

	return ""
}

func (ptr *QAndroidJniObject) GetFieldStringCaught(fieldName string) (string, error) {

	return "", nil
}

func (ptr *QAndroidJniObject) GetObjectField2(fieldName string, signature string) *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidJniObject) GetObjectField2Caught(fieldName string, signature string) (*QAndroidJniObject, error) {

	return nil, nil
}

func (ptr *QAndroidJniObject) GetFieldString2(fieldName string, signature string) string {

	return ""
}

func (ptr *QAndroidJniObject) GetFieldString2Caught(fieldName string, signature string) (string, error) {

	return "", nil
}

func (ptr *QAndroidJniObject) ToString() string {

	return ""
}

func (ptr *QAndroidJniObject) CallMethodInt(methodName string) int {

	return 0
}

func (ptr *QAndroidJniObject) CallMethodIntCaught(methodName string) (int, error) {

	return 0, nil
}

func (ptr *QAndroidJniObject) CallMethodBoolean(methodName string) bool {

	return false
}

func (ptr *QAndroidJniObject) CallMethodBooleanCaught(methodName string) (bool, error) {

	return false, nil
}

func (ptr *QAndroidJniObject) CallMethodVoid(methodName string) {

}

func (ptr *QAndroidJniObject) CallMethodVoidCaught(methodName string) error {

	return nil
}

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, signature string, v ...interface{}) int {

	return 0
}

func (ptr *QAndroidJniObject) CallMethodInt2Caught(methodName string, signature string, v ...interface{}) (int, error) {

	return 0, nil
}

func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, signature string, v ...interface{}) bool {

	return false
}

func (ptr *QAndroidJniObject) CallMethodBoolean2Caught(methodName string, signature string, v ...interface{}) (bool, error) {

	return false, nil
}

func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, signature string, v ...interface{}) {

}

func (ptr *QAndroidJniObject) CallMethodVoid2Caught(methodName string, signature string, v ...interface{}) error {

	return nil
}

func (ptr *QAndroidJniObject) GetFieldInt(fieldName string) int {

	return 0
}

func (ptr *QAndroidJniObject) GetFieldIntCaught(fieldName string) (int, error) {

	return 0, nil
}

func (ptr *QAndroidJniObject) GetFieldBoolean(fieldName string) bool {

	return false
}

func (ptr *QAndroidJniObject) GetFieldBooleanCaught(fieldName string) (bool, error) {

	return false, nil
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniObject) IsValid() bool {

	return false
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

func NewQtAndroidFromPointer(ptr unsafe.Pointer) *QtAndroid {
	var n = new(QtAndroid)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtAndroid) DestroyQtAndroid() {}

func QtAndroid_AndroidActivity() *QAndroidJniObject {

	return nil
}

func (ptr *QtAndroid) AndroidActivity() *QAndroidJniObject {

	return nil
}

func QtAndroid_AndroidContext() *QAndroidJniObject {

	return nil
}

func (ptr *QtAndroid) AndroidContext() *QAndroidJniObject {

	return nil
}

func QtAndroid_AndroidService() *QAndroidJniObject {

	return nil
}

func (ptr *QtAndroid) AndroidService() *QAndroidJniObject {

	return nil
}

func QtAndroid_AndroidSdkVersion() int {

	return 0
}

func (ptr *QtAndroid) AndroidSdkVersion() int {

	return 0
}

func QtAndroid_HideSplashScreen() {

}

func (ptr *QtAndroid) HideSplashScreen() {

}

func QtAndroid_StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func (ptr *QtAndroid) StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func QtAndroid_StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func (ptr *QtAndroid) StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}
