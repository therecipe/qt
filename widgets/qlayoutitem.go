package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QLayoutItem struct {
	ptr unsafe.Pointer
}

type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (p *QLayoutItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLayoutItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLayoutItem(ptr QLayoutItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayoutItem_PTR().Pointer()
	}
	return nil
}

func NewQLayoutItemFromPointer(ptr unsafe.Pointer) *QLayoutItem {
	var n = new(QLayoutItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLayoutItem_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::controlTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::expandingDirections")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) HasHeightForWidth() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::hasHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayoutItem_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) HeightForWidth(w int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayoutItem_Invalidate(ptr.Pointer())
	}
}

func (ptr *QLayoutItem) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayoutItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) Layout() *QLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::layout")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayoutItem_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) MinimumHeightForWidth(w int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::minimumHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayoutItem_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayoutItem) SpacerItem() *QSpacerItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::spacerItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QLayoutItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayoutItem_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayoutItem::~QLayoutItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayoutItem_DestroyQLayoutItem(ptr.Pointer())
	}
}
