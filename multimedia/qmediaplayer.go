package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMediaPlayer struct {
	QMediaObject
}

type QMediaPlayer_ITF interface {
	QMediaObject_ITF
	QMediaPlayer_PTR() *QMediaPlayer
}

func PointerFromQMediaPlayer(ptr QMediaPlayer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlayer_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlayerFromPointer(ptr unsafe.Pointer) *QMediaPlayer {
	var n = new(QMediaPlayer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlayer_") {
		n.SetObjectName("QMediaPlayer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlayer) QMediaPlayer_PTR() *QMediaPlayer {
	return ptr
}

//QMediaPlayer::Error
type QMediaPlayer__Error int64

const (
	QMediaPlayer__NoError             = QMediaPlayer__Error(0)
	QMediaPlayer__ResourceError       = QMediaPlayer__Error(1)
	QMediaPlayer__FormatError         = QMediaPlayer__Error(2)
	QMediaPlayer__NetworkError        = QMediaPlayer__Error(3)
	QMediaPlayer__AccessDeniedError   = QMediaPlayer__Error(4)
	QMediaPlayer__ServiceMissingError = QMediaPlayer__Error(5)
	QMediaPlayer__MediaIsPlaylist     = QMediaPlayer__Error(6)
)

//QMediaPlayer::Flag
type QMediaPlayer__Flag int64

const (
	QMediaPlayer__LowLatency     = QMediaPlayer__Flag(0x01)
	QMediaPlayer__StreamPlayback = QMediaPlayer__Flag(0x02)
	QMediaPlayer__VideoSurface   = QMediaPlayer__Flag(0x04)
)

//QMediaPlayer::MediaStatus
type QMediaPlayer__MediaStatus int64

const (
	QMediaPlayer__UnknownMediaStatus = QMediaPlayer__MediaStatus(0)
	QMediaPlayer__NoMedia            = QMediaPlayer__MediaStatus(1)
	QMediaPlayer__LoadingMedia       = QMediaPlayer__MediaStatus(2)
	QMediaPlayer__LoadedMedia        = QMediaPlayer__MediaStatus(3)
	QMediaPlayer__StalledMedia       = QMediaPlayer__MediaStatus(4)
	QMediaPlayer__BufferingMedia     = QMediaPlayer__MediaStatus(5)
	QMediaPlayer__BufferedMedia      = QMediaPlayer__MediaStatus(6)
	QMediaPlayer__EndOfMedia         = QMediaPlayer__MediaStatus(7)
	QMediaPlayer__InvalidMedia       = QMediaPlayer__MediaStatus(8)
)

//QMediaPlayer::State
type QMediaPlayer__State int64

const (
	QMediaPlayer__StoppedState = QMediaPlayer__State(0)
	QMediaPlayer__PlayingState = QMediaPlayer__State(1)
	QMediaPlayer__PausedState  = QMediaPlayer__State(2)
)

func (ptr *QMediaPlayer) BufferStatus() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::bufferStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayer_BufferStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlayer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaPlayer) IsAudioAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::isAudioAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsAudioAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsMuted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::isMuted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsSeekable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::isSeekable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsSeekable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) IsVideoAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::isVideoAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlayer_IsVideoAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlayer) MediaStatus() QMediaPlayer__MediaStatus {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mediaStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaPlayer__MediaStatus(C.QMediaPlayer_MediaStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) PlaybackRate() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::playbackRate")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QMediaPlayer_PlaybackRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) Playlist() *QMediaPlaylist {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::playlist")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaPlaylistFromPointer(C.QMediaPlayer_Playlist(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) SetMuted(muted bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setMuted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QMediaPlayer) SetPlaybackRate(rate float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setPlaybackRate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetPlaybackRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QMediaPlayer) SetPlaylist(playlist QMediaPlaylist_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setPlaylist")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetPlaylist(ptr.Pointer(), PointerFromQMediaPlaylist(playlist))
	}
}

func (ptr *QMediaPlayer) SetVolume(volume int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVolume(ptr.Pointer(), C.int(volume))
	}
}

func (ptr *QMediaPlayer) Volume() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlayer_Volume(ptr.Pointer()))
	}
	return 0
}

