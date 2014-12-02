package qt

//#include "qspaceritem.h"
import "C"

type qspaceritem struct {
	qlayoutitem
}

type QSpacerItem interface {
	QLayoutItem
}

func (p *qspaceritem) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qspaceritem) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qspaceritem) Destroy() {
	if p.Pointer() != nil {
		C.QSpacerItem_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
