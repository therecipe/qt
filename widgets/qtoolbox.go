package widgets

//#include "qtoolbox.h"
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

type QToolBoxITF interface {
	QFrameITF
	QToolBoxPTR() *QToolBox
}

func PointerFromQToolBox(ptr QToolBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolBoxPTR().Pointer()
	}
	return nil
}

func QToolBoxFromPointer(ptr unsafe.Pointer) *QToolBox {
	var n = new(QToolBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QToolBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QToolBox) QToolBoxPTR() *QToolBox {
	return ptr
}

func (ptr *QToolBox) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolBox) CurrentIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_CurrentIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QToolBox) SetCurrentIndex(index int) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func NewQToolBox(parent QWidgetITF, f core.Qt__WindowType) *QToolBox {
	return QToolBoxFromPointer(unsafe.Pointer(C.QToolBox_NewQToolBox(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func (ptr *QToolBox) AddItem2(w QWidgetITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w)), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) AddItem(widget QWidgetITF, iconSet gui.QIconITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.QtObjectPtr(gui.PointerFromQIcon(iconSet)), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) ConnectCurrentChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QToolBox_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QToolBox) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QToolBox_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQToolBoxCurrentChanged
func callbackQToolBoxCurrentChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QToolBox) CurrentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QToolBox_CurrentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QToolBox) IndexOf(widget QWidgetITF) int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_IndexOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))))
	}
	return 0
}

func (ptr *QToolBox) InsertItem(index int, widget QWidgetITF, icon gui.QIconITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)), C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) InsertItem2(index int, widget QWidgetITF, text string) int {
	if ptr.Pointer() != nil {
		return int(C.QToolBox_InsertItem2(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)), C.CString(text)))
	}
	return 0
}

func (ptr *QToolBox) IsItemEnabled(index int) bool {
	if ptr.Pointer() != nil {
		return C.QToolBox_IsItemEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index)) != 0
	}
	return false
}

func (ptr *QToolBox) ItemText(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemText(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) ItemToolTip(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QToolBox_ItemToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QToolBox) RemoveItem(index int) {
	if ptr.Pointer() != nil {
		C.QToolBox_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QToolBox) SetCurrentWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetCurrentWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QToolBox) SetItemEnabled(index int, enabled bool) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetItemEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QToolBox) SetItemIcon(index int, icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetItemIcon(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QToolBox) SetItemText(index int, text string) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetItemText(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(text))
	}
}

func (ptr *QToolBox) SetItemToolTip(index int, toolTip string) {
	if ptr.Pointer() != nil {
		C.QToolBox_SetItemToolTip(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.CString(toolTip))
	}
}

func (ptr *QToolBox) Widget(index int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QToolBox_Widget(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QToolBox) DestroyQToolBox() {
	if ptr.Pointer() != nil {
		C.QToolBox_DestroyQToolBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
