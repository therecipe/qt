package widgets

//#include "qscrollarea.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScrollArea struct {
	QAbstractScrollArea
}

type QScrollAreaITF interface {
	QAbstractScrollAreaITF
	QScrollAreaPTR() *QScrollArea
}

func PointerFromQScrollArea(ptr QScrollAreaITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollAreaPTR().Pointer()
	}
	return nil
}

func QScrollAreaFromPointer(ptr unsafe.Pointer) *QScrollArea {
	var n = new(QScrollArea)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScrollArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScrollArea) QScrollAreaPTR() *QScrollArea {
	return ptr
}

func (ptr *QScrollArea) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QScrollArea_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScrollArea) SetAlignment(v core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QScrollArea_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QScrollArea) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QScrollArea) SetWidgetResizable(resizable bool) {
	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidgetResizable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(resizable)))
	}
}

func (ptr *QScrollArea) WidgetResizable() bool {
	if ptr.Pointer() != nil {
		return C.QScrollArea_WidgetResizable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQScrollArea(parent QWidgetITF) *QScrollArea {
	return QScrollAreaFromPointer(unsafe.Pointer(C.QScrollArea_NewQScrollArea(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureVisible(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) EnsureWidgetVisible(childWidget QWidgetITF, xmargin int, ymargin int) {
	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureWidgetVisible(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(childWidget)), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) FocusNextPrevChild(next bool) bool {
	if ptr.Pointer() != nil {
		return C.QScrollArea_FocusNextPrevChild(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QScrollArea) TakeWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QScrollArea_TakeWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScrollArea) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QScrollArea_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QScrollArea) DestroyQScrollArea() {
	if ptr.Pointer() != nil {
		C.QScrollArea_DestroyQScrollArea(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
