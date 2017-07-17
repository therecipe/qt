// +build !minimal

package designer

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "designer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtDesigner_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QAbstractExtensionFactory struct {
	ptr unsafe.Pointer
}

type QAbstractExtensionFactory_ITF interface {
	QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory
}

func (ptr *QAbstractExtensionFactory) QAbstractExtensionFactory_PTR() *QAbstractExtensionFactory {
	return ptr
}

func (ptr *QAbstractExtensionFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractExtensionFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

type QAbstractExtensionManager struct {
	ptr unsafe.Pointer
}

type QAbstractExtensionManager_ITF interface {
	QAbstractExtensionManager_PTR() *QAbstractExtensionManager
}

func (ptr *QAbstractExtensionManager) QAbstractExtensionManager_PTR() *QAbstractExtensionManager {
	return ptr
}

func (ptr *QAbstractExtensionManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractExtensionManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQAbstractExtensionManager_RegisterExtensions
func callbackQAbstractExtensionManager_RegisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "registerExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	}

}

func (ptr *QAbstractExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerExtensions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "registerExtensions", func(factory *QAbstractExtensionFactory, iid string) {
				signal.(func(*QAbstractExtensionFactory, string))(factory, iid)
				f(factory, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerExtensions", f)
		}
	}
}

func (ptr *QAbstractExtensionManager) DisconnectRegisterExtensions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerExtensions")
	}
}

func (ptr *QAbstractExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QAbstractExtensionManager_RegisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

//export callbackQAbstractExtensionManager_UnregisterExtensions
func callbackQAbstractExtensionManager_UnregisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "unregisterExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	}

}

func (ptr *QAbstractExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unregisterExtensions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unregisterExtensions", func(factory *QAbstractExtensionFactory, iid string) {
				signal.(func(*QAbstractExtensionFactory, string))(factory, iid)
				f(factory, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unregisterExtensions", f)
		}
	}
}

func (ptr *QAbstractExtensionManager) DisconnectUnregisterExtensions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unregisterExtensions")
	}
}

func (ptr *QAbstractExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QAbstractExtensionManager_UnregisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

//export callbackQAbstractExtensionManager_DestroyQAbstractExtensionManager
func callbackQAbstractExtensionManager_DestroyQAbstractExtensionManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractExtensionManager"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractExtensionManagerFromPointer(ptr).DestroyQAbstractExtensionManagerDefault()
	}
}

func (ptr *QAbstractExtensionManager) ConnectDestroyQAbstractExtensionManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractExtensionManager"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractExtensionManager", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractExtensionManager", f)
		}
	}
}

func (ptr *QAbstractExtensionManager) DisconnectDestroyQAbstractExtensionManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractExtensionManager")
	}
}

func (ptr *QAbstractExtensionManager) DestroyQAbstractExtensionManager() {
	if ptr.Pointer() != nil {
		C.QAbstractExtensionManager_DestroyQAbstractExtensionManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractExtensionManager) DestroyQAbstractExtensionManagerDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractExtensionManager_DestroyQAbstractExtensionManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractExtensionManager_Extension
func callbackQAbstractExtensionManager_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid C.struct_QtDesigner_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), cGoUnpackString(iid)))
	}

	return core.PointerFromQObject(core.NewQObject(nil))
}

func (ptr *QAbstractExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "extension", func(object *core.QObject, iid string) *core.QObject {
				signal.(func(*core.QObject, string) *core.QObject)(object, iid)
				return f(object, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", f)
		}
	}
}

func (ptr *QAbstractExtensionManager) DisconnectExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "extension")
	}
}

func (ptr *QAbstractExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QAbstractExtensionManager_Extension(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QAbstractFormBuilder struct {
	ptr unsafe.Pointer
}

type QAbstractFormBuilder_ITF interface {
	QAbstractFormBuilder_PTR() *QAbstractFormBuilder
}

func (ptr *QAbstractFormBuilder) QAbstractFormBuilder_PTR() *QAbstractFormBuilder {
	return ptr
}

func (ptr *QAbstractFormBuilder) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QAbstractFormBuilder) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQAbstractFormBuilder() *QAbstractFormBuilder {
	return NewQAbstractFormBuilderFromPointer(C.QAbstractFormBuilder_NewQAbstractFormBuilder())
}

//export callbackQAbstractFormBuilder_Load
func callbackQAbstractFormBuilder_Load(ptr unsafe.Pointer, device unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*core.QIODevice, *widgets.QWidget) *widgets.QWidget)(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
	}

	return widgets.PointerFromQWidget(NewQAbstractFormBuilderFromPointer(ptr).LoadDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(parent)))
}

func (ptr *QAbstractFormBuilder) ConnectLoad(f func(device *core.QIODevice, parent *widgets.QWidget) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load", func(device *core.QIODevice, parent *widgets.QWidget) *widgets.QWidget {
				signal.(func(*core.QIODevice, *widgets.QWidget) *widgets.QWidget)(device, parent)
				return f(device, parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", f)
		}
	}
}

func (ptr *QAbstractFormBuilder) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QAbstractFormBuilder) Load(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QAbstractFormBuilder_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractFormBuilder) LoadDefault(device core.QIODevice_ITF, parent widgets.QWidget_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QAbstractFormBuilder_LoadDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractFormBuilder_Save
func callbackQAbstractFormBuilder_Save(ptr unsafe.Pointer, device unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "save"); signal != nil {
		signal.(func(*core.QIODevice, *widgets.QWidget))(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQAbstractFormBuilderFromPointer(ptr).SaveDefault(core.NewQIODeviceFromPointer(device), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QAbstractFormBuilder) ConnectSave(f func(device *core.QIODevice, widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "save"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "save", func(device *core.QIODevice, widget *widgets.QWidget) {
				signal.(func(*core.QIODevice, *widgets.QWidget))(device, widget)
				f(device, widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "save", f)
		}
	}
}

func (ptr *QAbstractFormBuilder) DisconnectSave() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "save")
	}
}

func (ptr *QAbstractFormBuilder) Save(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_Save(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractFormBuilder) SaveDefault(device core.QIODevice_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_SaveDefault(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractFormBuilder) SetWorkingDirectory(directory core.QDir_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_SetWorkingDirectory(ptr.Pointer(), core.PointerFromQDir(directory))
	}
}

//export callbackQAbstractFormBuilder_DestroyQAbstractFormBuilder
func callbackQAbstractFormBuilder_DestroyQAbstractFormBuilder(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAbstractFormBuilder"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractFormBuilderFromPointer(ptr).DestroyQAbstractFormBuilderDefault()
	}
}

func (ptr *QAbstractFormBuilder) ConnectDestroyQAbstractFormBuilder(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAbstractFormBuilder"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractFormBuilder", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAbstractFormBuilder", f)
		}
	}
}

func (ptr *QAbstractFormBuilder) DisconnectDestroyQAbstractFormBuilder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAbstractFormBuilder")
	}
}

func (ptr *QAbstractFormBuilder) DestroyQAbstractFormBuilder() {
	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_DestroyQAbstractFormBuilder(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractFormBuilder) DestroyQAbstractFormBuilderDefault() {
	if ptr.Pointer() != nil {
		C.QAbstractFormBuilder_DestroyQAbstractFormBuilderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAbstractFormBuilder) WorkingDirectory() *core.QDir {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDirFromPointer(C.QAbstractFormBuilder_WorkingDirectory(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractFormBuilder) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAbstractFormBuilder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractFormBuilder) __propertyMap_properties_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstractFormBuilder___propertyMap_properties_newList(ptr.Pointer()))
}

func (ptr *QAbstractFormBuilder) __computeProperties_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstractFormBuilder___computeProperties_newList(ptr.Pointer()))
}

func (ptr *QAbstractFormBuilder) __applyProperties_properties_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QAbstractFormBuilder___applyProperties_properties_newList(ptr.Pointer()))
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

func NewQDesignerActionEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerActionEditorInterface {
	var n = new(QDesignerActionEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerActionEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerActionEditorInterface {
	var tmpValue = NewQDesignerActionEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_NewQDesignerActionEditorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDesignerActionEditorInterface_ManageAction
func callbackQDesignerActionEditorInterface_ManageAction(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "manageAction"); signal != nil {
		signal.(func(*widgets.QAction))(widgets.NewQActionFromPointer(action))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectManageAction(f func(action *widgets.QAction)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "manageAction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "manageAction", func(action *widgets.QAction) {
				signal.(func(*widgets.QAction))(action)
				f(action)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "manageAction", f)
		}
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectManageAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "manageAction")
	}
}

func (ptr *QDesignerActionEditorInterface) ManageAction(action widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ManageAction(ptr.Pointer(), widgets.PointerFromQAction(action))
	}
}

//export callbackQDesignerActionEditorInterface_SetFormWindow
func callbackQDesignerActionEditorInterface_SetFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFormWindow", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFormWindow", f)
		}
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectSetFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFormWindow")
	}
}

func (ptr *QDesignerActionEditorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerActionEditorInterface_UnmanageAction
func callbackQDesignerActionEditorInterface_UnmanageAction(ptr unsafe.Pointer, action unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "unmanageAction"); signal != nil {
		signal.(func(*widgets.QAction))(widgets.NewQActionFromPointer(action))
	}

}

func (ptr *QDesignerActionEditorInterface) ConnectUnmanageAction(f func(action *widgets.QAction)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unmanageAction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unmanageAction", func(action *widgets.QAction) {
				signal.(func(*widgets.QAction))(action)
				f(action)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unmanageAction", f)
		}
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectUnmanageAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unmanageAction")
	}
}

func (ptr *QDesignerActionEditorInterface) UnmanageAction(action widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UnmanageAction(ptr.Pointer(), widgets.PointerFromQAction(action))
	}
}

//export callbackQDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface
func callbackQDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerActionEditorInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DestroyQDesignerActionEditorInterfaceDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectDestroyQDesignerActionEditorInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerActionEditorInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerActionEditorInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerActionEditorInterface", f)
		}
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectDestroyQDesignerActionEditorInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerActionEditorInterface")
	}
}

func (ptr *QDesignerActionEditorInterface) DestroyQDesignerActionEditorInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DestroyQDesignerActionEditorInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerActionEditorInterface) DestroyQDesignerActionEditorInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DestroyQDesignerActionEditorInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerActionEditorInterface_Core
func callbackQDesignerActionEditorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerActionEditorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerActionEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "core"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "core", func() *QDesignerFormEditorInterface {
				signal.(func() *QDesignerFormEditorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "core", f)
		}
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "core")
	}
}

func (ptr *QDesignerActionEditorInterface) Core() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) CoreDefault() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerActionEditorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerActionEditorInterface___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerActionEditorInterface___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerActionEditorInterface___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerActionEditorInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerActionEditorInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerActionEditorInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerActionEditorInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerActionEditorInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerActionEditorInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerActionEditorInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerActionEditorInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerActionEditorInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerActionEditorInterface_Close
func callbackQDesignerActionEditorInterface_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerActionEditorInterface) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_Event
func callbackQDesignerActionEditorInterface_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerActionEditorInterface) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_FocusNextPrevChild
func callbackQDesignerActionEditorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerActionEditorInterface) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_NativeEvent
func callbackQDesignerActionEditorInterface_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QDesignerActionEditorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_ActionEvent
func callbackQDesignerActionEditorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ChangeEvent
func callbackQDesignerActionEditorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_CloseEvent
func callbackQDesignerActionEditorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ContextMenuEvent
func callbackQDesignerActionEditorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_CustomContextMenuRequested
func callbackQDesignerActionEditorInterface_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQDesignerActionEditorInterface_DragEnterEvent
func callbackQDesignerActionEditorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DragLeaveEvent
func callbackQDesignerActionEditorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DragMoveEvent
func callbackQDesignerActionEditorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DropEvent
func callbackQDesignerActionEditorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_EnterEvent
func callbackQDesignerActionEditorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_FocusInEvent
func callbackQDesignerActionEditorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_FocusOutEvent
func callbackQDesignerActionEditorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Hide
func callbackQDesignerActionEditorInterface_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) HideDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_HideEvent
func callbackQDesignerActionEditorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_InputMethodEvent
func callbackQDesignerActionEditorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_KeyPressEvent
func callbackQDesignerActionEditorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_KeyReleaseEvent
func callbackQDesignerActionEditorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_LeaveEvent
func callbackQDesignerActionEditorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Lower
func callbackQDesignerActionEditorInterface_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_MouseDoubleClickEvent
func callbackQDesignerActionEditorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MouseMoveEvent
func callbackQDesignerActionEditorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MousePressEvent
func callbackQDesignerActionEditorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MouseReleaseEvent
func callbackQDesignerActionEditorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MoveEvent
func callbackQDesignerActionEditorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_PaintEvent
func callbackQDesignerActionEditorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Raise
func callbackQDesignerActionEditorInterface_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_Repaint
func callbackQDesignerActionEditorInterface_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ResizeEvent
func callbackQDesignerActionEditorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_SetDisabled
func callbackQDesignerActionEditorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerActionEditorInterface_SetEnabled
func callbackQDesignerActionEditorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerActionEditorInterface_SetFocus2
func callbackQDesignerActionEditorInterface_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerActionEditorInterface) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_SetHidden
func callbackQDesignerActionEditorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerActionEditorInterface_SetStyleSheet
func callbackQDesignerActionEditorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QDesignerActionEditorInterface) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QDesignerActionEditorInterface_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQDesignerActionEditorInterface_SetVisible
func callbackQDesignerActionEditorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerActionEditorInterface_SetWindowModified
func callbackQDesignerActionEditorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerActionEditorInterface_SetWindowTitle
func callbackQDesignerActionEditorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QDesignerActionEditorInterface) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QDesignerActionEditorInterface_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQDesignerActionEditorInterface_Show
func callbackQDesignerActionEditorInterface_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowEvent
func callbackQDesignerActionEditorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ShowFullScreen
func callbackQDesignerActionEditorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowMaximized
func callbackQDesignerActionEditorInterface_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowMinimized
func callbackQDesignerActionEditorInterface_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_ShowNormal
func callbackQDesignerActionEditorInterface_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_TabletEvent
func callbackQDesignerActionEditorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_Update
func callbackQDesignerActionEditorInterface_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_UpdateMicroFocus
func callbackQDesignerActionEditorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerActionEditorInterface_WheelEvent
func callbackQDesignerActionEditorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_WindowIconChanged
func callbackQDesignerActionEditorInterface_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQDesignerActionEditorInterface_WindowTitleChanged
func callbackQDesignerActionEditorInterface_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQDesignerActionEditorInterface_PaintEngine
func callbackQDesignerActionEditorInterface_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQDesignerActionEditorInterfaceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QDesignerActionEditorInterface) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QDesignerActionEditorInterface_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_MinimumSizeHint
func callbackQDesignerActionEditorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerActionEditorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerActionEditorInterface) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_SizeHint
func callbackQDesignerActionEditorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerActionEditorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerActionEditorInterface) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerActionEditorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_InputMethodQuery
func callbackQDesignerActionEditorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerActionEditorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerActionEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerActionEditorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerActionEditorInterface_HasHeightForWidth
func callbackQDesignerActionEditorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerActionEditorInterface) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_HeightForWidth
func callbackQDesignerActionEditorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerActionEditorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerActionEditorInterface) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerActionEditorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerActionEditorInterface_Metric
func callbackQDesignerActionEditorInterface_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQDesignerActionEditorInterfaceFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QDesignerActionEditorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerActionEditorInterface_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQDesignerActionEditorInterface_EventFilter
func callbackQDesignerActionEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerActionEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerActionEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerActionEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerActionEditorInterface_ChildEvent
func callbackQDesignerActionEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_ConnectNotify
func callbackQDesignerActionEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerActionEditorInterface_CustomEvent
func callbackQDesignerActionEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_DeleteLater
func callbackQDesignerActionEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerActionEditorInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerActionEditorInterface_Destroyed
func callbackQDesignerActionEditorInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerActionEditorInterface_DisconnectNotify
func callbackQDesignerActionEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerActionEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerActionEditorInterface_ObjectNameChanged
func callbackQDesignerActionEditorInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerActionEditorInterface_TimerEvent
func callbackQDesignerActionEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerActionEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerActionEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerActionEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerActionEditorInterface_MetaObject
func callbackQDesignerActionEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerActionEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerActionEditorInterface) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QDesignerContainerExtension) QDesignerContainerExtension_PTR() *QDesignerContainerExtension {
	return ptr
}

func (ptr *QDesignerContainerExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerContainerExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
	if signal := qt.GetSignal(ptr, "addWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(page))
	}

}

func (ptr *QDesignerContainerExtension) ConnectAddWidget(f func(page *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addWidget", func(page *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(page)
				f(page)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addWidget", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectAddWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addWidget")
	}
}

func (ptr *QDesignerContainerExtension) AddWidget(page widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_AddWidget(ptr.Pointer(), widgets.PointerFromQWidget(page))
	}
}

//export callbackQDesignerContainerExtension_InsertWidget
func callbackQDesignerContainerExtension_InsertWidget(ptr unsafe.Pointer, index C.int, page unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "insertWidget"); signal != nil {
		signal.(func(int, *widgets.QWidget))(int(int32(index)), widgets.NewQWidgetFromPointer(page))
	}

}

func (ptr *QDesignerContainerExtension) ConnectInsertWidget(f func(index int, page *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "insertWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "insertWidget", func(index int, page *widgets.QWidget) {
				signal.(func(int, *widgets.QWidget))(index, page)
				f(index, page)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "insertWidget", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectInsertWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "insertWidget")
	}
}

func (ptr *QDesignerContainerExtension) InsertWidget(index int, page widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_InsertWidget(ptr.Pointer(), C.int(int32(index)), widgets.PointerFromQWidget(page))
	}
}

//export callbackQDesignerContainerExtension_Remove
func callbackQDesignerContainerExtension_Remove(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "remove"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QDesignerContainerExtension) ConnectRemove(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "remove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "remove", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "remove", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "remove")
	}
}

func (ptr *QDesignerContainerExtension) Remove(index int) {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_Remove(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQDesignerContainerExtension_SetCurrentIndex
func callbackQDesignerContainerExtension_SetCurrentIndex(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentIndex"); signal != nil {
		signal.(func(int))(int(int32(index)))
	}

}

func (ptr *QDesignerContainerExtension) ConnectSetCurrentIndex(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentIndex"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", func(index int) {
				signal.(func(int))(index)
				f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentIndex", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectSetCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentIndex")
	}
}

func (ptr *QDesignerContainerExtension) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_SetCurrentIndex(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQDesignerContainerExtension_DestroyQDesignerContainerExtension
func callbackQDesignerContainerExtension_DestroyQDesignerContainerExtension(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerContainerExtension"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerContainerExtensionFromPointer(ptr).DestroyQDesignerContainerExtensionDefault()
	}
}

func (ptr *QDesignerContainerExtension) ConnectDestroyQDesignerContainerExtension(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerContainerExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerContainerExtension", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerContainerExtension", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectDestroyQDesignerContainerExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerContainerExtension")
	}
}

func (ptr *QDesignerContainerExtension) DestroyQDesignerContainerExtension() {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_DestroyQDesignerContainerExtension(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerContainerExtension) DestroyQDesignerContainerExtensionDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerContainerExtension_DestroyQDesignerContainerExtensionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerContainerExtension_Widget
func callbackQDesignerContainerExtension_Widget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerContainerExtension) ConnectWidget(f func(index int) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "widget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widget", func(index int) *widgets.QWidget {
				signal.(func(int) *widgets.QWidget)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widget", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "widget")
	}
}

func (ptr *QDesignerContainerExtension) Widget(index int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerContainerExtension_Widget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerContainerExtension_CanAddWidget
func callbackQDesignerContainerExtension_CanAddWidget(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canAddWidget"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerContainerExtensionFromPointer(ptr).CanAddWidgetDefault())))
}

func (ptr *QDesignerContainerExtension) ConnectCanAddWidget(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "canAddWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canAddWidget", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canAddWidget", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCanAddWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "canAddWidget")
	}
}

func (ptr *QDesignerContainerExtension) CanAddWidget() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanAddWidget(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerContainerExtension) CanAddWidgetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanAddWidgetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerContainerExtension_CanRemove
func callbackQDesignerContainerExtension_CanRemove(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "canRemove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerContainerExtensionFromPointer(ptr).CanRemoveDefault(int(int32(index))))))
}

func (ptr *QDesignerContainerExtension) ConnectCanRemove(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "canRemove"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canRemove", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canRemove", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCanRemove() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "canRemove")
	}
}

func (ptr *QDesignerContainerExtension) CanRemove(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanRemove(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QDesignerContainerExtension) CanRemoveDefault(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerContainerExtension_CanRemoveDefault(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerContainerExtension_Count
func callbackQDesignerContainerExtension_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerContainerExtension) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "count", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QDesignerContainerExtension) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerContainerExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerContainerExtension_CurrentIndex
func callbackQDesignerContainerExtension_CurrentIndex(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "currentIndex"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerContainerExtension) ConnectCurrentIndex(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentIndex"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "currentIndex", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentIndex", f)
		}
	}
}

