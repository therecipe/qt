package uitools

//#include "uitools.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

type QUiLoader struct {
	core.QObject
}

type QUiLoader_ITF interface {
	core.QObject_ITF
	QUiLoader_PTR() *QUiLoader
}

func (p *QUiLoader) QUiLoader_PTR() *QUiLoader {
	return p
}

func (p *QUiLoader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QUiLoader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
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

func newQUiLoaderFromPointer(ptr unsafe.Pointer) *QUiLoader {
	var n = NewQUiLoaderFromPointer(ptr)
	for len(n.ObjectName()) < len("QUiLoader_") {
		n.SetObjectName("QUiLoader_" + qt.Identifier())
	}
	return n
}

func NewQUiLoader(parent core.QObject_ITF) *QUiLoader {
	defer qt.Recovering("QUiLoader::QUiLoader")

	return newQUiLoaderFromPointer(C.QUiLoader_NewQUiLoader(core.PointerFromQObject(parent)))
}

func (ptr *QUiLoader) AddPluginPath(path string) {
	defer qt.Recovering("QUiLoader::addPluginPath")

	if ptr.Pointer() != nil {
		C.QUiLoader_AddPluginPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QUiLoader) AvailableLayouts() []string {
	defer qt.Recovering("QUiLoader::availableLayouts")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUiLoader_AvailableLayouts(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) AvailableWidgets() []string {
	defer qt.Recovering("QUiLoader::availableWidgets")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUiLoader_AvailableWidgets(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) ClearPluginPaths() {
	defer qt.Recovering("QUiLoader::clearPluginPaths")

	if ptr.Pointer() != nil {
		C.QUiLoader_ClearPluginPaths(ptr.Pointer())
	}
}

//export callbackQUiLoader_CreateAction
func callbackQUiLoader_CreateAction(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QUiLoader::createAction")

	if signal := qt.GetSignal(C.GoString(ptrName), "createAction"); signal != nil {
		return widgets.PointerFromQAction(signal.(func(*core.QObject, string) *widgets.QAction)(core.NewQObjectFromPointer(parent), C.GoString(name)))
	}

	return widgets.PointerFromQAction(NewQUiLoaderFromPointer(ptr).CreateActionDefault(core.NewQObjectFromPointer(parent), C.GoString(name)))
}

func (ptr *QUiLoader) ConnectCreateAction(f func(parent *core.QObject, name string) *widgets.QAction) {
	defer qt.Recovering("connect QUiLoader::createAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createAction", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateAction() {
	defer qt.Recovering("disconnect QUiLoader::createAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createAction")
	}
}

func (ptr *QUiLoader) CreateAction(parent core.QObject_ITF, name string) *widgets.QAction {
	defer qt.Recovering("QUiLoader::createAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QUiLoader_CreateAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateActionDefault(parent core.QObject_ITF, name string) *widgets.QAction {
	defer qt.Recovering("QUiLoader::createAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QUiLoader_CreateActionDefault(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

//export callbackQUiLoader_CreateActionGroup
func callbackQUiLoader_CreateActionGroup(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QUiLoader::createActionGroup")

	if signal := qt.GetSignal(C.GoString(ptrName), "createActionGroup"); signal != nil {
		return widgets.PointerFromQActionGroup(signal.(func(*core.QObject, string) *widgets.QActionGroup)(core.NewQObjectFromPointer(parent), C.GoString(name)))
	}

	return widgets.PointerFromQActionGroup(NewQUiLoaderFromPointer(ptr).CreateActionGroupDefault(core.NewQObjectFromPointer(parent), C.GoString(name)))
}

func (ptr *QUiLoader) ConnectCreateActionGroup(f func(parent *core.QObject, name string) *widgets.QActionGroup) {
	defer qt.Recovering("connect QUiLoader::createActionGroup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createActionGroup", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateActionGroup() {
	defer qt.Recovering("disconnect QUiLoader::createActionGroup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createActionGroup")
	}
}

func (ptr *QUiLoader) CreateActionGroup(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	defer qt.Recovering("QUiLoader::createActionGroup")

	if ptr.Pointer() != nil {
		return widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroup(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateActionGroupDefault(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	defer qt.Recovering("QUiLoader::createActionGroup")

	if ptr.Pointer() != nil {
		return widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroupDefault(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

//export callbackQUiLoader_CreateLayout
func callbackQUiLoader_CreateLayout(ptr unsafe.Pointer, ptrName *C.char, className *C.char, parent unsafe.Pointer, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QUiLoader::createLayout")

	if signal := qt.GetSignal(C.GoString(ptrName), "createLayout"); signal != nil {
		return widgets.PointerFromQLayout(signal.(func(string, *core.QObject, string) *widgets.QLayout)(C.GoString(className), core.NewQObjectFromPointer(parent), C.GoString(name)))
	}

	return widgets.PointerFromQLayout(NewQUiLoaderFromPointer(ptr).CreateLayoutDefault(C.GoString(className), core.NewQObjectFromPointer(parent), C.GoString(name)))
}

func (ptr *QUiLoader) ConnectCreateLayout(f func(className string, parent *core.QObject, name string) *widgets.QLayout) {
	defer qt.Recovering("connect QUiLoader::createLayout")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createLayout", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateLayout() {
	defer qt.Recovering("disconnect QUiLoader::createLayout")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createLayout")
	}
}

func (ptr *QUiLoader) CreateLayout(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	defer qt.Recovering("QUiLoader::createLayout")

	if ptr.Pointer() != nil {
		return widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayout(ptr.Pointer(), C.CString(className), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateLayoutDefault(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	defer qt.Recovering("QUiLoader::createLayout")

	if ptr.Pointer() != nil {
		return widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayoutDefault(ptr.Pointer(), C.CString(className), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

//export callbackQUiLoader_CreateWidget
func callbackQUiLoader_CreateWidget(ptr unsafe.Pointer, ptrName *C.char, className *C.char, parent unsafe.Pointer, name *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QUiLoader::createWidget")

	if signal := qt.GetSignal(C.GoString(ptrName), "createWidget"); signal != nil {
		return widgets.PointerFromQWidget(signal.(func(string, *widgets.QWidget, string) *widgets.QWidget)(C.GoString(className), widgets.NewQWidgetFromPointer(parent), C.GoString(name)))
	}

	return widgets.PointerFromQWidget(NewQUiLoaderFromPointer(ptr).CreateWidgetDefault(C.GoString(className), widgets.NewQWidgetFromPointer(parent), C.GoString(name)))
}

func (ptr *QUiLoader) ConnectCreateWidget(f func(className string, parent *widgets.QWidget, name string) *widgets.QWidget) {
	defer qt.Recovering("connect QUiLoader::createWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createWidget", f)
	}
}

func (ptr *QUiLoader) DisconnectCreateWidget() {
	defer qt.Recovering("disconnect QUiLoader::createWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createWidget")
	}
}

func (ptr *QUiLoader) CreateWidget(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	defer qt.Recovering("QUiLoader::createWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidget(ptr.Pointer(), C.CString(className), widgets.PointerFromQWidget(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateWidgetDefault(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	defer qt.Recovering("QUiLoader::createWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidgetDefault(ptr.Pointer(), C.CString(className), widgets.PointerFromQWidget(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) ErrorString() string {
	defer qt.Recovering("QUiLoader::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUiLoader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUiLoader) IsLanguageChangeEnabled() bool {
	defer qt.Recovering("QUiLoader::isLanguageChangeEnabled")

	if ptr.Pointer() != nil {
		return C.QUiLoader_IsLanguageChangeEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUiLoader) Load(device core.QIODevice_ITF, parentWidget widgets.QWidget_ITF) *widgets.QWidget {
	defer qt.Recovering("QUiLoader::load")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QUiLoader_Load(ptr.Pointer(), core.PointerFromQIODevice(device), widgets.PointerFromQWidget(parentWidget)))
	}
	return nil
}

func (ptr *QUiLoader) PluginPaths() []string {
	defer qt.Recovering("QUiLoader::pluginPaths")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUiLoader_PluginPaths(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) SetLanguageChangeEnabled(enabled bool) {
	defer qt.Recovering("QUiLoader::setLanguageChangeEnabled")

	if ptr.Pointer() != nil {
		C.QUiLoader_SetLanguageChangeEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QUiLoader) SetWorkingDirectory(dir core.QDir_ITF) {
	defer qt.Recovering("QUiLoader::setWorkingDirectory")

	if ptr.Pointer() != nil {
		C.QUiLoader_SetWorkingDirectory(ptr.Pointer(), core.PointerFromQDir(dir))
	}
}

func (ptr *QUiLoader) WorkingDirectory() *core.QDir {
	defer qt.Recovering("QUiLoader::workingDirectory")

	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QUiLoader_WorkingDirectory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUiLoader) DestroyQUiLoader() {
	defer qt.Recovering("QUiLoader::~QUiLoader")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QUiLoader_DestroyQUiLoader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQUiLoader_TimerEvent
func callbackQUiLoader_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QUiLoader::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QUiLoader::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QUiLoader) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUiLoader::timerEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUiLoader) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUiLoader::timerEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQUiLoader_ChildEvent
func callbackQUiLoader_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QUiLoader::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QUiLoader::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QUiLoader) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUiLoader::childEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUiLoader) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUiLoader::childEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQUiLoader_ConnectNotify
func callbackQUiLoader_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QUiLoader::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QUiLoader) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QUiLoader::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QUiLoader) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUiLoader::connectNotify")

	if ptr.Pointer() != nil {
		C.QUiLoader_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUiLoader) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUiLoader::connectNotify")

	if ptr.Pointer() != nil {
		C.QUiLoader_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_CustomEvent
func callbackQUiLoader_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUiLoader) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUiLoader::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QUiLoader) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QUiLoader::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QUiLoader) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUiLoader::customEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUiLoader) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUiLoader::customEvent")

	if ptr.Pointer() != nil {
		C.QUiLoader_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQUiLoader_DeleteLater
func callbackQUiLoader_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUiLoader::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQUiLoaderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QUiLoader) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QUiLoader::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QUiLoader) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QUiLoader::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QUiLoader) DeleteLater() {
	defer qt.Recovering("QUiLoader::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QUiLoader_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUiLoader) DeleteLaterDefault() {
	defer qt.Recovering("QUiLoader::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QUiLoader_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQUiLoader_DisconnectNotify
func callbackQUiLoader_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUiLoaderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUiLoader) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QUiLoader::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QUiLoader) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QUiLoader::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QUiLoader) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUiLoader::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QUiLoader_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUiLoader) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUiLoader::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QUiLoader_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUiLoader_Event
func callbackQUiLoader_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QUiLoader::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QUiLoader) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QUiLoader::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QUiLoader) DisconnectEvent() {
	defer qt.Recovering("disconnect QUiLoader::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QUiLoader) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QUiLoader::event")

	if ptr.Pointer() != nil {
		return C.QUiLoader_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QUiLoader) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QUiLoader::event")

	if ptr.Pointer() != nil {
		return C.QUiLoader_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQUiLoader_EventFilter
func callbackQUiLoader_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QUiLoader::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQUiLoaderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QUiLoader) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QUiLoader::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QUiLoader) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QUiLoader::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QUiLoader) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QUiLoader::eventFilter")

	if ptr.Pointer() != nil {
		return C.QUiLoader_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QUiLoader) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QUiLoader::eventFilter")

	if ptr.Pointer() != nil {
		return C.QUiLoader_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQUiLoader_MetaObject
func callbackQUiLoader_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QUiLoader::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQUiLoaderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QUiLoader) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QUiLoader::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QUiLoader) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QUiLoader::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QUiLoader) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QUiLoader::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUiLoader_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUiLoader) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QUiLoader::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUiLoader_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
