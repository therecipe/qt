package quick

//#include "quick.h"
import "C"
import (
	"unsafe"
)

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterial {
	var n = new(QSGSimpleMaterial)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return ptr
}
