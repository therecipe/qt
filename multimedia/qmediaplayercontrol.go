package multimedia

//#include "multimedia.h"
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
	for len(n.ObjectName()) < len("QMediaPlayerControl_") {
		n.SetObjectName("QMediaPlayerControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaPlayerControl) QMediaPlayerControl_PTR() *QMediaPlayerControl {
	return ptr
}

func (ptr *QMediaPlayerControl) ConnectAudioAvailableChanged(f func(audio bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAudioAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::audioAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerControlAudioAvailableChanged
func callbackQMediaPlayerControlAudioAvailableChanged(ptrName *C.char, audio C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::audioAvailableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged")
	if signal != nil {
		signal.(func(bool))(int(audio) != 0)
	}

}

func (ptr *QMediaPlayerControl) BufferStatus() int {
	defer qt.Recovering("QMediaPlayerControl::bufferStatus")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectBufferStatusChanged(f func(progress int)) {
	defer qt.Recovering("connect QMediaPlayerControl::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectBufferStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::bufferStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerControlBufferStatusChanged
func callbackQMediaPlayerControlBufferStatusChanged(ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::bufferStatusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged")
	if signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QMediaPlayerControl) Duration() int64 {
	defer qt.Recovering("QMediaPlayerControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayerControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QMediaPlayerControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQMediaPlayerControlDurationChanged
func callbackQMediaPlayerControlDurationChanged(ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QMediaPlayerControl::durationChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "durationChanged")
	if signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QMediaPlayerControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QMediaPlayerControl::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectError() {
	defer qt.Recovering("disconnect QMediaPlayerControl::error")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaPlayerControlError
func callbackQMediaPlayerControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QMediaPlayerControl::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error")
	if signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QMediaPlayerControl) IsAudioAvailable() bool {
	defer qt.Recovering("QMediaPlayerControl::isAudioAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsMuted() bool {
	defer qt.Recovering("QMediaPlayerControl::isMuted")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsSeekable() bool {
	defer qt.Recovering("QMediaPlayerControl::isSeekable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsVideoAvailable() bool {
	defer qt.Recovering("QMediaPlayerControl::isVideoAvailable")

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) Media() *QMediaContent {
	defer qt.Recovering("QMediaPlayerControl::media")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlayerControl_Media(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMediaChanged(f func(content *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlayerControl::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlayerControlMediaChanged
func callbackQMediaPlayerControlMediaChanged(ptrName *C.char, content unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlayerControl::mediaChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "mediaChanged")
	if signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(content))
	}

}

func (ptr *QMediaPlayerControl) MediaStatus() QMediaPlayer__MediaStatus {
	defer qt.Recovering("QMediaPlayerControl::mediaStatus")

	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayerControl_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	defer qt.Recovering("connect QMediaPlayerControl::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaStatusChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mediaStatusChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerControlMediaStatusChanged
func callbackQMediaPlayerControlMediaStatusChanged(ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::mediaStatusChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged")
	if signal != nil {
		signal.(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
	}

}

func (ptr *QMediaPlayerControl) MediaStream() *core.QIODevice {
	defer qt.Recovering("QMediaPlayerControl::mediaStream")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayerControl_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMutedChanged(f func(mute bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMutedChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::mutedChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerControlMutedChanged
func callbackQMediaPlayerControlMutedChanged(ptrName *C.char, mute C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::mutedChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "mutedChanged")
	if signal != nil {
		signal.(func(bool))(int(mute) != 0)
	}

}

func (ptr *QMediaPlayerControl) Pause() {
	defer qt.Recovering("QMediaPlayerControl::pause")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) Play() {
	defer qt.Recovering("QMediaPlayerControl::play")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) PlaybackRate() float64 {
	defer qt.Recovering("QMediaPlayerControl::playbackRate")

	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayerControl_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectPlaybackRateChanged(f func(rate float64)) {
	defer qt.Recovering("connect QMediaPlayerControl::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectPlaybackRateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackRateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectPlaybackRateChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::playbackRateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectPlaybackRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackRateChanged")
	}
}

//export callbackQMediaPlayerControlPlaybackRateChanged
func callbackQMediaPlayerControlPlaybackRateChanged(ptrName *C.char, rate C.double) {
	defer qt.Recovering("callback QMediaPlayerControl::playbackRateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "playbackRateChanged")
	if signal != nil {
		signal.(func(float64))(float64(rate))
	}

}

func (ptr *QMediaPlayerControl) Position() int64 {
	defer qt.Recovering("QMediaPlayerControl::position")

	if ptr.Pointer() != nil {
		return int64(C.QMediaPlayerControl_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QMediaPlayerControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQMediaPlayerControlPositionChanged
func callbackQMediaPlayerControlPositionChanged(ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QMediaPlayerControl::positionChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "positionChanged")
	if signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QMediaPlayerControl) ConnectSeekableChanged(f func(seekable bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectSeekableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::seekableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerControlSeekableChanged
func callbackQMediaPlayerControlSeekableChanged(ptrName *C.char, seekable C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::seekableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "seekableChanged")
	if signal != nil {
		signal.(func(bool))(int(seekable) != 0)
	}

}

func (ptr *QMediaPlayerControl) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	defer qt.Recovering("QMediaPlayerControl::setMedia")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayerControl) SetMuted(mute bool) {
	defer qt.Recovering("QMediaPlayerControl::setMuted")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) SetPlaybackRate(rate float64) {
	defer qt.Recovering("QMediaPlayerControl::setPlaybackRate")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayerControl) SetPosition(position int64) {
	defer qt.Recovering("QMediaPlayerControl::setPosition")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPosition(ptr.Pointer(), C.longlong(position))
	}
}

func (ptr *QMediaPlayerControl) SetVolume(volume int) {
	defer qt.Recovering("QMediaPlayerControl::setVolume")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) State() QMediaPlayer__State {
	defer qt.Recovering("QMediaPlayerControl::state")

	if ptr.Pointer() != nil {
		return QMediaPlayer__State(C.QMediaPlayerControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	defer qt.Recovering("connect QMediaPlayerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerControlStateChanged
func callbackQMediaPlayerControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QMediaPlayer__State))(QMediaPlayer__State(state))
	}

}

func (ptr *QMediaPlayerControl) Stop() {
	defer qt.Recovering("QMediaPlayerControl::stop")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) ConnectVideoAvailableChanged(f func(video bool)) {
	defer qt.Recovering("connect QMediaPlayerControl::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVideoAvailableChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::videoAvailableChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerControlVideoAvailableChanged
func callbackQMediaPlayerControlVideoAvailableChanged(ptrName *C.char, video C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::videoAvailableChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged")
	if signal != nil {
		signal.(func(bool))(int(video) != 0)
	}

}

func (ptr *QMediaPlayerControl) Volume() int {
	defer qt.Recovering("QMediaPlayerControl::volume")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectVolumeChanged(f func(volume int)) {
	defer qt.Recovering("connect QMediaPlayerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVolumeChanged() {
	defer qt.Recovering("disconnect QMediaPlayerControl::volumeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerControlVolumeChanged
func callbackQMediaPlayerControlVolumeChanged(ptrName *C.char, volume C.int) {
	defer qt.Recovering("callback QMediaPlayerControl::volumeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "volumeChanged")
	if signal != nil {
		signal.(func(int))(int(volume))
	}

}

func (ptr *QMediaPlayerControl) DestroyQMediaPlayerControl() {
	defer qt.Recovering("QMediaPlayerControl::~QMediaPlayerControl")

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DestroyQMediaPlayerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
