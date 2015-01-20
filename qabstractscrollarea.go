package qt

//#include "qabstractscrollarea.h"
import "C"

type qabstractscrollarea struct {
	qframe
}

type QAbstractScrollArea interface {
	QFrame
	AddScrollBarWidget(widget QWidget, alignment AlignmentFlag)
	CornerWidget() QWidget
	HorizontalScrollBarPolicy() ScrollBarPolicy
	SetCornerWidget(widget QWidget)
	SetHorizontalScrollBarPolicy(ScrollBarPolicy ScrollBarPolicy)
	SetVerticalScrollBarPolicy(ScrollBarPolicy ScrollBarPolicy)
	SetViewport(widget QWidget)
	SetupViewport(viewport QWidget)
	VerticalScrollBarPolicy() ScrollBarPolicy
	Viewport() QWidget
}

func (p *qabstractscrollarea) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qabstractscrollarea) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQAbstractScrollArea(parent QWidget) QAbstractScrollArea {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qabstractscrollarea = new(qabstractscrollarea)
	qabstractscrollarea.SetPointer(C.QAbstractScrollArea_New_QWidget(parentPtr))
	qabstractscrollarea.SetObjectName("QAbstractScrollArea_" + randomIdentifier())
	return qabstractscrollarea
}

func (p *qabstractscrollarea) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAbstractScrollArea_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qabstractscrollarea) AddScrollBarWidget(widget QWidget, alignment AlignmentFlag) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QAbstractScrollArea_AddScrollBarWidget_QWidget_AlignmentFlag(p.Pointer(), widgetPtr, C.int(alignment))
	}
}

func (p *qabstractscrollarea) CornerWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QAbstractScrollArea_CornerWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qabstractscrollarea) HorizontalScrollBarPolicy() ScrollBarPolicy {
	if p.Pointer() == nil {
		return 0
	}
	return ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(p.Pointer()))
}

func (p *qabstractscrollarea) SetCornerWidget(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QAbstractScrollArea_SetCornerWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qabstractscrollarea) SetHorizontalScrollBarPolicy(ScrollBarPolicy ScrollBarPolicy) {
	if p.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy_ScrollBarPolicy(p.Pointer(), C.int(ScrollBarPolicy))
	}
}

func (p *qabstractscrollarea) SetVerticalScrollBarPolicy(ScrollBarPolicy ScrollBarPolicy) {
	if p.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy_ScrollBarPolicy(p.Pointer(), C.int(ScrollBarPolicy))
	}
}

func (p *qabstractscrollarea) SetViewport(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QAbstractScrollArea_SetViewport_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qabstractscrollarea) SetupViewport(viewport QWidget) {
	if p.Pointer() != nil {
		var viewportPtr C.QtObjectPtr
		if viewport != nil {
			viewportPtr = viewport.Pointer()
		}
		C.QAbstractScrollArea_SetupViewport_QWidget(p.Pointer(), viewportPtr)
	}
}

func (p *qabstractscrollarea) VerticalScrollBarPolicy() ScrollBarPolicy {
	if p.Pointer() == nil {
		return 0
	}
	return ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(p.Pointer()))
}

func (p *qabstractscrollarea) Viewport() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QAbstractScrollArea_Viewport(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}
