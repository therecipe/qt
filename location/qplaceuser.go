package location

//#include "qplaceuser.h"
import "C"
import (
	"unsafe"
)

type QPlaceUser struct {
	ptr unsafe.Pointer
}

type QPlaceUserITF interface {
	QPlaceUserPTR() *QPlaceUser
}

func (p *QPlaceUser) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceUser) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceUser(ptr QPlaceUserITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceUserPTR().Pointer()
	}
	return nil
}

func QPlaceUserFromPointer(ptr unsafe.Pointer) *QPlaceUser {
	var n = new(QPlaceUser)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceUser) QPlaceUserPTR() *QPlaceUser {
	return ptr
}
