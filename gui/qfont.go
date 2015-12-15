package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFont struct {
	ptr unsafe.Pointer
}

type QFont_ITF interface {
	QFont_PTR() *QFont
}

func (p *QFont) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFont) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFont(ptr QFont_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFont_PTR().Pointer()
	}
	return nil
}

func NewQFontFromPointer(ptr unsafe.Pointer) *QFont {
	var n = new(QFont)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFont) QFont_PTR() *QFont {
	return ptr
}

//QFont::Capitalization
type QFont__Capitalization int64

const (
	QFont__MixedCase    = QFont__Capitalization(0)
	QFont__AllUppercase = QFont__Capitalization(1)
	QFont__AllLowercase = QFont__Capitalization(2)
	QFont__SmallCaps    = QFont__Capitalization(3)
	QFont__Capitalize   = QFont__Capitalization(4)
)

//QFont::HintingPreference
type QFont__HintingPreference int64

const (
	QFont__PreferDefaultHinting  = QFont__HintingPreference(0)
	QFont__PreferNoHinting       = QFont__HintingPreference(1)
	QFont__PreferVerticalHinting = QFont__HintingPreference(2)
	QFont__PreferFullHinting     = QFont__HintingPreference(3)
)

//QFont::SpacingType
type QFont__SpacingType int64

const (
	QFont__PercentageSpacing = QFont__SpacingType(0)
	QFont__AbsoluteSpacing   = QFont__SpacingType(1)
)

//QFont::Stretch
type QFont__Stretch int64

const (
	QFont__UltraCondensed = QFont__Stretch(50)
	QFont__ExtraCondensed = QFont__Stretch(62)
	QFont__Condensed      = QFont__Stretch(75)
	QFont__SemiCondensed  = QFont__Stretch(87)
	QFont__Unstretched    = QFont__Stretch(100)
	QFont__SemiExpanded   = QFont__Stretch(112)
	QFont__Expanded       = QFont__Stretch(125)
	QFont__ExtraExpanded  = QFont__Stretch(150)
	QFont__UltraExpanded  = QFont__Stretch(200)
)

//QFont::Style
type QFont__Style int64

var (
	QFont__StyleNormal  = QFont__Style(0)
	QFont__StyleItalic  = QFont__Style(1)
	QFont__StyleOblique = QFont__Style(2)
)

//QFont::StyleHint
type QFont__StyleHint int64

var (
	QFont__Helvetica  = QFont__StyleHint(0)
	QFont__SansSerif  = QFont__StyleHint(QFont__Helvetica)
	QFont__Times      = QFont__StyleHint(C.QFont_Times_Type())
	QFont__Serif      = QFont__StyleHint(QFont__Times)
	QFont__Courier    = QFont__StyleHint(C.QFont_Courier_Type())
	QFont__TypeWriter = QFont__StyleHint(QFont__Courier)
	QFont__OldEnglish = QFont__StyleHint(C.QFont_OldEnglish_Type())
	QFont__Decorative = QFont__StyleHint(QFont__OldEnglish)
	QFont__System     = QFont__StyleHint(C.QFont_System_Type())
	QFont__AnyStyle   = QFont__StyleHint(C.QFont_AnyStyle_Type())
	QFont__Cursive    = QFont__StyleHint(C.QFont_Cursive_Type())
	QFont__Monospace  = QFont__StyleHint(C.QFont_Monospace_Type())
	QFont__Fantasy    = QFont__StyleHint(C.QFont_Fantasy_Type())
)

//QFont::StyleStrategy
type QFont__StyleStrategy int64

var (
	QFont__PreferDefault       = QFont__StyleStrategy(0x0001)
	QFont__PreferBitmap        = QFont__StyleStrategy(0x0002)
	QFont__PreferDevice        = QFont__StyleStrategy(0x0004)
	QFont__PreferOutline       = QFont__StyleStrategy(0x0008)
	QFont__ForceOutline        = QFont__StyleStrategy(0x0010)
	QFont__PreferMatch         = QFont__StyleStrategy(0x0020)
	QFont__PreferQuality       = QFont__StyleStrategy(0x0040)
	QFont__PreferAntialias     = QFont__StyleStrategy(0x0080)
	QFont__NoAntialias         = QFont__StyleStrategy(0x0100)
	QFont__OpenGLCompatible    = QFont__StyleStrategy(0x0200)
	QFont__ForceIntegerMetrics = QFont__StyleStrategy(0x0400)
	QFont__NoSubpixelAntialias = QFont__StyleStrategy(0x0800)
	QFont__NoFontMerging       = QFont__StyleStrategy(0x8000)
)

