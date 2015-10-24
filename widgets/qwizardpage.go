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

type QWizardPageITF interface {
	QWidgetITF
	QWizardPagePTR() *QWizardPage
}

func PointerFromQWizardPage(ptr QWizardPageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWizardPagePTR().Pointer()
	}
	return nil
}

func QWizardPageFromPointer(ptr unsafe.Pointer) *QWizardPage {
	var n = new(QWizardPage)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWizardPage_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizardPage) QWizardPagePTR() *QWizardPage {
	return ptr
}

func (ptr *QWizardPage) SetSubTitle(subTitle string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetSubTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(subTitle))
	}
}

func (ptr *QWizardPage) SetTitle(title string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(title))
	}
}

func (ptr *QWizardPage) SubTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_SubTitle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWizardPage) Title() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_Title(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQWizardPage(parent QWidgetITF) *QWizardPage {
	return QWizardPageFromPointer(unsafe.Pointer(C.QWizardPage_NewQWizardPage(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_ButtonText(C.QtObjectPtr(ptr.Pointer()), C.int(which)))
	}
	return ""
}

func (ptr *QWizardPage) CleanupPage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_CleanupPage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWizardPage) ConnectCompleteChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QWizardPage_ConnectCompleteChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "completeChanged", f)
	}
}

func (ptr *QWizardPage) DisconnectCompleteChanged() {
	if ptr.Pointer() != nil {
		C.QWizardPage_DisconnectCompleteChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "completeChanged")
	}
}

//export callbackQWizardPageCompleteChanged
func callbackQWizardPageCompleteChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "completeChanged").(func())()
}

func (ptr *QWizardPage) InitializePage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_InitializePage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWizardPage) IsCommitPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsCommitPage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWizardPage) IsComplete() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsComplete(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWizardPage) IsFinalPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_IsFinalPage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWizardPage) NextId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizardPage_NextId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetButtonText(C.QtObjectPtr(ptr.Pointer()), C.int(which), C.CString(text))
	}
}

func (ptr *QWizardPage) SetCommitPage(commitPage bool) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetCommitPage(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(commitPage)))
	}
}

func (ptr *QWizardPage) SetFinalPage(finalPage bool) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetFinalPage(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(finalPage)))
	}
}

func (ptr *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QWizardPage_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.int(which), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QWizardPage) ValidatePage() bool {
	if ptr.Pointer() != nil {
		return C.QWizardPage_ValidatePage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWizardPage) DestroyQWizardPage() {
	if ptr.Pointer() != nil {
		C.QWizardPage_DestroyQWizardPage(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
