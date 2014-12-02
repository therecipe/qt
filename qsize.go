package qt

//#include "qsize.h"
import "C"

type qsize struct {
	ptr C.QtObjectPtr
}

type QSize interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Height() int
	IsEmpty() bool
	IsNull() bool
	IsValid() bool
	Rheight() int
	Rwidth() int
	Scale_Int_Int_AspectRatioMode(width int, height int, mode AspectRatioMode)
	SetHeight_Int(height int)
	SetWidth_Int(width int)
	Transpose()
	Width() int
}

func (p *qsize) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qsize) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qsize) Height() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSize_Height(p.Pointer()))
	}
}

func (p *qsize) IsEmpty() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSize_IsEmpty(p.Pointer()) != 0
	}
}

func (p *qsize) IsNull() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSize_IsNull(p.Pointer()) != 0
	}
}

func (p *qsize) IsValid() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QSize_IsValid(p.Pointer()) != 0
	}
}

func (p *qsize) Rheight() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSize_Rheight(p.Pointer()))
	}
}

func (p *qsize) Rwidth() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSize_Rwidth(p.Pointer()))
	}
}

func (p *qsize) Scale_Int_Int_AspectRatioMode(width int, height int, mode AspectRatioMode) {
	if p.Pointer() != nil {
		C.QSize_Scale_Int_Int_AspectRatioMode(p.Pointer(), C.int(width), C.int(height), C.int(mode))
	}
}

func (p *qsize) SetHeight_Int(height int) {
	if p.Pointer() != nil {
		C.QSize_SetHeight_Int(p.Pointer(), C.int(height))
	}
}

func (p *qsize) SetWidth_Int(width int) {
	if p.Pointer() != nil {
		C.QSize_SetWidth_Int(p.Pointer(), C.int(width))
	}
}

func (p *qsize) Transpose() {
	if p.Pointer() != nil {
		C.QSize_Transpose(p.Pointer())
	}
}

func (p *qsize) Width() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QSize_Width(p.Pointer()))
	}
}