func (ptr *QDesignerContainerExtension) DisconnectCurrentIndex() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "currentIndex")
	}
}

func (ptr *QDesignerContainerExtension) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerContainerExtension_CurrentIndex(ptr.Pointer())))
	}
	return 0
}

type QDesignerCustomWidgetCollectionInterface struct {
	ptr unsafe.Pointer
}

type QDesignerCustomWidgetCollectionInterface_ITF interface {
	QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface
}

func (ptr *QDesignerCustomWidgetCollectionInterface) QDesignerCustomWidgetCollectionInterface_PTR() *QDesignerCustomWidgetCollectionInterface {
	return ptr
}

func (ptr *QDesignerCustomWidgetCollectionInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerCustomWidgetCollectionInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface
func callbackQDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerCustomWidgetCollectionInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerCustomWidgetCollectionInterfaceFromPointer(ptr).DestroyQDesignerCustomWidgetCollectionInterfaceDefault()
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) ConnectDestroyQDesignerCustomWidgetCollectionInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerCustomWidgetCollectionInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetCollectionInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetCollectionInterface", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DisconnectDestroyQDesignerCustomWidgetCollectionInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetCollectionInterface")
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DestroyQDesignerCustomWidgetCollectionInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DestroyQDesignerCustomWidgetCollectionInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetCollectionInterface_DestroyQDesignerCustomWidgetCollectionInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerCustomWidgetCollectionInterface_CustomWidgets
func callbackQDesignerCustomWidgetCollectionInterface_CustomWidgets(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "customWidgets"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQDesignerCustomWidgetCollectionInterfaceFromPointer(NewQDesignerCustomWidgetCollectionInterfaceFromPointer(nil).__customWidgets_newList())
			for _, v := range signal.(func() []*QDesignerCustomWidgetInterface)() {
				tmpList.__customWidgets_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQDesignerCustomWidgetCollectionInterfaceFromPointer(NewQDesignerCustomWidgetCollectionInterfaceFromPointer(nil).__customWidgets_newList())
		for _, v := range make([]*QDesignerCustomWidgetInterface, 0) {
			tmpList.__customWidgets_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QDesignerCustomWidgetCollectionInterface) ConnectCustomWidgets(f func() []*QDesignerCustomWidgetInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "customWidgets"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "customWidgets", func() []*QDesignerCustomWidgetInterface {
				signal.(func() []*QDesignerCustomWidgetInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "customWidgets", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) DisconnectCustomWidgets() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "customWidgets")
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) CustomWidgets() []*QDesignerCustomWidgetInterface {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDesigner_PackedList) []*QDesignerCustomWidgetInterface {
			var out = make([]*QDesignerCustomWidgetInterface, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDesignerCustomWidgetCollectionInterfaceFromPointer(l.data).__customWidgets_atList(i)
			}
			return out
		}(C.QDesignerCustomWidgetCollectionInterface_CustomWidgets(ptr.Pointer()))
	}
	return make([]*QDesignerCustomWidgetInterface, 0)
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_atList(i int) *QDesignerCustomWidgetInterface {
	if ptr.Pointer() != nil {
		return NewQDesignerCustomWidgetInterfaceFromPointer(C.QDesignerCustomWidgetCollectionInterface___customWidgets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_setList(i QDesignerCustomWidgetInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetCollectionInterface___customWidgets_setList(ptr.Pointer(), PointerFromQDesignerCustomWidgetInterface(i))
	}
}

func (ptr *QDesignerCustomWidgetCollectionInterface) __customWidgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerCustomWidgetCollectionInterface___customWidgets_newList(ptr.Pointer()))
}

type QDesignerCustomWidgetInterface struct {
	ptr unsafe.Pointer
}

type QDesignerCustomWidgetInterface_ITF interface {
	QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface
}

func (ptr *QDesignerCustomWidgetInterface) QDesignerCustomWidgetInterface_PTR() *QDesignerCustomWidgetInterface {
	return ptr
}

func (ptr *QDesignerCustomWidgetInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerCustomWidgetInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerCustomWidgetInterface_CreateWidget
func callbackQDesignerCustomWidgetInterface_CreateWidget(ptr unsafe.Pointer, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(*widgets.QWidget) *widgets.QWidget)(widgets.NewQWidgetFromPointer(parent)))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerCustomWidgetInterface) ConnectCreateWidget(f func(parent *widgets.QWidget) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createWidget", func(parent *widgets.QWidget) *widgets.QWidget {
				signal.(func(*widgets.QWidget) *widgets.QWidget)(parent)
				return f(parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createWidget", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCreateWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createWidget")
	}
}

func (ptr *QDesignerCustomWidgetInterface) CreateWidget(parent widgets.QWidget_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerCustomWidgetInterface_CreateWidget(ptr.Pointer(), widgets.PointerFromQWidget(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerCustomWidgetInterface_Initialize
func callbackQDesignerCustomWidgetInterface_Initialize(ptr unsafe.Pointer, formEditor unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initialize"); signal != nil {
		signal.(func(*QDesignerFormEditorInterface))(NewQDesignerFormEditorInterfaceFromPointer(formEditor))
	} else {
		NewQDesignerCustomWidgetInterfaceFromPointer(ptr).InitializeDefault(NewQDesignerFormEditorInterfaceFromPointer(formEditor))
	}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectInitialize(f func(formEditor *QDesignerFormEditorInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "initialize"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "initialize", func(formEditor *QDesignerFormEditorInterface) {
				signal.(func(*QDesignerFormEditorInterface))(formEditor)
				f(formEditor)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "initialize", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectInitialize() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "initialize")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Initialize(formEditor QDesignerFormEditorInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_Initialize(ptr.Pointer(), PointerFromQDesignerFormEditorInterface(formEditor))
	}
}

func (ptr *QDesignerCustomWidgetInterface) InitializeDefault(formEditor QDesignerFormEditorInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_InitializeDefault(ptr.Pointer(), PointerFromQDesignerFormEditorInterface(formEditor))
	}
}

//export callbackQDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface
func callbackQDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerCustomWidgetInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerCustomWidgetInterfaceFromPointer(ptr).DestroyQDesignerCustomWidgetInterfaceDefault()
	}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectDestroyQDesignerCustomWidgetInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerCustomWidgetInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetInterface", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectDestroyQDesignerCustomWidgetInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerCustomWidgetInterface")
	}
}

func (ptr *QDesignerCustomWidgetInterface) DestroyQDesignerCustomWidgetInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerCustomWidgetInterface) DestroyQDesignerCustomWidgetInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerCustomWidgetInterface_DestroyQDesignerCustomWidgetInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerCustomWidgetInterface_Icon
func callbackQDesignerCustomWidgetInterface_Icon(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "icon"); signal != nil {
		return gui.PointerFromQIcon(signal.(func() *gui.QIcon)())
	}

	return gui.PointerFromQIcon(gui.NewQIcon())
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIcon(f func() *gui.QIcon) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "icon"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "icon", func() *gui.QIcon {
				signal.(func() *gui.QIcon)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "icon", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIcon() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "icon")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQIconFromPointer(C.QDesignerCustomWidgetInterface_Icon(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerCustomWidgetInterface_CodeTemplate
func callbackQDesignerCustomWidgetInterface_CodeTemplate(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "codeTemplate"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQDesignerCustomWidgetInterfaceFromPointer(ptr).CodeTemplateDefault()
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectCodeTemplate(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "codeTemplate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "codeTemplate", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "codeTemplate", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectCodeTemplate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "codeTemplate")
	}
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplate() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_CodeTemplate(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerCustomWidgetInterface) CodeTemplateDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_CodeTemplateDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_DomXml
func callbackQDesignerCustomWidgetInterface_DomXml(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "domXml"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQDesignerCustomWidgetInterfaceFromPointer(ptr).DomXmlDefault()
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectDomXml(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "domXml"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "domXml", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "domXml", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectDomXml() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "domXml")
	}
}

func (ptr *QDesignerCustomWidgetInterface) DomXml() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_DomXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerCustomWidgetInterface) DomXmlDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_DomXmlDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_Group
func callbackQDesignerCustomWidgetInterface_Group(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "group"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectGroup(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "group"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "group", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "group", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "group")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Group() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_Group(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_IncludeFile
func callbackQDesignerCustomWidgetInterface_IncludeFile(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "includeFile"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIncludeFile(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "includeFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "includeFile", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "includeFile", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIncludeFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "includeFile")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IncludeFile() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_IncludeFile(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_Name
func callbackQDesignerCustomWidgetInterface_Name(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "name", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QDesignerCustomWidgetInterface) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_ToolTip
func callbackQDesignerCustomWidgetInterface_ToolTip(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "toolTip"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectToolTip(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "toolTip"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "toolTip", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "toolTip", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectToolTip() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "toolTip")
	}
}

func (ptr *QDesignerCustomWidgetInterface) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_WhatsThis
func callbackQDesignerCustomWidgetInterface_WhatsThis(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "whatsThis"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerCustomWidgetInterface) ConnectWhatsThis(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "whatsThis"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "whatsThis", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "whatsThis", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectWhatsThis() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "whatsThis")
	}
}

func (ptr *QDesignerCustomWidgetInterface) WhatsThis() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerCustomWidgetInterface_WhatsThis(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerCustomWidgetInterface_IsContainer
func callbackQDesignerCustomWidgetInterface_IsContainer(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isContainer"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsContainer(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isContainer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isContainer", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isContainer", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsContainer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isContainer")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IsContainer() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsContainer(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerCustomWidgetInterface_IsInitialized
func callbackQDesignerCustomWidgetInterface_IsInitialized(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isInitialized"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerCustomWidgetInterfaceFromPointer(ptr).IsInitializedDefault())))
}

func (ptr *QDesignerCustomWidgetInterface) ConnectIsInitialized(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isInitialized"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isInitialized", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isInitialized", f)
		}
	}
}

func (ptr *QDesignerCustomWidgetInterface) DisconnectIsInitialized() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isInitialized")
	}
}

func (ptr *QDesignerCustomWidgetInterface) IsInitialized() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerCustomWidgetInterface) IsInitializedDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerCustomWidgetInterface_IsInitializedDefault(ptr.Pointer()) != 0
	}
	return false
}

type QDesignerDynamicPropertySheetExtension struct {
	ptr unsafe.Pointer
}

type QDesignerDynamicPropertySheetExtension_ITF interface {
	QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension
}

func (ptr *QDesignerDynamicPropertySheetExtension) QDesignerDynamicPropertySheetExtension_PTR() *QDesignerDynamicPropertySheetExtension {
	return ptr
}

func (ptr *QDesignerDynamicPropertySheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerDynamicPropertySheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "removeDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectRemoveDynamicProperty(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeDynamicProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeDynamicProperty", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeDynamicProperty", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectRemoveDynamicProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) RemoveDynamicProperty(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_RemoveDynamicProperty(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_AddDynamicProperty(ptr unsafe.Pointer, propertyName C.struct_QtDesigner_PackedString, value unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "addDynamicProperty"); signal != nil {
		return C.int(int32(signal.(func(string, *core.QVariant) int)(cGoUnpackString(propertyName), core.NewQVariantFromPointer(value))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectAddDynamicProperty(f func(propertyName string, value *core.QVariant) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addDynamicProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addDynamicProperty", func(propertyName string, value *core.QVariant) int {
				signal.(func(string, *core.QVariant) int)(propertyName, value)
				return f(propertyName, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addDynamicProperty", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectAddDynamicProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) AddDynamicProperty(propertyName string, value core.QVariant_ITF) int {
	if ptr.Pointer() != nil {
		var propertyNameC *C.char
		if propertyName != "" {
			propertyNameC = C.CString(propertyName)
			defer C.free(unsafe.Pointer(propertyNameC))
		}
		return int(int32(C.QDesignerDynamicPropertySheetExtension_AddDynamicProperty(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: propertyNameC, len: C.longlong(len(propertyName))}, core.PointerFromQVariant(value))))
	}
	return 0
}

//export callbackQDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension
func callbackQDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerDynamicPropertySheetExtension"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerDynamicPropertySheetExtensionFromPointer(ptr).DestroyQDesignerDynamicPropertySheetExtensionDefault()
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectDestroyQDesignerDynamicPropertySheetExtension(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerDynamicPropertySheetExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerDynamicPropertySheetExtension", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerDynamicPropertySheetExtension", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectDestroyQDesignerDynamicPropertySheetExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerDynamicPropertySheetExtension")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DestroyQDesignerDynamicPropertySheetExtension() {
	if ptr.Pointer() != nil {
		C.QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtension(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DestroyQDesignerDynamicPropertySheetExtensionDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerDynamicPropertySheetExtension_DestroyQDesignerDynamicPropertySheetExtensionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(ptr unsafe.Pointer, propertyName C.struct_QtDesigner_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "canAddDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(propertyName)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectCanAddDynamicProperty(f func(propertyName string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "canAddDynamicProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canAddDynamicProperty", func(propertyName string) bool {
				signal.(func(string) bool)(propertyName)
				return f(propertyName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canAddDynamicProperty", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectCanAddDynamicProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "canAddDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) CanAddDynamicProperty(propertyName string) bool {
	if ptr.Pointer() != nil {
		var propertyNameC *C.char
		if propertyName != "" {
			propertyNameC = C.CString(propertyName)
			defer C.free(unsafe.Pointer(propertyNameC))
		}
		return C.QDesignerDynamicPropertySheetExtension_CanAddDynamicProperty(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: propertyNameC, len: C.longlong(len(propertyName))}) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed
func callbackQDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dynamicPropertiesAllowed"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectDynamicPropertiesAllowed(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "dynamicPropertiesAllowed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dynamicPropertiesAllowed", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dynamicPropertiesAllowed", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectDynamicPropertiesAllowed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "dynamicPropertiesAllowed")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DynamicPropertiesAllowed() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_DynamicPropertiesAllowed(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty
func callbackQDesignerDynamicPropertySheetExtension_IsDynamicProperty(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isDynamicProperty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerDynamicPropertySheetExtension) ConnectIsDynamicProperty(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isDynamicProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isDynamicProperty", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isDynamicProperty", f)
		}
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) DisconnectIsDynamicProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isDynamicProperty")
	}
}

func (ptr *QDesignerDynamicPropertySheetExtension) IsDynamicProperty(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerDynamicPropertySheetExtension_IsDynamicProperty(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
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

func NewQDesignerFormEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormEditorInterface {
	var n = new(QDesignerFormEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerFormEditorInterface(parent core.QObject_ITF) *QDesignerFormEditorInterface {
	var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_NewQDesignerFormEditorInterface(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDesignerFormEditorInterface) SetActionEditor(actionEditor QDesignerActionEditorInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetActionEditor(ptr.Pointer(), PointerFromQDesignerActionEditorInterface(actionEditor))
	}
}

func (ptr *QDesignerFormEditorInterface) SetObjectInspector(objectInspector QDesignerObjectInspectorInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetObjectInspector(ptr.Pointer(), PointerFromQDesignerObjectInspectorInterface(objectInspector))
	}
}

func (ptr *QDesignerFormEditorInterface) SetPropertyEditor(propertyEditor QDesignerPropertyEditorInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetPropertyEditor(ptr.Pointer(), PointerFromQDesignerPropertyEditorInterface(propertyEditor))
	}
}

func (ptr *QDesignerFormEditorInterface) SetWidgetBox(widgetBox QDesignerWidgetBoxInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_SetWidgetBox(ptr.Pointer(), PointerFromQDesignerWidgetBoxInterface(widgetBox))
	}
}

//export callbackQDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface
func callbackQDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerFormEditorInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).DestroyQDesignerFormEditorInterfaceDefault()
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectDestroyQDesignerFormEditorInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerFormEditorInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormEditorInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormEditorInterface", f)
		}
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectDestroyQDesignerFormEditorInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerFormEditorInterface")
	}
}

func (ptr *QDesignerFormEditorInterface) DestroyQDesignerFormEditorInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerFormEditorInterface) DestroyQDesignerFormEditorInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DestroyQDesignerFormEditorInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerFormEditorInterface) ActionEditor() *QDesignerActionEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerActionEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_ActionEditor(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) FormWindowManager() *QDesignerFormWindowManagerInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowManagerInterfaceFromPointer(C.QDesignerFormEditorInterface_FormWindowManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) ObjectInspector() *QDesignerObjectInspectorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerObjectInspectorInterfaceFromPointer(C.QDesignerFormEditorInterface_ObjectInspector(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) PropertyEditor() *QDesignerPropertyEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerPropertyEditorInterfaceFromPointer(C.QDesignerFormEditorInterface_PropertyEditor(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) WidgetBox() *QDesignerWidgetBoxInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerWidgetBoxInterfaceFromPointer(C.QDesignerFormEditorInterface_WidgetBox(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) ExtensionManager() *QExtensionManager {
	if ptr.Pointer() != nil {
		var tmpValue = NewQExtensionManagerFromPointer(C.QDesignerFormEditorInterface_ExtensionManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) TopLevel() *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormEditorInterface_TopLevel(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __setOptionsPages_optionsPages_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___setOptionsPages_optionsPages_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __optionsPages_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___optionsPages_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormEditorInterface___pluginInstances_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___pluginInstances_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __pluginInstances_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___pluginInstances_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerFormEditorInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormEditorInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormEditorInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormEditorInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormEditorInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormEditorInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormEditorInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormEditorInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormEditorInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerFormEditorInterface_Event
func callbackQDesignerFormEditorInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormEditorInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDesignerFormEditorInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDesignerFormEditorInterface_EventFilter
func callbackQDesignerFormEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormEditorInterface_ChildEvent
func callbackQDesignerFormEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_ConnectNotify
func callbackQDesignerFormEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormEditorInterface_CustomEvent
func callbackQDesignerFormEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_DeleteLater
func callbackQDesignerFormEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormEditorInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerFormEditorInterface_Destroyed
func callbackQDesignerFormEditorInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerFormEditorInterface_DisconnectNotify
func callbackQDesignerFormEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormEditorInterface_ObjectNameChanged
func callbackQDesignerFormEditorInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerFormEditorInterface_TimerEvent
func callbackQDesignerFormEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormEditorInterface_MetaObject
func callbackQDesignerFormEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormEditorInterface) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormEditorInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDesignerFormWindowCursorInterface struct {
	ptr unsafe.Pointer
}

type QDesignerFormWindowCursorInterface_ITF interface {
	QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface
}

func (ptr *QDesignerFormWindowCursorInterface) QDesignerFormWindowCursorInterface_PTR() *QDesignerFormWindowCursorInterface {
	return ptr
}

func (ptr *QDesignerFormWindowCursorInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerFormWindowCursorInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QDesignerFormWindowCursorInterface__MoveMode
//QDesignerFormWindowCursorInterface::MoveMode
type QDesignerFormWindowCursorInterface__MoveMode int64

const (
	QDesignerFormWindowCursorInterface__MoveAnchor QDesignerFormWindowCursorInterface__MoveMode = QDesignerFormWindowCursorInterface__MoveMode(0)
	QDesignerFormWindowCursorInterface__KeepAnchor QDesignerFormWindowCursorInterface__MoveMode = QDesignerFormWindowCursorInterface__MoveMode(1)
)

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

//export callbackQDesignerFormWindowCursorInterface_MovePosition
func callbackQDesignerFormWindowCursorInterface_MovePosition(ptr unsafe.Pointer, operation C.longlong, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "movePosition"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QDesignerFormWindowCursorInterface__MoveOperation, QDesignerFormWindowCursorInterface__MoveMode) bool)(QDesignerFormWindowCursorInterface__MoveOperation(operation), QDesignerFormWindowCursorInterface__MoveMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectMovePosition(f func(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "movePosition"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "movePosition", func(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool {
				signal.(func(QDesignerFormWindowCursorInterface__MoveOperation, QDesignerFormWindowCursorInterface__MoveMode) bool)(operation, mode)
				return f(operation, mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "movePosition", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectMovePosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "movePosition")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) MovePosition(operation QDesignerFormWindowCursorInterface__MoveOperation, mode QDesignerFormWindowCursorInterface__MoveMode) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_MovePosition(ptr.Pointer(), C.longlong(operation), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty
func callbackQDesignerFormWindowCursorInterface_ResetWidgetProperty(ptr unsafe.Pointer, widget unsafe.Pointer, name C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "resetWidgetProperty"); signal != nil {
		signal.(func(*widgets.QWidget, string))(widgets.NewQWidgetFromPointer(widget), cGoUnpackString(name))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectResetWidgetProperty(f func(widget *widgets.QWidget, name string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resetWidgetProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resetWidgetProperty", func(widget *widgets.QWidget, name string) {
				signal.(func(*widgets.QWidget, string))(widget, name)
				f(widget, name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resetWidgetProperty", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectResetWidgetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resetWidgetProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) ResetWidgetProperty(widget widgets.QWidget_ITF, name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDesignerFormWindowCursorInterface_ResetWidgetProperty(ptr.Pointer(), widgets.PointerFromQWidget(widget), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQDesignerFormWindowCursorInterface_SetPosition
func callbackQDesignerFormWindowCursorInterface_SetPosition(ptr unsafe.Pointer, position C.int, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "setPosition"); signal != nil {
		signal.(func(int, QDesignerFormWindowCursorInterface__MoveMode))(int(int32(position)), QDesignerFormWindowCursorInterface__MoveMode(mode))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetPosition(f func(position int, mode QDesignerFormWindowCursorInterface__MoveMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPosition"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPosition", func(position int, mode QDesignerFormWindowCursorInterface__MoveMode) {
				signal.(func(int, QDesignerFormWindowCursorInterface__MoveMode))(position, mode)
				f(position, mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPosition", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetPosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPosition")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetPosition(position int, mode QDesignerFormWindowCursorInterface__MoveMode) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowCursorInterface_SetPosition(ptr.Pointer(), C.int(int32(position)), C.longlong(mode))
	}
}

//export callbackQDesignerFormWindowCursorInterface_SetProperty
func callbackQDesignerFormWindowCursorInterface_SetProperty(ptr unsafe.Pointer, name C.struct_QtDesigner_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		signal.(func(string, *core.QVariant))(cGoUnpackString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetProperty(f func(name string, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(name string, value *core.QVariant) {
				signal.(func(string, *core.QVariant))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetProperty(name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDesignerFormWindowCursorInterface_SetProperty(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerFormWindowCursorInterface_SetWidgetProperty
func callbackQDesignerFormWindowCursorInterface_SetWidgetProperty(ptr unsafe.Pointer, widget unsafe.Pointer, name C.struct_QtDesigner_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setWidgetProperty"); signal != nil {
		signal.(func(*widgets.QWidget, string, *core.QVariant))(widgets.NewQWidgetFromPointer(widget), cGoUnpackString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSetWidgetProperty(f func(widget *widgets.QWidget, name string, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWidgetProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setWidgetProperty", func(widget *widgets.QWidget, name string, value *core.QVariant) {
				signal.(func(*widgets.QWidget, string, *core.QVariant))(widget, name, value)
				f(widget, name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWidgetProperty", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSetWidgetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWidgetProperty")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SetWidgetProperty(widget widgets.QWidget_ITF, name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDesignerFormWindowCursorInterface_SetWidgetProperty(ptr.Pointer(), widgets.PointerFromQWidget(widget), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface
func callbackQDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerFormWindowCursorInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowCursorInterfaceFromPointer(ptr).DestroyQDesignerFormWindowCursorInterfaceDefault()
	}
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectDestroyQDesignerFormWindowCursorInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerFormWindowCursorInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowCursorInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowCursorInterface", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectDestroyQDesignerFormWindowCursorInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerFormWindowCursorInterface")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DestroyQDesignerFormWindowCursorInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DestroyQDesignerFormWindowCursorInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowCursorInterface_DestroyQDesignerFormWindowCursorInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerFormWindowCursorInterface_FormWindow
func callbackQDesignerFormWindowCursorInterface_FormWindow(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "formWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func() *QDesignerFormWindowInterface)())
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectFormWindow(f func() *QDesignerFormWindowInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "formWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindow", func() *QDesignerFormWindowInterface {
				signal.(func() *QDesignerFormWindowInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "formWindow")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) FormWindow() *QDesignerFormWindowInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowCursorInterface_FormWindow(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_Current
func callbackQDesignerFormWindowCursorInterface_Current(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "current"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func() *widgets.QWidget)())
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectCurrent(f func() *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "current"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "current", func() *widgets.QWidget {
				signal.(func() *widgets.QWidget)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "current", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectCurrent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "current")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Current() *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_Current(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_SelectedWidget
func callbackQDesignerFormWindowCursorInterface_SelectedWidget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "selectedWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidget(f func(index int) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectedWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectedWidget", func(index int) *widgets.QWidget {
				signal.(func(int) *widgets.QWidget)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedWidget", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectedWidget")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidget(index int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_SelectedWidget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_Widget
func callbackQDesignerFormWindowCursorInterface_Widget(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(int) *widgets.QWidget)(int(int32(index))))
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidget(f func(index int) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "widget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widget", func(index int) *widgets.QWidget {
				signal.(func(int) *widgets.QWidget)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widget", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "widget")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Widget(index int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowCursorInterface_Widget(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowCursorInterface_HasSelection
func callbackQDesignerFormWindowCursorInterface_HasSelection(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasSelection"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectHasSelection(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasSelection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasSelection", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasSelection", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectHasSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasSelection")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) HasSelection() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowCursorInterface) IsWidgetSelected(widget widgets.QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowCursorInterface_IsWidgetSelected(ptr.Pointer(), widgets.PointerFromQWidget(widget)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowCursorInterface_Position
func callbackQDesignerFormWindowCursorInterface_Position(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "position"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectPosition(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "position"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "position", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "position", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectPosition() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "position")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) Position() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_Position(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount
func callbackQDesignerFormWindowCursorInterface_SelectedWidgetCount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "selectedWidgetCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectSelectedWidgetCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectedWidgetCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectedWidgetCount", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedWidgetCount", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectSelectedWidgetCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectedWidgetCount")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) SelectedWidgetCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_SelectedWidgetCount(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerFormWindowCursorInterface_WidgetCount
func callbackQDesignerFormWindowCursorInterface_WidgetCount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "widgetCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowCursorInterface) ConnectWidgetCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "widgetCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widgetCount", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widgetCount", f)
		}
	}
}

func (ptr *QDesignerFormWindowCursorInterface) DisconnectWidgetCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "widgetCount")
	}
}

func (ptr *QDesignerFormWindowCursorInterface) WidgetCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowCursorInterface_WidgetCount(ptr.Pointer())))
	}
	return 0
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

func NewQDesignerFormWindowInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormWindowInterface {
	var n = new(QDesignerFormWindowInterface)
	n.SetPointer(ptr)
	return n
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

func QDesignerFormWindowInterface_FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {
	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow2(object core.QObject_ITF) *QDesignerFormWindowInterface {
	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow2(core.PointerFromQObject(object)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QDesignerFormWindowInterface_FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {
	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(widgets.PointerFromQWidget(widget)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDesignerFormWindowInterface) FindFormWindow(widget widgets.QWidget_ITF) *QDesignerFormWindowInterface {
	var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowInterface_QDesignerFormWindowInterface_FindFormWindow(widgets.PointerFromQWidget(widget)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDesignerFormWindowInterface_SetContents
func callbackQDesignerFormWindowInterface_SetContents(ptr unsafe.Pointer, device unsafe.Pointer, errorMessage C.struct_QtDesigner_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setContents"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QIODevice, string) bool)(core.NewQIODeviceFromPointer(device), cGoUnpackString(errorMessage)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents(f func(device *core.QIODevice, errorMessage string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setContents"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setContents", func(device *core.QIODevice, errorMessage string) bool {
				signal.(func(*core.QIODevice, string) bool)(device, errorMessage)
				return f(device, errorMessage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setContents", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setContents")
	}
}

func (ptr *QDesignerFormWindowInterface) SetContents(device core.QIODevice_ITF, errorMessage string) bool {
	if ptr.Pointer() != nil {
		var errorMessageC *C.char
		if errorMessage != "" {
			errorMessageC = C.CString(errorMessage)
			defer C.free(unsafe.Pointer(errorMessageC))
		}
		return C.QDesignerFormWindowInterface_SetContents(ptr.Pointer(), core.PointerFromQIODevice(device), C.struct_QtDesigner_PackedString{data: errorMessageC, len: C.longlong(len(errorMessage))}) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_SetContents2
func callbackQDesignerFormWindowInterface_SetContents2(ptr unsafe.Pointer, contents C.struct_QtDesigner_PackedString) C.char {
	if signal := qt.GetSignal(ptr, "setContents2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(cGoUnpackString(contents)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectSetContents2(f func(contents string) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setContents2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setContents2", func(contents string) bool {
				signal.(func(string) bool)(contents)
				return f(contents)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setContents2", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetContents2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setContents2")
	}
}

func (ptr *QDesignerFormWindowInterface) SetContents2(contents string) bool {
	if ptr.Pointer() != nil {
		var contentsC *C.char
		if contents != "" {
			contentsC = C.CString(contents)
			defer C.free(unsafe.Pointer(contentsC))
		}
		return C.QDesignerFormWindowInterface_SetContents2(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: contentsC, len: C.longlong(len(contents))}) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_AboutToUnmanageWidget
func callbackQDesignerFormWindowInterface_AboutToUnmanageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToUnmanageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectAboutToUnmanageWidget(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "aboutToUnmanageWidget") {
			C.QDesignerFormWindowInterface_ConnectAboutToUnmanageWidget(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "aboutToUnmanageWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "aboutToUnmanageWidget", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "aboutToUnmanageWidget", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAboutToUnmanageWidget() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectAboutToUnmanageWidget(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "aboutToUnmanageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) AboutToUnmanageWidget(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_AboutToUnmanageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_ActivateResourceFilePaths
func callbackQDesignerFormWindowInterface_ActivateResourceFilePaths(ptr unsafe.Pointer, paths C.struct_QtDesigner_PackedString, errorCount C.int, errorMessages C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "activateResourceFilePaths"); signal != nil {
		signal.(func([]string, int, string))(strings.Split(cGoUnpackString(paths), "|"), int(int32(errorCount)), cGoUnpackString(errorMessages))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ActivateResourceFilePathsDefault(strings.Split(cGoUnpackString(paths), "|"), int(int32(errorCount)), cGoUnpackString(errorMessages))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectActivateResourceFilePaths(f func(paths []string, errorCount int, errorMessages string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activateResourceFilePaths"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activateResourceFilePaths", func(paths []string, errorCount int, errorMessages string) {
				signal.(func([]string, int, string))(paths, errorCount, errorMessages)
				f(paths, errorCount, errorMessages)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activateResourceFilePaths", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivateResourceFilePaths() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activateResourceFilePaths")
	}
}

func (ptr *QDesignerFormWindowInterface) ActivateResourceFilePaths(paths []string, errorCount int, errorMessages string) {
	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		var errorMessagesC *C.char
		if errorMessages != "" {
			errorMessagesC = C.CString(errorMessages)
			defer C.free(unsafe.Pointer(errorMessagesC))
		}
		C.QDesignerFormWindowInterface_ActivateResourceFilePaths(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))}, C.int(int32(errorCount)), C.struct_QtDesigner_PackedString{data: errorMessagesC, len: C.longlong(len(errorMessages))})
	}
}

func (ptr *QDesignerFormWindowInterface) ActivateResourceFilePathsDefault(paths []string, errorCount int, errorMessages string) {
	if ptr.Pointer() != nil {
		var pathsC = C.CString(strings.Join(paths, "|"))
		defer C.free(unsafe.Pointer(pathsC))
		var errorMessagesC *C.char
		if errorMessages != "" {
			errorMessagesC = C.CString(errorMessages)
			defer C.free(unsafe.Pointer(errorMessagesC))
		}
		C.QDesignerFormWindowInterface_ActivateResourceFilePathsDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pathsC, len: C.longlong(len(strings.Join(paths, "|")))}, C.int(int32(errorCount)), C.struct_QtDesigner_PackedString{data: errorMessagesC, len: C.longlong(len(errorMessages))})
	}
}

//export callbackQDesignerFormWindowInterface_Activated
func callbackQDesignerFormWindowInterface_Activated(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectActivated(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activated") {
			C.QDesignerFormWindowInterface_ConnectActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activated", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activated", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activated")
	}
}

func (ptr *QDesignerFormWindowInterface) Activated(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Activated(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_AddResourceFile
func callbackQDesignerFormWindowInterface_AddResourceFile(ptr unsafe.Pointer, path C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "addResourceFile"); signal != nil {
		signal.(func(string))(cGoUnpackString(path))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectAddResourceFile(f func(path string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addResourceFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addResourceFile", func(path string) {
				signal.(func(string))(path)
				f(path)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addResourceFile", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAddResourceFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addResourceFile")
	}
}

func (ptr *QDesignerFormWindowInterface) AddResourceFile(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QDesignerFormWindowInterface_AddResourceFile(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

//export callbackQDesignerFormWindowInterface_Changed
func callbackQDesignerFormWindowInterface_Changed(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changed") {
			C.QDesignerFormWindowInterface_ConnectChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changed"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "changed", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changed", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changed")
	}
}

func (ptr *QDesignerFormWindowInterface) Changed() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_Changed(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ClearSelection
func callbackQDesignerFormWindowInterface_ClearSelection(ptr unsafe.Pointer, update C.char) {
	if signal := qt.GetSignal(ptr, "clearSelection"); signal != nil {
		signal.(func(bool))(int8(update) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectClearSelection(f func(update bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clearSelection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clearSelection", func(update bool) {
				signal.(func(bool))(update)
				f(update)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearSelection", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectClearSelection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clearSelection")
	}
}

func (ptr *QDesignerFormWindowInterface) ClearSelection(update bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ClearSelection(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(update))))
	}
}

//export callbackQDesignerFormWindowInterface_EmitSelectionChanged
func callbackQDesignerFormWindowInterface_EmitSelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "emitSelectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectEmitSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "emitSelectionChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "emitSelectionChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "emitSelectionChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectEmitSelectionChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "emitSelectionChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) EmitSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_EmitSelectionChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_FeatureChanged
func callbackQDesignerFormWindowInterface_FeatureChanged(ptr unsafe.Pointer, feature C.longlong) {
	if signal := qt.GetSignal(ptr, "featureChanged"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__FeatureFlag))(QDesignerFormWindowInterface__FeatureFlag(feature))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectFeatureChanged(f func(feature QDesignerFormWindowInterface__FeatureFlag)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "featureChanged") {
			C.QDesignerFormWindowInterface_ConnectFeatureChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "featureChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "featureChanged", func(feature QDesignerFormWindowInterface__FeatureFlag) {
				signal.(func(QDesignerFormWindowInterface__FeatureFlag))(feature)
				f(feature)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "featureChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatureChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectFeatureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "featureChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) FeatureChanged(feature QDesignerFormWindowInterface__FeatureFlag) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FeatureChanged(ptr.Pointer(), C.longlong(feature))
	}
}

//export callbackQDesignerFormWindowInterface_FileNameChanged
func callbackQDesignerFormWindowInterface_FileNameChanged(ptr unsafe.Pointer, fileName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "fileNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectFileNameChanged(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fileNameChanged") {
			C.QDesignerFormWindowInterface_ConnectFileNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fileNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileNameChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectFileNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fileNameChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) FileNameChanged(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QDesignerFormWindowInterface_FileNameChanged(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQDesignerFormWindowInterface_GeometryChanged
func callbackQDesignerFormWindowInterface_GeometryChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectGeometryChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "geometryChanged") {
			C.QDesignerFormWindowInterface_ConnectGeometryChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "geometryChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "geometryChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "geometryChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectGeometryChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "geometryChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) GeometryChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_GeometryChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_LayoutDefault
func callbackQDesignerFormWindowInterface_LayoutDefault(ptr unsafe.Pointer, margin C.int, spacing C.int) {
	if signal := qt.GetSignal(ptr, "layoutDefault"); signal != nil {
		signal.(func(int, int))(int(int32(margin)), int(int32(spacing)))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutDefault(f func(margin int, spacing int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "layoutDefault"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "layoutDefault", func(margin int, spacing int) {
				signal.(func(int, int))(margin, spacing)
				f(margin, spacing)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "layoutDefault", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "layoutDefault")
	}
}

func (ptr *QDesignerFormWindowInterface) LayoutDefault(margin int, spacing int) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LayoutDefault(ptr.Pointer(), C.int(int32(margin)), C.int(int32(spacing)))
	}
}

//export callbackQDesignerFormWindowInterface_LayoutFunction
func callbackQDesignerFormWindowInterface_LayoutFunction(ptr unsafe.Pointer, margin C.struct_QtDesigner_PackedString, spacing C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "layoutFunction"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(margin), cGoUnpackString(spacing))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectLayoutFunction(f func(margin string, spacing string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "layoutFunction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "layoutFunction", func(margin string, spacing string) {
				signal.(func(string, string))(margin, spacing)
				f(margin, spacing)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "layoutFunction", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectLayoutFunction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "layoutFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) LayoutFunction(margin string, spacing string) {
	if ptr.Pointer() != nil {
		var marginC *C.char
		if margin != "" {
			marginC = C.CString(margin)
			defer C.free(unsafe.Pointer(marginC))
		}
		var spacingC *C.char
		if spacing != "" {
			spacingC = C.CString(spacing)
			defer C.free(unsafe.Pointer(spacingC))
		}
		C.QDesignerFormWindowInterface_LayoutFunction(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: marginC, len: C.longlong(len(margin))}, C.struct_QtDesigner_PackedString{data: spacingC, len: C.longlong(len(spacing))})
	}
}

//export callbackQDesignerFormWindowInterface_MainContainerChanged
func callbackQDesignerFormWindowInterface_MainContainerChanged(ptr unsafe.Pointer, mainContainer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mainContainerChanged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(mainContainer))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectMainContainerChanged(f func(mainContainer *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "mainContainerChanged") {
			C.QDesignerFormWindowInterface_ConnectMainContainerChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "mainContainerChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "mainContainerChanged", func(mainContainer *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(mainContainer)
				f(mainContainer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "mainContainerChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectMainContainerChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectMainContainerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "mainContainerChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) MainContainerChanged(mainContainer widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MainContainerChanged(ptr.Pointer(), widgets.PointerFromQWidget(mainContainer))
	}
}

//export callbackQDesignerFormWindowInterface_ManageWidget
func callbackQDesignerFormWindowInterface_ManageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "manageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectManageWidget(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "manageWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "manageWidget", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "manageWidget", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectManageWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "manageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) ManageWidget(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ManageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_ObjectRemoved
func callbackQDesignerFormWindowInterface_ObjectRemoved(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "objectRemoved"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectObjectRemoved(f func(object *core.QObject)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "objectRemoved") {
			C.QDesignerFormWindowInterface_ConnectObjectRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "objectRemoved"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "objectRemoved", func(object *core.QObject) {
				signal.(func(*core.QObject))(object)
				f(object)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "objectRemoved", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectObjectRemoved() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectObjectRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "objectRemoved")
	}
}

func (ptr *QDesignerFormWindowInterface) ObjectRemoved(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ObjectRemoved(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQDesignerFormWindowInterface_RemoveResourceFile
func callbackQDesignerFormWindowInterface_RemoveResourceFile(ptr unsafe.Pointer, path C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "removeResourceFile"); signal != nil {
		signal.(func(string))(cGoUnpackString(path))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectRemoveResourceFile(f func(path string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeResourceFile"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeResourceFile", func(path string) {
				signal.(func(string))(path)
				f(path)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeResourceFile", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectRemoveResourceFile() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeResourceFile")
	}
}

func (ptr *QDesignerFormWindowInterface) RemoveResourceFile(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QDesignerFormWindowInterface_RemoveResourceFile(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

//export callbackQDesignerFormWindowInterface_ResourceFilesChanged
func callbackQDesignerFormWindowInterface_ResourceFilesChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resourceFilesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFilesChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "resourceFilesChanged") {
			C.QDesignerFormWindowInterface_ConnectResourceFilesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "resourceFilesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resourceFilesChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resourceFilesChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFilesChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectResourceFilesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "resourceFilesChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFilesChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ResourceFilesChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SelectWidget
func callbackQDesignerFormWindowInterface_SelectWidget(ptr unsafe.Pointer, widget unsafe.Pointer, sele C.char) {
	if signal := qt.GetSignal(ptr, "selectWidget"); signal != nil {
		signal.(func(*widgets.QWidget, bool))(widgets.NewQWidgetFromPointer(widget), int8(sele) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSelectWidget(f func(widget *widgets.QWidget, sele bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectWidget", func(widget *widgets.QWidget, sele bool) {
				signal.(func(*widgets.QWidget, bool))(widget, sele)
				f(widget, sele)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectWidget", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) SelectWidget(widget widgets.QWidget_ITF, sele bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SelectWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

//export callbackQDesignerFormWindowInterface_SelectionChanged
func callbackQDesignerFormWindowInterface_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "selectionChanged") {
			C.QDesignerFormWindowInterface_ConnectSelectionChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "selectionChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectionChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "selectionChanged")
	}
}

func (ptr *QDesignerFormWindowInterface) SelectionChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SelectionChanged(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SetAuthor
func callbackQDesignerFormWindowInterface_SetAuthor(ptr unsafe.Pointer, author C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setAuthor"); signal != nil {
		signal.(func(string))(cGoUnpackString(author))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetAuthor(f func(author string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAuthor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setAuthor", func(author string) {
				signal.(func(string))(author)
				f(author)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAuthor", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetAuthor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAuthor")
	}
}

func (ptr *QDesignerFormWindowInterface) SetAuthor(author string) {
	if ptr.Pointer() != nil {
		var authorC *C.char
		if author != "" {
			authorC = C.CString(author)
			defer C.free(unsafe.Pointer(authorC))
		}
		C.QDesignerFormWindowInterface_SetAuthor(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: authorC, len: C.longlong(len(author))})
	}
}

//export callbackQDesignerFormWindowInterface_SetComment
func callbackQDesignerFormWindowInterface_SetComment(ptr unsafe.Pointer, comment C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setComment"); signal != nil {
		signal.(func(string))(cGoUnpackString(comment))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetComment(f func(comment string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setComment"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setComment", func(comment string) {
				signal.(func(string))(comment)
				f(comment)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setComment", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setComment")
	}
}

func (ptr *QDesignerFormWindowInterface) SetComment(comment string) {
	if ptr.Pointer() != nil {
		var commentC *C.char
		if comment != "" {
			commentC = C.CString(comment)
			defer C.free(unsafe.Pointer(commentC))
		}
		C.QDesignerFormWindowInterface_SetComment(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: commentC, len: C.longlong(len(comment))})
	}
}

//export callbackQDesignerFormWindowInterface_SetDirty
func callbackQDesignerFormWindowInterface_SetDirty(ptr unsafe.Pointer, dirty C.char) {
	if signal := qt.GetSignal(ptr, "setDirty"); signal != nil {
		signal.(func(bool))(int8(dirty) != 0)
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetDirty(f func(dirty bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDirty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setDirty", func(dirty bool) {
				signal.(func(bool))(dirty)
				f(dirty)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDirty", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetDirty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDirty")
	}
}

func (ptr *QDesignerFormWindowInterface) SetDirty(dirty bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetDirty(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(dirty))))
	}
}

//export callbackQDesignerFormWindowInterface_SetExportMacro
func callbackQDesignerFormWindowInterface_SetExportMacro(ptr unsafe.Pointer, exportMacro C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setExportMacro"); signal != nil {
		signal.(func(string))(cGoUnpackString(exportMacro))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetExportMacro(f func(exportMacro string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setExportMacro"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setExportMacro", func(exportMacro string) {
				signal.(func(string))(exportMacro)
				f(exportMacro)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setExportMacro", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetExportMacro() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setExportMacro")
	}
}

func (ptr *QDesignerFormWindowInterface) SetExportMacro(exportMacro string) {
	if ptr.Pointer() != nil {
		var exportMacroC *C.char
		if exportMacro != "" {
			exportMacroC = C.CString(exportMacro)
			defer C.free(unsafe.Pointer(exportMacroC))
		}
		C.QDesignerFormWindowInterface_SetExportMacro(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: exportMacroC, len: C.longlong(len(exportMacro))})
	}
}

//export callbackQDesignerFormWindowInterface_SetFeatures
func callbackQDesignerFormWindowInterface_SetFeatures(ptr unsafe.Pointer, features C.longlong) {
	if signal := qt.GetSignal(ptr, "setFeatures"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__FeatureFlag))(QDesignerFormWindowInterface__FeatureFlag(features))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetFeatures(f func(features QDesignerFormWindowInterface__FeatureFlag)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFeatures"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFeatures", func(features QDesignerFormWindowInterface__FeatureFlag) {
				signal.(func(QDesignerFormWindowInterface__FeatureFlag))(features)
				f(features)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFeatures", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFeatures() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFeatures")
	}
}

func (ptr *QDesignerFormWindowInterface) SetFeatures(features QDesignerFormWindowInterface__FeatureFlag) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetFeatures(ptr.Pointer(), C.longlong(features))
	}
}

//export callbackQDesignerFormWindowInterface_SetFileName
func callbackQDesignerFormWindowInterface_SetFileName(ptr unsafe.Pointer, fileName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setFileName"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetFileName(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFileName")
	}
}

func (ptr *QDesignerFormWindowInterface) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QDesignerFormWindowInterface_SetFileName(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQDesignerFormWindowInterface_SetGrid
func callbackQDesignerFormWindowInterface_SetGrid(ptr unsafe.Pointer, grid unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGrid"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(grid))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetGrid(f func(grid *core.QPoint)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGrid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGrid", func(grid *core.QPoint) {
				signal.(func(*core.QPoint))(grid)
				f(grid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGrid", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetGrid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGrid")
	}
}

func (ptr *QDesignerFormWindowInterface) SetGrid(grid core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetGrid(ptr.Pointer(), core.PointerFromQPoint(grid))
	}
}

//export callbackQDesignerFormWindowInterface_SetIncludeHints
func callbackQDesignerFormWindowInterface_SetIncludeHints(ptr unsafe.Pointer, includeHints C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setIncludeHints"); signal != nil {
		signal.(func([]string))(strings.Split(cGoUnpackString(includeHints), "|"))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetIncludeHints(f func(includeHints []string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIncludeHints"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setIncludeHints", func(includeHints []string) {
				signal.(func([]string))(includeHints)
				f(includeHints)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIncludeHints", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetIncludeHints() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIncludeHints")
	}
}

func (ptr *QDesignerFormWindowInterface) SetIncludeHints(includeHints []string) {
	if ptr.Pointer() != nil {
		var includeHintsC = C.CString(strings.Join(includeHints, "|"))
		defer C.free(unsafe.Pointer(includeHintsC))
		C.QDesignerFormWindowInterface_SetIncludeHints(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: includeHintsC, len: C.longlong(len(strings.Join(includeHints, "|")))})
	}
}

//export callbackQDesignerFormWindowInterface_SetLayoutDefault
func callbackQDesignerFormWindowInterface_SetLayoutDefault(ptr unsafe.Pointer, margin C.int, spacing C.int) {
	if signal := qt.GetSignal(ptr, "setLayoutDefault"); signal != nil {
		signal.(func(int, int))(int(int32(margin)), int(int32(spacing)))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutDefault(f func(margin int, spacing int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLayoutDefault"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLayoutDefault", func(margin int, spacing int) {
				signal.(func(int, int))(margin, spacing)
				f(margin, spacing)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLayoutDefault", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutDefault() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLayoutDefault")
	}
}

func (ptr *QDesignerFormWindowInterface) SetLayoutDefault(margin int, spacing int) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetLayoutDefault(ptr.Pointer(), C.int(int32(margin)), C.int(int32(spacing)))
	}
}

//export callbackQDesignerFormWindowInterface_SetLayoutFunction
func callbackQDesignerFormWindowInterface_SetLayoutFunction(ptr unsafe.Pointer, margin C.struct_QtDesigner_PackedString, spacing C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setLayoutFunction"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(margin), cGoUnpackString(spacing))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetLayoutFunction(f func(margin string, spacing string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLayoutFunction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLayoutFunction", func(margin string, spacing string) {
				signal.(func(string, string))(margin, spacing)
				f(margin, spacing)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLayoutFunction", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetLayoutFunction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLayoutFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) SetLayoutFunction(margin string, spacing string) {
	if ptr.Pointer() != nil {
		var marginC *C.char
		if margin != "" {
			marginC = C.CString(margin)
			defer C.free(unsafe.Pointer(marginC))
		}
		var spacingC *C.char
		if spacing != "" {
			spacingC = C.CString(spacing)
			defer C.free(unsafe.Pointer(spacingC))
		}
		C.QDesignerFormWindowInterface_SetLayoutFunction(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: marginC, len: C.longlong(len(margin))}, C.struct_QtDesigner_PackedString{data: spacingC, len: C.longlong(len(spacing))})
	}
}

//export callbackQDesignerFormWindowInterface_SetMainContainer
func callbackQDesignerFormWindowInterface_SetMainContainer(ptr unsafe.Pointer, mainContainer unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setMainContainer"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(mainContainer))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetMainContainer(f func(mainContainer *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMainContainer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMainContainer", func(mainContainer *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(mainContainer)
				f(mainContainer)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMainContainer", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetMainContainer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMainContainer")
	}
}

func (ptr *QDesignerFormWindowInterface) SetMainContainer(mainContainer widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetMainContainer(ptr.Pointer(), widgets.PointerFromQWidget(mainContainer))
	}
}

//export callbackQDesignerFormWindowInterface_SetPixmapFunction
func callbackQDesignerFormWindowInterface_SetPixmapFunction(ptr unsafe.Pointer, pixmapFunction C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setPixmapFunction"); signal != nil {
		signal.(func(string))(cGoUnpackString(pixmapFunction))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetPixmapFunction(f func(pixmapFunction string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPixmapFunction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPixmapFunction", func(pixmapFunction string) {
				signal.(func(string))(pixmapFunction)
				f(pixmapFunction)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPixmapFunction", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetPixmapFunction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPixmapFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) SetPixmapFunction(pixmapFunction string) {
	if ptr.Pointer() != nil {
		var pixmapFunctionC *C.char
		if pixmapFunction != "" {
			pixmapFunctionC = C.CString(pixmapFunction)
			defer C.free(unsafe.Pointer(pixmapFunctionC))
		}
		C.QDesignerFormWindowInterface_SetPixmapFunction(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pixmapFunctionC, len: C.longlong(len(pixmapFunction))})
	}
}

//export callbackQDesignerFormWindowInterface_SetResourceFileSaveMode
func callbackQDesignerFormWindowInterface_SetResourceFileSaveMode(ptr unsafe.Pointer, behaviour C.longlong) {
	if signal := qt.GetSignal(ptr, "setResourceFileSaveMode"); signal != nil {
		signal.(func(QDesignerFormWindowInterface__ResourceFileSaveMode))(QDesignerFormWindowInterface__ResourceFileSaveMode(behaviour))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectSetResourceFileSaveMode(f func(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setResourceFileSaveMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setResourceFileSaveMode", func(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode) {
				signal.(func(QDesignerFormWindowInterface__ResourceFileSaveMode))(behaviour)
				f(behaviour)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setResourceFileSaveMode", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectSetResourceFileSaveMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setResourceFileSaveMode")
	}
}

func (ptr *QDesignerFormWindowInterface) SetResourceFileSaveMode(behaviour QDesignerFormWindowInterface__ResourceFileSaveMode) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetResourceFileSaveMode(ptr.Pointer(), C.longlong(behaviour))
	}
}

//export callbackQDesignerFormWindowInterface_UnmanageWidget
func callbackQDesignerFormWindowInterface_UnmanageWidget(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "unmanageWidget"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectUnmanageWidget(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unmanageWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unmanageWidget", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unmanageWidget", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectUnmanageWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unmanageWidget")
	}
}

func (ptr *QDesignerFormWindowInterface) UnmanageWidget(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UnmanageWidget(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetManaged
func callbackQDesignerFormWindowInterface_WidgetManaged(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "widgetManaged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetManaged(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "widgetManaged") {
			C.QDesignerFormWindowInterface_ConnectWidgetManaged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "widgetManaged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widgetManaged", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widgetManaged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetManaged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetManaged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "widgetManaged")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetManaged(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetManaged(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetRemoved
func callbackQDesignerFormWindowInterface_WidgetRemoved(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "widgetRemoved"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetRemoved(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "widgetRemoved") {
			C.QDesignerFormWindowInterface_ConnectWidgetRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "widgetRemoved"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widgetRemoved", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widgetRemoved", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetRemoved() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "widgetRemoved")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetRemoved(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetRemoved(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_WidgetUnmanaged
func callbackQDesignerFormWindowInterface_WidgetUnmanaged(ptr unsafe.Pointer, widget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "widgetUnmanaged"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(widget))
	}

}

func (ptr *QDesignerFormWindowInterface) ConnectWidgetUnmanaged(f func(widget *widgets.QWidget)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "widgetUnmanaged") {
			C.QDesignerFormWindowInterface_ConnectWidgetUnmanaged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "widgetUnmanaged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "widgetUnmanaged", func(widget *widgets.QWidget) {
				signal.(func(*widgets.QWidget))(widget)
				f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "widgetUnmanaged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectWidgetUnmanaged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectWidgetUnmanaged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "widgetUnmanaged")
	}
}

func (ptr *QDesignerFormWindowInterface) WidgetUnmanaged(widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WidgetUnmanaged(ptr.Pointer(), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface
func callbackQDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerFormWindowInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DestroyQDesignerFormWindowInterfaceDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectDestroyQDesignerFormWindowInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerFormWindowInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowInterface", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectDestroyQDesignerFormWindowInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerFormWindowInterface")
	}
}

func (ptr *QDesignerFormWindowInterface) DestroyQDesignerFormWindowInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DestroyQDesignerFormWindowInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerFormWindowInterface) DestroyQDesignerFormWindowInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DestroyQDesignerFormWindowInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerFormWindowInterface_Features
func callbackQDesignerFormWindowInterface_Features(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "features"); signal != nil {
		return C.longlong(signal.(func() QDesignerFormWindowInterface__FeatureFlag)())
	}

	return C.longlong(0)
}

func (ptr *QDesignerFormWindowInterface) ConnectFeatures(f func() QDesignerFormWindowInterface__FeatureFlag) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "features"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "features", func() QDesignerFormWindowInterface__FeatureFlag {
				signal.(func() QDesignerFormWindowInterface__FeatureFlag)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "features", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFeatures() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "features")
	}
}

func (ptr *QDesignerFormWindowInterface) Features() QDesignerFormWindowInterface__FeatureFlag {
	if ptr.Pointer() != nil {
		return QDesignerFormWindowInterface__FeatureFlag(C.QDesignerFormWindowInterface_Features(ptr.Pointer()))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_Core
func callbackQDesignerFormWindowInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerFormWindowInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerFormWindowInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "core"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "core", func() *QDesignerFormEditorInterface {
				signal.(func() *QDesignerFormEditorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "core", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "core")
	}
}

func (ptr *QDesignerFormWindowInterface) Core() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) CoreDefault() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_Cursor
func callbackQDesignerFormWindowInterface_Cursor(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "cursor"); signal != nil {
		return PointerFromQDesignerFormWindowCursorInterface(signal.(func() *QDesignerFormWindowCursorInterface)())
	}

	return PointerFromQDesignerFormWindowCursorInterface(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectCursor(f func() *QDesignerFormWindowCursorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "cursor"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "cursor", func() *QDesignerFormWindowCursorInterface {
				signal.(func() *QDesignerFormWindowCursorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "cursor", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCursor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "cursor")
	}
}

func (ptr *QDesignerFormWindowInterface) Cursor() *QDesignerFormWindowCursorInterface {
	if ptr.Pointer() != nil {
		return NewQDesignerFormWindowCursorInterfaceFromPointer(C.QDesignerFormWindowInterface_Cursor(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_AbsoluteDir
func callbackQDesignerFormWindowInterface_AbsoluteDir(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "absoluteDir"); signal != nil {
		return core.PointerFromQDir(signal.(func() *core.QDir)())
	}

	return core.PointerFromQDir(core.NewQDir(nil))
}

func (ptr *QDesignerFormWindowInterface) ConnectAbsoluteDir(f func() *core.QDir) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "absoluteDir"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "absoluteDir", func() *core.QDir {
				signal.(func() *core.QDir)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "absoluteDir", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAbsoluteDir() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "absoluteDir")
	}
}

func (ptr *QDesignerFormWindowInterface) AbsoluteDir() *core.QDir {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDirFromPointer(C.QDesignerFormWindowInterface_AbsoluteDir(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_Grid
func callbackQDesignerFormWindowInterface_Grid(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "grid"); signal != nil {
		return core.PointerFromQPoint(signal.(func() *core.QPoint)())
	}

	return core.PointerFromQPoint(core.NewQPoint())
}

func (ptr *QDesignerFormWindowInterface) ConnectGrid(f func() *core.QPoint) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "grid"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grid", func() *core.QPoint {
				signal.(func() *core.QPoint)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grid", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectGrid() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "grid")
	}
}

func (ptr *QDesignerFormWindowInterface) Grid() *core.QPoint {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQPointFromPointer(C.QDesignerFormWindowInterface_Grid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_Author
func callbackQDesignerFormWindowInterface_Author(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "author"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectAuthor(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "author"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "author", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "author", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectAuthor() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "author")
	}
}

func (ptr *QDesignerFormWindowInterface) Author() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_Author(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_Comment
func callbackQDesignerFormWindowInterface_Comment(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "comment"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectComment(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "comment"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "comment", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "comment", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectComment() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "comment")
	}
}

func (ptr *QDesignerFormWindowInterface) Comment() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_Comment(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_Contents
func callbackQDesignerFormWindowInterface_Contents(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "contents"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectContents(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "contents"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "contents", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "contents", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "contents")
	}
}

func (ptr *QDesignerFormWindowInterface) Contents() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_Contents(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_ExportMacro
func callbackQDesignerFormWindowInterface_ExportMacro(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "exportMacro"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectExportMacro(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exportMacro"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "exportMacro", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exportMacro", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectExportMacro() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exportMacro")
	}
}

func (ptr *QDesignerFormWindowInterface) ExportMacro() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_ExportMacro(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_FileName
func callbackQDesignerFormWindowInterface_FileName(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "fileName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectFileName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fileName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileName", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fileName")
	}
}

func (ptr *QDesignerFormWindowInterface) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_FileName(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerFormWindowInterface_PixmapFunction
func callbackQDesignerFormWindowInterface_PixmapFunction(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "pixmapFunction"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerFormWindowInterface) ConnectPixmapFunction(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pixmapFunction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pixmapFunction", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pixmapFunction", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectPixmapFunction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pixmapFunction")
	}
}

func (ptr *QDesignerFormWindowInterface) PixmapFunction() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerFormWindowInterface_PixmapFunction(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerFormWindowInterface) ActiveResourceFilePaths() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDesignerFormWindowInterface_ActiveResourceFilePaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_CheckContents
func callbackQDesignerFormWindowInterface_CheckContents(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "checkContents"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QDesignerFormWindowInterface) ConnectCheckContents(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "checkContents"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "checkContents", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "checkContents", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectCheckContents() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "checkContents")
	}
}

func (ptr *QDesignerFormWindowInterface) CheckContents() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDesignerFormWindowInterface_CheckContents(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_IncludeHints
func callbackQDesignerFormWindowInterface_IncludeHints(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "includeHints"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QDesignerFormWindowInterface) ConnectIncludeHints(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "includeHints"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "includeHints", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "includeHints", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIncludeHints() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "includeHints")
	}
}

func (ptr *QDesignerFormWindowInterface) IncludeHints() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDesignerFormWindowInterface_IncludeHints(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_ResourceFiles
func callbackQDesignerFormWindowInterface_ResourceFiles(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "resourceFiles"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_QtDesigner_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFiles(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resourceFiles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resourceFiles", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resourceFiles", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFiles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resourceFiles")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFiles() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QDesignerFormWindowInterface_ResourceFiles(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQDesignerFormWindowInterface_FormContainer
func callbackQDesignerFormWindowInterface_FormContainer(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "formContainer"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func() *widgets.QWidget)())
	}

	return widgets.PointerFromQWidget(nil)
}

func (ptr *QDesignerFormWindowInterface) ConnectFormContainer(f func() *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "formContainer"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formContainer", func() *widgets.QWidget {
				signal.(func() *widgets.QWidget)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formContainer", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectFormContainer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "formContainer")
	}
}

func (ptr *QDesignerFormWindowInterface) FormContainer() *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowInterface_FormContainer(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_ResourceFileSaveMode
func callbackQDesignerFormWindowInterface_ResourceFileSaveMode(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "resourceFileSaveMode"); signal != nil {
		return C.longlong(signal.(func() QDesignerFormWindowInterface__ResourceFileSaveMode)())
	}

	return C.longlong(0)
}

func (ptr *QDesignerFormWindowInterface) ConnectResourceFileSaveMode(f func() QDesignerFormWindowInterface__ResourceFileSaveMode) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resourceFileSaveMode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resourceFileSaveMode", func() QDesignerFormWindowInterface__ResourceFileSaveMode {
				signal.(func() QDesignerFormWindowInterface__ResourceFileSaveMode)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resourceFileSaveMode", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectResourceFileSaveMode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resourceFileSaveMode")
	}
}

func (ptr *QDesignerFormWindowInterface) ResourceFileSaveMode() QDesignerFormWindowInterface__ResourceFileSaveMode {
	if ptr.Pointer() != nil {
		return QDesignerFormWindowInterface__ResourceFileSaveMode(C.QDesignerFormWindowInterface_ResourceFileSaveMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_HasFeature
func callbackQDesignerFormWindowInterface_HasFeature(ptr unsafe.Pointer, feature C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QDesignerFormWindowInterface__FeatureFlag) bool)(QDesignerFormWindowInterface__FeatureFlag(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectHasFeature(f func(feature QDesignerFormWindowInterface__FeatureFlag) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasFeature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", func(feature QDesignerFormWindowInterface__FeatureFlag) bool {
				signal.(func(QDesignerFormWindowInterface__FeatureFlag) bool)(feature)
				return f(feature)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasFeature", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectHasFeature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasFeature")
	}
}

func (ptr *QDesignerFormWindowInterface) HasFeature(feature QDesignerFormWindowInterface__FeatureFlag) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_HasFeature(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_IsDirty
func callbackQDesignerFormWindowInterface_IsDirty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isDirty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectIsDirty(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isDirty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isDirty", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isDirty", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsDirty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isDirty")
	}
}

func (ptr *QDesignerFormWindowInterface) IsDirty() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_IsDirty(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_IsManaged
func callbackQDesignerFormWindowInterface_IsManaged(ptr unsafe.Pointer, widget unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isManaged"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QWidget) bool)(widgets.NewQWidgetFromPointer(widget)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerFormWindowInterface) ConnectIsManaged(f func(widget *widgets.QWidget) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isManaged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isManaged", func(widget *widgets.QWidget) bool {
				signal.(func(*widgets.QWidget) bool)(widget)
				return f(widget)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isManaged", f)
		}
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectIsManaged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isManaged")
	}
}

func (ptr *QDesignerFormWindowInterface) IsManaged(widget widgets.QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_IsManaged(ptr.Pointer(), widgets.PointerFromQWidget(widget)) != 0
	}
	return false
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_atList(i int) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QDesignerFormWindowInterface___simplifySelection_widgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_setList(i widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___simplifySelection_widgets_setList(ptr.Pointer(), widgets.PointerFromQWidget(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __simplifySelection_widgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___simplifySelection_widgets_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerFormWindowInterface___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerFormWindowInterface___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerFormWindowInterface___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerFormWindowInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerFormWindowInterface_Close
func callbackQDesignerFormWindowInterface_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerFormWindowInterface) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_Event
func callbackQDesignerFormWindowInterface_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormWindowInterface) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_FocusNextPrevChild
func callbackQDesignerFormWindowInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerFormWindowInterface) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_NativeEvent
func callbackQDesignerFormWindowInterface_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QDesignerFormWindowInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_ActionEvent
func callbackQDesignerFormWindowInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ChangeEvent
func callbackQDesignerFormWindowInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_CloseEvent
func callbackQDesignerFormWindowInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ContextMenuEvent
func callbackQDesignerFormWindowInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_CustomContextMenuRequested
func callbackQDesignerFormWindowInterface_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQDesignerFormWindowInterface_DragEnterEvent
func callbackQDesignerFormWindowInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DragLeaveEvent
func callbackQDesignerFormWindowInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DragMoveEvent
func callbackQDesignerFormWindowInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DropEvent
func callbackQDesignerFormWindowInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_EnterEvent
func callbackQDesignerFormWindowInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_FocusInEvent
func callbackQDesignerFormWindowInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_FocusOutEvent
func callbackQDesignerFormWindowInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Hide
func callbackQDesignerFormWindowInterface_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) HideDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_HideEvent
func callbackQDesignerFormWindowInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_InputMethodEvent
func callbackQDesignerFormWindowInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_KeyPressEvent
func callbackQDesignerFormWindowInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_KeyReleaseEvent
func callbackQDesignerFormWindowInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_LeaveEvent
func callbackQDesignerFormWindowInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Lower
func callbackQDesignerFormWindowInterface_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_MouseDoubleClickEvent
func callbackQDesignerFormWindowInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MouseMoveEvent
func callbackQDesignerFormWindowInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MousePressEvent
func callbackQDesignerFormWindowInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MouseReleaseEvent
func callbackQDesignerFormWindowInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MoveEvent
func callbackQDesignerFormWindowInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_PaintEvent
func callbackQDesignerFormWindowInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Raise
func callbackQDesignerFormWindowInterface_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_Repaint
func callbackQDesignerFormWindowInterface_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ResizeEvent
func callbackQDesignerFormWindowInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_SetDisabled
func callbackQDesignerFormWindowInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerFormWindowInterface_SetEnabled
func callbackQDesignerFormWindowInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerFormWindowInterface_SetFocus2
func callbackQDesignerFormWindowInterface_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerFormWindowInterface) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_SetHidden
func callbackQDesignerFormWindowInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerFormWindowInterface_SetStyleSheet
func callbackQDesignerFormWindowInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QDesignerFormWindowInterface) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QDesignerFormWindowInterface_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQDesignerFormWindowInterface_SetVisible
func callbackQDesignerFormWindowInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerFormWindowInterface_SetWindowModified
func callbackQDesignerFormWindowInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerFormWindowInterface_SetWindowTitle
func callbackQDesignerFormWindowInterface_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QDesignerFormWindowInterface) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QDesignerFormWindowInterface_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQDesignerFormWindowInterface_Show
func callbackQDesignerFormWindowInterface_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowEvent
func callbackQDesignerFormWindowInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ShowFullScreen
func callbackQDesignerFormWindowInterface_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowMaximized
func callbackQDesignerFormWindowInterface_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowMinimized
func callbackQDesignerFormWindowInterface_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_ShowNormal
func callbackQDesignerFormWindowInterface_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_TabletEvent
func callbackQDesignerFormWindowInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_Update
func callbackQDesignerFormWindowInterface_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_UpdateMicroFocus
func callbackQDesignerFormWindowInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowInterface_WheelEvent
func callbackQDesignerFormWindowInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_WindowIconChanged
func callbackQDesignerFormWindowInterface_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQDesignerFormWindowInterface_WindowTitleChanged
func callbackQDesignerFormWindowInterface_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQDesignerFormWindowInterface_PaintEngine
func callbackQDesignerFormWindowInterface_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQDesignerFormWindowInterfaceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QDesignerFormWindowInterface) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QDesignerFormWindowInterface_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_MinimumSizeHint
func callbackQDesignerFormWindowInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerFormWindowInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerFormWindowInterface) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_SizeHint
func callbackQDesignerFormWindowInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerFormWindowInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerFormWindowInterface) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerFormWindowInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_InputMethodQuery
func callbackQDesignerFormWindowInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerFormWindowInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerFormWindowInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerFormWindowInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowInterface_HasHeightForWidth
func callbackQDesignerFormWindowInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerFormWindowInterface) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_HeightForWidth
func callbackQDesignerFormWindowInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerFormWindowInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerFormWindowInterface) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_Metric
func callbackQDesignerFormWindowInterface_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQDesignerFormWindowInterfaceFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QDesignerFormWindowInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowInterface_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQDesignerFormWindowInterface_EventFilter
func callbackQDesignerFormWindowInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormWindowInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowInterface_ChildEvent
func callbackQDesignerFormWindowInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_ConnectNotify
func callbackQDesignerFormWindowInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowInterface_CustomEvent
func callbackQDesignerFormWindowInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_DeleteLater
func callbackQDesignerFormWindowInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormWindowInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerFormWindowInterface_Destroyed
func callbackQDesignerFormWindowInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerFormWindowInterface_DisconnectNotify
func callbackQDesignerFormWindowInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowInterface_ObjectNameChanged
func callbackQDesignerFormWindowInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerFormWindowInterface_TimerEvent
func callbackQDesignerFormWindowInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormWindowInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormWindowInterface_MetaObject
func callbackQDesignerFormWindowInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormWindowInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormWindowInterface) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDesignerFormWindowInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

func NewQDesignerFormWindowManagerInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerFormWindowManagerInterface {
	var n = new(QDesignerFormWindowManagerInterface)
	n.SetPointer(ptr)
	return n
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

//export callbackQDesignerFormWindowManagerInterface_CreateFormWindow
func callbackQDesignerFormWindowManagerInterface_CreateFormWindow(ptr unsafe.Pointer, parent unsafe.Pointer, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createFormWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func(*widgets.QWidget, core.Qt__WindowType) *QDesignerFormWindowInterface)(widgets.NewQWidgetFromPointer(parent), core.Qt__WindowType(flags)))
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreateFormWindow(f func(parent *widgets.QWidget, flags core.Qt__WindowType) *QDesignerFormWindowInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createFormWindow", func(parent *widgets.QWidget, flags core.Qt__WindowType) *QDesignerFormWindowInterface {
				signal.(func(*widgets.QWidget, core.Qt__WindowType) *QDesignerFormWindowInterface)(parent, flags)
				return f(parent, flags)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createFormWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreateFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CreateFormWindow(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerFormWindowInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_CreateFormWindow(ptr.Pointer(), widgets.PointerFromQWidget(parent), C.longlong(flags)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged
func callbackQDesignerFormWindowManagerInterface_ActiveFormWindowChanged(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "activeFormWindowChanged"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindowChanged(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activeFormWindowChanged") {
			C.QDesignerFormWindowManagerInterface_ConnectActiveFormWindowChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activeFormWindowChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeFormWindowChanged", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeFormWindowChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindowChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectActiveFormWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activeFormWindowChanged")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindowChanged(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ActiveFormWindowChanged(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_AddFormWindow
func callbackQDesignerFormWindowManagerInterface_AddFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectAddFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addFormWindow", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addFormWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAddFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) AddFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_AddFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_CloseAllPreviews
func callbackQDesignerFormWindowManagerInterface_CloseAllPreviews(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeAllPreviews"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCloseAllPreviews(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "closeAllPreviews"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "closeAllPreviews", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "closeAllPreviews", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCloseAllPreviews() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "closeAllPreviews")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CloseAllPreviews() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_CloseAllPreviews(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowAdded
func callbackQDesignerFormWindowManagerInterface_FormWindowAdded(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "formWindowAdded"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowAdded(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "formWindowAdded") {
			C.QDesignerFormWindowManagerInterface_ConnectFormWindowAdded(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "formWindowAdded"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindowAdded", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindowAdded", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowAdded() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "formWindowAdded")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowAdded(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowAdded(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowRemoved
func callbackQDesignerFormWindowManagerInterface_FormWindowRemoved(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "formWindowRemoved"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowRemoved(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "formWindowRemoved") {
			C.QDesignerFormWindowManagerInterface_ConnectFormWindowRemoved(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "formWindowRemoved"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindowRemoved", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindowRemoved", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowRemoved() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "formWindowRemoved")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowRemoved(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowRemoved(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged
func callbackQDesignerFormWindowManagerInterface_FormWindowSettingsChanged(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "formWindowSettingsChanged"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowSettingsChanged(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "formWindowSettingsChanged") {
			C.QDesignerFormWindowManagerInterface_ConnectFormWindowSettingsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "formWindowSettingsChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindowSettingsChanged", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindowSettingsChanged", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowSettingsChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectFormWindowSettingsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "formWindowSettingsChanged")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowSettingsChanged(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_FormWindowSettingsChanged(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_RemoveFormWindow
func callbackQDesignerFormWindowManagerInterface_RemoveFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "removeFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectRemoveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeFormWindow", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeFormWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectRemoveFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) RemoveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_RemoveFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow
func callbackQDesignerFormWindowManagerInterface_SetActiveFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setActiveFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectSetActiveFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setActiveFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setActiveFormWindow", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setActiveFormWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectSetActiveFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setActiveFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) SetActiveFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_SetActiveFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ShowPluginDialog
func callbackQDesignerFormWindowManagerInterface_ShowPluginDialog(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showPluginDialog"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPluginDialog(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showPluginDialog"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "showPluginDialog", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showPluginDialog", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPluginDialog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showPluginDialog")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPluginDialog() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ShowPluginDialog(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowManagerInterface_ShowPreview
func callbackQDesignerFormWindowManagerInterface_ShowPreview(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showPreview"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDesignerFormWindowManagerInterface) ConnectShowPreview(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "showPreview"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "showPreview", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "showPreview", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectShowPreview() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "showPreview")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ShowPreview() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ShowPreview(ptr.Pointer())
	}
}

//export callbackQDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface
func callbackQDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerFormWindowManagerInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).DestroyQDesignerFormWindowManagerInterfaceDefault()
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectDestroyQDesignerFormWindowManagerInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerFormWindowManagerInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowManagerInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerFormWindowManagerInterface", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectDestroyQDesignerFormWindowManagerInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerFormWindowManagerInterface")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DestroyQDesignerFormWindowManagerInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DestroyQDesignerFormWindowManagerInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DestroyQDesignerFormWindowManagerInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerFormWindowManagerInterface_Action
func callbackQDesignerFormWindowManagerInterface_Action(ptr unsafe.Pointer, action C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "action"); signal != nil {
		return widgets.PointerFromQAction(signal.(func(QDesignerFormWindowManagerInterface__Action) *widgets.QAction)(QDesignerFormWindowManagerInterface__Action(action)))
	}

	return widgets.PointerFromQAction(widgets.NewQAction(nil))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectAction(f func(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "action"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "action", func(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction {
				signal.(func(QDesignerFormWindowManagerInterface__Action) *widgets.QAction)(action)
				return f(action)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "action", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "action")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) Action(action QDesignerFormWindowManagerInterface__Action) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerFormWindowManagerInterface_Action(ptr.Pointer(), C.longlong(action)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActionGroup
func callbackQDesignerFormWindowManagerInterface_ActionGroup(ptr unsafe.Pointer, actionGroup C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "actionGroup"); signal != nil {
		return widgets.PointerFromQActionGroup(signal.(func(QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup)(QDesignerFormWindowManagerInterface__ActionGroup(actionGroup)))
	}

	return widgets.PointerFromQActionGroup(widgets.NewQActionGroup(nil))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActionGroup(f func(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "actionGroup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "actionGroup", func(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup {
				signal.(func(QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup)(actionGroup)
				return f(actionGroup)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "actionGroup", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActionGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "actionGroup")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActionGroup(actionGroup QDesignerFormWindowManagerInterface__ActionGroup) *widgets.QActionGroup {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionGroupFromPointer(C.QDesignerFormWindowManagerInterface_ActionGroup(ptr.Pointer(), C.longlong(actionGroup)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_Core
func callbackQDesignerFormWindowManagerInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerFormEditorInterface(nil))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "core"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "core", func() *QDesignerFormEditorInterface {
				signal.(func() *QDesignerFormEditorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "core", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "core")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) Core() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_ActiveFormWindow
func callbackQDesignerFormWindowManagerInterface_ActiveFormWindow(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "activeFormWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func() *QDesignerFormWindowInterface)())
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectActiveFormWindow(f func() *QDesignerFormWindowInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "activeFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "activeFormWindow", func() *QDesignerFormWindowInterface {
				signal.(func() *QDesignerFormWindowInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activeFormWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectActiveFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "activeFormWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ActiveFormWindow() *QDesignerFormWindowInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_ActiveFormWindow(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_FormWindow
func callbackQDesignerFormWindowManagerInterface_FormWindow(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "formWindow"); signal != nil {
		return PointerFromQDesignerFormWindowInterface(signal.(func(int) *QDesignerFormWindowInterface)(int(int32(index))))
	}

	return PointerFromQDesignerFormWindowInterface(nil)
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindow(f func(index int) *QDesignerFormWindowInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "formWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindow", func(index int) *QDesignerFormWindowInterface {
				signal.(func(int) *QDesignerFormWindowInterface)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindow", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "formWindow")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindow(index int) *QDesignerFormWindowInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormWindowInterfaceFromPointer(C.QDesignerFormWindowManagerInterface_FormWindow(ptr.Pointer(), C.int(int32(index))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap
func callbackQDesignerFormWindowManagerInterface_CreatePreviewPixmap(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPreviewPixmap"); signal != nil {
		return gui.PointerFromQPixmap(signal.(func() *gui.QPixmap)())
	}

	return gui.PointerFromQPixmap(gui.NewQPixmap())
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectCreatePreviewPixmap(f func() *gui.QPixmap) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createPreviewPixmap"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createPreviewPixmap", func() *gui.QPixmap {
				signal.(func() *gui.QPixmap)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createPreviewPixmap", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectCreatePreviewPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createPreviewPixmap")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CreatePreviewPixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPixmapFromPointer(C.QDesignerFormWindowManagerInterface_CreatePreviewPixmap(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerFormWindowManagerInterface_FormWindowCount
func callbackQDesignerFormWindowManagerInterface_FormWindowCount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "formWindowCount"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectFormWindowCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "formWindowCount"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "formWindowCount", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "formWindowCount", f)
		}
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectFormWindowCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "formWindowCount")
	}
}

func (ptr *QDesignerFormWindowManagerInterface) FormWindowCount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerFormWindowManagerInterface_FormWindowCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDesignerFormWindowManagerInterface) __dragItems_item_list_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___dragItems_item_list_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerFormWindowManagerInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowManagerInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowManagerInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowManagerInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerFormWindowManagerInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerFormWindowManagerInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerFormWindowManagerInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerFormWindowManagerInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerFormWindowManagerInterface_Event
func callbackQDesignerFormWindowManagerInterface_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDesignerFormWindowManagerInterface) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowManagerInterface_EventFilter
func callbackQDesignerFormWindowManagerInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerFormWindowManagerInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerFormWindowManagerInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerFormWindowManagerInterface_ChildEvent
func callbackQDesignerFormWindowManagerInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ConnectNotify
func callbackQDesignerFormWindowManagerInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowManagerInterface_CustomEvent
func callbackQDesignerFormWindowManagerInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_DeleteLater
func callbackQDesignerFormWindowManagerInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerFormWindowManagerInterface_Destroyed
func callbackQDesignerFormWindowManagerInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerFormWindowManagerInterface_DisconnectNotify
func callbackQDesignerFormWindowManagerInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerFormWindowManagerInterface_ObjectNameChanged
func callbackQDesignerFormWindowManagerInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerFormWindowManagerInterface_TimerEvent
func callbackQDesignerFormWindowManagerInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerFormWindowManagerInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerFormWindowManagerInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerFormWindowManagerInterface_MetaObject
func callbackQDesignerFormWindowManagerInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerFormWindowManagerInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerFormWindowManagerInterface) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QDesignerMemberSheetExtension) QDesignerMemberSheetExtension_PTR() *QDesignerMemberSheetExtension {
	return ptr
}

func (ptr *QDesignerMemberSheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerMemberSheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerMemberSheetExtension_SetMemberGroup
func callbackQDesignerMemberSheetExtension_SetMemberGroup(ptr unsafe.Pointer, index C.int, group C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setMemberGroup"); signal != nil {
		signal.(func(int, string))(int(int32(index)), cGoUnpackString(group))
	}

}

func (ptr *QDesignerMemberSheetExtension) ConnectSetMemberGroup(f func(index int, group string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMemberGroup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setMemberGroup", func(index int, group string) {
				signal.(func(int, string))(index, group)
				f(index, group)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMemberGroup", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetMemberGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMemberGroup")
	}
}

func (ptr *QDesignerMemberSheetExtension) SetMemberGroup(index int, group string) {
	if ptr.Pointer() != nil {
		var groupC *C.char
		if group != "" {
			groupC = C.CString(group)
			defer C.free(unsafe.Pointer(groupC))
		}
		C.QDesignerMemberSheetExtension_SetMemberGroup(ptr.Pointer(), C.int(int32(index)), C.struct_QtDesigner_PackedString{data: groupC, len: C.longlong(len(group))})
	}
}

//export callbackQDesignerMemberSheetExtension_SetVisible
func callbackQDesignerMemberSheetExtension_SetVisible(ptr unsafe.Pointer, index C.int, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(visible) != 0)
	}

}

func (ptr *QDesignerMemberSheetExtension) ConnectSetVisible(f func(index int, visible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", func(index int, visible bool) {
				signal.(func(int, bool))(index, visible)
				f(index, visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVisible")
	}
}

func (ptr *QDesignerMemberSheetExtension) SetVisible(index int, visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension_SetVisible(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension
func callbackQDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerMemberSheetExtension"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerMemberSheetExtensionFromPointer(ptr).DestroyQDesignerMemberSheetExtensionDefault()
	}
}

func (ptr *QDesignerMemberSheetExtension) ConnectDestroyQDesignerMemberSheetExtension(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerMemberSheetExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerMemberSheetExtension", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerMemberSheetExtension", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectDestroyQDesignerMemberSheetExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerMemberSheetExtension")
	}
}

func (ptr *QDesignerMemberSheetExtension) DestroyQDesignerMemberSheetExtension() {
	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtension(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerMemberSheetExtension) DestroyQDesignerMemberSheetExtensionDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension_DestroyQDesignerMemberSheetExtensionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerMemberSheetExtension_ParameterNames
func callbackQDesignerMemberSheetExtension_ParameterNames(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parameterNames"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQDesignerMemberSheetExtensionFromPointer(NewQDesignerMemberSheetExtensionFromPointer(nil).__parameterNames_newList())
			for _, v := range signal.(func(int) []*core.QByteArray)(int(int32(index))) {
				tmpList.__parameterNames_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQDesignerMemberSheetExtensionFromPointer(NewQDesignerMemberSheetExtensionFromPointer(nil).__parameterNames_newList())
		for _, v := range make([]*core.QByteArray, 0) {
			tmpList.__parameterNames_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QDesignerMemberSheetExtension) ConnectParameterNames(f func(index int) []*core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parameterNames"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parameterNames", func(index int) []*core.QByteArray {
				signal.(func(int) []*core.QByteArray)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parameterNames", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectParameterNames() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parameterNames")
	}
}

func (ptr *QDesignerMemberSheetExtension) ParameterNames(index int) []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDesigner_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDesignerMemberSheetExtensionFromPointer(l.data).__parameterNames_atList(i)
			}
			return out
		}(C.QDesignerMemberSheetExtension_ParameterNames(ptr.Pointer(), C.int(int32(index))))
	}
	return make([]*core.QByteArray, 0)
}

//export callbackQDesignerMemberSheetExtension_ParameterTypes
func callbackQDesignerMemberSheetExtension_ParameterTypes(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parameterTypes"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQDesignerMemberSheetExtensionFromPointer(NewQDesignerMemberSheetExtensionFromPointer(nil).__parameterTypes_newList())
			for _, v := range signal.(func(int) []*core.QByteArray)(int(int32(index))) {
				tmpList.__parameterTypes_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQDesignerMemberSheetExtensionFromPointer(NewQDesignerMemberSheetExtensionFromPointer(nil).__parameterTypes_newList())
		for _, v := range make([]*core.QByteArray, 0) {
			tmpList.__parameterTypes_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QDesignerMemberSheetExtension) ConnectParameterTypes(f func(index int) []*core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "parameterTypes"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "parameterTypes", func(index int) []*core.QByteArray {
				signal.(func(int) []*core.QByteArray)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "parameterTypes", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectParameterTypes() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "parameterTypes")
	}
}

func (ptr *QDesignerMemberSheetExtension) ParameterTypes(index int) []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDesigner_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDesignerMemberSheetExtensionFromPointer(l.data).__parameterTypes_atList(i)
			}
			return out
		}(C.QDesignerMemberSheetExtension_ParameterTypes(ptr.Pointer(), C.int(int32(index))))
	}
	return make([]*core.QByteArray, 0)
}

//export callbackQDesignerMemberSheetExtension_DeclaredInClass
func callbackQDesignerMemberSheetExtension_DeclaredInClass(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "declaredInClass"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerMemberSheetExtension) ConnectDeclaredInClass(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "declaredInClass"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "declaredInClass", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "declaredInClass", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectDeclaredInClass() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "declaredInClass")
	}
}

func (ptr *QDesignerMemberSheetExtension) DeclaredInClass(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerMemberSheetExtension_DeclaredInClass(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_MemberGroup
func callbackQDesignerMemberSheetExtension_MemberGroup(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "memberGroup"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberGroup(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "memberGroup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "memberGroup", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "memberGroup", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "memberGroup")
	}
}

func (ptr *QDesignerMemberSheetExtension) MemberGroup(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerMemberSheetExtension_MemberGroup(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_MemberName
func callbackQDesignerMemberSheetExtension_MemberName(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "memberName"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerMemberSheetExtension) ConnectMemberName(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "memberName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "memberName", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "memberName", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectMemberName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "memberName")
	}
}

func (ptr *QDesignerMemberSheetExtension) MemberName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerMemberSheetExtension_MemberName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_Signature
func callbackQDesignerMemberSheetExtension_Signature(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "signature"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerMemberSheetExtension) ConnectSignature(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "signature"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "signature", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "signature", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectSignature() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "signature")
	}
}

func (ptr *QDesignerMemberSheetExtension) Signature(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerMemberSheetExtension_Signature(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerMemberSheetExtension_InheritedFromWidget
func callbackQDesignerMemberSheetExtension_InheritedFromWidget(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "inheritedFromWidget"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectInheritedFromWidget(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inheritedFromWidget"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "inheritedFromWidget", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inheritedFromWidget", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectInheritedFromWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inheritedFromWidget")
	}
}

func (ptr *QDesignerMemberSheetExtension) InheritedFromWidget(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_InheritedFromWidget(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsSignal
func callbackQDesignerMemberSheetExtension_IsSignal(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isSignal"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSignal(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isSignal"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isSignal", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isSignal", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSignal() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isSignal")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsSignal(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsSignal(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsSlot
func callbackQDesignerMemberSheetExtension_IsSlot(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isSlot"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsSlot(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isSlot"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isSlot", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isSlot", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsSlot() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isSlot")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsSlot(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsSlot(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_IsVisible
func callbackQDesignerMemberSheetExtension_IsVisible(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIsVisible(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isVisible", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isVisible", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIsVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isVisible")
	}
}

func (ptr *QDesignerMemberSheetExtension) IsVisible(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerMemberSheetExtension_IsVisible(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerMemberSheetExtension_Count
func callbackQDesignerMemberSheetExtension_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerMemberSheetExtension) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "count", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QDesignerMemberSheetExtension) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerMemberSheetExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerMemberSheetExtension_IndexOf
func callbackQDesignerMemberSheetExtension_IndexOf(ptr unsafe.Pointer, name C.struct_QtDesigner_PackedString) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(string) int)(cGoUnpackString(name))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerMemberSheetExtension) ConnectIndexOf(f func(name string) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexOf"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexOf", func(name string) int {
				signal.(func(string) int)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexOf", f)
		}
	}
}

func (ptr *QDesignerMemberSheetExtension) DisconnectIndexOf() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexOf")
	}
}

func (ptr *QDesignerMemberSheetExtension) IndexOf(name string) int {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int(int32(C.QDesignerMemberSheetExtension_IndexOf(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))})))
	}
	return 0
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerMemberSheetExtension___parameterNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension___parameterNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerMemberSheetExtension) __parameterNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerMemberSheetExtension___parameterNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerMemberSheetExtension___parameterTypes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerMemberSheetExtension___parameterTypes_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerMemberSheetExtension) __parameterTypes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerMemberSheetExtension___parameterTypes_newList(ptr.Pointer()))
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

func NewQDesignerObjectInspectorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerObjectInspectorInterface {
	var n = new(QDesignerObjectInspectorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerObjectInspectorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerObjectInspectorInterface {
	var tmpValue = NewQDesignerObjectInspectorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_NewQDesignerObjectInspectorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDesignerObjectInspectorInterface_SetFormWindow
func callbackQDesignerObjectInspectorInterface_SetFormWindow(ptr unsafe.Pointer, formWindow unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFormWindow"); signal != nil {
		signal.(func(*QDesignerFormWindowInterface))(NewQDesignerFormWindowInterfaceFromPointer(formWindow))
	}

}

func (ptr *QDesignerObjectInspectorInterface) ConnectSetFormWindow(f func(formWindow *QDesignerFormWindowInterface)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFormWindow"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFormWindow", func(formWindow *QDesignerFormWindowInterface) {
				signal.(func(*QDesignerFormWindowInterface))(formWindow)
				f(formWindow)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFormWindow", f)
		}
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectSetFormWindow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFormWindow")
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetFormWindow(formWindow QDesignerFormWindowInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetFormWindow(ptr.Pointer(), PointerFromQDesignerFormWindowInterface(formWindow))
	}
}

//export callbackQDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface
func callbackQDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerObjectInspectorInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DestroyQDesignerObjectInspectorInterfaceDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectDestroyQDesignerObjectInspectorInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerObjectInspectorInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerObjectInspectorInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerObjectInspectorInterface", f)
		}
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectDestroyQDesignerObjectInspectorInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerObjectInspectorInterface")
	}
}

func (ptr *QDesignerObjectInspectorInterface) DestroyQDesignerObjectInspectorInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerObjectInspectorInterface) DestroyQDesignerObjectInspectorInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DestroyQDesignerObjectInspectorInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerObjectInspectorInterface_Core
func callbackQDesignerObjectInspectorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerObjectInspectorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "core"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "core", func() *QDesignerFormEditorInterface {
				signal.(func() *QDesignerFormEditorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "core", f)
		}
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "core")
	}
}

func (ptr *QDesignerObjectInspectorInterface) Core() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) CoreDefault() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerObjectInspectorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerObjectInspectorInterface___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerObjectInspectorInterface___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerObjectInspectorInterface___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerObjectInspectorInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerObjectInspectorInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerObjectInspectorInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerObjectInspectorInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerObjectInspectorInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerObjectInspectorInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerObjectInspectorInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerObjectInspectorInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerObjectInspectorInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerObjectInspectorInterface_Close
func callbackQDesignerObjectInspectorInterface_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerObjectInspectorInterface) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_Event
func callbackQDesignerObjectInspectorInterface_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerObjectInspectorInterface) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_FocusNextPrevChild
func callbackQDesignerObjectInspectorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerObjectInspectorInterface) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_NativeEvent
func callbackQDesignerObjectInspectorInterface_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QDesignerObjectInspectorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_ActionEvent
func callbackQDesignerObjectInspectorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ChangeEvent
func callbackQDesignerObjectInspectorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_CloseEvent
func callbackQDesignerObjectInspectorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ContextMenuEvent
func callbackQDesignerObjectInspectorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_CustomContextMenuRequested
func callbackQDesignerObjectInspectorInterface_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQDesignerObjectInspectorInterface_DragEnterEvent
func callbackQDesignerObjectInspectorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DragLeaveEvent
func callbackQDesignerObjectInspectorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DragMoveEvent
func callbackQDesignerObjectInspectorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DropEvent
func callbackQDesignerObjectInspectorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_EnterEvent
func callbackQDesignerObjectInspectorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_FocusInEvent
func callbackQDesignerObjectInspectorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_FocusOutEvent
func callbackQDesignerObjectInspectorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Hide
func callbackQDesignerObjectInspectorInterface_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) HideDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_HideEvent
func callbackQDesignerObjectInspectorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_InputMethodEvent
func callbackQDesignerObjectInspectorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_KeyPressEvent
func callbackQDesignerObjectInspectorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_KeyReleaseEvent
func callbackQDesignerObjectInspectorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_LeaveEvent
func callbackQDesignerObjectInspectorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Lower
func callbackQDesignerObjectInspectorInterface_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent
func callbackQDesignerObjectInspectorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseMoveEvent
func callbackQDesignerObjectInspectorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MousePressEvent
func callbackQDesignerObjectInspectorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MouseReleaseEvent
func callbackQDesignerObjectInspectorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MoveEvent
func callbackQDesignerObjectInspectorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_PaintEvent
func callbackQDesignerObjectInspectorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Raise
func callbackQDesignerObjectInspectorInterface_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_Repaint
func callbackQDesignerObjectInspectorInterface_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ResizeEvent
func callbackQDesignerObjectInspectorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetDisabled
func callbackQDesignerObjectInspectorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetEnabled
func callbackQDesignerObjectInspectorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetFocus2
func callbackQDesignerObjectInspectorInterface_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_SetHidden
func callbackQDesignerObjectInspectorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetStyleSheet
func callbackQDesignerObjectInspectorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QDesignerObjectInspectorInterface_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQDesignerObjectInspectorInterface_SetVisible
func callbackQDesignerObjectInspectorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetWindowModified
func callbackQDesignerObjectInspectorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerObjectInspectorInterface_SetWindowTitle
func callbackQDesignerObjectInspectorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QDesignerObjectInspectorInterface) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QDesignerObjectInspectorInterface_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQDesignerObjectInspectorInterface_Show
func callbackQDesignerObjectInspectorInterface_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowEvent
func callbackQDesignerObjectInspectorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowFullScreen
func callbackQDesignerObjectInspectorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowMaximized
func callbackQDesignerObjectInspectorInterface_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowMinimized
func callbackQDesignerObjectInspectorInterface_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_ShowNormal
func callbackQDesignerObjectInspectorInterface_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_TabletEvent
func callbackQDesignerObjectInspectorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_Update
func callbackQDesignerObjectInspectorInterface_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_UpdateMicroFocus
func callbackQDesignerObjectInspectorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerObjectInspectorInterface_WheelEvent
func callbackQDesignerObjectInspectorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_WindowIconChanged
func callbackQDesignerObjectInspectorInterface_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQDesignerObjectInspectorInterface_WindowTitleChanged
func callbackQDesignerObjectInspectorInterface_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQDesignerObjectInspectorInterface_PaintEngine
func callbackQDesignerObjectInspectorInterface_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QDesignerObjectInspectorInterface) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QDesignerObjectInspectorInterface_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_MinimumSizeHint
func callbackQDesignerObjectInspectorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerObjectInspectorInterface) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_SizeHint
func callbackQDesignerObjectInspectorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerObjectInspectorInterface) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerObjectInspectorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_InputMethodQuery
func callbackQDesignerObjectInspectorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerObjectInspectorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerObjectInspectorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerObjectInspectorInterface_HasHeightForWidth
func callbackQDesignerObjectInspectorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerObjectInspectorInterface) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_HeightForWidth
func callbackQDesignerObjectInspectorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerObjectInspectorInterface) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerObjectInspectorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerObjectInspectorInterface_Metric
func callbackQDesignerObjectInspectorInterface_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QDesignerObjectInspectorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerObjectInspectorInterface_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQDesignerObjectInspectorInterface_EventFilter
func callbackQDesignerObjectInspectorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerObjectInspectorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerObjectInspectorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerObjectInspectorInterface_ChildEvent
func callbackQDesignerObjectInspectorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_ConnectNotify
func callbackQDesignerObjectInspectorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerObjectInspectorInterface_CustomEvent
func callbackQDesignerObjectInspectorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_DeleteLater
func callbackQDesignerObjectInspectorInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerObjectInspectorInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerObjectInspectorInterface_Destroyed
func callbackQDesignerObjectInspectorInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerObjectInspectorInterface_DisconnectNotify
func callbackQDesignerObjectInspectorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerObjectInspectorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerObjectInspectorInterface_ObjectNameChanged
func callbackQDesignerObjectInspectorInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerObjectInspectorInterface_TimerEvent
func callbackQDesignerObjectInspectorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerObjectInspectorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerObjectInspectorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerObjectInspectorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerObjectInspectorInterface_MetaObject
func callbackQDesignerObjectInspectorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerObjectInspectorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerObjectInspectorInterface) MetaObjectDefault() *core.QMetaObject {
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

func NewQDesignerPropertyEditorInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerPropertyEditorInterface {
	var n = new(QDesignerPropertyEditorInterface)
	n.SetPointer(ptr)
	return n
}
func NewQDesignerPropertyEditorInterface(parent widgets.QWidget_ITF, flags core.Qt__WindowType) *QDesignerPropertyEditorInterface {
	var tmpValue = NewQDesignerPropertyEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_NewQDesignerPropertyEditorInterface(widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDesignerPropertyEditorInterface_PropertyChanged
func callbackQDesignerPropertyEditorInterface_PropertyChanged(ptr unsafe.Pointer, name C.struct_QtDesigner_PackedString, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "propertyChanged"); signal != nil {
		signal.(func(string, *core.QVariant))(cGoUnpackString(name), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectPropertyChanged(f func(name string, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "propertyChanged") {
			C.QDesignerPropertyEditorInterface_ConnectPropertyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "propertyChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "propertyChanged", func(name string, value *core.QVariant) {
				signal.(func(string, *core.QVariant))(name, value)
				f(name, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "propertyChanged", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectPropertyChanged() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DisconnectPropertyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "propertyChanged")
	}
}

func (ptr *QDesignerPropertyEditorInterface) PropertyChanged(name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDesignerPropertyEditorInterface_PropertyChanged(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetObject
func callbackQDesignerPropertyEditorInterface_SetObject(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setObject"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetObject(f func(object *core.QObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setObject"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setObject", func(object *core.QObject) {
				signal.(func(*core.QObject))(object)
				f(object)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setObject", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setObject")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetObject(object core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetPropertyValue
func callbackQDesignerPropertyEditorInterface_SetPropertyValue(ptr unsafe.Pointer, name C.struct_QtDesigner_PackedString, value unsafe.Pointer, changed C.char) {
	if signal := qt.GetSignal(ptr, "setPropertyValue"); signal != nil {
		signal.(func(string, *core.QVariant, bool))(cGoUnpackString(name), core.NewQVariantFromPointer(value), int8(changed) != 0)
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetPropertyValue(f func(name string, value *core.QVariant, changed bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPropertyValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPropertyValue", func(name string, value *core.QVariant, changed bool) {
				signal.(func(string, *core.QVariant, bool))(name, value, changed)
				f(name, value, changed)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPropertyValue", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetPropertyValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPropertyValue")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetPropertyValue(name string, value core.QVariant_ITF, changed bool) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QDesignerPropertyEditorInterface_SetPropertyValue(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))}, core.PointerFromQVariant(value), C.char(int8(qt.GoBoolToInt(changed))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetReadOnly
func callbackQDesignerPropertyEditorInterface_SetReadOnly(ptr unsafe.Pointer, readOnly C.char) {
	if signal := qt.GetSignal(ptr, "setReadOnly"); signal != nil {
		signal.(func(bool))(int8(readOnly) != 0)
	}

}

func (ptr *QDesignerPropertyEditorInterface) ConnectSetReadOnly(f func(readOnly bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setReadOnly"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setReadOnly", func(readOnly bool) {
				signal.(func(bool))(readOnly)
				f(readOnly)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setReadOnly", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectSetReadOnly() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setReadOnly")
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetReadOnly(readOnly bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(readOnly))))
	}
}

//export callbackQDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface
func callbackQDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerPropertyEditorInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DestroyQDesignerPropertyEditorInterfaceDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectDestroyQDesignerPropertyEditorInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerPropertyEditorInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerPropertyEditorInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerPropertyEditorInterface", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectDestroyQDesignerPropertyEditorInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerPropertyEditorInterface")
	}
}

func (ptr *QDesignerPropertyEditorInterface) DestroyQDesignerPropertyEditorInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerPropertyEditorInterface) DestroyQDesignerPropertyEditorInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DestroyQDesignerPropertyEditorInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerPropertyEditorInterface_Core
func callbackQDesignerPropertyEditorInterface_Core(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "core"); signal != nil {
		return PointerFromQDesignerFormEditorInterface(signal.(func() *QDesignerFormEditorInterface)())
	}

	return PointerFromQDesignerFormEditorInterface(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CoreDefault())
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCore(f func() *QDesignerFormEditorInterface) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "core"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "core", func() *QDesignerFormEditorInterface {
				signal.(func() *QDesignerFormEditorInterface)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "core", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCore() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "core")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Core() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_Core(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) CoreDefault() *QDesignerFormEditorInterface {
	if ptr.Pointer() != nil {
		var tmpValue = NewQDesignerFormEditorInterfaceFromPointer(C.QDesignerPropertyEditorInterface_CoreDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_Object
func callbackQDesignerPropertyEditorInterface_Object(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "object"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(core.NewQObject(nil))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectObject(f func() *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "object"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "object", func() *core.QObject {
				signal.(func() *core.QObject)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "object", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "object")
	}
}

func (ptr *QDesignerPropertyEditorInterface) Object() *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_CurrentPropertyName
func callbackQDesignerPropertyEditorInterface_CurrentPropertyName(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "currentPropertyName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectCurrentPropertyName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentPropertyName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "currentPropertyName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentPropertyName", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectCurrentPropertyName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "currentPropertyName")
	}
}

func (ptr *QDesignerPropertyEditorInterface) CurrentPropertyName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerPropertyEditorInterface_CurrentPropertyName(ptr.Pointer()))
	}
	return ""
}

//export callbackQDesignerPropertyEditorInterface_IsReadOnly
func callbackQDesignerPropertyEditorInterface_IsReadOnly(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isReadOnly"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertyEditorInterface) ConnectIsReadOnly(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isReadOnly"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isReadOnly", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isReadOnly", f)
		}
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectIsReadOnly() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isReadOnly")
	}
}

func (ptr *QDesignerPropertyEditorInterface) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerPropertyEditorInterface___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerPropertyEditorInterface___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerPropertyEditorInterface___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerPropertyEditorInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerPropertyEditorInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerPropertyEditorInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerPropertyEditorInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerPropertyEditorInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerPropertyEditorInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerPropertyEditorInterface_Close
func callbackQDesignerPropertyEditorInterface_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerPropertyEditorInterface) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_Event
func callbackQDesignerPropertyEditorInterface_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerPropertyEditorInterface) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_FocusNextPrevChild
func callbackQDesignerPropertyEditorInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerPropertyEditorInterface) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_NativeEvent
func callbackQDesignerPropertyEditorInterface_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QDesignerPropertyEditorInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_ActionEvent
func callbackQDesignerPropertyEditorInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ChangeEvent
func callbackQDesignerPropertyEditorInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_CloseEvent
func callbackQDesignerPropertyEditorInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ContextMenuEvent
func callbackQDesignerPropertyEditorInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_CustomContextMenuRequested
func callbackQDesignerPropertyEditorInterface_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQDesignerPropertyEditorInterface_DragEnterEvent
func callbackQDesignerPropertyEditorInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DragLeaveEvent
func callbackQDesignerPropertyEditorInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DragMoveEvent
func callbackQDesignerPropertyEditorInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DropEvent
func callbackQDesignerPropertyEditorInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_EnterEvent
func callbackQDesignerPropertyEditorInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_FocusInEvent
func callbackQDesignerPropertyEditorInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_FocusOutEvent
func callbackQDesignerPropertyEditorInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Hide
func callbackQDesignerPropertyEditorInterface_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) HideDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_HideEvent
func callbackQDesignerPropertyEditorInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_InputMethodEvent
func callbackQDesignerPropertyEditorInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_KeyPressEvent
func callbackQDesignerPropertyEditorInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_KeyReleaseEvent
func callbackQDesignerPropertyEditorInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_LeaveEvent
func callbackQDesignerPropertyEditorInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Lower
func callbackQDesignerPropertyEditorInterface_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent
func callbackQDesignerPropertyEditorInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseMoveEvent
func callbackQDesignerPropertyEditorInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MousePressEvent
func callbackQDesignerPropertyEditorInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MouseReleaseEvent
func callbackQDesignerPropertyEditorInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MoveEvent
func callbackQDesignerPropertyEditorInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_PaintEvent
func callbackQDesignerPropertyEditorInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Raise
func callbackQDesignerPropertyEditorInterface_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_Repaint
func callbackQDesignerPropertyEditorInterface_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ResizeEvent
func callbackQDesignerPropertyEditorInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetDisabled
func callbackQDesignerPropertyEditorInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetEnabled
func callbackQDesignerPropertyEditorInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetFocus2
func callbackQDesignerPropertyEditorInterface_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_SetHidden
func callbackQDesignerPropertyEditorInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetStyleSheet
func callbackQDesignerPropertyEditorInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QDesignerPropertyEditorInterface_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQDesignerPropertyEditorInterface_SetVisible
func callbackQDesignerPropertyEditorInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetWindowModified
func callbackQDesignerPropertyEditorInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerPropertyEditorInterface_SetWindowTitle
func callbackQDesignerPropertyEditorInterface_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QDesignerPropertyEditorInterface) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QDesignerPropertyEditorInterface_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQDesignerPropertyEditorInterface_Show
func callbackQDesignerPropertyEditorInterface_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowEvent
func callbackQDesignerPropertyEditorInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowFullScreen
func callbackQDesignerPropertyEditorInterface_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowMaximized
func callbackQDesignerPropertyEditorInterface_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowMinimized
func callbackQDesignerPropertyEditorInterface_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_ShowNormal
func callbackQDesignerPropertyEditorInterface_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_TabletEvent
func callbackQDesignerPropertyEditorInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_Update
func callbackQDesignerPropertyEditorInterface_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_UpdateMicroFocus
func callbackQDesignerPropertyEditorInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerPropertyEditorInterface_WheelEvent
func callbackQDesignerPropertyEditorInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_WindowIconChanged
func callbackQDesignerPropertyEditorInterface_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQDesignerPropertyEditorInterface_WindowTitleChanged
func callbackQDesignerPropertyEditorInterface_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQDesignerPropertyEditorInterface_PaintEngine
func callbackQDesignerPropertyEditorInterface_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QDesignerPropertyEditorInterface) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QDesignerPropertyEditorInterface_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_MinimumSizeHint
func callbackQDesignerPropertyEditorInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerPropertyEditorInterface) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_SizeHint
func callbackQDesignerPropertyEditorInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerPropertyEditorInterface) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerPropertyEditorInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_InputMethodQuery
func callbackQDesignerPropertyEditorInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerPropertyEditorInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerPropertyEditorInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertyEditorInterface_HasHeightForWidth
func callbackQDesignerPropertyEditorInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerPropertyEditorInterface) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_HeightForWidth
func callbackQDesignerPropertyEditorInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerPropertyEditorInterface) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertyEditorInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerPropertyEditorInterface_Metric
func callbackQDesignerPropertyEditorInterface_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QDesignerPropertyEditorInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertyEditorInterface_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQDesignerPropertyEditorInterface_EventFilter
func callbackQDesignerPropertyEditorInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerPropertyEditorInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertyEditorInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerPropertyEditorInterface_ChildEvent
func callbackQDesignerPropertyEditorInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_ConnectNotify
func callbackQDesignerPropertyEditorInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerPropertyEditorInterface_CustomEvent
func callbackQDesignerPropertyEditorInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_DeleteLater
func callbackQDesignerPropertyEditorInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerPropertyEditorInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerPropertyEditorInterface_Destroyed
func callbackQDesignerPropertyEditorInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerPropertyEditorInterface_DisconnectNotify
func callbackQDesignerPropertyEditorInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerPropertyEditorInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerPropertyEditorInterface_ObjectNameChanged
func callbackQDesignerPropertyEditorInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerPropertyEditorInterface_TimerEvent
func callbackQDesignerPropertyEditorInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerPropertyEditorInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerPropertyEditorInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertyEditorInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerPropertyEditorInterface_MetaObject
func callbackQDesignerPropertyEditorInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerPropertyEditorInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerPropertyEditorInterface) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QDesignerPropertySheetExtension) QDesignerPropertySheetExtension_PTR() *QDesignerPropertySheetExtension {
	return ptr
}

