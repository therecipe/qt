package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

type QHelpContentItem struct {
	ptr unsafe.Pointer
}

type QHelpContentItem_ITF interface {
	QHelpContentItem_PTR() *QHelpContentItem
}

func (p *QHelpContentItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpContentItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpContentItem(ptr QHelpContentItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentItem_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = new(QHelpContentItem)
	n.SetPointer(ptr)
	return n
}

func newQHelpContentItemFromPointer(ptr unsafe.Pointer) *QHelpContentItem {
	var n = NewQHelpContentItemFromPointer(ptr)
	return n
}

func (ptr *QHelpContentItem) QHelpContentItem_PTR() *QHelpContentItem {
	return ptr
}

func (ptr *QHelpContentItem) Child(row int) *QHelpContentItem {
	defer qt.Recovering("QHelpContentItem::child")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Child(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QHelpContentItem) ChildCount() int {
	defer qt.Recovering("QHelpContentItem::childCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentItem) ChildPosition(child QHelpContentItem_ITF) int {
	defer qt.Recovering("QHelpContentItem::childPosition")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_ChildPosition(ptr.Pointer(), PointerFromQHelpContentItem(child)))
	}
	return 0
}

func (ptr *QHelpContentItem) Parent() *QHelpContentItem {
	defer qt.Recovering("QHelpContentItem::parent")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) Row() int {
	defer qt.Recovering("QHelpContentItem::row")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentItem_Row(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpContentItem) Title() string {
	defer qt.Recovering("QHelpContentItem::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpContentItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpContentItem) Url() *core.QUrl {
	defer qt.Recovering("QHelpContentItem::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpContentItem_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpContentItem) DestroyQHelpContentItem() {
	defer qt.Recovering("QHelpContentItem::~QHelpContentItem")

	if ptr.Pointer() != nil {
		C.QHelpContentItem_DestroyQHelpContentItem(ptr.Pointer())
	}
}

type QHelpContentModel struct {
	core.QAbstractItemModel
}

type QHelpContentModel_ITF interface {
	core.QAbstractItemModel_ITF
	QHelpContentModel_PTR() *QHelpContentModel
}

func PointerFromQHelpContentModel(ptr QHelpContentModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = new(QHelpContentModel)
	n.SetPointer(ptr)
	return n
}

func newQHelpContentModelFromPointer(ptr unsafe.Pointer) *QHelpContentModel {
	var n = NewQHelpContentModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpContentModel_") {
		n.SetObjectName("QHelpContentModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpContentModel) QHelpContentModel_PTR() *QHelpContentModel {
	return ptr
}

func (ptr *QHelpContentModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::columnCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) ContentItemAt(index core.QModelIndex_ITF) *QHelpContentItem {
	defer qt.Recovering("QHelpContentModel::contentItemAt")

	if ptr.Pointer() != nil {
		return NewQHelpContentItemFromPointer(C.QHelpContentModel_ContentItemAt(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QHelpContentModel) ConnectContentsCreated(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreated", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreated() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreated")
	}
}

//export callbackQHelpContentModelContentsCreated
func callbackQHelpContentModelContentsCreated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ContentsCreated() {
	defer qt.Recovering("QHelpContentModel::contentsCreated")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreated(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ConnectContentsCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ConnectContentsCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsCreationStarted", f)
	}
}

func (ptr *QHelpContentModel) DisconnectContentsCreationStarted() {
	defer qt.Recovering("disconnect QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DisconnectContentsCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsCreationStarted")
	}
}

//export callbackQHelpContentModelContentsCreationStarted
func callbackQHelpContentModelContentsCreationStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpContentModel::contentsCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpContentModel) ContentsCreationStarted() {
	defer qt.Recovering("QHelpContentModel::contentsCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ContentsCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) CreateContents(customFilterName string) {
	defer qt.Recovering("QHelpContentModel::createContents")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CreateContents(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpContentModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QHelpContentModel::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpContentModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QHelpContentModel) IsCreatingContents() bool {
	defer qt.Recovering("QHelpContentModel::isCreatingContents")

	if ptr.Pointer() != nil {
		return C.QHelpContentModel_IsCreatingContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpContentModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QHelpContentModel::rowCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpContentModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QHelpContentModel) DestroyQHelpContentModel() {
	defer qt.Recovering("QHelpContentModel::~QHelpContentModel")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_DestroyQHelpContentModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpContentModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QHelpContentModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpContentModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQHelpContentModelRevert
func callbackQHelpContentModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpContentModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpContentModel) Revert() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) RevertDefault() {
	defer qt.Recovering("QHelpContentModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QHelpContentModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpContentModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQHelpContentModelSort
func callbackQHelpContentModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QHelpContentModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQHelpContentModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpContentModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpContentModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpContentModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpContentModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpContentModelTimerEvent
func callbackQHelpContentModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpContentModelChildEvent
func callbackQHelpContentModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpContentModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpContentModelCustomEvent
func callbackQHelpContentModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpContentWidget struct {
	widgets.QTreeView
}

type QHelpContentWidget_ITF interface {
	widgets.QTreeView_ITF
	QHelpContentWidget_PTR() *QHelpContentWidget
}

func PointerFromQHelpContentWidget(ptr QHelpContentWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpContentWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = new(QHelpContentWidget)
	n.SetPointer(ptr)
	return n
}

func newQHelpContentWidgetFromPointer(ptr unsafe.Pointer) *QHelpContentWidget {
	var n = NewQHelpContentWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpContentWidget_") {
		n.SetObjectName("QHelpContentWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpContentWidget) QHelpContentWidget_PTR() *QHelpContentWidget {
	return ptr
}

func (ptr *QHelpContentWidget) ConnectLinkActivated(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpContentWidgetLinkActivated
func callbackQHelpContentWidgetLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpContentWidget) LinkActivated(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpContentWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpContentWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpContentWidgetDragMoveEvent
func callbackQHelpContentWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpContentWidgetKeyPressEvent
func callbackQHelpContentWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQHelpContentWidgetKeyboardSearch
func callbackQHelpContentWidgetKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QHelpContentWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpContentWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpContentWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpContentWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpContentWidgetMouseDoubleClickEvent
func callbackQHelpContentWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpContentWidgetMouseMoveEvent
func callbackQHelpContentWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpContentWidgetMousePressEvent
func callbackQHelpContentWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpContentWidgetMouseReleaseEvent
func callbackQHelpContentWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpContentWidgetPaintEvent
func callbackQHelpContentWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpContentWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQHelpContentWidgetReset
func callbackQHelpContentWidgetReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpContentWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) Reset() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ResetDefault() {
	defer qt.Recovering("QHelpContentWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQHelpContentWidgetScrollContentsBy
func callbackQHelpContentWidgetScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpContentWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpContentWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpContentWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpContentWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQHelpContentWidgetSelectAll
func callbackQHelpContentWidgetSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpContentWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) SelectAll() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpContentWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQHelpContentWidgetSetModel
func callbackQHelpContentWidgetSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpContentWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQHelpContentWidgetSetSelection
func callbackQHelpContentWidgetSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QHelpContentWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpContentWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpContentWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpContentWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpContentWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQHelpContentWidgetSetSelectionModel
func callbackQHelpContentWidgetSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpContentWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpContentWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpContentWidgetTimerEvent
func callbackQHelpContentWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQHelpContentWidgetUpdateGeometries
func callbackQHelpContentWidgetUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpContentWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpContentWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QHelpContentWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpContentWidgetDragLeaveEvent
func callbackQHelpContentWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQHelpContentWidgetCloseEditor
func callbackQHelpContentWidgetCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QHelpContentWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpContentWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpContentWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQHelpContentWidgetCommitData
func callbackQHelpContentWidgetCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpContentWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpContentWidgetDragEnterEvent
func callbackQHelpContentWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpContentWidgetDropEvent
func callbackQHelpContentWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQHelpContentWidgetEditorDestroyed
func callbackQHelpContentWidgetEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpContentWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpContentWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpContentWidgetFocusInEvent
func callbackQHelpContentWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpContentWidgetFocusOutEvent
func callbackQHelpContentWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpContentWidgetInputMethodEvent
func callbackQHelpContentWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpContentWidgetResizeEvent
func callbackQHelpContentWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQHelpContentWidgetStartDrag
func callbackQHelpContentWidgetStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QHelpContentWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpContentWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpContentWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpContentWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpContentWidgetContextMenuEvent
func callbackQHelpContentWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpContentWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQHelpContentWidgetSetupViewport
func callbackQHelpContentWidgetSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpContentWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpContentWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpContentWidgetWheelEvent
func callbackQHelpContentWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpContentWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpContentWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQHelpContentWidgetChangeEvent
func callbackQHelpContentWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpContentWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpContentWidgetActionEvent
func callbackQHelpContentWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpContentWidgetEnterEvent
func callbackQHelpContentWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpContentWidgetHideEvent
func callbackQHelpContentWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpContentWidgetLeaveEvent
func callbackQHelpContentWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpContentWidgetMoveEvent
func callbackQHelpContentWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpContentWidgetSetVisible
func callbackQHelpContentWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpContentWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpContentWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpContentWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpContentWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpContentWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpContentWidgetShowEvent
func callbackQHelpContentWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpContentWidgetCloseEvent
func callbackQHelpContentWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpContentWidgetInitPainter
func callbackQHelpContentWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpContentWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpContentWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpContentWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpContentWidgetKeyReleaseEvent
func callbackQHelpContentWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpContentWidgetTabletEvent
func callbackQHelpContentWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpContentWidgetChildEvent
func callbackQHelpContentWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpContentWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpContentWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpContentWidgetCustomEvent
func callbackQHelpContentWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpContentWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpContentWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpContentWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpContentWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpContentWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpContentWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpEngine struct {
	QHelpEngineCore
}

type QHelpEngine_ITF interface {
	QHelpEngineCore_ITF
	QHelpEngine_PTR() *QHelpEngine
}

func PointerFromQHelpEngine(ptr QHelpEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = new(QHelpEngine)
	n.SetPointer(ptr)
	return n
}

func newQHelpEngineFromPointer(ptr unsafe.Pointer) *QHelpEngine {
	var n = NewQHelpEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpEngine_") {
		n.SetObjectName("QHelpEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpEngine) QHelpEngine_PTR() *QHelpEngine {
	return ptr
}

func NewQHelpEngine(collectionFile string, parent core.QObject_ITF) *QHelpEngine {
	defer qt.Recovering("QHelpEngine::QHelpEngine")

	return newQHelpEngineFromPointer(C.QHelpEngine_NewQHelpEngine(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngine) ContentModel() *QHelpContentModel {
	defer qt.Recovering("QHelpEngine::contentModel")

	if ptr.Pointer() != nil {
		return NewQHelpContentModelFromPointer(C.QHelpEngine_ContentModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) ContentWidget() *QHelpContentWidget {
	defer qt.Recovering("QHelpEngine::contentWidget")

	if ptr.Pointer() != nil {
		return NewQHelpContentWidgetFromPointer(C.QHelpEngine_ContentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexModel() *QHelpIndexModel {
	defer qt.Recovering("QHelpEngine::indexModel")

	if ptr.Pointer() != nil {
		return NewQHelpIndexModelFromPointer(C.QHelpEngine_IndexModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) IndexWidget() *QHelpIndexWidget {
	defer qt.Recovering("QHelpEngine::indexWidget")

	if ptr.Pointer() != nil {
		return NewQHelpIndexWidgetFromPointer(C.QHelpEngine_IndexWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) SearchEngine() *QHelpSearchEngine {
	defer qt.Recovering("QHelpEngine::searchEngine")

	if ptr.Pointer() != nil {
		return NewQHelpSearchEngineFromPointer(C.QHelpEngine_SearchEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpEngine) DestroyQHelpEngine() {
	defer qt.Recovering("QHelpEngine::~QHelpEngine")

	if ptr.Pointer() != nil {
		C.QHelpEngine_DestroyQHelpEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpEngineTimerEvent
func callbackQHelpEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpEngineChildEvent
func callbackQHelpEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpEngineCustomEvent
func callbackQHelpEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpEngineCore struct {
	core.QObject
}

type QHelpEngineCore_ITF interface {
	core.QObject_ITF
	QHelpEngineCore_PTR() *QHelpEngineCore
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	return n
}

func newQHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = NewQHelpEngineCoreFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpEngineCore_") {
		n.SetObjectName("QHelpEngineCore_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return ptr
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	defer qt.Recovering("QHelpEngineCore::autoSaveFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	defer qt.Recovering("QHelpEngineCore::collectionFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	defer qt.Recovering("QHelpEngineCore::currentFilter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	defer qt.Recovering("QHelpEngineCore::setAutoSaveFilter")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.int(qt.GoBoolToInt(save)))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	defer qt.Recovering("QHelpEngineCore::setCollectionFile")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	defer qt.Recovering("QHelpEngineCore::setCurrentFilter")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), C.CString(filterName))
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	defer qt.Recovering("QHelpEngineCore::QHelpEngineCore")

	return newQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	defer qt.Recovering("QHelpEngineCore::addCustomFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), C.CString(filterName), C.CString(strings.Join(attributes, "|"))) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	defer qt.Recovering("QHelpEngineCore::copyCollectionFile")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	defer qt.Recovering("connect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	defer qt.Recovering("disconnect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentFilterChanged")
	}
}

//export callbackQHelpEngineCoreCurrentFilterChanged
func callbackQHelpEngineCoreCurrentFilterChanged(ptr unsafe.Pointer, ptrName *C.char, newFilter *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::currentFilterChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentFilterChanged"); signal != nil {
		signal.(func(string))(C.GoString(newFilter))
	}

}

func (ptr *QHelpEngineCore) CurrentFilterChanged(newFilter string) {
	defer qt.Recovering("QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CurrentFilterChanged(ptr.Pointer(), C.CString(newFilter))
	}
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	defer qt.Recovering("QHelpEngineCore::customFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::customValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	defer qt.Recovering("QHelpEngineCore::documentationFileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), C.CString(namespaceName)))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	defer qt.Recovering("QHelpEngineCore::error")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) string {
	defer qt.Recovering("QHelpEngineCore::fileData")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return ""
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), C.CString(filterName))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QHelpEngineCore::findFile")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::metaData")

	return core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.CString(documentationFileName), C.CString(name)))
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	defer qt.Recovering("QHelpEngineCore::namespaceName")

	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.CString(documentationFileName)))
}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	defer qt.Recovering("disconnect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated")
	}
}

//export callbackQHelpEngineCoreReadersAboutToBeInvalidated
func callbackQHelpEngineCoreReadersAboutToBeInvalidated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::readersAboutToBeInvalidated")

	if signal := qt.GetSignal(C.GoString(ptrName), "readersAboutToBeInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ReadersAboutToBeInvalidated() {
	defer qt.Recovering("QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ReadersAboutToBeInvalidated(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	defer qt.Recovering("QHelpEngineCore::registerDocumentation")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), C.CString(documentationFileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	defer qt.Recovering("QHelpEngineCore::registeredDocumentations")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), C.CString(filterName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomValue")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::setCustomValue")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	defer qt.Recovering("QHelpEngineCore::setupData")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupFinished")
	}
}

//export callbackQHelpEngineCoreSetupFinished
func callbackQHelpEngineCoreSetupFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::setupFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) SetupFinished() {
	defer qt.Recovering("QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupFinished(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupStarted")
	}
}

//export callbackQHelpEngineCoreSetupStarted
func callbackQHelpEngineCoreSetupStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::setupStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) SetupStarted() {
	defer qt.Recovering("QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetupStarted(ptr.Pointer())
	}
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	defer qt.Recovering("QHelpEngineCore::unregisterDocumentation")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), C.CString(namespaceName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	defer qt.Recovering("connect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	defer qt.Recovering("disconnect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "warning")
	}
}

//export callbackQHelpEngineCoreWarning
func callbackQHelpEngineCoreWarning(ptr unsafe.Pointer, ptrName *C.char, msg *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::warning")

	if signal := qt.GetSignal(C.GoString(ptrName), "warning"); signal != nil {
		signal.(func(string))(C.GoString(msg))
	}

}

func (ptr *QHelpEngineCore) Warning(msg string) {
	defer qt.Recovering("QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_Warning(ptr.Pointer(), C.CString(msg))
	}
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	defer qt.Recovering("QHelpEngineCore::~QHelpEngineCore")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpEngineCoreTimerEvent
func callbackQHelpEngineCoreTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngineCore) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpEngineCore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpEngineCoreChildEvent
func callbackQHelpEngineCoreChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngineCore) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpEngineCore) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpEngineCoreCustomEvent
func callbackQHelpEngineCoreCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpEngineCore::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpEngineCoreFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpEngineCore) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpEngineCore) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpIndexModel struct {
	core.QStringListModel
}

type QHelpIndexModel_ITF interface {
	core.QStringListModel_ITF
	QHelpIndexModel_PTR() *QHelpIndexModel
}

func PointerFromQHelpIndexModel(ptr QHelpIndexModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexModel_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = new(QHelpIndexModel)
	n.SetPointer(ptr)
	return n
}

func newQHelpIndexModelFromPointer(ptr unsafe.Pointer) *QHelpIndexModel {
	var n = NewQHelpIndexModelFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpIndexModel_") {
		n.SetObjectName("QHelpIndexModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpIndexModel) QHelpIndexModel_PTR() *QHelpIndexModel {
	return ptr
}

func (ptr *QHelpIndexModel) CreateIndex(customFilterName string) {
	defer qt.Recovering("QHelpIndexModel::createIndex")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CreateIndex(ptr.Pointer(), C.CString(customFilterName))
	}
}

func (ptr *QHelpIndexModel) ConnectIndexCreated(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreated", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreated() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreated")
	}
}

//export callbackQHelpIndexModelIndexCreated
func callbackQHelpIndexModelIndexCreated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) IndexCreated() {
	defer qt.Recovering("QHelpIndexModel::indexCreated")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreated(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ConnectIndexCreationStarted(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ConnectIndexCreationStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexCreationStarted", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectIndexCreationStarted() {
	defer qt.Recovering("disconnect QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_DisconnectIndexCreationStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexCreationStarted")
	}
}

//export callbackQHelpIndexModelIndexCreationStarted
func callbackQHelpIndexModelIndexCreationStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpIndexModel::indexCreationStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexCreationStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpIndexModel) IndexCreationStarted() {
	defer qt.Recovering("QHelpIndexModel::indexCreationStarted")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_IndexCreationStarted(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) IsCreatingIndex() bool {
	defer qt.Recovering("QHelpIndexModel::isCreatingIndex")

	if ptr.Pointer() != nil {
		return C.QHelpIndexModel_IsCreatingIndex(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpIndexModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sort", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectSort() {
	defer qt.Recovering("disconnect QHelpIndexModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sort")
	}
}

//export callbackQHelpIndexModelSort
func callbackQHelpIndexModelSort(ptr unsafe.Pointer, ptrName *C.char, column C.int, order C.int) {
	defer qt.Recovering("callback QHelpIndexModel::sort")

	if signal := qt.GetSignal(C.GoString(ptrName), "sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(column), core.Qt__SortOrder(order))
	} else {
		NewQHelpIndexModelFromPointer(ptr).SortDefault(int(column), core.Qt__SortOrder(order))
	}
}

func (ptr *QHelpIndexModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpIndexModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHelpIndexModel::sort")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_SortDefault(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QHelpIndexModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "revert", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QHelpIndexModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "revert")
	}
}

//export callbackQHelpIndexModelRevert
func callbackQHelpIndexModelRevert(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexModel::revert")

	if signal := qt.GetSignal(C.GoString(ptrName), "revert"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexModel) Revert() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_Revert(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) RevertDefault() {
	defer qt.Recovering("QHelpIndexModel::revert")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_RevertDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpIndexModelTimerEvent
func callbackQHelpIndexModelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpIndexModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpIndexModelChildEvent
func callbackQHelpIndexModelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpIndexModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpIndexModelCustomEvent
func callbackQHelpIndexModelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexModel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexModel::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpIndexWidget struct {
	widgets.QListView
}

type QHelpIndexWidget_ITF interface {
	widgets.QListView_ITF
	QHelpIndexWidget_PTR() *QHelpIndexWidget
}

func PointerFromQHelpIndexWidget(ptr QHelpIndexWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpIndexWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = new(QHelpIndexWidget)
	n.SetPointer(ptr)
	return n
}

func newQHelpIndexWidgetFromPointer(ptr unsafe.Pointer) *QHelpIndexWidget {
	var n = NewQHelpIndexWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpIndexWidget_") {
		n.SetObjectName("QHelpIndexWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpIndexWidget) QHelpIndexWidget_PTR() *QHelpIndexWidget {
	return ptr
}

func (ptr *QHelpIndexWidget) ActivateCurrentItem() {
	defer qt.Recovering("QHelpIndexWidget::activateCurrentItem")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActivateCurrentItem(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) FilterIndices(filter string, wildcard string) {
	defer qt.Recovering("QHelpIndexWidget::filterIndices")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FilterIndices(ptr.Pointer(), C.CString(filter), C.CString(wildcard))
	}
}

func (ptr *QHelpIndexWidget) ConnectLinkActivated(f func(link *core.QUrl, keyword string)) {
	defer qt.Recovering("connect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQHelpIndexWidgetLinkActivated
func callbackQHelpIndexWidgetLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer, keyword *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(*core.QUrl, string))(core.NewQUrlFromPointer(link), C.GoString(keyword))
	}

}

func (ptr *QHelpIndexWidget) LinkActivated(link core.QUrl_ITF, keyword string) {
	defer qt.Recovering("QHelpIndexWidget::linkActivated")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LinkActivated(ptr.Pointer(), core.PointerFromQUrl(link), C.CString(keyword))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpIndexWidgetDragLeaveEvent
func callbackQHelpIndexWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpIndexWidgetDragMoveEvent
func callbackQHelpIndexWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpIndexWidgetDropEvent
func callbackQHelpIndexWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpIndexWidgetMouseMoveEvent
func callbackQHelpIndexWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpIndexWidgetMouseReleaseEvent
func callbackQHelpIndexWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpIndexWidgetPaintEvent
func callbackQHelpIndexWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpIndexWidgetResizeEvent
func callbackQHelpIndexWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQHelpIndexWidgetSetSelection
func callbackQHelpIndexWidgetSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHelpIndexWidget::setSelection")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QHelpIndexWidget) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQHelpIndexWidgetStartDrag
func callbackQHelpIndexWidgetStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHelpIndexWidget::startDrag")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHelpIndexWidget) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpIndexWidgetTimerEvent
func callbackQHelpIndexWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QHelpIndexWidget) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQHelpIndexWidgetUpdateGeometries
func callbackQHelpIndexWidgetUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexWidget::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) UpdateGeometries() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) UpdateGeometriesDefault() {
	defer qt.Recovering("QHelpIndexWidget::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEditor(f func(editor *widgets.QWidget, hint widgets.QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQHelpIndexWidgetCloseEditor
func callbackQHelpIndexWidgetCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QHelpIndexWidget::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*widgets.QWidget, widgets.QAbstractItemDelegate__EndEditHint))(widgets.NewQWidgetFromPointer(editor), widgets.QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) CloseEditor(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) CloseEditorDefault(editor widgets.QWidget_ITF, hint widgets.QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHelpIndexWidget::closeEditor")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHelpIndexWidget) ConnectCommitData(f func(editor *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQHelpIndexWidgetCommitData
func callbackQHelpIndexWidgetCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpIndexWidget::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) CommitData(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitData(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) CommitDataDefault(editor widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::commitData")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CommitDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpIndexWidgetDragEnterEvent
func callbackQHelpIndexWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQHelpIndexWidgetEditorDestroyed
func callbackQHelpIndexWidgetEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpIndexWidget::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHelpIndexWidget::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpIndexWidgetFocusInEvent
func callbackQHelpIndexWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpIndexWidgetFocusOutEvent
func callbackQHelpIndexWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpIndexWidgetInputMethodEvent
func callbackQHelpIndexWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpIndexWidgetKeyPressEvent
func callbackQHelpIndexWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQHelpIndexWidgetKeyboardSearch
func callbackQHelpIndexWidgetKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QHelpIndexWidget::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearch(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpIndexWidget) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHelpIndexWidget::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHelpIndexWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpIndexWidgetMouseDoubleClickEvent
func callbackQHelpIndexWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpIndexWidgetMousePressEvent
func callbackQHelpIndexWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectReset(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectReset() {
	defer qt.Recovering("disconnect QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQHelpIndexWidgetReset
func callbackQHelpIndexWidgetReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexWidget::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) Reset() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_Reset(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ResetDefault() {
	defer qt.Recovering("QHelpIndexWidget::reset")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQHelpIndexWidgetSelectAll
func callbackQHelpIndexWidgetSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHelpIndexWidget::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) SelectAll() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) SelectAllDefault() {
	defer qt.Recovering("QHelpIndexWidget::selectAll")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QHelpIndexWidget) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQHelpIndexWidgetSetModel
func callbackQHelpIndexWidgetSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHelpIndexWidget) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQHelpIndexWidgetSetSelectionModel
func callbackQHelpIndexWidgetSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHelpIndexWidget) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpIndexWidgetContextMenuEvent
func callbackQHelpIndexWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQHelpIndexWidgetScrollContentsBy
func callbackQHelpIndexWidgetScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHelpIndexWidget::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpIndexWidget) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHelpIndexWidget::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetupViewport(f func(viewport *widgets.QWidget)) {
	defer qt.Recovering("connect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQHelpIndexWidgetSetupViewport
func callbackQHelpIndexWidgetSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*widgets.QWidget))(widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).SetupViewportDefault(widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewport(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewport(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) SetupViewportDefault(viewport widgets.QWidget_ITF) {
	defer qt.Recovering("QHelpIndexWidget::setupViewport")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetupViewportDefault(ptr.Pointer(), widgets.PointerFromQWidget(viewport))
	}
}

func (ptr *QHelpIndexWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpIndexWidgetWheelEvent
func callbackQHelpIndexWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHelpIndexWidget) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQHelpIndexWidgetChangeEvent
func callbackQHelpIndexWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHelpIndexWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpIndexWidgetActionEvent
func callbackQHelpIndexWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpIndexWidgetEnterEvent
func callbackQHelpIndexWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpIndexWidgetHideEvent
func callbackQHelpIndexWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpIndexWidgetLeaveEvent
func callbackQHelpIndexWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpIndexWidgetMoveEvent
func callbackQHelpIndexWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpIndexWidgetSetVisible
func callbackQHelpIndexWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpIndexWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpIndexWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpIndexWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpIndexWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpIndexWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpIndexWidgetShowEvent
func callbackQHelpIndexWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpIndexWidgetCloseEvent
func callbackQHelpIndexWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpIndexWidgetInitPainter
func callbackQHelpIndexWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpIndexWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpIndexWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpIndexWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpIndexWidgetKeyReleaseEvent
func callbackQHelpIndexWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpIndexWidgetTabletEvent
func callbackQHelpIndexWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpIndexWidgetChildEvent
func callbackQHelpIndexWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpIndexWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpIndexWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpIndexWidgetCustomEvent
func callbackQHelpIndexWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpIndexWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpIndexWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpIndexWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpIndexWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpIndexWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpSearchEngine struct {
	core.QObject
}

type QHelpSearchEngine_ITF interface {
	core.QObject_ITF
	QHelpSearchEngine_PTR() *QHelpSearchEngine
}

func PointerFromQHelpSearchEngine(ptr QHelpSearchEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchEngine_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = new(QHelpSearchEngine)
	n.SetPointer(ptr)
	return n
}

func newQHelpSearchEngineFromPointer(ptr unsafe.Pointer) *QHelpSearchEngine {
	var n = NewQHelpSearchEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchEngine_") {
		n.SetObjectName("QHelpSearchEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchEngine) QHelpSearchEngine_PTR() *QHelpSearchEngine {
	return ptr
}

func NewQHelpSearchEngine(helpEngine QHelpEngineCore_ITF, parent core.QObject_ITF) *QHelpSearchEngine {
	defer qt.Recovering("QHelpSearchEngine::QHelpSearchEngine")

	return newQHelpSearchEngineFromPointer(C.QHelpSearchEngine_NewQHelpSearchEngine(PointerFromQHelpEngineCore(helpEngine), core.PointerFromQObject(parent)))
}

func (ptr *QHelpSearchEngine) CancelIndexing() {
	defer qt.Recovering("QHelpSearchEngine::cancelIndexing")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelIndexing(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) CancelSearching() {
	defer qt.Recovering("QHelpSearchEngine::cancelSearching")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CancelSearching(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) HitCount() int {
	defer qt.Recovering("QHelpSearchEngine::hitCount")

	if ptr.Pointer() != nil {
		return int(C.QHelpSearchEngine_HitCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHelpSearchEngine) ConnectIndexingFinished(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingFinished")
	}
}

//export callbackQHelpSearchEngineIndexingFinished
func callbackQHelpSearchEngineIndexingFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexingFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) IndexingFinished() {
	defer qt.Recovering("QHelpSearchEngine::indexingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingFinished(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ConnectIndexingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectIndexingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectIndexingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectIndexingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexingStarted")
	}
}

//export callbackQHelpSearchEngineIndexingStarted
func callbackQHelpSearchEngineIndexingStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::indexingStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) IndexingStarted() {
	defer qt.Recovering("QHelpSearchEngine::indexingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_IndexingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) QueryWidget() *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchEngine::queryWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchEngine_QueryWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ReindexDocumentation() {
	defer qt.Recovering("QHelpSearchEngine::reindexDocumentation")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ReindexDocumentation(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) ResultWidget() *QHelpSearchResultWidget {
	defer qt.Recovering("QHelpSearchEngine::resultWidget")

	if ptr.Pointer() != nil {
		return NewQHelpSearchResultWidgetFromPointer(C.QHelpSearchEngine_ResultWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHelpSearchEngine) ConnectSearchingFinished(f func(hits int)) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingFinished", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingFinished() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingFinished")
	}
}

//export callbackQHelpSearchEngineSearchingFinished
func callbackQHelpSearchEngineSearchingFinished(ptr unsafe.Pointer, ptrName *C.char, hits C.int) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "searchingFinished"); signal != nil {
		signal.(func(int))(int(hits))
	}

}

func (ptr *QHelpSearchEngine) SearchingFinished(hits int) {
	defer qt.Recovering("QHelpSearchEngine::searchingFinished")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingFinished(ptr.Pointer(), C.int(hits))
	}
}

func (ptr *QHelpSearchEngine) ConnectSearchingStarted(f func()) {
	defer qt.Recovering("connect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ConnectSearchingStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "searchingStarted", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectSearchingStarted() {
	defer qt.Recovering("disconnect QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DisconnectSearchingStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "searchingStarted")
	}
}

//export callbackQHelpSearchEngineSearchingStarted
func callbackQHelpSearchEngineSearchingStarted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchEngine::searchingStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "searchingStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchEngine) SearchingStarted() {
	defer qt.Recovering("QHelpSearchEngine::searchingStarted")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_SearchingStarted(ptr.Pointer())
	}
}

func (ptr *QHelpSearchEngine) DestroyQHelpSearchEngine() {
	defer qt.Recovering("QHelpSearchEngine::~QHelpSearchEngine")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_DestroyQHelpSearchEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpSearchEngineTimerEvent
func callbackQHelpSearchEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpSearchEngineChildEvent
func callbackQHelpSearchEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpSearchEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpSearchEngineCustomEvent
func callbackQHelpSearchEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (p *QHelpSearchQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpSearchQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return n
}

func newQHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = NewQHelpSearchQueryFromPointer(ptr)
	return n
}

func (ptr *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return ptr
}

//QHelpSearchQuery::FieldName
type QHelpSearchQuery__FieldName int64

const (
	QHelpSearchQuery__DEFAULT = QHelpSearchQuery__FieldName(0)
	QHelpSearchQuery__FUZZY   = QHelpSearchQuery__FieldName(1)
	QHelpSearchQuery__WITHOUT = QHelpSearchQuery__FieldName(2)
	QHelpSearchQuery__PHRASE  = QHelpSearchQuery__FieldName(3)
	QHelpSearchQuery__ALL     = QHelpSearchQuery__FieldName(4)
	QHelpSearchQuery__ATLEAST = QHelpSearchQuery__FieldName(5)
)

func NewQHelpSearchQuery() *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	return newQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
}

func NewQHelpSearchQuery2(field QHelpSearchQuery__FieldName, wordList []string) *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	return newQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery2(C.int(field), C.CString(strings.Join(wordList, "|"))))
}

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	return n
}

func newQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = NewQHelpSearchQueryWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchQueryWidget_") {
		n.SetObjectName("QHelpSearchQueryWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	defer qt.Recovering("QHelpSearchQueryWidget::isCompactMode")

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	defer qt.Recovering("QHelpSearchQueryWidget::QHelpSearchQueryWidget")

	return newQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::collapseExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	defer qt.Recovering("QHelpSearchQueryWidget::expandExtendedSearch")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "search")
	}
}

//export callbackQHelpSearchQueryWidgetSearch
func callbackQHelpSearchQueryWidgetSearch(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::search")

	if signal := qt.GetSignal(C.GoString(ptrName), "search"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpSearchQueryWidget) Search() {
	defer qt.Recovering("QHelpSearchQueryWidget::search")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_Search(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	defer qt.Recovering("QHelpSearchQueryWidget::~QHelpSearchQueryWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpSearchQueryWidgetActionEvent
func callbackQHelpSearchQueryWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragEnterEvent
func callbackQHelpSearchQueryWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragLeaveEvent
func callbackQHelpSearchQueryWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDragMoveEvent
func callbackQHelpSearchQueryWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpSearchQueryWidgetDropEvent
func callbackQHelpSearchQueryWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpSearchQueryWidgetEnterEvent
func callbackQHelpSearchQueryWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpSearchQueryWidgetFocusOutEvent
func callbackQHelpSearchQueryWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpSearchQueryWidgetHideEvent
func callbackQHelpSearchQueryWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetLeaveEvent
func callbackQHelpSearchQueryWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMoveEvent
func callbackQHelpSearchQueryWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpSearchQueryWidgetPaintEvent
func callbackQHelpSearchQueryWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpSearchQueryWidgetSetVisible
func callbackQHelpSearchQueryWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpSearchQueryWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpSearchQueryWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchQueryWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpSearchQueryWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpSearchQueryWidgetShowEvent
func callbackQHelpSearchQueryWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpSearchQueryWidgetCloseEvent
func callbackQHelpSearchQueryWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpSearchQueryWidgetContextMenuEvent
func callbackQHelpSearchQueryWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpSearchQueryWidgetInitPainter
func callbackQHelpSearchQueryWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchQueryWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchQueryWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpSearchQueryWidgetInputMethodEvent
func callbackQHelpSearchQueryWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpSearchQueryWidgetKeyPressEvent
func callbackQHelpSearchQueryWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpSearchQueryWidgetKeyReleaseEvent
func callbackQHelpSearchQueryWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseDoubleClickEvent
func callbackQHelpSearchQueryWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseMoveEvent
func callbackQHelpSearchQueryWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMousePressEvent
func callbackQHelpSearchQueryWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpSearchQueryWidgetMouseReleaseEvent
func callbackQHelpSearchQueryWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpSearchQueryWidgetResizeEvent
func callbackQHelpSearchQueryWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpSearchQueryWidgetTabletEvent
func callbackQHelpSearchQueryWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpSearchQueryWidgetWheelEvent
func callbackQHelpSearchQueryWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpSearchQueryWidgetTimerEvent
func callbackQHelpSearchQueryWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpSearchQueryWidgetChildEvent
func callbackQHelpSearchQueryWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpSearchQueryWidgetCustomEvent
func callbackQHelpSearchQueryWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchQueryWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchQueryWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchQueryWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchQueryWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHelpSearchResultWidget struct {
	widgets.QWidget
}

type QHelpSearchResultWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget
}

func PointerFromQHelpSearchResultWidget(ptr QHelpSearchResultWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchResultWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = new(QHelpSearchResultWidget)
	n.SetPointer(ptr)
	return n
}

func newQHelpSearchResultWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchResultWidget {
	var n = NewQHelpSearchResultWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchResultWidget_") {
		n.SetObjectName("QHelpSearchResultWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpSearchResultWidget) QHelpSearchResultWidget_PTR() *QHelpSearchResultWidget {
	return ptr
}

func (ptr *QHelpSearchResultWidget) LinkAt(point core.QPoint_ITF) *core.QUrl {
	defer qt.Recovering("QHelpSearchResultWidget::linkAt")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpSearchResultWidget_LinkAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QHelpSearchResultWidget) ConnectRequestShowLink(f func(link *core.QUrl)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ConnectRequestShowLink(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestShowLink", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectRequestShowLink() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DisconnectRequestShowLink(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestShowLink")
	}
}

//export callbackQHelpSearchResultWidgetRequestShowLink
func callbackQHelpSearchResultWidgetRequestShowLink(ptr unsafe.Pointer, ptrName *C.char, link unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::requestShowLink")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestShowLink"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(link))
	}

}

func (ptr *QHelpSearchResultWidget) RequestShowLink(link core.QUrl_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::requestShowLink")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_RequestShowLink(ptr.Pointer(), core.PointerFromQUrl(link))
	}
}

