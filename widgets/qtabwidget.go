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

type QTabWidgetITF interface {
	QWidgetITF
	QTabWidgetPTR() *QTabWidget
}

func PointerFromQTabWidget(ptr QTabWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabWidgetPTR().Pointer()
	}
	return nil
}

func QTabWidgetFromPointer(ptr unsafe.Pointer) *QTabWidget {
	var n = new(QTabWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTabWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTabWidget) QTabWidgetPTR() *QTabWidget {
	return ptr
}

//QTabWidget::TabPosition
type QTabWidget__TabPosition int

var (
	QTabWidget__North = QTabWidget__TabPosition(0)
	QTabWidget__South = QTabWidget__TabPosition(1)
	QTabWidget__West  = QTabWidget__TabPosition(2)
	QTabWidget__East  = QTabWidget__TabPosition(3)
)

//QTabWidget::TabShape
type QTabWidget__TabShape int

var (
	QTabWidget__Rounded    = QTabWidget__TabShape(0)
	QTabWidget__Triangular = QTabWidget__TabShape(1)
)

func (ptr *QTabWidget) AddTab2(page QWidgetITF, icon gui.QIconITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(page)), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) AddTab(page QWidgetITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_AddTab(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(page)), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabWidget) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabWidget) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_DocumentMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabWidget) ElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QTabWidget_ElideMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabWidget) InsertTab2(index int, page QWidgetITF, icon gui.QIconITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(page)), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) InsertTab(index int, page QWidgetITF, label string) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_InsertTab(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(page)), C.CString(label)))
	}
	return 0
}

func (ptr *QTabWidget) IsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_IsMovable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabWidget) SetCornerWidget(widget QWidgetITF, corner core.Qt__Corner) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCornerWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(corner))
	}
}

func (ptr *QTabWidget) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QTabWidget) SetDocumentMode(set bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetDocumentMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QTabWidget) SetElideMode(v core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetElideMode(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QTabWidget) SetIconSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QTabWidget) SetMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetMovable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QTabWidget) SetTabBarAutoHide(enabled bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabBarAutoHide(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QTabWidget) SetTabPosition(v QTabWidget__TabPosition) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabPosition(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QTabWidget) SetTabShape(s QTabWidget__TabShape) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabShape(C.QtObjectPtr(ptr.Pointer()), C.int(s))
	}
}

func (ptr *QTabWidget) SetTabsClosable(closeable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabsClosable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(closeable)))
	}
}

func (ptr *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetUsesScrollButtons(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(useButtons)))
	}
}

func (ptr *QTabWidget) TabBarAutoHide() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_TabBarAutoHide(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabWidget) TabPosition() QTabWidget__TabPosition {
	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QTabWidget_TabPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabWidget) TabShape() QTabWidget__TabShape {
	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QTabWidget_TabShape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabWidget) TabsClosable() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_TabsClosable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabWidget) UsesScrollButtons() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_UsesScrollButtons(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQTabWidget(parent QWidgetITF) *QTabWidget {
	return QTabWidgetFromPointer(unsafe.Pointer(C.QTabWidget_NewQTabWidget(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QTabWidget) Clear() {
	if ptr.Pointer() != nil {
		C.QTabWidget_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTabWidget) CornerWidget(corner core.Qt__Corner) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTabWidget_CornerWidget(C.QtObjectPtr(ptr.Pointer()), C.int(corner))))
	}
	return nil
}

func (ptr *QTabWidget) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QTabWidget) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQTabWidgetCurrentChanged
func callbackQTabWidgetCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QTabWidget) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTabWidget_CurrentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTabWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTabWidget) HeightForWidth(width int) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width)))
	}
	return 0
}

func (ptr *QTabWidget) IndexOf(w QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QTabWidget_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w))))
	}
	return 0
}

func (ptr *QTabWidget) IsTabEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QTabWidget_IsTabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QTabWidget) RemoveTab(index int) {
	if ptr.Pointer() != nil {
		C.QTabWidget_RemoveTab(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QTabWidget) SetCurrentWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetCurrentWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QTabWidget) SetTabEnabled(index int, enable bool) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTabWidget) SetTabIcon(index int, icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabIcon(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QTabWidget) SetTabText(index int, label string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabText(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(label))
	}
}

func (ptr *QTabWidget) SetTabToolTip(index int, tip string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(tip))
	}
}

func (ptr *QTabWidget) SetTabWhatsThis(index int, text string) {
	if ptr.Pointer() != nil {
		C.QTabWidget_SetTabWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text))
	}
}

func (ptr *QTabWidget) TabBar() *QTabBar {
	if ptr.Pointer() != nil {
		return QTabBarFromPointer(unsafe.Pointer(C.QTabWidget_TabBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTabWidget) ConnectTabBarClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabBarClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarClicked() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarClicked")
	}
}

//export callbackQTabWidgetTabBarClicked
func callbackQTabWidgetTabBarClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarClicked").(func(int))(int(index))
}

func (ptr *QTabWidget) ConnectTabBarDoubleClicked(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabBarDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabBarDoubleClicked", f)
	}
}

func (ptr *QTabWidget) DisconnectTabBarDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabBarDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabBarDoubleClicked")
	}
}

//export callbackQTabWidgetTabBarDoubleClicked
func callbackQTabWidgetTabBarDoubleClicked(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabBarDoubleClicked").(func(int))(int(index))
}

func (ptr *QTabWidget) ConnectTabCloseRequested(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QTabWidget_ConnectTabCloseRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "tabCloseRequested", f)
	}
}

func (ptr *QTabWidget) DisconnectTabCloseRequested() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DisconnectTabCloseRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "tabCloseRequested")
	}
}

//export callbackQTabWidgetTabCloseRequested
func callbackQTabWidgetTabCloseRequested(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "tabCloseRequested").(func(int))(int(index))
}

func (ptr *QTabWidget) TabText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabText(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabToolTip(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) TabWhatsThis(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTabWidget_TabWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QTabWidget) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QTabWidget_Widget(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QTabWidget) DestroyQTabWidget() {
	if ptr.Pointer() != nil {
		C.QTabWidget_DestroyQTabWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
