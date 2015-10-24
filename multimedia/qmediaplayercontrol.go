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

type QMediaPlayerControlITF interface {
	QMediaControlITF
	QMediaPlayerControlPTR() *QMediaPlayerControl
}

func PointerFromQMediaPlayerControl(ptr QMediaPlayerControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlayerControlPTR().Pointer()
	}
	return nil
}

func QMediaPlayerControlFromPointer(ptr unsafe.Pointer) *QMediaPlayerControl {
	var n = new(QMediaPlayerControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaPlayerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlayerControl) QMediaPlayerControlPTR() *QMediaPlayerControl {
	return ptr
}

func (ptr *QMediaPlayerControl) ConnectAudioAvailableChanged(f func(audio bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAudioAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAudioAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAudioAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerControlAudioAvailableChanged
func callbackQMediaPlayerControlAudioAvailableChanged(ptrName *C.char, audio C.int) {
	qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged").(func(bool))(int(audio) != 0)
}

func (ptr *QMediaPlayerControl) BufferStatus() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_BufferStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectBufferStatusChanged(f func(progress int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectBufferStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectBufferStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectBufferStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerControlBufferStatusChanged
func callbackQMediaPlayerControlBufferStatusChanged(ptrName *C.char, progress C.int) {
	qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged").(func(int))(int(progress))
}

func (ptr *QMediaPlayerControl) ConnectError(f func(error int, errorString string)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaPlayerControlError
func callbackQMediaPlayerControlError(ptrName *C.char, error C.int, errorString *C.char) {
	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaPlayerControl) IsAudioAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsAudioAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsMuted() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsMuted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsSeekable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsSeekable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsVideoAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsVideoAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) MediaStatus() QMediaPlayer__MediaStatus {
	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayerControl_MediaStatus(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaStatusChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaStatusChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerControlMediaStatusChanged
func callbackQMediaPlayerControlMediaStatusChanged(ptrName *C.char, status C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged").(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
}

func (ptr *QMediaPlayerControl) MediaStream() *core.QIODevice {
	if ptr.Pointer() != nil {
		return core.QIODeviceFromPointer(unsafe.Pointer(C.QMediaPlayerControl_MediaStream(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMutedChanged(f func(mute bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMutedChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMutedChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerControlMutedChanged
func callbackQMediaPlayerControlMutedChanged(ptrName *C.char, mute C.int) {
	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(mute) != 0)
}

func (ptr *QMediaPlayerControl) Pause() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Pause(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlayerControl) Play() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Play(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlayerControl) ConnectSeekableChanged(f func(seekable bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectSeekableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectSeekableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectSeekableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerControlSeekableChanged
func callbackQMediaPlayerControlSeekableChanged(ptrName *C.char, seekable C.int) {
	qt.GetSignal(C.GoString(ptrName), "seekableChanged").(func(bool))(int(seekable) != 0)
}

func (ptr *QMediaPlayerControl) SetMedia(media QMediaContentITF, stream core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMedia(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaContent(media)), C.QtObjectPtr(core.PointerFromQIODevice(stream)))
	}
}

func (ptr *QMediaPlayerControl) SetMuted(mute bool) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMuted(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) SetVolume(volume int) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetVolume(C.QtObjectPtr(ptr.Pointer()), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerControlStateChanged
func callbackQMediaPlayerControlStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaPlayer__State))(QMediaPlayer__State(state))
}

func (ptr *QMediaPlayerControl) Stop() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlayerControl) ConnectVideoAvailableChanged(f func(video bool)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVideoAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVideoAvailableChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVideoAvailableChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerControlVideoAvailableChanged
func callbackQMediaPlayerControlVideoAvailableChanged(ptrName *C.char, video C.int) {
	qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged").(func(bool))(int(video) != 0)
}

func (ptr *QMediaPlayerControl) Volume() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_Volume(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVolumeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerControlVolumeChanged
func callbackQMediaPlayerControlVolumeChanged(ptrName *C.char, volume C.int) {
	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QMediaPlayerControl) DestroyQMediaPlayerControl() {
	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DestroyQMediaPlayerControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
