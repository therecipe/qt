package widgets

//#include "qstyleoptioncomplex.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionComplex struct {
	QStyleOption
}

type QStyleOptionComplex_ITF interface {
	QStyleOption_ITF
	QStyleOptionComplex_PTR() *QStyleOptionComplex
}

func PointerFromQStyleOptionComplex(ptr QStyleOptionComplex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionComplex_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionComplexFromPointer(ptr unsafe.Pointer) *QStyleOptionComplex {
	var n = new(QStyleOptionComplex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionComplex) QStyleOptionComplex_PTR() *QStyleOptionComplex {
	return ptr
}

//QStyleOptionComplex::StyleOptionType
type QStyleOptionComplex__StyleOptionType int64

var (
	QStyleOptionComplex__Type = QStyleOptionComplex__StyleOptionType(QStyleOption__SO_Complex)
)

//QStyleOptionComplex::StyleOptionVersion
type QStyleOptionComplex__StyleOptionVersion int64

var (
	QStyleOptionComplex__Version = QStyleOptionComplex__StyleOptionVersion(1)
)

func NewQStyleOptionComplex2(other QStyleOptionComplex_ITF) *QStyleOptionComplex {
	return NewQStyleOptionComplexFromPointer(C.QStyleOptionComplex_NewQStyleOptionComplex2(PointerFromQStyleOptionComplex(other)))
}

func NewQStyleOptionComplex(version int, ty int) *QStyleOptionComplex {
	return NewQStyleOptionComplexFromPointer(C.QStyleOptionComplex_NewQStyleOptionComplex(C.int(version), C.int(ty)))
}
