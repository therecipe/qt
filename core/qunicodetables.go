package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QUnicodeTables struct {
	ptr unsafe.Pointer
}

type QUnicodeTables_ITF interface {
	QUnicodeTables_PTR() *QUnicodeTables
}

func (p *QUnicodeTables) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUnicodeTables) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUnicodeTables(ptr QUnicodeTables_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUnicodeTables_PTR().Pointer()
	}
	return nil
}

func NewQUnicodeTablesFromPointer(ptr unsafe.Pointer) *QUnicodeTables {
	var n = new(QUnicodeTables)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUnicodeTables) QUnicodeTables_PTR() *QUnicodeTables {
	return ptr
}

//QUnicodeTables::GraphemeBreakClass
type QUnicodeTables__GraphemeBreakClass int64

const (
	QUnicodeTables__GraphemeBreak_Other             = QUnicodeTables__GraphemeBreakClass(0)
	QUnicodeTables__GraphemeBreak_CR                = QUnicodeTables__GraphemeBreakClass(1)
	QUnicodeTables__GraphemeBreak_LF                = QUnicodeTables__GraphemeBreakClass(2)
	QUnicodeTables__GraphemeBreak_Control           = QUnicodeTables__GraphemeBreakClass(3)
	QUnicodeTables__GraphemeBreak_Extend            = QUnicodeTables__GraphemeBreakClass(4)
	QUnicodeTables__GraphemeBreak_RegionalIndicator = QUnicodeTables__GraphemeBreakClass(5)
	QUnicodeTables__GraphemeBreak_Prepend           = QUnicodeTables__GraphemeBreakClass(6)
	QUnicodeTables__GraphemeBreak_SpacingMark       = QUnicodeTables__GraphemeBreakClass(7)
	QUnicodeTables__GraphemeBreak_L                 = QUnicodeTables__GraphemeBreakClass(8)
	QUnicodeTables__GraphemeBreak_V                 = QUnicodeTables__GraphemeBreakClass(9)
	QUnicodeTables__GraphemeBreak_T                 = QUnicodeTables__GraphemeBreakClass(10)
	QUnicodeTables__GraphemeBreak_LV                = QUnicodeTables__GraphemeBreakClass(11)
	QUnicodeTables__GraphemeBreak_LVT               = QUnicodeTables__GraphemeBreakClass(12)
)

//QUnicodeTables::LineBreakClass
type QUnicodeTables__LineBreakClass int64

const (
	QUnicodeTables__LineBreak_OP = QUnicodeTables__LineBreakClass(0)
	QUnicodeTables__LineBreak_CL = QUnicodeTables__LineBreakClass(1)
	QUnicodeTables__LineBreak_CP = QUnicodeTables__LineBreakClass(2)
	QUnicodeTables__LineBreak_QU = QUnicodeTables__LineBreakClass(3)
	QUnicodeTables__LineBreak_GL = QUnicodeTables__LineBreakClass(4)
	QUnicodeTables__LineBreak_NS = QUnicodeTables__LineBreakClass(5)
	QUnicodeTables__LineBreak_EX = QUnicodeTables__LineBreakClass(6)
	QUnicodeTables__LineBreak_SY = QUnicodeTables__LineBreakClass(7)
	QUnicodeTables__LineBreak_IS = QUnicodeTables__LineBreakClass(8)
	QUnicodeTables__LineBreak_PR = QUnicodeTables__LineBreakClass(9)
	QUnicodeTables__LineBreak_PO = QUnicodeTables__LineBreakClass(10)
	QUnicodeTables__LineBreak_NU = QUnicodeTables__LineBreakClass(11)
	QUnicodeTables__LineBreak_AL = QUnicodeTables__LineBreakClass(12)
	QUnicodeTables__LineBreak_HL = QUnicodeTables__LineBreakClass(13)
	QUnicodeTables__LineBreak_ID = QUnicodeTables__LineBreakClass(14)
	QUnicodeTables__LineBreak_IN = QUnicodeTables__LineBreakClass(15)
	QUnicodeTables__LineBreak_HY = QUnicodeTables__LineBreakClass(16)
	QUnicodeTables__LineBreak_BA = QUnicodeTables__LineBreakClass(17)
	QUnicodeTables__LineBreak_BB = QUnicodeTables__LineBreakClass(18)
	QUnicodeTables__LineBreak_B2 = QUnicodeTables__LineBreakClass(19)
	QUnicodeTables__LineBreak_ZW = QUnicodeTables__LineBreakClass(20)
	QUnicodeTables__LineBreak_CM = QUnicodeTables__LineBreakClass(21)
	QUnicodeTables__LineBreak_WJ = QUnicodeTables__LineBreakClass(22)
	QUnicodeTables__LineBreak_H2 = QUnicodeTables__LineBreakClass(23)
	QUnicodeTables__LineBreak_H3 = QUnicodeTables__LineBreakClass(24)
	QUnicodeTables__LineBreak_JL = QUnicodeTables__LineBreakClass(25)
	QUnicodeTables__LineBreak_JV = QUnicodeTables__LineBreakClass(26)
	QUnicodeTables__LineBreak_JT = QUnicodeTables__LineBreakClass(27)
	QUnicodeTables__LineBreak_RI = QUnicodeTables__LineBreakClass(28)
	QUnicodeTables__LineBreak_CB = QUnicodeTables__LineBreakClass(29)
	QUnicodeTables__LineBreak_SA = QUnicodeTables__LineBreakClass(30)
	QUnicodeTables__LineBreak_SG = QUnicodeTables__LineBreakClass(31)
	QUnicodeTables__LineBreak_SP = QUnicodeTables__LineBreakClass(32)
	QUnicodeTables__LineBreak_CR = QUnicodeTables__LineBreakClass(33)
	QUnicodeTables__LineBreak_LF = QUnicodeTables__LineBreakClass(34)
	QUnicodeTables__LineBreak_BK = QUnicodeTables__LineBreakClass(35)
)

