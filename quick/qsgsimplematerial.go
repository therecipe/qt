package quick

//#include "qsgsimplematerial.h"
import "C"
import (
	"unsafe"
)

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterialITF interface {
	QSGMaterialITF
	QSGSimpleMaterialPTR() *QSGSimpleMaterial
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterialITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialPTR().Pointer()
	}
	return nil
}

func QSGSimpleMaterialFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterial {
	var n = new(QSGSimpleMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterial) QSGSimpleMaterialPTR() *QSGSimpleMaterial {
	return ptr
}
