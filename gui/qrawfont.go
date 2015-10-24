package gui

//#include "qrawfont.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRawFont struct {
	ptr unsafe.Pointer
}

type QRawFontITF interface {
	QRawFontPTR() *QRawFont
}

func (p *QRawFont) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRawFont) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRawFont(ptr QRawFontITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRawFontPTR().Pointer()
	}
	return nil
}

func QRawFontFromPointer(ptr unsafe.Pointer) *QRawFont {
	var n = new(QRawFont)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRawFont) QRawFontPTR() *QRawFont {
	return ptr
}

//QRawFont::AntialiasingType
type QRawFont__AntialiasingType int

var (
	QRawFont__PixelAntialiasing    = QRawFont__AntialiasingType(0)
	QRawFont__SubPixelAntialiasing = QRawFont__AntialiasingType(1)
)

//QRawFont::LayoutFlag
type QRawFont__LayoutFlag int

var (
	QRawFont__SeparateAdvances = QRawFont__LayoutFlag(0)
	QRawFont__KernedAdvances   = QRawFont__LayoutFlag(1)
	QRawFont__UseDesignMetrics = QRawFont__LayoutFlag(2)
)

func NewQRawFont() *QRawFont {
	return QRawFontFromPointer(unsafe.Pointer(C.QRawFont_NewQRawFont()))
}

func NewQRawFont4(other QRawFontITF) *QRawFont {
	return QRawFontFromPointer(unsafe.Pointer(C.QRawFont_NewQRawFont4(C.QtObjectPtr(PointerFromQRawFont(other)))))
}

func (ptr *QRawFont) FamilyName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRawFont_FamilyName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRawFont) HintingPreference() QFont__HintingPreference {
	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QRawFont_HintingPreference(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRawFont) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QRawFont_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRawFont) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QRawFont_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRawFont) StyleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRawFont_StyleName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRawFont) SupportsCharacter(character core.QCharITF) bool {
	if ptr.Pointer() != nil {
		return C.QRawFont_SupportsCharacter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQChar(character))) != 0
	}
	return false
}

func (ptr *QRawFont) Swap(other QRawFontITF) {
	if ptr.Pointer() != nil {
		C.QRawFont_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRawFont(other)))
	}
}

func (ptr *QRawFont) Weight() int {
	if ptr.Pointer() != nil {
		return int(C.QRawFont_Weight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRawFont) DestroyQRawFont() {
	if ptr.Pointer() != nil {
		C.QRawFont_DestroyQRawFont(C.QtObjectPtr(ptr.Pointer()))
	}
}
