package webkit

//#include "webkit.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	"github.com/therecipe/qt/printsupport"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

type QGraphicsWebView struct {
	widgets.QGraphicsWidget
}

type QGraphicsWebView_ITF interface {
	widgets.QGraphicsWidget_ITF
	QGraphicsWebView_PTR() *QGraphicsWebView
}

func (p *QGraphicsWebView) QGraphicsWebView_PTR() *QGraphicsWebView {
	return p
}

func (p *QGraphicsWebView) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func (p *QGraphicsWebView) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QGraphicsWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQGraphicsWebView(ptr QGraphicsWebView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWebView_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWebViewFromPointer(ptr unsafe.Pointer) *QGraphicsWebView {
	var n = new(QGraphicsWebView)
	n.SetPointer(ptr)
	return n
}

func newQGraphicsWebViewFromPointer(ptr unsafe.Pointer) *QGraphicsWebView {
	var n = NewQGraphicsWebViewFromPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsWebView_") {
		n.SetObjectName("QGraphicsWebView_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsWebView) Icon() *gui.QIcon {
	defer qt.Recovering("QGraphicsWebView::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QGraphicsWebView_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) IsModified() bool {
	defer qt.Recovering("QGraphicsWebView::isModified")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) IsTiledBackingStoreFrozen() bool {
	defer qt.Recovering("QGraphicsWebView::isTiledBackingStoreFrozen")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsTiledBackingStoreFrozen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) Load2(request network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body string) {
	defer qt.Recovering("QGraphicsWebView::load")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.int(operation), C.CString(body))
	}
}

func (ptr *QGraphicsWebView) ResizesToContents() bool {
	defer qt.Recovering("QGraphicsWebView::resizesToContents")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_ResizesToContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) SetResizesToContents(enabled bool) {
	defer qt.Recovering("QGraphicsWebView::setResizesToContents")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetResizesToContents(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWebView) SetTiledBackingStoreFrozen(frozen bool) {
	defer qt.Recovering("QGraphicsWebView::setTiledBackingStoreFrozen")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetTiledBackingStoreFrozen(ptr.Pointer(), C.int(qt.GoBoolToInt(frozen)))
	}
}

func (ptr *QGraphicsWebView) SetUrl(vqu core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::setUrl")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

func (ptr *QGraphicsWebView) SetZoomFactor(vqr float64) {
	defer qt.Recovering("QGraphicsWebView::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetZoomFactor(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QGraphicsWebView) Title() string {
	defer qt.Recovering("QGraphicsWebView::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsWebView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsWebView) Url() *core.QUrl {
	defer qt.Recovering("QGraphicsWebView::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QGraphicsWebView_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) ZoomFactor() float64 {
	defer qt.Recovering("QGraphicsWebView::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsWebView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsWebView(parent widgets.QGraphicsItem_ITF) *QGraphicsWebView {
	defer qt.Recovering("QGraphicsWebView::QGraphicsWebView")

	return newQGraphicsWebViewFromPointer(C.QGraphicsWebView_NewQGraphicsWebView(widgets.PointerFromQGraphicsItem(parent)))
}

//export callbackQGraphicsWebView_Back
func callbackQGraphicsWebView_Back(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::back")

	if signal := qt.GetSignal(C.GoString(ptrName), "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectBack(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::back")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "back", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectBack() {
	defer qt.Recovering("disconnect QGraphicsWebView::back")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "back")
	}
}

func (ptr *QGraphicsWebView) Back() {
	defer qt.Recovering("QGraphicsWebView::back")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Back(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_ContextMenuEvent
func callbackQGraphicsWebView_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneContextMenuEvent))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectContextMenuEvent(f func(ev *widgets.QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QGraphicsWebView) ContextMenuEvent(ev widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ContextMenuEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(ev))
	}
}

func (ptr *QGraphicsWebView) ContextMenuEventDefault(ev widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragEnterEvent
func callbackQGraphicsWebView_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectDragEnterEvent(f func(ev *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QGraphicsWebView) DragEnterEvent(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DragEnterEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragLeaveEvent
func callbackQGraphicsWebView_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectDragLeaveEvent(f func(ev *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QGraphicsWebView) DragLeaveEvent(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DragLeaveEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DragMoveEvent
func callbackQGraphicsWebView_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectDragMoveEvent(f func(ev *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QGraphicsWebView) DragMoveEvent(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DragMoveEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_DropEvent
func callbackQGraphicsWebView_DropEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectDropEvent(f func(ev *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QGraphicsWebView) DropEvent(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DropEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DropEventDefault(ev widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(ev))
	}
}

//export callbackQGraphicsWebView_Event
func callbackQGraphicsWebView_Event(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsWebView) ConnectEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsWebView::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QGraphicsWebView) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) EventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) FindText(subString string, options QWebPage__FindFlag) bool {
	defer qt.Recovering("QGraphicsWebView::findText")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_FindText(ptr.Pointer(), C.CString(subString), C.int(options)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_FocusInEvent
func callbackQGraphicsWebView_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QGraphicsWebView) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QGraphicsWebView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQGraphicsWebView_FocusNextPrevChild
func callbackQGraphicsWebView_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QGraphicsWebView::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QGraphicsWebView) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QGraphicsWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QGraphicsWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QGraphicsWebView) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QGraphicsWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QGraphicsWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQGraphicsWebView_FocusOutEvent
func callbackQGraphicsWebView_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QGraphicsWebView) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QGraphicsWebView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQGraphicsWebView_Forward
func callbackQGraphicsWebView_Forward(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::forward")

	if signal := qt.GetSignal(C.GoString(ptrName), "forward"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectForward(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::forward")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "forward", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectForward() {
	defer qt.Recovering("disconnect QGraphicsWebView::forward")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "forward")
	}
}

func (ptr *QGraphicsWebView) Forward() {
	defer qt.Recovering("QGraphicsWebView::forward")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Forward(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) History() *QWebHistory {
	defer qt.Recovering("QGraphicsWebView::history")

	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QGraphicsWebView_History(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_HoverLeaveEvent
func callbackQGraphicsWebView_HoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectHoverLeaveEvent(f func(ev *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

func (ptr *QGraphicsWebView) HoverLeaveEvent(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

func (ptr *QGraphicsWebView) HoverLeaveEventDefault(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

//export callbackQGraphicsWebView_HoverMoveEvent
func callbackQGraphicsWebView_HoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectHoverMoveEvent(f func(ev *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

func (ptr *QGraphicsWebView) HoverMoveEvent(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

func (ptr *QGraphicsWebView) HoverMoveEventDefault(ev widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(ev))
	}
}

//export callbackQGraphicsWebView_IconChanged
func callbackQGraphicsWebView_IconChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::iconChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectIconChanged(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectIconChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconChanged", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectIconChanged() {
	defer qt.Recovering("disconnect QGraphicsWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconChanged")
	}
}

func (ptr *QGraphicsWebView) IconChanged() {
	defer qt.Recovering("QGraphicsWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_IconChanged(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_InputMethodEvent
func callbackQGraphicsWebView_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectInputMethodEvent(f func(ev *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QGraphicsWebView) InputMethodEvent(ev gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(ev))
	}
}

func (ptr *QGraphicsWebView) InputMethodEventDefault(ev gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(ev))
	}
}

//export callbackQGraphicsWebView_InputMethodQuery
func callbackQGraphicsWebView_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQGraphicsWebViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QGraphicsWebView) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsWebView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QGraphicsWebView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QGraphicsWebView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsWebView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsWebView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsWebView) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsWebView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsWebView_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQGraphicsWebView_ItemChange
func callbackQGraphicsWebView_ItemChange(ptr unsafe.Pointer, ptrName *C.char, change C.int, value unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::itemChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "itemChange"); signal != nil {
		return core.PointerFromQVariant(signal.(func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQGraphicsWebViewFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QGraphicsWebView) ConnectItemChange(f func(change widgets.QGraphicsItem__GraphicsItemChange, value *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsWebView::itemChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "itemChange", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectItemChange() {
	defer qt.Recovering("disconnect QGraphicsWebView::itemChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "itemChange")
	}
}

func (ptr *QGraphicsWebView) ItemChange(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsWebView::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsWebView_ItemChange(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QGraphicsWebView) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsWebView::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsWebView_ItemChangeDefault(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
}

//export callbackQGraphicsWebView_KeyPressEvent
func callbackQGraphicsWebView_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QGraphicsWebView) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QGraphicsWebView) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQGraphicsWebView_KeyReleaseEvent
func callbackQGraphicsWebView_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectKeyReleaseEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QGraphicsWebView) KeyReleaseEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QGraphicsWebView) KeyReleaseEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQGraphicsWebView_LinkClicked
func callbackQGraphicsWebView_LinkClicked(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::linkClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QGraphicsWebView) ConnectLinkClicked(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QGraphicsWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectLinkClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkClicked", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectLinkClicked() {
	defer qt.Recovering("disconnect QGraphicsWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkClicked")
	}
}

func (ptr *QGraphicsWebView) LinkClicked(url core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QGraphicsWebView) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::load")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQGraphicsWebView_LoadFinished
func callbackQGraphicsWebView_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QGraphicsWebView::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QGraphicsWebView) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QGraphicsWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QGraphicsWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QGraphicsWebView) LoadFinished(ok bool) {
	defer qt.Recovering("QGraphicsWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQGraphicsWebView_LoadProgress
func callbackQGraphicsWebView_LoadProgress(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QGraphicsWebView::loadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadProgress"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QGraphicsWebView) ConnectLoadProgress(f func(progress int)) {
	defer qt.Recovering("connect QGraphicsWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadProgress", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadProgress() {
	defer qt.Recovering("disconnect QGraphicsWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadProgress")
	}
}

func (ptr *QGraphicsWebView) LoadProgress(progress int) {
	defer qt.Recovering("QGraphicsWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadProgress(ptr.Pointer(), C.int(progress))
	}
}

//export callbackQGraphicsWebView_LoadStarted
func callbackQGraphicsWebView_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QGraphicsWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QGraphicsWebView) LoadStarted() {
	defer qt.Recovering("QGraphicsWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_MouseDoubleClickEvent
func callbackQGraphicsWebView_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectMouseDoubleClickEvent(f func(ev *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QGraphicsWebView) MouseDoubleClickEvent(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseDoubleClickEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

func (ptr *QGraphicsWebView) MouseDoubleClickEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MouseMoveEvent
func callbackQGraphicsWebView_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectMouseMoveEvent(f func(ev *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QGraphicsWebView) MouseMoveEvent(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

func (ptr *QGraphicsWebView) MouseMoveEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MousePressEvent
func callbackQGraphicsWebView_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectMousePressEvent(f func(ev *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QGraphicsWebView) MousePressEvent(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MousePressEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

func (ptr *QGraphicsWebView) MousePressEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

//export callbackQGraphicsWebView_MouseReleaseEvent
func callbackQGraphicsWebView_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectMouseReleaseEvent(f func(ev *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QGraphicsWebView) MouseReleaseEvent(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseReleaseEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

func (ptr *QGraphicsWebView) MouseReleaseEventDefault(ev widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(ev))
	}
}

func (ptr *QGraphicsWebView) Page() *QWebPage {
	defer qt.Recovering("QGraphicsWebView::page")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QGraphicsWebView_Page(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) PageAction(action QWebPage__WebAction) *widgets.QAction {
	defer qt.Recovering("QGraphicsWebView::pageAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QGraphicsWebView_PageAction(ptr.Pointer(), C.int(action)))
	}
	return nil
}

//export callbackQGraphicsWebView_Paint
func callbackQGraphicsWebView_Paint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsWebView) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsWebView::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsWebView::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

func (ptr *QGraphicsWebView) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsWebView::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWebView) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsWebView::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsWebView_Reload
func callbackQGraphicsWebView_Reload(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::reload")

	if signal := qt.GetSignal(C.GoString(ptrName), "reload"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectReload(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::reload")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reload", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectReload() {
	defer qt.Recovering("disconnect QGraphicsWebView::reload")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reload")
	}
}

func (ptr *QGraphicsWebView) Reload() {
	defer qt.Recovering("QGraphicsWebView::reload")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Reload(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) RenderHints() gui.QPainter__RenderHint {
	defer qt.Recovering("QGraphicsWebView::renderHints")

	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QGraphicsWebView_RenderHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsWebView_SceneEvent
func callbackQGraphicsWebView_SceneEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::sceneEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsWebView) ConnectSceneEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsWebView::sceneEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sceneEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectSceneEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::sceneEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sceneEvent")
	}
}

func (ptr *QGraphicsWebView) SceneEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) SceneEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) SetContent(data string, mimeType string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::setContent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetContent(ptr.Pointer(), C.CString(data), C.CString(mimeType), core.PointerFromQUrl(baseUrl))
	}
}

//export callbackQGraphicsWebView_SetGeometry
func callbackQGraphicsWebView_SetGeometry(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::setGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "setGeometry"); signal != nil {
		signal.(func(*core.QRectF))(core.NewQRectFFromPointer(rect))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).SetGeometryDefault(core.NewQRectFFromPointer(rect))
	}
}

func (ptr *QGraphicsWebView) ConnectSetGeometry(f func(rect *core.QRectF)) {
	defer qt.Recovering("connect QGraphicsWebView::setGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setGeometry", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectSetGeometry() {
	defer qt.Recovering("disconnect QGraphicsWebView::setGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setGeometry")
	}
}

func (ptr *QGraphicsWebView) SetGeometry(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsWebView::setGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsWebView) SetGeometryDefault(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsWebView::setGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetGeometryDefault(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsWebView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::setHtml")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetHtml(ptr.Pointer(), C.CString(html), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QGraphicsWebView) SetPage(page QWebPage_ITF) {
	defer qt.Recovering("QGraphicsWebView::setPage")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

func (ptr *QGraphicsWebView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	defer qt.Recovering("QGraphicsWebView::setRenderHint")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetRenderHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWebView) SetRenderHints(hints gui.QPainter__RenderHint) {
	defer qt.Recovering("QGraphicsWebView::setRenderHints")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_SetRenderHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QGraphicsWebView) Settings() *QWebSettings {
	defer qt.Recovering("QGraphicsWebView::settings")

	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QGraphicsWebView_Settings(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_SizeHint
func callbackQGraphicsWebView_SizeHint(ptr unsafe.Pointer, ptrName *C.char, which C.int, constraint unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSizeF(signal.(func(core.Qt__SizeHint, *core.QSizeF) *core.QSizeF)(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
	}

	return core.PointerFromQSizeF(NewQGraphicsWebViewFromPointer(ptr).SizeHintDefault(core.Qt__SizeHint(which), core.NewQSizeFFromPointer(constraint)))
}

func (ptr *QGraphicsWebView) ConnectSizeHint(f func(which core.Qt__SizeHint, constraint *core.QSizeF) *core.QSizeF) {
	defer qt.Recovering("connect QGraphicsWebView::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QGraphicsWebView::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QGraphicsWebView) SizeHint(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	defer qt.Recovering("QGraphicsWebView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFFromPointer(C.QGraphicsWebView_SizeHint(ptr.Pointer(), C.int(which), core.PointerFromQSizeF(constraint)))
	}
	return nil
}

func (ptr *QGraphicsWebView) SizeHintDefault(which core.Qt__SizeHint, constraint core.QSizeF_ITF) *core.QSizeF {
	defer qt.Recovering("QGraphicsWebView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFFromPointer(C.QGraphicsWebView_SizeHintDefault(ptr.Pointer(), C.int(which), core.PointerFromQSizeF(constraint)))
	}
	return nil
}

//export callbackQGraphicsWebView_StatusBarMessage
func callbackQGraphicsWebView_StatusBarMessage(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::statusBarMessage")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusBarMessage"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QGraphicsWebView) ConnectStatusBarMessage(f func(text string)) {
	defer qt.Recovering("connect QGraphicsWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectStatusBarMessage(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusBarMessage", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectStatusBarMessage() {
	defer qt.Recovering("disconnect QGraphicsWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusBarMessage")
	}
}

func (ptr *QGraphicsWebView) StatusBarMessage(text string) {
	defer qt.Recovering("QGraphicsWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_StatusBarMessage(ptr.Pointer(), C.CString(text))
	}
}

//export callbackQGraphicsWebView_Stop
func callbackQGraphicsWebView_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWebView) ConnectStop(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectStop() {
	defer qt.Recovering("disconnect QGraphicsWebView::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QGraphicsWebView) Stop() {
	defer qt.Recovering("QGraphicsWebView::stop")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Stop(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_TitleChanged
func callbackQGraphicsWebView_TitleChanged(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::titleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "titleChanged"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	}

}

func (ptr *QGraphicsWebView) ConnectTitleChanged(f func(vqs string)) {
	defer qt.Recovering("connect QGraphicsWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "titleChanged", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectTitleChanged() {
	defer qt.Recovering("disconnect QGraphicsWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "titleChanged")
	}
}

func (ptr *QGraphicsWebView) TitleChanged(vqs string) {
	defer qt.Recovering("QGraphicsWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TitleChanged(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QGraphicsWebView) TriggerPageAction(action QWebPage__WebAction, checked bool) {
	defer qt.Recovering("QGraphicsWebView::triggerPageAction")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TriggerPageAction(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

//export callbackQGraphicsWebView_UpdateGeometry
func callbackQGraphicsWebView_UpdateGeometry(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::updateGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectUpdateGeometry(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::updateGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometry", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectUpdateGeometry() {
	defer qt.Recovering("disconnect QGraphicsWebView::updateGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometry")
	}
}

func (ptr *QGraphicsWebView) UpdateGeometry() {
	defer qt.Recovering("QGraphicsWebView::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateGeometry(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) UpdateGeometryDefault() {
	defer qt.Recovering("QGraphicsWebView::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateGeometryDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_UrlChanged
func callbackQGraphicsWebView_UrlChanged(ptr unsafe.Pointer, ptrName *C.char, vqu unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::urlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(vqu))
	}

}

func (ptr *QGraphicsWebView) ConnectUrlChanged(f func(vqu *core.QUrl)) {
	defer qt.Recovering("connect QGraphicsWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlChanged", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectUrlChanged() {
	defer qt.Recovering("disconnect QGraphicsWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlChanged")
	}
}

func (ptr *QGraphicsWebView) UrlChanged(vqu core.QUrl_ITF) {
	defer qt.Recovering("QGraphicsWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(vqu))
	}
}

//export callbackQGraphicsWebView_WheelEvent
func callbackQGraphicsWebView_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneWheelEvent))(widgets.NewQGraphicsSceneWheelEventFromPointer(ev))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(ev))
	}
}

func (ptr *QGraphicsWebView) ConnectWheelEvent(f func(ev *widgets.QGraphicsSceneWheelEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QGraphicsWebView) WheelEvent(ev widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_WheelEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(ev))
	}
}

func (ptr *QGraphicsWebView) WheelEventDefault(ev widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(ev))
	}
}

func (ptr *QGraphicsWebView) DestroyQGraphicsWebView() {
	defer qt.Recovering("QGraphicsWebView::~QGraphicsWebView")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DestroyQGraphicsWebView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsWebView_BoundingRect
func callbackQGraphicsWebView_BoundingRect(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::boundingRect")

	if signal := qt.GetSignal(C.GoString(ptrName), "boundingRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQGraphicsWebViewFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsWebView) ConnectBoundingRect(f func() *core.QRectF) {
	defer qt.Recovering("connect QGraphicsWebView::boundingRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "boundingRect", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectBoundingRect() {
	defer qt.Recovering("disconnect QGraphicsWebView::boundingRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "boundingRect")
	}
}

func (ptr *QGraphicsWebView) BoundingRect() *core.QRectF {
	defer qt.Recovering("QGraphicsWebView::boundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QGraphicsWebView_BoundingRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) BoundingRectDefault() *core.QRectF {
	defer qt.Recovering("QGraphicsWebView::boundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QGraphicsWebView_BoundingRectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_ChangeEvent
func callbackQGraphicsWebView_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *QGraphicsWebView) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_Close
func callbackQGraphicsWebView_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QGraphicsWebView::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CloseDefault()))
}

func (ptr *QGraphicsWebView) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QGraphicsWebView::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectClose() {
	defer qt.Recovering("disconnect QGraphicsWebView::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QGraphicsWebView) Close() bool {
	defer qt.Recovering("QGraphicsWebView::close")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) CloseDefault() bool {
	defer qt.Recovering("QGraphicsWebView::close")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQGraphicsWebView_CloseEvent
func callbackQGraphicsWebView_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *QGraphicsWebView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsWebView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQGraphicsWebView_GetContentsMargins
func callbackQGraphicsWebView_GetContentsMargins(ptr unsafe.Pointer, ptrName *C.char, left C.double, top C.double, right C.double, bottom C.double) {
	defer qt.Recovering("callback QGraphicsWebView::getContentsMargins")

	if signal := qt.GetSignal(C.GoString(ptrName), "getContentsMargins"); signal != nil {
		signal.(func(float64, float64, float64, float64))(float64(left), float64(top), float64(right), float64(bottom))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GetContentsMarginsDefault(float64(left), float64(top), float64(right), float64(bottom))
	}
}

func (ptr *QGraphicsWebView) ConnectGetContentsMargins(f func(left float64, top float64, right float64, bottom float64)) {
	defer qt.Recovering("connect QGraphicsWebView::getContentsMargins")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "getContentsMargins", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectGetContentsMargins() {
	defer qt.Recovering("disconnect QGraphicsWebView::getContentsMargins")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "getContentsMargins")
	}
}

func (ptr *QGraphicsWebView) GetContentsMargins(left float64, top float64, right float64, bottom float64) {
	defer qt.Recovering("QGraphicsWebView::getContentsMargins")

	if ptr.Pointer() != nil {

	}
}

func (ptr *QGraphicsWebView) GetContentsMarginsDefault(left float64, top float64, right float64, bottom float64) {
	defer qt.Recovering("QGraphicsWebView::getContentsMargins")

	if ptr.Pointer() != nil {

	}
}

//export callbackQGraphicsWebView_GrabKeyboardEvent
func callbackQGraphicsWebView_GrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::grabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectGrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectGrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabKeyboardEvent")
	}
}

func (ptr *QGraphicsWebView) GrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_GrabMouseEvent
func callbackQGraphicsWebView_GrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::grabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectGrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabMouseEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectGrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabMouseEvent")
	}
}

func (ptr *QGraphicsWebView) GrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) GrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_HideEvent
func callbackQGraphicsWebView_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *QGraphicsWebView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsWebView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQGraphicsWebView_InitStyleOption
func callbackQGraphicsWebView_InitStyleOption(ptr unsafe.Pointer, ptrName *C.char, option unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*widgets.QStyleOption))(widgets.NewQStyleOptionFromPointer(option))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).InitStyleOptionDefault(widgets.NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QGraphicsWebView) ConnectInitStyleOption(f func(option *widgets.QStyleOption)) {
	defer qt.Recovering("connect QGraphicsWebView::initStyleOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initStyleOption", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectInitStyleOption() {
	defer qt.Recovering("disconnect QGraphicsWebView::initStyleOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initStyleOption")
	}
}

func (ptr *QGraphicsWebView) InitStyleOption(option widgets.QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsWebView::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InitStyleOption(ptr.Pointer(), widgets.PointerFromQStyleOption(option))
	}
}

func (ptr *QGraphicsWebView) InitStyleOptionDefault(option widgets.QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsWebView::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_InitStyleOptionDefault(ptr.Pointer(), widgets.PointerFromQStyleOption(option))
	}
}

//export callbackQGraphicsWebView_MoveEvent
func callbackQGraphicsWebView_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMoveEvent))(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).MoveEventDefault(widgets.NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectMoveEvent(f func(event *widgets.QGraphicsSceneMoveEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *QGraphicsWebView) MoveEvent(event widgets.QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMoveEvent(event))
	}
}

func (ptr *QGraphicsWebView) MoveEventDefault(event widgets.QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_MoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMoveEvent(event))
	}
}

//export callbackQGraphicsWebView_PaintWindowFrame
func callbackQGraphicsWebView_PaintWindowFrame(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::paintWindowFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsWebView) ConnectPaintWindowFrame(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsWebView::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintWindowFrame", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectPaintWindowFrame() {
	defer qt.Recovering("disconnect QGraphicsWebView::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintWindowFrame")
	}
}

func (ptr *QGraphicsWebView) PaintWindowFrame(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsWebView::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PaintWindowFrame(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWebView) PaintWindowFrameDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsWebView::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

//export callbackQGraphicsWebView_PolishEvent
func callbackQGraphicsWebView_PolishEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::polishEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "polishEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectPolishEvent(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::polishEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polishEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectPolishEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::polishEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polishEvent")
	}
}

func (ptr *QGraphicsWebView) PolishEvent() {
	defer qt.Recovering("QGraphicsWebView::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PolishEvent(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) PolishEventDefault() {
	defer qt.Recovering("QGraphicsWebView::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_PolishEventDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_ResizeEvent
func callbackQGraphicsWebView_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneResizeEvent))(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ResizeEventDefault(widgets.NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectResizeEvent(f func(event *widgets.QGraphicsSceneResizeEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *QGraphicsWebView) ResizeEvent(event widgets.QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ResizeEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsWebView) ResizeEventDefault(event widgets.QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ResizeEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneResizeEvent(event))
	}
}

//export callbackQGraphicsWebView_Shape
func callbackQGraphicsWebView_Shape(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::shape")

	if signal := qt.GetSignal(C.GoString(ptrName), "shape"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsWebViewFromPointer(ptr).ShapeDefault())
}

func (ptr *QGraphicsWebView) ConnectShape(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsWebView::shape")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "shape", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectShape() {
	defer qt.Recovering("disconnect QGraphicsWebView::shape")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "shape")
	}
}

func (ptr *QGraphicsWebView) Shape() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsWebView::shape")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsWebView_Shape(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) ShapeDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsWebView::shape")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsWebView_ShapeDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_ShowEvent
func callbackQGraphicsWebView_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *QGraphicsWebView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsWebView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQGraphicsWebView_Type
func callbackQGraphicsWebView_Type(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QGraphicsWebView::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(NewQGraphicsWebViewFromPointer(ptr).TypeDefault())
}

func (ptr *QGraphicsWebView) ConnectType(f func() int) {
	defer qt.Recovering("connect QGraphicsWebView::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "type", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectType() {
	defer qt.Recovering("disconnect QGraphicsWebView::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "type")
	}
}

func (ptr *QGraphicsWebView) Type() int {
	defer qt.Recovering("QGraphicsWebView::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWebView_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWebView) TypeDefault() int {
	defer qt.Recovering("QGraphicsWebView::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWebView_TypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQGraphicsWebView_UngrabKeyboardEvent
func callbackQGraphicsWebView_UngrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::ungrabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectUngrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectUngrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent")
	}
}

func (ptr *QGraphicsWebView) UngrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_UngrabMouseEvent
func callbackQGraphicsWebView_UngrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::ungrabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectUngrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabMouseEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectUngrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabMouseEvent")
	}
}

func (ptr *QGraphicsWebView) UngrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) UngrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_WindowFrameEvent
func callbackQGraphicsWebView_WindowFrameEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::windowFrameEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowFrameEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).WindowFrameEventDefault(core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsWebView) ConnectWindowFrameEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsWebView::windowFrameEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "windowFrameEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectWindowFrameEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::windowFrameEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "windowFrameEvent")
	}
}

func (ptr *QGraphicsWebView) WindowFrameEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::windowFrameEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_WindowFrameEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) WindowFrameEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::windowFrameEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_WindowFrameEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_WindowFrameSectionAt
func callbackQGraphicsWebView_WindowFrameSectionAt(ptr unsafe.Pointer, ptrName *C.char, pos unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::windowFrameSectionAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowFrameSectionAt"); signal != nil {
		return C.int(signal.(func(*core.QPointF) core.Qt__WindowFrameSection)(core.NewQPointFFromPointer(pos)))
	}

	return C.int(NewQGraphicsWebViewFromPointer(ptr).WindowFrameSectionAtDefault(core.NewQPointFFromPointer(pos)))
}

func (ptr *QGraphicsWebView) ConnectWindowFrameSectionAt(f func(pos *core.QPointF) core.Qt__WindowFrameSection) {
	defer qt.Recovering("connect QGraphicsWebView::windowFrameSectionAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "windowFrameSectionAt", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectWindowFrameSectionAt() {
	defer qt.Recovering("disconnect QGraphicsWebView::windowFrameSectionAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "windowFrameSectionAt")
	}
}

func (ptr *QGraphicsWebView) WindowFrameSectionAt(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	defer qt.Recovering("QGraphicsWebView::windowFrameSectionAt")

	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QGraphicsWebView_WindowFrameSectionAt(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
}

func (ptr *QGraphicsWebView) WindowFrameSectionAtDefault(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	defer qt.Recovering("QGraphicsWebView::windowFrameSectionAt")

	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QGraphicsWebView_WindowFrameSectionAtDefault(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
}

//export callbackQGraphicsWebView_UpdateMicroFocus
func callbackQGraphicsWebView_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QGraphicsWebView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QGraphicsWebView) UpdateMicroFocus() {
	defer qt.Recovering("QGraphicsWebView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QGraphicsWebView) UpdateMicroFocusDefault() {
	defer qt.Recovering("QGraphicsWebView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsWebView_Advance
func callbackQGraphicsWebView_Advance(ptr unsafe.Pointer, ptrName *C.char, phase C.int) {
	defer qt.Recovering("callback QGraphicsWebView::advance")

	if signal := qt.GetSignal(C.GoString(ptrName), "advance"); signal != nil {
		signal.(func(int))(int(phase))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).AdvanceDefault(int(phase))
	}
}

func (ptr *QGraphicsWebView) ConnectAdvance(f func(phase int)) {
	defer qt.Recovering("connect QGraphicsWebView::advance")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "advance", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectAdvance() {
	defer qt.Recovering("disconnect QGraphicsWebView::advance")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "advance")
	}
}

func (ptr *QGraphicsWebView) Advance(phase int) {
	defer qt.Recovering("QGraphicsWebView::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_Advance(ptr.Pointer(), C.int(phase))
	}
}

func (ptr *QGraphicsWebView) AdvanceDefault(phase int) {
	defer qt.Recovering("QGraphicsWebView::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_AdvanceDefault(ptr.Pointer(), C.int(phase))
	}
}

//export callbackQGraphicsWebView_CollidesWithItem
func callbackQGraphicsWebView_CollidesWithItem(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer, mode C.int) C.int {
	defer qt.Recovering("callback QGraphicsWebView::collidesWithItem")

	if signal := qt.GetSignal(C.GoString(ptrName), "collidesWithItem"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode))))
}

func (ptr *QGraphicsWebView) ConnectCollidesWithItem(f func(other *widgets.QGraphicsItem, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsWebView::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "collidesWithItem", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectCollidesWithItem() {
	defer qt.Recovering("disconnect QGraphicsWebView::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "collidesWithItem")
	}
}

func (ptr *QGraphicsWebView) CollidesWithItem(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsWebView::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithItem(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.int(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsWebView::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.int(mode)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_CollidesWithPath
func callbackQGraphicsWebView_CollidesWithPath(ptr unsafe.Pointer, ptrName *C.char, path unsafe.Pointer, mode C.int) C.int {
	defer qt.Recovering("callback QGraphicsWebView::collidesWithPath")

	if signal := qt.GetSignal(C.GoString(ptrName), "collidesWithPath"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode))))
}

func (ptr *QGraphicsWebView) ConnectCollidesWithPath(f func(path *gui.QPainterPath, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsWebView::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "collidesWithPath", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectCollidesWithPath() {
	defer qt.Recovering("disconnect QGraphicsWebView::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "collidesWithPath")
	}
}

func (ptr *QGraphicsWebView) CollidesWithPath(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsWebView::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithPath(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsWebView::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_Contains
func callbackQGraphicsWebView_Contains(ptr unsafe.Pointer, ptrName *C.char, point unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::contains")

	if signal := qt.GetSignal(C.GoString(ptrName), "contains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point))))
}

func (ptr *QGraphicsWebView) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QGraphicsWebView::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contains", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectContains() {
	defer qt.Recovering("disconnect QGraphicsWebView::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contains")
	}
}

func (ptr *QGraphicsWebView) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_HoverEnterEvent
func callbackQGraphicsWebView_HoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectHoverEnterEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

func (ptr *QGraphicsWebView) HoverEnterEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsWebView) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsWebView_IsObscuredBy
func callbackQGraphicsWebView_IsObscuredBy(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::isObscuredBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "isObscuredBy"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem) bool)(widgets.NewQGraphicsItemFromPointer(item))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item))))
}

