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

type QScrollBar_ITF interface {
	QAbstractSlider_ITF
	QScrollBar_PTR() *QScrollBar
}

func PointerFromQScrollBar(ptr QScrollBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollBar_PTR().Pointer()
	}
	return nil
}

func NewQScrollBarFromPointer(ptr unsafe.Pointer) *QScrollBar {
	var n = new(QScrollBar)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QScrollBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScrollBar) QScrollBar_PTR() *QScrollBar {
	return ptr
}

func NewQScrollBar(parent QWidget_ITF) *QScrollBar {
	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar(PointerFromQWidget(parent)))
}

func NewQScrollBar2(orientation core.Qt__Orientation, parent QWidget_ITF) *QScrollBar {
	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QScrollBar) Event(event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QScrollBar_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScrollBar) DestroyQScrollBar() {
	if ptr.Pointer() != nil {
		C.QScrollBar_DestroyQScrollBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
