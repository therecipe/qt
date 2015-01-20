package qt

//#include "qformlayout.h"
import "C"

type qformlayout struct {
	qlayout
}

type QFormLayout interface {
	QLayout
	AddRow1(label QWidget, field QWidget)
	AddRow2(label QWidget, field QLayout)
	AddRow3(labelText string, field QWidget)
	AddRow4(labelText string, field QLayout)
	AddRow5(widget QWidget)
	AddRow6(layout QLayout)
	FormAlignment() AlignmentFlag
	HorizontalSpacing() int
	InsertRow1(row int, label QWidget, field QWidget)
	InsertRow2(row int, label QWidget, field QLayout)
	InsertRow3(row int, labelText string, field QWidget)
	InsertRow4(row int, labelText string, field QLayout)
	InsertRow5(row int, widget QWidget)
	InsertRow6(row int, layout QLayout)
	LabelAlignment() AlignmentFlag
	LabelForField(field QWidget) QWidget
	RowCount() int
	SetFormAlignment(alignment AlignmentFlag)
	SetHorizontalSpacing(spacing int)
	SetLabelAlignment(alignment AlignmentFlag)
	SetVerticalSpacing(spacing int)
	VerticalSpacing() int
}

func (p *qformlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qformlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQFormLayout(parent QWidget) QFormLayout {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qformlayout = new(qformlayout)
	qformlayout.SetPointer(C.QFormLayout_New_QWidget(parentPtr))
	qformlayout.SetObjectName("QFormLayout_" + randomIdentifier())
	return qformlayout
}

func (p *qformlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QFormLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qformlayout) AddRow1(label QWidget, field QWidget) {
	if p.Pointer() != nil {
		var labelPtr C.QtObjectPtr
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_QWidget_QWidget(p.Pointer(), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) AddRow2(label QWidget, field QLayout) {
	if p.Pointer() != nil {
		var labelPtr C.QtObjectPtr
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_QWidget_QLayout(p.Pointer(), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) AddRow3(labelText string, field QWidget) {
	if p.Pointer() != nil {
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_String_QWidget(p.Pointer(), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) AddRow4(labelText string, field QLayout) {
	if p.Pointer() != nil {
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_String_QLayout(p.Pointer(), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) AddRow5(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QFormLayout_AddRow_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qformlayout) AddRow6(layout QLayout) {
	if p.Pointer() != nil {
		var layoutPtr C.QtObjectPtr
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QFormLayout_AddRow_QLayout(p.Pointer(), layoutPtr)
	}
}

func (p *qformlayout) FormAlignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QFormLayout_FormAlignment(p.Pointer()))
}

func (p *qformlayout) HorizontalSpacing() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFormLayout_HorizontalSpacing(p.Pointer()))
}

func (p *qformlayout) InsertRow1(row int, label QWidget, field QWidget) {
	if p.Pointer() != nil {
		var labelPtr C.QtObjectPtr
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget_QWidget(p.Pointer(), C.int(row), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) InsertRow2(row int, label QWidget, field QLayout) {
	if p.Pointer() != nil {
		var labelPtr C.QtObjectPtr
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget_QLayout(p.Pointer(), C.int(row), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) InsertRow3(row int, labelText string, field QWidget) {
	if p.Pointer() != nil {
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_String_QWidget(p.Pointer(), C.int(row), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) InsertRow4(row int, labelText string, field QLayout) {
	if p.Pointer() != nil {
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_String_QLayout(p.Pointer(), C.int(row), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) InsertRow5(row int, widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget(p.Pointer(), C.int(row), widgetPtr)
	}
}

func (p *qformlayout) InsertRow6(row int, layout QLayout) {
	if p.Pointer() != nil {
		var layoutPtr C.QtObjectPtr
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QLayout(p.Pointer(), C.int(row), layoutPtr)
	}
}

func (p *qformlayout) LabelAlignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QFormLayout_LabelAlignment(p.Pointer()))
}

func (p *qformlayout) LabelForField(field QWidget) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var fieldPtr C.QtObjectPtr
		if field != nil {
			fieldPtr = field.Pointer()
		}
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QFormLayout_LabelForField_QWidget(p.Pointer(), fieldPtr))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qformlayout) RowCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFormLayout_RowCount(p.Pointer()))
}

func (p *qformlayout) SetFormAlignment(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QFormLayout_SetFormAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qformlayout) SetHorizontalSpacing(spacing int) {
	if p.Pointer() != nil {
		C.QFormLayout_SetHorizontalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qformlayout) SetLabelAlignment(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QFormLayout_SetLabelAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qformlayout) SetVerticalSpacing(spacing int) {
	if p.Pointer() != nil {
		C.QFormLayout_SetVerticalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qformlayout) VerticalSpacing() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QFormLayout_VerticalSpacing(p.Pointer()))
}
