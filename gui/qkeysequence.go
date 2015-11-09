package gui

//#include "qkeysequence.h"
import "C"
import (
	"unsafe"
)

type QKeySequence struct {
	ptr unsafe.Pointer
}

type QKeySequence_ITF interface {
	QKeySequence_PTR() *QKeySequence
}

func (p *QKeySequence) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QKeySequence) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQKeySequence(ptr QKeySequence_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequence_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceFromPointer(ptr unsafe.Pointer) *QKeySequence {
	var n = new(QKeySequence)
	n.SetPointer(ptr)
	return n
}

func (ptr *QKeySequence) QKeySequence_PTR() *QKeySequence {
	return ptr
}

//QKeySequence::SequenceFormat
type QKeySequence__SequenceFormat int64

const (
	QKeySequence__NativeText   = QKeySequence__SequenceFormat(0)
	QKeySequence__PortableText = QKeySequence__SequenceFormat(1)
)

//QKeySequence::SequenceMatch
type QKeySequence__SequenceMatch int64

const (
	QKeySequence__NoMatch      = QKeySequence__SequenceMatch(0)
	QKeySequence__PartialMatch = QKeySequence__SequenceMatch(1)
	QKeySequence__ExactMatch   = QKeySequence__SequenceMatch(2)
)

//QKeySequence::StandardKey
type QKeySequence__StandardKey int64

const (
	QKeySequence__UnknownKey               = QKeySequence__StandardKey(0)
	QKeySequence__HelpContents             = QKeySequence__StandardKey(1)
	QKeySequence__WhatsThis                = QKeySequence__StandardKey(2)
	QKeySequence__Open                     = QKeySequence__StandardKey(3)
	QKeySequence__Close                    = QKeySequence__StandardKey(4)
	QKeySequence__Save                     = QKeySequence__StandardKey(5)
	QKeySequence__New                      = QKeySequence__StandardKey(6)
	QKeySequence__Delete                   = QKeySequence__StandardKey(7)
	QKeySequence__Cut                      = QKeySequence__StandardKey(8)
	QKeySequence__Copy                     = QKeySequence__StandardKey(9)
	QKeySequence__Paste                    = QKeySequence__StandardKey(10)
	QKeySequence__Undo                     = QKeySequence__StandardKey(11)
	QKeySequence__Redo                     = QKeySequence__StandardKey(12)
	QKeySequence__Back                     = QKeySequence__StandardKey(13)
	QKeySequence__Forward                  = QKeySequence__StandardKey(14)
	QKeySequence__Refresh                  = QKeySequence__StandardKey(15)
	QKeySequence__ZoomIn                   = QKeySequence__StandardKey(16)
	QKeySequence__ZoomOut                  = QKeySequence__StandardKey(17)
	QKeySequence__Print                    = QKeySequence__StandardKey(18)
	QKeySequence__AddTab                   = QKeySequence__StandardKey(19)
	QKeySequence__NextChild                = QKeySequence__StandardKey(20)
	QKeySequence__PreviousChild            = QKeySequence__StandardKey(21)
	QKeySequence__Find                     = QKeySequence__StandardKey(22)
	QKeySequence__FindNext                 = QKeySequence__StandardKey(23)
	QKeySequence__FindPrevious             = QKeySequence__StandardKey(24)
	QKeySequence__Replace                  = QKeySequence__StandardKey(25)
	QKeySequence__SelectAll                = QKeySequence__StandardKey(26)
	QKeySequence__Bold                     = QKeySequence__StandardKey(27)
	QKeySequence__Italic                   = QKeySequence__StandardKey(28)
	QKeySequence__Underline                = QKeySequence__StandardKey(29)
	QKeySequence__MoveToNextChar           = QKeySequence__StandardKey(30)
	QKeySequence__MoveToPreviousChar       = QKeySequence__StandardKey(31)
	QKeySequence__MoveToNextWord           = QKeySequence__StandardKey(32)
	QKeySequence__MoveToPreviousWord       = QKeySequence__StandardKey(33)
	QKeySequence__MoveToNextLine           = QKeySequence__StandardKey(34)
	QKeySequence__MoveToPreviousLine       = QKeySequence__StandardKey(35)
	QKeySequence__MoveToNextPage           = QKeySequence__StandardKey(36)
	QKeySequence__MoveToPreviousPage       = QKeySequence__StandardKey(37)
	QKeySequence__MoveToStartOfLine        = QKeySequence__StandardKey(38)
	QKeySequence__MoveToEndOfLine          = QKeySequence__StandardKey(39)
	QKeySequence__MoveToStartOfBlock       = QKeySequence__StandardKey(40)
	QKeySequence__MoveToEndOfBlock         = QKeySequence__StandardKey(41)
	QKeySequence__MoveToStartOfDocument    = QKeySequence__StandardKey(42)
	QKeySequence__MoveToEndOfDocument      = QKeySequence__StandardKey(43)
	QKeySequence__SelectNextChar           = QKeySequence__StandardKey(44)
	QKeySequence__SelectPreviousChar       = QKeySequence__StandardKey(45)
	QKeySequence__SelectNextWord           = QKeySequence__StandardKey(46)
	QKeySequence__SelectPreviousWord       = QKeySequence__StandardKey(47)
	QKeySequence__SelectNextLine           = QKeySequence__StandardKey(48)
	QKeySequence__SelectPreviousLine       = QKeySequence__StandardKey(49)
	QKeySequence__SelectNextPage           = QKeySequence__StandardKey(50)
	QKeySequence__SelectPreviousPage       = QKeySequence__StandardKey(51)
	QKeySequence__SelectStartOfLine        = QKeySequence__StandardKey(52)
	QKeySequence__SelectEndOfLine          = QKeySequence__StandardKey(53)
	QKeySequence__SelectStartOfBlock       = QKeySequence__StandardKey(54)
	QKeySequence__SelectEndOfBlock         = QKeySequence__StandardKey(55)
	QKeySequence__SelectStartOfDocument    = QKeySequence__StandardKey(56)
	QKeySequence__SelectEndOfDocument      = QKeySequence__StandardKey(57)
	QKeySequence__DeleteStartOfWord        = QKeySequence__StandardKey(58)
	QKeySequence__DeleteEndOfWord          = QKeySequence__StandardKey(59)
	QKeySequence__DeleteEndOfLine          = QKeySequence__StandardKey(60)
	QKeySequence__InsertParagraphSeparator = QKeySequence__StandardKey(61)
	QKeySequence__InsertLineSeparator      = QKeySequence__StandardKey(62)
	QKeySequence__SaveAs                   = QKeySequence__StandardKey(63)
	QKeySequence__Preferences              = QKeySequence__StandardKey(64)
	QKeySequence__Quit                     = QKeySequence__StandardKey(65)
	QKeySequence__FullScreen               = QKeySequence__StandardKey(66)
	QKeySequence__Deselect                 = QKeySequence__StandardKey(67)
	QKeySequence__DeleteCompleteLine       = QKeySequence__StandardKey(68)
	QKeySequence__Backspace                = QKeySequence__StandardKey(69)
)

func NewQKeySequence() *QKeySequence {
	return NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence())
}

