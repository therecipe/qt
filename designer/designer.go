// +build !minimal

package designer

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/widgets"
	"strings"
	"unsafe"
)

type AppFontDialog struct {
	widgets.QDialog
}

type AppFontDialog_ITF interface {
	widgets.QDialog_ITF
	AppFontDialog_PTR() *AppFontDialog
}

func (ptr *AppFontDialog) AppFontDialog_PTR() *AppFontDialog {
	return ptr
}

func (ptr *AppFontDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *AppFontDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromAppFontDialog(ptr AppFontDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AppFontDialog_PTR().Pointer()
	}
	return nil
}

func (n *AppFontDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *AppFontDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewAppFontDialogFromPointer(ptr unsafe.Pointer) (n *AppFontDialog) {
	n = new(AppFontDialog)
	n.InitFromInternal(uintptr(ptr), "designer.AppFontDialog")
	return
}

func (ptr *AppFontDialog) DestroyAppFontDialog() {
}

type AppFontWidget struct {
	widgets.QGroupBox
}

type AppFontWidget_ITF interface {
	widgets.QGroupBox_ITF
	AppFontWidget_PTR() *AppFontWidget
}

func (ptr *AppFontWidget) AppFontWidget_PTR() *AppFontWidget {
	return ptr
}

func (ptr *AppFontWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBox_PTR().Pointer()
	}
	return nil
}

func (ptr *AppFontWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGroupBox_PTR().SetPointer(p)
	}
}

func PointerFromAppFontWidget(ptr AppFontWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AppFontWidget_PTR().Pointer()
	}
	return nil
}

func (n *AppFontWidget) InitFromInternal(ptr uintptr, name string) {
	n.QGroupBox_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *AppFontWidget) ClassNameInternalF() string {
	return n.QGroupBox_PTR().ClassNameInternalF()
}

func NewAppFontWidgetFromPointer(ptr unsafe.Pointer) (n *AppFontWidget) {
	n = new(AppFontWidget)
	n.InitFromInternal(uintptr(ptr), "designer.AppFontWidget")
	return
}

func (ptr *AppFontWidget) DestroyAppFontWidget() {
}

type AppearanceOptions struct {
	internal.Internal
}

type AppearanceOptions_ITF interface {
	AppearanceOptions_PTR() *AppearanceOptions
}

func (ptr *AppearanceOptions) AppearanceOptions_PTR() *AppearanceOptions {
	return ptr
}

func (ptr *AppearanceOptions) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *AppearanceOptions) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromAppearanceOptions(ptr AppearanceOptions_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AppearanceOptions_PTR().Pointer()
	}
	return nil
}

func (n *AppearanceOptions) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewAppearanceOptionsFromPointer(ptr unsafe.Pointer) (n *AppearanceOptions) {
	n = new(AppearanceOptions)
	n.InitFromInternal(uintptr(ptr), "designer.AppearanceOptions")
	return
}

func (ptr *AppearanceOptions) DestroyAppearanceOptions() {
}

type AssistantClient struct {
	core.QObject
}

type AssistantClient_ITF interface {
	core.QObject_ITF
	AssistantClient_PTR() *AssistantClient
}

func (ptr *AssistantClient) AssistantClient_PTR() *AssistantClient {
	return ptr
}

func (ptr *AssistantClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *AssistantClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromAssistantClient(ptr AssistantClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AssistantClient_PTR().Pointer()
	}
	return nil
}

func (n *AssistantClient) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *AssistantClient) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewAssistantClientFromPointer(ptr unsafe.Pointer) (n *AssistantClient) {
	n = new(AssistantClient)
	n.InitFromInternal(uintptr(ptr), "designer.AssistantClient")
	return
}

func (ptr *AssistantClient) DestroyAssistantClient() {
}

type DockedMainWindow struct {
	MainWindowBase
}

type DockedMainWindow_ITF interface {
	MainWindowBase_ITF
	DockedMainWindow_PTR() *DockedMainWindow
}

func (ptr *DockedMainWindow) DockedMainWindow_PTR() *DockedMainWindow {
	return ptr
}

func (ptr *DockedMainWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.MainWindowBase_PTR().Pointer()
	}
	return nil
}

func (ptr *DockedMainWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.MainWindowBase_PTR().SetPointer(p)
	}
}

func PointerFromDockedMainWindow(ptr DockedMainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DockedMainWindow_PTR().Pointer()
	}
	return nil
}

func (n *DockedMainWindow) InitFromInternal(ptr uintptr, name string) {
	n.MainWindowBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *DockedMainWindow) ClassNameInternalF() string {
	return n.MainWindowBase_PTR().ClassNameInternalF()
}

func NewDockedMainWindowFromPointer(ptr unsafe.Pointer) (n *DockedMainWindow) {
	n = new(DockedMainWindow)
	n.InitFromInternal(uintptr(ptr), "designer.DockedMainWindow")
	return
}

func (ptr *DockedMainWindow) DestroyDockedMainWindow() {
}

type DockedMdiArea struct {
	widgets.QMdiArea
}

type DockedMdiArea_ITF interface {
	widgets.QMdiArea_ITF
	DockedMdiArea_PTR() *DockedMdiArea
}

func (ptr *DockedMdiArea) DockedMdiArea_PTR() *DockedMdiArea {
	return ptr
}

func (ptr *DockedMdiArea) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiArea_PTR().Pointer()
	}
	return nil
}

func (ptr *DockedMdiArea) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMdiArea_PTR().SetPointer(p)
	}
}

func PointerFromDockedMdiArea(ptr DockedMdiArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.DockedMdiArea_PTR().Pointer()
	}
	return nil
}

func (n *DockedMdiArea) InitFromInternal(ptr uintptr, name string) {
	n.QMdiArea_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *DockedMdiArea) ClassNameInternalF() string {
	return n.QMdiArea_PTR().ClassNameInternalF()
}

func NewDockedMdiAreaFromPointer(ptr unsafe.Pointer) (n *DockedMdiArea) {
	n = new(DockedMdiArea)
	n.InitFromInternal(uintptr(ptr), "designer.DockedMdiArea")
	return
}

func (ptr *DockedMdiArea) DestroyDockedMdiArea() {
}

type MainWindowBase struct {
	widgets.QMainWindow
}

type MainWindowBase_ITF interface {
	widgets.QMainWindow_ITF
	MainWindowBase_PTR() *MainWindowBase
}

func (ptr *MainWindowBase) MainWindowBase_PTR() *MainWindowBase {
	return ptr
}

func (ptr *MainWindowBase) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *MainWindowBase) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromMainWindowBase(ptr MainWindowBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.MainWindowBase_PTR().Pointer()
	}
	return nil
}

func (n *MainWindowBase) InitFromInternal(ptr uintptr, name string) {
	n.QMainWindow_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *MainWindowBase) ClassNameInternalF() string {
	return n.QMainWindow_PTR().ClassNameInternalF()
}

func NewMainWindowBaseFromPointer(ptr unsafe.Pointer) (n *MainWindowBase) {
	n = new(MainWindowBase)
	n.InitFromInternal(uintptr(ptr), "designer.MainWindowBase")
	return
}

func (ptr *MainWindowBase) DestroyMainWindowBase() {
}

type NewForm struct {
	widgets.QDialog
}

type NewForm_ITF interface {
	widgets.QDialog_ITF
	NewForm_PTR() *NewForm
}

func (ptr *NewForm) NewForm_PTR() *NewForm {
	return ptr
}

func (ptr *NewForm) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *NewForm) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromNewForm(ptr NewForm_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.NewForm_PTR().Pointer()
	}
	return nil
}

func (n *NewForm) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *NewForm) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewNewFormFromPointer(ptr unsafe.Pointer) (n *NewForm) {
	n = new(NewForm)
	n.InitFromInternal(uintptr(ptr), "designer.NewForm")
	return
}

func (ptr *NewForm) DestroyNewForm() {
}

type PreferencesDialog struct {
	widgets.QDialog
}

type PreferencesDialog_ITF interface {
	widgets.QDialog_ITF
	PreferencesDialog_PTR() *PreferencesDialog
}

func (ptr *PreferencesDialog) PreferencesDialog_PTR() *PreferencesDialog {
	return ptr
}

func (ptr *PreferencesDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *PreferencesDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromPreferencesDialog(ptr PreferencesDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.PreferencesDialog_PTR().Pointer()
	}
	return nil
}

func (n *PreferencesDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *PreferencesDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewPreferencesDialogFromPointer(ptr unsafe.Pointer) (n *PreferencesDialog) {
	n = new(PreferencesDialog)
	n.InitFromInternal(uintptr(ptr), "designer.PreferencesDialog")
	return
}

func (ptr *PreferencesDialog) DestroyPreferencesDialog() {
}

type QAbstractExtensionFactory struct {
	internal.Internal
}

type QAbstractExtensionFactory_ITF interface {
	QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory
}

func (ptr *QAbstractExtensionFactory) QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory {
	return ptr
}

func (ptr *QAbstractExtensionFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractExtensionFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractExtensionFactory(ptr QAbstractExtensionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractExtensionFactory_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractExtensionFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractExtensionFactoryFromPointer(ptr unsafe.Pointer) (n *QAbstractExtensionFactory) {
	n = new(QAbstractExtensionFactory)
	n.InitFromInternal(uintptr(ptr), "designer.QAbstractExtensionFactory")
	return
}

type QAbstractExtensionManager struct {
	internal.Internal
}

type QAbstractExtensionManager_ITF interface {
	QAbstractExtensionManager_PTR() *QAbstractExtensionManager
}

func (ptr *QAbstractExtensionManager) QAbstractExtensionManager_PTR() *QAbstractExtensionManager {
	return ptr
}

func (ptr *QAbstractExtensionManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractExtensionManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractExtensionManager(ptr QAbstractExtensionManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractExtensionManager_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractExtensionManager) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractExtensionManagerFromPointer(ptr unsafe.Pointer) (n *QAbstractExtensionManager) {
	n = new(QAbstractExtensionManager)
	n.InitFromInternal(uintptr(ptr), "designer.QAbstractExtensionManager")
	return
}
func (ptr *QAbstractExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractExtensionManager) DisconnectExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExtension"})
}

func (ptr *QAbstractExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Extension", object, iid}).(*core.QObject)
}

func (ptr *QAbstractExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRegisterExtensions", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractExtensionManager) DisconnectRegisterExtensions() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRegisterExtensions"})
}

func (ptr *QAbstractExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RegisterExtensions", factory, iid})
}

func (ptr *QAbstractExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUnregisterExtensions", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractExtensionManager) DisconnectUnregisterExtensions() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUnregisterExtensions"})
}

func (ptr *QAbstractExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnregisterExtensions", factory, iid})
}

func (ptr *QAbstractExtensionManager) ConnectDestroyQAbstractExtensionManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractExtensionManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractExtensionManager) DisconnectDestroyQAbstractExtensionManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractExtensionManager"})
}

func (ptr *QAbstractExtensionManager) DestroyQAbstractExtensionManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractExtensionManager"})
}

func (ptr *QAbstractExtensionManager) DestroyQAbstractExtensionManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractExtensionManagerDefault"})
}

type QAbstractFormBuilder struct {
	internal.Internal
}

type QAbstractFormBuilder_ITF interface {
	QAbstractFormBuilder_PTR() *QAbstractFormBuilder
}

func (ptr *QAbstractFormBuilder) QAbstractFormBuilder_PTR() *QAbstractFormBuilder {
	return ptr
}

func (ptr *QAbstractFormBuilder) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAbstractFormBuilder) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAbstractFormBuilder(ptr QAbstractFormBuilder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractFormBuilder_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractFormBuilder) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAbstractFormBuilderFromPointer(ptr unsafe.Pointer) (n *QAbstractFormBuilder) {
	n = new(QAbstractFormBuilder)
	n.InitFromInternal(uintptr(ptr), "designer.QAbstractFormBuilder")
	return
}
func NewQAbstractFormBuilder() *QAbstractFormBuilder {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQAbstractFormBuilder", ""}).(*QAbstractFormBuilder)
}

func (ptr *QAbstractFormBuilder) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QAbstractFormBuilder) ConnectLoad(f func(device *core.QIODevice, parent *widgets.QWidget) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractFormBuilder) DisconnectLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad"})
}

func (ptr *QAbstractFormBuilder) Load(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load", device, parent}).(*widgets.QWidget)
}

func (ptr *QAbstractFormBuilder) LoadDefault(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LoadDefault", device, parent}).(*widgets.QWidget)
}

func (ptr *QAbstractFormBuilder) ConnectSave(f func(device *core.QIODevice, widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSave", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractFormBuilder) DisconnectSave() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSave"})
}

func (ptr *QAbstractFormBuilder) Save(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Save", device, widget})
}

func (ptr *QAbstractFormBuilder) SaveDefault(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SaveDefault", device, widget})
}

func (ptr *QAbstractFormBuilder) SetWorkingDirectory(directory core.QDir_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWorkingDirectory", directory})
}

func (ptr *QAbstractFormBuilder) WorkingDirectory() *core.QDir {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WorkingDirectory"}).(*core.QDir)
}

func (ptr *QAbstractFormBuilder) ConnectDestroyQAbstractFormBuilder(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractFormBuilder", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractFormBuilder) DisconnectDestroyQAbstractFormBuilder() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractFormBuilder"})
}

func (ptr *QAbstractFormBuilder) DestroyQAbstractFormBuilder() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractFormBuilder"})
}

func (ptr *QAbstractFormBuilder) DestroyQAbstractFormBuilderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractFormBuilderDefault"})
}

func (ptr *QAbstractFormBuilder) __applyProperties_properties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__applyProperties_properties_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractFormBuilder) __computeProperties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__computeProperties_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractFormBuilder) __propertyMap_properties_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__propertyMap_properties_newList"}).(unsafe.Pointer)
}