func (ptr *QGraphicsWebView) ConnectIsObscuredBy(f func(item *widgets.QGraphicsItem) bool) {
	defer qt.Recovering("connect QGraphicsWebView::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isObscuredBy", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectIsObscuredBy() {
	defer qt.Recovering("disconnect QGraphicsWebView::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isObscuredBy")
	}
}

func (ptr *QGraphicsWebView) IsObscuredBy(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsObscuredBy(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_OpaqueArea
func callbackQGraphicsWebView_OpaqueArea(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::opaqueArea")

	if signal := qt.GetSignal(C.GoString(ptrName), "opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsWebViewFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QGraphicsWebView) ConnectOpaqueArea(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsWebView::opaqueArea")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "opaqueArea", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectOpaqueArea() {
	defer qt.Recovering("disconnect QGraphicsWebView::opaqueArea")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "opaqueArea")
	}
}

func (ptr *QGraphicsWebView) OpaqueArea() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsWebView::opaqueArea")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsWebView_OpaqueArea(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) OpaqueAreaDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsWebView::opaqueArea")

	if ptr.Pointer() != nil {
		return gui.NewQPainterPathFromPointer(C.QGraphicsWebView_OpaqueAreaDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsWebView_SceneEventFilter
func callbackQGraphicsWebView_SceneEventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::sceneEventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneEventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, *core.QEvent) bool)(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsWebView) ConnectSceneEventFilter(f func(watched *widgets.QGraphicsItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsWebView::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sceneEventFilter", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectSceneEventFilter() {
	defer qt.Recovering("disconnect QGraphicsWebView::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sceneEventFilter")
	}
}

func (ptr *QGraphicsWebView) SceneEventFilter(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEventFilter(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_TimerEvent
func callbackQGraphicsWebView_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QGraphicsWebView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsWebView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsWebView_ChildEvent
func callbackQGraphicsWebView_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QGraphicsWebView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsWebView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsWebView_ConnectNotify
func callbackQGraphicsWebView_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsWebView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsWebView::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QGraphicsWebView::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QGraphicsWebView) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsWebView::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsWebView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsWebView::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsWebView_CustomEvent
func callbackQGraphicsWebView_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWebView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWebView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsWebView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QGraphicsWebView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWebView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWebView::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsWebView_DeleteLater
func callbackQGraphicsWebView_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWebView::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsWebView) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QGraphicsWebView::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QGraphicsWebView::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QGraphicsWebView) DeleteLater() {
	defer qt.Recovering("QGraphicsWebView::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsWebView) DeleteLaterDefault() {
	defer qt.Recovering("QGraphicsWebView::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsWebView_DisconnectNotify
func callbackQGraphicsWebView_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWebView::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsWebViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsWebView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsWebView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QGraphicsWebView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QGraphicsWebView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsWebView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsWebView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsWebView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsWebView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsWebView_EventFilter
func callbackQGraphicsWebView_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsWebView::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQGraphicsWebViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QGraphicsWebView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsWebView::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QGraphicsWebView::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QGraphicsWebView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWebView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWebView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsWebView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsWebView_MetaObject
func callbackQGraphicsWebView_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsWebView::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGraphicsWebViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGraphicsWebView) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QGraphicsWebView::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QGraphicsWebView) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QGraphicsWebView::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QGraphicsWebView) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QGraphicsWebView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsWebView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWebView) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QGraphicsWebView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsWebView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebDatabase struct {
	ptr unsafe.Pointer
}

type QWebDatabase_ITF interface {
	QWebDatabase_PTR() *QWebDatabase
}

func (p *QWebDatabase) QWebDatabase_PTR() *QWebDatabase {
	return p
}

func (p *QWebDatabase) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebDatabase) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebDatabase(ptr QWebDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebDatabase_PTR().Pointer()
	}
	return nil
}

func NewQWebDatabaseFromPointer(ptr unsafe.Pointer) *QWebDatabase {
	var n = new(QWebDatabase)
	n.SetPointer(ptr)
	return n
}

func newQWebDatabaseFromPointer(ptr unsafe.Pointer) *QWebDatabase {
	var n = NewQWebDatabaseFromPointer(ptr)
	return n
}

func NewQWebDatabase(other QWebDatabase_ITF) *QWebDatabase {
	defer qt.Recovering("QWebDatabase::QWebDatabase")

	return newQWebDatabaseFromPointer(C.QWebDatabase_NewQWebDatabase(PointerFromQWebDatabase(other)))
}

