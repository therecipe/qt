package widgets

//#include "qstyleoptionheader.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionHeader struct {
	QStyleOption
}

type QStyleOptionHeaderITF interface {
	QStyleOptionITF
	QStyleOptionHeaderPTR() *QStyleOptionHeader
}

func PointerFromQStyleOptionHeader(ptr QStyleOptionHeaderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionHeaderPTR().Pointer()
	}
	return nil
}

func QStyleOptionHeaderFromPointer(ptr unsafe.Pointer) *QStyleOptionHeader {
	var n = new(QStyleOptionHeader)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionHeader) QStyleOptionHeaderPTR() *QStyleOptionHeader {
	return ptr
}

//QStyleOptionHeader::SectionPosition
type QStyleOptionHeader__SectionPosition int

var (
	QStyleOptionHeader__Beginning      = QStyleOptionHeader__SectionPosition(0)
	QStyleOptionHeader__Middle         = QStyleOptionHeader__SectionPosition(1)
	QStyleOptionHeader__End            = QStyleOptionHeader__SectionPosition(2)
	QStyleOptionHeader__OnlyOneSection = QStyleOptionHeader__SectionPosition(3)
)

//QStyleOptionHeader::SelectedPosition
type QStyleOptionHeader__SelectedPosition int

var (
	QStyleOptionHeader__NotAdjacent                = QStyleOptionHeader__SelectedPosition(0)
	QStyleOptionHeader__NextIsSelected             = QStyleOptionHeader__SelectedPosition(1)
	QStyleOptionHeader__PreviousIsSelected         = QStyleOptionHeader__SelectedPosition(2)
	QStyleOptionHeader__NextAndPreviousAreSelected = QStyleOptionHeader__SelectedPosition(3)
)

//QStyleOptionHeader::SortIndicator
type QStyleOptionHeader__SortIndicator int

var (
	QStyleOptionHeader__None     = QStyleOptionHeader__SortIndicator(0)
	QStyleOptionHeader__SortUp   = QStyleOptionHeader__SortIndicator(1)
	QStyleOptionHeader__SortDown = QStyleOptionHeader__SortIndicator(2)
)

//QStyleOptionHeader::StyleOptionType
type QStyleOptionHeader__StyleOptionType int

var (
	QStyleOptionHeader__Type = QStyleOptionHeader__StyleOptionType(QStyleOption__SO_Header)
)

//QStyleOptionHeader::StyleOptionVersion
type QStyleOptionHeader__StyleOptionVersion int

var (
	QStyleOptionHeader__Version = QStyleOptionHeader__StyleOptionVersion(1)
)

func NewQStyleOptionHeader() *QStyleOptionHeader {
	return QStyleOptionHeaderFromPointer(unsafe.Pointer(C.QStyleOptionHeader_NewQStyleOptionHeader()))
}

func NewQStyleOptionHeader2(other QStyleOptionHeaderITF) *QStyleOptionHeader {
	return QStyleOptionHeaderFromPointer(unsafe.Pointer(C.QStyleOptionHeader_NewQStyleOptionHeader2(C.QtObjectPtr(PointerFromQStyleOptionHeader(other)))))
}
