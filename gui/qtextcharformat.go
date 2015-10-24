package gui

//#include "qtextcharformat.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QTextCharFormat struct {
	QTextFormat
}

type QTextCharFormatITF interface {
	QTextFormatITF
	QTextCharFormatPTR() *QTextCharFormat
}

func PointerFromQTextCharFormat(ptr QTextCharFormatITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCharFormatPTR().Pointer()
	}
	return nil
}

func QTextCharFormatFromPointer(ptr unsafe.Pointer) *QTextCharFormat {
	var n = new(QTextCharFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextCharFormat) QTextCharFormatPTR() *QTextCharFormat {
	return ptr
}

//QTextCharFormat::FontPropertiesInheritanceBehavior
type QTextCharFormat__FontPropertiesInheritanceBehavior int

var (
	QTextCharFormat__FontPropertiesSpecifiedOnly = QTextCharFormat__FontPropertiesInheritanceBehavior(0)
	QTextCharFormat__FontPropertiesAll           = QTextCharFormat__FontPropertiesInheritanceBehavior(1)
)

//QTextCharFormat::UnderlineStyle
type QTextCharFormat__UnderlineStyle int

var (
	QTextCharFormat__NoUnderline         = QTextCharFormat__UnderlineStyle(0)
	QTextCharFormat__SingleUnderline     = QTextCharFormat__UnderlineStyle(1)
	QTextCharFormat__DashUnderline       = QTextCharFormat__UnderlineStyle(2)
	QTextCharFormat__DotLine             = QTextCharFormat__UnderlineStyle(3)
	QTextCharFormat__DashDotLine         = QTextCharFormat__UnderlineStyle(4)
	QTextCharFormat__DashDotDotLine      = QTextCharFormat__UnderlineStyle(5)
	QTextCharFormat__WaveUnderline       = QTextCharFormat__UnderlineStyle(6)
	QTextCharFormat__SpellCheckUnderline = QTextCharFormat__UnderlineStyle(7)
)

//QTextCharFormat::VerticalAlignment
type QTextCharFormat__VerticalAlignment int

var (
	QTextCharFormat__AlignNormal      = QTextCharFormat__VerticalAlignment(0)
	QTextCharFormat__AlignSuperScript = QTextCharFormat__VerticalAlignment(1)
	QTextCharFormat__AlignSubScript   = QTextCharFormat__VerticalAlignment(2)
	QTextCharFormat__AlignMiddle      = QTextCharFormat__VerticalAlignment(3)
	QTextCharFormat__AlignTop         = QTextCharFormat__VerticalAlignment(4)
	QTextCharFormat__AlignBottom      = QTextCharFormat__VerticalAlignment(5)
	QTextCharFormat__AlignBaseline    = QTextCharFormat__VerticalAlignment(6)
)

func NewQTextCharFormat() *QTextCharFormat {
	return QTextCharFormatFromPointer(unsafe.Pointer(C.QTextCharFormat_NewQTextCharFormat()))
}

func (ptr *QTextCharFormat) AnchorNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QTextCharFormat_AnchorNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QTextCharFormat) FontUnderline() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontUnderline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) SetUnderlineStyle(style QTextCharFormat__UnderlineStyle) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetUnderlineStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTextCharFormat) AnchorHref() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_AnchorHref(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextCharFormat) FontCapitalization() QFont__Capitalization {
	if ptr.Pointer() != nil {
		return QFont__Capitalization(C.QTextCharFormat_FontCapitalization(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontFamily() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_FontFamily(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextCharFormat) FontFixedPitch() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontFixedPitch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontHintingPreference() QFont__HintingPreference {
	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QTextCharFormat_FontHintingPreference(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontItalic() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontItalic(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontKerning() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontKerning(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontLetterSpacingType() QFont__SpacingType {
	if ptr.Pointer() != nil {
		return QFont__SpacingType(C.QTextCharFormat_FontLetterSpacingType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontOverline() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontOverline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontStretch() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCharFormat_FontStretch(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStrikeOut() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontStrikeOut(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontStyleHint() QFont__StyleHint {
	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QTextCharFormat_FontStyleHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStyleStrategy() QFont__StyleStrategy {
	if ptr.Pointer() != nil {
		return QFont__StyleStrategy(C.QTextCharFormat_FontStyleStrategy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) FontWeight() int {
	if ptr.Pointer() != nil {
		return int(C.QTextCharFormat_FontWeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) IsAnchor() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_IsAnchor(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextCharFormat_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextCharFormat) SetAnchor(anchor bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchor(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(anchor)))
	}
}

func (ptr *QTextCharFormat) SetAnchorHref(value string) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchorHref(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QTextCharFormat) SetAnchorNames(names []string) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchorNames(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(names, "|")))
	}
}

func (ptr *QTextCharFormat) SetFont2(font QFontITF) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QTextCharFormat) SetFont(font QFontITF, behavior QTextCharFormat__FontPropertiesInheritanceBehavior) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)), C.int(behavior))
	}
}

func (ptr *QTextCharFormat) SetFontCapitalization(capitalization QFont__Capitalization) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontCapitalization(C.QtObjectPtr(ptr.Pointer()), C.int(capitalization))
	}
}

func (ptr *QTextCharFormat) SetFontFamily(family string) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontFamily(C.QtObjectPtr(ptr.Pointer()), C.CString(family))
	}
}

func (ptr *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontFixedPitch(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(fixedPitch)))
	}
}

func (ptr *QTextCharFormat) SetFontHintingPreference(hintingPreference QFont__HintingPreference) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontHintingPreference(C.QtObjectPtr(ptr.Pointer()), C.int(hintingPreference))
	}
}

func (ptr *QTextCharFormat) SetFontItalic(italic bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontItalic(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextCharFormat) SetFontKerning(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontKerning(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType QFont__SpacingType) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontLetterSpacingType(C.QtObjectPtr(ptr.Pointer()), C.int(letterSpacingType))
	}
}

func (ptr *QTextCharFormat) SetFontOverline(overline bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontOverline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(overline)))
	}
}

func (ptr *QTextCharFormat) SetFontStretch(factor int) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStretch(C.QtObjectPtr(ptr.Pointer()), C.int(factor))
	}
}

func (ptr *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStrikeOut(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(strikeOut)))
	}
}

func (ptr *QTextCharFormat) SetFontStyleHint(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStyleHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.int(strategy))
	}
}

func (ptr *QTextCharFormat) SetFontStyleStrategy(strategy QFont__StyleStrategy) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStyleStrategy(C.QtObjectPtr(ptr.Pointer()), C.int(strategy))
	}
}

func (ptr *QTextCharFormat) SetFontUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontUnderline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextCharFormat) SetFontWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontWeight(C.QtObjectPtr(ptr.Pointer()), C.int(weight))
	}
}

func (ptr *QTextCharFormat) SetTextOutline(pen QPenITF) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetTextOutline(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPen(pen)))
	}
}

func (ptr *QTextCharFormat) SetToolTip(text string) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QTextCharFormat) SetUnderlineColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetUnderlineColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QTextCharFormat) SetVerticalAlignment(alignment QTextCharFormat__VerticalAlignment) {
	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetVerticalAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QTextCharFormat) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextCharFormat) UnderlineStyle() QTextCharFormat__UnderlineStyle {
	if ptr.Pointer() != nil {
		return QTextCharFormat__UnderlineStyle(C.QTextCharFormat_UnderlineStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextCharFormat) VerticalAlignment() QTextCharFormat__VerticalAlignment {
	if ptr.Pointer() != nil {
		return QTextCharFormat__VerticalAlignment(C.QTextCharFormat_VerticalAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
