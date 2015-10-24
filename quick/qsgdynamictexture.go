package quick

//#include "qsgdynamictexture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTextureITF interface {
	QSGTextureITF
	QSGDynamicTexturePTR() *QSGDynamicTexture
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTextureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexturePTR().Pointer()
	}
	return nil
}

func QSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSGDynamicTexture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGDynamicTexture) QSGDynamicTexturePTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}
