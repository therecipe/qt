package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSourceLocation struct {
	ptr unsafe.Pointer
}

type QSourceLocation_ITF interface {
	QSourceLocation_PTR() *QSourceLocation
}

func (p *QSourceLocation) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSourceLocation) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSourceLocation(ptr QSourceLocation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSourceLocation_PTR().Pointer()
	}
	return nil
}

func NewQSourceLocationFromPointer(ptr unsafe.Pointer) *QSourceLocation {
	var n = new(QSourceLocation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSourceLocation) QSourceLocation_PTR() *QSourceLocation {
	return ptr
}

func NewQSourceLocation() *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation())
}

func NewQSourceLocation2(other QSourceLocation_ITF) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation2(PointerFromQSourceLocation(other)))
}

func NewQSourceLocation3(u core.QUrl_ITF, l int, c int) *QSourceLocation {
	defer qt.Recovering("QSourceLocation::QSourceLocation")

	return NewQSourceLocationFromPointer(C.QSourceLocation_NewQSourceLocation3(core.PointerFromQUrl(u), C.int(l), C.int(c)))
}

func (ptr *QSourceLocation) IsNull() bool {
	defer qt.Recovering("QSourceLocation::isNull")

	if ptr.Pointer() != nil {
		return C.QSourceLocation_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSourceLocation) SetUri(newUri core.QUrl_ITF) {
	defer qt.Recovering("QSourceLocation::setUri")

	if ptr.Pointer() != nil {
		C.QSourceLocation_SetUri(ptr.Pointer(), core.PointerFromQUrl(newUri))
	}
}

func (ptr *QSourceLocation) DestroyQSourceLocation() {
	defer qt.Recovering("QSourceLocation::~QSourceLocation")

	if ptr.Pointer() != nil {
		C.QSourceLocation_DestroyQSourceLocation(ptr.Pointer())
	}
}
