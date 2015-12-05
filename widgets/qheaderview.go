package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QHeaderView_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::cascadingSectionResizes")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_CascadingSectionResizes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) DefaultAlignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::defaultAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QHeaderView_DefaultAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) DefaultSectionSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::defaultSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_DefaultSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) HighlightSections() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::highlightSections")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_HighlightSections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) IsSortIndicatorShown() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::isSortIndicatorShown")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSortIndicatorShown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) MaximumSectionSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::maximumSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MaximumSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) MinimumSectionSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::minimumSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_MinimumSectionSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ResetDefaultSectionSize() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::resetDefaultSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ResetDefaultSectionSize(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ResizeSection(logicalIndex int, size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::resizeSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSection(ptr.Pointer(), C.int(logicalIndex), C.int(size))
	}
}

func (ptr *QHeaderView) SetCascadingSectionResizes(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setCascadingSectionResizes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetCascadingSectionResizes(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QHeaderView) SetDefaultAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setDefaultAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QHeaderView) SetDefaultSectionSize(size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setDefaultSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetDefaultSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetHighlightSections(highlight bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setHighlightSections")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetHighlightSections(ptr.Pointer(), C.int(qt.GoBoolToInt(highlight)))
	}
}

func (ptr *QHeaderView) SetMaximumSectionSize(size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setMaximumSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetMaximumSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetMinimumSectionSize(size int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setMinimumSectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetMinimumSectionSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QHeaderView) SetOffset(offset int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setOffset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffset(ptr.Pointer(), C.int(offset))
	}
}

func (ptr *QHeaderView) SetSortIndicatorShown(show bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSortIndicatorShown")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QHeaderView) SetStretchLastSection(stretch bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setStretchLastSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetStretchLastSection(ptr.Pointer(), C.int(qt.GoBoolToInt(stretch)))
	}
}

func (ptr *QHeaderView) StretchLastSection() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::stretchLastSection")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_StretchLastSection(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ConnectGeometriesChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::geometriesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectGeometriesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometriesChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectGeometriesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::geometriesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectGeometriesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometriesChanged")
	}
}

//export callbackQHeaderViewGeometriesChanged
func callbackQHeaderViewGeometriesChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::geometriesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "geometriesChanged").(func())()
}

func (ptr *QHeaderView) HeaderDataChanged(orientation core.Qt__Orientation, logicalFirst int, logicalLast int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::headerDataChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_HeaderDataChanged(ptr.Pointer(), C.int(orientation), C.int(logicalFirst), C.int(logicalLast))
	}
}

func (ptr *QHeaderView) HiddenSectionCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::hiddenSectionCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_HiddenSectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) HideSection(logicalIndex int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::hideSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_HideSection(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::isSectionHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_IsSectionHidden(ptr.Pointer(), C.int(logicalIndex)) != 0
	}
	return false
}

func (ptr *QHeaderView) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndex(visualIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::logicalIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndex(ptr.Pointer(), C.int(visualIndex)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt3(pos core.QPoint_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::logicalIndexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt3(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt(position int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::logicalIndexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt(ptr.Pointer(), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) LogicalIndexAt2(x int, y int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::logicalIndexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_LogicalIndexAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return 0
}

func (ptr *QHeaderView) MoveSection(from int, to int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::moveSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_MoveSection(ptr.Pointer(), C.int(from), C.int(to))
	}
}

func (ptr *QHeaderView) Offset() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::offset")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_Offset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) Orientation() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::orientation")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QHeaderView_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_Reset(ptr.Pointer())
	}
}

func (ptr *QHeaderView) ResizeContentsPrecision() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::resizeContentsPrecision")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_ResizeContentsPrecision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) ResizeSections(mode QHeaderView__ResizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::resizeSections")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ResizeSections(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QHeaderView) RestoreState(state core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::restoreState")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QHeaderView) SaveState() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::saveState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHeaderView_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHeaderView) ConnectSectionClicked(f func(logicalIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionClicked")
	}
}

//export callbackQHeaderViewSectionClicked
func callbackQHeaderViewSectionClicked(ptrName *C.char, logicalIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionCountChanged(f func(oldCount int, newCount int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionCountChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionCountChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionCountChanged")
	}
}

//export callbackQHeaderViewSectionCountChanged
func callbackQHeaderViewSectionCountChanged(ptrName *C.char, oldCount C.int, newCount C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionCountChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionCountChanged").(func(int, int))(int(oldCount), int(newCount))
}

func (ptr *QHeaderView) ConnectSectionDoubleClicked(f func(logicalIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionDoubleClicked")
	}
}

//export callbackQHeaderViewSectionDoubleClicked
func callbackQHeaderViewSectionDoubleClicked(ptrName *C.char, logicalIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionDoubleClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionEntered(f func(logicalIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionEntered", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionEntered")
	}
}

//export callbackQHeaderViewSectionEntered
func callbackQHeaderViewSectionEntered(ptrName *C.char, logicalIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionEntered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionEntered").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionHandleDoubleClicked(f func(logicalIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionHandleDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionHandleDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionHandleDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionHandleDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionHandleDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionHandleDoubleClicked")
	}
}

