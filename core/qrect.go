package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRect struct {
	ptr unsafe.Pointer
}

type QRect_ITF interface {
	QRect_PTR() *QRect
}

func (p *QRect) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRect) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRect(ptr QRect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRect_PTR().Pointer()
	}
	return nil
}

func NewQRectFromPointer(ptr unsafe.Pointer) *QRect {
	var n = new(QRect)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRect) QRect_PTR() *QRect {
	return ptr
}

func (ptr *QRect) Contains(point QPoint_ITF, proper bool) bool {
	defer qt.Recovering("QRect::contains")

	if ptr.Pointer() != nil {
		return C.QRect_Contains(ptr.Pointer(), PointerFromQPoint(point), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) Contains4(rectangle QRect_ITF, proper bool) bool {
	defer qt.Recovering("QRect::contains")

	if ptr.Pointer() != nil {
		return C.QRect_Contains4(ptr.Pointer(), PointerFromQRect(rectangle), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) Intersects(rectangle QRect_ITF) bool {
	defer qt.Recovering("QRect::intersects")

	if ptr.Pointer() != nil {
		return C.QRect_Intersects(ptr.Pointer(), PointerFromQRect(rectangle)) != 0
	}
	return false
}

func NewQRect() *QRect {
	defer qt.Recovering("QRect::QRect")

	return NewQRectFromPointer(C.QRect_NewQRect())
}

func NewQRect2(topLeft QPoint_ITF, bottomRight QPoint_ITF) *QRect {
	defer qt.Recovering("QRect::QRect")

	return NewQRectFromPointer(C.QRect_NewQRect2(PointerFromQPoint(topLeft), PointerFromQPoint(bottomRight)))
}

func NewQRect3(topLeft QPoint_ITF, size QSize_ITF) *QRect {
	defer qt.Recovering("QRect::QRect")

	return NewQRectFromPointer(C.QRect_NewQRect3(PointerFromQPoint(topLeft), PointerFromQSize(size)))
}

func NewQRect4(x int, y int, width int, height int) *QRect {
	defer qt.Recovering("QRect::QRect")

	return NewQRectFromPointer(C.QRect_NewQRect4(C.int(x), C.int(y), C.int(width), C.int(height)))
}

func (ptr *QRect) Adjust(dx1 int, dy1 int, dx2 int, dy2 int) {
	defer qt.Recovering("QRect::adjust")

	if ptr.Pointer() != nil {
		C.QRect_Adjust(ptr.Pointer(), C.int(dx1), C.int(dy1), C.int(dx2), C.int(dy2))
	}
}

func (ptr *QRect) Bottom() int {
	defer qt.Recovering("QRect::bottom")

	if ptr.Pointer() != nil {
		return int(C.QRect_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) BottomLeft() *QPoint {
	defer qt.Recovering("QRect::bottomLeft")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QRect_BottomLeft(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRect) BottomRight() *QPoint {
	defer qt.Recovering("QRect::bottomRight")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QRect_BottomRight(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRect) Center() *QPoint {
	defer qt.Recovering("QRect::center")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QRect_Center(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRect) Contains3(x int, y int) bool {
	defer qt.Recovering("QRect::contains")

	if ptr.Pointer() != nil {
		return C.QRect_Contains3(ptr.Pointer(), C.int(x), C.int(y)) != 0
	}
	return false
}

func (ptr *QRect) Contains2(x int, y int, proper bool) bool {
	defer qt.Recovering("QRect::contains")

	if ptr.Pointer() != nil {
		return C.QRect_Contains2(ptr.Pointer(), C.int(x), C.int(y), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) GetCoords(x1 int, y1 int, x2 int, y2 int) {
	defer qt.Recovering("QRect::getCoords")

	if ptr.Pointer() != nil {
		C.QRect_GetCoords(ptr.Pointer(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QRect) GetRect(x int, y int, width int, height int) {
	defer qt.Recovering("QRect::getRect")

	if ptr.Pointer() != nil {
		C.QRect_GetRect(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRect) Height() int {
	defer qt.Recovering("QRect::height")

	if ptr.Pointer() != nil {
		return int(C.QRect_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) IsEmpty() bool {
	defer qt.Recovering("QRect::isEmpty")

	if ptr.Pointer() != nil {
		return C.QRect_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRect) IsNull() bool {
	defer qt.Recovering("QRect::isNull")

	if ptr.Pointer() != nil {
		return C.QRect_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRect) IsValid() bool {
	defer qt.Recovering("QRect::isValid")

	if ptr.Pointer() != nil {
		return C.QRect_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRect) Left() int {
	defer qt.Recovering("QRect::left")

	if ptr.Pointer() != nil {
		return int(C.QRect_Left(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) MoveBottom(y int) {
	defer qt.Recovering("QRect::moveBottom")

	if ptr.Pointer() != nil {
		C.QRect_MoveBottom(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QRect) MoveBottomLeft(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveBottomLeft")

	if ptr.Pointer() != nil {
		C.QRect_MoveBottomLeft(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) MoveBottomRight(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveBottomRight")

	if ptr.Pointer() != nil {
		C.QRect_MoveBottomRight(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) MoveCenter(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveCenter")

	if ptr.Pointer() != nil {
		C.QRect_MoveCenter(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) MoveLeft(x int) {
	defer qt.Recovering("QRect::moveLeft")

	if ptr.Pointer() != nil {
		C.QRect_MoveLeft(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QRect) MoveRight(x int) {
	defer qt.Recovering("QRect::moveRight")

	if ptr.Pointer() != nil {
		C.QRect_MoveRight(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QRect) MoveTo2(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveTo")

	if ptr.Pointer() != nil {
		C.QRect_MoveTo2(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) MoveTo(x int, y int) {
	defer qt.Recovering("QRect::moveTo")

	if ptr.Pointer() != nil {
		C.QRect_MoveTo(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QRect) MoveTop(y int) {
	defer qt.Recovering("QRect::moveTop")

	if ptr.Pointer() != nil {
		C.QRect_MoveTop(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QRect) MoveTopLeft(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveTopLeft")

	if ptr.Pointer() != nil {
		C.QRect_MoveTopLeft(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) MoveTopRight(position QPoint_ITF) {
	defer qt.Recovering("QRect::moveTopRight")

	if ptr.Pointer() != nil {
		C.QRect_MoveTopRight(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) Right() int {
	defer qt.Recovering("QRect::right")

	if ptr.Pointer() != nil {
		return int(C.QRect_Right(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) SetBottom(y int) {
	defer qt.Recovering("QRect::setBottom")

	if ptr.Pointer() != nil {
		C.QRect_SetBottom(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QRect) SetBottomLeft(position QPoint_ITF) {
	defer qt.Recovering("QRect::setBottomLeft")

	if ptr.Pointer() != nil {
		C.QRect_SetBottomLeft(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) SetBottomRight(position QPoint_ITF) {
	defer qt.Recovering("QRect::setBottomRight")

	if ptr.Pointer() != nil {
		C.QRect_SetBottomRight(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	defer qt.Recovering("QRect::setCoords")

	if ptr.Pointer() != nil {
		C.QRect_SetCoords(ptr.Pointer(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QRect) SetHeight(height int) {
	defer qt.Recovering("QRect::setHeight")

	if ptr.Pointer() != nil {
		C.QRect_SetHeight(ptr.Pointer(), C.int(height))
	}
}

func (ptr *QRect) SetLeft(x int) {
	defer qt.Recovering("QRect::setLeft")

	if ptr.Pointer() != nil {
		C.QRect_SetLeft(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QRect) SetRect(x int, y int, width int, height int) {
	defer qt.Recovering("QRect::setRect")

	if ptr.Pointer() != nil {
		C.QRect_SetRect(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRect) SetRight(x int) {
	defer qt.Recovering("QRect::setRight")

	if ptr.Pointer() != nil {
		C.QRect_SetRight(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QRect) SetSize(size QSize_ITF) {
	defer qt.Recovering("QRect::setSize")

	if ptr.Pointer() != nil {
		C.QRect_SetSize(ptr.Pointer(), PointerFromQSize(size))
	}
}

func (ptr *QRect) SetTop(y int) {
	defer qt.Recovering("QRect::setTop")

	if ptr.Pointer() != nil {
		C.QRect_SetTop(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QRect) SetTopLeft(position QPoint_ITF) {
	defer qt.Recovering("QRect::setTopLeft")

	if ptr.Pointer() != nil {
		C.QRect_SetTopLeft(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) SetTopRight(position QPoint_ITF) {
	defer qt.Recovering("QRect::setTopRight")

	if ptr.Pointer() != nil {
		C.QRect_SetTopRight(ptr.Pointer(), PointerFromQPoint(position))
	}
}

func (ptr *QRect) SetWidth(width int) {
	defer qt.Recovering("QRect::setWidth")

	if ptr.Pointer() != nil {
		C.QRect_SetWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QRect) SetX(x int) {
	defer qt.Recovering("QRect::setX")

	if ptr.Pointer() != nil {
		C.QRect_SetX(ptr.Pointer(), C.int(x))
	}
}

func (ptr *QRect) SetY(y int) {
	defer qt.Recovering("QRect::setY")

	if ptr.Pointer() != nil {
		C.QRect_SetY(ptr.Pointer(), C.int(y))
	}
}

func (ptr *QRect) Top() int {
	defer qt.Recovering("QRect::top")

	if ptr.Pointer() != nil {
		return int(C.QRect_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) TopLeft() *QPoint {
	defer qt.Recovering("QRect::topLeft")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QRect_TopLeft(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRect) TopRight() *QPoint {
	defer qt.Recovering("QRect::topRight")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QRect_TopRight(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRect) Translate2(offset QPoint_ITF) {
	defer qt.Recovering("QRect::translate")

	if ptr.Pointer() != nil {
		C.QRect_Translate2(ptr.Pointer(), PointerFromQPoint(offset))
	}
}

func (ptr *QRect) Translate(dx int, dy int) {
	defer qt.Recovering("QRect::translate")

	if ptr.Pointer() != nil {
		C.QRect_Translate(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QRect) Width() int {
	defer qt.Recovering("QRect::width")

	if ptr.Pointer() != nil {
		return int(C.QRect_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) X() int {
	defer qt.Recovering("QRect::x")

	if ptr.Pointer() != nil {
		return int(C.QRect_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRect) Y() int {
	defer qt.Recovering("QRect::y")

	if ptr.Pointer() != nil {
		return int(C.QRect_Y(ptr.Pointer()))
	}
	return 0
}
