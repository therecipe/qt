package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMetaDataReaderControl struct {
	QMediaControl
}

type QMetaDataReaderControl_ITF interface {
	QMediaControl_ITF
	QMetaDataReaderControl_PTR() *QMetaDataReaderControl
}

func PointerFromQMetaDataReaderControl(ptr QMetaDataReaderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaDataReaderControl_PTR().Pointer()
	}
	return nil
}

func NewQMetaDataReaderControlFromPointer(ptr unsafe.Pointer) *QMetaDataReaderControl {
	var n = new(QMetaDataReaderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMetaDataReaderControl_") {
		n.SetObjectName("QMetaDataReaderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMetaDataReaderControl) QMetaDataReaderControl_PTR() *QMetaDataReaderControl {
	return ptr
}

func (ptr *QMetaDataReaderControl) AvailableMetaData() []string {
	defer qt.Recovering("QMetaDataReaderControl::availableMetaData")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMetaDataReaderControl_AvailableMetaData(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMetaDataReaderControl) IsMetaDataAvailable() bool {
	defer qt.Recovering("QMetaDataReaderControl::isMetaDataAvailable")

	if ptr.Pointer() != nil {
		return C.QMetaDataReaderControl_IsMetaDataAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaDataReaderControl) MetaData(key string) *core.QVariant {
	defer qt.Recovering("QMetaDataReaderControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMetaDataReaderControl_MetaData(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataAvailableChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataAvailableChanged() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataAvailableChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataAvailableChanged
func callbackQMetaDataReaderControlMetaDataAvailableChanged(ptr unsafe.Pointer, ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QMetaDataReaderControl) MetaDataAvailableChanged(available bool) {
	defer qt.Recovering("QMetaDataReaderControl::metaDataAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataAvailableChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(available)))
	}
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQMetaDataReaderControlMetaDataChanged
func callbackQMetaDataReaderControlMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMetaDataReaderControl) MetaDataChanged() {
	defer qt.Recovering("QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QMetaDataReaderControl) ConnectMetaDataChanged2(f func(key string, value *core.QVariant)) {
	defer qt.Recovering("connect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ConnectMetaDataChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged2", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectMetaDataChanged2() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DisconnectMetaDataChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged2")
	}
}

//export callbackQMetaDataReaderControlMetaDataChanged2
func callbackQMetaDataReaderControlMetaDataChanged2(ptr unsafe.Pointer, ptrName *C.char, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged2"); signal != nil {
		signal.(func(string, *core.QVariant))(C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QMetaDataReaderControl) MetaDataChanged2(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_MetaDataChanged2(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QMetaDataReaderControl) DestroyQMetaDataReaderControl() {
	defer qt.Recovering("QMetaDataReaderControl::~QMetaDataReaderControl")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_DestroyQMetaDataReaderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMetaDataReaderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMetaDataReaderControlTimerEvent
func callbackQMetaDataReaderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMetaDataReaderControlChildEvent
func callbackQMetaDataReaderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMetaDataReaderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMetaDataReaderControlCustomEvent
func callbackQMetaDataReaderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMetaDataReaderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMetaDataReaderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMetaDataReaderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMetaDataReaderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMetaDataReaderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMetaDataReaderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
