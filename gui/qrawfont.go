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

type QRawFont_ITF interface {
	QRawFont_PTR() *QRawFont
}

func (p *QRawFont) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRawFont) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRawFont(ptr QRawFont_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRawFont_PTR().Pointer()
	}
	return nil
}

func NewQRawFontFromPointer(ptr unsafe.Pointer) *QRawFont {
	var n = new(QRawFont)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRawFont) QRawFont_PTR() *QRawFont {
	return ptr
}

//QRawFont::AntialiasingType
type QRawFont__AntialiasingType int64

const (
	QRawFont__PixelAntialiasing    = QRawFont__AntialiasingType(0)
	QRawFont__SubPixelAntialiasing = QRawFont__AntialiasingType(1)
)

//QRawFont::LayoutFlag
type QRawFont__LayoutFlag int64

const (
	QRawFont__SeparateAdvances = QRawFont__LayoutFlag(0)
	QRawFont__KernedAdvances   = QRawFont__LayoutFlag(1)
	QRawFont__UseDesignMetrics = QRawFont__LayoutFlag(2)
)

func NewQRawFont() *QRawFont {
	return NewQRawFontFromPointer(C.QRawFont_NewQRawFont())
}

func NewQRawFont3(fontData core.QByteArray_ITF, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	return NewQRawFontFromPointer(C.QRawFont_NewQRawFont3(core.PointerFromQByteArray(fontData), C.double(pixelSize), C.int(hintingPreference)))
}

func NewQRawFont4(other QRawFont_ITF) *QRawFont {
	return NewQRawFontFromPointer(C.QRawFont_NewQRawFont4(PointerFromQRawFont(other)))
}

func NewQRawFont2(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	return NewQRawFontFromPointer(C.QRawFont_NewQRawFont2(C.CString(fileName), C.double(pixelSize), C.int(hintingPreference)))
}

func (ptr *QRawFont) Ascent() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) AverageCharWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_AverageCharWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) Descent() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) FamilyName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRawFont_FamilyName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRawFont) FontTable(tagName string) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QRawFont_FontTable(ptr.Pointer(), C.CString(tagName)))
	}
	return nil
}

func (ptr *QRawFont) HintingPreference() QFont__HintingPreference {
	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QRawFont_HintingPreference(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QRawFont_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRawFont) Leading() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_Leading(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) LineThickness() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_LineThickness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) LoadFromData(fontData core.QByteArray_ITF, pixelSize float64, hintingPreference QFont__HintingPreference) {
	if ptr.Pointer() != nil {
		C.QRawFont_LoadFromData(ptr.Pointer(), core.PointerFromQByteArray(fontData), C.double(pixelSize), C.int(hintingPreference))
	}
}

func (ptr *QRawFont) LoadFromFile(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) {
	if ptr.Pointer() != nil {
		C.QRawFont_LoadFromFile(ptr.Pointer(), C.CString(fileName), C.double(pixelSize), C.int(hintingPreference))
	}
}

func (ptr *QRawFont) MaxCharWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_MaxCharWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) PixelSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_PixelSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) SetPixelSize(pixelSize float64) {
	if ptr.Pointer() != nil {
		C.QRawFont_SetPixelSize(ptr.Pointer(), C.double(pixelSize))
	}
}

func (ptr *QRawFont) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QRawFont_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) StyleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRawFont_StyleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRawFont) SupportsCharacter(character core.QChar_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRawFont_SupportsCharacter(ptr.Pointer(), core.PointerFromQChar(character)) != 0
	}
	return false
}

func (ptr *QRawFont) Swap(other QRawFont_ITF) {
	if ptr.Pointer() != nil {
		C.QRawFont_Swap(ptr.Pointer(), PointerFromQRawFont(other))
	}
}

func (ptr *QRawFont) UnderlinePosition() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_UnderlinePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) UnitsPerEm() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_UnitsPerEm(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) Weight() int {
	if ptr.Pointer() != nil {
		return int(C.QRawFont_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) XHeight() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QRawFont_XHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRawFont) DestroyQRawFont() {
	if ptr.Pointer() != nil {
		C.QRawFont_DestroyQRawFont(ptr.Pointer())
	}
}
