package qt

//#include "qhboxlayout.h"
import "C"

type qhboxlayout struct {
	qboxlayout
}

type QHBoxLayout interface {
	QBoxLayout
}

func (p *qhboxlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qhboxlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQHBoxLayout1() QHBoxLayout {
	var qhboxlayout = new(qhboxlayout)
	qhboxlayout.SetPointer(C.QHBoxLayout_New())
	qhboxlayout.SetObjectName("QHBoxLayout_" + randomIdentifier())
	return qhboxlayout
}

func NewQHBoxLayout2(parent QWidget) QHBoxLayout {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qhboxlayout = new(qhboxlayout)
	qhboxlayout.SetPointer(C.QHBoxLayout_New_QWidget(parentPtr))
	qhboxlayout.SetObjectName("QHBoxLayout_" + randomIdentifier())
	return qhboxlayout
}

func (p *qhboxlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QHBoxLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
