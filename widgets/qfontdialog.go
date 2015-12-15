package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QFontDialog_") {
		n.SetObjectName("QFontDialog_" + qt.Identifier())
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
	defer qt.Recovering("QFontDialog::options")

	if ptr.Pointer() != nil {
		return QFontDialog__FontDialogOption(C.QFontDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontDialog) SetOptions(options QFontDialog__FontDialogOption) {
	defer qt.Recovering("QFontDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQFontDialog(parent QWidget_ITF) *QFontDialog {
	defer qt.Recovering("QFontDialog::QFontDialog")

	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog(PointerFromQWidget(parent)))
}

func NewQFontDialog2(initial gui.QFont_ITF, parent QWidget_ITF) *QFontDialog {
	defer qt.Recovering("QFontDialog::QFontDialog")

	return NewQFontDialogFromPointer(C.QFontDialog_NewQFontDialog2(gui.PointerFromQFont(initial), PointerFromQWidget(parent)))
}

func (ptr *QFontDialog) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QFontDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFontDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFontDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFontDialogChangeEvent
func callbackQFontDialogChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFontDialog::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFontDialog) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QFontDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QFontDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QFontDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQFontDialogDone
func callbackQFontDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QFontDialog::done")

	var signal = qt.GetSignal(C.GoString(ptrName), "done")
	if signal != nil {
		defer signal.(func(int))(int(result))
		return true
	}
	return false

}

func (ptr *QFontDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QFontDialog::open")

	if ptr.Pointer() != nil {
		C.QFontDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFontDialog) SetCurrentFont(font gui.QFont_ITF) {
	defer qt.Recovering("QFontDialog::setCurrentFont")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QFontDialog) SetOption(option QFontDialog__FontDialogOption, on bool) {
	defer qt.Recovering("QFontDialog::setOption")

	if ptr.Pointer() != nil {
		C.QFontDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFontDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFontDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFontDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFontDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFontDialogSetVisible
func callbackQFontDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFontDialog::setVisible")

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFontDialog) TestOption(option QFontDialog__FontDialogOption) bool {
	defer qt.Recovering("QFontDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QFontDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}
