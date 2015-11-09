package widgets

//#include "qstyleoptionrubberband.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionRubberBand struct {
	QStyleOption
}

type QStyleOptionRubberBand_ITF interface {
	QStyleOption_ITF
	QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand
}

func PointerFromQStyleOptionRubberBand(ptr QStyleOptionRubberBand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionRubberBand_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionRubberBandFromPointer(ptr unsafe.Pointer) *QStyleOptionRubberBand {
	var n = new(QStyleOptionRubberBand)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionRubberBand) QStyleOptionRubberBand_PTR() *QStyleOptionRubberBand {
	return ptr
}

//QStyleOptionRubberBand::StyleOptionType
type QStyleOptionRubberBand__StyleOptionType int64

var (
	QStyleOptionRubberBand__Type = QStyleOptionRubberBand__StyleOptionType(QStyleOption__SO_RubberBand)
)

//QStyleOptionRubberBand::StyleOptionVersion
type QStyleOptionRubberBand__StyleOptionVersion int64

var (
	QStyleOptionRubberBand__Version = QStyleOptionRubberBand__StyleOptionVersion(1)
)

func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBandFromPointer(C.QStyleOptionRubberBand_NewQStyleOptionRubberBand())
}

func NewQStyleOptionRubberBand2(other QStyleOptionRubberBand_ITF) *QStyleOptionRubberBand {
	return NewQStyleOptionRubberBandFromPointer(C.QStyleOptionRubberBand_NewQStyleOptionRubberBand2(PointerFromQStyleOptionRubberBand(other)))
}
