package qt

//#include "qrect.h"
import "C"

type qrect struct {
	ptr C.QtObjectPtr
}

type QRect interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Adjust_Int_Int_Int_Int(dx1 int, dy1 int, dx2 int, dy2 int)
	Bottom() int
	Contains_Int_Int_Bool(x int, y int, proper bool) bool
	Contains_Int_Int(x int, y int) bool
	Height() int
	IsEmpty() bool
	IsNull() bool
	IsValid() bool
	Left() int
	MoveBottom_Int(y int)
	MoveLeft_Int(x int)
	MoveRight_Int(x int)
	MoveTo_Int_Int(x int, y int)
	MoveTop_Int(y int)
	Right() int
	SetBottom_Int(y int)
	SetCoords_Int_Int_Int_Int(x1 int, y1 int, x2 int, y2 int)
	SetHeight_Int(height int)
	SetLeft_Int(x int)
	SetRect_Int_Int_Int_Int(x int, y int, width int, height int)
	SetRight_Int(x int)
	SetTop_Int(y int)
	SetWidth_Int(width int)
	SetX_Int(x int)
	SetY_Int(y int)
	Top() int
	Translate_Int_Int(dx int, dy int)
	Width() int
	X() int
	Y() int
}

func (p *qrect) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qrect) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func (p *qrect) Adjust_Int_Int_Int_Int(dx1 int, dy1 int, dx2 int, dy2 int) {
	if p.Pointer() != nil {
		C.QRect_Adjust_Int_Int_Int_Int(p.Pointer(), C.int(dx1), C.int(dy1), C.int(dx2), C.int(dy2))
	}
}

func (p *qrect) Bottom() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Bottom(p.Pointer()))
	}
}

func (p *qrect) Contains_Int_Int_Bool(x int, y int, proper bool) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QRect_Contains_Int_Int_Bool(p.Pointer(), C.int(x), C.int(y), goBoolToCInt(proper)) != 0
	}
}

func (p *qrect) Contains_Int_Int(x int, y int) bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QRect_Contains_Int_Int(p.Pointer(), C.int(x), C.int(y)) != 0
	}
}

func (p *qrect) Height() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Height(p.Pointer()))
	}
}

func (p *qrect) IsEmpty() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QRect_IsEmpty(p.Pointer()) != 0
	}
}

func (p *qrect) IsNull() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QRect_IsNull(p.Pointer()) != 0
	}
}

func (p *qrect) IsValid() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QRect_IsValid(p.Pointer()) != 0
	}
}

func (p *qrect) Left() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Left(p.Pointer()))
	}
}

func (p *qrect) MoveBottom_Int(y int) {
	if p.Pointer() != nil {
		C.QRect_MoveBottom_Int(p.Pointer(), C.int(y))
	}
}

func (p *qrect) MoveLeft_Int(x int) {
	if p.Pointer() != nil {
		C.QRect_MoveLeft_Int(p.Pointer(), C.int(x))
	}
}

func (p *qrect) MoveRight_Int(x int) {
	if p.Pointer() != nil {
		C.QRect_MoveRight_Int(p.Pointer(), C.int(x))
	}
}

func (p *qrect) MoveTo_Int_Int(x int, y int) {
	if p.Pointer() != nil {
		C.QRect_MoveTo_Int_Int(p.Pointer(), C.int(x), C.int(y))
	}
}

func (p *qrect) MoveTop_Int(y int) {
	if p.Pointer() != nil {
		C.QRect_MoveTop_Int(p.Pointer(), C.int(y))
	}
}

func (p *qrect) Right() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Right(p.Pointer()))
	}
}

func (p *qrect) SetBottom_Int(y int) {
	if p.Pointer() != nil {
		C.QRect_SetBottom_Int(p.Pointer(), C.int(y))
	}
}

func (p *qrect) SetCoords_Int_Int_Int_Int(x1 int, y1 int, x2 int, y2 int) {
	if p.Pointer() != nil {
		C.QRect_SetCoords_Int_Int_Int_Int(p.Pointer(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (p *qrect) SetHeight_Int(height int) {
	if p.Pointer() != nil {
		C.QRect_SetHeight_Int(p.Pointer(), C.int(height))
	}
}

func (p *qrect) SetLeft_Int(x int) {
	if p.Pointer() != nil {
		C.QRect_SetLeft_Int(p.Pointer(), C.int(x))
	}
}

func (p *qrect) SetRect_Int_Int_Int_Int(x int, y int, width int, height int) {
	if p.Pointer() != nil {
		C.QRect_SetRect_Int_Int_Int_Int(p.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (p *qrect) SetRight_Int(x int) {
	if p.Pointer() != nil {
		C.QRect_SetRight_Int(p.Pointer(), C.int(x))
	}
}

func (p *qrect) SetTop_Int(y int) {
	if p.Pointer() != nil {
		C.QRect_SetTop_Int(p.Pointer(), C.int(y))
	}
}

func (p *qrect) SetWidth_Int(width int) {
	if p.Pointer() != nil {
		C.QRect_SetWidth_Int(p.Pointer(), C.int(width))
	}
}

func (p *qrect) SetX_Int(x int) {
	if p.Pointer() != nil {
		C.QRect_SetX_Int(p.Pointer(), C.int(x))
	}
}

func (p *qrect) SetY_Int(y int) {
	if p.Pointer() != nil {
		C.QRect_SetY_Int(p.Pointer(), C.int(y))
	}
}

func (p *qrect) Top() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Top(p.Pointer()))
	}
}

func (p *qrect) Translate_Int_Int(dx int, dy int) {
	if p.Pointer() != nil {
		C.QRect_Translate_Int_Int(p.Pointer(), C.int(dx), C.int(dy))
	}
}

func (p *qrect) Width() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Width(p.Pointer()))
	}
}

func (p *qrect) X() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_X(p.Pointer()))
	}
}

func (p *qrect) Y() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QRect_Y(p.Pointer()))
	}
}
