package widgets

//#include "qprogressdialog.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QProgressDialog struct {
	QDialog
}

type QProgressDialogITF interface {
	QDialogITF
	QProgressDialogPTR() *QProgressDialog
}

func PointerFromQProgressDialog(ptr QProgressDialogITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressDialogPTR().Pointer()
	}
	return nil
}

func QProgressDialogFromPointer(ptr unsafe.Pointer) *QProgressDialog {
	var n = new(QProgressDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QProgressDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProgressDialog) QProgressDialogPTR() *QProgressDialog {
	return ptr
}

func (ptr *QProgressDialog) AutoClose() bool {
	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoClose(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProgressDialog) AutoReset() bool {
	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoReset(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QProgressDialog) LabelText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressDialog_LabelText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QProgressDialog) Maximum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Maximum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressDialog) Minimum() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Minimum(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressDialog) MinimumDuration() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_MinimumDuration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressDialog) SetAutoClose(close bool) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoClose(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProgressDialog) SetAutoReset(reset bool) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoReset(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(reset)))
	}
}

func (ptr *QProgressDialog) SetLabelText(text string) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabelText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QProgressDialog) SetMaximum(maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMaximum(C.QtObjectPtr(ptr.Pointer()), C.int(maximum))
	}
}

func (ptr *QProgressDialog) SetMinimum(minimum int) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimum(C.QtObjectPtr(ptr.Pointer()), C.int(minimum))
	}
}

func (ptr *QProgressDialog) SetMinimumDuration(ms int) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimumDuration(C.QtObjectPtr(ptr.Pointer()), C.int(ms))
	}
}

func (ptr *QProgressDialog) SetValue(progress int) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetValue(C.QtObjectPtr(ptr.Pointer()), C.int(progress))
	}
}

func (ptr *QProgressDialog) Value() int {
	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QProgressDialog) WasCanceled() bool {
	if ptr.Pointer() != nil {
		return C.QProgressDialog_WasCanceled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQProgressDialog(parent QWidgetITF, f core.Qt__WindowType) *QProgressDialog {
	return QProgressDialogFromPointer(unsafe.Pointer(C.QProgressDialog_NewQProgressDialog(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func NewQProgressDialog2(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidgetITF, f core.Qt__WindowType) *QProgressDialog {
	return QProgressDialogFromPointer(unsafe.Pointer(C.QProgressDialog_NewQProgressDialog2(C.CString(labelText), C.CString(cancelButtonText), C.int(minimum), C.int(maximum), C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func (ptr *QProgressDialog) Cancel() {
	if ptr.Pointer() != nil {
		C.QProgressDialog_Cancel(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProgressDialog) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_ConnectCanceled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QProgressDialog) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QProgressDialog_DisconnectCanceled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQProgressDialogCanceled
func callbackQProgressDialogCanceled(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QProgressDialog) Open(receiver core.QObjectITF, member string) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_Open(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))
	}
}

func (ptr *QProgressDialog) Reset() {
	if ptr.Pointer() != nil {
		C.QProgressDialog_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QProgressDialog) SetBar(bar QProgressBarITF) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQProgressBar(bar)))
	}
}

func (ptr *QProgressDialog) SetCancelButton(cancelButton QPushButtonITF) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButton(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPushButton(cancelButton)))
	}
}

func (ptr *QProgressDialog) SetCancelButtonText(cancelButtonText string) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButtonText(C.QtObjectPtr(ptr.Pointer()), C.CString(cancelButtonText))
	}
}

func (ptr *QProgressDialog) SetLabel(label QLabelITF) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLabel(label)))
	}
}

func (ptr *QProgressDialog) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QProgressDialog_SetRange(C.QtObjectPtr(ptr.Pointer()), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressDialog) DestroyQProgressDialog() {
	if ptr.Pointer() != nil {
		C.QProgressDialog_DestroyQProgressDialog(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
