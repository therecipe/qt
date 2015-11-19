package widgets

//#include "qstackedlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QStackedLayout_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) SetCurrentWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetStackingMode(ptr.Pointer(), C.int(stackingMode))
	}
}

func (ptr *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	if ptr.Pointer() != nil {
		return QStackedLayout__StackingMode(C.QStackedLayout_StackingMode(ptr.Pointer()))
	}
	return 0
}

func NewQStackedLayout() *QStackedLayout {
	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout())
}

func NewQStackedLayout3(parentLayout QLayout_ITF) *QStackedLayout {
	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout3(PointerFromQLayout(parentLayout)))
}

func NewQStackedLayout2(parent QWidget_ITF) *QStackedLayout {
	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout2(PointerFromQWidget(parent)))
}

func (ptr *QStackedLayout) AddItem(item QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QStackedLayout) AddWidget(widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedLayout) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedLayoutCurrentChanged
func callbackQStackedLayoutCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QStackedLayout) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QStackedLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStackedLayout) HeightForWidth(width int) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QStackedLayout) InsertWidget(index int, widget QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) SetGeometry(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QStackedLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectWidgetRemoved(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedLayout) DisconnectWidgetRemoved() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedLayoutWidgetRemoved
func callbackQStackedLayoutWidgetRemoved(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "widgetRemoved").(func(int))(int(index))
}

func (ptr *QStackedLayout) DestroyQStackedLayout() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DestroyQStackedLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
