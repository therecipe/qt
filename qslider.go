package qt

//#include "qslider.h"
import "C"

type qslider struct {
	qabstractslider
}

type QSlider interface {
	QAbstractSlider
	SetTickInterval(ti int)
	TickInterval() int
}

func (p *qslider) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qslider) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSlider1(parent QWidget) QSlider {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qslider = new(qslider)
	qslider.SetPointer(C.QSlider_New_QWidget(parentPtr))
	qslider.SetObjectName("QSlider_" + randomIdentifier())
	return qslider
}

func NewQSlider2(orientation Orientation, parent QWidget) QSlider {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qslider = new(qslider)
	qslider.SetPointer(C.QSlider_New_Orientation_QWidget(C.int(orientation), parentPtr))
	qslider.SetObjectName("QSlider_" + randomIdentifier())
	return qslider
}

func (p *qslider) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QSlider_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qslider) SetTickInterval(ti int) {
	if p.Pointer() != nil {
		C.QSlider_SetTickInterval_Int(p.Pointer(), C.int(ti))
	}
}

func (p *qslider) TickInterval() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QSlider_TickInterval(p.Pointer()))
}
