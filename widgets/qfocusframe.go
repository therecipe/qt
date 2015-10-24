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

type QFocusFrameITF interface {
	QWidgetITF
	QFocusFramePTR() *QFocusFrame
}

func PointerFromQFocusFrame(ptr QFocusFrameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusFramePTR().Pointer()
	}
	return nil
}

func QFocusFrameFromPointer(ptr unsafe.Pointer) *QFocusFrame {
	var n = new(QFocusFrame)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFocusFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFramePTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidgetITF) *QFocusFrame {
	return QFocusFrameFromPointer(unsafe.Pointer(C.QFocusFrame_NewQFocusFrame(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QFocusFrame) SetWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QFocusFrame_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
