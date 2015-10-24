package gui

//#include "qpagesize.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPageSize struct {
	ptr unsafe.Pointer
}

type QPageSizeITF interface {
	QPageSizePTR() *QPageSize
}

func (p *QPageSize) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPageSize) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPageSize(ptr QPageSizeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPageSizePTR().Pointer()
	}
	return nil
}

func QPageSizeFromPointer(ptr unsafe.Pointer) *QPageSize {
	var n = new(QPageSize)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPageSize) QPageSizePTR() *QPageSize {
	return ptr
}

//QPageSize::PageSizeId
type QPageSize__PageSizeId int

var (
	QPageSize__A4                 = QPageSize__PageSizeId(0)
	QPageSize__B5                 = QPageSize__PageSizeId(1)
	QPageSize__Letter             = QPageSize__PageSizeId(2)
	QPageSize__Legal              = QPageSize__PageSizeId(3)
	QPageSize__Executive          = QPageSize__PageSizeId(4)
	QPageSize__A0                 = QPageSize__PageSizeId(5)
	QPageSize__A1                 = QPageSize__PageSizeId(6)
	QPageSize__A2                 = QPageSize__PageSizeId(7)
	QPageSize__A3                 = QPageSize__PageSizeId(8)
	QPageSize__A5                 = QPageSize__PageSizeId(9)
	QPageSize__A6                 = QPageSize__PageSizeId(10)
	QPageSize__A7                 = QPageSize__PageSizeId(11)
	QPageSize__A8                 = QPageSize__PageSizeId(12)
	QPageSize__A9                 = QPageSize__PageSizeId(13)
	QPageSize__B0                 = QPageSize__PageSizeId(14)
	QPageSize__B1                 = QPageSize__PageSizeId(15)
	QPageSize__B10                = QPageSize__PageSizeId(16)
	QPageSize__B2                 = QPageSize__PageSizeId(17)
	QPageSize__B3                 = QPageSize__PageSizeId(18)
	QPageSize__B4                 = QPageSize__PageSizeId(19)
	QPageSize__B6                 = QPageSize__PageSizeId(20)
	QPageSize__B7                 = QPageSize__PageSizeId(21)
	QPageSize__B8                 = QPageSize__PageSizeId(22)
	QPageSize__B9                 = QPageSize__PageSizeId(23)
	QPageSize__C5E                = QPageSize__PageSizeId(24)
	QPageSize__Comm10E            = QPageSize__PageSizeId(25)
	QPageSize__DLE                = QPageSize__PageSizeId(26)
	QPageSize__Folio              = QPageSize__PageSizeId(27)
	QPageSize__Ledger             = QPageSize__PageSizeId(28)
	QPageSize__Tabloid            = QPageSize__PageSizeId(29)
	QPageSize__Custom             = QPageSize__PageSizeId(30)
	QPageSize__A10                = QPageSize__PageSizeId(31)
	QPageSize__A3Extra            = QPageSize__PageSizeId(32)
	QPageSize__A4Extra            = QPageSize__PageSizeId(33)
	QPageSize__A4Plus             = QPageSize__PageSizeId(34)
	QPageSize__A4Small            = QPageSize__PageSizeId(35)
	QPageSize__A5Extra            = QPageSize__PageSizeId(36)
	QPageSize__B5Extra            = QPageSize__PageSizeId(37)
	QPageSize__JisB0              = QPageSize__PageSizeId(38)
	QPageSize__JisB1              = QPageSize__PageSizeId(39)
	QPageSize__JisB2              = QPageSize__PageSizeId(40)
	QPageSize__JisB3              = QPageSize__PageSizeId(41)
	QPageSize__JisB4              = QPageSize__PageSizeId(42)
	QPageSize__JisB5              = QPageSize__PageSizeId(43)
	QPageSize__JisB6              = QPageSize__PageSizeId(44)
	QPageSize__JisB7              = QPageSize__PageSizeId(45)
	QPageSize__JisB8              = QPageSize__PageSizeId(46)
	QPageSize__JisB9              = QPageSize__PageSizeId(47)
	QPageSize__JisB10             = QPageSize__PageSizeId(48)
	QPageSize__AnsiC              = QPageSize__PageSizeId(49)
	QPageSize__AnsiD              = QPageSize__PageSizeId(50)
	QPageSize__AnsiE              = QPageSize__PageSizeId(51)
	QPageSize__LegalExtra         = QPageSize__PageSizeId(52)
	QPageSize__LetterExtra        = QPageSize__PageSizeId(53)
	QPageSize__LetterPlus         = QPageSize__PageSizeId(54)
	QPageSize__LetterSmall        = QPageSize__PageSizeId(55)
	QPageSize__TabloidExtra       = QPageSize__PageSizeId(56)
	QPageSize__ArchA              = QPageSize__PageSizeId(57)
	QPageSize__ArchB              = QPageSize__PageSizeId(58)
	QPageSize__ArchC              = QPageSize__PageSizeId(59)
	QPageSize__ArchD              = QPageSize__PageSizeId(60)
	QPageSize__ArchE              = QPageSize__PageSizeId(61)
	QPageSize__Imperial7x9        = QPageSize__PageSizeId(62)
	QPageSize__Imperial8x10       = QPageSize__PageSizeId(63)
	QPageSize__Imperial9x11       = QPageSize__PageSizeId(64)
	QPageSize__Imperial9x12       = QPageSize__PageSizeId(65)
	QPageSize__Imperial10x11      = QPageSize__PageSizeId(66)
	QPageSize__Imperial10x13      = QPageSize__PageSizeId(67)
	QPageSize__Imperial10x14      = QPageSize__PageSizeId(68)
	QPageSize__Imperial12x11      = QPageSize__PageSizeId(69)
	QPageSize__Imperial15x11      = QPageSize__PageSizeId(70)
	QPageSize__ExecutiveStandard  = QPageSize__PageSizeId(71)
	QPageSize__Note               = QPageSize__PageSizeId(72)
	QPageSize__Quarto             = QPageSize__PageSizeId(73)
	QPageSize__Statement          = QPageSize__PageSizeId(74)
	QPageSize__SuperA             = QPageSize__PageSizeId(75)
	QPageSize__SuperB             = QPageSize__PageSizeId(76)
	QPageSize__Postcard           = QPageSize__PageSizeId(77)
	QPageSize__DoublePostcard     = QPageSize__PageSizeId(78)
	QPageSize__Prc16K             = QPageSize__PageSizeId(79)
	QPageSize__Prc32K             = QPageSize__PageSizeId(80)
	QPageSize__Prc32KBig          = QPageSize__PageSizeId(81)
	QPageSize__FanFoldUS          = QPageSize__PageSizeId(82)
	QPageSize__FanFoldGerman      = QPageSize__PageSizeId(83)
	QPageSize__FanFoldGermanLegal = QPageSize__PageSizeId(84)
	QPageSize__EnvelopeB4         = QPageSize__PageSizeId(85)
	QPageSize__EnvelopeB5         = QPageSize__PageSizeId(86)
	QPageSize__EnvelopeB6         = QPageSize__PageSizeId(87)
	QPageSize__EnvelopeC0         = QPageSize__PageSizeId(88)
	QPageSize__EnvelopeC1         = QPageSize__PageSizeId(89)
	QPageSize__EnvelopeC2         = QPageSize__PageSizeId(90)
	QPageSize__EnvelopeC3         = QPageSize__PageSizeId(91)
	QPageSize__EnvelopeC4         = QPageSize__PageSizeId(92)
	QPageSize__EnvelopeC6         = QPageSize__PageSizeId(93)
	QPageSize__EnvelopeC65        = QPageSize__PageSizeId(94)
	QPageSize__EnvelopeC7         = QPageSize__PageSizeId(95)
	QPageSize__Envelope9          = QPageSize__PageSizeId(96)
	QPageSize__Envelope11         = QPageSize__PageSizeId(97)
	QPageSize__Envelope12         = QPageSize__PageSizeId(98)
	QPageSize__Envelope14         = QPageSize__PageSizeId(99)
	QPageSize__EnvelopeMonarch    = QPageSize__PageSizeId(100)
	QPageSize__EnvelopePersonal   = QPageSize__PageSizeId(101)
	QPageSize__EnvelopeChou3      = QPageSize__PageSizeId(102)
	QPageSize__EnvelopeChou4      = QPageSize__PageSizeId(103)
	QPageSize__EnvelopeInvite     = QPageSize__PageSizeId(104)
	QPageSize__EnvelopeItalian    = QPageSize__PageSizeId(105)
	QPageSize__EnvelopeKaku2      = QPageSize__PageSizeId(106)
	QPageSize__EnvelopeKaku3      = QPageSize__PageSizeId(107)
	QPageSize__EnvelopePrc1       = QPageSize__PageSizeId(108)
	QPageSize__EnvelopePrc2       = QPageSize__PageSizeId(109)
	QPageSize__EnvelopePrc3       = QPageSize__PageSizeId(110)
	QPageSize__EnvelopePrc4       = QPageSize__PageSizeId(111)
	QPageSize__EnvelopePrc5       = QPageSize__PageSizeId(112)
	QPageSize__EnvelopePrc6       = QPageSize__PageSizeId(113)
	QPageSize__EnvelopePrc7       = QPageSize__PageSizeId(114)
	QPageSize__EnvelopePrc8       = QPageSize__PageSizeId(115)
	QPageSize__EnvelopePrc9       = QPageSize__PageSizeId(116)
	QPageSize__EnvelopePrc10      = QPageSize__PageSizeId(117)
	QPageSize__EnvelopeYou4       = QPageSize__PageSizeId(118)
	QPageSize__LastPageSize       = QPageSize__PageSizeId(QPageSize__EnvelopeYou4)
	QPageSize__NPageSize          = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__NPaperSize         = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__AnsiA              = QPageSize__PageSizeId(QPageSize__Letter)
	QPageSize__AnsiB              = QPageSize__PageSizeId(QPageSize__Ledger)
	QPageSize__EnvelopeC5         = QPageSize__PageSizeId(QPageSize__C5E)
	QPageSize__EnvelopeDL         = QPageSize__PageSizeId(QPageSize__DLE)
	QPageSize__Envelope10         = QPageSize__PageSizeId(QPageSize__Comm10E)
)

