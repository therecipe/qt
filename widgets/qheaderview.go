package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QHeaderView struct {
	QAbstractItemView
}

type QHeaderView_ITF interface {
	QAbstractItemView_ITF
	QHeaderView_PTR() *QHeaderView
}

func PointerFromQHeaderView(ptr QHeaderView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHeaderView_PTR().Pointer()
	}
	return nil
}

func NewQHeaderViewFromPointer(ptr unsafe.Pointer) *QHeaderView {
	var n = new(QHeaderView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHeaderView_") {
		n.SetObjectName("QHeaderView_" + qt.Identifier())
	}
	return n
}

func (ptr *QHeaderView) QHeaderView_PTR() *QHeaderView {
	return ptr
}

//QHeaderView::ResizeMode
type QHeaderView__ResizeMode int64

const (
	QHeaderView__Interactive      = QHeaderView__ResizeMode(0)
	QHeaderView__Stretch          = QHeaderView__ResizeMode(1)
	QHeaderView__Fixed            = QHeaderView__ResizeMode(2)
	QHeaderView__ResizeToContents = QHeaderView__ResizeMode(3)
	QHeaderView__Custom           = QHeaderView__ResizeMode(QHeaderView__Fixed)
)

func (ptr *QHeaderView) CascadingSectionResizes() bool {
	defer qt.Recovering("QHeaderView::cascadingSectionResizes")

	if ptr.Pointer() != nil {
		return C.QHeaderView_CascadingSectionResizes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) DefaultAlignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QHeaderView::defaultAlignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QHeaderView_DefaultAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) DefaultSectionSize() int {
	defer qt.Recovering("QHeaderView::defaultSectionSize")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_DefaultSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) HighlightSections() bool {
	defer qt.Recovering("QHeaderView::highlightSections")

	if ptr.Pointer() != nil {
		return C.QHeaderView_HighlightSections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) IsSortIndicatorShown() bool {
	defer qt.Recovering("QHeaderView::isSortIndicatorShown")

	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSortIndicatorShown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) MaximumSectionSize() int {
	defer qt.Recovering("QHeaderView::maximumSectionSize")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MaximumSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) MinimumSectionSize() int {
	defer qt.Recovering("QHeaderView::minimumSectionSize")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MinimumSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ResetDefaultSectionSize() {
	defer qt.Recovering("QHeaderView::resetDefaultSectionSize")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResetDefaultSectionSize(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ResizeSection(logicalIndex int, size int) {
	defer qt.Recovering("QHeaderView::resizeSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSection(ptr.Pointer(), C.int(logicalIndex), C.int(size))
	}
}

func (ptr *QHeaderView) SetCascadingSectionResizes(enable bool) {
	defer qt.Recovering("QHeaderView::setCascadingSectionResizes")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetCascadingSectionResizes(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QHeaderView) SetDefaultAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QHeaderView::setDefaultAlignment")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QHeaderView) SetDefaultSectionSize(size int) {
	defer qt.Recovering("QHeaderView::setDefaultSectionSize")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetHighlightSections(highlight bool) {
	defer qt.Recovering("QHeaderView::setHighlightSections")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetHighlightSections(ptr.Pointer(), C.int(qt.GoBoolToInt(highlight)))
	}
}

func (ptr *QHeaderView) SetMaximumSectionSize(size int) {
	defer qt.Recovering("QHeaderView::setMaximumSectionSize")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetMaximumSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetMinimumSectionSize(size int) {
	defer qt.Recovering("QHeaderView::setMinimumSectionSize")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetMinimumSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetOffset(offset int) {
	defer qt.Recovering("QHeaderView::setOffset")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffset(ptr.Pointer(), C.int(offset))
	}
}

func (ptr *QHeaderView) SetSortIndicatorShown(show bool) {
	defer qt.Recovering("QHeaderView::setSortIndicatorShown")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QHeaderView) SetStretchLastSection(stretch bool) {
	defer qt.Recovering("QHeaderView::setStretchLastSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetStretchLastSection(ptr.Pointer(), C.int(qt.GoBoolToInt(stretch)))
	}
}

func (ptr *QHeaderView) StretchLastSection() bool {
	defer qt.Recovering("QHeaderView::stretchLastSection")

	if ptr.Pointer() != nil {
		return C.QHeaderView_StretchLastSection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) Count() int {
	defer qt.Recovering("QHeaderView::count")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ConnectCurrentChanged(f func(current *core.QModelIndex, old *core.QModelIndex)) {
	defer qt.Recovering("connect QHeaderView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QHeaderView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQHeaderViewCurrentChanged
func callbackQHeaderViewCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, old unsafe.Pointer) bool {
	defer qt.Recovering("callback QHeaderView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(old))
		return true
	}
	return false

}

func (ptr *QHeaderView) CurrentChanged(current core.QModelIndex_ITF, old core.QModelIndex_ITF) {
	defer qt.Recovering("QHeaderView::currentChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(old))
	}
}

func (ptr *QHeaderView) CurrentChangedDefault(current core.QModelIndex_ITF, old core.QModelIndex_ITF) {
	defer qt.Recovering("QHeaderView::currentChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(old))
	}
}

func (ptr *QHeaderView) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHeaderView::event")

	if ptr.Pointer() != nil {
		return C.QHeaderView_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHeaderView) ConnectGeometriesChanged(f func()) {
	defer qt.Recovering("connect QHeaderView::geometriesChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectGeometriesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometriesChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectGeometriesChanged() {
	defer qt.Recovering("disconnect QHeaderView::geometriesChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectGeometriesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometriesChanged")
	}
}

//export callbackQHeaderViewGeometriesChanged
func callbackQHeaderViewGeometriesChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QHeaderView::geometriesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometriesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHeaderView) GeometriesChanged() {
	defer qt.Recovering("QHeaderView::geometriesChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_GeometriesChanged(ptr.Pointer())
	}
}

func (ptr *QHeaderView) HeaderDataChanged(orientation core.Qt__Orientation, logicalFirst int, logicalLast int) {
	defer qt.Recovering("QHeaderView::headerDataChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_HeaderDataChanged(ptr.Pointer(), C.int(orientation), C.int(logicalFirst), C.int(logicalLast))
	}
}

func (ptr *QHeaderView) HiddenSectionCount() int {
	defer qt.Recovering("QHeaderView::hiddenSectionCount")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_HiddenSectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) HideSection(logicalIndex int) {
	defer qt.Recovering("QHeaderView::hideSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_HideSection(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) HorizontalOffset() int {
	defer qt.Recovering("QHeaderView::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_HorizontalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	defer qt.Recovering("QHeaderView::isSectionHidden")

	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSectionHidden(ptr.Pointer(), C.int(logicalIndex)) != 0
	}
	return false
}

func (ptr *QHeaderView) Length() int {
	defer qt.Recovering("QHeaderView::length")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndex(visualIndex int) int {
	defer qt.Recovering("QHeaderView::logicalIndex")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndex(ptr.Pointer(), C.int(visualIndex)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt3(pos core.QPoint_ITF) int {
	defer qt.Recovering("QHeaderView::logicalIndexAt")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt3(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt(position int) int {
	defer qt.Recovering("QHeaderView::logicalIndexAt")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt(ptr.Pointer(), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt2(x int, y int) int {
	defer qt.Recovering("QHeaderView::logicalIndexAt")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHeaderView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QHeaderView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQHeaderViewMouseDoubleClickEvent
func callbackQHeaderViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHeaderView) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHeaderView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QHeaderView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQHeaderViewMouseMoveEvent
func callbackQHeaderViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHeaderView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHeaderView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QHeaderView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQHeaderViewMousePressEvent
func callbackQHeaderViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHeaderView) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QHeaderView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QHeaderView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQHeaderViewMouseReleaseEvent
func callbackQHeaderViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QHeaderView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QHeaderView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QHeaderView) MoveSection(from int, to int) {
	defer qt.Recovering("QHeaderView::moveSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_MoveSection(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QHeaderView) Offset() int {
	defer qt.Recovering("QHeaderView::offset")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Offset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QHeaderView::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHeaderView_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QHeaderView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QHeaderView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQHeaderViewPaintEvent
func callbackQHeaderViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QHeaderView) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHeaderView::paintEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHeaderView) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QHeaderView::paintEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QHeaderView) ConnectPaintSection(f func(painter *gui.QPainter, rect *core.QRect, logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::paintSection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintSection", f)
	}
}

func (ptr *QHeaderView) DisconnectPaintSection() {
	defer qt.Recovering("disconnect QHeaderView::paintSection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintSection")
	}
}

//export callbackQHeaderViewPaintSection
func callbackQHeaderViewPaintSection(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, rect unsafe.Pointer, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::paintSection")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintSection"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRect, int))(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), int(logicalIndex))
	} else {
		NewQHeaderViewFromPointer(ptr).PaintSectionDefault(gui.NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), int(logicalIndex))
	}
}

func (ptr *QHeaderView) PaintSection(painter gui.QPainter_ITF, rect core.QRect_ITF, logicalIndex int) {
	defer qt.Recovering("QHeaderView::paintSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_PaintSection(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) PaintSectionDefault(painter gui.QPainter_ITF, rect core.QRect_ITF, logicalIndex int) {
	defer qt.Recovering("QHeaderView::paintSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_PaintSectionDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRect(rect), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectReset(f func()) {
	defer qt.Recovering("connect QHeaderView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QHeaderView) DisconnectReset() {
	defer qt.Recovering("disconnect QHeaderView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQHeaderViewReset
func callbackQHeaderViewReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHeaderView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHeaderView) Reset() {
	defer qt.Recovering("QHeaderView::reset")

	if ptr.Pointer() != nil {
		C.QHeaderView_Reset(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ResetDefault() {
	defer qt.Recovering("QHeaderView::reset")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ResizeContentsPrecision() int {
	defer qt.Recovering("QHeaderView::resizeContentsPrecision")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_ResizeContentsPrecision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ResizeSections(mode QHeaderView__ResizeMode) {
	defer qt.Recovering("QHeaderView::resizeSections")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSections(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QHeaderView) RestoreState(state core.QByteArray_ITF) bool {
	defer qt.Recovering("QHeaderView::restoreState")

	if ptr.Pointer() != nil {
		return C.QHeaderView_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QHeaderView) SaveState() *core.QByteArray {
	defer qt.Recovering("QHeaderView::saveState")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHeaderView_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHeaderView) ConnectSectionClicked(f func(logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionClicked() {
	defer qt.Recovering("disconnect QHeaderView::sectionClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionClicked")
	}
}

//export callbackQHeaderViewSectionClicked
func callbackQHeaderViewSectionClicked(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionClicked"); signal != nil {
		signal.(func(int))(int(logicalIndex))
	}

}

func (ptr *QHeaderView) SectionClicked(logicalIndex int) {
	defer qt.Recovering("QHeaderView::sectionClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionClicked(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSectionCountChanged(f func(oldCount int, newCount int)) {
	defer qt.Recovering("connect QHeaderView::sectionCountChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionCountChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionCountChanged() {
	defer qt.Recovering("disconnect QHeaderView::sectionCountChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionCountChanged")
	}
}

//export callbackQHeaderViewSectionCountChanged
func callbackQHeaderViewSectionCountChanged(ptr unsafe.Pointer, ptrName *C.char, oldCount C.int, newCount C.int) {
	defer qt.Recovering("callback QHeaderView::sectionCountChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionCountChanged"); signal != nil {
		signal.(func(int, int))(int(oldCount), int(newCount))
	}

}

func (ptr *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	defer qt.Recovering("QHeaderView::sectionCountChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionCountChanged(ptr.Pointer(), C.int(oldCount), C.int(newCount))
	}
}

func (ptr *QHeaderView) ConnectSectionDoubleClicked(f func(logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionDoubleClicked() {
	defer qt.Recovering("disconnect QHeaderView::sectionDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionDoubleClicked")
	}
}

//export callbackQHeaderViewSectionDoubleClicked
func callbackQHeaderViewSectionDoubleClicked(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionDoubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionDoubleClicked"); signal != nil {
		signal.(func(int))(int(logicalIndex))
	}

}

func (ptr *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	defer qt.Recovering("QHeaderView::sectionDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionDoubleClicked(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSectionEntered(f func(logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionEntered")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionEntered", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionEntered() {
	defer qt.Recovering("disconnect QHeaderView::sectionEntered")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionEntered")
	}
}

//export callbackQHeaderViewSectionEntered
func callbackQHeaderViewSectionEntered(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionEntered")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionEntered"); signal != nil {
		signal.(func(int))(int(logicalIndex))
	}

}

func (ptr *QHeaderView) SectionEntered(logicalIndex int) {
	defer qt.Recovering("QHeaderView::sectionEntered")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionEntered(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSectionHandleDoubleClicked(f func(logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionHandleDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionHandleDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionHandleDoubleClicked() {
	defer qt.Recovering("disconnect QHeaderView::sectionHandleDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionHandleDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked")
	}
}

//export callbackQHeaderViewSectionHandleDoubleClicked
func callbackQHeaderViewSectionHandleDoubleClicked(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionHandleDoubleClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionHandleDoubleClicked"); signal != nil {
		signal.(func(int))(int(logicalIndex))
	}

}

func (ptr *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	defer qt.Recovering("QHeaderView::sectionHandleDoubleClicked")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionHandleDoubleClicked(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSectionMoved(f func(logicalIndex int, oldVisualIndex int, newVisualIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionMoved")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionMoved", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionMoved() {
	defer qt.Recovering("disconnect QHeaderView::sectionMoved")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionMoved")
	}
}

//export callbackQHeaderViewSectionMoved
func callbackQHeaderViewSectionMoved(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int, oldVisualIndex C.int, newVisualIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionMoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionMoved"); signal != nil {
		signal.(func(int, int, int))(int(logicalIndex), int(oldVisualIndex), int(newVisualIndex))
	}

}

func (ptr *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	defer qt.Recovering("QHeaderView::sectionMoved")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionMoved(ptr.Pointer(), C.int(logicalIndex), C.int(oldVisualIndex), C.int(newVisualIndex))
	}
}

func (ptr *QHeaderView) SectionPosition(logicalIndex int) int {
	defer qt.Recovering("QHeaderView::sectionPosition")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionPosition(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionPressed(f func(logicalIndex int)) {
	defer qt.Recovering("connect QHeaderView::sectionPressed")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionPressed", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionPressed() {
	defer qt.Recovering("disconnect QHeaderView::sectionPressed")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionPressed")
	}
}

//export callbackQHeaderViewSectionPressed
func callbackQHeaderViewSectionPressed(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int) {
	defer qt.Recovering("callback QHeaderView::sectionPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionPressed"); signal != nil {
		signal.(func(int))(int(logicalIndex))
	}

}

func (ptr *QHeaderView) SectionPressed(logicalIndex int) {
	defer qt.Recovering("QHeaderView::sectionPressed")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionPressed(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) SectionResizeMode(logicalIndex int) QHeaderView__ResizeMode {
	defer qt.Recovering("QHeaderView::sectionResizeMode")

	if ptr.Pointer() != nil {
		return QHeaderView__ResizeMode(C.QHeaderView_SectionResizeMode(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionResized(f func(logicalIndex int, oldSize int, newSize int)) {
	defer qt.Recovering("connect QHeaderView::sectionResized")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionResized", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionResized() {
	defer qt.Recovering("disconnect QHeaderView::sectionResized")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionResized")
	}
}

//export callbackQHeaderViewSectionResized
func callbackQHeaderViewSectionResized(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int, oldSize C.int, newSize C.int) {
	defer qt.Recovering("callback QHeaderView::sectionResized")

	if signal := qt.GetSignal(C.GoString(ptrName), "sectionResized"); signal != nil {
		signal.(func(int, int, int))(int(logicalIndex), int(oldSize), int(newSize))
	}

}

func (ptr *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	defer qt.Recovering("QHeaderView::sectionResized")

	if ptr.Pointer() != nil {
		C.QHeaderView_SectionResized(ptr.Pointer(), C.int(logicalIndex), C.int(oldSize), C.int(newSize))
	}
}

func (ptr *QHeaderView) SectionSize(logicalIndex int) int {
	defer qt.Recovering("QHeaderView::sectionSize")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSize(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionSizeFromContents(logicalIndex int) *core.QSize {
	defer qt.Recovering("QHeaderView::sectionSizeFromContents")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QHeaderView_SectionSizeFromContents(ptr.Pointer(), C.int(logicalIndex)))
	}
	return nil
}

func (ptr *QHeaderView) SectionSizeHint(logicalIndex int) int {
	defer qt.Recovering("QHeaderView::sectionSizeHint")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSizeHint(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	defer qt.Recovering("QHeaderView::sectionViewportPosition")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionViewportPosition(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionsClickable() bool {
	defer qt.Recovering("QHeaderView::sectionsClickable")

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsClickable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsHidden() bool {
	defer qt.Recovering("QHeaderView::sectionsHidden")

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMovable() bool {
	defer qt.Recovering("QHeaderView::sectionsMovable")

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMoved() bool {
	defer qt.Recovering("QHeaderView::sectionsMoved")

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMoved(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QHeaderView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QHeaderView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QHeaderView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQHeaderViewSetModel
func callbackQHeaderViewSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQHeaderViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QHeaderView) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHeaderView::setModel")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHeaderView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QHeaderView::setModel")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHeaderView) SetOffsetToLastSection() {
	defer qt.Recovering("QHeaderView::setOffsetToLastSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToLastSection(ptr.Pointer())
	}
}

func (ptr *QHeaderView) SetOffsetToSectionPosition(visualSectionNumber int) {
	defer qt.Recovering("QHeaderView::setOffsetToSectionPosition")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToSectionPosition(ptr.Pointer(), C.int(visualSectionNumber))
	}
}

func (ptr *QHeaderView) SetResizeContentsPrecision(precision int) {
	defer qt.Recovering("QHeaderView::setResizeContentsPrecision")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetResizeContentsPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	defer qt.Recovering("QHeaderView::setSectionHidden")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionHidden(ptr.Pointer(), C.int(logicalIndex), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode(mode QHeaderView__ResizeMode) {
	defer qt.Recovering("QHeaderView::setSectionResizeMode")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode2(logicalIndex int, mode QHeaderView__ResizeMode) {
	defer qt.Recovering("QHeaderView::setSectionResizeMode")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode2(ptr.Pointer(), C.int(logicalIndex), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionsClickable(clickable bool) {
	defer qt.Recovering("QHeaderView::setSectionsClickable")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsClickable(ptr.Pointer(), C.int(qt.GoBoolToInt(clickable)))
	}
}

func (ptr *QHeaderView) SetSectionsMovable(movable bool) {
	defer qt.Recovering("QHeaderView::setSectionsMovable")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QHeaderView) ConnectSetSelection(f func(rect *core.QRect, flags core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QHeaderView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QHeaderView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QHeaderView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQHeaderViewSetSelection
func callbackQHeaderViewSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, flags C.int) {
	defer qt.Recovering("callback QHeaderView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
	} else {
		NewQHeaderViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(flags))
	}
}

func (ptr *QHeaderView) SetSelection(rect core.QRect_ITF, flags core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHeaderView::setSelection")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(flags))
	}
}

func (ptr *QHeaderView) SetSelectionDefault(rect core.QRect_ITF, flags core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QHeaderView::setSelection")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(flags))
	}
}

func (ptr *QHeaderView) SetSortIndicator(logicalIndex int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHeaderView::setSortIndicator")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicator(ptr.Pointer(), C.int(logicalIndex), C.int(order))
	}
}

func (ptr *QHeaderView) ConnectSetVisible(f func(v bool)) {
	defer qt.Recovering("connect QHeaderView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QHeaderView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QHeaderView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQHeaderViewSetVisible
func callbackQHeaderViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, v C.int) bool {
	defer qt.Recovering("callback QHeaderView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(v) != 0)
		return true
	}
	return false

}

func (ptr *QHeaderView) SetVisible(v bool) {
	defer qt.Recovering("QHeaderView::setVisible")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QHeaderView) SetVisibleDefault(v bool) {
	defer qt.Recovering("QHeaderView::setVisible")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QHeaderView) ShowSection(logicalIndex int) {
	defer qt.Recovering("QHeaderView::showSection")

	if ptr.Pointer() != nil {
		C.QHeaderView_ShowSection(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) SizeHint() *core.QSize {
	defer qt.Recovering("QHeaderView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QHeaderView_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHeaderView) ConnectSortIndicatorChanged(f func(logicalIndex int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QHeaderView::sortIndicatorChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSortIndicatorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sortIndicatorChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSortIndicatorChanged() {
	defer qt.Recovering("disconnect QHeaderView::sortIndicatorChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSortIndicatorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sortIndicatorChanged")
	}
}

//export callbackQHeaderViewSortIndicatorChanged
func callbackQHeaderViewSortIndicatorChanged(ptr unsafe.Pointer, ptrName *C.char, logicalIndex C.int, order C.int) {
	defer qt.Recovering("callback QHeaderView::sortIndicatorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sortIndicatorChanged"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(logicalIndex), core.Qt__SortOrder(order))
	}

}

func (ptr *QHeaderView) SortIndicatorChanged(logicalIndex int, order core.Qt__SortOrder) {
	defer qt.Recovering("QHeaderView::sortIndicatorChanged")

	if ptr.Pointer() != nil {
		C.QHeaderView_SortIndicatorChanged(ptr.Pointer(), C.int(logicalIndex), C.int(order))
	}
}

func (ptr *QHeaderView) SortIndicatorOrder() core.Qt__SortOrder {
	defer qt.Recovering("QHeaderView::sortIndicatorOrder")

	if ptr.Pointer() != nil {
		return core.Qt__SortOrder(C.QHeaderView_SortIndicatorOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) SortIndicatorSection() int {
	defer qt.Recovering("QHeaderView::sortIndicatorSection")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SortIndicatorSection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) StretchSectionCount() int {
	defer qt.Recovering("QHeaderView::stretchSectionCount")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_StretchSectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) SwapSections(first int, second int) {
	defer qt.Recovering("QHeaderView::swapSections")

	if ptr.Pointer() != nil {
		C.QHeaderView_SwapSections(ptr.Pointer(), C.int(first), C.int(second))
	}
}

func (ptr *QHeaderView) VerticalOffset() int {
	defer qt.Recovering("QHeaderView::verticalOffset")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VerticalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ViewportEvent(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHeaderView::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QHeaderView_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHeaderView) VisualIndex(logicalIndex int) int {
	defer qt.Recovering("QHeaderView::visualIndex")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndex(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) VisualIndexAt(position int) int {
	defer qt.Recovering("QHeaderView::visualIndexAt")

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndexAt(ptr.Pointer(), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) DestroyQHeaderView() {
	defer qt.Recovering("QHeaderView::~QHeaderView")

	if ptr.Pointer() != nil {
		C.QHeaderView_DestroyQHeaderView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHeaderView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QHeaderView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QHeaderView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQHeaderViewDragLeaveEvent
func callbackQHeaderViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QHeaderView) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHeaderView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QHeaderView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QHeaderView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QHeaderView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QHeaderView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQHeaderViewCloseEditor
func callbackQHeaderViewCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QHeaderView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QHeaderView) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHeaderView::closeEditor")

	if ptr.Pointer() != nil {
		C.QHeaderView_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHeaderView) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QHeaderView::closeEditor")

	if ptr.Pointer() != nil {
		C.QHeaderView_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QHeaderView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QHeaderView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QHeaderView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QHeaderView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQHeaderViewCommitData
func callbackQHeaderViewCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHeaderView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHeaderView) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QHeaderView::commitData")

	if ptr.Pointer() != nil {
		C.QHeaderView_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QHeaderView) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QHeaderView::commitData")

	if ptr.Pointer() != nil {
		C.QHeaderView_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QHeaderView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QHeaderView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QHeaderView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQHeaderViewDragEnterEvent
func callbackQHeaderViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QHeaderView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHeaderView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QHeaderView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QHeaderView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QHeaderView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQHeaderViewDragMoveEvent
func callbackQHeaderViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QHeaderView) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHeaderView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QHeaderView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QHeaderView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QHeaderView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QHeaderView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQHeaderViewDropEvent
func callbackQHeaderViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QHeaderView) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHeaderView::dropEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHeaderView) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QHeaderView::dropEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QHeaderView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QHeaderView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QHeaderView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QHeaderView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQHeaderViewEditorDestroyed
func callbackQHeaderViewEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QHeaderView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QHeaderView) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QHeaderView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHeaderView_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHeaderView) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QHeaderView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QHeaderView_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QHeaderView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHeaderView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QHeaderView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQHeaderViewFocusInEvent
func callbackQHeaderViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHeaderView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHeaderView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHeaderView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHeaderView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHeaderView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QHeaderView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QHeaderView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQHeaderViewFocusOutEvent
func callbackQHeaderViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QHeaderView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHeaderView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHeaderView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QHeaderView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QHeaderView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QHeaderView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QHeaderView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQHeaderViewInputMethodEvent
func callbackQHeaderViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QHeaderView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHeaderView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHeaderView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QHeaderView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QHeaderView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHeaderView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QHeaderView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQHeaderViewKeyPressEvent
func callbackQHeaderViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHeaderView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHeaderView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHeaderView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHeaderView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHeaderView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QHeaderView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QHeaderView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QHeaderView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQHeaderViewKeyboardSearch
func callbackQHeaderViewKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QHeaderView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQHeaderViewFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QHeaderView) KeyboardSearch(search string) {
	defer qt.Recovering("QHeaderView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHeaderView) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QHeaderView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QHeaderView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QHeaderView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QHeaderView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQHeaderViewResizeEvent
func callbackQHeaderViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QHeaderView) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHeaderView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHeaderView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QHeaderView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QHeaderView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHeaderView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QHeaderView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QHeaderView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQHeaderViewRowsAboutToBeRemoved
func callbackQHeaderViewRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QHeaderView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QHeaderView) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHeaderView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHeaderView_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHeaderView) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHeaderView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QHeaderView_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHeaderView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QHeaderView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QHeaderView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QHeaderView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQHeaderViewRowsInserted
func callbackQHeaderViewRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QHeaderView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QHeaderView) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHeaderView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHeaderView_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHeaderView) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QHeaderView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QHeaderView_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QHeaderView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QHeaderView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QHeaderView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QHeaderView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQHeaderViewSelectAll
func callbackQHeaderViewSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHeaderView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHeaderView) SelectAll() {
	defer qt.Recovering("QHeaderView::selectAll")

	if ptr.Pointer() != nil {
		C.QHeaderView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QHeaderView) SelectAllDefault() {
	defer qt.Recovering("QHeaderView::selectAll")

	if ptr.Pointer() != nil {
		C.QHeaderView_SelectAllDefault(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QHeaderView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QHeaderView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QHeaderView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQHeaderViewSetRootIndex
func callbackQHeaderViewSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QHeaderView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QHeaderView) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHeaderView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHeaderView) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QHeaderView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QHeaderView) ConnectSetSelectionModel(f func(selectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QHeaderView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QHeaderView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QHeaderView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQHeaderViewSetSelectionModel
func callbackQHeaderViewSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, selectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(selectionModel))
	} else {
		NewQHeaderViewFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(selectionModel))
	}
}

func (ptr *QHeaderView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHeaderView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHeaderView) SetSelectionModelDefault(selectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QHeaderView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QHeaderView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QHeaderView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QHeaderView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QHeaderView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQHeaderViewStartDrag
func callbackQHeaderViewStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QHeaderView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQHeaderViewFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QHeaderView) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHeaderView::startDrag")

	if ptr.Pointer() != nil {
		C.QHeaderView_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHeaderView) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QHeaderView::startDrag")

	if ptr.Pointer() != nil {
		C.QHeaderView_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QHeaderView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHeaderView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHeaderView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHeaderViewTimerEvent
func callbackQHeaderViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHeaderView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHeaderView::timerEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHeaderView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHeaderView::timerEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHeaderView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QHeaderView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QHeaderView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QHeaderView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQHeaderViewUpdateGeometries
func callbackQHeaderViewUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QHeaderView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QHeaderView) UpdateGeometries() {
	defer qt.Recovering("QHeaderView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHeaderView_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QHeaderView) UpdateGeometriesDefault() {
	defer qt.Recovering("QHeaderView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QHeaderView_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QHeaderView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QHeaderView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQHeaderViewContextMenuEvent
func callbackQHeaderViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QHeaderView) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHeaderView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHeaderView) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QHeaderView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QHeaderView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QHeaderView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QHeaderView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QHeaderView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQHeaderViewScrollContentsBy
func callbackQHeaderViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QHeaderView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQHeaderViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QHeaderView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QHeaderView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHeaderView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHeaderView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QHeaderView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QHeaderView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QHeaderView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QHeaderView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QHeaderView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QHeaderView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQHeaderViewSetupViewport
func callbackQHeaderViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQHeaderViewFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QHeaderView) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QHeaderView::setupViewport")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QHeaderView) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QHeaderView::setupViewport")

	if ptr.Pointer() != nil {
		C.QHeaderView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QHeaderView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QHeaderView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QHeaderView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQHeaderViewWheelEvent
func callbackQHeaderViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQHeaderViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QHeaderView) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHeaderView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHeaderView) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QHeaderView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QHeaderView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QHeaderView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QHeaderView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQHeaderViewChangeEvent
func callbackQHeaderViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQHeaderViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QHeaderView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::changeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHeaderView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::changeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QHeaderView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QHeaderView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QHeaderView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQHeaderViewActionEvent
func callbackQHeaderViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QHeaderView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHeaderView::actionEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHeaderView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QHeaderView::actionEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QHeaderView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHeaderView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QHeaderView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQHeaderViewEnterEvent
func callbackQHeaderViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHeaderView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::enterEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHeaderView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::enterEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHeaderView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QHeaderView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QHeaderView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQHeaderViewHideEvent
func callbackQHeaderViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QHeaderView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHeaderView::hideEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHeaderView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QHeaderView::hideEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QHeaderView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHeaderView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QHeaderView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQHeaderViewLeaveEvent
func callbackQHeaderViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHeaderView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHeaderView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHeaderView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QHeaderView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QHeaderView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQHeaderViewMoveEvent
func callbackQHeaderViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QHeaderView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHeaderView::moveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHeaderView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QHeaderView::moveEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QHeaderView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QHeaderView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QHeaderView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQHeaderViewShowEvent
func callbackQHeaderViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QHeaderView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHeaderView::showEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHeaderView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QHeaderView::showEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QHeaderView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QHeaderView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QHeaderView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQHeaderViewCloseEvent
func callbackQHeaderViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QHeaderView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHeaderView::closeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHeaderView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QHeaderView::closeEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QHeaderView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QHeaderView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QHeaderView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QHeaderView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQHeaderViewInitPainter
func callbackQHeaderViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQHeaderViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QHeaderView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHeaderView::initPainter")

	if ptr.Pointer() != nil {
		C.QHeaderView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHeaderView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QHeaderView::initPainter")

	if ptr.Pointer() != nil {
		C.QHeaderView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QHeaderView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QHeaderView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QHeaderView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQHeaderViewKeyReleaseEvent
func callbackQHeaderViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QHeaderView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHeaderView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHeaderView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QHeaderView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QHeaderView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QHeaderView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QHeaderView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQHeaderViewTabletEvent
func callbackQHeaderViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QHeaderView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHeaderView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHeaderView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QHeaderView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QHeaderView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHeaderView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHeaderView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHeaderViewChildEvent
func callbackQHeaderViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHeaderView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHeaderView::childEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHeaderView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHeaderView::childEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHeaderView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHeaderView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHeaderView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHeaderView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHeaderViewCustomEvent
func callbackQHeaderViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHeaderView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHeaderViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHeaderView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::customEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHeaderView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHeaderView::customEvent")

	if ptr.Pointer() != nil {
		C.QHeaderView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
