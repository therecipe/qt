package xmlpatterns

//#include "qsourcelocation.h"
import "C"
import (
	"unsafe"
)

type QSourceLocation struct {
	ptr unsafe.Pointer
}

type QSourceLocationITF interface {
	QSourceLocationPTR() *QSourceLocation
}

func (p *QSourceLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSourceLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSourceLocation(ptr QSourceLocationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocationPTR().Pointer()
	}
	return nil
}

func QSourceLocationFromPointer(ptr unsafe.Pointer) *QSourceLocation {
	var n = new(QSourceLocation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSourceLocation) QSourceLocationPTR() *QSourceLocation {
	return ptr
}

func NewQSourceLocation() *QSourceLocation {
	return QSourceLocationFromPointer(unsafe.Pointer(C.QSourceLocation_NewQSourceLocation()))
}

func NewQSourceLocation2(other QSourceLocationITF) *QSourceLocation {
	return QSourceLocationFromPointer(unsafe.Pointer(C.QSourceLocation_NewQSourceLocation2(C.QtObjectPtr(PointerFromQSourceLocation(other)))))
}

func NewQSourceLocation3(u string, l int, c int) *QSourceLocation {
	return QSourceLocationFromPointer(unsafe.Pointer(C.QSourceLocation_NewQSourceLocation3(C.CString(u), C.int(l), C.int(c))))
}

func (ptr *QSourceLocation) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSourceLocation_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSourceLocation) SetUri(newUri string) {
	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(C.QtObjectPtr(ptr.Pointer()), C.CString(newUri))
	}
}

func (ptr *QSourceLocation) Uri() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSourceLocation_Uri(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	if ptr.Pointer() != nil {
		C.QSourceLocation_DestroyQSourceLocation(C.QtObjectPtr(ptr.Pointer()))
	}
}
