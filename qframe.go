package qt

//#include "qframe.h"
import "C"

type qframe struct {
	qwidget
}

type QFrame interface {
	QWidget
	FrameStyle() int
	FrameWidth() int
	LineWidth() int
	MidLineWidth() int
	SetFrameStyle(style int)
	SetLineWidth(width int)
	SetMidLineWidth(width int)
}

func (p *qframe) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qframe) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQFrame(parent QWidget, f WindowType) QFrame {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qframe = new(qframe)
	qframe.SetPointer(C.QFrame_New_QWidget_WindowType(parentPtr, C.int(f)))
	qframe.SetObjectName("QFrame_" + randomIdentifier())
	return qframe
}

func (p *qframe) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QFrame_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qframe) FrameStyle() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFrame_FrameStyle(p.Pointer()))
}

func (p *qframe) FrameWidth() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFrame_FrameWidth(p.Pointer()))
}

func (p *qframe) LineWidth() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFrame_LineWidth(p.Pointer()))
}

func (p *qframe) MidLineWidth() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFrame_MidLineWidth(p.Pointer()))
}

func (p *qframe) SetFrameStyle(style int) {
	if p.Pointer() != nil {
		C.QFrame_SetFrameStyle_Int(p.Pointer(), C.int(style))
	}
}

func (p *qframe) SetLineWidth(width int) {
	if p.Pointer() != nil {
		C.QFrame_SetLineWidth_Int(p.Pointer(), C.int(width))
	}
}

func (p *qframe) SetMidLineWidth(width int) {
	if p.Pointer() != nil {
		C.QFrame_SetMidLineWidth_Int(p.Pointer(), C.int(width))
	}
}
