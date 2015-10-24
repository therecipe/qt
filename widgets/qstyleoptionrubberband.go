package widgets

//#include "qstyleoptionrubberband.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionRubberBand struct {
	QStyleOption
}

type QStyleOptionRubberBandITF interface {
	QStyleOptionITF
	QStyleOptionRubberBandPTR() *QStyleOptionRubberBand
}

func PointerFromQStyleOptionRubberBand(ptr QStyleOptionRubberBandITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionRubberBandPTR().Pointer()
	}
	return nil
}

func QStyleOptionRubberBandFromPointer(ptr unsafe.Pointer) *QStyleOptionRubberBand {
	var n = new(QStyleOptionRubberBand)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionRubberBand) QStyleOptionRubberBandPTR() *QStyleOptionRubberBand {
	return ptr
}

//QStyleOptionRubberBand::StyleOptionType
type QStyleOptionRubberBand__StyleOptionType int

var (
	QStyleOptionRubberBand__Type = QStyleOptionRubberBand__StyleOptionType(QStyleOption__SO_RubberBand)
)

//QStyleOptionRubberBand::StyleOptionVersion
type QStyleOptionRubberBand__StyleOptionVersion int

var (
	QStyleOptionRubberBand__Version = QStyleOptionRubberBand__StyleOptionVersion(1)
)

func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	return QStyleOptionRubberBandFromPointer(unsafe.Pointer(C.QStyleOptionRubberBand_NewQStyleOptionRubberBand()))
}

func NewQStyleOptionRubberBand2(other QStyleOptionRubberBandITF) *QStyleOptionRubberBand {
	return QStyleOptionRubberBandFromPointer(unsafe.Pointer(C.QStyleOptionRubberBand_NewQStyleOptionRubberBand2(C.QtObjectPtr(PointerFromQStyleOptionRubberBand(other)))))
}
