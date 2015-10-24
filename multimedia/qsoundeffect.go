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

type QSoundEffectITF interface {
	core.QObjectITF
	QSoundEffectPTR() *QSoundEffect
}

func PointerFromQSoundEffect(ptr QSoundEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSoundEffectPTR().Pointer()
	}
	return nil
}

func QSoundEffectFromPointer(ptr unsafe.Pointer) *QSoundEffect {
	var n = new(QSoundEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSoundEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSoundEffect) QSoundEffectPTR() *QSoundEffect {
	return ptr
}

//QSoundEffect::Loop
type QSoundEffect__Loop int

var (
	QSoundEffect__Infinite = QSoundEffect__Loop(-2)
)

//QSoundEffect::Status
type QSoundEffect__Status int

var (
	QSoundEffect__Null    = QSoundEffect__Status(0)
	QSoundEffect__Loading = QSoundEffect__Status(1)
	QSoundEffect__Ready   = QSoundEffect__Status(2)
	QSoundEffect__Error   = QSoundEffect__Status(3)
)

func (ptr *QSoundEffect) IsLoaded() bool {
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsLoaded(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSoundEffect) LoopsRemaining() int {
	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopsRemaining(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSoundEffect) Play() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_Play(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSoundEffect) Stop() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QSoundEffect_SupportedMimeTypes() []string {
	return strings.Split(C.GoString(C.QSoundEffect_QSoundEffect_SupportedMimeTypes()), "|")
}

func NewQSoundEffect(parent core.QObjectITF) *QSoundEffect {
	return QSoundEffectFromPointer(unsafe.Pointer(C.QSoundEffect_NewQSoundEffect(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSoundEffect) Category() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Category(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectCategoryChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectCategoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "categoryChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectCategoryChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectCategoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "categoryChanged")
	}
}

//export callbackQSoundEffectCategoryChanged
func callbackQSoundEffectCategoryChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "categoryChanged").(func())()
}

func (ptr *QSoundEffect) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSoundEffect) IsPlaying() bool {
	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsPlaying(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSoundEffect) ConnectLoadedChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoadedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "loadedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoadedChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoadedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "loadedChanged")
	}
}

//export callbackQSoundEffectLoadedChanged
func callbackQSoundEffectLoadedChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loadedChanged").(func())()
}

func (ptr *QSoundEffect) LoopCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectLoopCountChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "loopCountChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopCountChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "loopCountChanged")
	}
}

//export callbackQSoundEffectLoopCountChanged
func callbackQSoundEffectLoopCountChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loopCountChanged").(func())()
}

func (ptr *QSoundEffect) ConnectLoopsRemainingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopsRemainingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "loopsRemainingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopsRemainingChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopsRemainingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "loopsRemainingChanged")
	}
}

//export callbackQSoundEffectLoopsRemainingChanged
func callbackQSoundEffectLoopsRemainingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged").(func())()
}

func (ptr *QSoundEffect) ConnectMutedChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQSoundEffectMutedChanged
func callbackQSoundEffectMutedChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func())()
}

func (ptr *QSoundEffect) ConnectPlayingChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectPlayingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "playingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectPlayingChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectPlayingChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "playingChanged")
	}
}

//export callbackQSoundEffectPlayingChanged
func callbackQSoundEffectPlayingChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "playingChanged").(func())()
}

func (ptr *QSoundEffect) SetCategory(category string) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetCategory(C.QtObjectPtr(ptr.Pointer()), C.CString(category))
	}
}

func (ptr *QSoundEffect) SetLoopCount(loopCount int) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetLoopCount(C.QtObjectPtr(ptr.Pointer()), C.int(loopCount))
	}
}

func (ptr *QSoundEffect) SetMuted(muted bool) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QSoundEffect) SetSource(url string) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_SetSource(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QSoundEffect) Source() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectSourceChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectSourceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQSoundEffectSourceChanged
func callbackQSoundEffectSourceChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QSoundEffect) Status() QSoundEffect__Status {
	if ptr.Pointer() != nil {
		return QSoundEffect__Status(C.QSoundEffect_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectStatusChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQSoundEffectStatusChanged
func callbackQSoundEffectStatusChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func())()
}

func (ptr *QSoundEffect) ConnectVolumeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQSoundEffectVolumeChanged
func callbackQSoundEffectVolumeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func())()
}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
