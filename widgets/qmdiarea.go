package widgets

//#include "qmdiarea.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMdiArea struct {
	QAbstractScrollArea
}

type QMdiAreaITF interface {
	QAbstractScrollAreaITF
	QMdiAreaPTR() *QMdiArea
}

func PointerFromQMdiArea(ptr QMdiAreaITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiAreaPTR().Pointer()
	}
	return nil
}

func QMdiAreaFromPointer(ptr unsafe.Pointer) *QMdiArea {
	var n = new(QMdiArea)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMdiArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMdiArea) QMdiAreaPTR() *QMdiArea {
	return ptr
}

//QMdiArea::AreaOption
type QMdiArea__AreaOption int

var (
	QMdiArea__DontMaximizeSubWindowOnActivation = QMdiArea__AreaOption(0x1)
)

//QMdiArea::ViewMode
type QMdiArea__ViewMode int

var (
	QMdiArea__SubWindowView = QMdiArea__ViewMode(0)
	QMdiArea__TabbedView    = QMdiArea__ViewMode(1)
)

//QMdiArea::WindowOrder
type QMdiArea__WindowOrder int

var (
	QMdiArea__CreationOrder          = QMdiArea__WindowOrder(0)
	QMdiArea__StackingOrder          = QMdiArea__WindowOrder(1)
	QMdiArea__ActivationHistoryOrder = QMdiArea__WindowOrder(2)
)

func (ptr *QMdiArea) ActivationOrder() QMdiArea__WindowOrder {
	if ptr.Pointer() != nil {
		return QMdiArea__WindowOrder(C.QMdiArea_ActivationOrder(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMdiArea) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QMdiArea_DocumentMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMdiArea) SetActivationOrder(order QMdiArea__WindowOrder) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetActivationOrder(C.QtObjectPtr(ptr.Pointer()), C.int(order))
	}
}

func (ptr *QMdiArea) SetBackground(background gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(background)))
	}
}

func (ptr *QMdiArea) SetDocumentMode(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetDocumentMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabPosition(C.QtObjectPtr(ptr.Pointer()), C.int(position))
	}
}

func (ptr *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabShape(C.QtObjectPtr(ptr.Pointer()), C.int(shape))
	}
}

func (ptr *QMdiArea) SetTabsClosable(closable bool) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsClosable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QMdiArea) SetTabsMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsMovable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QMdiArea) SetViewMode(mode QMdiArea__ViewMode) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetViewMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QMdiArea) TabPosition() QTabWidget__TabPosition {
	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMdiArea_TabPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMdiArea) TabShape() QTabWidget__TabShape {
	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMdiArea_TabShape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMdiArea) TabsClosable() bool {
	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsClosable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMdiArea) TabsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsMovable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMdiArea) ViewMode() QMdiArea__ViewMode {
	if ptr.Pointer() != nil {
		return QMdiArea__ViewMode(C.QMdiArea_ViewMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQMdiArea(parent QWidgetITF) *QMdiArea {
	return QMdiAreaFromPointer(unsafe.Pointer(C.QMdiArea_NewQMdiArea(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QMdiArea) ActivateNextSubWindow() {
	if ptr.Pointer() != nil {
		C.QMdiArea_ActivateNextSubWindow(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) ActivatePreviousSubWindow() {
	if ptr.Pointer() != nil {
		C.QMdiArea_ActivatePreviousSubWindow(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	if ptr.Pointer() != nil {
		return QMdiSubWindowFromPointer(unsafe.Pointer(C.QMdiArea_ActiveSubWindow(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMdiArea) AddSubWindow(widget QWidgetITF, windowFlags core.Qt__WindowType) *QMdiSubWindow {
	if ptr.Pointer() != nil {
		return QMdiSubWindowFromPointer(unsafe.Pointer(C.QMdiArea_AddSubWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(windowFlags))))
	}
	return nil
}

func (ptr *QMdiArea) CascadeSubWindows() {
	if ptr.Pointer() != nil {
		C.QMdiArea_CascadeSubWindows(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) CloseActiveSubWindow() {
	if ptr.Pointer() != nil {
		C.QMdiArea_CloseActiveSubWindow(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) CloseAllSubWindows() {
	if ptr.Pointer() != nil {
		C.QMdiArea_CloseAllSubWindows(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	if ptr.Pointer() != nil {
		return QMdiSubWindowFromPointer(unsafe.Pointer(C.QMdiArea_CurrentSubWindow(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMdiArea) RemoveSubWindow(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMdiArea_RemoveSubWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QMdiArea) SetActiveSubWindow(window QMdiSubWindowITF) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetActiveSubWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMdiSubWindow(window)))
	}
}

func (ptr *QMdiArea) SetOption(option QMdiArea__AreaOption, on bool) {
	if ptr.Pointer() != nil {
		C.QMdiArea_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiArea) ConnectSubWindowActivated(f func(window QMdiSubWindowITF)) {
	if ptr.Pointer() != nil {
		C.QMdiArea_ConnectSubWindowActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "subWindowActivated", f)
	}
}

func (ptr *QMdiArea) DisconnectSubWindowActivated() {
	if ptr.Pointer() != nil {
		C.QMdiArea_DisconnectSubWindowActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "subWindowActivated")
	}
}

//export callbackQMdiAreaSubWindowActivated
func callbackQMdiAreaSubWindowActivated(ptrName *C.char, window unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "subWindowActivated").(func(*QMdiSubWindow))(QMdiSubWindowFromPointer(window))
}

func (ptr *QMdiArea) TestOption(option QMdiArea__AreaOption) bool {
	if ptr.Pointer() != nil {
		return C.QMdiArea_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiArea) TileSubWindows() {
	if ptr.Pointer() != nil {
		C.QMdiArea_TileSubWindows(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMdiArea) DestroyQMdiArea() {
	if ptr.Pointer() != nil {
		C.QMdiArea_DestroyQMdiArea(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
