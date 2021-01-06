// +build !minimal

package uitools

import (
	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/internal"
	"github.com/dev-drprasad/qt/widgets"
	"unsafe"
)

type QUiLoader struct {
	core.QObject
}

type QUiLoader_ITF interface {
	core.QObject_ITF
	QUiLoader_PTR() *QUiLoader
}

func (ptr *QUiLoader) QUiLoader_PTR() *QUiLoader {
	return ptr
}

func (ptr *QUiLoader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QUiLoader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQUiLoader(ptr QUiLoader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUiLoader_PTR().Pointer()
	}
	return nil
}

func (n *QUiLoader) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QUiLoader) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQUiLoaderFromPointer(ptr unsafe.Pointer) (n *QUiLoader) {
	n = new(QUiLoader)
	n.InitFromInternal(uintptr(ptr), "uitools.QUiLoader")
	return
}
func NewQUiLoader(parent core.QObject_ITF) *QUiLoader {

	return internal.CallLocalFunction([]interface{}{"", "", "uitools.NewQUiLoader", "", parent}).(*QUiLoader)
}

func (ptr *QUiLoader) AddPluginPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddPluginPath", path})
}

func (ptr *QUiLoader) AvailableLayouts() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableLayouts"}).([]string)
}

func (ptr *QUiLoader) AvailableWidgets() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableWidgets"}).([]string)
}

func (ptr *QUiLoader) ClearPluginPaths() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearPluginPaths"})
}

func (ptr *QUiLoader) ConnectCreateAction(f func(parent *core.QObject, name string) *widgets.QAction) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUiLoader) DisconnectCreateAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateAction"})
}

func (ptr *QUiLoader) CreateAction(parent core.QObject_ITF, name string) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateAction", parent, name}).(*widgets.QAction)
}

func (ptr *QUiLoader) CreateActionDefault(parent core.QObject_ITF, name string) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateActionDefault", parent, name}).(*widgets.QAction)
}

func (ptr *QUiLoader) ConnectCreateActionGroup(f func(parent *core.QObject, name string) *widgets.QActionGroup) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateActionGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUiLoader) DisconnectCreateActionGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateActionGroup"})
}

func (ptr *QUiLoader) CreateActionGroup(parent core.QObject_ITF, name string) *widgets.QActionGroup {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateActionGroup", parent, name}).(*widgets.QActionGroup)
}

func (ptr *QUiLoader) CreateActionGroupDefault(parent core.QObject_ITF, name string) *widgets.QActionGroup {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateActionGroupDefault", parent, name}).(*widgets.QActionGroup)
}

func (ptr *QUiLoader) ConnectCreateLayout(f func(className string, parent *core.QObject, name string) *widgets.QLayout) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateLayout", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUiLoader) DisconnectCreateLayout() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateLayout"})
}

func (ptr *QUiLoader) CreateLayout(className string, parent core.QObject_ITF, name string) *widgets.QLayout {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateLayout", className, parent, name}).(*widgets.QLayout)
}

func (ptr *QUiLoader) CreateLayoutDefault(className string, parent core.QObject_ITF, name string) *widgets.QLayout {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateLayoutDefault", className, parent, name}).(*widgets.QLayout)
}

func (ptr *QUiLoader) ConnectCreateWidget(f func(className string, parent *widgets.QWidget, name string) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUiLoader) DisconnectCreateWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateWidget"})
}

func (ptr *QUiLoader) CreateWidget(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateWidget", className, parent, name}).(*widgets.QWidget)
}

func (ptr *QUiLoader) CreateWidgetDefault(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateWidgetDefault", className, parent, name}).(*widgets.QWidget)
}

func (ptr *QUiLoader) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QUiLoader) IsLanguageChangeEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLanguageChangeEnabled"}).(bool)
}

func (ptr *QUiLoader) Load(device core.QIODevice_ITF, parentWidget widgets.QWidget_ITF) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", device, parentWidget}).(*widgets.QWidget)
}

func (ptr *QUiLoader) PluginPaths() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PluginPaths"}).([]string)
}

func (ptr *QUiLoader) SetLanguageChangeEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLanguageChangeEnabled", enabled})
}

func (ptr *QUiLoader) SetWorkingDirectory(dir core.QDir_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWorkingDirectory", dir})
}

func (ptr *QUiLoader) WorkingDirectory() *core.QDir {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WorkingDirectory"}).(*core.QDir)
}

func (ptr *QUiLoader) ConnectDestroyQUiLoader(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQUiLoader", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUiLoader) DisconnectDestroyQUiLoader() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQUiLoader"})
}

func (ptr *QUiLoader) DestroyQUiLoader() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQUiLoader"})
}

func (ptr *QUiLoader) DestroyQUiLoaderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQUiLoaderDefault"})
}

func (ptr *QUiLoader) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QUiLoader) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QUiLoader) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QUiLoader) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QUiLoader) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QUiLoader) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QUiLoader) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QUiLoader) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QUiLoader) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QUiLoader) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QUiLoader) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QUiLoader) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QUiLoader) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QUiLoader) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QUiLoader) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QUiLoader) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QUiLoader) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QUiLoader) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QUiLoader) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QUiLoader) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QUiLoader) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

func init() {
	internal.ConstructorTable["uitools.QUiLoader"] = NewQUiLoaderFromPointer
}
