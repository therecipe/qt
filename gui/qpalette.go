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

type QPaletteITF interface {
	QPalettePTR() *QPalette
}

func (p *QPalette) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPalette) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPalette(ptr QPaletteITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPalettePTR().Pointer()
	}
	return nil
}

func QPaletteFromPointer(ptr unsafe.Pointer) *QPalette {
	var n = new(QPalette)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPalette) QPalettePTR() *QPalette {
	return ptr
}

//QPalette::ColorGroup
type QPalette__ColorGroup int

var (
	QPalette__Active       = QPalette__ColorGroup(0)
	QPalette__Disabled     = QPalette__ColorGroup(1)
	QPalette__Inactive     = QPalette__ColorGroup(2)
	QPalette__NColorGroups = QPalette__ColorGroup(3)
	QPalette__Current      = QPalette__ColorGroup(4)
	QPalette__All          = QPalette__ColorGroup(5)
	QPalette__Normal       = QPalette__ColorGroup(QPalette__Active)
)

//QPalette::ColorRole
type QPalette__ColorRole int

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

func (ptr *QPalette) IsEqual(cg1 QPalette__ColorGroup, cg2 QPalette__ColorGroup) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsEqual(C.QtObjectPtr(ptr.Pointer()), C.int(cg1), C.int(cg2)) != 0
	}
	return false
}

func (ptr *QPalette) SetBrush2(group QPalette__ColorGroup, role QPalette__ColorRole, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetBrush2(C.QtObjectPtr(ptr.Pointer()), C.int(group), C.int(role), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func NewQPalette() *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette()))
}

func NewQPalette8(other QPaletteITF) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette8(C.QtObjectPtr(PointerFromQPalette(other)))))
}

func NewQPalette3(button core.Qt__GlobalColor) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette3(C.int(button))))
}

func NewQPalette5(windowText QBrushITF, button QBrushITF, light QBrushITF, dark QBrushITF, mid QBrushITF, text QBrushITF, bright_text QBrushITF, base QBrushITF, window QBrushITF) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette5(C.QtObjectPtr(PointerFromQBrush(windowText)), C.QtObjectPtr(PointerFromQBrush(button)), C.QtObjectPtr(PointerFromQBrush(light)), C.QtObjectPtr(PointerFromQBrush(dark)), C.QtObjectPtr(PointerFromQBrush(mid)), C.QtObjectPtr(PointerFromQBrush(text)), C.QtObjectPtr(PointerFromQBrush(bright_text)), C.QtObjectPtr(PointerFromQBrush(base)), C.QtObjectPtr(PointerFromQBrush(window)))))
}

func NewQPalette2(button QColorITF) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette2(C.QtObjectPtr(PointerFromQColor(button)))))
}

func NewQPalette4(button QColorITF, window QColorITF) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette4(C.QtObjectPtr(PointerFromQColor(button)), C.QtObjectPtr(PointerFromQColor(window)))))
}

func NewQPalette7(p QPaletteITF) *QPalette {
	return QPaletteFromPointer(unsafe.Pointer(C.QPalette_NewQPalette7(C.QtObjectPtr(PointerFromQPalette(p)))))
}

func (ptr *QPalette) CurrentColorGroup() QPalette__ColorGroup {
	if ptr.Pointer() != nil {
		return QPalette__ColorGroup(C.QPalette_CurrentColorGroup(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPalette) IsBrushSet(cg QPalette__ColorGroup, cr QPalette__ColorRole) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsBrushSet(C.QtObjectPtr(ptr.Pointer()), C.int(cg), C.int(cr)) != 0
	}
	return false
}

func (ptr *QPalette) IsCopyOf(p QPaletteITF) bool {
	if ptr.Pointer() != nil {
		return C.QPalette_IsCopyOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPalette(p))) != 0
	}
	return false
}

func (ptr *QPalette) SetBrush(role QPalette__ColorRole, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetBrush(C.QtObjectPtr(ptr.Pointer()), C.int(role), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPalette) SetColor(group QPalette__ColorGroup, role QPalette__ColorRole, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColor(C.QtObjectPtr(ptr.Pointer()), C.int(group), C.int(role), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPalette) SetColor2(role QPalette__ColorRole, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColor2(C.QtObjectPtr(ptr.Pointer()), C.int(role), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPalette) SetColorGroup(cg QPalette__ColorGroup, windowText QBrushITF, button QBrushITF, light QBrushITF, dark QBrushITF, mid QBrushITF, text QBrushITF, bright_text QBrushITF, base QBrushITF, window QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPalette_SetColorGroup(C.QtObjectPtr(ptr.Pointer()), C.int(cg), C.QtObjectPtr(PointerFromQBrush(windowText)), C.QtObjectPtr(PointerFromQBrush(button)), C.QtObjectPtr(PointerFromQBrush(light)), C.QtObjectPtr(PointerFromQBrush(dark)), C.QtObjectPtr(PointerFromQBrush(mid)), C.QtObjectPtr(PointerFromQBrush(text)), C.QtObjectPtr(PointerFromQBrush(bright_text)), C.QtObjectPtr(PointerFromQBrush(base)), C.QtObjectPtr(PointerFromQBrush(window)))
	}
}

func (ptr *QPalette) SetCurrentColorGroup(cg QPalette__ColorGroup) {
	if ptr.Pointer() != nil {
		C.QPalette_SetCurrentColorGroup(C.QtObjectPtr(ptr.Pointer()), C.int(cg))
	}
}

func (ptr *QPalette) Swap(other QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QPalette_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPalette(other)))
	}
}

func (ptr *QPalette) DestroyQPalette() {
	if ptr.Pointer() != nil {
		C.QPalette_DestroyQPalette(C.QtObjectPtr(ptr.Pointer()))
	}
}