func (ptr *QWebDatabase) DisplayName() string {
	defer qt.Recovering("QWebDatabase::displayName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebDatabase_DisplayName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) ExpectedSize() int64 {
	defer qt.Recovering("QWebDatabase::expectedSize")

	if ptr.Pointer() != nil {
		return int64(C.QWebDatabase_ExpectedSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebDatabase) FileName() string {
	defer qt.Recovering("QWebDatabase::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebDatabase_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) Name() string {
	defer qt.Recovering("QWebDatabase::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebDatabase_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebDatabase) Origin() *QWebSecurityOrigin {
	defer qt.Recovering("QWebDatabase::origin")

	if ptr.Pointer() != nil {
		return NewQWebSecurityOriginFromPointer(C.QWebDatabase_Origin(ptr.Pointer()))
	}
	return nil
}

func QWebDatabase_RemoveAllDatabases() {
	defer qt.Recovering("QWebDatabase::removeAllDatabases")

	C.QWebDatabase_QWebDatabase_RemoveAllDatabases()
}

func (ptr *QWebDatabase) RemoveAllDatabases() {
	defer qt.Recovering("QWebDatabase::removeAllDatabases")

	C.QWebDatabase_QWebDatabase_RemoveAllDatabases()
}

func QWebDatabase_RemoveDatabase(db QWebDatabase_ITF) {
	defer qt.Recovering("QWebDatabase::removeDatabase")

	C.QWebDatabase_QWebDatabase_RemoveDatabase(PointerFromQWebDatabase(db))
}

func (ptr *QWebDatabase) RemoveDatabase(db QWebDatabase_ITF) {
	defer qt.Recovering("QWebDatabase::removeDatabase")

	C.QWebDatabase_QWebDatabase_RemoveDatabase(PointerFromQWebDatabase(db))
}

func (ptr *QWebDatabase) Size() int64 {
	defer qt.Recovering("QWebDatabase::size")

	if ptr.Pointer() != nil {
		return int64(C.QWebDatabase_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebDatabase) DestroyQWebDatabase() {
	defer qt.Recovering("QWebDatabase::~QWebDatabase")

	if ptr.Pointer() != nil {
		C.QWebDatabase_DestroyQWebDatabase(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebElement::StyleResolveStrategy
type QWebElement__StyleResolveStrategy int64

var (
	QWebElement__InlineStyle   = QWebElement__StyleResolveStrategy(0)
	QWebElement__CascadedStyle = QWebElement__StyleResolveStrategy(1)
	QWebElement__ComputedStyle = QWebElement__StyleResolveStrategy(2)
)

type QWebElement struct {
	ptr unsafe.Pointer
}

type QWebElement_ITF interface {
	QWebElement_PTR() *QWebElement
}

func (p *QWebElement) QWebElement_PTR() *QWebElement {
	return p
}

func (p *QWebElement) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebElement) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebElement(ptr QWebElement_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebElement_PTR().Pointer()
	}
	return nil
}

func NewQWebElementFromPointer(ptr unsafe.Pointer) *QWebElement {
	var n = new(QWebElement)
	n.SetPointer(ptr)
	return n
}

func newQWebElementFromPointer(ptr unsafe.Pointer) *QWebElement {
	var n = NewQWebElementFromPointer(ptr)
	return n
}

func NewQWebElement() *QWebElement {
	defer qt.Recovering("QWebElement::QWebElement")

	return newQWebElementFromPointer(C.QWebElement_NewQWebElement())
}

func NewQWebElement2(other QWebElement_ITF) *QWebElement {
	defer qt.Recovering("QWebElement::QWebElement")

	return newQWebElementFromPointer(C.QWebElement_NewQWebElement2(PointerFromQWebElement(other)))
}

func (ptr *QWebElement) AddClass(name string) {
	defer qt.Recovering("QWebElement::addClass")

	if ptr.Pointer() != nil {
		C.QWebElement_AddClass(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWebElement) AppendInside(markup string) {
	defer qt.Recovering("QWebElement::appendInside")

	if ptr.Pointer() != nil {
		C.QWebElement_AppendInside(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) AppendInside2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::appendInside")

	if ptr.Pointer() != nil {
		C.QWebElement_AppendInside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) AppendOutside(markup string) {
	defer qt.Recovering("QWebElement::appendOutside")

	if ptr.Pointer() != nil {
		C.QWebElement_AppendOutside(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) AppendOutside2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::appendOutside")

	if ptr.Pointer() != nil {
		C.QWebElement_AppendOutside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) Attribute(name string, defaultValue string) string {
	defer qt.Recovering("QWebElement::attribute")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_Attribute(ptr.Pointer(), C.CString(name), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QWebElement) AttributeNS(namespaceUri string, name string, defaultValue string) string {
	defer qt.Recovering("QWebElement::attributeNS")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_AttributeNS(ptr.Pointer(), C.CString(namespaceUri), C.CString(name), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QWebElement) AttributeNames(namespaceUri string) []string {
	defer qt.Recovering("QWebElement::attributeNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QWebElement_AttributeNames(ptr.Pointer(), C.CString(namespaceUri))), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebElement) Classes() []string {
	defer qt.Recovering("QWebElement::classes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QWebElement_Classes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebElement) Clone() *QWebElement {
	defer qt.Recovering("QWebElement::clone")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) Document() *QWebElement {
	defer qt.Recovering("QWebElement::document")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) EncloseContentsWith2(markup string) {
	defer qt.Recovering("QWebElement::encloseContentsWith")

	if ptr.Pointer() != nil {
		C.QWebElement_EncloseContentsWith2(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) EncloseContentsWith(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::encloseContentsWith")

	if ptr.Pointer() != nil {
		C.QWebElement_EncloseContentsWith(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) EncloseWith(markup string) {
	defer qt.Recovering("QWebElement::encloseWith")

	if ptr.Pointer() != nil {
		C.QWebElement_EncloseWith(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) EncloseWith2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::encloseWith")

	if ptr.Pointer() != nil {
		C.QWebElement_EncloseWith2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) EvaluateJavaScript(scriptSource string) *core.QVariant {
	defer qt.Recovering("QWebElement::evaluateJavaScript")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebElement_EvaluateJavaScript(ptr.Pointer(), C.CString(scriptSource)))
	}
	return nil
}

func (ptr *QWebElement) FindAll(selectorQuery string) *QWebElementCollection {
	defer qt.Recovering("QWebElement::findAll")

	if ptr.Pointer() != nil {
		return NewQWebElementCollectionFromPointer(C.QWebElement_FindAll(ptr.Pointer(), C.CString(selectorQuery)))
	}
	return nil
}

func (ptr *QWebElement) FindFirst(selectorQuery string) *QWebElement {
	defer qt.Recovering("QWebElement::findFirst")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_FindFirst(ptr.Pointer(), C.CString(selectorQuery)))
	}
	return nil
}

func (ptr *QWebElement) FirstChild() *QWebElement {
	defer qt.Recovering("QWebElement::firstChild")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) Geometry() *core.QRect {
	defer qt.Recovering("QWebElement::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWebElement_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) HasAttribute(name string) bool {
	defer qt.Recovering("QWebElement::hasAttribute")

	if ptr.Pointer() != nil {
		return C.QWebElement_HasAttribute(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QWebElement) HasAttributeNS(namespaceUri string, name string) bool {
	defer qt.Recovering("QWebElement::hasAttributeNS")

	if ptr.Pointer() != nil {
		return C.QWebElement_HasAttributeNS(ptr.Pointer(), C.CString(namespaceUri), C.CString(name)) != 0
	}
	return false
}

func (ptr *QWebElement) HasAttributes() bool {
	defer qt.Recovering("QWebElement::hasAttributes")

	if ptr.Pointer() != nil {
		return C.QWebElement_HasAttributes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebElement) HasClass(name string) bool {
	defer qt.Recovering("QWebElement::hasClass")

	if ptr.Pointer() != nil {
		return C.QWebElement_HasClass(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QWebElement) HasFocus() bool {
	defer qt.Recovering("QWebElement::hasFocus")

	if ptr.Pointer() != nil {
		return C.QWebElement_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebElement) IsNull() bool {
	defer qt.Recovering("QWebElement::isNull")

	if ptr.Pointer() != nil {
		return C.QWebElement_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebElement) LastChild() *QWebElement {
	defer qt.Recovering("QWebElement::lastChild")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) LocalName() string {
	defer qt.Recovering("QWebElement::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) NamespaceUri() string {
	defer qt.Recovering("QWebElement::namespaceUri")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_NamespaceUri(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) NextSibling() *QWebElement {
	defer qt.Recovering("QWebElement::nextSibling")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_NextSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) Parent() *QWebElement {
	defer qt.Recovering("QWebElement::parent")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) Prefix() string {
	defer qt.Recovering("QWebElement::prefix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_Prefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) PrependInside(markup string) {
	defer qt.Recovering("QWebElement::prependInside")

	if ptr.Pointer() != nil {
		C.QWebElement_PrependInside(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) PrependInside2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::prependInside")

	if ptr.Pointer() != nil {
		C.QWebElement_PrependInside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) PrependOutside(markup string) {
	defer qt.Recovering("QWebElement::prependOutside")

	if ptr.Pointer() != nil {
		C.QWebElement_PrependOutside(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) PrependOutside2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::prependOutside")

	if ptr.Pointer() != nil {
		C.QWebElement_PrependOutside2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) PreviousSibling() *QWebElement {
	defer qt.Recovering("QWebElement::previousSibling")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_PreviousSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) RemoveAllChildren() {
	defer qt.Recovering("QWebElement::removeAllChildren")

	if ptr.Pointer() != nil {
		C.QWebElement_RemoveAllChildren(ptr.Pointer())
	}
}

func (ptr *QWebElement) RemoveAttribute(name string) {
	defer qt.Recovering("QWebElement::removeAttribute")

	if ptr.Pointer() != nil {
		C.QWebElement_RemoveAttribute(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWebElement) RemoveAttributeNS(namespaceUri string, name string) {
	defer qt.Recovering("QWebElement::removeAttributeNS")

	if ptr.Pointer() != nil {
		C.QWebElement_RemoveAttributeNS(ptr.Pointer(), C.CString(namespaceUri), C.CString(name))
	}
}

func (ptr *QWebElement) RemoveClass(name string) {
	defer qt.Recovering("QWebElement::removeClass")

	if ptr.Pointer() != nil {
		C.QWebElement_RemoveClass(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWebElement) RemoveFromDocument() {
	defer qt.Recovering("QWebElement::removeFromDocument")

	if ptr.Pointer() != nil {
		C.QWebElement_RemoveFromDocument(ptr.Pointer())
	}
}

func (ptr *QWebElement) Render(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWebElement::render")

	if ptr.Pointer() != nil {
		C.QWebElement_Render(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWebElement) Render2(painter gui.QPainter_ITF, clip core.QRect_ITF) {
	defer qt.Recovering("QWebElement::render")

	if ptr.Pointer() != nil {
		C.QWebElement_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(clip))
	}
}

func (ptr *QWebElement) Replace(markup string) {
	defer qt.Recovering("QWebElement::replace")

	if ptr.Pointer() != nil {
		C.QWebElement_Replace(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) Replace2(element QWebElement_ITF) {
	defer qt.Recovering("QWebElement::replace")

	if ptr.Pointer() != nil {
		C.QWebElement_Replace2(ptr.Pointer(), PointerFromQWebElement(element))
	}
}

func (ptr *QWebElement) SetAttribute(name string, value string) {
	defer qt.Recovering("QWebElement::setAttribute")

	if ptr.Pointer() != nil {
		C.QWebElement_SetAttribute(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QWebElement) SetAttributeNS(namespaceUri string, name string, value string) {
	defer qt.Recovering("QWebElement::setAttributeNS")

	if ptr.Pointer() != nil {
		C.QWebElement_SetAttributeNS(ptr.Pointer(), C.CString(namespaceUri), C.CString(name), C.CString(value))
	}
}

func (ptr *QWebElement) SetFocus() {
	defer qt.Recovering("QWebElement::setFocus")

	if ptr.Pointer() != nil {
		C.QWebElement_SetFocus(ptr.Pointer())
	}
}

func (ptr *QWebElement) SetInnerXml(markup string) {
	defer qt.Recovering("QWebElement::setInnerXml")

	if ptr.Pointer() != nil {
		C.QWebElement_SetInnerXml(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) SetOuterXml(markup string) {
	defer qt.Recovering("QWebElement::setOuterXml")

	if ptr.Pointer() != nil {
		C.QWebElement_SetOuterXml(ptr.Pointer(), C.CString(markup))
	}
}

func (ptr *QWebElement) SetPlainText(text string) {
	defer qt.Recovering("QWebElement::setPlainText")

	if ptr.Pointer() != nil {
		C.QWebElement_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QWebElement) SetStyleProperty(name string, value string) {
	defer qt.Recovering("QWebElement::setStyleProperty")

	if ptr.Pointer() != nil {
		C.QWebElement_SetStyleProperty(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QWebElement) StyleProperty(name string, strategy QWebElement__StyleResolveStrategy) string {
	defer qt.Recovering("QWebElement::styleProperty")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_StyleProperty(ptr.Pointer(), C.CString(name), C.int(strategy)))
	}
	return ""
}

func (ptr *QWebElement) TagName() string {
	defer qt.Recovering("QWebElement::tagName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_TagName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) TakeFromDocument() *QWebElement {
	defer qt.Recovering("QWebElement::takeFromDocument")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElement_TakeFromDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) ToInnerXml() string {
	defer qt.Recovering("QWebElement::toInnerXml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_ToInnerXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToOuterXml() string {
	defer qt.Recovering("QWebElement::toOuterXml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_ToOuterXml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToPlainText() string {
	defer qt.Recovering("QWebElement::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebElement_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebElement) ToggleClass(name string) {
	defer qt.Recovering("QWebElement::toggleClass")

	if ptr.Pointer() != nil {
		C.QWebElement_ToggleClass(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWebElement) WebFrame() *QWebFrame {
	defer qt.Recovering("QWebElement::webFrame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebElement_WebFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElement) DestroyQWebElement() {
	defer qt.Recovering("QWebElement::~QWebElement")

	if ptr.Pointer() != nil {
		C.QWebElement_DestroyQWebElement(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebElementCollection struct {
	ptr unsafe.Pointer
}

type QWebElementCollection_ITF interface {
	QWebElementCollection_PTR() *QWebElementCollection
}

func (p *QWebElementCollection) QWebElementCollection_PTR() *QWebElementCollection {
	return p
}

func (p *QWebElementCollection) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebElementCollection) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebElementCollection(ptr QWebElementCollection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebElementCollection_PTR().Pointer()
	}
	return nil
}

func NewQWebElementCollectionFromPointer(ptr unsafe.Pointer) *QWebElementCollection {
	var n = new(QWebElementCollection)
	n.SetPointer(ptr)
	return n
}

func newQWebElementCollectionFromPointer(ptr unsafe.Pointer) *QWebElementCollection {
	var n = NewQWebElementCollectionFromPointer(ptr)
	return n
}

func NewQWebElementCollection() *QWebElementCollection {
	defer qt.Recovering("QWebElementCollection::QWebElementCollection")

	return newQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection())
}

func NewQWebElementCollection2(contextElement QWebElement_ITF, query string) *QWebElementCollection {
	defer qt.Recovering("QWebElementCollection::QWebElementCollection")

	return newQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection2(PointerFromQWebElement(contextElement), C.CString(query)))
}

func NewQWebElementCollection3(other QWebElementCollection_ITF) *QWebElementCollection {
	defer qt.Recovering("QWebElementCollection::QWebElementCollection")

	return newQWebElementCollectionFromPointer(C.QWebElementCollection_NewQWebElementCollection3(PointerFromQWebElementCollection(other)))
}

func (ptr *QWebElementCollection) Append(other QWebElementCollection_ITF) {
	defer qt.Recovering("QWebElementCollection::append")

	if ptr.Pointer() != nil {
		C.QWebElementCollection_Append(ptr.Pointer(), PointerFromQWebElementCollection(other))
	}
}

func (ptr *QWebElementCollection) At(i int) *QWebElement {
	defer qt.Recovering("QWebElementCollection::at")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElementCollection_At(ptr.Pointer(), C.int(i)))
	}
	return nil
}

func (ptr *QWebElementCollection) Count() int {
	defer qt.Recovering("QWebElementCollection::count")

	if ptr.Pointer() != nil {
		return int(C.QWebElementCollection_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebElementCollection) First() *QWebElement {
	defer qt.Recovering("QWebElementCollection::first")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElementCollection_First(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElementCollection) Last() *QWebElement {
	defer qt.Recovering("QWebElementCollection::last")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebElementCollection_Last(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebElementCollection) DestroyQWebElementCollection() {
	defer qt.Recovering("QWebElementCollection::~QWebElementCollection")

	if ptr.Pointer() != nil {
		C.QWebElementCollection_DestroyQWebElementCollection(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebFrame::RenderLayer
type QWebFrame__RenderLayer int64

const (
	QWebFrame__ContentsLayer  = QWebFrame__RenderLayer(0x10)
	QWebFrame__ScrollBarLayer = QWebFrame__RenderLayer(0x20)
	QWebFrame__PanIconLayer   = QWebFrame__RenderLayer(0x40)
	QWebFrame__AllLayers      = QWebFrame__RenderLayer(0xff)
)

//QWebFrame::ValueOwnership
type QWebFrame__ValueOwnership int64

const (
	QWebFrame__QtOwnership     = QWebFrame__ValueOwnership(0)
	QWebFrame__ScriptOwnership = QWebFrame__ValueOwnership(1)
	QWebFrame__AutoOwnership   = QWebFrame__ValueOwnership(2)
)

type QWebFrame struct {
	core.QObject
}

type QWebFrame_ITF interface {
	core.QObject_ITF
	QWebFrame_PTR() *QWebFrame
}

func (p *QWebFrame) QWebFrame_PTR() *QWebFrame {
	return p
}

func (p *QWebFrame) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebFrame) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebFrame(ptr QWebFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebFrame_PTR().Pointer()
	}
	return nil
}

func NewQWebFrameFromPointer(ptr unsafe.Pointer) *QWebFrame {
	var n = new(QWebFrame)
	n.SetPointer(ptr)
	return n
}

func newQWebFrameFromPointer(ptr unsafe.Pointer) *QWebFrame {
	var n = NewQWebFrameFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebFrame_") {
		n.SetObjectName("QWebFrame_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebFrame) AddToJavaScriptWindowObject(name string, object core.QObject_ITF, own QWebFrame__ValueOwnership) {
	defer qt.Recovering("QWebFrame::addToJavaScriptWindowObject")

	if ptr.Pointer() != nil {
		C.QWebFrame_AddToJavaScriptWindowObject(ptr.Pointer(), C.CString(name), core.PointerFromQObject(object), C.int(own))
	}
}

func (ptr *QWebFrame) BaseUrl() *core.QUrl {
	defer qt.Recovering("QWebFrame::baseUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebFrame_BaseUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) ContentsSize() *core.QSize {
	defer qt.Recovering("QWebFrame::contentsSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebFrame_ContentsSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) HasFocus() bool {
	defer qt.Recovering("QWebFrame::hasFocus")

	if ptr.Pointer() != nil {
		return C.QWebFrame_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebFrame) Icon() *gui.QIcon {
	defer qt.Recovering("QWebFrame::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QWebFrame_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) RequestedUrl() *core.QUrl {
	defer qt.Recovering("QWebFrame::requestedUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebFrame_RequestedUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) ScrollPosition() *core.QPoint {
	defer qt.Recovering("QWebFrame::scrollPosition")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWebFrame_ScrollPosition(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) SetScrollPosition(pos core.QPoint_ITF) {
	defer qt.Recovering("QWebFrame::setScrollPosition")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollPosition(ptr.Pointer(), core.PointerFromQPoint(pos))
	}
}

func (ptr *QWebFrame) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QWebFrame::setUrl")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebFrame) SetZoomFactor(factor float64) {
	defer qt.Recovering("QWebFrame::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebFrame) Title() string {
	defer qt.Recovering("QWebFrame::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebFrame_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) Url() *core.QUrl {
	defer qt.Recovering("QWebFrame::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebFrame_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) ZoomFactor() float64 {
	defer qt.Recovering("QWebFrame::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QWebFrame_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebFrame_ContentsSizeChanged
func callbackQWebFrame_ContentsSizeChanged(ptr unsafe.Pointer, ptrName *C.char, size unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::contentsSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(size))
	}

}

func (ptr *QWebFrame) ConnectContentsSizeChanged(f func(size *core.QSize)) {
	defer qt.Recovering("connect QWebFrame::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectContentsSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsSizeChanged", f)
	}
}

func (ptr *QWebFrame) DisconnectContentsSizeChanged() {
	defer qt.Recovering("disconnect QWebFrame::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsSizeChanged")
	}
}

func (ptr *QWebFrame) ContentsSizeChanged(size core.QSize_ITF) {
	defer qt.Recovering("QWebFrame::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ContentsSizeChanged(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWebFrame) DocumentElement() *QWebElement {
	defer qt.Recovering("QWebFrame::documentElement")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebFrame_DocumentElement(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebFrame_EvaluateJavaScript
func callbackQWebFrame_EvaluateJavaScript(ptr unsafe.Pointer, ptrName *C.char, scriptSource *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebFrame::evaluateJavaScript")

	if signal := qt.GetSignal(C.GoString(ptrName), "evaluateJavaScript"); signal != nil {
		return core.PointerFromQVariant(signal.(func(string) *core.QVariant)(C.GoString(scriptSource)))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QWebFrame) ConnectEvaluateJavaScript(f func(scriptSource string) *core.QVariant) {
	defer qt.Recovering("connect QWebFrame::evaluateJavaScript")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "evaluateJavaScript", f)
	}
}

func (ptr *QWebFrame) DisconnectEvaluateJavaScript(scriptSource string) {
	defer qt.Recovering("disconnect QWebFrame::evaluateJavaScript")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "evaluateJavaScript")
	}
}

func (ptr *QWebFrame) EvaluateJavaScript(scriptSource string) *core.QVariant {
	defer qt.Recovering("QWebFrame::evaluateJavaScript")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebFrame_EvaluateJavaScript(ptr.Pointer(), C.CString(scriptSource)))
	}
	return nil
}

//export callbackQWebFrame_Event
func callbackQWebFrame_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebFrame::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebFrameFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebFrame) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebFrame::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebFrame) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebFrame::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebFrame) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebFrame::event")

	if ptr.Pointer() != nil {
		return C.QWebFrame_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebFrame) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebFrame::event")

	if ptr.Pointer() != nil {
		return C.QWebFrame_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebFrame) FindAllElements(selectorQuery string) *QWebElementCollection {
	defer qt.Recovering("QWebFrame::findAllElements")

	if ptr.Pointer() != nil {
		return NewQWebElementCollectionFromPointer(C.QWebFrame_FindAllElements(ptr.Pointer(), C.CString(selectorQuery)))
	}
	return nil
}

func (ptr *QWebFrame) FindFirstElement(selectorQuery string) *QWebElement {
	defer qt.Recovering("QWebFrame::findFirstElement")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebFrame_FindFirstElement(ptr.Pointer(), C.CString(selectorQuery)))
	}
	return nil
}

func (ptr *QWebFrame) FrameName() string {
	defer qt.Recovering("QWebFrame::frameName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebFrame_FrameName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) Geometry() *core.QRect {
	defer qt.Recovering("QWebFrame::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWebFrame_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) HitTestContent(pos core.QPoint_ITF) *QWebHitTestResult {
	defer qt.Recovering("QWebFrame::hitTestContent")

	if ptr.Pointer() != nil {
		return NewQWebHitTestResultFromPointer(C.QWebFrame_HitTestContent(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

//export callbackQWebFrame_IconChanged
func callbackQWebFrame_IconChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::iconChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectIconChanged(f func()) {
	defer qt.Recovering("connect QWebFrame::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectIconChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconChanged", f)
	}
}

func (ptr *QWebFrame) DisconnectIconChanged() {
	defer qt.Recovering("disconnect QWebFrame::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconChanged")
	}
}

func (ptr *QWebFrame) IconChanged() {
	defer qt.Recovering("QWebFrame::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_IconChanged(ptr.Pointer())
	}
}

//export callbackQWebFrame_InitialLayoutCompleted
func callbackQWebFrame_InitialLayoutCompleted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::initialLayoutCompleted")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialLayoutCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectInitialLayoutCompleted(f func()) {
	defer qt.Recovering("connect QWebFrame::initialLayoutCompleted")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectInitialLayoutCompleted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "initialLayoutCompleted", f)
	}
}

func (ptr *QWebFrame) DisconnectInitialLayoutCompleted() {
	defer qt.Recovering("disconnect QWebFrame::initialLayoutCompleted")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectInitialLayoutCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "initialLayoutCompleted")
	}
}

func (ptr *QWebFrame) InitialLayoutCompleted() {
	defer qt.Recovering("QWebFrame::initialLayoutCompleted")

	if ptr.Pointer() != nil {
		C.QWebFrame_InitialLayoutCompleted(ptr.Pointer())
	}
}

//export callbackQWebFrame_JavaScriptWindowObjectCleared
func callbackQWebFrame_JavaScriptWindowObjectCleared(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::javaScriptWindowObjectCleared")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptWindowObjectCleared"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectJavaScriptWindowObjectCleared(f func()) {
	defer qt.Recovering("connect QWebFrame::javaScriptWindowObjectCleared")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectJavaScriptWindowObjectCleared(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "javaScriptWindowObjectCleared", f)
	}
}

func (ptr *QWebFrame) DisconnectJavaScriptWindowObjectCleared() {
	defer qt.Recovering("disconnect QWebFrame::javaScriptWindowObjectCleared")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectJavaScriptWindowObjectCleared(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptWindowObjectCleared")
	}
}

func (ptr *QWebFrame) JavaScriptWindowObjectCleared() {
	defer qt.Recovering("QWebFrame::javaScriptWindowObjectCleared")

	if ptr.Pointer() != nil {
		C.QWebFrame_JavaScriptWindowObjectCleared(ptr.Pointer())
	}
}

func (ptr *QWebFrame) Load2(req network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body string) {
	defer qt.Recovering("QWebFrame::load")

	if ptr.Pointer() != nil {
		C.QWebFrame_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(req), C.int(operation), C.CString(body))
	}
}

func (ptr *QWebFrame) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QWebFrame::load")

	if ptr.Pointer() != nil {
		C.QWebFrame_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebFrame_LoadFinished
func callbackQWebFrame_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QWebFrame::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QWebFrame) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QWebFrame::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QWebFrame) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QWebFrame::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QWebFrame) LoadFinished(ok bool) {
	defer qt.Recovering("QWebFrame::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebFrame_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQWebFrame_LoadStarted
func callbackQWebFrame_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QWebFrame::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QWebFrame) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QWebFrame::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QWebFrame) LoadStarted() {
	defer qt.Recovering("QWebFrame::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebFrame_LoadStarted(ptr.Pointer())
	}
}

func (ptr *QWebFrame) Page() *QWebPage {
	defer qt.Recovering("QWebFrame::page")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QWebFrame_Page(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebFrame_PageChanged
func callbackQWebFrame_PageChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::pageChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "pageChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebFrame) ConnectPageChanged(f func()) {
	defer qt.Recovering("connect QWebFrame::pageChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectPageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageChanged", f)
	}
}

func (ptr *QWebFrame) DisconnectPageChanged() {
	defer qt.Recovering("disconnect QWebFrame::pageChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectPageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageChanged")
	}
}

func (ptr *QWebFrame) PageChanged() {
	defer qt.Recovering("QWebFrame::pageChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_PageChanged(ptr.Pointer())
	}
}

func (ptr *QWebFrame) ParentFrame() *QWebFrame {
	defer qt.Recovering("QWebFrame::parentFrame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebFrame_ParentFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) Pos() *core.QPoint {
	defer qt.Recovering("QWebFrame::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWebFrame_Pos(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebFrame_Print
func callbackQWebFrame_Print(ptr unsafe.Pointer, ptrName *C.char, printer unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::print")

	if signal := qt.GetSignal(C.GoString(ptrName), "print"); signal != nil {
		signal.(func(*printsupport.QPrinter))(printsupport.NewQPrinterFromPointer(printer))
	}

}

func (ptr *QWebFrame) ConnectPrint(f func(printer *printsupport.QPrinter)) {
	defer qt.Recovering("connect QWebFrame::print")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "print", f)
	}
}

func (ptr *QWebFrame) DisconnectPrint(printer printsupport.QPrinter_ITF) {
	defer qt.Recovering("disconnect QWebFrame::print")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "print")
	}
}

func (ptr *QWebFrame) Print(printer printsupport.QPrinter_ITF) {
	defer qt.Recovering("QWebFrame::print")

	if ptr.Pointer() != nil {
		C.QWebFrame_Print(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

func (ptr *QWebFrame) Render2(painter gui.QPainter_ITF, layer QWebFrame__RenderLayer, clip gui.QRegion_ITF) {
	defer qt.Recovering("QWebFrame::render")

	if ptr.Pointer() != nil {
		C.QWebFrame_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), C.int(layer), gui.PointerFromQRegion(clip))
	}
}

func (ptr *QWebFrame) Render(painter gui.QPainter_ITF, clip gui.QRegion_ITF) {
	defer qt.Recovering("QWebFrame::render")

	if ptr.Pointer() != nil {
		C.QWebFrame_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), gui.PointerFromQRegion(clip))
	}
}

func (ptr *QWebFrame) Scroll(dx int, dy int) {
	defer qt.Recovering("QWebFrame::scroll")

	if ptr.Pointer() != nil {
		C.QWebFrame_Scroll(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QWebFrame) ScrollBarGeometry(orientation core.Qt__Orientation) *core.QRect {
	defer qt.Recovering("QWebFrame::scrollBarGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWebFrame_ScrollBarGeometry(ptr.Pointer(), C.int(orientation)))
	}
	return nil
}

func (ptr *QWebFrame) ScrollBarMaximum(orientation core.Qt__Orientation) int {
	defer qt.Recovering("QWebFrame::scrollBarMaximum")

	if ptr.Pointer() != nil {
		return int(C.QWebFrame_ScrollBarMaximum(ptr.Pointer(), C.int(orientation)))
	}
	return 0
}

func (ptr *QWebFrame) ScrollBarMinimum(orientation core.Qt__Orientation) int {
	defer qt.Recovering("QWebFrame::scrollBarMinimum")

	if ptr.Pointer() != nil {
		return int(C.QWebFrame_ScrollBarMinimum(ptr.Pointer(), C.int(orientation)))
	}
	return 0
}

func (ptr *QWebFrame) ScrollBarPolicy(orientation core.Qt__Orientation) core.Qt__ScrollBarPolicy {
	defer qt.Recovering("QWebFrame::scrollBarPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollBarPolicy(C.QWebFrame_ScrollBarPolicy(ptr.Pointer(), C.int(orientation)))
	}
	return 0
}

func (ptr *QWebFrame) ScrollBarValue(orientation core.Qt__Orientation) int {
	defer qt.Recovering("QWebFrame::scrollBarValue")

	if ptr.Pointer() != nil {
		return int(C.QWebFrame_ScrollBarValue(ptr.Pointer(), C.int(orientation)))
	}
	return 0
}

func (ptr *QWebFrame) ScrollToAnchor(anchor string) {
	defer qt.Recovering("QWebFrame::scrollToAnchor")

	if ptr.Pointer() != nil {
		C.QWebFrame_ScrollToAnchor(ptr.Pointer(), C.CString(anchor))
	}
}

func (ptr *QWebFrame) SecurityOrigin() *QWebSecurityOrigin {
	defer qt.Recovering("QWebFrame::securityOrigin")

	if ptr.Pointer() != nil {
		return NewQWebSecurityOriginFromPointer(C.QWebFrame_SecurityOrigin(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) SetContent(data string, mimeType string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebFrame::setContent")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetContent(ptr.Pointer(), C.CString(data), C.CString(mimeType), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebFrame) SetFocus() {
	defer qt.Recovering("QWebFrame::setFocus")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetFocus(ptr.Pointer())
	}
}

func (ptr *QWebFrame) SetHtml(html string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebFrame::setHtml")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetHtml(ptr.Pointer(), C.CString(html), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebFrame) SetScrollBarPolicy(orientation core.Qt__Orientation, policy core.Qt__ScrollBarPolicy) {
	defer qt.Recovering("QWebFrame::setScrollBarPolicy")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollBarPolicy(ptr.Pointer(), C.int(orientation), C.int(policy))
	}
}

func (ptr *QWebFrame) SetScrollBarValue(orientation core.Qt__Orientation, value int) {
	defer qt.Recovering("QWebFrame::setScrollBarValue")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetScrollBarValue(ptr.Pointer(), C.int(orientation), C.int(value))
	}
}

func (ptr *QWebFrame) SetTextSizeMultiplier(factor float64) {
	defer qt.Recovering("QWebFrame::setTextSizeMultiplier")

	if ptr.Pointer() != nil {
		C.QWebFrame_SetTextSizeMultiplier(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebFrame) TextSizeMultiplier() float64 {
	defer qt.Recovering("QWebFrame::textSizeMultiplier")

	if ptr.Pointer() != nil {
		return float64(C.QWebFrame_TextSizeMultiplier(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebFrame_TitleChanged
func callbackQWebFrame_TitleChanged(ptr unsafe.Pointer, ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWebFrame::titleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "titleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWebFrame) ConnectTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWebFrame::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "titleChanged", f)
	}
}

func (ptr *QWebFrame) DisconnectTitleChanged() {
	defer qt.Recovering("disconnect QWebFrame::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "titleChanged")
	}
}

func (ptr *QWebFrame) TitleChanged(title string) {
	defer qt.Recovering("QWebFrame::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_TitleChanged(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QWebFrame) ToHtml() string {
	defer qt.Recovering("QWebFrame::toHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebFrame_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebFrame) ToPlainText() string {
	defer qt.Recovering("QWebFrame::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebFrame_ToPlainText(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebFrame_UrlChanged
func callbackQWebFrame_UrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::urlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebFrame) ConnectUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebFrame::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlChanged", f)
	}
}

func (ptr *QWebFrame) DisconnectUrlChanged() {
	defer qt.Recovering("disconnect QWebFrame::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlChanged")
	}
}

func (ptr *QWebFrame) UrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebFrame::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebFrame_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebFrame_TimerEvent
func callbackQWebFrame_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebFrame) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebFrame) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebFrame::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebFrame) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebFrame) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebFrame::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebFrame_ChildEvent
func callbackQWebFrame_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebFrame) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebFrame) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebFrame::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebFrame) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebFrame) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebFrame::childEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebFrame_ConnectNotify
func callbackQWebFrame_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebFrameFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebFrame) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebFrame::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebFrame) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebFrame::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebFrame) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebFrame::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebFrame) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebFrame::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebFrame_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebFrame_CustomEvent
func callbackQWebFrame_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebFrameFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebFrame) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebFrame) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebFrame::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebFrame) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebFrame) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebFrame::customEvent")

	if ptr.Pointer() != nil {
		C.QWebFrame_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebFrame_DeleteLater
func callbackQWebFrame_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebFrame::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebFrameFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebFrame) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebFrame::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebFrame) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebFrame::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebFrame) DeleteLater() {
	defer qt.Recovering("QWebFrame::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebFrame_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebFrame) DeleteLaterDefault() {
	defer qt.Recovering("QWebFrame::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebFrame_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebFrame_DisconnectNotify
func callbackQWebFrame_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebFrame::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebFrameFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebFrame) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebFrame::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebFrame) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebFrame::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebFrame) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebFrame::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebFrame) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebFrame::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebFrame_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebFrame_EventFilter
func callbackQWebFrame_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebFrame::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebFrameFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebFrame) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebFrame::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebFrame) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebFrame::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebFrame) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebFrame::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebFrame_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebFrame) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebFrame::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebFrame_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebFrame_MetaObject
func callbackQWebFrame_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebFrame::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebFrameFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebFrame) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebFrame::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebFrame) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebFrame::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebFrame) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebFrame::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebFrame_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebFrame) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebFrame::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebFrame_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebHistory struct {
	ptr unsafe.Pointer
}

type QWebHistory_ITF interface {
	QWebHistory_PTR() *QWebHistory
}

func (p *QWebHistory) QWebHistory_PTR() *QWebHistory {
	return p
}

func (p *QWebHistory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebHistory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebHistory(ptr QWebHistory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistory_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryFromPointer(ptr unsafe.Pointer) *QWebHistory {
	var n = new(QWebHistory)
	n.SetPointer(ptr)
	return n
}

func newQWebHistoryFromPointer(ptr unsafe.Pointer) *QWebHistory {
	var n = NewQWebHistoryFromPointer(ptr)
	return n
}

func (ptr *QWebHistory) Back() {
	defer qt.Recovering("QWebHistory::back")

	if ptr.Pointer() != nil {
		C.QWebHistory_Back(ptr.Pointer())
	}
}

func (ptr *QWebHistory) BackItem() *QWebHistoryItem {
	defer qt.Recovering("QWebHistory::backItem")

	if ptr.Pointer() != nil {
		return NewQWebHistoryItemFromPointer(C.QWebHistory_BackItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistory) CanGoBack() bool {
	defer qt.Recovering("QWebHistory::canGoBack")

	if ptr.Pointer() != nil {
		return C.QWebHistory_CanGoBack(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistory) CanGoForward() bool {
	defer qt.Recovering("QWebHistory::canGoForward")

	if ptr.Pointer() != nil {
		return C.QWebHistory_CanGoForward(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistory) Clear() {
	defer qt.Recovering("QWebHistory::clear")

	if ptr.Pointer() != nil {
		C.QWebHistory_Clear(ptr.Pointer())
	}
}

func (ptr *QWebHistory) Count() int {
	defer qt.Recovering("QWebHistory::count")

	if ptr.Pointer() != nil {
		return int(C.QWebHistory_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebHistory) CurrentItem() *QWebHistoryItem {
	defer qt.Recovering("QWebHistory::currentItem")

	if ptr.Pointer() != nil {
		return NewQWebHistoryItemFromPointer(C.QWebHistory_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistory) CurrentItemIndex() int {
	defer qt.Recovering("QWebHistory::currentItemIndex")

	if ptr.Pointer() != nil {
		return int(C.QWebHistory_CurrentItemIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebHistory) Forward() {
	defer qt.Recovering("QWebHistory::forward")

	if ptr.Pointer() != nil {
		C.QWebHistory_Forward(ptr.Pointer())
	}
}

func (ptr *QWebHistory) ForwardItem() *QWebHistoryItem {
	defer qt.Recovering("QWebHistory::forwardItem")

	if ptr.Pointer() != nil {
		return NewQWebHistoryItemFromPointer(C.QWebHistory_ForwardItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistory) GoToItem(item QWebHistoryItem_ITF) {
	defer qt.Recovering("QWebHistory::goToItem")

	if ptr.Pointer() != nil {
		C.QWebHistory_GoToItem(ptr.Pointer(), PointerFromQWebHistoryItem(item))
	}
}

func (ptr *QWebHistory) ItemAt(i int) *QWebHistoryItem {
	defer qt.Recovering("QWebHistory::itemAt")

	if ptr.Pointer() != nil {
		return NewQWebHistoryItemFromPointer(C.QWebHistory_ItemAt(ptr.Pointer(), C.int(i)))
	}
	return nil
}

func (ptr *QWebHistory) MaximumItemCount() int {
	defer qt.Recovering("QWebHistory::maximumItemCount")

	if ptr.Pointer() != nil {
		return int(C.QWebHistory_MaximumItemCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebHistory) SetMaximumItemCount(count int) {
	defer qt.Recovering("QWebHistory::setMaximumItemCount")

	if ptr.Pointer() != nil {
		C.QWebHistory_SetMaximumItemCount(ptr.Pointer(), C.int(count))
	}
}

type QWebHistoryInterface struct {
	core.QObject
}

type QWebHistoryInterface_ITF interface {
	core.QObject_ITF
	QWebHistoryInterface_PTR() *QWebHistoryInterface
}

func (p *QWebHistoryInterface) QWebHistoryInterface_PTR() *QWebHistoryInterface {
	return p
}

func (p *QWebHistoryInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebHistoryInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebHistoryInterface(ptr QWebHistoryInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistoryInterface_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryInterfaceFromPointer(ptr unsafe.Pointer) *QWebHistoryInterface {
	var n = new(QWebHistoryInterface)
	n.SetPointer(ptr)
	return n
}

func newQWebHistoryInterfaceFromPointer(ptr unsafe.Pointer) *QWebHistoryInterface {
	var n = NewQWebHistoryInterfaceFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebHistoryInterface_") {
		n.SetObjectName("QWebHistoryInterface_" + qt.Identifier())
	}
	return n
}

func NewQWebHistoryInterface(parent core.QObject_ITF) *QWebHistoryInterface {
	defer qt.Recovering("QWebHistoryInterface::QWebHistoryInterface")

	return newQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_NewQWebHistoryInterface(core.PointerFromQObject(parent)))
}

//export callbackQWebHistoryInterface_AddHistoryEntry
func callbackQWebHistoryInterface_AddHistoryEntry(ptr unsafe.Pointer, ptrName *C.char, url *C.char) {
	defer qt.Recovering("callback QWebHistoryInterface::addHistoryEntry")

	if signal := qt.GetSignal(C.GoString(ptrName), "addHistoryEntry"); signal != nil {
		signal.(func(string))(C.GoString(url))
	}

}

func (ptr *QWebHistoryInterface) ConnectAddHistoryEntry(f func(url string)) {
	defer qt.Recovering("connect QWebHistoryInterface::addHistoryEntry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "addHistoryEntry", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectAddHistoryEntry(url string) {
	defer qt.Recovering("disconnect QWebHistoryInterface::addHistoryEntry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "addHistoryEntry")
	}
}

func (ptr *QWebHistoryInterface) AddHistoryEntry(url string) {
	defer qt.Recovering("QWebHistoryInterface::addHistoryEntry")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_AddHistoryEntry(ptr.Pointer(), C.CString(url))
	}
}

func QWebHistoryInterface_DefaultInterface() *QWebHistoryInterface {
	defer qt.Recovering("QWebHistoryInterface::defaultInterface")

	return NewQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_QWebHistoryInterface_DefaultInterface())
}

func (ptr *QWebHistoryInterface) DefaultInterface() *QWebHistoryInterface {
	defer qt.Recovering("QWebHistoryInterface::defaultInterface")

	return NewQWebHistoryInterfaceFromPointer(C.QWebHistoryInterface_QWebHistoryInterface_DefaultInterface())
}

//export callbackQWebHistoryInterface_HistoryContains
func callbackQWebHistoryInterface_HistoryContains(ptr unsafe.Pointer, ptrName *C.char, url *C.char) C.int {
	defer qt.Recovering("callback QWebHistoryInterface::historyContains")

	if signal := qt.GetSignal(C.GoString(ptrName), "historyContains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(url))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QWebHistoryInterface) ConnectHistoryContains(f func(url string) bool) {
	defer qt.Recovering("connect QWebHistoryInterface::historyContains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "historyContains", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectHistoryContains(url string) {
	defer qt.Recovering("disconnect QWebHistoryInterface::historyContains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "historyContains")
	}
}

func (ptr *QWebHistoryInterface) HistoryContains(url string) bool {
	defer qt.Recovering("QWebHistoryInterface::historyContains")

	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_HistoryContains(ptr.Pointer(), C.CString(url)) != 0
	}
	return false
}

func QWebHistoryInterface_SetDefaultInterface(defaultInterface QWebHistoryInterface_ITF) {
	defer qt.Recovering("QWebHistoryInterface::setDefaultInterface")

	C.QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(PointerFromQWebHistoryInterface(defaultInterface))
}

func (ptr *QWebHistoryInterface) SetDefaultInterface(defaultInterface QWebHistoryInterface_ITF) {
	defer qt.Recovering("QWebHistoryInterface::setDefaultInterface")

	C.QWebHistoryInterface_QWebHistoryInterface_SetDefaultInterface(PointerFromQWebHistoryInterface(defaultInterface))
}

func (ptr *QWebHistoryInterface) DestroyQWebHistoryInterface() {
	defer qt.Recovering("QWebHistoryInterface::~QWebHistoryInterface")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DestroyQWebHistoryInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebHistoryInterface_TimerEvent
func callbackQWebHistoryInterface_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebHistoryInterface::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebHistoryInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebHistoryInterface::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebHistoryInterface) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebHistoryInterface) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebHistoryInterface_ChildEvent
func callbackQWebHistoryInterface_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebHistoryInterface::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebHistoryInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebHistoryInterface::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebHistoryInterface) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebHistoryInterface) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::childEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebHistoryInterface_ConnectNotify
func callbackQWebHistoryInterface_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebHistoryInterface::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebHistoryInterface) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebHistoryInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebHistoryInterface::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebHistoryInterface) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebHistoryInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebHistoryInterface) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebHistoryInterface::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebHistoryInterface_CustomEvent
func callbackQWebHistoryInterface_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebHistoryInterface::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebHistoryInterface) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebHistoryInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebHistoryInterface::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebHistoryInterface) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebHistoryInterface) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebHistoryInterface::customEvent")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebHistoryInterface_DeleteLater
func callbackQWebHistoryInterface_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebHistoryInterface::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebHistoryInterface) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebHistoryInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebHistoryInterface::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebHistoryInterface) DeleteLater() {
	defer qt.Recovering("QWebHistoryInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebHistoryInterface) DeleteLaterDefault() {
	defer qt.Recovering("QWebHistoryInterface::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebHistoryInterface_DisconnectNotify
func callbackQWebHistoryInterface_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebHistoryInterface::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebHistoryInterfaceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebHistoryInterface) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebHistoryInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebHistoryInterface::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebHistoryInterface) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebHistoryInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebHistoryInterface) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebHistoryInterface::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebHistoryInterface_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebHistoryInterface_Event
func callbackQWebHistoryInterface_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebHistoryInterface::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebHistoryInterfaceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebHistoryInterface) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebHistoryInterface::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebHistoryInterface::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebHistoryInterface) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebHistoryInterface::event")

	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebHistoryInterface) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebHistoryInterface::event")

	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebHistoryInterface_EventFilter
func callbackQWebHistoryInterface_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebHistoryInterface::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebHistoryInterfaceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebHistoryInterface) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebHistoryInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebHistoryInterface::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebHistoryInterface) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebHistoryInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebHistoryInterface) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebHistoryInterface::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebHistoryInterface_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebHistoryInterface_MetaObject
func callbackQWebHistoryInterface_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebHistoryInterface::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebHistoryInterfaceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebHistoryInterface) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebHistoryInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebHistoryInterface) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebHistoryInterface::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebHistoryInterface) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebHistoryInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebHistoryInterface_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryInterface) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebHistoryInterface::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebHistoryInterface_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebHistoryItem struct {
	ptr unsafe.Pointer
}

type QWebHistoryItem_ITF interface {
	QWebHistoryItem_PTR() *QWebHistoryItem
}

func (p *QWebHistoryItem) QWebHistoryItem_PTR() *QWebHistoryItem {
	return p
}

func (p *QWebHistoryItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebHistoryItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebHistoryItem(ptr QWebHistoryItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHistoryItem_PTR().Pointer()
	}
	return nil
}

func NewQWebHistoryItemFromPointer(ptr unsafe.Pointer) *QWebHistoryItem {
	var n = new(QWebHistoryItem)
	n.SetPointer(ptr)
	return n
}

func newQWebHistoryItemFromPointer(ptr unsafe.Pointer) *QWebHistoryItem {
	var n = NewQWebHistoryItemFromPointer(ptr)
	return n
}

func NewQWebHistoryItem(other QWebHistoryItem_ITF) *QWebHistoryItem {
	defer qt.Recovering("QWebHistoryItem::QWebHistoryItem")

	return newQWebHistoryItemFromPointer(C.QWebHistoryItem_NewQWebHistoryItem(PointerFromQWebHistoryItem(other)))
}

func (ptr *QWebHistoryItem) Icon() *gui.QIcon {
	defer qt.Recovering("QWebHistoryItem::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QWebHistoryItem_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryItem) IsValid() bool {
	defer qt.Recovering("QWebHistoryItem::isValid")

	if ptr.Pointer() != nil {
		return C.QWebHistoryItem_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHistoryItem) LastVisited() *core.QDateTime {
	defer qt.Recovering("QWebHistoryItem::lastVisited")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QWebHistoryItem_LastVisited(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryItem) OriginalUrl() *core.QUrl {
	defer qt.Recovering("QWebHistoryItem::originalUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebHistoryItem_OriginalUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryItem) SetUserData(userData core.QVariant_ITF) {
	defer qt.Recovering("QWebHistoryItem::setUserData")

	if ptr.Pointer() != nil {
		C.QWebHistoryItem_SetUserData(ptr.Pointer(), core.PointerFromQVariant(userData))
	}
}

func (ptr *QWebHistoryItem) Title() string {
	defer qt.Recovering("QWebHistoryItem::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebHistoryItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHistoryItem) Url() *core.QUrl {
	defer qt.Recovering("QWebHistoryItem::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebHistoryItem_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryItem) UserData() *core.QVariant {
	defer qt.Recovering("QWebHistoryItem::userData")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebHistoryItem_UserData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHistoryItem) DestroyQWebHistoryItem() {
	defer qt.Recovering("QWebHistoryItem::~QWebHistoryItem")

	if ptr.Pointer() != nil {
		C.QWebHistoryItem_DestroyQWebHistoryItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebHitTestResult struct {
	ptr unsafe.Pointer
}

type QWebHitTestResult_ITF interface {
	QWebHitTestResult_PTR() *QWebHitTestResult
}

func (p *QWebHitTestResult) QWebHitTestResult_PTR() *QWebHitTestResult {
	return p
}

func (p *QWebHitTestResult) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebHitTestResult) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebHitTestResult(ptr QWebHitTestResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebHitTestResult_PTR().Pointer()
	}
	return nil
}

func NewQWebHitTestResultFromPointer(ptr unsafe.Pointer) *QWebHitTestResult {
	var n = new(QWebHitTestResult)
	n.SetPointer(ptr)
	return n
}

func newQWebHitTestResultFromPointer(ptr unsafe.Pointer) *QWebHitTestResult {
	var n = NewQWebHitTestResultFromPointer(ptr)
	return n
}

func NewQWebHitTestResult() *QWebHitTestResult {
	defer qt.Recovering("QWebHitTestResult::QWebHitTestResult")

	return newQWebHitTestResultFromPointer(C.QWebHitTestResult_NewQWebHitTestResult())
}

func NewQWebHitTestResult2(other QWebHitTestResult_ITF) *QWebHitTestResult {
	defer qt.Recovering("QWebHitTestResult::QWebHitTestResult")

	return newQWebHitTestResultFromPointer(C.QWebHitTestResult_NewQWebHitTestResult2(PointerFromQWebHitTestResult(other)))
}

func (ptr *QWebHitTestResult) AlternateText() string {
	defer qt.Recovering("QWebHitTestResult::alternateText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebHitTestResult_AlternateText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) BoundingRect() *core.QRect {
	defer qt.Recovering("QWebHitTestResult::boundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWebHitTestResult_BoundingRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) Element() *QWebElement {
	defer qt.Recovering("QWebHitTestResult::element")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebHitTestResult_Element(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) EnclosingBlockElement() *QWebElement {
	defer qt.Recovering("QWebHitTestResult::enclosingBlockElement")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebHitTestResult_EnclosingBlockElement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) Frame() *QWebFrame {
	defer qt.Recovering("QWebHitTestResult::frame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebHitTestResult_Frame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) ImageUrl() *core.QUrl {
	defer qt.Recovering("QWebHitTestResult::imageUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebHitTestResult_ImageUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) IsContentEditable() bool {
	defer qt.Recovering("QWebHitTestResult::isContentEditable")

	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsContentEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHitTestResult) IsContentSelected() bool {
	defer qt.Recovering("QWebHitTestResult::isContentSelected")

	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsContentSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHitTestResult) IsNull() bool {
	defer qt.Recovering("QWebHitTestResult::isNull")

	if ptr.Pointer() != nil {
		return C.QWebHitTestResult_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebHitTestResult) LinkElement() *QWebElement {
	defer qt.Recovering("QWebHitTestResult::linkElement")

	if ptr.Pointer() != nil {
		return NewQWebElementFromPointer(C.QWebHitTestResult_LinkElement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) LinkTargetFrame() *QWebFrame {
	defer qt.Recovering("QWebHitTestResult::linkTargetFrame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebHitTestResult_LinkTargetFrame(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) LinkText() string {
	defer qt.Recovering("QWebHitTestResult::linkText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebHitTestResult_LinkText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) LinkTitleString() string {
	defer qt.Recovering("QWebHitTestResult::linkTitleString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebHitTestResult_LinkTitleString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) LinkUrl() *core.QUrl {
	defer qt.Recovering("QWebHitTestResult::linkUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebHitTestResult_LinkUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) MediaUrl() *core.QUrl {
	defer qt.Recovering("QWebHitTestResult::mediaUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebHitTestResult_MediaUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) Pixmap() *gui.QPixmap {
	defer qt.Recovering("QWebHitTestResult::pixmap")

	if ptr.Pointer() != nil {
		return gui.NewQPixmapFromPointer(C.QWebHitTestResult_Pixmap(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) Pos() *core.QPoint {
	defer qt.Recovering("QWebHitTestResult::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWebHitTestResult_Pos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebHitTestResult) Title() string {
	defer qt.Recovering("QWebHitTestResult::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebHitTestResult_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebHitTestResult) DestroyQWebHitTestResult() {
	defer qt.Recovering("QWebHitTestResult::~QWebHitTestResult")

	if ptr.Pointer() != nil {
		C.QWebHitTestResult_DestroyQWebHitTestResult(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWebInspector struct {
	widgets.QWidget
}

type QWebInspector_ITF interface {
	widgets.QWidget_ITF
	QWebInspector_PTR() *QWebInspector
}

func (p *QWebInspector) QWebInspector_PTR() *QWebInspector {
	return p
}

func (p *QWebInspector) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QWebInspector) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebInspector(ptr QWebInspector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebInspector_PTR().Pointer()
	}
	return nil
}

func NewQWebInspectorFromPointer(ptr unsafe.Pointer) *QWebInspector {
	var n = new(QWebInspector)
	n.SetPointer(ptr)
	return n
}

func newQWebInspectorFromPointer(ptr unsafe.Pointer) *QWebInspector {
	var n = NewQWebInspectorFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebInspector_") {
		n.SetObjectName("QWebInspector_" + qt.Identifier())
	}
	return n
}

func NewQWebInspector(parent widgets.QWidget_ITF) *QWebInspector {
	defer qt.Recovering("QWebInspector::QWebInspector")

	return newQWebInspectorFromPointer(C.QWebInspector_NewQWebInspector(widgets.PointerFromQWidget(parent)))
}

//export callbackQWebInspector_CloseEvent
func callbackQWebInspector_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWebInspector::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWebInspector::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *QWebInspector) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebInspector::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWebInspector) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebInspector::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebInspector_Event
func callbackQWebInspector_Event(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebInspector::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev))))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev))))
}

func (ptr *QWebInspector) ConnectEvent(f func(ev *core.QEvent) bool) {
	defer qt.Recovering("connect QWebInspector::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebInspector) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebInspector::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebInspector) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebInspector::event")

	if ptr.Pointer() != nil {
		return C.QWebInspector_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebInspector) EventDefault(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebInspector::event")

	if ptr.Pointer() != nil {
		return C.QWebInspector_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQWebInspector_HideEvent
func callbackQWebInspector_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWebInspector::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWebInspector::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *QWebInspector) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebInspector::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWebInspector) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebInspector::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWebInspector) Page() *QWebPage {
	defer qt.Recovering("QWebInspector::page")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QWebInspector_Page(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebInspector_ResizeEvent
func callbackQWebInspector_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWebInspector::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWebInspector::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *QWebInspector) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebInspector::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWebInspector) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebInspector::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QWebInspector) SetPage(page QWebPage_ITF) {
	defer qt.Recovering("QWebInspector::setPage")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

//export callbackQWebInspector_ShowEvent
func callbackQWebInspector_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWebInspector::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWebInspector::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *QWebInspector) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebInspector::showEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWebInspector) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebInspector::showEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebInspector_SizeHint
func callbackQWebInspector_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebInspector::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebInspectorFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebInspector) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebInspector::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QWebInspector) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QWebInspector::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QWebInspector) SizeHint() *core.QSize {
	defer qt.Recovering("QWebInspector::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebInspector_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebInspector) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebInspector::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebInspector_SizeHintDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebInspector) DestroyQWebInspector() {
	defer qt.Recovering("QWebInspector::~QWebInspector")

	if ptr.Pointer() != nil {
		C.QWebInspector_DestroyQWebInspector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebInspector_ActionEvent
func callbackQWebInspector_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWebInspector::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWebInspector::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *QWebInspector) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebInspector::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWebInspector) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebInspector::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebInspector_DragEnterEvent
func callbackQWebInspector_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWebInspector::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWebInspector::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QWebInspector) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QWebInspector) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQWebInspector_DragLeaveEvent
func callbackQWebInspector_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWebInspector::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWebInspector::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QWebInspector) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QWebInspector) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQWebInspector_DragMoveEvent
func callbackQWebInspector_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWebInspector::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWebInspector::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QWebInspector) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QWebInspector) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebInspector::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQWebInspector_DropEvent
func callbackQWebInspector_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QWebInspector::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWebInspector::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QWebInspector) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebInspector::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QWebInspector) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebInspector::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQWebInspector_EnterEvent
func callbackQWebInspector_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebInspector::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWebInspector::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *QWebInspector) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebInspector) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_FocusInEvent
func callbackQWebInspector_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebInspector::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWebInspector::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QWebInspector) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebInspector::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebInspector) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebInspector::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebInspector_FocusOutEvent
func callbackQWebInspector_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebInspector::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWebInspector::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QWebInspector) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebInspector::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QWebInspector) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebInspector::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQWebInspector_LeaveEvent
func callbackQWebInspector_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebInspector::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWebInspector::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *QWebInspector) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebInspector) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_Metric
func callbackQWebInspector_Metric(ptr unsafe.Pointer, ptrName *C.char, m C.int) C.int {
	defer qt.Recovering("callback QWebInspector::metric")

	if signal := qt.GetSignal(C.GoString(ptrName), "metric"); signal != nil {
		return C.int(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m)))
	}

	return C.int(NewQWebInspectorFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m)))
}

func (ptr *QWebInspector) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QWebInspector::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metric", f)
	}
}

func (ptr *QWebInspector) DisconnectMetric() {
	defer qt.Recovering("disconnect QWebInspector::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metric")
	}
}

func (ptr *QWebInspector) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QWebInspector::metric")

	if ptr.Pointer() != nil {
		return int(C.QWebInspector_Metric(ptr.Pointer(), C.int(m)))
	}
	return 0
}

func (ptr *QWebInspector) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QWebInspector::metric")

	if ptr.Pointer() != nil {
		return int(C.QWebInspector_MetricDefault(ptr.Pointer(), C.int(m)))
	}
	return 0
}

//export callbackQWebInspector_MinimumSizeHint
func callbackQWebInspector_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebInspector::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebInspectorFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebInspector) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebInspector::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QWebInspector) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QWebInspector::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QWebInspector) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QWebInspector::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebInspector_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebInspector) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebInspector::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebInspector_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebInspector_MoveEvent
func callbackQWebInspector_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWebInspector::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWebInspector::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *QWebInspector) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebInspector::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWebInspector) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebInspector::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebInspector_PaintEngine
func callbackQWebInspector_PaintEngine(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebInspector::paintEngine")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQWebInspectorFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebInspector) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	defer qt.Recovering("connect QWebInspector::paintEngine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEngine", f)
	}
}

func (ptr *QWebInspector) DisconnectPaintEngine() {
	defer qt.Recovering("disconnect QWebInspector::paintEngine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEngine")
	}
}

func (ptr *QWebInspector) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QWebInspector::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebInspector_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebInspector) PaintEngineDefault() *gui.QPaintEngine {
	defer qt.Recovering("QWebInspector::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebInspector_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebInspector_PaintEvent
func callbackQWebInspector_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWebInspector::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWebInspector::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *QWebInspector) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebInspector::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QWebInspector) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebInspector::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

//export callbackQWebInspector_SetEnabled
func callbackQWebInspector_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebInspector::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QWebInspector) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QWebInspector::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QWebInspector) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QWebInspector::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QWebInspector) SetEnabled(vbo bool) {
	defer qt.Recovering("QWebInspector::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebInspector) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QWebInspector::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebInspector_SetStyleSheet
func callbackQWebInspector_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QWebInspector::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQWebInspectorFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QWebInspector) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QWebInspector::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QWebInspector) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QWebInspector::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QWebInspector) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QWebInspector::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QWebInspector) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QWebInspector::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQWebInspector_SetVisible
func callbackQWebInspector_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebInspector::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QWebInspector) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWebInspector::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWebInspector) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWebInspector::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QWebInspector) SetVisible(visible bool) {
	defer qt.Recovering("QWebInspector::setVisible")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWebInspector) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QWebInspector::setVisible")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQWebInspector_SetWindowModified
