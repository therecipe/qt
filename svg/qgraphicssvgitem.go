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

type QGraphicsSvgItemITF interface {
	widgets.QGraphicsObjectITF
	QGraphicsSvgItemPTR() *QGraphicsSvgItem
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItemPTR().Pointer()
	}
	return nil
}

func QGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsSvgItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItemPTR() *QGraphicsSvgItem {
	return ptr
}

func NewQGraphicsSvgItem(parent widgets.QGraphicsItemITF) *QGraphicsSvgItem {
	return QGraphicsSvgItemFromPointer(unsafe.Pointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem(C.QtObjectPtr(widgets.PointerFromQGraphicsItem(parent)))))
}

func NewQGraphicsSvgItem2(fileName string, parent widgets.QGraphicsItemITF) *QGraphicsSvgItem {
	return QGraphicsSvgItemFromPointer(unsafe.Pointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem2(C.CString(fileName), C.QtObjectPtr(widgets.PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSvgItem_ElementId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainterITF, option widgets.QStyleOptionGraphicsItemITF, widget widgets.QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(widgets.PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(widgets.PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	if ptr.Pointer() != nil {
		return QSvgRendererFromPointer(unsafe.Pointer(C.QGraphicsSvgItem_Renderer(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetElementId(C.QtObjectPtr(ptr.Pointer()), C.CString(id))
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRendererITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSvgRenderer(renderer)))
	}
}

func (ptr *QGraphicsSvgItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsSvgItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