type QDesigner struct {
	widgets.QApplication
}

type QDesigner_ITF interface {
	widgets.QApplication_ITF
	QDesigner_PTR() *QDesigner
}

func (ptr *QDesigner) QDesigner_PTR() *QDesigner {
	return ptr
}

func (ptr *QDesigner) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesigner) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QApplication_PTR().SetPointer(p)
	}
}

func PointerFromQDesigner(ptr QDesigner_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesigner_PTR().Pointer()
	}
	return nil
}

func (n *QDesigner) InitFromInternal(ptr uintptr, name string) {
	n.QApplication_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesigner) ClassNameInternalF() string {
	return n.QApplication_PTR().ClassNameInternalF()
}

func NewQDesignerFromPointer(ptr unsafe.Pointer) (n *QDesigner) {
	n = new(QDesigner)
	n.InitFromInternal(uintptr(ptr), "designer.QDesigner")
	return
}

func (ptr *QDesigner) DestroyQDesigner() {
}

type QDesignerActionEditorInterface struct {
	widgets.QWidget
}

type QDesignerActionEditorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerActionEditorInterface_PTR() *QDesignerActionEditorInterface
}

func (ptr *QDesignerActionEditorInterface) QDesignerActionEditorInterface_PTR() *QDesignerActionEditorInterface {
	return ptr
}

func (ptr *QDesignerActionEditorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerActionEditorInterface(ptr QDesignerActionEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerActionEditorInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerActionEditorInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerActionEditorInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerActionEditorInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerActionEditorInterface) {
	n = new(QDesignerActionEditorInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerActionEditorInterface")
	return
}
func NewQDesignerActionEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerActionEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQDesignerActionEditorInterface", "", parent, flags}).(*QDesignerActionEditorInterface)
}

func (ptr *QDesignerActionEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerActionEditorInterface) DisconnectCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCore"})
}

func (ptr *QDesignerActionEditorInterface) Core() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Core"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerActionEditorInterface) CoreDefault() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoreDefault"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerActionEditorInterface) ConnectManageAction(f func(action *widgets.QAction)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectManageAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerActionEditorInterface) DisconnectManageAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectManageAction"})
}

func (ptr *QDesignerActionEditorInterface) ManageAction(action widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManageAction", action})
}

func (ptr *QDesignerActionEditorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFormWindow"})
}

func (ptr *QDesignerActionEditorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFormWindow", formWindow})
}

func (ptr *QDesignerActionEditorInterface) ConnectUnmanageAction(f func(action *widgets.QAction)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUnmanageAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerActionEditorInterface) DisconnectUnmanageAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUnmanageAction"})
}

func (ptr *QDesignerActionEditorInterface) UnmanageAction(action widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnmanageAction", action})
}

func (ptr *QDesignerActionEditorInterface) ConnectDestroyQDesignerActionEditorInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerActionEditorInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerActionEditorInterface) DisconnectDestroyQDesignerActionEditorInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerActionEditorInterface"})
}

func (ptr *QDesignerActionEditorInterface) DestroyQDesignerActionEditorInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerActionEditorInterface"})
}

func (ptr *QDesignerActionEditorInterface) DestroyQDesignerActionEditorInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerActionEditorInterfaceDefault"})
}

func (ptr *QDesignerActionEditorInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerActionEditorInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerActionEditorInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerActionEditorInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerActionEditorInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerActionEditorInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerActionEditorInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerActionEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerActionEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerActionEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerActionEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerActionEditorInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerActionEditorInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerActionEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerActionEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerActionEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerActionEditorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerActionEditorInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerActionEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerActionEditorInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerActionEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerActionEditorInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerActionEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerActionEditorInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerActionEditorInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerActionEditorInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerActionEditorInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerActionEditorInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerActionEditorInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerActionEditorInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerActionEditorInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerActionEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerActionEditorInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerActionEditorInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerActionEditorInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerActionEditorInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerActionEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerActionEditorInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerActionEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerActionEditorInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerActionEditorInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerActionEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerActionEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerActionEditorInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerActionEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerActions struct {
	core.QObject
}

type QDesignerActions_ITF interface {
	core.QObject_ITF
	QDesignerActions_PTR() *QDesignerActions
}

func (ptr *QDesignerActions) QDesignerActions_PTR() *QDesignerActions {
	return ptr
}

func (ptr *QDesignerActions) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerActions) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerActions(ptr QDesignerActions_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerActions_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerActions) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerActions) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerActionsFromPointer(ptr unsafe.Pointer) (n *QDesignerActions) {
	n = new(QDesignerActions)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerActions")
	return
}

func (ptr *QDesignerActions) DestroyQDesignerActions() {
}

type QDesignerAppearanceOptionsPage struct {
	core.QObject
}

type QDesignerAppearanceOptionsPage_ITF interface {
	core.QObject_ITF
	QDesignerAppearanceOptionsPage_PTR() *QDesignerAppearanceOptionsPage
}

func (ptr *QDesignerAppearanceOptionsPage) QDesignerAppearanceOptionsPage_PTR() *QDesignerAppearanceOptionsPage {
	return ptr
}

func (ptr *QDesignerAppearanceOptionsPage) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerAppearanceOptionsPage) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerAppearanceOptionsPage(ptr QDesignerAppearanceOptionsPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerAppearanceOptionsPage_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerAppearanceOptionsPage) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerAppearanceOptionsPage) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerAppearanceOptionsPageFromPointer(ptr unsafe.Pointer) (n *QDesignerAppearanceOptionsPage) {
	n = new(QDesignerAppearanceOptionsPage)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerAppearanceOptionsPage")
	return
}

func (ptr *QDesignerAppearanceOptionsPage) DestroyQDesignerAppearanceOptionsPage() {
}

type QDesignerAppearanceOptionsWidget struct {
	widgets.QWidget
}

type QDesignerAppearanceOptionsWidget_ITF interface {
	widgets.QWidget_ITF
	QDesignerAppearanceOptionsWidget_PTR() *QDesignerAppearanceOptionsWidget
}

func (ptr *QDesignerAppearanceOptionsWidget) QDesignerAppearanceOptionsWidget_PTR() *QDesignerAppearanceOptionsWidget {
	return ptr
}

func (ptr *QDesignerAppearanceOptionsWidget) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerAppearanceOptionsWidget) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerAppearanceOptionsWidget(ptr QDesignerAppearanceOptionsWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerAppearanceOptionsWidget_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerAppearanceOptionsWidget) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerAppearanceOptionsWidget) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerAppearanceOptionsWidgetFromPointer(ptr unsafe.Pointer) (n *QDesignerAppearanceOptionsWidget) {
	n = new(QDesignerAppearanceOptionsWidget)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerAppearanceOptionsWidget")
	return
}

func (ptr *QDesignerAppearanceOptionsWidget) DestroyQDesignerAppearanceOptionsWidget() {
}

type QDesignerClient struct {
	core.QObject
}

type QDesignerClient_ITF interface {
	core.QObject_ITF
	QDesignerClient_PTR() *QDesignerClient
}

func (ptr *QDesignerClient) QDesignerClient_PTR() *QDesignerClient {
	return ptr
}

func (ptr *QDesignerClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerClient(ptr QDesignerClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerClient_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerClient) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerClient) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerClientFromPointer(ptr unsafe.Pointer) (n *QDesignerClient) {
	n = new(QDesignerClient)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerClient")
	return
}

func (ptr *QDesignerClient) DestroyQDesignerClient() {
}

type QDesignerComponents struct {
	internal.Internal
}

type QDesignerComponents_ITF interface {
	QDesignerComponents_PTR() *QDesignerComponents
}

func (ptr *QDesignerComponents) QDesignerComponents_PTR() *QDesignerComponents {
	return ptr
}

func (ptr *QDesignerComponents) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerComponents) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerComponents(ptr QDesignerComponents_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerComponents_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerComponents) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerComponentsFromPointer(ptr unsafe.Pointer) (n *QDesignerComponents) {
	n = new(QDesignerComponents)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerComponents")
	return
}

func (ptr *QDesignerComponents) DestroyQDesignerComponents() {
}

type QDesignerContainerExtension struct {
	internal.Internal
}

type QDesignerContainerExtension_ITF interface {
	QDesignerContainerExtension_PTR() *QDesignerContainerExtension
}

func (ptr *QDesignerContainerExtension) QDesignerContainerExtension_PTR() *QDesignerContainerExtension {
	return ptr
}

func (ptr *QDesignerContainerExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerContainerExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerContainerExtension(ptr QDesignerContainerExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerContainerExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerContainerExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerContainerExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerContainerExtension) {
	n = new(QDesignerContainerExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerContainerExtension")
	return
}
func (ptr *QDesignerContainerExtension) ConnectAddWidget(f func(page *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAddWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectAddWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAddWidget"})
}

func (ptr *QDesignerContainerExtension) AddWidget(page widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddWidget", page})
}

func (ptr *QDesignerContainerExtension) ConnectCanAddWidget(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanAddWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectCanAddWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanAddWidget"})
}

func (ptr *QDesignerContainerExtension) CanAddWidget() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanAddWidget"}).(bool)
}

func (ptr *QDesignerContainerExtension) CanAddWidgetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanAddWidgetDefault"}).(bool)
}

func (ptr *QDesignerContainerExtension) ConnectCanRemove(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanRemove", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectCanRemove() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanRemove"})
}

func (ptr *QDesignerContainerExtension) CanRemove(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanRemove", index}).(bool)
}

func (ptr *QDesignerContainerExtension) CanRemoveDefault(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanRemoveDefault", index}).(bool)
}

func (ptr *QDesignerContainerExtension) ConnectCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCount"})
}

func (ptr *QDesignerContainerExtension) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QDesignerContainerExtension) ConnectCurrentIndex(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCurrentIndex", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectCurrentIndex() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCurrentIndex"})
}

func (ptr *QDesignerContainerExtension) CurrentIndex() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentIndex"}).(float64))
}

func (ptr *QDesignerContainerExtension) ConnectInsertWidget(f func(index int, page *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInsertWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectInsertWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInsertWidget"})
}

func (ptr *QDesignerContainerExtension) InsertWidget(index int, page widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertWidget", index, page})
}

func (ptr *QDesignerContainerExtension) ConnectRemove(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemove", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectRemove() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemove"})
}

func (ptr *QDesignerContainerExtension) Remove(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", index})
}

func (ptr *QDesignerContainerExtension) ConnectSetCurrentIndex(f func(index int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetCurrentIndex", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectSetCurrentIndex() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetCurrentIndex"})
}

func (ptr *QDesignerContainerExtension) SetCurrentIndex(index int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCurrentIndex", index})
}

func (ptr *QDesignerContainerExtension) ConnectWidget(f func(index int) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidget"})
}

func (ptr *QDesignerContainerExtension) Widget(index int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Widget", index}).(*widgets.QWidget)
}

func (ptr *QDesignerContainerExtension) ConnectDestroyQDesignerContainerExtension(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerContainerExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerContainerExtension) DisconnectDestroyQDesignerContainerExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerContainerExtension"})
}

func (ptr *QDesignerContainerExtension) DestroyQDesignerContainerExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerContainerExtension"})
}

func (ptr *QDesignerContainerExtension) DestroyQDesignerContainerExtensionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerContainerExtensionDefault"})
}

type QDesignerCustomWidgetCollectionInterface struct {
	internal.Internal
}

type QDesignerCustomWidgetCollectionInterface_ITF interface {
	QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface
}

func (ptr *QDesignerCustomWidgetCollectionInterface) QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface {
	return ptr
}

func (ptr *QDesignerCustomWidgetCollectionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerCustomWidgetCollectionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerCustomWidgetCollectionInterface(ptr QDesignerCustomWidgetCollectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerCustomWidgetCollectionInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerCustomWidgetCollectionInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerCustomWidgetCollectionInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerCustomWidgetCollectionInterface) {
	n = new(QDesignerCustomWidgetCollectionInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerCustomWidgetCollectionInterface")
	return
}
func (ptr *QDesignerCustomWidgetCollectionInterface) ConnectCustomWidgets(f func() []*QDesignerCustomWidgetInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCustomWidgets", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DisconnectCustomWidgets() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCustomWidgets"})
}

func (ptr *QDesignerCustomWidgetCollectionInterface) CustomWidgets() []*QDesignerCustomWidgetInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomWidgets"}).([]*QDesignerCustomWidgetInterface)
}

func (ptr *QDesignerCustomWidgetCollectionInterface) ConnectDestroyQDesignerCustomWidgetCollectionInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerCustomWidgetCollectionInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DisconnectDestroyQDesignerCustomWidgetCollectionInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerCustomWidgetCollectionInterface"})
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DestroyQDesignerCustomWidgetCollectionInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerCustomWidgetCollectionInterface"})
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DestroyQDesignerCustomWidgetCollectionInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerCustomWidgetCollectionInterfaceDefault"})
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_atList(i int) *QDesignerCustomWidgetInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_atList", i}).(*QDesignerCustomWidgetInterface)
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_setList(i QDesignerCustomWidgetInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_setList", i})
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_newList"}).(unsafe.Pointer)
}

type QDesignerCustomWidgetInterface struct {
	internal.Internal
}

type QDesignerCustomWidgetInterface_ITF interface {
	QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface
}

func (ptr *QDesignerCustomWidgetInterface) QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface {
	return ptr
}

func (ptr *QDesignerCustomWidgetInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerCustomWidgetInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerCustomWidgetInterface(ptr QDesignerCustomWidgetInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerCustomWidgetInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerCustomWidgetInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerCustomWidgetInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerCustomWidgetInterface) {
	n = new(QDesignerCustomWidgetInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerCustomWidgetInterface")
	return
}
func (ptr *QDesignerCustomWidgetInterface) ConnectCodeTemplate(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCodeTemplate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCodeTemplate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCodeTemplate"})
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplate() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CodeTemplate"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplateDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CodeTemplateDefault"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectCreateWidget(f func(parent *widgets.QWidget) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCreateWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateWidget"})
}

func (ptr *QDesignerCustomWidgetInterface) CreateWidget(parent widgets.QWidget_ITF) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateWidget", parent}).(*widgets.QWidget)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectDomXml(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDomXml", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectDomXml() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDomXml"})
}

func (ptr *QDesignerCustomWidgetInterface) DomXml() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DomXml"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) DomXmlDefault() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DomXmlDefault"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectGroup(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGroup"})
}

func (ptr *QDesignerCustomWidgetInterface) Group() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Group"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIcon(f func() *gui.QIcon) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIcon", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIcon() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIcon"})
}

func (ptr *QDesignerCustomWidgetInterface) Icon() *gui.QIcon {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Icon"}).(*gui.QIcon)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIncludeFile(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIncludeFile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIncludeFile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIncludeFile"})
}

func (ptr *QDesignerCustomWidgetInterface) IncludeFile() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludeFile"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectInitialize(f func(formEditor *QDesignerFormEditorInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInitialize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectInitialize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInitialize"})
}

func (ptr *QDesignerCustomWidgetInterface) Initialize(formEditor QDesignerFormEditorInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Initialize", formEditor})
}

func (ptr *QDesignerCustomWidgetInterface) InitializeDefault(formEditor QDesignerFormEditorInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitializeDefault", formEditor})
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsContainer(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsContainer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsContainer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsContainer"})
}

func (ptr *QDesignerCustomWidgetInterface) IsContainer() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsContainer"}).(bool)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsInitialized(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsInitialized", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsInitialized() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsInitialized"})
}

func (ptr *QDesignerCustomWidgetInterface) IsInitialized() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInitialized"}).(bool)
}

func (ptr *QDesignerCustomWidgetInterface) IsInitializedDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInitializedDefault"}).(bool)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectName"})
}

func (ptr *QDesignerCustomWidgetInterface) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectToolTip(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToolTip", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectToolTip() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectToolTip"})
}

func (ptr *QDesignerCustomWidgetInterface) ToolTip() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToolTip"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectWhatsThis(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWhatsThis", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectWhatsThis() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWhatsThis"})
}

func (ptr *QDesignerCustomWidgetInterface) WhatsThis() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WhatsThis"}).(string)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectDestroyQDesignerCustomWidgetInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerCustomWidgetInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectDestroyQDesignerCustomWidgetInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerCustomWidgetInterface"})
}

func (ptr *QDesignerCustomWidgetInterface) DestroyQDesignerCustomWidgetInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerCustomWidgetInterface"})
}

func (ptr *QDesignerCustomWidgetInterface) DestroyQDesignerCustomWidgetInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerCustomWidgetInterfaceDefault"})
}

type QDesignerDynamicPropertySheetExtension struct {
	internal.Internal
}

type QDesignerDynamicPropertySheetExtension_ITF interface {
	QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension
}

func (ptr *QDesignerDynamicPropertySheetExtension) QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension {
	return ptr
}

func (ptr *QDesignerDynamicPropertySheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerDynamicPropertySheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerDynamicPropertySheetExtension(ptr QDesignerDynamicPropertySheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerDynamicPropertySheetExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerDynamicPropertySheetExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerDynamicPropertySheetExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerDynamicPropertySheetExtension) {
	n = new(QDesignerDynamicPropertySheetExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerDynamicPropertySheetExtension")
	return
}
func (ptr *QDesignerDynamicPropertySheetExtension) ConnectAddDynamicProperty(f func(propertyName string, value *core.QVariant) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAddDynamicProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectAddDynamicProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAddDynamicProperty"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) AddDynamicProperty(propertyName string, value core.QVariant_ITF) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddDynamicProperty", propertyName, value}).(float64))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectCanAddDynamicProperty(f func(propertyName string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCanAddDynamicProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectCanAddDynamicProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCanAddDynamicProperty"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) CanAddDynamicProperty(propertyName string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanAddDynamicProperty", propertyName}).(bool)
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectDynamicPropertiesAllowed(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDynamicPropertiesAllowed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectDynamicPropertiesAllowed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDynamicPropertiesAllowed"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) DynamicPropertiesAllowed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DynamicPropertiesAllowed"}).(bool)
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectIsDynamicProperty(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsDynamicProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectIsDynamicProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsDynamicProperty"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) IsDynamicProperty(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDynamicProperty", index}).(bool)
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectRemoveDynamicProperty(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemoveDynamicProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectRemoveDynamicProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoveDynamicProperty"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) RemoveDynamicProperty(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveDynamicProperty", index}).(bool)
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectDestroyQDesignerDynamicPropertySheetExtension(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerDynamicPropertySheetExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectDestroyQDesignerDynamicPropertySheetExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerDynamicPropertySheetExtension"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) DestroyQDesignerDynamicPropertySheetExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerDynamicPropertySheetExtension"})
}

func (ptr *QDesignerDynamicPropertySheetExtension) DestroyQDesignerDynamicPropertySheetExtensionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerDynamicPropertySheetExtensionDefault"})
}

type QDesignerFormEditorInterface struct {
	core.QObject
}

type QDesignerFormEditorInterface_ITF interface {
	core.QObject_ITF
	QDesignerFormEditorInterface_PTR() *QDesignerFormEditorInterface
}

func (ptr *QDesignerFormEditorInterface) QDesignerFormEditorInterface_PTR() *QDesignerFormEditorInterface {
	return ptr
}

func (ptr *QDesignerFormEditorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerFormEditorInterface(ptr QDesignerFormEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormEditorInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerFormEditorInterface) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerFormEditorInterface) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerFormEditorInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerFormEditorInterface) {
	n = new(QDesignerFormEditorInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerFormEditorInterface")
	return
}
func NewQDesignerFormEditorInterface(parent core.QObject_ITF) *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQDesignerFormEditorInterface", "", parent}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerFormEditorInterface) ActionEditor() *QDesignerActionEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEditor"}).(*QDesignerActionEditorInterface)
}

func (ptr *QDesignerFormEditorInterface) ExtensionManager() *QExtensionManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExtensionManager"}).(*QExtensionManager)
}

func (ptr *QDesignerFormEditorInterface) FormWindowManager() *QDesignerFormWindowManagerInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindowManager"}).(*QDesignerFormWindowManagerInterface)
}

func (ptr *QDesignerFormEditorInterface) ObjectInspector() *QDesignerObjectInspectorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ObjectInspector"}).(*QDesignerObjectInspectorInterface)
}

func (ptr *QDesignerFormEditorInterface) PropertyEditor() *QDesignerPropertyEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyEditor"}).(*QDesignerPropertyEditorInterface)
}

func (ptr *QDesignerFormEditorInterface) SetActionEditor(actionEditor QDesignerActionEditorInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActionEditor", actionEditor})
}

func (ptr *QDesignerFormEditorInterface) SetObjectInspector(objectInspector QDesignerObjectInspectorInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetObjectInspector", objectInspector})
}

func (ptr *QDesignerFormEditorInterface) SetPropertyEditor(propertyEditor QDesignerPropertyEditorInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPropertyEditor", propertyEditor})
}

func (ptr *QDesignerFormEditorInterface) SetWidgetBox(widgetBox QDesignerWidgetBoxInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWidgetBox", widgetBox})
}

func (ptr *QDesignerFormEditorInterface) TopLevel() *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TopLevel"}).(*widgets.QWidget)
}

func (ptr *QDesignerFormEditorInterface) WidgetBox() *QDesignerWidgetBoxInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WidgetBox"}).(*QDesignerWidgetBoxInterface)
}

func (ptr *QDesignerFormEditorInterface) ConnectDestroyQDesignerFormEditorInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerFormEditorInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormEditorInterface) DisconnectDestroyQDesignerFormEditorInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerFormEditorInterface"})
}

func (ptr *QDesignerFormEditorInterface) DestroyQDesignerFormEditorInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormEditorInterface"})
}

func (ptr *QDesignerFormEditorInterface) DestroyQDesignerFormEditorInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormEditorInterfaceDefault"})
}

func (ptr *QDesignerFormEditorInterface) __optionsPages_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__optionsPages_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pluginInstances_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pluginInstances_setList", i})
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pluginInstances_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __setOptionsPages_optionsPages_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setOptionsPages_optionsPages_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormEditorInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerFormEditorInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormEditorInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerFormEditorInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerFormEditorInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerFormEditorInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerFormEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerFormEditorInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerFormEditorInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerFormEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerFormEditorInterface) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDesignerFormEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerFormEditorInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerFormEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerFormWindow struct {
	widgets.QWidget
}

type QDesignerFormWindow_ITF interface {
	widgets.QWidget_ITF
	QDesignerFormWindow_PTR() *QDesignerFormWindow
}

func (ptr *QDesignerFormWindow) QDesignerFormWindow_PTR() *QDesignerFormWindow {
	return ptr
}

func (ptr *QDesignerFormWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerFormWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerFormWindow(ptr QDesignerFormWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindow_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerFormWindow) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerFormWindow) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerFormWindowFromPointer(ptr unsafe.Pointer) (n *QDesignerFormWindow) {
	n = new(QDesignerFormWindow)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerFormWindow")
	return
}

func (ptr *QDesignerFormWindow) DestroyQDesignerFormWindow() {
}

type QDesignerFormWindowCursorInterface struct {
	internal.Internal
}

type QDesignerFormWindowCursorInterface_ITF interface {
	QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface
}

func (ptr *QDesignerFormWindowCursorInterface) QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface {
	return ptr
}

func (ptr *QDesignerFormWindowCursorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerFormWindowCursorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerFormWindowCursorInterface(ptr QDesignerFormWindowCursorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowCursorInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerFormWindowCursorInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerFormWindowCursorInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerFormWindowCursorInterface) {
	n = new(QDesignerFormWindowCursorInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerFormWindowCursorInterface")
	return
}

//go:generate stringer -type=QDesignerFormWindowCursorInterface__MoveOperation
//QDesignerFormWindowCursorInterface::MoveOperation
type QDesignerFormWindowCursorInterface__MoveOperation int64

const (
	QDesignerFormWindowCursorInterface__NoMove QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(0)
	QDesignerFormWindowCursorInterface__Start  QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(1)
	QDesignerFormWindowCursorInterface__End    QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(2)
	QDesignerFormWindowCursorInterface__Next   QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(3)
	QDesignerFormWindowCursorInterface__Prev   QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(4)
	QDesignerFormWindowCursorInterface__Left   QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(5)
	QDesignerFormWindowCursorInterface__Right  QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(6)
	QDesignerFormWindowCursorInterface__Up     QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(7)
	QDesignerFormWindowCursorInterface__Down   QDesignerFormWindowCursorInterface__MoveOperation = QDesignerFormWindowCursorInterface__MoveOperation(8)
)

//go:generate stringer -type=QDesignerFormWindowCursorInterface__MoveMode
//QDesignerFormWindowCursorInterface::MoveMode
type QDesignerFormWindowCursorInterface__MoveMode int64

const (
	QDesignerFormWindowCursorInterface__MoveAnchor QDesignerFormWindowCursorInterface__MoveMode = QDesignerFormWindowCursorInterface__MoveMode(0)
	QDesignerFormWindowCursorInterface__KeepAnchor QDesignerFormWindowCursorInterface__MoveMode = QDesignerFormWindowCursorInterface__MoveMode(1)
)

func (ptr *QDesignerFormWindowCursorInterface) ConnectCurrent(f func() *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCurrent", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectCurrent() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCurrent"})
}

func (ptr *QDesignerFormWindowCursorInterface) Current() *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Current"}).(*widgets.QWidget)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectFormWindow(f func() *QDesignerFormWindowInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindow"})
}

func (ptr *QDesignerFormWindowCursorInterface) FormWindow() *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindow"}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectHasSelection(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasSelection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectHasSelection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasSelection"})
}

func (ptr *QDesignerFormWindowCursorInterface) HasSelection() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasSelection"}).(bool)
}

func (ptr *QDesignerFormWindowCursorInterface) IsWidgetSelected(widget widgets.QWidget_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsWidgetSelected", widget}).(bool)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectMovePosition(f func(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMovePosition", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectMovePosition() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMovePosition"})
}

func (ptr *QDesignerFormWindowCursorInterface) MovePosition(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MovePosition", operation, mode}).(bool)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectPosition(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPosition", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectPosition() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPosition"})
}

func (ptr *QDesignerFormWindowCursorInterface) Position() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Position"}).(float64))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectResetWidgetProperty(f func(widget *widgets.QWidget, name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResetWidgetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectResetWidgetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResetWidgetProperty"})
}

func (ptr *QDesignerFormWindowCursorInterface) ResetWidgetProperty(widget widgets.QWidget_ITF, name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetWidgetProperty", widget, name})
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidget(f func(index int) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedWidget"})
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidget(index int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedWidget", index}).(*widgets.QWidget)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidgetCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectedWidgetCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidgetCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectedWidgetCount"})
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidgetCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectedWidgetCount"}).(float64))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetPosition(f func(position int, mode QDesignerFormWindowCursorInterface__MoveMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPosition", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetPosition() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPosition"})
}

