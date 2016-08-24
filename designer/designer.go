// +build !minimal

package designer

//#include <stdint.h>
//#include <stdlib.h>
//#include "designer.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

type QAbstractExtensionFactory struct {
	ptr unsafe.Pointer
}

type QAbstractExtensionFactory_ITF interface {
	QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory
}

func (p *QAbstractExtensionFactory) QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory {
	return p
}

func (p *QAbstractExtensionFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAbstractExtensionFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQAbstractExtensionFactory(ptr QAbstractExtensionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractExtensionFactory_PTR().Pointer()
	}
	return nil
}

func NewQAbstractExtensionFactoryFromPointer(ptr unsafe.Pointer) *QAbstractExtensionFactory {
	var n = new(QAbstractExtensionFactory)
	n.SetPointer(ptr)
	return n
}

//export callbackQAbstractExtensionFactory_Extension
func callbackQAbstractExtensionFactory_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractExtensionFactory::extension")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractExtensionFactory::extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), C.GoString(iid)))
	}

	return core.PointerFromQObject(nil)
}

func (ptr *QAbstractExtensionFactory) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	defer qt.Recovering("connect QAbstractExtensionFactory::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionFactory::extension", f)
	}
}

func (ptr *QAbstractExtensionFactory) DisconnectExtension(object core.QObject_ITF, iid string) {
	defer qt.Recovering("disconnect QAbstractExtensionFactory::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionFactory::extension")
	}
}

func (ptr *QAbstractExtensionFactory) Extension(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QAbstractExtensionFactory::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QAbstractExtensionFactory_Extension(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractExtensionFactory) DestroyQAbstractExtensionFactory() {
	defer qt.Recovering("QAbstractExtensionFactory::~QAbstractExtensionFactory")

	if ptr.Pointer() != nil {
		C.QAbstractExtensionFactory_DestroyQAbstractExtensionFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QAbstractExtensionManager struct {
	ptr unsafe.Pointer
}

type QAbstractExtensionManager_ITF interface {
	QAbstractExtensionManager_PTR() *QAbstractExtensionManager
}

func (p *QAbstractExtensionManager) QAbstractExtensionManager_PTR() *QAbstractExtensionManager {
	return p
}

func (p *QAbstractExtensionManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAbstractExtensionManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQAbstractExtensionManager(ptr QAbstractExtensionManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractExtensionManager_PTR().Pointer()
	}
	return nil
}

func NewQAbstractExtensionManagerFromPointer(ptr unsafe.Pointer) *QAbstractExtensionManager {
	var n = new(QAbstractExtensionManager)
	n.SetPointer(ptr)
	return n
}

//export callbackQAbstractExtensionManager_Extension
func callbackQAbstractExtensionManager_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractExtensionManager::extension")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractExtensionManager::extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), C.GoString(iid)))
	}

	return core.PointerFromQObject(nil)
}

func (ptr *QAbstractExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	defer qt.Recovering("connect QAbstractExtensionManager::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::extension", f)
	}
}

func (ptr *QAbstractExtensionManager) DisconnectExtension(object core.QObject_ITF, iid string) {
	defer qt.Recovering("disconnect QAbstractExtensionManager::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::extension")
	}
}

func (ptr *QAbstractExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QAbstractExtensionManager::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QAbstractExtensionManager_Extension(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractExtensionManager_RegisterExtensions
func callbackQAbstractExtensionManager_RegisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid *C.char) {
	defer qt.Recovering("callback QAbstractExtensionManager::registerExtensions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractExtensionManager::registerExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	}

}

func (ptr *QAbstractExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	defer qt.Recovering("connect QAbstractExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::registerExtensions", f)
	}
}

func (ptr *QAbstractExtensionManager) DisconnectRegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("disconnect QAbstractExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::registerExtensions")
	}
}

func (ptr *QAbstractExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QAbstractExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QAbstractExtensionManager_RegisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

//export callbackQAbstractExtensionManager_UnregisterExtensions
func callbackQAbstractExtensionManager_UnregisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid *C.char) {
	defer qt.Recovering("callback QAbstractExtensionManager::unregisterExtensions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractExtensionManager::unregisterExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	}

}

func (ptr *QAbstractExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	defer qt.Recovering("connect QAbstractExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::unregisterExtensions", f)
	}
}

func (ptr *QAbstractExtensionManager) DisconnectUnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("disconnect QAbstractExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractExtensionManager::unregisterExtensions")
	}
}

func (ptr *QAbstractExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QAbstractExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QAbstractExtensionManager_UnregisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

func (ptr *QAbstractExtensionManager) DestroyQAbstractExtensionManager() {
	defer qt.Recovering("QAbstractExtensionManager::~QAbstractExtensionManager")

	if ptr.Pointer() != nil {
		C.QAbstractExtensionManager_DestroyQAbstractExtensionManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QAbstractFormBuilder struct {
	ptr unsafe.Pointer
}

type QAbstractFormBuilder_ITF interface {
	QAbstractFormBuilder_PTR() *QAbstractFormBuilder
}

func (p *QAbstractFormBuilder) QAbstractFormBuilder_PTR() *QAbstractFormBuilder {
	return p
}

func (p *QAbstractFormBuilder) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAbstractFormBuilder) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQAbstractFormBuilder(ptr QAbstractFormBuilder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractFormBuilder_PTR().Pointer()
	}
	return nil
}

func NewQAbstractFormBuilderFromPointer(ptr unsafe.Pointer) *QAbstractFormBuilder {
	var n = new(QAbstractFormBuilder)
	n.SetPointer(ptr)
	return n
}

//export callbackQAbstractFormBuilder_Load
func callbackQAbstractFormBuilder_Load(ptr unsafe.Pointer, device unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractFormBuilder::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractFormBuilder::load"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*core.QIODevice, *widgets.QWidget) *widgets.QWidget)(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
	}

	return widgets.PointerFromQWidget(NewQAbstractFormBuilderFromPointer(ptr).LoadDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
}

func (ptr *QAbstractFormBuilder) ConnectLoad(f func(device *core.QIODevice, parent *widgets.QWidget) *widgets.QWidget) {
	defer qt.Recovering("connect QAbstractFormBuilder::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractFormBuilder::load", f)
	}
}

func (ptr *QAbstractFormBuilder) DisconnectLoad() {
	defer qt.Recovering("disconnect QAbstractFormBuilder::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractFormBuilder::load")
	}
}

func (ptr *QAbstractFormBuilder) Load(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QAbstractFormBuilder::load")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QAbstractFormBuilder_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractFormBuilder) LoadDefault(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QAbstractFormBuilder::load")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QAbstractFormBuilder_LoadDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractFormBuilder_Save
func callbackQAbstractFormBuilder_Save(ptr unsafe.Pointer, device unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractFormBuilder::save")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractFormBuilder::save"); signal != nil {
		signal.(func(*core.QIODevice, *widgets.QWidget))(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQAbstractFormBuilderFromPointer(ptr).SaveDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QAbstractFormBuilder) ConnectSave(f func(device *core.QIODevice, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QAbstractFormBuilder::save")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractFormBuilder::save", f)
	}
}

func (ptr *QAbstractFormBuilder) DisconnectSave() {
	defer qt.Recovering("disconnect QAbstractFormBuilder::save")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractFormBuilder::save")
	}
}

func (ptr *QAbstractFormBuilder) Save(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QAbstractFormBuilder::save")

	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_Save(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractFormBuilder) SaveDefault(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QAbstractFormBuilder::save")

	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_SaveDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}

func NewQAbstractFormBuilder() *QAbstractFormBuilder {
	defer qt.Recovering("QAbstractFormBuilder::QAbstractFormBuilder")

	var tmpValue = NewQAbstractFormBuilderFromPointer(C.QAbstractFormBuilder_NewQAbstractFormBuilder())
	runtime.SetFinalizer(tmpValue, (*QAbstractFormBuilder).DestroyQAbstractFormBuilder)
	return tmpValue
}

func (ptr *QAbstractFormBuilder) ErrorString() string {
	defer qt.Recovering("QAbstractFormBuilder::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractFormBuilder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractFormBuilder) SetWorkingDirectory(directory core.QDir_ITF) {
	defer qt.Recovering("QAbstractFormBuilder::setWorkingDirectory")

	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_SetWorkingDirectory(ptr.Pointer(), core.PointerFromQDir(directory))
	}
}

func (ptr *QAbstractFormBuilder) WorkingDirectory() *core.QDir {
	defer qt.Recovering("QAbstractFormBuilder::workingDirectory")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDirFromPointer(C.QAbstractFormBuilder_WorkingDirectory(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractFormBuilder) DestroyQAbstractFormBuilder() {
	defer qt.Recovering("QAbstractFormBuilder::~QAbstractFormBuilder")

	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_DestroyQAbstractFormBuilder(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerActionEditorInterface struct {
	widgets.QWidget
}

type QDesignerActionEditorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerActionEditorInterface_PTR() *QDesignerActionEditorInterface
}

func (p *QDesignerActionEditorInterface) QDesignerActionEditorInterface_PTR() *QDesignerActionEditorInterface {
	return p
}

func (p *QDesignerActionEditorInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerActionEditorInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerActionEditorInterface(ptr QDesignerActionEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerActionEditorInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerActionEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerActionEditorInterface {
	var n = new(QDesignerActionEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerActionEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerActionEditorInterface {
	defer qt.Recovering("QDesignerActionEditorInterface::QDesignerActionEditorInterface")

	var tmpValue = NewQDesignerActionEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_NewQDesignerActionEditorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQDesignerActionEditorInterface_Core
func callbackQDesignerActionEditorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerActionEditorInterface::core")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerActionEditorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerActionEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::core")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::core", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectCore() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::core")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::core")
	}
}

func (ptr *QDesignerActionEditorInterface) Core() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerActionEditorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) CoreDefault() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerActionEditorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_ManageAction
func callbackQDesignerActionEditorInterface_ManageAction(ptr unsafe.Pointer, action unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::manageAction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::manageAction"); signal != nil {
		signal.(func(*widgets.QAction))(widgets.NewQActionFromPointer(action))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectManageAction(f func(action *widgets.QAction)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::manageAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::manageAction", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectManageAction(action widgets.QAction_ITF) {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::manageAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::manageAction")
	}
}

func (ptr *QDesignerActionEditorInterface) ManageAction(action widgets.QAction_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::manageAction")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ManageAction(ptr.Pointer(), widgets.PointerFromQAction(action))
	}
}

//export callbackQDesignerActionEditorInterface_SetFormWindow
func callbackQDesignerActionEditorInterface_SetFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setFormWindow", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setFormWindow")
	}
}

func (ptr *QDesignerActionEditorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::setFormWindow")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerActionEditorInterface_UnmanageAction
func callbackQDesignerActionEditorInterface_UnmanageAction(ptr unsafe.Pointer, action unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::unmanageAction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::unmanageAction"); signal != nil {
		signal.(func(*widgets.QAction))(widgets.NewQActionFromPointer(action))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectUnmanageAction(f func(action *widgets.QAction)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::unmanageAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::unmanageAction", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectUnmanageAction(action widgets.QAction_ITF) {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::unmanageAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::unmanageAction")
	}
}

func (ptr *QDesignerActionEditorInterface) UnmanageAction(action widgets.QAction_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::unmanageAction")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UnmanageAction(ptr.Pointer(), widgets.PointerFromQAction(action))
	}
}

func (ptr *QDesignerActionEditorInterface) DestroyQDesignerActionEditorInterface() {
	defer qt.Recovering("QDesignerActionEditorInterface::~QDesignerActionEditorInterface")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerActionEditorInterface_ActionEvent
func callbackQDesignerActionEditorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::actionEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::actionEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DragEnterEvent
func callbackQDesignerActionEditorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragEnterEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragEnterEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DragLeaveEvent
func callbackQDesignerActionEditorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragLeaveEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragLeaveEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DragMoveEvent
func callbackQDesignerActionEditorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragMoveEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dragMoveEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DropEvent
func callbackQDesignerActionEditorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dropEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::dropEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_EnterEvent
func callbackQDesignerActionEditorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::enterEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::enterEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_FocusInEvent
func callbackQDesignerActionEditorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusInEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusInEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_FocusOutEvent
func callbackQDesignerActionEditorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusOutEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusOutEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_HideEvent
func callbackQDesignerActionEditorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hideEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hideEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_LeaveEvent
func callbackQDesignerActionEditorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::leaveEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::leaveEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MinimumSizeHint
func callbackQDesignerActionEditorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerActionEditorInterface::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerActionEditorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerActionEditorInterface) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::minimumSizeHint", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::minimumSizeHint")
	}
}

func (ptr *QDesignerActionEditorInterface) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDesignerActionEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerActionEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_MoveEvent
func callbackQDesignerActionEditorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::moveEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::moveEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_PaintEvent
func callbackQDesignerActionEditorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::paintEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::paintEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_SetEnabled
func callbackQDesignerActionEditorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setEnabled", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setEnabled")
	}
}

func (ptr *QDesignerActionEditorInterface) SetEnabled(vbo bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerActionEditorInterface) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerActionEditorInterface_SetStyleSheet
func callbackQDesignerActionEditorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setStyleSheet", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setStyleSheet")
	}
}

func (ptr *QDesignerActionEditorInterface) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QDesignerActionEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerActionEditorInterface_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QDesignerActionEditorInterface) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QDesignerActionEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerActionEditorInterface_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQDesignerActionEditorInterface_SetVisible
func callbackQDesignerActionEditorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setVisible", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setVisible")
	}
}

func (ptr *QDesignerActionEditorInterface) SetVisible(visible bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerActionEditorInterface) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerActionEditorInterface_SetWindowModified
func callbackQDesignerActionEditorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setWindowModified", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setWindowModified")
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowModified(vbo bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerActionEditorInterface_SetWindowTitle
func callbackQDesignerActionEditorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setWindowTitle", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setWindowTitle")
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowTitle(vqs string) {
	defer qt.Recovering("QDesignerActionEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerActionEditorInterface_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QDesignerActionEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerActionEditorInterface_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQDesignerActionEditorInterface_ShowEvent
func callbackQDesignerActionEditorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_SizeHint
func callbackQDesignerActionEditorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerActionEditorInterface::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerActionEditorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerActionEditorInterface) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::sizeHint", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::sizeHint")
	}
}

func (ptr *QDesignerActionEditorInterface) SizeHint() *core.QSize {
	defer qt.Recovering("QDesignerActionEditorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerActionEditorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_ChangeEvent
func callbackQDesignerActionEditorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::changeEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::changeEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Close
func callbackQDesignerActionEditorInterface_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerActionEditorInterface::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerActionEditorInterface) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::close", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectClose() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::close")
	}
}

func (ptr *QDesignerActionEditorInterface) Close() bool {
	defer qt.Recovering("QDesignerActionEditorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerActionEditorInterface) CloseDefault() bool {
	defer qt.Recovering("QDesignerActionEditorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_CloseEvent
func callbackQDesignerActionEditorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::closeEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::closeEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ContextMenuEvent
func callbackQDesignerActionEditorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::contextMenuEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::contextMenuEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_FocusNextPrevChild
func callbackQDesignerActionEditorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QDesignerActionEditorInterface::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerActionEditorInterface) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusNextPrevChild", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::focusNextPrevChild")
	}
}

func (ptr *QDesignerActionEditorInterface) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QDesignerActionEditorInterface) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_HasHeightForWidth
func callbackQDesignerActionEditorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerActionEditorInterface::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerActionEditorInterface) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hasHeightForWidth", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hasHeightForWidth")
	}
}

func (ptr *QDesignerActionEditorInterface) HasHeightForWidth() bool {
	defer qt.Recovering("QDesignerActionEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerActionEditorInterface) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QDesignerActionEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_HeightForWidth
func callbackQDesignerActionEditorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QDesignerActionEditorInterface::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerActionEditorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerActionEditorInterface) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::heightForWidth", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::heightForWidth")
	}
}

func (ptr *QDesignerActionEditorInterface) HeightForWidth(w int) int {
	defer qt.Recovering("QDesignerActionEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerActionEditorInterface_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QDesignerActionEditorInterface) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QDesignerActionEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerActionEditorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerActionEditorInterface_Hide
func callbackQDesignerActionEditorInterface_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectHide(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hide", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectHide() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::hide")
	}
}

func (ptr *QDesignerActionEditorInterface) Hide() {
	defer qt.Recovering("QDesignerActionEditorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Hide(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) HideDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_InputMethodEvent
func callbackQDesignerActionEditorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::inputMethodEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::inputMethodEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_InputMethodQuery
func callbackQDesignerActionEditorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerActionEditorInterface::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerActionEditorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerActionEditorInterface) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::inputMethodQuery", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::inputMethodQuery")
	}
}

func (ptr *QDesignerActionEditorInterface) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerActionEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerActionEditorInterface_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerActionEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerActionEditorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_KeyPressEvent
func callbackQDesignerActionEditorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::keyPressEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::keyPressEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_KeyReleaseEvent
func callbackQDesignerActionEditorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::keyReleaseEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::keyReleaseEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Lower
func callbackQDesignerActionEditorInterface_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectLower(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::lower", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectLower() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::lower")
	}
}

func (ptr *QDesignerActionEditorInterface) Lower() {
	defer qt.Recovering("QDesignerActionEditorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Lower(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) LowerDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_MouseDoubleClickEvent
func callbackQDesignerActionEditorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseDoubleClickEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MouseMoveEvent
func callbackQDesignerActionEditorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseMoveEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseMoveEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MousePressEvent
func callbackQDesignerActionEditorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mousePressEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mousePressEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MouseReleaseEvent
func callbackQDesignerActionEditorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseReleaseEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::mouseReleaseEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_NativeEvent
func callbackQDesignerActionEditorInterface_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QDesignerActionEditorInterface::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QDesignerActionEditorInterface) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::nativeEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::nativeEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerActionEditorInterface_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QDesignerActionEditorInterface) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerActionEditorInterface_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_Raise
func callbackQDesignerActionEditorInterface_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectRaise(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::raise", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectRaise() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::raise")
	}
}

func (ptr *QDesignerActionEditorInterface) Raise() {
	defer qt.Recovering("QDesignerActionEditorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Raise(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) RaiseDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_Repaint
func callbackQDesignerActionEditorInterface_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::repaint", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectRepaint() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::repaint")
	}
}

func (ptr *QDesignerActionEditorInterface) Repaint() {
	defer qt.Recovering("QDesignerActionEditorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Repaint(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) RepaintDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ResizeEvent
func callbackQDesignerActionEditorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::resizeEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::resizeEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_SetDisabled
func callbackQDesignerActionEditorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setDisabled", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setDisabled")
	}
}

func (ptr *QDesignerActionEditorInterface) SetDisabled(disable bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QDesignerActionEditorInterface) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerActionEditorInterface_SetFocus2
func callbackQDesignerActionEditorInterface_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setFocus2", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setFocus2")
	}
}

func (ptr *QDesignerActionEditorInterface) SetFocus2() {
	defer qt.Recovering("QDesignerActionEditorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) SetFocus2Default() {
	defer qt.Recovering("QDesignerActionEditorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_SetHidden
func callbackQDesignerActionEditorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setHidden", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::setHidden")
	}
}

func (ptr *QDesignerActionEditorInterface) SetHidden(hidden bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QDesignerActionEditorInterface) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QDesignerActionEditorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerActionEditorInterface_Show
func callbackQDesignerActionEditorInterface_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShow(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::show", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShow() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::show")
	}
}

func (ptr *QDesignerActionEditorInterface) Show() {
	defer qt.Recovering("QDesignerActionEditorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Show(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) ShowDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowFullScreen
func callbackQDesignerActionEditorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showFullScreen", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showFullScreen")
	}
}

func (ptr *QDesignerActionEditorInterface) ShowFullScreen() {
	defer qt.Recovering("QDesignerActionEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) ShowFullScreenDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowMaximized
func callbackQDesignerActionEditorInterface_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showMaximized", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showMaximized")
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMaximized() {
	defer qt.Recovering("QDesignerActionEditorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMaximizedDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowMinimized
func callbackQDesignerActionEditorInterface_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showMinimized", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showMinimized")
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMinimized() {
	defer qt.Recovering("QDesignerActionEditorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMinimizedDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowNormal
func callbackQDesignerActionEditorInterface_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showNormal", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::showNormal")
	}
}

func (ptr *QDesignerActionEditorInterface) ShowNormal() {
	defer qt.Recovering("QDesignerActionEditorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) ShowNormalDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_TabletEvent
func callbackQDesignerActionEditorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::tabletEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::tabletEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Update
func callbackQDesignerActionEditorInterface_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::update", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectUpdate() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::update")
	}
}

func (ptr *QDesignerActionEditorInterface) Update() {
	defer qt.Recovering("QDesignerActionEditorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_Update(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) UpdateDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_UpdateMicroFocus
func callbackQDesignerActionEditorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::updateMicroFocus", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::updateMicroFocus")
	}
}

func (ptr *QDesignerActionEditorInterface) UpdateMicroFocus() {
	defer qt.Recovering("QDesignerActionEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QDesignerActionEditorInterface) UpdateMicroFocusDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_WheelEvent
func callbackQDesignerActionEditorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::wheelEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::wheelEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_TimerEvent
func callbackQDesignerActionEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::timerEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::timerEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ChildEvent
func callbackQDesignerActionEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::childEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::childEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ConnectNotify
func callbackQDesignerActionEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::connectNotify", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::connectNotify")
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerActionEditorInterface_CustomEvent
func callbackQDesignerActionEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::customEvent", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::customEvent")
	}
}

func (ptr *QDesignerActionEditorInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerActionEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DeleteLater
func callbackQDesignerActionEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::deleteLater", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::deleteLater")
	}
}

func (ptr *QDesignerActionEditorInterface) DeleteLater() {
	defer qt.Recovering("QDesignerActionEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerActionEditorInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerActionEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerActionEditorInterface_DisconnectNotify
func callbackQDesignerActionEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerActionEditorInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::disconnectNotify")
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerActionEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerActionEditorInterface_EventFilter
func callbackQDesignerActionEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerActionEditorInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerActionEditorInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::eventFilter", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::eventFilter")
	}
}

func (ptr *QDesignerActionEditorInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerActionEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerActionEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_MetaObject
func callbackQDesignerActionEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerActionEditorInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerActionEditorInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerActionEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerActionEditorInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerActionEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::metaObject", f)
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerActionEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerActionEditorInterface::metaObject")
	}
}

