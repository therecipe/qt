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

type QAbstractScrollAreaITF interface {
	QFrameITF
	QAbstractScrollAreaPTR() *QAbstractScrollArea
}

func PointerFromQAbstractScrollArea(ptr QAbstractScrollAreaITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractScrollAreaPTR().Pointer()
	}
	return nil
}

func QAbstractScrollAreaFromPointer(ptr unsafe.Pointer) *QAbstractScrollArea {
	var n = new(QAbstractScrollArea)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractScrollArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractScrollArea) QAbstractScrollAreaPTR() *QAbstractScrollArea {
	return ptr
}

//QAbstractScrollArea::SizeAdjustPolicy
type QAbstractScrollArea__SizeAdjustPolicy int

var (
	QAbstractScrollArea__AdjustIgnored               = QAbstractScrollArea__SizeAdjustPolicy(0)
	QAbstractScrollArea__AdjustToContentsOnFirstShow = QAbstractScrollArea__SizeAdjustPolicy(1)
	QAbstractScrollArea__AdjustToContents            = QAbstractScrollArea__SizeAdjustPolicy(2)
)

func (ptr *QAbstractScrollArea) HorizontalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_HorizontalScrollBarPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBarPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetSizeAdjustPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBarPolicy(v core.Qt__ScrollBarPolicy) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBarPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	if ptr.Pointer() != nil {
		return QAbstractScrollArea__SizeAdjustPolicy(C.QAbstractScrollArea_SizeAdjustPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractScrollArea) VerticalScrollBarPolicy() core.Qt__ScrollBarPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QAbstractScrollArea_VerticalScrollBarPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQAbstractScrollArea(parent QWidgetITF) *QAbstractScrollArea {
	return QAbstractScrollAreaFromPointer(unsafe.Pointer(C.QAbstractScrollArea_NewQAbstractScrollArea(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QAbstractScrollArea) AddScrollBarWidget(widget QWidgetITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_AddScrollBarWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(alignment))
	}
}

func (ptr *QAbstractScrollArea) CornerWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QAbstractScrollArea_CornerWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		return QScrollBarFromPointer(unsafe.Pointer(C.QAbstractScrollArea_HorizontalScrollBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractScrollArea) SetCornerWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetCornerWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QAbstractScrollArea) SetHorizontalScrollBar(scrollBar QScrollBarITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetHorizontalScrollBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScrollBar(scrollBar)))
	}
}

func (ptr *QAbstractScrollArea) SetVerticalScrollBar(scrollBar QScrollBarITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetVerticalScrollBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQScrollBar(scrollBar)))
	}
}

func (ptr *QAbstractScrollArea) SetViewport(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetViewport(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QAbstractScrollArea) SetupViewport(viewport QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_SetupViewport(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(viewport)))
	}
}

func (ptr *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	if ptr.Pointer() != nil {
		return QScrollBarFromPointer(unsafe.Pointer(C.QAbstractScrollArea_VerticalScrollBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractScrollArea) Viewport() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QAbstractScrollArea_Viewport(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractScrollArea) DestroyQAbstractScrollArea() {
	if ptr.Pointer() != nil {
		C.QAbstractScrollArea_DestroyQAbstractScrollArea(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