func (ptr *QDesignerFormWindowCursorInterface) SetPosition(position int, mode QDesignerFormWindowCursorInterface__MoveMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPosition", position, mode})
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetProperty(f func(name string, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QDesignerFormWindowCursorInterface) SetProperty(name string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", name, value})
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetWidgetProperty(f func(widget *widgets.QWidget, name string, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetWidgetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetWidgetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetWidgetProperty"})
}

func (ptr *QDesignerFormWindowCursorInterface) SetWidgetProperty(widget widgets.QWidget_ITF, name string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWidgetProperty", widget, name, value})
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidget(f func(index int) *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidget"})
}

func (ptr *QDesignerFormWindowCursorInterface) Widget(index int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Widget", index}).(*widgets.QWidget)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidgetCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidgetCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidgetCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidgetCount"})
}

func (ptr *QDesignerFormWindowCursorInterface) WidgetCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WidgetCount"}).(float64))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectDestroyQDesignerFormWindowCursorInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerFormWindowCursorInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectDestroyQDesignerFormWindowCursorInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerFormWindowCursorInterface"})
}

func (ptr *QDesignerFormWindowCursorInterface) DestroyQDesignerFormWindowCursorInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowCursorInterface"})
}

func (ptr *QDesignerFormWindowCursorInterface) DestroyQDesignerFormWindowCursorInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowCursorInterfaceDefault"})
}

type QDesignerFormWindowInterface struct {
	widgets.QWidget
}

type QDesignerFormWindowInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerFormWindowInterface_PTR() *QDesignerFormWindowInterface
}

func (ptr *QDesignerFormWindowInterface) QDesignerFormWindowInterface_PTR() *QDesignerFormWindowInterface {
	return ptr
}

func (ptr *QDesignerFormWindowInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerFormWindowInterface(ptr QDesignerFormWindowInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerFormWindowInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerFormWindowInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerFormWindowInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerFormWindowInterface) {
	n = new(QDesignerFormWindowInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerFormWindowInterface")
	return
}

//go:generate stringer -type=QDesignerFormWindowInterface__FeatureFlag
//QDesignerFormWindowInterface::FeatureFlag
type QDesignerFormWindowInterface__FeatureFlag int64

const (
	QDesignerFormWindowInterface__EditFeature     QDesignerFormWindowInterface__FeatureFlag = QDesignerFormWindowInterface__FeatureFlag(0x01)
	QDesignerFormWindowInterface__GridFeature     QDesignerFormWindowInterface__FeatureFlag = QDesignerFormWindowInterface__FeatureFlag(0x02)
	QDesignerFormWindowInterface__TabOrderFeature QDesignerFormWindowInterface__FeatureFlag = QDesignerFormWindowInterface__FeatureFlag(0x04)
	QDesignerFormWindowInterface__DefaultFeature  QDesignerFormWindowInterface__FeatureFlag = QDesignerFormWindowInterface__FeatureFlag(QDesignerFormWindowInterface__EditFeature | QDesignerFormWindowInterface__GridFeature)
)

//go:generate stringer -type=QDesignerFormWindowInterface__ResourceFileSaveMode
//QDesignerFormWindowInterface::ResourceFileSaveMode
type QDesignerFormWindowInterface__ResourceFileSaveMode int64

const (
	QDesignerFormWindowInterface__SaveAllResourceFiles      QDesignerFormWindowInterface__ResourceFileSaveMode = QDesignerFormWindowInterface__ResourceFileSaveMode(0)
	QDesignerFormWindowInterface__SaveOnlyUsedResourceFiles QDesignerFormWindowInterface__ResourceFileSaveMode = QDesignerFormWindowInterface__ResourceFileSaveMode(1)
	QDesignerFormWindowInterface__DontSaveResourceFiles     QDesignerFormWindowInterface__ResourceFileSaveMode = QDesignerFormWindowInterface__ResourceFileSaveMode(2)
)

func (ptr *QDesignerFormWindowInterface) ConnectAboutToUnmanageWidget(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAboutToUnmanageWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectAboutToUnmanageWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAboutToUnmanageWidget"})
}

func (ptr *QDesignerFormWindowInterface) AboutToUnmanageWidget(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AboutToUnmanageWidget", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectAbsoluteDir(f func() *core.QDir) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbsoluteDir", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectAbsoluteDir() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbsoluteDir"})
}

func (ptr *QDesignerFormWindowInterface) AbsoluteDir() *core.QDir {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbsoluteDir"}).(*core.QDir)
}

func (ptr *QDesignerFormWindowInterface) ConnectActivateResourceFilePaths(f func(paths []string, errorCount int, errorMessages string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActivateResourceFilePaths", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivateResourceFilePaths() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActivateResourceFilePaths"})
}

func (ptr *QDesignerFormWindowInterface) ActivateResourceFilePaths(paths []string, errorCount int, errorMessages string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivateResourceFilePaths", paths, errorCount, errorMessages})
}

func (ptr *QDesignerFormWindowInterface) ActivateResourceFilePathsDefault(paths []string, errorCount int, errorMessages string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActivateResourceFilePathsDefault", paths, errorCount, errorMessages})
}

func (ptr *QDesignerFormWindowInterface) ConnectActivated(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActivated"})
}

func (ptr *QDesignerFormWindowInterface) Activated(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Activated", widget})
}

func (ptr *QDesignerFormWindowInterface) ActiveResourceFilePaths() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveResourceFilePaths"}).([]string)
}

func (ptr *QDesignerFormWindowInterface) ConnectAddResourceFile(f func(path string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAddResourceFile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectAddResourceFile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAddResourceFile"})
}

func (ptr *QDesignerFormWindowInterface) AddResourceFile(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddResourceFile", path})
}

func (ptr *QDesignerFormWindowInterface) ConnectAuthor(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAuthor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectAuthor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAuthor"})
}

func (ptr *QDesignerFormWindowInterface) Author() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Author"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectChanged"})
}

func (ptr *QDesignerFormWindowInterface) Changed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Changed"})
}

func (ptr *QDesignerFormWindowInterface) ConnectCheckContents(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCheckContents", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectCheckContents() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCheckContents"})
}

func (ptr *QDesignerFormWindowInterface) CheckContents() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CheckContents"}).([]string)
}

func (ptr *QDesignerFormWindowInterface) ConnectClearSelection(f func(update bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClearSelection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectClearSelection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClearSelection"})
}

func (ptr *QDesignerFormWindowInterface) ClearSelection(update bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearSelection", update})
}

func (ptr *QDesignerFormWindowInterface) ConnectComment(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectComment"})
}

func (ptr *QDesignerFormWindowInterface) Comment() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Comment"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectContents(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectContents", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectContents() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectContents"})
}

func (ptr *QDesignerFormWindowInterface) Contents() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Contents"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCore"})
}

func (ptr *QDesignerFormWindowInterface) Core() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Core"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerFormWindowInterface) CoreDefault() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoreDefault"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerFormWindowInterface) ConnectCursor(f func() *QDesignerFormWindowCursorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCursor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectCursor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCursor"})
}

func (ptr *QDesignerFormWindowInterface) Cursor() *QDesignerFormWindowCursorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Cursor"}).(*QDesignerFormWindowCursorInterface)
}

func (ptr *QDesignerFormWindowInterface) ConnectEmitSelectionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEmitSelectionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectEmitSelectionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEmitSelectionChanged"})
}

func (ptr *QDesignerFormWindowInterface) EmitSelectionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EmitSelectionChanged"})
}

func (ptr *QDesignerFormWindowInterface) ConnectExportMacro(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExportMacro", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectExportMacro() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExportMacro"})
}

func (ptr *QDesignerFormWindowInterface) ExportMacro() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExportMacro"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectFeatureChanged(f func(feature QDesignerFormWindowInterface__FeatureFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeatureChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatureChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeatureChanged"})
}

func (ptr *QDesignerFormWindowInterface) FeatureChanged(feature QDesignerFormWindowInterface__FeatureFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FeatureChanged", feature})
}

func (ptr *QDesignerFormWindowInterface) ConnectFeatures(f func() QDesignerFormWindowInterface__FeatureFlag) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFeatures", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatures() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFeatures"})
}

func (ptr *QDesignerFormWindowInterface) Features() QDesignerFormWindowInterface__FeatureFlag {

	return QDesignerFormWindowInterface__FeatureFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Features"}).(float64))
}

func (ptr *QDesignerFormWindowInterface) ConnectFileName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFileName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFileName"})
}

func (ptr *QDesignerFormWindowInterface) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectFileNameChanged(f func(fileName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFileNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFileNameChanged"})
}

func (ptr *QDesignerFormWindowInterface) FileNameChanged(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileNameChanged", fileName})
}

func QDesignerFormWindowInterface_FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.QDesignerFormWindowInterface_FindFormWindow", "", widget}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.QDesignerFormWindowInterface_FindFormWindow", "", widget}).(*QDesignerFormWindowInterface)
}

func QDesignerFormWindowInterface_FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.QDesignerFormWindowInterface_FindFormWindow2", "", object}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.QDesignerFormWindowInterface_FindFormWindow2", "", object}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowInterface) ConnectFormContainer(f func() *widgets.QWidget) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormContainer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectFormContainer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormContainer"})
}

func (ptr *QDesignerFormWindowInterface) FormContainer() *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormContainer"}).(*widgets.QWidget)
}

func (ptr *QDesignerFormWindowInterface) ConnectGeometryChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGeometryChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectGeometryChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGeometryChanged"})
}

func (ptr *QDesignerFormWindowInterface) GeometryChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "GeometryChanged"})
}

func (ptr *QDesignerFormWindowInterface) ConnectGrid(f func() *core.QPoint) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectGrid", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectGrid() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectGrid"})
}

func (ptr *QDesignerFormWindowInterface) Grid() *core.QPoint {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Grid"}).(*core.QPoint)
}

func (ptr *QDesignerFormWindowInterface) ConnectHasFeature(f func(feature QDesignerFormWindowInterface__FeatureFlag) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasFeature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectHasFeature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasFeature"})
}

func (ptr *QDesignerFormWindowInterface) HasFeature(feature QDesignerFormWindowInterface__FeatureFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasFeature", feature}).(bool)
}

func (ptr *QDesignerFormWindowInterface) ConnectIncludeHints(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIncludeHints", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectIncludeHints() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIncludeHints"})
}

func (ptr *QDesignerFormWindowInterface) IncludeHints() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludeHints"}).([]string)
}

func (ptr *QDesignerFormWindowInterface) ConnectIsDirty(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsDirty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsDirty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsDirty"})
}

func (ptr *QDesignerFormWindowInterface) IsDirty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsDirty"}).(bool)
}

func (ptr *QDesignerFormWindowInterface) ConnectIsManaged(f func(widget *widgets.QWidget) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsManaged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsManaged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsManaged"})
}

func (ptr *QDesignerFormWindowInterface) IsManaged(widget widgets.QWidget_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsManaged", widget}).(bool)
}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutDefault(f func(margin int, spacing int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLayoutDefault", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutDefault() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLayoutDefault"})
}

func (ptr *QDesignerFormWindowInterface) LayoutDefault(margin int, spacing int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LayoutDefault", margin, spacing})
}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutFunction(f func(margin string, spacing string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLayoutFunction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutFunction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLayoutFunction"})
}

func (ptr *QDesignerFormWindowInterface) LayoutFunction(margin string, spacing string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LayoutFunction", margin, spacing})
}

func (ptr *QDesignerFormWindowInterface) ConnectMainContainerChanged(f func(mainContainer *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMainContainerChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectMainContainerChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMainContainerChanged"})
}

func (ptr *QDesignerFormWindowInterface) MainContainerChanged(mainContainer widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MainContainerChanged", mainContainer})
}

func (ptr *QDesignerFormWindowInterface) ConnectManageWidget(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectManageWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectManageWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectManageWidget"})
}

func (ptr *QDesignerFormWindowInterface) ManageWidget(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ManageWidget", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectObjectRemoved(f func(object *core.QObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectObjectRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectObjectRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectObjectRemoved"})
}

func (ptr *QDesignerFormWindowInterface) ObjectRemoved(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ObjectRemoved", object})
}

func (ptr *QDesignerFormWindowInterface) ConnectPixmapFunction(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPixmapFunction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectPixmapFunction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPixmapFunction"})
}

func (ptr *QDesignerFormWindowInterface) PixmapFunction() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PixmapFunction"}).(string)
}

func (ptr *QDesignerFormWindowInterface) ConnectRemoveResourceFile(f func(path string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemoveResourceFile", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectRemoveResourceFile() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoveResourceFile"})
}

func (ptr *QDesignerFormWindowInterface) RemoveResourceFile(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveResourceFile", path})
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFileSaveMode(f func() QDesignerFormWindowInterface__ResourceFileSaveMode) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResourceFileSaveMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFileSaveMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResourceFileSaveMode"})
}

func (ptr *QDesignerFormWindowInterface) ResourceFileSaveMode() QDesignerFormWindowInterface__ResourceFileSaveMode {

	return QDesignerFormWindowInterface__ResourceFileSaveMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResourceFileSaveMode"}).(float64))
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFiles(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResourceFiles", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFiles() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResourceFiles"})
}

func (ptr *QDesignerFormWindowInterface) ResourceFiles() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResourceFiles"}).([]string)
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFilesChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResourceFilesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFilesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResourceFilesChanged"})
}

func (ptr *QDesignerFormWindowInterface) ResourceFilesChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResourceFilesChanged"})
}

func (ptr *QDesignerFormWindowInterface) ConnectSelectWidget(f func(widget *widgets.QWidget, sele bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectWidget"})
}

func (ptr *QDesignerFormWindowInterface) SelectWidget(widget widgets.QWidget_ITF, sele bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectWidget", widget, sele})
}

func (ptr *QDesignerFormWindowInterface) ConnectSelectionChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSelectionChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectionChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSelectionChanged"})
}

