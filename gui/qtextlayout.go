package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTextLayout struct {
	ptr unsafe.Pointer
}

type QTextLayout_ITF interface {
	QTextLayout_PTR() *QTextLayout
}

func (p *QTextLayout) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextLayout) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextLayout(ptr QTextLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextLayout_PTR().Pointer()
	}
	return nil
}

func NewQTextLayoutFromPointer(ptr unsafe.Pointer) *QTextLayout {
	var n = new(QTextLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextLayout) QTextLayout_PTR() *QTextLayout {
	return ptr
}

//QTextLayout::CursorMode
type QTextLayout__CursorMode int64

const (
	QTextLayout__SkipCharacters = QTextLayout__CursorMode(0)
	QTextLayout__SkipWords      = QTextLayout__CursorMode(1)
)

func (ptr *QTextLayout) DrawCursor2(painter QPainter_ITF, position core.QPointF_ITF, cursorPosition int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::drawCursor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_DrawCursor2(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQPointF(position), C.int(cursorPosition))
	}
}

func (ptr *QTextLayout) DrawCursor(painter QPainter_ITF, position core.QPointF_ITF, cursorPosition int, width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::drawCursor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_DrawCursor(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQPointF(position), C.int(cursorPosition), C.int(width))
	}
}

func NewQTextLayout() *QTextLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::QTextLayout")
		}
	}()

	return NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout())
}

func NewQTextLayout2(text string) *QTextLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::QTextLayout")
		}
	}()

	return NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout2(C.CString(text)))
}

func NewQTextLayout3(text string, font QFont_ITF, paintdevice QPaintDevice_ITF) *QTextLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::QTextLayout")
		}
	}()

	return NewQTextLayoutFromPointer(C.QTextLayout_NewQTextLayout3(C.CString(text), PointerFromQFont(font), PointerFromQPaintDevice(paintdevice)))
}

func (ptr *QTextLayout) BeginLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::beginLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_BeginLayout(ptr.Pointer())
	}
}

func (ptr *QTextLayout) CacheEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::cacheEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextLayout_CacheEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextLayout) ClearAdditionalFormats() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::clearAdditionalFormats")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_ClearAdditionalFormats(ptr.Pointer())
	}
}

func (ptr *QTextLayout) ClearLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::clearLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_ClearLayout(ptr.Pointer())
	}
}

func (ptr *QTextLayout) CursorMoveStyle() core.Qt__CursorMoveStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::cursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CursorMoveStyle(C.QTextLayout_CursorMoveStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) EndLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::endLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_EndLayout(ptr.Pointer())
	}
}

func (ptr *QTextLayout) IsValidCursorPosition(pos int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::isValidCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextLayout_IsValidCursorPosition(ptr.Pointer(), C.int(pos)) != 0
	}
	return false
}

func (ptr *QTextLayout) LeftCursorPosition(oldPos int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::leftCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_LeftCursorPosition(ptr.Pointer(), C.int(oldPos)))
	}
	return 0
}

func (ptr *QTextLayout) LineCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::lineCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_LineCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) MaximumWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::maximumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLayout_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) MinimumWidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::minimumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextLayout_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) NextCursorPosition(oldPos int, mode QTextLayout__CursorMode) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::nextCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_NextCursorPosition(ptr.Pointer(), C.int(oldPos), C.int(mode)))
	}
	return 0
}

func (ptr *QTextLayout) PreeditAreaPosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::preeditAreaPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_PreeditAreaPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextLayout) PreeditAreaText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::preeditAreaText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextLayout_PreeditAreaText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextLayout) PreviousCursorPosition(oldPos int, mode QTextLayout__CursorMode) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::previousCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_PreviousCursorPosition(ptr.Pointer(), C.int(oldPos), C.int(mode)))
	}
	return 0
}

func (ptr *QTextLayout) RightCursorPosition(oldPos int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::rightCursorPosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextLayout_RightCursorPosition(ptr.Pointer(), C.int(oldPos)))
	}
	return 0
}

func (ptr *QTextLayout) SetCacheEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setCacheEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetCacheEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QTextLayout) SetCursorMoveStyle(style core.Qt__CursorMoveStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setCursorMoveStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetCursorMoveStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextLayout) SetFont(font QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QTextLayout) SetPosition(p core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setPosition")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetPosition(ptr.Pointer(), core.PointerFromQPointF(p))
	}
}

func (ptr *QTextLayout) SetPreeditArea(position int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setPreeditArea")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetPreeditArea(ptr.Pointer(), C.int(position), C.CString(text))
	}
}

func (ptr *QTextLayout) SetText(stri string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetText(ptr.Pointer(), C.CString(stri))
	}
}

func (ptr *QTextLayout) SetTextOption(option QTextOption_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::setTextOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_SetTextOption(ptr.Pointer(), PointerFromQTextOption(option))
	}
}

func (ptr *QTextLayout) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextLayout_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextLayout) DestroyQTextLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextLayout::~QTextLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextLayout_DestroyQTextLayout(ptr.Pointer())
	}
}
