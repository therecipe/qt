package qt

//#include "qvboxlayout.h"
import "C"

type qvboxlayout struct {
	qboxlayout
}

type QVBoxLayout interface {
	QBoxLayout
}

func (p *qvboxlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qvboxlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQVBoxLayout() QVBoxLayout {
	var qvboxlayout = new(qvboxlayout)
	qvboxlayout.SetPointer(C.QVBoxLayout_New())
	qvboxlayout.SetObjectName_String("QVBoxLayout_" + randomIdentifier())
	return qvboxlayout
}

func NewQVBoxLayout_QWidget(parent QWidget) QVBoxLayout {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qvboxlayout = new(qvboxlayout)
	qvboxlayout.SetPointer(C.QVBoxLayout_New_QWidget(parentPtr))
	qvboxlayout.SetObjectName_String("QVBoxLayout_" + randomIdentifier())
	return qvboxlayout
}

func (p *qvboxlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QVBoxLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