func (ptr *QDesignerFormWindowInterface) SelectionChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SelectionChanged"})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetAuthor(f func(author string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetAuthor", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetAuthor() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetAuthor"})
}

func (ptr *QDesignerFormWindowInterface) SetAuthor(author string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAuthor", author})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetComment(f func(comment string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetComment", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetComment() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetComment"})
}

func (ptr *QDesignerFormWindowInterface) SetComment(comment string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetComment", comment})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents(f func(device *core.QIODevice, errorMessage string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetContents", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetContents"})
}

func (ptr *QDesignerFormWindowInterface) SetContents(device core.QIODevice_ITF, errorMessage string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContents", device, errorMessage}).(bool)
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents2(f func(contents string) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetContents2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetContents2"})
}

func (ptr *QDesignerFormWindowInterface) SetContents2(contents string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContents2", contents}).(bool)
}

func (ptr *QDesignerFormWindowInterface) ConnectSetDirty(f func(dirty bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetDirty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetDirty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetDirty"})
}

func (ptr *QDesignerFormWindowInterface) SetDirty(dirty bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDirty", dirty})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetExportMacro(f func(exportMacro string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetExportMacro", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetExportMacro() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetExportMacro"})
}

func (ptr *QDesignerFormWindowInterface) SetExportMacro(exportMacro string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExportMacro", exportMacro})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetFeatures(f func(features QDesignerFormWindowInterface__FeatureFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFeatures", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFeatures() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFeatures"})
}

func (ptr *QDesignerFormWindowInterface) SetFeatures(features QDesignerFormWindowInterface__FeatureFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFeatures", features})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetFileName(f func(fileName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFileName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFileName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFileName"})
}

func (ptr *QDesignerFormWindowInterface) SetFileName(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFileName", fileName})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetGrid(f func(grid *core.QPoint)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetGrid", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetGrid() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetGrid"})
}

func (ptr *QDesignerFormWindowInterface) SetGrid(grid core.QPoint_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetGrid", grid})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetIncludeHints(f func(includeHints []string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetIncludeHints", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetIncludeHints() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetIncludeHints"})
}

func (ptr *QDesignerFormWindowInterface) SetIncludeHints(includeHints []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncludeHints", includeHints})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutDefault(f func(margin int, spacing int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLayoutDefault", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutDefault() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLayoutDefault"})
}

func (ptr *QDesignerFormWindowInterface) SetLayoutDefault(margin int, spacing int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLayoutDefault", margin, spacing})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutFunction(f func(margin string, spacing string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLayoutFunction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutFunction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLayoutFunction"})
}

func (ptr *QDesignerFormWindowInterface) SetLayoutFunction(margin string, spacing string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLayoutFunction", margin, spacing})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetMainContainer(f func(mainContainer *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetMainContainer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetMainContainer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetMainContainer"})
}

func (ptr *QDesignerFormWindowInterface) SetMainContainer(mainContainer widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMainContainer", mainContainer})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetPixmapFunction(f func(pixmapFunction string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPixmapFunction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetPixmapFunction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPixmapFunction"})
}

func (ptr *QDesignerFormWindowInterface) SetPixmapFunction(pixmapFunction string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPixmapFunction", pixmapFunction})
}

func (ptr *QDesignerFormWindowInterface) ConnectSetResourceFileSaveMode(f func(behavior QDesignerFormWindowInterface__ResourceFileSaveMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetResourceFileSaveMode", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetResourceFileSaveMode() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetResourceFileSaveMode"})
}

func (ptr *QDesignerFormWindowInterface) SetResourceFileSaveMode(behavior QDesignerFormWindowInterface__ResourceFileSaveMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetResourceFileSaveMode", behavior})
}

func (ptr *QDesignerFormWindowInterface) ConnectUnmanageWidget(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUnmanageWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectUnmanageWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUnmanageWidget"})
}

func (ptr *QDesignerFormWindowInterface) UnmanageWidget(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UnmanageWidget", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetManaged(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidgetManaged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetManaged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidgetManaged"})
}

func (ptr *QDesignerFormWindowInterface) WidgetManaged(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WidgetManaged", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetRemoved(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidgetRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidgetRemoved"})
}

func (ptr *QDesignerFormWindowInterface) WidgetRemoved(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WidgetRemoved", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetUnmanaged(f func(widget *widgets.QWidget)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWidgetUnmanaged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetUnmanaged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWidgetUnmanaged"})
}

func (ptr *QDesignerFormWindowInterface) WidgetUnmanaged(widget widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WidgetUnmanaged", widget})
}

func (ptr *QDesignerFormWindowInterface) ConnectDestroyQDesignerFormWindowInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerFormWindowInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowInterface) DisconnectDestroyQDesignerFormWindowInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerFormWindowInterface"})
}

func (ptr *QDesignerFormWindowInterface) DestroyQDesignerFormWindowInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowInterface"})
}

func (ptr *QDesignerFormWindowInterface) DestroyQDesignerFormWindowInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowInterfaceDefault"})
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_atList(i int) *widgets.QWidget {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__simplifySelection_widgets_atList", i}).(*widgets.QWidget)
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_setList(i widgets.QWidget_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__simplifySelection_widgets_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__simplifySelection_widgets_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerFormWindowInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerFormWindowInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerFormWindowInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerFormWindowInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerFormWindowInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerFormWindowInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerFormWindowInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerFormWindowInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerFormWindowInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerFormWindowInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerFormWindowInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerFormWindowInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerFormWindowInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerFormWindowInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerFormWindowInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerFormWindowInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerFormWindowInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerFormWindowInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerFormWindowInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerFormWindowInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerFormWindowInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerFormWindowInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerFormWindowInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerFormWindowInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerFormWindowInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerFormWindowInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerFormWindowInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerFormWindowInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerFormWindowInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerFormWindowInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerFormWindowInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerFormWindowInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerFormWindowInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerFormWindowInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerFormWindowInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerFormWindowInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerFormWindowInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerFormWindowInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerFormWindowInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerFormWindowInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerFormWindowManagerInterface struct {
	core.QObject
}

type QDesignerFormWindowManagerInterface_ITF interface {
	core.QObject_ITF
	QDesignerFormWindowManagerInterface_PTR() *QDesignerFormWindowManagerInterface
}

func (ptr *QDesignerFormWindowManagerInterface) QDesignerFormWindowManagerInterface_PTR() *QDesignerFormWindowManagerInterface {
	return ptr
}

func (ptr *QDesignerFormWindowManagerInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerFormWindowManagerInterface(ptr QDesignerFormWindowManagerInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowManagerInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerFormWindowManagerInterface) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerFormWindowManagerInterface) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerFormWindowManagerInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerFormWindowManagerInterface) {
	n = new(QDesignerFormWindowManagerInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerFormWindowManagerInterface")
	return
}

//go:generate stringer -type=QDesignerFormWindowManagerInterface__Action
//QDesignerFormWindowManagerInterface::Action
type QDesignerFormWindowManagerInterface__Action int64

const (
	QDesignerFormWindowManagerInterface__CutAction                      QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(100)
	QDesignerFormWindowManagerInterface__CopyAction                     QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(101)
	QDesignerFormWindowManagerInterface__PasteAction                    QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(102)
	QDesignerFormWindowManagerInterface__DeleteAction                   QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(103)
	QDesignerFormWindowManagerInterface__SelectAllAction                QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(104)
	QDesignerFormWindowManagerInterface__LowerAction                    QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(200)
	QDesignerFormWindowManagerInterface__RaiseAction                    QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(201)
	QDesignerFormWindowManagerInterface__UndoAction                     QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(300)
	QDesignerFormWindowManagerInterface__RedoAction                     QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(301)
	QDesignerFormWindowManagerInterface__HorizontalLayoutAction         QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(400)
	QDesignerFormWindowManagerInterface__VerticalLayoutAction           QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(401)
	QDesignerFormWindowManagerInterface__SplitHorizontalAction          QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(402)
	QDesignerFormWindowManagerInterface__SplitVerticalAction            QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(403)
	QDesignerFormWindowManagerInterface__GridLayoutAction               QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(404)
	QDesignerFormWindowManagerInterface__FormLayoutAction               QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(405)
	QDesignerFormWindowManagerInterface__BreakLayoutAction              QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(406)
	QDesignerFormWindowManagerInterface__AdjustSizeAction               QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(407)
	QDesignerFormWindowManagerInterface__SimplifyLayoutAction           QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(408)
	QDesignerFormWindowManagerInterface__DefaultPreviewAction           QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(500)
	QDesignerFormWindowManagerInterface__FormWindowSettingsDialogAction QDesignerFormWindowManagerInterface__Action = QDesignerFormWindowManagerInterface__Action(600)
)

//go:generate stringer -type=QDesignerFormWindowManagerInterface__ActionGroup
//QDesignerFormWindowManagerInterface::ActionGroup
type QDesignerFormWindowManagerInterface__ActionGroup int64

const (
	QDesignerFormWindowManagerInterface__StyledPreviewActionGroup QDesignerFormWindowManagerInterface__ActionGroup = QDesignerFormWindowManagerInterface__ActionGroup(100)
)

func (ptr *QDesignerFormWindowManagerInterface) ConnectAction(f func(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAction"})
}

func (ptr *QDesignerFormWindowManagerInterface) Action(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Action", action}).(*widgets.QAction)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActionGroup(f func(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActionGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActionGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActionGroup"})
}

func (ptr *QDesignerFormWindowManagerInterface) ActionGroup(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionGroup", actionGroup}).(*widgets.QActionGroup)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindow(f func() *QDesignerFormWindowInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindow() *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveFormWindow"}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindowChanged(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectActiveFormWindowChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindowChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectActiveFormWindowChanged"})
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindowChanged(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveFormWindowChanged", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectAddFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAddFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAddFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAddFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) AddFormWindow(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddFormWindow", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCloseAllPreviews(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCloseAllPreviews", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCloseAllPreviews() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCloseAllPreviews"})
}

func (ptr *QDesignerFormWindowManagerInterface) CloseAllPreviews() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseAllPreviews"})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCore"})
}

func (ptr *QDesignerFormWindowManagerInterface) Core() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Core"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreateFormWindow(f func(parent *widgets.QWidget, flags core.Qt__WindowType) *QDesignerFormWindowInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreateFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) CreateFormWindow(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateFormWindow", parent, flags}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreatePreviewPixmap(f func() *gui.QPixmap) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreatePreviewPixmap", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreatePreviewPixmap() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreatePreviewPixmap"})
}

func (ptr *QDesignerFormWindowManagerInterface) CreatePreviewPixmap() *gui.QPixmap {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreatePreviewPixmap"}).(*gui.QPixmap)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindow(f func(index int) *QDesignerFormWindowInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindow(index int) *QDesignerFormWindowInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindow", index}).(*QDesignerFormWindowInterface)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowAdded(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindowAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindowAdded"})
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowAdded(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindowAdded", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindowCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindowCount"})
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowCount() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindowCount"}).(float64))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowRemoved(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindowRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindowRemoved"})
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowRemoved(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindowRemoved", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowSettingsChanged(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFormWindowSettingsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowSettingsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFormWindowSettingsChanged"})
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowSettingsChanged(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FormWindowSettingsChanged", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectRemoveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemoveFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectRemoveFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemoveFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) RemoveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveFormWindow", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectSetActiveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetActiveFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectSetActiveFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetActiveFormWindow"})
}

func (ptr *QDesignerFormWindowManagerInterface) SetActiveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetActiveFormWindow", formWindow})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPluginDialog(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShowPluginDialog", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPluginDialog() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShowPluginDialog"})
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPluginDialog() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowPluginDialog"})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPreview(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectShowPreview", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPreview() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectShowPreview"})
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPreview() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowPreview"})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectDestroyQDesignerFormWindowManagerInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerFormWindowManagerInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectDestroyQDesignerFormWindowManagerInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerFormWindowManagerInterface"})
}

func (ptr *QDesignerFormWindowManagerInterface) DestroyQDesignerFormWindowManagerInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowManagerInterface"})
}

func (ptr *QDesignerFormWindowManagerInterface) DestroyQDesignerFormWindowManagerInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerFormWindowManagerInterfaceDefault"})
}

func (ptr *QDesignerFormWindowManagerInterface) __dragItems_item_list_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dragItems_item_list_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowManagerInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowManagerInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerFormWindowManagerInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerFormWindowManagerInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerFormWindowManagerInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerFormWindowManagerInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerFormWindowManagerInterface) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDesignerFormWindowManagerInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerFormWindowManagerInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerFormWindowManagerInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerLanguageExtension struct {
	internal.Internal
}

type QDesignerLanguageExtension_ITF interface {
	QDesignerLanguageExtension_PTR() *QDesignerLanguageExtension
}

func (ptr *QDesignerLanguageExtension) QDesignerLanguageExtension_PTR() *QDesignerLanguageExtension {
	return ptr
}

func (ptr *QDesignerLanguageExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerLanguageExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerLanguageExtension(ptr QDesignerLanguageExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerLanguageExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerLanguageExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerLanguageExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerLanguageExtension) {
	n = new(QDesignerLanguageExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerLanguageExtension")
	return
}

func (ptr *QDesignerLanguageExtension) DestroyQDesignerLanguageExtension() {
}

type QDesignerMemberSheetExtension struct {
	internal.Internal
}

type QDesignerMemberSheetExtension_ITF interface {
	QDesignerMemberSheetExtension_PTR() *QDesignerMemberSheetExtension
}

func (ptr *QDesignerMemberSheetExtension) QDesignerMemberSheetExtension_PTR() *QDesignerMemberSheetExtension {
	return ptr
}

func (ptr *QDesignerMemberSheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerMemberSheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerMemberSheetExtension(ptr QDesignerMemberSheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerMemberSheetExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerMemberSheetExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerMemberSheetExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerMemberSheetExtension) {
	n = new(QDesignerMemberSheetExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerMemberSheetExtension")
	return
}
func (ptr *QDesignerMemberSheetExtension) ConnectCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCount"})
}

func (ptr *QDesignerMemberSheetExtension) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QDesignerMemberSheetExtension) ConnectDeclaredInClass(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeclaredInClass", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectDeclaredInClass() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeclaredInClass"})
}

