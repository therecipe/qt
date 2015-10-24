package widgets

//#include "qscrollbar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScrollBar struct {
	QAbstractSlider
}

type QScrollBarITF interface {
	QAbstractSliderITF
	QScrollBarPTR() *QScrollBar
}

func PointerFromQScrollBar(ptr QScrollBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollBarPTR().Pointer()
	}
	return nil
}

func QScrollBarFromPointer(ptr unsafe.Pointer) *QScrollBar {
	var n = new(QScrollBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QScrollBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScrollBar) QScrollBarPTR() *QScrollBar {
	return ptr
}

func NewQScrollBar(parent QWidgetITF) *QScrollBar {
	return QScrollBarFromPointer(unsafe.Pointer(C.QScrollBar_NewQScrollBar(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQScrollBar2(orientation core.Qt__Orientation, parent QWidgetITF) *QScrollBar {
	return QScrollBarFromPointer(unsafe.Pointer(C.QScrollBar_NewQScrollBar2(C.int(orientation), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QScrollBar) Event(event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QScrollBar_Event(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QScrollBar) DestroyQScrollBar() {
	if ptr.Pointer() != nil {
		C.QScrollBar_DestroyQScrollBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