//export callbackQHeaderViewSectionHandleDoubleClicked
func callbackQHeaderViewSectionHandleDoubleClicked(ptrName *C.char, logicalIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionHandleDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionHandleDoubleClicked").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) ConnectSectionMoved(f func(logicalIndex int, oldVisualIndex int, newVisualIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionMoved", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionMoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionMoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionMoved")
	}
}

//export callbackQHeaderViewSectionMoved
func callbackQHeaderViewSectionMoved(ptrName *C.char, logicalIndex C.int, oldVisualIndex C.int, newVisualIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionMoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionMoved").(func(int, int, int))(int(logicalIndex), int(oldVisualIndex), int(newVisualIndex))
}

func (ptr *QHeaderView) SectionPosition(logicalIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionPosition(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionPressed(f func(logicalIndex int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionPressed", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionPressed")
	}
}

//export callbackQHeaderViewSectionPressed
func callbackQHeaderViewSectionPressed(ptrName *C.char, logicalIndex C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionPressed").(func(int))(int(logicalIndex))
}

func (ptr *QHeaderView) SectionResizeMode(logicalIndex int) QHeaderView__ResizeMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionResizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QHeaderView__ResizeMode(C.QHeaderView_SectionResizeMode(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) ConnectSectionResized(f func(logicalIndex int, oldSize int, newSize int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionResized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSectionResized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sectionResized", f)
	}
}

func (ptr *QHeaderView) DisconnectSectionResized() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionResized")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSectionResized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sectionResized")
	}
}

//export callbackQHeaderViewSectionResized
func callbackQHeaderViewSectionResized(ptrName *C.char, logicalIndex C.int, oldSize C.int, newSize C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionResized")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sectionResized").(func(int, int, int))(int(logicalIndex), int(oldSize), int(newSize))
}

func (ptr *QHeaderView) SectionSize(logicalIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSize(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionSizeHint(logicalIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionSizeHint")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionSizeHint(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionViewportPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SectionViewportPosition(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) SectionsClickable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionsClickable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsClickable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsHidden() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionsHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMovable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionsMovable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMovable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SectionsMoved() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sectionsMoved")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHeaderView_SectionsMoved(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHeaderView) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QHeaderView) SetOffsetToLastSection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setOffsetToLastSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToLastSection(ptr.Pointer())
	}
}

func (ptr *QHeaderView) SetOffsetToSectionPosition(visualSectionNumber int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setOffsetToSectionPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetOffsetToSectionPosition(ptr.Pointer(), C.int(visualSectionNumber))
	}
}

func (ptr *QHeaderView) SetResizeContentsPrecision(precision int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setResizeContentsPrecision")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetResizeContentsPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSectionHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionHidden(ptr.Pointer(), C.int(logicalIndex), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode(mode QHeaderView__ResizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSectionResizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionResizeMode2(logicalIndex int, mode QHeaderView__ResizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSectionResizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionResizeMode2(ptr.Pointer(), C.int(logicalIndex), C.int(mode))
	}
}

func (ptr *QHeaderView) SetSectionsClickable(clickable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSectionsClickable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsClickable(ptr.Pointer(), C.int(qt.GoBoolToInt(clickable)))
	}
}

func (ptr *QHeaderView) SetSectionsMovable(movable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSectionsMovable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSectionsMovable(ptr.Pointer(), C.int(qt.GoBoolToInt(movable)))
	}
}

func (ptr *QHeaderView) SetSortIndicator(logicalIndex int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setSortIndicator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetSortIndicator(ptr.Pointer(), C.int(logicalIndex), C.int(order))
	}
}

func (ptr *QHeaderView) SetVisible(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QHeaderView) ShowSection(logicalIndex int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::showSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ShowSection(ptr.Pointer(), C.int(logicalIndex))
	}
}

func (ptr *QHeaderView) ConnectSortIndicatorChanged(f func(logicalIndex int, order core.Qt__SortOrder)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sortIndicatorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_ConnectSortIndicatorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sortIndicatorChanged", f)
	}
}

func (ptr *QHeaderView) DisconnectSortIndicatorChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sortIndicatorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DisconnectSortIndicatorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sortIndicatorChanged")
	}
}

//export callbackQHeaderViewSortIndicatorChanged
func callbackQHeaderViewSortIndicatorChanged(ptrName *C.char, logicalIndex C.int, order C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sortIndicatorChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sortIndicatorChanged").(func(int, core.Qt__SortOrder))(int(logicalIndex), core.Qt__SortOrder(order))
}

func (ptr *QHeaderView) SortIndicatorOrder() core.Qt__SortOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sortIndicatorOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__SortOrder(C.QHeaderView_SortIndicatorOrder(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) SortIndicatorSection() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::sortIndicatorSection")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_SortIndicatorSection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) StretchSectionCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::stretchSectionCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_StretchSectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHeaderView) SwapSections(first int, second int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::swapSections")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_SwapSections(ptr.Pointer(), C.int(first), C.int(second))
	}
}

func (ptr *QHeaderView) VisualIndex(logicalIndex int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::visualIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndex(ptr.Pointer(), C.int(logicalIndex)))
	}
	return 0
}

func (ptr *QHeaderView) VisualIndexAt(position int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::visualIndexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QHeaderView_VisualIndexAt(ptr.Pointer(), C.int(position)))
	}
	return 0
}

func (ptr *QHeaderView) DestroyQHeaderView() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHeaderView::~QHeaderView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHeaderView_DestroyQHeaderView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
