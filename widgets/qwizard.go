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

type QWizard_ITF interface {
	QDialog_ITF
	QWizard_PTR() *QWizard
}

func PointerFromQWizard(ptr QWizard_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWizard_PTR().Pointer()
	}
	return nil
}

func NewQWizardFromPointer(ptr unsafe.Pointer) *QWizard {
	var n = new(QWizard)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QWizard_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizard) QWizard_PTR() *QWizard {
	return ptr
}

//QWizard::WizardButton
type QWizard__WizardButton int64

const (
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
type QWizard__WizardOption int64

const (
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
type QWizard__WizardPixmap int64

const (
	QWizard__WatermarkPixmap  = QWizard__WizardPixmap(0)
	QWizard__LogoPixmap       = QWizard__WizardPixmap(1)
	QWizard__BannerPixmap     = QWizard__WizardPixmap(2)
	QWizard__BackgroundPixmap = QWizard__WizardPixmap(3)
	QWizard__NPixmaps         = QWizard__WizardPixmap(4)
)

//QWizard::WizardStyle
type QWizard__WizardStyle int64

var (
	QWizard__ClassicStyle = QWizard__WizardStyle(0)
	QWizard__ModernStyle  = QWizard__WizardStyle(1)
	QWizard__MacStyle     = QWizard__WizardStyle(2)
	QWizard__AeroStyle    = QWizard__WizardStyle(3)
	QWizard__NStyles      = QWizard__WizardStyle(4)
)

func (ptr *QWizard) CurrentId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_CurrentId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) HasVisitedPage(id int) bool {
	if ptr.Pointer() != nil {
		return C.QWizard_HasVisitedPage(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QWizard) Options() QWizard__WizardOption {
	if ptr.Pointer() != nil {
		return QWizard__WizardOption(C.QWizard_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) Page(id int) *QWizardPage {
	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_Page(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QWizard) SetOptions(options QWizard__WizardOption) {
	if ptr.Pointer() != nil {
		C.QWizard_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QWizard) SetPage(id int, page QWizardPage_ITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetPage(ptr.Pointer(), C.int(id), PointerFromQWizardPage(page))
	}
}

func (ptr *QWizard) SetStartId(id int) {
	if ptr.Pointer() != nil {
		C.QWizard_SetStartId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) SetSubTitleFormat(format core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QWizard_SetSubTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetTitleFormat(format core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QWizard_SetTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetWizardStyle(style QWizard__WizardStyle) {
	if ptr.Pointer() != nil {
		C.QWizard_SetWizardStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QWizard) StartId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_StartId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) SubTitleFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_SubTitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) TitleFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_TitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) WizardStyle() QWizard__WizardStyle {
	if ptr.Pointer() != nil {
		return QWizard__WizardStyle(C.QWizard_WizardStyle(ptr.Pointer()))
	}
	return 0
}

func NewQWizard(parent QWidget_ITF, flags core.Qt__WindowType) *QWizard {
	return NewQWizardFromPointer(C.QWizard_NewQWizard(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QWizard) AddPage(page QWizardPage_ITF) int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_AddPage(ptr.Pointer(), PointerFromQWizardPage(page)))
	}
	return 0
}

func (ptr *QWizard) Back() {
	if ptr.Pointer() != nil {
		C.QWizard_Back(ptr.Pointer())
	}
}

func (ptr *QWizard) Button(which QWizard__WizardButton) *QAbstractButton {
	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QWizard_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QWizard) ButtonText(which QWizard__WizardButton) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWizard_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizard) ConnectCurrentIdChanged(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectCurrentIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIdChanged", f)
	}
}

func (ptr *QWizard) DisconnectCurrentIdChanged() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCurrentIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIdChanged")
	}
}

//export callbackQWizardCurrentIdChanged
func callbackQWizardCurrentIdChanged(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "currentIdChanged").(func(int))(int(id))
}

func (ptr *QWizard) CurrentPage() *QWizardPage {
	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_CurrentPage(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) ConnectCustomButtonClicked(f func(which int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectCustomButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "customButtonClicked", f)
	}
}

func (ptr *QWizard) DisconnectCustomButtonClicked() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCustomButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "customButtonClicked")
	}
}

//export callbackQWizardCustomButtonClicked
func callbackQWizardCustomButtonClicked(ptrName *C.char, which C.int) {
	qt.GetSignal(C.GoString(ptrName), "customButtonClicked").(func(int))(int(which))
}

func (ptr *QWizard) Field(name string) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWizard_Field(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QWizard) ConnectHelpRequested(f func()) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectHelpRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QWizard) DisconnectHelpRequested() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectHelpRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQWizardHelpRequested
func callbackQWizardHelpRequested(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "helpRequested").(func())()
}

func (ptr *QWizard) Next() {
	if ptr.Pointer() != nil {
		C.QWizard_Next(ptr.Pointer())
	}
}

func (ptr *QWizard) NextId() int {
	if ptr.Pointer() != nil {
		return int(C.QWizard_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) ConnectPageAdded(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageAdded", f)
	}
}

func (ptr *QWizard) DisconnectPageAdded() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageAdded")
	}
}

//export callbackQWizardPageAdded
func callbackQWizardPageAdded(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "pageAdded").(func(int))(int(id))
}

func (ptr *QWizard) ConnectPageRemoved(f func(id int)) {
	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageRemoved", f)
	}
}

func (ptr *QWizard) DisconnectPageRemoved() {
	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageRemoved")
	}
}

//export callbackQWizardPageRemoved
func callbackQWizardPageRemoved(ptrName *C.char, id C.int) {
	qt.GetSignal(C.GoString(ptrName), "pageRemoved").(func(int))(int(id))
}

func (ptr *QWizard) RemovePage(id int) {
	if ptr.Pointer() != nil {
		C.QWizard_RemovePage(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) Restart() {
	if ptr.Pointer() != nil {
		C.QWizard_Restart(ptr.Pointer())
	}
}

func (ptr *QWizard) SetButton(which QWizard__WizardButton, button QAbstractButton_ITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetButton(ptr.Pointer(), C.int(which), PointerFromQAbstractButton(button))
	}
}

func (ptr *QWizard) SetButtonText(which QWizard__WizardButton, text string) {
	if ptr.Pointer() != nil {
		C.QWizard_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	if ptr.Pointer() != nil {
		C.QWizard_SetDefaultProperty(ptr.Pointer(), C.CString(className), C.CString(property), C.CString(changedSignal))
	}
}

func (ptr *QWizard) SetField(name string, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetField(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QWizard) SetOption(option QWizard__WizardOption, on bool) {
	if ptr.Pointer() != nil {
		C.QWizard_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWizard) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizard) SetSideWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWizard_SetSideWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWizard) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWizard_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizard) SideWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWizard_SideWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) TestOption(option QWizard__WizardOption) bool {
	if ptr.Pointer() != nil {
		return C.QWizard_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QWizard) ValidateCurrentPage() bool {
	if ptr.Pointer() != nil {
		return C.QWizard_ValidateCurrentPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizard) DestroyQWizard() {
	if ptr.Pointer() != nil {
		C.QWizard_DestroyQWizard(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
