package quick

//#include "qsgvertexcolormaterial.h"
import "C"
import (
	"unsafe"
)

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterialITF interface {
	QSGMaterialITF
	QSGVertexColorMaterialPTR() *QSGVertexColorMaterial
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterialPTR().Pointer()
	}
	return nil
}

func QSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) *QSGVertexColorMaterial {
	var n = new(QSGVertexColorMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterialPTR() *QSGVertexColorMaterial {
	return ptr
}
