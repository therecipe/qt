package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFileSelector struct {
	QObject
}

type QFileSelector_ITF interface {
	QObject_ITF
	QFileSelector_PTR() *QFileSelector
}

func PointerFromQFileSelector(ptr QFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQFileSelectorFromPointer(ptr unsafe.Pointer) *QFileSelector {
	var n = new(QFileSelector)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileSelector_") {
		n.SetObjectName("QFileSelector_" + qt.Identifier())
	}
	return n
}

func (ptr *QFileSelector) QFileSelector_PTR() *QFileSelector {
	return ptr
}

func NewQFileSelector(parent QObject_ITF) *QFileSelector {
	defer qt.Recovering("QFileSelector::QFileSelector")

	return NewQFileSelectorFromPointer(C.QFileSelector_NewQFileSelector(PointerFromQObject(parent)))
}

func (ptr *QFileSelector) AllSelectors() []string {
	defer qt.Recovering("QFileSelector::allSelectors")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_AllSelectors(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) ExtraSelectors() []string {
	defer qt.Recovering("QFileSelector::extraSelectors")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSelector_ExtraSelectors(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSelector) Select(filePath string) string {
	defer qt.Recovering("QFileSelector::select")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSelector_Select(ptr.Pointer(), C.CString(filePath)))
	}
	return ""
}

func (ptr *QFileSelector) Select2(filePath QUrl_ITF) *QUrl {
	defer qt.Recovering("QFileSelector::select")

	if ptr.Pointer() != nil {
		return NewQUrlFromPointer(C.QFileSelector_Select2(ptr.Pointer(), PointerFromQUrl(filePath)))
	}
	return nil
}

func (ptr *QFileSelector) SetExtraSelectors(list []string) {
	defer qt.Recovering("QFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QFileSelector_SetExtraSelectors(ptr.Pointer(), C.CString(strings.Join(list, ",,,")))
	}
}

func (ptr *QFileSelector) DestroyQFileSelector() {
	defer qt.Recovering("QFileSelector::~QFileSelector")

	if ptr.Pointer() != nil {
		C.QFileSelector_DestroyQFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileSelector) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFileSelector) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFileSelectorTimerEvent
func callbackQFileSelectorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSelector::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQFileSelectorFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFileSelector) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSelector) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileSelector::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileSelector) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFileSelector) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFileSelectorChildEvent
func callbackQFileSelectorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSelector::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQFileSelectorFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QFileSelector) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSelector) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QFileSelector::childEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileSelector) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFileSelector) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFileSelectorCustomEvent
func callbackQFileSelectorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileSelector::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFileSelectorFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFileSelector) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFileSelector) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QFileSelector::customEvent")

	if ptr.Pointer() != nil {
		C.QFileSelector_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
