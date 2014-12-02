package qt

//#include "qgridlayout.h"
import "C"

type qgridlayout struct {
	qlayout
}

type QGridLayout interface {
	QLayout
	AddItem_QLayoutItem_Int_Int_Int_Int_AlignmentFlag(item QLayoutItem, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag)
	AddLayout_QLayout_Int_Int_AlignmentFlag(layout QLayout, row int, column int, alignment AlignmentFlag)
	AddLayout_QLayout_Int_Int_Int_Int_AlignmentFlag(layout QLayout, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag)
	AddWidget_QWidget_Int_Int_AlignmentFlag(widget QWidget, row int, column int, alignment AlignmentFlag)
	AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(widget QWidget, fromRow int, fromColumn int, rowSpan int, columnSpan int, alignment AlignmentFlag)
	ColumnCount() int
	ColumnMinimumWidth_Int(column int) int
	ColumnStretch_Int(column int) int
	HorizontalSpacing() int
	OriginCorner() Corner
	RowCount() int
	RowMinimumHeight_Int(row int) int
	RowStretch_Int(row int) int
	SetColumnMinimumWidth_Int_Int(column int, minSize int)
	SetColumnStretch_Int_Int(column int, stretch int)
	SetHorizontalSpacing_Int(spacing int)
	SetOriginCorner_Corner(corner Corner)
	SetRowMinimumHeight_Int_Int(row int, minSize int)
	SetRowStretch_Int_Int(row int, stretch int)
	SetVerticalSpacing_Int(spacing int)
	VerticalSpacing() int
}

func (p *qgridlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qgridlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQGridLayout_QWidget(parent QWidget) QGridLayout {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qgridlayout = new(qgridlayout)
	qgridlayout.SetPointer(C.QGridLayout_New_QWidget(parentPtr))
	qgridlayout.SetObjectName_String("QGridLayout_" + randomIdentifier())
	return qgridlayout
}

func NewQGridLayout() QGridLayout {
	var qgridlayout = new(qgridlayout)
	qgridlayout.SetPointer(C.QGridLayout_New())
	qgridlayout.SetObjectName_String("QGridLayout_" + randomIdentifier())
	return qgridlayout
}

func (p *qgridlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QGridLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qgridlayout) AddItem_QLayoutItem_Int_Int_Int_Int_AlignmentFlag(item QLayoutItem, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var itemPtr C.QtObjectPtr = nil
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QGridLayout_AddItem_QLayoutItem_Int_Int_Int_Int_AlignmentFlag(p.Pointer(), itemPtr, C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (p *qgridlayout) AddLayout_QLayout_Int_Int_AlignmentFlag(layout QLayout, row int, column int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QGridLayout_AddLayout_QLayout_Int_Int_AlignmentFlag(p.Pointer(), layoutPtr, C.int(row), C.int(column), C.int(alignment))
	}
}

func (p *qgridlayout) AddLayout_QLayout_Int_Int_Int_Int_AlignmentFlag(layout QLayout, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QGridLayout_AddLayout_QLayout_Int_Int_Int_Int_AlignmentFlag(p.Pointer(), layoutPtr, C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (p *qgridlayout) AddWidget_QWidget_Int_Int_AlignmentFlag(widget QWidget, row int, column int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QGridLayout_AddWidget_QWidget_Int_Int_AlignmentFlag(p.Pointer(), widgetPtr, C.int(row), C.int(column), C.int(alignment))
	}
}

func (p *qgridlayout) AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(widget QWidget, fromRow int, fromColumn int, rowSpan int, columnSpan int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QGridLayout_AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(p.Pointer(), widgetPtr, C.int(fromRow), C.int(fromColumn), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (p *qgridlayout) ColumnCount() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_ColumnCount(p.Pointer()))
	}
}

func (p *qgridlayout) ColumnMinimumWidth_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_ColumnMinimumWidth_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qgridlayout) ColumnStretch_Int(column int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_ColumnStretch_Int(p.Pointer(), C.int(column)))
	}
}

func (p *qgridlayout) HorizontalSpacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_HorizontalSpacing(p.Pointer()))
	}
}

func (p *qgridlayout) OriginCorner() Corner {
	if p.Pointer() == nil {
		return 0
	} else {
		return Corner(C.QGridLayout_OriginCorner(p.Pointer()))
	}
}

func (p *qgridlayout) RowCount() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_RowCount(p.Pointer()))
	}
}

func (p *qgridlayout) RowMinimumHeight_Int(row int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_RowMinimumHeight_Int(p.Pointer(), C.int(row)))
	}
}

func (p *qgridlayout) RowStretch_Int(row int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_RowStretch_Int(p.Pointer(), C.int(row)))
	}
}

func (p *qgridlayout) SetColumnMinimumWidth_Int_Int(column int, minSize int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetColumnMinimumWidth_Int_Int(p.Pointer(), C.int(column), C.int(minSize))
	}
}

func (p *qgridlayout) SetColumnStretch_Int_Int(column int, stretch int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetColumnStretch_Int_Int(p.Pointer(), C.int(column), C.int(stretch))
	}
}

func (p *qgridlayout) SetHorizontalSpacing_Int(spacing int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetHorizontalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qgridlayout) SetOriginCorner_Corner(corner Corner) {
	if p.Pointer() != nil {
		C.QGridLayout_SetOriginCorner_Corner(p.Pointer(), C.int(corner))
	}
}

func (p *qgridlayout) SetRowMinimumHeight_Int_Int(row int, minSize int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetRowMinimumHeight_Int_Int(p.Pointer(), C.int(row), C.int(minSize))
	}
}

func (p *qgridlayout) SetRowStretch_Int_Int(row int, stretch int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetRowStretch_Int_Int(p.Pointer(), C.int(row), C.int(stretch))
	}
}

func (p *qgridlayout) SetVerticalSpacing_Int(spacing int) {
	if p.Pointer() != nil {
		C.QGridLayout_SetVerticalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qgridlayout) VerticalSpacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QGridLayout_VerticalSpacing(p.Pointer()))
	}
}
