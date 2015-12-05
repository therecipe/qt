package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QRectF struct {
	ptr unsafe.Pointer
}

type QRectF_ITF interface {
	QRectF_PTR() *QRectF
}

func (p *QRectF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRectF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRectF(ptr QRectF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRectF_PTR().Pointer()
	}
	return nil
}

func NewQRectFFromPointer(ptr unsafe.Pointer) *QRectF {
	var n = new(QRectF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRectF) QRectF_PTR() *QRectF {
	return ptr
}

func (ptr *QRectF) Contains(point QPointF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_Contains(ptr.Pointer(), PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QRectF) Contains3(rectangle QRectF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_Contains3(ptr.Pointer(), PointerFromQRectF(rectangle)) != 0
	}
	return false
}

func (ptr *QRectF) Intersects(rectangle QRectF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::intersects")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_Intersects(ptr.Pointer(), PointerFromQRectF(rectangle)) != 0
	}
	return false
}

func NewQRectF() *QRectF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::QRectF")
		}
	}()

	return NewQRectFFromPointer(C.QRectF_NewQRectF())
}

func NewQRectF3(topLeft QPointF_ITF, bottomRight QPointF_ITF) *QRectF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::QRectF")
		}
	}()

	return NewQRectFFromPointer(C.QRectF_NewQRectF3(PointerFromQPointF(topLeft), PointerFromQPointF(bottomRight)))
}

func NewQRectF2(topLeft QPointF_ITF, size QSizeF_ITF) *QRectF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::QRectF")
		}
	}()

	return NewQRectFFromPointer(C.QRectF_NewQRectF2(PointerFromQPointF(topLeft), PointerFromQSizeF(size)))
}

func NewQRectF5(rectangle QRect_ITF) *QRectF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::QRectF")
		}
	}()

	return NewQRectFFromPointer(C.QRectF_NewQRectF5(PointerFromQRect(rectangle)))
}

func NewQRectF4(x float64, y float64, width float64, height float64) *QRectF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::QRectF")
		}
	}()

	return NewQRectFFromPointer(C.QRectF_NewQRectF4(C.double(x), C.double(y), C.double(width), C.double(height)))
}

func (ptr *QRectF) Adjust(dx1 float64, dy1 float64, dx2 float64, dy2 float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::adjust")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_Adjust(ptr.Pointer(), C.double(dx1), C.double(dy1), C.double(dx2), C.double(dy2))
	}
}

func (ptr *QRectF) Bottom() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::bottom")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Bottom(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Contains2(x float64, y float64) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_Contains2(ptr.Pointer(), C.double(x), C.double(y)) != 0
	}
	return false
}

func (ptr *QRectF) Height() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::height")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRectF) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRectF) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRectF_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRectF) Left() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::left")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Left(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) MoveBottom(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveBottom")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveBottom(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) MoveBottomLeft(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveBottomLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveBottomLeft(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) MoveBottomRight(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveBottomRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveBottomRight(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) MoveCenter(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveCenter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveCenter(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) MoveLeft(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveLeft(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QRectF) MoveRight(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveRight(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QRectF) MoveTo2(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveTo2(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) MoveTo(x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveTo(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QRectF) MoveTop(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveTop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveTop(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) MoveTopLeft(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveTopLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveTopLeft(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) MoveTopRight(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::moveTopRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_MoveTopRight(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) Right() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::right")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Right(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) SetBottom(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setBottom")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetBottom(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) SetBottomLeft(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setBottomLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetBottomLeft(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) SetBottomRight(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setBottomRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetBottomRight(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setCoords")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetCoords(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	}
}

func (ptr *QRectF) SetHeight(height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetHeight(ptr.Pointer(), C.double(height))
	}
}

func (ptr *QRectF) SetLeft(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetLeft(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QRectF) SetRect(x float64, y float64, width float64, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetRect(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QRectF) SetRight(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetRight(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QRectF) SetSize(size QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetSize(ptr.Pointer(), PointerFromQSizeF(size))
	}
}

func (ptr *QRectF) SetTop(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setTop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetTop(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) SetTopLeft(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setTopLeft")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetTopLeft(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) SetTopRight(position QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setTopRight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetTopRight(ptr.Pointer(), PointerFromQPointF(position))
	}
}

func (ptr *QRectF) SetWidth(width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QRectF) SetX(x float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setX")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetX(ptr.Pointer(), C.double(x))
	}
}

func (ptr *QRectF) SetY(y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::setY")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_SetY(ptr.Pointer(), C.double(y))
	}
}

func (ptr *QRectF) Top() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::top")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Top(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Translate2(offset QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_Translate2(ptr.Pointer(), PointerFromQPointF(offset))
	}
}

func (ptr *QRectF) Translate(dx float64, dy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRectF_Translate(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QRectF) Width() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::width")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) X() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::x")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRectF) Y() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRectF::y")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QRectF_Y(ptr.Pointer()))
	}
	return 0
}
