package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDialog struct {
	QWidget
}

type QDialog_ITF interface {
	QWidget_ITF
	QDialog_PTR() *QDialog
}

func PointerFromQDialog(ptr QDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDialog_PTR().Pointer()
	}
	return nil
}

func NewQDialogFromPointer(ptr unsafe.Pointer) *QDialog {
	var n = new(QDialog)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDialog_") {
		n.SetObjectName("QDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDialog) QDialog_PTR() *QDialog {
	return ptr
}

//QDialog::DialogCode
type QDialog__DialogCode int64

const (
	QDialog__Rejected = QDialog__DialogCode(0)
	QDialog__Accepted = QDialog__DialogCode(1)
)

func (ptr *QDialog) IsSizeGripEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::isSizeGripEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDialog_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDialog) SetModal(modal bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::setModal")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_SetModal(ptr.Pointer(), C.int(qt.GoBoolToInt(modal)))
	}
}

func (ptr *QDialog) SetResult(i int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::setResult")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_SetResult(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QDialog) SetSizeGripEnabled(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::setSizeGripEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQDialog(parent QWidget_ITF, f core.Qt__WindowType) *QDialog {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::QDialog")
		}
	}()

	return NewQDialogFromPointer(C.QDialog_NewQDialog(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QDialog) Accept() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::accept")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_Accept(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectAccepted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::accepted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_ConnectAccepted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "accepted", f)
	}
}

func (ptr *QDialog) DisconnectAccepted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::accepted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectAccepted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "accepted")
	}
}

//export callbackQDialogAccepted
func callbackQDialogAccepted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::accepted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "accepted").(func())()
}

func (ptr *QDialog) Done(r int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::done")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_Done(ptr.Pointer(), C.int(r))
	}
}

func (ptr *QDialog) Exec() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::exec")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDialog_Exec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) ConnectFinished(f func(result int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDialog) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDialogFinished
func callbackQDialogFinished(ptrName *C.char, result C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(int))(int(result))
}

func (ptr *QDialog) Open() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::open")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_Open(ptr.Pointer())
	}
}

func (ptr *QDialog) Reject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::reject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_Reject(ptr.Pointer())
	}
}

func (ptr *QDialog) ConnectRejected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::rejected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_ConnectRejected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rejected", f)
	}
}

func (ptr *QDialog) DisconnectRejected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::rejected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_DisconnectRejected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rejected")
	}
}

//export callbackQDialogRejected
func callbackQDialogRejected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::rejected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "rejected").(func())()
}

func (ptr *QDialog) Result() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::result")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDialog_Result(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDialog) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDialog) DestroyQDialog() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDialog::~QDialog")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDialog_DestroyQDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
