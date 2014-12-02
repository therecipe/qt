package qt

//#include "qslider.h"
import "C"

type qslider struct {
	qabstractslider
}

type QSlider interface {
	QAbstractSlider
	SetTickInterval_Int(ti int)
	TickInterval() int
}

func (p *qslider) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qslider) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSlider_QWidget(parent QWidget) QSlider {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qslider = new(qslider)
	qslider.SetPointer(C.QSlider_New_QWidget(parentPtr))
	qslider.SetObjectName_String("QSlider_" + randomIdentifier())
	return qslider
}

func NewQSlider_Orientation_QWidget(orientation Orientation, parent QWidget) QSlider {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qslider = new(qslider)
	qslider.SetPointer(C.QSlider_New_Orientation_QWidget(C.int(orientation), parentPtr))
	qslider.SetObjectName_String("QSlider_" + randomIdentifier())
	return qslider
}

func (p *qslider) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QSlider_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qslider) SetTickInterval_Int(ti int) {
	if p.Pointer() != nil {
		C.QSlider_SetTickInterval_Int(p.Pointer(), C.int(ti))
	}
}

func (p *qslider) TickInterval() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSlider_TickInterval(p.Pointer()))
	}
}
