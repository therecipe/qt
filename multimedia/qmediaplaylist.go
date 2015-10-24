package multimedia

//#include "qmediaplaylist.h"
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

type QMediaPlaylistITF interface {
	core.QObjectITF
	QMediaBindableInterfaceITF
	QMediaPlaylistPTR() *QMediaPlaylist
}

func (p *QMediaPlaylist) Pointer() unsafe.Pointer {
	return p.QObjectPTR().Pointer()
}

func (p *QMediaPlaylist) SetPointer(ptr unsafe.Pointer) {
	p.QObjectPTR().SetPointer(ptr)
	p.QMediaBindableInterfacePTR().SetPointer(ptr)
}

func PointerFromQMediaPlaylist(ptr QMediaPlaylistITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlaylistPTR().Pointer()
	}
	return nil
}

func QMediaPlaylistFromPointer(ptr unsafe.Pointer) *QMediaPlaylist {
	var n = new(QMediaPlaylist)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaPlaylist_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlaylist) QMediaPlaylistPTR() *QMediaPlaylist {
	return ptr
}

//QMediaPlaylist::Error
type QMediaPlaylist__Error int

var (
	QMediaPlaylist__NoError                 = QMediaPlaylist__Error(0)
	QMediaPlaylist__FormatError             = QMediaPlaylist__Error(1)
	QMediaPlaylist__FormatNotSupportedError = QMediaPlaylist__Error(2)
	QMediaPlaylist__NetworkError            = QMediaPlaylist__Error(3)
	QMediaPlaylist__AccessDeniedError       = QMediaPlaylist__Error(4)
)

//QMediaPlaylist::PlaybackMode
type QMediaPlaylist__PlaybackMode int

var (
	QMediaPlaylist__CurrentItemOnce   = QMediaPlaylist__PlaybackMode(0)
	QMediaPlaylist__CurrentItemInLoop = QMediaPlaylist__PlaybackMode(1)
	QMediaPlaylist__Sequential        = QMediaPlaylist__PlaybackMode(2)
	QMediaPlaylist__Loop              = QMediaPlaylist__PlaybackMode(3)
	QMediaPlaylist__Random            = QMediaPlaylist__PlaybackMode(4)
)

func (ptr *QMediaPlaylist) PlaybackMode() QMediaPlaylist__PlaybackMode {
	if ptr.Pointer() != nil {
		return QMediaPlaylist__PlaybackMode(C.QMediaPlaylist_PlaybackMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlaylist) SetPlaybackMode(mode QMediaPlaylist__PlaybackMode) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetPlaybackMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func NewQMediaPlaylist(parent core.QObjectITF) *QMediaPlaylist {
	return QMediaPlaylistFromPointer(unsafe.Pointer(C.QMediaPlaylist_NewQMediaPlaylist(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMediaPlaylist) AddMedia(content QMediaContentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_AddMedia(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMediaContent(content))) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Clear() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Clear(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectCurrentIndexChanged(f func(position int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentIndexChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQMediaPlaylistCurrentIndexChanged
func callbackQMediaPlaylistCurrentIndexChanged(ptrName *C.char, position C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(position))
}

func (ptr *QMediaPlaylist) Error() QMediaPlaylist__Error {
	if ptr.Pointer() != nil {
		return QMediaPlaylist__Error(C.QMediaPlaylist_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlaylist) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlaylist_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaPlaylist) InsertMedia(pos int, content QMediaContentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_InsertMedia(C.QtObjectPtr(ptr.Pointer()), C.int(pos), C.QtObjectPtr(PointerFromQMediaContent(content))) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Load3(device core.QIODeviceITF, format string) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load(request network.QNetworkRequestITF, format string) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkRequest(request)), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load2(location string, format string) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load2(C.QtObjectPtr(ptr.Pointer()), C.CString(location), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) ConnectLoadFailed(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoadFailed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "loadFailed", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoadFailed() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoadFailed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "loadFailed")
	}
}

//export callbackQMediaPlaylistLoadFailed
func callbackQMediaPlaylistLoadFailed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loadFailed").(func())()
}

func (ptr *QMediaPlaylist) ConnectLoaded(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoaded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "loaded", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoaded() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoaded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "loaded")
	}
}

//export callbackQMediaPlaylistLoaded
func callbackQMediaPlaylistLoaded(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "loaded").(func())()
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeInserted(f func(start int, end int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeInserted() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeInserted
func callbackQMediaPlaylistMediaAboutToBeInserted(ptrName *C.char, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeInserted").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeRemoved(f func(start int, end int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeRemoved() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeRemoved
func callbackQMediaPlaylistMediaAboutToBeRemoved(ptrName *C.char, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeRemoved").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) ConnectMediaChanged(f func(start int, end int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlaylistMediaChanged
func callbackQMediaPlaylistMediaChanged(ptrName *C.char, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaChanged").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) MediaCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_MediaCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectMediaInserted(f func(start int, end int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaInserted() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaInserted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaInserted")
	}
}

//export callbackQMediaPlaylistMediaInserted
func callbackQMediaPlaylistMediaInserted(ptrName *C.char, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaInserted").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return QMediaObjectFromPointer(unsafe.Pointer(C.QMediaPlaylist_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaRemoved(f func(start int, end int)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "mediaRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaRemoved() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "mediaRemoved")
	}
}

//export callbackQMediaPlaylistMediaRemoved
func callbackQMediaPlaylistMediaRemoved(ptrName *C.char, start C.int, end C.int) {
	qt.GetSignal(C.GoString(ptrName), "mediaRemoved").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) Next() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Next(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlaylist) NextIndex(steps int) int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_NextIndex(C.QtObjectPtr(ptr.Pointer()), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectPlaybackModeChanged(f func(mode QMediaPlaylist__PlaybackMode)) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectPlaybackModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "playbackModeChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectPlaybackModeChanged() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectPlaybackModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "playbackModeChanged")
	}
}

//export callbackQMediaPlaylistPlaybackModeChanged
func callbackQMediaPlaylistPlaybackModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "playbackModeChanged").(func(QMediaPlaylist__PlaybackMode))(QMediaPlaylist__PlaybackMode(mode))
}

func (ptr *QMediaPlaylist) Previous() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Previous(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlaylist) PreviousIndex(steps int) int {
	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_PreviousIndex(C.QtObjectPtr(ptr.Pointer()), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) RemoveMedia(pos int) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia(C.QtObjectPtr(ptr.Pointer()), C.int(pos)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) RemoveMedia2(start int, end int) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia2(C.QtObjectPtr(ptr.Pointer()), C.int(start), C.int(end)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save2(device core.QIODeviceITF, format string) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save(location string, format string) bool {
	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save(C.QtObjectPtr(ptr.Pointer()), C.CString(location), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) SetCurrentIndex(playlistPosition int) {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(playlistPosition))
	}
}

func (ptr *QMediaPlaylist) Shuffle() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Shuffle(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMediaPlaylist) DestroyQMediaPlaylist() {
	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DestroyQMediaPlaylist(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
