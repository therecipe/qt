package widgets

//#include "qgraphicssimpletextitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsSimpleTextItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsSimpleTextItemITF interface {
	QAbstractGraphicsShapeItemITF
	QGraphicsSimpleTextItemPTR() *QGraphicsSimpleTextItem
}

func PointerFromQGraphicsSimpleTextItem(ptr QGraphicsSimpleTextItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSimpleTextItemPTR().Pointer()
	}
	return nil
}

func QGraphicsSimpleTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsSimpleTextItem {
	var n = new(QGraphicsSimpleTextItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSimpleTextItem) QGraphicsSimpleTextItemPTR() *QGraphicsSimpleTextItem {
	return ptr
}

func (ptr *QGraphicsSimpleTextItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QGraphicsSimpleTextItem) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSimpleTextItem_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGraphicsSimpleTextItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsSimpleTextItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSimpleTextItem) DestroyQGraphicsSimpleTextItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