//QFont::Weight
type QFont__Weight int64

const (
	QFont__Thin       = QFont__Weight(0)
	QFont__ExtraLight = QFont__Weight(12)
	QFont__Light      = QFont__Weight(25)
	QFont__Normal     = QFont__Weight(50)
	QFont__Medium     = QFont__Weight(57)
	QFont__DemiBold   = QFont__Weight(63)
	QFont__Bold       = QFont__Weight(75)
	QFont__ExtraBold  = QFont__Weight(81)
	QFont__Black      = QFont__Weight(87)
)

func (ptr *QFont) DefaultFamily() string {
	defer qt.Recovering("QFont::defaultFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_DefaultFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) LastResortFamily() string {
	defer qt.Recovering("QFont::lastResortFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_LastResortFamily(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) LastResortFont() string {
	defer qt.Recovering("QFont::lastResortFont")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_LastResortFont(ptr.Pointer()))
	}
	return ""
}

func NewQFont() *QFont {
	defer qt.Recovering("QFont::QFont")

	return NewQFontFromPointer(C.QFont_NewQFont())
}

func NewQFont4(font QFont_ITF) *QFont {
	defer qt.Recovering("QFont::QFont")

	return NewQFontFromPointer(C.QFont_NewQFont4(PointerFromQFont(font)))
}

func NewQFont3(font QFont_ITF, pd QPaintDevice_ITF) *QFont {
	defer qt.Recovering("QFont::QFont")

	return NewQFontFromPointer(C.QFont_NewQFont3(PointerFromQFont(font), PointerFromQPaintDevice(pd)))
}

func NewQFont2(family string, pointSize int, weight int, italic bool) *QFont {
	defer qt.Recovering("QFont::QFont")

	return NewQFontFromPointer(C.QFont_NewQFont2(C.CString(family), C.int(pointSize), C.int(weight), C.int(qt.GoBoolToInt(italic))))
}

