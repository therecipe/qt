package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaStreamsControl struct {
	QMediaControl
}

type QMediaStreamsControl_ITF interface {
	QMediaControl_ITF
	QMediaStreamsControl_PTR() *QMediaStreamsControl
}

func PointerFromQMediaStreamsControl(ptr QMediaStreamsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaStreamsControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaStreamsControlFromPointer(ptr unsafe.Pointer) *QMediaStreamsControl {
	var n = new(QMediaStreamsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaStreamsControl_") {
		n.SetObjectName("QMediaStreamsControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaStreamsControl) QMediaStreamsControl_PTR() *QMediaStreamsControl {
	return ptr
}

//QMediaStreamsControl::StreamType
type QMediaStreamsControl__StreamType int64

const (
	QMediaStreamsControl__UnknownStream    = QMediaStreamsControl__StreamType(0)
	QMediaStreamsControl__VideoStream      = QMediaStreamsControl__StreamType(1)
	QMediaStreamsControl__AudioStream      = QMediaStreamsControl__StreamType(2)
	QMediaStreamsControl__SubPictureStream = QMediaStreamsControl__StreamType(3)
	QMediaStreamsControl__DataStream       = QMediaStreamsControl__StreamType(4)
)

func (ptr *QMediaStreamsControl) ConnectActiveStreamsChanged(f func()) {
	defer qt.Recovering("connect QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectActiveStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeStreamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectActiveStreamsChanged() {
	defer qt.Recovering("disconnect QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectActiveStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeStreamsChanged")
	}
}

//export callbackQMediaStreamsControlActiveStreamsChanged
func callbackQMediaStreamsControlActiveStreamsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaStreamsControl::activeStreamsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeStreamsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaStreamsControl) ActiveStreamsChanged() {
	defer qt.Recovering("QMediaStreamsControl::activeStreamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ActiveStreamsChanged(ptr.Pointer())
	}
}

func (ptr *QMediaStreamsControl) IsActive(stream int) bool {
	defer qt.Recovering("QMediaStreamsControl::isActive")

	if ptr.Pointer() != nil {
		return C.QMediaStreamsControl_IsActive(ptr.Pointer(), C.int(stream)) != 0
	}
	return false
}

func (ptr *QMediaStreamsControl) MetaData(stream int, key string) *core.QVariant {
	defer qt.Recovering("QMediaStreamsControl::metaData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaStreamsControl_MetaData(ptr.Pointer(), C.int(stream), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaStreamsControl) SetActive(stream int, state bool) {
	defer qt.Recovering("QMediaStreamsControl::setActive")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_SetActive(ptr.Pointer(), C.int(stream), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QMediaStreamsControl) StreamCount() int {
	defer qt.Recovering("QMediaStreamsControl::streamCount")

	if ptr.Pointer() != nil {
		return int(C.QMediaStreamsControl_StreamCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaStreamsControl) StreamType(stream int) QMediaStreamsControl__StreamType {
	defer qt.Recovering("QMediaStreamsControl::streamType")

	if ptr.Pointer() != nil {
		return QMediaStreamsControl__StreamType(C.QMediaStreamsControl_StreamType(ptr.Pointer(), C.int(stream)))
	}
	return 0
}

func (ptr *QMediaStreamsControl) ConnectStreamsChanged(f func()) {
	defer qt.Recovering("connect QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "streamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectStreamsChanged() {
	defer qt.Recovering("disconnect QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "streamsChanged")
	}
}

//export callbackQMediaStreamsControlStreamsChanged
func callbackQMediaStreamsControlStreamsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaStreamsControl::streamsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "streamsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaStreamsControl) StreamsChanged() {
	defer qt.Recovering("QMediaStreamsControl::streamsChanged")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_StreamsChanged(ptr.Pointer())
	}
}

func (ptr *QMediaStreamsControl) DestroyQMediaStreamsControl() {
	defer qt.Recovering("QMediaStreamsControl::~QMediaStreamsControl")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DestroyQMediaStreamsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaStreamsControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaStreamsControlTimerEvent
func callbackQMediaStreamsControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaStreamsControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaStreamsControlChildEvent
func callbackQMediaStreamsControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaStreamsControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaStreamsControlCustomEvent
func callbackQMediaStreamsControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaStreamsControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaStreamsControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaStreamsControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaStreamsControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaStreamsControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
