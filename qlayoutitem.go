package qt

//#include "qlayoutitem.h"
import "C"

type qlayoutitem struct {
	ptr C.QtObjectPtr
}

type QLayoutItem interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Destroy()
	Alignment() AlignmentFlag
	ExpandingDirections() Orientation
	HasHeightForWidth() bool
	HeightForWidth(w int) int
	Invalidate()
	IsEmpty() bool
	Layout() QLayout
	MinimumHeightForWidth(w int) int
	SetAlignment(alignment AlignmentFlag)
	Widget() QWidget
}

func (p *qlayoutitem) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlayoutitem) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qlayoutitem) Destroy() {
	if p.Pointer() != nil {
		C.QLayoutItem_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlayoutitem) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QLayoutItem_Alignment(p.Pointer()))
}

func (p *qlayoutitem) ExpandingDirections() Orientation {
	if p.Pointer() == nil {
		return 0
	}
	return Orientation(C.QLayoutItem_ExpandingDirections(p.Pointer()))
}

func (p *qlayoutitem) HasHeightForWidth() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLayoutItem_HasHeightForWidth(p.Pointer()) != 0
}

func (p *qlayoutitem) HeightForWidth(w int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLayoutItem_HeightForWidth_Int(p.Pointer(), C.int(w)))
}

func (p *qlayoutitem) Invalidate() {
	if p.Pointer() != nil {
		C.QLayoutItem_Invalidate(p.Pointer())
	}
}

func (p *qlayoutitem) IsEmpty() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLayoutItem_IsEmpty(p.Pointer()) != 0
}

func (p *qlayoutitem) Layout() QLayout {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlayout = new(qlayout)
		qlayout.SetPointer(C.QLayoutItem_Layout(p.Pointer()))
		if qlayout.ObjectName() == "" {
			qlayout.SetObjectName("QLayout_" + randomIdentifier())
		}
		return qlayout
	}
}

func (p *qlayoutitem) MinimumHeightForWidth(w int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLayoutItem_MinimumHeightForWidth_Int(p.Pointer(), C.int(w)))
}

func (p *qlayoutitem) SetAlignment(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLayoutItem_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qlayoutitem) Widget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLayoutItem_Widget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}
