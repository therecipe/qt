package quick

//#include "quick.h"
import "C"
import (
	"unsafe"
)

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterialShader {
	var n = new(QSGSimpleMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return ptr
}
