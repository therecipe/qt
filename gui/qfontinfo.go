package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFontInfo struct {
	ptr unsafe.Pointer
}

type QFontInfo_ITF interface {
	QFontInfo_PTR() *QFontInfo
}

func (p *QFontInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontInfo(ptr QFontInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontInfo_PTR().Pointer()
	}
	return nil
}

func NewQFontInfoFromPointer(ptr unsafe.Pointer) *QFontInfo {
	var n = new(QFontInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontInfo) QFontInfo_PTR() *QFontInfo {
	return ptr
}

func NewQFontInfo(font QFont_ITF) *QFontInfo {
	defer qt.Recovering("QFontInfo::QFontInfo")

	return NewQFontInfoFromPointer(C.QFontInfo_NewQFontInfo(PointerFromQFont(font)))
}

func NewQFontInfo2(fi QFontInfo_ITF) *QFontInfo {
	defer qt.Recovering("QFontInfo::QFontInfo")

	return NewQFontInfoFromPointer(C.QFontInfo_NewQFontInfo2(PointerFromQFontInfo(fi)))
}

func (ptr *QFontInfo) Bold() bool {
	defer qt.Recovering("QFontInfo::bold")

	if ptr.Pointer() != nil {
		return C.QFontInfo_Bold(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) ExactMatch() bool {
	defer qt.Recovering("QFontInfo::exactMatch")

	if ptr.Pointer() != nil {
		return C.QFontInfo_ExactMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) Family() string {
	defer qt.Recovering("QFontInfo::family")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_Family(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFontInfo) FixedPitch() bool {
	defer qt.Recovering("QFontInfo::fixedPitch")

	if ptr.Pointer() != nil {
		return C.QFontInfo_FixedPitch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) Italic() bool {
	defer qt.Recovering("QFontInfo::italic")

	if ptr.Pointer() != nil {
		return C.QFontInfo_Italic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) PixelSize() int {
	defer qt.Recovering("QFontInfo::pixelSize")

	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PixelSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) PointSize() int {
	defer qt.Recovering("QFontInfo::pointSize")

	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) PointSizeF() float64 {
	defer qt.Recovering("QFontInfo::pointSizeF")

	if ptr.Pointer() != nil {
		return float64(C.QFontInfo_PointSizeF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) Style() QFont__Style {
	defer qt.Recovering("QFontInfo::style")

	if ptr.Pointer() != nil {
		return QFont__Style(C.QFontInfo_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) StyleName() string {
	defer qt.Recovering("QFontInfo::styleName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_StyleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFontInfo) Swap(other QFontInfo_ITF) {
	defer qt.Recovering("QFontInfo::swap")

	if ptr.Pointer() != nil {
		C.QFontInfo_Swap(ptr.Pointer(), PointerFromQFontInfo(other))
	}
}

func (ptr *QFontInfo) StyleHint() QFont__StyleHint {
	defer qt.Recovering("QFontInfo::styleHint")

	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QFontInfo_StyleHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) Weight() int {
	defer qt.Recovering("QFontInfo::weight")

	if ptr.Pointer() != nil {
		return int(C.QFontInfo_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) DestroyQFontInfo() {
	defer qt.Recovering("QFontInfo::~QFontInfo")

	if ptr.Pointer() != nil {
		C.QFontInfo_DestroyQFontInfo(ptr.Pointer())
	}
}
