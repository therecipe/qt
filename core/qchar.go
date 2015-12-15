package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QChar struct {
	ptr unsafe.Pointer
}

type QChar_ITF interface {
	QChar_PTR() *QChar
}

func (p *QChar) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QChar) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQChar(ptr QChar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QChar_PTR().Pointer()
	}
	return nil
}

func NewQCharFromPointer(ptr unsafe.Pointer) *QChar {
	var n = new(QChar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QChar) QChar_PTR() *QChar {
	return ptr
}

//QChar::Category
type QChar__Category int64

const (
	QChar__Mark_NonSpacing          = QChar__Category(0)
	QChar__Mark_SpacingCombining    = QChar__Category(1)
	QChar__Mark_Enclosing           = QChar__Category(2)
	QChar__Number_DecimalDigit      = QChar__Category(3)
	QChar__Number_Letter            = QChar__Category(4)
	QChar__Number_Other             = QChar__Category(5)
	QChar__Separator_Space          = QChar__Category(6)
	QChar__Separator_Line           = QChar__Category(7)
	QChar__Separator_Paragraph      = QChar__Category(8)
	QChar__Other_Control            = QChar__Category(9)
	QChar__Other_Format             = QChar__Category(10)
	QChar__Other_Surrogate          = QChar__Category(11)
	QChar__Other_PrivateUse         = QChar__Category(12)
	QChar__Other_NotAssigned        = QChar__Category(13)
	QChar__Letter_Uppercase         = QChar__Category(14)
	QChar__Letter_Lowercase         = QChar__Category(15)
	QChar__Letter_Titlecase         = QChar__Category(16)
	QChar__Letter_Modifier          = QChar__Category(17)
	QChar__Letter_Other             = QChar__Category(18)
	QChar__Punctuation_Connector    = QChar__Category(19)
	QChar__Punctuation_Dash         = QChar__Category(20)
	QChar__Punctuation_Open         = QChar__Category(21)
	QChar__Punctuation_Close        = QChar__Category(22)
	QChar__Punctuation_InitialQuote = QChar__Category(23)
	QChar__Punctuation_FinalQuote   = QChar__Category(24)
	QChar__Punctuation_Other        = QChar__Category(25)
	QChar__Symbol_Math              = QChar__Category(26)
	QChar__Symbol_Currency          = QChar__Category(27)
	QChar__Symbol_Modifier          = QChar__Category(28)
	QChar__Symbol_Other             = QChar__Category(29)
)

//QChar::Decomposition
type QChar__Decomposition int64

const (
	QChar__NoDecomposition = QChar__Decomposition(0)
	QChar__Canonical       = QChar__Decomposition(1)
	QChar__Font            = QChar__Decomposition(2)
	QChar__NoBreak         = QChar__Decomposition(3)
	QChar__Initial         = QChar__Decomposition(4)
	QChar__Medial          = QChar__Decomposition(5)
	QChar__Final           = QChar__Decomposition(6)
	QChar__Isolated        = QChar__Decomposition(7)
	QChar__Circle          = QChar__Decomposition(8)
	QChar__Super           = QChar__Decomposition(9)
	QChar__Sub             = QChar__Decomposition(10)
	QChar__Vertical        = QChar__Decomposition(11)
	QChar__Wide            = QChar__Decomposition(12)
	QChar__Narrow          = QChar__Decomposition(13)
	QChar__Small           = QChar__Decomposition(14)
	QChar__Square          = QChar__Decomposition(15)
	QChar__Compat          = QChar__Decomposition(16)
	QChar__Fraction        = QChar__Decomposition(17)
)

//QChar::Direction
type QChar__Direction int64

const (
	QChar__DirL   = QChar__Direction(0)
	QChar__DirR   = QChar__Direction(1)
	QChar__DirEN  = QChar__Direction(2)
	QChar__DirES  = QChar__Direction(3)
	QChar__DirET  = QChar__Direction(4)
	QChar__DirAN  = QChar__Direction(5)
	QChar__DirCS  = QChar__Direction(6)
	QChar__DirB   = QChar__Direction(7)
	QChar__DirS   = QChar__Direction(8)
	QChar__DirWS  = QChar__Direction(9)
	QChar__DirON  = QChar__Direction(10)
	QChar__DirLRE = QChar__Direction(11)
	QChar__DirLRO = QChar__Direction(12)
	QChar__DirAL  = QChar__Direction(13)
	QChar__DirRLE = QChar__Direction(14)
	QChar__DirRLO = QChar__Direction(15)
	QChar__DirPDF = QChar__Direction(16)
	QChar__DirNSM = QChar__Direction(17)
	QChar__DirBN  = QChar__Direction(18)
	QChar__DirLRI = QChar__Direction(19)
	QChar__DirRLI = QChar__Direction(20)
	QChar__DirFSI = QChar__Direction(21)
	QChar__DirPDI = QChar__Direction(22)
)

//QChar::JoiningType
type QChar__JoiningType int64

const (
	QChar__Joining_None        = QChar__JoiningType(0)
	QChar__Joining_Causing     = QChar__JoiningType(1)
	QChar__Joining_Dual        = QChar__JoiningType(2)
	QChar__Joining_Right       = QChar__JoiningType(3)
	QChar__Joining_Left        = QChar__JoiningType(4)
	QChar__Joining_Transparent = QChar__JoiningType(5)
)

//QChar::Script
type QChar__Script int64

const (
	QChar__Script_Unknown               = QChar__Script(0)
	QChar__Script_Inherited             = QChar__Script(1)
	QChar__Script_Common                = QChar__Script(2)
	QChar__Script_Latin                 = QChar__Script(3)
	QChar__Script_Greek                 = QChar__Script(4)
	QChar__Script_Cyrillic              = QChar__Script(5)
	QChar__Script_Armenian              = QChar__Script(6)
	QChar__Script_Hebrew                = QChar__Script(7)
	QChar__Script_Arabic                = QChar__Script(8)
	QChar__Script_Syriac                = QChar__Script(9)
	QChar__Script_Thaana                = QChar__Script(10)
	QChar__Script_Devanagari            = QChar__Script(11)
	QChar__Script_Bengali               = QChar__Script(12)
	QChar__Script_Gurmukhi              = QChar__Script(13)
	QChar__Script_Gujarati              = QChar__Script(14)
	QChar__Script_Oriya                 = QChar__Script(15)
	QChar__Script_Tamil                 = QChar__Script(16)
	QChar__Script_Telugu                = QChar__Script(17)
	QChar__Script_Kannada               = QChar__Script(18)
	QChar__Script_Malayalam             = QChar__Script(19)
	QChar__Script_Sinhala               = QChar__Script(20)
	QChar__Script_Thai                  = QChar__Script(21)
	QChar__Script_Lao                   = QChar__Script(22)
	QChar__Script_Tibetan               = QChar__Script(23)
	QChar__Script_Myanmar               = QChar__Script(24)
	QChar__Script_Georgian              = QChar__Script(25)
	QChar__Script_Hangul                = QChar__Script(26)
	QChar__Script_Ethiopic              = QChar__Script(27)
	QChar__Script_Cherokee              = QChar__Script(28)
	QChar__Script_CanadianAboriginal    = QChar__Script(29)
	QChar__Script_Ogham                 = QChar__Script(30)
	QChar__Script_Runic                 = QChar__Script(31)
	QChar__Script_Khmer                 = QChar__Script(32)
	QChar__Script_Mongolian             = QChar__Script(33)
	QChar__Script_Hiragana              = QChar__Script(34)
	QChar__Script_Katakana              = QChar__Script(35)
	QChar__Script_Bopomofo              = QChar__Script(36)
	QChar__Script_Han                   = QChar__Script(37)
	QChar__Script_Yi                    = QChar__Script(38)
	QChar__Script_OldItalic             = QChar__Script(39)
	QChar__Script_Gothic                = QChar__Script(40)
	QChar__Script_Deseret               = QChar__Script(41)
	QChar__Script_Tagalog               = QChar__Script(42)
	QChar__Script_Hanunoo               = QChar__Script(43)
	QChar__Script_Buhid                 = QChar__Script(44)
	QChar__Script_Tagbanwa              = QChar__Script(45)
	QChar__Script_Coptic                = QChar__Script(46)
	QChar__Script_Limbu                 = QChar__Script(47)
	QChar__Script_TaiLe                 = QChar__Script(48)
	QChar__Script_LinearB               = QChar__Script(49)
	QChar__Script_Ugaritic              = QChar__Script(50)
	QChar__Script_Shavian               = QChar__Script(51)
	QChar__Script_Osmanya               = QChar__Script(52)
	QChar__Script_Cypriot               = QChar__Script(53)
	QChar__Script_Braille               = QChar__Script(54)
	QChar__Script_Buginese              = QChar__Script(55)
	QChar__Script_NewTaiLue             = QChar__Script(56)
	QChar__Script_Glagolitic            = QChar__Script(57)
	QChar__Script_Tifinagh              = QChar__Script(58)
	QChar__Script_SylotiNagri           = QChar__Script(59)
	QChar__Script_OldPersian            = QChar__Script(60)
	QChar__Script_Kharoshthi            = QChar__Script(61)
	QChar__Script_Balinese              = QChar__Script(62)
	QChar__Script_Cuneiform             = QChar__Script(63)
	QChar__Script_Phoenician            = QChar__Script(64)
	QChar__Script_PhagsPa               = QChar__Script(65)
	QChar__Script_Nko                   = QChar__Script(66)
	QChar__Script_Sundanese             = QChar__Script(67)
	QChar__Script_Lepcha                = QChar__Script(68)
	QChar__Script_OlChiki               = QChar__Script(69)
	QChar__Script_Vai                   = QChar__Script(70)
	QChar__Script_Saurashtra            = QChar__Script(71)
	QChar__Script_KayahLi               = QChar__Script(72)
	QChar__Script_Rejang                = QChar__Script(73)
	QChar__Script_Lycian                = QChar__Script(74)
	QChar__Script_Carian                = QChar__Script(75)
	QChar__Script_Lydian                = QChar__Script(76)
	QChar__Script_Cham                  = QChar__Script(77)
	QChar__Script_TaiTham               = QChar__Script(78)
	QChar__Script_TaiViet               = QChar__Script(79)
	QChar__Script_Avestan               = QChar__Script(80)
	QChar__Script_EgyptianHieroglyphs   = QChar__Script(81)
	QChar__Script_Samaritan             = QChar__Script(82)
	QChar__Script_Lisu                  = QChar__Script(83)
	QChar__Script_Bamum                 = QChar__Script(84)
	QChar__Script_Javanese              = QChar__Script(85)
	QChar__Script_MeeteiMayek           = QChar__Script(86)
	QChar__Script_ImperialAramaic       = QChar__Script(87)
	QChar__Script_OldSouthArabian       = QChar__Script(88)
	QChar__Script_InscriptionalParthian = QChar__Script(89)
	QChar__Script_InscriptionalPahlavi  = QChar__Script(90)
	QChar__Script_OldTurkic             = QChar__Script(91)
	QChar__Script_Kaithi                = QChar__Script(92)
	QChar__Script_Batak                 = QChar__Script(93)
	QChar__Script_Brahmi                = QChar__Script(94)
	QChar__Script_Mandaic               = QChar__Script(95)
	QChar__Script_Chakma                = QChar__Script(96)
	QChar__Script_MeroiticCursive       = QChar__Script(97)
	QChar__Script_MeroiticHieroglyphs   = QChar__Script(98)
	QChar__Script_Miao                  = QChar__Script(99)
	QChar__Script_Sharada               = QChar__Script(100)
	QChar__Script_SoraSompeng           = QChar__Script(101)
	QChar__Script_Takri                 = QChar__Script(102)
	QChar__Script_CaucasianAlbanian     = QChar__Script(103)
	QChar__Script_BassaVah              = QChar__Script(104)
	QChar__Script_Duployan              = QChar__Script(105)
	QChar__Script_Elbasan               = QChar__Script(106)
	QChar__Script_Grantha               = QChar__Script(107)
	QChar__Script_PahawhHmong           = QChar__Script(108)
	QChar__Script_Khojki                = QChar__Script(109)
	QChar__Script_LinearA               = QChar__Script(110)
	QChar__Script_Mahajani              = QChar__Script(111)
	QChar__Script_Manichaean            = QChar__Script(112)
	QChar__Script_MendeKikakui          = QChar__Script(113)
	QChar__Script_Modi                  = QChar__Script(114)
	QChar__Script_Mro                   = QChar__Script(115)
	QChar__Script_OldNorthArabian       = QChar__Script(116)
	QChar__Script_Nabataean             = QChar__Script(117)
	QChar__Script_Palmyrene             = QChar__Script(118)
	QChar__Script_PauCinHau             = QChar__Script(119)
	QChar__Script_OldPermic             = QChar__Script(120)
	QChar__Script_PsalterPahlavi        = QChar__Script(121)
	QChar__Script_Siddham               = QChar__Script(122)
	QChar__Script_Khudawadi             = QChar__Script(123)
	QChar__Script_Tirhuta               = QChar__Script(124)
	QChar__Script_WarangCiti            = QChar__Script(125)
	QChar__ScriptCount                  = QChar__Script(126)
)

//QChar::SpecialCharacter
type QChar__SpecialCharacter int64

const (
	QChar__Null                       = QChar__SpecialCharacter(0x0000)
	QChar__Tabulation                 = QChar__SpecialCharacter(0x0009)
	QChar__LineFeed                   = QChar__SpecialCharacter(0x000a)
	QChar__CarriageReturn             = QChar__SpecialCharacter(0x000d)
	QChar__Space                      = QChar__SpecialCharacter(0x0020)
	QChar__Nbsp                       = QChar__SpecialCharacter(0x00a0)
	QChar__SoftHyphen                 = QChar__SpecialCharacter(0x00ad)
	QChar__ReplacementCharacter       = QChar__SpecialCharacter(0xfffd)
	QChar__ObjectReplacementCharacter = QChar__SpecialCharacter(0xfffc)
	QChar__ByteOrderMark              = QChar__SpecialCharacter(0xfeff)
	QChar__ByteOrderSwapped           = QChar__SpecialCharacter(0xfffe)
	QChar__ParagraphSeparator         = QChar__SpecialCharacter(0x2029)
	QChar__LineSeparator              = QChar__SpecialCharacter(0x2028)
	QChar__LastValidCodePoint         = QChar__SpecialCharacter(0x10ffff)
)

//QChar::UnicodeVersion
type QChar__UnicodeVersion int64

const (
	QChar__Unicode_Unassigned = QChar__UnicodeVersion(0)
	QChar__Unicode_1_1        = QChar__UnicodeVersion(1)
	QChar__Unicode_2_0        = QChar__UnicodeVersion(2)
	QChar__Unicode_2_1_2      = QChar__UnicodeVersion(3)
	QChar__Unicode_3_0        = QChar__UnicodeVersion(4)
	QChar__Unicode_3_1        = QChar__UnicodeVersion(5)
	QChar__Unicode_3_2        = QChar__UnicodeVersion(6)
	QChar__Unicode_4_0        = QChar__UnicodeVersion(7)
	QChar__Unicode_4_1        = QChar__UnicodeVersion(8)
	QChar__Unicode_5_0        = QChar__UnicodeVersion(9)
	QChar__Unicode_5_1        = QChar__UnicodeVersion(10)
	QChar__Unicode_5_2        = QChar__UnicodeVersion(11)
	QChar__Unicode_6_0        = QChar__UnicodeVersion(12)
	QChar__Unicode_6_1        = QChar__UnicodeVersion(13)
	QChar__Unicode_6_2        = QChar__UnicodeVersion(14)
	QChar__Unicode_6_3        = QChar__UnicodeVersion(15)
	QChar__Unicode_7_0        = QChar__UnicodeVersion(16)
)

func NewQChar() *QChar {
	defer qt.Recovering("QChar::QChar")

	return NewQCharFromPointer(C.QChar_NewQChar())
}

func NewQChar8(ch QLatin1Char_ITF) *QChar {
	defer qt.Recovering("QChar::QChar")

	return NewQCharFromPointer(C.QChar_NewQChar8(PointerFromQLatin1Char(ch)))
}

func NewQChar7(ch QChar__SpecialCharacter) *QChar {
	defer qt.Recovering("QChar::QChar")

	return NewQCharFromPointer(C.QChar_NewQChar7(C.int(ch)))
}

func NewQChar9(ch string) *QChar {
	defer qt.Recovering("QChar::QChar")

	return NewQCharFromPointer(C.QChar_NewQChar9(C.CString(ch)))
}

func NewQChar6(code int) *QChar {
	defer qt.Recovering("QChar::QChar")

	return NewQCharFromPointer(C.QChar_NewQChar6(C.int(code)))
}

func (ptr *QChar) Category() QChar__Category {
	defer qt.Recovering("QChar::category")

	if ptr.Pointer() != nil {
		return QChar__Category(C.QChar_Category(ptr.Pointer()))
	}
	return 0
}

func QChar_CurrentUnicodeVersion() QChar__UnicodeVersion {
	defer qt.Recovering("QChar::currentUnicodeVersion")

	return QChar__UnicodeVersion(C.QChar_QChar_CurrentUnicodeVersion())
}

func (ptr *QChar) Decomposition() string {
	defer qt.Recovering("QChar::decomposition")

	if ptr.Pointer() != nil {
		return C.GoString(C.QChar_Decomposition(ptr.Pointer()))
	}
	return ""
}

func (ptr *QChar) DecompositionTag() QChar__Decomposition {
	defer qt.Recovering("QChar::decompositionTag")

	if ptr.Pointer() != nil {
		return QChar__Decomposition(C.QChar_DecompositionTag(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChar) DigitValue() int {
	defer qt.Recovering("QChar::digitValue")

	if ptr.Pointer() != nil {
		return int(C.QChar_DigitValue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChar) Direction() QChar__Direction {
	defer qt.Recovering("QChar::direction")

	if ptr.Pointer() != nil {
		return QChar__Direction(C.QChar_Direction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChar) HasMirrored() bool {
	defer qt.Recovering("QChar::hasMirrored")

	if ptr.Pointer() != nil {
		return C.QChar_HasMirrored(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsDigit() bool {
	defer qt.Recovering("QChar::isDigit")

	if ptr.Pointer() != nil {
		return C.QChar_IsDigit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsHighSurrogate() bool {
	defer qt.Recovering("QChar::isHighSurrogate")

	if ptr.Pointer() != nil {
		return C.QChar_IsHighSurrogate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsLetter() bool {
	defer qt.Recovering("QChar::isLetter")

	if ptr.Pointer() != nil {
		return C.QChar_IsLetter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsLetterOrNumber() bool {
	defer qt.Recovering("QChar::isLetterOrNumber")

	if ptr.Pointer() != nil {
		return C.QChar_IsLetterOrNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsLower() bool {
	defer qt.Recovering("QChar::isLower")

	if ptr.Pointer() != nil {
		return C.QChar_IsLower(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsLowSurrogate() bool {
	defer qt.Recovering("QChar::isLowSurrogate")

	if ptr.Pointer() != nil {
		return C.QChar_IsLowSurrogate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsMark() bool {
	defer qt.Recovering("QChar::isMark")

	if ptr.Pointer() != nil {
		return C.QChar_IsMark(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsNonCharacter() bool {
	defer qt.Recovering("QChar::isNonCharacter")

	if ptr.Pointer() != nil {
		return C.QChar_IsNonCharacter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsNull() bool {
	defer qt.Recovering("QChar::isNull")

	if ptr.Pointer() != nil {
		return C.QChar_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsNumber() bool {
	defer qt.Recovering("QChar::isNumber")

	if ptr.Pointer() != nil {
		return C.QChar_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsPrint() bool {
	defer qt.Recovering("QChar::isPrint")

	if ptr.Pointer() != nil {
		return C.QChar_IsPrint(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsPunct() bool {
	defer qt.Recovering("QChar::isPunct")

	if ptr.Pointer() != nil {
		return C.QChar_IsPunct(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsSpace() bool {
	defer qt.Recovering("QChar::isSpace")

	if ptr.Pointer() != nil {
		return C.QChar_IsSpace(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsSurrogate() bool {
	defer qt.Recovering("QChar::isSurrogate")

	if ptr.Pointer() != nil {
		return C.QChar_IsSurrogate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsSymbol() bool {
	defer qt.Recovering("QChar::isSymbol")

	if ptr.Pointer() != nil {
		return C.QChar_IsSymbol(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsTitleCase() bool {
	defer qt.Recovering("QChar::isTitleCase")

	if ptr.Pointer() != nil {
		return C.QChar_IsTitleCase(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) IsUpper() bool {
	defer qt.Recovering("QChar::isUpper")

	if ptr.Pointer() != nil {
		return C.QChar_IsUpper(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QChar) JoiningType() QChar__JoiningType {
	defer qt.Recovering("QChar::joiningType")

	if ptr.Pointer() != nil {
		return QChar__JoiningType(C.QChar_JoiningType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChar) Script() QChar__Script {
	defer qt.Recovering("QChar::script")

	if ptr.Pointer() != nil {
		return QChar__Script(C.QChar_Script(ptr.Pointer()))
	}
	return 0
}

func (ptr *QChar) UnicodeVersion() QChar__UnicodeVersion {
	defer qt.Recovering("QChar::unicodeVersion")

	if ptr.Pointer() != nil {
		return QChar__UnicodeVersion(C.QChar_UnicodeVersion(ptr.Pointer()))
	}
	return 0
}
