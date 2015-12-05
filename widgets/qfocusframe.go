package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
	for len(n.ObjectName()) < len("QFocusFrame_") {
		n.SetObjectName("QFocusFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame {
	return ptr
}

func NewQFocusFrame(parent QWidget_ITF) *QFocusFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusFrame::QFocusFrame")
		}
	}()

	return NewQFocusFrameFromPointer(C.QFocusFrame_NewQFocusFrame(PointerFromQWidget(parent)))
}

func (ptr *QFocusFrame) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusFrame::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFocusFrame_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFocusFrame) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusFrame::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFocusFrame_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFocusFrame) DestroyQFocusFrame() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFocusFrame::~QFocusFrame")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFocusFrame_DestroyQFocusFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