func (ptr *QDesignerPropertySheetExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerPropertySheetExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerPropertySheetExtension_Reset
func callbackQDesignerPropertySheetExtension_Reset(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectReset(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reset"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "reset", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reset", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reset")
	}
}

func (ptr *QDesignerPropertySheetExtension) Reset(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_Reset(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_SetAttribute
func callbackQDesignerPropertySheetExtension_SetAttribute(ptr unsafe.Pointer, index C.int, attribute C.char) {
	if signal := qt.GetSignal(ptr, "setAttribute"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(attribute) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetAttribute(f func(index int, attribute bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAttribute"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setAttribute", func(index int, attribute bool) {
				signal.(func(int, bool))(index, attribute)
				f(index, attribute)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAttribute", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAttribute")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetAttribute(index int, attribute bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetAttribute(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(attribute))))
	}
}

//export callbackQDesignerPropertySheetExtension_SetChanged
func callbackQDesignerPropertySheetExtension_SetChanged(ptr unsafe.Pointer, index C.int, changed C.char) {
	if signal := qt.GetSignal(ptr, "setChanged"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(changed) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetChanged(f func(index int, changed bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setChanged", func(index int, changed bool) {
				signal.(func(int, bool))(index, changed)
				f(index, changed)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setChanged", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setChanged")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetChanged(index int, changed bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetChanged(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(changed))))
	}
}

//export callbackQDesignerPropertySheetExtension_SetProperty
func callbackQDesignerPropertySheetExtension_SetProperty(ptr unsafe.Pointer, index C.int, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setProperty"); signal != nil {
		signal.(func(int, *core.QVariant))(int(int32(index)), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetProperty(f func(index int, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProperty"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", func(index int, value *core.QVariant) {
				signal.(func(int, *core.QVariant))(index, value)
				f(index, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProperty", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProperty")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetProperty(index int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetProperty(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(value))
	}
}

//export callbackQDesignerPropertySheetExtension_SetPropertyGroup
func callbackQDesignerPropertySheetExtension_SetPropertyGroup(ptr unsafe.Pointer, index C.int, group C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setPropertyGroup"); signal != nil {
		signal.(func(int, string))(int(int32(index)), cGoUnpackString(group))
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetPropertyGroup(f func(index int, group string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPropertyGroup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPropertyGroup", func(index int, group string) {
				signal.(func(int, string))(index, group)
				f(index, group)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPropertyGroup", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetPropertyGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPropertyGroup")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetPropertyGroup(index int, group string) {
	if ptr.Pointer() != nil {
		var groupC *C.char
		if group != "" {
			groupC = C.CString(group)
			defer C.free(unsafe.Pointer(groupC))
		}
		C.QDesignerPropertySheetExtension_SetPropertyGroup(ptr.Pointer(), C.int(int32(index)), C.struct_QtDesigner_PackedString{data: groupC, len: C.longlong(len(group))})
	}
}

//export callbackQDesignerPropertySheetExtension_SetVisible
func callbackQDesignerPropertySheetExtension_SetVisible(ptr unsafe.Pointer, index C.int, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(int, bool))(int(int32(index)), int8(visible) != 0)
	}

}

func (ptr *QDesignerPropertySheetExtension) ConnectSetVisible(f func(index int, visible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", func(index int, visible bool) {
				signal.(func(int, bool))(index, visible)
				f(index, visible)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVisible")
	}
}

func (ptr *QDesignerPropertySheetExtension) SetVisible(index int, visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_SetVisible(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension
func callbackQDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerPropertySheetExtension"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerPropertySheetExtensionFromPointer(ptr).DestroyQDesignerPropertySheetExtensionDefault()
	}
}

func (ptr *QDesignerPropertySheetExtension) ConnectDestroyQDesignerPropertySheetExtension(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerPropertySheetExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerPropertySheetExtension", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerPropertySheetExtension", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectDestroyQDesignerPropertySheetExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerPropertySheetExtension")
	}
}

func (ptr *QDesignerPropertySheetExtension) DestroyQDesignerPropertySheetExtension() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtension(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerPropertySheetExtension) DestroyQDesignerPropertySheetExtensionDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerPropertySheetExtension_DestroyQDesignerPropertySheetExtensionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerPropertySheetExtension_PropertyGroup
func callbackQDesignerPropertySheetExtension_PropertyGroup(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "propertyGroup"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyGroup(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "propertyGroup"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "propertyGroup", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "propertyGroup", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "propertyGroup")
	}
}

func (ptr *QDesignerPropertySheetExtension) PropertyGroup(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerPropertySheetExtension_PropertyGroup(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerPropertySheetExtension_PropertyName
func callbackQDesignerPropertySheetExtension_PropertyName(ptr unsafe.Pointer, index C.int) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "propertyName"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(index)))
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerPropertySheetExtension) ConnectPropertyName(f func(index int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "propertyName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "propertyName", func(index int) string {
				signal.(func(int) string)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "propertyName", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectPropertyName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "propertyName")
	}
}

func (ptr *QDesignerPropertySheetExtension) PropertyName(index int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerPropertySheetExtension_PropertyName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

//export callbackQDesignerPropertySheetExtension_Property
func callbackQDesignerPropertySheetExtension_Property(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "property"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(index))))
	}

	return core.PointerFromQVariant(core.NewQVariant())
}

func (ptr *QDesignerPropertySheetExtension) ConnectProperty(f func(index int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "property"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "property", func(index int) *core.QVariant {
				signal.(func(int) *core.QVariant)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "property", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectProperty() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "property")
	}
}

func (ptr *QDesignerPropertySheetExtension) Property(index int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerPropertySheetExtension_Property(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerPropertySheetExtension_HasReset
func callbackQDesignerPropertySheetExtension_HasReset(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "hasReset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectHasReset(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hasReset"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hasReset", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hasReset", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectHasReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hasReset")
	}
}

func (ptr *QDesignerPropertySheetExtension) HasReset(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_HasReset(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsAttribute
func callbackQDesignerPropertySheetExtension_IsAttribute(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isAttribute"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsAttribute(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isAttribute"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isAttribute", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isAttribute", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsAttribute() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isAttribute")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsAttribute(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsAttribute(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsChanged
func callbackQDesignerPropertySheetExtension_IsChanged(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isChanged"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsChanged(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isChanged", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isChanged", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsChanged() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isChanged")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsChanged(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsChanged(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsEnabled
func callbackQDesignerPropertySheetExtension_IsEnabled(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isEnabled"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerPropertySheetExtensionFromPointer(ptr).IsEnabledDefault(int(int32(index))))))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsEnabled(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isEnabled"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isEnabled", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isEnabled", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isEnabled")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsEnabled(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QDesignerPropertySheetExtension) IsEnabledDefault(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsEnabledDefault(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_IsVisible
func callbackQDesignerPropertySheetExtension_IsVisible(ptr unsafe.Pointer, index C.int) C.char {
	if signal := qt.GetSignal(ptr, "isVisible"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIsVisible(f func(index int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isVisible"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isVisible", func(index int) bool {
				signal.(func(int) bool)(index)
				return f(index)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isVisible", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIsVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isVisible")
	}
}

func (ptr *QDesignerPropertySheetExtension) IsVisible(index int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerPropertySheetExtension_IsVisible(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQDesignerPropertySheetExtension_Count
func callbackQDesignerPropertySheetExtension_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerPropertySheetExtension) ConnectCount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "count"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "count", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "count", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectCount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "count")
	}
}

func (ptr *QDesignerPropertySheetExtension) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerPropertySheetExtension_Count(ptr.Pointer())))
	}
	return 0
}

//export callbackQDesignerPropertySheetExtension_IndexOf
func callbackQDesignerPropertySheetExtension_IndexOf(ptr unsafe.Pointer, name C.struct_QtDesigner_PackedString) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(string) int)(cGoUnpackString(name))))
	}

	return C.int(int32(0))
}

