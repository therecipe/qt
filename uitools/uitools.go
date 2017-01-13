// +build !minimal

package uitools

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "uitools.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtUiTools_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQUiLoaderFromPointer(ptr unsafe.Pointer) *QUiLoader {
	var n = new(QUiLoader)
	n.SetPointer(ptr)
	return n
}
func NewQUiLoader(parent core.QObject_ITF) *QUiLoader {
	var tmpValue = NewQUiLoaderFromPointer(C.QUiLoader_NewQUiLoader(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QUiLoader) AddPluginPath(path string) {
	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QUiLoader_AddPluginPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QUiLoader) AvailableLayouts() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QUiLoader_AvailableLayouts(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) AvailableWidgets() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QUiLoader_AvailableWidgets(ptr.Pointer())), "|")
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

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::createAction"); signal != nil {
		return widgets.PointerFromQAction(signal.(func(*core.QObject, string) *widgets.QAction)(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQAction(NewQUiLoaderFromPointer(ptr).CreateActionDefault(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateAction(f func(parent *core.QObject, name string) *widgets.QAction) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createAction", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateAction() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createAction")
	}
}

func (ptr *QUiLoader) CreateAction(parent core.QObject_ITF, name string) *widgets.QAction {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQActionFromPointer(C.QUiLoader_CreateAction(ptr.Pointer(), core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateActionDefault(parent core.QObject_ITF, name string) *widgets.QAction {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQActionFromPointer(C.QUiLoader_CreateActionDefault(ptr.Pointer(), core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateActionGroup
func callbackQUiLoader_CreateActionGroup(ptr unsafe.Pointer, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::createActionGroup"); signal != nil {
		return widgets.PointerFromQActionGroup(signal.(func(*core.QObject, string) *widgets.QActionGroup)(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQActionGroup(NewQUiLoaderFromPointer(ptr).CreateActionGroupDefault(core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateActionGroup(f func(parent *core.QObject, name string) *widgets.QActionGroup) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createActionGroup", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateActionGroup() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createActionGroup")
	}
}

func (ptr *QUiLoader) CreateActionGroup(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroup(ptr.Pointer(), core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateActionGroupDefault(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroupDefault(ptr.Pointer(), core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateLayout
func callbackQUiLoader_CreateLayout(ptr unsafe.Pointer, className C.struct_QtUiTools_PackedString, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::createLayout"); signal != nil {
		return widgets.PointerFromQLayout(signal.(func(string, *core.QObject, string) *widgets.QLayout)(cGoUnpackString(className), core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQLayout(NewQUiLoaderFromPointer(ptr).CreateLayoutDefault(cGoUnpackString(className), core.NewQObjectFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateLayout(f func(className string, parent *core.QObject, name string) *widgets.QLayout) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createLayout", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateLayout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createLayout")
	}
}

func (ptr *QUiLoader) CreateLayout(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	if ptr.Pointer() != nil {
		var classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayout(ptr.Pointer(), classNameC, core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateLayoutDefault(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	if ptr.Pointer() != nil {
		var classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayoutDefault(ptr.Pointer(), classNameC, core.PointerFromQObject(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_CreateWidget
func callbackQUiLoader_CreateWidget(ptr unsafe.Pointer, className C.struct_QtUiTools_PackedString, parent unsafe.Pointer, name C.struct_QtUiTools_PackedString) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::createWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(string, *widgets.QWidget, string) *widgets.QWidget)(cGoUnpackString(className), widgets.NewQWidgetFromPointer(parent), cGoUnpackString(name)))
	}

	return widgets.PointerFromQWidget(NewQUiLoaderFromPointer(ptr).CreateWidgetDefault(cGoUnpackString(className), widgets.NewQWidgetFromPointer(parent), cGoUnpackString(name)))
}

func (ptr *QUiLoader) ConnectCreateWidget(f func(className string, parent *widgets.QWidget, name string) *widgets.QWidget) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createWidget", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateWidget() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::createWidget")
	}
}

func (ptr *QUiLoader) CreateWidget(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidget(ptr.Pointer(), classNameC, widgets.PointerFromQWidget(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) CreateWidgetDefault(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var classNameC = C.CString(className)
		defer C.free(unsafe.Pointer(classNameC))
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidgetDefault(ptr.Pointer(), classNameC, widgets.PointerFromQWidget(parent), nameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
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
		return C.QUiLoader_IsLanguageChangeEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUiLoader) Load(device core.QIODevice_ITF, parentWidget widgets.QWidget_ITF) *widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QUiLoader_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parentWidget)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QUiLoader) PluginPaths() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QUiLoader_PluginPaths(ptr.Pointer())), "|")
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
		var tmpValue = core.NewQDirFromPointer(C.QUiLoader_WorkingDirectory(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDir).DestroyQDir)
		return tmpValue
	}
	return nil
}

//export callbackQUiLoader_DestroyQUiLoader
func callbackQUiLoader_DestroyQUiLoader(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::~QUiLoader"); signal != nil {
		signal.(func())()
	} else {
		NewQUiLoaderFromPointer(ptr).DestroyQUiLoaderDefault()
	}
}

func (ptr *QUiLoader) ConnectDestroyQUiLoader(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::~QUiLoader", f)
	}
}

func (ptr *QUiLoader) DisconnectDestroyQUiLoader() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::~QUiLoader")
	}
}

func (ptr *QUiLoader) DestroyQUiLoader() {
	if ptr.Pointer() != nil {
		C.QUiLoader_DestroyQUiLoader(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUiLoader) DestroyQUiLoaderDefault() {
	if ptr.Pointer() != nil {
		C.QUiLoader_DestroyQUiLoaderDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUiLoader_TimerEvent
func callbackQUiLoader_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::timerEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::timerEvent")
	}
}

func (ptr *QUiLoader) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUiLoader) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQUiLoader_ChildEvent
func callbackQUiLoader_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::childEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::childEvent")
	}
}

func (ptr *QUiLoader) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUiLoader) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQUiLoader_ConnectNotify
func callbackQUiLoader_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::connectNotify", f)
	}
}

func (ptr *QUiLoader) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::connectNotify")
	}
}

func (ptr *QUiLoader) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUiLoader) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_CustomEvent
func callbackQUiLoader_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::customEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::customEvent")
	}
}

