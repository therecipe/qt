package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	for len(n.ObjectNameAbs()) < len("QLayoutItem_") {
		n.SetObjectNameAbs("QLayoutItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem {
	return ptr
}

func (ptr *QLayoutItem) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QLayoutItem::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLayoutItem_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	defer qt.Recovering("QLayoutItem::controlTypes")

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayoutItem_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) ExpandingDirections() core.Qt__Orientation {
	defer qt.Recovering("QLayoutItem::expandingDirections")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayoutItem_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayoutItem) HasHeightForWidth() bool {
	defer qt.Recovering("QLayoutItem::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QLayoutItem_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) HeightForWidth(w int) int {
	defer qt.Recovering("QLayoutItem::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) ConnectInvalidate(f func()) {
	defer qt.Recovering("connect QLayoutItem::invalidate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "invalidate", f)
	}
}

func (ptr *QLayoutItem) DisconnectInvalidate() {
	defer qt.Recovering("disconnect QLayoutItem::invalidate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "invalidate")
	}
}

//export callbackQLayoutItemInvalidate
func callbackQLayoutItemInvalidate(ptrName *C.char) bool {
	defer qt.Recovering("callback QLayoutItem::invalidate")

	var signal = qt.GetSignal(C.GoString(ptrName), "invalidate")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QLayoutItem) IsEmpty() bool {
	defer qt.Recovering("QLayoutItem::isEmpty")

	if ptr.Pointer() != nil {
		return C.QLayoutItem_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayoutItem) Layout() *QLayout {
	defer qt.Recovering("QLayoutItem::layout")

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayoutItem_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) MinimumHeightForWidth(w int) int {
	defer qt.Recovering("QLayoutItem::minimumHeightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QLayoutItem_MinimumHeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLayoutItem) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QLayoutItem::setAlignment")

	if ptr.Pointer() != nil {
		C.QLayoutItem_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QLayoutItem) SetGeometry(r core.QRect_ITF) {
	defer qt.Recovering("QLayoutItem::setGeometry")

	if ptr.Pointer() != nil {
		C.QLayoutItem_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayoutItem) SpacerItem() *QSpacerItem {
	defer qt.Recovering("QLayoutItem::spacerItem")

	if ptr.Pointer() != nil {
		return NewQSpacerItemFromPointer(C.QLayoutItem_SpacerItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) Widget() *QWidget {
	defer qt.Recovering("QLayoutItem::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayoutItem_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayoutItem) DestroyQLayoutItem() {
	defer qt.Recovering("QLayoutItem::~QLayoutItem")

	if ptr.Pointer() != nil {
		C.QLayoutItem_DestroyQLayoutItem(ptr.Pointer())
	}
}

func (ptr *QLayoutItem) ObjectNameAbs() string {
	defer qt.Recovering("QLayoutItem::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLayoutItem_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLayoutItem) SetObjectNameAbs(name string) {
	defer qt.Recovering("QLayoutItem::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QLayoutItem_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
