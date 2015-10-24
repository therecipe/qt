package core

//#include "qrect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRect struct {
	ptr unsafe.Pointer
}

type QRectITF interface {
	QRectPTR() *QRect
}

func (p *QRect) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRect) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRect(ptr QRectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRectPTR().Pointer()
	}
	return nil
}

func QRectFromPointer(ptr unsafe.Pointer) *QRect {
	var n = new(QRect)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRect) QRectPTR() *QRect {
	return ptr
}

func (ptr *QRect) Contains(point QPointITF, proper bool) bool {
	if ptr.Pointer() != nil {
		return C.QRect_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(point)), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) Contains4(rectangle QRectITF, proper bool) bool {
	if ptr.Pointer() != nil {
		return C.QRect_Contains4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRect(rectangle)), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) Intersects(rectangle QRectITF) bool {
	if ptr.Pointer() != nil {
		return C.QRect_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRect(rectangle))) != 0
	}
	return false
}

func NewQRect() *QRect {
	return QRectFromPointer(unsafe.Pointer(C.QRect_NewQRect()))
}

func NewQRect2(topLeft QPointITF, bottomRight QPointITF) *QRect {
	return QRectFromPointer(unsafe.Pointer(C.QRect_NewQRect2(C.QtObjectPtr(PointerFromQPoint(topLeft)), C.QtObjectPtr(PointerFromQPoint(bottomRight)))))
}

func NewQRect3(topLeft QPointITF, size QSizeITF) *QRect {
	return QRectFromPointer(unsafe.Pointer(C.QRect_NewQRect3(C.QtObjectPtr(PointerFromQPoint(topLeft)), C.QtObjectPtr(PointerFromQSize(size)))))
}

func NewQRect4(x int, y int, width int, height int) *QRect {
	return QRectFromPointer(unsafe.Pointer(C.QRect_NewQRect4(C.int(x), C.int(y), C.int(width), C.int(height))))
}

func (ptr *QRect) Adjust(dx1 int, dy1 int, dx2 int, dy2 int) {
	if ptr.Pointer() != nil {
		C.QRect_Adjust(C.QtObjectPtr(ptr.Pointer()), C.int(dx1), C.int(dy1), C.int(dx2), C.int(dy2))
	}
}

func (ptr *QRect) Bottom() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Bottom(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Contains3(x int, y int) bool {
	if ptr.Pointer() != nil {
		return C.QRect_Contains3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y)) != 0
	}
	return false
}

func (ptr *QRect) Contains2(x int, y int, proper bool) bool {
	if ptr.Pointer() != nil {
		return C.QRect_Contains2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(qt.GoBoolToInt(proper))) != 0
	}
	return false
}

func (ptr *QRect) GetCoords(x1 int, y1 int, x2 int, y2 int) {
	if ptr.Pointer() != nil {
		C.QRect_GetCoords(C.QtObjectPtr(ptr.Pointer()), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QRect) GetRect(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QRect_GetRect(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRect) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QRect_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRect) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QRect_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRect) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QRect_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRect) Left() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Left(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) MoveBottom(y int) {
	if ptr.Pointer() != nil {
		C.QRect_MoveBottom(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QRect) MoveBottomLeft(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveBottomLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) MoveBottomRight(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveBottomRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) MoveCenter(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveCenter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) MoveLeft(x int) {
	if ptr.Pointer() != nil {
		C.QRect_MoveLeft(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QRect) MoveRight(x int) {
	if ptr.Pointer() != nil {
		C.QRect_MoveRight(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QRect) MoveTo2(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveTo2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) MoveTo(x int, y int) {
	if ptr.Pointer() != nil {
		C.QRect_MoveTo(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))
	}
}

func (ptr *QRect) MoveTop(y int) {
	if ptr.Pointer() != nil {
		C.QRect_MoveTop(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QRect) MoveTopLeft(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveTopLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) MoveTopRight(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_MoveTopRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) Right() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Right(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) SetBottom(y int) {
	if ptr.Pointer() != nil {
		C.QRect_SetBottom(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QRect) SetBottomLeft(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetBottomLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) SetBottomRight(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetBottomRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	if ptr.Pointer() != nil {
		C.QRect_SetCoords(C.QtObjectPtr(ptr.Pointer()), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QRect) SetHeight(height int) {
	if ptr.Pointer() != nil {
		C.QRect_SetHeight(C.QtObjectPtr(ptr.Pointer()), C.int(height))
	}
}

func (ptr *QRect) SetLeft(x int) {
	if ptr.Pointer() != nil {
		C.QRect_SetLeft(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QRect) SetRect(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QRect_SetRect(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QRect) SetRight(x int) {
	if ptr.Pointer() != nil {
		C.QRect_SetRight(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QRect) SetSize(size QSizeITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSize(size)))
	}
}

func (ptr *QRect) SetTop(y int) {
	if ptr.Pointer() != nil {
		C.QRect_SetTop(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QRect) SetTopLeft(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetTopLeft(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) SetTopRight(position QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_SetTopRight(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(position)))
	}
}

func (ptr *QRect) SetWidth(width int) {
	if ptr.Pointer() != nil {
		C.QRect_SetWidth(C.QtObjectPtr(ptr.Pointer()), C.int(width))
	}
}

func (ptr *QRect) SetX(x int) {
	if ptr.Pointer() != nil {
		C.QRect_SetX(C.QtObjectPtr(ptr.Pointer()), C.int(x))
	}
}

func (ptr *QRect) SetY(y int) {
	if ptr.Pointer() != nil {
		C.QRect_SetY(C.QtObjectPtr(ptr.Pointer()), C.int(y))
	}
}

func (ptr *QRect) Top() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Top(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Translate2(offset QPointITF) {
	if ptr.Pointer() != nil {
		C.QRect_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(offset)))
	}
}

func (ptr *QRect) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QRect_Translate(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy))
	}
}

func (ptr *QRect) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) X() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRect) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QRect_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
