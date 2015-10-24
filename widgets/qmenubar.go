package widgets

//#include "qmenubar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMenuBar struct {
	QWidget
}

type QMenuBarITF interface {
	QWidgetITF
	QMenuBarPTR() *QMenuBar
}

func PointerFromQMenuBar(ptr QMenuBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBarPTR().Pointer()
	}
	return nil
}

func QMenuBarFromPointer(ptr unsafe.Pointer) *QMenuBar {
	var n = new(QMenuBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMenuBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMenuBar) QMenuBarPTR() *QMenuBar {
	return ptr
}

func (ptr *QMenuBar) IsDefaultUp() bool {
	if ptr.Pointer() != nil {
		return C.QMenuBar_IsDefaultUp(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenuBar) IsNativeMenuBar() bool {
	if ptr.Pointer() != nil {
		return C.QMenuBar_IsNativeMenuBar(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMenuBar) SetCornerWidget(widget QWidgetITF, corner core.Qt__Corner) {
	if ptr.Pointer() != nil {
		C.QMenuBar_SetCornerWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(corner))
	}
}

func (ptr *QMenuBar) SetDefaultUp(v bool) {
	if ptr.Pointer() != nil {
		C.QMenuBar_SetDefaultUp(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	if ptr.Pointer() != nil {
		C.QMenuBar_SetNativeMenuBar(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(nativeMenuBar)))
	}
}

func NewQMenuBar(parent QWidgetITF) *QMenuBar {
	return QMenuBarFromPointer(unsafe.Pointer(C.QMenuBar_NewQMenuBar(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QMenuBar) ActionAt(pt core.QPointITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_ActionAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pt)))))
	}
	return nil
}

func (ptr *QMenuBar) ActiveAction() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_ActiveAction(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenuBar) AddAction(text string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_AddAction(C.QtObjectPtr(ptr.Pointer()), C.CString(text))))
	}
	return nil
}

func (ptr *QMenuBar) AddAction2(text string, receiver core.QObjectITF, member string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_AddAction2(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu(menu QMenuITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_AddMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu3(icon gui.QIconITF, title string) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMenuBar_AddMenu3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(title))))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu2(title string) *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMenuBar_AddMenu2(C.QtObjectPtr(ptr.Pointer()), C.CString(title))))
	}
	return nil
}

func (ptr *QMenuBar) AddSeparator() *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_AddSeparator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMenuBar) Clear() {
	if ptr.Pointer() != nil {
		C.QMenuBar_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMenuBar) CornerWidget(corner core.Qt__Corner) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QMenuBar_CornerWidget(C.QtObjectPtr(ptr.Pointer()), C.int(corner))))
	}
	return nil
}

func (ptr *QMenuBar) HeightForWidth(v int) int {
	if ptr.Pointer() != nil {
		return int(C.QMenuBar_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(v)))
	}
	return 0
}

func (ptr *QMenuBar) ConnectHovered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenuBar) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuBarHovered
func callbackQMenuBarHovered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QMenuBar) InsertMenu(before QActionITF, menu QMenuITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_InsertMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(PointerFromQMenu(menu)))))
	}
	return nil
}

func (ptr *QMenuBar) InsertSeparator(before QActionITF) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QMenuBar_InsertSeparator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)))))
	}
	return nil
}

func (ptr *QMenuBar) SetActiveAction(act QActionITF) {
	if ptr.Pointer() != nil {
		C.QMenuBar_SetActiveAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(act)))
	}
}

func (ptr *QMenuBar) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QMenuBar_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenuBar) ConnectTriggered(f func(action QActionITF)) {
	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenuBar) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuBarTriggered
func callbackQMenuBarTriggered(ptrName *C.char, action unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(*QAction))(QActionFromPointer(action))
}

func (ptr *QMenuBar) DestroyQMenuBar() {
	if ptr.Pointer() != nil {
		C.QMenuBar_DestroyQMenuBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
