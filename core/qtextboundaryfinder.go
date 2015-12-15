package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTextBoundaryFinder struct {
	ptr unsafe.Pointer
}

type QTextBoundaryFinder_ITF interface {
	QTextBoundaryFinder_PTR() *QTextBoundaryFinder
}

func (p *QTextBoundaryFinder) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextBoundaryFinder) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextBoundaryFinder(ptr QTextBoundaryFinder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextBoundaryFinder_PTR().Pointer()
	}
	return nil
}

func NewQTextBoundaryFinderFromPointer(ptr unsafe.Pointer) *QTextBoundaryFinder {
	var n = new(QTextBoundaryFinder)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextBoundaryFinder) QTextBoundaryFinder_PTR() *QTextBoundaryFinder {
	return ptr
}

//QTextBoundaryFinder::BoundaryReason
type QTextBoundaryFinder__BoundaryReason int64

const (
	QTextBoundaryFinder__NotAtBoundary    = QTextBoundaryFinder__BoundaryReason(0)
	QTextBoundaryFinder__BreakOpportunity = QTextBoundaryFinder__BoundaryReason(0x1f)
	QTextBoundaryFinder__StartOfItem      = QTextBoundaryFinder__BoundaryReason(0x20)
	QTextBoundaryFinder__EndOfItem        = QTextBoundaryFinder__BoundaryReason(0x40)
	QTextBoundaryFinder__MandatoryBreak   = QTextBoundaryFinder__BoundaryReason(0x80)
	QTextBoundaryFinder__SoftHyphen       = QTextBoundaryFinder__BoundaryReason(0x100)
)

//QTextBoundaryFinder::BoundaryType
type QTextBoundaryFinder__BoundaryType int64

const (
	QTextBoundaryFinder__Grapheme = QTextBoundaryFinder__BoundaryType(0)
	QTextBoundaryFinder__Word     = QTextBoundaryFinder__BoundaryType(1)
	QTextBoundaryFinder__Sentence = QTextBoundaryFinder__BoundaryType(2)
	QTextBoundaryFinder__Line     = QTextBoundaryFinder__BoundaryType(3)
)

func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	defer qt.Recovering("QTextBoundaryFinder::QTextBoundaryFinder")

	return NewQTextBoundaryFinderFromPointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder())
}

func NewQTextBoundaryFinder3(ty QTextBoundaryFinder__BoundaryType, stri string) *QTextBoundaryFinder {
	defer qt.Recovering("QTextBoundaryFinder::QTextBoundaryFinder")

	return NewQTextBoundaryFinderFromPointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder3(C.int(ty), C.CString(stri)))
}

func NewQTextBoundaryFinder2(other QTextBoundaryFinder_ITF) *QTextBoundaryFinder {
	defer qt.Recovering("QTextBoundaryFinder::QTextBoundaryFinder")

	return NewQTextBoundaryFinderFromPointer(C.QTextBoundaryFinder_NewQTextBoundaryFinder2(PointerFromQTextBoundaryFinder(other)))
}

func (ptr *QTextBoundaryFinder) BoundaryReasons() QTextBoundaryFinder__BoundaryReason {
	defer qt.Recovering("QTextBoundaryFinder::boundaryReasons")

	if ptr.Pointer() != nil {
		return QTextBoundaryFinder__BoundaryReason(C.QTextBoundaryFinder_BoundaryReasons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) IsAtBoundary() bool {
	defer qt.Recovering("QTextBoundaryFinder::isAtBoundary")

	if ptr.Pointer() != nil {
		return C.QTextBoundaryFinder_IsAtBoundary(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBoundaryFinder) IsValid() bool {
	defer qt.Recovering("QTextBoundaryFinder::isValid")

	if ptr.Pointer() != nil {
		return C.QTextBoundaryFinder_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextBoundaryFinder) Position() int {
	defer qt.Recovering("QTextBoundaryFinder::position")

	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) SetPosition(position int) {
	defer qt.Recovering("QTextBoundaryFinder::setPosition")

	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_SetPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QTextBoundaryFinder) String() string {
	defer qt.Recovering("QTextBoundaryFinder::string")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextBoundaryFinder_String(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextBoundaryFinder) ToEnd() {
	defer qt.Recovering("QTextBoundaryFinder::toEnd")

	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_ToEnd(ptr.Pointer())
	}
}

func (ptr *QTextBoundaryFinder) ToNextBoundary() int {
	defer qt.Recovering("QTextBoundaryFinder::toNextBoundary")

	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_ToNextBoundary(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) ToPreviousBoundary() int {
	defer qt.Recovering("QTextBoundaryFinder::toPreviousBoundary")

	if ptr.Pointer() != nil {
		return int(C.QTextBoundaryFinder_ToPreviousBoundary(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) ToStart() {
	defer qt.Recovering("QTextBoundaryFinder::toStart")

	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_ToStart(ptr.Pointer())
	}
}

func (ptr *QTextBoundaryFinder) Type() QTextBoundaryFinder__BoundaryType {
	defer qt.Recovering("QTextBoundaryFinder::type")

	if ptr.Pointer() != nil {
		return QTextBoundaryFinder__BoundaryType(C.QTextBoundaryFinder_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextBoundaryFinder) DestroyQTextBoundaryFinder() {
	defer qt.Recovering("QTextBoundaryFinder::~QTextBoundaryFinder")

	if ptr.Pointer() != nil {
		C.QTextBoundaryFinder_DestroyQTextBoundaryFinder(ptr.Pointer())
	}
}
