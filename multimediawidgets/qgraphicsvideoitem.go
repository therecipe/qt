package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"log"
	"unsafe"
)

type QGraphicsVideoItem struct {
	multimedia.QMediaBindableInterface
	widgets.QGraphicsObject
}

type QGraphicsVideoItem_ITF interface {
	multimedia.QMediaBindableInterface_ITF
	widgets.QGraphicsObject_ITF
	QGraphicsVideoItem_PTR() *QGraphicsVideoItem
}

func (p *QGraphicsVideoItem) Pointer() unsafe.Pointer {
	return p.QMediaBindableInterface_PTR().Pointer()
}

func (p *QGraphicsVideoItem) SetPointer(ptr unsafe.Pointer) {
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
	p.QGraphicsObject_PTR().SetPointer(ptr)
}

func PointerFromQGraphicsVideoItem(ptr QGraphicsVideoItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsVideoItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsVideoItemFromPointer(ptr unsafe.Pointer) *QGraphicsVideoItem {
	var n = new(QGraphicsVideoItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsVideoItem_") {
		n.SetObjectName("QGraphicsVideoItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsVideoItem) QGraphicsVideoItem_PTR() *QGraphicsVideoItem {
	return ptr
}

func NewQGraphicsVideoItem(parent widgets.QGraphicsItem_ITF) *QGraphicsVideoItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::QGraphicsVideoItem")
		}
	}()

	return NewQGraphicsVideoItemFromPointer(C.QGraphicsVideoItem_NewQGraphicsVideoItem(widgets.PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsVideoItem) AspectRatioMode() core.Qt__AspectRatioMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::aspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QGraphicsVideoItem_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsVideoItem) MediaObject() *multimedia.QMediaObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::mediaObject")
		}
	}()

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QGraphicsVideoItem_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsVideoItem) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::setAspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsVideoItem) SetOffset(offset core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::setOffset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsVideoItem) SetSize(size core.QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::setSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsVideoItem) DestroyQGraphicsVideoItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsVideoItem::~QGraphicsVideoItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DestroyQGraphicsVideoItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
