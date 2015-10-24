package widgets

//#include "qgraphicsgridlayout.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsGridLayout struct {
	QGraphicsLayout
}

type QGraphicsGridLayoutITF interface {
	QGraphicsLayoutITF
	QGraphicsGridLayoutPTR() *QGraphicsGridLayout
}

func PointerFromQGraphicsGridLayout(ptr QGraphicsGridLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsGridLayoutPTR().Pointer()
	}
	return nil
}

func QGraphicsGridLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsGridLayout {
	var n = new(QGraphicsGridLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsGridLayout) QGraphicsGridLayoutPTR() *QGraphicsGridLayout {
	return ptr
}

func NewQGraphicsGridLayout(parent QGraphicsLayoutItemITF) *QGraphicsGridLayout {
	return QGraphicsGridLayoutFromPointer(unsafe.Pointer(C.QGraphicsGridLayout_NewQGraphicsGridLayout(C.QtObjectPtr(PointerFromQGraphicsLayoutItem(parent)))))
}

func (ptr *QGraphicsGridLayout) AddItem2(item QGraphicsLayoutItemITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) AddItem(item QGraphicsLayoutItemITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) Alignment(item QGraphicsLayoutItemITF) core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_Alignment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item))))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnAlignment(column int) core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_ColumnAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnCount() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) Invalidate() {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_Invalidate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsGridLayout) ItemAt2(index int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return QGraphicsLayoutItemFromPointer(unsafe.Pointer(C.QGraphicsGridLayout_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem {
	if ptr.Pointer() != nil {
		return QGraphicsLayoutItemFromPointer(unsafe.Pointer(C.QGraphicsGridLayout_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column))))
	}
	return nil
}

func (ptr *QGraphicsGridLayout) RemoveAt(index int) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveAt(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QGraphicsGridLayout) RemoveItem(item QGraphicsLayoutItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)))
	}
}

func (ptr *QGraphicsGridLayout) RowAlignment(row int) core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_RowAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowCount() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowStretchFactor(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) SetAlignment(item QGraphicsLayoutItemITF, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsLayoutItem(item)), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnAlignment(column int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) SetGeometry(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsGridLayout) SetRowAlignment(row int, alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowStretchFactor(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) DestroyQGraphicsGridLayout() {
	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_DestroyQGraphicsGridLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}
