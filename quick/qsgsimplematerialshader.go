package quick

//#include "qsgsimplematerialshader.h"
import "C"
import (
	"unsafe"
)

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShaderITF interface {
	QSGMaterialShaderITF
	QSGSimpleMaterialShaderPTR() *QSGSimpleMaterialShader
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShaderPTR().Pointer()
	}
	return nil
}

func QSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterialShader {
	var n = new(QSGSimpleMaterialShader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterialShader) QSGSimpleMaterialShaderPTR() *QSGSimpleMaterialShader {
	return ptr
}