func callbackQWebInspector_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebInspector::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QWebInspector) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QWebInspector::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QWebInspector) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QWebInspector::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QWebInspector) SetWindowModified(vbo bool) {
	defer qt.Recovering("QWebInspector::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebInspector) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QWebInspector::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebInspector_SetWindowTitle
func callbackQWebInspector_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QWebInspector::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQWebInspectorFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QWebInspector) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QWebInspector::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QWebInspector) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QWebInspector::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QWebInspector) SetWindowTitle(vqs string) {
	defer qt.Recovering("QWebInspector::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QWebInspector) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QWebInspector::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQWebInspector_ChangeEvent
func callbackQWebInspector_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebInspector::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QWebInspector::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *QWebInspector) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebInspector) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_Close
func callbackQWebInspector_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebInspector::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).CloseDefault()))
}

func (ptr *QWebInspector) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QWebInspector::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QWebInspector) DisconnectClose() {
	defer qt.Recovering("disconnect QWebInspector::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QWebInspector) Close() bool {
	defer qt.Recovering("QWebInspector::close")

	if ptr.Pointer() != nil {
		return C.QWebInspector_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebInspector) CloseDefault() bool {
	defer qt.Recovering("QWebInspector::close")

	if ptr.Pointer() != nil {
		return C.QWebInspector_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebInspector_ContextMenuEvent
func callbackQWebInspector_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWebInspector::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWebInspector::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QWebInspector) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebInspector::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QWebInspector) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebInspector::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQWebInspector_FocusNextPrevChild
func callbackQWebInspector_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QWebInspector::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QWebInspector) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QWebInspector::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QWebInspector) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QWebInspector::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QWebInspector) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QWebInspector::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebInspector_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QWebInspector) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QWebInspector::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebInspector_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQWebInspector_HasHeightForWidth
func callbackQWebInspector_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebInspector::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QWebInspector) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QWebInspector::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QWebInspector) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QWebInspector::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QWebInspector) HasHeightForWidth() bool {
	defer qt.Recovering("QWebInspector::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebInspector_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebInspector) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QWebInspector::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebInspector_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebInspector_HeightForWidth
func callbackQWebInspector_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QWebInspector::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQWebInspectorFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QWebInspector) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QWebInspector::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QWebInspector) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QWebInspector::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QWebInspector) HeightForWidth(w int) int {
	defer qt.Recovering("QWebInspector::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebInspector_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWebInspector) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QWebInspector::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebInspector_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQWebInspector_Hide
func callbackQWebInspector_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebInspector) ConnectHide(f func()) {
	defer qt.Recovering("connect QWebInspector::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QWebInspector) DisconnectHide() {
	defer qt.Recovering("disconnect QWebInspector::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QWebInspector) Hide() {
	defer qt.Recovering("QWebInspector::hide")

	if ptr.Pointer() != nil {
		C.QWebInspector_Hide(ptr.Pointer())
	}
}

func (ptr *QWebInspector) HideDefault() {
	defer qt.Recovering("QWebInspector::hide")

	if ptr.Pointer() != nil {
		C.QWebInspector_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_InitPainter
func callbackQWebInspector_InitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQWebInspectorFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QWebInspector) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QWebInspector::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QWebInspector) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QWebInspector::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

func (ptr *QWebInspector) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWebInspector::initPainter")

	if ptr.Pointer() != nil {
		C.QWebInspector_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWebInspector) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWebInspector::initPainter")

	if ptr.Pointer() != nil {
		C.QWebInspector_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQWebInspector_InputMethodEvent
func callbackQWebInspector_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWebInspector::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWebInspector::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QWebInspector) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebInspector::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QWebInspector) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebInspector::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQWebInspector_InputMethodQuery
func callbackQWebInspector_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebInspector::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQWebInspectorFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QWebInspector) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QWebInspector::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QWebInspector) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QWebInspector::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QWebInspector) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebInspector::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebInspector_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QWebInspector) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebInspector::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebInspector_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQWebInspector_KeyPressEvent
func callbackQWebInspector_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebInspector::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWebInspector::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QWebInspector) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebInspector::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebInspector) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebInspector::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebInspector_KeyReleaseEvent
func callbackQWebInspector_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebInspector::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWebInspector::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QWebInspector) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebInspector::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QWebInspector) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebInspector::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQWebInspector_Lower
func callbackQWebInspector_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebInspector) ConnectLower(f func()) {
	defer qt.Recovering("connect QWebInspector::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QWebInspector) DisconnectLower() {
	defer qt.Recovering("disconnect QWebInspector::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QWebInspector) Lower() {
	defer qt.Recovering("QWebInspector::lower")

	if ptr.Pointer() != nil {
		C.QWebInspector_Lower(ptr.Pointer())
	}
}

func (ptr *QWebInspector) LowerDefault() {
	defer qt.Recovering("QWebInspector::lower")

	if ptr.Pointer() != nil {
		C.QWebInspector_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_MouseDoubleClickEvent
func callbackQWebInspector_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebInspector::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWebInspector::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QWebInspector) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebInspector) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MouseMoveEvent
func callbackQWebInspector_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebInspector::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWebInspector::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QWebInspector) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebInspector) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MousePressEvent
func callbackQWebInspector_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebInspector::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWebInspector::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QWebInspector) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebInspector) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_MouseReleaseEvent
func callbackQWebInspector_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebInspector::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWebInspector::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QWebInspector) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QWebInspector) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebInspector::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQWebInspector_NativeEvent
func callbackQWebInspector_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QWebInspector::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QWebInspector) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QWebInspector::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QWebInspector::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QWebInspector) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebInspector::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebInspector_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QWebInspector) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebInspector::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebInspector_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQWebInspector_Raise
func callbackQWebInspector_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebInspector) ConnectRaise(f func()) {
	defer qt.Recovering("connect QWebInspector::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QWebInspector) DisconnectRaise() {
	defer qt.Recovering("disconnect QWebInspector::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QWebInspector) Raise() {
	defer qt.Recovering("QWebInspector::raise")

	if ptr.Pointer() != nil {
		C.QWebInspector_Raise(ptr.Pointer())
	}
}

func (ptr *QWebInspector) RaiseDefault() {
	defer qt.Recovering("QWebInspector::raise")

	if ptr.Pointer() != nil {
		C.QWebInspector_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_Repaint
func callbackQWebInspector_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebInspector) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QWebInspector::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QWebInspector) DisconnectRepaint() {
	defer qt.Recovering("disconnect QWebInspector::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QWebInspector) Repaint() {
	defer qt.Recovering("QWebInspector::repaint")

	if ptr.Pointer() != nil {
		C.QWebInspector_Repaint(ptr.Pointer())
	}
}

func (ptr *QWebInspector) RepaintDefault() {
	defer qt.Recovering("QWebInspector::repaint")

	if ptr.Pointer() != nil {
		C.QWebInspector_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_SetDisabled
func callbackQWebInspector_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QWebInspector::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QWebInspector) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QWebInspector::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QWebInspector) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QWebInspector::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QWebInspector) SetDisabled(disable bool) {
	defer qt.Recovering("QWebInspector::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWebInspector) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QWebInspector::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQWebInspector_SetFocus2
func callbackQWebInspector_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebInspector) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QWebInspector::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QWebInspector) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QWebInspector::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QWebInspector) SetFocus2() {
	defer qt.Recovering("QWebInspector::setFocus")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWebInspector) SetFocus2Default() {
	defer qt.Recovering("QWebInspector::setFocus")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebInspector_SetHidden
func callbackQWebInspector_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QWebInspector::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQWebInspectorFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QWebInspector) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QWebInspector::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QWebInspector) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QWebInspector::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QWebInspector) SetHidden(hidden bool) {
	defer qt.Recovering("QWebInspector::setHidden")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWebInspector) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QWebInspector::setHidden")

	if ptr.Pointer() != nil {
		C.QWebInspector_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQWebInspector_Show
func callbackQWebInspector_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebInspector) ConnectShow(f func()) {
	defer qt.Recovering("connect QWebInspector::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QWebInspector) DisconnectShow() {
	defer qt.Recovering("disconnect QWebInspector::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QWebInspector) Show() {
	defer qt.Recovering("QWebInspector::show")

	if ptr.Pointer() != nil {
		C.QWebInspector_Show(ptr.Pointer())
	}
}

func (ptr *QWebInspector) ShowDefault() {
	defer qt.Recovering("QWebInspector::show")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowFullScreen
func callbackQWebInspector_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebInspector) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QWebInspector::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QWebInspector) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QWebInspector::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QWebInspector) ShowFullScreen() {
	defer qt.Recovering("QWebInspector::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWebInspector) ShowFullScreenDefault() {
	defer qt.Recovering("QWebInspector::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowMaximized
func callbackQWebInspector_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebInspector) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QWebInspector::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QWebInspector) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QWebInspector::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QWebInspector) ShowMaximized() {
	defer qt.Recovering("QWebInspector::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWebInspector) ShowMaximizedDefault() {
	defer qt.Recovering("QWebInspector::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowMinimized
func callbackQWebInspector_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebInspector) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QWebInspector::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QWebInspector) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QWebInspector::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QWebInspector) ShowMinimized() {
	defer qt.Recovering("QWebInspector::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWebInspector) ShowMinimizedDefault() {
	defer qt.Recovering("QWebInspector::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_ShowNormal
func callbackQWebInspector_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebInspector) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QWebInspector::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QWebInspector) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QWebInspector::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QWebInspector) ShowNormal() {
	defer qt.Recovering("QWebInspector::showNormal")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWebInspector) ShowNormalDefault() {
	defer qt.Recovering("QWebInspector::showNormal")

	if ptr.Pointer() != nil {
		C.QWebInspector_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_TabletEvent
func callbackQWebInspector_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWebInspector::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWebInspector::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *QWebInspector) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebInspector::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWebInspector) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebInspector::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebInspector_Update
func callbackQWebInspector_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebInspector) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QWebInspector::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QWebInspector) DisconnectUpdate() {
	defer qt.Recovering("disconnect QWebInspector::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QWebInspector) Update() {
	defer qt.Recovering("QWebInspector::update")

	if ptr.Pointer() != nil {
		C.QWebInspector_Update(ptr.Pointer())
	}
}

func (ptr *QWebInspector) UpdateDefault() {
	defer qt.Recovering("QWebInspector::update")

	if ptr.Pointer() != nil {
		C.QWebInspector_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_UpdateMicroFocus
func callbackQWebInspector_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebInspector) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QWebInspector::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QWebInspector) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QWebInspector::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QWebInspector) UpdateMicroFocus() {
	defer qt.Recovering("QWebInspector::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebInspector_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QWebInspector) UpdateMicroFocusDefault() {
	defer qt.Recovering("QWebInspector::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebInspector_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebInspector_WheelEvent
func callbackQWebInspector_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWebInspector::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWebInspector::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QWebInspector) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebInspector::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QWebInspector) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebInspector::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQWebInspector_TimerEvent
func callbackQWebInspector_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebInspector::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebInspector::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebInspector) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebInspector::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebInspector) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebInspector::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebInspector_ChildEvent
func callbackQWebInspector_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebInspector::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebInspector::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebInspector) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebInspector::childEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebInspector) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebInspector::childEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebInspector_ConnectNotify
func callbackQWebInspector_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebInspectorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebInspector) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebInspector::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebInspector) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebInspector::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebInspector) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebInspector::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebInspector_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebInspector) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebInspector::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebInspector_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebInspector_CustomEvent
func callbackQWebInspector_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebInspectorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebInspector) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebInspector::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebInspector) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebInspector::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebInspector) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::customEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebInspector) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebInspector::customEvent")

	if ptr.Pointer() != nil {
		C.QWebInspector_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebInspector_DeleteLater
func callbackQWebInspector_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebInspector::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebInspectorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebInspector) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebInspector::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebInspector) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebInspector::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebInspector) DeleteLater() {
	defer qt.Recovering("QWebInspector::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebInspector_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebInspector) DeleteLaterDefault() {
	defer qt.Recovering("QWebInspector::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebInspector_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebInspector_DisconnectNotify
func callbackQWebInspector_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebInspector::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebInspectorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebInspector) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebInspector::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebInspector) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebInspector::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebInspector) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebInspector::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebInspector_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebInspector) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebInspector::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebInspector_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebInspector_EventFilter
func callbackQWebInspector_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebInspector::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebInspectorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebInspector) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebInspector::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebInspector) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebInspector::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebInspector) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebInspector::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebInspector_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebInspector) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebInspector::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebInspector_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebInspector_MetaObject
func callbackQWebInspector_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebInspector::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebInspectorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebInspector) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebInspector::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebInspector) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebInspector::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebInspector) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebInspector::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebInspector_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebInspector) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebInspector::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebInspector_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebPage::ErrorDomain
type QWebPage__ErrorDomain int64

const (
	QWebPage__QtNetwork = QWebPage__ErrorDomain(0)
	QWebPage__Http      = QWebPage__ErrorDomain(1)
	QWebPage__WebKit    = QWebPage__ErrorDomain(2)
)

//QWebPage::Extension
type QWebPage__Extension int64

const (
	QWebPage__ChooseMultipleFilesExtension = QWebPage__Extension(0)
	QWebPage__ErrorPageExtension           = QWebPage__Extension(1)
)

//QWebPage::Feature
type QWebPage__Feature int64

const (
	QWebPage__Notifications = QWebPage__Feature(0)
	QWebPage__Geolocation   = QWebPage__Feature(1)
)

//QWebPage::FindFlag
type QWebPage__FindFlag int64

const (
	QWebPage__FindBackward                      = QWebPage__FindFlag(1)
	QWebPage__FindCaseSensitively               = QWebPage__FindFlag(2)
	QWebPage__FindWrapsAroundDocument           = QWebPage__FindFlag(4)
	QWebPage__HighlightAllOccurrences           = QWebPage__FindFlag(8)
	QWebPage__FindAtWordBeginningsOnly          = QWebPage__FindFlag(16)
	QWebPage__TreatMedialCapitalAsWordBeginning = QWebPage__FindFlag(32)
	QWebPage__FindBeginsInSelection             = QWebPage__FindFlag(64)
)

//QWebPage::LinkDelegationPolicy
type QWebPage__LinkDelegationPolicy int64

const (
	QWebPage__DontDelegateLinks     = QWebPage__LinkDelegationPolicy(0)
	QWebPage__DelegateExternalLinks = QWebPage__LinkDelegationPolicy(1)
	QWebPage__DelegateAllLinks      = QWebPage__LinkDelegationPolicy(2)
)

//QWebPage::NavigationType
type QWebPage__NavigationType int64

const (
	QWebPage__NavigationTypeLinkClicked     = QWebPage__NavigationType(0)
	QWebPage__NavigationTypeFormSubmitted   = QWebPage__NavigationType(1)
	QWebPage__NavigationTypeBackOrForward   = QWebPage__NavigationType(2)
	QWebPage__NavigationTypeReload          = QWebPage__NavigationType(3)
	QWebPage__NavigationTypeFormResubmitted = QWebPage__NavigationType(4)
	QWebPage__NavigationTypeOther           = QWebPage__NavigationType(5)
)

//QWebPage::PermissionPolicy
type QWebPage__PermissionPolicy int64

const (
	QWebPage__PermissionUnknown       = QWebPage__PermissionPolicy(0)
	QWebPage__PermissionGrantedByUser = QWebPage__PermissionPolicy(1)
	QWebPage__PermissionDeniedByUser  = QWebPage__PermissionPolicy(2)
)

//QWebPage::VisibilityState
type QWebPage__VisibilityState int64

const (
	QWebPage__VisibilityStateVisible   = QWebPage__VisibilityState(0)
	QWebPage__VisibilityStateHidden    = QWebPage__VisibilityState(1)
	QWebPage__VisibilityStatePrerender = QWebPage__VisibilityState(2)
	QWebPage__VisibilityStateUnloaded  = QWebPage__VisibilityState(3)
)

//QWebPage::WebAction
type QWebPage__WebAction int64

const (
	QWebPage__NoWebAction                 = QWebPage__WebAction(-1)
	QWebPage__OpenLink                    = QWebPage__WebAction(0)
	QWebPage__OpenLinkInNewWindow         = QWebPage__WebAction(1)
	QWebPage__OpenFrameInNewWindow        = QWebPage__WebAction(2)
	QWebPage__DownloadLinkToDisk          = QWebPage__WebAction(3)
	QWebPage__CopyLinkToClipboard         = QWebPage__WebAction(4)
	QWebPage__OpenImageInNewWindow        = QWebPage__WebAction(5)
	QWebPage__DownloadImageToDisk         = QWebPage__WebAction(6)
	QWebPage__CopyImageToClipboard        = QWebPage__WebAction(7)
	QWebPage__Back                        = QWebPage__WebAction(8)
	QWebPage__Forward                     = QWebPage__WebAction(9)
	QWebPage__Stop                        = QWebPage__WebAction(10)
	QWebPage__Reload                      = QWebPage__WebAction(11)
	QWebPage__Cut                         = QWebPage__WebAction(12)
	QWebPage__Copy                        = QWebPage__WebAction(13)
	QWebPage__Paste                       = QWebPage__WebAction(14)
	QWebPage__Undo                        = QWebPage__WebAction(15)
	QWebPage__Redo                        = QWebPage__WebAction(16)
	QWebPage__MoveToNextChar              = QWebPage__WebAction(17)
	QWebPage__MoveToPreviousChar          = QWebPage__WebAction(18)
	QWebPage__MoveToNextWord              = QWebPage__WebAction(19)
	QWebPage__MoveToPreviousWord          = QWebPage__WebAction(20)
	QWebPage__MoveToNextLine              = QWebPage__WebAction(21)
	QWebPage__MoveToPreviousLine          = QWebPage__WebAction(22)
	QWebPage__MoveToStartOfLine           = QWebPage__WebAction(23)
	QWebPage__MoveToEndOfLine             = QWebPage__WebAction(24)
	QWebPage__MoveToStartOfBlock          = QWebPage__WebAction(25)
	QWebPage__MoveToEndOfBlock            = QWebPage__WebAction(26)
	QWebPage__MoveToStartOfDocument       = QWebPage__WebAction(27)
	QWebPage__MoveToEndOfDocument         = QWebPage__WebAction(28)
	QWebPage__SelectNextChar              = QWebPage__WebAction(29)
	QWebPage__SelectPreviousChar          = QWebPage__WebAction(30)
	QWebPage__SelectNextWord              = QWebPage__WebAction(31)
	QWebPage__SelectPreviousWord          = QWebPage__WebAction(32)
	QWebPage__SelectNextLine              = QWebPage__WebAction(33)
	QWebPage__SelectPreviousLine          = QWebPage__WebAction(34)
	QWebPage__SelectStartOfLine           = QWebPage__WebAction(35)
	QWebPage__SelectEndOfLine             = QWebPage__WebAction(36)
	QWebPage__SelectStartOfBlock          = QWebPage__WebAction(37)
	QWebPage__SelectEndOfBlock            = QWebPage__WebAction(38)
	QWebPage__SelectStartOfDocument       = QWebPage__WebAction(39)
	QWebPage__SelectEndOfDocument         = QWebPage__WebAction(40)
	QWebPage__DeleteStartOfWord           = QWebPage__WebAction(41)
	QWebPage__DeleteEndOfWord             = QWebPage__WebAction(42)
	QWebPage__SetTextDirectionDefault     = QWebPage__WebAction(43)
	QWebPage__SetTextDirectionLeftToRight = QWebPage__WebAction(44)
	QWebPage__SetTextDirectionRightToLeft = QWebPage__WebAction(45)
	QWebPage__ToggleBold                  = QWebPage__WebAction(46)
	QWebPage__ToggleItalic                = QWebPage__WebAction(47)
	QWebPage__ToggleUnderline             = QWebPage__WebAction(48)
	QWebPage__InspectElement              = QWebPage__WebAction(49)
	QWebPage__InsertParagraphSeparator    = QWebPage__WebAction(50)
	QWebPage__InsertLineSeparator         = QWebPage__WebAction(51)
	QWebPage__SelectAll                   = QWebPage__WebAction(52)
	QWebPage__ReloadAndBypassCache        = QWebPage__WebAction(53)
	QWebPage__PasteAndMatchStyle          = QWebPage__WebAction(54)
	QWebPage__RemoveFormat                = QWebPage__WebAction(55)
	QWebPage__ToggleStrikethrough         = QWebPage__WebAction(56)
	QWebPage__ToggleSubscript             = QWebPage__WebAction(57)
	QWebPage__ToggleSuperscript           = QWebPage__WebAction(58)
	QWebPage__InsertUnorderedList         = QWebPage__WebAction(59)
	QWebPage__InsertOrderedList           = QWebPage__WebAction(60)
	QWebPage__Indent                      = QWebPage__WebAction(61)
	QWebPage__Outdent                     = QWebPage__WebAction(62)
	QWebPage__AlignCenter                 = QWebPage__WebAction(63)
	QWebPage__AlignJustified              = QWebPage__WebAction(64)
	QWebPage__AlignLeft                   = QWebPage__WebAction(65)
	QWebPage__AlignRight                  = QWebPage__WebAction(66)
	QWebPage__StopScheduledPageRefresh    = QWebPage__WebAction(67)
	QWebPage__CopyImageUrlToClipboard     = QWebPage__WebAction(68)
	QWebPage__OpenLinkInThisWindow        = QWebPage__WebAction(69)
	QWebPage__DownloadMediaToDisk         = QWebPage__WebAction(70)
	QWebPage__CopyMediaUrlToClipboard     = QWebPage__WebAction(71)
	QWebPage__ToggleMediaControls         = QWebPage__WebAction(72)
	QWebPage__ToggleMediaLoop             = QWebPage__WebAction(73)
	QWebPage__ToggleMediaPlayPause        = QWebPage__WebAction(74)
	QWebPage__ToggleMediaMute             = QWebPage__WebAction(75)
	QWebPage__ToggleVideoFullscreen       = QWebPage__WebAction(76)
	QWebPage__WebActionCount              = QWebPage__WebAction(77)
)

//QWebPage::WebWindowType
type QWebPage__WebWindowType int64

const (
	QWebPage__WebBrowserWindow = QWebPage__WebWindowType(0)
	QWebPage__WebModalDialog   = QWebPage__WebWindowType(1)
)

