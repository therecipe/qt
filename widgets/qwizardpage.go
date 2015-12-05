package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"log"
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
		n.SetObjectName("QWizardPage_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizardPage) QWizardPage_PTR() *QWizardPage {
	return ptr
}

func (ptr *QWizardPage) SetSubTitle(subTitle string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setSubTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetSubTitle(ptr.Pointer(), C.CString(subTitle))
	}
}

func (ptr *QWizardPage) SetTitle(title string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setTitle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QWizardPage) SubTitle() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::subTitle")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_SubTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWizardPage) Title() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::title")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_Title(ptr.Pointer()))
	}
	return ""
}

func NewQWizardPage(parent QWidget_ITF) *QWizardPage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::QWizardPage")
		}
	}()

	return NewQWizardPageFromPointer(C.QWizardPage_NewQWizardPage(PointerFromQWidget(parent)))
}

func (ptr *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::buttonText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizardPage_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizardPage) CleanupPage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::cleanupPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_CleanupPage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) ConnectCompleteChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::completeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_ConnectCompleteChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "completeChanged", f)
	}
}

func (ptr *QWizardPage) DisconnectCompleteChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::completeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_DisconnectCompleteChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "completeChanged")
	}
}

//export callbackQWizardPageCompleteChanged
func callbackQWizardPageCompleteChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::completeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "completeChanged").(func())()
}

func (ptr *QWizardPage) InitializePage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::initializePage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_InitializePage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) IsCommitPage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::isCommitPage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsCommitPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsComplete() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::isComplete")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) IsFinalPage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::isFinalPage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizardPage_IsFinalPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) NextId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::nextId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QWizardPage_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setButtonText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizardPage) SetCommitPage(commitPage bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setCommitPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetCommitPage(ptr.Pointer(), C.int(qt.GoBoolToInt(commitPage)))
	}
}

func (ptr *QWizardPage) SetFinalPage(finalPage bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setFinalPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetFinalPage(ptr.Pointer(), C.int(qt.GoBoolToInt(finalPage)))
	}
}

func (ptr *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::setPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizardPage) ValidatePage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::validatePage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizardPage_ValidatePage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizardPage) DestroyQWizardPage() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizardPage::~QWizardPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizardPage_DestroyQWizardPage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
