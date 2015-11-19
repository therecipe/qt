package multimedia

//#include "qsoundeffect.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QSoundEffect_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) LoopsRemaining() int {
	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) Play() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_Play(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Stop() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_Stop(ptr.Pointer())
	}
}

func QSoundEffect_SupportedMimeTypes() []string {
	return strings.Split(C.GoString(C.QSoundEffect_QSoundEffect_SupportedMimeTypes()), "|")
}

func NewQSoundEffect(parent core.QObject_ITF) *QSoundEffect {
	return NewQSoundEffectFromPointer(C.QSoundEffect_NewQSoundEffect(core.PointerFromQObject(parent)))
}

func (ptr *QSoundEffect) Category() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectCategoryChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectCategoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "categoryChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectCategoryChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectCategoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "categoryChanged")
	}
}

//export callbackQSoundEffectCategoryChanged
func callbackQSoundEffectCategoryChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "categoryChanged").(func())()
}

func (ptr *QSoundEffect) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) IsPlaying() bool {
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsPlaying(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) ConnectLoadedChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoadedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoadedChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoadedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadedChanged")
	}
}

//export callbackQSoundEffectLoadedChanged
func callbackQSoundEffectLoadedChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loadedChanged").(func())()
}

func (ptr *QSoundEffect) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectLoopCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopCountChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopCountChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopCountChanged")
	}
}

//export callbackQSoundEffectLoopCountChanged
func callbackQSoundEffectLoopCountChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loopCountChanged").(func())()
}

func (ptr *QSoundEffect) ConnectLoopsRemainingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopsRemainingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopsRemainingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopsRemainingChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopsRemainingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopsRemainingChanged")
	}
}

//export callbackQSoundEffectLoopsRemainingChanged
func callbackQSoundEffectLoopsRemainingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged").(func())()
}

func (ptr *QSoundEffect) ConnectMutedChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQSoundEffectMutedChanged
func callbackQSoundEffectMutedChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func())()
}

func (ptr *QSoundEffect) ConnectPlayingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectPlayingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectPlayingChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectPlayingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playingChanged")
	}
}

//export callbackQSoundEffectPlayingChanged
func callbackQSoundEffectPlayingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "playingChanged").(func())()
}

func (ptr *QSoundEffect) SetCategory(category string) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QSoundEffect) SetLoopCount(loopCount int) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QSoundEffect) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QSoundEffect) SetSource(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QSoundEffect) SetVolume(volume float64) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QSoundEffect) ConnectSourceChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQSoundEffectSourceChanged
func callbackQSoundEffectSourceChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QSoundEffect) Status() QSoundEffect__Status {
	if ptr.Pointer() != nil {
		return QSoundEffect__Status(C.QSoundEffect_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectStatusChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQSoundEffectStatusChanged
func callbackQSoundEffectStatusChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func())()
}

func (ptr *QSoundEffect) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QSoundEffect_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectVolumeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQSoundEffectVolumeChanged
func callbackQSoundEffectVolumeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func())()
}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
