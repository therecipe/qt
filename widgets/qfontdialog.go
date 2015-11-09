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

type QFontDialog_ITF interface {
	QDialog_ITF
	QFontDialog_PTR() *QFontDialog
}

func PointerFromQFontDialog(ptr QFontDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontDialog_PTR().Pointer()
	}
	return nil
}

func NewQFontDialogFromPointer(ptr unsafe.Pointer) *QFontDialog {
	var n = new(QFontDialog)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QFontDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFontDialog) QFontDialog_PTR() *QFontDialog {
	return ptr
}

//QFontDialog::FontDialogOption
type QFontDialog__FontDialogOption int64

const (
	QFontDialog__NoButtons           = QFontDialog__FontDialogOption(0x00000001)
	QFontDialog__DontUseNativeDialog = QFontDialog__FontDialogOption(0x00000002)
	QFontDialog__ScalableFonts       = QFontDialog__FontDialogOption(0x00000004)
	QFontDialog__NonScalableFonts    = QFontDialog__FontDialogOption(0x00000008)
	QFontDialog__MonospacedFonts     = QFontDialog__FontDialogOption(0x00000010)
	QFontDialog__ProportionalFonts   = QFontDialog__FontDialogOption(0x00000020)
)

func (ptr *QFontDialog) Options() QFontDialog__FontDialogOption {
	if ptr.Pointer() != nil {
		return QFontDialog__FontDialogOption(C.QFontDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontDialog) SetOptions(options QFontDialog__FontDialogOption) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQFontDialog(parent QWidget_ITF) *QFontDialog {
	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog(PointerFromQWidget(parent)))
}

func NewQFontDialog2(initial gui.QFont_ITF, parent QWidget_ITF) *QFontDialog {
	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog2(gui.PointerFromQFont(initial), PointerFromQWidget(parent)))
}

func (ptr *QFontDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		C.QFontDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFontDialog) SetCurrentFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QFontDialog) SetOption(option QFontDialog__FontDialogOption, on bool) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFontDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QFontDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFontDialog) TestOption(option QFontDialog__FontDialogOption) bool {
	if ptr.Pointer() != nil {
		return C.QFontDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}
