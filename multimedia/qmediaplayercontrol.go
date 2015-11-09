package multimedia

//#include "qmediaplayercontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaPlayerControl struct {
	QMediaControl
}

type QMediaPlayerControl_ITF interface {
	QMediaControl_ITF
	QMediaPlayerControl_PTR() *QMediaPlayerControl
}

func PointerFromQMediaPlayerControl(ptr QMediaPlayerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlayerControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlayerControlFromPointer(ptr unsafe.Pointer) *QMediaPlayerControl {
	var n = new(QMediaPlayerControl)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QMediaPlayerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlayerControl) QMediaPlayerControl_PTR() *QMediaPlayerControl {
	return ptr
}

func (ptr *QMediaPlayerControl) ConnectAudioAvailableChanged(f func(audio bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAudioAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerControlAudioAvailableChanged
func callbackQMediaPlayerControlAudioAvailableChanged(ptrName *C.char, audio C.int) {
	qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged").(func(bool))(int(audio) != 0)
}

func (ptr *QMediaPlayerControl) BufferStatus() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectBufferStatusChanged(f func(progress int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectBufferStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerControlBufferStatusChanged
func callbackQMediaPlayerControlBufferStatusChanged(ptrName *C.char, progress C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged").(func(int))(int(progress))
}

func (ptr *QMediaPlayerControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaPlayerControlError
func callbackQMediaPlayerControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaPlayerControl) IsAudioAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsSeekable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsVideoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) MediaStatus() QMediaPlayer__MediaStatus {
	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayerControl_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerControlMediaStatusChanged
func callbackQMediaPlayerControlMediaStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged").(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
}

func (ptr *QMediaPlayerControl) MediaStream() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayerControl_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMutedChanged(f func(mute bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerControlMutedChanged
func callbackQMediaPlayerControlMutedChanged(ptrName *C.char, mute C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(mute) != 0)
}

func (ptr *QMediaPlayerControl) Pause() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) Play() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) PlaybackRate() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayerControl_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectSeekableChanged(f func(seekable bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectSeekableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerControlSeekableChanged
func callbackQMediaPlayerControlSeekableChanged(ptrName *C.char, seekable C.int) {
	qt.GetSignal(C.GoString(ptrName), "seekableChanged").(func(bool))(int(seekable) != 0)
}

func (ptr *QMediaPlayerControl) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayerControl) SetMuted(mute bool) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) SetPlaybackRate(rate float64) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayerControl) SetVolume(volume int) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerControlStateChanged
func callbackQMediaPlayerControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaPlayer__State))(QMediaPlayer__State(state))
}

func (ptr *QMediaPlayerControl) Stop() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) ConnectVideoAvailableChanged(f func(video bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVideoAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerControlVideoAvailableChanged
func callbackQMediaPlayerControlVideoAvailableChanged(ptrName *C.char, video C.int) {
	qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged").(func(bool))(int(video) != 0)
}

func (ptr *QMediaPlayerControl) Volume() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerControlVolumeChanged
func callbackQMediaPlayerControlVolumeChanged(ptrName *C.char, volume C.int) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QMediaPlayerControl) DestroyQMediaPlayerControl() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DestroyQMediaPlayerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
