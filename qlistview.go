package qt

//#include "qlistview.h"
import "C"

type qlistview struct {
	qabstractitemview
}

type QListView interface {
	QAbstractItemView
	BatchSize() int
	ClearPropertyFlags()
	IsRowHidden(row int) bool
	IsSelectionRectVisible() bool
	IsWrapping() bool
	ModelColumn() int
	SetBatchSize(batchSize int)
	SetModelColumn(column int)
	SetRowHidden(row int, hide bool)
	SetSelectionRectVisible(show bool)
	SetSpacing(space int)
	SetUniformItemSizes(enable bool)
	SetWordWrap(on bool)
	SetWrapping(enable bool)
	Spacing() int
	UniformItemSizes() bool
	WordWrap() bool
}

func (p *qlistview) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlistview) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQListView(parent QWidget) QListView {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlistview = new(qlistview)
	qlistview.SetPointer(C.QListView_New_QWidget(parentPtr))
	qlistview.SetObjectName("QListView_" + randomIdentifier())
	return qlistview
}

func (p *qlistview) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QListView_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlistview) BatchSize() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListView_BatchSize(p.Pointer()))
}

func (p *qlistview) ClearPropertyFlags() {
	if p.Pointer() != nil {
		C.QListView_ClearPropertyFlags(p.Pointer())
	}
}

func (p *qlistview) IsRowHidden(row int) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListView_IsRowHidden_Int(p.Pointer(), C.int(row)) != 0
}

func (p *qlistview) IsSelectionRectVisible() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListView_IsSelectionRectVisible(p.Pointer()) != 0
}

func (p *qlistview) IsWrapping() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListView_IsWrapping(p.Pointer()) != 0
}

func (p *qlistview) ModelColumn() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListView_ModelColumn(p.Pointer()))
}

func (p *qlistview) SetBatchSize(batchSize int) {
	if p.Pointer() != nil {
		C.QListView_SetBatchSize_Int(p.Pointer(), C.int(batchSize))
	}
}

func (p *qlistview) SetModelColumn(column int) {
	if p.Pointer() != nil {
		C.QListView_SetModelColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qlistview) SetRowHidden(row int, hide bool) {
	if p.Pointer() != nil {
		C.QListView_SetRowHidden_Int_Bool(p.Pointer(), C.int(row), goBoolToCInt(hide))
	}
}

func (p *qlistview) SetSelectionRectVisible(show bool) {
	if p.Pointer() != nil {
		C.QListView_SetSelectionRectVisible_Bool(p.Pointer(), goBoolToCInt(show))
	}
}

func (p *qlistview) SetSpacing(space int) {
	if p.Pointer() != nil {
		C.QListView_SetSpacing_Int(p.Pointer(), C.int(space))
	}
}

func (p *qlistview) SetUniformItemSizes(enable bool) {
	if p.Pointer() != nil {
		C.QListView_SetUniformItemSizes_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlistview) SetWordWrap(on bool) {
	if p.Pointer() != nil {
		C.QListView_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qlistview) SetWrapping(enable bool) {
	if p.Pointer() != nil {
		C.QListView_SetWrapping_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlistview) Spacing() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListView_Spacing(p.Pointer()))
}

func (p *qlistview) UniformItemSizes() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListView_UniformItemSizes(p.Pointer()) != 0
}

func (p *qlistview) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListView_WordWrap(p.Pointer()) != 0
}
