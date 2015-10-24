package gui

//#include "qglyphrun.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGlyphRun struct {
	ptr unsafe.Pointer
}

type QGlyphRunITF interface {
	QGlyphRunPTR() *QGlyphRun
}

func (p *QGlyphRun) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGlyphRun) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGlyphRun(ptr QGlyphRunITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGlyphRunPTR().Pointer()
	}
	return nil
}

func QGlyphRunFromPointer(ptr unsafe.Pointer) *QGlyphRun {
	var n = new(QGlyphRun)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGlyphRun) QGlyphRunPTR() *QGlyphRun {
	return ptr
}

//QGlyphRun::GlyphRunFlag
type QGlyphRun__GlyphRunFlag int

var (
	QGlyphRun__Overline      = QGlyphRun__GlyphRunFlag(0x01)
	QGlyphRun__Underline     = QGlyphRun__GlyphRunFlag(0x02)
	QGlyphRun__StrikeOut     = QGlyphRun__GlyphRunFlag(0x04)
	QGlyphRun__RightToLeft   = QGlyphRun__GlyphRunFlag(0x08)
	QGlyphRun__SplitLigature = QGlyphRun__GlyphRunFlag(0x10)
)

func NewQGlyphRun() *QGlyphRun {
	return QGlyphRunFromPointer(unsafe.Pointer(C.QGlyphRun_NewQGlyphRun()))
}

func NewQGlyphRun2(other QGlyphRunITF) *QGlyphRun {
	return QGlyphRunFromPointer(unsafe.Pointer(C.QGlyphRun_NewQGlyphRun2(C.QtObjectPtr(PointerFromQGlyphRun(other)))))
}

func (ptr *QGlyphRun) Clear() {
	if ptr.Pointer() != nil {
		C.QGlyphRun_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGlyphRun) Flags() QGlyphRun__GlyphRunFlag {
	if ptr.Pointer() != nil {
		return QGlyphRun__GlyphRunFlag(C.QGlyphRun_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGlyphRun) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QGlyphRun_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) IsRightToLeft() bool {
	if ptr.Pointer() != nil {
		return C.QGlyphRun_IsRightToLeft(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) Overline() bool {
	if ptr.Pointer() != nil {
		return C.QGlyphRun_Overline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) SetBoundingRect(boundingRect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetBoundingRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(boundingRect)))
	}
}

func (ptr *QGlyphRun) SetFlag(flag QGlyphRun__GlyphRunFlag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetFlag(C.QtObjectPtr(ptr.Pointer()), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGlyphRun) SetFlags(flags QGlyphRun__GlyphRunFlag) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QGlyphRun) SetOverline(overline bool) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetOverline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(overline)))
	}
}

func (ptr *QGlyphRun) SetRawFont(rawFont QRawFontITF) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetRawFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRawFont(rawFont)))
	}
}

func (ptr *QGlyphRun) SetRightToLeft(rightToLeft bool) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetRightToLeft(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(rightToLeft)))
	}
}

func (ptr *QGlyphRun) SetStrikeOut(strikeOut bool) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetStrikeOut(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(strikeOut)))
	}
}

func (ptr *QGlyphRun) SetUnderline(underline bool) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_SetUnderline(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QGlyphRun) StrikeOut() bool {
	if ptr.Pointer() != nil {
		return C.QGlyphRun_StrikeOut(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) Swap(other QGlyphRunITF) {
	if ptr.Pointer() != nil {
		C.QGlyphRun_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGlyphRun(other)))
	}
}

func (ptr *QGlyphRun) Underline() bool {
	if ptr.Pointer() != nil {
		return C.QGlyphRun_Underline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGlyphRun) DestroyQGlyphRun() {
	if ptr.Pointer() != nil {
		C.QGlyphRun_DestroyQGlyphRun(C.QtObjectPtr(ptr.Pointer()))
	}
}
