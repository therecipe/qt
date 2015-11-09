package widgets

//#include "qsizepolicy.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSizePolicy struct {
	ptr unsafe.Pointer
}

type QSizePolicy_ITF interface {
	QSizePolicy_PTR() *QSizePolicy
}

func (p *QSizePolicy) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSizePolicy) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSizePolicy(ptr QSizePolicy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizePolicy_PTR().Pointer()
	}
	return nil
}

func NewQSizePolicyFromPointer(ptr unsafe.Pointer) *QSizePolicy {
	var n = new(QSizePolicy)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSizePolicy) QSizePolicy_PTR() *QSizePolicy {
	return ptr
}

//QSizePolicy::ControlType
type QSizePolicy__ControlType int64

const (
	QSizePolicy__DefaultType = QSizePolicy__ControlType(0x00000001)
	QSizePolicy__ButtonBox   = QSizePolicy__ControlType(0x00000002)
	QSizePolicy__CheckBox    = QSizePolicy__ControlType(0x00000004)
	QSizePolicy__ComboBox    = QSizePolicy__ControlType(0x00000008)
	QSizePolicy__Frame       = QSizePolicy__ControlType(0x00000010)
	QSizePolicy__GroupBox    = QSizePolicy__ControlType(0x00000020)
	QSizePolicy__Label       = QSizePolicy__ControlType(0x00000040)
	QSizePolicy__Line        = QSizePolicy__ControlType(0x00000080)
	QSizePolicy__LineEdit    = QSizePolicy__ControlType(0x00000100)
	QSizePolicy__PushButton  = QSizePolicy__ControlType(0x00000200)
	QSizePolicy__RadioButton = QSizePolicy__ControlType(0x00000400)
	QSizePolicy__Slider      = QSizePolicy__ControlType(0x00000800)
	QSizePolicy__SpinBox     = QSizePolicy__ControlType(0x00001000)
	QSizePolicy__TabWidget   = QSizePolicy__ControlType(0x00002000)
	QSizePolicy__ToolButton  = QSizePolicy__ControlType(0x00004000)
)

//QSizePolicy::Policy
type QSizePolicy__Policy int64

const (
	QSizePolicy__Fixed            = QSizePolicy__Policy(0)
	QSizePolicy__Minimum          = QSizePolicy__Policy(QSizePolicy__GrowFlag)
	QSizePolicy__Maximum          = QSizePolicy__Policy(QSizePolicy__ShrinkFlag)
	QSizePolicy__Preferred        = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag)
	QSizePolicy__MinimumExpanding = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Expanding        = QSizePolicy__Policy(QSizePolicy__GrowFlag | QSizePolicy__ShrinkFlag | QSizePolicy__ExpandFlag)
	QSizePolicy__Ignored          = QSizePolicy__Policy(QSizePolicy__ShrinkFlag | QSizePolicy__GrowFlag | QSizePolicy__IgnoreFlag)
)

//QSizePolicy::PolicyFlag
type QSizePolicy__PolicyFlag int64

const (
	QSizePolicy__GrowFlag   = QSizePolicy__PolicyFlag(1)
	QSizePolicy__ExpandFlag = QSizePolicy__PolicyFlag(2)
	QSizePolicy__ShrinkFlag = QSizePolicy__PolicyFlag(4)
	QSizePolicy__IgnoreFlag = QSizePolicy__PolicyFlag(8)
)

func NewQSizePolicy() *QSizePolicy {
	return NewQSizePolicyFromPointer(C.QSizePolicy_NewQSizePolicy())
}

func NewQSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy, ty QSizePolicy__ControlType) *QSizePolicy {
	return NewQSizePolicyFromPointer(C.QSizePolicy_NewQSizePolicy2(C.int(horizontal), C.int(vertical), C.int(ty)))
}

func (ptr *QSizePolicy) ControlType() QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QSizePolicy_ControlType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) ExpandingDirections() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSizePolicy_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QSizePolicy_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizePolicy) HasWidthForHeight() bool {
	if ptr.Pointer() != nil {
		return C.QSizePolicy_HasWidthForHeight(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizePolicy) HorizontalPolicy() QSizePolicy__Policy {
	if ptr.Pointer() != nil {
		return QSizePolicy__Policy(C.QSizePolicy_HorizontalPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) HorizontalStretch() int {
	if ptr.Pointer() != nil {
		return int(C.QSizePolicy_HorizontalStretch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) RetainSizeWhenHidden() bool {
	if ptr.Pointer() != nil {
		return C.QSizePolicy_RetainSizeWhenHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizePolicy) SetControlType(ty QSizePolicy__ControlType) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetControlType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QSizePolicy) SetHeightForWidth(dependent bool) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetHeightForWidth(ptr.Pointer(), C.int(qt.GoBoolToInt(dependent)))
	}
}

func (ptr *QSizePolicy) SetHorizontalPolicy(policy QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetHorizontalPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetHorizontalStretch(ptr.Pointer(), C.int(stretchFactor))
	}
}

func (ptr *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetRetainSizeWhenHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(retainSize)))
	}
}

func (ptr *QSizePolicy) SetVerticalPolicy(policy QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetVerticalPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetVerticalStretch(ptr.Pointer(), C.int(stretchFactor))
	}
}

func (ptr *QSizePolicy) SetWidthForHeight(dependent bool) {
	if ptr.Pointer() != nil {
		C.QSizePolicy_SetWidthForHeight(ptr.Pointer(), C.int(qt.GoBoolToInt(dependent)))
	}
}

func (ptr *QSizePolicy) Transpose() {
	if ptr.Pointer() != nil {
		C.QSizePolicy_Transpose(ptr.Pointer())
	}
}

func (ptr *QSizePolicy) VerticalPolicy() QSizePolicy__Policy {
	if ptr.Pointer() != nil {
		return QSizePolicy__Policy(C.QSizePolicy_VerticalPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizePolicy) VerticalStretch() int {
	if ptr.Pointer() != nil {
		return int(C.QSizePolicy_VerticalStretch(ptr.Pointer()))
	}
	return 0
}