func (ptr *QDesignerActionEditorInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerActionEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerActionEditorInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerActionEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerActionEditorInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDesignerContainerExtension struct {
	ptr unsafe.Pointer
}

type QDesignerContainerExtension_ITF interface {
	QDesignerContainerExtension_PTR() *QDesignerContainerExtension
}

func (p *QDesignerContainerExtension) QDesignerContainerExtension_PTR() *QDesignerContainerExtension {
	return p
}

func (p *QDesignerContainerExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerContainerExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerContainerExtension(ptr QDesignerContainerExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerContainerExtension_PTR().Pointer()
	}
	return nil
}

func NewQDesignerContainerExtensionFromPointer(ptr unsafe.Pointer) *QDesignerContainerExtension {
	var n = new(QDesignerContainerExtension)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerContainerExtension_AddWidget
func callbackQDesignerContainerExtension_AddWidget(ptr unsafe.Pointer, page unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerContainerExtension::addWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::addWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(page))
	}

}

func (ptr *QDesignerContainerExtension) ConnectAddWidget(f func(page *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerContainerExtension::addWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::addWidget", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectAddWidget(page widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerContainerExtension::addWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::addWidget")
	}
}

func (ptr *QDesignerContainerExtension) AddWidget(page widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerContainerExtension::addWidget")

	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_AddWidget(ptr.Pointer(), widgets.PointerFromQWidget(page))
	}
}

//export callbackQDesignerContainerExtension_CanAddWidget
func callbackQDesignerContainerExtension_CanAddWidget(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerContainerExtension::canAddWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::canAddWidget"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerContainerExtensionFromPointer(ptr).CanAddWidgetDefault())))
}

