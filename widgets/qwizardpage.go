package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
func callbackQWizardPageCleanupPage(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWizardPage::cleanupPage")

	if signal := qt.GetSignal(C.GoString(ptrName), "cleanupPage"); signal != nil {
		signal.(func())()
	} else {
		NewQWizardPageFromPointer(ptr).CleanupPageDefault()
	}
}

func (ptr *QWizardPage) CleanupPage() {
	defer qt.Recovering("QWizardPage::cleanupPage")

	if ptr.Pointer() != nil {
		C.QWizardPage_CleanupPage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) CleanupPageDefault() {
	defer qt.Recovering("QWizardPage::cleanupPage")

	if ptr.Pointer() != nil {
		C.QWizardPage_CleanupPageDefault(ptr.Pointer())
	}
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
func callbackQWizardPageCompleteChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWizardPage::completeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "completeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWizardPage) CompleteChanged() {
	defer qt.Recovering("QWizardPage::completeChanged")

	if ptr.Pointer() != nil {
		C.QWizardPage_CompleteChanged(ptr.Pointer())
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
func callbackQWizardPageInitializePage(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWizardPage::initializePage")

	if signal := qt.GetSignal(C.GoString(ptrName), "initializePage"); signal != nil {
		signal.(func())()
	} else {
		NewQWizardPageFromPointer(ptr).InitializePageDefault()
	}
}

func (ptr *QWizardPage) InitializePage() {
	defer qt.Recovering("QWizardPage::initializePage")

	if ptr.Pointer() != nil {
		C.QWizardPage_InitializePage(ptr.Pointer())
	}
}

func (ptr *QWizardPage) InitializePageDefault() {
	defer qt.Recovering("QWizardPage::initializePage")

	if ptr.Pointer() != nil {
		C.QWizardPage_InitializePageDefault(ptr.Pointer())
	}
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

func (ptr *QWizardPage) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWizardPage::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWizardPage::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQWizardPageActionEvent
func callbackQWizardPageActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWizardPage::actionEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWizardPage) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWizardPage::actionEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWizardPage) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWizardPage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWizardPage::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQWizardPageDragEnterEvent
func callbackQWizardPageDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWizardPage) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWizardPage) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWizardPage) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWizardPage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWizardPage::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQWizardPageDragLeaveEvent
func callbackQWizardPageDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWizardPage) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWizardPage) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWizardPage) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWizardPage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWizardPage::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQWizardPageDragMoveEvent
func callbackQWizardPageDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWizardPage) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWizardPage) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWizardPage::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWizardPage) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QWizardPage::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWizardPage::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQWizardPageDropEvent
func callbackQWizardPageDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWizardPage) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWizardPage::dropEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWizardPage) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWizardPage::dropEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWizardPage) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizardPage::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWizardPage::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQWizardPageEnterEvent
func callbackQWizardPageEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizardPage) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::enterEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::enterEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWizardPage::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWizardPage::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQWizardPageFocusInEvent
func callbackQWizardPageFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWizardPage) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizardPage::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizardPage) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizardPage::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizardPage) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWizardPage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWizardPage::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQWizardPageFocusOutEvent
func callbackQWizardPageFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWizardPage) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizardPage::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizardPage) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWizardPage::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWizardPage) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWizardPage::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWizardPage::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQWizardPageHideEvent
func callbackQWizardPageHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWizardPage) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWizardPage::hideEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWizardPage) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWizardPage::hideEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWizardPage) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizardPage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWizardPage::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQWizardPageLeaveEvent
func callbackQWizardPageLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizardPage) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWizardPage::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWizardPage::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQWizardPageMoveEvent
func callbackQWizardPageMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWizardPage) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWizardPage::moveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWizardPage) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWizardPage::moveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWizardPage) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWizardPage::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWizardPage::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQWizardPagePaintEvent
func callbackQWizardPagePaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWizardPage) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWizardPage::paintEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWizardPage) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWizardPage::paintEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWizardPage) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWizardPage::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWizardPage) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWizardPage::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQWizardPageSetVisible
func callbackQWizardPageSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QWizardPage::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QWizardPage) SetVisible(visible bool) {
	defer qt.Recovering("QWizardPage::setVisible")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizardPage) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QWizardPage::setVisible")

	if ptr.Pointer() != nil {
		C.QWizardPage_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizardPage) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWizardPage::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWizardPage::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQWizardPageShowEvent
func callbackQWizardPageShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWizardPage::showEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWizardPage) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWizardPage::showEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWizardPage) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizardPage::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QWizardPage::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQWizardPageChangeEvent
func callbackQWizardPageChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::changeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::changeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWizardPage::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWizardPage::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQWizardPageCloseEvent
func callbackQWizardPageCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWizardPage) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWizardPage::closeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWizardPage) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWizardPage::closeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWizardPage) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWizardPage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWizardPage::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQWizardPageContextMenuEvent
func callbackQWizardPageContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWizardPage::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWizardPage) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWizardPage::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWizardPage) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QWizardPage::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QWizardPage) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QWizardPage::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQWizardPageInitPainter
func callbackQWizardPageInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQWizardPageFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QWizardPage) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWizardPage::initPainter")

	if ptr.Pointer() != nil {
		C.QWizardPage_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWizardPage) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWizardPage::initPainter")

	if ptr.Pointer() != nil {
		C.QWizardPage_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWizardPage) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWizardPage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWizardPage::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQWizardPageInputMethodEvent
func callbackQWizardPageInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWizardPage) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWizardPage::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWizardPage) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWizardPage::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWizardPage) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWizardPage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWizardPage::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQWizardPageKeyPressEvent
func callbackQWizardPageKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWizardPage) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizardPage::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizardPage) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizardPage::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizardPage) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWizardPage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWizardPage::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQWizardPageKeyReleaseEvent
func callbackQWizardPageKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWizardPage) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizardPage::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizardPage) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWizardPage::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWizardPage) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizardPage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWizardPage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQWizardPageMouseDoubleClickEvent
func callbackQWizardPageMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizardPage) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizardPage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWizardPage::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQWizardPageMouseMoveEvent
func callbackQWizardPageMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizardPage) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizardPage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWizardPage::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQWizardPageMousePressEvent
func callbackQWizardPageMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizardPage) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWizardPage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWizardPage::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQWizardPageMouseReleaseEvent
func callbackQWizardPageMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWizardPage) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWizardPage::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWizardPage) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWizardPage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWizardPage::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQWizardPageResizeEvent
func callbackQWizardPageResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWizardPage::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWizardPage) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWizardPage::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWizardPage) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWizardPage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWizardPage::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQWizardPageTabletEvent
func callbackQWizardPageTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWizardPage) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWizardPage::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWizardPage) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWizardPage::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWizardPage) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWizardPage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWizardPage::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQWizardPageWheelEvent
func callbackQWizardPageWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWizardPage) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWizardPage::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWizardPage) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWizardPage::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWizardPage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWizardPage::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWizardPage::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWizardPageTimerEvent
func callbackQWizardPageTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWizardPage) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWizardPage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWizardPage) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWizardPage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWizardPage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWizardPage::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWizardPage::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWizardPageChildEvent
func callbackQWizardPageChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWizardPage) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWizardPage::childEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWizardPage) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWizardPage::childEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWizardPage) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWizardPage::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWizardPage) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWizardPage::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWizardPageCustomEvent
func callbackQWizardPageCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWizardPage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWizardPageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWizardPage) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::customEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWizardPage) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWizardPage::customEvent")

	if ptr.Pointer() != nil {
		C.QWizardPage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
