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

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemInserted"); signal != nil {
		signal.(func(int))(int(index))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "itemRemoved"); signal != nil {
		signal.(func(int))(int(index))
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

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(e))
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
func callbackQToolBoxPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

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
func callbackQToolBoxActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QToolBox::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQToolBoxCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQToolBoxInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQToolBoxCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QToolBox::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
