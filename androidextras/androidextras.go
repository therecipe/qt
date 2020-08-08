// +build !android,!android_emulator

package androidextras

import (
	"github.com/StarAurryon/qt"
	"github.com/StarAurryon/qt/core"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}
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

func NewQAndroidActivityResultReceiverFromPointer(ptr unsafe.Pointer) (n *QAndroidActivityResultReceiver) {
	n = new(QAndroidActivityResultReceiver)
	n.SetPointer(ptr)
	return
}

func (ptr *QAndroidActivityResultReceiver) DestroyQAndroidActivityResultReceiver() {
}

func (ptr *QAndroidActivityResultReceiver) ConnectHandleActivityResult(f func(receiverRequestCode int, resultCode int, data *QAndroidJniObject)) {

}

func (ptr *QAndroidActivityResultReceiver) DisconnectHandleActivityResult() {

}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {

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

func (ptr *QAndroidBinder) DestroyQAndroidBinder() {
}

//go:generate stringer -type=QAndroidBinder__CallType
//QAndroidBinder::CallType
type QAndroidBinder__CallType int64

const (
	QAndroidBinder__Normal QAndroidBinder__CallType = QAndroidBinder__CallType(0)
	QAndroidBinder__OneWay QAndroidBinder__CallType = QAndroidBinder__CallType(1)
)

func NewQAndroidBinder() *QAndroidBinder {

	return nil
}

func NewQAndroidBinder2(binder QAndroidJniObject_ITF) *QAndroidBinder {

	return nil
}

func (ptr *QAndroidBinder) Handle() *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidBinder) Transact(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF, flags QAndroidBinder__CallType) bool {

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

func (ptr *QAndroidIntent) DestroyQAndroidIntent() {
}

func NewQAndroidIntent() *QAndroidIntent {

	return nil
}

func NewQAndroidIntent2(intent QAndroidJniObject_ITF) *QAndroidIntent {

	return nil
}

func NewQAndroidIntent3(action string) *QAndroidIntent {

	return nil
}

func NewQAndroidIntent4(packageContext QAndroidJniObject_ITF, className string) *QAndroidIntent {

	return nil
}

func (ptr *QAndroidIntent) ExtraBytes(key string) *core.QByteArray {

	return nil
}

func (ptr *QAndroidIntent) ExtraVariant(key string) *core.QVariant {

	return nil
}

func (ptr *QAndroidIntent) Handle() *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidIntent) PutExtra(key string, data core.QByteArray_ITF) {

}

func (ptr *QAndroidIntent) PutExtra2(key string, value core.QVariant_ITF) {

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

	return nil
}

func (ptr *QAndroidJniEnvironment) FindClass(className string) unsafe.Pointer {

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

	return nil
}

func (ptr *QAndroidJniExceptionCleaner) Clean() {

}

