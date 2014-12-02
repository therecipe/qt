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
	HeightForWidth_Int(w int) int
	Invalidate()
	IsEmpty() bool
	Layout() QLayout
	MinimumHeightForWidth_Int(w int) int
	SetAlignment_AlignmentFlag(alignment AlignmentFlag)
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
	} else {
		return AlignmentFlag(C.QLayoutItem_Alignment(p.Pointer()))
	}
}

func (p *qlayoutitem) ExpandingDirections() Orientation {
	if p.Pointer() == nil {
		return 0
	} else {
		return Orientation(C.QLayoutItem_ExpandingDirections(p.Pointer()))
	}
}

func (p *qlayoutitem) HasHeightForWidth() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLayoutItem_HasHeightForWidth(p.Pointer()) != 0
	}
}

func (p *qlayoutitem) HeightForWidth_Int(w int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLayoutItem_HeightForWidth_Int(p.Pointer(), C.int(w)))
	}
}

func (p *qlayoutitem) Invalidate() {
	if p.Pointer() != nil {
		C.QLayoutItem_Invalidate(p.Pointer())
	}
}

func (p *qlayoutitem) IsEmpty() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLayoutItem_IsEmpty(p.Pointer()) != 0
	}
}

func (p *qlayoutitem) Layout() QLayout {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlayout = new(qlayout)
		qlayout.SetPointer(C.QLayoutItem_Layout(p.Pointer()))
		if qlayout.ObjectName() == "" {
			qlayout.SetObjectName_String("QLayout_" + randomIdentifier())
		}
		return qlayout
	}
}

func (p *qlayoutitem) MinimumHeightForWidth_Int(w int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLayoutItem_MinimumHeightForWidth_Int(p.Pointer(), C.int(w)))
	}
}

func (p *qlayoutitem) SetAlignment_AlignmentFlag(alignment AlignmentFlag) {
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
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}
