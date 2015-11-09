package quick

//#include "qsgmaterialtype.h"
import "C"
import (
	"unsafe"
)

type QSGMaterialType struct {
	ptr unsafe.Pointer
}

type QSGMaterialType_ITF interface {
	QSGMaterialType_PTR() *QSGMaterialType
}

func (p *QSGMaterialType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialType(ptr QSGMaterialType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialType_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialTypeFromPointer(ptr unsafe.Pointer) *QSGMaterialType {
	var n = new(QSGMaterialType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return ptr
}