func (ptr *QAndroidJniExceptionCleaner) DestroyQAndroidJniExceptionCleaner() {

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

func (ptr *QAndroidJniObject) CallMethodInt2(methodName string, sig string, v ...interface{}) int {

	return 0
}

func (ptr *QAndroidJniObject) CallMethodInt2Caught(methodName string, sig string, v ...interface{}) (int, error) {

	return 0, nil
}

func (ptr *QAndroidJniObject) CallMethodBoolean2(methodName string, sig string, v ...interface{}) bool {

	return false
}

func (ptr *QAndroidJniObject) CallMethodBoolean2Caught(methodName string, sig string, v ...interface{}) (bool, error) {

	return false, nil
}

func (ptr *QAndroidJniObject) CallMethodVoid2(methodName string, sig string, v ...interface{}) {

}

func (ptr *QAndroidJniObject) CallMethodVoid2Caught(methodName string, sig string, v ...interface{}) error {

	return nil
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

func QAndroidJniObject_IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) IsClassAvailable(className string) bool {

	return false
}

func (ptr *QAndroidJniObject) IsValid() bool {

	return false
}

func (ptr *QAndroidJniObject) Object() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidJniObject) SetField(fieldName string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetField2(fieldName string, signature string, value interface{}) {

}

func QAndroidJniObject_SetStaticField(className string, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetStaticField(className string, fieldName string, signature string, value interface{}) {

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

func QAndroidJniObject_SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

}

func (ptr *QAndroidJniObject) SetStaticField3(clazz unsafe.Pointer, fieldName string, signature string, value interface{}) {

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

func (ptr *QAndroidJniObject) ToString() string {

	return ""
}

func (ptr *QAndroidJniObject) DestroyQAndroidJniObject() {

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

func (ptr *QAndroidParcel) DestroyQAndroidParcel() {
}

func NewQAndroidParcel() *QAndroidParcel {

	return nil
}

func NewQAndroidParcel2(parcel QAndroidJniObject_ITF) *QAndroidParcel {

	return nil
}

func (ptr *QAndroidParcel) Handle() *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidParcel) ReadBinder() *QAndroidBinder {

	return nil
}

func (ptr *QAndroidParcel) ReadData() *core.QByteArray {

	return nil
}

func (ptr *QAndroidParcel) ReadFileDescriptor() int {

	return 0
}

func (ptr *QAndroidParcel) ReadVariant() *core.QVariant {

	return nil
}

func (ptr *QAndroidParcel) WriteBinder(binder QAndroidBinder_ITF) {

}

func (ptr *QAndroidParcel) WriteData(data core.QByteArray_ITF) {

}

func (ptr *QAndroidParcel) WriteFileDescriptor(fd int) {

}

func (ptr *QAndroidParcel) WriteVariant(value core.QVariant_ITF) {

}

type QAndroidService struct {
	core.QCoreApplication
}

type QAndroidService_ITF interface {
	core.QCoreApplication_ITF
	QAndroidService_PTR() *QAndroidService
}

func (ptr *QAndroidService) QAndroidService_PTR() *QAndroidService {
	return ptr
}

func (ptr *QAndroidService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QAndroidService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCoreApplication_PTR().SetPointer(p)
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

func (ptr *QAndroidService) DestroyQAndroidService() {
}

func NewQAndroidService(argc int, argv []string) *QAndroidService {

	return nil
}

func (ptr *QAndroidService) ConnectOnBind(f func(intent *QAndroidIntent) *QAndroidBinder) {

}

func (ptr *QAndroidService) DisconnectOnBind() {

}

func (ptr *QAndroidService) OnBind(intent QAndroidIntent_ITF) *QAndroidBinder {

	return nil
}

func (ptr *QAndroidService) OnBindDefault(intent QAndroidIntent_ITF) *QAndroidBinder {

	return nil
}

func (ptr *QAndroidService) __children_atList(i int) *core.QObject {

	return nil
}

func (ptr *QAndroidService) __children_setList(i core.QObject_ITF) {

}

func (ptr *QAndroidService) __children_newList() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidService) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return nil
}

func (ptr *QAndroidService) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

}

func (ptr *QAndroidService) __dynamicPropertyNames_newList() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidService) __findChildren_atList(i int) *core.QObject {

	return nil
}

func (ptr *QAndroidService) __findChildren_setList(i core.QObject_ITF) {

}

func (ptr *QAndroidService) __findChildren_newList() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidService) __findChildren_atList3(i int) *core.QObject {

	return nil
}

func (ptr *QAndroidService) __findChildren_setList3(i core.QObject_ITF) {

}

func (ptr *QAndroidService) __findChildren_newList3() unsafe.Pointer {

	return nil
}

func (ptr *QAndroidService) EventDefault(e core.QEvent_ITF) bool {

	return false
}

func (ptr *QAndroidService) QuitDefault() {

}

func (ptr *QAndroidService) ChildEventDefault(event core.QChildEvent_ITF) {

}

func (ptr *QAndroidService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

}

func (ptr *QAndroidService) CustomEventDefault(event core.QEvent_ITF) {

}

func (ptr *QAndroidService) DeleteLaterDefault() {

}

func (ptr *QAndroidService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

}

func (ptr *QAndroidService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return false
}

