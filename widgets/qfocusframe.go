package widgets

//#include "qfocusframe.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFocusFrame struct {
	QWidget
}

type QFocusFrame_ITF interface {
	QWidget_ITF
	QFocusFrame_PTR() *QFocusFrame
}

func PointerFromQFocusFrame(ptr QFocusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusFrame_PTR().Pointer()
	}
	return nil
}

func NewQFocusFrameFromPointer(ptr unsafe.Pointer) *QFocusFrame {
	var n = new(QFocusFrame)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QFocusFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidget_ITF) *QFocusFrame {
	return NewQFocusFrameFromPointer(C.QFocusFrame_NewQFocusFrame(PointerFromQWidget(parent)))
}

func (ptr *QFocusFrame) SetWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFocusFrame_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