//QUnicodeTables::SentenceBreakClass
type QUnicodeTables__SentenceBreakClass int64

const (
	QUnicodeTables__SentenceBreak_Other     = QUnicodeTables__SentenceBreakClass(0)
	QUnicodeTables__SentenceBreak_CR        = QUnicodeTables__SentenceBreakClass(1)
	QUnicodeTables__SentenceBreak_LF        = QUnicodeTables__SentenceBreakClass(2)
	QUnicodeTables__SentenceBreak_Sep       = QUnicodeTables__SentenceBreakClass(3)
	QUnicodeTables__SentenceBreak_Extend    = QUnicodeTables__SentenceBreakClass(4)
	QUnicodeTables__SentenceBreak_Sp        = QUnicodeTables__SentenceBreakClass(5)
	QUnicodeTables__SentenceBreak_Lower     = QUnicodeTables__SentenceBreakClass(6)
	QUnicodeTables__SentenceBreak_Upper     = QUnicodeTables__SentenceBreakClass(7)
	QUnicodeTables__SentenceBreak_OLetter   = QUnicodeTables__SentenceBreakClass(8)
	QUnicodeTables__SentenceBreak_Numeric   = QUnicodeTables__SentenceBreakClass(9)
	QUnicodeTables__SentenceBreak_ATerm     = QUnicodeTables__SentenceBreakClass(10)
	QUnicodeTables__SentenceBreak_SContinue = QUnicodeTables__SentenceBreakClass(11)
	QUnicodeTables__SentenceBreak_STerm     = QUnicodeTables__SentenceBreakClass(12)
	QUnicodeTables__SentenceBreak_Close     = QUnicodeTables__SentenceBreakClass(13)
)

//QUnicodeTables::WordBreakClass
type QUnicodeTables__WordBreakClass int64

const (
	QUnicodeTables__WordBreak_Other             = QUnicodeTables__WordBreakClass(0)
	QUnicodeTables__WordBreak_CR                = QUnicodeTables__WordBreakClass(1)
	QUnicodeTables__WordBreak_LF                = QUnicodeTables__WordBreakClass(2)
	QUnicodeTables__WordBreak_Newline           = QUnicodeTables__WordBreakClass(3)
	QUnicodeTables__WordBreak_Extend            = QUnicodeTables__WordBreakClass(4)
	QUnicodeTables__WordBreak_RegionalIndicator = QUnicodeTables__WordBreakClass(5)
	QUnicodeTables__WordBreak_Katakana          = QUnicodeTables__WordBreakClass(6)
	QUnicodeTables__WordBreak_HebrewLetter      = QUnicodeTables__WordBreakClass(7)
	QUnicodeTables__WordBreak_ALetter           = QUnicodeTables__WordBreakClass(8)
	QUnicodeTables__WordBreak_SingleQuote       = QUnicodeTables__WordBreakClass(9)
	QUnicodeTables__WordBreak_DoubleQuote       = QUnicodeTables__WordBreakClass(10)
	QUnicodeTables__WordBreak_MidNumLet         = QUnicodeTables__WordBreakClass(11)
	QUnicodeTables__WordBreak_MidLetter         = QUnicodeTables__WordBreakClass(12)
	QUnicodeTables__WordBreak_MidNum            = QUnicodeTables__WordBreakClass(13)
	QUnicodeTables__WordBreak_Numeric           = QUnicodeTables__WordBreakClass(14)
	QUnicodeTables__WordBreak_ExtendNumLet      = QUnicodeTables__WordBreakClass(15)
)
