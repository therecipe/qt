package qt

//#include "qlayout.h"
import "C"

type qlayout struct {
	qobject
}

type QLayout interface {
	QObject
	Activate() bool
	AddWidget(w QWidget)
	Count() int
	IndexOf(widget QWidget) int
	IsEnabled() bool
	MenuBar() QWidget
	ParentWidget() QWidget
	RemoveItem(item QLayoutItem)
	RemoveWidget(widget QWidget)
	SetAlignment1(w QWidget, alignment AlignmentFlag) bool
	SetAlignment2(alignment AlignmentFlag)
	SetAlignment3(l QLayout, alignment AlignmentFlag) bool
	SetContentsMargins(left int, top int, right int, bottom int)
	SetEnabled(enable bool)
	SetMenuBar(widget QWidget)
	SetSizeConstraint(Si SizeConstraint)
	SetSpacing(spacing int)
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
	}
	return C.QLayout_Activate(p.Pointer()) != 0
}

func (p *qlayout) AddWidget(w QWidget) {
	if p.Pointer() != nil {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		C.QLayout_AddWidget_QWidget(p.Pointer(), wPtr)
	}
}

func (p *qlayout) Count() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLayout_Count(p.Pointer()))
}

func (p *qlayout) IndexOf(widget QWidget) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return int(C.QLayout_IndexOf_QWidget(p.Pointer(), widgetPtr))
	}
}

func (p *qlayout) IsEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLayout_IsEnabled(p.Pointer()) != 0
}

func (p *qlayout) MenuBar() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLayout_MenuBar(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
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
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlayout) RemoveItem(item QLayoutItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QLayout_RemoveItem_QLayoutItem(p.Pointer(), itemPtr)
	}
}

func (p *qlayout) RemoveWidget(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QLayout_RemoveWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qlayout) SetAlignment1(w QWidget, alignment AlignmentFlag) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		return C.QLayout_SetAlignment_QWidget_AlignmentFlag(p.Pointer(), wPtr, C.int(alignment)) != 0
	}
}

func (p *qlayout) SetAlignment2(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLayout_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qlayout) SetAlignment3(l QLayout, alignment AlignmentFlag) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var lPtr C.QtObjectPtr
		if l != nil {
			lPtr = l.Pointer()
		}
		return C.QLayout_SetAlignment_QLayout_AlignmentFlag(p.Pointer(), lPtr, C.int(alignment)) != 0
	}
}

func (p *qlayout) SetContentsMargins(left int, top int, right int, bottom int) {
	if p.Pointer() != nil {
		C.QLayout_SetContentsMargins_Int_Int_Int_Int(p.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (p *qlayout) SetEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QLayout_SetEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlayout) SetMenuBar(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QLayout_SetMenuBar_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qlayout) SetSizeConstraint(Si SizeConstraint) {
	if p.Pointer() != nil {
		C.QLayout_SetSizeConstraint_SizeConstraint(p.Pointer(), C.int(Si))
	}
}

func (p *qlayout) SetSpacing(spacing int) {
	if p.Pointer() != nil {
		C.QLayout_SetSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qlayout) SizeConstraint() SizeConstraint {
	if p.Pointer() == nil {
		return 0
	}
	return SizeConstraint(C.QLayout_SizeConstraint(p.Pointer()))
}

func (p *qlayout) Spacing() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLayout_Spacing(p.Pointer()))
}

func (p *qlayout) Update() {
	if p.Pointer() != nil {
		C.QLayout_Update(p.Pointer())
	}
}