func (ptr *QDesignerMemberSheetExtension) DeclaredInClass(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeclaredInClass", index}).(string)
}

func (ptr *QDesignerMemberSheetExtension) ConnectIndexOf(f func(name string) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexOf", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIndexOf() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexOf"})
}

func (ptr *QDesignerMemberSheetExtension) IndexOf(name string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexOf", name}).(float64))
}

func (ptr *QDesignerMemberSheetExtension) ConnectInheritedFromWidget(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInheritedFromWidget", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectInheritedFromWidget() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInheritedFromWidget"})
}

func (ptr *QDesignerMemberSheetExtension) InheritedFromWidget(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InheritedFromWidget", index}).(bool)
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSignal(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsSignal", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSignal() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsSignal"})
}

func (ptr *QDesignerMemberSheetExtension) IsSignal(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSignal", index}).(bool)
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSlot(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsSlot", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSlot() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsSlot"})
}

func (ptr *QDesignerMemberSheetExtension) IsSlot(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSlot", index}).(bool)
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsVisible(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsVisible", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsVisible() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsVisible"})
}

func (ptr *QDesignerMemberSheetExtension) IsVisible(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible", index}).(bool)
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberGroup(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMemberGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMemberGroup"})
}

func (ptr *QDesignerMemberSheetExtension) MemberGroup(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MemberGroup", index}).(string)
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberName(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMemberName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMemberName"})
}

func (ptr *QDesignerMemberSheetExtension) MemberName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MemberName", index}).(string)
}

func (ptr *QDesignerMemberSheetExtension) ConnectParameterNames(f func(index int) []*core.QByteArray) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParameterNames", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectParameterNames() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParameterNames"})
}

func (ptr *QDesignerMemberSheetExtension) ParameterNames(index int) []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParameterNames", index}).([]*core.QByteArray)
}

func (ptr *QDesignerMemberSheetExtension) ConnectParameterTypes(f func(index int) []*core.QByteArray) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectParameterTypes", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectParameterTypes() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectParameterTypes"})
}

func (ptr *QDesignerMemberSheetExtension) ParameterTypes(index int) []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ParameterTypes", index}).([]*core.QByteArray)
}

func (ptr *QDesignerMemberSheetExtension) ConnectSetMemberGroup(f func(index int, group string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetMemberGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetMemberGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetMemberGroup"})
}

func (ptr *QDesignerMemberSheetExtension) SetMemberGroup(index int, group string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMemberGroup", index, group})
}

func (ptr *QDesignerMemberSheetExtension) ConnectSetVisible(f func(index int, visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVisible", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetVisible() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVisible"})
}

func (ptr *QDesignerMemberSheetExtension) SetVisible(index int, visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", index, visible})
}

func (ptr *QDesignerMemberSheetExtension) ConnectSignature(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSignature", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSignature() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSignature"})
}

func (ptr *QDesignerMemberSheetExtension) Signature(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Signature", index}).(string)
}

func (ptr *QDesignerMemberSheetExtension) ConnectDestroyQDesignerMemberSheetExtension(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerMemberSheetExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerMemberSheetExtension) DisconnectDestroyQDesignerMemberSheetExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerMemberSheetExtension"})
}

func (ptr *QDesignerMemberSheetExtension) DestroyQDesignerMemberSheetExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerMemberSheetExtension"})
}

func (ptr *QDesignerMemberSheetExtension) DestroyQDesignerMemberSheetExtensionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerMemberSheetExtensionDefault"})
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterNames_setList", i})
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterTypes_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterTypes_setList", i})
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parameterTypes_newList"}).(unsafe.Pointer)
}

type QDesignerObjectInspectorInterface struct {
	widgets.QWidget
}

type QDesignerObjectInspectorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerObjectInspectorInterface_PTR() *QDesignerObjectInspectorInterface
}

func (ptr *QDesignerObjectInspectorInterface) QDesignerObjectInspectorInterface_PTR() *QDesignerObjectInspectorInterface {
	return ptr
}

func (ptr *QDesignerObjectInspectorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerObjectInspectorInterface(ptr QDesignerObjectInspectorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerObjectInspectorInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerObjectInspectorInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerObjectInspectorInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerObjectInspectorInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerObjectInspectorInterface) {
	n = new(QDesignerObjectInspectorInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerObjectInspectorInterface")
	return
}
func NewQDesignerObjectInspectorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerObjectInspectorInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQDesignerObjectInspectorInterface", "", parent, flags}).(*QDesignerObjectInspectorInterface)
}

func (ptr *QDesignerObjectInspectorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCore"})
}

func (ptr *QDesignerObjectInspectorInterface) Core() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Core"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerObjectInspectorInterface) CoreDefault() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoreDefault"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFormWindow", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetFormWindow() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFormWindow"})
}

func (ptr *QDesignerObjectInspectorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFormWindow", formWindow})
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDestroyQDesignerObjectInspectorInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerObjectInspectorInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDestroyQDesignerObjectInspectorInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerObjectInspectorInterface"})
}

func (ptr *QDesignerObjectInspectorInterface) DestroyQDesignerObjectInspectorInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerObjectInspectorInterface"})
}

func (ptr *QDesignerObjectInspectorInterface) DestroyQDesignerObjectInspectorInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerObjectInspectorInterfaceDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerObjectInspectorInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerObjectInspectorInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerObjectInspectorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerObjectInspectorInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerObjectInspectorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerObjectInspectorInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerObjectInspectorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerObjectInspectorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerObjectInspectorInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerObjectInspectorInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerObjectInspectorInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerObjectInspectorInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerObjectInspectorInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerObjectInspectorInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerObjectInspectorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerObjectInspectorInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerObjectInspectorInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerObjectInspectorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerObjectInspectorInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerObjectInspectorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerPropertyEditorInterface struct {
	widgets.QWidget
}

type QDesignerPropertyEditorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerPropertyEditorInterface_PTR() *QDesignerPropertyEditorInterface
}

func (ptr *QDesignerPropertyEditorInterface) QDesignerPropertyEditorInterface_PTR() *QDesignerPropertyEditorInterface {
	return ptr
}

func (ptr *QDesignerPropertyEditorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerPropertyEditorInterface(ptr QDesignerPropertyEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerPropertyEditorInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerPropertyEditorInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerPropertyEditorInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerPropertyEditorInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerPropertyEditorInterface) {
	n = new(QDesignerPropertyEditorInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerPropertyEditorInterface")
	return
}
func NewQDesignerPropertyEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerPropertyEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQDesignerPropertyEditorInterface", "", parent, flags}).(*QDesignerPropertyEditorInterface)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCore"})
}

func (ptr *QDesignerPropertyEditorInterface) Core() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Core"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerPropertyEditorInterface) CoreDefault() *QDesignerFormEditorInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CoreDefault"}).(*QDesignerFormEditorInterface)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCurrentPropertyName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCurrentPropertyName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCurrentPropertyName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCurrentPropertyName"})
}

func (ptr *QDesignerPropertyEditorInterface) CurrentPropertyName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CurrentPropertyName"}).(string)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectIsReadOnly(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsReadOnly", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectIsReadOnly() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsReadOnly"})
}

func (ptr *QDesignerPropertyEditorInterface) IsReadOnly() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsReadOnly"}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectObject(f func() *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectObject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectObject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectObject"})
}

func (ptr *QDesignerPropertyEditorInterface) Object() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Object"}).(*core.QObject)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectPropertyChanged(f func(name string, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPropertyChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectPropertyChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPropertyChanged"})
}

func (ptr *QDesignerPropertyEditorInterface) PropertyChanged(name string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyChanged", name, value})
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetObject(f func(object *core.QObject)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetObject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetObject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetObject"})
}

func (ptr *QDesignerPropertyEditorInterface) SetObject(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetObject", object})
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetPropertyValue(f func(name string, value *core.QVariant, changed bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPropertyValue", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetPropertyValue() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPropertyValue"})
}

func (ptr *QDesignerPropertyEditorInterface) SetPropertyValue(name string, value core.QVariant_ITF, changed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPropertyValue", name, value, changed})
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetReadOnly(f func(readOnly bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetReadOnly", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetReadOnly() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetReadOnly"})
}

func (ptr *QDesignerPropertyEditorInterface) SetReadOnly(readOnly bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadOnly", readOnly})
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDestroyQDesignerPropertyEditorInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerPropertyEditorInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDestroyQDesignerPropertyEditorInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerPropertyEditorInterface"})
}

func (ptr *QDesignerPropertyEditorInterface) DestroyQDesignerPropertyEditorInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerPropertyEditorInterface"})
}

func (ptr *QDesignerPropertyEditorInterface) DestroyQDesignerPropertyEditorInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerPropertyEditorInterfaceDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerPropertyEditorInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerPropertyEditorInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerPropertyEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerPropertyEditorInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerPropertyEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerPropertyEditorInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerPropertyEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerPropertyEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerPropertyEditorInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerPropertyEditorInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerPropertyEditorInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerPropertyEditorInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerPropertyEditorInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerPropertyEditorInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerPropertyEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerPropertyEditorInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerPropertyEditorInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerPropertyEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerPropertyEditorInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerPropertyEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerPropertySheetExtension struct {
	internal.Internal
}

type QDesignerPropertySheetExtension_ITF interface {
	QDesignerPropertySheetExtension_PTR() *QDesignerPropertySheetExtension
}

func (ptr *QDesignerPropertySheetExtension) QDesignerPropertySheetExtension_PTR() *QDesignerPropertySheetExtension {
	return ptr
}

func (ptr *QDesignerPropertySheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerPropertySheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerPropertySheetExtension(ptr QDesignerPropertySheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerPropertySheetExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerPropertySheetExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerPropertySheetExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerPropertySheetExtension) {
	n = new(QDesignerPropertySheetExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerPropertySheetExtension")
	return
}
func (ptr *QDesignerPropertySheetExtension) ConnectCount(f func() int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCount", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectCount() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCount"})
}

func (ptr *QDesignerPropertySheetExtension) Count() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Count"}).(float64))
}

func (ptr *QDesignerPropertySheetExtension) ConnectHasReset(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectHasReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasReset"})
}

func (ptr *QDesignerPropertySheetExtension) HasReset(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasReset", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectIndexOf(f func(name string) int) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIndexOf", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIndexOf() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIndexOf"})
}

func (ptr *QDesignerPropertySheetExtension) IndexOf(name string) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IndexOf", name}).(float64))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsAttribute(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsAttribute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsAttribute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsAttribute"})
}

func (ptr *QDesignerPropertySheetExtension) IsAttribute(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsAttribute", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsChanged(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsChanged"})
}

func (ptr *QDesignerPropertySheetExtension) IsChanged(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsChanged", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsEnabled(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsEnabled", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsEnabled() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsEnabled"})
}

func (ptr *QDesignerPropertySheetExtension) IsEnabled(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEnabled", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) IsEnabledDefault(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEnabledDefault", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsVisible(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIsVisible", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsVisible() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIsVisible"})
}

func (ptr *QDesignerPropertySheetExtension) IsVisible(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsVisible", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectProperty(f func(index int) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProperty"})
}

func (ptr *QDesignerPropertySheetExtension) Property(index int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Property", index}).(*core.QVariant)
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyGroup(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPropertyGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPropertyGroup"})
}

func (ptr *QDesignerPropertySheetExtension) PropertyGroup(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyGroup", index}).(string)
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyName(f func(index int) string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPropertyName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPropertyName"})
}

func (ptr *QDesignerPropertySheetExtension) PropertyName(index int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PropertyName", index}).(string)
}

func (ptr *QDesignerPropertySheetExtension) ConnectReset(f func(index int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReset", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectReset() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReset"})
}

func (ptr *QDesignerPropertySheetExtension) Reset(index int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reset", index}).(bool)
}

func (ptr *QDesignerPropertySheetExtension) ConnectSetAttribute(f func(index int, attribute bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetAttribute", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetAttribute() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetAttribute"})
}

func (ptr *QDesignerPropertySheetExtension) SetAttribute(index int, attribute bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", index, attribute})
}

func (ptr *QDesignerPropertySheetExtension) ConnectSetChanged(f func(index int, changed bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetChanged"})
}

func (ptr *QDesignerPropertySheetExtension) SetChanged(index int, changed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetChanged", index, changed})
}

func (ptr *QDesignerPropertySheetExtension) ConnectSetProperty(f func(index int, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetProperty", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetProperty() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetProperty"})
}

func (ptr *QDesignerPropertySheetExtension) SetProperty(index int, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProperty", index, value})
}

func (ptr *QDesignerPropertySheetExtension) ConnectSetPropertyGroup(f func(index int, group string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPropertyGroup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetPropertyGroup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPropertyGroup"})
}

func (ptr *QDesignerPropertySheetExtension) SetPropertyGroup(index int, group string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPropertyGroup", index, group})
}

func (ptr *QDesignerPropertySheetExtension) ConnectSetVisible(f func(index int, visible bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVisible", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetVisible() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVisible"})
}

func (ptr *QDesignerPropertySheetExtension) SetVisible(index int, visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisible", index, visible})
}