func (ptr *QUiLoader) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUiLoader) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQUiLoader_DeleteLater
func callbackQUiLoader_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQUiLoaderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QUiLoader) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::deleteLater", f)
	}
}

func (ptr *QUiLoader) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::deleteLater")
	}
}

func (ptr *QUiLoader) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QUiLoader_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUiLoader) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QUiLoader_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUiLoader_DisconnectNotify
func callbackQUiLoader_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::disconnectNotify", f)
	}
}

func (ptr *QUiLoader) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::disconnectNotify")
	}
}

func (ptr *QUiLoader) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUiLoader) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QUiLoader_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_Event
func callbackQUiLoader_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QUiLoader) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::event", f)
	}
}

func (ptr *QUiLoader) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::event")
	}
}

func (ptr *QUiLoader) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUiLoader_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QUiLoader) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUiLoader_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQUiLoader_EventFilter
func callbackQUiLoader_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QUiLoader) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::eventFilter", f)
	}
}

func (ptr *QUiLoader) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::eventFilter")
	}
}

func (ptr *QUiLoader) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUiLoader_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QUiLoader) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QUiLoader_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQUiLoader_MetaObject
func callbackQUiLoader_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUiLoader::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQUiLoaderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QUiLoader) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::metaObject", f)
	}
}

func (ptr *QUiLoader) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUiLoader::metaObject")
	}
}

func (ptr *QUiLoader) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUiLoader_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUiLoader) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUiLoader_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
