package svg

//#include "qgraphicssvgitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsSvgItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return ptr
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSvgItem_ElementId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QGraphicsSvgItem_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetElementId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(ptr.Pointer(), PointerFromQSvgRenderer(renderer))
	}
}

func (ptr *QGraphicsSvgItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsSvgItem_Type(ptr.Pointer()))
	}
	return 0
}
