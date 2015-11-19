package widgets

//#include "qgroupbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGroupBox struct {
	QWidget
}

type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func PointerFromQGroupBox(ptr QGroupBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBox_PTR().Pointer()
	}
	return nil
}

func NewQGroupBoxFromPointer(ptr unsafe.Pointer) *QGroupBox {
	var n = new(QGroupBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGroupBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox {
	return ptr
}

func (ptr *QGroupBox) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGroupBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGroupBox) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsFlat() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) SetAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QGroupBox) SetCheckable(checkable bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QGroupBox) SetChecked(checked bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QGroupBox) SetFlat(flat bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(flat)))
	}
}

func (ptr *QGroupBox) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QGroupBox) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGroupBox_Title(ptr.Pointer()))
	}
	return ""
}

func NewQGroupBox(parent QWidget_ITF) *QGroupBox {
	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox(PointerFromQWidget(parent)))
}

func NewQGroupBox2(title string, parent QWidget_ITF) *QGroupBox {
	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QGroupBox) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QGroupBox) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQGroupBoxClicked
func callbackQGroupBoxClicked(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(bool))(int(checked) != 0)
}

func (ptr *QGroupBox) ConnectToggled(f func(on bool)) {
	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QGroupBox) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQGroupBoxToggled
func callbackQGroupBoxToggled(ptrName *C.char, on C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(on) != 0)
}

func (ptr *QGroupBox) DestroyQGroupBox() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DestroyQGroupBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
