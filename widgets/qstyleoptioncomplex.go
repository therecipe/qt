package widgets

//#include "qstyleoptioncomplex.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionComplex struct {
	QStyleOption
}

type QStyleOptionComplexITF interface {
	QStyleOptionITF
	QStyleOptionComplexPTR() *QStyleOptionComplex
}

func PointerFromQStyleOptionComplex(ptr QStyleOptionComplexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionComplexPTR().Pointer()
	}
	return nil
}

func QStyleOptionComplexFromPointer(ptr unsafe.Pointer) *QStyleOptionComplex {
	var n = new(QStyleOptionComplex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionComplex) QStyleOptionComplexPTR() *QStyleOptionComplex {
	return ptr
}

//QStyleOptionComplex::StyleOptionType
type QStyleOptionComplex__StyleOptionType int

var (
	QStyleOptionComplex__Type = QStyleOptionComplex__StyleOptionType(QStyleOption__SO_Complex)
)

//QStyleOptionComplex::StyleOptionVersion
type QStyleOptionComplex__StyleOptionVersion int

var (
	QStyleOptionComplex__Version = QStyleOptionComplex__StyleOptionVersion(1)
)

func NewQStyleOptionComplex2(other QStyleOptionComplexITF) *QStyleOptionComplex {
	return QStyleOptionComplexFromPointer(unsafe.Pointer(C.QStyleOptionComplex_NewQStyleOptionComplex2(C.QtObjectPtr(PointerFromQStyleOptionComplex(other)))))
}

func NewQStyleOptionComplex(version int, ty int) *QStyleOptionComplex {
	return QStyleOptionComplexFromPointer(unsafe.Pointer(C.QStyleOptionComplex_NewQStyleOptionComplex(C.int(version), C.int(ty))))
}
