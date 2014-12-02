package qt

//#include "qformlayout.h"
import "C"

type qformlayout struct {
	qlayout
}

type QFormLayout interface {
	QLayout
	AddRow_QWidget_QWidget(label QWidget, field QWidget)
	AddRow_QWidget_QLayout(label QWidget, field QLayout)
	AddRow_String_QWidget(labelText string, field QWidget)
	AddRow_String_QLayout(labelText string, field QLayout)
	AddRow_QWidget(widget QWidget)
	AddRow_QLayout(layout QLayout)
	FormAlignment() AlignmentFlag
	HorizontalSpacing() int
	InsertRow_Int_QWidget_QWidget(row int, label QWidget, field QWidget)
	InsertRow_Int_QWidget_QLayout(row int, label QWidget, field QLayout)
	InsertRow_Int_String_QWidget(row int, labelText string, field QWidget)
	InsertRow_Int_String_QLayout(row int, labelText string, field QLayout)
	InsertRow_Int_QWidget(row int, widget QWidget)
	InsertRow_Int_QLayout(row int, layout QLayout)
	LabelAlignment() AlignmentFlag
	LabelForField_QWidget(field QWidget) QWidget
	RowCount() int
	SetFormAlignment_AlignmentFlag(alignment AlignmentFlag)
	SetHorizontalSpacing_Int(spacing int)
	SetLabelAlignment_AlignmentFlag(alignment AlignmentFlag)
	SetVerticalSpacing_Int(spacing int)
	VerticalSpacing() int
}

func (p *qformlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qformlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQFormLayout_QWidget(parent QWidget) QFormLayout {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qformlayout = new(qformlayout)
	qformlayout.SetPointer(C.QFormLayout_New_QWidget(parentPtr))
	qformlayout.SetObjectName_String("QFormLayout_" + randomIdentifier())
	return qformlayout
}

func (p *qformlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QFormLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qformlayout) AddRow_QWidget_QWidget(label QWidget, field QWidget) {
	if p.Pointer() == nil {
	} else {
		var labelPtr C.QtObjectPtr = nil
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_QWidget_QWidget(p.Pointer(), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) AddRow_QWidget_QLayout(label QWidget, field QLayout) {
	if p.Pointer() == nil {
	} else {
		var labelPtr C.QtObjectPtr = nil
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_QWidget_QLayout(p.Pointer(), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) AddRow_String_QWidget(labelText string, field QWidget) {
	if p.Pointer() == nil {
	} else {
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_String_QWidget(p.Pointer(), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) AddRow_String_QLayout(labelText string, field QLayout) {
	if p.Pointer() == nil {
	} else {
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_AddRow_String_QLayout(p.Pointer(), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) AddRow_QWidget(widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QFormLayout_AddRow_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qformlayout) AddRow_QLayout(layout QLayout) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QFormLayout_AddRow_QLayout(p.Pointer(), layoutPtr)
	}
}

func (p *qformlayout) FormAlignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return AlignmentFlag(C.QFormLayout_FormAlignment(p.Pointer()))
	}
}

func (p *qformlayout) HorizontalSpacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QFormLayout_HorizontalSpacing(p.Pointer()))
	}
}

func (p *qformlayout) InsertRow_Int_QWidget_QWidget(row int, label QWidget, field QWidget) {
	if p.Pointer() == nil {
	} else {
		var labelPtr C.QtObjectPtr = nil
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget_QWidget(p.Pointer(), C.int(row), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) InsertRow_Int_QWidget_QLayout(row int, label QWidget, field QLayout) {
	if p.Pointer() == nil {
	} else {
		var labelPtr C.QtObjectPtr = nil
		if label != nil {
			labelPtr = label.Pointer()
		}
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget_QLayout(p.Pointer(), C.int(row), labelPtr, fieldPtr)
	}
}

func (p *qformlayout) InsertRow_Int_String_QWidget(row int, labelText string, field QWidget) {
	if p.Pointer() == nil {
	} else {
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_String_QWidget(p.Pointer(), C.int(row), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) InsertRow_Int_String_QLayout(row int, labelText string, field QLayout) {
	if p.Pointer() == nil {
	} else {
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		C.QFormLayout_InsertRow_Int_String_QLayout(p.Pointer(), C.int(row), C.CString(labelText), fieldPtr)
	}
}

func (p *qformlayout) InsertRow_Int_QWidget(row int, widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QWidget(p.Pointer(), C.int(row), widgetPtr)
	}
}

func (p *qformlayout) InsertRow_Int_QLayout(row int, layout QLayout) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QFormLayout_InsertRow_Int_QLayout(p.Pointer(), C.int(row), layoutPtr)
	}
}

func (p *qformlayout) LabelAlignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return AlignmentFlag(C.QFormLayout_LabelAlignment(p.Pointer()))
	}
}

func (p *qformlayout) LabelForField_QWidget(field QWidget) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var fieldPtr C.QtObjectPtr = nil
		if field != nil {
			fieldPtr = field.Pointer()
		}
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QFormLayout_LabelForField_QWidget(p.Pointer(), fieldPtr))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qformlayout) RowCount() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QFormLayout_RowCount(p.Pointer()))
	}
}

func (p *qformlayout) SetFormAlignment_AlignmentFlag(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QFormLayout_SetFormAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qformlayout) SetHorizontalSpacing_Int(spacing int) {
	if p.Pointer() != nil {
		C.QFormLayout_SetHorizontalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qformlayout) SetLabelAlignment_AlignmentFlag(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QFormLayout_SetLabelAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qformlayout) SetVerticalSpacing_Int(spacing int) {
	if p.Pointer() != nil {
		C.QFormLayout_SetVerticalSpacing_Int(p.Pointer(), C.int(spacing))
	}
}

func (p *qformlayout) VerticalSpacing() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QFormLayout_VerticalSpacing(p.Pointer()))
	}
}