func (ptr *QDesignerPropertySheetExtension) ConnectDestroyQDesignerPropertySheetExtension(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerPropertySheetExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerPropertySheetExtension) DisconnectDestroyQDesignerPropertySheetExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerPropertySheetExtension"})
}

func (ptr *QDesignerPropertySheetExtension) DestroyQDesignerPropertySheetExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerPropertySheetExtension"})
}

func (ptr *QDesignerPropertySheetExtension) DestroyQDesignerPropertySheetExtensionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerPropertySheetExtensionDefault"})
}

type QDesignerResourceBrowserInterface struct {
	widgets.QWidget
}

type QDesignerResourceBrowserInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerResourceBrowserInterface_PTR() *QDesignerResourceBrowserInterface
}

func (ptr *QDesignerResourceBrowserInterface) QDesignerResourceBrowserInterface_PTR() *QDesignerResourceBrowserInterface {
	return ptr
}

func (ptr *QDesignerResourceBrowserInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerResourceBrowserInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerResourceBrowserInterface(ptr QDesignerResourceBrowserInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerResourceBrowserInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerResourceBrowserInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerResourceBrowserInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerResourceBrowserInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerResourceBrowserInterface) {
	n = new(QDesignerResourceBrowserInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerResourceBrowserInterface")
	return
}

func (ptr *QDesignerResourceBrowserInterface) DestroyQDesignerResourceBrowserInterface() {
}

func (ptr *QDesignerResourceBrowserInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerResourceBrowserInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerResourceBrowserInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerResourceBrowserInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerResourceBrowserInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerResourceBrowserInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerResourceBrowserInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerResourceBrowserInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerResourceBrowserInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerResourceBrowserInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerResourceBrowserInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerResourceBrowserInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerResourceBrowserInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerResourceBrowserInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerResourceBrowserInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerResourceBrowserInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerResourceBrowserInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerResourceBrowserInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerResourceBrowserInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerResourceBrowserInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerResourceBrowserInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerResourceBrowserInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerResourceBrowserInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerResourceBrowserInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerResourceBrowserInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerResourceBrowserInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerResourceBrowserInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerResourceBrowserInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerResourceBrowserInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerServer struct {
	core.QObject
}

type QDesignerServer_ITF interface {
	core.QObject_ITF
	QDesignerServer_PTR() *QDesignerServer
}

func (ptr *QDesignerServer) QDesignerServer_PTR() *QDesignerServer {
	return ptr
}

func (ptr *QDesignerServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerServer(ptr QDesignerServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerServer_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerServerFromPointer(ptr unsafe.Pointer) (n *QDesignerServer) {
	n = new(QDesignerServer)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerServer")
	return
}

func (ptr *QDesignerServer) DestroyQDesignerServer() {
}

type QDesignerSettings struct {
	internal.Internal
}

type QDesignerSettings_ITF interface {
	QDesignerSettings_PTR() *QDesignerSettings
}

func (ptr *QDesignerSettings) QDesignerSettings_PTR() *QDesignerSettings {
	return ptr
}

func (ptr *QDesignerSettings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerSettings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerSettings(ptr QDesignerSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerSettings_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerSettings) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerSettingsFromPointer(ptr unsafe.Pointer) (n *QDesignerSettings) {
	n = new(QDesignerSettings)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerSettings")
	return
}

func (ptr *QDesignerSettings) DestroyQDesignerSettings() {
}

type QDesignerTaskMenuExtension struct {
	internal.Internal
}

type QDesignerTaskMenuExtension_ITF interface {
	QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension
}

func (ptr *QDesignerTaskMenuExtension) QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension {
	return ptr
}

func (ptr *QDesignerTaskMenuExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDesignerTaskMenuExtension(ptr QDesignerTaskMenuExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerTaskMenuExtension_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerTaskMenuExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDesignerTaskMenuExtensionFromPointer(ptr unsafe.Pointer) (n *QDesignerTaskMenuExtension) {
	n = new(QDesignerTaskMenuExtension)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerTaskMenuExtension")
	return
}
func (ptr *QDesignerTaskMenuExtension) ConnectPreferredEditAction(f func() *widgets.QAction) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreferredEditAction", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerTaskMenuExtension) DisconnectPreferredEditAction() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreferredEditAction"})
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditAction() *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreferredEditAction"}).(*widgets.QAction)
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditActionDefault() *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreferredEditActionDefault"}).(*widgets.QAction)
}

func (ptr *QDesignerTaskMenuExtension) ConnectTaskActions(f func() []*widgets.QAction) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTaskActions", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerTaskMenuExtension) DisconnectTaskActions() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTaskActions"})
}

func (ptr *QDesignerTaskMenuExtension) TaskActions() []*widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TaskActions"}).([]*widgets.QAction)
}

func (ptr *QDesignerTaskMenuExtension) ConnectDestroyQDesignerTaskMenuExtension(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerTaskMenuExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerTaskMenuExtension) DisconnectDestroyQDesignerTaskMenuExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerTaskMenuExtension"})
}

func (ptr *QDesignerTaskMenuExtension) DestroyQDesignerTaskMenuExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerTaskMenuExtension"})
}

func (ptr *QDesignerTaskMenuExtension) DestroyQDesignerTaskMenuExtensionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerTaskMenuExtensionDefault"})
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__taskActions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__taskActions_setList", i})
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__taskActions_newList"}).(unsafe.Pointer)
}

type QDesignerToolWindow struct {
	MainWindowBase
}

type QDesignerToolWindow_ITF interface {
	MainWindowBase_ITF
	QDesignerToolWindow_PTR() *QDesignerToolWindow
}

func (ptr *QDesignerToolWindow) QDesignerToolWindow_PTR() *QDesignerToolWindow {
	return ptr
}

func (ptr *QDesignerToolWindow) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.MainWindowBase_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerToolWindow) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.MainWindowBase_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerToolWindow(ptr QDesignerToolWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerToolWindow_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerToolWindow) InitFromInternal(ptr uintptr, name string) {
	n.MainWindowBase_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerToolWindow) ClassNameInternalF() string {
	return n.MainWindowBase_PTR().ClassNameInternalF()
}

func NewQDesignerToolWindowFromPointer(ptr unsafe.Pointer) (n *QDesignerToolWindow) {
	n = new(QDesignerToolWindow)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerToolWindow")
	return
}

func (ptr *QDesignerToolWindow) DestroyQDesignerToolWindow() {
}

type QDesignerWidgetBoxInterface struct {
	widgets.QWidget
}

type QDesignerWidgetBoxInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerWidgetBoxInterface_PTR() *QDesignerWidgetBoxInterface
}

func (ptr *QDesignerWidgetBoxInterface) QDesignerWidgetBoxInterface_PTR() *QDesignerWidgetBoxInterface {
	return ptr
}

func (ptr *QDesignerWidgetBoxInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerWidgetBoxInterface(ptr QDesignerWidgetBoxInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerWidgetBoxInterface_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerWidgetBoxInterface) InitFromInternal(ptr uintptr, name string) {
	n.QWidget_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerWidgetBoxInterface) ClassNameInternalF() string {
	return n.QWidget_PTR().ClassNameInternalF()
}

func NewQDesignerWidgetBoxInterfaceFromPointer(ptr unsafe.Pointer) (n *QDesignerWidgetBoxInterface) {
	n = new(QDesignerWidgetBoxInterface)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerWidgetBoxInterface")
	return
}
func (ptr *QDesignerWidgetBoxInterface) ConnectFileName(f func() string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFileName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFileName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFileName"})
}

func (ptr *QDesignerWidgetBoxInterface) FileName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileName"}).(string)
}

func (ptr *QDesignerWidgetBoxInterface) ConnectLoad(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLoad", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectLoad() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLoad"})
}

func (ptr *QDesignerWidgetBoxInterface) Load() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Load"}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSave(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSave", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSave() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSave"})
}

func (ptr *QDesignerWidgetBoxInterface) Save() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Save"}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetFileName(f func(fileName string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetFileName", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetFileName() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetFileName"})
}

func (ptr *QDesignerWidgetBoxInterface) SetFileName(fileName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFileName", fileName})
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDestroyQDesignerWidgetBoxInterface(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDesignerWidgetBoxInterface", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDestroyQDesignerWidgetBoxInterface() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDesignerWidgetBoxInterface"})
}

func (ptr *QDesignerWidgetBoxInterface) DestroyQDesignerWidgetBoxInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerWidgetBoxInterface"})
}

func (ptr *QDesignerWidgetBoxInterface) DestroyQDesignerWidgetBoxInterfaceDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDesignerWidgetBoxInterfaceDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) __dropWidgets_item_list_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dropWidgets_item_list_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerWidgetBoxInterface) __actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_atList(i int) *widgets.QAction {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_atList", i}).(*widgets.QAction)
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__insertActions_actions_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDesignerWidgetBoxInterface) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDesignerWidgetBoxInterface) ActionEventDefault(event gui.QActionEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActionEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) ChangeEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChangeEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) CloseDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ContextMenuEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragEnterEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragLeaveEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DragMoveEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) DropEventDefault(event gui.QDropEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DropEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) EnterEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnterEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) EventDefault(event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", event}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusInEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) FocusNextPrevChildDefault(next bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusNextPrevChildDefault", next}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FocusOutEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) HasHeightForWidthDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasHeightForWidthDefault"}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) HeightForWidthDefault(w int) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HeightForWidthDefault", w}).(float64))
}

func (ptr *QDesignerWidgetBoxInterface) HideDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) HideEventDefault(event gui.QHideEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HideEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) InitPainterDefault(painter gui.QPainter_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InitPainterDefault", painter})
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InputMethodQueryDefault", query}).(*core.QVariant)
}

func (ptr *QDesignerWidgetBoxInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyPressEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyReleaseEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) LeaveEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) LowerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LowerDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetricDefault", m}).(float64))
}

func (ptr *QDesignerWidgetBoxInterface) MinimumSizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MinimumSizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerWidgetBoxInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseDoubleClickEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseMoveEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MousePressEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MouseReleaseEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MoveEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NativeEventDefault", eventType, message, result}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) PaintEngineDefault() *gui.QPaintEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEngineDefault"}).(*gui.QPaintEngine)
}

func (ptr *QDesignerWidgetBoxInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PaintEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) RaiseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RaiseDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) RepaintDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RepaintDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResizeEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) SetDisabledDefault(disable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDisabledDefault", disable})
}

func (ptr *QDesignerWidgetBoxInterface) SetEnabledDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEnabledDefault", vbo})
}

func (ptr *QDesignerWidgetBoxInterface) SetFocus2Default() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFocus2Default"})
}

func (ptr *QDesignerWidgetBoxInterface) SetHiddenDefault(hidden bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHiddenDefault", hidden})
}

func (ptr *QDesignerWidgetBoxInterface) SetStyleSheetDefault(styleSheet string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStyleSheetDefault", styleSheet})
}

func (ptr *QDesignerWidgetBoxInterface) SetVisibleDefault(visible bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVisibleDefault", visible})
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowModifiedDefault(vbo bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowModifiedDefault", vbo})
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowTitleDefault(vqs string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetWindowTitleDefault", vqs})
}

func (ptr *QDesignerWidgetBoxInterface) ShowDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) ShowEventDefault(event gui.QShowEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) ShowFullScreenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowFullScreenDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) ShowMaximizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMaximizedDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) ShowMinimizedDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowMinimizedDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) ShowNormalDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShowNormalDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) SizeHintDefault() *core.QSize {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeHintDefault"}).(*core.QSize)
}

func (ptr *QDesignerWidgetBoxInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TabletEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) UpdateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) UpdateMicroFocusDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMicroFocusDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WheelEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDesignerWidgetBoxInterface) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDesignerWidgetBoxInterface) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDesignerWidgetBoxInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDesignerWidgetBoxInterface) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDesignerWidgetBoxInterface) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDesignerWorkbench struct {
	core.QObject
}

type QDesignerWorkbench_ITF interface {
	core.QObject_ITF
	QDesignerWorkbench_PTR() *QDesignerWorkbench
}

func (ptr *QDesignerWorkbench) QDesignerWorkbench_PTR() *QDesignerWorkbench {
	return ptr
}

func (ptr *QDesignerWorkbench) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDesignerWorkbench) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDesignerWorkbench(ptr QDesignerWorkbench_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerWorkbench_PTR().Pointer()
	}
	return nil
}

func (n *QDesignerWorkbench) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDesignerWorkbench) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDesignerWorkbenchFromPointer(ptr unsafe.Pointer) (n *QDesignerWorkbench) {
	n = new(QDesignerWorkbench)
	n.InitFromInternal(uintptr(ptr), "designer.QDesignerWorkbench")
	return
}

func (ptr *QDesignerWorkbench) DestroyQDesignerWorkbench() {
}

type QExtensionFactory struct {
	core.QObject
	QAbstractExtensionFactory
}

type QExtensionFactory_ITF interface {
	core.QObject_ITF
	QAbstractExtensionFactory_ITF
	QExtensionFactory_PTR() *QExtensionFactory
}

func (ptr *QExtensionFactory) QExtensionFactory_PTR() *QExtensionFactory {
	return ptr
}

func (ptr *QExtensionFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QExtensionFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QAbstractExtensionFactory_PTR().SetPointer(p)
	}
}

func PointerFromQExtensionFactory(ptr QExtensionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExtensionFactory_PTR().Pointer()
	}
	return nil
}

func (n *QExtensionFactory) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QAbstractExtensionFactory_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QExtensionFactory) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQExtensionFactoryFromPointer(ptr unsafe.Pointer) (n *QExtensionFactory) {
	n = new(QExtensionFactory)
	n.InitFromInternal(uintptr(ptr), "designer.QExtensionFactory")
	return
}

func (ptr *QExtensionFactory) DestroyQExtensionFactory() {
}

