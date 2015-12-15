package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStackedLayout struct {
	QLayout
}

type QStackedLayout_ITF interface {
	QLayout_ITF
	QStackedLayout_PTR() *QStackedLayout
}

func PointerFromQStackedLayout(ptr QStackedLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedLayout_PTR().Pointer()
	}
	return nil
}

func NewQStackedLayoutFromPointer(ptr unsafe.Pointer) *QStackedLayout {
	var n = new(QStackedLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStackedLayout_") {
		n.SetObjectName("QStackedLayout_" + qt.Identifier())
	}
	return n
}

func (ptr *QStackedLayout) QStackedLayout_PTR() *QStackedLayout {
	return ptr
}

//QStackedLayout::StackingMode
type QStackedLayout__StackingMode int64

const (
	QStackedLayout__StackOne = QStackedLayout__StackingMode(0)
	QStackedLayout__StackAll = QStackedLayout__StackingMode(1)
)

func (ptr *QStackedLayout) Count() int {
	defer qt.Recovering("QStackedLayout::count")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) CurrentIndex() int {
	defer qt.Recovering("QStackedLayout::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) SetCurrentIndex(index int) {
	defer qt.Recovering("QStackedLayout::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) SetCurrentWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStackedLayout::setCurrentWidget")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	defer qt.Recovering("QStackedLayout::setStackingMode")

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetStackingMode(ptr.Pointer(), C.int(stackingMode))
	}
}

func (ptr *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	defer qt.Recovering("QStackedLayout::stackingMode")

	if ptr.Pointer() != nil {
		return QStackedLayout__StackingMode(C.QStackedLayout_StackingMode(ptr.Pointer()))
	}
	return 0
}

func NewQStackedLayout() *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout())
}

func NewQStackedLayout3(parentLayout QLayout_ITF) *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout3(PointerFromQLayout(parentLayout)))
}

func NewQStackedLayout2(parent QWidget_ITF) *QStackedLayout {
	defer qt.Recovering("QStackedLayout::QStackedLayout")

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout2(PointerFromQWidget(parent)))
}

func (ptr *QStackedLayout) ConnectAddItem(f func(item *QLayoutItem)) {
	defer qt.Recovering("connect QStackedLayout::addItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addItem", f)
	}
}

func (ptr *QStackedLayout) DisconnectAddItem() {
	defer qt.Recovering("disconnect QStackedLayout::addItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addItem")
	}
}

//export callbackQStackedLayoutAddItem
func callbackQStackedLayoutAddItem(ptrName *C.char, item unsafe.Pointer) bool {
	defer qt.Recovering("callback QStackedLayout::addItem")

	var signal = qt.GetSignal(C.GoString(ptrName), "addItem")
	if signal != nil {
		defer signal.(func(*QLayoutItem))(NewQLayoutItemFromPointer(item))
		return true
	}
	return false

}

func (ptr *QStackedLayout) AddWidget(widget QWidget_ITF) int {
	defer qt.Recovering("QStackedLayout::addWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QStackedLayout::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedLayout) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QStackedLayout::currentChanged")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedLayoutCurrentChanged
func callbackQStackedLayoutCurrentChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedLayout::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedLayout) CurrentWidget() *QWidget {
	defer qt.Recovering("QStackedLayout::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) HasHeightForWidth() bool {
	defer qt.Recovering("QStackedLayout::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QStackedLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStackedLayout) HeightForWidth(width int) int {
	defer qt.Recovering("QStackedLayout::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QStackedLayout) InsertWidget(index int, widget QWidget_ITF) int {
	defer qt.Recovering("QStackedLayout::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ItemAt(index int) *QLayoutItem {
	defer qt.Recovering("QStackedLayout::itemAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) TakeAt(index int) *QLayoutItem {
	defer qt.Recovering("QStackedLayout::takeAt")

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) Widget(index int) *QWidget {
	defer qt.Recovering("QStackedLayout::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectWidgetRemoved(f func(index int)) {
	defer qt.Recovering("connect QStackedLayout::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedLayout) DisconnectWidgetRemoved() {
	defer qt.Recovering("disconnect QStackedLayout::widgetRemoved")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedLayoutWidgetRemoved
func callbackQStackedLayoutWidgetRemoved(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QStackedLayout::widgetRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "widgetRemoved")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QStackedLayout) DestroyQStackedLayout() {
	defer qt.Recovering("QStackedLayout::~QStackedLayout")

	if ptr.Pointer() != nil {
		C.QStackedLayout_DestroyQStackedLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