func (ptr *QDesignerContainerExtension) ConnectCanAddWidget(f func() bool) {
	defer qt.Recovering("connect QDesignerContainerExtension::canAddWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::canAddWidget", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCanAddWidget() {
	defer qt.Recovering("disconnect QDesignerContainerExtension::canAddWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::canAddWidget")
	}
}

func (ptr *QDesignerContainerExtension) CanAddWidget() bool {
	defer qt.Recovering("QDesignerContainerExtension::canAddWidget")

	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanAddWidget(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerContainerExtension) CanAddWidgetDefault() bool {
	defer qt.Recovering("QDesignerContainerExtension::canAddWidget")

	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanAddWidgetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerContainerExtension_CanRemove
func callbackQDesignerContainerExtension_CanRemove(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerContainerExtension::canRemove")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::canRemove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerContainerExtensionFromPointer(ptr).CanRemoveDefault(int(int32(index))))))
}

func (ptr *QDesignerContainerExtension) ConnectCanRemove(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerContainerExtension::canRemove")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::canRemove", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCanRemove() {
	defer qt.Recovering("disconnect QDesignerContainerExtension::canRemove")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::canRemove")
	}
}

func (ptr *QDesignerContainerExtension) CanRemove(index int) bool {
	defer qt.Recovering("QDesignerContainerExtension::canRemove")

	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanRemove(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QDesignerContainerExtension) CanRemoveDefault(index int) bool {
	defer qt.Recovering("QDesignerContainerExtension::canRemove")

	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanRemoveDefault(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerContainerExtension_Count
func callbackQDesignerContainerExtension_Count(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerContainerExtension::count")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerContainerExtension) ConnectCount(f func() int) {
	defer qt.Recovering("connect QDesignerContainerExtension::count")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::count", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCount() {
	defer qt.Recovering("disconnect QDesignerContainerExtension::count")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::count")
	}
}

func (ptr *QDesignerContainerExtension) Count() int {
	defer qt.Recovering("QDesignerContainerExtension::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerContainerExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerContainerExtension_CurrentIndex
func callbackQDesignerContainerExtension_CurrentIndex(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerContainerExtension::currentIndex")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::currentIndex"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerContainerExtension) ConnectCurrentIndex(f func() int) {
	defer qt.Recovering("connect QDesignerContainerExtension::currentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::currentIndex", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCurrentIndex() {
	defer qt.Recovering("disconnect QDesignerContainerExtension::currentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::currentIndex")
	}
}

func (ptr *QDesignerContainerExtension) CurrentIndex() int {
	defer qt.Recovering("QDesignerContainerExtension::currentIndex")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerContainerExtension_CurrentIndex(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerContainerExtension_InsertWidget
func callbackQDesignerContainerExtension_InsertWidget(ptr unsafe.Pointer, index C.int, page unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerContainerExtension::insertWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::insertWidget"); signal != nil {
		signal.(func(int, *widgets.QWidget))(int(int32(index)), widgets.NewQWidgetFromPointer(page))
	}

}

func (ptr *QDesignerContainerExtension) ConnectInsertWidget(f func(index int, page *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerContainerExtension::insertWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::insertWidget", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectInsertWidget(index int, page widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerContainerExtension::insertWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::insertWidget")
	}
}

func (ptr *QDesignerContainerExtension) InsertWidget(index int, page widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerContainerExtension::insertWidget")

	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_InsertWidget(ptr.Pointer(), C.int(int32(index)), widgets.PointerFromQWidget(page))
	}
}

//export callbackQDesignerContainerExtension_Remove
func callbackQDesignerContainerExtension_Remove(ptr unsafe.Pointer, index C.int) {
	defer qt.Recovering("callback QDesignerContainerExtension::remove")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::remove"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QDesignerContainerExtension) ConnectRemove(f func(index int)) {
	defer qt.Recovering("connect QDesignerContainerExtension::remove")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::remove", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectRemove(index int) {
	defer qt.Recovering("disconnect QDesignerContainerExtension::remove")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::remove")
	}
}

func (ptr *QDesignerContainerExtension) Remove(index int) {
	defer qt.Recovering("QDesignerContainerExtension::remove")

	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_Remove(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQDesignerContainerExtension_SetCurrentIndex
func callbackQDesignerContainerExtension_SetCurrentIndex(ptr unsafe.Pointer, index C.int) {
	defer qt.Recovering("callback QDesignerContainerExtension::setCurrentIndex")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::setCurrentIndex"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QDesignerContainerExtension) ConnectSetCurrentIndex(f func(index int)) {
	defer qt.Recovering("connect QDesignerContainerExtension::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::setCurrentIndex", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectSetCurrentIndex(index int) {
	defer qt.Recovering("disconnect QDesignerContainerExtension::setCurrentIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::setCurrentIndex")
	}
}

func (ptr *QDesignerContainerExtension) SetCurrentIndex(index int) {
	defer qt.Recovering("QDesignerContainerExtension::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_SetCurrentIndex(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQDesignerContainerExtension_Widget
func callbackQDesignerContainerExtension_Widget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerContainerExtension::widget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerContainerExtension::widget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerContainerExtension) ConnectWidget(f func(index int) *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerContainerExtension::widget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::widget", f)
	}
}

func (ptr *QDesignerContainerExtension) DisconnectWidget(index int) {
	defer qt.Recovering("disconnect QDesignerContainerExtension::widget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerContainerExtension::widget")
	}
}

func (ptr *QDesignerContainerExtension) Widget(index int) *widgets.QWidget {
	defer qt.Recovering("QDesignerContainerExtension::widget")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerContainerExtension_Widget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerContainerExtension) DestroyQDesignerContainerExtension() {
	defer qt.Recovering("QDesignerContainerExtension::~QDesignerContainerExtension")

	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_DestroyQDesignerContainerExtension(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerCustomWidgetCollectionInterface struct {
	ptr unsafe.Pointer
}

type QDesignerCustomWidgetCollectionInterface_ITF interface {
	QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface
}

func (p *QDesignerCustomWidgetCollectionInterface) QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface {
	return p
}

func (p *QDesignerCustomWidgetCollectionInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerCustomWidgetCollectionInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerCustomWidgetCollectionInterface(ptr QDesignerCustomWidgetCollectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerCustomWidgetCollectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerCustomWidgetCollectionInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerCustomWidgetCollectionInterface {
	var n = new(QDesignerCustomWidgetCollectionInterface)
	n.SetPointer(ptr)
	return n
}
func (ptr *QDesignerCustomWidgetCollectionInterface) DestroyQDesignerCustomWidgetCollectionInterface() {
	defer qt.Recovering("QDesignerCustomWidgetCollectionInterface::~QDesignerCustomWidgetCollectionInterface")

	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerCustomWidgetInterface struct {
	ptr unsafe.Pointer
}

type QDesignerCustomWidgetInterface_ITF interface {
	QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface
}

func (p *QDesignerCustomWidgetInterface) QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface {
	return p
}

func (p *QDesignerCustomWidgetInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerCustomWidgetInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerCustomWidgetInterface(ptr QDesignerCustomWidgetInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerCustomWidgetInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerCustomWidgetInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerCustomWidgetInterface {
	var n = new(QDesignerCustomWidgetInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerCustomWidgetInterface_CodeTemplate
func callbackQDesignerCustomWidgetInterface_CodeTemplate(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::codeTemplate")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::codeTemplate"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQDesignerCustomWidgetInterfaceFromPointer(ptr).CodeTemplateDefault())
}

func (ptr *QDesignerCustomWidgetInterface) ConnectCodeTemplate(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::codeTemplate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::codeTemplate", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCodeTemplate() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::codeTemplate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::codeTemplate")
	}
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplate() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::codeTemplate")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_CodeTemplate(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplateDefault() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::codeTemplate")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_CodeTemplateDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_CreateWidget
func callbackQDesignerCustomWidgetInterface_CreateWidget(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::createWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::createWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*widgets.QWidget) *widgets.QWidget)(widgets.NewQWidgetFromPointer(parent)))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectCreateWidget(f func(parent *widgets.QWidget) *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::createWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::createWidget", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCreateWidget(parent widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::createWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::createWidget")
	}
}

func (ptr *QDesignerCustomWidgetInterface) CreateWidget(parent widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QDesignerCustomWidgetInterface::createWidget")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerCustomWidgetInterface_CreateWidget(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerCustomWidgetInterface_DomXml
func callbackQDesignerCustomWidgetInterface_DomXml(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::domXml")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::domXml"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQDesignerCustomWidgetInterfaceFromPointer(ptr).DomXmlDefault())
}

func (ptr *QDesignerCustomWidgetInterface) ConnectDomXml(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::domXml")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::domXml", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectDomXml() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::domXml")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::domXml")
	}
}

func (ptr *QDesignerCustomWidgetInterface) DomXml() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::domXml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_DomXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerCustomWidgetInterface) DomXmlDefault() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::domXml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_DomXmlDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_Group
func callbackQDesignerCustomWidgetInterface_Group(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::group")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::group"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerCustomWidgetInterface) ConnectGroup(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::group")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::group", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectGroup() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::group")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::group")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Group() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::group")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_Group(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_Icon
func callbackQDesignerCustomWidgetInterface_Icon(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::icon")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::icon"); signal != nil {
		return gui.PointerFromQIcon(signal.(func() *gui.QIcon)())
	}

	return gui.PointerFromQIcon(nil)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIcon(f func() *gui.QIcon) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::icon")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::icon", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIcon() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::icon")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::icon")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Icon() *gui.QIcon {
	defer qt.Recovering("QDesignerCustomWidgetInterface::icon")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QDesignerCustomWidgetInterface_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerCustomWidgetInterface_IncludeFile
func callbackQDesignerCustomWidgetInterface_IncludeFile(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::includeFile")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::includeFile"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIncludeFile(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::includeFile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::includeFile", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIncludeFile() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::includeFile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::includeFile")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IncludeFile() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::includeFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_IncludeFile(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_Initialize
func callbackQDesignerCustomWidgetInterface_Initialize(ptr unsafe.Pointer, formEditor unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::initialize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::initialize"); signal != nil {
		signal.(func(*QDesignerFormEditorInterface))(NewQDesignerFormEditorInterfaceFromPointer(formEditor))
	} else {
		NewQDesignerCustomWidgetInterfaceFromPointer(ptr).InitializeDefault(NewQDesignerFormEditorInterfaceFromPointer(formEditor))
	}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectInitialize(f func(formEditor *QDesignerFormEditorInterface)) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::initialize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::initialize", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectInitialize() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::initialize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::initialize")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Initialize(formEditor QDesignerFormEditorInterface_ITF) {
	defer qt.Recovering("QDesignerCustomWidgetInterface::initialize")

	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_Initialize(ptr.Pointer(), PointerFromQDesignerFormEditorInterface(formEditor))
	}
}

func (ptr *QDesignerCustomWidgetInterface) InitializeDefault(formEditor QDesignerFormEditorInterface_ITF) {
	defer qt.Recovering("QDesignerCustomWidgetInterface::initialize")

	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_InitializeDefault(ptr.Pointer(), PointerFromQDesignerFormEditorInterface(formEditor))
	}
}

//export callbackQDesignerCustomWidgetInterface_IsContainer
func callbackQDesignerCustomWidgetInterface_IsContainer(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::isContainer")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::isContainer"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsContainer(f func() bool) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::isContainer")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::isContainer", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsContainer() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::isContainer")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::isContainer")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IsContainer() bool {
	defer qt.Recovering("QDesignerCustomWidgetInterface::isContainer")

	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsContainer(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerCustomWidgetInterface_IsInitialized
func callbackQDesignerCustomWidgetInterface_IsInitialized(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::isInitialized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::isInitialized"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerCustomWidgetInterfaceFromPointer(ptr).IsInitializedDefault())))
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsInitialized(f func() bool) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::isInitialized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::isInitialized", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsInitialized() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::isInitialized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::isInitialized")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IsInitialized() bool {
	defer qt.Recovering("QDesignerCustomWidgetInterface::isInitialized")

	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerCustomWidgetInterface) IsInitializedDefault() bool {
	defer qt.Recovering("QDesignerCustomWidgetInterface::isInitialized")

	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsInitializedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerCustomWidgetInterface_Name
func callbackQDesignerCustomWidgetInterface_Name(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::name")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::name"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerCustomWidgetInterface) ConnectName(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::name")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::name", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectName() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::name")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::name")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Name() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_ToolTip
func callbackQDesignerCustomWidgetInterface_ToolTip(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::toolTip")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::toolTip"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerCustomWidgetInterface) ConnectToolTip(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::toolTip")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::toolTip", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectToolTip() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::toolTip")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::toolTip")
	}
}

func (ptr *QDesignerCustomWidgetInterface) ToolTip() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_WhatsThis
func callbackQDesignerCustomWidgetInterface_WhatsThis(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerCustomWidgetInterface::whatsThis")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerCustomWidgetInterface::whatsThis"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerCustomWidgetInterface) ConnectWhatsThis(f func() string) {
	defer qt.Recovering("connect QDesignerCustomWidgetInterface::whatsThis")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::whatsThis", f)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectWhatsThis() {
	defer qt.Recovering("disconnect QDesignerCustomWidgetInterface::whatsThis")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerCustomWidgetInterface::whatsThis")
	}
}

func (ptr *QDesignerCustomWidgetInterface) WhatsThis() string {
	defer qt.Recovering("QDesignerCustomWidgetInterface::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerCustomWidgetInterface_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerCustomWidgetInterface) DestroyQDesignerCustomWidgetInterface() {
	defer qt.Recovering("QDesignerCustomWidgetInterface::~QDesignerCustomWidgetInterface")

	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerDynamicPropertySheetExtension struct {
	ptr unsafe.Pointer
}

type QDesignerDynamicPropertySheetExtension_ITF interface {
	QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension
}

func (p *QDesignerDynamicPropertySheetExtension) QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension {
	return p
}

func (p *QDesignerDynamicPropertySheetExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerDynamicPropertySheetExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerDynamicPropertySheetExtension(ptr QDesignerDynamicPropertySheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerDynamicPropertySheetExtension_PTR().Pointer()
	}
	return nil
}

func NewQDesignerDynamicPropertySheetExtensionFromPointer(ptr unsafe.Pointer) *QDesignerDynamicPropertySheetExtension {
	var n = new(QDesignerDynamicPropertySheetExtension)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty(ptr unsafe.Pointer, propertyName *C.char, value unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerDynamicPropertySheetExtension::addDynamicProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerDynamicPropertySheetExtension::addDynamicProperty"); signal != nil {
		return C.int(int32(signal.(func(string, *core.QVariant) int)(C.GoString(propertyName), core.NewQVariantFromPointer(value))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectAddDynamicProperty(f func(propertyName string, value *core.QVariant) int) {
	defer qt.Recovering("connect QDesignerDynamicPropertySheetExtension::addDynamicProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::addDynamicProperty", f)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectAddDynamicProperty(propertyName string, value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QDesignerDynamicPropertySheetExtension::addDynamicProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::addDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) AddDynamicProperty(propertyName string, value core.QVariant_ITF) int {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::addDynamicProperty")

	if ptr.Pointer() != nil {
		var propertyNameC = C.CString(propertyName)
		defer C.free(unsafe.Pointer(propertyNameC))
		return int(int32(C.QDesignerDynamicPropertySheetExtension_AddDynamicProperty(ptr.Pointer(), propertyNameC, core.PointerFromQVariant(value))))
	}
	return 0
}

//export callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(ptr unsafe.Pointer, propertyName *C.char) C.char {
	defer qt.Recovering("callback QDesignerDynamicPropertySheetExtension::canAddDynamicProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerDynamicPropertySheetExtension::canAddDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(propertyName)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectCanAddDynamicProperty(f func(propertyName string) bool) {
	defer qt.Recovering("connect QDesignerDynamicPropertySheetExtension::canAddDynamicProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::canAddDynamicProperty", f)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectCanAddDynamicProperty(propertyName string) {
	defer qt.Recovering("disconnect QDesignerDynamicPropertySheetExtension::canAddDynamicProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::canAddDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) CanAddDynamicProperty(propertyName string) bool {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::canAddDynamicProperty")

	if ptr.Pointer() != nil {
		var propertyNameC = C.CString(propertyName)
		defer C.free(unsafe.Pointer(propertyNameC))
		return C.QDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(ptr.Pointer(), propertyNameC) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed
func callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectDynamicPropertiesAllowed(f func() bool) {
	defer qt.Recovering("connect QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed", f)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectDynamicPropertiesAllowed() {
	defer qt.Recovering("disconnect QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DynamicPropertiesAllowed() bool {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::dynamicPropertiesAllowed")

	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerDynamicPropertySheetExtension::isDynamicProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerDynamicPropertySheetExtension::isDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectIsDynamicProperty(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerDynamicPropertySheetExtension::isDynamicProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::isDynamicProperty", f)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectIsDynamicProperty(index int) {
	defer qt.Recovering("disconnect QDesignerDynamicPropertySheetExtension::isDynamicProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::isDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) IsDynamicProperty(index int) bool {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::isDynamicProperty")

	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_IsDynamicProperty(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerDynamicPropertySheetExtension::removeDynamicProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerDynamicPropertySheetExtension::removeDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectRemoveDynamicProperty(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerDynamicPropertySheetExtension::removeDynamicProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::removeDynamicProperty", f)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectRemoveDynamicProperty(index int) {
	defer qt.Recovering("disconnect QDesignerDynamicPropertySheetExtension::removeDynamicProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerDynamicPropertySheetExtension::removeDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) RemoveDynamicProperty(index int) bool {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::removeDynamicProperty")

	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QDesignerDynamicPropertySheetExtension) DestroyQDesignerDynamicPropertySheetExtension() {
	defer qt.Recovering("QDesignerDynamicPropertySheetExtension::~QDesignerDynamicPropertySheetExtension")

	if ptr.Pointer() != nil {
		C.QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerFormEditorInterface struct {
	core.QObject
}

type QDesignerFormEditorInterface_ITF interface {
	core.QObject_ITF
	QDesignerFormEditorInterface_PTR() *QDesignerFormEditorInterface
}

func (p *QDesignerFormEditorInterface) QDesignerFormEditorInterface_PTR() *QDesignerFormEditorInterface {
	return p
}

func (p *QDesignerFormEditorInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerFormEditorInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerFormEditorInterface(ptr QDesignerFormEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormEditorInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerFormEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormEditorInterface {
	var n = new(QDesignerFormEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerFormEditorInterface(parent core.QObject_ITF) *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::QDesignerFormEditorInterface")

	var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_NewQDesignerFormEditorInterface(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QDesignerFormEditorInterface) ActionEditor() *QDesignerActionEditorInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::actionEditor")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerActionEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_ActionEditor(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) ExtensionManager() *QExtensionManager {
	defer qt.Recovering("QDesignerFormEditorInterface::extensionManager")

	if ptr.Pointer() != nil {
		var tmpValue = NewQExtensionManagerFromPointer(C.QDesignerFormEditorInterface_ExtensionManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) FormWindowManager() *QDesignerFormWindowManagerInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::formWindowManager")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowManagerInterfaceFromPointer(C.QDesignerFormEditorInterface_FormWindowManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) ObjectInspector() *QDesignerObjectInspectorInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::objectInspector")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerObjectInspectorInterfaceFromPointer(C.QDesignerFormEditorInterface_ObjectInspector(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) PropertyEditor() *QDesignerPropertyEditorInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::propertyEditor")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerPropertyEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_PropertyEditor(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) SetActionEditor(actionEditor QDesignerActionEditorInterface_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::setActionEditor")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetActionEditor(ptr.Pointer(), PointerFromQDesignerActionEditorInterface(actionEditor))
	}
}

func (ptr *QDesignerFormEditorInterface) SetObjectInspector(objectInspector QDesignerObjectInspectorInterface_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::setObjectInspector")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetObjectInspector(ptr.Pointer(), PointerFromQDesignerObjectInspectorInterface(objectInspector))
	}
}

func (ptr *QDesignerFormEditorInterface) SetPropertyEditor(propertyEditor QDesignerPropertyEditorInterface_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::setPropertyEditor")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetPropertyEditor(ptr.Pointer(), PointerFromQDesignerPropertyEditorInterface(propertyEditor))
	}
}

func (ptr *QDesignerFormEditorInterface) SetWidgetBox(widgetBox QDesignerWidgetBoxInterface_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::setWidgetBox")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetWidgetBox(ptr.Pointer(), PointerFromQDesignerWidgetBoxInterface(widgetBox))
	}
}

func (ptr *QDesignerFormEditorInterface) TopLevel() *widgets.QWidget {
	defer qt.Recovering("QDesignerFormEditorInterface::topLevel")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormEditorInterface_TopLevel(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) WidgetBox() *QDesignerWidgetBoxInterface {
	defer qt.Recovering("QDesignerFormEditorInterface::widgetBox")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerWidgetBoxInterfaceFromPointer(C.QDesignerFormEditorInterface_WidgetBox(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) DestroyQDesignerFormEditorInterface() {
	defer qt.Recovering("QDesignerFormEditorInterface::~QDesignerFormEditorInterface")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormEditorInterface_TimerEvent
func callbackQDesignerFormEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::timerEvent", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::timerEvent")
	}
}

func (ptr *QDesignerFormEditorInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerFormEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_ChildEvent
func callbackQDesignerFormEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::childEvent", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::childEvent")
	}
}

func (ptr *QDesignerFormEditorInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerFormEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_ConnectNotify
func callbackQDesignerFormEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::connectNotify", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::connectNotify")
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormEditorInterface_CustomEvent
func callbackQDesignerFormEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::customEvent", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::customEvent")
	}
}

func (ptr *QDesignerFormEditorInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_DeleteLater
func callbackQDesignerFormEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::deleteLater", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::deleteLater")
	}
}

func (ptr *QDesignerFormEditorInterface) DeleteLater() {
	defer qt.Recovering("QDesignerFormEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerFormEditorInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerFormEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormEditorInterface_DisconnectNotify
func callbackQDesignerFormEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormEditorInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::disconnectNotify")
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormEditorInterface_Event
func callbackQDesignerFormEditorInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormEditorInterface::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormEditorInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDesignerFormEditorInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::event", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::event")
	}
}

func (ptr *QDesignerFormEditorInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormEditorInterface::event")

	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDesignerFormEditorInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormEditorInterface::event")

	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDesignerFormEditorInterface_EventFilter
func callbackQDesignerFormEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormEditorInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormEditorInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::eventFilter", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::eventFilter")
	}
}

func (ptr *QDesignerFormEditorInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerFormEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormEditorInterface_MetaObject
func callbackQDesignerFormEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormEditorInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormEditorInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormEditorInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerFormEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::metaObject", f)
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerFormEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormEditorInterface::metaObject")
	}
}

func (ptr *QDesignerFormEditorInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormEditorInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormEditorInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QDesignerFormWindowCursorInterface::MoveMode
type QDesignerFormWindowCursorInterface__MoveMode int64

const (
	QDesignerFormWindowCursorInterface__MoveAnchor = QDesignerFormWindowCursorInterface__MoveMode(0)
	QDesignerFormWindowCursorInterface__KeepAnchor = QDesignerFormWindowCursorInterface__MoveMode(1)
)

//QDesignerFormWindowCursorInterface::MoveOperation
type QDesignerFormWindowCursorInterface__MoveOperation int64

const (
	QDesignerFormWindowCursorInterface__NoMove = QDesignerFormWindowCursorInterface__MoveOperation(0)
	QDesignerFormWindowCursorInterface__Start  = QDesignerFormWindowCursorInterface__MoveOperation(1)
	QDesignerFormWindowCursorInterface__End    = QDesignerFormWindowCursorInterface__MoveOperation(2)
	QDesignerFormWindowCursorInterface__Next   = QDesignerFormWindowCursorInterface__MoveOperation(3)
	QDesignerFormWindowCursorInterface__Prev   = QDesignerFormWindowCursorInterface__MoveOperation(4)
	QDesignerFormWindowCursorInterface__Left   = QDesignerFormWindowCursorInterface__MoveOperation(5)
	QDesignerFormWindowCursorInterface__Right  = QDesignerFormWindowCursorInterface__MoveOperation(6)
	QDesignerFormWindowCursorInterface__Up     = QDesignerFormWindowCursorInterface__MoveOperation(7)
	QDesignerFormWindowCursorInterface__Down   = QDesignerFormWindowCursorInterface__MoveOperation(8)
)

type QDesignerFormWindowCursorInterface struct {
	ptr unsafe.Pointer
}

type QDesignerFormWindowCursorInterface_ITF interface {
	QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface
}

func (p *QDesignerFormWindowCursorInterface) QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface {
	return p
}

func (p *QDesignerFormWindowCursorInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerFormWindowCursorInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerFormWindowCursorInterface(ptr QDesignerFormWindowCursorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowCursorInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerFormWindowCursorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormWindowCursorInterface {
	var n = new(QDesignerFormWindowCursorInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerFormWindowCursorInterface_Current
func callbackQDesignerFormWindowCursorInterface_Current(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::current")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::current"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func() *widgets.QWidget)())
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectCurrent(f func() *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::current")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::current", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectCurrent() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::current")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::current")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Current() *widgets.QWidget {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::current")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_Current(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_FormWindow
func callbackQDesignerFormWindowCursorInterface_FormWindow(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::formWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::formWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func() *QDesignerFormWindowInterface)())
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectFormWindow(f func() *QDesignerFormWindowInterface) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::formWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::formWindow", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectFormWindow() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::formWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::formWindow")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) FormWindow() *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::formWindow")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowCursorInterface_FormWindow(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_HasSelection
func callbackQDesignerFormWindowCursorInterface_HasSelection(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::hasSelection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::hasSelection"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectHasSelection(f func() bool) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::hasSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::hasSelection", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectHasSelection() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::hasSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::hasSelection")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) HasSelection() bool {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::hasSelection")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowCursorInterface) IsWidgetSelected(widget widgets.QWidget_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::isWidgetSelected")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_IsWidgetSelected(ptr.Pointer(), widgets.PointerFromQWidget(widget)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowCursorInterface_MovePosition
func callbackQDesignerFormWindowCursorInterface_MovePosition(ptr unsafe.Pointer, operation C.longlong, mode C.longlong) C.char {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::movePosition")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::movePosition"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QDesignerFormWindowCursorInterface__MoveOperation, QDesignerFormWindowCursorInterface__MoveMode) bool)(QDesignerFormWindowCursorInterface__MoveOperation(operation), QDesignerFormWindowCursorInterface__MoveMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectMovePosition(f func(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::movePosition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::movePosition", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectMovePosition(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::movePosition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::movePosition")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) MovePosition(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::movePosition")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_MovePosition(ptr.Pointer(), C.longlong(operation), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowCursorInterface_Position
func callbackQDesignerFormWindowCursorInterface_Position(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::position")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::position"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectPosition(f func() int) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::position")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::position", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectPosition() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::position")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::position")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Position() int {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::position")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_Position(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty
func callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty(ptr unsafe.Pointer, widget unsafe.Pointer, name *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::resetWidgetProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::resetWidgetProperty"); signal != nil {
		signal.(func(*widgets.QWidget, string))(widgets.NewQWidgetFromPointer(widget), C.GoString(name))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectResetWidgetProperty(f func(widget *widgets.QWidget, name string)) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::resetWidgetProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::resetWidgetProperty", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectResetWidgetProperty(widget widgets.QWidget_ITF, name string) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::resetWidgetProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::resetWidgetProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) ResetWidgetProperty(widget widgets.QWidget_ITF, name string) {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::resetWidgetProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDesignerFormWindowCursorInterface_ResetWidgetProperty(ptr.Pointer(), widgets.PointerFromQWidget(widget), nameC)
	}
}

//export callbackQDesignerFormWindowCursorInterface_SelectedWidget
func callbackQDesignerFormWindowCursorInterface_SelectedWidget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::selectedWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::selectedWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidget(f func(index int) *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::selectedWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::selectedWidget", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidget(index int) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::selectedWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::selectedWidget")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidget(index int) *widgets.QWidget {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::selectedWidget")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_SelectedWidget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount
func callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::selectedWidgetCount")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::selectedWidgetCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidgetCount(f func() int) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::selectedWidgetCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::selectedWidgetCount", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidgetCount() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::selectedWidgetCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::selectedWidgetCount")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidgetCount() int {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::selectedWidgetCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_SelectedWidgetCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerFormWindowCursorInterface_SetPosition
func callbackQDesignerFormWindowCursorInterface_SetPosition(ptr unsafe.Pointer, position C.int, mode C.longlong) {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::setPosition")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::setPosition"); signal != nil {
		signal.(func(int, QDesignerFormWindowCursorInterface__MoveMode))(int(int32(position)), QDesignerFormWindowCursorInterface__MoveMode(mode))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetPosition(f func(position int, mode QDesignerFormWindowCursorInterface__MoveMode)) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::setPosition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setPosition", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetPosition(position int, mode QDesignerFormWindowCursorInterface__MoveMode) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::setPosition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setPosition")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetPosition(position int, mode QDesignerFormWindowCursorInterface__MoveMode) {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::setPosition")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowCursorInterface_SetPosition(ptr.Pointer(), C.int(int32(position)), C.longlong(mode))
	}
}

//export callbackQDesignerFormWindowCursorInterface_SetProperty
func callbackQDesignerFormWindowCursorInterface_SetProperty(ptr unsafe.Pointer, name *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::setProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::setProperty"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetProperty(f func(name string, value *core.QVariant)) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::setProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setProperty", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetProperty(name string, value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::setProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetProperty(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::setProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDesignerFormWindowCursorInterface_SetProperty(ptr.Pointer(), nameC, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerFormWindowCursorInterface_SetWidgetProperty
func callbackQDesignerFormWindowCursorInterface_SetWidgetProperty(ptr unsafe.Pointer, widget unsafe.Pointer, name *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::setWidgetProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::setWidgetProperty"); signal != nil {
		signal.(func(*widgets.QWidget, string, *core.QVariant))(widgets.NewQWidgetFromPointer(widget), C.GoString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetWidgetProperty(f func(widget *widgets.QWidget, name string, value *core.QVariant)) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::setWidgetProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setWidgetProperty", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetWidgetProperty(widget widgets.QWidget_ITF, name string, value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::setWidgetProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::setWidgetProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetWidgetProperty(widget widgets.QWidget_ITF, name string, value core.QVariant_ITF) {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::setWidgetProperty")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDesignerFormWindowCursorInterface_SetWidgetProperty(ptr.Pointer(), widgets.PointerFromQWidget(widget), nameC, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerFormWindowCursorInterface_Widget
func callbackQDesignerFormWindowCursorInterface_Widget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::widget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::widget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidget(f func(index int) *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::widget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::widget", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidget(index int) {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::widget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::widget")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Widget(index int) *widgets.QWidget {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::widget")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_Widget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_WidgetCount
func callbackQDesignerFormWindowCursorInterface_WidgetCount(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerFormWindowCursorInterface::widgetCount")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowCursorInterface::widgetCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidgetCount(f func() int) {
	defer qt.Recovering("connect QDesignerFormWindowCursorInterface::widgetCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::widgetCount", f)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidgetCount() {
	defer qt.Recovering("disconnect QDesignerFormWindowCursorInterface::widgetCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowCursorInterface::widgetCount")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) WidgetCount() int {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::widgetCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_WidgetCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDesignerFormWindowCursorInterface) DestroyQDesignerFormWindowCursorInterface() {
	defer qt.Recovering("QDesignerFormWindowCursorInterface::~QDesignerFormWindowCursorInterface")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//QDesignerFormWindowInterface::FeatureFlag
type QDesignerFormWindowInterface__FeatureFlag int64

const (
	QDesignerFormWindowInterface__EditFeature     = QDesignerFormWindowInterface__FeatureFlag(0x01)
	QDesignerFormWindowInterface__GridFeature     = QDesignerFormWindowInterface__FeatureFlag(0x02)
	QDesignerFormWindowInterface__TabOrderFeature = QDesignerFormWindowInterface__FeatureFlag(0x04)
	QDesignerFormWindowInterface__DefaultFeature  = QDesignerFormWindowInterface__FeatureFlag(QDesignerFormWindowInterface__EditFeature | QDesignerFormWindowInterface__GridFeature)
)

//QDesignerFormWindowInterface::ResourceFileSaveMode
type QDesignerFormWindowInterface__ResourceFileSaveMode int64

const (
	QDesignerFormWindowInterface__SaveAllResourceFiles      = QDesignerFormWindowInterface__ResourceFileSaveMode(0)
	QDesignerFormWindowInterface__SaveOnlyUsedResourceFiles = QDesignerFormWindowInterface__ResourceFileSaveMode(1)
	QDesignerFormWindowInterface__DontSaveResourceFiles     = QDesignerFormWindowInterface__ResourceFileSaveMode(2)
)

type QDesignerFormWindowInterface struct {
	widgets.QWidget
}

type QDesignerFormWindowInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerFormWindowInterface_PTR() *QDesignerFormWindowInterface
}

func (p *QDesignerFormWindowInterface) QDesignerFormWindowInterface_PTR() *QDesignerFormWindowInterface {
	return p
}

func (p *QDesignerFormWindowInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerFormWindowInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerFormWindowInterface(ptr QDesignerFormWindowInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerFormWindowInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormWindowInterface {
	var n = new(QDesignerFormWindowInterface)
	n.SetPointer(ptr)
	return n
}
func QDesignerFormWindowInterface_FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::findFormWindow")

	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::findFormWindow")

	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func QDesignerFormWindowInterface_FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::findFormWindow")

	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(widgets.PointerFromQWidget(widget)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::findFormWindow")

	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(widgets.PointerFromQWidget(widget)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQDesignerFormWindowInterface_AboutToUnmanageWidget
func callbackQDesignerFormWindowInterface_AboutToUnmanageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::aboutToUnmanageWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::aboutToUnmanageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectAboutToUnmanageWidget(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::aboutToUnmanageWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectAboutToUnmanageWidget(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::aboutToUnmanageWidget", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAboutToUnmanageWidget() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::aboutToUnmanageWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectAboutToUnmanageWidget(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::aboutToUnmanageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) AboutToUnmanageWidget(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::aboutToUnmanageWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_AboutToUnmanageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_AbsoluteDir
func callbackQDesignerFormWindowInterface_AbsoluteDir(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::absoluteDir")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::absoluteDir"); signal != nil {
		return core.PointerFromQDir(signal.(func() *core.QDir)())
	}

	return core.PointerFromQDir(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectAbsoluteDir(f func() *core.QDir) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::absoluteDir")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::absoluteDir", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAbsoluteDir() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::absoluteDir")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::absoluteDir")
	}
}

func (ptr *QDesignerFormWindowInterface) AbsoluteDir() *core.QDir {
	defer qt.Recovering("QDesignerFormWindowInterface::absoluteDir")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDirFromPointer(C.QDesignerFormWindowInterface_AbsoluteDir(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_ActivateResourceFilePaths
func callbackQDesignerFormWindowInterface_ActivateResourceFilePaths(ptr unsafe.Pointer, paths *C.char, errorCount C.int, errorMessages *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::activateResourceFilePaths")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::activateResourceFilePaths"); signal != nil {
		signal.(func([]string, int, string))(strings.Split(C.GoString(paths), "|"), int(int32(errorCount)), C.GoString(errorMessages))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectActivateResourceFilePaths(f func(paths []string, errorCount int, errorMessages string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::activateResourceFilePaths")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::activateResourceFilePaths", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivateResourceFilePaths(paths []string, errorCount int, errorMessages string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::activateResourceFilePaths")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::activateResourceFilePaths")
	}
}

func (ptr *QDesignerFormWindowInterface) ActivateResourceFilePaths(paths []string, errorCount int, errorMessages string) {
	defer qt.Recovering("QDesignerFormWindowInterface::activateResourceFilePaths")

	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		var errorMessagesC = C.CString(errorMessages)
		defer C.free(unsafe.Pointer(errorMessagesC))
		C.QDesignerFormWindowInterface_ActivateResourceFilePaths(ptr.Pointer(), pathsC, C.int(int32(errorCount)), errorMessagesC)
	}
}

//export callbackQDesignerFormWindowInterface_Activated
func callbackQDesignerFormWindowInterface_Activated(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::activated")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::activated"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectActivated(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::activated")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::activated", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivated() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::activated")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::activated")
	}
}

func (ptr *QDesignerFormWindowInterface) Activated(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::activated")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Activated(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QDesignerFormWindowInterface) ActiveResourceFilePaths() []string {
	defer qt.Recovering("QDesignerFormWindowInterface::activeResourceFilePaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDesignerFormWindowInterface_ActiveResourceFilePaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_AddResourceFile
func callbackQDesignerFormWindowInterface_AddResourceFile(ptr unsafe.Pointer, path *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::addResourceFile")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::addResourceFile"); signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectAddResourceFile(f func(path string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::addResourceFile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::addResourceFile", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAddResourceFile(path string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::addResourceFile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::addResourceFile")
	}
}

func (ptr *QDesignerFormWindowInterface) AddResourceFile(path string) {
	defer qt.Recovering("QDesignerFormWindowInterface::addResourceFile")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDesignerFormWindowInterface_AddResourceFile(ptr.Pointer(), pathC)
	}
}

//export callbackQDesignerFormWindowInterface_Author
func callbackQDesignerFormWindowInterface_Author(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::author")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::author"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectAuthor(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::author")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::author", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAuthor() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::author")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::author")
	}
}

func (ptr *QDesignerFormWindowInterface) Author() string {
	defer qt.Recovering("QDesignerFormWindowInterface::author")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_Author(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_Changed
func callbackQDesignerFormWindowInterface_Changed(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::changed")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::changed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectChanged(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::changed")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::changed", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::changed")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::changed")
	}
}

func (ptr *QDesignerFormWindowInterface) Changed() {
	defer qt.Recovering("QDesignerFormWindowInterface::changed")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Changed(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_CheckContents
func callbackQDesignerFormWindowInterface_CheckContents(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::checkContents")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::checkContents"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QDesignerFormWindowInterface) ConnectCheckContents(f func() []string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::checkContents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::checkContents", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCheckContents() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::checkContents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::checkContents")
	}
}

func (ptr *QDesignerFormWindowInterface) CheckContents() []string {
	defer qt.Recovering("QDesignerFormWindowInterface::checkContents")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDesignerFormWindowInterface_CheckContents(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_ClearSelection
func callbackQDesignerFormWindowInterface_ClearSelection(ptr unsafe.Pointer, update C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::clearSelection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::clearSelection"); signal != nil {
		signal.(func(bool))(int8(update) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectClearSelection(f func(update bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::clearSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::clearSelection", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectClearSelection(update bool) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::clearSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::clearSelection")
	}
}

func (ptr *QDesignerFormWindowInterface) ClearSelection(update bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::clearSelection")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ClearSelection(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(update))))
	}
}

//export callbackQDesignerFormWindowInterface_Comment
func callbackQDesignerFormWindowInterface_Comment(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::comment")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::comment"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectComment(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::comment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::comment", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectComment() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::comment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::comment")
	}
}

func (ptr *QDesignerFormWindowInterface) Comment() string {
	defer qt.Recovering("QDesignerFormWindowInterface::comment")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_Comment(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_Contents
func callbackQDesignerFormWindowInterface_Contents(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::contents")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::contents"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectContents(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::contents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::contents", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectContents() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::contents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::contents")
	}
}

func (ptr *QDesignerFormWindowInterface) Contents() string {
	defer qt.Recovering("QDesignerFormWindowInterface::contents")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_Contents(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_Core
func callbackQDesignerFormWindowInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::core")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerFormWindowInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerFormWindowInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::core")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::core", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCore() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::core")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::core")
	}
}

func (ptr *QDesignerFormWindowInterface) Core() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) CoreDefault() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_Cursor
func callbackQDesignerFormWindowInterface_Cursor(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::cursor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::cursor"); signal != nil {
		return PointerFromQDesignerFormWindowCursorInterface(signal.(func() *QDesignerFormWindowCursorInterface)())
	}

	return PointerFromQDesignerFormWindowCursorInterface(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectCursor(f func() *QDesignerFormWindowCursorInterface) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::cursor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::cursor", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCursor() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::cursor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::cursor")
	}
}

func (ptr *QDesignerFormWindowInterface) Cursor() *QDesignerFormWindowCursorInterface {
	defer qt.Recovering("QDesignerFormWindowInterface::cursor")

	if ptr.Pointer() != nil {
		return NewQDesignerFormWindowCursorInterfaceFromPointer(C.QDesignerFormWindowInterface_Cursor(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_EmitSelectionChanged
func callbackQDesignerFormWindowInterface_EmitSelectionChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::emitSelectionChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::emitSelectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectEmitSelectionChanged(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::emitSelectionChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::emitSelectionChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectEmitSelectionChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::emitSelectionChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::emitSelectionChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) EmitSelectionChanged() {
	defer qt.Recovering("QDesignerFormWindowInterface::emitSelectionChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_EmitSelectionChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ExportMacro
func callbackQDesignerFormWindowInterface_ExportMacro(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::exportMacro")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::exportMacro"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectExportMacro(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::exportMacro")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::exportMacro", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectExportMacro() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::exportMacro")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::exportMacro")
	}
}

func (ptr *QDesignerFormWindowInterface) ExportMacro() string {
	defer qt.Recovering("QDesignerFormWindowInterface::exportMacro")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_ExportMacro(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_FeatureChanged
func callbackQDesignerFormWindowInterface_FeatureChanged(ptr unsafe.Pointer, feature C.longlong) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::featureChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::featureChanged"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__FeatureFlag))(QDesignerFormWindowInterface__FeatureFlag(feature))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectFeatureChanged(f func(feature QDesignerFormWindowInterface__FeatureFlag)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::featureChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectFeatureChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::featureChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatureChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::featureChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectFeatureChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::featureChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) FeatureChanged(feature QDesignerFormWindowInterface__FeatureFlag) {
	defer qt.Recovering("QDesignerFormWindowInterface::featureChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FeatureChanged(ptr.Pointer(), C.longlong(feature))
	}
}

//export callbackQDesignerFormWindowInterface_Features
func callbackQDesignerFormWindowInterface_Features(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QDesignerFormWindowInterface::features")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::features"); signal != nil {
		return C.longlong(signal.(func() QDesignerFormWindowInterface__FeatureFlag)())
	}

	return C.longlong(0)
}

func (ptr *QDesignerFormWindowInterface) ConnectFeatures(f func() QDesignerFormWindowInterface__FeatureFlag) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::features")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::features", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatures() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::features")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::features")
	}
}

func (ptr *QDesignerFormWindowInterface) Features() QDesignerFormWindowInterface__FeatureFlag {
	defer qt.Recovering("QDesignerFormWindowInterface::features")

	if ptr.Pointer() != nil {
		return QDesignerFormWindowInterface__FeatureFlag(C.QDesignerFormWindowInterface_Features(ptr.Pointer()))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_FileName
func callbackQDesignerFormWindowInterface_FileName(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::fileName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::fileName"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectFileName(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::fileName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::fileName", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileName() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::fileName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::fileName")
	}
}

func (ptr *QDesignerFormWindowInterface) FileName() string {
	defer qt.Recovering("QDesignerFormWindowInterface::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_FileName(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_FileNameChanged
func callbackQDesignerFormWindowInterface_FileNameChanged(ptr unsafe.Pointer, fileName *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::fileNameChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::fileNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(fileName))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectFileNameChanged(f func(fileName string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::fileNameChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectFileNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::fileNameChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileNameChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::fileNameChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectFileNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::fileNameChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) FileNameChanged(fileName string) {
	defer qt.Recovering("QDesignerFormWindowInterface::fileNameChanged")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QDesignerFormWindowInterface_FileNameChanged(ptr.Pointer(), fileNameC)
	}
}

//export callbackQDesignerFormWindowInterface_FormContainer
func callbackQDesignerFormWindowInterface_FormContainer(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::formContainer")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::formContainer"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func() *widgets.QWidget)())
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectFormContainer(f func() *widgets.QWidget) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::formContainer")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::formContainer", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFormContainer() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::formContainer")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::formContainer")
	}
}

func (ptr *QDesignerFormWindowInterface) FormContainer() *widgets.QWidget {
	defer qt.Recovering("QDesignerFormWindowInterface::formContainer")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowInterface_FormContainer(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_GeometryChanged
func callbackQDesignerFormWindowInterface_GeometryChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::geometryChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::geometryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectGeometryChanged(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::geometryChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::geometryChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::geometryChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::geometryChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) GeometryChanged() {
	defer qt.Recovering("QDesignerFormWindowInterface::geometryChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_GeometryChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_Grid
func callbackQDesignerFormWindowInterface_Grid(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::grid")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::grid"); signal != nil {
		return core.PointerFromQPoint(signal.(func() *core.QPoint)())
	}

	return core.PointerFromQPoint(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectGrid(f func() *core.QPoint) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::grid")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::grid", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectGrid() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::grid")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::grid")
	}
}

func (ptr *QDesignerFormWindowInterface) Grid() *core.QPoint {
	defer qt.Recovering("QDesignerFormWindowInterface::grid")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QDesignerFormWindowInterface_Grid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_HasFeature
func callbackQDesignerFormWindowInterface_HasFeature(ptr unsafe.Pointer, feature C.longlong) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::hasFeature")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QDesignerFormWindowInterface__FeatureFlag) bool)(QDesignerFormWindowInterface__FeatureFlag(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectHasFeature(f func(feature QDesignerFormWindowInterface__FeatureFlag) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::hasFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hasFeature", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHasFeature(feature QDesignerFormWindowInterface__FeatureFlag) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::hasFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hasFeature")
	}
}

func (ptr *QDesignerFormWindowInterface) HasFeature(feature QDesignerFormWindowInterface__FeatureFlag) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::hasFeature")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_HasFeature(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_IncludeHints
func callbackQDesignerFormWindowInterface_IncludeHints(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::includeHints")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::includeHints"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QDesignerFormWindowInterface) ConnectIncludeHints(f func() []string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::includeHints")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::includeHints", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIncludeHints() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::includeHints")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::includeHints")
	}
}

func (ptr *QDesignerFormWindowInterface) IncludeHints() []string {
	defer qt.Recovering("QDesignerFormWindowInterface::includeHints")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDesignerFormWindowInterface_IncludeHints(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_IsDirty
func callbackQDesignerFormWindowInterface_IsDirty(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::isDirty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::isDirty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectIsDirty(f func() bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::isDirty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::isDirty", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsDirty() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::isDirty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::isDirty")
	}
}

func (ptr *QDesignerFormWindowInterface) IsDirty() bool {
	defer qt.Recovering("QDesignerFormWindowInterface::isDirty")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_IsDirty(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_IsManaged
func callbackQDesignerFormWindowInterface_IsManaged(ptr unsafe.Pointer, widget unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::isManaged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::isManaged"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QWidget) bool)(widgets.NewQWidgetFromPointer(widget)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectIsManaged(f func(widget *widgets.QWidget) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::isManaged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::isManaged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsManaged(widget widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::isManaged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::isManaged")
	}
}

func (ptr *QDesignerFormWindowInterface) IsManaged(widget widgets.QWidget_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::isManaged")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_IsManaged(ptr.Pointer(), widgets.PointerFromQWidget(widget)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_LayoutDefault
func callbackQDesignerFormWindowInterface_LayoutDefault(ptr unsafe.Pointer, margin C.int, spacing C.int) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::layoutDefault")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::layoutDefault"); signal != nil {
		signal.(func(int, int))(int(int32(margin)), int(int32(spacing)))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutDefault(f func(margin int, spacing int)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::layoutDefault")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::layoutDefault", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutDefault(margin int, spacing int) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::layoutDefault")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::layoutDefault")
	}
}

func (ptr *QDesignerFormWindowInterface) LayoutDefault(margin int, spacing int) {
	defer qt.Recovering("QDesignerFormWindowInterface::layoutDefault")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LayoutDefault(ptr.Pointer(), C.int(int32(margin)), C.int(int32(spacing)))
	}
}

//export callbackQDesignerFormWindowInterface_LayoutFunction
func callbackQDesignerFormWindowInterface_LayoutFunction(ptr unsafe.Pointer, margin *C.char, spacing *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::layoutFunction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::layoutFunction"); signal != nil {
		signal.(func(string, string))(C.GoString(margin), C.GoString(spacing))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutFunction(f func(margin string, spacing string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::layoutFunction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::layoutFunction", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutFunction(margin string, spacing string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::layoutFunction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::layoutFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) LayoutFunction(margin string, spacing string) {
	defer qt.Recovering("QDesignerFormWindowInterface::layoutFunction")

	if ptr.Pointer() != nil {
		var marginC = C.CString(margin)
		defer C.free(unsafe.Pointer(marginC))
		var spacingC = C.CString(spacing)
		defer C.free(unsafe.Pointer(spacingC))
		C.QDesignerFormWindowInterface_LayoutFunction(ptr.Pointer(), marginC, spacingC)
	}
}

//export callbackQDesignerFormWindowInterface_MainContainerChanged
func callbackQDesignerFormWindowInterface_MainContainerChanged(ptr unsafe.Pointer, mainContainer unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::mainContainerChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::mainContainerChanged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(mainContainer))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectMainContainerChanged(f func(mainContainer *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::mainContainerChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectMainContainerChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mainContainerChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMainContainerChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::mainContainerChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectMainContainerChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mainContainerChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) MainContainerChanged(mainContainer widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mainContainerChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MainContainerChanged(ptr.Pointer(), widgets.PointerFromQWidget(mainContainer))
	}
}

//export callbackQDesignerFormWindowInterface_ManageWidget
func callbackQDesignerFormWindowInterface_ManageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::manageWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::manageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectManageWidget(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::manageWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::manageWidget", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectManageWidget(widget widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::manageWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::manageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) ManageWidget(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::manageWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ManageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_ObjectRemoved
func callbackQDesignerFormWindowInterface_ObjectRemoved(ptr unsafe.Pointer, object unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::objectRemoved")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::objectRemoved"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectObjectRemoved(f func(object *core.QObject)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::objectRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectObjectRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::objectRemoved", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectObjectRemoved() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::objectRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectObjectRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::objectRemoved")
	}
}

func (ptr *QDesignerFormWindowInterface) ObjectRemoved(object core.QObject_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::objectRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ObjectRemoved(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQDesignerFormWindowInterface_PixmapFunction
func callbackQDesignerFormWindowInterface_PixmapFunction(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::pixmapFunction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::pixmapFunction"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerFormWindowInterface) ConnectPixmapFunction(f func() string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::pixmapFunction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::pixmapFunction", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectPixmapFunction() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::pixmapFunction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::pixmapFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) PixmapFunction() string {
	defer qt.Recovering("QDesignerFormWindowInterface::pixmapFunction")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerFormWindowInterface_PixmapFunction(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_RemoveResourceFile
func callbackQDesignerFormWindowInterface_RemoveResourceFile(ptr unsafe.Pointer, path *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::removeResourceFile")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::removeResourceFile"); signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectRemoveResourceFile(f func(path string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::removeResourceFile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::removeResourceFile", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectRemoveResourceFile(path string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::removeResourceFile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::removeResourceFile")
	}
}

func (ptr *QDesignerFormWindowInterface) RemoveResourceFile(path string) {
	defer qt.Recovering("QDesignerFormWindowInterface::removeResourceFile")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QDesignerFormWindowInterface_RemoveResourceFile(ptr.Pointer(), pathC)
	}
}

//export callbackQDesignerFormWindowInterface_ResourceFileSaveMode
func callbackQDesignerFormWindowInterface_ResourceFileSaveMode(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QDesignerFormWindowInterface::resourceFileSaveMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::resourceFileSaveMode"); signal != nil {
		return C.longlong(signal.(func() QDesignerFormWindowInterface__ResourceFileSaveMode)())
	}

	return C.longlong(0)
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFileSaveMode(f func() QDesignerFormWindowInterface__ResourceFileSaveMode) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::resourceFileSaveMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFileSaveMode", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFileSaveMode() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::resourceFileSaveMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFileSaveMode")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFileSaveMode() QDesignerFormWindowInterface__ResourceFileSaveMode {
	defer qt.Recovering("QDesignerFormWindowInterface::resourceFileSaveMode")

	if ptr.Pointer() != nil {
		return QDesignerFormWindowInterface__ResourceFileSaveMode(C.QDesignerFormWindowInterface_ResourceFileSaveMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_ResourceFiles
func callbackQDesignerFormWindowInterface_ResourceFiles(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::resourceFiles")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::resourceFiles"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFiles(f func() []string) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::resourceFiles")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFiles", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFiles() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::resourceFiles")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFiles")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFiles() []string {
	defer qt.Recovering("QDesignerFormWindowInterface::resourceFiles")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QDesignerFormWindowInterface_ResourceFiles(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_ResourceFilesChanged
func callbackQDesignerFormWindowInterface_ResourceFilesChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::resourceFilesChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::resourceFilesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFilesChanged(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::resourceFilesChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectResourceFilesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFilesChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFilesChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::resourceFilesChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectResourceFilesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resourceFilesChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFilesChanged() {
	defer qt.Recovering("QDesignerFormWindowInterface::resourceFilesChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ResourceFilesChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SelectWidget
func callbackQDesignerFormWindowInterface_SelectWidget(ptr unsafe.Pointer, widget unsafe.Pointer, sele C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::selectWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::selectWidget"); signal != nil {
		signal.(func(*widgets.QWidget, bool))(widgets.NewQWidgetFromPointer(widget), int8(sele) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSelectWidget(f func(widget *widgets.QWidget, sele bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::selectWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::selectWidget", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectWidget(widget widgets.QWidget_ITF, sele bool) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::selectWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::selectWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) SelectWidget(widget widgets.QWidget_ITF, sele bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::selectWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SelectWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

//export callbackQDesignerFormWindowInterface_SelectionChanged
func callbackQDesignerFormWindowInterface_SelectionChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::selectionChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::selectionChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::selectionChanged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::selectionChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::selectionChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) SelectionChanged() {
	defer qt.Recovering("QDesignerFormWindowInterface::selectionChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SelectionChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SetAuthor
func callbackQDesignerFormWindowInterface_SetAuthor(ptr unsafe.Pointer, author *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setAuthor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setAuthor"); signal != nil {
		signal.(func(string))(C.GoString(author))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetAuthor(f func(author string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setAuthor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setAuthor", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetAuthor(author string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setAuthor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setAuthor")
	}
}

func (ptr *QDesignerFormWindowInterface) SetAuthor(author string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setAuthor")

	if ptr.Pointer() != nil {
		var authorC = C.CString(author)
		defer C.free(unsafe.Pointer(authorC))
		C.QDesignerFormWindowInterface_SetAuthor(ptr.Pointer(), authorC)
	}
}

//export callbackQDesignerFormWindowInterface_SetComment
func callbackQDesignerFormWindowInterface_SetComment(ptr unsafe.Pointer, comment *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setComment")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setComment"); signal != nil {
		signal.(func(string))(C.GoString(comment))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetComment(f func(comment string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setComment")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setComment", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetComment(comment string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setComment")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setComment")
	}
}

func (ptr *QDesignerFormWindowInterface) SetComment(comment string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setComment")

	if ptr.Pointer() != nil {
		var commentC = C.CString(comment)
		defer C.free(unsafe.Pointer(commentC))
		C.QDesignerFormWindowInterface_SetComment(ptr.Pointer(), commentC)
	}
}

//export callbackQDesignerFormWindowInterface_SetContents
func callbackQDesignerFormWindowInterface_SetContents(ptr unsafe.Pointer, device unsafe.Pointer, errorMessage *C.char) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setContents")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setContents"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QIODevice, string) bool)(core.NewQIODeviceFromPointer(device), C.GoString(errorMessage)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents(f func(device *core.QIODevice, errorMessage string) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setContents", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents(device core.QIODevice_ITF, errorMessage string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setContents")
	}
}

func (ptr *QDesignerFormWindowInterface) SetContents(device core.QIODevice_ITF, errorMessage string) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {
		var errorMessageC = C.CString(errorMessage)
		defer C.free(unsafe.Pointer(errorMessageC))
		return C.QDesignerFormWindowInterface_SetContents(ptr.Pointer(), core.PointerFromQIODevice(device), errorMessageC) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_SetContents2
func callbackQDesignerFormWindowInterface_SetContents2(ptr unsafe.Pointer, contents *C.char) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setContents")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setContents2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(contents)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents2(f func(contents string) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setContents2", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents2(contents string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setContents2")
	}
}

func (ptr *QDesignerFormWindowInterface) SetContents2(contents string) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::setContents")

	if ptr.Pointer() != nil {
		var contentsC = C.CString(contents)
		defer C.free(unsafe.Pointer(contentsC))
		return C.QDesignerFormWindowInterface_SetContents2(ptr.Pointer(), contentsC) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_SetDirty
func callbackQDesignerFormWindowInterface_SetDirty(ptr unsafe.Pointer, dirty C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setDirty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setDirty"); signal != nil {
		signal.(func(bool))(int8(dirty) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetDirty(f func(dirty bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setDirty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setDirty", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetDirty(dirty bool) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setDirty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setDirty")
	}
}

func (ptr *QDesignerFormWindowInterface) SetDirty(dirty bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setDirty")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetDirty(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(dirty))))
	}
}

//export callbackQDesignerFormWindowInterface_SetExportMacro
func callbackQDesignerFormWindowInterface_SetExportMacro(ptr unsafe.Pointer, exportMacro *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setExportMacro")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setExportMacro"); signal != nil {
		signal.(func(string))(C.GoString(exportMacro))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetExportMacro(f func(exportMacro string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setExportMacro")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setExportMacro", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetExportMacro(exportMacro string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setExportMacro")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setExportMacro")
	}
}

func (ptr *QDesignerFormWindowInterface) SetExportMacro(exportMacro string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setExportMacro")

	if ptr.Pointer() != nil {
		var exportMacroC = C.CString(exportMacro)
		defer C.free(unsafe.Pointer(exportMacroC))
		C.QDesignerFormWindowInterface_SetExportMacro(ptr.Pointer(), exportMacroC)
	}
}

//export callbackQDesignerFormWindowInterface_SetFeatures
func callbackQDesignerFormWindowInterface_SetFeatures(ptr unsafe.Pointer, features C.longlong) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setFeatures")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setFeatures"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__FeatureFlag))(QDesignerFormWindowInterface__FeatureFlag(features))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetFeatures(f func(features QDesignerFormWindowInterface__FeatureFlag)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setFeatures")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFeatures", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFeatures(features QDesignerFormWindowInterface__FeatureFlag) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setFeatures")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFeatures")
	}
}

func (ptr *QDesignerFormWindowInterface) SetFeatures(features QDesignerFormWindowInterface__FeatureFlag) {
	defer qt.Recovering("QDesignerFormWindowInterface::setFeatures")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetFeatures(ptr.Pointer(), C.longlong(features))
	}
}

//export callbackQDesignerFormWindowInterface_SetFileName
func callbackQDesignerFormWindowInterface_SetFileName(ptr unsafe.Pointer, fileName *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setFileName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setFileName"); signal != nil {
		signal.(func(string))(C.GoString(fileName))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetFileName(f func(fileName string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setFileName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFileName", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFileName(fileName string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setFileName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFileName")
	}
}

func (ptr *QDesignerFormWindowInterface) SetFileName(fileName string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setFileName")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QDesignerFormWindowInterface_SetFileName(ptr.Pointer(), fileNameC)
	}
}

//export callbackQDesignerFormWindowInterface_SetGrid
func callbackQDesignerFormWindowInterface_SetGrid(ptr unsafe.Pointer, grid unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setGrid")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setGrid"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(grid))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetGrid(f func(grid *core.QPoint)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setGrid")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setGrid", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetGrid(grid core.QPoint_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setGrid")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setGrid")
	}
}

func (ptr *QDesignerFormWindowInterface) SetGrid(grid core.QPoint_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::setGrid")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetGrid(ptr.Pointer(), core.PointerFromQPoint(grid))
	}
}

//export callbackQDesignerFormWindowInterface_SetIncludeHints
func callbackQDesignerFormWindowInterface_SetIncludeHints(ptr unsafe.Pointer, includeHints *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setIncludeHints")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setIncludeHints"); signal != nil {
		signal.(func([]string))(strings.Split(C.GoString(includeHints), "|"))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetIncludeHints(f func(includeHints []string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setIncludeHints")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setIncludeHints", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetIncludeHints(includeHints []string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setIncludeHints")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setIncludeHints")
	}
}

func (ptr *QDesignerFormWindowInterface) SetIncludeHints(includeHints []string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setIncludeHints")

	if ptr.Pointer() != nil {
		var includeHintsC = C.CString(strings.Join(includeHints, "|"))
		defer C.free(unsafe.Pointer(includeHintsC))
		C.QDesignerFormWindowInterface_SetIncludeHints(ptr.Pointer(), includeHintsC)
	}
}

//export callbackQDesignerFormWindowInterface_SetLayoutDefault
func callbackQDesignerFormWindowInterface_SetLayoutDefault(ptr unsafe.Pointer, margin C.int, spacing C.int) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setLayoutDefault")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setLayoutDefault"); signal != nil {
		signal.(func(int, int))(int(int32(margin)), int(int32(spacing)))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutDefault(f func(margin int, spacing int)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setLayoutDefault")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setLayoutDefault", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutDefault(margin int, spacing int) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setLayoutDefault")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setLayoutDefault")
	}
}

func (ptr *QDesignerFormWindowInterface) SetLayoutDefault(margin int, spacing int) {
	defer qt.Recovering("QDesignerFormWindowInterface::setLayoutDefault")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetLayoutDefault(ptr.Pointer(), C.int(int32(margin)), C.int(int32(spacing)))
	}
}

//export callbackQDesignerFormWindowInterface_SetLayoutFunction
func callbackQDesignerFormWindowInterface_SetLayoutFunction(ptr unsafe.Pointer, margin *C.char, spacing *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setLayoutFunction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setLayoutFunction"); signal != nil {
		signal.(func(string, string))(C.GoString(margin), C.GoString(spacing))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutFunction(f func(margin string, spacing string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setLayoutFunction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setLayoutFunction", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutFunction(margin string, spacing string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setLayoutFunction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setLayoutFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) SetLayoutFunction(margin string, spacing string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setLayoutFunction")

	if ptr.Pointer() != nil {
		var marginC = C.CString(margin)
		defer C.free(unsafe.Pointer(marginC))
		var spacingC = C.CString(spacing)
		defer C.free(unsafe.Pointer(spacingC))
		C.QDesignerFormWindowInterface_SetLayoutFunction(ptr.Pointer(), marginC, spacingC)
	}
}

//export callbackQDesignerFormWindowInterface_SetMainContainer
func callbackQDesignerFormWindowInterface_SetMainContainer(ptr unsafe.Pointer, mainContainer unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setMainContainer")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setMainContainer"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(mainContainer))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetMainContainer(f func(mainContainer *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setMainContainer")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setMainContainer", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetMainContainer(mainContainer widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setMainContainer")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setMainContainer")
	}
}

func (ptr *QDesignerFormWindowInterface) SetMainContainer(mainContainer widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::setMainContainer")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetMainContainer(ptr.Pointer(), widgets.PointerFromQWidget(mainContainer))
	}
}

//export callbackQDesignerFormWindowInterface_SetPixmapFunction
func callbackQDesignerFormWindowInterface_SetPixmapFunction(ptr unsafe.Pointer, pixmapFunction *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setPixmapFunction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setPixmapFunction"); signal != nil {
		signal.(func(string))(C.GoString(pixmapFunction))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetPixmapFunction(f func(pixmapFunction string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setPixmapFunction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setPixmapFunction", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetPixmapFunction(pixmapFunction string) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setPixmapFunction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setPixmapFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) SetPixmapFunction(pixmapFunction string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setPixmapFunction")

	if ptr.Pointer() != nil {
		var pixmapFunctionC = C.CString(pixmapFunction)
		defer C.free(unsafe.Pointer(pixmapFunctionC))
		C.QDesignerFormWindowInterface_SetPixmapFunction(ptr.Pointer(), pixmapFunctionC)
	}
}

//export callbackQDesignerFormWindowInterface_SetResourceFileSaveMode
func callbackQDesignerFormWindowInterface_SetResourceFileSaveMode(ptr unsafe.Pointer, behaviour C.longlong) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setResourceFileSaveMode")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setResourceFileSaveMode"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__ResourceFileSaveMode))(QDesignerFormWindowInterface__ResourceFileSaveMode(behaviour))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetResourceFileSaveMode(f func(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setResourceFileSaveMode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setResourceFileSaveMode", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetResourceFileSaveMode(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setResourceFileSaveMode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setResourceFileSaveMode")
	}
}

func (ptr *QDesignerFormWindowInterface) SetResourceFileSaveMode(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode) {
	defer qt.Recovering("QDesignerFormWindowInterface::setResourceFileSaveMode")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetResourceFileSaveMode(ptr.Pointer(), C.longlong(behaviour))
	}
}

//export callbackQDesignerFormWindowInterface_UnmanageWidget
func callbackQDesignerFormWindowInterface_UnmanageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::unmanageWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::unmanageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectUnmanageWidget(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::unmanageWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::unmanageWidget", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectUnmanageWidget(widget widgets.QWidget_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::unmanageWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::unmanageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) UnmanageWidget(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::unmanageWidget")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UnmanageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetManaged
func callbackQDesignerFormWindowInterface_WidgetManaged(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::widgetManaged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::widgetManaged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetManaged(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::widgetManaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectWidgetManaged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetManaged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetManaged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::widgetManaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetManaged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetManaged")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetManaged(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::widgetManaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetManaged(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetRemoved
func callbackQDesignerFormWindowInterface_WidgetRemoved(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::widgetRemoved")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::widgetRemoved"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetRemoved(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetRemoved", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetRemoved() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetRemoved")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetRemoved(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetRemoved(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetUnmanaged
func callbackQDesignerFormWindowInterface_WidgetUnmanaged(ptr unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::widgetUnmanaged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::widgetUnmanaged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetUnmanaged(f func(widget *widgets.QWidget)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::widgetUnmanaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectWidgetUnmanaged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetUnmanaged", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetUnmanaged() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::widgetUnmanaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetUnmanaged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::widgetUnmanaged")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetUnmanaged(widget widgets.QWidget_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::widgetUnmanaged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetUnmanaged(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QDesignerFormWindowInterface) DestroyQDesignerFormWindowInterface() {
	defer qt.Recovering("QDesignerFormWindowInterface::~QDesignerFormWindowInterface")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormWindowInterface_ActionEvent
func callbackQDesignerFormWindowInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::actionEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::actionEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DragEnterEvent
func callbackQDesignerFormWindowInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragEnterEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragEnterEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DragLeaveEvent
func callbackQDesignerFormWindowInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragLeaveEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragLeaveEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DragMoveEvent
func callbackQDesignerFormWindowInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragMoveEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dragMoveEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DropEvent
func callbackQDesignerFormWindowInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dropEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::dropEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_EnterEvent
func callbackQDesignerFormWindowInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::enterEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::enterEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_FocusInEvent
func callbackQDesignerFormWindowInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusInEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusInEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_FocusOutEvent
func callbackQDesignerFormWindowInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusOutEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusOutEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_HideEvent
func callbackQDesignerFormWindowInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hideEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hideEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_LeaveEvent
func callbackQDesignerFormWindowInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::leaveEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::leaveEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MinimumSizeHint
func callbackQDesignerFormWindowInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerFormWindowInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerFormWindowInterface) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::minimumSizeHint", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::minimumSizeHint")
	}
}

func (ptr *QDesignerFormWindowInterface) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDesignerFormWindowInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerFormWindowInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_MoveEvent
func callbackQDesignerFormWindowInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::moveEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::moveEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_PaintEvent
func callbackQDesignerFormWindowInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::paintEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::paintEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_SetEnabled
func callbackQDesignerFormWindowInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setEnabled", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setEnabled")
	}
}

func (ptr *QDesignerFormWindowInterface) SetEnabled(vbo bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerFormWindowInterface) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerFormWindowInterface_SetStyleSheet
func callbackQDesignerFormWindowInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setStyleSheet", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setStyleSheet")
	}
}

func (ptr *QDesignerFormWindowInterface) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerFormWindowInterface_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QDesignerFormWindowInterface) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerFormWindowInterface_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQDesignerFormWindowInterface_SetVisible
func callbackQDesignerFormWindowInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setVisible", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setVisible")
	}
}

func (ptr *QDesignerFormWindowInterface) SetVisible(visible bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerFormWindowInterface) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerFormWindowInterface_SetWindowModified
func callbackQDesignerFormWindowInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setWindowModified", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setWindowModified")
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowModified(vbo bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerFormWindowInterface_SetWindowTitle
func callbackQDesignerFormWindowInterface_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setWindowTitle", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setWindowTitle")
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowTitle(vqs string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerFormWindowInterface_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QDesignerFormWindowInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerFormWindowInterface_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQDesignerFormWindowInterface_ShowEvent
func callbackQDesignerFormWindowInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_SizeHint
func callbackQDesignerFormWindowInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerFormWindowInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerFormWindowInterface) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::sizeHint", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::sizeHint")
	}
}

func (ptr *QDesignerFormWindowInterface) SizeHint() *core.QSize {
	defer qt.Recovering("QDesignerFormWindowInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerFormWindowInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_ChangeEvent
func callbackQDesignerFormWindowInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::changeEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::changeEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Close
func callbackQDesignerFormWindowInterface_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerFormWindowInterface) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::close", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectClose() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::close")
	}
}

func (ptr *QDesignerFormWindowInterface) Close() bool {
	defer qt.Recovering("QDesignerFormWindowInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) CloseDefault() bool {
	defer qt.Recovering("QDesignerFormWindowInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_CloseEvent
func callbackQDesignerFormWindowInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::closeEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::closeEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ContextMenuEvent
func callbackQDesignerFormWindowInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::contextMenuEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::contextMenuEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_FocusNextPrevChild
func callbackQDesignerFormWindowInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerFormWindowInterface) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusNextPrevChild", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::focusNextPrevChild")
	}
}

func (ptr *QDesignerFormWindowInterface) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_HasHeightForWidth
func callbackQDesignerFormWindowInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerFormWindowInterface) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hasHeightForWidth", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hasHeightForWidth")
	}
}

func (ptr *QDesignerFormWindowInterface) HasHeightForWidth() bool {
	defer qt.Recovering("QDesignerFormWindowInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QDesignerFormWindowInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_HeightForWidth
func callbackQDesignerFormWindowInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QDesignerFormWindowInterface::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerFormWindowInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerFormWindowInterface) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::heightForWidth", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::heightForWidth")
	}
}

func (ptr *QDesignerFormWindowInterface) HeightForWidth(w int) int {
	defer qt.Recovering("QDesignerFormWindowInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowInterface_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QDesignerFormWindowInterface) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QDesignerFormWindowInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_Hide
func callbackQDesignerFormWindowInterface_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectHide(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hide", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHide() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::hide")
	}
}

func (ptr *QDesignerFormWindowInterface) Hide() {
	defer qt.Recovering("QDesignerFormWindowInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Hide(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) HideDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_InputMethodEvent
func callbackQDesignerFormWindowInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::inputMethodEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::inputMethodEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_InputMethodQuery
func callbackQDesignerFormWindowInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerFormWindowInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerFormWindowInterface) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::inputMethodQuery", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::inputMethodQuery")
	}
}

func (ptr *QDesignerFormWindowInterface) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerFormWindowInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerFormWindowInterface_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerFormWindowInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerFormWindowInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_KeyPressEvent
func callbackQDesignerFormWindowInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::keyPressEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::keyPressEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_KeyReleaseEvent
func callbackQDesignerFormWindowInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::keyReleaseEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::keyReleaseEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Lower
func callbackQDesignerFormWindowInterface_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectLower(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::lower", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLower() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::lower")
	}
}

func (ptr *QDesignerFormWindowInterface) Lower() {
	defer qt.Recovering("QDesignerFormWindowInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Lower(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) LowerDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_MouseDoubleClickEvent
func callbackQDesignerFormWindowInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseDoubleClickEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MouseMoveEvent
func callbackQDesignerFormWindowInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseMoveEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseMoveEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MousePressEvent
func callbackQDesignerFormWindowInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mousePressEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mousePressEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MouseReleaseEvent
func callbackQDesignerFormWindowInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseReleaseEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::mouseReleaseEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_NativeEvent
func callbackQDesignerFormWindowInterface_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QDesignerFormWindowInterface) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::nativeEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::nativeEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerFormWindowInterface_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerFormWindowInterface_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_Raise
func callbackQDesignerFormWindowInterface_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectRaise(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::raise", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectRaise() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::raise")
	}
}

func (ptr *QDesignerFormWindowInterface) Raise() {
	defer qt.Recovering("QDesignerFormWindowInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Raise(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) RaiseDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_Repaint
func callbackQDesignerFormWindowInterface_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::repaint", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectRepaint() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::repaint")
	}
}