type QWebPage struct {
	core.QObject
}

type QWebPage_ITF interface {
	core.QObject_ITF
	QWebPage_PTR() *QWebPage
}

func (p *QWebPage) QWebPage_PTR() *QWebPage {
	return p
}

func (p *QWebPage) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebPage) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebPage(ptr QWebPage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebPage_PTR().Pointer()
	}
	return nil
}

func NewQWebPageFromPointer(ptr unsafe.Pointer) *QWebPage {
	var n = new(QWebPage)
	n.SetPointer(ptr)
	return n
}

func newQWebPageFromPointer(ptr unsafe.Pointer) *QWebPage {
	var n = NewQWebPageFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebPage_") {
		n.SetObjectName("QWebPage_" + qt.Identifier())
	}
	return n
}

//export callbackQWebPage_AcceptNavigationRequest
func callbackQWebPage_AcceptNavigationRequest(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, request unsafe.Pointer, ty C.int) C.int {
	defer qt.Recovering("callback QWebPage::acceptNavigationRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "acceptNavigationRequest"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QWebFrame, *network.QNetworkRequest, QWebPage__NavigationType) bool)(NewQWebFrameFromPointer(frame), network.NewQNetworkRequestFromPointer(request), QWebPage__NavigationType(ty))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).AcceptNavigationRequestDefault(NewQWebFrameFromPointer(frame), network.NewQNetworkRequestFromPointer(request), QWebPage__NavigationType(ty))))
}

func (ptr *QWebPage) ConnectAcceptNavigationRequest(f func(frame *QWebFrame, request *network.QNetworkRequest, ty QWebPage__NavigationType) bool) {
	defer qt.Recovering("connect QWebPage::acceptNavigationRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "acceptNavigationRequest", f)
	}
}

func (ptr *QWebPage) DisconnectAcceptNavigationRequest() {
	defer qt.Recovering("disconnect QWebPage::acceptNavigationRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "acceptNavigationRequest")
	}
}

func (ptr *QWebPage) AcceptNavigationRequest(frame QWebFrame_ITF, request network.QNetworkRequest_ITF, ty QWebPage__NavigationType) bool {
	defer qt.Recovering("QWebPage::acceptNavigationRequest")

	if ptr.Pointer() != nil {
		return C.QWebPage_AcceptNavigationRequest(ptr.Pointer(), PointerFromQWebFrame(frame), network.PointerFromQNetworkRequest(request), C.int(ty)) != 0
	}
	return false
}

func (ptr *QWebPage) AcceptNavigationRequestDefault(frame QWebFrame_ITF, request network.QNetworkRequest_ITF, ty QWebPage__NavigationType) bool {
	defer qt.Recovering("QWebPage::acceptNavigationRequest")

	if ptr.Pointer() != nil {
		return C.QWebPage_AcceptNavigationRequestDefault(ptr.Pointer(), PointerFromQWebFrame(frame), network.PointerFromQNetworkRequest(request), C.int(ty)) != 0
	}
	return false
}

func (ptr *QWebPage) ForwardUnsupportedContent() bool {
	defer qt.Recovering("QWebPage::forwardUnsupportedContent")

	if ptr.Pointer() != nil {
		return C.QWebPage_ForwardUnsupportedContent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) HasSelection() bool {
	defer qt.Recovering("QWebPage::hasSelection")

	if ptr.Pointer() != nil {
		return C.QWebPage_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) IsContentEditable() bool {
	defer qt.Recovering("QWebPage::isContentEditable")

	if ptr.Pointer() != nil {
		return C.QWebPage_IsContentEditable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) IsModified() bool {
	defer qt.Recovering("QWebPage::isModified")

	if ptr.Pointer() != nil {
		return C.QWebPage_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) LinkDelegationPolicy() QWebPage__LinkDelegationPolicy {
	defer qt.Recovering("QWebPage::linkDelegationPolicy")

	if ptr.Pointer() != nil {
		return QWebPage__LinkDelegationPolicy(C.QWebPage_LinkDelegationPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebPage) Palette() *gui.QPalette {
	defer qt.Recovering("QWebPage::palette")

	if ptr.Pointer() != nil {
		return gui.NewQPaletteFromPointer(C.QWebPage_Palette(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) PreferredContentsSize() *core.QSize {
	defer qt.Recovering("QWebPage::preferredContentsSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebPage_PreferredContentsSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) SelectedHtml() string {
	defer qt.Recovering("QWebPage::selectedHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_SelectedHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebPage) SelectedText() string {
	defer qt.Recovering("QWebPage::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebPage) SetActualVisibleContentRect(rect core.QRect_ITF) {
	defer qt.Recovering("QWebPage::setActualVisibleContentRect")

	if ptr.Pointer() != nil {
		C.QWebPage_SetActualVisibleContentRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWebPage) SetContentEditable(editable bool) {
	defer qt.Recovering("QWebPage::setContentEditable")

	if ptr.Pointer() != nil {
		C.QWebPage_SetContentEditable(ptr.Pointer(), C.int(qt.GoBoolToInt(editable)))
	}
}

func (ptr *QWebPage) SetFeaturePermission(frame QWebFrame_ITF, feature QWebPage__Feature, policy QWebPage__PermissionPolicy) {
	defer qt.Recovering("QWebPage::setFeaturePermission")

	if ptr.Pointer() != nil {
		C.QWebPage_SetFeaturePermission(ptr.Pointer(), PointerFromQWebFrame(frame), C.int(feature), C.int(policy))
	}
}

func (ptr *QWebPage) SetForwardUnsupportedContent(forward bool) {
	defer qt.Recovering("QWebPage::setForwardUnsupportedContent")

	if ptr.Pointer() != nil {
		C.QWebPage_SetForwardUnsupportedContent(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QWebPage) SetLinkDelegationPolicy(policy QWebPage__LinkDelegationPolicy) {
	defer qt.Recovering("QWebPage::setLinkDelegationPolicy")

	if ptr.Pointer() != nil {
		C.QWebPage_SetLinkDelegationPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWebPage) SetPalette(palette gui.QPalette_ITF) {
	defer qt.Recovering("QWebPage::setPalette")

	if ptr.Pointer() != nil {
		C.QWebPage_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QWebPage) SetPreferredContentsSize(size core.QSize_ITF) {
	defer qt.Recovering("QWebPage::setPreferredContentsSize")

	if ptr.Pointer() != nil {
		C.QWebPage_SetPreferredContentsSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWebPage) SetViewportSize(size core.QSize_ITF) {
	defer qt.Recovering("QWebPage::setViewportSize")

	if ptr.Pointer() != nil {
		C.QWebPage_SetViewportSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QWebPage) SetVisibilityState(vvi QWebPage__VisibilityState) {
	defer qt.Recovering("QWebPage::setVisibilityState")

	if ptr.Pointer() != nil {
		C.QWebPage_SetVisibilityState(ptr.Pointer(), C.int(vvi))
	}
}

//export callbackQWebPage_ShouldInterruptJavaScript
func callbackQWebPage_ShouldInterruptJavaScript(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebPage::shouldInterruptJavaScript")

	if signal := qt.GetSignal(C.GoString(ptrName), "shouldInterruptJavaScript"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).ShouldInterruptJavaScriptDefault()))
}

func (ptr *QWebPage) ConnectShouldInterruptJavaScript(f func() bool) {
	defer qt.Recovering("connect QWebPage::shouldInterruptJavaScript")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "shouldInterruptJavaScript", f)
	}
}

func (ptr *QWebPage) DisconnectShouldInterruptJavaScript() {
	defer qt.Recovering("disconnect QWebPage::shouldInterruptJavaScript")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "shouldInterruptJavaScript")
	}
}

func (ptr *QWebPage) ShouldInterruptJavaScript() bool {
	defer qt.Recovering("QWebPage::shouldInterruptJavaScript")

	if ptr.Pointer() != nil {
		return C.QWebPage_ShouldInterruptJavaScript(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) ShouldInterruptJavaScriptDefault() bool {
	defer qt.Recovering("QWebPage::shouldInterruptJavaScript")

	if ptr.Pointer() != nil {
		return C.QWebPage_ShouldInterruptJavaScriptDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebPage) ViewportSize() *core.QSize {
	defer qt.Recovering("QWebPage::viewportSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebPage_ViewportSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) VisibilityState() QWebPage__VisibilityState {
	defer qt.Recovering("QWebPage::visibilityState")

	if ptr.Pointer() != nil {
		return QWebPage__VisibilityState(C.QWebPage_VisibilityState(ptr.Pointer()))
	}
	return 0
}

func NewQWebPage(parent core.QObject_ITF) *QWebPage {
	defer qt.Recovering("QWebPage::QWebPage")

	return newQWebPageFromPointer(C.QWebPage_NewQWebPage(core.PointerFromQObject(parent)))
}

func (ptr *QWebPage) Action(action QWebPage__WebAction) *widgets.QAction {
	defer qt.Recovering("QWebPage::action")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QWebPage_Action(ptr.Pointer(), C.int(action)))
	}
	return nil
}

//export callbackQWebPage_ChooseFile
func callbackQWebPage_ChooseFile(ptr unsafe.Pointer, ptrName *C.char, parentFrame unsafe.Pointer, suggestedFile *C.char) *C.char {
	defer qt.Recovering("callback QWebPage::chooseFile")

	if signal := qt.GetSignal(C.GoString(ptrName), "chooseFile"); signal != nil {
		return C.CString(signal.(func(*QWebFrame, string) string)(NewQWebFrameFromPointer(parentFrame), C.GoString(suggestedFile)))
	}

	return C.CString(NewQWebPageFromPointer(ptr).ChooseFileDefault(NewQWebFrameFromPointer(parentFrame), C.GoString(suggestedFile)))
}

func (ptr *QWebPage) ConnectChooseFile(f func(parentFrame *QWebFrame, suggestedFile string) string) {
	defer qt.Recovering("connect QWebPage::chooseFile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "chooseFile", f)
	}
}

func (ptr *QWebPage) DisconnectChooseFile() {
	defer qt.Recovering("disconnect QWebPage::chooseFile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "chooseFile")
	}
}

func (ptr *QWebPage) ChooseFile(parentFrame QWebFrame_ITF, suggestedFile string) string {
	defer qt.Recovering("QWebPage::chooseFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_ChooseFile(ptr.Pointer(), PointerFromQWebFrame(parentFrame), C.CString(suggestedFile)))
	}
	return ""
}

func (ptr *QWebPage) ChooseFileDefault(parentFrame QWebFrame_ITF, suggestedFile string) string {
	defer qt.Recovering("QWebPage::chooseFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_ChooseFileDefault(ptr.Pointer(), PointerFromQWebFrame(parentFrame), C.CString(suggestedFile)))
	}
	return ""
}

//export callbackQWebPage_ContentsChanged
func callbackQWebPage_ContentsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::contentsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectContentsChanged(f func()) {
	defer qt.Recovering("connect QWebPage::contentsChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectContentsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsChanged", f)
	}
}

func (ptr *QWebPage) DisconnectContentsChanged() {
	defer qt.Recovering("disconnect QWebPage::contentsChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectContentsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsChanged")
	}
}

func (ptr *QWebPage) ContentsChanged() {
	defer qt.Recovering("QWebPage::contentsChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_ContentsChanged(ptr.Pointer())
	}
}

//export callbackQWebPage_CreatePlugin
func callbackQWebPage_CreatePlugin(ptr unsafe.Pointer, ptrName *C.char, classid *C.char, url unsafe.Pointer, paramNames *C.char, paramValues *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebPage::createPlugin")

	if signal := qt.GetSignal(C.GoString(ptrName), "createPlugin"); signal != nil {
		return core.PointerFromQObject(signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(C.GoString(classid), core.NewQUrlFromPointer(url), strings.Split(C.GoString(paramNames), "|"), strings.Split(C.GoString(paramValues), "|")))
	}

	return core.PointerFromQObject(NewQWebPageFromPointer(ptr).CreatePluginDefault(C.GoString(classid), core.NewQUrlFromPointer(url), strings.Split(C.GoString(paramNames), "|"), strings.Split(C.GoString(paramValues), "|")))
}

func (ptr *QWebPage) ConnectCreatePlugin(f func(classid string, url *core.QUrl, paramNames []string, paramValues []string) *core.QObject) {
	defer qt.Recovering("connect QWebPage::createPlugin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createPlugin", f)
	}
}

func (ptr *QWebPage) DisconnectCreatePlugin() {
	defer qt.Recovering("disconnect QWebPage::createPlugin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createPlugin")
	}
}

func (ptr *QWebPage) CreatePlugin(classid string, url core.QUrl_ITF, paramNames []string, paramValues []string) *core.QObject {
	defer qt.Recovering("QWebPage::createPlugin")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QWebPage_CreatePlugin(ptr.Pointer(), C.CString(classid), core.PointerFromQUrl(url), C.CString(strings.Join(paramNames, "|")), C.CString(strings.Join(paramValues, "|"))))
	}
	return nil
}

func (ptr *QWebPage) CreatePluginDefault(classid string, url core.QUrl_ITF, paramNames []string, paramValues []string) *core.QObject {
	defer qt.Recovering("QWebPage::createPlugin")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QWebPage_CreatePluginDefault(ptr.Pointer(), C.CString(classid), core.PointerFromQUrl(url), C.CString(strings.Join(paramNames, "|")), C.CString(strings.Join(paramValues, "|"))))
	}
	return nil
}

func (ptr *QWebPage) CreateStandardContextMenu() *widgets.QMenu {
	defer qt.Recovering("QWebPage::createStandardContextMenu")

	if ptr.Pointer() != nil {
		return widgets.NewQMenuFromPointer(C.QWebPage_CreateStandardContextMenu(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_CreateWindow
func callbackQWebPage_CreateWindow(ptr unsafe.Pointer, ptrName *C.char, ty C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebPage::createWindow")

	if signal := qt.GetSignal(C.GoString(ptrName), "createWindow"); signal != nil {
		return PointerFromQWebPage(signal.(func(QWebPage__WebWindowType) *QWebPage)(QWebPage__WebWindowType(ty)))
	}

	return PointerFromQWebPage(NewQWebPageFromPointer(ptr).CreateWindowDefault(QWebPage__WebWindowType(ty)))
}

func (ptr *QWebPage) ConnectCreateWindow(f func(ty QWebPage__WebWindowType) *QWebPage) {
	defer qt.Recovering("connect QWebPage::createWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createWindow", f)
	}
}

func (ptr *QWebPage) DisconnectCreateWindow() {
	defer qt.Recovering("disconnect QWebPage::createWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createWindow")
	}
}

func (ptr *QWebPage) CreateWindow(ty QWebPage__WebWindowType) *QWebPage {
	defer qt.Recovering("QWebPage::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QWebPage_CreateWindow(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebPage) CreateWindowDefault(ty QWebPage__WebWindowType) *QWebPage {
	defer qt.Recovering("QWebPage::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QWebPage_CreateWindowDefault(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebPage) CurrentFrame() *QWebFrame {
	defer qt.Recovering("QWebPage::currentFrame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebPage_CurrentFrame(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_DatabaseQuotaExceeded
func callbackQWebPage_DatabaseQuotaExceeded(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, databaseName *C.char) {
	defer qt.Recovering("callback QWebPage::databaseQuotaExceeded")

	if signal := qt.GetSignal(C.GoString(ptrName), "databaseQuotaExceeded"); signal != nil {
		signal.(func(*QWebFrame, string))(NewQWebFrameFromPointer(frame), C.GoString(databaseName))
	}

}

func (ptr *QWebPage) ConnectDatabaseQuotaExceeded(f func(frame *QWebFrame, databaseName string)) {
	defer qt.Recovering("connect QWebPage::databaseQuotaExceeded")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectDatabaseQuotaExceeded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "databaseQuotaExceeded", f)
	}
}

func (ptr *QWebPage) DisconnectDatabaseQuotaExceeded() {
	defer qt.Recovering("disconnect QWebPage::databaseQuotaExceeded")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectDatabaseQuotaExceeded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "databaseQuotaExceeded")
	}
}

func (ptr *QWebPage) DatabaseQuotaExceeded(frame QWebFrame_ITF, databaseName string) {
	defer qt.Recovering("QWebPage::databaseQuotaExceeded")

	if ptr.Pointer() != nil {
		C.QWebPage_DatabaseQuotaExceeded(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(databaseName))
	}
}

//export callbackQWebPage_DownloadRequested
func callbackQWebPage_DownloadRequested(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::downloadRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "downloadRequested"); signal != nil {
		signal.(func(*network.QNetworkRequest))(network.NewQNetworkRequestFromPointer(request))
	}

}

func (ptr *QWebPage) ConnectDownloadRequested(f func(request *network.QNetworkRequest)) {
	defer qt.Recovering("connect QWebPage::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectDownloadRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "downloadRequested", f)
	}
}

func (ptr *QWebPage) DisconnectDownloadRequested() {
	defer qt.Recovering("disconnect QWebPage::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectDownloadRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "downloadRequested")
	}
}

func (ptr *QWebPage) DownloadRequested(request network.QNetworkRequest_ITF) {
	defer qt.Recovering("QWebPage::downloadRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DownloadRequested(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

//export callbackQWebPage_Event
func callbackQWebPage_Event(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebPage::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(ev))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).EventDefault(core.NewQEventFromPointer(ev))))
}

func (ptr *QWebPage) ConnectEvent(f func(ev *core.QEvent) bool) {
	defer qt.Recovering("connect QWebPage::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebPage) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebPage::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebPage) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPage::event")

	if ptr.Pointer() != nil {
		return C.QWebPage_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QWebPage) EventDefault(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPage::event")

	if ptr.Pointer() != nil {
		return C.QWebPage_EventDefault(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

//export callbackQWebPage_FeaturePermissionRequestCanceled
func callbackQWebPage_FeaturePermissionRequestCanceled(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, feature C.int) {
	defer qt.Recovering("callback QWebPage::featurePermissionRequestCanceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "featurePermissionRequestCanceled"); signal != nil {
		signal.(func(*QWebFrame, QWebPage__Feature))(NewQWebFrameFromPointer(frame), QWebPage__Feature(feature))
	}

}

func (ptr *QWebPage) ConnectFeaturePermissionRequestCanceled(f func(frame *QWebFrame, feature QWebPage__Feature)) {
	defer qt.Recovering("connect QWebPage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featurePermissionRequestCanceled", f)
	}
}

func (ptr *QWebPage) DisconnectFeaturePermissionRequestCanceled() {
	defer qt.Recovering("disconnect QWebPage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFeaturePermissionRequestCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featurePermissionRequestCanceled")
	}
}

func (ptr *QWebPage) FeaturePermissionRequestCanceled(frame QWebFrame_ITF, feature QWebPage__Feature) {
	defer qt.Recovering("QWebPage::featurePermissionRequestCanceled")

	if ptr.Pointer() != nil {
		C.QWebPage_FeaturePermissionRequestCanceled(ptr.Pointer(), PointerFromQWebFrame(frame), C.int(feature))
	}
}

//export callbackQWebPage_FeaturePermissionRequested
func callbackQWebPage_FeaturePermissionRequested(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, feature C.int) {
	defer qt.Recovering("callback QWebPage::featurePermissionRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "featurePermissionRequested"); signal != nil {
		signal.(func(*QWebFrame, QWebPage__Feature))(NewQWebFrameFromPointer(frame), QWebPage__Feature(feature))
	}

}

func (ptr *QWebPage) ConnectFeaturePermissionRequested(f func(frame *QWebFrame, feature QWebPage__Feature)) {
	defer qt.Recovering("connect QWebPage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectFeaturePermissionRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "featurePermissionRequested", f)
	}
}

func (ptr *QWebPage) DisconnectFeaturePermissionRequested() {
	defer qt.Recovering("disconnect QWebPage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFeaturePermissionRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "featurePermissionRequested")
	}
}

func (ptr *QWebPage) FeaturePermissionRequested(frame QWebFrame_ITF, feature QWebPage__Feature) {
	defer qt.Recovering("QWebPage::featurePermissionRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_FeaturePermissionRequested(ptr.Pointer(), PointerFromQWebFrame(frame), C.int(feature))
	}
}

func (ptr *QWebPage) FindText(subString string, options QWebPage__FindFlag) bool {
	defer qt.Recovering("QWebPage::findText")

	if ptr.Pointer() != nil {
		return C.QWebPage_FindText(ptr.Pointer(), C.CString(subString), C.int(options)) != 0
	}
	return false
}

func (ptr *QWebPage) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QWebPage::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebPage_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QWebPage) FrameAt(pos core.QPoint_ITF) *QWebFrame {
	defer qt.Recovering("QWebPage::frameAt")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebPage_FrameAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

//export callbackQWebPage_FrameCreated
func callbackQWebPage_FrameCreated(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::frameCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "frameCreated"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectFrameCreated(f func(frame *QWebFrame)) {
	defer qt.Recovering("connect QWebPage::frameCreated")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectFrameCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameCreated", f)
	}
}

func (ptr *QWebPage) DisconnectFrameCreated() {
	defer qt.Recovering("disconnect QWebPage::frameCreated")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectFrameCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameCreated")
	}
}

func (ptr *QWebPage) FrameCreated(frame QWebFrame_ITF) {
	defer qt.Recovering("QWebPage::frameCreated")

	if ptr.Pointer() != nil {
		C.QWebPage_FrameCreated(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_GeometryChangeRequested
func callbackQWebPage_GeometryChangeRequested(ptr unsafe.Pointer, ptrName *C.char, geom unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::geometryChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChangeRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(geom))
	}

}

func (ptr *QWebPage) ConnectGeometryChangeRequested(f func(geom *core.QRect)) {
	defer qt.Recovering("connect QWebPage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectGeometryChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometryChangeRequested", f)
	}
}

func (ptr *QWebPage) DisconnectGeometryChangeRequested() {
	defer qt.Recovering("disconnect QWebPage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectGeometryChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChangeRequested")
	}
}

func (ptr *QWebPage) GeometryChangeRequested(geom core.QRect_ITF) {
	defer qt.Recovering("QWebPage::geometryChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_GeometryChangeRequested(ptr.Pointer(), core.PointerFromQRect(geom))
	}
}

func (ptr *QWebPage) History() *QWebHistory {
	defer qt.Recovering("QWebPage::history")

	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QWebPage_History(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebPage::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebPage_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

//export callbackQWebPage_JavaScriptAlert
func callbackQWebPage_JavaScriptAlert(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, msg *C.char) {
	defer qt.Recovering("callback QWebPage::javaScriptAlert")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptAlert"); signal != nil {
		signal.(func(*QWebFrame, string))(NewQWebFrameFromPointer(frame), C.GoString(msg))
	} else {
		NewQWebPageFromPointer(ptr).JavaScriptAlertDefault(NewQWebFrameFromPointer(frame), C.GoString(msg))
	}
}

func (ptr *QWebPage) ConnectJavaScriptAlert(f func(frame *QWebFrame, msg string)) {
	defer qt.Recovering("connect QWebPage::javaScriptAlert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptAlert", f)
	}
}

func (ptr *QWebPage) DisconnectJavaScriptAlert() {
	defer qt.Recovering("disconnect QWebPage::javaScriptAlert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptAlert")
	}
}

func (ptr *QWebPage) JavaScriptAlert(frame QWebFrame_ITF, msg string) {
	defer qt.Recovering("QWebPage::javaScriptAlert")

	if ptr.Pointer() != nil {
		C.QWebPage_JavaScriptAlert(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg))
	}
}

func (ptr *QWebPage) JavaScriptAlertDefault(frame QWebFrame_ITF, msg string) {
	defer qt.Recovering("QWebPage::javaScriptAlert")

	if ptr.Pointer() != nil {
		C.QWebPage_JavaScriptAlertDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg))
	}
}

//export callbackQWebPage_JavaScriptConfirm
func callbackQWebPage_JavaScriptConfirm(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, msg *C.char) C.int {
	defer qt.Recovering("callback QWebPage::javaScriptConfirm")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptConfirm"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QWebFrame, string) bool)(NewQWebFrameFromPointer(frame), C.GoString(msg))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).JavaScriptConfirmDefault(NewQWebFrameFromPointer(frame), C.GoString(msg))))
}

func (ptr *QWebPage) ConnectJavaScriptConfirm(f func(frame *QWebFrame, msg string) bool) {
	defer qt.Recovering("connect QWebPage::javaScriptConfirm")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptConfirm", f)
	}
}

func (ptr *QWebPage) DisconnectJavaScriptConfirm() {
	defer qt.Recovering("disconnect QWebPage::javaScriptConfirm")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptConfirm")
	}
}

func (ptr *QWebPage) JavaScriptConfirm(frame QWebFrame_ITF, msg string) bool {
	defer qt.Recovering("QWebPage::javaScriptConfirm")

	if ptr.Pointer() != nil {
		return C.QWebPage_JavaScriptConfirm(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg)) != 0
	}
	return false
}

func (ptr *QWebPage) JavaScriptConfirmDefault(frame QWebFrame_ITF, msg string) bool {
	defer qt.Recovering("QWebPage::javaScriptConfirm")

	if ptr.Pointer() != nil {
		return C.QWebPage_JavaScriptConfirmDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg)) != 0
	}
	return false
}

//export callbackQWebPage_JavaScriptConsoleMessage
func callbackQWebPage_JavaScriptConsoleMessage(ptr unsafe.Pointer, ptrName *C.char, message *C.char, lineNumber C.int, sourceID *C.char) {
	defer qt.Recovering("callback QWebPage::javaScriptConsoleMessage")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptConsoleMessage"); signal != nil {
		signal.(func(string, int, string))(C.GoString(message), int(lineNumber), C.GoString(sourceID))
	} else {
		NewQWebPageFromPointer(ptr).JavaScriptConsoleMessageDefault(C.GoString(message), int(lineNumber), C.GoString(sourceID))
	}
}

func (ptr *QWebPage) ConnectJavaScriptConsoleMessage(f func(message string, lineNumber int, sourceID string)) {
	defer qt.Recovering("connect QWebPage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptConsoleMessage", f)
	}
}

func (ptr *QWebPage) DisconnectJavaScriptConsoleMessage() {
	defer qt.Recovering("disconnect QWebPage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptConsoleMessage")
	}
}

func (ptr *QWebPage) JavaScriptConsoleMessage(message string, lineNumber int, sourceID string) {
	defer qt.Recovering("QWebPage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {
		C.QWebPage_JavaScriptConsoleMessage(ptr.Pointer(), C.CString(message), C.int(lineNumber), C.CString(sourceID))
	}
}

func (ptr *QWebPage) JavaScriptConsoleMessageDefault(message string, lineNumber int, sourceID string) {
	defer qt.Recovering("QWebPage::javaScriptConsoleMessage")

	if ptr.Pointer() != nil {
		C.QWebPage_JavaScriptConsoleMessageDefault(ptr.Pointer(), C.CString(message), C.int(lineNumber), C.CString(sourceID))
	}
}

//export callbackQWebPage_JavaScriptPrompt
func callbackQWebPage_JavaScriptPrompt(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, msg *C.char, defaultValue *C.char, result *C.char) C.int {
	defer qt.Recovering("callback QWebPage::javaScriptPrompt")

	if signal := qt.GetSignal(C.GoString(ptrName), "javaScriptPrompt"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QWebFrame, string, string, string) bool)(NewQWebFrameFromPointer(frame), C.GoString(msg), C.GoString(defaultValue), C.GoString(result))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).JavaScriptPromptDefault(NewQWebFrameFromPointer(frame), C.GoString(msg), C.GoString(defaultValue), C.GoString(result))))
}

func (ptr *QWebPage) ConnectJavaScriptPrompt(f func(frame *QWebFrame, msg string, defaultValue string, result string) bool) {
	defer qt.Recovering("connect QWebPage::javaScriptPrompt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "javaScriptPrompt", f)
	}
}

func (ptr *QWebPage) DisconnectJavaScriptPrompt() {
	defer qt.Recovering("disconnect QWebPage::javaScriptPrompt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "javaScriptPrompt")
	}
}

func (ptr *QWebPage) JavaScriptPrompt(frame QWebFrame_ITF, msg string, defaultValue string, result string) bool {
	defer qt.Recovering("QWebPage::javaScriptPrompt")

	if ptr.Pointer() != nil {
		return C.QWebPage_JavaScriptPrompt(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg), C.CString(defaultValue), C.CString(result)) != 0
	}
	return false
}

func (ptr *QWebPage) JavaScriptPromptDefault(frame QWebFrame_ITF, msg string, defaultValue string, result string) bool {
	defer qt.Recovering("QWebPage::javaScriptPrompt")

	if ptr.Pointer() != nil {
		return C.QWebPage_JavaScriptPromptDefault(ptr.Pointer(), PointerFromQWebFrame(frame), C.CString(msg), C.CString(defaultValue), C.CString(result)) != 0
	}
	return false
}

//export callbackQWebPage_LinkClicked
func callbackQWebPage_LinkClicked(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::linkClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebPage) ConnectLinkClicked(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebPage::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectLinkClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkClicked", f)
	}
}

func (ptr *QWebPage) DisconnectLinkClicked() {
	defer qt.Recovering("disconnect QWebPage::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkClicked")
	}
}

func (ptr *QWebPage) LinkClicked(url core.QUrl_ITF) {
	defer qt.Recovering("QWebPage::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebPage_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebPage_LinkHovered
func callbackQWebPage_LinkHovered(ptr unsafe.Pointer, ptrName *C.char, link *C.char, title *C.char, textContent *C.char) {
	defer qt.Recovering("callback QWebPage::linkHovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkHovered"); signal != nil {
		signal.(func(string, string, string))(C.GoString(link), C.GoString(title), C.GoString(textContent))
	}

}

func (ptr *QWebPage) ConnectLinkHovered(f func(link string, title string, textContent string)) {
	defer qt.Recovering("connect QWebPage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QWebPage) DisconnectLinkHovered() {
	defer qt.Recovering("disconnect QWebPage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

func (ptr *QWebPage) LinkHovered(link string, title string, textContent string) {
	defer qt.Recovering("QWebPage::linkHovered")

	if ptr.Pointer() != nil {
		C.QWebPage_LinkHovered(ptr.Pointer(), C.CString(link), C.CString(title), C.CString(textContent))
	}
}

//export callbackQWebPage_LoadFinished
func callbackQWebPage_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QWebPage::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QWebPage) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QWebPage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QWebPage) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QWebPage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QWebPage) LoadFinished(ok bool) {
	defer qt.Recovering("QWebPage::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebPage_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQWebPage_LoadProgress
func callbackQWebPage_LoadProgress(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QWebPage::loadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadProgress"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QWebPage) ConnectLoadProgress(f func(progress int)) {
	defer qt.Recovering("connect QWebPage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadProgress", f)
	}
}

func (ptr *QWebPage) DisconnectLoadProgress() {
	defer qt.Recovering("disconnect QWebPage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadProgress")
	}
}

func (ptr *QWebPage) LoadProgress(progress int) {
	defer qt.Recovering("QWebPage::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebPage_LoadProgress(ptr.Pointer(), C.int(progress))
	}
}

//export callbackQWebPage_LoadStarted
func callbackQWebPage_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QWebPage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QWebPage) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QWebPage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QWebPage) LoadStarted() {
	defer qt.Recovering("QWebPage::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebPage_LoadStarted(ptr.Pointer())
	}
}

