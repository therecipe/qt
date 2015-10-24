package gui

//#include "qfontinfo.h"
import "C"
import (
	"unsafe"
)

type QFontInfo struct {
	ptr unsafe.Pointer
}

type QFontInfoITF interface {
	QFontInfoPTR() *QFontInfo
}

func (p *QFontInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFontInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFontInfo(ptr QFontInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontInfoPTR().Pointer()
	}
	return nil
}

func QFontInfoFromPointer(ptr unsafe.Pointer) *QFontInfo {
	var n = new(QFontInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFontInfo) QFontInfoPTR() *QFontInfo {
	return ptr
}

func NewQFontInfo(font QFontITF) *QFontInfo {
	return QFontInfoFromPointer(unsafe.Pointer(C.QFontInfo_NewQFontInfo(C.QtObjectPtr(PointerFromQFont(font)))))
}

func NewQFontInfo2(fi QFontInfoITF) *QFontInfo {
	return QFontInfoFromPointer(unsafe.Pointer(C.QFontInfo_NewQFontInfo2(C.QtObjectPtr(PointerFromQFontInfo(fi)))))
}

func (ptr *QFontInfo) Bold() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_Bold(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFontInfo) ExactMatch() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_ExactMatch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFontInfo) Family() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_Family(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFontInfo) FixedPitch() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_FixedPitch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFontInfo) Italic() bool {
	if ptr.Pointer() != nil {
		return C.QFontInfo_Italic(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFontInfo) PixelSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PixelSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontInfo) PointSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_PointSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontInfo) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QFontInfo_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontInfo) StyleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFontInfo_StyleName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFontInfo) Swap(other QFontInfoITF) {
	if ptr.Pointer() != nil {
		C.QFontInfo_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFontInfo(other)))
	}
}

func (ptr *QFontInfo) StyleHint() QFont__StyleHint {
	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QFontInfo_StyleHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontInfo) Weight() int {
	if ptr.Pointer() != nil {
		return int(C.QFontInfo_Weight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontInfo) DestroyQFontInfo() {
	if ptr.Pointer() != nil {
		C.QFontInfo_DestroyQFontInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
