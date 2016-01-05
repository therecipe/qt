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

func PointerFromQUiLoader(ptr QUiLoader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUiLoader_PTR().Pointer()
	}
	return nil
}

func NewQUiLoaderFromPointer(ptr unsafe.Pointer) *QUiLoader {
	var n = new(QUiLoader)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUiLoader_") {
		n.SetObjectName("QUiLoader_" + qt.Identifier())
	}
	return n
}

func (ptr *QUiLoader) QUiLoader_PTR() *QUiLoader {
	return ptr
}

func NewQUiLoader(parent core.QObject_ITF) *QUiLoader {
	defer qt.Recovering("QUiLoader::QUiLoader")

	return NewQUiLoaderFromPointer(C.QUiLoader_NewQUiLoader(core.PointerFromQObject(parent)))
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
		return strings.Split(C.GoString(C.QUiLoader_AvailableLayouts(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) AvailableWidgets() []string {
	defer qt.Recovering("QUiLoader::availableWidgets")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUiLoader_AvailableWidgets(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QUiLoader) ClearPluginPaths() {
	defer qt.Recovering("QUiLoader::clearPluginPaths")

	if ptr.Pointer() != nil {
		C.QUiLoader_ClearPluginPaths(ptr.Pointer())
	}
}

func (ptr *QUiLoader) CreateAction(parent core.QObject_ITF, name string) *widgets.QAction {
	defer qt.Recovering("QUiLoader::createAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QUiLoader_CreateAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateActionGroup(parent core.QObject_ITF, name string) *widgets.QActionGroup {
	defer qt.Recovering("QUiLoader::createActionGroup")

	if ptr.Pointer() != nil {
		return widgets.NewQActionGroupFromPointer(C.QUiLoader_CreateActionGroup(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateLayout(className string, parent core.QObject_ITF, name string) *widgets.QLayout {
	defer qt.Recovering("QUiLoader::createLayout")

	if ptr.Pointer() != nil {
		return widgets.NewQLayoutFromPointer(C.QUiLoader_CreateLayout(ptr.Pointer(), C.CString(className), core.PointerFromQObject(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QUiLoader) CreateWidget(className string, parent widgets.QWidget_ITF, name string) *widgets.QWidget {
	defer qt.Recovering("QUiLoader::createWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QUiLoader_CreateWidget(ptr.Pointer(), C.CString(className), widgets.PointerFromQWidget(parent), C.CString(name)))
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
		return strings.Split(C.GoString(C.QUiLoader_PluginPaths(ptr.Pointer())), ",,,")
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
		C.QUiLoader_DestroyQUiLoader(ptr.Pointer())
		ptr.SetPointer(nil)
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

//export callbackQUiLoaderTimerEvent
func callbackQUiLoaderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQUiLoaderChildEvent
func callbackQUiLoaderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQUiLoaderCustomEvent
func callbackQUiLoaderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUiLoader::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUiLoaderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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