func NewQMediaPlayer(parent core.QObject_ITF, flags QMediaPlayer__Flag) *QMediaPlayer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::QMediaPlayer")
		}
	}()

	return NewQMediaPlayerFromPointer(C.QMediaPlayer_NewQMediaPlayer(core.PointerFromQObject(parent), C.int(flags)))
}

func (ptr *QMediaPlayer) ConnectAudioAvailableChanged(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::audioAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectAudioAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioAvailableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectAudioAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::audioAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectAudioAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioAvailableChanged")
	}
}

//export callbackQMediaPlayerAudioAvailableChanged
func callbackQMediaPlayerAudioAvailableChanged(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::audioAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "audioAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QMediaPlayer) ConnectBufferStatusChanged(f func(percentFilled int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::bufferStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectBufferStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferStatusChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectBufferStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::bufferStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectBufferStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferStatusChanged")
	}
}

//export callbackQMediaPlayerBufferStatusChanged
func callbackQMediaPlayerBufferStatusChanged(ptrName *C.char, percentFilled C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::bufferStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferStatusChanged").(func(int))(int(percentFilled))
}

func (ptr *QMediaPlayer) Error() QMediaPlayer__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaPlayer__Error(C.QMediaPlayer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlayer) ConnectMediaStatusChanged(f func(status QMediaPlayer__MediaStatus)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mediaStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectMediaStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaStatusChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectMediaStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mediaStatusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectMediaStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaStatusChanged")
	}
}

//export callbackQMediaPlayerMediaStatusChanged
func callbackQMediaPlayerMediaStatusChanged(ptrName *C.char, status C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mediaStatusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaStatusChanged").(func(QMediaPlayer__MediaStatus))(QMediaPlayer__MediaStatus(status))
}

func (ptr *QMediaPlayer) MediaStream() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mediaStream")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QMediaPlayer_MediaStream(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlayer) ConnectMutedChanged(f func(muted bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectMutedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQMediaPlayerMutedChanged
func callbackQMediaPlayerMutedChanged(ptrName *C.char, muted C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::mutedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func(bool))(int(muted) != 0)
}

func (ptr *QMediaPlayer) Pause() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::pause")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Pause(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) Play() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::play")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Play(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) ConnectSeekableChanged(f func(seekable bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::seekableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectSeekableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "seekableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectSeekableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::seekableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectSeekableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "seekableChanged")
	}
}

//export callbackQMediaPlayerSeekableChanged
func callbackQMediaPlayerSeekableChanged(ptrName *C.char, seekable C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::seekableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "seekableChanged").(func(bool))(int(seekable) != 0)
}

func (ptr *QMediaPlayer) SetMedia(media QMediaContent_ITF, stream core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setMedia")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetMedia(ptr.Pointer(), PointerFromQMediaContent(media), core.PointerFromQIODevice(stream))
	}
}

func (ptr *QMediaPlayer) SetVideoOutput3(surface QAbstractVideoSurface_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::setVideoOutput")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_SetVideoOutput3(ptr.Pointer(), PointerFromQAbstractVideoSurface(surface))
	}
}

func (ptr *QMediaPlayer) ConnectStateChanged(f func(state QMediaPlayer__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQMediaPlayerStateChanged
func callbackQMediaPlayerStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QMediaPlayer__State))(QMediaPlayer__State(state))
}

func (ptr *QMediaPlayer) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_Stop(ptr.Pointer())
	}
}

func (ptr *QMediaPlayer) ConnectVideoAvailableChanged(f func(videoAvailable bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::videoAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectVideoAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "videoAvailableChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectVideoAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::videoAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectVideoAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "videoAvailableChanged")
	}
}

//export callbackQMediaPlayerVideoAvailableChanged
func callbackQMediaPlayerVideoAvailableChanged(ptrName *C.char, videoAvailable C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::videoAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "videoAvailableChanged").(func(bool))(int(videoAvailable) != 0)
}

func (ptr *QMediaPlayer) ConnectVolumeChanged(f func(volume int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QMediaPlayer) DisconnectVolumeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQMediaPlayerVolumeChanged
func callbackQMediaPlayerVolumeChanged(ptrName *C.char, volume C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::volumeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func(int))(int(volume))
}

func (ptr *QMediaPlayer) DestroyQMediaPlayer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlayer::~QMediaPlayer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlayer_DestroyQMediaPlayer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
