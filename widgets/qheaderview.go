package widgets

//#include "qheaderview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHeaderView struct {
	QAbstractItemView
}

type QHeaderViewITF interface {
	QAbstractItemViewITF
	QHeaderViewPTR() *QHeaderView
}

func PointerFromQHeaderView(ptr QHeaderViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHeaderViewPTR().Pointer()
	}
	return nil
}

func QHeaderViewFromPointer(ptr unsafe.Pointer) *QHeaderView {
	var n = new(QHeaderView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHeaderView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHeaderView) QHeaderViewPTR() *QHeaderView {
	return ptr
}

//QHeaderView::ResizeMode
type QHeaderView__ResizeMode int

var (
	QHeaderView__Interactive      = QHeaderView__ResizeMode(0)
	QHeaderView__Stretch          = QHeaderView__ResizeMode(1)
	QHeaderView__Fixed            = QHeaderView__ResizeMode(2)
	QHeaderView__ResizeToContents = QHeaderView__ResizeMode(3)
	QHeaderView__Custom           = QHeaderView__ResizeMode(QHeaderView__Fixed)
)

func (ptr *QHeaderView) CascadingSectionResizes() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_CascadingSectionResizes(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) DefaultAlignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QHeaderView_DefaultAlignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) DefaultSectionSize() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_DefaultSectionSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) HighlightSections() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_HighlightSections(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) IsSortIndicatorShown() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSortIndicatorShown(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) MaximumSectionSize() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MaximumSectionSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) MinimumSectionSize() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MinimumSectionSize(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) ResetDefaultSectionSize() {
	if ptr.Pointer() != nil {
		C.QHeaderView_ResetDefaultSectionSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHeaderView) ResizeSection(logicalIndex int, size int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSection(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex), C.int(size))
	}
}

func (ptr *QHeaderView) SetCascadingSectionResizes(enable bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetCascadingSectionResizes(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QHeaderView) SetDefaultAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QHeaderView) SetDefaultSectionSize(size int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultSectionSize(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QHeaderView) SetHighlightSections(highlight bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetHighlightSections(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(highlight)))
	}
}

func (ptr *QHeaderView) SetMaximumSectionSize(size int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetMaximumSectionSize(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QHeaderView) SetMinimumSectionSize(size int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetMinimumSectionSize(C.QtObjectPtr(ptr.Pointer()), C.int(size))
	}
}

func (ptr *QHeaderView) SetOffset(offset int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffset(C.QtObjectPtr(ptr.Pointer()), C.int(offset))
	}
}

func (ptr *QHeaderView) SetSortIndicatorShown(show bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicatorShown(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QHeaderView) SetStretchLastSection(stretch bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetStretchLastSection(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(stretch)))
	}
}

func (ptr *QHeaderView) StretchLastSection() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_StretchLastSection(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) ConnectGeometriesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectGeometriesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "geometriesChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectGeometriesChanged() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectGeometriesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "geometriesChanged")
	}
}

//export callbackQHeaderViewGeometriesChanged
func callbackQHeaderViewGeometriesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "geometriesChanged").(func())()
}

func (ptr *QHeaderView) HeaderDataChanged(orientation core.Qt__Orientation, logicalFirst int, logicalLast int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_HeaderDataChanged(C.QtObjectPtr(ptr.Pointer()), C.int(orientation), C.int(logicalFirst), C.int(logicalLast))
	}
}

func (ptr *QHeaderView) HiddenSectionCount() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_HiddenSectionCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) HideSection(logicalIndex int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_HideSection(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSectionHidden(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)) != 0
	}
	return false
}

func (ptr *QHeaderView) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndex(visualIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndex(C.QtObjectPtr(ptr.Pointer()), C.int(visualIndex)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt3(pos core.QPointITF) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos))))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt(position int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt(C.QtObjectPtr(ptr.Pointer()), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt2(x int, y int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y)))
	}
	return 0
}

func (ptr *QHeaderView) MoveSection(from int, to int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_MoveSection(C.QtObjectPtr(ptr.Pointer()), C.int(from), C.int(to))
	}
}

func (ptr *QHeaderView) Offset() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Offset(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHeaderView_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) Reset() {
	if ptr.Pointer() != nil {
		C.QHeaderView_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHeaderView) ResizeContentsPrecision() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_ResizeContentsPrecision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) ResizeSections(mode QHeaderView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSections(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QHeaderView) RestoreState(state core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_RestoreState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(state))) != 0
	}
	return false
}

func (ptr *QHeaderView) ConnectSectionClicked(f func(logicalIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionClicked() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionClicked")
	}
}

