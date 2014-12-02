package qt

//#include "qlayout.h"
import "C"

type qlayout struct {
	qobject
}

type QLayout interface {
	QObject
	Activate() bool
	AddItem_QLayoutItem(item QLayoutItem)
	AddWidget_QWidget(w QWidget)
	Count() int
	IndexOf_QWidget(widget QWidget) int
	IsEnabled() bool
	MenuBar() QWidget
	ParentWidget() QWidget
	RemoveItem_QLayoutItem(item QLayoutItem)
	RemoveWidget_QWidget(widget QWidget)
	SetAlignment_QWidget_AlignmentFlag(w QWidget, alignment AlignmentFlag) bool
	SetAlignment_AlignmentFlag(alignment AlignmentFlag)
	SetAlignment_QLayout_AlignmentFlag(l QLayout, alignment AlignmentFlag) bool
	SetContentsMargins_Int_Int_Int_Int(left int, top int, right int, bottom int)
	SetEnabled_Bool(enable bool)
	SetMenuBar_QWidget(widget QWidget)
	SetSizeConstraint_SizeConstraint(Si SizeConstraint)
	SetSpacing_Int(spacing int)
	SizeConstraint() SizeConstraint
	Spacing() int
	Update()
}

func (p *qlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//SizeConstraint
type SizeConstraint int

var (
	SETDEFAULTCONSTRAINT = SizeConstraint(C.QLayout_SetDefaultConstraint())
	SETFIXEDSIZE         = SizeConstraint(C.QLayout_SetFixedSize())
	SETMINIMUMSIZE       = SizeConstraint(C.QLayout_SetMinimumSize())
	SETMAXIMUMSIZE       = SizeConstraint(C.QLayout_SetMaximumSize())
	SETMINANDMAXSIZE     = SizeConstraint(C.QLayout_SetMinAndMaxSize())
	SETNOCONSTRAINT      = SizeConstraint(C.QLayout_SetNoConstraint())
)

func (p *qlayout) Activate() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLayout_Activate(p.Pointer()) != 0
	}
}

func (p *qlayout) AddItem_QLayoutItem(item QLayoutItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QLayout_AddItem_QLayoutItem(p.Pointer(), itemPtr)
	}
}

func (p *qlayout) AddWidget_QWidget(w QWidget) {
	if p.Pointer() == nil {
	} else {
		var wPtr C.QtObjectPtr = nil
		if w != nil {
			wPtr = w.Pointer()
		}
		C.QLayout_AddWidget_QWidget(p.Pointer(), wPtr)
	}
}

func (p *qlayout) Count() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLayout_Count(p.Pointer()))
	}
}

func (p *qlayout) IndexOf_QWidget(widget QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QLayout_IndexOf_QWidget(p.Pointer(), widgetPtr))
	}
}

func (p *qlayout) IsEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLayout_IsEnabled(p.Pointer()) != 0
	}
}

func (p *qlayout) MenuBar() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLayout_MenuBar(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlayout) ParentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLayout_ParentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlayout) RemoveItem_QLayoutItem(item QLayoutItem) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QLayout_RemoveItem_QLayoutItem(p.Pointer(), itemPtr)
	}
}

func (p *qlayout) RemoveWidget_QWidget(widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QLayout_RemoveWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qlayout) SetAlignment_QWidget_AlignmentFlag(w QWidget, alignment AlignmentFlag) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var wPtr C.QtObjectPtr = nil
		if w != nil {
			wPtr = w.Pointer()
		}
		return C.QLayout_SetAlignment_QWidget_AlignmentFlag(p.Pointer(), wPtr, C.int(alignment)) != 0
	}
}

func (p *qlayout) SetAlignment_AlignmentFlag(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLayout_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qlayout) SetAlignment_QLayout_AlignmentFlag(l QLayout, alignment AlignmentFlag) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var lPtr C.QtObjectPtr = nil
		if l != nil {
			lPtr = l.Pointer()
		}
		return C.QLayout_SetAlignment_QLayout_AlignmentFlag(p.Pointer(), lPtr, C.int(alignment)) != 0
	}
}

func (p *qlayout) SetContentsMargins_Int_Int_Int_Int(left int, top int, right int, bottom int) {
	if p.Pointer() != nil {
		C.QLayout_SetContentsMargins_Int_Int_Int_Int(p.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (p *qlayout) SetEnabled_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QLayout_SetEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlayout) SetMenuBar_QWidget(widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QLayout_SetMenuBar_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qlayout) SetSizeConstraint_SizeConstraint(Si SizeConstraint) {
	if p.Pointer() != nil {
		C.QLayout_SetSizeConstraint_SizeConstraint(p.Pointer(), C.int(Si))
	}
}

func (p *qlayout) SetSpacing_Int(spacing int) {
	if p.Pointer() != nil {
		C.QLayout_SetSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qlayout) SizeConstraint() SizeConstraint {
	if p.Pointer() == nil {
		return 0
	} else {
		return SizeConstraint(C.QLayout_SizeConstraint(p.Pointer()))
	}
}

func (p *qlayout) Spacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLayout_Spacing(p.Pointer()))
	}
}

func (p *qlayout) Update() {
	if p.Pointer() != nil {
		C.QLayout_Update(p.Pointer())
	}
}