func (ptr *QHelpSearchResultWidget) DestroyQHelpSearchResultWidget() {
	defer qt.Recovering("QHelpSearchResultWidget::~QHelpSearchResultWidget")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpSearchResultWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHelpSearchResultWidgetActionEvent
func callbackQHelpSearchResultWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragEnterEvent
func callbackQHelpSearchResultWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragLeaveEvent
func callbackQHelpSearchResultWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHelpSearchResultWidgetDragMoveEvent
func callbackQHelpSearchResultWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHelpSearchResultWidgetDropEvent
func callbackQHelpSearchResultWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHelpSearchResultWidgetEnterEvent
func callbackQHelpSearchResultWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHelpSearchResultWidgetFocusInEvent
func callbackQHelpSearchResultWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHelpSearchResultWidgetFocusOutEvent
func callbackQHelpSearchResultWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHelpSearchResultWidgetHideEvent
func callbackQHelpSearchResultWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHelpSearchResultWidgetLeaveEvent
func callbackQHelpSearchResultWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHelpSearchResultWidgetMoveEvent
func callbackQHelpSearchResultWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHelpSearchResultWidgetPaintEvent
func callbackQHelpSearchResultWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHelpSearchResultWidgetSetVisible
func callbackQHelpSearchResultWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QHelpSearchResultWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QHelpSearchResultWidget) SetVisible(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchResultWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QHelpSearchResultWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHelpSearchResultWidgetShowEvent
func callbackQHelpSearchResultWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHelpSearchResultWidgetCloseEvent
func callbackQHelpSearchResultWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHelpSearchResultWidgetContextMenuEvent
func callbackQHelpSearchResultWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHelpSearchResultWidgetInitPainter
func callbackQHelpSearchResultWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchResultWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHelpSearchResultWidgetInputMethodEvent
func callbackQHelpSearchResultWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHelpSearchResultWidgetKeyPressEvent
func callbackQHelpSearchResultWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHelpSearchResultWidgetKeyReleaseEvent
func callbackQHelpSearchResultWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseDoubleClickEvent
func callbackQHelpSearchResultWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseMoveEvent
func callbackQHelpSearchResultWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHelpSearchResultWidgetMousePressEvent
func callbackQHelpSearchResultWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHelpSearchResultWidgetMouseReleaseEvent
func callbackQHelpSearchResultWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHelpSearchResultWidgetResizeEvent
func callbackQHelpSearchResultWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHelpSearchResultWidgetTabletEvent
func callbackQHelpSearchResultWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHelpSearchResultWidgetWheelEvent
func callbackQHelpSearchResultWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpSearchResultWidgetTimerEvent
func callbackQHelpSearchResultWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpSearchResultWidgetChildEvent
func callbackQHelpSearchResultWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpSearchResultWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpSearchResultWidgetCustomEvent
func callbackQHelpSearchResultWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHelpSearchResultWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHelpSearchResultWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHelpSearchResultWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHelpSearchResultWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QHelpSearchResultWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
