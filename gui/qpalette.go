package gui

//#include "qpalette.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPalette struct {
	ptr unsafe.Pointer
}

type QPalette_ITF interface {
	QPalette_PTR() *QPalette
}

func (p *QPalette) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPalette) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPalette(ptr QPalette_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPalette_PTR().Pointer()
	}
	return nil
}

func NewQPaletteFromPointer(ptr unsafe.Pointer) *QPalette {
	var n = new(QPalette)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPalette) QPalette_PTR() *QPalette {
	return ptr
}

//QPalette::ColorGroup
type QPalette__ColorGroup int64

const (
	QPalette__Active       = QPalette__ColorGroup(0)
	QPalette__Disabled     = QPalette__ColorGroup(1)
	QPalette__Inactive     = QPalette__ColorGroup(2)
	QPalette__NColorGroups = QPalette__ColorGroup(3)
	QPalette__Current      = QPalette__ColorGroup(4)
	QPalette__All          = QPalette__ColorGroup(5)
	QPalette__Normal       = QPalette__ColorGroup(QPalette__Active)
)

//QPalette::ColorRole
type QPalette__ColorRole int64

var (
	QPalette__WindowText      = QPalette__ColorRole(0)
	QPalette__Button          = QPalette__ColorRole(1)
	QPalette__Light           = QPalette__ColorRole(2)
	QPalette__Midlight        = QPalette__ColorRole(3)
	QPalette__Dark            = QPalette__ColorRole(4)
	QPalette__Mid             = QPalette__ColorRole(5)
	QPalette__Text            = QPalette__ColorRole(6)
	QPalette__BrightText      = QPalette__ColorRole(7)
	QPalette__ButtonText      = QPalette__ColorRole(8)
	QPalette__Base            = QPalette__ColorRole(9)
	QPalette__Window          = QPalette__ColorRole(10)
	QPalette__Shadow          = QPalette__ColorRole(11)
	QPalette__Highlight       = QPalette__ColorRole(12)
	QPalette__HighlightedText = QPalette__ColorRole(13)
	QPalette__Link            = QPalette__ColorRole(14)
	QPalette__LinkVisited     = QPalette__ColorRole(15)
	QPalette__AlternateBase   = QPalette__ColorRole(16)
	QPalette__NoRole          = QPalette__ColorRole(17)
	QPalette__ToolTipBase     = QPalette__ColorRole(18)
	QPalette__ToolTipText     = QPalette__ColorRole(19)
	QPalette__NColorRoles     = QPalette__ColorRole(C.QPalette_NColorRoles_Type())
	QPalette__Foreground      = QPalette__ColorRole(QPalette__WindowText)
	QPalette__Background      = QPalette__ColorRole(QPalette__Window)
)

func (ptr *QPalette) Brush(group QPalette__ColorGroup, role QPalette__ColorRole) *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Brush(ptr.Pointer(), C.int(group), C.int(role)))
	}
	return nil
}

func (ptr *QPalette) IsEqual(cg1 QPalette__ColorGroup, cg2 QPalette__ColorGroup) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsEqual(ptr.Pointer(), C.int(cg1), C.int(cg2)) != 0
	}
	return false
}

func (ptr *QPalette) SetBrush2(group QPalette__ColorGroup, role QPalette__ColorRole, brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetBrush2(ptr.Pointer(), C.int(group), C.int(role), PointerFromQBrush(brush))
	}
}

func NewQPalette() *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette())
}

func NewQPalette8(other QPalette_ITF) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette8(PointerFromQPalette(other)))
}

func NewQPalette3(button core.Qt__GlobalColor) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette3(C.int(button)))
}

func NewQPalette5(windowText QBrush_ITF, button QBrush_ITF, light QBrush_ITF, dark QBrush_ITF, mid QBrush_ITF, text QBrush_ITF, bright_text QBrush_ITF, base QBrush_ITF, window QBrush_ITF) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette5(PointerFromQBrush(windowText), PointerFromQBrush(button), PointerFromQBrush(light), PointerFromQBrush(dark), PointerFromQBrush(mid), PointerFromQBrush(text), PointerFromQBrush(bright_text), PointerFromQBrush(base), PointerFromQBrush(window)))
}

