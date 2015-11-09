package widgets

//#include "qstyleoptionsizegrip.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionSizeGrip struct {
	QStyleOptionComplex
}

type QStyleOptionSizeGrip_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSizeGrip_PTR() *QStyleOptionSizeGrip
}

func PointerFromQStyleOptionSizeGrip(ptr QStyleOptionSizeGrip_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSizeGrip_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionSizeGripFromPointer(ptr unsafe.Pointer) *QStyleOptionSizeGrip {
	var n = new(QStyleOptionSizeGrip)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSizeGrip) QStyleOptionSizeGrip_PTR() *QStyleOptionSizeGrip {
	return ptr
}

//QStyleOptionSizeGrip::StyleOptionType
type QStyleOptionSizeGrip__StyleOptionType int64

var (
	QStyleOptionSizeGrip__Type = QStyleOptionSizeGrip__StyleOptionType(QStyleOption__SO_SizeGrip)
)

//QStyleOptionSizeGrip::StyleOptionVersion
type QStyleOptionSizeGrip__StyleOptionVersion int64

var (
	QStyleOptionSizeGrip__Version = QStyleOptionSizeGrip__StyleOptionVersion(1)
)

func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGripFromPointer(C.QStyleOptionSizeGrip_NewQStyleOptionSizeGrip())
}

func NewQStyleOptionSizeGrip2(other QStyleOptionSizeGrip_ITF) *QStyleOptionSizeGrip {
	return NewQStyleOptionSizeGripFromPointer(C.QStyleOptionSizeGrip_NewQStyleOptionSizeGrip2(PointerFromQStyleOptionSizeGrip(other)))
}
