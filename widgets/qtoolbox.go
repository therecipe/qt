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
func callbackQToolBoxChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQToolBoxFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QToolBox) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QToolBox) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
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
func callbackQToolBoxCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QToolBox::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QToolBox) CurrentChanged(index int) {
	defer qt.Recovering("QToolBox::currentChanged")

	if ptr.Pointer() != nil {
		C.QToolBox_CurrentChanged(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QToolBox) CurrentWidget() *QWidget {
	defer qt.Recovering("QToolBox::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QToolBox_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolBox) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QToolBox::event")

	if ptr.Pointer() != nil {
		return C.QToolBox_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQToolBoxItemInserted(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QToolBox::itemInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemInserted"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQToolBoxFromPointer(ptr).ItemInsertedDefault(int(index))
	}
}

func (ptr *QToolBox) ItemInserted(index int) {
	defer qt.Recovering("QToolBox::itemInserted")

	if ptr.Pointer() != nil {
		C.QToolBox_ItemInserted(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QToolBox) ItemInsertedDefault(index int) {
	defer qt.Recovering("QToolBox::itemInserted")

	if ptr.Pointer() != nil {
		C.QToolBox_ItemInsertedDefault(ptr.Pointer(), C.int(index))
	}
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
func callbackQToolBoxItemRemoved(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QToolBox::itemRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemRemoved"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQToolBoxFromPointer(ptr).ItemRemovedDefault(int(index))
	}
}

func (ptr *QToolBox) ItemRemoved(index int) {
	defer qt.Recovering("QToolBox::itemRemoved")

	if ptr.Pointer() != nil {
		C.QToolBox_ItemRemoved(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QToolBox) ItemRemovedDefault(index int) {
	defer qt.Recovering("QToolBox::itemRemoved")

	if ptr.Pointer() != nil {
		C.QToolBox_ItemRemovedDefault(ptr.Pointer(), C.int(index))
	}
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
func callbackQToolBoxShowEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
	} else {
		NewQToolBoxFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(e))
	}
}

func (ptr *QToolBox) ShowEvent(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolBox::showEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
}

func (ptr *QToolBox) ShowEventDefault(e gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolBox::showEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(e))
	}
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

func (ptr *QToolBox) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QToolBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QToolBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QToolBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQToolBoxPaintEvent
func callbackQToolBoxPaintEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
	} else {
		NewQToolBoxFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(v))
	}
}

func (ptr *QToolBox) PaintEvent(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QToolBox) PaintEventDefault(v gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolBox::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(v))
	}
}

func (ptr *QToolBox) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QToolBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QToolBox) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QToolBox::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQToolBoxActionEvent
func callbackQToolBoxActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QToolBox) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QToolBox) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolBox::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QToolBox) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QToolBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QToolBox) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QToolBox::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQToolBoxDragEnterEvent
func callbackQToolBoxDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QToolBox) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolBox) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolBox::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolBox) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QToolBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QToolBox) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QToolBox::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQToolBoxDragLeaveEvent
func callbackQToolBoxDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QToolBox) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolBox) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolBox::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolBox) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QToolBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QToolBox) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QToolBox::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQToolBoxDragMoveEvent
func callbackQToolBoxDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QToolBox) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolBox) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolBox::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolBox) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QToolBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QToolBox) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QToolBox::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQToolBoxDropEvent
func callbackQToolBoxDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QToolBox) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolBox) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolBox::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolBox) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QToolBox) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QToolBox::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQToolBoxEnterEvent
func callbackQToolBoxEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBox) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBox) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBox) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QToolBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QToolBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQToolBoxFocusInEvent
func callbackQToolBoxFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QToolBox) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBox) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBox::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBox) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QToolBox) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QToolBox::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQToolBoxFocusOutEvent
func callbackQToolBoxFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QToolBox) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBox) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolBox::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QToolBox) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QToolBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QToolBox) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QToolBox::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQToolBoxHideEvent
func callbackQToolBoxHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QToolBox) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolBox) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolBox::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolBox) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QToolBox) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QToolBox::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQToolBoxLeaveEvent
func callbackQToolBoxLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBox) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBox) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBox) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QToolBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QToolBox) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QToolBox::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQToolBoxMoveEvent
func callbackQToolBoxMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QToolBox) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolBox) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolBox::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolBox) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QToolBox::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QToolBox) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QToolBox::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQToolBoxSetVisible
func callbackQToolBoxSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QToolBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QToolBox) SetVisible(visible bool) {
	defer qt.Recovering("QToolBox::setVisible")

	if ptr.Pointer() != nil {
		C.QToolBox_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolBox) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QToolBox::setVisible")

	if ptr.Pointer() != nil {
		C.QToolBox_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolBox) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QToolBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QToolBox) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QToolBox::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQToolBoxCloseEvent
func callbackQToolBoxCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QToolBox) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolBox) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolBox::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolBox) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QToolBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QToolBox) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QToolBox::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQToolBoxContextMenuEvent
func callbackQToolBoxContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QToolBox) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolBox) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolBox::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolBox) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QToolBox::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QToolBox) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QToolBox::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQToolBoxInitPainter
func callbackQToolBoxInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQToolBoxFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QToolBox) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolBox::initPainter")

	if ptr.Pointer() != nil {
		C.QToolBox_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolBox) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolBox::initPainter")

	if ptr.Pointer() != nil {
		C.QToolBox_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolBox) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QToolBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QToolBox) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QToolBox::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQToolBoxInputMethodEvent
func callbackQToolBoxInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QToolBox) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolBox) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolBox::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolBox) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QToolBox) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QToolBox::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQToolBoxKeyPressEvent
func callbackQToolBoxKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QToolBox) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBox) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBox::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBox) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QToolBox) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QToolBox::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQToolBoxKeyReleaseEvent
func callbackQToolBoxKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QToolBox) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBox) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolBox::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QToolBox) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QToolBox) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QToolBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQToolBoxMouseDoubleClickEvent
func callbackQToolBoxMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBox) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QToolBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QToolBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQToolBoxMouseMoveEvent
func callbackQToolBoxMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBox) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QToolBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QToolBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQToolBoxMousePressEvent
func callbackQToolBoxMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBox) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QToolBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QToolBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQToolBoxMouseReleaseEvent
func callbackQToolBoxMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolBox) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolBox) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QToolBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QToolBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QToolBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQToolBoxResizeEvent
func callbackQToolBoxResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QToolBox) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolBox) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolBox::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolBox) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QToolBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QToolBox) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QToolBox::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQToolBoxTabletEvent
func callbackQToolBoxTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QToolBox) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolBox) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolBox::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolBox) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QToolBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QToolBox) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QToolBox::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQToolBoxWheelEvent
func callbackQToolBoxWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QToolBox) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolBox) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolBox::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolBox) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QToolBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QToolBox) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QToolBox::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQToolBoxTimerEvent
func callbackQToolBoxTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QToolBox) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QToolBox) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolBox::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QToolBox) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QToolBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QToolBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QToolBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQToolBoxChildEvent
func callbackQToolBoxChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QToolBox) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolBox::childEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolBox) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolBox::childEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolBox) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolBox::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QToolBox) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QToolBox::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQToolBoxCustomEvent
func callbackQToolBoxCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolBoxFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolBox) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::customEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolBox) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolBox::customEvent")

	if ptr.Pointer() != nil {
		C.QToolBox_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
