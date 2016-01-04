package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSoundEffect struct {
	core.QObject
}

type QSoundEffect_ITF interface {
	core.QObject_ITF
	QSoundEffect_PTR() *QSoundEffect
}

func PointerFromQSoundEffect(ptr QSoundEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSoundEffect_PTR().Pointer()
	}
	return nil
}

func NewQSoundEffectFromPointer(ptr unsafe.Pointer) *QSoundEffect {
	var n = new(QSoundEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSoundEffect_") {
		n.SetObjectName("QSoundEffect_" + qt.Identifier())
	}
	return n
}

func (ptr *QSoundEffect) QSoundEffect_PTR() *QSoundEffect {
	return ptr
}

//QSoundEffect::Loop
type QSoundEffect__Loop int64

const (
	QSoundEffect__Infinite = QSoundEffect__Loop(-2)
)

//QSoundEffect::Status
type QSoundEffect__Status int64

const (
	QSoundEffect__Null    = QSoundEffect__Status(0)
	QSoundEffect__Loading = QSoundEffect__Status(1)
	QSoundEffect__Ready   = QSoundEffect__Status(2)
	QSoundEffect__Error   = QSoundEffect__Status(3)
)

func (ptr *QSoundEffect) IsLoaded() bool {
	defer qt.Recovering("QSoundEffect::isLoaded")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) LoopsRemaining() int {
	defer qt.Recovering("QSoundEffect::loopsRemaining")

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) Play() {
	defer qt.Recovering("QSoundEffect::play")

	if ptr.Pointer() != nil {
		C.QSoundEffect_Play(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Stop() {
	defer qt.Recovering("QSoundEffect::stop")

	if ptr.Pointer() != nil {
		C.QSoundEffect_Stop(ptr.Pointer())
	}
}

func QSoundEffect_SupportedMimeTypes() []string {
	defer qt.Recovering("QSoundEffect::supportedMimeTypes")

	return strings.Split(C.GoString(C.QSoundEffect_QSoundEffect_SupportedMimeTypes()), ",,,")
}

func NewQSoundEffect(parent core.QObject_ITF) *QSoundEffect {
	defer qt.Recovering("QSoundEffect::QSoundEffect")

	return NewQSoundEffectFromPointer(C.QSoundEffect_NewQSoundEffect(core.PointerFromQObject(parent)))
}

func (ptr *QSoundEffect) Category() string {
	defer qt.Recovering("QSoundEffect::category")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectCategoryChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectCategoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "categoryChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectCategoryChanged() {
	defer qt.Recovering("disconnect QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectCategoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "categoryChanged")
	}
}

//export callbackQSoundEffectCategoryChanged
func callbackQSoundEffectCategoryChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::categoryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "categoryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) CategoryChanged() {
	defer qt.Recovering("QSoundEffect::categoryChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CategoryChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) IsMuted() bool {
	defer qt.Recovering("QSoundEffect::isMuted")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) IsPlaying() bool {
	defer qt.Recovering("QSoundEffect::isPlaying")

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsPlaying(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) ConnectLoadedChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoadedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoadedChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoadedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadedChanged")
	}
}

//export callbackQSoundEffectLoadedChanged
func callbackQSoundEffectLoadedChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loadedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadedChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoadedChanged() {
	defer qt.Recovering("QSoundEffect::loadedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoadedChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) LoopCount() int {
	defer qt.Recovering("QSoundEffect::loopCount")

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectLoopCountChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopCountChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopCountChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopCountChanged")
	}
}

//export callbackQSoundEffectLoopCountChanged
func callbackQSoundEffectLoopCountChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopCountChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoopCountChanged() {
	defer qt.Recovering("QSoundEffect::loopCountChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoopCountChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectLoopsRemainingChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopsRemainingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopsRemainingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopsRemainingChanged() {
	defer qt.Recovering("disconnect QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopsRemainingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopsRemainingChanged")
	}
}

//export callbackQSoundEffectLoopsRemainingChanged
func callbackQSoundEffectLoopsRemainingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopsRemainingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) LoopsRemainingChanged() {
	defer qt.Recovering("QSoundEffect::loopsRemainingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_LoopsRemainingChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectMutedChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQSoundEffectMutedChanged
func callbackQSoundEffectMutedChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) MutedChanged() {
	defer qt.Recovering("QSoundEffect::mutedChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_MutedChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) ConnectPlayingChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectPlayingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectPlayingChanged() {
	defer qt.Recovering("disconnect QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectPlayingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playingChanged")
	}
}

//export callbackQSoundEffectPlayingChanged
func callbackQSoundEffectPlayingChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::playingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playingChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) PlayingChanged() {
	defer qt.Recovering("QSoundEffect::playingChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_PlayingChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) SetCategory(category string) {
	defer qt.Recovering("QSoundEffect::setCategory")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QSoundEffect) SetLoopCount(loopCount int) {
	defer qt.Recovering("QSoundEffect::setLoopCount")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QSoundEffect) SetMuted(muted bool) {
	defer qt.Recovering("QSoundEffect::setMuted")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QSoundEffect) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QSoundEffect::setSource")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QSoundEffect) SetVolume(volume float64) {
	defer qt.Recovering("QSoundEffect::setVolume")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QSoundEffect) Source() *core.QUrl {
	defer qt.Recovering("QSoundEffect::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QSoundEffect_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSoundEffect) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQSoundEffectSourceChanged
func callbackQSoundEffectSourceChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) SourceChanged() {
	defer qt.Recovering("QSoundEffect::sourceChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_SourceChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Status() QSoundEffect__Status {
	defer qt.Recovering("QSoundEffect::status")

	if ptr.Pointer() != nil {
		return QSoundEffect__Status(C.QSoundEffect_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectStatusChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQSoundEffectStatusChanged
func callbackQSoundEffectStatusChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) StatusChanged() {
	defer qt.Recovering("QSoundEffect::statusChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_StatusChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Volume() float64 {
	defer qt.Recovering("QSoundEffect::volume")

	if ptr.Pointer() != nil {
		return float64(C.QSoundEffect_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectVolumeChanged(f func()) {
	defer qt.Recovering("connect QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQSoundEffectVolumeChanged
func callbackQSoundEffectVolumeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) VolumeChanged() {
	defer qt.Recovering("QSoundEffect::volumeChanged")

	if ptr.Pointer() != nil {
		C.QSoundEffect_VolumeChanged(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	defer qt.Recovering("QSoundEffect::~QSoundEffect")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSoundEffect) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSoundEffectTimerEvent
func callbackQSoundEffectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSoundEffect) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSoundEffect::timerEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSoundEffect) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSoundEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSoundEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSoundEffectChildEvent
func callbackQSoundEffectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSoundEffect::childEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSoundEffect) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSoundEffect::childEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSoundEffect) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSoundEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSoundEffect) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSoundEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSoundEffectCustomEvent
func callbackQSoundEffectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSoundEffect::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSoundEffectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSoundEffect) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSoundEffect::customEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSoundEffect) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSoundEffect::customEvent")

	if ptr.Pointer() != nil {
		C.QSoundEffect_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
