package quick

//#include "quick.h"
import "C"
import (
	"unsafe"
)

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) *QSGVertexColorMaterial {
	var n = new(QSGVertexColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return ptr
}
