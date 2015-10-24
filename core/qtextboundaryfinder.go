package core

//#include "qtextboundaryfinder.h"
import "C"
import (
	"unsafe"
)

type QTextBoundaryFinder struct {
	ptr unsafe.Pointer
}

type QTextBoundaryFinderITF interface {
	QTextBoundaryFinderPTR() *QTextBoundaryFinder
}

func (p *QTextBoundaryFinder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBoundaryFinder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBoundaryFinder(ptr QTextBoundaryFinderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBoundaryFinderPTR().Pointer()
	}
	return nil
}

func QTextBoundaryFinderFromPointer(ptr unsafe.Pointer) *QTextBoundaryFinder {
	var n = new(QTextBoundaryFinder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBoundaryFinder) QTextBoundaryFinderPTR() *QTextBoundaryFinder {
	return ptr
}

//QTextBoundaryFinder::BoundaryReason
type QTextBoundaryFinder__BoundaryReason int

var (
	QTextBoundaryFinder__NotAtBoundary    = QTextBoundaryFinder__BoundaryReason(0)
	QTextBoundaryFinder__BreakOpportunity = QTextBoundaryFinder__BoundaryReason(0x1f)
	QTextBoundaryFinder__StartOfItem      = QTextBoundaryFinder__BoundaryReason(0x20)
	QTextBoundaryFinder__EndOfItem        = QTextBoundaryFinder__BoundaryReason(0x40)
	QTextBoundaryFinder__MandatoryBreak   = QTextBoundaryFinder__BoundaryReason(0x80)
	QTextBoundaryFinder__SoftHyphen       = QTextBoundaryFinder__BoundaryReason(0x100)
)

//QTextBoundaryFinder::BoundaryType
type QTextBoundaryFinder__BoundaryType int

var (
	QTextBoundaryFinder__Grapheme = QTextBoundaryFinder__BoundaryType(0)
	QTextBoundaryFinder__Word     = QTextBoundaryFinder__BoundaryType(1)
	QTextBoundaryFinder__Sentence = QTextBoundaryFinder__BoundaryType(2)
	QTextBoundaryFinder__Line     = QTextBoundaryFinder__BoundaryType(3)
)

func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	return QTextBoundaryFinderFromPointer(unsafe.Pointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder()))
}

func NewQTextBoundaryFinder3(ty QTextBoundaryFinder__BoundaryType, stri string) *QTextBoundaryFinder {
	return QTextBoundaryFinderFromPointer(unsafe.Pointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder3(C.int(ty), C.CString(stri))))
}

func NewQTextBoundaryFinder2(other QTextBoundaryFinderITF) *QTextBoundaryFinder {
	return QTextBoundaryFinderFromPointer(unsafe.Pointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder2(C.QtObjectPtr(PointerFromQTextBoundaryFinder(other)))))
}

func (ptr *QTextBoundaryFinder) BoundaryReasons() QTextBoundaryFinder__BoundaryReason {
	if ptr.Pointer() != nil {
		return QTextBoundaryFinder__BoundaryReason(C.QTextBoundaryFinder_BoundaryReasons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) IsAtBoundary() bool {
	if ptr.Pointer() != nil {
		return C.QTextBoundaryFinder_IsAtBoundary(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBoundaryFinder) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextBoundaryFinder_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextBoundaryFinder) Position() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_Position(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) SetPosition(position int) {
	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.int(position))
	}
}

func (ptr *QTextBoundaryFinder) String() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBoundaryFinder_String(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextBoundaryFinder) ToEnd() {
	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_ToEnd(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBoundaryFinder) ToNextBoundary() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_ToNextBoundary(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) ToPreviousBoundary() int {
	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_ToPreviousBoundary(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) ToStart() {
	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_ToStart(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextBoundaryFinder) Type() QTextBoundaryFinder__BoundaryType {
	if ptr.Pointer() != nil {
		return QTextBoundaryFinder__BoundaryType(C.QTextBoundaryFinder_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) DestroyQTextBoundaryFinder() {
	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_DestroyQTextBoundaryFinder(C.QtObjectPtr(ptr.Pointer()))
	}
}
