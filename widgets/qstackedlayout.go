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

type QStackedLayoutITF interface {
	QLayoutITF
	QStackedLayoutPTR() *QStackedLayout
}

func PointerFromQStackedLayout(ptr QStackedLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedLayoutPTR().Pointer()
	}
	return nil
}

func QStackedLayoutFromPointer(ptr unsafe.Pointer) *QStackedLayout {
	var n = new(QStackedLayout)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStackedLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStackedLayout) QStackedLayoutPTR() *QStackedLayout {
	return ptr
}

//QStackedLayout::StackingMode
type QStackedLayout__StackingMode int

var (
	QStackedLayout__StackOne = QStackedLayout__StackingMode(0)
	QStackedLayout__StackAll = QStackedLayout__StackingMode(1)
)

func (ptr *QStackedLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStackedLayout) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QStackedLayout) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QStackedLayout) SetCurrentWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetStackingMode(C.QtObjectPtr(ptr.Pointer()), C.int(stackingMode))
	}
}

func (ptr *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	if ptr.Pointer() != nil {
		return QStackedLayout__StackingMode(C.QStackedLayout_StackingMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQStackedLayout() *QStackedLayout {
	return QStackedLayoutFromPointer(unsafe.Pointer(C.QStackedLayout_NewQStackedLayout()))
}

func NewQStackedLayout3(parentLayout QLayoutITF) *QStackedLayout {
	return QStackedLayoutFromPointer(unsafe.Pointer(C.QStackedLayout_NewQStackedLayout3(C.QtObjectPtr(PointerFromQLayout(parentLayout)))))
}

func NewQStackedLayout2(parent QWidgetITF) *QStackedLayout {
	return QStackedLayoutFromPointer(unsafe.Pointer(C.QStackedLayout_NewQStackedLayout2(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QStackedLayout) AddItem(item QLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayoutItem(item)))
	}
}

func (ptr *QStackedLayout) AddWidget(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QStackedLayout) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedLayout) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedLayoutCurrentChanged
func callbackQStackedLayoutCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QStackedLayout) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QStackedLayout_CurrentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QStackedLayout) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QStackedLayout_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStackedLayout) HeightForWidth(width int) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width)))
	}
	return 0
}

func (ptr *QStackedLayout) InsertWidget(index int, widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QStackedLayout) ItemAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QStackedLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QStackedLayout) SetGeometry(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QStackedLayout) TakeAt(index int) *QLayoutItem {
	if ptr.Pointer() != nil {
		return QLayoutItemFromPointer(unsafe.Pointer(C.QStackedLayout_TakeAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QStackedLayout) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QStackedLayout_Widget(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectWidgetRemoved(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectWidgetRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedLayout) DisconnectWidgetRemoved() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectWidgetRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedLayoutWidgetRemoved
func callbackQStackedLayoutWidgetRemoved(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "widgetRemoved").(func(int))(int(index))
}

func (ptr *QStackedLayout) DestroyQStackedLayout() {
	if ptr.Pointer() != nil {
		C.QStackedLayout_DestroyQStackedLayout(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
