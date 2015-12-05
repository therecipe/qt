package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QProgressDialog struct {
	QDialog
}

type QProgressDialog_ITF interface {
	QDialog_ITF
	QProgressDialog_PTR() *QProgressDialog
}

func PointerFromQProgressDialog(ptr QProgressDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QProgressDialog_PTR().Pointer()
	}
	return nil
}

func NewQProgressDialogFromPointer(ptr unsafe.Pointer) *QProgressDialog {
	var n = new(QProgressDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QProgressDialog_") {
		n.SetObjectName("QProgressDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QProgressDialog) QProgressDialog_PTR() *QProgressDialog {
	return ptr
}

func (ptr *QProgressDialog) AutoClose() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::autoClose")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoClose(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) AutoReset() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::autoReset")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProgressDialog_AutoReset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QProgressDialog) LabelText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::labelText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QProgressDialog_LabelText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QProgressDialog) Maximum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::maximum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) Minimum() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::minimum")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) MinimumDuration() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::minimumDuration")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_MinimumDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) SetAutoClose(close bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setAutoClose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoClose(ptr.Pointer(), C.int(qt.GoBoolToInt(close)))
	}
}

func (ptr *QProgressDialog) SetAutoReset(reset bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setAutoReset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetAutoReset(ptr.Pointer(), C.int(qt.GoBoolToInt(reset)))
	}
}

func (ptr *QProgressDialog) SetLabelText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setLabelText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabelText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QProgressDialog) SetMaximum(maximum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setMaximum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMaximum(ptr.Pointer(), C.int(maximum))
	}
}

func (ptr *QProgressDialog) SetMinimum(minimum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setMinimum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimum(ptr.Pointer(), C.int(minimum))
	}
}

func (ptr *QProgressDialog) SetMinimumDuration(ms int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setMinimumDuration")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetMinimumDuration(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QProgressDialog) SetValue(progress int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setValue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetValue(ptr.Pointer(), C.int(progress))
	}
}

func (ptr *QProgressDialog) Value() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::value")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QProgressDialog_Value(ptr.Pointer()))
	}
	return 0
}

func (ptr *QProgressDialog) WasCanceled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::wasCanceled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QProgressDialog_WasCanceled(ptr.Pointer()) != 0
	}
	return false
}

func NewQProgressDialog(parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::QProgressDialog")
		}
	}()

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog(PointerFromQWidget(parent), C.int(f)))
}

func NewQProgressDialog2(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF, f core.Qt__WindowType) *QProgressDialog {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::QProgressDialog")
		}
	}()

	return NewQProgressDialogFromPointer(C.QProgressDialog_NewQProgressDialog2(C.CString(labelText), C.CString(cancelButtonText), C.int(minimum), C.int(maximum), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QProgressDialog) Cancel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::cancel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_Cancel(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) ConnectCanceled(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::canceled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QProgressDialog) DisconnectCanceled() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::canceled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQProgressDialogCanceled
func callbackQProgressDialogCanceled(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::canceled")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QProgressDialog) Open(receiver core.QObject_ITF, member string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::open")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QProgressDialog) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_Reset(ptr.Pointer())
	}
}

func (ptr *QProgressDialog) SetBar(bar QProgressBar_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetBar(ptr.Pointer(), PointerFromQProgressBar(bar))
	}
}

func (ptr *QProgressDialog) SetCancelButton(cancelButton QPushButton_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setCancelButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButton(ptr.Pointer(), PointerFromQPushButton(cancelButton))
	}
}

func (ptr *QProgressDialog) SetCancelButtonText(cancelButtonText string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setCancelButtonText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetCancelButtonText(ptr.Pointer(), C.CString(cancelButtonText))
	}
}

func (ptr *QProgressDialog) SetLabel(label QLabel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setLabel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetLabel(ptr.Pointer(), PointerFromQLabel(label))
	}
}

func (ptr *QProgressDialog) SetRange(minimum int, maximum int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::setRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_SetRange(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QProgressDialog) DestroyQProgressDialog() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QProgressDialog::~QProgressDialog")
		}
	}()

	if ptr.Pointer() != nil {
		C.QProgressDialog_DestroyQProgressDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
