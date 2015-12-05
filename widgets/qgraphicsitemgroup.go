package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsItemGroup struct {
	QGraphicsItem
}

type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func PointerFromQGraphicsItemGroup(ptr QGraphicsItemGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItemGroup_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemGroupFromPointer(ptr unsafe.Pointer) *QGraphicsItemGroup {
	var n = new(QGraphicsItemGroup)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup {
	return ptr
}

func NewQGraphicsItemGroup(parent QGraphicsItem_ITF) *QGraphicsItemGroup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::QGraphicsItemGroup")
		}
	}()

	return NewQGraphicsItemGroupFromPointer(C.QGraphicsItemGroup_NewQGraphicsItemGroup(PointerFromQGraphicsItem(parent)))
}

func (ptr *QGraphicsItemGroup) AddToGroup(item QGraphicsItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::addToGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_AddToGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::isObscuredBy")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QGraphicsItemGroup_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsItemGroup) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsItemGroup) RemoveFromGroup(item QGraphicsItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::removeFromGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_RemoveFromGroup(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsItemGroup) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsItemGroup_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsItemGroup) DestroyQGraphicsItemGroup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsItemGroup::~QGraphicsItemGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsItemGroup_DestroyQGraphicsItemGroup(ptr.Pointer())
	}
}
