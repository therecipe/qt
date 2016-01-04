package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpEngine_") {
		n.SetObjectName("QHelpEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return ptr
}

func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	defer qt.Recovering("QHelpEngine::QHelpEngine")

	return NewQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	defer qt.Recovering("QHelpEngine::contentModel")

	if ptr.Pointer() != nil {
		return NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	defer qt.Recovering("QHelpEngine::contentWidget")

	if ptr.Pointer() != nil {
		return NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	defer qt.Recovering("QHelpEngine::indexModel")

	if ptr.Pointer() != nil {
		return NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	defer qt.Recovering("QHelpEngine::indexWidget")

	if ptr.Pointer() != nil {
		return NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	defer qt.Recovering("QHelpEngine::searchEngine")

	if ptr.Pointer() != nil {
		return NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	defer qt.Recovering("QHelpEngine::~QHelpEngine")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpEngineTimerEvent
func callbackQHelpEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpEngineChildEvent
func callbackQHelpEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpEngineCustomEvent
func callbackQHelpEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
