package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGDynamicTexture_") {
		n.SetObjectName("QSGDynamicTexture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGDynamicTexture::updateTexture")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(ptr.Pointer()) != 0
	}
	return false
}
