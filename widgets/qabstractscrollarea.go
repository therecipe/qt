package widgets

//#include "qabstractscrollarea.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractScrollArea struct {
	QFrame
}

type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQAbstractScrollAreaFromPointer(ptr unsafe.Pointer) *QAbstractScrollArea {
	var n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractScrollArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea {
	return ptr
}

//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int64

const (
	QAbstractScrollArea__AdjustIgnored               = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetSizeAdjustPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	if ptr.Pointer() != nil {
		return QAbstractScrollArea__SizeAdjustPolicy(C.QAbstractScrollArea_SizeAdjustPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractScrollArea(parent QWidget_ITF) *QAbstractScrollArea {
	return NewQAbstractScrollAreaFromPointer(C.QAbstractScrollArea_NewQAbstractScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QAbstractScrollArea) AddScrollBarWidget(widget QWidget_ITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_AddScrollBarWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(alignment))
	}
}

func (ptr *QAbstractScrollArea) CornerWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_CornerWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_HorizontalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) SetCornerWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBar_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBar_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(ptr.Pointer(), PointerFromQScrollBar(scrollBar))
	}
}

func (ptr *QAbstractScrollArea) SetViewport(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetViewport(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractScrollArea) SetupViewport(viewport QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		return NewQScrollBarFromPointer(C.QAbstractScrollArea_VerticalScrollBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) Viewport() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractScrollArea_Viewport(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
