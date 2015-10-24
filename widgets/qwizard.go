package widgets

//#include "qwizard.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QWizard struct {
	QDialog
}

type QWizardITF interface {
	QDialogITF
	QWizardPTR() *QWizard
}

func PointerFromQWizard(ptr QWizardITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWizardPTR().Pointer()
	}
	return nil
}

func QWizardFromPointer(ptr unsafe.Pointer) *QWizard {
	var n = new(QWizard)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWizard_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizard) QWizardPTR() *QWizard {
	return ptr
}

//QWizard::WizardButton
type QWizard__WizardButton int

var (
	QWizard__BackButton       = QWizard__WizardButton(0)
	QWizard__NextButton       = QWizard__WizardButton(1)
	QWizard__CommitButton     = QWizard__WizardButton(2)
	QWizard__FinishButton     = QWizard__WizardButton(3)
	QWizard__CancelButton     = QWizard__WizardButton(4)
	QWizard__HelpButton       = QWizard__WizardButton(5)
	QWizard__CustomButton1    = QWizard__WizardButton(6)
	QWizard__CustomButton2    = QWizard__WizardButton(7)
	QWizard__CustomButton3    = QWizard__WizardButton(8)
	QWizard__Stretch          = QWizard__WizardButton(9)
	QWizard__NoButton         = QWizard__WizardButton(-1)
	QWizard__NStandardButtons = QWizard__WizardButton(6)
	QWizard__NButtons         = QWizard__WizardButton(9)
)

//QWizard::WizardOption
type QWizard__WizardOption int

var (
	QWizard__IndependentPages             = QWizard__WizardOption(0x00000001)
	QWizard__IgnoreSubTitles              = QWizard__WizardOption(0x00000002)
	QWizard__ExtendedWatermarkPixmap      = QWizard__WizardOption(0x00000004)
	QWizard__NoDefaultButton              = QWizard__WizardOption(0x00000008)
	QWizard__NoBackButtonOnStartPage      = QWizard__WizardOption(0x00000010)
	QWizard__NoBackButtonOnLastPage       = QWizard__WizardOption(0x00000020)
	QWizard__DisabledBackButtonOnLastPage = QWizard__WizardOption(0x00000040)
	QWizard__HaveNextButtonOnLastPage     = QWizard__WizardOption(0x00000080)
	QWizard__HaveFinishButtonOnEarlyPages = QWizard__WizardOption(0x00000100)
	QWizard__NoCancelButton               = QWizard__WizardOption(0x00000200)
	QWizard__CancelButtonOnLeft           = QWizard__WizardOption(0x00000400)
	QWizard__HaveHelpButton               = QWizard__WizardOption(0x00000800)
	QWizard__HelpButtonOnRight            = QWizard__WizardOption(0x00001000)
	QWizard__HaveCustomButton1            = QWizard__WizardOption(0x00002000)
	QWizard__HaveCustomButton2            = QWizard__WizardOption(0x00004000)
	QWizard__HaveCustomButton3            = QWizard__WizardOption(0x00008000)
	QWizard__NoCancelButtonOnLastPage     = QWizard__WizardOption(0x00010000)
)

//QWizard::WizardPixmap
type QWizard__WizardPixmap int

var (
	QWizard__WatermarkPixmap  = QWizard__WizardPixmap(0)
	QWizard__LogoPixmap       = QWizard__WizardPixmap(1)
	QWizard__BannerPixmap     = QWizard__WizardPixmap(2)
	QWizard__BackgroundPixmap = QWizard__WizardPixmap(3)
	QWizard__NPixmaps         = QWizard__WizardPixmap(4)
)

//QWizard::WizardStyle
type QWizard__WizardStyle int

var (
	QWizard__ClassicStyle = QWizard__WizardStyle(0)
	QWizard__ModernStyle  = QWizard__WizardStyle(1)
	QWizard__MacStyle     = QWizard__WizardStyle(2)
	QWizard__AeroStyle    = QWizard__WizardStyle(3)
	QWizard__NStyles      = QWizard__WizardStyle(4)
)

func (ptr *QWizard) CurrentId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_CurrentId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) HasVisitedPage(id int) bool {
	if ptr.Pointer() != nil {
		return C.QWizard_HasVisitedPage(C.QtObjectPtr(ptr.Pointer()), C.int(id)) != 0
	}
	return false
}

