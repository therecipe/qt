package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPluginLoader struct {
	QObject
}

type QPluginLoader_ITF interface {
	QObject_ITF
	QPluginLoader_PTR() *QPluginLoader
}

func PointerFromQPluginLoader(ptr QPluginLoader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPluginLoader_PTR().Pointer()
	}
	return nil
}

func NewQPluginLoaderFromPointer(ptr unsafe.Pointer) *QPluginLoader {
	var n = new(QPluginLoader)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPluginLoader_") {
		n.SetObjectName("QPluginLoader_" + qt.Identifier())
	}
	return n
}

func (ptr *QPluginLoader) QPluginLoader_PTR() *QPluginLoader {
	return ptr
}

func (ptr *QPluginLoader) FileName() string {
	defer qt.Recovering("QPluginLoader::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPluginLoader) LoadHints() QLibrary__LoadHint {
	defer qt.Recovering("QPluginLoader::loadHints")

	if ptr.Pointer() != nil {
		return QLibrary__LoadHint(C.QPluginLoader_LoadHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPluginLoader) SetFileName(fileName string) {
	defer qt.Recovering("QPluginLoader::setFileName")

	if ptr.Pointer() != nil {
		C.QPluginLoader_SetFileName(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QPluginLoader) SetLoadHints(loadHints QLibrary__LoadHint) {
	defer qt.Recovering("QPluginLoader::setLoadHints")

	if ptr.Pointer() != nil {
		C.QPluginLoader_SetLoadHints(ptr.Pointer(), C.int(loadHints))
	}
}

func NewQPluginLoader(parent QObject_ITF) *QPluginLoader {
	defer qt.Recovering("QPluginLoader::QPluginLoader")

	return NewQPluginLoaderFromPointer(C.QPluginLoader_NewQPluginLoader(PointerFromQObject(parent)))
}

func NewQPluginLoader2(fileName string, parent QObject_ITF) *QPluginLoader {
	defer qt.Recovering("QPluginLoader::QPluginLoader")

	return NewQPluginLoaderFromPointer(C.QPluginLoader_NewQPluginLoader2(C.CString(fileName), PointerFromQObject(parent)))
}

func (ptr *QPluginLoader) ErrorString() string {
	defer qt.Recovering("QPluginLoader::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPluginLoader_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPluginLoader) Instance() *QObject {
	defer qt.Recovering("QPluginLoader::instance")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QPluginLoader_Instance(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPluginLoader) IsLoaded() bool {
	defer qt.Recovering("QPluginLoader::isLoaded")

	if ptr.Pointer() != nil {
		return C.QPluginLoader_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) Load() bool {
	defer qt.Recovering("QPluginLoader::load")

	if ptr.Pointer() != nil {
		return C.QPluginLoader_Load(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) MetaData() *QJsonObject {
	defer qt.Recovering("QPluginLoader::metaData")

	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QPluginLoader_MetaData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPluginLoader) Unload() bool {
	defer qt.Recovering("QPluginLoader::unload")

	if ptr.Pointer() != nil {
		return C.QPluginLoader_Unload(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPluginLoader) DestroyQPluginLoader() {
	defer qt.Recovering("QPluginLoader::~QPluginLoader")

	if ptr.Pointer() != nil {
		C.QPluginLoader_DestroyQPluginLoader(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPluginLoader) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QPluginLoader::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPluginLoader) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPluginLoader::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPluginLoaderTimerEvent
func callbackQPluginLoaderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPluginLoader::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQPluginLoaderFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPluginLoader) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QPluginLoader::timerEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPluginLoader) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QPluginLoader::timerEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPluginLoader) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QPluginLoader::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPluginLoader) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPluginLoader::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPluginLoaderChildEvent
func callbackQPluginLoaderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPluginLoader::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQPluginLoaderFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QPluginLoader) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QPluginLoader::childEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPluginLoader) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QPluginLoader::childEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPluginLoader) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QPluginLoader::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPluginLoader) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPluginLoader::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPluginLoaderCustomEvent
func callbackQPluginLoaderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPluginLoader::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQPluginLoaderFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QPluginLoader) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QPluginLoader::customEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QPluginLoader) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QPluginLoader::customEvent")

	if ptr.Pointer() != nil {
		C.QPluginLoader_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
