package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolBox struct {
	QFrame
}

type QToolBox_ITF interface {
	QFrame_ITF
	QToolBox_PTR() *QToolBox
}

func PointerFromQToolBox(ptr QToolBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBox_PTR().Pointer()
	}
	return nil
}

func NewQToolBoxFromPointer(ptr unsafe.Pointer) *QToolBox {
	var n = new(QToolBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolBox_") {
		n.SetObjectName("QToolBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QToolBox) QToolBox_PTR() *QToolBox {
	return ptr
}

func (ptr *QToolBox) Count() int {
	defer qt.Recovering("QToolBox::count")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBox) CurrentIndex() int {
	defer qt.Recovering("QToolBox::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolBox) SetCurrentIndex(index int) {
	defer qt.Recovering("QToolBox::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func NewQToolBox(parent QWidget_ITF, f core.Qt__WindowType) *QToolBox {
	defer qt.Recovering("QToolBox::QToolBox")

	return NewQToolBoxFromPointer(C.QToolBox_NewQToolBox(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QToolBox) AddItem2(w QWidget_ITF, text string) int {
	defer qt.Recovering("QToolBox::addItem")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem2(ptr.Pointer(), PointerFromQWidget(w), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) AddItem(widget QWidget_ITF, iconSet gui.QIcon_ITF, text string) int {
	defer qt.Recovering("QToolBox::addItem")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem(ptr.Pointer(), PointerFromQWidget(widget), gui.PointerFromQIcon(iconSet), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QToolBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QToolBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QToolBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQToolBoxChangeEvent
func callbackQToolBoxChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QToolBox) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QToolBox::currentChanged")

	if ptr.Pointer() != nil {
		C.QToolBox_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QToolBox) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QToolBox::currentChanged")

	if ptr.Pointer() != nil {
		C.QToolBox_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQToolBoxCurrentChanged
func callbackQToolBoxCurrentChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QToolBox::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QToolBox) CurrentWidget() *QWidget {
	defer qt.Recovering("QToolBox::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBox_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBox) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QToolBox::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QToolBox) InsertItem(index int, widget QWidget_ITF, icon gui.QIcon_ITF, text string) int {
	defer qt.Recovering("QToolBox::insertItem")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) InsertItem2(index int, widget QWidget_ITF, text string) int {
	defer qt.Recovering("QToolBox::insertItem")

	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem2(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) IsItemEnabled(index int) bool {
	defer qt.Recovering("QToolBox::isItemEnabled")

	if ptr.Pointer() != nil {
		return C.QToolBox_IsItemEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QToolBox) ItemIcon(index int) *gui.QIcon {
	defer qt.Recovering("QToolBox::itemIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QToolBox_ItemIcon(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QToolBox) ConnectItemInserted(f func(index int)) {
	defer qt.Recovering("connect QToolBox::itemInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "itemInserted", f)
	}
}

func (ptr *QToolBox) DisconnectItemInserted() {
	defer qt.Recovering("disconnect QToolBox::itemInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "itemInserted")
	}
}

//export callbackQToolBoxItemInserted
func callbackQToolBoxItemInserted(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QToolBox::itemInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemInserted")
	if signal != nil {
		defer signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QToolBox) ConnectItemRemoved(f func(index int)) {
	defer qt.Recovering("connect QToolBox::itemRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "itemRemoved", f)
	}
}

func (ptr *QToolBox) DisconnectItemRemoved() {
	defer qt.Recovering("disconnect QToolBox::itemRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "itemRemoved")
	}
}

//export callbackQToolBoxItemRemoved
func callbackQToolBoxItemRemoved(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QToolBox::itemRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "itemRemoved")
	if signal != nil {
		defer signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QToolBox) ItemText(index int) string {
	defer qt.Recovering("QToolBox::itemText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) ItemToolTip(index int) string {
	defer qt.Recovering("QToolBox::itemToolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) RemoveItem(index int) {
	defer qt.Recovering("QToolBox::removeItem")

	if ptr.Pointer() != nil {
		C.QToolBox_RemoveItem(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QToolBox) SetCurrentWidget(widget QWidget_ITF) {
	defer qt.Recovering("QToolBox::setCurrentWidget")

	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QToolBox) SetItemEnabled(index int, enabled bool) {
	defer qt.Recovering("QToolBox::setItemEnabled")

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QToolBox) SetItemIcon(index int, icon gui.QIcon_ITF) {
	defer qt.Recovering("QToolBox::setItemIcon")

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QToolBox) SetItemText(index int, text string) {
	defer qt.Recovering("QToolBox::setItemText")

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemText(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QToolBox) SetItemToolTip(index int, toolTip string) {
	defer qt.Recovering("QToolBox::setItemToolTip")

	if ptr.Pointer() != nil {
		C.QToolBox_SetItemToolTip(ptr.Pointer(), C.int(index), C.CString(toolTip))
	}
}

func (ptr *QToolBox) ConnectShowEvent(f func(e *gui.QShowEvent)) {
	defer qt.Recovering("connect QToolBox::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QToolBox) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QToolBox::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQToolBoxShowEvent
func callbackQToolBoxShowEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QToolBox) Widget(index int) *QWidget {
	defer qt.Recovering("QToolBox::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBox_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QToolBox) DestroyQToolBox() {
	defer qt.Recovering("QToolBox::~QToolBox")

	if ptr.Pointer() != nil {
		C.QToolBox_DestroyQToolBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
