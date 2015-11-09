package gui

//#include "qfontinfo.h"
import "C"
import (
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
	return NewQFontInfoFromPointer(C.QFontInfo_NewQFontInfo(PointerFromQFont(font)))
}

func NewQFontInfo2(fi QFontInfo_ITF) *QFontInfo {
	return NewQFontInfoFromPointer(C.QFontInfo_NewQFontInfo2(PointerFromQFontInfo(fi)))
}

func (ptr *QFontInfo) Bold() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_Bold(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) ExactMatch() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_ExactMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) Family() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_Family(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFontInfo) FixedPitch() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_FixedPitch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) Italic() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_Italic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFontInfo) PixelSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PixelSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) PointSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) PointSizeF() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QFontInfo_PointSizeF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QFontInfo_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) StyleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_StyleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFontInfo) Swap(other QFontInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QFontInfo_Swap(ptr.Pointer(), PointerFromQFontInfo(other))
	}
}

func (ptr *QFontInfo) StyleHint() QFont__StyleHint {
	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QFontInfo_StyleHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) Weight() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontInfo) DestroyQFontInfo() {
	if ptr.Pointer() != nil {
		C.QFontInfo_DestroyQFontInfo(ptr.Pointer())
	}
}
