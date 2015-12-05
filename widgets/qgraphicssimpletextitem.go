package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsSimpleTextItem struct {
	QAbstractGraphicsShapeItem
}

type QGraphicsSimpleTextItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsSimpleTextItem_PTR() *QGraphicsSimpleTextItem
}

func PointerFromQGraphicsSimpleTextItem(ptr QGraphicsSimpleTextItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSimpleTextItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSimpleTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsSimpleTextItem {
	var n = new(QGraphicsSimpleTextItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSimpleTextItem) QGraphicsSimpleTextItem_PTR() *QGraphicsSimpleTextItem {
	return ptr
}

func (ptr *QGraphicsSimpleTextItem) Contains(point core.QPointF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::isObscuredBy")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsSimpleTextItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsSimpleTextItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetFont(font gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsSimpleTextItem) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsSimpleTextItem) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSimpleTextItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSimpleTextItem) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsSimpleTextItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSimpleTextItem) DestroyQGraphicsSimpleTextItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSimpleTextItem::~QGraphicsSimpleTextItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(ptr.Pointer())
	}
}
