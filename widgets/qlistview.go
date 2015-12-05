package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QListView struct {
	QAbstractItemView
}

type QListView_ITF interface {
	QAbstractItemView_ITF
	QListView_PTR() *QListView
}

func PointerFromQListView(ptr QListView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListView_PTR().Pointer()
	}
	return nil
}

func NewQListViewFromPointer(ptr unsafe.Pointer) *QListView {
	var n = new(QListView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QListView_") {
		n.SetObjectName("QListView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QListView) QListView_PTR() *QListView {
	return ptr
}

//QListView::Flow
type QListView__Flow int64

const (
	QListView__LeftToRight = QListView__Flow(0)
	QListView__TopToBottom = QListView__Flow(1)
)

//QListView::LayoutMode
type QListView__LayoutMode int64

const (
	QListView__SinglePass = QListView__LayoutMode(0)
	QListView__Batched    = QListView__LayoutMode(1)
)

//QListView::Movement
type QListView__Movement int64

const (
	QListView__Static = QListView__Movement(0)
	QListView__Free   = QListView__Movement(1)
	QListView__Snap   = QListView__Movement(2)
)

//QListView::ResizeMode
type QListView__ResizeMode int64

const (
	QListView__Fixed  = QListView__ResizeMode(0)
	QListView__Adjust = QListView__ResizeMode(1)
)

//QListView::ViewMode
type QListView__ViewMode int64

const (
	QListView__ListMode = QListView__ViewMode(0)
	QListView__IconMode = QListView__ViewMode(1)
)

func (ptr *QListView) BatchSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::batchSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QListView_BatchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) Flow() QListView__Flow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::flow")
		}
	}()

	if ptr.Pointer() != nil {
		return QListView__Flow(C.QListView_Flow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) IsSelectionRectVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::isSelectionRectVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListView_IsSelectionRectVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) IsWrapping() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::isWrapping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListView_IsWrapping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) LayoutMode() QListView__LayoutMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::layoutMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QListView__LayoutMode(C.QListView_LayoutMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) ModelColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::modelColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QListView_ModelColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) Movement() QListView__Movement {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::movement")
		}
	}()

	if ptr.Pointer() != nil {
		return QListView__Movement(C.QListView_Movement(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) ResizeMode() QListView__ResizeMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::resizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QListView__ResizeMode(C.QListView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) SetBatchSize(batchSize int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setBatchSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetBatchSize(ptr.Pointer(), C.int(batchSize))
	}
}

func (ptr *QListView) SetFlow(flow QListView__Flow) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setFlow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetFlow(ptr.Pointer(), C.int(flow))
	}
}

func (ptr *QListView) SetGridSize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setGridSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetGridSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QListView) SetLayoutMode(mode QListView__LayoutMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setLayoutMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetLayoutMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetModelColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setModelColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetModelColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QListView) SetMovement(movement QListView__Movement) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setMovement")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetMovement(ptr.Pointer(), C.int(movement))
	}
}

func (ptr *QListView) SetResizeMode(mode QListView__ResizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setResizeMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetResizeMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetSelectionRectVisible(show bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setSelectionRectVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetSelectionRectVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(show)))
	}
}

func (ptr *QListView) SetSpacing(space int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetSpacing(ptr.Pointer(), C.int(space))
	}
}

func (ptr *QListView) SetUniformItemSizes(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setUniformItemSizes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetUniformItemSizes(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QListView) SetViewMode(mode QListView__ViewMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setViewMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QListView) SetWordWrap(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setWordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QListView) SetWrapping(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setWrapping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QListView) Spacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::spacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QListView_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) UniformItemSizes() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::uniformItemSizes")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListView_UniformItemSizes(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QListView) ViewMode() QListView__ViewMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::viewMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QListView__ViewMode(C.QListView_ViewMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QListView) WordWrap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::wordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListView_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQListView(parent QWidget_ITF) *QListView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::QListView")
		}
	}()

	return NewQListViewFromPointer(C.QListView_NewQListView(PointerFromQWidget(parent)))
}

func (ptr *QListView) ClearPropertyFlags() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::clearPropertyFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_ClearPropertyFlags(ptr.Pointer())
	}
}

func (ptr *QListView) IndexAt(p core.QPoint_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::indexAt")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QListView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QListView) IsRowHidden(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::isRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QListView_IsRowHidden(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QListView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::scrollTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QListView) SetRowHidden(row int, hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::setRowHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_SetRowHidden(ptr.Pointer(), C.int(row), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QListView) DestroyQListView() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QListView::~QListView")
		}
	}()

	if ptr.Pointer() != nil {
		C.QListView_DestroyQListView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