func (ptr *QWebPage) MainFrame() *QWebFrame {
	defer qt.Recovering("QWebPage::mainFrame")

	if ptr.Pointer() != nil {
		return NewQWebFrameFromPointer(C.QWebPage_MainFrame(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_MenuBarVisibilityChangeRequested
func callbackQWebPage_MenuBarVisibilityChangeRequested(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebPage::menuBarVisibilityChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "menuBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectMenuBarVisibilityChangeRequested(f func(visible bool)) {
	defer qt.Recovering("connect QWebPage::menuBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectMenuBarVisibilityChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "menuBarVisibilityChangeRequested", f)
	}
}

func (ptr *QWebPage) DisconnectMenuBarVisibilityChangeRequested() {
	defer qt.Recovering("disconnect QWebPage::menuBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectMenuBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "menuBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) MenuBarVisibilityChangeRequested(visible bool) {
	defer qt.Recovering("QWebPage::menuBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_MenuBarVisibilityChangeRequested(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQWebPage_MicroFocusChanged
func callbackQWebPage_MicroFocusChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::microFocusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "microFocusChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectMicroFocusChanged(f func()) {
	defer qt.Recovering("connect QWebPage::microFocusChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectMicroFocusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "microFocusChanged", f)
	}
}

func (ptr *QWebPage) DisconnectMicroFocusChanged() {
	defer qt.Recovering("disconnect QWebPage::microFocusChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectMicroFocusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "microFocusChanged")
	}
}

func (ptr *QWebPage) MicroFocusChanged() {
	defer qt.Recovering("QWebPage::microFocusChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_MicroFocusChanged(ptr.Pointer())
	}
}

func (ptr *QWebPage) NetworkAccessManager() *network.QNetworkAccessManager {
	defer qt.Recovering("QWebPage::networkAccessManager")

	if ptr.Pointer() != nil {
		return network.NewQNetworkAccessManagerFromPointer(C.QWebPage_NetworkAccessManager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) PluginFactory() *QWebPluginFactory {
	defer qt.Recovering("QWebPage::pluginFactory")

	if ptr.Pointer() != nil {
		return NewQWebPluginFactoryFromPointer(C.QWebPage_PluginFactory(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_PrintRequested
func callbackQWebPage_PrintRequested(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::printRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "printRequested"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectPrintRequested(f func(frame *QWebFrame)) {
	defer qt.Recovering("connect QWebPage::printRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectPrintRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "printRequested", f)
	}
}

func (ptr *QWebPage) DisconnectPrintRequested() {
	defer qt.Recovering("disconnect QWebPage::printRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectPrintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "printRequested")
	}
}

func (ptr *QWebPage) PrintRequested(frame QWebFrame_ITF) {
	defer qt.Recovering("QWebPage::printRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_PrintRequested(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_RepaintRequested
func callbackQWebPage_RepaintRequested(ptr unsafe.Pointer, ptrName *C.char, dirtyRect unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::repaintRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaintRequested"); signal != nil {
		signal.(func(*core.QRect))(core.NewQRectFromPointer(dirtyRect))
	}

}

func (ptr *QWebPage) ConnectRepaintRequested(f func(dirtyRect *core.QRect)) {
	defer qt.Recovering("connect QWebPage::repaintRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectRepaintRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "repaintRequested", f)
	}
}

func (ptr *QWebPage) DisconnectRepaintRequested() {
	defer qt.Recovering("disconnect QWebPage::repaintRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectRepaintRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "repaintRequested")
	}
}

func (ptr *QWebPage) RepaintRequested(dirtyRect core.QRect_ITF) {
	defer qt.Recovering("QWebPage::repaintRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_RepaintRequested(ptr.Pointer(), core.PointerFromQRect(dirtyRect))
	}
}

//export callbackQWebPage_RestoreFrameStateRequested
func callbackQWebPage_RestoreFrameStateRequested(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::restoreFrameStateRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "restoreFrameStateRequested"); signal != nil {
		signal.(func(*QWebFrame))(NewQWebFrameFromPointer(frame))
	}

}

func (ptr *QWebPage) ConnectRestoreFrameStateRequested(f func(frame *QWebFrame)) {
	defer qt.Recovering("connect QWebPage::restoreFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectRestoreFrameStateRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "restoreFrameStateRequested", f)
	}
}

func (ptr *QWebPage) DisconnectRestoreFrameStateRequested() {
	defer qt.Recovering("disconnect QWebPage::restoreFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectRestoreFrameStateRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "restoreFrameStateRequested")
	}
}

func (ptr *QWebPage) RestoreFrameStateRequested(frame QWebFrame_ITF) {
	defer qt.Recovering("QWebPage::restoreFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_RestoreFrameStateRequested(ptr.Pointer(), PointerFromQWebFrame(frame))
	}
}

//export callbackQWebPage_SaveFrameStateRequested
func callbackQWebPage_SaveFrameStateRequested(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, item unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::saveFrameStateRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "saveFrameStateRequested"); signal != nil {
		signal.(func(*QWebFrame, *QWebHistoryItem))(NewQWebFrameFromPointer(frame), NewQWebHistoryItemFromPointer(item))
	}

}

func (ptr *QWebPage) ConnectSaveFrameStateRequested(f func(frame *QWebFrame, item *QWebHistoryItem)) {
	defer qt.Recovering("connect QWebPage::saveFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectSaveFrameStateRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saveFrameStateRequested", f)
	}
}

func (ptr *QWebPage) DisconnectSaveFrameStateRequested() {
	defer qt.Recovering("disconnect QWebPage::saveFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectSaveFrameStateRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saveFrameStateRequested")
	}
}

func (ptr *QWebPage) SaveFrameStateRequested(frame QWebFrame_ITF, item QWebHistoryItem_ITF) {
	defer qt.Recovering("QWebPage::saveFrameStateRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_SaveFrameStateRequested(ptr.Pointer(), PointerFromQWebFrame(frame), PointerFromQWebHistoryItem(item))
	}
}

//export callbackQWebPage_ScrollRequested
func callbackQWebPage_ScrollRequested(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int, rectToScroll unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::scrollRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollRequested"); signal != nil {
		signal.(func(int, int, *core.QRect))(int(dx), int(dy), core.NewQRectFromPointer(rectToScroll))
	}

}

func (ptr *QWebPage) ConnectScrollRequested(f func(dx int, dy int, rectToScroll *core.QRect)) {
	defer qt.Recovering("connect QWebPage::scrollRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectScrollRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "scrollRequested", f)
	}
}

func (ptr *QWebPage) DisconnectScrollRequested() {
	defer qt.Recovering("disconnect QWebPage::scrollRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectScrollRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "scrollRequested")
	}
}

func (ptr *QWebPage) ScrollRequested(dx int, dy int, rectToScroll core.QRect_ITF) {
	defer qt.Recovering("QWebPage::scrollRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ScrollRequested(ptr.Pointer(), C.int(dx), C.int(dy), core.PointerFromQRect(rectToScroll))
	}
}

//export callbackQWebPage_SelectionChanged
func callbackQWebPage_SelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QWebPage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QWebPage) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QWebPage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

func (ptr *QWebPage) SelectionChanged() {
	defer qt.Recovering("QWebPage::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebPage_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebPage) SetNetworkAccessManager(manager network.QNetworkAccessManager_ITF) {
	defer qt.Recovering("QWebPage::setNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QWebPage_SetNetworkAccessManager(ptr.Pointer(), network.PointerFromQNetworkAccessManager(manager))
	}
}

func (ptr *QWebPage) SetPluginFactory(factory QWebPluginFactory_ITF) {
	defer qt.Recovering("QWebPage::setPluginFactory")

	if ptr.Pointer() != nil {
		C.QWebPage_SetPluginFactory(ptr.Pointer(), PointerFromQWebPluginFactory(factory))
	}
}

func (ptr *QWebPage) SetView(view widgets.QWidget_ITF) {
	defer qt.Recovering("QWebPage::setView")

	if ptr.Pointer() != nil {
		C.QWebPage_SetView(ptr.Pointer(), widgets.PointerFromQWidget(view))
	}
}

func (ptr *QWebPage) Settings() *QWebSettings {
	defer qt.Recovering("QWebPage::settings")

	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QWebPage_Settings(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_StatusBarMessage
func callbackQWebPage_StatusBarMessage(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QWebPage::statusBarMessage")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusBarMessage"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QWebPage) ConnectStatusBarMessage(f func(text string)) {
	defer qt.Recovering("connect QWebPage::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectStatusBarMessage(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusBarMessage", f)
	}
}

func (ptr *QWebPage) DisconnectStatusBarMessage() {
	defer qt.Recovering("disconnect QWebPage::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusBarMessage")
	}
}

func (ptr *QWebPage) StatusBarMessage(text string) {
	defer qt.Recovering("QWebPage::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebPage_StatusBarMessage(ptr.Pointer(), C.CString(text))
	}
}

//export callbackQWebPage_StatusBarVisibilityChangeRequested
func callbackQWebPage_StatusBarVisibilityChangeRequested(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebPage::statusBarVisibilityChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectStatusBarVisibilityChangeRequested(f func(visible bool)) {
	defer qt.Recovering("connect QWebPage::statusBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectStatusBarVisibilityChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusBarVisibilityChangeRequested", f)
	}
}

func (ptr *QWebPage) DisconnectStatusBarVisibilityChangeRequested() {
	defer qt.Recovering("disconnect QWebPage::statusBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectStatusBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) StatusBarVisibilityChangeRequested(visible bool) {
	defer qt.Recovering("QWebPage::statusBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_StatusBarVisibilityChangeRequested(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWebPage) SupportedContentTypes() []string {
	defer qt.Recovering("QWebPage::supportedContentTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QWebPage_SupportedContentTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QWebPage) SupportsContentType(mimeType string) bool {
	defer qt.Recovering("QWebPage::supportsContentType")

	if ptr.Pointer() != nil {
		return C.QWebPage_SupportsContentType(ptr.Pointer(), C.CString(mimeType)) != 0
	}
	return false
}

//export callbackQWebPage_SupportsExtension
func callbackQWebPage_SupportsExtension(ptr unsafe.Pointer, ptrName *C.char, extension C.int) C.int {
	defer qt.Recovering("callback QWebPage::supportsExtension")

	if signal := qt.GetSignal(C.GoString(ptrName), "supportsExtension"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(QWebPage__Extension) bool)(QWebPage__Extension(extension))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).SupportsExtensionDefault(QWebPage__Extension(extension))))
}

func (ptr *QWebPage) ConnectSupportsExtension(f func(extension QWebPage__Extension) bool) {
	defer qt.Recovering("connect QWebPage::supportsExtension")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "supportsExtension", f)
	}
}

func (ptr *QWebPage) DisconnectSupportsExtension() {
	defer qt.Recovering("disconnect QWebPage::supportsExtension")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "supportsExtension")
	}
}

func (ptr *QWebPage) SupportsExtension(extension QWebPage__Extension) bool {
	defer qt.Recovering("QWebPage::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QWebPage_SupportsExtension(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QWebPage) SupportsExtensionDefault(extension QWebPage__Extension) bool {
	defer qt.Recovering("QWebPage::supportsExtension")

	if ptr.Pointer() != nil {
		return C.QWebPage_SupportsExtensionDefault(ptr.Pointer(), C.int(extension)) != 0
	}
	return false
}

func (ptr *QWebPage) SwallowContextMenuEvent(event gui.QContextMenuEvent_ITF) bool {
	defer qt.Recovering("QWebPage::swallowContextMenuEvent")

	if ptr.Pointer() != nil {
		return C.QWebPage_SwallowContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event)) != 0
	}
	return false
}

//export callbackQWebPage_ToolBarVisibilityChangeRequested
func callbackQWebPage_ToolBarVisibilityChangeRequested(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebPage::toolBarVisibilityChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "toolBarVisibilityChangeRequested"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	}

}

func (ptr *QWebPage) ConnectToolBarVisibilityChangeRequested(f func(visible bool)) {
	defer qt.Recovering("connect QWebPage::toolBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectToolBarVisibilityChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toolBarVisibilityChangeRequested", f)
	}
}

func (ptr *QWebPage) DisconnectToolBarVisibilityChangeRequested() {
	defer qt.Recovering("disconnect QWebPage::toolBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectToolBarVisibilityChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toolBarVisibilityChangeRequested")
	}
}

func (ptr *QWebPage) ToolBarVisibilityChangeRequested(visible bool) {
	defer qt.Recovering("QWebPage::toolBarVisibilityChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ToolBarVisibilityChangeRequested(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQWebPage_TriggerAction
func callbackQWebPage_TriggerAction(ptr unsafe.Pointer, ptrName *C.char, action C.int, checked C.int) {
	defer qt.Recovering("callback QWebPage::triggerAction")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggerAction"); signal != nil {
		signal.(func(QWebPage__WebAction, bool))(QWebPage__WebAction(action), int(checked) != 0)
	} else {
		NewQWebPageFromPointer(ptr).TriggerActionDefault(QWebPage__WebAction(action), int(checked) != 0)
	}
}

func (ptr *QWebPage) ConnectTriggerAction(f func(action QWebPage__WebAction, checked bool)) {
	defer qt.Recovering("connect QWebPage::triggerAction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "triggerAction", f)
	}
}

func (ptr *QWebPage) DisconnectTriggerAction() {
	defer qt.Recovering("disconnect QWebPage::triggerAction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "triggerAction")
	}
}

func (ptr *QWebPage) TriggerAction(action QWebPage__WebAction, checked bool) {
	defer qt.Recovering("QWebPage::triggerAction")

	if ptr.Pointer() != nil {
		C.QWebPage_TriggerAction(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QWebPage) TriggerActionDefault(action QWebPage__WebAction, checked bool) {
	defer qt.Recovering("QWebPage::triggerAction")

	if ptr.Pointer() != nil {
		C.QWebPage_TriggerActionDefault(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QWebPage) UndoStack() *widgets.QUndoStack {
	defer qt.Recovering("QWebPage::undoStack")

	if ptr.Pointer() != nil {
		return widgets.NewQUndoStackFromPointer(C.QWebPage_UndoStack(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_UnsupportedContent
func callbackQWebPage_UnsupportedContent(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::unsupportedContent")

	if signal := qt.GetSignal(C.GoString(ptrName), "unsupportedContent"); signal != nil {
		signal.(func(*network.QNetworkReply))(network.NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QWebPage) ConnectUnsupportedContent(f func(reply *network.QNetworkReply)) {
	defer qt.Recovering("connect QWebPage::unsupportedContent")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectUnsupportedContent(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "unsupportedContent", f)
	}
}

func (ptr *QWebPage) DisconnectUnsupportedContent() {
	defer qt.Recovering("disconnect QWebPage::unsupportedContent")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectUnsupportedContent(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "unsupportedContent")
	}
}

func (ptr *QWebPage) UnsupportedContent(reply network.QNetworkReply_ITF) {
	defer qt.Recovering("QWebPage::unsupportedContent")

	if ptr.Pointer() != nil {
		C.QWebPage_UnsupportedContent(ptr.Pointer(), network.PointerFromQNetworkReply(reply))
	}
}

func (ptr *QWebPage) UpdatePositionDependentActions(pos core.QPoint_ITF) {
	defer qt.Recovering("QWebPage::updatePositionDependentActions")

	if ptr.Pointer() != nil {
		C.QWebPage_UpdatePositionDependentActions(ptr.Pointer(), core.PointerFromQPoint(pos))
	}
}

//export callbackQWebPage_UserAgentForUrl
func callbackQWebPage_UserAgentForUrl(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QWebPage::userAgentForUrl")

	if signal := qt.GetSignal(C.GoString(ptrName), "userAgentForUrl"); signal != nil {
		return C.CString(signal.(func(*core.QUrl) string)(core.NewQUrlFromPointer(url)))
	}

	return C.CString(NewQWebPageFromPointer(ptr).UserAgentForUrlDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QWebPage) ConnectUserAgentForUrl(f func(url *core.QUrl) string) {
	defer qt.Recovering("connect QWebPage::userAgentForUrl")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "userAgentForUrl", f)
	}
}

func (ptr *QWebPage) DisconnectUserAgentForUrl() {
	defer qt.Recovering("disconnect QWebPage::userAgentForUrl")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "userAgentForUrl")
	}
}

func (ptr *QWebPage) UserAgentForUrl(url core.QUrl_ITF) string {
	defer qt.Recovering("QWebPage::userAgentForUrl")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_UserAgentForUrl(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return ""
}

func (ptr *QWebPage) UserAgentForUrlDefault(url core.QUrl_ITF) string {
	defer qt.Recovering("QWebPage::userAgentForUrl")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebPage_UserAgentForUrlDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return ""
}

func (ptr *QWebPage) View() *widgets.QWidget {
	defer qt.Recovering("QWebPage::view")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QWebPage_View(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebPage_ViewportChangeRequested
func callbackQWebPage_ViewportChangeRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::viewportChangeRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "viewportChangeRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectViewportChangeRequested(f func()) {
	defer qt.Recovering("connect QWebPage::viewportChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectViewportChangeRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "viewportChangeRequested", f)
	}
}

func (ptr *QWebPage) DisconnectViewportChangeRequested() {
	defer qt.Recovering("disconnect QWebPage::viewportChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectViewportChangeRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "viewportChangeRequested")
	}
}

func (ptr *QWebPage) ViewportChangeRequested() {
	defer qt.Recovering("QWebPage::viewportChangeRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ViewportChangeRequested(ptr.Pointer())
	}
}

//export callbackQWebPage_WindowCloseRequested
func callbackQWebPage_WindowCloseRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::windowCloseRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowCloseRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebPage) ConnectWindowCloseRequested(f func()) {
	defer qt.Recovering("connect QWebPage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectWindowCloseRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowCloseRequested", f)
	}
}

func (ptr *QWebPage) DisconnectWindowCloseRequested() {
	defer qt.Recovering("disconnect QWebPage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectWindowCloseRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowCloseRequested")
	}
}

func (ptr *QWebPage) WindowCloseRequested() {
	defer qt.Recovering("QWebPage::windowCloseRequested")

	if ptr.Pointer() != nil {
		C.QWebPage_WindowCloseRequested(ptr.Pointer())
	}
}

func (ptr *QWebPage) DestroyQWebPage() {
	defer qt.Recovering("QWebPage::~QWebPage")

	if ptr.Pointer() != nil {
		C.QWebPage_DestroyQWebPage(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebPage_TimerEvent
func callbackQWebPage_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebPage) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebPage::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebPage) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebPage::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebPage) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebPage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebPage) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebPage::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebPage_ChildEvent
func callbackQWebPage_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebPage) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebPage::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebPage) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebPage::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebPage) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebPage::childEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebPage) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebPage::childEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebPage_ConnectNotify
func callbackQWebPage_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPageFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPage) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebPage::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebPage) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebPage::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebPage) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPage::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebPage) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPage::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebPage_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPage_CustomEvent
func callbackQWebPage_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebPageFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebPage) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebPage::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebPage) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebPage::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebPage) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebPage::customEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebPage) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebPage::customEvent")

	if ptr.Pointer() != nil {
		C.QWebPage_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebPage_DeleteLater
func callbackQWebPage_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPage::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebPage) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebPage::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebPage) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebPage::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebPage) DeleteLater() {
	defer qt.Recovering("QWebPage::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebPage_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebPage) DeleteLaterDefault() {
	defer qt.Recovering("QWebPage::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebPage_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebPage_DisconnectNotify
func callbackQWebPage_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebPage::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPageFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPage) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebPage::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebPage) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebPage::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebPage) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPage::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebPage) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPage::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebPage_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPage_EventFilter
func callbackQWebPage_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebPage::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPageFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebPage) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebPage::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebPage) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebPage::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebPage) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPage::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebPage_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebPage) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPage::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebPage_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebPage_MetaObject
func callbackQWebPage_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebPage::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebPageFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebPage) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebPage::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebPage) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebPage::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebPage) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebPage::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPage_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPage) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebPage::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPage_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebPluginFactory struct {
	core.QObject
}

type QWebPluginFactory_ITF interface {
	core.QObject_ITF
	QWebPluginFactory_PTR() *QWebPluginFactory
}

func (p *QWebPluginFactory) QWebPluginFactory_PTR() *QWebPluginFactory {
	return p
}

func (p *QWebPluginFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebPluginFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebPluginFactory(ptr QWebPluginFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebPluginFactory_PTR().Pointer()
	}
	return nil
}

func NewQWebPluginFactoryFromPointer(ptr unsafe.Pointer) *QWebPluginFactory {
	var n = new(QWebPluginFactory)
	n.SetPointer(ptr)
	return n
}

func newQWebPluginFactoryFromPointer(ptr unsafe.Pointer) *QWebPluginFactory {
	var n = NewQWebPluginFactoryFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebPluginFactory_") {
		n.SetObjectName("QWebPluginFactory_" + qt.Identifier())
	}
	return n
}

//export callbackQWebPluginFactory_Create
func callbackQWebPluginFactory_Create(ptr unsafe.Pointer, ptrName *C.char, mimeType *C.char, url unsafe.Pointer, argumentNames *C.char, argumentValues *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebPluginFactory::create")

	if signal := qt.GetSignal(C.GoString(ptrName), "create"); signal != nil {
		return core.PointerFromQObject(signal.(func(string, *core.QUrl, []string, []string) *core.QObject)(C.GoString(mimeType), core.NewQUrlFromPointer(url), strings.Split(C.GoString(argumentNames), "|"), strings.Split(C.GoString(argumentValues), "|")))
	}

	return core.PointerFromQObject(nil)
}

func (ptr *QWebPluginFactory) ConnectCreate(f func(mimeType string, url *core.QUrl, argumentNames []string, argumentValues []string) *core.QObject) {
	defer qt.Recovering("connect QWebPluginFactory::create")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "create", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectCreate(mimeType string, url core.QUrl_ITF, argumentNames []string, argumentValues []string) {
	defer qt.Recovering("disconnect QWebPluginFactory::create")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "create")
	}
}

func (ptr *QWebPluginFactory) Create(mimeType string, url core.QUrl_ITF, argumentNames []string, argumentValues []string) *core.QObject {
	defer qt.Recovering("QWebPluginFactory::create")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QWebPluginFactory_Create(ptr.Pointer(), C.CString(mimeType), core.PointerFromQUrl(url), C.CString(strings.Join(argumentNames, "|")), C.CString(strings.Join(argumentValues, "|"))))
	}
	return nil
}

//export callbackQWebPluginFactory_RefreshPlugins
func callbackQWebPluginFactory_RefreshPlugins(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPluginFactory::refreshPlugins")

	if signal := qt.GetSignal(C.GoString(ptrName), "refreshPlugins"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPluginFactoryFromPointer(ptr).RefreshPluginsDefault()
	}
}

func (ptr *QWebPluginFactory) ConnectRefreshPlugins(f func()) {
	defer qt.Recovering("connect QWebPluginFactory::refreshPlugins")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "refreshPlugins", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectRefreshPlugins() {
	defer qt.Recovering("disconnect QWebPluginFactory::refreshPlugins")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "refreshPlugins")
	}
}

func (ptr *QWebPluginFactory) RefreshPlugins() {
	defer qt.Recovering("QWebPluginFactory::refreshPlugins")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_RefreshPlugins(ptr.Pointer())
	}
}

func (ptr *QWebPluginFactory) RefreshPluginsDefault() {
	defer qt.Recovering("QWebPluginFactory::refreshPlugins")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_RefreshPluginsDefault(ptr.Pointer())
	}
}

func (ptr *QWebPluginFactory) DestroyQWebPluginFactory() {
	defer qt.Recovering("QWebPluginFactory::~QWebPluginFactory")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DestroyQWebPluginFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebPluginFactory_TimerEvent
func callbackQWebPluginFactory_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPluginFactory::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebPluginFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebPluginFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebPluginFactory) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebPluginFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebPluginFactory_ChildEvent
func callbackQWebPluginFactory_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPluginFactory::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebPluginFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebPluginFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebPluginFactory) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebPluginFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebPluginFactory_ConnectNotify
func callbackQWebPluginFactory_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebPluginFactory::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPluginFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebPluginFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebPluginFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebPluginFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPluginFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebPluginFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPluginFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPluginFactory_CustomEvent
func callbackQWebPluginFactory_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebPluginFactory::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebPluginFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebPluginFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebPluginFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebPluginFactory) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebPluginFactory) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebPluginFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebPluginFactory_DeleteLater
func callbackQWebPluginFactory_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebPluginFactory::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebPluginFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebPluginFactory) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebPluginFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebPluginFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebPluginFactory) DeleteLater() {
	defer qt.Recovering("QWebPluginFactory::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebPluginFactory) DeleteLaterDefault() {
	defer qt.Recovering("QWebPluginFactory::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebPluginFactory_DisconnectNotify
func callbackQWebPluginFactory_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebPluginFactory::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebPluginFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebPluginFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebPluginFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebPluginFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebPluginFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPluginFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebPluginFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebPluginFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebPluginFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebPluginFactory_Event
func callbackQWebPluginFactory_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebPluginFactory::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPluginFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebPluginFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebPluginFactory::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebPluginFactory::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebPluginFactory) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPluginFactory::event")

	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebPluginFactory) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPluginFactory::event")

	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebPluginFactory_EventFilter
func callbackQWebPluginFactory_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebPluginFactory::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebPluginFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebPluginFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebPluginFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebPluginFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebPluginFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPluginFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebPluginFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebPluginFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebPluginFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebPluginFactory_MetaObject
func callbackQWebPluginFactory_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebPluginFactory::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebPluginFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebPluginFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebPluginFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebPluginFactory) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebPluginFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebPluginFactory) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebPluginFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPluginFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebPluginFactory) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebPluginFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebPluginFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QWebSecurityOrigin::SubdomainSetting
type QWebSecurityOrigin__SubdomainSetting int64

const (
	QWebSecurityOrigin__AllowSubdomains    = QWebSecurityOrigin__SubdomainSetting(0)
	QWebSecurityOrigin__DisallowSubdomains = QWebSecurityOrigin__SubdomainSetting(1)
)

type QWebSecurityOrigin struct {
	ptr unsafe.Pointer
}

type QWebSecurityOrigin_ITF interface {
	QWebSecurityOrigin_PTR() *QWebSecurityOrigin
}

func (p *QWebSecurityOrigin) QWebSecurityOrigin_PTR() *QWebSecurityOrigin {
	return p
}

func (p *QWebSecurityOrigin) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebSecurityOrigin) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebSecurityOrigin(ptr QWebSecurityOrigin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSecurityOrigin_PTR().Pointer()
	}
	return nil
}

func NewQWebSecurityOriginFromPointer(ptr unsafe.Pointer) *QWebSecurityOrigin {
	var n = new(QWebSecurityOrigin)
	n.SetPointer(ptr)
	return n
}

func newQWebSecurityOriginFromPointer(ptr unsafe.Pointer) *QWebSecurityOrigin {
	var n = NewQWebSecurityOriginFromPointer(ptr)
	return n
}

func (ptr *QWebSecurityOrigin) SetApplicationCacheQuota(quota int64) {
	defer qt.Recovering("QWebSecurityOrigin::setApplicationCacheQuota")

	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_SetApplicationCacheQuota(ptr.Pointer(), C.longlong(quota))
	}
}

func NewQWebSecurityOrigin(url core.QUrl_ITF) *QWebSecurityOrigin {
	defer qt.Recovering("QWebSecurityOrigin::QWebSecurityOrigin")

	return newQWebSecurityOriginFromPointer(C.QWebSecurityOrigin_NewQWebSecurityOrigin(core.PointerFromQUrl(url)))
}

func NewQWebSecurityOrigin2(other QWebSecurityOrigin_ITF) *QWebSecurityOrigin {
	defer qt.Recovering("QWebSecurityOrigin::QWebSecurityOrigin")

	return newQWebSecurityOriginFromPointer(C.QWebSecurityOrigin_NewQWebSecurityOrigin2(PointerFromQWebSecurityOrigin(other)))
}

func (ptr *QWebSecurityOrigin) AddAccessWhitelistEntry(scheme string, host string, subdomainSetting QWebSecurityOrigin__SubdomainSetting) {
	defer qt.Recovering("QWebSecurityOrigin::addAccessWhitelistEntry")

	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_AddAccessWhitelistEntry(ptr.Pointer(), C.CString(scheme), C.CString(host), C.int(subdomainSetting))
	}
}

func QWebSecurityOrigin_AddLocalScheme(scheme string) {
	defer qt.Recovering("QWebSecurityOrigin::addLocalScheme")

	C.QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(C.CString(scheme))
}

func (ptr *QWebSecurityOrigin) AddLocalScheme(scheme string) {
	defer qt.Recovering("QWebSecurityOrigin::addLocalScheme")

	C.QWebSecurityOrigin_QWebSecurityOrigin_AddLocalScheme(C.CString(scheme))
}