func (ptr *QDesignerFormWindowInterface) Repaint() {
	defer qt.Recovering("QDesignerFormWindowInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Repaint(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) RepaintDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ResizeEvent
func callbackQDesignerFormWindowInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resizeEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::resizeEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_SetDisabled
func callbackQDesignerFormWindowInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setDisabled", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setDisabled")
	}
}

func (ptr *QDesignerFormWindowInterface) SetDisabled(disable bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QDesignerFormWindowInterface) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerFormWindowInterface_SetFocus2
func callbackQDesignerFormWindowInterface_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFocus2", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setFocus2")
	}
}

func (ptr *QDesignerFormWindowInterface) SetFocus2() {
	defer qt.Recovering("QDesignerFormWindowInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) SetFocus2Default() {
	defer qt.Recovering("QDesignerFormWindowInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SetHidden
func callbackQDesignerFormWindowInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setHidden", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::setHidden")
	}
}

func (ptr *QDesignerFormWindowInterface) SetHidden(hidden bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QDesignerFormWindowInterface) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QDesignerFormWindowInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerFormWindowInterface_Show
func callbackQDesignerFormWindowInterface_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShow(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::show", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShow() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::show")
	}
}

func (ptr *QDesignerFormWindowInterface) Show() {
	defer qt.Recovering("QDesignerFormWindowInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Show(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) ShowDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowFullScreen
func callbackQDesignerFormWindowInterface_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showFullScreen", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showFullScreen")
	}
}