func NewQKeySequence5(key QKeySequence__StandardKey) *QKeySequence {
	return NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence5(C.int(key)))
}

func NewQKeySequence4(keysequence QKeySequence_ITF) *QKeySequence {
	return NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence4(PointerFromQKeySequence(keysequence)))
}

func NewQKeySequence2(key string, format QKeySequence__SequenceFormat) *QKeySequence {
	return NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence2(C.CString(key), C.int(format)))
}

func NewQKeySequence3(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	return NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence3(C.int(k1), C.int(k2), C.int(k3), C.int(k4)))
}

func (ptr *QKeySequence) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QKeySequence_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeySequence) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QKeySequence_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QKeySequence) Matches(seq QKeySequence_ITF) QKeySequence__SequenceMatch {
	if ptr.Pointer() != nil {
		return QKeySequence__SequenceMatch(C.QKeySequence_Matches(ptr.Pointer(), PointerFromQKeySequence(seq)))
	}
	return 0
}

func (ptr *QKeySequence) Swap(other QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence_Swap(ptr.Pointer(), PointerFromQKeySequence(other))
	}
}

func (ptr *QKeySequence) ToString(format QKeySequence__SequenceFormat) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QKeySequence_ToString(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QKeySequence) DestroyQKeySequence() {
	if ptr.Pointer() != nil {
		C.QKeySequence_DestroyQKeySequence(ptr.Pointer())
	}
}