func (ptr *QDesignerPropertySheetExtension) ConnectIndexOf(f func(name string) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "indexOf"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "indexOf", func(name string) int {
				signal.(func(string) int)(name)
				return f(name)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "indexOf", f)
		}
	}
}

func (ptr *QDesignerPropertySheetExtension) DisconnectIndexOf() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "indexOf")
	}
}

func (ptr *QDesignerPropertySheetExtension) IndexOf(name string) int {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		return int(int32(C.QDesignerPropertySheetExtension_IndexOf(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: nameC, len: C.longlong(len(name))})))
	}
	return 0
}

type QDesignerTaskMenuExtension struct {
	ptr unsafe.Pointer
}

type QDesignerTaskMenuExtension_ITF interface {
	QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension
}

func (ptr *QDesignerTaskMenuExtension) QDesignerTaskMenuExtension_PTR() *QDesignerTaskMenuExtension {
	return ptr
}

func (ptr *QDesignerTaskMenuExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//export callbackQDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension
func callbackQDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerTaskMenuExtension"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerTaskMenuExtensionFromPointer(ptr).DestroyQDesignerTaskMenuExtensionDefault()
	}
}

func (ptr *QDesignerTaskMenuExtension) ConnectDestroyQDesignerTaskMenuExtension(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerTaskMenuExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerTaskMenuExtension", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerTaskMenuExtension", f)
		}
	}
}

