package widgets

//#include "qcolordialog.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QColorDialog struct {
	QDialog
}

type QColorDialog_ITF interface {
	QDialog_ITF
	QColorDialog_PTR() *QColorDialog
}

func PointerFromQColorDialog(ptr QColorDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColorDialog_PTR().Pointer()
	}
	return nil
}

func NewQColorDialogFromPointer(ptr unsafe.Pointer) *QColorDialog {
	var n = new(QColorDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QColorDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QColorDialog) QColorDialog_PTR() *QColorDialog {
	return ptr
}

//QColorDialog::ColorDialogOption
type QColorDialog__ColorDialogOption int64

const (
	QColorDialog__ShowAlphaChannel    = QColorDialog__ColorDialogOption(0x00000001)
	QColorDialog__NoButtons           = QColorDialog__ColorDialogOption(0x00000002)
	QColorDialog__DontUseNativeDialog = QColorDialog__ColorDialogOption(0x00000004)
)

func (ptr *QColorDialog) CurrentColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QColorDialog_CurrentColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColorDialog) Options() QColorDialog__ColorDialogOption {
	if ptr.Pointer() != nil {
		return QColorDialog__ColorDialogOption(C.QColorDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColorDialog) SetCurrentColor(color gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetCurrentColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QColorDialog) SetOptions(options QColorDialog__ColorDialogOption) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQColorDialog(parent QWidget_ITF) *QColorDialog {
	return NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog(PointerFromQWidget(parent)))
}

func NewQColorDialog2(initial gui.QColor_ITF, parent QWidget_ITF) *QColorDialog {
	return NewQColorDialogFromPointer(C.QColorDialog_NewQColorDialog2(gui.PointerFromQColor(initial), PointerFromQWidget(parent)))
}

func (ptr *QColorDialog) ConnectColorSelected(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QColorDialog_ConnectColorSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorSelected", f)
	}
}

func (ptr *QColorDialog) DisconnectColorSelected() {
	if ptr.Pointer() != nil {
		C.QColorDialog_DisconnectColorSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorSelected")
	}
}

//export callbackQColorDialogColorSelected
func callbackQColorDialogColorSelected(ptrName *C.char, color unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "colorSelected").(func(*gui.QColor))(gui.NewQColorFromPointer(color))
}

func (ptr *QColorDialog) ConnectCurrentColorChanged(f func(color *gui.QColor)) {
	if ptr.Pointer() != nil {
		C.QColorDialog_ConnectCurrentColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentColorChanged", f)
	}
}

func (ptr *QColorDialog) DisconnectCurrentColorChanged() {
	if ptr.Pointer() != nil {
		C.QColorDialog_DisconnectCurrentColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentColorChanged")
	}
}

//export callbackQColorDialogCurrentColorChanged
func callbackQColorDialogCurrentColorChanged(ptrName *C.char, color unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "currentColorChanged").(func(*gui.QColor))(gui.NewQColorFromPointer(color))
}

func QColorDialog_CustomColor(index int) *gui.QColor {
	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_CustomColor(C.int(index)))
}

func QColorDialog_CustomCount() int {
	return int(C.QColorDialog_QColorDialog_CustomCount())
}

func QColorDialog_GetColor(initial gui.QColor_ITF, parent QWidget_ITF, title string, options QColorDialog__ColorDialogOption) *gui.QColor {
	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_GetColor(gui.PointerFromQColor(initial), PointerFromQWidget(parent), C.CString(title), C.int(options)))
}

func (ptr *QColorDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		C.QColorDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QColorDialog) SelectedColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QColorDialog_SelectedColor(ptr.Pointer()))
	}
	return nil
}

func QColorDialog_SetCustomColor(index int, color gui.QColor_ITF) {
	C.QColorDialog_QColorDialog_SetCustomColor(C.int(index), gui.PointerFromQColor(color))
}

func (ptr *QColorDialog) SetOption(option QColorDialog__ColorDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func QColorDialog_SetStandardColor(index int, color gui.QColor_ITF) {
	C.QColorDialog_QColorDialog_SetStandardColor(C.int(index), gui.PointerFromQColor(color))
}

func (ptr *QColorDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func QColorDialog_StandardColor(index int) *gui.QColor {
	return gui.NewQColorFromPointer(C.QColorDialog_QColorDialog_StandardColor(C.int(index)))
}

func (ptr *QColorDialog) TestOption(option QColorDialog__ColorDialogOption) bool {
	if ptr.Pointer() != nil {
		return C.QColorDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QColorDialog) DestroyQColorDialog() {
	if ptr.Pointer() != nil {
		C.QColorDialog_DestroyQColorDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
