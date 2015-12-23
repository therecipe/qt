package multimediawidgets

//#include "multimediawidgets.h"
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
		n.SetObjectName("QGraphicsVideoItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsVideoItem) QGraphicsVideoItem_PTR() *QGraphicsVideoItem {
	return ptr
}

func NewQGraphicsVideoItem(parent widgets.QGraphicsItem_ITF) *QGraphicsVideoItem {
	defer qt.Recovering("QGraphicsVideoItem::QGraphicsVideoItem")

	return NewQGraphicsVideoItemFromPointer(C.QGraphicsVideoItem_NewQGraphicsVideoItem(widgets.PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsVideoItem) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QGraphicsVideoItem::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QGraphicsVideoItem_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsVideoItem) MediaObject() *multimedia.QMediaObject {
	defer qt.Recovering("QGraphicsVideoItem::mediaObject")

	if ptr.Pointer() != nil {
		return multimedia.NewQMediaObjectFromPointer(C.QGraphicsVideoItem_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsVideoItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsVideoItemPaint
func callbackQGraphicsVideoItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsVideoItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsVideoItem) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsVideoItem::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsVideoItem) SetOffset(offset core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetOffset(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QGraphicsVideoItem) SetSize(size core.QSizeF_ITF) {
	defer qt.Recovering("QGraphicsVideoItem::setSize")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_SetSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsVideoItem) DestroyQGraphicsVideoItem() {
	defer qt.Recovering("QGraphicsVideoItem::~QGraphicsVideoItem")

	if ptr.Pointer() != nil {
		C.QGraphicsVideoItem_DestroyQGraphicsVideoItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsVideoItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsVideoItemTimerEvent
func callbackQGraphicsVideoItemTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsVideoItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsVideoItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsVideoItemChildEvent
func callbackQGraphicsVideoItemChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsVideoItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsVideoItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsVideoItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsVideoItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsVideoItemCustomEvent
func callbackQGraphicsVideoItemCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsVideoItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