func (ptr *QDesignerTaskMenuExtension) DisconnectDestroyQDesignerTaskMenuExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerTaskMenuExtension")
	}
}

func (ptr *QDesignerTaskMenuExtension) DestroyQDesignerTaskMenuExtension() {
	if ptr.Pointer() != nil {
		C.QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtension(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDesignerTaskMenuExtension) DestroyQDesignerTaskMenuExtensionDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerTaskMenuExtension_DestroyQDesignerTaskMenuExtensionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQDesignerTaskMenuExtension_PreferredEditAction
func callbackQDesignerTaskMenuExtension_PreferredEditAction(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "preferredEditAction"); signal != nil {
		return widgets.PointerFromQAction(signal.(func() *widgets.QAction)())
	}

	return widgets.PointerFromQAction(NewQDesignerTaskMenuExtensionFromPointer(ptr).PreferredEditActionDefault())
}

func (ptr *QDesignerTaskMenuExtension) ConnectPreferredEditAction(f func() *widgets.QAction) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "preferredEditAction"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "preferredEditAction", func() *widgets.QAction {
				signal.(func() *widgets.QAction)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "preferredEditAction", f)
		}
	}
}

func (ptr *QDesignerTaskMenuExtension) DisconnectPreferredEditAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "preferredEditAction")
	}
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditAction() *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerTaskMenuExtension_PreferredEditAction(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) PreferredEditActionDefault() *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerTaskMenuExtension_PreferredEditActionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQDesignerTaskMenuExtension_TaskActions
func callbackQDesignerTaskMenuExtension_TaskActions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "taskActions"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQDesignerTaskMenuExtensionFromPointer(NewQDesignerTaskMenuExtensionFromPointer(nil).__taskActions_newList())
			for _, v := range signal.(func() []*widgets.QAction)() {
				tmpList.__taskActions_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQDesignerTaskMenuExtensionFromPointer(NewQDesignerTaskMenuExtensionFromPointer(nil).__taskActions_newList())
		for _, v := range make([]*widgets.QAction, 0) {
			tmpList.__taskActions_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QDesignerTaskMenuExtension) ConnectTaskActions(f func() []*widgets.QAction) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "taskActions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "taskActions", func() []*widgets.QAction {
				signal.(func() []*widgets.QAction)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "taskActions", f)
		}
	}
}

