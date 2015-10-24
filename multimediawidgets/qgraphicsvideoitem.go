package multimediawidgets

//#include "qgraphicsvideoitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QGraphicsVideoItem struct {
	multimedia.QMediaBindableInterface
	widgets.QGraphicsObject
}

type QGraphicsVideoItemITF interface {
	multimedia.QMediaBindableInterfaceITF
	widgets.QGraphicsObjectITF
	QGraphicsVideoItemPTR() *QGraphicsVideoItem
}

func (p *QGraphicsVideoItem) Pointer() unsafe.Pointer {
	return p.QMediaBindableInterfacePTR().Pointer()
}

func (p *QGraphicsVideoItem) SetPointer(ptr unsafe.Pointer) {
	p.QMediaBindableInterfacePTR().SetPointer(ptr)
	p.QGraphicsObjectPTR().SetPointer(ptr)
}

func PointerFromQGraphicsVideoItem(ptr QGraphicsVideoItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsVideoItemPTR().Pointer()
	}
	return nil
}

func QGraphicsVideoItemFromPointer(ptr unsafe.Pointer) *QGraphicsVideoItem {
	var n = new(QGraphicsVideoItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsVideoItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsVideoItem) QGraphicsVideoItemPTR() *QGraphicsVideoItem {
	return ptr
}

func NewQGraphicsVideoItem(parent widgets.QGraphicsItemITF) *QGraphicsVideoItem {
	return QGraphicsVideoItemFromPointer(unsafe.Pointer(C.QGraphicsVideoItem_NewQGraphicsVideoItem(C.QtObjectPtr(widgets.PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsVideoItem) AspectRatioMode() core.Qt__AspectRatioMode {
	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QGraphicsVideoItem_AspectRatioMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsVideoItem) MediaObject() *multimedia.QMediaObject {
	if ptr.Pointer() != nil {
		return multimedia.QMediaObjectFromPointer(unsafe.Pointer(C.QGraphicsVideoItem_MediaObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) Paint(painter gui.QPainterITF, option widgets.QStyleOptionGraphicsItemITF, widget widgets.QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(widgets.PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(widgets.PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsVideoItem) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetAspectRatioMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsVideoItem) SetOffset(offset core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(offset)))
	}
}

func (ptr *QGraphicsVideoItem) SetSize(size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QGraphicsVideoItem) DestroyQGraphicsVideoItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DestroyQGraphicsVideoItem(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