func (ptr *QWizard) Options() QWizard__WizardOption {
	if ptr.Pointer() != nil {
		return QWizard__WizardOption(C.QWizard_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) Page(id int) *QWizardPage {
	if ptr.Pointer() != nil {
		return QWizardPageFromPointer(unsafe.Pointer(C.QWizard_Page(C.QtObjectPtr(ptr.Pointer()), C.int(id))))
	}
	return nil
}

func (ptr *QWizard) SetOptions(options QWizard__WizardOption) {
	if ptr.Pointer() != nil {
		C.QWizard_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func (ptr *QWizard) SetPage(id int, page QWizardPageITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetPage(C.QtObjectPtr(ptr.Pointer()), C.int(id), C.QtObjectPtr(PointerFromQWizardPage(page)))
	}
}

func (ptr *QWizard) SetStartId(id int) {
	if ptr.Pointer() != nil {
		C.QWizard_SetStartId(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QWizard) SetSubTitleFormat(format core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QWizard_SetSubTitleFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QWizard) SetTitleFormat(format core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QWizard_SetTitleFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QWizard) SetWizardStyle(style QWizard__WizardStyle) {
	if ptr.Pointer() != nil {
		C.QWizard_SetWizardStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QWizard) StartId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_StartId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) SubTitleFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_SubTitleFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) TitleFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_TitleFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) WizardStyle() QWizard__WizardStyle {
	if ptr.Pointer() != nil {
		return QWizard__WizardStyle(C.QWizard_WizardStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQWizard(parent QWidgetITF, flags core.Qt__WindowType) *QWizard {
	return QWizardFromPointer(unsafe.Pointer(C.QWizard_NewQWizard(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QWizard) AddPage(page QWizardPageITF) int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_AddPage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWizardPage(page))))
	}
	return 0
}

func (ptr *QWizard) Back() {
	if ptr.Pointer() != nil {
		C.QWizard_Back(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWizard) Button(which QWizard__WizardButton) *QAbstractButton {
	if ptr.Pointer() != nil {
		return QAbstractButtonFromPointer(unsafe.Pointer(C.QWizard_Button(C.QtObjectPtr(ptr.Pointer()), C.int(which))))
	}
	return nil
}

func (ptr *QWizard) ButtonText(which QWizard__WizardButton) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizard_ButtonText(C.QtObjectPtr(ptr.Pointer()), C.int(which)))
	}
	return ""
}

func (ptr *QWizard) ConnectCurrentIdChanged(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectCurrentIdChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentIdChanged", f)
	}
}

func (ptr *QWizard) DisconnectCurrentIdChanged() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCurrentIdChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentIdChanged")
	}
}

//export callbackQWizardCurrentIdChanged
func callbackQWizardCurrentIdChanged(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIdChanged").(func(int))(int(id))
}

func (ptr *QWizard) CurrentPage() *QWizardPage {
	if ptr.Pointer() != nil {
		return QWizardPageFromPointer(unsafe.Pointer(C.QWizard_CurrentPage(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWizard) ConnectCustomButtonClicked(f func(which int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectCustomButtonClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "customButtonClicked", f)
	}
}

func (ptr *QWizard) DisconnectCustomButtonClicked() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCustomButtonClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "customButtonClicked")
	}
}

//export callbackQWizardCustomButtonClicked
func callbackQWizardCustomButtonClicked(ptrName *C.char, which C.int) {
	qt.GetSignal(C.GoString(ptrName), "customButtonClicked").(func(int))(int(which))
}

func (ptr *QWizard) Field(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizard_Field(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QWizard) ConnectHelpRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectHelpRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QWizard) DisconnectHelpRequested() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectHelpRequested(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQWizardHelpRequested
func callbackQWizardHelpRequested(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "helpRequested").(func())()
}

func (ptr *QWizard) Next() {
	if ptr.Pointer() != nil {
		C.QWizard_Next(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWizard) NextId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_NextId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWizard) ConnectPageAdded(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "pageAdded", f)
	}
}

func (ptr *QWizard) DisconnectPageAdded() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageAdded(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "pageAdded")
	}
}

//export callbackQWizardPageAdded
func callbackQWizardPageAdded(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "pageAdded").(func(int))(int(id))
}

func (ptr *QWizard) ConnectPageRemoved(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "pageRemoved", f)
	}
}

func (ptr *QWizard) DisconnectPageRemoved() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageRemoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "pageRemoved")
	}
}

//export callbackQWizardPageRemoved
func callbackQWizardPageRemoved(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "pageRemoved").(func(int))(int(id))
}

func (ptr *QWizard) RemovePage(id int) {
	if ptr.Pointer() != nil {
		C.QWizard_RemovePage(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QWizard) Restart() {
	if ptr.Pointer() != nil {
		C.QWizard_Restart(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWizard) SetButton(which QWizard__WizardButton, button QAbstractButtonITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetButton(C.QtObjectPtr(ptr.Pointer()), C.int(which), C.QtObjectPtr(PointerFromQAbstractButton(button)))
	}
}

func (ptr *QWizard) SetButtonText(which QWizard__WizardButton, text string) {
	if ptr.Pointer() != nil {
		C.QWizard_SetButtonText(C.QtObjectPtr(ptr.Pointer()), C.int(which), C.CString(text))
	}
}

func (ptr *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	if ptr.Pointer() != nil {
		C.QWizard_SetDefaultProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(className), C.CString(property), C.CString(changedSignal))
	}
}

func (ptr *QWizard) SetField(name string, value string) {
	if ptr.Pointer() != nil {
		C.QWizard_SetField(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value))
	}
}

func (ptr *QWizard) SetOption(option QWizard__WizardOption, on bool) {
	if ptr.Pointer() != nil {
		C.QWizard_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWizard) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.int(which), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QWizard) SetSideWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetSideWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QWizard) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWizard_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizard) SideWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWizard_SideWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWizard) TestOption(option QWizard__WizardOption) bool {
	if ptr.Pointer() != nil {
		return C.QWizard_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QWizard) ValidateCurrentPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizard_ValidateCurrentPage(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWizard) DestroyQWizard() {
	if ptr.Pointer() != nil {
		C.QWizard_DestroyQWizard(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
