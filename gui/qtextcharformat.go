package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QTextCharFormat struct {
	QTextFormat
}

type QTextCharFormat_ITF interface {
	QTextFormat_ITF
	QTextCharFormat_PTR() *QTextCharFormat
}

func PointerFromQTextCharFormat(ptr QTextCharFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextCharFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextCharFormatFromPointer(ptr unsafe.Pointer) *QTextCharFormat {
	var n = new(QTextCharFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextCharFormat) QTextCharFormat_PTR() *QTextCharFormat {
	return ptr
}

//QTextCharFormat::FontPropertiesInheritanceBehavior
type QTextCharFormat__FontPropertiesInheritanceBehavior int64

const (
	QTextCharFormat__FontPropertiesSpecifiedOnly = QTextCharFormat__FontPropertiesInheritanceBehavior(0)
	QTextCharFormat__FontPropertiesAll           = QTextCharFormat__FontPropertiesInheritanceBehavior(1)
)

//QTextCharFormat::UnderlineStyle
type QTextCharFormat__UnderlineStyle int64

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
type QTextCharFormat__VerticalAlignment int64

const (
	QTextCharFormat__AlignNormal      = QTextCharFormat__VerticalAlignment(0)
	QTextCharFormat__AlignSuperScript = QTextCharFormat__VerticalAlignment(1)
	QTextCharFormat__AlignSubScript   = QTextCharFormat__VerticalAlignment(2)
	QTextCharFormat__AlignMiddle      = QTextCharFormat__VerticalAlignment(3)
	QTextCharFormat__AlignTop         = QTextCharFormat__VerticalAlignment(4)
	QTextCharFormat__AlignBottom      = QTextCharFormat__VerticalAlignment(5)
	QTextCharFormat__AlignBaseline    = QTextCharFormat__VerticalAlignment(6)
)

func NewQTextCharFormat() *QTextCharFormat {
	defer qt.Recovering("QTextCharFormat::QTextCharFormat")

	return NewQTextCharFormatFromPointer(C.QTextCharFormat_NewQTextCharFormat())
}

func (ptr *QTextCharFormat) AnchorNames() []string {
	defer qt.Recovering("QTextCharFormat::anchorNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QTextCharFormat_AnchorNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QTextCharFormat) FontUnderline() bool {
	defer qt.Recovering("QTextCharFormat::fontUnderline")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontUnderline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) SetUnderlineStyle(style QTextCharFormat__UnderlineStyle) {
	defer qt.Recovering("QTextCharFormat::setUnderlineStyle")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetUnderlineStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextCharFormat) AnchorHref() string {
	defer qt.Recovering("QTextCharFormat::anchorHref")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_AnchorHref(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCharFormat) FontCapitalization() QFont__Capitalization {
	defer qt.Recovering("QTextCharFormat::fontCapitalization")

	if ptr.Pointer() != nil {
		return QFont__Capitalization(C.QTextCharFormat_FontCapitalization(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontFamily() string {
	defer qt.Recovering("QTextCharFormat::fontFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_FontFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCharFormat) FontFixedPitch() bool {
	defer qt.Recovering("QTextCharFormat::fontFixedPitch")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontFixedPitch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontHintingPreference() QFont__HintingPreference {
	defer qt.Recovering("QTextCharFormat::fontHintingPreference")

	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QTextCharFormat_FontHintingPreference(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontItalic() bool {
	defer qt.Recovering("QTextCharFormat::fontItalic")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontItalic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontKerning() bool {
	defer qt.Recovering("QTextCharFormat::fontKerning")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontKerning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontLetterSpacing() float64 {
	defer qt.Recovering("QTextCharFormat::fontLetterSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QTextCharFormat_FontLetterSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontLetterSpacingType() QFont__SpacingType {
	defer qt.Recovering("QTextCharFormat::fontLetterSpacingType")

	if ptr.Pointer() != nil {
		return QFont__SpacingType(C.QTextCharFormat_FontLetterSpacingType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontOverline() bool {
	defer qt.Recovering("QTextCharFormat::fontOverline")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontOverline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontPointSize() float64 {
	defer qt.Recovering("QTextCharFormat::fontPointSize")

	if ptr.Pointer() != nil {
		return float64(C.QTextCharFormat_FontPointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStretch() int {
	defer qt.Recovering("QTextCharFormat::fontStretch")

	if ptr.Pointer() != nil {
		return int(C.QTextCharFormat_FontStretch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStrikeOut() bool {
	defer qt.Recovering("QTextCharFormat::fontStrikeOut")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_FontStrikeOut(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) FontStyleHint() QFont__StyleHint {
	defer qt.Recovering("QTextCharFormat::fontStyleHint")

	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QTextCharFormat_FontStyleHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontStyleStrategy() QFont__StyleStrategy {
	defer qt.Recovering("QTextCharFormat::fontStyleStrategy")

	if ptr.Pointer() != nil {
		return QFont__StyleStrategy(C.QTextCharFormat_FontStyleStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontWeight() int {
	defer qt.Recovering("QTextCharFormat::fontWeight")

	if ptr.Pointer() != nil {
		return int(C.QTextCharFormat_FontWeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) FontWordSpacing() float64 {
	defer qt.Recovering("QTextCharFormat::fontWordSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QTextCharFormat_FontWordSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) IsAnchor() bool {
	defer qt.Recovering("QTextCharFormat::isAnchor")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_IsAnchor(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) IsValid() bool {
	defer qt.Recovering("QTextCharFormat::isValid")

	if ptr.Pointer() != nil {
		return C.QTextCharFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextCharFormat) SetAnchor(anchor bool) {
	defer qt.Recovering("QTextCharFormat::setAnchor")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchor(ptr.Pointer(), C.int(qt.GoBoolToInt(anchor)))
	}
}

func (ptr *QTextCharFormat) SetAnchorHref(value string) {
	defer qt.Recovering("QTextCharFormat::setAnchorHref")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchorHref(ptr.Pointer(), C.CString(value))
	}
}

func (ptr *QTextCharFormat) SetAnchorNames(names []string) {
	defer qt.Recovering("QTextCharFormat::setAnchorNames")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetAnchorNames(ptr.Pointer(), C.CString(strings.Join(names, ",,,")))
	}
}

func (ptr *QTextCharFormat) SetFont2(font QFont_ITF) {
	defer qt.Recovering("QTextCharFormat::setFont")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont2(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextCharFormat) SetFont(font QFont_ITF, behavior QTextCharFormat__FontPropertiesInheritanceBehavior) {
	defer qt.Recovering("QTextCharFormat::setFont")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFont(ptr.Pointer(), PointerFromQFont(font), C.int(behavior))
	}
}

func (ptr *QTextCharFormat) SetFontCapitalization(capitalization QFont__Capitalization) {
	defer qt.Recovering("QTextCharFormat::setFontCapitalization")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontCapitalization(ptr.Pointer(), C.int(capitalization))
	}
}

func (ptr *QTextCharFormat) SetFontFamily(family string) {
	defer qt.Recovering("QTextCharFormat::setFontFamily")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontFamily(ptr.Pointer(), C.CString(family))
	}
}

func (ptr *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	defer qt.Recovering("QTextCharFormat::setFontFixedPitch")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontFixedPitch(ptr.Pointer(), C.int(qt.GoBoolToInt(fixedPitch)))
	}
}

func (ptr *QTextCharFormat) SetFontHintingPreference(hintingPreference QFont__HintingPreference) {
	defer qt.Recovering("QTextCharFormat::setFontHintingPreference")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontHintingPreference(ptr.Pointer(), C.int(hintingPreference))
	}
}

func (ptr *QTextCharFormat) SetFontItalic(italic bool) {
	defer qt.Recovering("QTextCharFormat::setFontItalic")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontItalic(ptr.Pointer(), C.int(qt.GoBoolToInt(italic)))
	}
}

func (ptr *QTextCharFormat) SetFontKerning(enable bool) {
	defer qt.Recovering("QTextCharFormat::setFontKerning")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontKerning(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	defer qt.Recovering("QTextCharFormat::setFontLetterSpacing")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontLetterSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType QFont__SpacingType) {
	defer qt.Recovering("QTextCharFormat::setFontLetterSpacingType")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontLetterSpacingType(ptr.Pointer(), C.int(letterSpacingType))
	}
}

func (ptr *QTextCharFormat) SetFontOverline(overline bool) {
	defer qt.Recovering("QTextCharFormat::setFontOverline")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontOverline(ptr.Pointer(), C.int(qt.GoBoolToInt(overline)))
	}
}

func (ptr *QTextCharFormat) SetFontPointSize(size float64) {
	defer qt.Recovering("QTextCharFormat::setFontPointSize")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontPointSize(ptr.Pointer(), C.double(size))
	}
}

func (ptr *QTextCharFormat) SetFontStretch(factor int) {
	defer qt.Recovering("QTextCharFormat::setFontStretch")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStretch(ptr.Pointer(), C.int(factor))
	}
}

func (ptr *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	defer qt.Recovering("QTextCharFormat::setFontStrikeOut")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStrikeOut(ptr.Pointer(), C.int(qt.GoBoolToInt(strikeOut)))
	}
}

func (ptr *QTextCharFormat) SetFontStyleHint(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	defer qt.Recovering("QTextCharFormat::setFontStyleHint")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStyleHint(ptr.Pointer(), C.int(hint), C.int(strategy))
	}
}

func (ptr *QTextCharFormat) SetFontStyleStrategy(strategy QFont__StyleStrategy) {
	defer qt.Recovering("QTextCharFormat::setFontStyleStrategy")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontStyleStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QTextCharFormat) SetFontUnderline(underline bool) {
	defer qt.Recovering("QTextCharFormat::setFontUnderline")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QTextCharFormat) SetFontWeight(weight int) {
	defer qt.Recovering("QTextCharFormat::setFontWeight")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontWeight(ptr.Pointer(), C.int(weight))
	}
}

func (ptr *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	defer qt.Recovering("QTextCharFormat::setFontWordSpacing")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetFontWordSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QTextCharFormat) SetTextOutline(pen QPen_ITF) {
	defer qt.Recovering("QTextCharFormat::setTextOutline")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetTextOutline(ptr.Pointer(), PointerFromQPen(pen))
	}
}

func (ptr *QTextCharFormat) SetToolTip(text string) {
	defer qt.Recovering("QTextCharFormat::setToolTip")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetToolTip(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QTextCharFormat) SetUnderlineColor(color QColor_ITF) {
	defer qt.Recovering("QTextCharFormat::setUnderlineColor")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetUnderlineColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QTextCharFormat) SetVerticalAlignment(alignment QTextCharFormat__VerticalAlignment) {
	defer qt.Recovering("QTextCharFormat::setVerticalAlignment")

	if ptr.Pointer() != nil {
		C.QTextCharFormat_SetVerticalAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QTextCharFormat) ToolTip() string {
	defer qt.Recovering("QTextCharFormat::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextCharFormat_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextCharFormat) UnderlineColor() *QColor {
	defer qt.Recovering("QTextCharFormat::underlineColor")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QTextCharFormat_UnderlineColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextCharFormat) UnderlineStyle() QTextCharFormat__UnderlineStyle {
	defer qt.Recovering("QTextCharFormat::underlineStyle")

	if ptr.Pointer() != nil {
		return QTextCharFormat__UnderlineStyle(C.QTextCharFormat_UnderlineStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextCharFormat) VerticalAlignment() QTextCharFormat__VerticalAlignment {
	defer qt.Recovering("QTextCharFormat::verticalAlignment")

	if ptr.Pointer() != nil {
		return QTextCharFormat__VerticalAlignment(C.QTextCharFormat_VerticalAlignment(ptr.Pointer()))
	}
	return 0
}