func (ptr *QDesignerTaskMenuExtension) DisconnectTaskActions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "taskActions")
	}
}

func (ptr *QDesignerTaskMenuExtension) TaskActions() []*widgets.QAction {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDesigner_PackedList) []*widgets.QAction {
			var out = make([]*widgets.QAction, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQDesignerTaskMenuExtensionFromPointer(l.data).__taskActions_atList(i)
			}
			return out
		}(C.QDesignerTaskMenuExtension_TaskActions(ptr.Pointer()))
	}
	return make([]*widgets.QAction, 0)
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerTaskMenuExtension___taskActions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerTaskMenuExtension___taskActions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerTaskMenuExtension) __taskActions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerTaskMenuExtension___taskActions_newList(ptr.Pointer()))
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

func NewQDesignerWidgetBoxInterfaceFromPointer(ptr unsafe.Pointer) *QDesignerWidgetBoxInterface {
	var n = new(QDesignerWidgetBoxInterface)
	n.SetPointer(ptr)
	return n
}

//export callbackQDesignerWidgetBoxInterface_Load
func callbackQDesignerWidgetBoxInterface_Load(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "load"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectLoad(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "load"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "load", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "load", f)
		}
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectLoad() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "load")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Load() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_Load(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_Save
func callbackQDesignerWidgetBoxInterface_Save(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "save"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QDesignerWidgetBoxInterface) ConnectSave(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "save"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "save", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "save", f)
		}
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSave() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "save")
	}
}

func (ptr *QDesignerWidgetBoxInterface) Save() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_Save(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_SetFileName
func callbackQDesignerWidgetBoxInterface_SetFileName(ptr unsafe.Pointer, fileName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setFileName"); signal != nil {
		signal.(func(string))(cGoUnpackString(fileName))
	}

}

func (ptr *QDesignerWidgetBoxInterface) ConnectSetFileName(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", func(fileName string) {
				signal.(func(string))(fileName)
				f(fileName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", f)
		}
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectSetFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFileName")
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QDesignerWidgetBoxInterface_SetFileName(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface
func callbackQDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QDesignerWidgetBoxInterface"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DestroyQDesignerWidgetBoxInterfaceDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectDestroyQDesignerWidgetBoxInterface(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QDesignerWidgetBoxInterface"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerWidgetBoxInterface", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QDesignerWidgetBoxInterface", f)
		}
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectDestroyQDesignerWidgetBoxInterface() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QDesignerWidgetBoxInterface")
	}
}

func (ptr *QDesignerWidgetBoxInterface) DestroyQDesignerWidgetBoxInterface() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterface(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QDesignerWidgetBoxInterface) DestroyQDesignerWidgetBoxInterfaceDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DestroyQDesignerWidgetBoxInterfaceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerWidgetBoxInterface_FileName
func callbackQDesignerWidgetBoxInterface_FileName(ptr unsafe.Pointer) C.struct_QtDesigner_PackedString {
	if signal := qt.GetSignal(ptr, "fileName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtDesigner_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectFileName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fileName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "fileName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileName", f)
		}
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fileName")
	}
}

func (ptr *QDesignerWidgetBoxInterface) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QDesignerWidgetBoxInterface_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDesignerWidgetBoxInterface) __dropWidgets_item_list_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___dropWidgets_item_list_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerWidgetBoxInterface___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___addActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __addActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___addActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerWidgetBoxInterface___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___insertActions_actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __insertActions_actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___insertActions_actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __actions_atList(i int) *widgets.QAction {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQActionFromPointer(C.QDesignerWidgetBoxInterface___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __actions_setList(i widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___actions_setList(ptr.Pointer(), widgets.PointerFromQAction(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __actions_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___actions_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QDesignerWidgetBoxInterface___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerWidgetBoxInterface___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerWidgetBoxInterface___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerWidgetBoxInterface___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___findChildren_newList(ptr.Pointer()))
}

func (ptr *QDesignerWidgetBoxInterface) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QDesignerWidgetBoxInterface___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QDesignerWidgetBoxInterface) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QDesignerWidgetBoxInterface) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QDesignerWidgetBoxInterface___children_newList(ptr.Pointer()))
}

//export callbackQDesignerWidgetBoxInterface_Close
func callbackQDesignerWidgetBoxInterface_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CloseDefault())))
}

