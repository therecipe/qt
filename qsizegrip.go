package qt

//#include "qsizegrip.h"
import "C"

type qsizegrip struct {
	qwidget
}

type QSizeGrip interface {
	QWidget
}

func (p *qsizegrip) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qsizegrip) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQSizeGrip_QWidget(parent QWidget) QSizeGrip {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qsizegrip = new(qsizegrip)
	qsizegrip.SetPointer(C.QSizeGrip_New_QWidget(parentPtr))
	qsizegrip.SetObjectName_String("QSizeGrip_" + randomIdentifier())
	return qsizegrip
}

func (p *qsizegrip) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QSizeGrip_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
