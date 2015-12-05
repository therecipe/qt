package quick

//#include "quick.h"
import "C"
import (
	"unsafe"
)

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGTextureMaterial {
	var n = new(QSGTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return ptr
}
