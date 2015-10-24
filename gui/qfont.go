package gui

//#include "qfont.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFont struct {
	ptr unsafe.Pointer
}

type QFontITF interface {
	QFontPTR() *QFont
}

func (p *QFont) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFont) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFont(ptr QFontITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontPTR().Pointer()
	}
	return nil
}

func QFontFromPointer(ptr unsafe.Pointer) *QFont {
	var n = new(QFont)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFont) QFontPTR() *QFont {
	return ptr
}

//QFont::Capitalization
type QFont__Capitalization int

var (
	QFont__MixedCase    = QFont__Capitalization(0)
	QFont__AllUppercase = QFont__Capitalization(1)
	QFont__AllLowercase = QFont__Capitalization(2)
	QFont__SmallCaps    = QFont__Capitalization(3)
	QFont__Capitalize   = QFont__Capitalization(4)
)

//QFont::HintingPreference
type QFont__HintingPreference int

var (
	QFont__PreferDefaultHinting  = QFont__HintingPreference(0)
	QFont__PreferNoHinting       = QFont__HintingPreference(1)
	QFont__PreferVerticalHinting = QFont__HintingPreference(2)
	QFont__PreferFullHinting     = QFont__HintingPreference(3)
)

//QFont::SpacingType
type QFont__SpacingType int

var (
	QFont__PercentageSpacing = QFont__SpacingType(0)
	QFont__AbsoluteSpacing   = QFont__SpacingType(1)
)

//QFont::Stretch
type QFont__Stretch int

