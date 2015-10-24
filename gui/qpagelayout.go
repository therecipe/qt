package gui

//#include "qpagelayout.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPageLayout struct {
	ptr unsafe.Pointer
}

type QPageLayoutITF interface {
	QPageLayoutPTR() *QPageLayout
}

func (p *QPageLayout) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPageLayout) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPageLayout(ptr QPageLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageLayoutPTR().Pointer()
	}
	return nil
}

func QPageLayoutFromPointer(ptr unsafe.Pointer) *QPageLayout {
	var n = new(QPageLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPageLayout) QPageLayoutPTR() *QPageLayout {
	return ptr
}

//QPageLayout::Mode
type QPageLayout__Mode int

var (
	QPageLayout__StandardMode = QPageLayout__Mode(0)
	QPageLayout__FullPageMode = QPageLayout__Mode(1)
)

//QPageLayout::Orientation
type QPageLayout__Orientation int

var (
	QPageLayout__Portrait  = QPageLayout__Orientation(0)
	QPageLayout__Landscape = QPageLayout__Orientation(1)
)

//QPageLayout::Unit
type QPageLayout__Unit int

var (
	QPageLayout__Millimeter = QPageLayout__Unit(0)
	QPageLayout__Point      = QPageLayout__Unit(1)
	QPageLayout__Inch       = QPageLayout__Unit(2)
	QPageLayout__Pica       = QPageLayout__Unit(3)
	QPageLayout__Didot      = QPageLayout__Unit(4)
	QPageLayout__Cicero     = QPageLayout__Unit(5)
)

func NewQPageLayout() *QPageLayout {
	return QPageLayoutFromPointer(unsafe.Pointer(C.QPageLayout_NewQPageLayout()))
}

func NewQPageLayout3(other QPageLayoutITF) *QPageLayout {
	return QPageLayoutFromPointer(unsafe.Pointer(C.QPageLayout_NewQPageLayout3(C.QtObjectPtr(PointerFromQPageLayout(other)))))
}

func NewQPageLayout2(pageSize QPageSizeITF, orientation QPageLayout__Orientation, margins core.QMarginsFITF, units QPageLayout__Unit, minMargins core.QMarginsFITF) *QPageLayout {
	return QPageLayoutFromPointer(unsafe.Pointer(C.QPageLayout_NewQPageLayout2(C.QtObjectPtr(PointerFromQPageSize(pageSize)), C.int(orientation), C.QtObjectPtr(core.PointerFromQMarginsF(margins)), C.int(units), C.QtObjectPtr(core.PointerFromQMarginsF(minMargins)))))
}

func (ptr *QPageLayout) IsEquivalentTo(other QPageLayoutITF) bool {
	if ptr.Pointer() != nil {
		return C.QPageLayout_IsEquivalentTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageLayout(other))) != 0
	}
	return false
}

func (ptr *QPageLayout) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QPageLayout_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPageLayout) Mode() QPageLayout__Mode {
	if ptr.Pointer() != nil {
		return QPageLayout__Mode(C.QPageLayout_Mode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageLayout) Orientation() QPageLayout__Orientation {
	if ptr.Pointer() != nil {
		return QPageLayout__Orientation(C.QPageLayout_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageLayout) SetMargins(margins core.QMarginsFITF) bool {
	if ptr.Pointer() != nil {
		return C.QPageLayout_SetMargins(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMarginsF(margins))) != 0
	}
	return false
}

func (ptr *QPageLayout) SetMinimumMargins(minMargins core.QMarginsFITF) {
	if ptr.Pointer() != nil {
		C.QPageLayout_SetMinimumMargins(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMarginsF(minMargins)))
	}
}

func (ptr *QPageLayout) SetMode(mode QPageLayout__Mode) {
	if ptr.Pointer() != nil {
		C.QPageLayout_SetMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QPageLayout) SetOrientation(orientation QPageLayout__Orientation) {
	if ptr.Pointer() != nil {
		C.QPageLayout_SetOrientation(C.QtObjectPtr(ptr.Pointer()), C.int(orientation))
	}
}

func (ptr *QPageLayout) SetPageSize(pageSize QPageSizeITF, minMargins core.QMarginsFITF) {
	if ptr.Pointer() != nil {
		C.QPageLayout_SetPageSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageSize(pageSize)), C.QtObjectPtr(core.PointerFromQMarginsF(minMargins)))
	}
}

func (ptr *QPageLayout) SetUnits(units QPageLayout__Unit) {
	if ptr.Pointer() != nil {
		C.QPageLayout_SetUnits(C.QtObjectPtr(ptr.Pointer()), C.int(units))
	}
}

func (ptr *QPageLayout) Swap(other QPageLayoutITF) {
	if ptr.Pointer() != nil {
		C.QPageLayout_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageLayout(other)))
	}
}

func (ptr *QPageLayout) Units() QPageLayout__Unit {
	if ptr.Pointer() != nil {
		return QPageLayout__Unit(C.QPageLayout_Units(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageLayout) DestroyQPageLayout() {
	if ptr.Pointer() != nil {
		C.QPageLayout_DestroyQPageLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}
