package gui

//#include "qtextlayout.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextLayout struct {
	ptr unsafe.Pointer
}

type QTextLayoutITF interface {
	QTextLayoutPTR() *QTextLayout
}

func (p *QTextLayout) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLayout) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLayout(ptr QTextLayoutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLayoutPTR().Pointer()
	}
	return nil
}

func QTextLayoutFromPointer(ptr unsafe.Pointer) *QTextLayout {
	var n = new(QTextLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLayout) QTextLayoutPTR() *QTextLayout {
	return ptr
}

//QTextLayout::CursorMode
type QTextLayout__CursorMode int

var (
	QTextLayout__SkipCharacters = QTextLayout__CursorMode(0)
	QTextLayout__SkipWords      = QTextLayout__CursorMode(1)
)

func (ptr *QTextLayout) DrawCursor2(painter QPainterITF, position core.QPointFITF, cursorPosition int) {
	if ptr.Pointer() != nil {
		C.QTextLayout_DrawCursor2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQPointF(position)), C.int(cursorPosition))
	}
}

func (ptr *QTextLayout) DrawCursor(painter QPainterITF, position core.QPointFITF, cursorPosition int, width int) {
	if ptr.Pointer() != nil {
		C.QTextLayout_DrawCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQPointF(position)), C.int(cursorPosition), C.int(width))
	}
}

func NewQTextLayout() *QTextLayout {
	return QTextLayoutFromPointer(unsafe.Pointer(C.QTextLayout_NewQTextLayout()))
}

func NewQTextLayout2(text string) *QTextLayout {
	return QTextLayoutFromPointer(unsafe.Pointer(C.QTextLayout_NewQTextLayout2(C.CString(text))))
}

func NewQTextLayout3(text string, font QFontITF, paintdevice QPaintDeviceITF) *QTextLayout {
	return QTextLayoutFromPointer(unsafe.Pointer(C.QTextLayout_NewQTextLayout3(C.CString(text), C.QtObjectPtr(PointerFromQFont(font)), C.QtObjectPtr(PointerFromQPaintDevice(paintdevice)))))
}

func (ptr *QTextLayout) BeginLayout() {
	if ptr.Pointer() != nil {
		C.QTextLayout_BeginLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextLayout) CacheEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QTextLayout_CacheEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTextLayout) ClearAdditionalFormats() {
	if ptr.Pointer() != nil {
		C.QTextLayout_ClearAdditionalFormats(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextLayout) ClearLayout() {
	if ptr.Pointer() != nil {
		C.QTextLayout_ClearLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextLayout) CursorMoveStyle() core.Qt__CursorMoveStyle {
	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QTextLayout_CursorMoveStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextLayout) EndLayout() {
	if ptr.Pointer() != nil {
		C.QTextLayout_EndLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTextLayout) IsValidCursorPosition(pos int) bool {
	if ptr.Pointer() != nil {
		return C.QTextLayout_IsValidCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(pos)) != 0
	}
	return false
}

func (ptr *QTextLayout) LeftCursorPosition(oldPos int) int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_LeftCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(oldPos)))
	}
	return 0
}

func (ptr *QTextLayout) LineCount() int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_LineCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextLayout) NextCursorPosition(oldPos int, mode QTextLayout__CursorMode) int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_NextCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(oldPos), C.int(mode)))
	}
	return 0
}

func (ptr *QTextLayout) PreeditAreaPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_PreeditAreaPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTextLayout) PreeditAreaText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextLayout_PreeditAreaText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextLayout) PreviousCursorPosition(oldPos int, mode QTextLayout__CursorMode) int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_PreviousCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(oldPos), C.int(mode)))
	}
	return 0
}

func (ptr *QTextLayout) RightCursorPosition(oldPos int) int {
	if ptr.Pointer() != nil {
		return int(C.QTextLayout_RightCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(oldPos)))
	}
	return 0
}

func (ptr *QTextLayout) SetCacheEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetCacheEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextLayout) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetCursorMoveStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QTextLayout) SetFont(font QFontITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QTextLayout) SetPosition(p core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetPosition(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(p)))
	}
}

func (ptr *QTextLayout) SetPreeditArea(position int, text string) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetPreeditArea(C.QtObjectPtr(ptr.Pointer()), C.int(position), C.CString(text))
	}
}

func (ptr *QTextLayout) SetText(stri string) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(stri))
	}
}

func (ptr *QTextLayout) SetTextOption(option QTextOptionITF) {
	if ptr.Pointer() != nil {
		C.QTextLayout_SetTextOption(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextOption(option)))
	}
}

func (ptr *QTextLayout) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTextLayout_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTextLayout) DestroyQTextLayout() {
	if ptr.Pointer() != nil {
		C.QTextLayout_DestroyQTextLayout(C.QtObjectPtr(ptr.Pointer()))
	}
}
