package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QMdiArea struct {
	QAbstractScrollArea
}

type QMdiArea_ITF interface {
	QAbstractScrollArea_ITF
	QMdiArea_PTR() *QMdiArea
}

func PointerFromQMdiArea(ptr QMdiArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMdiArea_PTR().Pointer()
	}
	return nil
}

func NewQMdiAreaFromPointer(ptr unsafe.Pointer) *QMdiArea {
	var n = new(QMdiArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMdiArea_") {
		n.SetObjectName("QMdiArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMdiArea) QMdiArea_PTR() *QMdiArea {
	return ptr
}

//QMdiArea::AreaOption
type QMdiArea__AreaOption int64

const (
	QMdiArea__DontMaximizeSubWindowOnActivation = QMdiArea__AreaOption(0x1)
)

//QMdiArea::ViewMode
type QMdiArea__ViewMode int64

const (
	QMdiArea__SubWindowView = QMdiArea__ViewMode(0)
	QMdiArea__TabbedView    = QMdiArea__ViewMode(1)
)

//QMdiArea::WindowOrder
type QMdiArea__WindowOrder int64

const (
	QMdiArea__CreationOrder          = QMdiArea__WindowOrder(0)
	QMdiArea__StackingOrder          = QMdiArea__WindowOrder(1)
	QMdiArea__ActivationHistoryOrder = QMdiArea__WindowOrder(2)
)

func (ptr *QMdiArea) ActivationOrder() QMdiArea__WindowOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::activationOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return QMdiArea__WindowOrder(C.QMdiArea_ActivationOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) Background() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::background")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QMdiArea_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) DocumentMode() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::documentMode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiArea_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) SetActivationOrder(order QMdiArea__WindowOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setActivationOrder")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetActivationOrder(ptr.Pointer(), C.int(order))
	}
}

func (ptr *QMdiArea) SetBackground(background gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetBackground(ptr.Pointer(), gui.PointerFromQBrush(background))
	}
}

func (ptr *QMdiArea) SetDocumentMode(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setDocumentMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setTabPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setTabShape")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QMdiArea) SetTabsClosable(closable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setTabsClosable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsClosable(ptr.Pointer(), C.int(qt.GoBoolToInt(closable)))
	}
}

func (ptr *QMdiArea) SetTabsMovable(movable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setTabsMovable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetTabsMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QMdiArea) SetViewMode(mode QMdiArea__ViewMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setViewMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QMdiArea) TabPosition() QTabWidget__TabPosition {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::tabPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMdiArea_TabPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) TabShape() QTabWidget__TabShape {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::tabShape")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMdiArea_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMdiArea) TabsClosable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::tabsClosable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsClosable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) TabsMovable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::tabsMovable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiArea_TabsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMdiArea) ViewMode() QMdiArea__ViewMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::viewMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QMdiArea__ViewMode(C.QMdiArea_ViewMode(ptr.Pointer()))
	}
	return 0
}

func NewQMdiArea(parent QWidget_ITF) *QMdiArea {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::QMdiArea")
		}
	}()

	return NewQMdiAreaFromPointer(C.QMdiArea_NewQMdiArea(PointerFromQWidget(parent)))
}

func (ptr *QMdiArea) ActivateNextSubWindow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::activateNextSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_ActivateNextSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ActivatePreviousSubWindow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::activatePreviousSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_ActivatePreviousSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::activeSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_ActiveSubWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) AddSubWindow(widget QWidget_ITF, windowFlags core.Qt__WindowType) *QMdiSubWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::addSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_AddSubWindow(ptr.Pointer(), PointerFromQWidget(widget), C.int(windowFlags)))
	}
	return nil
}

func (ptr *QMdiArea) CascadeSubWindows() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::cascadeSubWindows")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_CascadeSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) CloseActiveSubWindow() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::closeActiveSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseActiveSubWindow(ptr.Pointer())
	}
}

func (ptr *QMdiArea) CloseAllSubWindows() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::closeAllSubWindows")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_CloseAllSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::currentSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMdiSubWindowFromPointer(C.QMdiArea_CurrentSubWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMdiArea) RemoveSubWindow(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::removeSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_RemoveSubWindow(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMdiArea) SetActiveSubWindow(window QMdiSubWindow_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setActiveSubWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetActiveSubWindow(ptr.Pointer(), PointerFromQMdiSubWindow(window))
	}
}

func (ptr *QMdiArea) SetOption(option QMdiArea__AreaOption, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::setOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QMdiArea) ConnectSubWindowActivated(f func(window *QMdiSubWindow)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::subWindowActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_ConnectSubWindowActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "subWindowActivated", f)
	}
}

func (ptr *QMdiArea) DisconnectSubWindowActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::subWindowActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_DisconnectSubWindowActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "subWindowActivated")
	}
}

//export callbackQMdiAreaSubWindowActivated
func callbackQMdiAreaSubWindowActivated(ptrName *C.char, window unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::subWindowActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "subWindowActivated").(func(*QMdiSubWindow))(NewQMdiSubWindowFromPointer(window))
}

func (ptr *QMdiArea) TestOption(option QMdiArea__AreaOption) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::testOption")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMdiArea_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QMdiArea) TileSubWindows() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::tileSubWindows")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_TileSubWindows(ptr.Pointer())
	}
}

func (ptr *QMdiArea) DestroyQMdiArea() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMdiArea::~QMdiArea")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMdiArea_DestroyQMdiArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
