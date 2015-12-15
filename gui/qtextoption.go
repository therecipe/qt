package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextOption struct {
	ptr unsafe.Pointer
}

type QTextOption_ITF interface {
	QTextOption_PTR() *QTextOption
}

func (p *QTextOption) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextOption) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextOption(ptr QTextOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextOption_PTR().Pointer()
	}
	return nil
}

func NewQTextOptionFromPointer(ptr unsafe.Pointer) *QTextOption {
	var n = new(QTextOption)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextOption) QTextOption_PTR() *QTextOption {
	return ptr
}

//QTextOption::Flag
type QTextOption__Flag int64

const (
	QTextOption__ShowTabsAndSpaces                     = QTextOption__Flag(0x1)
	QTextOption__ShowLineAndParagraphSeparators        = QTextOption__Flag(0x2)
	QTextOption__AddSpaceForLineAndParagraphSeparators = QTextOption__Flag(0x4)
	QTextOption__SuppressColors                        = QTextOption__Flag(0x8)
	QTextOption__IncludeTrailingSpaces                 = QTextOption__Flag(0x80000000)
)

//QTextOption::TabType
type QTextOption__TabType int64

const (
	QTextOption__LeftTab      = QTextOption__TabType(0)
	QTextOption__RightTab     = QTextOption__TabType(1)
	QTextOption__CenterTab    = QTextOption__TabType(2)
	QTextOption__DelimiterTab = QTextOption__TabType(3)
)

//QTextOption::WrapMode
type QTextOption__WrapMode int64

const (
	QTextOption__NoWrap                       = QTextOption__WrapMode(0)
	QTextOption__WordWrap                     = QTextOption__WrapMode(1)
	QTextOption__ManualWrap                   = QTextOption__WrapMode(2)
	QTextOption__WrapAnywhere                 = QTextOption__WrapMode(3)
	QTextOption__WrapAtWordBoundaryOrAnywhere = QTextOption__WrapMode(4)
)

func NewQTextOption3(other QTextOption_ITF) *QTextOption {
	defer qt.Recovering("QTextOption::QTextOption")

	return NewQTextOptionFromPointer(C.QTextOption_NewQTextOption3(PointerFromQTextOption(other)))
}

func NewQTextOption() *QTextOption {
	defer qt.Recovering("QTextOption::QTextOption")

	return NewQTextOptionFromPointer(C.QTextOption_NewQTextOption())
}

func NewQTextOption2(alignment core.Qt__AlignmentFlag) *QTextOption {
	defer qt.Recovering("QTextOption::QTextOption")

	return NewQTextOptionFromPointer(C.QTextOption_NewQTextOption2(C.int(alignment)))
}

func (ptr *QTextOption) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QTextOption::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextOption_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) Flags() QTextOption__Flag {
	defer qt.Recovering("QTextOption::flags")

	if ptr.Pointer() != nil {
		return QTextOption__Flag(C.QTextOption_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QTextOption::setAlignment")

	if ptr.Pointer() != nil {
		C.QTextOption_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QTextOption) SetFlags(flags QTextOption__Flag) {
	defer qt.Recovering("QTextOption::setFlags")

	if ptr.Pointer() != nil {
		C.QTextOption_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QTextOption) SetTabStop(tabStop float64) {
	defer qt.Recovering("QTextOption::setTabStop")

	if ptr.Pointer() != nil {
		C.QTextOption_SetTabStop(ptr.Pointer(), C.double(tabStop))
	}
}

func (ptr *QTextOption) SetTextDirection(direction core.Qt__LayoutDirection) {
	defer qt.Recovering("QTextOption::setTextDirection")

	if ptr.Pointer() != nil {
		C.QTextOption_SetTextDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QTextOption) SetUseDesignMetrics(enable bool) {
	defer qt.Recovering("QTextOption::setUseDesignMetrics")

	if ptr.Pointer() != nil {
		C.QTextOption_SetUseDesignMetrics(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextOption) SetWrapMode(mode QTextOption__WrapMode) {
	defer qt.Recovering("QTextOption::setWrapMode")

	if ptr.Pointer() != nil {
		C.QTextOption_SetWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QTextOption) TabStop() float64 {
	defer qt.Recovering("QTextOption::tabStop")

	if ptr.Pointer() != nil {
		return float64(C.QTextOption_TabStop(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) TextDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QTextOption::textDirection")

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextOption_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) UseDesignMetrics() bool {
	defer qt.Recovering("QTextOption::useDesignMetrics")

	if ptr.Pointer() != nil {
		return C.QTextOption_UseDesignMetrics(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextOption) WrapMode() QTextOption__WrapMode {
	defer qt.Recovering("QTextOption::wrapMode")

	if ptr.Pointer() != nil {
		return QTextOption__WrapMode(C.QTextOption_WrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextOption) DestroyQTextOption() {
	defer qt.Recovering("QTextOption::~QTextOption")

	if ptr.Pointer() != nil {
		C.QTextOption_DestroyQTextOption(ptr.Pointer())
	}
}
