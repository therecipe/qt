package widgets

//#include "qstyleoptionsizegrip.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionSizeGrip struct {
	QStyleOptionComplex
}

type QStyleOptionSizeGripITF interface {
	QStyleOptionComplexITF
	QStyleOptionSizeGripPTR() *QStyleOptionSizeGrip
}

func PointerFromQStyleOptionSizeGrip(ptr QStyleOptionSizeGripITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSizeGripPTR().Pointer()
	}
	return nil
}

func QStyleOptionSizeGripFromPointer(ptr unsafe.Pointer) *QStyleOptionSizeGrip {
	var n = new(QStyleOptionSizeGrip)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSizeGrip) QStyleOptionSizeGripPTR() *QStyleOptionSizeGrip {
	return ptr
}

//QStyleOptionSizeGrip::StyleOptionType
type QStyleOptionSizeGrip__StyleOptionType int

var (
	QStyleOptionSizeGrip__Type = QStyleOptionSizeGrip__StyleOptionType(QStyleOption__SO_SizeGrip)
)

//QStyleOptionSizeGrip::StyleOptionVersion
type QStyleOptionSizeGrip__StyleOptionVersion int

var (
	QStyleOptionSizeGrip__Version = QStyleOptionSizeGrip__StyleOptionVersion(1)
)

func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	return QStyleOptionSizeGripFromPointer(unsafe.Pointer(C.QStyleOptionSizeGrip_NewQStyleOptionSizeGrip()))
}

func NewQStyleOptionSizeGrip2(other QStyleOptionSizeGripITF) *QStyleOptionSizeGrip {
	return QStyleOptionSizeGripFromPointer(unsafe.Pointer(C.QStyleOptionSizeGrip_NewQStyleOptionSizeGrip2(C.QtObjectPtr(PointerFromQStyleOptionSizeGrip(other)))))
}
