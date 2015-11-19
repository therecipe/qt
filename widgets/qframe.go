package widgets

//#include "qframe.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFrame struct {
	QWidget
}

type QFrame_ITF interface {
	QWidget_ITF
	QFrame_PTR() *QFrame
}

func PointerFromQFrame(ptr QFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFrame_PTR().Pointer()
	}
	return nil
}

func NewQFrameFromPointer(ptr unsafe.Pointer) *QFrame {
	var n = new(QFrame)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFrame_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFrame) QFrame_PTR() *QFrame {
	return ptr
}

//QFrame::Shadow
type QFrame__Shadow int64

const (
	QFrame__Plain  = QFrame__Shadow(0x0010)
	QFrame__Raised = QFrame__Shadow(0x0020)
	QFrame__Sunken = QFrame__Shadow(0x0030)
)

//QFrame::Shape
type QFrame__Shape int64

const (
	QFrame__NoFrame     = QFrame__Shape(0)
	QFrame__Box         = QFrame__Shape(0x0001)
	QFrame__Panel       = QFrame__Shape(0x0002)
	QFrame__WinPanel    = QFrame__Shape(0x0003)
	QFrame__HLine       = QFrame__Shape(0x0004)
	QFrame__VLine       = QFrame__Shape(0x0005)
	QFrame__StyledPanel = QFrame__Shape(0x0006)
)

//QFrame::StyleMask
type QFrame__StyleMask int64

var (
	QFrame__Shadow_Mask = QFrame__StyleMask(0x00f0)
	QFrame__Shape_Mask  = QFrame__StyleMask(0x000f)
)

func (ptr *QFrame) FrameShadow() QFrame__Shadow {
	if ptr.Pointer() != nil {
		return QFrame__Shadow(C.QFrame_FrameShadow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameShape() QFrame__Shape {
	if ptr.Pointer() != nil {
		return QFrame__Shape(C.QFrame_FrameShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) FrameWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) LineWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFrame_LineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) MidLineWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QFrame_MidLineWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) SetFrameRect(v core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QFrame_SetFrameRect(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QFrame) SetFrameShadow(v QFrame__Shadow) {
	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShadow(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetFrameShape(v QFrame__Shape) {
	if ptr.Pointer() != nil {
		C.QFrame_SetFrameShape(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetLineWidth(v int) {
	if ptr.Pointer() != nil {
		C.QFrame_SetLineWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QFrame) SetMidLineWidth(v int) {
	if ptr.Pointer() != nil {
		C.QFrame_SetMidLineWidth(ptr.Pointer(), C.int(v))
	}
}

func NewQFrame(parent QWidget_ITF, f core.Qt__WindowType) *QFrame {
	return NewQFrameFromPointer(C.QFrame_NewQFrame(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QFrame) FrameStyle() int {
	if ptr.Pointer() != nil {
		return int(C.QFrame_FrameStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFrame) SetFrameStyle(style int) {
	if ptr.Pointer() != nil {
		C.QFrame_SetFrameStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QFrame) DestroyQFrame() {
	if ptr.Pointer() != nil {
		C.QFrame_DestroyQFrame(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
