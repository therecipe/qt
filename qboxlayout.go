package qt

//#include "qboxlayout.h"
import "C"

type qboxlayout struct {
	qlayout
}

type QBoxLayout interface {
	QLayout
	AddLayout(layout QLayout, stretch int)
	AddSpacing(size int)
	AddStretch(stretch int)
	AddStrut(size int)
	InsertLayout(index int, layout QLayout, stretch int)
	InsertSpacing(index int, size int)
	InsertStretch(index int, stretch int)
	InsertWidget(index int, widget QWidget, stretch int, alignment AlignmentFlag)
	SetStretch(index int, stretch int)
	SetStretchFactor(widget QWidget, stretch int) bool
	Stretch(index int) int
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

func (p *qboxlayout) AddLayout(layout QLayout, stretch int) {
	if p.Pointer() != nil {
		var layoutPtr C.QtObjectPtr
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QBoxLayout_AddLayout_QLayout_Int(p.Pointer(), layoutPtr, C.int(stretch))
	}
}

func (p *qboxlayout) AddSpacing(size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddSpacing_Int(p.Pointer(), C.int(size))
	}
}

func (p *qboxlayout) AddStretch(stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddStretch_Int(p.Pointer(), C.int(stretch))
	}
}

func (p *qboxlayout) AddStrut(size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_AddStrut_Int(p.Pointer(), C.int(size))
	}
}

func (p *qboxlayout) InsertLayout(index int, layout QLayout, stretch int) {
	if p.Pointer() != nil {
		var layoutPtr C.QtObjectPtr
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QBoxLayout_InsertLayout_Int_QLayout_Int(p.Pointer(), C.int(index), layoutPtr, C.int(stretch))
	}
}

func (p *qboxlayout) InsertSpacing(index int, size int) {
	if p.Pointer() != nil {
		C.QBoxLayout_InsertSpacing_Int_Int(p.Pointer(), C.int(index), C.int(size))
	}
}

func (p *qboxlayout) InsertStretch(index int, stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_InsertStretch_Int_Int(p.Pointer(), C.int(index), C.int(stretch))
	}
}

func (p *qboxlayout) InsertWidget(index int, widget QWidget, stretch int, alignment AlignmentFlag) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QBoxLayout_InsertWidget_Int_QWidget_Int_AlignmentFlag(p.Pointer(), C.int(index), widgetPtr, C.int(stretch), C.int(alignment))
	}
}

func (p *qboxlayout) SetStretch(index int, stretch int) {
	if p.Pointer() != nil {
		C.QBoxLayout_SetStretch_Int_Int(p.Pointer(), C.int(index), C.int(stretch))
	}
}

func (p *qboxlayout) SetStretchFactor(widget QWidget, stretch int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return C.QBoxLayout_SetStretchFactor_QWidget_Int(p.Pointer(), widgetPtr, C.int(stretch)) != 0
	}
}

func (p *qboxlayout) Stretch(index int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QBoxLayout_Stretch_Int(p.Pointer(), C.int(index)))
}
