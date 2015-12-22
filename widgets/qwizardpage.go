package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QWizardPage_") {
		n.SetObjectName("QWizardPage_" + qt.Identifier())
	}
	return n
}

func (ptr *QWizardPage) QWizardPage_PTR() *QWizardPage {
	return ptr
}

func (ptr *QWizardPage) SetSubTitle(subTitle string) {
	defer qt.Recovering("QWizardPage::setSubTitle")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetSubTitle(ptr.Pointer(), C.CString(subTitle))
	}
}

func (ptr *QWizardPage) SetTitle(title string) {
	defer qt.Recovering("QWizardPage::setTitle")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QWizardPage) SubTitle() string {
	defer qt.Recovering("QWizardPage::subTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_SubTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWizardPage) Title() string {
	defer qt.Recovering("QWizardPage::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_Title(ptr.Pointer()))
	}
	return ""
}

func NewQWizardPage(parent QWidget_ITF) *QWizardPage {
	defer qt.Recovering("QWizardPage::QWizardPage")

	return NewQWizardPageFromPointer(C.QWizardPage_NewQWizardPage(PointerFromQWidget(parent)))
}

func (ptr *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	defer qt.Recovering("QWizardPage::buttonText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizardPage) ConnectCleanupPage(f func()) {
	defer qt.Recovering("connect QWizardPage::cleanupPage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "cleanupPage", f)
	}
}

func (ptr *QWizardPage) DisconnectCleanupPage() {
	defer qt.Recovering("disconnect QWizardPage::cleanupPage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "cleanupPage")
	}
}

//export callbackQWizardPageCleanupPage
func callbackQWizardPageCleanupPage(ptrName *C.char) bool {
	defer qt.Recovering("callback QWizardPage::cleanupPage")

	if signal := qt.GetSignal(C.GoString(ptrName), "cleanupPage"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QWizardPage) ConnectCompleteChanged(f func()) {
	defer qt.Recovering("connect QWizardPage::completeChanged")

	if ptr.Pointer() != nil {
		C.QWizardPage_ConnectCompleteChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "completeChanged", f)
	}
}

func (ptr *QWizardPage) DisconnectCompleteChanged() {
	defer qt.Recovering("disconnect QWizardPage::completeChanged")

	if ptr.Pointer() != nil {
		C.QWizardPage_DisconnectCompleteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "completeChanged")
	}
}

//export callbackQWizardPageCompleteChanged
func callbackQWizardPageCompleteChanged(ptrName *C.char) {
	defer qt.Recovering("callback QWizardPage::completeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "completeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWizardPage) ConnectInitializePage(f func()) {
	defer qt.Recovering("connect QWizardPage::initializePage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initializePage", f)
	}
}

func (ptr *QWizardPage) DisconnectInitializePage() {
	defer qt.Recovering("disconnect QWizardPage::initializePage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initializePage")
	}
}

//export callbackQWizardPageInitializePage
func callbackQWizardPageInitializePage(ptrName *C.char) bool {
	defer qt.Recovering("callback QWizardPage::initializePage")

	if signal := qt.GetSignal(C.GoString(ptrName), "initializePage"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QWizardPage) IsCommitPage() bool {
	defer qt.Recovering("QWizardPage::isCommitPage")

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsCommitPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsComplete() bool {
	defer qt.Recovering("QWizardPage::isComplete")

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsFinalPage() bool {
	defer qt.Recovering("QWizardPage::isFinalPage")

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsFinalPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) NextId() int {
	defer qt.Recovering("QWizardPage::nextId")

	if ptr.Pointer() != nil {
		return int(C.QWizardPage_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	defer qt.Recovering("QWizardPage::setButtonText")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizardPage) SetCommitPage(commitPage bool) {
	defer qt.Recovering("QWizardPage::setCommitPage")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetCommitPage(ptr.Pointer(), C.int(qt.GoBoolToInt(commitPage)))
	}
}

func (ptr *QWizardPage) SetFinalPage(finalPage bool) {
	defer qt.Recovering("QWizardPage::setFinalPage")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetFinalPage(ptr.Pointer(), C.int(qt.GoBoolToInt(finalPage)))
	}
}

func (ptr *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QWizardPage::setPixmap")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizardPage) ValidatePage() bool {
	defer qt.Recovering("QWizardPage::validatePage")

	if ptr.Pointer() != nil {
		return C.QWizardPage_ValidatePage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) DestroyQWizardPage() {
	defer qt.Recovering("QWizardPage::~QWizardPage")

	if ptr.Pointer() != nil {
		C.QWizardPage_DestroyQWizardPage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
