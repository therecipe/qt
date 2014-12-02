package qt

//#include "qwidgetitem.h"
import "C"

type qwidgetitem struct {
	qlayoutitem
}

type QWidgetItem interface {
	QLayoutItem
}

func (p *qwidgetitem) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qwidgetitem) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQWidgetItem_QWidget(widget QWidget) QWidgetItem {
	var widgetPtr C.QtObjectPtr = nil
	if widget != nil {
		widgetPtr = widget.Pointer()
	}
	var qwidgetitem = new(qwidgetitem)
	qwidgetitem.SetPointer(C.QWidgetItem_New_QWidget(widgetPtr))
	return qwidgetitem
}

func (p *qwidgetitem) Destroy() {
	if p.Pointer() != nil {
		C.QWidgetItem_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}
