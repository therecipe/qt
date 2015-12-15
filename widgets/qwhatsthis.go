package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWhatsThis struct {
	ptr unsafe.Pointer
}

type QWhatsThis_ITF interface {
	QWhatsThis_PTR() *QWhatsThis
}

func (p *QWhatsThis) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWhatsThis) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWhatsThis(ptr QWhatsThis_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWhatsThis_PTR().Pointer()
	}
	return nil
}

func NewQWhatsThisFromPointer(ptr unsafe.Pointer) *QWhatsThis {
	var n = new(QWhatsThis)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWhatsThis) QWhatsThis_PTR() *QWhatsThis {
	return ptr
}

func QWhatsThis_CreateAction(parent core.QObject_ITF) *QAction {
	defer qt.Recovering("QWhatsThis::createAction")

	return NewQActionFromPointer(C.QWhatsThis_QWhatsThis_CreateAction(core.PointerFromQObject(parent)))
}

func QWhatsThis_EnterWhatsThisMode() {
	defer qt.Recovering("QWhatsThis::enterWhatsThisMode")

	C.QWhatsThis_QWhatsThis_EnterWhatsThisMode()
}

func QWhatsThis_HideText() {
	defer qt.Recovering("QWhatsThis::hideText")

	C.QWhatsThis_QWhatsThis_HideText()
}

func QWhatsThis_InWhatsThisMode() bool {
	defer qt.Recovering("QWhatsThis::inWhatsThisMode")

	return C.QWhatsThis_QWhatsThis_InWhatsThisMode() != 0
}

func QWhatsThis_LeaveWhatsThisMode() {
	defer qt.Recovering("QWhatsThis::leaveWhatsThisMode")

	C.QWhatsThis_QWhatsThis_LeaveWhatsThisMode()
}

func QWhatsThis_ShowText(pos core.QPoint_ITF, text string, w QWidget_ITF) {
	defer qt.Recovering("QWhatsThis::showText")

	C.QWhatsThis_QWhatsThis_ShowText(core.PointerFromQPoint(pos), C.CString(text), PointerFromQWidget(w))
}
