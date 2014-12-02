package qt

//#include "qbrush.h"
import "C"

type qbrush struct {
	ptr C.QtObjectPtr
}

type QBrush interface {
	Pointer() (ptr C.QtObjectPtr)
	SetPointer(ptr C.QtObjectPtr)
	Destroy()
	IsOpaque() bool
	SetColor_GlobalColor(color GlobalColor)
	SetStyle_BrushStyle(style BrushStyle)
	Style() BrushStyle
}

func (p *qbrush) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qbrush) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQBrush() QBrush {
	var qbrush = new(qbrush)
	qbrush.SetPointer(C.QBrush_New())
	return qbrush
}

func NewQBrush_BrushStyle(style BrushStyle) QBrush {
	var qbrush = new(qbrush)
	qbrush.SetPointer(C.QBrush_New_BrushStyle(C.int(style)))
	return qbrush
}

func NewQBrush_GlobalColor_BrushStyle(color GlobalColor, style BrushStyle) QBrush {
	var qbrush = new(qbrush)
	qbrush.SetPointer(C.QBrush_New_GlobalColor_BrushStyle(C.int(color), C.int(style)))
	return qbrush
}

func (p *qbrush) Destroy() {
	if p.Pointer() != nil {
		C.QBrush_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qbrush) IsOpaque() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QBrush_IsOpaque(p.Pointer()) != 0
	}
}

func (p *qbrush) SetColor_GlobalColor(color GlobalColor) {
	if p.Pointer() != nil {
		C.QBrush_SetColor_GlobalColor(p.Pointer(), C.int(color))
	}
}

func (p *qbrush) SetStyle_BrushStyle(style BrushStyle) {
	if p.Pointer() != nil {
		C.QBrush_SetStyle_BrushStyle(p.Pointer(), C.int(style))
	}
}

func (p *qbrush) Style() BrushStyle {
	if p.Pointer() == nil {
		return 0
	} else {
		return BrushStyle(C.QBrush_Style(p.Pointer()))
	}
}
