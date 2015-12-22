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
func callbackQSoundEffectCategoryChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::categoryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "categoryChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectLoadedChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loadedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadedChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectLoopCountChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopCountChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectLoopsRemainingChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::loopsRemainingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectMutedChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::mutedChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mutedChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectPlayingChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::playingChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playingChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectSourceChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectStatusChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func())()
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
func callbackQSoundEffectVolumeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QSoundEffect::volumeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "volumeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	defer qt.Recovering("QSoundEffect::~QSoundEffect")

	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
