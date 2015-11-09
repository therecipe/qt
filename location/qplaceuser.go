package location

//#include "qplaceuser.h"
import "C"
import (
	"unsafe"
)

type QPlaceUser struct {
	ptr unsafe.Pointer
}

type QPlaceUser_ITF interface {
	QPlaceUser_PTR() *QPlaceUser
}

func (p *QPlaceUser) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceUser) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceUser(ptr QPlaceUser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceUser_PTR().Pointer()
	}
	return nil
}

func NewQPlaceUserFromPointer(ptr unsafe.Pointer) *QPlaceUser {
	var n = new(QPlaceUser)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceUser) QPlaceUser_PTR() *QPlaceUser {
	return ptr
}