//QPageSize::SizeMatchPolicy
type QPageSize__SizeMatchPolicy int

var (
	QPageSize__FuzzyMatch            = QPageSize__SizeMatchPolicy(0)
	QPageSize__FuzzyOrientationMatch = QPageSize__SizeMatchPolicy(1)
	QPageSize__ExactMatch            = QPageSize__SizeMatchPolicy(2)
)

//QPageSize::Unit
type QPageSize__Unit int

var (
	QPageSize__Millimeter = QPageSize__Unit(0)
	QPageSize__Point      = QPageSize__Unit(1)
	QPageSize__Inch       = QPageSize__Unit(2)
	QPageSize__Pica       = QPageSize__Unit(3)
	QPageSize__Didot      = QPageSize__Unit(4)
	QPageSize__Cicero     = QPageSize__Unit(5)
)

func NewQPageSize() *QPageSize {
	return QPageSizeFromPointer(unsafe.Pointer(C.QPageSize_NewQPageSize()))
}

func NewQPageSize2(pageSize QPageSize__PageSizeId) *QPageSize {
	return QPageSizeFromPointer(unsafe.Pointer(C.QPageSize_NewQPageSize2(C.int(pageSize))))
}

func NewQPageSize5(other QPageSizeITF) *QPageSize {
	return QPageSizeFromPointer(unsafe.Pointer(C.QPageSize_NewQPageSize5(C.QtObjectPtr(PointerFromQPageSize(other)))))
}