func (ptr *QDesignerFormWindowInterface) ShowFullScreen() {
	defer qt.Recovering("QDesignerFormWindowInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) ShowFullScreenDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowMaximized
func callbackQDesignerFormWindowInterface_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showMaximized", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showMaximized")
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMaximized() {
	defer qt.Recovering("QDesignerFormWindowInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMaximizedDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowMinimized
func callbackQDesignerFormWindowInterface_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showMinimized", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showMinimized")
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMinimized() {
	defer qt.Recovering("QDesignerFormWindowInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMinimizedDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowNormal
func callbackQDesignerFormWindowInterface_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showNormal", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::showNormal")
	}
}

func (ptr *QDesignerFormWindowInterface) ShowNormal() {
	defer qt.Recovering("QDesignerFormWindowInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) ShowNormalDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_TabletEvent
func callbackQDesignerFormWindowInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::tabletEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::tabletEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Update
func callbackQDesignerFormWindowInterface_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::update", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectUpdate() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::update")
	}
}

func (ptr *QDesignerFormWindowInterface) Update() {
	defer qt.Recovering("QDesignerFormWindowInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Update(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) UpdateDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_UpdateMicroFocus
func callbackQDesignerFormWindowInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::updateMicroFocus", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::updateMicroFocus")
	}
}

func (ptr *QDesignerFormWindowInterface) UpdateMicroFocus() {
	defer qt.Recovering("QDesignerFormWindowInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowInterface) UpdateMicroFocusDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_WheelEvent
func callbackQDesignerFormWindowInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::wheelEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::wheelEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_TimerEvent
func callbackQDesignerFormWindowInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::timerEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::timerEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ChildEvent
func callbackQDesignerFormWindowInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::childEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::childEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ConnectNotify
func callbackQDesignerFormWindowInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::connectNotify", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::connectNotify")
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowInterface_CustomEvent
func callbackQDesignerFormWindowInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::customEvent", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::customEvent")
	}
}

func (ptr *QDesignerFormWindowInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormWindowInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DeleteLater
func callbackQDesignerFormWindowInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::deleteLater", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::deleteLater")
	}
}

func (ptr *QDesignerFormWindowInterface) DeleteLater() {
	defer qt.Recovering("QDesignerFormWindowInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerFormWindowInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerFormWindowInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormWindowInterface_DisconnectNotify
func callbackQDesignerFormWindowInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::disconnectNotify")
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowInterface_EventFilter
func callbackQDesignerFormWindowInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormWindowInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::eventFilter", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::eventFilter")
	}
}

func (ptr *QDesignerFormWindowInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_MetaObject
func callbackQDesignerFormWindowInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormWindowInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormWindowInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerFormWindowInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::metaObject", f)
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerFormWindowInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowInterface::metaObject")
	}
}