//export callbackQHeaderViewSectionClicked
func callbackQHeaderViewSectionClicked(ptrName *C.char, logicalIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionCountChanged(f func(oldCount int, newCount int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionCountChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionCountChanged() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionCountChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionCountChanged")
	}
}

//export callbackQHeaderViewSectionCountChanged
func callbackQHeaderViewSectionCountChanged(ptrName *C.char, oldCount C.int, newCount C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionCountChanged").(func(int, int))(int(oldCount), int(newCount))
}

func (ptr *QHeaderView) ConnectSectionDoubleClicked(f func(logicalIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionDoubleClicked")
	}
}

//export callbackQHeaderViewSectionDoubleClicked
func callbackQHeaderViewSectionDoubleClicked(ptrName *C.char, logicalIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionDoubleClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionEntered(f func(logicalIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionEntered", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionEntered() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionEntered")
	}
}

//export callbackQHeaderViewSectionEntered
func callbackQHeaderViewSectionEntered(ptrName *C.char, logicalIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionEntered").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionHandleDoubleClicked(f func(logicalIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionHandleDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionHandleDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionHandleDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked")
	}
}

//export callbackQHeaderViewSectionHandleDoubleClicked
func callbackQHeaderViewSectionHandleDoubleClicked(ptrName *C.char, logicalIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionHandleDoubleClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionMoved(f func(logicalIndex int, oldVisualIndex int, newVisualIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionMoved", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionMoved() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionMoved(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionMoved")
	}
}

//export callbackQHeaderViewSectionMoved
func callbackQHeaderViewSectionMoved(ptrName *C.char, logicalIndex C.int, oldVisualIndex C.int, newVisualIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionMoved").(func(int, int, int))(int(logicalIndex), int(oldVisualIndex), int(newVisualIndex))
}

func (ptr *QHeaderView) SectionPosition(logicalIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionPosition(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionPressed(f func(logicalIndex int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionPressed", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionPressed() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionPressed")
	}
}

//export callbackQHeaderViewSectionPressed
func callbackQHeaderViewSectionPressed(ptrName *C.char, logicalIndex C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionPressed").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) SectionResizeMode(logicalIndex int) QHeaderView__ResizeMode {
	if ptr.Pointer() != nil {
		return QHeaderView__ResizeMode(C.QHeaderView_SectionResizeMode(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionResized(f func(logicalIndex int, oldSize int, newSize int)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionResized(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sectionResized", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionResized() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionResized(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sectionResized")
	}
}

//export callbackQHeaderViewSectionResized
func callbackQHeaderViewSectionResized(ptrName *C.char, logicalIndex C.int, oldSize C.int, newSize C.int) {
	qt.GetSignal(C.GoString(ptrName), "sectionResized").(func(int, int, int))(int(logicalIndex), int(oldSize), int(newSize))
}

func (ptr *QHeaderView) SectionSize(logicalIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSize(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionSizeHint(logicalIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSizeHint(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionViewportPosition(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionsClickable() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsClickable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMovable() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMovable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMoved() bool {
	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMoved(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHeaderView) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QHeaderView) SetOffsetToLastSection() {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToLastSection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QHeaderView) SetOffsetToSectionPosition(visualSectionNumber int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToSectionPosition(C.QtObjectPtr(ptr.Pointer()), C.int(visualSectionNumber))
	}
}

func (ptr *QHeaderView) SetResizeContentsPrecision(precision int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetResizeContentsPrecision(C.QtObjectPtr(ptr.Pointer()), C.int(precision))
	}
}

func (ptr *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionHidden(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode(mode QHeaderView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode2(logicalIndex int, mode QHeaderView__ResizeMode) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode2(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionsClickable(clickable bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsClickable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(clickable)))
	}
}

func (ptr *QHeaderView) SetSectionsMovable(movable bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsMovable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QHeaderView) SetSortIndicator(logicalIndex int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicator(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex), C.int(order))
	}
}

func (ptr *QHeaderView) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QHeaderView) ShowSection(logicalIndex int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ShowSection(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSortIndicatorChanged(f func(logicalIndex int, order core.Qt__SortOrder)) {
	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSortIndicatorChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "sortIndicatorChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSortIndicatorChanged() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSortIndicatorChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "sortIndicatorChanged")
	}
}

//export callbackQHeaderViewSortIndicatorChanged
func callbackQHeaderViewSortIndicatorChanged(ptrName *C.char, logicalIndex C.int, order C.int) {
	qt.GetSignal(C.GoString(ptrName), "sortIndicatorChanged").(func(int, core.Qt__SortOrder))(int(logicalIndex), core.Qt__SortOrder(order))
}

func (ptr *QHeaderView) SortIndicatorOrder() core.Qt__SortOrder {
	if ptr.Pointer() != nil {
		return core.Qt__SortOrder(C.QHeaderView_SortIndicatorOrder(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) SortIndicatorSection() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SortIndicatorSection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) StretchSectionCount() int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_StretchSectionCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHeaderView) SwapSections(first int, second int) {
	if ptr.Pointer() != nil {
		C.QHeaderView_SwapSections(C.QtObjectPtr(ptr.Pointer()), C.int(first), C.int(second))
	}
}

func (ptr *QHeaderView) VisualIndex(logicalIndex int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndex(C.QtObjectPtr(ptr.Pointer()), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) VisualIndexAt(position int) int {
	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndexAt(C.QtObjectPtr(ptr.Pointer()), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) DestroyQHeaderView() {
	if ptr.Pointer() != nil {
		C.QHeaderView_DestroyQHeaderView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
