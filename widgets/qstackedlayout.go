package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedLayout) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedLayout) SetCurrentWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::setCurrentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::setStackingMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetStackingMode(ptr.Pointer(), C.int(stackingMode))
	}
}

func (ptr *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::stackingMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QStackedLayout__StackingMode(C.QStackedLayout_StackingMode(ptr.Pointer()))
	}
	return 0
}

func NewQStackedLayout() *QStackedLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::QStackedLayout")
		}
	}()

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout())
}

func NewQStackedLayout3(parentLayout QLayout_ITF) *QStackedLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::QStackedLayout")
		}
	}()

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout3(PointerFromQLayout(parentLayout)))
}

func NewQStackedLayout2(parent QWidget_ITF) *QStackedLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::QStackedLayout")
		}
	}()

	return NewQStackedLayoutFromPointer(C.QStackedLayout_NewQStackedLayout2(PointerFromQWidget(parent)))
}

func (ptr *QStackedLayout) AddItem(item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QStackedLayout) AddWidget(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ConnectCurrentChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedLayout) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedLayoutCurrentChanged
func callbackQStackedLayoutCurrentChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QStackedLayout) CurrentWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::currentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedLayout) HasHeightForWidth() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::hasHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QStackedLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStackedLayout) HeightForWidth(width int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QStackedLayout) InsertWidget(index int, widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedLayout_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedLayout) ItemAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) SetGeometry(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QStackedLayout) TakeAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::takeAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QStackedLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) Widget(index int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedLayout_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedLayout) ConnectWidgetRemoved(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::widgetRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedLayout) DisconnectWidgetRemoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::widgetRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedLayoutWidgetRemoved
func callbackQStackedLayoutWidgetRemoved(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::widgetRemoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "widgetRemoved").(func(int))(int(index))
}

func (ptr *QStackedLayout) DestroyQStackedLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedLayout::~QStackedLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedLayout_DestroyQStackedLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
