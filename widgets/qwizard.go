package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QWizard_") {
		n.SetObjectName("QWizard_" + qt.Identifier())
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

func (ptr *QWizard) ConnectCleanupPage(f func(id int)) {
	defer qt.Recovering("connect QWizard::cleanupPage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "cleanupPage", f)
	}
}

func (ptr *QWizard) DisconnectCleanupPage() {
	defer qt.Recovering("disconnect QWizard::cleanupPage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "cleanupPage")
	}
}

//export callbackQWizardCleanupPage
func callbackQWizardCleanupPage(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QWizard::cleanupPage")

	if signal := qt.GetSignal(C.GoString(ptrName), "cleanupPage"); signal != nil {
		signal.(func(int))(int(id))
	} else {
		NewQWizardFromPointer(ptr).CleanupPageDefault(int(id))
	}
}

func (ptr *QWizard) CleanupPage(id int) {
	defer qt.Recovering("QWizard::cleanupPage")

	if ptr.Pointer() != nil {
		C.QWizard_CleanupPage(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) CleanupPageDefault(id int) {
	defer qt.Recovering("QWizard::cleanupPage")

	if ptr.Pointer() != nil {
		C.QWizard_CleanupPageDefault(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) CurrentId() int {
	defer qt.Recovering("QWizard::currentId")

	if ptr.Pointer() != nil {
		return int(C.QWizard_CurrentId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) HasVisitedPage(id int) bool {
	defer qt.Recovering("QWizard::hasVisitedPage")

	if ptr.Pointer() != nil {
		return C.QWizard_HasVisitedPage(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QWizard) ConnectInitializePage(f func(id int)) {
	defer qt.Recovering("connect QWizard::initializePage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initializePage", f)
	}
}

func (ptr *QWizard) DisconnectInitializePage() {
	defer qt.Recovering("disconnect QWizard::initializePage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initializePage")
	}
}

//export callbackQWizardInitializePage
func callbackQWizardInitializePage(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QWizard::initializePage")

	if signal := qt.GetSignal(C.GoString(ptrName), "initializePage"); signal != nil {
		signal.(func(int))(int(id))
	} else {
		NewQWizardFromPointer(ptr).InitializePageDefault(int(id))
	}
}

func (ptr *QWizard) InitializePage(id int) {
	defer qt.Recovering("QWizard::initializePage")

	if ptr.Pointer() != nil {
		C.QWizard_InitializePage(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) InitializePageDefault(id int) {
	defer qt.Recovering("QWizard::initializePage")

	if ptr.Pointer() != nil {
		C.QWizard_InitializePageDefault(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) Options() QWizard__WizardOption {
	defer qt.Recovering("QWizard::options")

	if ptr.Pointer() != nil {
		return QWizard__WizardOption(C.QWizard_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) Page(id int) *QWizardPage {
	defer qt.Recovering("QWizard::page")

	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_Page(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QWizard) SetOptions(options QWizard__WizardOption) {
	defer qt.Recovering("QWizard::setOptions")

	if ptr.Pointer() != nil {
		C.QWizard_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QWizard) SetPage(id int, page QWizardPage_ITF) {
	defer qt.Recovering("QWizard::setPage")

	if ptr.Pointer() != nil {
		C.QWizard_SetPage(ptr.Pointer(), C.int(id), PointerFromQWizardPage(page))
	}
}

func (ptr *QWizard) SetStartId(id int) {
	defer qt.Recovering("QWizard::setStartId")

	if ptr.Pointer() != nil {
		C.QWizard_SetStartId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) SetSubTitleFormat(format core.Qt__TextFormat) {
	defer qt.Recovering("QWizard::setSubTitleFormat")

	if ptr.Pointer() != nil {
		C.QWizard_SetSubTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetTitleFormat(format core.Qt__TextFormat) {
	defer qt.Recovering("QWizard::setTitleFormat")

	if ptr.Pointer() != nil {
		C.QWizard_SetTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetWizardStyle(style QWizard__WizardStyle) {
	defer qt.Recovering("QWizard::setWizardStyle")

	if ptr.Pointer() != nil {
		C.QWizard_SetWizardStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QWizard) StartId() int {
	defer qt.Recovering("QWizard::startId")

	if ptr.Pointer() != nil {
		return int(C.QWizard_StartId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) SubTitleFormat() core.Qt__TextFormat {
	defer qt.Recovering("QWizard::subTitleFormat")

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_SubTitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) TitleFormat() core.Qt__TextFormat {
	defer qt.Recovering("QWizard::titleFormat")

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_TitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) WizardStyle() QWizard__WizardStyle {
	defer qt.Recovering("QWizard::wizardStyle")

	if ptr.Pointer() != nil {
		return QWizard__WizardStyle(C.QWizard_WizardStyle(ptr.Pointer()))
	}
	return 0
}

func NewQWizard(parent QWidget_ITF, flags core.Qt__WindowType) *QWizard {
	defer qt.Recovering("QWizard::QWizard")

	return NewQWizardFromPointer(C.QWizard_NewQWizard(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QWizard) AddPage(page QWizardPage_ITF) int {
	defer qt.Recovering("QWizard::addPage")

	if ptr.Pointer() != nil {
		return int(C.QWizard_AddPage(ptr.Pointer(), PointerFromQWizardPage(page)))
	}
	return 0
}

func (ptr *QWizard) Back() {
	defer qt.Recovering("QWizard::back")

	if ptr.Pointer() != nil {
		C.QWizard_Back(ptr.Pointer())
	}
}

func (ptr *QWizard) Button(which QWizard__WizardButton) *QAbstractButton {
	defer qt.Recovering("QWizard::button")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QWizard_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QWizard) ButtonText(which QWizard__WizardButton) string {
	defer qt.Recovering("QWizard::buttonText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizard_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizard) ConnectCurrentIdChanged(f func(id int)) {
	defer qt.Recovering("connect QWizard::currentIdChanged")

	if ptr.Pointer() != nil {
		C.QWizard_ConnectCurrentIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIdChanged", f)
	}
}

func (ptr *QWizard) DisconnectCurrentIdChanged() {
	defer qt.Recovering("disconnect QWizard::currentIdChanged")

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCurrentIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIdChanged")
	}
}

//export callbackQWizardCurrentIdChanged
func callbackQWizardCurrentIdChanged(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QWizard::currentIdChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentIdChanged"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QWizard) CurrentIdChanged(id int) {
	defer qt.Recovering("QWizard::currentIdChanged")

	if ptr.Pointer() != nil {
		C.QWizard_CurrentIdChanged(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) CurrentPage() *QWizardPage {
	defer qt.Recovering("QWizard::currentPage")

	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_CurrentPage(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) ConnectCustomButtonClicked(f func(which int)) {
	defer qt.Recovering("connect QWizard::customButtonClicked")

	if ptr.Pointer() != nil {
		C.QWizard_ConnectCustomButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "customButtonClicked", f)
	}
}

func (ptr *QWizard) DisconnectCustomButtonClicked() {
	defer qt.Recovering("disconnect QWizard::customButtonClicked")

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCustomButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "customButtonClicked")
	}
}

//export callbackQWizardCustomButtonClicked
func callbackQWizardCustomButtonClicked(ptr unsafe.Pointer, ptrName *C.char, which C.int) {
	defer qt.Recovering("callback QWizard::customButtonClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "customButtonClicked"); signal != nil {
		signal.(func(int))(int(which))
	}

}

func (ptr *QWizard) CustomButtonClicked(which int) {
	defer qt.Recovering("QWizard::customButtonClicked")

	if ptr.Pointer() != nil {
		C.QWizard_CustomButtonClicked(ptr.Pointer(), C.int(which))
	}
}

func (ptr *QWizard) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QWizard::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QWizard) DisconnectDone() {
	defer qt.Recovering("disconnect QWizard::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQWizardDone
func callbackQWizardDone(ptr unsafe.Pointer, ptrName *C.char, result C.int) {
	defer qt.Recovering("callback QWizard::done")

	if signal := qt.GetSignal(C.GoString(ptrName), "done"); signal != nil {
		signal.(func(int))(int(result))
	} else {
		NewQWizardFromPointer(ptr).DoneDefault(int(result))
	}
}

func (ptr *QWizard) Done(result int) {
	defer qt.Recovering("QWizard::done")

	if ptr.Pointer() != nil {
		C.QWizard_Done(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QWizard) DoneDefault(result int) {
	defer qt.Recovering("QWizard::done")

	if ptr.Pointer() != nil {
		C.QWizard_DoneDefault(ptr.Pointer(), C.int(result))
	}
}

func (ptr *QWizard) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QWizard::event")

	if ptr.Pointer() != nil {
		return C.QWizard_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWizard) Field(name string) *core.QVariant {
	defer qt.Recovering("QWizard::field")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWizard_Field(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QWizard) ConnectHelpRequested(f func()) {
	defer qt.Recovering("connect QWizard::helpRequested")

	if ptr.Pointer() != nil {
		C.QWizard_ConnectHelpRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QWizard) DisconnectHelpRequested() {
	defer qt.Recovering("disconnect QWizard::helpRequested")

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectHelpRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQWizardHelpRequested
func callbackQWizardHelpRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWizard::helpRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "helpRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWizard) HelpRequested() {
	defer qt.Recovering("QWizard::helpRequested")

	if ptr.Pointer() != nil {
		C.QWizard_HelpRequested(ptr.Pointer())
	}
}

func (ptr *QWizard) Next() {
	defer qt.Recovering("QWizard::next")

	if ptr.Pointer() != nil {
		C.QWizard_Next(ptr.Pointer())
	}
}

func (ptr *QWizard) NextId() int {
	defer qt.Recovering("QWizard::nextId")

	if ptr.Pointer() != nil {
		return int(C.QWizard_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) ConnectPageAdded(f func(id int)) {
	defer qt.Recovering("connect QWizard::pageAdded")

	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageAdded", f)
	}
}

func (ptr *QWizard) DisconnectPageAdded() {
	defer qt.Recovering("disconnect QWizard::pageAdded")

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageAdded")
	}
}

//export callbackQWizardPageAdded
func callbackQWizardPageAdded(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QWizard::pageAdded")

	if signal := qt.GetSignal(C.GoString(ptrName), "pageAdded"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QWizard) PageAdded(id int) {
	defer qt.Recovering("QWizard::pageAdded")

	if ptr.Pointer() != nil {
		C.QWizard_PageAdded(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) ConnectPageRemoved(f func(id int)) {
	defer qt.Recovering("connect QWizard::pageRemoved")

	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageRemoved", f)
	}
}

func (ptr *QWizard) DisconnectPageRemoved() {
	defer qt.Recovering("disconnect QWizard::pageRemoved")

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageRemoved")
	}
}

//export callbackQWizardPageRemoved
func callbackQWizardPageRemoved(ptr unsafe.Pointer, ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QWizard::pageRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "pageRemoved"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QWizard) PageRemoved(id int) {
	defer qt.Recovering("QWizard::pageRemoved")

	if ptr.Pointer() != nil {
		C.QWizard_PageRemoved(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWizard::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWizard) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWizard::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQWizardPaintEvent
func callbackQWizardPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWizard) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWizard::paintEvent")

	if ptr.Pointer() != nil {
		C.QWizard_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWizard) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWizard::paintEvent")

	if ptr.Pointer() != nil {
		C.QWizard_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWizard) RemovePage(id int) {
	defer qt.Recovering("QWizard::removePage")

	if ptr.Pointer() != nil {
		C.QWizard_RemovePage(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWizard::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWizard) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWizard::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQWizardResizeEvent
func callbackQWizardResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWizard) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWizard::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWizard) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWizard::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWizard) Restart() {
	defer qt.Recovering("QWizard::restart")

	if ptr.Pointer() != nil {
		C.QWizard_Restart(ptr.Pointer())
	}
}

func (ptr *QWizard) SetButton(which QWizard__WizardButton, button QAbstractButton_ITF) {
	defer qt.Recovering("QWizard::setButton")

	if ptr.Pointer() != nil {
		C.QWizard_SetButton(ptr.Pointer(), C.int(which), PointerFromQAbstractButton(button))
	}
}

func (ptr *QWizard) SetButtonText(which QWizard__WizardButton, text string) {
	defer qt.Recovering("QWizard::setButtonText")

	if ptr.Pointer() != nil {
		C.QWizard_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	defer qt.Recovering("QWizard::setDefaultProperty")

	if ptr.Pointer() != nil {
		C.QWizard_SetDefaultProperty(ptr.Pointer(), C.CString(className), C.CString(property), C.CString(changedSignal))
	}
}

func (ptr *QWizard) SetField(name string, value core.QVariant_ITF) {
	defer qt.Recovering("QWizard::setField")

	if ptr.Pointer() != nil {
		C.QWizard_SetField(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QWizard) SetOption(option QWizard__WizardOption, on bool) {
	defer qt.Recovering("QWizard::setOption")

	if ptr.Pointer() != nil {
		C.QWizard_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWizard) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QWizard::setPixmap")

	if ptr.Pointer() != nil {
		C.QWizard_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizard) SetSideWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWizard::setSideWidget")

	if ptr.Pointer() != nil {
		C.QWizard_SetSideWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWizard) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWizard::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWizard) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWizard::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQWizardSetVisible
func callbackQWizardSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWizard::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQWizardFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QWizard) SetVisible(visible bool) {
	defer qt.Recovering("QWizard::setVisible")

	if ptr.Pointer() != nil {
		C.QWizard_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizard) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QWizard::setVisible")

	if ptr.Pointer() != nil {
		C.QWizard_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizard) SideWidget() *QWidget {
	defer qt.Recovering("QWizard::sideWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWizard_SideWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) SizeHint() *core.QSize {
	defer qt.Recovering("QWizard::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWizard_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) TestOption(option QWizard__WizardOption) bool {
	defer qt.Recovering("QWizard::testOption")

	if ptr.Pointer() != nil {
		return C.QWizard_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QWizard) ValidateCurrentPage() bool {
	defer qt.Recovering("QWizard::validateCurrentPage")

	if ptr.Pointer() != nil {
		return C.QWizard_ValidateCurrentPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizard) DestroyQWizard() {
	defer qt.Recovering("QWizard::~QWizard")

	if ptr.Pointer() != nil {
		C.QWizard_DestroyQWizard(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWizard) ConnectAccept(f func()) {
	defer qt.Recovering("connect QWizard::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QWizard) DisconnectAccept() {
	defer qt.Recovering("disconnect QWizard::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQWizardAccept
func callbackQWizardAccept(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QWizard::accept")

	if signal := qt.GetSignal(C.GoString(ptrName), "accept"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QWizard) Accept() {
	defer qt.Recovering("QWizard::accept")

	if ptr.Pointer() != nil {
		C.QWizard_Accept(ptr.Pointer())
	}
}

func (ptr *QWizard) AcceptDefault() {
	defer qt.Recovering("QWizard::accept")

	if ptr.Pointer() != nil {
		C.QWizard_AcceptDefault(ptr.Pointer())
	}
}

func (ptr *QWizard) ConnectCloseEvent(f func(e *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWizard::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWizard) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWizard::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQWizardCloseEvent
func callbackQWizardCloseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(e))
	} else {
		NewQWizardFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(e))
	}
}

func (ptr *QWizard) CloseEvent(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWizard::closeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QWizard) CloseEventDefault(e gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWizard::closeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(e))
	}
}

func (ptr *QWizard) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWizard::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWizard) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWizard::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQWizardContextMenuEvent
func callbackQWizardContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQWizardFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QWizard) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWizard::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QWizard) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWizard::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QWizard) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWizard::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWizard) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWizard::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQWizardKeyPressEvent
func callbackQWizardKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQWizardFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QWizard) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizard::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWizard_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QWizard) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizard::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWizard_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QWizard) ConnectOpen(f func()) {
	defer qt.Recovering("connect QWizard::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QWizard) DisconnectOpen() {
	defer qt.Recovering("disconnect QWizard::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

//export callbackQWizardOpen
func callbackQWizardOpen(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QWizard::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QWizard) Open() {
	defer qt.Recovering("QWizard::open")

	if ptr.Pointer() != nil {
		C.QWizard_Open(ptr.Pointer())
	}
}

func (ptr *QWizard) OpenDefault() {
	defer qt.Recovering("QWizard::open")

	if ptr.Pointer() != nil {
		C.QWizard_OpenDefault(ptr.Pointer())
	}
}

func (ptr *QWizard) ConnectReject(f func()) {
	defer qt.Recovering("connect QWizard::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reject", f)
	}
}

func (ptr *QWizard) DisconnectReject() {
	defer qt.Recovering("disconnect QWizard::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reject")
	}
}

//export callbackQWizardReject
func callbackQWizardReject(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QWizard::reject")

	if signal := qt.GetSignal(C.GoString(ptrName), "reject"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QWizard) Reject() {
	defer qt.Recovering("QWizard::reject")

	if ptr.Pointer() != nil {
		C.QWizard_Reject(ptr.Pointer())
	}
}

func (ptr *QWizard) RejectDefault() {
	defer qt.Recovering("QWizard::reject")

	if ptr.Pointer() != nil {
		C.QWizard_RejectDefault(ptr.Pointer())
	}
}

func (ptr *QWizard) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWizard::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWizard) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWizard::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQWizardShowEvent
func callbackQWizardShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWizard) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWizard::showEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWizard) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWizard::showEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWizard) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWizard::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWizard) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWizard::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQWizardActionEvent
func callbackQWizardActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWizard) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWizard::actionEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWizard) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWizard::actionEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWizard) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWizard::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWizard) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWizard::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQWizardDragEnterEvent
func callbackQWizardDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWizard) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWizard::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWizard) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWizard::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWizard) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWizard::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWizard) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWizard::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQWizardDragLeaveEvent
func callbackQWizardDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWizard) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWizard::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWizard) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWizard::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWizard) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWizard::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWizard) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWizard::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQWizardDragMoveEvent
func callbackQWizardDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWizard) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWizard::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWizard) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWizard::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWizard) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QWizard::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWizard) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWizard::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQWizardDropEvent
func callbackQWizardDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWizard) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWizard::dropEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWizard) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWizard::dropEvent")

	if ptr.Pointer() != nil {
		C.QWizard_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWizard) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizard::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWizard) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWizard::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQWizardEnterEvent
func callbackQWizardEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizard) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::enterEvent")

	if ptr.Pointer() != nil {
		C.QWizard_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::enterEvent")

	if ptr.Pointer() != nil {
		C.QWizard_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWizard::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWizard) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWizard::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQWizardFocusInEvent
func callbackQWizardFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWizard) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizard::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWizard_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizard) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizard::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWizard_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizard) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWizard::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWizard) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWizard::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQWizardFocusOutEvent
func callbackQWizardFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWizard) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizard::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWizard_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizard) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizard::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWizard_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizard) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWizard::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWizard) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWizard::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQWizardHideEvent
func callbackQWizardHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWizard) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWizard::hideEvent")

	if ptr.Pointer() != nil {
		C.QWizard_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWizard) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWizard::hideEvent")

	if ptr.Pointer() != nil {
		C.QWizard_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWizard) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizard::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWizard) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWizard::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQWizardLeaveEvent
func callbackQWizardLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizard) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWizard::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWizard) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWizard::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQWizardMoveEvent
func callbackQWizardMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWizard) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWizard::moveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWizard) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWizard::moveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWizard) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizard::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QWizard) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QWizard::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQWizardChangeEvent
func callbackQWizardChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizard) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::changeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::changeEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QWizard::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QWizard) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QWizard::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQWizardInitPainter
func callbackQWizardInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQWizardFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QWizard) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWizard::initPainter")

	if ptr.Pointer() != nil {
		C.QWizard_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWizard) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWizard::initPainter")

	if ptr.Pointer() != nil {
		C.QWizard_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWizard) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWizard::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWizard) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWizard::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQWizardInputMethodEvent
func callbackQWizardInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWizard) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWizard::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWizard_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWizard) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWizard::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWizard_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWizard) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWizard::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWizard) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWizard::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQWizardKeyReleaseEvent
func callbackQWizardKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWizard) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizard::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizard_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizard) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizard::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizard_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizard) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizard::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWizard) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWizard::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQWizardMouseDoubleClickEvent
func callbackQWizardMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizard) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizard::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWizard) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWizard::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQWizardMouseMoveEvent
func callbackQWizardMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizard) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizard::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWizard) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWizard::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQWizardMousePressEvent
func callbackQWizardMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizard) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizard::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWizard) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWizard::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQWizardMouseReleaseEvent
func callbackQWizardMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizard) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizard::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizard_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizard) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWizard::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWizard) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWizard::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQWizardTabletEvent
func callbackQWizardTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWizard) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWizard::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWizard_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWizard) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWizard::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWizard_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWizard) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWizard::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWizard) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWizard::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQWizardWheelEvent
func callbackQWizardWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWizard) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWizard::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWizard_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWizard) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWizard::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWizard_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWizard) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWizard::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWizard) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWizard::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWizardTimerEvent
func callbackQWizardTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWizard) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWizard::timerEvent")

	if ptr.Pointer() != nil {
		C.QWizard_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWizard) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWizard::timerEvent")

	if ptr.Pointer() != nil {
		C.QWizard_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWizard) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWizard::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWizard) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWizard::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWizardChildEvent
func callbackQWizardChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWizard) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWizard::childEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWizard) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWizard::childEvent")

	if ptr.Pointer() != nil {
		C.QWizard_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWizard) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizard::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWizard) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWizard::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWizardCustomEvent
func callbackQWizardCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizard::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizard) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::customEvent")

	if ptr.Pointer() != nil {
		C.QWizard_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizard) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizard::customEvent")

	if ptr.Pointer() != nil {
		C.QWizard_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