func NewQPageSize3(pointSize core.QSizeITF, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	return QPageSizeFromPointer(unsafe.Pointer(C.QPageSize_NewQPageSize3(C.QtObjectPtr(core.PointerFromQSize(pointSize)), C.CString(name), C.int(matchPolicy))))
}

func NewQPageSize4(size core.QSizeFITF, units QPageSize__Unit, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	return QPageSizeFromPointer(unsafe.Pointer(C.QPageSize_NewQPageSize4(C.QtObjectPtr(core.PointerFromQSizeF(size)), C.int(units), C.CString(name), C.int(matchPolicy))))
}

func QPageSize_DefinitionUnits2(pageSizeId QPageSize__PageSizeId) QPageSize__Unit {
	return QPageSize__Unit(C.QPageSize_QPageSize_DefinitionUnits2(C.int(pageSizeId)))
}

func (ptr *QPageSize) DefinitionUnits() QPageSize__Unit {
	if ptr.Pointer() != nil {
		return QPageSize__Unit(C.QPageSize_DefinitionUnits(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QPageSize_Id2(pointSize core.QSizeITF, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id2(C.QtObjectPtr(core.PointerFromQSize(pointSize)), C.int(matchPolicy)))
}

func QPageSize_Id3(size core.QSizeFITF, units QPageSize__Unit, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id3(C.QtObjectPtr(core.PointerFromQSizeF(size)), C.int(units), C.int(matchPolicy)))
}

func QPageSize_Id4(windowsId int) QPageSize__PageSizeId {
	return QPageSize__PageSizeId(C.QPageSize_QPageSize_Id4(C.int(windowsId)))
}

func (ptr *QPageSize) Id() QPageSize__PageSizeId {
	if ptr.Pointer() != nil {
		return QPageSize__PageSizeId(C.QPageSize_Id(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageSize) IsEquivalentTo(other QPageSizeITF) bool {
	if ptr.Pointer() != nil {
		return C.QPageSize_IsEquivalentTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageSize(other))) != 0
	}
	return false
}

func (ptr *QPageSize) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QPageSize_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QPageSize_Key2(pageSizeId QPageSize__PageSizeId) string {
	return C.GoString(C.QPageSize_QPageSize_Key2(C.int(pageSizeId)))
}

func (ptr *QPageSize) Key() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPageSize_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QPageSize_Name2(pageSizeId QPageSize__PageSizeId) string {
	return C.GoString(C.QPageSize_QPageSize_Name2(C.int(pageSizeId)))
}

func (ptr *QPageSize) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QPageSize_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QPageSize) Swap(other QPageSizeITF) {
	if ptr.Pointer() != nil {
		C.QPageSize_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPageSize(other)))
	}
}

func QPageSize_WindowsId2(pageSizeId QPageSize__PageSizeId) int {
	return int(C.QPageSize_QPageSize_WindowsId2(C.int(pageSizeId)))
}

func (ptr *QPageSize) WindowsId() int {
	if ptr.Pointer() != nil {
		return int(C.QPageSize_WindowsId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPageSize) DestroyQPageSize() {
	if ptr.Pointer() != nil {
		C.QPageSize_DestroyQPageSize(C.QtObjectPtr(ptr.Pointer()))
	}
}
