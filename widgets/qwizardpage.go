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
func callbackQWizardPageActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPagePaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QWizardPage::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQWizardPageShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQWizardPageInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQWizardPageCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWizardPage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