func (ptr *QFont) Bold() bool {
	defer qt.Recovering("QFont::bold")

	if ptr.Pointer() != nil {
		return C.QFont_Bold(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Capitalization() QFont__Capitalization {
	defer qt.Recovering("QFont::capitalization")

	if ptr.Pointer() != nil {
		return QFont__Capitalization(C.QFont_Capitalization(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) ExactMatch() bool {
	defer qt.Recovering("QFont::exactMatch")

	if ptr.Pointer() != nil {
		return C.QFont_ExactMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Family() string {
	defer qt.Recovering("QFont::family")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_Family(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) FixedPitch() bool {
	defer qt.Recovering("QFont::fixedPitch")

	if ptr.Pointer() != nil {
		return C.QFont_FixedPitch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) FromString(descrip string) bool {
	defer qt.Recovering("QFont::fromString")

	if ptr.Pointer() != nil {
		return C.QFont_FromString(ptr.Pointer(), C.CString(descrip)) != 0
	}
	return false
}

func (ptr *QFont) HintingPreference() QFont__HintingPreference {
	defer qt.Recovering("QFont::hintingPreference")

	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QFont_HintingPreference(ptr.Pointer()))
	}
	return 0
}

func QFont_InsertSubstitution(familyName string, substituteName string) {
	defer qt.Recovering("QFont::insertSubstitution")

	C.QFont_QFont_InsertSubstitution(C.CString(familyName), C.CString(substituteName))
}

func QFont_InsertSubstitutions(familyName string, substituteNames []string) {
	defer qt.Recovering("QFont::insertSubstitutions")

	C.QFont_QFont_InsertSubstitutions(C.CString(familyName), C.CString(strings.Join(substituteNames, ",,,")))
}

func (ptr *QFont) IsCopyOf(f QFont_ITF) bool {
	defer qt.Recovering("QFont::isCopyOf")

	if ptr.Pointer() != nil {
		return C.QFont_IsCopyOf(ptr.Pointer(), PointerFromQFont(f)) != 0
	}
	return false
}

func (ptr *QFont) Italic() bool {
	defer qt.Recovering("QFont::italic")

	if ptr.Pointer() != nil {
		return C.QFont_Italic(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Kerning() bool {
	defer qt.Recovering("QFont::kerning")

	if ptr.Pointer() != nil {
		return C.QFont_Kerning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Key() string {
	defer qt.Recovering("QFont::key")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) LetterSpacing() float64 {
	defer qt.Recovering("QFont::letterSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QFont_LetterSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) LetterSpacingType() QFont__SpacingType {
	defer qt.Recovering("QFont::letterSpacingType")

	if ptr.Pointer() != nil {
		return QFont__SpacingType(C.QFont_LetterSpacingType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) Overline() bool {
	defer qt.Recovering("QFont::overline")

	if ptr.Pointer() != nil {
		return C.QFont_Overline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) PixelSize() int {
	defer qt.Recovering("QFont::pixelSize")

	if ptr.Pointer() != nil {
		return int(C.QFont_PixelSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) PointSize() int {
	defer qt.Recovering("QFont::pointSize")

	if ptr.Pointer() != nil {
		return int(C.QFont_PointSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) PointSizeF() float64 {
	defer qt.Recovering("QFont::pointSizeF")

	if ptr.Pointer() != nil {
		return float64(C.QFont_PointSizeF(ptr.Pointer()))
	}
	return 0
}

func QFont_RemoveSubstitutions(familyName string) {
	defer qt.Recovering("QFont::removeSubstitutions")

	C.QFont_QFont_RemoveSubstitutions(C.CString(familyName))
}

func (ptr *QFont) SetBold(enable bool) {
	defer qt.Recovering("QFont::setBold")

	if ptr.Pointer() != nil {
		C.QFont_SetBold(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetCapitalization(caps QFont__Capitalization) {
	defer qt.Recovering("QFont::setCapitalization")

	if ptr.Pointer() != nil {
		C.QFont_SetCapitalization(ptr.Pointer(), C.int(caps))
	}
}

func (ptr *QFont) SetFamily(family string) {
	defer qt.Recovering("QFont::setFamily")

	if ptr.Pointer() != nil {
		C.QFont_SetFamily(ptr.Pointer(), C.CString(family))
	}
}

func (ptr *QFont) SetFixedPitch(enable bool) {
	defer qt.Recovering("QFont::setFixedPitch")

	if ptr.Pointer() != nil {
		C.QFont_SetFixedPitch(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetHintingPreference(hintingPreference QFont__HintingPreference) {
	defer qt.Recovering("QFont::setHintingPreference")

	if ptr.Pointer() != nil {
		C.QFont_SetHintingPreference(ptr.Pointer(), C.int(hintingPreference))
	}
}

func (ptr *QFont) SetItalic(enable bool) {
	defer qt.Recovering("QFont::setItalic")

	if ptr.Pointer() != nil {
		C.QFont_SetItalic(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetKerning(enable bool) {
	defer qt.Recovering("QFont::setKerning")

	if ptr.Pointer() != nil {
		C.QFont_SetKerning(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetLetterSpacing(ty QFont__SpacingType, spacing float64) {
	defer qt.Recovering("QFont::setLetterSpacing")

	if ptr.Pointer() != nil {
		C.QFont_SetLetterSpacing(ptr.Pointer(), C.int(ty), C.double(spacing))
	}
}

func (ptr *QFont) SetOverline(enable bool) {
	defer qt.Recovering("QFont::setOverline")

	if ptr.Pointer() != nil {
		C.QFont_SetOverline(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetPixelSize(pixelSize int) {
	defer qt.Recovering("QFont::setPixelSize")

	if ptr.Pointer() != nil {
		C.QFont_SetPixelSize(ptr.Pointer(), C.int(pixelSize))
	}
}

func (ptr *QFont) SetPointSize(pointSize int) {
	defer qt.Recovering("QFont::setPointSize")

	if ptr.Pointer() != nil {
		C.QFont_SetPointSize(ptr.Pointer(), C.int(pointSize))
	}
}

func (ptr *QFont) SetPointSizeF(pointSize float64) {
	defer qt.Recovering("QFont::setPointSizeF")

	if ptr.Pointer() != nil {
		C.QFont_SetPointSizeF(ptr.Pointer(), C.double(pointSize))
	}
}

func (ptr *QFont) SetStretch(factor int) {
	defer qt.Recovering("QFont::setStretch")

	if ptr.Pointer() != nil {
		C.QFont_SetStretch(ptr.Pointer(), C.int(factor))
	}
}

func (ptr *QFont) SetStrikeOut(enable bool) {
	defer qt.Recovering("QFont::setStrikeOut")

	if ptr.Pointer() != nil {
		C.QFont_SetStrikeOut(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetStyle(style QFont__Style) {
	defer qt.Recovering("QFont::setStyle")

	if ptr.Pointer() != nil {
		C.QFont_SetStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QFont) SetStyleHint(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	defer qt.Recovering("QFont::setStyleHint")

	if ptr.Pointer() != nil {
		C.QFont_SetStyleHint(ptr.Pointer(), C.int(hint), C.int(strategy))
	}
}

func (ptr *QFont) SetStyleName(styleName string) {
	defer qt.Recovering("QFont::setStyleName")

	if ptr.Pointer() != nil {
		C.QFont_SetStyleName(ptr.Pointer(), C.CString(styleName))
	}
}

func (ptr *QFont) SetStyleStrategy(s QFont__StyleStrategy) {
	defer qt.Recovering("QFont::setStyleStrategy")

	if ptr.Pointer() != nil {
		C.QFont_SetStyleStrategy(ptr.Pointer(), C.int(s))
	}
}

func (ptr *QFont) SetUnderline(enable bool) {
	defer qt.Recovering("QFont::setUnderline")

	if ptr.Pointer() != nil {
		C.QFont_SetUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetWeight(weight int) {
	defer qt.Recovering("QFont::setWeight")

	if ptr.Pointer() != nil {
		C.QFont_SetWeight(ptr.Pointer(), C.int(weight))
	}
}

func (ptr *QFont) SetWordSpacing(spacing float64) {
	defer qt.Recovering("QFont::setWordSpacing")

	if ptr.Pointer() != nil {
		C.QFont_SetWordSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QFont) Stretch() int {
	defer qt.Recovering("QFont::stretch")

	if ptr.Pointer() != nil {
		return int(C.QFont_Stretch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) StrikeOut() bool {
	defer qt.Recovering("QFont::strikeOut")

	if ptr.Pointer() != nil {
		return C.QFont_StrikeOut(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Style() QFont__Style {
	defer qt.Recovering("QFont::style")

	if ptr.Pointer() != nil {
		return QFont__Style(C.QFont_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) StyleHint() QFont__StyleHint {
	defer qt.Recovering("QFont::styleHint")

	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QFont_StyleHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) StyleName() string {
	defer qt.Recovering("QFont::styleName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_StyleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) StyleStrategy() QFont__StyleStrategy {
	defer qt.Recovering("QFont::styleStrategy")

	if ptr.Pointer() != nil {
		return QFont__StyleStrategy(C.QFont_StyleStrategy(ptr.Pointer()))
	}
	return 0
}

func QFont_Substitute(familyName string) string {
	defer qt.Recovering("QFont::substitute")

	return C.GoString(C.QFont_QFont_Substitute(C.CString(familyName)))
}

func QFont_Substitutes(familyName string) []string {
	defer qt.Recovering("QFont::substitutes")

	return strings.Split(C.GoString(C.QFont_QFont_Substitutes(C.CString(familyName))), ",,,")
}

func QFont_Substitutions() []string {
	defer qt.Recovering("QFont::substitutions")

	return strings.Split(C.GoString(C.QFont_QFont_Substitutions()), ",,,")
}

func (ptr *QFont) Swap(other QFont_ITF) {
	defer qt.Recovering("QFont::swap")

	if ptr.Pointer() != nil {
		C.QFont_Swap(ptr.Pointer(), PointerFromQFont(other))
	}
}

func (ptr *QFont) ToString() string {
	defer qt.Recovering("QFont::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFont) Underline() bool {
	defer qt.Recovering("QFont::underline")

	if ptr.Pointer() != nil {
		return C.QFont_Underline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFont) Weight() int {
	defer qt.Recovering("QFont::weight")

	if ptr.Pointer() != nil {
		return int(C.QFont_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) WordSpacing() float64 {
	defer qt.Recovering("QFont::wordSpacing")

	if ptr.Pointer() != nil {
		return float64(C.QFont_WordSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFont) DestroyQFont() {
	defer qt.Recovering("QFont::~QFont")

	if ptr.Pointer() != nil {
		C.QFont_DestroyQFont(ptr.Pointer())
	}
}
