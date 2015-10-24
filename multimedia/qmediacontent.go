package multimedia

//#include "qmediacontent.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QMediaContent struct {
	ptr unsafe.Pointer
}

type QMediaContentITF interface {
	QMediaContentPTR() *QMediaContent
}

func (p *QMediaContent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaContent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaContent(ptr QMediaContentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContentPTR().Pointer()
	}
	return nil
}

func QMediaContentFromPointer(ptr unsafe.Pointer) *QMediaContent {
	var n = new(QMediaContent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaContent) QMediaContentPTR() *QMediaContent {
	return ptr
}

func NewQMediaContent() *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent()))
}

func NewQMediaContent7(playlist QMediaPlaylistITF, contentUrl string, takeOwnership bool) *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent7(C.QtObjectPtr(PointerFromQMediaPlaylist(playlist)), C.CString(contentUrl), C.int(qt.GoBoolToInt(takeOwnership)))))
}

func NewQMediaContent6(other QMediaContentITF) *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent6(C.QtObjectPtr(PointerFromQMediaContent(other)))))
}

func NewQMediaContent4(resource QMediaResourceITF) *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent4(C.QtObjectPtr(PointerFromQMediaResource(resource)))))
}

func NewQMediaContent3(request network.QNetworkRequestITF) *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent3(C.QtObjectPtr(network.PointerFromQNetworkRequest(request)))))
}

func NewQMediaContent2(url string) *QMediaContent {
	return QMediaContentFromPointer(unsafe.Pointer(C.QMediaContent_NewQMediaContent2(C.CString(url))))
}

func (ptr *QMediaContent) CanonicalUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContent_CanonicalUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QMediaContent) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QMediaContent_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMediaContent) Playlist() *QMediaPlaylist {
	if ptr.Pointer() != nil {
		return QMediaPlaylistFromPointer(unsafe.Pointer(C.QMediaContent_Playlist(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMediaContent) DestroyQMediaContent() {
	if ptr.Pointer() != nil {
		C.QMediaContent_DestroyQMediaContent(C.QtObjectPtr(ptr.Pointer()))
	}
}