func (ptr *QDesignerFormWindowInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormWindowInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormWindowInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormWindowInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormWindowInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QDesignerFormWindowManagerInterface::Action
type QDesignerFormWindowManagerInterface__Action int64

const (
	QDesignerFormWindowManagerInterface__CutAction                      = QDesignerFormWindowManagerInterface__Action(100)
	QDesignerFormWindowManagerInterface__CopyAction                     = QDesignerFormWindowManagerInterface__Action(101)
	QDesignerFormWindowManagerInterface__PasteAction                    = QDesignerFormWindowManagerInterface__Action(102)
	QDesignerFormWindowManagerInterface__DeleteAction                   = QDesignerFormWindowManagerInterface__Action(103)
	QDesignerFormWindowManagerInterface__SelectAllAction                = QDesignerFormWindowManagerInterface__Action(104)
	QDesignerFormWindowManagerInterface__LowerAction                    = QDesignerFormWindowManagerInterface__Action(200)
	QDesignerFormWindowManagerInterface__RaiseAction                    = QDesignerFormWindowManagerInterface__Action(201)
	QDesignerFormWindowManagerInterface__UndoAction                     = QDesignerFormWindowManagerInterface__Action(300)
	QDesignerFormWindowManagerInterface__RedoAction                     = QDesignerFormWindowManagerInterface__Action(301)
	QDesignerFormWindowManagerInterface__HorizontalLayoutAction         = QDesignerFormWindowManagerInterface__Action(400)
	QDesignerFormWindowManagerInterface__VerticalLayoutAction           = QDesignerFormWindowManagerInterface__Action(401)
	QDesignerFormWindowManagerInterface__SplitHorizontalAction          = QDesignerFormWindowManagerInterface__Action(402)
	QDesignerFormWindowManagerInterface__SplitVerticalAction            = QDesignerFormWindowManagerInterface__Action(403)
	QDesignerFormWindowManagerInterface__GridLayoutAction               = QDesignerFormWindowManagerInterface__Action(404)
	QDesignerFormWindowManagerInterface__FormLayoutAction               = QDesignerFormWindowManagerInterface__Action(405)
	QDesignerFormWindowManagerInterface__BreakLayoutAction              = QDesignerFormWindowManagerInterface__Action(406)
	QDesignerFormWindowManagerInterface__AdjustSizeAction               = QDesignerFormWindowManagerInterface__Action(407)
	QDesignerFormWindowManagerInterface__SimplifyLayoutAction           = QDesignerFormWindowManagerInterface__Action(408)
	QDesignerFormWindowManagerInterface__DefaultPreviewAction           = QDesignerFormWindowManagerInterface__Action(500)
	QDesignerFormWindowManagerInterface__FormWindowSettingsDialogAction = QDesignerFormWindowManagerInterface__Action(600)
)

//QDesignerFormWindowManagerInterface::ActionGroup
type QDesignerFormWindowManagerInterface__ActionGroup int64

const (
	QDesignerFormWindowManagerInterface__StyledPreviewActionGroup = QDesignerFormWindowManagerInterface__ActionGroup(100)
)

type QDesignerFormWindowManagerInterface struct {
	core.QObject
}

type QDesignerFormWindowManagerInterface_ITF interface {
	core.QObject_ITF
	QDesignerFormWindowManagerInterface_PTR() *QDesignerFormWindowManagerInterface
}

func (p *QDesignerFormWindowManagerInterface) QDesignerFormWindowManagerInterface_PTR() *QDesignerFormWindowManagerInterface {
	return p
}

func (p *QDesignerFormWindowManagerInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerFormWindowManagerInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerFormWindowManagerInterface(ptr QDesignerFormWindowManagerInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerFormWindowManagerInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerFormWindowManagerInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormWindowManagerInterface {
	var n = new(QDesignerFormWindowManagerInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerFormWindowManagerInterface_Action
func callbackQDesignerFormWindowManagerInterface_Action(ptr unsafe.Pointer, action C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::action")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::action"); signal != nil {
		return widgets.PointerFromQAction(signal.(func(QDesignerFormWindowManagerInterface__Action) *widgets.QAction)(QDesignerFormWindowManagerInterface__Action(action)))
	}

	return widgets.PointerFromQAction(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectAction(f func(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::action")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::action", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAction(action QDesignerFormWindowManagerInterface__Action) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::action")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::action")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) Action(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::action")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerFormWindowManagerInterface_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActionGroup
func callbackQDesignerFormWindowManagerInterface_ActionGroup(ptr unsafe.Pointer, actionGroup C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::actionGroup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::actionGroup"); signal != nil {
		return widgets.PointerFromQActionGroup(signal.(func(QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup)(QDesignerFormWindowManagerInterface__ActionGroup(actionGroup)))
	}

	return widgets.PointerFromQActionGroup(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActionGroup(f func(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::actionGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::actionGroup", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActionGroup(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::actionGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::actionGroup")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActionGroup(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::actionGroup")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionGroupFromPointer(C.QDesignerFormWindowManagerInterface_ActionGroup(ptr.Pointer(), C.longlong(actionGroup)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActiveFormWindow
func callbackQDesignerFormWindowManagerInterface_ActiveFormWindow(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::activeFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::activeFormWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func() *QDesignerFormWindowInterface)())
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindow(f func() *QDesignerFormWindowInterface) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::activeFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::activeFormWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindow() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::activeFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::activeFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindow() *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::activeFormWindow")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_ActiveFormWindow(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged
func callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::activeFormWindowChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::activeFormWindowChanged"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindowChanged(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::activeFormWindowChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectActiveFormWindowChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::activeFormWindowChanged", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindowChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::activeFormWindowChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectActiveFormWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::activeFormWindowChanged")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindowChanged(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::activeFormWindowChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ActiveFormWindowChanged(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_AddFormWindow
func callbackQDesignerFormWindowManagerInterface_AddFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::addFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::addFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectAddFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::addFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::addFormWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAddFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::addFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::addFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) AddFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::addFormWindow")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_AddFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_CloseAllPreviews
func callbackQDesignerFormWindowManagerInterface_CloseAllPreviews(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::closeAllPreviews")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::closeAllPreviews"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCloseAllPreviews(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::closeAllPreviews")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::closeAllPreviews", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCloseAllPreviews() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::closeAllPreviews")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::closeAllPreviews")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CloseAllPreviews() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::closeAllPreviews")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_CloseAllPreviews(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowManagerInterface_Core
func callbackQDesignerFormWindowManagerInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::core")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::core")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::core", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCore() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::core")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::core")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) Core() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_CreateFormWindow
func callbackQDesignerFormWindowManagerInterface_CreateFormWindow(ptr unsafe.Pointer, parent unsafe.Pointer, flags C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::createFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::createFormWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func(*widgets.QWidget, core.Qt__WindowType) *QDesignerFormWindowInterface)(widgets.NewQWidgetFromPointer(parent), core.Qt__WindowType(flags)))
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreateFormWindow(f func(parent *widgets.QWidget, flags core.Qt__WindowType) *QDesignerFormWindowInterface) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::createFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::createFormWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreateFormWindow(parent widgets.QWidget_ITF, flags core.Qt__WindowType) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::createFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::createFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CreateFormWindow(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::createFormWindow")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_CreateFormWindow(ptr.Pointer(), widgets.PointerFromQWidget(parent), C.longlong(flags)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap
func callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::createPreviewPixmap")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::createPreviewPixmap"); signal != nil {
		return gui.PointerFromQPixmap(signal.(func() *gui.QPixmap)())
	}

	return gui.PointerFromQPixmap(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreatePreviewPixmap(f func() *gui.QPixmap) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::createPreviewPixmap")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::createPreviewPixmap", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreatePreviewPixmap() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::createPreviewPixmap")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::createPreviewPixmap")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CreatePreviewPixmap() *gui.QPixmap {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::createPreviewPixmap")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPixmapFromPointer(C.QDesignerFormWindowManagerInterface_CreatePreviewPixmap(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_FormWindow
func callbackQDesignerFormWindowManagerInterface_FormWindow(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::formWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::formWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func(int) *QDesignerFormWindowInterface)(int(int32(index))))
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindow(f func(index int) *QDesignerFormWindowInterface) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::formWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindow(index int) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::formWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindow(index int) *QDesignerFormWindowInterface {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::formWindow")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_FormWindow(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowAdded
func callbackQDesignerFormWindowManagerInterface_FormWindowAdded(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::formWindowAdded")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::formWindowAdded"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowAdded(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::formWindowAdded")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectFormWindowAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowAdded", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowAdded() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::formWindowAdded")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowAdded")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowAdded(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::formWindowAdded")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowAdded(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowCount
func callbackQDesignerFormWindowManagerInterface_FormWindowCount(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::formWindowCount")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::formWindowCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowCount(f func() int) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::formWindowCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowCount", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowCount() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::formWindowCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowCount")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowCount() int {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::formWindowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowManagerInterface_FormWindowCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowRemoved
func callbackQDesignerFormWindowManagerInterface_FormWindowRemoved(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::formWindowRemoved")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::formWindowRemoved"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowRemoved(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::formWindowRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectFormWindowRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowRemoved", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowRemoved() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::formWindowRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowRemoved")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowRemoved(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::formWindowRemoved")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowRemoved(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged
func callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::formWindowSettingsChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::formWindowSettingsChanged"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowSettingsChanged(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::formWindowSettingsChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectFormWindowSettingsChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowSettingsChanged", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowSettingsChanged() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::formWindowSettingsChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowSettingsChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::formWindowSettingsChanged")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowSettingsChanged(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::formWindowSettingsChanged")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowSettingsChanged(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_RemoveFormWindow
func callbackQDesignerFormWindowManagerInterface_RemoveFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::removeFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::removeFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectRemoveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::removeFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::removeFormWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectRemoveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::removeFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::removeFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) RemoveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::removeFormWindow")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_RemoveFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow
func callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::setActiveFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::setActiveFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectSetActiveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::setActiveFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::setActiveFormWindow", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectSetActiveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::setActiveFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::setActiveFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) SetActiveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::setActiveFormWindow")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_SetActiveFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ShowPluginDialog
func callbackQDesignerFormWindowManagerInterface_ShowPluginDialog(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::showPluginDialog")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::showPluginDialog"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPluginDialog(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::showPluginDialog")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::showPluginDialog", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPluginDialog() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::showPluginDialog")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::showPluginDialog")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPluginDialog() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::showPluginDialog")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ShowPluginDialog(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowManagerInterface_ShowPreview
func callbackQDesignerFormWindowManagerInterface_ShowPreview(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::showPreview")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::showPreview"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPreview(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::showPreview")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::showPreview", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPreview() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::showPreview")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::showPreview")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPreview() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::showPreview")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ShowPreview(ptr.Pointer())
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DestroyQDesignerFormWindowManagerInterface() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::~QDesignerFormWindowManagerInterface")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormWindowManagerInterface_TimerEvent
func callbackQDesignerFormWindowManagerInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::timerEvent", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::timerEvent")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ChildEvent
func callbackQDesignerFormWindowManagerInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::childEvent", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::childEvent")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ConnectNotify
func callbackQDesignerFormWindowManagerInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::connectNotify", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::connectNotify")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowManagerInterface_CustomEvent
func callbackQDesignerFormWindowManagerInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::customEvent", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::customEvent")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_DeleteLater
func callbackQDesignerFormWindowManagerInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::deleteLater", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::deleteLater")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DeleteLater() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormWindowManagerInterface_DisconnectNotify
func callbackQDesignerFormWindowManagerInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::disconnectNotify")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowManagerInterface_Event
func callbackQDesignerFormWindowManagerInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::event", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::event")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::event")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowManagerInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::event")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowManagerInterface_EventFilter
func callbackQDesignerFormWindowManagerInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::eventFilter", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::eventFilter")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowManagerInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowManagerInterface_MetaObject
func callbackQDesignerFormWindowManagerInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerFormWindowManagerInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerFormWindowManagerInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerFormWindowManagerInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::metaObject", f)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerFormWindowManagerInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerFormWindowManagerInterface::metaObject")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormWindowManagerInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerFormWindowManagerInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormWindowManagerInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDesignerMemberSheetExtension struct {
	ptr unsafe.Pointer
}

type QDesignerMemberSheetExtension_ITF interface {
	QDesignerMemberSheetExtension_PTR() *QDesignerMemberSheetExtension
}

func (p *QDesignerMemberSheetExtension) QDesignerMemberSheetExtension_PTR() *QDesignerMemberSheetExtension {
	return p
}

func (p *QDesignerMemberSheetExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerMemberSheetExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerMemberSheetExtension(ptr QDesignerMemberSheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerMemberSheetExtension_PTR().Pointer()
	}
	return nil
}

func NewQDesignerMemberSheetExtensionFromPointer(ptr unsafe.Pointer) *QDesignerMemberSheetExtension {
	var n = new(QDesignerMemberSheetExtension)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerMemberSheetExtension_Count
func callbackQDesignerMemberSheetExtension_Count(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::count")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerMemberSheetExtension) ConnectCount(f func() int) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::count")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::count", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectCount() {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::count")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::count")
	}
}

func (ptr *QDesignerMemberSheetExtension) Count() int {
	defer qt.Recovering("QDesignerMemberSheetExtension::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerMemberSheetExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerMemberSheetExtension_DeclaredInClass
func callbackQDesignerMemberSheetExtension_DeclaredInClass(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::declaredInClass")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::declaredInClass"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerMemberSheetExtension) ConnectDeclaredInClass(f func(index int) string) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::declaredInClass")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::declaredInClass", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectDeclaredInClass(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::declaredInClass")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::declaredInClass")
	}
}

func (ptr *QDesignerMemberSheetExtension) DeclaredInClass(index int) string {
	defer qt.Recovering("QDesignerMemberSheetExtension::declaredInClass")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerMemberSheetExtension_DeclaredInClass(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_IndexOf
func callbackQDesignerMemberSheetExtension_IndexOf(ptr unsafe.Pointer, name *C.char) C.int {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::indexOf")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::indexOf"); signal != nil {
		return C.int(int32(signal.(func(string) int)(C.GoString(name))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIndexOf(f func(name string) int) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::indexOf")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::indexOf", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIndexOf(name string) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::indexOf")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::indexOf")
	}
}

func (ptr *QDesignerMemberSheetExtension) IndexOf(name string) int {
	defer qt.Recovering("QDesignerMemberSheetExtension::indexOf")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return int(int32(C.QDesignerMemberSheetExtension_IndexOf(ptr.Pointer(), nameC)))
	}
	return 0
}

//export callbackQDesignerMemberSheetExtension_InheritedFromWidget
func callbackQDesignerMemberSheetExtension_InheritedFromWidget(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::inheritedFromWidget")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::inheritedFromWidget"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectInheritedFromWidget(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::inheritedFromWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::inheritedFromWidget", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectInheritedFromWidget(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::inheritedFromWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::inheritedFromWidget")
	}
}

func (ptr *QDesignerMemberSheetExtension) InheritedFromWidget(index int) bool {
	defer qt.Recovering("QDesignerMemberSheetExtension::inheritedFromWidget")

	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_InheritedFromWidget(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsSignal
func callbackQDesignerMemberSheetExtension_IsSignal(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::isSignal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::isSignal"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSignal(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::isSignal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isSignal", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSignal(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::isSignal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isSignal")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsSignal(index int) bool {
	defer qt.Recovering("QDesignerMemberSheetExtension::isSignal")

	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsSignal(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsSlot
func callbackQDesignerMemberSheetExtension_IsSlot(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::isSlot")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::isSlot"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSlot(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::isSlot")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isSlot", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSlot(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::isSlot")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isSlot")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsSlot(index int) bool {
	defer qt.Recovering("QDesignerMemberSheetExtension::isSlot")

	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsSlot(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsVisible
func callbackQDesignerMemberSheetExtension_IsVisible(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::isVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::isVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsVisible(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::isVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isVisible", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsVisible(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::isVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::isVisible")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsVisible(index int) bool {
	defer qt.Recovering("QDesignerMemberSheetExtension::isVisible")

	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsVisible(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_MemberGroup
func callbackQDesignerMemberSheetExtension_MemberGroup(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::memberGroup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::memberGroup"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberGroup(f func(index int) string) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::memberGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::memberGroup", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberGroup(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::memberGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::memberGroup")
	}
}

func (ptr *QDesignerMemberSheetExtension) MemberGroup(index int) string {
	defer qt.Recovering("QDesignerMemberSheetExtension::memberGroup")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerMemberSheetExtension_MemberGroup(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_MemberName
func callbackQDesignerMemberSheetExtension_MemberName(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::memberName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::memberName"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberName(f func(index int) string) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::memberName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::memberName", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberName(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::memberName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::memberName")
	}
}

func (ptr *QDesignerMemberSheetExtension) MemberName(index int) string {
	defer qt.Recovering("QDesignerMemberSheetExtension::memberName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerMemberSheetExtension_MemberName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_SetMemberGroup
func callbackQDesignerMemberSheetExtension_SetMemberGroup(ptr unsafe.Pointer, index C.int, group *C.char) {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::setMemberGroup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::setMemberGroup"); signal != nil {
		signal.(func(int, string))(int(int32(index)), C.GoString(group))
	}

}

func (ptr *QDesignerMemberSheetExtension) ConnectSetMemberGroup(f func(index int, group string)) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::setMemberGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::setMemberGroup", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetMemberGroup(index int, group string) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::setMemberGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::setMemberGroup")
	}
}

func (ptr *QDesignerMemberSheetExtension) SetMemberGroup(index int, group string) {
	defer qt.Recovering("QDesignerMemberSheetExtension::setMemberGroup")

	if ptr.Pointer() != nil {
		var groupC = C.CString(group)
		defer C.free(unsafe.Pointer(groupC))
		C.QDesignerMemberSheetExtension_SetMemberGroup(ptr.Pointer(), C.int(int32(index)), groupC)
	}
}

//export callbackQDesignerMemberSheetExtension_SetVisible
func callbackQDesignerMemberSheetExtension_SetVisible(ptr unsafe.Pointer, index C.int, visible C.char) {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::setVisible"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(visible) != 0)
	}

}

func (ptr *QDesignerMemberSheetExtension) ConnectSetVisible(f func(index int, visible bool)) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::setVisible", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetVisible(index int, visible bool) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::setVisible")
	}
}

func (ptr *QDesignerMemberSheetExtension) SetVisible(index int, visible bool) {
	defer qt.Recovering("QDesignerMemberSheetExtension::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension_SetVisible(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerMemberSheetExtension_Signature
func callbackQDesignerMemberSheetExtension_Signature(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerMemberSheetExtension::signature")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerMemberSheetExtension::signature"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerMemberSheetExtension) ConnectSignature(f func(index int) string) {
	defer qt.Recovering("connect QDesignerMemberSheetExtension::signature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::signature", f)
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSignature(index int) {
	defer qt.Recovering("disconnect QDesignerMemberSheetExtension::signature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerMemberSheetExtension::signature")
	}
}

func (ptr *QDesignerMemberSheetExtension) Signature(index int) string {
	defer qt.Recovering("QDesignerMemberSheetExtension::signature")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerMemberSheetExtension_Signature(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QDesignerMemberSheetExtension) DestroyQDesignerMemberSheetExtension() {
	defer qt.Recovering("QDesignerMemberSheetExtension::~QDesignerMemberSheetExtension")

	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerObjectInspectorInterface struct {
	widgets.QWidget
}

type QDesignerObjectInspectorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerObjectInspectorInterface_PTR() *QDesignerObjectInspectorInterface
}

func (p *QDesignerObjectInspectorInterface) QDesignerObjectInspectorInterface_PTR() *QDesignerObjectInspectorInterface {
	return p
}

func (p *QDesignerObjectInspectorInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerObjectInspectorInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerObjectInspectorInterface(ptr QDesignerObjectInspectorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerObjectInspectorInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerObjectInspectorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerObjectInspectorInterface {
	var n = new(QDesignerObjectInspectorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerObjectInspectorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerObjectInspectorInterface {
	defer qt.Recovering("QDesignerObjectInspectorInterface::QDesignerObjectInspectorInterface")

	var tmpValue = NewQDesignerObjectInspectorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_NewQDesignerObjectInspectorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQDesignerObjectInspectorInterface_Core
func callbackQDesignerObjectInspectorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::core")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerObjectInspectorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::core")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::core", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectCore() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::core")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::core")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Core() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerObjectInspectorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) CoreDefault() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerObjectInspectorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_SetFormWindow
func callbackQDesignerObjectInspectorInterface_SetFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setFormWindow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setFormWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setFormWindow", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setFormWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setFormWindow")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setFormWindow")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DestroyQDesignerObjectInspectorInterface() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::~QDesignerObjectInspectorInterface")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerObjectInspectorInterface_ActionEvent
func callbackQDesignerObjectInspectorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::actionEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::actionEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DragEnterEvent
func callbackQDesignerObjectInspectorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragEnterEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragEnterEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DragLeaveEvent
func callbackQDesignerObjectInspectorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragLeaveEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragLeaveEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DragMoveEvent
func callbackQDesignerObjectInspectorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragMoveEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dragMoveEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DropEvent
func callbackQDesignerObjectInspectorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dropEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::dropEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_EnterEvent
func callbackQDesignerObjectInspectorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::enterEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::enterEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_FocusInEvent
func callbackQDesignerObjectInspectorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusInEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusInEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_FocusOutEvent
func callbackQDesignerObjectInspectorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusOutEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusOutEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_HideEvent
func callbackQDesignerObjectInspectorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hideEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hideEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_LeaveEvent
func callbackQDesignerObjectInspectorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::leaveEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::leaveEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MinimumSizeHint
func callbackQDesignerObjectInspectorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::minimumSizeHint", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::minimumSizeHint")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDesignerObjectInspectorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerObjectInspectorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_MoveEvent
func callbackQDesignerObjectInspectorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::moveEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::moveEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_PaintEvent
func callbackQDesignerObjectInspectorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::paintEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::paintEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetEnabled
func callbackQDesignerObjectInspectorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setEnabled", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setEnabled")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetEnabled(vbo bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetStyleSheet
func callbackQDesignerObjectInspectorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setStyleSheet", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setStyleSheet")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerObjectInspectorInterface_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerObjectInspectorInterface_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQDesignerObjectInspectorInterface_SetVisible
func callbackQDesignerObjectInspectorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setVisible", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setVisible")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetVisible(visible bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetWindowModified
func callbackQDesignerObjectInspectorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setWindowModified", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setWindowModified")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowModified(vbo bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetWindowTitle
func callbackQDesignerObjectInspectorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setWindowTitle", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setWindowTitle")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowTitle(vqs string) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerObjectInspectorInterface_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerObjectInspectorInterface_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowEvent
func callbackQDesignerObjectInspectorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_SizeHint
func callbackQDesignerObjectInspectorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::sizeHint", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::sizeHint")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SizeHint() *core.QSize {
	defer qt.Recovering("QDesignerObjectInspectorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerObjectInspectorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_ChangeEvent
func callbackQDesignerObjectInspectorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::changeEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::changeEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Close
func callbackQDesignerObjectInspectorInterface_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::close", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectClose() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::close")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Close() bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerObjectInspectorInterface) CloseDefault() bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_CloseEvent
func callbackQDesignerObjectInspectorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::closeEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::closeEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ContextMenuEvent
func callbackQDesignerObjectInspectorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::contextMenuEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::contextMenuEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_FocusNextPrevChild
func callbackQDesignerObjectInspectorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusNextPrevChild", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::focusNextPrevChild")
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QDesignerObjectInspectorInterface) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_HasHeightForWidth
func callbackQDesignerObjectInspectorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hasHeightForWidth", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hasHeightForWidth")
	}
}

func (ptr *QDesignerObjectInspectorInterface) HasHeightForWidth() bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerObjectInspectorInterface) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_HeightForWidth
func callbackQDesignerObjectInspectorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::heightForWidth", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::heightForWidth")
	}
}

func (ptr *QDesignerObjectInspectorInterface) HeightForWidth(w int) int {
	defer qt.Recovering("QDesignerObjectInspectorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerObjectInspectorInterface_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QDesignerObjectInspectorInterface) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QDesignerObjectInspectorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerObjectInspectorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerObjectInspectorInterface_Hide
func callbackQDesignerObjectInspectorInterface_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectHide(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hide", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectHide() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::hide")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Hide() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Hide(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) HideDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_InputMethodEvent
func callbackQDesignerObjectInspectorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::inputMethodEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::inputMethodEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_InputMethodQuery
func callbackQDesignerObjectInspectorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::inputMethodQuery", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::inputMethodQuery")
	}
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerObjectInspectorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerObjectInspectorInterface_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerObjectInspectorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerObjectInspectorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_KeyPressEvent
func callbackQDesignerObjectInspectorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::keyPressEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::keyPressEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_KeyReleaseEvent
func callbackQDesignerObjectInspectorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::keyReleaseEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::keyReleaseEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Lower
func callbackQDesignerObjectInspectorInterface_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectLower(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::lower", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectLower() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::lower")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Lower() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Lower(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) LowerDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent
func callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseDoubleClickEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseMoveEvent
func callbackQDesignerObjectInspectorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseMoveEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseMoveEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MousePressEvent
func callbackQDesignerObjectInspectorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mousePressEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mousePressEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseReleaseEvent
func callbackQDesignerObjectInspectorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseReleaseEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::mouseReleaseEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_NativeEvent
func callbackQDesignerObjectInspectorInterface_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::nativeEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::nativeEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerObjectInspectorInterface_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QDesignerObjectInspectorInterface) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerObjectInspectorInterface_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_Raise
func callbackQDesignerObjectInspectorInterface_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectRaise(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::raise", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectRaise() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::raise")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Raise() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Raise(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) RaiseDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_Repaint
func callbackQDesignerObjectInspectorInterface_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::repaint", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectRepaint() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::repaint")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Repaint() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Repaint(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) RepaintDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ResizeEvent
func callbackQDesignerObjectInspectorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::resizeEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::resizeEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetDisabled
func callbackQDesignerObjectInspectorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setDisabled", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setDisabled")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetDisabled(disable bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetFocus2
func callbackQDesignerObjectInspectorInterface_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setFocus2", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setFocus2")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetFocus2() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetFocus2Default() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_SetHidden
func callbackQDesignerObjectInspectorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setHidden", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::setHidden")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetHidden(hidden bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerObjectInspectorInterface_Show
func callbackQDesignerObjectInspectorInterface_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShow(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::show", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShow() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::show")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Show() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Show(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowFullScreen
func callbackQDesignerObjectInspectorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showFullScreen", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showFullScreen")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowFullScreen() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowFullScreenDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowMaximized
func callbackQDesignerObjectInspectorInterface_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showMaximized", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showMaximized")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMaximized() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMaximizedDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowMinimized
func callbackQDesignerObjectInspectorInterface_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showMinimized", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showMinimized")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMinimized() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMinimizedDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowNormal
func callbackQDesignerObjectInspectorInterface_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showNormal", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::showNormal")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowNormal() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowNormalDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_TabletEvent
func callbackQDesignerObjectInspectorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::tabletEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::tabletEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Update
func callbackQDesignerObjectInspectorInterface_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::update", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectUpdate() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::update")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Update() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_Update(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) UpdateDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_UpdateMicroFocus
func callbackQDesignerObjectInspectorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::updateMicroFocus", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::updateMicroFocus")
	}
}

func (ptr *QDesignerObjectInspectorInterface) UpdateMicroFocus() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QDesignerObjectInspectorInterface) UpdateMicroFocusDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_WheelEvent
func callbackQDesignerObjectInspectorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::wheelEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::wheelEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_TimerEvent
func callbackQDesignerObjectInspectorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::timerEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::timerEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ChildEvent
func callbackQDesignerObjectInspectorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::childEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::childEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ConnectNotify
func callbackQDesignerObjectInspectorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::connectNotify", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::connectNotify")
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerObjectInspectorInterface_CustomEvent
func callbackQDesignerObjectInspectorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::customEvent", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::customEvent")
	}
}

func (ptr *QDesignerObjectInspectorInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DeleteLater
func callbackQDesignerObjectInspectorInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::deleteLater", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::deleteLater")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DeleteLater() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerObjectInspectorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerObjectInspectorInterface_DisconnectNotify
func callbackQDesignerObjectInspectorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::disconnectNotify")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerObjectInspectorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerObjectInspectorInterface_EventFilter
func callbackQDesignerObjectInspectorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerObjectInspectorInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::eventFilter", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::eventFilter")
	}
}

func (ptr *QDesignerObjectInspectorInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerObjectInspectorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerObjectInspectorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_MetaObject
func callbackQDesignerObjectInspectorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerObjectInspectorInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerObjectInspectorInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerObjectInspectorInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerObjectInspectorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::metaObject", f)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerObjectInspectorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerObjectInspectorInterface::metaObject")
	}
}

func (ptr *QDesignerObjectInspectorInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerObjectInspectorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerObjectInspectorInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerObjectInspectorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerObjectInspectorInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDesignerPropertyEditorInterface struct {
	widgets.QWidget
}

type QDesignerPropertyEditorInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerPropertyEditorInterface_PTR() *QDesignerPropertyEditorInterface
}

func (p *QDesignerPropertyEditorInterface) QDesignerPropertyEditorInterface_PTR() *QDesignerPropertyEditorInterface {
	return p
}

func (p *QDesignerPropertyEditorInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerPropertyEditorInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerPropertyEditorInterface(ptr QDesignerPropertyEditorInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerPropertyEditorInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerPropertyEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerPropertyEditorInterface {
	var n = new(QDesignerPropertyEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerPropertyEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerPropertyEditorInterface {
	defer qt.Recovering("QDesignerPropertyEditorInterface::QDesignerPropertyEditorInterface")

	var tmpValue = NewQDesignerPropertyEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_NewQDesignerPropertyEditorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQDesignerPropertyEditorInterface_Core
func callbackQDesignerPropertyEditorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::core")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::core")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::core", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCore() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::core")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::core")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Core() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerPropertyEditorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) CoreDefault() *QDesignerFormEditorInterface {
	defer qt.Recovering("QDesignerPropertyEditorInterface::core")

	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_CurrentPropertyName
func callbackQDesignerPropertyEditorInterface_CurrentPropertyName(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::currentPropertyName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::currentPropertyName"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCurrentPropertyName(f func() string) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::currentPropertyName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::currentPropertyName", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCurrentPropertyName() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::currentPropertyName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::currentPropertyName")
	}
}

func (ptr *QDesignerPropertyEditorInterface) CurrentPropertyName() string {
	defer qt.Recovering("QDesignerPropertyEditorInterface::currentPropertyName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerPropertyEditorInterface_CurrentPropertyName(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerPropertyEditorInterface_IsReadOnly
func callbackQDesignerPropertyEditorInterface_IsReadOnly(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::isReadOnly")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::isReadOnly"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectIsReadOnly(f func() bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::isReadOnly")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::isReadOnly", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectIsReadOnly() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::isReadOnly")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::isReadOnly")
	}
}

func (ptr *QDesignerPropertyEditorInterface) IsReadOnly() bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_Object
func callbackQDesignerPropertyEditorInterface_Object(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::object")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::object"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(nil)
}

func (ptr *QDesignerPropertyEditorInterface) ConnectObject(f func() *core.QObject) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::object")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::object", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectObject() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::object")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::object")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Object() *core.QObject {
	defer qt.Recovering("QDesignerPropertyEditorInterface::object")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface_Object(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_PropertyChanged
func callbackQDesignerPropertyEditorInterface_PropertyChanged(ptr unsafe.Pointer, name *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::propertyChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::propertyChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectPropertyChanged(f func(name string, value *core.QVariant)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::propertyChanged")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ConnectPropertyChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::propertyChanged", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectPropertyChanged() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::propertyChanged")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DisconnectPropertyChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::propertyChanged")
	}
}

func (ptr *QDesignerPropertyEditorInterface) PropertyChanged(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::propertyChanged")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDesignerPropertyEditorInterface_PropertyChanged(ptr.Pointer(), nameC, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetObject
func callbackQDesignerPropertyEditorInterface_SetObject(ptr unsafe.Pointer, object unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setObject"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetObject(f func(object *core.QObject)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setObject", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetObject(object core.QObject_ITF) {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setObject")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetObject(object core.QObject_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setObject")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetPropertyValue
func callbackQDesignerPropertyEditorInterface_SetPropertyValue(ptr unsafe.Pointer, name *C.char, value unsafe.Pointer, changed C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setPropertyValue")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setPropertyValue"); signal != nil {
		signal.(func(string, *core.QVariant, bool))(C.GoString(name), core.NewQVariantFromPointer(value), int8(changed) != 0)
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetPropertyValue(f func(name string, value *core.QVariant, changed bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setPropertyValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setPropertyValue", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetPropertyValue(name string, value core.QVariant_ITF, changed bool) {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setPropertyValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setPropertyValue")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetPropertyValue(name string, value core.QVariant_ITF, changed bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setPropertyValue")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDesignerPropertyEditorInterface_SetPropertyValue(ptr.Pointer(), nameC, core.PointerFromQVariant(value), C.char(int8(qt.GoBoolToInt(changed))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetReadOnly
func callbackQDesignerPropertyEditorInterface_SetReadOnly(ptr unsafe.Pointer, readOnly C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setReadOnly")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setReadOnly"); signal != nil {
		signal.(func(bool))(int8(readOnly) != 0)
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetReadOnly(f func(readOnly bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setReadOnly")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setReadOnly", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetReadOnly(readOnly bool) {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setReadOnly")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setReadOnly")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetReadOnly(readOnly bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setReadOnly")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(readOnly))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DestroyQDesignerPropertyEditorInterface() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::~QDesignerPropertyEditorInterface")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerPropertyEditorInterface_ActionEvent
func callbackQDesignerPropertyEditorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::actionEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::actionEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DragEnterEvent
func callbackQDesignerPropertyEditorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragEnterEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragEnterEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DragLeaveEvent
func callbackQDesignerPropertyEditorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragLeaveEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragLeaveEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DragMoveEvent
func callbackQDesignerPropertyEditorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragMoveEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dragMoveEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DropEvent
func callbackQDesignerPropertyEditorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dropEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::dropEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_EnterEvent
func callbackQDesignerPropertyEditorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::enterEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::enterEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_FocusInEvent
func callbackQDesignerPropertyEditorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusInEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusInEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_FocusOutEvent
func callbackQDesignerPropertyEditorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusOutEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusOutEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_HideEvent
func callbackQDesignerPropertyEditorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hideEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hideEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_LeaveEvent
func callbackQDesignerPropertyEditorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::leaveEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::leaveEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MinimumSizeHint
func callbackQDesignerPropertyEditorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::minimumSizeHint", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::minimumSizeHint")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDesignerPropertyEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerPropertyEditorInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_MoveEvent
func callbackQDesignerPropertyEditorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::moveEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::moveEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_PaintEvent
func callbackQDesignerPropertyEditorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::paintEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::paintEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetEnabled
func callbackQDesignerPropertyEditorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setEnabled", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setEnabled")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetEnabled(vbo bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetStyleSheet
func callbackQDesignerPropertyEditorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setStyleSheet", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setStyleSheet")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerPropertyEditorInterface_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerPropertyEditorInterface_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQDesignerPropertyEditorInterface_SetVisible
func callbackQDesignerPropertyEditorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setVisible", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setVisible")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetVisible(visible bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetWindowModified
func callbackQDesignerPropertyEditorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setWindowModified", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setWindowModified")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowModified(vbo bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetWindowTitle
func callbackQDesignerPropertyEditorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setWindowTitle", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setWindowTitle")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowTitle(vqs string) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerPropertyEditorInterface_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerPropertyEditorInterface_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowEvent
func callbackQDesignerPropertyEditorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_SizeHint
func callbackQDesignerPropertyEditorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::sizeHint", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::sizeHint")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SizeHint() *core.QSize {
	defer qt.Recovering("QDesignerPropertyEditorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerPropertyEditorInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_ChangeEvent
func callbackQDesignerPropertyEditorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::changeEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::changeEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Close
func callbackQDesignerPropertyEditorInterface_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::close", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectClose() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::close")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Close() bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) CloseDefault() bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_CloseEvent
func callbackQDesignerPropertyEditorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::closeEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::closeEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ContextMenuEvent
func callbackQDesignerPropertyEditorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::contextMenuEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::contextMenuEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_FocusNextPrevChild
func callbackQDesignerPropertyEditorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusNextPrevChild", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::focusNextPrevChild")
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_HasHeightForWidth
func callbackQDesignerPropertyEditorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hasHeightForWidth", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hasHeightForWidth")
	}
}

func (ptr *QDesignerPropertyEditorInterface) HasHeightForWidth() bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_HeightForWidth
func callbackQDesignerPropertyEditorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::heightForWidth", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::heightForWidth")
	}
}