func (ptr *QWebSecurityOrigin) DatabaseQuota() int64 {
	defer qt.Recovering("QWebSecurityOrigin::databaseQuota")

	if ptr.Pointer() != nil {
		return int64(C.QWebSecurityOrigin_DatabaseQuota(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) DatabaseUsage() int64 {
	defer qt.Recovering("QWebSecurityOrigin::databaseUsage")

	if ptr.Pointer() != nil {
		return int64(C.QWebSecurityOrigin_DatabaseUsage(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) Host() string {
	defer qt.Recovering("QWebSecurityOrigin::host")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSecurityOrigin_Host(ptr.Pointer()))
	}
	return ""
}

func QWebSecurityOrigin_LocalSchemes() []string {
	defer qt.Recovering("QWebSecurityOrigin::localSchemes")

	return strings.Split(C.GoString(C.QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()), "|")
}

func (ptr *QWebSecurityOrigin) LocalSchemes() []string {
	defer qt.Recovering("QWebSecurityOrigin::localSchemes")

	return strings.Split(C.GoString(C.QWebSecurityOrigin_QWebSecurityOrigin_LocalSchemes()), "|")
}

func (ptr *QWebSecurityOrigin) Port() int {
	defer qt.Recovering("QWebSecurityOrigin::port")

	if ptr.Pointer() != nil {
		return int(C.QWebSecurityOrigin_Port(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSecurityOrigin) RemoveAccessWhitelistEntry(scheme string, host string, subdomainSetting QWebSecurityOrigin__SubdomainSetting) {
	defer qt.Recovering("QWebSecurityOrigin::removeAccessWhitelistEntry")

	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_RemoveAccessWhitelistEntry(ptr.Pointer(), C.CString(scheme), C.CString(host), C.int(subdomainSetting))
	}
}

func QWebSecurityOrigin_RemoveLocalScheme(scheme string) {
	defer qt.Recovering("QWebSecurityOrigin::removeLocalScheme")

	C.QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(C.CString(scheme))
}

func (ptr *QWebSecurityOrigin) RemoveLocalScheme(scheme string) {
	defer qt.Recovering("QWebSecurityOrigin::removeLocalScheme")

	C.QWebSecurityOrigin_QWebSecurityOrigin_RemoveLocalScheme(C.CString(scheme))
}

func (ptr *QWebSecurityOrigin) Scheme() string {
	defer qt.Recovering("QWebSecurityOrigin::scheme")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSecurityOrigin_Scheme(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSecurityOrigin) SetDatabaseQuota(quota int64) {
	defer qt.Recovering("QWebSecurityOrigin::setDatabaseQuota")

	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_SetDatabaseQuota(ptr.Pointer(), C.longlong(quota))
	}
}

func (ptr *QWebSecurityOrigin) DestroyQWebSecurityOrigin() {
	defer qt.Recovering("QWebSecurityOrigin::~QWebSecurityOrigin")

	if ptr.Pointer() != nil {
		C.QWebSecurityOrigin_DestroyQWebSecurityOrigin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebSettings::FontFamily
type QWebSettings__FontFamily int64

const (
	QWebSettings__StandardFont  = QWebSettings__FontFamily(0)
	QWebSettings__FixedFont     = QWebSettings__FontFamily(1)
	QWebSettings__SerifFont     = QWebSettings__FontFamily(2)
	QWebSettings__SansSerifFont = QWebSettings__FontFamily(3)
	QWebSettings__CursiveFont   = QWebSettings__FontFamily(4)
	QWebSettings__FantasyFont   = QWebSettings__FontFamily(5)
)

//QWebSettings::FontSize
type QWebSettings__FontSize int64

const (
	QWebSettings__MinimumFontSize        = QWebSettings__FontSize(0)
	QWebSettings__MinimumLogicalFontSize = QWebSettings__FontSize(1)
	QWebSettings__DefaultFontSize        = QWebSettings__FontSize(2)
	QWebSettings__DefaultFixedFontSize   = QWebSettings__FontSize(3)
)

//QWebSettings::ThirdPartyCookiePolicy
type QWebSettings__ThirdPartyCookiePolicy int64

const (
	QWebSettings__AlwaysAllowThirdPartyCookies       = QWebSettings__ThirdPartyCookiePolicy(0)
	QWebSettings__AlwaysBlockThirdPartyCookies       = QWebSettings__ThirdPartyCookiePolicy(1)
	QWebSettings__AllowThirdPartyWithExistingCookies = QWebSettings__ThirdPartyCookiePolicy(2)
)

//QWebSettings::WebAttribute
type QWebSettings__WebAttribute int64

var (
	QWebSettings__AutoLoadImages                    = QWebSettings__WebAttribute(0)
	QWebSettings__JavascriptEnabled                 = QWebSettings__WebAttribute(1)
	QWebSettings__JavaEnabled                       = QWebSettings__WebAttribute(2)
	QWebSettings__PluginsEnabled                    = QWebSettings__WebAttribute(3)
	QWebSettings__PrivateBrowsingEnabled            = QWebSettings__WebAttribute(4)
	QWebSettings__JavascriptCanOpenWindows          = QWebSettings__WebAttribute(5)
	QWebSettings__JavascriptCanAccessClipboard      = QWebSettings__WebAttribute(6)
	QWebSettings__DeveloperExtrasEnabled            = QWebSettings__WebAttribute(7)
	QWebSettings__LinksIncludedInFocusChain         = QWebSettings__WebAttribute(8)
	QWebSettings__ZoomTextOnly                      = QWebSettings__WebAttribute(9)
	QWebSettings__PrintElementBackgrounds           = QWebSettings__WebAttribute(10)
	QWebSettings__OfflineStorageDatabaseEnabled     = QWebSettings__WebAttribute(11)
	QWebSettings__OfflineWebApplicationCacheEnabled = QWebSettings__WebAttribute(12)
	QWebSettings__LocalStorageEnabled               = QWebSettings__WebAttribute(13)
	QWebSettings__LocalStorageDatabaseEnabled       = QWebSettings__WebAttribute(QWebSettings__LocalStorageEnabled)
	QWebSettings__LocalContentCanAccessRemoteUrls   = QWebSettings__WebAttribute(C.QWebSettings_LocalContentCanAccessRemoteUrls_Type())
	QWebSettings__DnsPrefetchEnabled                = QWebSettings__WebAttribute(C.QWebSettings_DnsPrefetchEnabled_Type())
	QWebSettings__XSSAuditingEnabled                = QWebSettings__WebAttribute(C.QWebSettings_XSSAuditingEnabled_Type())
	QWebSettings__AcceleratedCompositingEnabled     = QWebSettings__WebAttribute(C.QWebSettings_AcceleratedCompositingEnabled_Type())
	QWebSettings__SpatialNavigationEnabled          = QWebSettings__WebAttribute(C.QWebSettings_SpatialNavigationEnabled_Type())
	QWebSettings__LocalContentCanAccessFileUrls     = QWebSettings__WebAttribute(C.QWebSettings_LocalContentCanAccessFileUrls_Type())
	QWebSettings__TiledBackingStoreEnabled          = QWebSettings__WebAttribute(C.QWebSettings_TiledBackingStoreEnabled_Type())
	QWebSettings__FrameFlatteningEnabled            = QWebSettings__WebAttribute(C.QWebSettings_FrameFlatteningEnabled_Type())
	QWebSettings__SiteSpecificQuirksEnabled         = QWebSettings__WebAttribute(C.QWebSettings_SiteSpecificQuirksEnabled_Type())
	QWebSettings__JavascriptCanCloseWindows         = QWebSettings__WebAttribute(C.QWebSettings_JavascriptCanCloseWindows_Type())
	QWebSettings__WebGLEnabled                      = QWebSettings__WebAttribute(C.QWebSettings_WebGLEnabled_Type())
	QWebSettings__CSSRegionsEnabled                 = QWebSettings__WebAttribute(C.QWebSettings_CSSRegionsEnabled_Type())
	QWebSettings__HyperlinkAuditingEnabled          = QWebSettings__WebAttribute(C.QWebSettings_HyperlinkAuditingEnabled_Type())
	QWebSettings__CSSGridLayoutEnabled              = QWebSettings__WebAttribute(C.QWebSettings_CSSGridLayoutEnabled_Type())
	QWebSettings__ScrollAnimatorEnabled             = QWebSettings__WebAttribute(C.QWebSettings_ScrollAnimatorEnabled_Type())
	QWebSettings__CaretBrowsingEnabled              = QWebSettings__WebAttribute(C.QWebSettings_CaretBrowsingEnabled_Type())
	QWebSettings__NotificationsEnabled              = QWebSettings__WebAttribute(C.QWebSettings_NotificationsEnabled_Type())
	QWebSettings__WebAudioEnabled                   = QWebSettings__WebAttribute(C.QWebSettings_WebAudioEnabled_Type())
	QWebSettings__Accelerated2dCanvasEnabled        = QWebSettings__WebAttribute(C.QWebSettings_Accelerated2dCanvasEnabled_Type())
)

//QWebSettings::WebGraphic
type QWebSettings__WebGraphic int64

const (
	QWebSettings__MissingImageGraphic              = QWebSettings__WebGraphic(0)
	QWebSettings__MissingPluginGraphic             = QWebSettings__WebGraphic(1)
	QWebSettings__DefaultFrameIconGraphic          = QWebSettings__WebGraphic(2)
	QWebSettings__TextAreaSizeGripCornerGraphic    = QWebSettings__WebGraphic(3)
	QWebSettings__DeleteButtonGraphic              = QWebSettings__WebGraphic(4)
	QWebSettings__InputSpeechButtonGraphic         = QWebSettings__WebGraphic(5)
	QWebSettings__SearchCancelButtonGraphic        = QWebSettings__WebGraphic(6)
	QWebSettings__SearchCancelButtonPressedGraphic = QWebSettings__WebGraphic(7)
)

type QWebSettings struct {
	ptr unsafe.Pointer
}

type QWebSettings_ITF interface {
	QWebSettings_PTR() *QWebSettings
}

func (p *QWebSettings) QWebSettings_PTR() *QWebSettings {
	return p
}

func (p *QWebSettings) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebSettings) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebSettings(ptr QWebSettings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSettings_PTR().Pointer()
	}
	return nil
}

func NewQWebSettingsFromPointer(ptr unsafe.Pointer) *QWebSettings {
	var n = new(QWebSettings)
	n.SetPointer(ptr)
	return n
}

func newQWebSettingsFromPointer(ptr unsafe.Pointer) *QWebSettings {
	var n = NewQWebSettingsFromPointer(ptr)
	return n
}

func (ptr *QWebSettings) ResetAttribute(attribute QWebSettings__WebAttribute) {
	defer qt.Recovering("QWebSettings::resetAttribute")

	if ptr.Pointer() != nil {
		C.QWebSettings_ResetAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QWebSettings) SetAttribute(attribute QWebSettings__WebAttribute, on bool) {
	defer qt.Recovering("QWebSettings::setAttribute")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWebSettings) TestAttribute(attribute QWebSettings__WebAttribute) bool {
	defer qt.Recovering("QWebSettings::testAttribute")

	if ptr.Pointer() != nil {
		return C.QWebSettings_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func QWebSettings_ClearIconDatabase() {
	defer qt.Recovering("QWebSettings::clearIconDatabase")

	C.QWebSettings_QWebSettings_ClearIconDatabase()
}

func (ptr *QWebSettings) ClearIconDatabase() {
	defer qt.Recovering("QWebSettings::clearIconDatabase")

	C.QWebSettings_QWebSettings_ClearIconDatabase()
}

func QWebSettings_ClearMemoryCaches() {
	defer qt.Recovering("QWebSettings::clearMemoryCaches")

	C.QWebSettings_QWebSettings_ClearMemoryCaches()
}

func (ptr *QWebSettings) ClearMemoryCaches() {
	defer qt.Recovering("QWebSettings::clearMemoryCaches")

	C.QWebSettings_QWebSettings_ClearMemoryCaches()
}

func (ptr *QWebSettings) CssMediaType() string {
	defer qt.Recovering("QWebSettings::cssMediaType")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSettings_CssMediaType(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSettings) DefaultTextEncoding() string {
	defer qt.Recovering("QWebSettings::defaultTextEncoding")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSettings_DefaultTextEncoding(ptr.Pointer()))
	}
	return ""
}

func QWebSettings_EnablePersistentStorage(path string) {
	defer qt.Recovering("QWebSettings::enablePersistentStorage")

	C.QWebSettings_QWebSettings_EnablePersistentStorage(C.CString(path))
}

func (ptr *QWebSettings) EnablePersistentStorage(path string) {
	defer qt.Recovering("QWebSettings::enablePersistentStorage")

	C.QWebSettings_QWebSettings_EnablePersistentStorage(C.CString(path))
}

func (ptr *QWebSettings) FontFamily(which QWebSettings__FontFamily) string {
	defer qt.Recovering("QWebSettings::fontFamily")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSettings_FontFamily(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWebSettings) FontSize(ty QWebSettings__FontSize) int {
	defer qt.Recovering("QWebSettings::fontSize")

	if ptr.Pointer() != nil {
		return int(C.QWebSettings_FontSize(ptr.Pointer(), C.int(ty)))
	}
	return 0
}

func QWebSettings_GlobalSettings() *QWebSettings {
	defer qt.Recovering("QWebSettings::globalSettings")

	return NewQWebSettingsFromPointer(C.QWebSettings_QWebSettings_GlobalSettings())
}

func (ptr *QWebSettings) GlobalSettings() *QWebSettings {
	defer qt.Recovering("QWebSettings::globalSettings")

	return NewQWebSettingsFromPointer(C.QWebSettings_QWebSettings_GlobalSettings())
}

func QWebSettings_IconDatabasePath() string {
	defer qt.Recovering("QWebSettings::iconDatabasePath")

	return C.GoString(C.QWebSettings_QWebSettings_IconDatabasePath())
}

func (ptr *QWebSettings) IconDatabasePath() string {
	defer qt.Recovering("QWebSettings::iconDatabasePath")

	return C.GoString(C.QWebSettings_QWebSettings_IconDatabasePath())
}

func QWebSettings_IconForUrl(url core.QUrl_ITF) *gui.QIcon {
	defer qt.Recovering("QWebSettings::iconForUrl")

	return gui.NewQIconFromPointer(C.QWebSettings_QWebSettings_IconForUrl(core.PointerFromQUrl(url)))
}

func (ptr *QWebSettings) IconForUrl(url core.QUrl_ITF) *gui.QIcon {
	defer qt.Recovering("QWebSettings::iconForUrl")

	return gui.NewQIconFromPointer(C.QWebSettings_QWebSettings_IconForUrl(core.PointerFromQUrl(url)))
}

func (ptr *QWebSettings) LocalStoragePath() string {
	defer qt.Recovering("QWebSettings::localStoragePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSettings_LocalStoragePath(ptr.Pointer()))
	}
	return ""
}

func QWebSettings_MaximumPagesInCache() int {
	defer qt.Recovering("QWebSettings::maximumPagesInCache")

	return int(C.QWebSettings_QWebSettings_MaximumPagesInCache())
}

func (ptr *QWebSettings) MaximumPagesInCache() int {
	defer qt.Recovering("QWebSettings::maximumPagesInCache")

	return int(C.QWebSettings_QWebSettings_MaximumPagesInCache())
}

func QWebSettings_OfflineStorageDefaultQuota() int64 {
	defer qt.Recovering("QWebSettings::offlineStorageDefaultQuota")

	return int64(C.QWebSettings_QWebSettings_OfflineStorageDefaultQuota())
}

func (ptr *QWebSettings) OfflineStorageDefaultQuota() int64 {
	defer qt.Recovering("QWebSettings::offlineStorageDefaultQuota")

	return int64(C.QWebSettings_QWebSettings_OfflineStorageDefaultQuota())
}

func QWebSettings_OfflineStoragePath() string {
	defer qt.Recovering("QWebSettings::offlineStoragePath")

	return C.GoString(C.QWebSettings_QWebSettings_OfflineStoragePath())
}

func (ptr *QWebSettings) OfflineStoragePath() string {
	defer qt.Recovering("QWebSettings::offlineStoragePath")

	return C.GoString(C.QWebSettings_QWebSettings_OfflineStoragePath())
}

func QWebSettings_OfflineWebApplicationCachePath() string {
	defer qt.Recovering("QWebSettings::offlineWebApplicationCachePath")

	return C.GoString(C.QWebSettings_QWebSettings_OfflineWebApplicationCachePath())
}

func (ptr *QWebSettings) OfflineWebApplicationCachePath() string {
	defer qt.Recovering("QWebSettings::offlineWebApplicationCachePath")

	return C.GoString(C.QWebSettings_QWebSettings_OfflineWebApplicationCachePath())
}

func QWebSettings_OfflineWebApplicationCacheQuota() int64 {
	defer qt.Recovering("QWebSettings::offlineWebApplicationCacheQuota")

	return int64(C.QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota())
}

func (ptr *QWebSettings) OfflineWebApplicationCacheQuota() int64 {
	defer qt.Recovering("QWebSettings::offlineWebApplicationCacheQuota")

	return int64(C.QWebSettings_QWebSettings_OfflineWebApplicationCacheQuota())
}

func (ptr *QWebSettings) ResetFontFamily(which QWebSettings__FontFamily) {
	defer qt.Recovering("QWebSettings::resetFontFamily")

	if ptr.Pointer() != nil {
		C.QWebSettings_ResetFontFamily(ptr.Pointer(), C.int(which))
	}
}

func (ptr *QWebSettings) ResetFontSize(ty QWebSettings__FontSize) {
	defer qt.Recovering("QWebSettings::resetFontSize")

	if ptr.Pointer() != nil {
		C.QWebSettings_ResetFontSize(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QWebSettings) SetCSSMediaType(ty string) {
	defer qt.Recovering("QWebSettings::setCSSMediaType")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetCSSMediaType(ptr.Pointer(), C.CString(ty))
	}
}

func (ptr *QWebSettings) SetDefaultTextEncoding(encoding string) {
	defer qt.Recovering("QWebSettings::setDefaultTextEncoding")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetDefaultTextEncoding(ptr.Pointer(), C.CString(encoding))
	}
}

func (ptr *QWebSettings) SetFontFamily(which QWebSettings__FontFamily, family string) {
	defer qt.Recovering("QWebSettings::setFontFamily")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetFontFamily(ptr.Pointer(), C.int(which), C.CString(family))
	}
}

func (ptr *QWebSettings) SetFontSize(ty QWebSettings__FontSize, size int) {
	defer qt.Recovering("QWebSettings::setFontSize")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetFontSize(ptr.Pointer(), C.int(ty), C.int(size))
	}
}

func QWebSettings_SetIconDatabasePath(path string) {
	defer qt.Recovering("QWebSettings::setIconDatabasePath")

	C.QWebSettings_QWebSettings_SetIconDatabasePath(C.CString(path))
}

func (ptr *QWebSettings) SetIconDatabasePath(path string) {
	defer qt.Recovering("QWebSettings::setIconDatabasePath")

	C.QWebSettings_QWebSettings_SetIconDatabasePath(C.CString(path))
}

func (ptr *QWebSettings) SetLocalStoragePath(path string) {
	defer qt.Recovering("QWebSettings::setLocalStoragePath")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetLocalStoragePath(ptr.Pointer(), C.CString(path))
	}
}

func QWebSettings_SetMaximumPagesInCache(pages int) {
	defer qt.Recovering("QWebSettings::setMaximumPagesInCache")

	C.QWebSettings_QWebSettings_SetMaximumPagesInCache(C.int(pages))
}

func (ptr *QWebSettings) SetMaximumPagesInCache(pages int) {
	defer qt.Recovering("QWebSettings::setMaximumPagesInCache")

	C.QWebSettings_QWebSettings_SetMaximumPagesInCache(C.int(pages))
}

func QWebSettings_SetObjectCacheCapacities(cacheMinDeadCapacity int, cacheMaxDead int, totalCapacity int) {
	defer qt.Recovering("QWebSettings::setObjectCacheCapacities")

	C.QWebSettings_QWebSettings_SetObjectCacheCapacities(C.int(cacheMinDeadCapacity), C.int(cacheMaxDead), C.int(totalCapacity))
}

func (ptr *QWebSettings) SetObjectCacheCapacities(cacheMinDeadCapacity int, cacheMaxDead int, totalCapacity int) {
	defer qt.Recovering("QWebSettings::setObjectCacheCapacities")

	C.QWebSettings_QWebSettings_SetObjectCacheCapacities(C.int(cacheMinDeadCapacity), C.int(cacheMaxDead), C.int(totalCapacity))
}

func QWebSettings_SetOfflineStorageDefaultQuota(maximumSize int64) {
	defer qt.Recovering("QWebSettings::setOfflineStorageDefaultQuota")

	C.QWebSettings_QWebSettings_SetOfflineStorageDefaultQuota(C.longlong(maximumSize))
}

func (ptr *QWebSettings) SetOfflineStorageDefaultQuota(maximumSize int64) {
	defer qt.Recovering("QWebSettings::setOfflineStorageDefaultQuota")

	C.QWebSettings_QWebSettings_SetOfflineStorageDefaultQuota(C.longlong(maximumSize))
}

func QWebSettings_SetOfflineStoragePath(path string) {
	defer qt.Recovering("QWebSettings::setOfflineStoragePath")

	C.QWebSettings_QWebSettings_SetOfflineStoragePath(C.CString(path))
}

func (ptr *QWebSettings) SetOfflineStoragePath(path string) {
	defer qt.Recovering("QWebSettings::setOfflineStoragePath")

	C.QWebSettings_QWebSettings_SetOfflineStoragePath(C.CString(path))
}

func QWebSettings_SetOfflineWebApplicationCachePath(path string) {
	defer qt.Recovering("QWebSettings::setOfflineWebApplicationCachePath")

	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(C.CString(path))
}

func (ptr *QWebSettings) SetOfflineWebApplicationCachePath(path string) {
	defer qt.Recovering("QWebSettings::setOfflineWebApplicationCachePath")

	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCachePath(C.CString(path))
}

func QWebSettings_SetOfflineWebApplicationCacheQuota(maximumSize int64) {
	defer qt.Recovering("QWebSettings::setOfflineWebApplicationCacheQuota")

	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(C.longlong(maximumSize))
}

func (ptr *QWebSettings) SetOfflineWebApplicationCacheQuota(maximumSize int64) {
	defer qt.Recovering("QWebSettings::setOfflineWebApplicationCacheQuota")

	C.QWebSettings_QWebSettings_SetOfflineWebApplicationCacheQuota(C.longlong(maximumSize))
}

func (ptr *QWebSettings) SetThirdPartyCookiePolicy(policy QWebSettings__ThirdPartyCookiePolicy) {
	defer qt.Recovering("QWebSettings::setThirdPartyCookiePolicy")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetThirdPartyCookiePolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWebSettings) SetUserStyleSheetUrl(location core.QUrl_ITF) {
	defer qt.Recovering("QWebSettings::setUserStyleSheetUrl")

	if ptr.Pointer() != nil {
		C.QWebSettings_SetUserStyleSheetUrl(ptr.Pointer(), core.PointerFromQUrl(location))
	}
}

func QWebSettings_SetWebGraphic(ty QWebSettings__WebGraphic, graphic gui.QPixmap_ITF) {
	defer qt.Recovering("QWebSettings::setWebGraphic")

	C.QWebSettings_QWebSettings_SetWebGraphic(C.int(ty), gui.PointerFromQPixmap(graphic))
}

func (ptr *QWebSettings) SetWebGraphic(ty QWebSettings__WebGraphic, graphic gui.QPixmap_ITF) {
	defer qt.Recovering("QWebSettings::setWebGraphic")

	C.QWebSettings_QWebSettings_SetWebGraphic(C.int(ty), gui.PointerFromQPixmap(graphic))
}

func (ptr *QWebSettings) ThirdPartyCookiePolicy() QWebSettings__ThirdPartyCookiePolicy {
	defer qt.Recovering("QWebSettings::thirdPartyCookiePolicy")

	if ptr.Pointer() != nil {
		return QWebSettings__ThirdPartyCookiePolicy(C.QWebSettings_ThirdPartyCookiePolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSettings) UserStyleSheetUrl() *core.QUrl {
	defer qt.Recovering("QWebSettings::userStyleSheetUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebSettings_UserStyleSheetUrl(ptr.Pointer()))
	}
	return nil
}

func QWebSettings_WebGraphic(ty QWebSettings__WebGraphic) *gui.QPixmap {
	defer qt.Recovering("QWebSettings::webGraphic")

	return gui.NewQPixmapFromPointer(C.QWebSettings_QWebSettings_WebGraphic(C.int(ty)))
}

func (ptr *QWebSettings) WebGraphic(ty QWebSettings__WebGraphic) *gui.QPixmap {
	defer qt.Recovering("QWebSettings::webGraphic")

	return gui.NewQPixmapFromPointer(C.QWebSettings_QWebSettings_WebGraphic(C.int(ty)))
}

type QWebView struct {
	widgets.QWidget
}

type QWebView_ITF interface {
	widgets.QWidget_ITF
	QWebView_PTR() *QWebView
}

func (p *QWebView) QWebView_PTR() *QWebView {
	return p
}

func (p *QWebView) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QWebView) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebView(ptr QWebView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebView_PTR().Pointer()
	}
	return nil
}

func NewQWebViewFromPointer(ptr unsafe.Pointer) *QWebView {
	var n = new(QWebView)
	n.SetPointer(ptr)
	return n
}

func newQWebViewFromPointer(ptr unsafe.Pointer) *QWebView {
	var n = NewQWebViewFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebView_") {
		n.SetObjectName("QWebView_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebView) HasSelection() bool {
	defer qt.Recovering("QWebView::hasSelection")

	if ptr.Pointer() != nil {
		return C.QWebView_HasSelection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) Icon() *gui.QIcon {
	defer qt.Recovering("QWebView::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QWebView_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) IsModified() bool {
	defer qt.Recovering("QWebView::isModified")

	if ptr.Pointer() != nil {
		return C.QWebView_IsModified(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) Load2(request network.QNetworkRequest_ITF, operation network.QNetworkAccessManager__Operation, body string) {
	defer qt.Recovering("QWebView::load")

	if ptr.Pointer() != nil {
		C.QWebView_Load2(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.int(operation), C.CString(body))
	}
}

func (ptr *QWebView) SelectedHtml() string {
	defer qt.Recovering("QWebView::selectedHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebView_SelectedHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) SelectedText() string {
	defer qt.Recovering("QWebView::selectedText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebView_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QWebView::setUrl")

	if ptr.Pointer() != nil {
		C.QWebView_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebView) SetZoomFactor(factor float64) {
	defer qt.Recovering("QWebView::setZoomFactor")

	if ptr.Pointer() != nil {
		C.QWebView_SetZoomFactor(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebView) Title() string {
	defer qt.Recovering("QWebView::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebView_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebView) Url() *core.QUrl {
	defer qt.Recovering("QWebView::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebView_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) ZoomFactor() float64 {
	defer qt.Recovering("QWebView::zoomFactor")

	if ptr.Pointer() != nil {
		return float64(C.QWebView_ZoomFactor(ptr.Pointer()))
	}
	return 0
}

func NewQWebView(parent widgets.QWidget_ITF) *QWebView {
	defer qt.Recovering("QWebView::QWebView")

	return newQWebViewFromPointer(C.QWebView_NewQWebView(widgets.PointerFromQWidget(parent)))
}

//export callbackQWebView_Back
func callbackQWebView_Back(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::back")

	if signal := qt.GetSignal(C.GoString(ptrName), "back"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectBack(f func()) {
	defer qt.Recovering("connect QWebView::back")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "back", f)
	}
}

func (ptr *QWebView) DisconnectBack() {
	defer qt.Recovering("disconnect QWebView::back")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "back")
	}
}

func (ptr *QWebView) Back() {
	defer qt.Recovering("QWebView::back")

	if ptr.Pointer() != nil {
		C.QWebView_Back(ptr.Pointer())
	}
}

//export callbackQWebView_ChangeEvent
func callbackQWebView_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QWebView) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QWebView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QWebView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QWebView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

func (ptr *QWebView) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QWebView::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QWebView) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QWebView::changeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

//export callbackQWebView_ContextMenuEvent
func callbackQWebView_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectContextMenuEvent(f func(ev *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWebView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWebView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWebView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

func (ptr *QWebView) ContextMenuEvent(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

func (ptr *QWebView) ContextMenuEventDefault(ev gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QWebView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(ev))
	}
}

//export callbackQWebView_CreateWindow
func callbackQWebView_CreateWindow(ptr unsafe.Pointer, ptrName *C.char, ty C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::createWindow")

	if signal := qt.GetSignal(C.GoString(ptrName), "createWindow"); signal != nil {
		return PointerFromQWebView(signal.(func(QWebPage__WebWindowType) *QWebView)(QWebPage__WebWindowType(ty)))
	}

	return PointerFromQWebView(NewQWebViewFromPointer(ptr).CreateWindowDefault(QWebPage__WebWindowType(ty)))
}

func (ptr *QWebView) ConnectCreateWindow(f func(ty QWebPage__WebWindowType) *QWebView) {
	defer qt.Recovering("connect QWebView::createWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createWindow", f)
	}
}

func (ptr *QWebView) DisconnectCreateWindow() {
	defer qt.Recovering("disconnect QWebView::createWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createWindow")
	}
}

func (ptr *QWebView) CreateWindow(ty QWebPage__WebWindowType) *QWebView {
	defer qt.Recovering("QWebView::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebViewFromPointer(C.QWebView_CreateWindow(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QWebView) CreateWindowDefault(ty QWebPage__WebWindowType) *QWebView {
	defer qt.Recovering("QWebView::createWindow")

	if ptr.Pointer() != nil {
		return NewQWebViewFromPointer(C.QWebView_CreateWindowDefault(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

//export callbackQWebView_DragEnterEvent
func callbackQWebView_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectDragEnterEvent(f func(ev *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWebView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWebView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWebView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

func (ptr *QWebView) DragEnterEvent(ev gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(ev))
	}
}

func (ptr *QWebView) DragEnterEventDefault(ev gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QWebView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(ev))
	}
}

//export callbackQWebView_DragLeaveEvent
func callbackQWebView_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectDragLeaveEvent(f func(ev *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWebView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

func (ptr *QWebView) DragLeaveEvent(ev gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(ev))
	}
}

func (ptr *QWebView) DragLeaveEventDefault(ev gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QWebView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(ev))
	}
}

//export callbackQWebView_DragMoveEvent
func callbackQWebView_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectDragMoveEvent(f func(ev *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWebView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWebView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWebView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

func (ptr *QWebView) DragMoveEvent(ev gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(ev))
	}
}

func (ptr *QWebView) DragMoveEventDefault(ev gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QWebView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(ev))
	}
}

//export callbackQWebView_DropEvent
func callbackQWebView_DropEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectDropEvent(f func(ev *gui.QDropEvent)) {
	defer qt.Recovering("connect QWebView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWebView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWebView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

func (ptr *QWebView) DropEvent(ev gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebView::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(ev))
	}
}

func (ptr *QWebView) DropEventDefault(ev gui.QDropEvent_ITF) {
	defer qt.Recovering("QWebView::dropEvent")

	if ptr.Pointer() != nil {
		C.QWebView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(ev))
	}
}

//export callbackQWebView_Event
func callbackQWebView_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebView::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebView) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebView::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebView) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebView::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebView) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebView::event")

	if ptr.Pointer() != nil {
		return C.QWebView_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebView) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebView::event")

	if ptr.Pointer() != nil {
		return C.QWebView_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebView) FindText(subString string, options QWebPage__FindFlag) bool {
	defer qt.Recovering("QWebView::findText")

	if ptr.Pointer() != nil {
		return C.QWebView_FindText(ptr.Pointer(), C.CString(subString), C.int(options)) != 0
	}
	return false
}

//export callbackQWebView_FocusInEvent
func callbackQWebView_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QWebView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QWebView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

func (ptr *QWebView) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QWebView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QWebView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQWebView_FocusNextPrevChild
func callbackQWebView_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QWebView::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QWebView) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QWebView) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QWebView) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebView_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QWebView) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QWebView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QWebView_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQWebView_FocusOutEvent
func callbackQWebView_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWebView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWebView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWebView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

func (ptr *QWebView) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QWebView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QWebView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QWebView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQWebView_Forward
func callbackQWebView_Forward(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::forward")

	if signal := qt.GetSignal(C.GoString(ptrName), "forward"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectForward(f func()) {
	defer qt.Recovering("connect QWebView::forward")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "forward", f)
	}
}

func (ptr *QWebView) DisconnectForward() {
	defer qt.Recovering("disconnect QWebView::forward")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "forward")
	}
}

func (ptr *QWebView) Forward() {
	defer qt.Recovering("QWebView::forward")

	if ptr.Pointer() != nil {
		C.QWebView_Forward(ptr.Pointer())
	}
}

func (ptr *QWebView) History() *QWebHistory {
	defer qt.Recovering("QWebView::history")

	if ptr.Pointer() != nil {
		return NewQWebHistoryFromPointer(C.QWebView_History(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_IconChanged
func callbackQWebView_IconChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::iconChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectIconChanged(f func()) {
	defer qt.Recovering("connect QWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectIconChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconChanged", f)
	}
}

func (ptr *QWebView) DisconnectIconChanged() {
	defer qt.Recovering("disconnect QWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconChanged")
	}
}

func (ptr *QWebView) IconChanged() {
	defer qt.Recovering("QWebView::iconChanged")

	if ptr.Pointer() != nil {
		C.QWebView_IconChanged(ptr.Pointer())
	}
}

//export callbackQWebView_InputMethodEvent
func callbackQWebView_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *QWebView) ConnectInputMethodEvent(f func(e *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWebView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWebView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWebView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

func (ptr *QWebView) InputMethodEvent(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

func (ptr *QWebView) InputMethodEventDefault(e gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QWebView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QWebView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(e))
	}
}

//export callbackQWebView_InputMethodQuery
func callbackQWebView_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, property C.int) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(property)))
	}

	return core.PointerFromQVariant(NewQWebViewFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(property)))
}

