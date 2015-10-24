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

type QSGTextureITF interface {
	core.QObjectITF
	QSGTexturePTR() *QSGTexture
}

func PointerFromQSGTexture(ptr QSGTextureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexturePTR().Pointer()
	}
	return nil
}

func QSGTextureFromPointer(ptr unsafe.Pointer) *QSGTexture {
	var n = new(QSGTexture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGTexture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGTexture) QSGTexturePTR() *QSGTexture {
	return ptr
}

//QSGTexture::Filtering
type QSGTexture__Filtering int

var (
	QSGTexture__None    = QSGTexture__Filtering(0)
	QSGTexture__Nearest = QSGTexture__Filtering(1)
	QSGTexture__Linear  = QSGTexture__Filtering(2)
)

//QSGTexture::WrapMode
type QSGTexture__WrapMode int

var (
	QSGTexture__Repeat      = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge = QSGTexture__WrapMode(1)
)

func (ptr *QSGTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QSGTexture_Bind(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_Filtering(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasAlphaChannel(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGTexture) HasMipmaps() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_HasMipmaps(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_HorizontalWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTexture(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {
	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_MipmapFiltering(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	if ptr.Pointer() != nil {
		return QSGTextureFromPointer(unsafe.Pointer(C.QSGTexture_RemovedFromAtlas(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSGTexture) SetFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetFiltering(C.QtObjectPtr(ptr.Pointer()), C.int(filter))
	}
}

func (ptr *QSGTexture) SetHorizontalWrapMode(hwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetHorizontalWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(hwrap))
	}
}

func (ptr *QSGTexture) SetMipmapFiltering(filter QSGTexture__Filtering) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetMipmapFiltering(C.QtObjectPtr(ptr.Pointer()), C.int(filter))
	}
}

func (ptr *QSGTexture) SetVerticalWrapMode(vwrap QSGTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QSGTexture_SetVerticalWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(vwrap))
	}
}

func (ptr *QSGTexture) TextureId() int {
	if ptr.Pointer() != nil {
		return int(C.QSGTexture_TextureId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {
	if ptr.Pointer() != nil {
		C.QSGTexture_UpdateBindOptions(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(force)))
	}
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_VerticalWrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGTexture) DestroyQSGTexture() {
	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTexture(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
