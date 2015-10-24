package quick

//#include "qsgtexturematerial.h"
import "C"
import (
	"unsafe"
)

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterialITF interface {
	QSGOpaqueTextureMaterialITF
	QSGTextureMaterialPTR() *QSGTextureMaterial
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterialPTR().Pointer()
	}
	return nil
}

func QSGTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGTextureMaterial {
	var n = new(QSGTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGTextureMaterial) QSGTextureMaterialPTR() *QSGTextureMaterial {
	return ptr
}
