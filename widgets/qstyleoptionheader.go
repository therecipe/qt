package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleOptionHeader struct {
	QStyleOption
}

type QStyleOptionHeader_ITF interface {
	QStyleOption_ITF
	QStyleOptionHeader_PTR() *QStyleOptionHeader
}

func PointerFromQStyleOptionHeader(ptr QStyleOptionHeader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionHeader_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionHeaderFromPointer(ptr unsafe.Pointer) *QStyleOptionHeader {
	var n = new(QStyleOptionHeader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionHeader) QStyleOptionHeader_PTR() *QStyleOptionHeader {
	return ptr
}

//QStyleOptionHeader::SectionPosition
type QStyleOptionHeader__SectionPosition int64

const (
	QStyleOptionHeader__Beginning      = QStyleOptionHeader__SectionPosition(0)
	QStyleOptionHeader__Middle         = QStyleOptionHeader__SectionPosition(1)
	QStyleOptionHeader__End            = QStyleOptionHeader__SectionPosition(2)
	QStyleOptionHeader__OnlyOneSection = QStyleOptionHeader__SectionPosition(3)
)

//QStyleOptionHeader::SelectedPosition
type QStyleOptionHeader__SelectedPosition int64

const (
	QStyleOptionHeader__NotAdjacent                = QStyleOptionHeader__SelectedPosition(0)
	QStyleOptionHeader__NextIsSelected             = QStyleOptionHeader__SelectedPosition(1)
	QStyleOptionHeader__PreviousIsSelected         = QStyleOptionHeader__SelectedPosition(2)
	QStyleOptionHeader__NextAndPreviousAreSelected = QStyleOptionHeader__SelectedPosition(3)
)

//QStyleOptionHeader::SortIndicator
type QStyleOptionHeader__SortIndicator int64

const (
	QStyleOptionHeader__None     = QStyleOptionHeader__SortIndicator(0)
	QStyleOptionHeader__SortUp   = QStyleOptionHeader__SortIndicator(1)
	QStyleOptionHeader__SortDown = QStyleOptionHeader__SortIndicator(2)
)

//QStyleOptionHeader::StyleOptionType
type QStyleOptionHeader__StyleOptionType int64

var (
	QStyleOptionHeader__Type = QStyleOptionHeader__StyleOptionType(QStyleOption__SO_Header)
)

//QStyleOptionHeader::StyleOptionVersion
type QStyleOptionHeader__StyleOptionVersion int64

var (
	QStyleOptionHeader__Version = QStyleOptionHeader__StyleOptionVersion(1)
)

func NewQStyleOptionHeader() *QStyleOptionHeader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionHeader::QStyleOptionHeader")
		}
	}()

	return NewQStyleOptionHeaderFromPointer(C.QStyleOptionHeader_NewQStyleOptionHeader())
}

func NewQStyleOptionHeader2(other QStyleOptionHeader_ITF) *QStyleOptionHeader {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionHeader::QStyleOptionHeader")
		}
	}()

	return NewQStyleOptionHeaderFromPointer(C.QStyleOptionHeader_NewQStyleOptionHeader2(PointerFromQStyleOptionHeader(other)))
}
