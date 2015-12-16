package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QTabWidget struct {
	QWidget
}

type QTabWidget_ITF interface {
	QWidget_ITF
	QTabWidget_PTR() *QTabWidget
}

func PointerFromQTabWidget(ptr QTabWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabWidget_PTR().Pointer()
	}
	return nil
}

func NewQTabWidgetFromPointer(ptr unsafe.Pointer) *QTabWidget {
	var n = new(QTabWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTabWidget_") {
		n.SetObjectName("QTabWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QTabWidget) QTabWidget_PTR() *QTabWidget {
	return ptr
}

//QTabWidget::TabPosition
type QTabWidget__TabPosition int64

const (
	QTabWidget__North = QTabWidget__TabPosition(0)
	QTabWidget__South = QTabWidget__TabPosition(1)
	QTabWidget__West  = QTabWidget__TabPosition(2)
	QTabWidget__East  = QTabWidget__TabPosition(3)
)

//QTabWidget::TabShape
type QTabWidget__TabShape int64

const (
	QTabWidget__Rounded    = QTabWidget__TabShape(0)
	QTabWidget__Triangular = QTabWidget__TabShape(1)
)

func (ptr *QTabWidget) AddTab2(page QWidget_ITF, icon gui.QIcon_ITF, label string) int {
	defer qt.Recovering("QTabWidget::addTab")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab2(ptr.Pointer(), PointerFromQWidget(page), gui.PointerFromQIcon(icon), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) AddTab(page QWidget_ITF, label string) int {
	defer qt.Recovering("QTabWidget::addTab")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab(ptr.Pointer(), PointerFromQWidget(page), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) Count() int {
	defer qt.Recovering("QTabWidget::count")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) CurrentIndex() int {
	defer qt.Recovering("QTabWidget::currentIndex")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) DocumentMode() bool {
	defer qt.Recovering("QTabWidget::documentMode")

	if ptr.Pointer() != nil {
		return C.QTabWidget_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) ElideMode() core.Qt__TextElideMode {
	defer qt.Recovering("QTabWidget::elideMode")

	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabWidget_ElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) IconSize() *core.QSize {
	defer qt.Recovering("QTabWidget::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabWidget_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) InsertTab2(index int, page QWidget_ITF, icon gui.QIcon_ITF, label string) int {
	defer qt.Recovering("QTabWidget::insertTab")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab2(ptr.Pointer(), C.int(index), PointerFromQWidget(page), gui.PointerFromQIcon(icon), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) InsertTab(index int, page QWidget_ITF, label string) int {
	defer qt.Recovering("QTabWidget::insertTab")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab(ptr.Pointer(), C.int(index), PointerFromQWidget(page), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) IsMovable() bool {
	defer qt.Recovering("QTabWidget::isMovable")

	if ptr.Pointer() != nil {
		return C.QTabWidget_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QTabWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QTabWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QTabWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQTabWidgetPaintEvent
func callbackQTabWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTabWidget::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTabWidget) SetCornerWidget(widget QWidget_ITF, corner core.Qt__Corner) {
	defer qt.Recovering("QTabWidget::setCornerWidget")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(corner))
	}
}

func (ptr *QTabWidget) SetCurrentIndex(index int) {
	defer qt.Recovering("QTabWidget::setCurrentIndex")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabWidget) SetDocumentMode(set bool) {
	defer qt.Recovering("QTabWidget::setDocumentMode")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabWidget) SetElideMode(v core.Qt__TextElideMode) {
	defer qt.Recovering("QTabWidget::setElideMode")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetElideMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabWidget) SetIconSize(size core.QSize_ITF) {
	defer qt.Recovering("QTabWidget::setIconSize")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTabWidget) SetMovable(movable bool) {
	defer qt.Recovering("QTabWidget::setMovable")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabWidget) SetTabBarAutoHide(enabled bool) {
	defer qt.Recovering("QTabWidget::setTabBarAutoHide")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabBarAutoHide(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabWidget) SetTabPosition(v QTabWidget__TabPosition) {
	defer qt.Recovering("QTabWidget::setTabPosition")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabWidget) SetTabShape(s QTabWidget__TabShape) {
	defer qt.Recovering("QTabWidget::setTabShape")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabShape(ptr.Pointer(), C.int(s))
	}
}

func (ptr *QTabWidget) SetTabsClosable(closeable bool) {
	defer qt.Recovering("QTabWidget::setTabsClosable")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closeable)))
	}
}

func (ptr *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	defer qt.Recovering("QTabWidget::setUsesScrollButtons")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetUsesScrollButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabWidget) TabBarAutoHide() bool {
	defer qt.Recovering("QTabWidget::tabBarAutoHide")

	if ptr.Pointer() != nil {
		return C.QTabWidget_TabBarAutoHide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) TabPosition() QTabWidget__TabPosition {
	defer qt.Recovering("QTabWidget::tabPosition")

	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QTabWidget_TabPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) TabShape() QTabWidget__TabShape {
	defer qt.Recovering("QTabWidget::tabShape")

	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QTabWidget_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) TabsClosable() bool {
	defer qt.Recovering("QTabWidget::tabsClosable")

	if ptr.Pointer() != nil {
		return C.QTabWidget_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) UsesScrollButtons() bool {
	defer qt.Recovering("QTabWidget::usesScrollButtons")

	if ptr.Pointer() != nil {
		return C.QTabWidget_UsesScrollButtons(ptr.Pointer()) != 0
	}
	return false
}

func NewQTabWidget(parent QWidget_ITF) *QTabWidget {
	defer qt.Recovering("QTabWidget::QTabWidget")

	return NewQTabWidgetFromPointer(C.QTabWidget_NewQTabWidget(PointerFromQWidget(parent)))
}

func (ptr *QTabWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QTabWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QTabWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QTabWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQTabWidgetChangeEvent
func callbackQTabWidgetChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QTabWidget::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QTabWidget) Clear() {
	defer qt.Recovering("QTabWidget::clear")

	if ptr.Pointer() != nil {
		C.QTabWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTabWidget) CornerWidget(corner core.Qt__Corner) *QWidget {
	defer qt.Recovering("QTabWidget::cornerWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_CornerWidget(ptr.Pointer(), C.int(corner)))
	}
	return nil
}

func (ptr *QTabWidget) ConnectCurrentChanged(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabWidget) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QTabWidget::currentChanged")

	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabWidgetCurrentChanged
func callbackQTabWidgetCurrentChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabWidget::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabWidget) CurrentWidget() *QWidget {
	defer qt.Recovering("QTabWidget::currentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QTabWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QTabWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) HeightForWidth(width int) int {
	defer qt.Recovering("QTabWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QTabWidget) IndexOf(w QWidget_ITF) int {
	defer qt.Recovering("QTabWidget::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QTabWidget_IndexOf(ptr.Pointer(), PointerFromQWidget(w)))
	}
	return 0
}

func (ptr *QTabWidget) IsTabEnabled(index int) bool {
	defer qt.Recovering("QTabWidget::isTabEnabled")

	if ptr.Pointer() != nil {
		return C.QTabWidget_IsTabEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabWidget) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QTabWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QTabWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QTabWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQTabWidgetKeyPressEvent
func callbackQTabWidgetKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTabWidget::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTabWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QTabWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabWidget_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) RemoveTab(index int) {
	defer qt.Recovering("QTabWidget::removeTab")

	if ptr.Pointer() != nil {
		C.QTabWidget_RemoveTab(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QTabWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QTabWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QTabWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQTabWidgetResizeEvent
func callbackQTabWidgetResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTabWidget::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTabWidget) SetCurrentWidget(widget QWidget_ITF) {
	defer qt.Recovering("QTabWidget::setCurrentWidget")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QTabWidget) SetTabEnabled(index int, enable bool) {
	defer qt.Recovering("QTabWidget::setTabEnabled")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTabWidget) SetTabIcon(index int, icon gui.QIcon_ITF) {
	defer qt.Recovering("QTabWidget::setTabIcon")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTabWidget) SetTabText(index int, label string) {
	defer qt.Recovering("QTabWidget::setTabText")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabText(ptr.Pointer(), C.int(index), C.CString(label))
	}
}

func (ptr *QTabWidget) SetTabToolTip(index int, tip string) {
	defer qt.Recovering("QTabWidget::setTabToolTip")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabToolTip(ptr.Pointer(), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabWidget) SetTabWhatsThis(index int, text string) {
	defer qt.Recovering("QTabWidget::setTabWhatsThis")

	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabWhatsThis(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabWidget) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QTabWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QTabWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QTabWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQTabWidgetShowEvent
func callbackQTabWidgetShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QTabWidget::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QTabWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QTabWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QTabWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) TabBar() *QTabBar {
	defer qt.Recovering("QTabWidget::tabBar")

	if ptr.Pointer() != nil {
		return NewQTabBarFromPointer(C.QTabWidget_TabBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) ConnectTabBarClicked(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::tabBarClicked")

	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarClicked() {
	defer qt.Recovering("disconnect QTabWidget::tabBarClicked")

	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabWidgetTabBarClicked
func callbackQTabWidgetTabBarClicked(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabWidget::tabBarClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "tabBarClicked")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabWidget) ConnectTabBarDoubleClicked(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::tabBarDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarDoubleClicked() {
	defer qt.Recovering("disconnect QTabWidget::tabBarDoubleClicked")

	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabWidgetTabBarDoubleClicked
func callbackQTabWidgetTabBarDoubleClicked(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabWidget::tabBarDoubleClicked")

	var signal = qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabWidget) ConnectTabCloseRequested(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::tabCloseRequested")

	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabWidget) DisconnectTabCloseRequested() {
	defer qt.Recovering("disconnect QTabWidget::tabCloseRequested")

	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabWidgetTabCloseRequested
func callbackQTabWidgetTabCloseRequested(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QTabWidget::tabCloseRequested")

	var signal = qt.GetSignal(C.GoString(ptrName), "tabCloseRequested")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QTabWidget) TabIcon(index int) *gui.QIcon {
	defer qt.Recovering("QTabWidget::tabIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QTabWidget_TabIcon(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabWidget) ConnectTabInserted(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::tabInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabInserted", f)
	}
}

func (ptr *QTabWidget) DisconnectTabInserted() {
	defer qt.Recovering("disconnect QTabWidget::tabInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabInserted")
	}
}

//export callbackQTabWidgetTabInserted
func callbackQTabWidgetTabInserted(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QTabWidget::tabInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "tabInserted")
	if signal != nil {
		defer signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QTabWidget) ConnectTabRemoved(f func(index int)) {
	defer qt.Recovering("connect QTabWidget::tabRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabRemoved", f)
	}
}

func (ptr *QTabWidget) DisconnectTabRemoved() {
	defer qt.Recovering("disconnect QTabWidget::tabRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabRemoved")
	}
}

//export callbackQTabWidgetTabRemoved
func callbackQTabWidgetTabRemoved(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QTabWidget::tabRemoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "tabRemoved")
	if signal != nil {
		defer signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QTabWidget) TabText(index int) string {
	defer qt.Recovering("QTabWidget::tabText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabToolTip(index int) string {
	defer qt.Recovering("QTabWidget::tabToolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabWhatsThis(index int) string {
	defer qt.Recovering("QTabWidget::tabWhatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabWhatsThis(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) Widget(index int) *QWidget {
	defer qt.Recovering("QTabWidget::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabWidget) DestroyQTabWidget() {
	defer qt.Recovering("QTabWidget::~QTabWidget")

	if ptr.Pointer() != nil {
		C.QTabWidget_DestroyQTabWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
