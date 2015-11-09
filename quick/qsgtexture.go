package quick

//#include "qsgtexture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func PointerFromQSGTexture(ptr QSGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureFromPointer(ptr unsafe.Pointer) *QSGTexture {
	var n = new(QSGTexture)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSGTexture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return ptr
}

//QSGTexture::Filtering
type QSGTexture__Filtering int64

const (
	QSGTexture__None    = QSGTexture__Filtering(0)
	QSGTexture__Nearest = QSGTexture__Filtering(1)
	QSGTexture__Linear  = QSGTexture__Filtering(2)
)

//QSGTexture::WrapMode
type QSGTexture__WrapMode int64

const (
	QSGTexture__Repeat      = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge = QSGTexture__WrapMode(1)
)

func (ptr *QSGTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGTexture_Bind(ptr.Pointer())
	}
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasMipmaps(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlas(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) SetFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetFiltering(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QSGTexture) SetHorizontalWrapMode(hwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetHorizontalWrapMode(ptr.Pointer(), C.int(hwrap))
	}
}

func (ptr *QSGTexture) SetMipmapFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetMipmapFiltering(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QSGTexture) SetVerticalWrapMode(vwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetVerticalWrapMode(ptr.Pointer(), C.int(vwrap))
	}
}

func (ptr *QSGTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(C.QSGTexture_TextureId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {
	if ptr.Pointer() != nil {
		C.QSGTexture_UpdateBindOptions(ptr.Pointer(), C.int(qt.GoBoolToInt(force)))
	}
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) DestroyQSGTexture() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTexture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
