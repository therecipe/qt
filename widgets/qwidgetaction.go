package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWidgetAction struct {
	QAction
}

type QWidgetAction_ITF interface {
	QAction_ITF
	QWidgetAction_PTR() *QWidgetAction
}

func PointerFromQWidgetAction(ptr QWidgetAction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetAction_PTR().Pointer()
	}
	return nil
}

func NewQWidgetActionFromPointer(ptr unsafe.Pointer) *QWidgetAction {
	var n = new(QWidgetAction)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWidgetAction_") {
		n.SetObjectName("QWidgetAction_" + qt.Identifier())
	}
	return n
}

func (ptr *QWidgetAction) QWidgetAction_PTR() *QWidgetAction {
	return ptr
}

func NewQWidgetAction(parent core.QObject_ITF) *QWidgetAction {
	defer qt.Recovering("QWidgetAction::QWidgetAction")

	return NewQWidgetActionFromPointer(C.QWidgetAction_NewQWidgetAction(core.PointerFromQObject(parent)))
}

func (ptr *QWidgetAction) CreateWidget(parent QWidget_ITF) *QWidget {
	defer qt.Recovering("QWidgetAction::createWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetAction_CreateWidget(ptr.Pointer(), PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QWidgetAction) DefaultWidget() *QWidget {
	defer qt.Recovering("QWidgetAction::defaultWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetAction_DefaultWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidgetAction) ConnectDeleteWidget(f func(widget *QWidget)) {
	defer qt.Recovering("connect QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteWidget", f)
	}
}

func (ptr *QWidgetAction) DisconnectDeleteWidget() {
	defer qt.Recovering("disconnect QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteWidget")
	}
}

//export callbackQWidgetActionDeleteWidget
func callbackQWidgetActionDeleteWidget(ptr unsafe.Pointer, ptrName *C.char, widget unsafe.Pointer) {
	defer qt.Recovering("callback QWidgetAction::deleteWidget")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteWidget"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
	} else {
		NewQWidgetActionFromPointer(ptr).DeleteWidgetDefault(NewQWidgetFromPointer(widget))
	}
}

func (ptr *QWidgetAction) DeleteWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_DeleteWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) DeleteWidgetDefault(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_DeleteWidgetDefault(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QWidgetAction::event")

	if ptr.Pointer() != nil {
		return C.QWidgetAction_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWidgetAction) EventFilter(obj core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWidgetAction::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWidgetAction_EventFilter(ptr.Pointer(), core.PointerFromQObject(obj), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWidgetAction) ReleaseWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::releaseWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_ReleaseWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) RequestWidget(parent QWidget_ITF) *QWidget {
	defer qt.Recovering("QWidgetAction::requestWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetAction_RequestWidget(ptr.Pointer(), PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QWidgetAction) SetDefaultWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::setDefaultWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_SetDefaultWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) DestroyQWidgetAction() {
	defer qt.Recovering("QWidgetAction::~QWidgetAction")

	if ptr.Pointer() != nil {
		C.QWidgetAction_DestroyQWidgetAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWidgetAction) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWidgetAction::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWidgetAction) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWidgetAction::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWidgetActionTimerEvent
func callbackQWidgetActionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWidgetAction::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWidgetActionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWidgetAction) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWidgetAction::timerEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWidgetAction) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWidgetAction::timerEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWidgetAction) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWidgetAction::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWidgetAction) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWidgetAction::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWidgetActionChildEvent
func callbackQWidgetActionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWidgetAction::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWidgetActionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWidgetAction) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWidgetAction::childEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWidgetAction) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWidgetAction::childEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWidgetAction) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWidgetAction::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWidgetAction) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWidgetAction::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWidgetActionCustomEvent
func callbackQWidgetActionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWidgetAction::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWidgetActionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWidgetAction) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWidgetAction::customEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWidgetAction) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWidgetAction::customEvent")

	if ptr.Pointer() != nil {
		C.QWidgetAction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
