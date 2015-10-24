package widgets

//#include "qwhatsthis.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWhatsThis struct {
	ptr unsafe.Pointer
}

type QWhatsThisITF interface {
	QWhatsThisPTR() *QWhatsThis
}

func (p *QWhatsThis) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWhatsThis) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWhatsThis(ptr QWhatsThisITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWhatsThisPTR().Pointer()
	}
	return nil
}

func QWhatsThisFromPointer(ptr unsafe.Pointer) *QWhatsThis {
	var n = new(QWhatsThis)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWhatsThis) QWhatsThisPTR() *QWhatsThis {
	return ptr
}

func QWhatsThis_CreateAction(parent core.QObjectITF) *QAction {
	return QActionFromPointer(unsafe.Pointer(C.QWhatsThis_QWhatsThis_CreateAction(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func QWhatsThis_EnterWhatsThisMode() {
	C.QWhatsThis_QWhatsThis_EnterWhatsThisMode()
}

func QWhatsThis_HideText() {
	C.QWhatsThis_QWhatsThis_HideText()
}

func QWhatsThis_InWhatsThisMode() bool {
	return C.QWhatsThis_QWhatsThis_InWhatsThisMode() != 0
}

func QWhatsThis_LeaveWhatsThisMode() {
	C.QWhatsThis_QWhatsThis_LeaveWhatsThisMode()
}

func QWhatsThis_ShowText(pos core.QPointITF, text string, w QWidgetITF) {
	C.QWhatsThis_QWhatsThis_ShowText(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.CString(text), C.QtObjectPtr(PointerFromQWidget(w)))
}
