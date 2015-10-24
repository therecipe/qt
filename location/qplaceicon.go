package location

//#include "qplaceicon.h"
import "C"
import (
	"unsafe"
)

type QPlaceIcon struct {
	ptr unsafe.Pointer
}

type QPlaceIconITF interface {
	QPlaceIconPTR() *QPlaceIcon
}

func (p *QPlaceIcon) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceIcon) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceIcon(ptr QPlaceIconITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceIconPTR().Pointer()
	}
	return nil
}

func QPlaceIconFromPointer(ptr unsafe.Pointer) *QPlaceIcon {
	var n = new(QPlaceIcon)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceIcon) QPlaceIconPTR() *QPlaceIcon {
	return ptr
}