func (ptr *QAndroidService) MetaObjectDefault() *core.QMetaObject {

	return nil
}

func (ptr *QAndroidService) TimerEventDefault(event core.QTimerEvent_ITF) {

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

func (ptr *QAndroidServiceConnection) DestroyQAndroidServiceConnection() {
}

func NewQAndroidServiceConnection() *QAndroidServiceConnection {

	return nil
}

func NewQAndroidServiceConnection2(serviceConnection QAndroidJniObject_ITF) *QAndroidServiceConnection {

	return nil
}

func (ptr *QAndroidServiceConnection) Handle() *QAndroidJniObject {

	return nil
}

func (ptr *QAndroidServiceConnection) ConnectOnServiceConnected(f func(name string, serviceBinder *QAndroidBinder)) {

}

func (ptr *QAndroidServiceConnection) DisconnectOnServiceConnected() {

}

func (ptr *QAndroidServiceConnection) OnServiceConnected(name string, serviceBinder QAndroidBinder_ITF) {

}

func (ptr *QAndroidServiceConnection) ConnectOnServiceDisconnected(f func(name string)) {

}

func (ptr *QAndroidServiceConnection) DisconnectOnServiceDisconnected() {

}

func (ptr *QAndroidServiceConnection) OnServiceDisconnected(name string) {

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

func QtAndroid_AndroidSdkVersion() int {

	return 0
}

func (ptr *QtAndroid) AndroidSdkVersion() int {

	return 0
}

func QtAndroid_AndroidService() *QAndroidJniObject {

	return nil
}

func (ptr *QtAndroid) AndroidService() *QAndroidJniObject {

	return nil
}

func QtAndroid_BindService(serviceIntent QAndroidIntent_ITF, serviceConnection QAndroidServiceConnection_ITF, flags QtAndroid__BindFlag) bool {

	return false
}

func (ptr *QtAndroid) BindService(serviceIntent QAndroidIntent_ITF, serviceConnection QAndroidServiceConnection_ITF, flags QtAndroid__BindFlag) bool {

	return false
}

func QtAndroid_HideSplashScreen() {

}

func (ptr *QtAndroid) HideSplashScreen() {

}

func QtAndroid_HideSplashScreen2(duration int) {

}

func (ptr *QtAndroid) HideSplashScreen2(duration int) {

}

func QtAndroid_ShouldShowRequestPermissionRationale(permission string) bool {

	return false
}

func (ptr *QtAndroid) ShouldShowRequestPermissionRationale(permission string) bool {

	return false
}

func QtAndroid_StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func (ptr *QtAndroid) StartActivity(intent QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func QtAndroid_StartActivity2(intent QAndroidIntent_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func (ptr *QtAndroid) StartActivity2(intent QAndroidIntent_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func QtAndroid_StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func (ptr *QtAndroid) StartIntentSender(intentSender QAndroidJniObject_ITF, receiverRequestCode int, resultReceiver QAndroidActivityResultReceiver_ITF) {

}

func init() {
	qt.ItfMap["androidextras.QAndroidActivityResultReceiver_ITF"] = QAndroidActivityResultReceiver{}
	qt.ItfMap["androidextras.QAndroidBinder_ITF"] = QAndroidBinder{}
	qt.FuncMap["androidextras.NewQAndroidBinder"] = NewQAndroidBinder
	qt.FuncMap["androidextras.NewQAndroidBinder2"] = NewQAndroidBinder2
	qt.EnumMap["androidextras.QAndroidBinder__Normal"] = int64(QAndroidBinder__Normal)
	qt.EnumMap["androidextras.QAndroidBinder__OneWay"] = int64(QAndroidBinder__OneWay)
	qt.ItfMap["androidextras.QAndroidIntent_ITF"] = QAndroidIntent{}
	qt.FuncMap["androidextras.NewQAndroidIntent"] = NewQAndroidIntent
	qt.FuncMap["androidextras.NewQAndroidIntent2"] = NewQAndroidIntent2
	qt.FuncMap["androidextras.NewQAndroidIntent3"] = NewQAndroidIntent3
	qt.FuncMap["androidextras.NewQAndroidIntent4"] = NewQAndroidIntent4
	qt.ItfMap["androidextras.QAndroidJniEnvironment_ITF"] = QAndroidJniEnvironment{}
	qt.FuncMap["androidextras.NewQAndroidJniEnvironment"] = NewQAndroidJniEnvironment
	qt.FuncMap["androidextras.QAndroidJniEnvironment_JavaVM"] = QAndroidJniEnvironment_JavaVM
	qt.FuncMap["androidextras.QAndroidJniEnvironment_ExceptionCheck"] = QAndroidJniEnvironment_ExceptionCheck
	qt.FuncMap["androidextras.QAndroidJniEnvironment_ExceptionDescribe"] = QAndroidJniEnvironment_ExceptionDescribe
	qt.FuncMap["androidextras.QAndroidJniEnvironment_ExceptionClear"] = QAndroidJniEnvironment_ExceptionClear
	qt.FuncMap["androidextras.QAndroidJniEnvironment_ExceptionOccurred"] = QAndroidJniEnvironment_ExceptionOccurred
	qt.ItfMap["androidextras.QAndroidJniExceptionCleaner_ITF"] = QAndroidJniExceptionCleaner{}
	qt.FuncMap["androidextras.NewQAndroidJniExceptionCleaner"] = NewQAndroidJniExceptionCleaner
	qt.EnumMap["androidextras.QAndroidJniExceptionCleaner__Silent"] = int64(QAndroidJniExceptionCleaner__Silent)
	qt.EnumMap["androidextras.QAndroidJniExceptionCleaner__Verbose"] = int64(QAndroidJniExceptionCleaner__Verbose)
	qt.ItfMap["androidextras.QAndroidJniObject_ITF"] = QAndroidJniObject{}
	qt.FuncMap["androidextras.NewQAndroidJniObject"] = NewQAndroidJniObject
	qt.FuncMap["androidextras.NewQAndroidJniObject2"] = NewQAndroidJniObject2
	qt.FuncMap["androidextras.NewQAndroidJniObject3"] = NewQAndroidJniObject3
	qt.FuncMap["androidextras.NewQAndroidJniObject4"] = NewQAndroidJniObject4
	qt.FuncMap["androidextras.NewQAndroidJniObject5"] = NewQAndroidJniObject5
	qt.FuncMap["androidextras.NewQAndroidJniObject6"] = NewQAndroidJniObject6
	qt.FuncMap["androidextras.QAndroidJniObject_CallStaticObjectMethod"] = QAndroidJniObject_CallStaticObjectMethod
	qt.FuncMap["androidextras.QAndroidJniObject_CallStaticObjectMethod2"] = QAndroidJniObject_CallStaticObjectMethod2
	qt.FuncMap["androidextras.QAndroidJniObject_CallStaticObjectMethod3"] = QAndroidJniObject_CallStaticObjectMethod3
	qt.FuncMap["androidextras.QAndroidJniObject_CallStaticObjectMethod4"] = QAndroidJniObject_CallStaticObjectMethod4
	qt.FuncMap["androidextras.QAndroidJniObject_FromLocalRef"] = QAndroidJniObject_FromLocalRef
	qt.FuncMap["androidextras.QAndroidJniObject_FromString"] = QAndroidJniObject_FromString
	qt.FuncMap["androidextras.QAndroidJniObject_GetStaticObjectField"] = QAndroidJniObject_GetStaticObjectField
	qt.FuncMap["androidextras.QAndroidJniObject_GetStaticObjectField2"] = QAndroidJniObject_GetStaticObjectField2
	qt.FuncMap["androidextras.QAndroidJniObject_GetStaticObjectField3"] = QAndroidJniObject_GetStaticObjectField3
	qt.FuncMap["androidextras.QAndroidJniObject_GetStaticObjectField4"] = QAndroidJniObject_GetStaticObjectField4
	qt.FuncMap["androidextras.QAndroidJniObject_IsClassAvailable"] = QAndroidJniObject_IsClassAvailable
	qt.FuncMap["androidextras.QAndroidJniObject_SetStaticField"] = QAndroidJniObject_SetStaticField
	qt.FuncMap["androidextras.QAndroidJniObject_SetStaticField3"] = QAndroidJniObject_SetStaticField3
	qt.ItfMap["androidextras.QAndroidParcel_ITF"] = QAndroidParcel{}
	qt.FuncMap["androidextras.NewQAndroidParcel"] = NewQAndroidParcel
	qt.FuncMap["androidextras.NewQAndroidParcel2"] = NewQAndroidParcel2
	qt.ItfMap["androidextras.QAndroidService_ITF"] = QAndroidService{}
	qt.FuncMap["androidextras.NewQAndroidService"] = NewQAndroidService
	qt.ItfMap["androidextras.QAndroidServiceConnection_ITF"] = QAndroidServiceConnection{}
	qt.FuncMap["androidextras.NewQAndroidServiceConnection"] = NewQAndroidServiceConnection
	qt.FuncMap["androidextras.NewQAndroidServiceConnection2"] = NewQAndroidServiceConnection2
	qt.ItfMap["androidextras.QtAndroid_ITF"] = QtAndroid{}
	qt.FuncMap["androidextras.QtAndroid_AndroidActivity"] = QtAndroid_AndroidActivity
	qt.FuncMap["androidextras.QtAndroid_AndroidContext"] = QtAndroid_AndroidContext
	qt.FuncMap["androidextras.QtAndroid_AndroidSdkVersion"] = QtAndroid_AndroidSdkVersion
	qt.FuncMap["androidextras.QtAndroid_AndroidService"] = QtAndroid_AndroidService
	qt.FuncMap["androidextras.QtAndroid_BindService"] = QtAndroid_BindService
	qt.FuncMap["androidextras.QtAndroid_HideSplashScreen"] = QtAndroid_HideSplashScreen
	qt.FuncMap["androidextras.QtAndroid_HideSplashScreen2"] = QtAndroid_HideSplashScreen2
	qt.FuncMap["androidextras.QtAndroid_ShouldShowRequestPermissionRationale"] = QtAndroid_ShouldShowRequestPermissionRationale
	qt.FuncMap["androidextras.QtAndroid_StartActivity"] = QtAndroid_StartActivity
	qt.FuncMap["androidextras.QtAndroid_StartActivity2"] = QtAndroid_StartActivity2
	qt.FuncMap["androidextras.QtAndroid_StartIntentSender"] = QtAndroid_StartIntentSender
	qt.EnumMap["androidextras.QtAndroid__None"] = int64(QtAndroid__None)
	qt.EnumMap["androidextras.QtAndroid__AutoCreate"] = int64(QtAndroid__AutoCreate)
	qt.EnumMap["androidextras.QtAndroid__DebugUnbind"] = int64(QtAndroid__DebugUnbind)
	qt.EnumMap["androidextras.QtAndroid__NotForeground"] = int64(QtAndroid__NotForeground)
	qt.EnumMap["androidextras.QtAndroid__AboveClient"] = int64(QtAndroid__AboveClient)
	qt.EnumMap["androidextras.QtAndroid__AllowOomManagement"] = int64(QtAndroid__AllowOomManagement)
	qt.EnumMap["androidextras.QtAndroid__WaivePriority"] = int64(QtAndroid__WaivePriority)
	qt.EnumMap["androidextras.QtAndroid__Important"] = int64(QtAndroid__Important)
	qt.EnumMap["androidextras.QtAndroid__AdjustWithActivity"] = int64(QtAndroid__AdjustWithActivity)
	qt.EnumMap["androidextras.QtAndroid__ExternalService"] = int64(QtAndroid__ExternalService)
	qt.EnumMap["androidextras.QtAndroid__Granted"] = int64(QtAndroid__Granted)
	qt.EnumMap["androidextras.QtAndroid__Denied"] = int64(QtAndroid__Denied)
}
