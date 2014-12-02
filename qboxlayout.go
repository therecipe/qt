package qt

//#include "qboxlayout.h"
import "C"

type qboxlayout struct {
	qlayout
}

type QBoxLayout interface {
	QLayout
	AddLayout_QLayout_Int(layout QLayout, stretch int)
	AddSpacing_Int(size int)
	AddStretch_Int(stretch int)
	AddStrut_Int(size int)
	InsertLayout_Int_QLayout_Int(index int, layout QLayout, stretch int)
	InsertSpacing_Int_Int(index int, size int)
	InsertStretch_Int_Int(index int, stretch int)
	InsertWidget_Int_QWidget_Int_AlignmentFlag(index int, widget QWidget, stretch int, alignment AlignmentFlag)
	SetStretch_Int_Int(index int, stretch int)
	SetStretchFactor_QWidget_Int(widget QWidget, stretch int) bool
	Stretch_Int(index int) int
}

func (p *qboxlayout) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qboxlayout) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qboxlayout) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QBoxLayout_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qboxlayout) AddLayout_QLayout_Int(layout QLayout, stretch int) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QBoxLayout_AddLayout_QLayout_Int(p.Pointer(), layoutPtr, C.int(stretch))
	}
}

func (p *qboxlayout) AddSpacing_Int(size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddSpacing_Int(p.Pointer(), C.int(size))
	}
}

func (p *qboxlayout) AddStretch_Int(stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddStretch_Int(p.Pointer(), C.int(stretch))
	}
}

func (p *qboxlayout) AddStrut_Int(size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddStrut_Int(p.Pointer(), C.int(size))
	}
}

func (p *qboxlayout) InsertLayout_Int_QLayout_Int(index int, layout QLayout, stretch int) {
	if p.Pointer() == nil {
	} else {
		var layoutPtr C.QtObjectPtr = nil
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QBoxLayout_InsertLayout_Int_QLayout_Int(p.Pointer(), C.int(index), layoutPtr, C.int(stretch))
	}
}

func (p *qboxlayout) InsertSpacing_Int_Int(index int, size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_InsertSpacing_Int_Int(p.Pointer(), C.int(index), C.int(size))
	}
}

func (p *qboxlayout) InsertStretch_Int_Int(index int, stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_InsertStretch_Int_Int(p.Pointer(), C.int(index), C.int(stretch))
	}
}

func (p *qboxlayout) InsertWidget_Int_QWidget_Int_AlignmentFlag(index int, widget QWidget, stretch int, alignment AlignmentFlag) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QBoxLayout_InsertWidget_Int_QWidget_Int_AlignmentFlag(p.Pointer(), C.int(index), widgetPtr, C.int(stretch), C.int(alignment))
	}
}

func (p *qboxlayout) SetStretch_Int_Int(index int, stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_SetStretch_Int_Int(p.Pointer(), C.int(index), C.int(stretch))
	}
}

func (p *qboxlayout) SetStretchFactor_QWidget_Int(widget QWidget, stretch int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return C.QBoxLayout_SetStretchFactor_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch)) != 0
	}
}

func (p *qboxlayout) Stretch_Int(index int) int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QBoxLayout_Stretch_Int(p.Pointer(), C.int(index)))
	}
}
