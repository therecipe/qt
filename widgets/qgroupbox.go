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

type QGroupBoxITF interface {
	QWidgetITF
	QGroupBoxPTR() *QGroupBox
}

func PointerFromQGroupBox(ptr QGroupBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBoxPTR().Pointer()
	}
	return nil
}

func QGroupBoxFromPointer(ptr unsafe.Pointer) *QGroupBox {
	var n = new(QGroupBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGroupBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGroupBox) QGroupBoxPTR() *QGroupBox {
	return ptr
}

func (ptr *QGroupBox) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGroupBox_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGroupBox) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsCheckable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGroupBox) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsChecked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGroupBox) IsFlat() bool {
	if ptr.Pointer() != nil {
		return C.QGroupBox_IsFlat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGroupBox) SetAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QGroupBox) SetCheckable(checkable bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetCheckable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QGroupBox) SetChecked(checked bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetChecked(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QGroupBox) SetFlat(flat bool) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetFlat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(flat)))
	}
}

func (ptr *QGroupBox) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QGroupBox_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QGroupBox) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGroupBox_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQGroupBox(parent QWidgetITF) *QGroupBox {
	return QGroupBoxFromPointer(unsafe.Pointer(C.QGroupBox_NewQGroupBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQGroupBox2(title string, parent QWidgetITF) *QGroupBox {
	return QGroupBoxFromPointer(unsafe.Pointer(C.QGroupBox_NewQGroupBox2(C.CString(title), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QGroupBox) ConnectClicked(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QGroupBox) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQGroupBoxClicked
func callbackQGroupBoxClicked(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(bool))(int(checked) != 0)
}

func (ptr *QGroupBox) ConnectToggled(f func(on bool)) {
	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QGroupBox) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQGroupBoxToggled
func callbackQGroupBoxToggled(ptrName *C.char, on C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(on) != 0)
}

func (ptr *QGroupBox) DestroyQGroupBox() {
	if ptr.Pointer() != nil {
		C.QGroupBox_DestroyQGroupBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