func (ptr *QDesignerPropertyEditorInterface) HeightForWidth(w int) int {
	defer qt.Recovering("QDesignerPropertyEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertyEditorInterface_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QDesignerPropertyEditorInterface) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QDesignerPropertyEditorInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertyEditorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerPropertyEditorInterface_Hide
func callbackQDesignerPropertyEditorInterface_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectHide(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hide", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectHide() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::hide")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Hide() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Hide(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) HideDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_InputMethodEvent
func callbackQDesignerPropertyEditorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::inputMethodEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::inputMethodEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_InputMethodQuery
func callbackQDesignerPropertyEditorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::inputMethodQuery", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::inputMethodQuery")
	}
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerPropertyEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerPropertyEditorInterface_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerPropertyEditorInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerPropertyEditorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_KeyPressEvent
func callbackQDesignerPropertyEditorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::keyPressEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::keyPressEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_KeyReleaseEvent
func callbackQDesignerPropertyEditorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::keyReleaseEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::keyReleaseEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Lower
func callbackQDesignerPropertyEditorInterface_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectLower(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::lower", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectLower() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::lower")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Lower() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Lower(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) LowerDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent
func callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseDoubleClickEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseMoveEvent
func callbackQDesignerPropertyEditorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseMoveEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseMoveEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MousePressEvent
func callbackQDesignerPropertyEditorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mousePressEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mousePressEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseReleaseEvent
func callbackQDesignerPropertyEditorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseReleaseEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::mouseReleaseEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_NativeEvent
func callbackQDesignerPropertyEditorInterface_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::nativeEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::nativeEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerPropertyEditorInterface_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerPropertyEditorInterface_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_Raise
func callbackQDesignerPropertyEditorInterface_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectRaise(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::raise", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectRaise() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::raise")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Raise() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Raise(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) RaiseDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_Repaint
func callbackQDesignerPropertyEditorInterface_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::repaint", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectRepaint() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::repaint")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Repaint() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Repaint(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) RepaintDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ResizeEvent
func callbackQDesignerPropertyEditorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::resizeEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::resizeEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetDisabled
func callbackQDesignerPropertyEditorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setDisabled", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setDisabled")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetDisabled(disable bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetFocus2
func callbackQDesignerPropertyEditorInterface_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setFocus2", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setFocus2")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetFocus2() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetFocus2Default() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_SetHidden
func callbackQDesignerPropertyEditorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setHidden", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::setHidden")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetHidden(hidden bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerPropertyEditorInterface_Show
func callbackQDesignerPropertyEditorInterface_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShow(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::show", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShow() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::show")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Show() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Show(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowFullScreen
func callbackQDesignerPropertyEditorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showFullScreen", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showFullScreen")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowFullScreen() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowFullScreenDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowMaximized
func callbackQDesignerPropertyEditorInterface_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showMaximized", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showMaximized")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMaximized() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMaximizedDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowMinimized
func callbackQDesignerPropertyEditorInterface_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showMinimized", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showMinimized")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMinimized() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMinimizedDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowNormal
func callbackQDesignerPropertyEditorInterface_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showNormal", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::showNormal")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowNormal() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowNormalDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_TabletEvent
func callbackQDesignerPropertyEditorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::tabletEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::tabletEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Update
func callbackQDesignerPropertyEditorInterface_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::update", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectUpdate() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::update")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Update() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_Update(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) UpdateDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_UpdateMicroFocus
func callbackQDesignerPropertyEditorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::updateMicroFocus", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::updateMicroFocus")
	}
}

func (ptr *QDesignerPropertyEditorInterface) UpdateMicroFocus() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QDesignerPropertyEditorInterface) UpdateMicroFocusDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_WheelEvent
func callbackQDesignerPropertyEditorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::wheelEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::wheelEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_TimerEvent
func callbackQDesignerPropertyEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::timerEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::timerEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ChildEvent
func callbackQDesignerPropertyEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::childEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::childEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ConnectNotify
func callbackQDesignerPropertyEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::connectNotify", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::connectNotify")
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerPropertyEditorInterface_CustomEvent
func callbackQDesignerPropertyEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::customEvent", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::customEvent")
	}
}

func (ptr *QDesignerPropertyEditorInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DeleteLater
func callbackQDesignerPropertyEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::deleteLater", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::deleteLater")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DeleteLater() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerPropertyEditorInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerPropertyEditorInterface_DisconnectNotify
func callbackQDesignerPropertyEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::disconnectNotify")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerPropertyEditorInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerPropertyEditorInterface_EventFilter
func callbackQDesignerPropertyEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::eventFilter", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::eventFilter")
	}
}

func (ptr *QDesignerPropertyEditorInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerPropertyEditorInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_MetaObject
func callbackQDesignerPropertyEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertyEditorInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertyEditorInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerPropertyEditorInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerPropertyEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::metaObject", f)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerPropertyEditorInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertyEditorInterface::metaObject")
	}
}

func (ptr *QDesignerPropertyEditorInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerPropertyEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerPropertyEditorInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerPropertyEditorInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerPropertyEditorInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDesignerPropertySheetExtension struct {
	ptr unsafe.Pointer
}

type QDesignerPropertySheetExtension_ITF interface {
	QDesignerPropertySheetExtension_PTR() *QDesignerPropertySheetExtension
}

func (p *QDesignerPropertySheetExtension) QDesignerPropertySheetExtension_PTR() *QDesignerPropertySheetExtension {
	return p
}

func (p *QDesignerPropertySheetExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerPropertySheetExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerPropertySheetExtension(ptr QDesignerPropertySheetExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerPropertySheetExtension_PTR().Pointer()
	}
	return nil
}

func NewQDesignerPropertySheetExtensionFromPointer(ptr unsafe.Pointer) *QDesignerPropertySheetExtension {
	var n = new(QDesignerPropertySheetExtension)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerPropertySheetExtension_Count
func callbackQDesignerPropertySheetExtension_Count(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::count")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerPropertySheetExtension) ConnectCount(f func() int) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::count")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::count", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectCount() {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::count")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::count")
	}
}

func (ptr *QDesignerPropertySheetExtension) Count() int {
	defer qt.Recovering("QDesignerPropertySheetExtension::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertySheetExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerPropertySheetExtension_HasReset
func callbackQDesignerPropertySheetExtension_HasReset(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::hasReset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::hasReset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectHasReset(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::hasReset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::hasReset", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectHasReset(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::hasReset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::hasReset")
	}
}

func (ptr *QDesignerPropertySheetExtension) HasReset(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::hasReset")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_HasReset(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IndexOf
func callbackQDesignerPropertySheetExtension_IndexOf(ptr unsafe.Pointer, name *C.char) C.int {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::indexOf")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::indexOf"); signal != nil {
		return C.int(int32(signal.(func(string) int)(C.GoString(name))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIndexOf(f func(name string) int) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::indexOf")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::indexOf", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIndexOf(name string) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::indexOf")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::indexOf")
	}
}

func (ptr *QDesignerPropertySheetExtension) IndexOf(name string) int {
	defer qt.Recovering("QDesignerPropertySheetExtension::indexOf")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return int(int32(C.QDesignerPropertySheetExtension_IndexOf(ptr.Pointer(), nameC)))
	}
	return 0
}

//export callbackQDesignerPropertySheetExtension_IsAttribute
func callbackQDesignerPropertySheetExtension_IsAttribute(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::isAttribute")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::isAttribute"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsAttribute(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::isAttribute")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isAttribute", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsAttribute(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::isAttribute")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isAttribute")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsAttribute(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::isAttribute")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsAttribute(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsChanged
func callbackQDesignerPropertySheetExtension_IsChanged(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::isChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::isChanged"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsChanged(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::isChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isChanged", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsChanged(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::isChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isChanged")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsChanged(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::isChanged")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsChanged(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsEnabled
func callbackQDesignerPropertySheetExtension_IsEnabled(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::isEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::isEnabled"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertySheetExtensionFromPointer(ptr).IsEnabledDefault(int(int32(index))))))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsEnabled(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::isEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isEnabled", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsEnabled() {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::isEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isEnabled")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsEnabled(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::isEnabled")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsEnabled(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QDesignerPropertySheetExtension) IsEnabledDefault(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::isEnabled")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsEnabledDefault(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsVisible
func callbackQDesignerPropertySheetExtension_IsVisible(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::isVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::isVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsVisible(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::isVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isVisible", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsVisible(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::isVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::isVisible")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsVisible(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::isVisible")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsVisible(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_Property
func callbackQDesignerPropertySheetExtension_Property(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::property")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::property"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(index))))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QDesignerPropertySheetExtension) ConnectProperty(f func(index int) *core.QVariant) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::property")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::property", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectProperty(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::property")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::property")
	}
}

func (ptr *QDesignerPropertySheetExtension) Property(index int) *core.QVariant {
	defer qt.Recovering("QDesignerPropertySheetExtension::property")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerPropertySheetExtension_Property(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertySheetExtension_PropertyGroup
func callbackQDesignerPropertySheetExtension_PropertyGroup(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::propertyGroup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::propertyGroup"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyGroup(f func(index int) string) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::propertyGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::propertyGroup", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyGroup(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::propertyGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::propertyGroup")
	}
}

func (ptr *QDesignerPropertySheetExtension) PropertyGroup(index int) string {
	defer qt.Recovering("QDesignerPropertySheetExtension::propertyGroup")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerPropertySheetExtension_PropertyGroup(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerPropertySheetExtension_PropertyName
func callbackQDesignerPropertySheetExtension_PropertyName(ptr unsafe.Pointer, index C.int) *C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::propertyName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::propertyName"); signal != nil {
		return C.CString(signal.(func(int) string)(int(int32(index))))
	}

	return C.CString("")
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyName(f func(index int) string) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::propertyName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::propertyName", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyName(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::propertyName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::propertyName")
	}
}

func (ptr *QDesignerPropertySheetExtension) PropertyName(index int) string {
	defer qt.Recovering("QDesignerPropertySheetExtension::propertyName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerPropertySheetExtension_PropertyName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerPropertySheetExtension_Reset
func callbackQDesignerPropertySheetExtension_Reset(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectReset(f func(index int) bool) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::reset", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectReset(index int) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::reset")
	}
}

func (ptr *QDesignerPropertySheetExtension) Reset(index int) bool {
	defer qt.Recovering("QDesignerPropertySheetExtension::reset")

	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_Reset(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_SetAttribute
func callbackQDesignerPropertySheetExtension_SetAttribute(ptr unsafe.Pointer, index C.int, attribute C.char) {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::setAttribute")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::setAttribute"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(attribute) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetAttribute(f func(index int, attribute bool)) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::setAttribute")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setAttribute", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetAttribute(index int, attribute bool) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::setAttribute")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setAttribute")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetAttribute(index int, attribute bool) {
	defer qt.Recovering("QDesignerPropertySheetExtension::setAttribute")

	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetAttribute(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(attribute))))
	}
}

//export callbackQDesignerPropertySheetExtension_SetChanged
func callbackQDesignerPropertySheetExtension_SetChanged(ptr unsafe.Pointer, index C.int, changed C.char) {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::setChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::setChanged"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(changed) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetChanged(f func(index int, changed bool)) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::setChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setChanged", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetChanged(index int, changed bool) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::setChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setChanged")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetChanged(index int, changed bool) {
	defer qt.Recovering("QDesignerPropertySheetExtension::setChanged")

	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetChanged(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(changed))))
	}
}

//export callbackQDesignerPropertySheetExtension_SetProperty
func callbackQDesignerPropertySheetExtension_SetProperty(ptr unsafe.Pointer, index C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::setProperty")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::setProperty"); signal != nil {
		signal.(func(int, *core.QVariant))(int(int32(index)), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetProperty(f func(index int, value *core.QVariant)) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::setProperty")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setProperty", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetProperty(index int, value core.QVariant_ITF) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::setProperty")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setProperty")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetProperty(index int, value core.QVariant_ITF) {
	defer qt.Recovering("QDesignerPropertySheetExtension::setProperty")

	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetProperty(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerPropertySheetExtension_SetPropertyGroup
func callbackQDesignerPropertySheetExtension_SetPropertyGroup(ptr unsafe.Pointer, index C.int, group *C.char) {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::setPropertyGroup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::setPropertyGroup"); signal != nil {
		signal.(func(int, string))(int(int32(index)), C.GoString(group))
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetPropertyGroup(f func(index int, group string)) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::setPropertyGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setPropertyGroup", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetPropertyGroup(index int, group string) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::setPropertyGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setPropertyGroup")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetPropertyGroup(index int, group string) {
	defer qt.Recovering("QDesignerPropertySheetExtension::setPropertyGroup")

	if ptr.Pointer() != nil {
		var groupC = C.CString(group)
		defer C.free(unsafe.Pointer(groupC))
		C.QDesignerPropertySheetExtension_SetPropertyGroup(ptr.Pointer(), C.int(int32(index)), groupC)
	}
}

//export callbackQDesignerPropertySheetExtension_SetVisible
func callbackQDesignerPropertySheetExtension_SetVisible(ptr unsafe.Pointer, index C.int, visible C.char) {
	defer qt.Recovering("callback QDesignerPropertySheetExtension::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerPropertySheetExtension::setVisible"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(visible) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetVisible(f func(index int, visible bool)) {
	defer qt.Recovering("connect QDesignerPropertySheetExtension::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setVisible", f)
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetVisible(index int, visible bool) {
	defer qt.Recovering("disconnect QDesignerPropertySheetExtension::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerPropertySheetExtension::setVisible")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetVisible(index int, visible bool) {
	defer qt.Recovering("QDesignerPropertySheetExtension::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetVisible(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerPropertySheetExtension) DestroyQDesignerPropertySheetExtension() {
	defer qt.Recovering("QDesignerPropertySheetExtension::~QDesignerPropertySheetExtension")

	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerTaskMenuExtension struct {
	ptr unsafe.Pointer
}

type QDesignerTaskMenuExtension_ITF interface {
	QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension
}

func (p *QDesignerTaskMenuExtension) QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension {
	return p
}

func (p *QDesignerTaskMenuExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDesignerTaskMenuExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDesignerTaskMenuExtension(ptr QDesignerTaskMenuExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerTaskMenuExtension_PTR().Pointer()
	}
	return nil
}

func NewQDesignerTaskMenuExtensionFromPointer(ptr unsafe.Pointer) *QDesignerTaskMenuExtension {
	var n = new(QDesignerTaskMenuExtension)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerTaskMenuExtension_PreferredEditAction
func callbackQDesignerTaskMenuExtension_PreferredEditAction(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerTaskMenuExtension::preferredEditAction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerTaskMenuExtension::preferredEditAction"); signal != nil {
		return widgets.PointerFromQAction(signal.(func() *widgets.QAction)())
	}

	return widgets.PointerFromQAction(NewQDesignerTaskMenuExtensionFromPointer(ptr).PreferredEditActionDefault())
}

func (ptr *QDesignerTaskMenuExtension) ConnectPreferredEditAction(f func() *widgets.QAction) {
	defer qt.Recovering("connect QDesignerTaskMenuExtension::preferredEditAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerTaskMenuExtension::preferredEditAction", f)
	}
}

func (ptr *QDesignerTaskMenuExtension) DisconnectPreferredEditAction() {
	defer qt.Recovering("disconnect QDesignerTaskMenuExtension::preferredEditAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerTaskMenuExtension::preferredEditAction")
	}
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditAction() *widgets.QAction {
	defer qt.Recovering("QDesignerTaskMenuExtension::preferredEditAction")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerTaskMenuExtension_PreferredEditAction(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditActionDefault() *widgets.QAction {
	defer qt.Recovering("QDesignerTaskMenuExtension::preferredEditAction")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerTaskMenuExtension_PreferredEditActionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) DestroyQDesignerTaskMenuExtension() {
	defer qt.Recovering("QDesignerTaskMenuExtension::~QDesignerTaskMenuExtension")

	if ptr.Pointer() != nil {
		C.QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QDesignerWidgetBoxInterface struct {
	widgets.QWidget
}

type QDesignerWidgetBoxInterface_ITF interface {
	widgets.QWidget_ITF
	QDesignerWidgetBoxInterface_PTR() *QDesignerWidgetBoxInterface
}

func (p *QDesignerWidgetBoxInterface) QDesignerWidgetBoxInterface_PTR() *QDesignerWidgetBoxInterface {
	return p
}

func (p *QDesignerWidgetBoxInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QDesignerWidgetBoxInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQDesignerWidgetBoxInterface(ptr QDesignerWidgetBoxInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDesignerWidgetBoxInterface_PTR().Pointer()
	}
	return nil
}

func NewQDesignerWidgetBoxInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerWidgetBoxInterface {
	var n = new(QDesignerWidgetBoxInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerWidgetBoxInterface_FileName
func callbackQDesignerWidgetBoxInterface_FileName(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::fileName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::fileName"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QDesignerWidgetBoxInterface) ConnectFileName(f func() string) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::fileName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::fileName", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFileName() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::fileName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::fileName")
	}
}

func (ptr *QDesignerWidgetBoxInterface) FileName() string {
	defer qt.Recovering("QDesignerWidgetBoxInterface::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDesignerWidgetBoxInterface_FileName(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerWidgetBoxInterface_Load
func callbackQDesignerWidgetBoxInterface_Load(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::load"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectLoad(f func() bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::load", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectLoad() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::load")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Load() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::load")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_Load(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_Save
func callbackQDesignerWidgetBoxInterface_Save(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::save")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::save"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSave(f func() bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::save")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::save", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSave() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::save")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::save")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Save() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::save")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_Save(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_SetFileName
func callbackQDesignerWidgetBoxInterface_SetFileName(ptr unsafe.Pointer, fileName *C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setFileName")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setFileName"); signal != nil {
		signal.(func(string))(C.GoString(fileName))
	}

}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetFileName(f func(fileName string)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setFileName")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setFileName", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetFileName(fileName string) {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setFileName")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setFileName")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetFileName(fileName string) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setFileName")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QDesignerWidgetBoxInterface_SetFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DestroyQDesignerWidgetBoxInterface() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::~QDesignerWidgetBoxInterface")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerWidgetBoxInterface_ActionEvent
func callbackQDesignerWidgetBoxInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::actionEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::actionEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::actionEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DragEnterEvent
func callbackQDesignerWidgetBoxInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragEnterEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragEnterEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DragLeaveEvent
func callbackQDesignerWidgetBoxInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragLeaveEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragLeaveEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DragMoveEvent
func callbackQDesignerWidgetBoxInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragMoveEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dragMoveEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DropEvent
func callbackQDesignerWidgetBoxInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dropEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::dropEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::dropEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_EnterEvent
func callbackQDesignerWidgetBoxInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::enterEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::enterEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::enterEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_FocusInEvent
func callbackQDesignerWidgetBoxInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusInEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusInEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_FocusOutEvent
func callbackQDesignerWidgetBoxInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusOutEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusOutEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_HideEvent
func callbackQDesignerWidgetBoxInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hideEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hideEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hideEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_LeaveEvent
func callbackQDesignerWidgetBoxInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::leaveEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::leaveEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MinimumSizeHint
func callbackQDesignerWidgetBoxInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::minimumSizeHint", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::minimumSizeHint")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDesignerWidgetBoxInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerWidgetBoxInterface::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_MoveEvent
func callbackQDesignerWidgetBoxInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::moveEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::moveEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::moveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_PaintEvent
func callbackQDesignerWidgetBoxInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::paintEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::paintEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::paintEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::paintEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetEnabled
func callbackQDesignerWidgetBoxInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setEnabled", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setEnabled")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetEnabled(vbo bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setEnabled")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetStyleSheet
func callbackQDesignerWidgetBoxInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setStyleSheet", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setStyleSheet")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerWidgetBoxInterface_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QDesignerWidgetBoxInterface_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQDesignerWidgetBoxInterface_SetVisible
func callbackQDesignerWidgetBoxInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setVisible", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setVisible")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetVisible(visible bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setVisible")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetWindowModified
func callbackQDesignerWidgetBoxInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setWindowModified", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setWindowModified")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowModified(vbo bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setWindowModified")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetWindowTitle
func callbackQDesignerWidgetBoxInterface_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setWindowTitle", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setWindowTitle")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowTitle(vqs string) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerWidgetBoxInterface_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QDesignerWidgetBoxInterface_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowEvent
func callbackQDesignerWidgetBoxInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_SizeHint
func callbackQDesignerWidgetBoxInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::sizeHint", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::sizeHint")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SizeHint() *core.QSize {
	defer qt.Recovering("QDesignerWidgetBoxInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QDesignerWidgetBoxInterface::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_ChangeEvent
func callbackQDesignerWidgetBoxInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::changeEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::changeEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::changeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Close
func callbackQDesignerWidgetBoxInterface_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::close", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectClose() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::close")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Close() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerWidgetBoxInterface) CloseDefault() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::close")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_CloseEvent
func callbackQDesignerWidgetBoxInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::closeEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::closeEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::closeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ContextMenuEvent
func callbackQDesignerWidgetBoxInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::contextMenuEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::contextMenuEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_FocusNextPrevChild
func callbackQDesignerWidgetBoxInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusNextPrevChild", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::focusNextPrevChild")
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QDesignerWidgetBoxInterface) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_HasHeightForWidth
func callbackQDesignerWidgetBoxInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hasHeightForWidth", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hasHeightForWidth")
	}
}

func (ptr *QDesignerWidgetBoxInterface) HasHeightForWidth() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerWidgetBoxInterface) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_HeightForWidth
func callbackQDesignerWidgetBoxInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::heightForWidth", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::heightForWidth")
	}
}

func (ptr *QDesignerWidgetBoxInterface) HeightForWidth(w int) int {
	defer qt.Recovering("QDesignerWidgetBoxInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerWidgetBoxInterface_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QDesignerWidgetBoxInterface) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QDesignerWidgetBoxInterface::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerWidgetBoxInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerWidgetBoxInterface_Hide
func callbackQDesignerWidgetBoxInterface_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectHide(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hide", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectHide() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::hide")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Hide() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Hide(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) HideDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::hide")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_InputMethodEvent
func callbackQDesignerWidgetBoxInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::inputMethodEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::inputMethodEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_InputMethodQuery
func callbackQDesignerWidgetBoxInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::inputMethodQuery", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::inputMethodQuery")
	}
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerWidgetBoxInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerWidgetBoxInterface_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QDesignerWidgetBoxInterface::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerWidgetBoxInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_KeyPressEvent
func callbackQDesignerWidgetBoxInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::keyPressEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::keyPressEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_KeyReleaseEvent
func callbackQDesignerWidgetBoxInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::keyReleaseEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::keyReleaseEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Lower
func callbackQDesignerWidgetBoxInterface_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectLower(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::lower", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectLower() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::lower")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Lower() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Lower(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) LowerDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::lower")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent
func callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseDoubleClickEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseDoubleClickEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseMoveEvent
func callbackQDesignerWidgetBoxInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseMoveEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseMoveEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MousePressEvent
func callbackQDesignerWidgetBoxInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mousePressEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mousePressEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseReleaseEvent
func callbackQDesignerWidgetBoxInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseReleaseEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::mouseReleaseEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_NativeEvent
func callbackQDesignerWidgetBoxInterface_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::nativeEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::nativeEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerWidgetBoxInterface_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QDesignerWidgetBoxInterface) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QDesignerWidgetBoxInterface_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_Raise
func callbackQDesignerWidgetBoxInterface_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectRaise(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::raise", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectRaise() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::raise")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Raise() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Raise(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) RaiseDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::raise")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_Repaint
func callbackQDesignerWidgetBoxInterface_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::repaint", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectRepaint() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::repaint")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Repaint() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Repaint(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) RepaintDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::repaint")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ResizeEvent
func callbackQDesignerWidgetBoxInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::resizeEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::resizeEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetDisabled
func callbackQDesignerWidgetBoxInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setDisabled", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setDisabled")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetDisabled(disable bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setDisabled")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetFocus2
func callbackQDesignerWidgetBoxInterface_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setFocus2", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setFocus2")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetFocus2() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetFocus2Default() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setFocus")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_SetHidden
func callbackQDesignerWidgetBoxInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setHidden", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::setHidden")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetHidden(hidden bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::setHidden")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerWidgetBoxInterface_Show
func callbackQDesignerWidgetBoxInterface_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShow(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::show", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShow() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::show")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Show() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Show(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::show")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowFullScreen
func callbackQDesignerWidgetBoxInterface_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showFullScreen", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showFullScreen")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowFullScreen() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowFullScreenDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showFullScreen")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowMaximized
func callbackQDesignerWidgetBoxInterface_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showMaximized", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showMaximized")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMaximized() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMaximizedDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showMaximized")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowMinimized
func callbackQDesignerWidgetBoxInterface_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showMinimized", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showMinimized")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMinimized() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMinimizedDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showMinimized")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowNormal
func callbackQDesignerWidgetBoxInterface_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showNormal", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::showNormal")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowNormal() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowNormalDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::showNormal")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_TabletEvent
func callbackQDesignerWidgetBoxInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::tabletEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::tabletEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Update
func callbackQDesignerWidgetBoxInterface_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::update", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectUpdate() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::update")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Update() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_Update(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) UpdateDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::update")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_UpdateMicroFocus
func callbackQDesignerWidgetBoxInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::updateMicroFocus", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::updateMicroFocus")
	}
}

