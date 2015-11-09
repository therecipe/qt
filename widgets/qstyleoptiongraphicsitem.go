package widgets

//#include "qstyleoptiongraphicsitem.h"
import "C"
import (
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStyleOptionGraphicsItem struct {
	QStyleOption
}

type QStyleOptionGraphicsItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem
}

func PointerFromQStyleOptionGraphicsItem(ptr QStyleOptionGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionGraphicsItemFromPointer(ptr unsafe.Pointer) *QStyleOptionGraphicsItem {
	var n = new(QStyleOptionGraphicsItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionGraphicsItem) QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem {
	return ptr
}

//QStyleOptionGraphicsItem::StyleOptionType
type QStyleOptionGraphicsItem__StyleOptionType int64

var (
	QStyleOptionGraphicsItem__Type = QStyleOptionGraphicsItem__StyleOptionType(QStyleOption__SO_GraphicsItem)
)

//QStyleOptionGraphicsItem::StyleOptionVersion
type QStyleOptionGraphicsItem__StyleOptionVersion int64

var (
	QStyleOptionGraphicsItem__Version = QStyleOptionGraphicsItem__StyleOptionVersion(1)
)

func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	return NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem())
}

func NewQStyleOptionGraphicsItem2(other QStyleOptionGraphicsItem_ITF) *QStyleOptionGraphicsItem {
	return NewQStyleOptionGraphicsItemFromPointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(PointerFromQStyleOptionGraphicsItem(other)))
}

func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform gui.QTransform_ITF) float64 {
	return float64(C.QStyleOptionGraphicsItem_QStyleOptionGraphicsItem_LevelOfDetailFromTransform(gui.PointerFromQTransform(worldTransform)))
}
