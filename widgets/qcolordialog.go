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

type QColorDialogITF interface {
	QDialogITF
	QColorDialogPTR() *QColorDialog
}

func PointerFromQColorDialog(ptr QColorDialogITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColorDialogPTR().Pointer()
	}
	return nil
}

func QColorDialogFromPointer(ptr unsafe.Pointer) *QColorDialog {
	var n = new(QColorDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QColorDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QColorDialog) QColorDialogPTR() *QColorDialog {
	return ptr
}

//QColorDialog::ColorDialogOption
type QColorDialog__ColorDialogOption int

var (
	QColorDialog__ShowAlphaChannel    = QColorDialog__ColorDialogOption(0x00000001)
	QColorDialog__NoButtons           = QColorDialog__ColorDialogOption(0x00000002)
	QColorDialog__DontUseNativeDialog = QColorDialog__ColorDialogOption(0x00000004)
)

func (ptr *QColorDialog) Options() QColorDialog__ColorDialogOption {
	if ptr.Pointer() != nil {
		return QColorDialog__ColorDialogOption(C.QColorDialog_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColorDialog) SetCurrentColor(color gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetCurrentColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(color)))
	}
}

func (ptr *QColorDialog) SetOptions(options QColorDialog__ColorDialogOption) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func NewQColorDialog(parent QWidgetITF) *QColorDialog {
	return QColorDialogFromPointer(unsafe.Pointer(C.QColorDialog_NewQColorDialog(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQColorDialog2(initial gui.QColorITF, parent QWidgetITF) *QColorDialog {
	return QColorDialogFromPointer(unsafe.Pointer(C.QColorDialog_NewQColorDialog2(C.QtObjectPtr(gui.PointerFromQColor(initial)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func QColorDialog_CustomCount() int {
	return int(C.QColorDialog_QColorDialog_CustomCount())
}

func (ptr *QColorDialog) Open(receiver core.QObjectITF, member string) {
	if ptr.Pointer() != nil {
		C.QColorDialog_Open(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))
	}
}

func QColorDialog_SetCustomColor(index int, color gui.QColorITF) {
	C.QColorDialog_QColorDialog_SetCustomColor(C.int(index), C.QtObjectPtr(gui.PointerFromQColor(color)))
}

func (ptr *QColorDialog) SetOption(option QColorDialog__ColorDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func QColorDialog_SetStandardColor(index int, color gui.QColorITF) {
	C.QColorDialog_QColorDialog_SetStandardColor(C.int(index), C.QtObjectPtr(gui.PointerFromQColor(color)))
}

func (ptr *QColorDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QColorDialog_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QColorDialog) TestOption(option QColorDialog__ColorDialogOption) bool {
	if ptr.Pointer() != nil {
		return C.QColorDialog_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QColorDialog) DestroyQColorDialog() {
	if ptr.Pointer() != nil {
		C.QColorDialog_DestroyQColorDialog(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