func (ptr *QDesignerWidgetBoxInterface) UpdateMicroFocus() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QDesignerWidgetBoxInterface) UpdateMicroFocusDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_WheelEvent
func callbackQDesignerWidgetBoxInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::wheelEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::wheelEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_TimerEvent
func callbackQDesignerWidgetBoxInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::timerEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::timerEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ChildEvent
func callbackQDesignerWidgetBoxInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::childEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::childEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ConnectNotify
func callbackQDesignerWidgetBoxInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::connectNotify", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::connectNotify")
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerWidgetBoxInterface_CustomEvent
func callbackQDesignerWidgetBoxInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::customEvent", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::customEvent")
	}
}

func (ptr *QDesignerWidgetBoxInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DeleteLater
func callbackQDesignerWidgetBoxInterface_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::deleteLater", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::deleteLater")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DeleteLater() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DeleteLaterDefault() {
	defer qt.Recovering("QDesignerWidgetBoxInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerWidgetBoxInterface_DisconnectNotify
func callbackQDesignerWidgetBoxInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::disconnectNotify", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::disconnectNotify")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDesignerWidgetBoxInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerWidgetBoxInterface_EventFilter
func callbackQDesignerWidgetBoxInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::eventFilter", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::eventFilter")
	}
}

func (ptr *QDesignerWidgetBoxInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDesignerWidgetBoxInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDesignerWidgetBoxInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_MetaObject
func callbackQDesignerWidgetBoxInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDesignerWidgetBoxInterface::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDesignerWidgetBoxInterface::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerWidgetBoxInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDesignerWidgetBoxInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::metaObject", f)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDesignerWidgetBoxInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDesignerWidgetBoxInterface::metaObject")
	}
}

func (ptr *QDesignerWidgetBoxInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDesignerWidgetBoxInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerWidgetBoxInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDesignerWidgetBoxInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerWidgetBoxInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (p *QExtensionFactory) QExtensionFactory_PTR() *QExtensionFactory {
	return p
}

func (p *QExtensionFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QExtensionFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QAbstractExtensionFactory_PTR().SetPointer(ptr)
	}
}

func PointerFromQExtensionFactory(ptr QExtensionFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExtensionFactory_PTR().Pointer()
	}
	return nil
}

func NewQExtensionFactoryFromPointer(ptr unsafe.Pointer) *QExtensionFactory {
	var n = new(QExtensionFactory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QExtensionFactory) DestroyQExtensionFactory() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQExtensionFactory(parent QExtensionManager_ITF) *QExtensionFactory {
	defer qt.Recovering("QExtensionFactory::QExtensionFactory")

	var tmpValue = NewQExtensionFactoryFromPointer(C.QExtensionFactory_NewQExtensionFactory(PointerFromQExtensionManager(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQExtensionFactory_CreateExtension
func callbackQExtensionFactory_CreateExtension(ptr unsafe.Pointer, object unsafe.Pointer, iid *C.char, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QExtensionFactory::createExtension")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::createExtension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string, *core.QObject) *core.QObject)(core.NewQObjectFromPointer(object), C.GoString(iid), core.NewQObjectFromPointer(parent)))
	}

	return core.PointerFromQObject(NewQExtensionFactoryFromPointer(ptr).CreateExtensionDefault(core.NewQObjectFromPointer(object), C.GoString(iid), core.NewQObjectFromPointer(parent)))
}

func (ptr *QExtensionFactory) ConnectCreateExtension(f func(object *core.QObject, iid string, parent *core.QObject) *core.QObject) {
	defer qt.Recovering("connect QExtensionFactory::createExtension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::createExtension", f)
	}
}

func (ptr *QExtensionFactory) DisconnectCreateExtension() {
	defer qt.Recovering("disconnect QExtensionFactory::createExtension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::createExtension")
	}
}

func (ptr *QExtensionFactory) CreateExtension(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {
	defer qt.Recovering("QExtensionFactory::createExtension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_CreateExtension(ptr.Pointer(), core.PointerFromQObject(object), iidC, core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) CreateExtensionDefault(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {
	defer qt.Recovering("QExtensionFactory::createExtension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_CreateExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), iidC, core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQExtensionFactory_Extension
func callbackQExtensionFactory_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QExtensionFactory::extension")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), C.GoString(iid)))
	}

	return core.PointerFromQObject(NewQExtensionFactoryFromPointer(ptr).ExtensionDefault(core.NewQObjectFromPointer(object), C.GoString(iid)))
}

func (ptr *QExtensionFactory) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	defer qt.Recovering("connect QExtensionFactory::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::extension", f)
	}
}

func (ptr *QExtensionFactory) DisconnectExtension() {
	defer qt.Recovering("disconnect QExtensionFactory::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::extension")
	}
}

func (ptr *QExtensionFactory) Extension(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QExtensionFactory::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_Extension(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QExtensionFactory::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_ExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) ExtensionManager() *QExtensionManager {
	defer qt.Recovering("QExtensionFactory::extensionManager")

	if ptr.Pointer() != nil {
		var tmpValue = NewQExtensionManagerFromPointer(C.QExtensionFactory_ExtensionManager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQExtensionFactory_TimerEvent
func callbackQExtensionFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QExtensionFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::timerEvent", f)
	}
}

func (ptr *QExtensionFactory) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QExtensionFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::timerEvent")
	}
}

func (ptr *QExtensionFactory) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QExtensionFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQExtensionFactory_ChildEvent
func callbackQExtensionFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QExtensionFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::childEvent", f)
	}
}

func (ptr *QExtensionFactory) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QExtensionFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::childEvent")
	}
}

func (ptr *QExtensionFactory) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QExtensionFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQExtensionFactory_ConnectNotify
func callbackQExtensionFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QExtensionFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::connectNotify", f)
	}
}

func (ptr *QExtensionFactory) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QExtensionFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::connectNotify")
	}
}

func (ptr *QExtensionFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionFactory_CustomEvent
func callbackQExtensionFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QExtensionFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::customEvent", f)
	}
}

func (ptr *QExtensionFactory) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QExtensionFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::customEvent")
	}
}

func (ptr *QExtensionFactory) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QExtensionFactory) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QExtensionFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQExtensionFactory_DeleteLater
func callbackQExtensionFactory_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQExtensionFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QExtensionFactory) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QExtensionFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::deleteLater", f)
	}
}

func (ptr *QExtensionFactory) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QExtensionFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::deleteLater")
	}
}

func (ptr *QExtensionFactory) DeleteLater() {
	defer qt.Recovering("QExtensionFactory::deleteLater")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QExtensionFactory) DeleteLaterDefault() {
	defer qt.Recovering("QExtensionFactory::deleteLater")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQExtensionFactory_DisconnectNotify
func callbackQExtensionFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionFactory::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QExtensionFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::disconnectNotify", f)
	}
}

func (ptr *QExtensionFactory) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QExtensionFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::disconnectNotify")
	}
}

func (ptr *QExtensionFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionFactory_Event
func callbackQExtensionFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QExtensionFactory::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QExtensionFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QExtensionFactory::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::event", f)
	}
}

func (ptr *QExtensionFactory) DisconnectEvent() {
	defer qt.Recovering("disconnect QExtensionFactory::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::event")
	}
}

func (ptr *QExtensionFactory) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionFactory::event")

	if ptr.Pointer() != nil {
		return C.QExtensionFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QExtensionFactory) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionFactory::event")

	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQExtensionFactory_EventFilter
func callbackQExtensionFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QExtensionFactory::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QExtensionFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QExtensionFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::eventFilter", f)
	}
}

func (ptr *QExtensionFactory) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QExtensionFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::eventFilter")
	}
}

func (ptr *QExtensionFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QExtensionFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQExtensionFactory_MetaObject
func callbackQExtensionFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QExtensionFactory::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionFactory::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQExtensionFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QExtensionFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QExtensionFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::metaObject", f)
	}
}

func (ptr *QExtensionFactory) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QExtensionFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionFactory::metaObject")
	}
}

func (ptr *QExtensionFactory) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QExtensionFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QExtensionFactory) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QExtensionFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func (p *QExtensionManager) QExtensionManager_PTR() *QExtensionManager {
	return p
}

func (p *QExtensionManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QExtensionManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QAbstractExtensionManager_PTR().SetPointer(ptr)
	}
}

func PointerFromQExtensionManager(ptr QExtensionManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExtensionManager_PTR().Pointer()
	}
	return nil
}

func NewQExtensionManagerFromPointer(ptr unsafe.Pointer) *QExtensionManager {
	var n = new(QExtensionManager)
	n.SetPointer(ptr)
	return n
}
func NewQExtensionManager(parent core.QObject_ITF) *QExtensionManager {
	defer qt.Recovering("QExtensionManager::QExtensionManager")

	var tmpValue = NewQExtensionManagerFromPointer(C.QExtensionManager_NewQExtensionManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) {})
	}
	return tmpValue
}

//export callbackQExtensionManager_Extension
func callbackQExtensionManager_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QExtensionManager::extension")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), C.GoString(iid)))
	}

	return core.PointerFromQObject(NewQExtensionManagerFromPointer(ptr).ExtensionDefault(core.NewQObjectFromPointer(object), C.GoString(iid)))
}

func (ptr *QExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	defer qt.Recovering("connect QExtensionManager::extension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::extension", f)
	}
}

func (ptr *QExtensionManager) DisconnectExtension() {
	defer qt.Recovering("disconnect QExtensionManager::extension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::extension")
	}
}

func (ptr *QExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QExtensionManager::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager_Extension(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {
	defer qt.Recovering("QExtensionManager::extension")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager_ExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), iidC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQExtensionManager_RegisterExtensions
func callbackQExtensionManager_RegisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid *C.char) {
	defer qt.Recovering("callback QExtensionManager::registerExtensions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::registerExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	} else {
		NewQExtensionManagerFromPointer(ptr).RegisterExtensionsDefault(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	}
}

func (ptr *QExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	defer qt.Recovering("connect QExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::registerExtensions", f)
	}
}

func (ptr *QExtensionManager) DisconnectRegisterExtensions() {
	defer qt.Recovering("disconnect QExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::registerExtensions")
	}
}

func (ptr *QExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QExtensionManager_RegisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

func (ptr *QExtensionManager) RegisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QExtensionManager::registerExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QExtensionManager_RegisterExtensionsDefault(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

//export callbackQExtensionManager_UnregisterExtensions
func callbackQExtensionManager_UnregisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid *C.char) {
	defer qt.Recovering("callback QExtensionManager::unregisterExtensions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::unregisterExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	} else {
		NewQExtensionManagerFromPointer(ptr).UnregisterExtensionsDefault(NewQAbstractExtensionFactoryFromPointer(factory), C.GoString(iid))
	}
}

func (ptr *QExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	defer qt.Recovering("connect QExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::unregisterExtensions", f)
	}
}

func (ptr *QExtensionManager) DisconnectUnregisterExtensions() {
	defer qt.Recovering("disconnect QExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::unregisterExtensions")
	}
}

func (ptr *QExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QExtensionManager_UnregisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

func (ptr *QExtensionManager) UnregisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {
	defer qt.Recovering("QExtensionManager::unregisterExtensions")

	if ptr.Pointer() != nil {
		var iidC = C.CString(iid)
		defer C.free(unsafe.Pointer(iidC))
		C.QExtensionManager_UnregisterExtensionsDefault(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), iidC)
	}
}

func (ptr *QExtensionManager) DestroyQExtensionManager() {
	defer qt.Recovering("QExtensionManager::~QExtensionManager")

	if ptr.Pointer() != nil {
		C.QExtensionManager_DestroyQExtensionManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQExtensionManager_TimerEvent
func callbackQExtensionManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QExtensionManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::timerEvent", f)
	}
}

func (ptr *QExtensionManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QExtensionManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::timerEvent")
	}
}

func (ptr *QExtensionManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QExtensionManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QExtensionManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QExtensionManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQExtensionManager_ChildEvent
func callbackQExtensionManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QExtensionManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::childEvent", f)
	}
}

func (ptr *QExtensionManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QExtensionManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::childEvent")
	}
}

func (ptr *QExtensionManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QExtensionManager::childEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QExtensionManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QExtensionManager::childEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQExtensionManager_ConnectNotify
func callbackQExtensionManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QExtensionManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::connectNotify", f)
	}
}

func (ptr *QExtensionManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QExtensionManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::connectNotify")
	}
}

func (ptr *QExtensionManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionManager_CustomEvent
func callbackQExtensionManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QExtensionManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::customEvent", f)
	}
}

func (ptr *QExtensionManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QExtensionManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::customEvent")
	}
}

func (ptr *QExtensionManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QExtensionManager::customEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QExtensionManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QExtensionManager::customEvent")

	if ptr.Pointer() != nil {
		C.QExtensionManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQExtensionManager_DeleteLater
func callbackQExtensionManager_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQExtensionManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QExtensionManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QExtensionManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::deleteLater", f)
	}
}

func (ptr *QExtensionManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QExtensionManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::deleteLater")
	}
}

func (ptr *QExtensionManager) DeleteLater() {
	defer qt.Recovering("QExtensionManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QExtensionManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QExtensionManager) DeleteLaterDefault() {
	defer qt.Recovering("QExtensionManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QExtensionManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQExtensionManager_DisconnectNotify
func callbackQExtensionManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QExtensionManager::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QExtensionManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::disconnectNotify", f)
	}
}

func (ptr *QExtensionManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QExtensionManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::disconnectNotify")
	}
}

func (ptr *QExtensionManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QExtensionManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QExtensionManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionManager_Event
func callbackQExtensionManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QExtensionManager::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QExtensionManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QExtensionManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::event", f)
	}
}

func (ptr *QExtensionManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QExtensionManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::event")
	}
}

func (ptr *QExtensionManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionManager::event")

	if ptr.Pointer() != nil {
		return C.QExtensionManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QExtensionManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionManager::event")

	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQExtensionManager_EventFilter
func callbackQExtensionManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QExtensionManager::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QExtensionManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QExtensionManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::eventFilter", f)
	}
}

func (ptr *QExtensionManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QExtensionManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::eventFilter")
	}
}

func (ptr *QExtensionManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QExtensionManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QExtensionManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQExtensionManager_MetaObject
func callbackQExtensionManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QExtensionManager::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QExtensionManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQExtensionManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QExtensionManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QExtensionManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::metaObject", f)
	}
}

func (ptr *QExtensionManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QExtensionManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QExtensionManager::metaObject")
	}
}

func (ptr *QExtensionManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QExtensionManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QExtensionManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QExtensionManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QFormBuilder struct {
	QAbstractFormBuilder
}

type QFormBuilder_ITF interface {
	QAbstractFormBuilder_ITF
	QFormBuilder_PTR() *QFormBuilder
}

func (p *QFormBuilder) QFormBuilder_PTR() *QFormBuilder {
	return p
}

func (p *QFormBuilder) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractFormBuilder_PTR().Pointer()
	}
	return nil
}

func (p *QFormBuilder) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractFormBuilder_PTR().SetPointer(ptr)
	}
}

func PointerFromQFormBuilder(ptr QFormBuilder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFormBuilder_PTR().Pointer()
	}
	return nil
}

func NewQFormBuilderFromPointer(ptr unsafe.Pointer) *QFormBuilder {
	var n = new(QFormBuilder)
	n.SetPointer(ptr)
	return n
}
func NewQFormBuilder() *QFormBuilder {
	defer qt.Recovering("QFormBuilder::QFormBuilder")

	var tmpValue = NewQFormBuilderFromPointer(C.QFormBuilder_NewQFormBuilder())
	runtime.SetFinalizer(tmpValue, (*QFormBuilder).DestroyQFormBuilder)
	return tmpValue
}

func (ptr *QFormBuilder) AddPluginPath(pluginPath string) {
	defer qt.Recovering("QFormBuilder::addPluginPath")

	if ptr.Pointer() != nil {
		var pluginPathC = C.CString(pluginPath)
		defer C.free(unsafe.Pointer(pluginPathC))
		C.QFormBuilder_AddPluginPath(ptr.Pointer(), pluginPathC)
	}
}

func (ptr *QFormBuilder) ClearPluginPaths() {
	defer qt.Recovering("QFormBuilder::clearPluginPaths")

	if ptr.Pointer() != nil {
		C.QFormBuilder_ClearPluginPaths(ptr.Pointer())
	}
}

func (ptr *QFormBuilder) PluginPaths() []string {
	defer qt.Recovering("QFormBuilder::pluginPaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFormBuilder_PluginPaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFormBuilder) SetPluginPath(pluginPaths []string) {
	defer qt.Recovering("QFormBuilder::setPluginPath")

	if ptr.Pointer() != nil {
		var pluginPathsC = C.CString(strings.Join(pluginPaths, "|"))
		defer C.free(unsafe.Pointer(pluginPathsC))
		C.QFormBuilder_SetPluginPath(ptr.Pointer(), pluginPathsC)
	}
}

func (ptr *QFormBuilder) DestroyQFormBuilder() {
	defer qt.Recovering("QFormBuilder::~QFormBuilder")

	if ptr.Pointer() != nil {
		C.QFormBuilder_DestroyQFormBuilder(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQFormBuilder_Load
func callbackQFormBuilder_Load(ptr unsafe.Pointer, device unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QFormBuilder::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QFormBuilder::load"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*core.QIODevice, *widgets.QWidget) *widgets.QWidget)(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
	}

	return widgets.PointerFromQWidget(NewQFormBuilderFromPointer(ptr).LoadDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
}

func (ptr *QFormBuilder) ConnectLoad(f func(device *core.QIODevice, parent *widgets.QWidget) *widgets.QWidget) {
	defer qt.Recovering("connect QFormBuilder::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QFormBuilder::load", f)
	}
}

func (ptr *QFormBuilder) DisconnectLoad() {
	defer qt.Recovering("disconnect QFormBuilder::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QFormBuilder::load")
	}
}

func (ptr *QFormBuilder) Load(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QFormBuilder::load")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QFormBuilder_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

func (ptr *QFormBuilder) LoadDefault(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QFormBuilder::load")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QFormBuilder_LoadDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) {})
		}
		return tmpValue
	}
	return nil
}

//export callbackQFormBuilder_Save
func callbackQFormBuilder_Save(ptr unsafe.Pointer, device unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QFormBuilder::save")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QFormBuilder::save"); signal != nil {
		signal.(func(*core.QIODevice, *widgets.QWidget))(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQFormBuilderFromPointer(ptr).SaveDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QFormBuilder) ConnectSave(f func(device *core.QIODevice, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QFormBuilder::save")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QFormBuilder::save", f)
	}
}

func (ptr *QFormBuilder) DisconnectSave() {
	defer qt.Recovering("disconnect QFormBuilder::save")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QFormBuilder::save")
	}
}

func (ptr *QFormBuilder) Save(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QFormBuilder::save")

	if ptr.Pointer() != nil {
		C.QFormBuilder_Save(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QFormBuilder) SaveDefault(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QFormBuilder::save")

	if ptr.Pointer() != nil {
		C.QFormBuilder_SaveDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}