func (ptr *QDesignerWidgetBoxInterface) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_Event
func callbackQDesignerWidgetBoxInterface_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerWidgetBoxInterface) EventDefault(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_FocusNextPrevChild
func callbackQDesignerWidgetBoxInterface_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QDesignerWidgetBoxInterface) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_NativeEvent
func callbackQDesignerWidgetBoxInterface_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result C.long) C.char {
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QByteArray, unsafe.Pointer, int) bool)(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).NativeEventDefault(core.NewQByteArrayFromPointer(eventType), message, int(int32(result))))))
}

func (ptr *QDesignerWidgetBoxInterface) NativeEventDefault(eventType core.QByteArray_ITF, message unsafe.Pointer, result int) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_NativeEventDefault(ptr.Pointer(), core.PointerFromQByteArray(eventType), message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_ActionEvent
func callbackQDesignerWidgetBoxInterface_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ActionEventDefault(event gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ChangeEvent
func callbackQDesignerWidgetBoxInterface_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChangeEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_CloseEvent
func callbackQDesignerWidgetBoxInterface_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) CloseEventDefault(event gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ContextMenuEvent
func callbackQDesignerWidgetBoxInterface_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_CustomContextMenuRequested
func callbackQDesignerWidgetBoxInterface_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

//export callbackQDesignerWidgetBoxInterface_DragEnterEvent
func callbackQDesignerWidgetBoxInterface_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DragLeaveEvent
func callbackQDesignerWidgetBoxInterface_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DragMoveEvent
func callbackQDesignerWidgetBoxInterface_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DropEvent
func callbackQDesignerWidgetBoxInterface_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DropEventDefault(event gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_EnterEvent
func callbackQDesignerWidgetBoxInterface_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) EnterEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_FocusInEvent
func callbackQDesignerWidgetBoxInterface_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_FocusOutEvent
func callbackQDesignerWidgetBoxInterface_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Hide
func callbackQDesignerWidgetBoxInterface_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HideDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) HideDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_HideDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_HideEvent
func callbackQDesignerWidgetBoxInterface_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) HideEventDefault(event gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_InputMethodEvent
func callbackQDesignerWidgetBoxInterface_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_KeyPressEvent
func callbackQDesignerWidgetBoxInterface_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_KeyReleaseEvent
func callbackQDesignerWidgetBoxInterface_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_LeaveEvent
func callbackQDesignerWidgetBoxInterface_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) LeaveEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Lower
func callbackQDesignerWidgetBoxInterface_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) LowerDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_LowerDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent
func callbackQDesignerWidgetBoxInterface_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseMoveEvent
func callbackQDesignerWidgetBoxInterface_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MousePressEvent
func callbackQDesignerWidgetBoxInterface_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MouseReleaseEvent
func callbackQDesignerWidgetBoxInterface_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MoveEvent
func callbackQDesignerWidgetBoxInterface_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) MoveEventDefault(event gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_PaintEvent
func callbackQDesignerWidgetBoxInterface_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) PaintEventDefault(event gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Raise
func callbackQDesignerWidgetBoxInterface_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_Repaint
func callbackQDesignerWidgetBoxInterface_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ResizeEvent
func callbackQDesignerWidgetBoxInterface_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetDisabled
func callbackQDesignerWidgetBoxInterface_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetEnabled
func callbackQDesignerWidgetBoxInterface_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetFocus2
func callbackQDesignerWidgetBoxInterface_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_SetHidden
func callbackQDesignerWidgetBoxInterface_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetStyleSheet
func callbackQDesignerWidgetBoxInterface_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.QDesignerWidgetBoxInterface_SetStyleSheetDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackQDesignerWidgetBoxInterface_SetVisible
func callbackQDesignerWidgetBoxInterface_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetWindowModified
func callbackQDesignerWidgetBoxInterface_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQDesignerWidgetBoxInterface_SetWindowTitle
func callbackQDesignerWidgetBoxInterface_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *QDesignerWidgetBoxInterface) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.QDesignerWidgetBoxInterface_SetWindowTitleDefault(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackQDesignerWidgetBoxInterface_Show
func callbackQDesignerWidgetBoxInterface_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowEvent
func callbackQDesignerWidgetBoxInterface_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowEventDefault(event gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowFullScreen
func callbackQDesignerWidgetBoxInterface_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowMaximized
func callbackQDesignerWidgetBoxInterface_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowMinimized
func callbackQDesignerWidgetBoxInterface_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_ShowNormal
func callbackQDesignerWidgetBoxInterface_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_TabletEvent
func callbackQDesignerWidgetBoxInterface_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) TabletEventDefault(event gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_Update
func callbackQDesignerWidgetBoxInterface_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_UpdateMicroFocus
func callbackQDesignerWidgetBoxInterface_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQDesignerWidgetBoxInterface_WheelEvent
func callbackQDesignerWidgetBoxInterface_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) WheelEventDefault(event gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_WindowIconChanged
func callbackQDesignerWidgetBoxInterface_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

//export callbackQDesignerWidgetBoxInterface_WindowTitleChanged
func callbackQDesignerWidgetBoxInterface_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackQDesignerWidgetBoxInterface_PaintEngine
func callbackQDesignerWidgetBoxInterface_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QDesignerWidgetBoxInterface) PaintEngineDefault() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QDesignerWidgetBoxInterface_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_MinimumSizeHint
func callbackQDesignerWidgetBoxInterface_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QDesignerWidgetBoxInterface) MinimumSizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_SizeHint
func callbackQDesignerWidgetBoxInterface_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).SizeHintDefault())
}

func (ptr *QDesignerWidgetBoxInterface) SizeHintDefault() *core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QDesignerWidgetBoxInterface_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_InputMethodQuery
func callbackQDesignerWidgetBoxInterface_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QDesignerWidgetBoxInterface) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QDesignerWidgetBoxInterface_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQDesignerWidgetBoxInterface_HasHeightForWidth
func callbackQDesignerWidgetBoxInterface_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QDesignerWidgetBoxInterface) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_HeightForWidth
func callbackQDesignerWidgetBoxInterface_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QDesignerWidgetBoxInterface) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerWidgetBoxInterface_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQDesignerWidgetBoxInterface_Metric
func callbackQDesignerWidgetBoxInterface_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *QDesignerWidgetBoxInterface) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QDesignerWidgetBoxInterface_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackQDesignerWidgetBoxInterface_EventFilter
func callbackQDesignerWidgetBoxInterface_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDesignerWidgetBoxInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QDesignerWidgetBoxInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDesignerWidgetBoxInterface_ChildEvent
func callbackQDesignerWidgetBoxInterface_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_ConnectNotify
func callbackQDesignerWidgetBoxInterface_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerWidgetBoxInterface_CustomEvent
func callbackQDesignerWidgetBoxInterface_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_DeleteLater
func callbackQDesignerWidgetBoxInterface_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDesignerWidgetBoxInterface) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQDesignerWidgetBoxInterface_Destroyed
func callbackQDesignerWidgetBoxInterface_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQDesignerWidgetBoxInterface_DisconnectNotify
func callbackQDesignerWidgetBoxInterface_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDesignerWidgetBoxInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDesignerWidgetBoxInterface_ObjectNameChanged
func callbackQDesignerWidgetBoxInterface_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQDesignerWidgetBoxInterface_TimerEvent
func callbackQDesignerWidgetBoxInterface_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDesignerWidgetBoxInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDesignerWidgetBoxInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QDesignerWidgetBoxInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDesignerWidgetBoxInterface_MetaObject
func callbackQDesignerWidgetBoxInterface_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDesignerWidgetBoxInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDesignerWidgetBoxInterface) MetaObjectDefault() *core.QMetaObject {
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

func NewQExtensionFactoryFromPointer(ptr unsafe.Pointer) *QExtensionFactory {
	var n = new(QExtensionFactory)
	n.SetPointer(ptr)
	return n
}
func NewQExtensionFactory(parent QExtensionManager_ITF) *QExtensionFactory {
	var tmpValue = NewQExtensionFactoryFromPointer(C.QExtensionFactory_NewQExtensionFactory(PointerFromQExtensionManager(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QExtensionFactory) ExtensionManager() *QExtensionManager {
	if ptr.Pointer() != nil {
		var tmpValue = NewQExtensionManagerFromPointer(C.QExtensionFactory_ExtensionManager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQExtensionFactory_CreateExtension
func callbackQExtensionFactory_CreateExtension(ptr unsafe.Pointer, object unsafe.Pointer, iid C.struct_QtDesigner_PackedString, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createExtension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string, *core.QObject) *core.QObject)(core.NewQObjectFromPointer(object), cGoUnpackString(iid), core.NewQObjectFromPointer(parent)))
	}

	return core.PointerFromQObject(NewQExtensionFactoryFromPointer(ptr).CreateExtensionDefault(core.NewQObjectFromPointer(object), cGoUnpackString(iid), core.NewQObjectFromPointer(parent)))
}

func (ptr *QExtensionFactory) ConnectCreateExtension(f func(object *core.QObject, iid string, parent *core.QObject) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createExtension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createExtension", func(object *core.QObject, iid string, parent *core.QObject) *core.QObject {
				signal.(func(*core.QObject, string, *core.QObject) *core.QObject)(object, iid, parent)
				return f(object, iid, parent)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createExtension", f)
		}
	}
}

func (ptr *QExtensionFactory) DisconnectCreateExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createExtension")
	}
}

func (ptr *QExtensionFactory) CreateExtension(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_CreateExtension(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}, core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) CreateExtensionDefault(object core.QObject_ITF, iid string, parent core.QObject_ITF) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_CreateExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}, core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQExtensionFactory_Extension
func callbackQExtensionFactory_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid C.struct_QtDesigner_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), cGoUnpackString(iid)))
	}

	return core.PointerFromQObject(NewQExtensionFactoryFromPointer(ptr).ExtensionDefault(core.NewQObjectFromPointer(object), cGoUnpackString(iid)))
}

func (ptr *QExtensionFactory) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "extension", func(object *core.QObject, iid string) *core.QObject {
				signal.(func(*core.QObject, string) *core.QObject)(object, iid)
				return f(object, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", f)
		}
	}
}

func (ptr *QExtensionFactory) DisconnectExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "extension")
	}
}

func (ptr *QExtensionFactory) Extension(object core.QObject_ITF, iid string) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_Extension(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory_ExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QExtensionFactory___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QExtensionFactory) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionFactory___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QExtensionFactory) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionFactory) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionFactory___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QExtensionFactory) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionFactory) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionFactory___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QExtensionFactory) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionFactory) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionFactory___findChildren_newList(ptr.Pointer()))
}

func (ptr *QExtensionFactory) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionFactory___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionFactory) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionFactory) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionFactory___children_newList(ptr.Pointer()))
}

//export callbackQExtensionFactory_Event
func callbackQExtensionFactory_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QExtensionFactory) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QExtensionFactory) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQExtensionFactory_EventFilter
func callbackQExtensionFactory_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QExtensionFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QExtensionFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQExtensionFactory_ChildEvent
func callbackQExtensionFactory_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QExtensionFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQExtensionFactory_ConnectNotify
func callbackQExtensionFactory_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionFactory_CustomEvent
func callbackQExtensionFactory_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QExtensionFactory) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQExtensionFactory_DeleteLater
func callbackQExtensionFactory_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQExtensionFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QExtensionFactory) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QExtensionFactory) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQExtensionFactory_Destroyed
func callbackQExtensionFactory_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQExtensionFactory_DisconnectNotify
func callbackQExtensionFactory_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionFactory_ObjectNameChanged
func callbackQExtensionFactory_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQExtensionFactory_TimerEvent
func callbackQExtensionFactory_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQExtensionFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QExtensionFactory) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QExtensionFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQExtensionFactory_MetaObject
func callbackQExtensionFactory_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQExtensionFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QExtensionFactory) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QExtensionFactory) MetaObjectDefault() *core.QMetaObject {
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

func NewQExtensionManagerFromPointer(ptr unsafe.Pointer) *QExtensionManager {
	var n = new(QExtensionManager)
	n.SetPointer(ptr)
	return n
}
func NewQExtensionManager(parent core.QObject_ITF) *QExtensionManager {
	var tmpValue = NewQExtensionManagerFromPointer(C.QExtensionManager_NewQExtensionManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQExtensionManager_RegisterExtensions
func callbackQExtensionManager_RegisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "registerExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	} else {
		NewQExtensionManagerFromPointer(ptr).RegisterExtensionsDefault(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	}
}

func (ptr *QExtensionManager) ConnectRegisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerExtensions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "registerExtensions", func(factory *QAbstractExtensionFactory, iid string) {
				signal.(func(*QAbstractExtensionFactory, string))(factory, iid)
				f(factory, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerExtensions", f)
		}
	}
}

func (ptr *QExtensionManager) DisconnectRegisterExtensions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerExtensions")
	}
}

func (ptr *QExtensionManager) RegisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QExtensionManager_RegisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

func (ptr *QExtensionManager) RegisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QExtensionManager_RegisterExtensionsDefault(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

//export callbackQExtensionManager_UnregisterExtensions
func callbackQExtensionManager_UnregisterExtensions(ptr unsafe.Pointer, factory unsafe.Pointer, iid C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "unregisterExtensions"); signal != nil {
		signal.(func(*QAbstractExtensionFactory, string))(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	} else {
		NewQExtensionManagerFromPointer(ptr).UnregisterExtensionsDefault(NewQAbstractExtensionFactoryFromPointer(factory), cGoUnpackString(iid))
	}
}

func (ptr *QExtensionManager) ConnectUnregisterExtensions(f func(factory *QAbstractExtensionFactory, iid string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "unregisterExtensions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "unregisterExtensions", func(factory *QAbstractExtensionFactory, iid string) {
				signal.(func(*QAbstractExtensionFactory, string))(factory, iid)
				f(factory, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "unregisterExtensions", f)
		}
	}
}

func (ptr *QExtensionManager) DisconnectUnregisterExtensions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "unregisterExtensions")
	}
}

func (ptr *QExtensionManager) UnregisterExtensions(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QExtensionManager_UnregisterExtensions(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

func (ptr *QExtensionManager) UnregisterExtensionsDefault(factory QAbstractExtensionFactory_ITF, iid string) {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		C.QExtensionManager_UnregisterExtensionsDefault(ptr.Pointer(), PointerFromQAbstractExtensionFactory(factory), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))})
	}
}

func (ptr *QExtensionManager) DestroyQExtensionManager() {
	if ptr.Pointer() != nil {
		C.QExtensionManager_DestroyQExtensionManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQExtensionManager_Extension
func callbackQExtensionManager_Extension(ptr unsafe.Pointer, object unsafe.Pointer, iid C.struct_QtDesigner_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "extension"); signal != nil {
		return core.PointerFromQObject(signal.(func(*core.QObject, string) *core.QObject)(core.NewQObjectFromPointer(object), cGoUnpackString(iid)))
	}

	return core.PointerFromQObject(NewQExtensionManagerFromPointer(ptr).ExtensionDefault(core.NewQObjectFromPointer(object), cGoUnpackString(iid)))
}

func (ptr *QExtensionManager) ConnectExtension(f func(object *core.QObject, iid string) *core.QObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "extension"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "extension", func(object *core.QObject, iid string) *core.QObject {
				signal.(func(*core.QObject, string) *core.QObject)(object, iid)
				return f(object, iid)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "extension", f)
		}
	}
}

func (ptr *QExtensionManager) DisconnectExtension() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "extension")
	}
}

func (ptr *QExtensionManager) Extension(object core.QObject_ITF, iid string) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager_Extension(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) ExtensionDefault(object core.QObject_ITF, iid string) *core.QObject {
	if ptr.Pointer() != nil {
		var iidC *C.char
		if iid != "" {
			iidC = C.CString(iid)
			defer C.free(unsafe.Pointer(iidC))
		}
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager_ExtensionDefault(ptr.Pointer(), core.PointerFromQObject(object), C.struct_QtDesigner_PackedString{data: iidC, len: C.longlong(len(iid))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QExtensionManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QExtensionManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionManager___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QExtensionManager) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionManager) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionManager___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QExtensionManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionManager) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionManager___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QExtensionManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionManager) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionManager___findChildren_newList(ptr.Pointer()))
}

func (ptr *QExtensionManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QExtensionManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QExtensionManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QExtensionManager) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QExtensionManager___children_newList(ptr.Pointer()))
}

//export callbackQExtensionManager_Event
func callbackQExtensionManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QExtensionManager) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QExtensionManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQExtensionManager_EventFilter
func callbackQExtensionManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQExtensionManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QExtensionManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QExtensionManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QExtensionManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQExtensionManager_ChildEvent
func callbackQExtensionManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QExtensionManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQExtensionManager_ConnectNotify
func callbackQExtensionManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionManager_CustomEvent
func callbackQExtensionManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QExtensionManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQExtensionManager_DeleteLater
func callbackQExtensionManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQExtensionManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QExtensionManager) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QExtensionManager_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QExtensionManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QExtensionManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQExtensionManager_Destroyed
func callbackQExtensionManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQExtensionManager_DisconnectNotify
func callbackQExtensionManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQExtensionManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QExtensionManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QExtensionManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQExtensionManager_ObjectNameChanged
func callbackQExtensionManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtDesigner_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQExtensionManager_TimerEvent
func callbackQExtensionManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQExtensionManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QExtensionManager) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QExtensionManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QExtensionManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQExtensionManager_MetaObject
func callbackQExtensionManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQExtensionManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QExtensionManager) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QExtensionManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QExtensionManager) MetaObjectDefault() *core.QMetaObject {
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

func NewQFormBuilderFromPointer(ptr unsafe.Pointer) *QFormBuilder {
	var n = new(QFormBuilder)
	n.SetPointer(ptr)
	return n
}
func NewQFormBuilder() *QFormBuilder {
	return NewQFormBuilderFromPointer(C.QFormBuilder_NewQFormBuilder())
}

func (ptr *QFormBuilder) AddPluginPath(pluginPath string) {
	if ptr.Pointer() != nil {
		var pluginPathC *C.char
		if pluginPath != "" {
			pluginPathC = C.CString(pluginPath)
			defer C.free(unsafe.Pointer(pluginPathC))
		}
		C.QFormBuilder_AddPluginPath(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pluginPathC, len: C.longlong(len(pluginPath))})
	}
}

func (ptr *QFormBuilder) ClearPluginPaths() {
	if ptr.Pointer() != nil {
		C.QFormBuilder_ClearPluginPaths(ptr.Pointer())
	}
}

func (ptr *QFormBuilder) SetPluginPath(pluginPaths []string) {
	if ptr.Pointer() != nil {
		var pluginPathsC = C.CString(strings.Join(pluginPaths, "|"))
		defer C.free(unsafe.Pointer(pluginPathsC))
		C.QFormBuilder_SetPluginPath(ptr.Pointer(), C.struct_QtDesigner_PackedString{data: pluginPathsC, len: C.longlong(len(strings.Join(pluginPaths, "|")))})
	}
}

//export callbackQFormBuilder_DestroyQFormBuilder
func callbackQFormBuilder_DestroyQFormBuilder(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QFormBuilder"); signal != nil {
		signal.(func())()
	} else {
		NewQFormBuilderFromPointer(ptr).DestroyQFormBuilderDefault()
	}
}

func (ptr *QFormBuilder) ConnectDestroyQFormBuilder(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QFormBuilder"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QFormBuilder", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QFormBuilder", f)
		}
	}
}

func (ptr *QFormBuilder) DisconnectDestroyQFormBuilder() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QFormBuilder")
	}
}

func (ptr *QFormBuilder) DestroyQFormBuilder() {
	if ptr.Pointer() != nil {
		C.QFormBuilder_DestroyQFormBuilder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFormBuilder) DestroyQFormBuilderDefault() {
	if ptr.Pointer() != nil {
		C.QFormBuilder_DestroyQFormBuilderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFormBuilder) CustomWidgets() []*QDesignerCustomWidgetInterface {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtDesigner_PackedList) []*QDesignerCustomWidgetInterface {
			var out = make([]*QDesignerCustomWidgetInterface, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQFormBuilderFromPointer(l.data).__customWidgets_atList(i)
			}
			return out
		}(C.QFormBuilder_CustomWidgets(ptr.Pointer()))
	}
	return make([]*QDesignerCustomWidgetInterface, 0)
}

func (ptr *QFormBuilder) PluginPaths() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QFormBuilder_PluginPaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFormBuilder) __customWidgets_atList(i int) *QDesignerCustomWidgetInterface {
	if ptr.Pointer() != nil {
		return NewQDesignerCustomWidgetInterfaceFromPointer(C.QFormBuilder___customWidgets_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QFormBuilder) __customWidgets_setList(i QDesignerCustomWidgetInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QFormBuilder___customWidgets_setList(ptr.Pointer(), PointerFromQDesignerCustomWidgetInterface(i))
	}
}

func (ptr *QFormBuilder) __customWidgets_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QFormBuilder___customWidgets_newList(ptr.Pointer()))
}
