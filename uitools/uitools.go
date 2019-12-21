// +build !minimal

package uitools

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "uitools.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtUiTools_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtUiTools_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

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

func NewQUiLoaderFromPointer(ptr unsafe.Pointer) (n *QUiLoader) {
	n = new(QUiLoader)
	n.SetPointer(ptr)
	return
}
func NewQUiLoader(parent core.QObject_ITF) *QUiLoader {
	tmpValue := NewQUiLoaderFromPointer(C.QUiLoader_NewQUiLoader(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QUiLoader) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QUiLoader_AddPluginPath(ptr.Pointer(), C.struct_QtUiTools_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

func (ptr *QUiLoader) AvailableLayouts() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QUiLoader_AvailableLayouts(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) AvailableWidgets() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QUiLoader_AvailableWidgets(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) ClearPluginPaths() {
	if ptr.Pointer() != nil {
		C.QUiLoader_ClearPluginPaths(ptr.Pointer())
	}
}

//export callbackQUiLoader_CreateAction
func callbackQUiLoader_CreateAction(ptr unsafe.Pointer, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createAction"); signal != nil {
		return widgets.PointerFromQAction((*(*func(*core.QObject, string) *widgets.QAction)(signal))(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQAction(NewQUiLoaderFromPointer(ptr).CreateActionDefault(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateAction(f func(parent *core.QObject, name string) *widgets.QAction) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createAction"); signal != nil {
			f := func(parent *core.QObject, name string) *widgets.QAction {
				(*(*func(*core.QObject, string) *widgets.QAction)(signal))(parent, name)
				return f(parent, name)
			}
			qt.ConnectSignal(ptr.Pointer(), "createAction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createAction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUiLoader) DisconnectCreateAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createAction")
	}
}

func (ptr *QUiLoader) CreateAction(parent core.QObject_ITF, name string) *widgets.QAction {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQActionFromPointer(C.QUiLoader_CreateAction(ptr.Pointer(), core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateActionDefault(parent core.QObject_ITF, name string) *widgets.QAction {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQActionFromPointer(C.QUiLoader_CreateActionDefault(ptr.Pointer(), core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateActionGroup
func callbackQUiLoader_CreateActionGroup(ptr unsafe.Pointer, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createActionGroup"); signal != nil {
		return widgets.PointerFromQActionGroup((*(*func(*core.QObject, string) *widgets.QActionGroup)(signal))(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQActionGroup(NewQUiLoaderFromPointer(ptr).CreateActionGroupDefault(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateActionGroup(f func(parent *core.QObject, name string) *widgets.QActionGroup) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createActionGroup"); signal != nil {
			f := func(parent *core.QObject, name string) *widgets.QActionGroup {
				(*(*func(*core.QObject, string) *widgets.QActionGroup)(signal))(parent, name)
				return f(parent, name)
			}
			qt.ConnectSignal(ptr.Pointer(), "createActionGroup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createActionGroup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUiLoader) DisconnectCreateActionGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createActionGroup")
	}
}

func (ptr *QUiLoader) CreateActionGroup(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroup(ptr.Pointer(), core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateActionGroupDefault(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroupDefault(ptr.Pointer(), core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateLayout
func callbackQUiLoader_CreateLayout(ptr unsafe.Pointer, className C.struct_QtUiTools_PackedString, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createLayout"); signal != nil {
		return widgets.PointerFromQLayout((*(*func(string, *core.QObject, string) *widgets.QLayout)(signal))(cGoUnpackString(className), core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQLayout(NewQUiLoaderFromPointer(ptr).CreateLayoutDefault(cGoUnpackString(className), core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateLayout(f func(className string, parent *core.QObject, name string) *widgets.QLayout) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createLayout"); signal != nil {
			f := func(className string, parent *core.QObject, name string) *widgets.QLayout {
				(*(*func(string, *core.QObject, string) *widgets.QLayout)(signal))(className, parent, name)
				return f(className, parent, name)
			}
			qt.ConnectSignal(ptr.Pointer(), "createLayout", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createLayout", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUiLoader) DisconnectCreateLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createLayout")
	}
}

func (ptr *QUiLoader) CreateLayout(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	if ptr.Pointer() != nil {
		var classNameC *C.char
		if className != "" {
			classNameC = C.CString(className)
			defer C.free(unsafe.Pointer(classNameC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayout(ptr.Pointer(), C.struct_QtUiTools_PackedString{data: classNameC, len: C.longlong(len(className))}, core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateLayoutDefault(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	if ptr.Pointer() != nil {
		var classNameC *C.char
		if className != "" {
			classNameC = C.CString(className)
			defer C.free(unsafe.Pointer(classNameC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayoutDefault(ptr.Pointer(), C.struct_QtUiTools_PackedString{data: classNameC, len: C.longlong(len(className))}, core.PointerFromQObject(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateWidget
func callbackQUiLoader_CreateWidget(ptr unsafe.Pointer, className C.struct_QtUiTools_PackedString, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createWidget"); signal != nil {
		return widgets.PointerFromQWidget((*(*func(string, *widgets.QWidget, string) *widgets.QWidget)(signal))(cGoUnpackString(className), widgets.NewQWidgetFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQWidget(NewQUiLoaderFromPointer(ptr).CreateWidgetDefault(cGoUnpackString(className), widgets.NewQWidgetFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateWidget(f func(className string, parent *widgets.QWidget, name string) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createWidget"); signal != nil {
			f := func(className string, parent *widgets.QWidget, name string) *widgets.QWidget {
				(*(*func(string, *widgets.QWidget, string) *widgets.QWidget)(signal))(className, parent, name)
				return f(className, parent, name)
			}
			qt.ConnectSignal(ptr.Pointer(), "createWidget", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createWidget", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUiLoader) DisconnectCreateWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createWidget")
	}
}

func (ptr *QUiLoader) CreateWidget(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var classNameC *C.char
		if className != "" {
			classNameC = C.CString(className)
			defer C.free(unsafe.Pointer(classNameC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidget(ptr.Pointer(), C.struct_QtUiTools_PackedString{data: classNameC, len: C.longlong(len(className))}, widgets.PointerFromQWidget(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateWidgetDefault(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var classNameC *C.char
		if className != "" {
			classNameC = C.CString(className)
			defer C.free(unsafe.Pointer(classNameC))
		}
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidgetDefault(ptr.Pointer(), C.struct_QtUiTools_PackedString{data: classNameC, len: C.longlong(len(className))}, widgets.PointerFromQWidget(parent), C.struct_QtUiTools_PackedString{data: nameC, len: C.longlong(len(name))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QUiLoader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUiLoader) IsLanguageChangeEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QUiLoader_IsLanguageChangeEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUiLoader) Load(device core.QIODevice_ITF, parentWidget widgets.QWidget_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := widgets.NewQWidgetFromPointer(C.QUiLoader_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parentWidget)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) PluginPaths() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QUiLoader_PluginPaths(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) SetLanguageChangeEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QUiLoader_SetLanguageChangeEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QUiLoader) SetWorkingDirectory(dir core.QDir_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_SetWorkingDirectory(ptr.Pointer(), core.PointerFromQDir(dir))
	}
}

func (ptr *QUiLoader) WorkingDirectory() *core.QDir {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQDirFromPointer(C.QUiLoader_WorkingDirectory(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_DestroyQUiLoader
func callbackQUiLoader_DestroyQUiLoader(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QUiLoader"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQUiLoaderFromPointer(ptr).DestroyQUiLoaderDefault()
	}
}

func (ptr *QUiLoader) ConnectDestroyQUiLoader(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QUiLoader"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QUiLoader", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QUiLoader", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QUiLoader) DisconnectDestroyQUiLoader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QUiLoader")
	}
}

func (ptr *QUiLoader) DestroyQUiLoader() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUiLoader_DestroyQUiLoader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUiLoader) DestroyQUiLoaderDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUiLoader_DestroyQUiLoaderDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUiLoader) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QUiLoader___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QUiLoader) __children_newList() unsafe.Pointer {
	return C.QUiLoader___children_newList(ptr.Pointer())
}

func (ptr *QUiLoader) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QUiLoader___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QUiLoader) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QUiLoader___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QUiLoader) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QUiLoader___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QUiLoader) __findChildren_newList() unsafe.Pointer {
	return C.QUiLoader___findChildren_newList(ptr.Pointer())
}

func (ptr *QUiLoader) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QUiLoader___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QUiLoader) __findChildren_newList3() unsafe.Pointer {
	return C.QUiLoader___findChildren_newList3(ptr.Pointer())
}

//export callbackQUiLoader_ChildEvent
func callbackQUiLoader_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQUiLoader_ConnectNotify
func callbackQUiLoader_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_CustomEvent
func callbackQUiLoader_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUiLoader) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQUiLoader_DeleteLater
func callbackQUiLoader_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQUiLoaderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QUiLoader) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QUiLoader_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQUiLoader_Destroyed
func callbackQUiLoader_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQUiLoader_DisconnectNotify
func callbackQUiLoader_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_Event
func callbackQUiLoader_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QUiLoader) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUiLoader_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQUiLoader_EventFilter
func callbackQUiLoader_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QUiLoader) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QUiLoader_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQUiLoader_MetaObject
func callbackQUiLoader_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQUiLoaderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QUiLoader) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUiLoader_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQUiLoader_ObjectNameChanged
func callbackQUiLoader_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtUiTools_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQUiLoader_TimerEvent
func callbackQUiLoader_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUiLoader) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["uitools.QUiLoader_ITF"] = QUiLoader{}
	qt.FuncMap["uitools.NewQUiLoader"] = NewQUiLoader
}
