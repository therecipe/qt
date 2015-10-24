package quick

//#include "qsgmaterialtype.h"
import "C"
import (
	"unsafe"
)

type QSGMaterialType struct {
	ptr unsafe.Pointer
}

type QSGMaterialTypeITF interface {
	QSGMaterialTypePTR() *QSGMaterialType
}

func (p *QSGMaterialType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialType(ptr QSGMaterialTypeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialTypePTR().Pointer()
	}
	return nil
}

func QSGMaterialTypeFromPointer(ptr unsafe.Pointer) *QSGMaterialType {
	var n = new(QSGMaterialType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialType) QSGMaterialTypePTR() *QSGMaterialType {
	return ptr
}
