package location

//#include "qplaceicon.h"
import "C"
import (
	"unsafe"
)

type QPlaceIcon struct {
	ptr unsafe.Pointer
}

type QPlaceIcon_ITF interface {
	QPlaceIcon_PTR() *QPlaceIcon
}

func (p *QPlaceIcon) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceIcon) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceIcon(ptr QPlaceIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIcon_PTR().Pointer()
	}
	return nil
}

func NewQPlaceIconFromPointer(ptr unsafe.Pointer) *QPlaceIcon {
	var n = new(QPlaceIcon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceIcon) QPlaceIcon_PTR() *QPlaceIcon {
	return ptr
}
