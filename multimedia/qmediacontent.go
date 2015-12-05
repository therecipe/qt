package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
	"unsafe"
)

type QMediaContent struct {
	ptr unsafe.Pointer
}

type QMediaContent_ITF interface {
	QMediaContent_PTR() *QMediaContent
}

func (p *QMediaContent) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaContent) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaContent(ptr QMediaContent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContent_PTR().Pointer()
	}
	return nil
}

func NewQMediaContentFromPointer(ptr unsafe.Pointer) *QMediaContent {
	var n = new(QMediaContent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaContent) QMediaContent_PTR() *QMediaContent {
	return ptr
}

func NewQMediaContent() *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent())
}

func NewQMediaContent7(playlist QMediaPlaylist_ITF, contentUrl core.QUrl_ITF, takeOwnership bool) *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent7(PointerFromQMediaPlaylist(playlist), core.PointerFromQUrl(contentUrl), C.int(qt.GoBoolToInt(takeOwnership))))
}

func NewQMediaContent6(other QMediaContent_ITF) *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent6(PointerFromQMediaContent(other)))
}

func NewQMediaContent4(resource QMediaResource_ITF) *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent4(PointerFromQMediaResource(resource)))
}

func NewQMediaContent3(request network.QNetworkRequest_ITF) *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent3(network.PointerFromQNetworkRequest(request)))
}

func NewQMediaContent2(url core.QUrl_ITF) *QMediaContent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::QMediaContent")
		}
	}()

	return NewQMediaContentFromPointer(C.QMediaContent_NewQMediaContent2(core.PointerFromQUrl(url)))
}

func (ptr *QMediaContent) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaContent_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaContent) Playlist() *QMediaPlaylist {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::playlist")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaPlaylistFromPointer(C.QMediaContent_Playlist(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaContent) DestroyQMediaContent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaContent::~QMediaContent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaContent_DestroyQMediaContent(ptr.Pointer())
	}
}
