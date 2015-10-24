package gui

//#include "qtextoption.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextOption struct {
	ptr unsafe.Pointer
}

type QTextOptionITF interface {
	QTextOptionPTR() *QTextOption
}

func (p *QTextOption) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextOption) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextOption(ptr QTextOptionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextOptionPTR().Pointer()
	}
	return nil
}

func QTextOptionFromPointer(ptr unsafe.Pointer) *QTextOption {
	var n = new(QTextOption)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextOption) QTextOptionPTR() *QTextOption {
	return ptr
}

//QTextOption::Flag
type QTextOption__Flag int

var (
	QTextOption__ShowTabsAndSpaces                     = QTextOption__Flag(0x1)
	QTextOption__ShowLineAndParagraphSeparators        = QTextOption__Flag(0x2)
	QTextOption__AddSpaceForLineAndParagraphSeparators = QTextOption__Flag(0x4)
	QTextOption__SuppressColors                        = QTextOption__Flag(0x8)
	QTextOption__IncludeTrailingSpaces                 = QTextOption__Flag(0x80000000)
)

//QTextOption::TabType
type QTextOption__TabType int

var (
	QTextOption__LeftTab      = QTextOption__TabType(0)
	QTextOption__RightTab     = QTextOption__TabType(1)
	QTextOption__CenterTab    = QTextOption__TabType(2)
	QTextOption__DelimiterTab = QTextOption__TabType(3)
)

//QTextOption::WrapMode
type QTextOption__WrapMode int

var (
	QTextOption__NoWrap                       = QTextOption__WrapMode(0)
	QTextOption__WordWrap                     = QTextOption__WrapMode(1)
	QTextOption__ManualWrap                   = QTextOption__WrapMode(2)
	QTextOption__WrapAnywhere                 = QTextOption__WrapMode(3)
	QTextOption__WrapAtWordBoundaryOrAnywhere = QTextOption__WrapMode(4)
)

func NewQTextOption3(other QTextOptionITF) *QTextOption {
	return QTextOptionFromPointer(unsafe.Pointer(C.QTextOption_NewQTextOption3(C.QtObjectPtr(PointerFromQTextOption(other)))))
}

func NewQTextOption() *QTextOption {
	return QTextOptionFromPointer(unsafe.Pointer(C.QTextOption_NewQTextOption()))
}

func NewQTextOption2(alignment core.Qt__AlignmentFlag) *QTextOption {
	return QTextOptionFromPointer(unsafe.Pointer(C.QTextOption_NewQTextOption2(C.int(alignment))))
}

func (ptr *QTextOption) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QTextOption_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextOption) Flags() QTextOption__Flag {
	if ptr.Pointer() != nil {
		return QTextOption__Flag(C.QTextOption_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextOption) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QTextOption) SetFlags(flags QTextOption__Flag) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QTextOption) SetTextDirection(direction core.Qt__LayoutDirection) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetTextDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QTextOption) SetUseDesignMetrics(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetUseDesignMetrics(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextOption) SetWrapMode(mode QTextOption__WrapMode) {
	if ptr.Pointer() != nil {
		C.QTextOption_SetWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QTextOption) TextDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextOption_TextDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextOption) UseDesignMetrics() bool {
	if ptr.Pointer() != nil {
		return C.QTextOption_UseDesignMetrics(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextOption) WrapMode() QTextOption__WrapMode {
	if ptr.Pointer() != nil {
		return QTextOption__WrapMode(C.QTextOption_WrapMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextOption) DestroyQTextOption() {
	if ptr.Pointer() != nil {
		C.QTextOption_DestroyQTextOption(C.QtObjectPtr(ptr.Pointer()))
	}
}