func NewQExtensionFactory(parent QExtensionManager_ITF) *QExtensionFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQExtensionFactory", "", parent}).(*QExtensionFactory)
}

func (ptr *QExtensionFactory) ConnectCreateExtension(f func(object *core.QObject, iid string, parent *core.QObject) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectCreateExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionFactory) DisconnectCreateExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectCreateExtension"})
}

func (ptr *QExtensionFactory) CreateExtension(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CreateExtension", object, iid, parent}).(*core.QObject)
}

func (ptr *QExtensionFactory) CreateExtensionDefault(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CreateExtensionDefault", object, iid, parent}).(*core.QObject)
}

func (ptr *QExtensionFactory) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionFactory) DisconnectExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectExtension"})
}

func (ptr *QExtensionFactory) Extension(object core.QObject_ITF, iid string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Extension", object, iid}).(*core.QObject)
}

func (ptr *QExtensionFactory) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ExtensionDefault", object, iid}).(*core.QObject)
}

func (ptr *QExtensionFactory) ExtensionManager() *QExtensionManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ExtensionManager"}).(*QExtensionManager)
}

func (ptr *QExtensionFactory) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QExtensionFactory) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QExtensionFactory) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionFactory) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QExtensionFactory) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QExtensionFactory) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionFactory) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QExtensionFactory) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QExtensionFactory) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QExtensionFactory) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QExtensionFactory) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QExtensionFactory) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QExtensionFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QExtensionFactory) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QExtensionFactory) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QExtensionFactory) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QExtensionFactory) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QExtensionFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QExtensionFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QExtensionFactory) Event(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", e}).(bool)
}

func (ptr *QExtensionFactory) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QExtensionFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QExtensionFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QExtensionFactory) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QExtensionFactory) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QExtensionFactory) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QExtensionFactory) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QExtensionManager struct {
	core.QObject
	QAbstractExtensionManager
}

type QExtensionManager_ITF interface {
	core.QObject_ITF
	QAbstractExtensionManager_ITF
	QExtensionManager_PTR() *QExtensionManager
}

func (ptr *QExtensionManager) QExtensionManager_PTR() *QExtensionManager {
	return ptr
}

func (ptr *QExtensionManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QExtensionManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QAbstractExtensionManager_PTR().SetPointer(p)
	}
}

func PointerFromQExtensionManager(ptr QExtensionManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExtensionManager_PTR().Pointer()
	}
	return nil
}

func (n *QExtensionManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QAbstractExtensionManager_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QExtensionManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQExtensionManagerFromPointer(ptr unsafe.Pointer) (n *QExtensionManager) {
	n = new(QExtensionManager)
	n.InitFromInternal(uintptr(ptr), "designer.QExtensionManager")
	return
}
func NewQExtensionManager(parent core.QObject_ITF) *QExtensionManager {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQExtensionManager", "", parent}).(*QExtensionManager)
}

func (ptr *QExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectExtension", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionManager) DisconnectExtension() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectExtension"})
}

func (ptr *QExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Extension", object, iid}).(*core.QObject)
}

func (ptr *QExtensionManager) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ExtensionDefault", object, iid}).(*core.QObject)
}

func (ptr *QExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectRegisterExtensions", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionManager) DisconnectRegisterExtensions() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectRegisterExtensions"})
}

func (ptr *QExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "RegisterExtensions", factory, iid})
}

func (ptr *QExtensionManager) RegisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "RegisterExtensionsDefault", factory, iid})
}

func (ptr *QExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectUnregisterExtensions", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionManager) DisconnectUnregisterExtensions() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectUnregisterExtensions"})
}

func (ptr *QExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UnregisterExtensions", factory, iid})
}

func (ptr *QExtensionManager) UnregisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "UnregisterExtensionsDefault", factory, iid})
}

func (ptr *QExtensionManager) ConnectDestroyQExtensionManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectDestroyQExtensionManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QExtensionManager) DisconnectDestroyQExtensionManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectDestroyQExtensionManager"})
}

func (ptr *QExtensionManager) DestroyQExtensionManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQExtensionManager"})
}

func (ptr *QExtensionManager) DestroyQExtensionManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DestroyQExtensionManagerDefault"})
}

func (ptr *QExtensionManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QExtensionManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QExtensionManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QExtensionManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QExtensionManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QExtensionManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QExtensionManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QExtensionManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QExtensionManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QExtensionManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QExtensionManager) ChildEvent(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEvent", event})
}

func (ptr *QExtensionManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QExtensionManager) ConnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotify", sign})
}

func (ptr *QExtensionManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QExtensionManager) CustomEvent(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEvent", event})
}

func (ptr *QExtensionManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QExtensionManager) DeleteLater() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLater"})
}

func (ptr *QExtensionManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QExtensionManager) DisconnectNotify(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotify", sign})
}

func (ptr *QExtensionManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QExtensionManager) Event(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "Event", e}).(bool)
}

func (ptr *QExtensionManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QExtensionManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilter", watched, event}).(bool)
}

func (ptr *QExtensionManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QExtensionManager) MetaObject() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObject"}).(*core.QMetaObject)
}

func (ptr *QExtensionManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QExtensionManager) TimerEvent(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEvent", event})
}

func (ptr *QExtensionManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.QObject_PTR().ClassNameInternalF(), "TimerEventDefault", event})
}

type QFormBuilder struct {
	QAbstractFormBuilder
}

type QFormBuilder_ITF interface {
	QAbstractFormBuilder_ITF
	QFormBuilder_PTR() *QFormBuilder
}

func (ptr *QFormBuilder) QFormBuilder_PTR() *QFormBuilder {
	return ptr
}

func (ptr *QFormBuilder) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractFormBuilder_PTR().Pointer()
	}
	return nil
}

func (ptr *QFormBuilder) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractFormBuilder_PTR().SetPointer(p)
	}
}

func PointerFromQFormBuilder(ptr QFormBuilder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFormBuilder_PTR().Pointer()
	}
	return nil
}

func (n *QFormBuilder) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractFormBuilder_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QFormBuilder) ClassNameInternalF() string {
	return n.QAbstractFormBuilder_PTR().ClassNameInternalF()
}

func NewQFormBuilderFromPointer(ptr unsafe.Pointer) (n *QFormBuilder) {
	n = new(QFormBuilder)
	n.InitFromInternal(uintptr(ptr), "designer.QFormBuilder")
	return
}
func NewQFormBuilder() *QFormBuilder {

	return internal.CallLocalFunction([]interface{}{"", "", "designer.NewQFormBuilder", ""}).(*QFormBuilder)
}

func (ptr *QFormBuilder) AddPluginPath(pluginPath string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddPluginPath", pluginPath})
}

func (ptr *QFormBuilder) ClearPluginPaths() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearPluginPaths"})
}

func (ptr *QFormBuilder) CustomWidgets() []*QDesignerCustomWidgetInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomWidgets"}).([]*QDesignerCustomWidgetInterface)
}

func (ptr *QFormBuilder) PluginPaths() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PluginPaths"}).([]string)
}

func (ptr *QFormBuilder) SetPluginPath(pluginPaths []string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPluginPath", pluginPaths})
}

func (ptr *QFormBuilder) ConnectDestroyQFormBuilder(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQFormBuilder", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QFormBuilder) DisconnectDestroyQFormBuilder() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQFormBuilder"})
}

func (ptr *QFormBuilder) DestroyQFormBuilder() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQFormBuilder"})
}

func (ptr *QFormBuilder) DestroyQFormBuilderDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQFormBuilderDefault"})
}

func (ptr *QFormBuilder) __customWidgets_atList(i int) *QDesignerCustomWidgetInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_atList", i}).(*QDesignerCustomWidgetInterface)
}

func (ptr *QFormBuilder) __customWidgets_setList(i QDesignerCustomWidgetInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_setList", i})
}

func (ptr *QFormBuilder) __customWidgets_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__customWidgets_newList"}).(unsafe.Pointer)
}

type SaveFormAsTemplate struct {
	widgets.QDialog
}

type SaveFormAsTemplate_ITF interface {
	widgets.QDialog_ITF
	SaveFormAsTemplate_PTR() *SaveFormAsTemplate
}

func (ptr *SaveFormAsTemplate) SaveFormAsTemplate_PTR() *SaveFormAsTemplate {
	return ptr
}

func (ptr *SaveFormAsTemplate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *SaveFormAsTemplate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromSaveFormAsTemplate(ptr SaveFormAsTemplate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SaveFormAsTemplate_PTR().Pointer()
	}
	return nil
}

func (n *SaveFormAsTemplate) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *SaveFormAsTemplate) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewSaveFormAsTemplateFromPointer(ptr unsafe.Pointer) (n *SaveFormAsTemplate) {
	n = new(SaveFormAsTemplate)
	n.InitFromInternal(uintptr(ptr), "designer.SaveFormAsTemplate")
	return
}

func (ptr *SaveFormAsTemplate) DestroySaveFormAsTemplate() {
}

type ToolBarManager struct {
	core.QObject
}

type ToolBarManager_ITF interface {
	core.QObject_ITF
	ToolBarManager_PTR() *ToolBarManager
}

func (ptr *ToolBarManager) ToolBarManager_PTR() *ToolBarManager {
	return ptr
}

func (ptr *ToolBarManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ToolBarManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromToolBarManager(ptr ToolBarManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ToolBarManager_PTR().Pointer()
	}
	return nil
}

func (n *ToolBarManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *ToolBarManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewToolBarManagerFromPointer(ptr unsafe.Pointer) (n *ToolBarManager) {
	n = new(ToolBarManager)
	n.InitFromInternal(uintptr(ptr), "designer.ToolBarManager")
	return
}

func (ptr *ToolBarManager) DestroyToolBarManager() {
}

type ToolWindowFontSettings struct {
	internal.Internal
}

type ToolWindowFontSettings_ITF interface {
	ToolWindowFontSettings_PTR() *ToolWindowFontSettings
}

func (ptr *ToolWindowFontSettings) ToolWindowFontSettings_PTR() *ToolWindowFontSettings {
	return ptr
}

func (ptr *ToolWindowFontSettings) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *ToolWindowFontSettings) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromToolWindowFontSettings(ptr ToolWindowFontSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ToolWindowFontSettings_PTR().Pointer()
	}
	return nil
}

func (n *ToolWindowFontSettings) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewToolWindowFontSettingsFromPointer(ptr unsafe.Pointer) (n *ToolWindowFontSettings) {
	n = new(ToolWindowFontSettings)
	n.InitFromInternal(uintptr(ptr), "designer.ToolWindowFontSettings")
	return
}

func (ptr *ToolWindowFontSettings) DestroyToolWindowFontSettings() {
}

type VersionDialog struct {
	widgets.QDialog
}

type VersionDialog_ITF interface {
	widgets.QDialog_ITF
	VersionDialog_PTR() *VersionDialog
}

func (ptr *VersionDialog) VersionDialog_PTR() *VersionDialog {
	return ptr
}

func (ptr *VersionDialog) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func (ptr *VersionDialog) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QDialog_PTR().SetPointer(p)
	}
}

func PointerFromVersionDialog(ptr VersionDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.VersionDialog_PTR().Pointer()
	}
	return nil
}

func (n *VersionDialog) InitFromInternal(ptr uintptr, name string) {
	n.QDialog_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *VersionDialog) ClassNameInternalF() string {
	return n.QDialog_PTR().ClassNameInternalF()
}

func NewVersionDialogFromPointer(ptr unsafe.Pointer) (n *VersionDialog) {
	n = new(VersionDialog)
	n.InitFromInternal(uintptr(ptr), "designer.VersionDialog")
	return
}

func (ptr *VersionDialog) DestroyVersionDialog() {
}

func init() {
	internal.ConstructorTable["designer.QAbstractExtensionManager"] = NewQAbstractExtensionManagerFromPointer
	internal.ConstructorTable["designer.QAbstractFormBuilder"] = NewQAbstractFormBuilderFromPointer
	internal.ConstructorTable["designer.QDesignerActionEditorInterface"] = NewQDesignerActionEditorInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerContainerExtension"] = NewQDesignerContainerExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerCustomWidgetCollectionInterface"] = NewQDesignerCustomWidgetCollectionInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerCustomWidgetInterface"] = NewQDesignerCustomWidgetInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerDynamicPropertySheetExtension"] = NewQDesignerDynamicPropertySheetExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerFormEditorInterface"] = NewQDesignerFormEditorInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerFormWindowCursorInterface"] = NewQDesignerFormWindowCursorInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerFormWindowInterface"] = NewQDesignerFormWindowInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerFormWindowManagerInterface"] = NewQDesignerFormWindowManagerInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerLanguageExtension"] = NewQDesignerLanguageExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerMemberSheetExtension"] = NewQDesignerMemberSheetExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerObjectInspectorInterface"] = NewQDesignerObjectInspectorInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerPropertyEditorInterface"] = NewQDesignerPropertyEditorInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerPropertySheetExtension"] = NewQDesignerPropertySheetExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerResourceBrowserInterface"] = NewQDesignerResourceBrowserInterfaceFromPointer
	internal.ConstructorTable["designer.QDesignerTaskMenuExtension"] = NewQDesignerTaskMenuExtensionFromPointer
	internal.ConstructorTable["designer.QDesignerWidgetBoxInterface"] = NewQDesignerWidgetBoxInterfaceFromPointer
	internal.ConstructorTable["designer.QExtensionFactory"] = NewQExtensionFactoryFromPointer
	internal.ConstructorTable["designer.QExtensionManager"] = NewQExtensionManagerFromPointer
	internal.ConstructorTable["designer.QFormBuilder"] = NewQFormBuilderFromPointer
}
