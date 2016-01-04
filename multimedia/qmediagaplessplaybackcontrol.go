package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaGaplessPlaybackControl struct {
	QMediaControl
}

type QMediaGaplessPlaybackControl_ITF interface {
	QMediaControl_ITF
	QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl
}

func PointerFromQMediaGaplessPlaybackControl(ptr QMediaGaplessPlaybackControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaGaplessPlaybackControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaGaplessPlaybackControlFromPointer(ptr unsafe.Pointer) *QMediaGaplessPlaybackControl {
	var n = new(QMediaGaplessPlaybackControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaGaplessPlaybackControl_") {
		n.SetObjectName("QMediaGaplessPlaybackControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaGaplessPlaybackControl) QMediaGaplessPlaybackControl_PTR() *QMediaGaplessPlaybackControl {
	return ptr
}

func (ptr *QMediaGaplessPlaybackControl) ConnectAdvancedToNextMedia(f func()) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "advancedToNextMedia", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectAdvancedToNextMedia() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "advancedToNextMedia")
	}
}

//export callbackQMediaGaplessPlaybackControlAdvancedToNextMedia
func callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::advancedToNextMedia")

	if signal := qt.GetSignal(C.GoString(ptrName), "advancedToNextMedia"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaGaplessPlaybackControl) AdvancedToNextMedia() {
	defer qt.Recovering("QMediaGaplessPlaybackControl::advancedToNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_AdvancedToNextMedia(ptr.Pointer())
	}
}

func (ptr *QMediaGaplessPlaybackControl) CrossfadeTime() float64 {
	defer qt.Recovering("QMediaGaplessPlaybackControl::crossfadeTime")

	if ptr.Pointer() != nil {
		return float64(C.QMediaGaplessPlaybackControl_CrossfadeTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaGaplessPlaybackControl) ConnectCrossfadeTimeChanged(f func(crossfadeTime float64)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectCrossfadeTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "crossfadeTimeChanged", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectCrossfadeTimeChanged() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectCrossfadeTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "crossfadeTimeChanged")
	}
}

//export callbackQMediaGaplessPlaybackControlCrossfadeTimeChanged
func callbackQMediaGaplessPlaybackControlCrossfadeTimeChanged(ptr unsafe.Pointer, ptrName *C.char, crossfadeTime C.double) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "crossfadeTimeChanged"); signal != nil {
		signal.(func(float64))(float64(crossfadeTime))
	}

}

func (ptr *QMediaGaplessPlaybackControl) CrossfadeTimeChanged(crossfadeTime float64) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::crossfadeTimeChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CrossfadeTimeChanged(ptr.Pointer(), C.double(crossfadeTime))
	}
}

func (ptr *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	defer qt.Recovering("QMediaGaplessPlaybackControl::isCrossfadeSupported")

	if ptr.Pointer() != nil {
		return C.QMediaGaplessPlaybackControl_IsCrossfadeSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaGaplessPlaybackControl) NextMedia() *QMediaContent {
	defer qt.Recovering("QMediaGaplessPlaybackControl::nextMedia")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaGaplessPlaybackControl_NextMedia(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaGaplessPlaybackControl) ConnectNextMediaChanged(f func(media *QMediaContent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ConnectNextMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nextMediaChanged", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectNextMediaChanged() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DisconnectNextMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nextMediaChanged")
	}
}

//export callbackQMediaGaplessPlaybackControlNextMediaChanged
func callbackQMediaGaplessPlaybackControlNextMediaChanged(ptr unsafe.Pointer, ptrName *C.char, media unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::nextMediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextMediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(media))
	}

}

func (ptr *QMediaGaplessPlaybackControl) NextMediaChanged(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::nextMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_NextMediaChanged(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaGaplessPlaybackControl) SetCrossfadeTime(crossfadeTime float64) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::setCrossfadeTime")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetCrossfadeTime(ptr.Pointer(), C.double(crossfadeTime))
	}
}

func (ptr *QMediaGaplessPlaybackControl) SetNextMedia(media QMediaContent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::setNextMedia")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_SetNextMedia(ptr.Pointer(), PointerFromQMediaContent(media))
	}
}

func (ptr *QMediaGaplessPlaybackControl) DestroyQMediaGaplessPlaybackControl() {
	defer qt.Recovering("QMediaGaplessPlaybackControl::~QMediaGaplessPlaybackControl")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlTimerEvent
func callbackQMediaGaplessPlaybackControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlChildEvent
func callbackQMediaGaplessPlaybackControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaGaplessPlaybackControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaGaplessPlaybackControlCustomEvent
func callbackQMediaGaplessPlaybackControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaGaplessPlaybackControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaGaplessPlaybackControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaGaplessPlaybackControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaGaplessPlaybackControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaGaplessPlaybackControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