var (
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
type QFont__Style int

var (
	QFont__StyleNormal  = QFont__Style(0)
	QFont__StyleItalic  = QFont__Style(1)
	QFont__StyleOblique = QFont__Style(2)
)

//QFont::StyleHint
type QFont__StyleHint int

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
type QFont__StyleStrategy int

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
type QFont__Weight int

var (
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
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_DefaultFamily(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) LastResortFamily() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_LastResortFamily(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) LastResortFont() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_LastResortFont(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQFont() *QFont {
	return QFontFromPointer(unsafe.Pointer(C.QFont_NewQFont()))
}

func NewQFont4(font QFontITF) *QFont {
	return QFontFromPointer(unsafe.Pointer(C.QFont_NewQFont4(C.QtObjectPtr(PointerFromQFont(font)))))
}

func NewQFont3(font QFontITF, pd QPaintDeviceITF) *QFont {
	return QFontFromPointer(unsafe.Pointer(C.QFont_NewQFont3(C.QtObjectPtr(PointerFromQFont(font)), C.QtObjectPtr(PointerFromQPaintDevice(pd)))))
}

func NewQFont2(family string, pointSize int, weight int, italic bool) *QFont {
	return QFontFromPointer(unsafe.Pointer(C.QFont_NewQFont2(C.CString(family), C.int(pointSize), C.int(weight), C.int(qt.GoBoolToInt(italic)))))
}

func (ptr *QFont) Bold() bool {
	if ptr.Pointer() != nil {
		return C.QFont_Bold(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Capitalization() QFont__Capitalization {
	if ptr.Pointer() != nil {
		return QFont__Capitalization(C.QFont_Capitalization(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) ExactMatch() bool {
	if ptr.Pointer() != nil {
		return C.QFont_ExactMatch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Family() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_Family(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) FixedPitch() bool {
	if ptr.Pointer() != nil {
		return C.QFont_FixedPitch(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) FromString(descrip string) bool {
	if ptr.Pointer() != nil {
		return C.QFont_FromString(C.QtObjectPtr(ptr.Pointer()), C.CString(descrip)) != 0
	}
	return false
}

func (ptr *QFont) HintingPreference() QFont__HintingPreference {
	if ptr.Pointer() != nil {
		return QFont__HintingPreference(C.QFont_HintingPreference(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QFont_InsertSubstitution(familyName string, substituteName string) {
	C.QFont_QFont_InsertSubstitution(C.CString(familyName), C.CString(substituteName))
}

func QFont_InsertSubstitutions(familyName string, substituteNames []string) {
	C.QFont_QFont_InsertSubstitutions(C.CString(familyName), C.CString(strings.Join(substituteNames, "|")))
}

func (ptr *QFont) IsCopyOf(f QFontITF) bool {
	if ptr.Pointer() != nil {
		return C.QFont_IsCopyOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(f))) != 0
	}
	return false
}

func (ptr *QFont) Italic() bool {
	if ptr.Pointer() != nil {
		return C.QFont_Italic(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Kerning() bool {
	if ptr.Pointer() != nil {
		return C.QFont_Kerning(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Key() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) LetterSpacingType() QFont__SpacingType {
	if ptr.Pointer() != nil {
		return QFont__SpacingType(C.QFont_LetterSpacingType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) Overline() bool {
	if ptr.Pointer() != nil {
		return C.QFont_Overline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) PixelSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFont_PixelSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) PointSize() int {
	if ptr.Pointer() != nil {
		return int(C.QFont_PointSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QFont_RemoveSubstitutions(familyName string) {
	C.QFont_QFont_RemoveSubstitutions(C.CString(familyName))
}

func (ptr *QFont) SetBold(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetBold(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetCapitalization(caps QFont__Capitalization) {
	if ptr.Pointer() != nil {
		C.QFont_SetCapitalization(C.QtObjectPtr(ptr.Pointer()), C.int(caps))
	}
}

func (ptr *QFont) SetFamily(family string) {
	if ptr.Pointer() != nil {
		C.QFont_SetFamily(C.QtObjectPtr(ptr.Pointer()), C.CString(family))
	}
}

func (ptr *QFont) SetFixedPitch(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetFixedPitch(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetHintingPreference(hintingPreference QFont__HintingPreference) {
	if ptr.Pointer() != nil {
		C.QFont_SetHintingPreference(C.QtObjectPtr(ptr.Pointer()), C.int(hintingPreference))
	}
}

func (ptr *QFont) SetItalic(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetItalic(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetKerning(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetKerning(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetOverline(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetOverline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetPixelSize(pixelSize int) {
	if ptr.Pointer() != nil {
		C.QFont_SetPixelSize(C.QtObjectPtr(ptr.Pointer()), C.int(pixelSize))
	}
}

func (ptr *QFont) SetPointSize(pointSize int) {
	if ptr.Pointer() != nil {
		C.QFont_SetPointSize(C.QtObjectPtr(ptr.Pointer()), C.int(pointSize))
	}
}

func (ptr *QFont) SetStretch(factor int) {
	if ptr.Pointer() != nil {
		C.QFont_SetStretch(C.QtObjectPtr(ptr.Pointer()), C.int(factor))
	}
}

func (ptr *QFont) SetStrikeOut(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetStrikeOut(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetStyle(style QFont__Style) {
	if ptr.Pointer() != nil {
		C.QFont_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QFont) SetStyleHint(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	if ptr.Pointer() != nil {
		C.QFont_SetStyleHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.int(strategy))
	}
}

func (ptr *QFont) SetStyleName(styleName string) {
	if ptr.Pointer() != nil {
		C.QFont_SetStyleName(C.QtObjectPtr(ptr.Pointer()), C.CString(styleName))
	}
}

func (ptr *QFont) SetStyleStrategy(s QFont__StyleStrategy) {
	if ptr.Pointer() != nil {
		C.QFont_SetStyleStrategy(C.QtObjectPtr(ptr.Pointer()), C.int(s))
	}
}

func (ptr *QFont) SetUnderline(enable bool) {
	if ptr.Pointer() != nil {
		C.QFont_SetUnderline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFont) SetWeight(weight int) {
	if ptr.Pointer() != nil {
		C.QFont_SetWeight(C.QtObjectPtr(ptr.Pointer()), C.int(weight))
	}
}

func (ptr *QFont) Stretch() int {
	if ptr.Pointer() != nil {
		return int(C.QFont_Stretch(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) StrikeOut() bool {
	if ptr.Pointer() != nil {
		return C.QFont_StrikeOut(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Style() QFont__Style {
	if ptr.Pointer() != nil {
		return QFont__Style(C.QFont_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) StyleHint() QFont__StyleHint {
	if ptr.Pointer() != nil {
		return QFont__StyleHint(C.QFont_StyleHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) StyleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_StyleName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) StyleStrategy() QFont__StyleStrategy {
	if ptr.Pointer() != nil {
		return QFont__StyleStrategy(C.QFont_StyleStrategy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QFont_Substitute(familyName string) string {
	return C.GoString(C.QFont_QFont_Substitute(C.CString(familyName)))
}

func QFont_Substitutes(familyName string) []string {
	return strings.Split(C.GoString(C.QFont_QFont_Substitutes(C.CString(familyName))), "|")
}

func QFont_Substitutions() []string {
	return strings.Split(C.GoString(C.QFont_QFont_Substitutions()), "|")
}

func (ptr *QFont) Swap(other QFontITF) {
	if ptr.Pointer() != nil {
		C.QFont_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(other)))
	}
}

func (ptr *QFont) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFont_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFont) Underline() bool {
	if ptr.Pointer() != nil {
		return C.QFont_Underline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFont) Weight() int {
	if ptr.Pointer() != nil {
		return int(C.QFont_Weight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFont) DestroyQFont() {
	if ptr.Pointer() != nil {
		C.QFont_DestroyQFont(C.QtObjectPtr(ptr.Pointer()))
	}
}
