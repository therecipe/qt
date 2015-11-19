package widgets

//#include "qtabwidget.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QTabWidget_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab2(ptr.Pointer(), PointerFromQWidget(page), gui.PointerFromQIcon(icon), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) AddTab(page QWidget_ITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab(ptr.Pointer(), PointerFromQWidget(page), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) ElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabWidget_ElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) InsertTab2(index int, page QWidget_ITF, icon gui.QIcon_ITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab2(ptr.Pointer(), C.int(index), PointerFromQWidget(page), gui.PointerFromQIcon(icon), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) InsertTab(index int, page QWidget_ITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab(ptr.Pointer(), C.int(index), PointerFromQWidget(page), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) IsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_IsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) SetCornerWidget(widget QWidget_ITF, corner core.Qt__Corner) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(corner))
	}
}

func (ptr *QTabWidget) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabWidget) SetDocumentMode(set bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabWidget) SetElideMode(v core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetElideMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabWidget) SetIconSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QTabWidget) SetMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabWidget) SetTabBarAutoHide(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabBarAutoHide(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabWidget) SetTabPosition(v QTabWidget__TabPosition) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QTabWidget) SetTabShape(s QTabWidget__TabShape) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabShape(ptr.Pointer(), C.int(s))
	}
}

func (ptr *QTabWidget) SetTabsClosable(closeable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closeable)))
	}
}

func (ptr *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetUsesScrollButtons(ptr.Pointer(), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabWidget) TabBarAutoHide() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_TabBarAutoHide(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) TabPosition() QTabWidget__TabPosition {
	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QTabWidget_TabPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) TabShape() QTabWidget__TabShape {
	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QTabWidget_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabWidget) TabsClosable() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) UsesScrollButtons() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_UsesScrollButtons(ptr.Pointer()) != 0
	}
	return false
}

func NewQTabWidget(parent QWidget_ITF) *QTabWidget {
	return NewQTabWidgetFromPointer(C.QTabWidget_NewQTabWidget(PointerFromQWidget(parent)))
}

func (ptr *QTabWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QTabWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTabWidget) CornerWidget(corner core.Qt__Corner) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_CornerWidget(ptr.Pointer(), C.int(corner)))
	}
	return nil
}

func (ptr *QTabWidget) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabWidget) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabWidgetCurrentChanged
func callbackQTabWidgetCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QTabWidget) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTabWidget) HeightForWidth(width int) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QTabWidget) IndexOf(w QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_IndexOf(ptr.Pointer(), PointerFromQWidget(w)))
	}
	return 0
}

func (ptr *QTabWidget) IsTabEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_IsTabEnabled(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabWidget) RemoveTab(index int) {
	if ptr.Pointer() != nil {
		C.QTabWidget_RemoveTab(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QTabWidget) SetCurrentWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QTabWidget) SetTabEnabled(index int, enable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabEnabled(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTabWidget) SetTabIcon(index int, icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabIcon(ptr.Pointer(), C.int(index), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTabWidget) SetTabText(index int, label string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabText(ptr.Pointer(), C.int(index), C.CString(label))
	}
}

func (ptr *QTabWidget) SetTabToolTip(index int, tip string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabToolTip(ptr.Pointer(), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabWidget) SetTabWhatsThis(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabWhatsThis(ptr.Pointer(), C.int(index), C.CString(text))
	}
}

func (ptr *QTabWidget) TabBar() *QTabBar {
	if ptr.Pointer() != nil {
		return NewQTabBarFromPointer(C.QTabWidget_TabBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTabWidget) ConnectTabBarClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarClicked() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabWidgetTabBarClicked
func callbackQTabWidgetTabBarClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarClicked").(func(int))(int(index))
}

func (ptr *QTabWidget) ConnectTabBarDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabWidgetTabBarDoubleClicked
func callbackQTabWidgetTabBarDoubleClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked").(func(int))(int(index))
}

func (ptr *QTabWidget) ConnectTabCloseRequested(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabWidget) DisconnectTabCloseRequested() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabWidgetTabCloseRequested
func callbackQTabWidgetTabCloseRequested(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabCloseRequested").(func(int))(int(index))
}

func (ptr *QTabWidget) TabText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabText(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabToolTip(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabToolTip(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabWhatsThis(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabWhatsThis(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTabWidget_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTabWidget) DestroyQTabWidget() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DestroyQTabWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
