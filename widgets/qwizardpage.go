package widgets

//#include "qwizardpage.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QWizardPage struct {
	QWidget
}

type QWizardPage_ITF interface {
	QWidget_ITF
	QWizardPage_PTR() *QWizardPage
}

func PointerFromQWizardPage(ptr QWizardPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWizardPage_PTR().Pointer()
	}
	return nil
}

func NewQWizardPageFromPointer(ptr unsafe.Pointer) *QWizardPage {
	var n = new(QWizardPage)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWizardPage_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizardPage) QWizardPage_PTR() *QWizardPage {
	return ptr
}

func (ptr *QWizardPage) SetSubTitle(subTitle string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetSubTitle(ptr.Pointer(), C.CString(subTitle))
	}
}

func (ptr *QWizardPage) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QWizardPage) SubTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_SubTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWizardPage) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_Title(ptr.Pointer()))
	}
	return ""
}

func NewQWizardPage(parent QWidget_ITF) *QWizardPage {
	return NewQWizardPageFromPointer(C.QWizardPage_NewQWizardPage(PointerFromQWidget(parent)))
}

func (ptr *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizardPage) CleanupPage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_CleanupPage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) ConnectCompleteChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWizardPage_ConnectCompleteChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "completeChanged", f)
	}
}

func (ptr *QWizardPage) DisconnectCompleteChanged() {
	if ptr.Pointer() != nil {
		C.QWizardPage_DisconnectCompleteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "completeChanged")
	}
}

//export callbackQWizardPageCompleteChanged
func callbackQWizardPageCompleteChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "completeChanged").(func())()
}

func (ptr *QWizardPage) InitializePage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_InitializePage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) IsCommitPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsCommitPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsComplete() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsFinalPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsFinalPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) NextId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizardPage_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizardPage) SetCommitPage(commitPage bool) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetCommitPage(ptr.Pointer(), C.int(qt.GoBoolToInt(commitPage)))
	}
}

func (ptr *QWizardPage) SetFinalPage(finalPage bool) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetFinalPage(ptr.Pointer(), C.int(qt.GoBoolToInt(finalPage)))
	}
}

func (ptr *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizardPage) ValidatePage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_ValidatePage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) DestroyQWizardPage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_DestroyQWizardPage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
