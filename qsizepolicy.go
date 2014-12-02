package qt

//#include "qsizepolicy.h"
import "C"

type qsizepolicy struct {
	ptr C.QtObjectPtr
}

type QSizePolicy interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	ExpandingDirections() Orientation
	HasHeightForWidth() bool
	HasWidthForHeight() bool
	HorizontalStretch() int
	RetainSizeWhenHidden() bool
	SetHeightForWidth_Bool(dependent bool)
	SetHorizontalStretch_Int(stretchFactor int)
	SetRetainSizeWhenHidden_Bool(retainSize bool)
	SetVerticalStretch_Int(stretchFactor int)
	SetWidthForHeight_Bool(dependent bool)
	Transpose()
	VerticalStretch() int
}

func (p *qsizepolicy) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qsizepolicy) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSizePolicy() QSizePolicy {
	var qsizepolicy = new(qsizepolicy)
	qsizepolicy.SetPointer(C.QSizePolicy_New())
	return qsizepolicy
}

func (p *qsizepolicy) ExpandingDirections() Orientation {
	if p.Pointer() == nil {
		return 0
	} else {
		return Orientation(C.QSizePolicy_ExpandingDirections(p.Pointer()))
	}
}

func (p *qsizepolicy) HasHeightForWidth() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSizePolicy_HasHeightForWidth(p.Pointer()) != 0
	}
}

func (p *qsizepolicy) HasWidthForHeight() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSizePolicy_HasWidthForHeight(p.Pointer()) != 0
	}
}

func (p *qsizepolicy) HorizontalStretch() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSizePolicy_HorizontalStretch(p.Pointer()))
	}
}

func (p *qsizepolicy) RetainSizeWhenHidden() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSizePolicy_RetainSizeWhenHidden(p.Pointer()) != 0
	}
}

func (p *qsizepolicy) SetHeightForWidth_Bool(dependent bool) {
	if p.Pointer() != nil {
		C.QSizePolicy_SetHeightForWidth_Bool(p.Pointer(), goBoolToCInt(dependent))
	}
}

func (p *qsizepolicy) SetHorizontalStretch_Int(stretchFactor int) {
	if p.Pointer() != nil {
		C.QSizePolicy_SetHorizontalStretch_Int(p.Pointer(), C.int(stretchFactor))
	}
}

func (p *qsizepolicy) SetRetainSizeWhenHidden_Bool(retainSize bool) {
	if p.Pointer() != nil {
		C.QSizePolicy_SetRetainSizeWhenHidden_Bool(p.Pointer(), goBoolToCInt(retainSize))
	}
}

func (p *qsizepolicy) SetVerticalStretch_Int(stretchFactor int) {
	if p.Pointer() != nil {
		C.QSizePolicy_SetVerticalStretch_Int(p.Pointer(), C.int(stretchFactor))
	}
}

func (p *qsizepolicy) SetWidthForHeight_Bool(dependent bool) {
	if p.Pointer() != nil {
		C.QSizePolicy_SetWidthForHeight_Bool(p.Pointer(), goBoolToCInt(dependent))
	}
}

func (p *qsizepolicy) Transpose() {
	if p.Pointer() != nil {
		C.QSizePolicy_Transpose(p.Pointer())
	}
}

func (p *qsizepolicy) VerticalStretch() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSizePolicy_VerticalStretch(p.Pointer()))
	}
}
