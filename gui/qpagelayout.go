package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPageLayout struct {
	ptr unsafe.Pointer
}

type QPageLayout_ITF interface {
	QPageLayout_PTR() *QPageLayout
}

func (p *QPageLayout) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPageLayout) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPageLayout(ptr QPageLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageLayout_PTR().Pointer()
	}
	return nil
}

func NewQPageLayoutFromPointer(ptr unsafe.Pointer) *QPageLayout {
	var n = new(QPageLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPageLayout) QPageLayout_PTR() *QPageLayout {
	return ptr
}

//QPageLayout::Mode
type QPageLayout__Mode int64

const (
	QPageLayout__StandardMode = QPageLayout__Mode(0)
	QPageLayout__FullPageMode = QPageLayout__Mode(1)
)

//QPageLayout::Orientation
type QPageLayout__Orientation int64

const (
	QPageLayout__Portrait  = QPageLayout__Orientation(0)
	QPageLayout__Landscape = QPageLayout__Orientation(1)
)

//QPageLayout::Unit
type QPageLayout__Unit int64

const (
	QPageLayout__Millimeter = QPageLayout__Unit(0)
	QPageLayout__Point      = QPageLayout__Unit(1)
	QPageLayout__Inch       = QPageLayout__Unit(2)
	QPageLayout__Pica       = QPageLayout__Unit(3)
	QPageLayout__Didot      = QPageLayout__Unit(4)
	QPageLayout__Cicero     = QPageLayout__Unit(5)
)

func NewQPageLayout() *QPageLayout {
	defer qt.Recovering("QPageLayout::QPageLayout")

	return NewQPageLayoutFromPointer(C.QPageLayout_NewQPageLayout())
}

func NewQPageLayout3(other QPageLayout_ITF) *QPageLayout {
	defer qt.Recovering("QPageLayout::QPageLayout")

	return NewQPageLayoutFromPointer(C.QPageLayout_NewQPageLayout3(PointerFromQPageLayout(other)))
}

func NewQPageLayout2(pageSize QPageSize_ITF, orientation QPageLayout__Orientation, margins core.QMarginsF_ITF, units QPageLayout__Unit, minMargins core.QMarginsF_ITF) *QPageLayout {
	defer qt.Recovering("QPageLayout::QPageLayout")

	return NewQPageLayoutFromPointer(C.QPageLayout_NewQPageLayout2(PointerFromQPageSize(pageSize), C.int(orientation), core.PointerFromQMarginsF(margins), C.int(units), core.PointerFromQMarginsF(minMargins)))
}

func (ptr *QPageLayout) IsEquivalentTo(other QPageLayout_ITF) bool {
	defer qt.Recovering("QPageLayout::isEquivalentTo")

	if ptr.Pointer() != nil {
		return C.QPageLayout_IsEquivalentTo(ptr.Pointer(), PointerFromQPageLayout(other)) != 0
	}
	return false
}

func (ptr *QPageLayout) IsValid() bool {
	defer qt.Recovering("QPageLayout::isValid")

	if ptr.Pointer() != nil {
		return C.QPageLayout_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPageLayout) Mode() QPageLayout__Mode {
	defer qt.Recovering("QPageLayout::mode")

	if ptr.Pointer() != nil {
		return QPageLayout__Mode(C.QPageLayout_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPageLayout) Orientation() QPageLayout__Orientation {
	defer qt.Recovering("QPageLayout::orientation")

	if ptr.Pointer() != nil {
		return QPageLayout__Orientation(C.QPageLayout_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	defer qt.Recovering("QPageLayout::setBottomMargin")

	if ptr.Pointer() != nil {
		return C.QPageLayout_SetBottomMargin(ptr.Pointer(), C.double(bottomMargin)) != 0
	}
	return false
}

func (ptr *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	defer qt.Recovering("QPageLayout::setLeftMargin")

	if ptr.Pointer() != nil {
		return C.QPageLayout_SetLeftMargin(ptr.Pointer(), C.double(leftMargin)) != 0
	}
	return false
}

func (ptr *QPageLayout) SetMargins(margins core.QMarginsF_ITF) bool {
	defer qt.Recovering("QPageLayout::setMargins")

	if ptr.Pointer() != nil {
		return C.QPageLayout_SetMargins(ptr.Pointer(), core.PointerFromQMarginsF(margins)) != 0
	}
	return false
}

func (ptr *QPageLayout) SetMinimumMargins(minMargins core.QMarginsF_ITF) {
	defer qt.Recovering("QPageLayout::setMinimumMargins")

	if ptr.Pointer() != nil {
		C.QPageLayout_SetMinimumMargins(ptr.Pointer(), core.PointerFromQMarginsF(minMargins))
	}
}

func (ptr *QPageLayout) SetMode(mode QPageLayout__Mode) {
	defer qt.Recovering("QPageLayout::setMode")

	if ptr.Pointer() != nil {
		C.QPageLayout_SetMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QPageLayout) SetOrientation(orientation QPageLayout__Orientation) {
	defer qt.Recovering("QPageLayout::setOrientation")

	if ptr.Pointer() != nil {
		C.QPageLayout_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QPageLayout) SetPageSize(pageSize QPageSize_ITF, minMargins core.QMarginsF_ITF) {
	defer qt.Recovering("QPageLayout::setPageSize")

	if ptr.Pointer() != nil {
		C.QPageLayout_SetPageSize(ptr.Pointer(), PointerFromQPageSize(pageSize), core.PointerFromQMarginsF(minMargins))
	}
}

func (ptr *QPageLayout) SetRightMargin(rightMargin float64) bool {
	defer qt.Recovering("QPageLayout::setRightMargin")

	if ptr.Pointer() != nil {
		return C.QPageLayout_SetRightMargin(ptr.Pointer(), C.double(rightMargin)) != 0
	}
	return false
}

func (ptr *QPageLayout) SetTopMargin(topMargin float64) bool {
	defer qt.Recovering("QPageLayout::setTopMargin")

	if ptr.Pointer() != nil {
		return C.QPageLayout_SetTopMargin(ptr.Pointer(), C.double(topMargin)) != 0
	}
	return false
}

func (ptr *QPageLayout) SetUnits(units QPageLayout__Unit) {
	defer qt.Recovering("QPageLayout::setUnits")

	if ptr.Pointer() != nil {
		C.QPageLayout_SetUnits(ptr.Pointer(), C.int(units))
	}
}

func (ptr *QPageLayout) Swap(other QPageLayout_ITF) {
	defer qt.Recovering("QPageLayout::swap")

	if ptr.Pointer() != nil {
		C.QPageLayout_Swap(ptr.Pointer(), PointerFromQPageLayout(other))
	}
}

func (ptr *QPageLayout) Units() QPageLayout__Unit {
	defer qt.Recovering("QPageLayout::units")

	if ptr.Pointer() != nil {
		return QPageLayout__Unit(C.QPageLayout_Units(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPageLayout) DestroyQPageLayout() {
	defer qt.Recovering("QPageLayout::~QPageLayout")

	if ptr.Pointer() != nil {
		C.QPageLayout_DestroyQPageLayout(ptr.Pointer())
	}
}