func NewQPalette2(button QColor_ITF) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette2(PointerFromQColor(button)))
}

func NewQPalette4(button QColor_ITF, window QColor_ITF) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette4(PointerFromQColor(button), PointerFromQColor(window)))
}

func NewQPalette7(p QPalette_ITF) *QPalette {
	return NewQPaletteFromPointer(C.QPalette_NewQPalette7(PointerFromQPalette(p)))
}

func (ptr *QPalette) AlternateBase() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_AlternateBase(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Base() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Base(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) BrightText() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_BrightText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Brush2(role QPalette__ColorRole) *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Brush2(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QPalette) Button() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Button(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) ButtonText() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_ButtonText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Color(group QPalette__ColorGroup, role QPalette__ColorRole) *QColor {
	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QPalette_Color(ptr.Pointer(), C.int(group), C.int(role)))
	}
	return nil
}

func (ptr *QPalette) Color2(role QPalette__ColorRole) *QColor {
	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QPalette_Color2(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QPalette) CurrentColorGroup() QPalette__ColorGroup {
	if ptr.Pointer() != nil {
		return QPalette__ColorGroup(C.QPalette_CurrentColorGroup(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPalette) Dark() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Dark(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Highlight() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Highlight(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) HighlightedText() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_HighlightedText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) IsBrushSet(cg QPalette__ColorGroup, cr QPalette__ColorRole) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsBrushSet(ptr.Pointer(), C.int(cg), C.int(cr)) != 0
	}
	return false
}

func (ptr *QPalette) IsCopyOf(p QPalette_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsCopyOf(ptr.Pointer(), PointerFromQPalette(p)) != 0
	}
	return false
}

func (ptr *QPalette) Light() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Light(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Link() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Link(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) LinkVisited() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_LinkVisited(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Mid() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Mid(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Midlight() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Midlight(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) SetBrush(role QPalette__ColorRole, brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetBrush(ptr.Pointer(), C.int(role), PointerFromQBrush(brush))
	}
}

func (ptr *QPalette) SetColor(group QPalette__ColorGroup, role QPalette__ColorRole, color QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColor(ptr.Pointer(), C.int(group), C.int(role), PointerFromQColor(color))
	}
}

func (ptr *QPalette) SetColor2(role QPalette__ColorRole, color QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColor2(ptr.Pointer(), C.int(role), PointerFromQColor(color))
	}
}

func (ptr *QPalette) SetColorGroup(cg QPalette__ColorGroup, windowText QBrush_ITF, button QBrush_ITF, light QBrush_ITF, dark QBrush_ITF, mid QBrush_ITF, text QBrush_ITF, bright_text QBrush_ITF, base QBrush_ITF, window QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColorGroup(ptr.Pointer(), C.int(cg), PointerFromQBrush(windowText), PointerFromQBrush(button), PointerFromQBrush(light), PointerFromQBrush(dark), PointerFromQBrush(mid), PointerFromQBrush(text), PointerFromQBrush(bright_text), PointerFromQBrush(base), PointerFromQBrush(window))
	}
}

func (ptr *QPalette) SetCurrentColorGroup(cg QPalette__ColorGroup) {
	if ptr.Pointer() != nil {
		C.QPalette_SetCurrentColorGroup(ptr.Pointer(), C.int(cg))
	}
}

func (ptr *QPalette) Shadow() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Shadow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Swap(other QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QPalette_Swap(ptr.Pointer(), PointerFromQPalette(other))
	}
}

func (ptr *QPalette) Text() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Text(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) ToolTipBase() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_ToolTipBase(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) ToolTipText() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_ToolTipText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) Window() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) WindowText() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPalette_WindowText(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPalette) DestroyQPalette() {
	if ptr.Pointer() != nil {
		C.QPalette_DestroyQPalette(ptr.Pointer())
	}
}
