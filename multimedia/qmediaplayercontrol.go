package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QMediaPlayerControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlayerControl) QMediaPlayerControl_PTR() *QMediaPlayerControl {
	return ptr
}

func (ptr *QMediaPlayerControl) ConnectAudioAvailableChanged(f func(audio bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::audioAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectAudioAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::audioAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerControlAudioAvailableChanged
func callbackQMediaPlayerControlAudioAvailableChanged(ptrName *C.char, audio C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::audioAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged").(func(bool))(int(audio) != 0)
}

func (ptr *QMediaPlayerControl) BufferStatus() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::bufferStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectBufferStatusChanged(f func(progress int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::bufferStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectBufferStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::bufferStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerControlBufferStatusChanged
func callbackQMediaPlayerControlBufferStatusChanged(ptrName *C.char, progress C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::bufferStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged").(func(int))(int(progress))
}

func (ptr *QMediaPlayerControl) ConnectError(f func(error int, errorString string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQMediaPlayerControlError
func callbackQMediaPlayerControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(int, string))(int(error), C.GoString(errorString))
}

func (ptr *QMediaPlayerControl) IsAudioAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::isAudioAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsMuted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::isMuted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsSeekable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::isSeekable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) IsVideoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::isVideoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayerControl_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayerControl) MediaStatus() QMediaPlayer__MediaStatus {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mediaStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayerControl_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mediaStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMediaStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mediaStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerControlMediaStatusChanged
func callbackQMediaPlayerControlMediaStatusChanged(ptrName *C.char, status C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mediaStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged").(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
}

func (ptr *QMediaPlayerControl) MediaStream() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mediaStream")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayerControl_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayerControl) ConnectMutedChanged(f func(mute bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectMutedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerControlMutedChanged
func callbackQMediaPlayerControlMutedChanged(ptrName *C.char, mute C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::mutedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(mute) != 0)
}

func (ptr *QMediaPlayerControl) Pause() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::pause")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) Play() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::play")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) PlaybackRate() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::playbackRate")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayerControl_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectSeekableChanged(f func(seekable bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::seekableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectSeekableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::seekableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerControlSeekableChanged
func callbackQMediaPlayerControlSeekableChanged(ptrName *C.char, seekable C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::seekableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "seekableChanged").(func(bool))(int(seekable) != 0)
}

func (ptr *QMediaPlayerControl) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::setMedia")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayerControl) SetMuted(mute bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::setMuted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(mute)))
	}
}

func (ptr *QMediaPlayerControl) SetPlaybackRate(rate float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::setPlaybackRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayerControl) SetVolume(volume int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayerControl) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerControlStateChanged
func callbackQMediaPlayerControlStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaPlayer__State))(QMediaPlayer__State(state))
}

func (ptr *QMediaPlayerControl) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayerControl) ConnectVideoAvailableChanged(f func(video bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::videoAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVideoAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::videoAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerControlVideoAvailableChanged
func callbackQMediaPlayerControlVideoAvailableChanged(ptrName *C.char, video C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::videoAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged").(func(bool))(int(video) != 0)
}

func (ptr *QMediaPlayerControl) Volume() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayerControl_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayerControl) ConnectVolumeChanged(f func(volume int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayerControl) DisconnectVolumeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerControlVolumeChanged
func callbackQMediaPlayerControlVolumeChanged(ptrName *C.char, volume C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::volumeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QMediaPlayerControl) DestroyQMediaPlayerControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayerControl::~QMediaPlayerControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayerControl_DestroyQMediaPlayerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
