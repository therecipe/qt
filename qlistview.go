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
	IsRowHidden_Int(row int) bool
	IsSelectionRectVisible() bool
	IsWrapping() bool
	ModelColumn() int
	SetBatchSize_Int(batchSize int)
	SetModelColumn_Int(column int)
	SetRowHidden_Int_Bool(row int, hide bool)
	SetSelectionRectVisible_Bool(show bool)
	SetSpacing_Int(space int)
	SetUniformItemSizes_Bool(enable bool)
	SetWordWrap_Bool(on bool)
	SetWrapping_Bool(enable bool)
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

func NewQListView_QWidget(parent QWidget) QListView {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlistview = new(qlistview)
	qlistview.SetPointer(C.QListView_New_QWidget(parentPtr))
	qlistview.SetObjectName_String("QListView_" + randomIdentifier())
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
	} else {
		return int(C.QListView_BatchSize(p.Pointer()))
	}
}

func (p *qlistview) ClearPropertyFlags() {
	if p.Pointer() != nil {
		C.QListView_ClearPropertyFlags(p.Pointer())
	}
}

func (p *qlistview) IsRowHidden_Int(row int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QListView_IsRowHidden_Int(p.Pointer(), C.int(row)) != 0
	}
}

func (p *qlistview) IsSelectionRectVisible() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QListView_IsSelectionRectVisible(p.Pointer()) != 0
	}
}

func (p *qlistview) IsWrapping() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QListView_IsWrapping(p.Pointer()) != 0
	}
}

func (p *qlistview) ModelColumn() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QListView_ModelColumn(p.Pointer()))
	}
}

func (p *qlistview) SetBatchSize_Int(batchSize int) {
	if p.Pointer() != nil {
		C.QListView_SetBatchSize_Int(p.Pointer(), C.int(batchSize))
	}
}

func (p *qlistview) SetModelColumn_Int(column int) {
	if p.Pointer() != nil {
		C.QListView_SetModelColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qlistview) SetRowHidden_Int_Bool(row int, hide bool) {
	if p.Pointer() != nil {
		C.QListView_SetRowHidden_Int_Bool(p.Pointer(), C.int(row), goBoolToCInt(hide))
	}
}

func (p *qlistview) SetSelectionRectVisible_Bool(show bool) {
	if p.Pointer() != nil {
		C.QListView_SetSelectionRectVisible_Bool(p.Pointer(), goBoolToCInt(show))
	}
}

func (p *qlistview) SetSpacing_Int(space int) {
	if p.Pointer() != nil {
		C.QListView_SetSpacing_Int(p.Pointer(), C.int(space))
	}
}

func (p *qlistview) SetUniformItemSizes_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QListView_SetUniformItemSizes_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlistview) SetWordWrap_Bool(on bool) {
	if p.Pointer() != nil {
		C.QListView_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qlistview) SetWrapping_Bool(enable bool) {
	if p.Pointer() != nil {
		C.QListView_SetWrapping_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlistview) Spacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QListView_Spacing(p.Pointer()))
	}
}

func (p *qlistview) UniformItemSizes() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QListView_UniformItemSizes(p.Pointer()) != 0
	}
}

func (p *qlistview) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QListView_WordWrap(p.Pointer()) != 0
	}
}
