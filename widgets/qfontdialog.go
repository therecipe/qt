package widgets

//#include "qfontdialog.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QFontDialog struct {
	QDialog
}

type QFontDialogITF interface {
	QDialogITF
	QFontDialogPTR() *QFontDialog
}

func PointerFromQFontDialog(ptr QFontDialogITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontDialogPTR().Pointer()
	}
	return nil
}

func QFontDialogFromPointer(ptr unsafe.Pointer) *QFontDialog {
	var n = new(QFontDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFontDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFontDialog) QFontDialogPTR() *QFontDialog {
	return ptr
}

//QFontDialog::FontDialogOption
type QFontDialog__FontDialogOption int

var (
	QFontDialog__NoButtons           = QFontDialog__FontDialogOption(0x00000001)
	QFontDialog__DontUseNativeDialog = QFontDialog__FontDialogOption(0x00000002)
	QFontDialog__ScalableFonts       = QFontDialog__FontDialogOption(0x00000004)
	QFontDialog__NonScalableFonts    = QFontDialog__FontDialogOption(0x00000008)
	QFontDialog__MonospacedFonts     = QFontDialog__FontDialogOption(0x00000010)
	QFontDialog__ProportionalFonts   = QFontDialog__FontDialogOption(0x00000020)
)

func (ptr *QFontDialog) Options() QFontDialog__FontDialogOption {
	if ptr.Pointer() != nil {
		return QFontDialog__FontDialogOption(C.QFontDialog_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFontDialog) SetOptions(options QFontDialog__FontDialogOption) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func NewQFontDialog(parent QWidgetITF) *QFontDialog {
	return QFontDialogFromPointer(unsafe.Pointer(C.QFontDialog_NewQFontDialog(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQFontDialog2(initial gui.QFontITF, parent QWidgetITF) *QFontDialog {
	return QFontDialogFromPointer(unsafe.Pointer(C.QFontDialog_NewQFontDialog2(C.QtObjectPtr(gui.PointerFromQFont(initial)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QFontDialog) Open(receiver core.QObjectITF, member string) {
	if ptr.Pointer() != nil {
		C.QFontDialog_Open(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))
	}
}

func (ptr *QFontDialog) SetCurrentFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetCurrentFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QFontDialog) SetOption(option QFontDialog__FontDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFontDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFontDialog) TestOption(option QFontDialog__FontDialogOption) bool {
	if ptr.Pointer() != nil {
		return C.QFontDialog_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}
