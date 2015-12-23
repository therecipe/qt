package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QMediaPlaylist struct {
	core.QObject
	QMediaBindableInterface
}

type QMediaPlaylist_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaPlaylist_PTR() *QMediaPlaylist
}

func (p *QMediaPlaylist) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QMediaPlaylist) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQMediaPlaylist(ptr QMediaPlaylist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlaylist_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlaylistFromPointer(ptr unsafe.Pointer) *QMediaPlaylist {
	var n = new(QMediaPlaylist)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlaylist_") {
		n.SetObjectName("QMediaPlaylist_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaPlaylist) QMediaPlaylist_PTR() *QMediaPlaylist {
	return ptr
}

//QMediaPlaylist::Error
type QMediaPlaylist__Error int64

const (
	QMediaPlaylist__NoError                 = QMediaPlaylist__Error(0)
	QMediaPlaylist__FormatError             = QMediaPlaylist__Error(1)
	QMediaPlaylist__FormatNotSupportedError = QMediaPlaylist__Error(2)
	QMediaPlaylist__NetworkError            = QMediaPlaylist__Error(3)
	QMediaPlaylist__AccessDeniedError       = QMediaPlaylist__Error(4)
)

//QMediaPlaylist::PlaybackMode
type QMediaPlaylist__PlaybackMode int64

const (
	QMediaPlaylist__CurrentItemOnce   = QMediaPlaylist__PlaybackMode(0)
	QMediaPlaylist__CurrentItemInLoop = QMediaPlaylist__PlaybackMode(1)
	QMediaPlaylist__Sequential        = QMediaPlaylist__PlaybackMode(2)
	QMediaPlaylist__Loop              = QMediaPlaylist__PlaybackMode(3)
	QMediaPlaylist__Random            = QMediaPlaylist__PlaybackMode(4)
)

func (ptr *QMediaPlaylist) PlaybackMode() QMediaPlaylist__PlaybackMode {
	defer qt.Recovering("QMediaPlaylist::playbackMode")

	if ptr.Pointer() != nil {
		return QMediaPlaylist__PlaybackMode(C.QMediaPlaylist_PlaybackMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) SetPlaybackMode(mode QMediaPlaylist__PlaybackMode) {
	defer qt.Recovering("QMediaPlaylist::setPlaybackMode")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetPlaybackMode(ptr.Pointer(), C.int(mode))
	}
}

func NewQMediaPlaylist(parent core.QObject_ITF) *QMediaPlaylist {
	defer qt.Recovering("QMediaPlaylist::QMediaPlaylist")

	return NewQMediaPlaylistFromPointer(C.QMediaPlaylist_NewQMediaPlaylist(core.PointerFromQObject(parent)))
}

func (ptr *QMediaPlaylist) AddMedia(content QMediaContent_ITF) bool {
	defer qt.Recovering("QMediaPlaylist::addMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_AddMedia(ptr.Pointer(), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Clear() bool {
	defer qt.Recovering("QMediaPlaylist::clear")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) CurrentIndex() int {
	defer qt.Recovering("QMediaPlaylist::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectCurrentIndexChanged(f func(position int)) {
	defer qt.Recovering("connect QMediaPlaylist::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentIndexChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::currentIndexChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQMediaPlaylistCurrentIndexChanged
func callbackQMediaPlaylistCurrentIndexChanged(ptrName *C.char, position C.int) {
	defer qt.Recovering("callback QMediaPlaylist::currentIndexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIndexChanged"); signal != nil {
		signal.(func(int))(int(position))
	}

}

func (ptr *QMediaPlaylist) CurrentMedia() *QMediaContent {
	defer qt.Recovering("QMediaPlaylist::currentMedia")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlaylist_CurrentMedia(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectCurrentMediaChanged(f func(content *QMediaContent)) {
	defer qt.Recovering("connect QMediaPlaylist::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentMediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::currentMediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentMediaChanged")
	}
}

//export callbackQMediaPlaylistCurrentMediaChanged
func callbackQMediaPlaylistCurrentMediaChanged(ptrName *C.char, content unsafe.Pointer) {
	defer qt.Recovering("callback QMediaPlaylist::currentMediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentMediaChanged"); signal != nil {
		signal.(func(*QMediaContent))(NewQMediaContentFromPointer(content))
	}

}

func (ptr *QMediaPlaylist) Error() QMediaPlaylist__Error {
	defer qt.Recovering("QMediaPlaylist::error")

	if ptr.Pointer() != nil {
		return QMediaPlaylist__Error(C.QMediaPlaylist_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ErrorString() string {
	defer qt.Recovering("QMediaPlaylist::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlaylist_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaPlaylist) InsertMedia(pos int, content QMediaContent_ITF) bool {
	defer qt.Recovering("QMediaPlaylist::insertMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_InsertMedia(ptr.Pointer(), C.int(pos), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsEmpty() bool {
	defer qt.Recovering("QMediaPlaylist::isEmpty")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsReadOnly() bool {
	defer qt.Recovering("QMediaPlaylist::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Load3(device core.QIODevice_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load3(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load(request network.QNetworkRequest_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load2(location core.QUrl_ITF, format string) {
	defer qt.Recovering("QMediaPlaylist::load")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load2(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) ConnectLoadFailed(f func()) {
	defer qt.Recovering("connect QMediaPlaylist::loadFailed")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoadFailed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFailed", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoadFailed() {
	defer qt.Recovering("disconnect QMediaPlaylist::loadFailed")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoadFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFailed")
	}
}

//export callbackQMediaPlaylistLoadFailed
func callbackQMediaPlaylistLoadFailed(ptrName *C.char) {
	defer qt.Recovering("callback QMediaPlaylist::loadFailed")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFailed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaPlaylist) ConnectLoaded(f func()) {
	defer qt.Recovering("connect QMediaPlaylist::loaded")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoaded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loaded", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoaded() {
	defer qt.Recovering("disconnect QMediaPlaylist::loaded")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoaded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loaded")
	}
}

//export callbackQMediaPlaylistLoaded
func callbackQMediaPlaylistLoaded(ptrName *C.char) {
	defer qt.Recovering("callback QMediaPlaylist::loaded")

	if signal := qt.GetSignal(C.GoString(ptrName), "loaded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaPlaylist) Media(index int) *QMediaContent {
	defer qt.Recovering("QMediaPlaylist::media")

	if ptr.Pointer() != nil {
		return NewQMediaContentFromPointer(C.QMediaPlaylist_Media(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeInserted(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeInserted() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaAboutToBeInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeInserted
func callbackQMediaPlaylistMediaAboutToBeInserted(ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaAboutToBeInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeInserted"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeRemoved(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeRemoved() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeRemoved
func callbackQMediaPlaylistMediaAboutToBeRemoved(ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeRemoved"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) ConnectMediaChanged(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlaylistMediaChanged
func callbackQMediaPlaylistMediaChanged(ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaChanged"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaCount() int {
	defer qt.Recovering("QMediaPlaylist::mediaCount")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_MediaCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectMediaInserted(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaInserted() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaInserted")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaInserted")
	}
}

//export callbackQMediaPlaylistMediaInserted
func callbackQMediaPlaylistMediaInserted(ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaInserted"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) MediaObject() *QMediaObject {
	defer qt.Recovering("QMediaPlaylist::mediaObject")

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaPlaylist_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaRemoved(f func(start int, end int)) {
	defer qt.Recovering("connect QMediaPlaylist::mediaRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaRemoved() {
	defer qt.Recovering("disconnect QMediaPlaylist::mediaRemoved")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaRemoved")
	}
}

//export callbackQMediaPlaylistMediaRemoved
func callbackQMediaPlaylistMediaRemoved(ptrName *C.char, start C.int, end C.int) {
	defer qt.Recovering("callback QMediaPlaylist::mediaRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "mediaRemoved"); signal != nil {
		signal.(func(int, int))(int(start), int(end))
	}

}

func (ptr *QMediaPlaylist) Next() {
	defer qt.Recovering("QMediaPlaylist::next")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Next(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) NextIndex(steps int) int {
	defer qt.Recovering("QMediaPlaylist::nextIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_NextIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectPlaybackModeChanged(f func(mode QMediaPlaylist__PlaybackMode)) {
	defer qt.Recovering("connect QMediaPlaylist::playbackModeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectPlaybackModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackModeChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectPlaybackModeChanged() {
	defer qt.Recovering("disconnect QMediaPlaylist::playbackModeChanged")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectPlaybackModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackModeChanged")
	}
}

//export callbackQMediaPlaylistPlaybackModeChanged
func callbackQMediaPlaylistPlaybackModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QMediaPlaylist::playbackModeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "playbackModeChanged"); signal != nil {
		signal.(func(QMediaPlaylist__PlaybackMode))(QMediaPlaylist__PlaybackMode(mode))
	}

}

func (ptr *QMediaPlaylist) Previous() {
	defer qt.Recovering("QMediaPlaylist::previous")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Previous(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) PreviousIndex(steps int) int {
	defer qt.Recovering("QMediaPlaylist::previousIndex")

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_PreviousIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) RemoveMedia(pos int) bool {
	defer qt.Recovering("QMediaPlaylist::removeMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia(ptr.Pointer(), C.int(pos)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) RemoveMedia2(start int, end int) bool {
	defer qt.Recovering("QMediaPlaylist::removeMedia")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia2(ptr.Pointer(), C.int(start), C.int(end)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save2(device core.QIODevice_ITF, format string) bool {
	defer qt.Recovering("QMediaPlaylist::save")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save(location core.QUrl_ITF, format string) bool {
	defer qt.Recovering("QMediaPlaylist::save")

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) SetCurrentIndex(playlistPosition int) {
	defer qt.Recovering("QMediaPlaylist::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetCurrentIndex(ptr.Pointer(), C.int(playlistPosition))
	}
}

func (ptr *QMediaPlaylist) Shuffle() {
	defer qt.Recovering("QMediaPlaylist::shuffle")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Shuffle(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) DestroyQMediaPlaylist() {
	defer qt.Recovering("QMediaPlaylist::~QMediaPlaylist")

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DestroyQMediaPlaylist(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaPlaylist) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaPlaylistTimerEvent
func callbackQMediaPlaylistTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaPlaylist::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaPlaylist) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaPlaylistChildEvent
func callbackQMediaPlaylistChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaPlaylist::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaPlaylist) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaPlaylist::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaPlaylistCustomEvent
func callbackQMediaPlaylistCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaPlaylist::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