func (ptr *QWebView) ConnectInputMethodQuery(f func(property core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QWebView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QWebView) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QWebView::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QWebView) InputMethodQuery(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebView_InputMethodQuery(ptr.Pointer(), C.int(property)))
	}
	return nil
}

func (ptr *QWebView) InputMethodQueryDefault(property core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWebView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWebView_InputMethodQueryDefault(ptr.Pointer(), C.int(property)))
	}
	return nil
}

//export callbackQWebView_KeyPressEvent
func callbackQWebView_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWebView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWebView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

func (ptr *QWebView) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QWebView) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QWebView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQWebView_KeyReleaseEvent
func callbackQWebView_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectKeyReleaseEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWebView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

func (ptr *QWebView) KeyReleaseEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QWebView) KeyReleaseEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QWebView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

//export callbackQWebView_LinkClicked
func callbackQWebView_LinkClicked(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::linkClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkClicked"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebView) ConnectLinkClicked(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectLinkClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkClicked", f)
	}
}

func (ptr *QWebView) DisconnectLinkClicked() {
	defer qt.Recovering("disconnect QWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLinkClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkClicked")
	}
}

func (ptr *QWebView) LinkClicked(url core.QUrl_ITF) {
	defer qt.Recovering("QWebView::linkClicked")

	if ptr.Pointer() != nil {
		C.QWebView_LinkClicked(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebView) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QWebView::load")

	if ptr.Pointer() != nil {
		C.QWebView_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebView_LoadFinished
func callbackQWebView_LoadFinished(ptr unsafe.Pointer, ptrName *C.char, ok C.int) {
	defer qt.Recovering("callback QWebView::loadFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadFinished"); signal != nil {
		signal.(func(bool))(int(ok) != 0)
	}

}

func (ptr *QWebView) ConnectLoadFinished(f func(ok bool)) {
	defer qt.Recovering("connect QWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectLoadFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFinished", f)
	}
}

func (ptr *QWebView) DisconnectLoadFinished() {
	defer qt.Recovering("disconnect QWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFinished")
	}
}

func (ptr *QWebView) LoadFinished(ok bool) {
	defer qt.Recovering("QWebView::loadFinished")

	if ptr.Pointer() != nil {
		C.QWebView_LoadFinished(ptr.Pointer(), C.int(qt.GoBoolToInt(ok)))
	}
}

//export callbackQWebView_LoadProgress
func callbackQWebView_LoadProgress(ptr unsafe.Pointer, ptrName *C.char, progress C.int) {
	defer qt.Recovering("callback QWebView::loadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadProgress"); signal != nil {
		signal.(func(int))(int(progress))
	}

}

func (ptr *QWebView) ConnectLoadProgress(f func(progress int)) {
	defer qt.Recovering("connect QWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectLoadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadProgress", f)
	}
}

func (ptr *QWebView) DisconnectLoadProgress() {
	defer qt.Recovering("disconnect QWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadProgress")
	}
}

func (ptr *QWebView) LoadProgress(progress int) {
	defer qt.Recovering("QWebView::loadProgress")

	if ptr.Pointer() != nil {
		C.QWebView_LoadProgress(ptr.Pointer(), C.int(progress))
	}
}

//export callbackQWebView_LoadStarted
func callbackQWebView_LoadStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::loadStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "loadStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectLoadStarted(f func()) {
	defer qt.Recovering("connect QWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectLoadStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadStarted", f)
	}
}

func (ptr *QWebView) DisconnectLoadStarted() {
	defer qt.Recovering("disconnect QWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectLoadStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadStarted")
	}
}

func (ptr *QWebView) LoadStarted() {
	defer qt.Recovering("QWebView::loadStarted")

	if ptr.Pointer() != nil {
		C.QWebView_LoadStarted(ptr.Pointer())
	}
}

//export callbackQWebView_MouseDoubleClickEvent
func callbackQWebView_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectMouseDoubleClickEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWebView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

func (ptr *QWebView) MouseDoubleClickEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWebView) MouseDoubleClickEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MouseMoveEvent
func callbackQWebView_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectMouseMoveEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWebView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

func (ptr *QWebView) MouseMoveEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWebView) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MousePressEvent
func callbackQWebView_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectMousePressEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWebView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWebView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

func (ptr *QWebView) MousePressEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWebView) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

//export callbackQWebView_MouseReleaseEvent
func callbackQWebView_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectMouseReleaseEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWebView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

func (ptr *QWebView) MouseReleaseEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWebView) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QWebView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QWebView) Page() *QWebPage {
	defer qt.Recovering("QWebView::page")

	if ptr.Pointer() != nil {
		return NewQWebPageFromPointer(C.QWebView_Page(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) PageAction(action QWebPage__WebAction) *widgets.QAction {
	defer qt.Recovering("QWebView::pageAction")

	if ptr.Pointer() != nil {
		return widgets.NewQActionFromPointer(C.QWebView_PageAction(ptr.Pointer(), C.int(action)))
	}
	return nil
}

//export callbackQWebView_PaintEvent
func callbackQWebView_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectPaintEvent(f func(ev *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWebView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWebView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWebView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

func (ptr *QWebView) PaintEvent(ev gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebView::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(ev))
	}
}

func (ptr *QWebView) PaintEventDefault(ev gui.QPaintEvent_ITF) {
	defer qt.Recovering("QWebView::paintEvent")

	if ptr.Pointer() != nil {
		C.QWebView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(ev))
	}
}

//export callbackQWebView_Print
func callbackQWebView_Print(ptr unsafe.Pointer, ptrName *C.char, printer unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::print")

	if signal := qt.GetSignal(C.GoString(ptrName), "print"); signal != nil {
		signal.(func(*printsupport.QPrinter))(printsupport.NewQPrinterFromPointer(printer))
	}

}

func (ptr *QWebView) ConnectPrint(f func(printer *printsupport.QPrinter)) {
	defer qt.Recovering("connect QWebView::print")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "print", f)
	}
}

func (ptr *QWebView) DisconnectPrint(printer printsupport.QPrinter_ITF) {
	defer qt.Recovering("disconnect QWebView::print")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "print")
	}
}

func (ptr *QWebView) Print(printer printsupport.QPrinter_ITF) {
	defer qt.Recovering("QWebView::print")

	if ptr.Pointer() != nil {
		C.QWebView_Print(ptr.Pointer(), printsupport.PointerFromQPrinter(printer))
	}
}

//export callbackQWebView_Reload
func callbackQWebView_Reload(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::reload")

	if signal := qt.GetSignal(C.GoString(ptrName), "reload"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectReload(f func()) {
	defer qt.Recovering("connect QWebView::reload")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reload", f)
	}
}

func (ptr *QWebView) DisconnectReload() {
	defer qt.Recovering("disconnect QWebView::reload")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reload")
	}
}

func (ptr *QWebView) Reload() {
	defer qt.Recovering("QWebView::reload")

	if ptr.Pointer() != nil {
		C.QWebView_Reload(ptr.Pointer())
	}
}

func (ptr *QWebView) RenderHints() gui.QPainter__RenderHint {
	defer qt.Recovering("QWebView::renderHints")

	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QWebView_RenderHints(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebView_ResizeEvent
func callbackQWebView_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQWebViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QWebView) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWebView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWebView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWebView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

func (ptr *QWebView) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QWebView) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QWebView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

//export callbackQWebView_SelectionChanged
func callbackQWebView_SelectionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QWebView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QWebView) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QWebView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

func (ptr *QWebView) SelectionChanged() {
	defer qt.Recovering("QWebView::selectionChanged")

	if ptr.Pointer() != nil {
		C.QWebView_SelectionChanged(ptr.Pointer())
	}
}

func (ptr *QWebView) SetContent(data string, mimeType string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebView::setContent")

	if ptr.Pointer() != nil {
		C.QWebView_SetContent(ptr.Pointer(), C.CString(data), C.CString(mimeType), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebView) SetHtml(html string, baseUrl core.QUrl_ITF) {
	defer qt.Recovering("QWebView::setHtml")

	if ptr.Pointer() != nil {
		C.QWebView_SetHtml(ptr.Pointer(), C.CString(html), core.PointerFromQUrl(baseUrl))
	}
}

func (ptr *QWebView) SetPage(page QWebPage_ITF) {
	defer qt.Recovering("QWebView::setPage")

	if ptr.Pointer() != nil {
		C.QWebView_SetPage(ptr.Pointer(), PointerFromQWebPage(page))
	}
}

func (ptr *QWebView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	defer qt.Recovering("QWebView::setRenderHint")

	if ptr.Pointer() != nil {
		C.QWebView_SetRenderHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QWebView) SetRenderHints(hints gui.QPainter__RenderHint) {
	defer qt.Recovering("QWebView::setRenderHints")

	if ptr.Pointer() != nil {
		C.QWebView_SetRenderHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QWebView) SetTextSizeMultiplier(factor float64) {
	defer qt.Recovering("QWebView::setTextSizeMultiplier")

	if ptr.Pointer() != nil {
		C.QWebView_SetTextSizeMultiplier(ptr.Pointer(), C.double(factor))
	}
}

func (ptr *QWebView) Settings() *QWebSettings {
	defer qt.Recovering("QWebView::settings")

	if ptr.Pointer() != nil {
		return NewQWebSettingsFromPointer(C.QWebView_Settings(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_SizeHint
func callbackQWebView_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebViewFromPointer(ptr).SizeHintDefault())
}

func (ptr *QWebView) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebView::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QWebView) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QWebView::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QWebView) SizeHint() *core.QSize {
	defer qt.Recovering("QWebView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebView_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebView_SizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_StatusBarMessage
func callbackQWebView_StatusBarMessage(ptr unsafe.Pointer, ptrName *C.char, text *C.char) {
	defer qt.Recovering("callback QWebView::statusBarMessage")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusBarMessage"); signal != nil {
		signal.(func(string))(C.GoString(text))
	}

}

func (ptr *QWebView) ConnectStatusBarMessage(f func(text string)) {
	defer qt.Recovering("connect QWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectStatusBarMessage(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusBarMessage", f)
	}
}

func (ptr *QWebView) DisconnectStatusBarMessage() {
	defer qt.Recovering("disconnect QWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectStatusBarMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusBarMessage")
	}
}

func (ptr *QWebView) StatusBarMessage(text string) {
	defer qt.Recovering("QWebView::statusBarMessage")

	if ptr.Pointer() != nil {
		C.QWebView_StatusBarMessage(ptr.Pointer(), C.CString(text))
	}
}

//export callbackQWebView_Stop
func callbackQWebView_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebView) ConnectStop(f func()) {
	defer qt.Recovering("connect QWebView::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QWebView) DisconnectStop() {
	defer qt.Recovering("disconnect QWebView::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QWebView) Stop() {
	defer qt.Recovering("QWebView::stop")

	if ptr.Pointer() != nil {
		C.QWebView_Stop(ptr.Pointer())
	}
}

func (ptr *QWebView) TextSizeMultiplier() float64 {
	defer qt.Recovering("QWebView::textSizeMultiplier")

	if ptr.Pointer() != nil {
		return float64(C.QWebView_TextSizeMultiplier(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebView_TitleChanged
func callbackQWebView_TitleChanged(ptr unsafe.Pointer, ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWebView::titleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "titleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWebView) ConnectTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "titleChanged", f)
	}
}

func (ptr *QWebView) DisconnectTitleChanged() {
	defer qt.Recovering("disconnect QWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "titleChanged")
	}
}

func (ptr *QWebView) TitleChanged(title string) {
	defer qt.Recovering("QWebView::titleChanged")

	if ptr.Pointer() != nil {
		C.QWebView_TitleChanged(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QWebView) TriggerPageAction(action QWebPage__WebAction, checked bool) {
	defer qt.Recovering("QWebView::triggerPageAction")

	if ptr.Pointer() != nil {
		C.QWebView_TriggerPageAction(ptr.Pointer(), C.int(action), C.int(qt.GoBoolToInt(checked)))
	}
}

//export callbackQWebView_UrlChanged
func callbackQWebView_UrlChanged(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::urlChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "urlChanged"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebView) ConnectUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlChanged", f)
	}
}

func (ptr *QWebView) DisconnectUrlChanged() {
	defer qt.Recovering("disconnect QWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlChanged")
	}
}

func (ptr *QWebView) UrlChanged(url core.QUrl_ITF) {
	defer qt.Recovering("QWebView::urlChanged")

	if ptr.Pointer() != nil {
		C.QWebView_UrlChanged(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQWebView_WheelEvent
func callbackQWebView_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(ev))
	} else {
		NewQWebViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(ev))
	}
}

func (ptr *QWebView) ConnectWheelEvent(f func(ev *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWebView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWebView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWebView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

func (ptr *QWebView) WheelEvent(ev gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(ev))
	}
}

func (ptr *QWebView) WheelEventDefault(ev gui.QWheelEvent_ITF) {
	defer qt.Recovering("QWebView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QWebView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(ev))
	}
}

func (ptr *QWebView) DestroyQWebView() {
	defer qt.Recovering("QWebView::~QWebView")

	if ptr.Pointer() != nil {
		C.QWebView_DestroyQWebView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebView_ActionEvent
func callbackQWebView_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWebView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWebView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWebView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

func (ptr *QWebView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebView::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QWebView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QWebView::actionEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQWebView_EnterEvent
func callbackQWebView_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWebView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWebView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

func (ptr *QWebView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::enterEvent")

	if ptr.Pointer() != nil {
		C.QWebView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_HideEvent
func callbackQWebView_HideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWebView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWebView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWebView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

func (ptr *QWebView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebView::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QWebView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QWebView::hideEvent")

	if ptr.Pointer() != nil {
		C.QWebView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQWebView_LeaveEvent
func callbackQWebView_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWebView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWebView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

func (ptr *QWebView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_Metric
func callbackQWebView_Metric(ptr unsafe.Pointer, ptrName *C.char, m C.int) C.int {
	defer qt.Recovering("callback QWebView::metric")

	if signal := qt.GetSignal(C.GoString(ptrName), "metric"); signal != nil {
		return C.int(signal.(func(gui.QPaintDevice__PaintDeviceMetric) int)(gui.QPaintDevice__PaintDeviceMetric(m)))
	}

	return C.int(NewQWebViewFromPointer(ptr).MetricDefault(gui.QPaintDevice__PaintDeviceMetric(m)))
}

func (ptr *QWebView) ConnectMetric(f func(m gui.QPaintDevice__PaintDeviceMetric) int) {
	defer qt.Recovering("connect QWebView::metric")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metric", f)
	}
}

func (ptr *QWebView) DisconnectMetric() {
	defer qt.Recovering("disconnect QWebView::metric")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metric")
	}
}

func (ptr *QWebView) Metric(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QWebView::metric")

	if ptr.Pointer() != nil {
		return int(C.QWebView_Metric(ptr.Pointer(), C.int(m)))
	}
	return 0
}

func (ptr *QWebView) MetricDefault(m gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QWebView::metric")

	if ptr.Pointer() != nil {
		return int(C.QWebView_MetricDefault(ptr.Pointer(), C.int(m)))
	}
	return 0
}

//export callbackQWebView_MinimumSizeHint
func callbackQWebView_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQWebViewFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QWebView) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QWebView::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QWebView) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QWebView::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QWebView) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QWebView::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebView_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QWebView::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWebView_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_MoveEvent
func callbackQWebView_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWebView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWebView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWebView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

func (ptr *QWebView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebView::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QWebView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QWebView::moveEvent")

	if ptr.Pointer() != nil {
		C.QWebView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQWebView_PaintEngine
func callbackQWebView_PaintEngine(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::paintEngine")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEngine"); signal != nil {
		return gui.PointerFromQPaintEngine(signal.(func() *gui.QPaintEngine)())
	}

	return gui.PointerFromQPaintEngine(NewQWebViewFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QWebView) ConnectPaintEngine(f func() *gui.QPaintEngine) {
	defer qt.Recovering("connect QWebView::paintEngine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEngine", f)
	}
}

func (ptr *QWebView) DisconnectPaintEngine() {
	defer qt.Recovering("disconnect QWebView::paintEngine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEngine")
	}
}

func (ptr *QWebView) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QWebView::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebView_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) PaintEngineDefault() *gui.QPaintEngine {
	defer qt.Recovering("QWebView::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWebView_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebView_SetEnabled
func callbackQWebView_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebView::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QWebView) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QWebView::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QWebView) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QWebView::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QWebView) SetEnabled(vbo bool) {
	defer qt.Recovering("QWebView::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebView_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebView) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QWebView::setEnabled")

	if ptr.Pointer() != nil {
		C.QWebView_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebView_SetStyleSheet
func callbackQWebView_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QWebView::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQWebViewFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QWebView) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QWebView::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QWebView) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QWebView::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QWebView) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QWebView::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebView_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QWebView) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QWebView::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWebView_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQWebView_SetVisible
func callbackQWebView_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QWebView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QWebView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWebView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWebView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWebView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QWebView) SetVisible(visible bool) {
	defer qt.Recovering("QWebView::setVisible")

	if ptr.Pointer() != nil {
		C.QWebView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWebView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QWebView::setVisible")

	if ptr.Pointer() != nil {
		C.QWebView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQWebView_SetWindowModified
func callbackQWebView_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QWebView::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QWebView) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QWebView::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QWebView) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QWebView::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QWebView) SetWindowModified(vbo bool) {
	defer qt.Recovering("QWebView::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebView_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QWebView) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QWebView::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWebView_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQWebView_SetWindowTitle
func callbackQWebView_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QWebView::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQWebViewFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QWebView) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QWebView::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QWebView) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QWebView::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QWebView) SetWindowTitle(vqs string) {
	defer qt.Recovering("QWebView::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebView_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QWebView) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QWebView::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWebView_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQWebView_ShowEvent
func callbackQWebView_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWebView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWebView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWebView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

func (ptr *QWebView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebView::showEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QWebView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QWebView::showEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQWebView_Close
func callbackQWebView_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebView::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).CloseDefault()))
}

func (ptr *QWebView) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QWebView::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QWebView) DisconnectClose() {
	defer qt.Recovering("disconnect QWebView::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QWebView) Close() bool {
	defer qt.Recovering("QWebView::close")

	if ptr.Pointer() != nil {
		return C.QWebView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) CloseDefault() bool {
	defer qt.Recovering("QWebView::close")

	if ptr.Pointer() != nil {
		return C.QWebView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebView_CloseEvent
func callbackQWebView_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWebView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWebView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWebView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

func (ptr *QWebView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebView::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QWebView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QWebView::closeEvent")

	if ptr.Pointer() != nil {
		C.QWebView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQWebView_HasHeightForWidth
func callbackQWebView_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QWebView::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QWebView) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QWebView::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QWebView) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QWebView::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QWebView) HasHeightForWidth() bool {
	defer qt.Recovering("QWebView::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebView_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebView) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QWebView::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWebView_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebView_HeightForWidth
func callbackQWebView_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QWebView::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQWebViewFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QWebView) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QWebView::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QWebView) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QWebView::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QWebView) HeightForWidth(w int) int {
	defer qt.Recovering("QWebView::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebView_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWebView) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QWebView::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWebView_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQWebView_Hide
func callbackQWebView_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWebView) ConnectHide(f func()) {
	defer qt.Recovering("connect QWebView::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QWebView) DisconnectHide() {
	defer qt.Recovering("disconnect QWebView::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QWebView) Hide() {
	defer qt.Recovering("QWebView::hide")

	if ptr.Pointer() != nil {
		C.QWebView_Hide(ptr.Pointer())
	}
}

func (ptr *QWebView) HideDefault() {
	defer qt.Recovering("QWebView::hide")

	if ptr.Pointer() != nil {
		C.QWebView_HideDefault(ptr.Pointer())
	}
}

//export callbackQWebView_InitPainter
func callbackQWebView_InitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQWebViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QWebView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QWebView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QWebView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QWebView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

func (ptr *QWebView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWebView::initPainter")

	if ptr.Pointer() != nil {
		C.QWebView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QWebView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QWebView::initPainter")

	if ptr.Pointer() != nil {
		C.QWebView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQWebView_Lower
func callbackQWebView_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QWebView) ConnectLower(f func()) {
	defer qt.Recovering("connect QWebView::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QWebView) DisconnectLower() {
	defer qt.Recovering("disconnect QWebView::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QWebView) Lower() {
	defer qt.Recovering("QWebView::lower")

	if ptr.Pointer() != nil {
		C.QWebView_Lower(ptr.Pointer())
	}
}

func (ptr *QWebView) LowerDefault() {
	defer qt.Recovering("QWebView::lower")

	if ptr.Pointer() != nil {
		C.QWebView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQWebView_NativeEvent
func callbackQWebView_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QWebView::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QWebView) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QWebView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QWebView) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QWebView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QWebView) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebView_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QWebView) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QWebView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QWebView_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQWebView_Raise
func callbackQWebView_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QWebView) ConnectRaise(f func()) {
	defer qt.Recovering("connect QWebView::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QWebView) DisconnectRaise() {
	defer qt.Recovering("disconnect QWebView::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QWebView) Raise() {
	defer qt.Recovering("QWebView::raise")

	if ptr.Pointer() != nil {
		C.QWebView_Raise(ptr.Pointer())
	}
}

func (ptr *QWebView) RaiseDefault() {
	defer qt.Recovering("QWebView::raise")

	if ptr.Pointer() != nil {
		C.QWebView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQWebView_Repaint
func callbackQWebView_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QWebView) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QWebView::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QWebView) DisconnectRepaint() {
	defer qt.Recovering("disconnect QWebView::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QWebView) Repaint() {
	defer qt.Recovering("QWebView::repaint")

	if ptr.Pointer() != nil {
		C.QWebView_Repaint(ptr.Pointer())
	}
}

func (ptr *QWebView) RepaintDefault() {
	defer qt.Recovering("QWebView::repaint")

	if ptr.Pointer() != nil {
		C.QWebView_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQWebView_SetDisabled
func callbackQWebView_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QWebView::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QWebView) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QWebView::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QWebView) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QWebView::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QWebView) SetDisabled(disable bool) {
	defer qt.Recovering("QWebView::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebView_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWebView) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QWebView::setDisabled")

	if ptr.Pointer() != nil {
		C.QWebView_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQWebView_SetFocus2
func callbackQWebView_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QWebView) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QWebView::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QWebView) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QWebView::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QWebView) SetFocus2() {
	defer qt.Recovering("QWebView::setFocus")

	if ptr.Pointer() != nil {
		C.QWebView_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWebView) SetFocus2Default() {
	defer qt.Recovering("QWebView::setFocus")

	if ptr.Pointer() != nil {
		C.QWebView_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQWebView_SetHidden
func callbackQWebView_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QWebView::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQWebViewFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QWebView) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QWebView::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QWebView) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QWebView::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QWebView) SetHidden(hidden bool) {
	defer qt.Recovering("QWebView::setHidden")

	if ptr.Pointer() != nil {
		C.QWebView_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWebView) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QWebView::setHidden")

	if ptr.Pointer() != nil {
		C.QWebView_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQWebView_Show
func callbackQWebView_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWebView) ConnectShow(f func()) {
	defer qt.Recovering("connect QWebView::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QWebView) DisconnectShow() {
	defer qt.Recovering("disconnect QWebView::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QWebView) Show() {
	defer qt.Recovering("QWebView::show")

	if ptr.Pointer() != nil {
		C.QWebView_Show(ptr.Pointer())
	}
}

func (ptr *QWebView) ShowDefault() {
	defer qt.Recovering("QWebView::show")

	if ptr.Pointer() != nil {
		C.QWebView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowFullScreen
func callbackQWebView_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QWebView) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QWebView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QWebView) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QWebView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QWebView) ShowFullScreen() {
	defer qt.Recovering("QWebView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebView_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWebView) ShowFullScreenDefault() {
	defer qt.Recovering("QWebView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWebView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowMaximized
func callbackQWebView_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QWebView) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QWebView::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QWebView) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QWebView::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QWebView) ShowMaximized() {
	defer qt.Recovering("QWebView::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebView_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWebView) ShowMaximizedDefault() {
	defer qt.Recovering("QWebView::showMaximized")

	if ptr.Pointer() != nil {
		C.QWebView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowMinimized
func callbackQWebView_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QWebView) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QWebView::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QWebView) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QWebView::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QWebView) ShowMinimized() {
	defer qt.Recovering("QWebView::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebView_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWebView) ShowMinimizedDefault() {
	defer qt.Recovering("QWebView::showMinimized")

	if ptr.Pointer() != nil {
		C.QWebView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQWebView_ShowNormal
func callbackQWebView_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QWebView) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QWebView::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QWebView) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QWebView::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QWebView) ShowNormal() {
	defer qt.Recovering("QWebView::showNormal")

	if ptr.Pointer() != nil {
		C.QWebView_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWebView) ShowNormalDefault() {
	defer qt.Recovering("QWebView::showNormal")

	if ptr.Pointer() != nil {
		C.QWebView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQWebView_TabletEvent
func callbackQWebView_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWebView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWebView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWebView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

func (ptr *QWebView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QWebView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QWebView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QWebView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQWebView_Update
func callbackQWebView_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QWebView) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QWebView::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QWebView) DisconnectUpdate() {
	defer qt.Recovering("disconnect QWebView::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QWebView) Update() {
	defer qt.Recovering("QWebView::update")

	if ptr.Pointer() != nil {
		C.QWebView_Update(ptr.Pointer())
	}
}

func (ptr *QWebView) UpdateDefault() {
	defer qt.Recovering("QWebView::update")

	if ptr.Pointer() != nil {
		C.QWebView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQWebView_UpdateMicroFocus
func callbackQWebView_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QWebView) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QWebView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QWebView) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QWebView::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QWebView) UpdateMicroFocus() {
	defer qt.Recovering("QWebView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebView_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QWebView) UpdateMicroFocusDefault() {
	defer qt.Recovering("QWebView::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QWebView_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQWebView_TimerEvent
func callbackQWebView_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebView::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebView::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebView_ChildEvent
func callbackQWebView_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebView::childEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebView::childEvent")

	if ptr.Pointer() != nil {
		C.QWebView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebView_ConnectNotify
func callbackQWebView_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebView::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebView) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebView::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebView) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebView::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebView::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebView_CustomEvent
func callbackQWebView_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::customEvent")

	if ptr.Pointer() != nil {
		C.QWebView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebView::customEvent")

	if ptr.Pointer() != nil {
		C.QWebView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebView_DeleteLater
func callbackQWebView_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebView::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebView) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebView::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebView) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebView::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebView) DeleteLater() {
	defer qt.Recovering("QWebView::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebView_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebView) DeleteLaterDefault() {
	defer qt.Recovering("QWebView::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebView_DisconnectNotify
func callbackQWebView_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebView::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebView) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebView_EventFilter
func callbackQWebView_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebView::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebView::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebView) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebView::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebView_MetaObject
func callbackQWebView_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebView::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebView) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebView::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebView) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebView::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebView) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebView) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
