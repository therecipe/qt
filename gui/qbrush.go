package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBrush struct {
	ptr unsafe.Pointer
}

type QBrush_ITF interface {
	QBrush_PTR() *QBrush
}

func (p *QBrush) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBrush) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBrush(ptr QBrush_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBrush_PTR().Pointer()
	}
	return nil
}

func NewQBrushFromPointer(ptr unsafe.Pointer) *QBrush {
	var n = new(QBrush)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBrush) QBrush_PTR() *QBrush {
	return ptr
}

func NewQBrush4(color core.Qt__GlobalColor, style core.Qt__BrushStyle) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush4(C.int(color), C.int(style)))
}

func (ptr *QBrush) SetColor(color QColor_ITF) {
	defer qt.Recovering("QBrush::setColor")

	if ptr.Pointer() != nil {
		C.QBrush_SetColor(ptr.Pointer(), PointerFromQColor(color))
	}
}

func NewQBrush() *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush())
}

func NewQBrush2(style core.Qt__BrushStyle) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush2(C.int(style)))
}

func NewQBrush6(color core.Qt__GlobalColor, pixmap QPixmap_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush6(C.int(color), PointerFromQPixmap(pixmap)))
}

func NewQBrush9(other QBrush_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush9(PointerFromQBrush(other)))
}

func NewQBrush3(color QColor_ITF, style core.Qt__BrushStyle) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush3(PointerFromQColor(color), C.int(style)))
}

func NewQBrush5(color QColor_ITF, pixmap QPixmap_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush5(PointerFromQColor(color), PointerFromQPixmap(pixmap)))
}

func NewQBrush10(gradient QGradient_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush10(PointerFromQGradient(gradient)))
}

func NewQBrush8(image QImage_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush8(PointerFromQImage(image)))
}

func NewQBrush7(pixmap QPixmap_ITF) *QBrush {
	defer qt.Recovering("QBrush::QBrush")

	return NewQBrushFromPointer(C.QBrush_NewQBrush7(PointerFromQPixmap(pixmap)))
}

func (ptr *QBrush) Color() *QColor {
	defer qt.Recovering("QBrush::color")

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QBrush_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBrush) Gradient() *QGradient {
	defer qt.Recovering("QBrush::gradient")

	if ptr.Pointer() != nil {
		return NewQGradientFromPointer(C.QBrush_Gradient(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBrush) IsOpaque() bool {
	defer qt.Recovering("QBrush::isOpaque")

	if ptr.Pointer() != nil {
		return C.QBrush_IsOpaque(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBrush) SetColor2(color core.Qt__GlobalColor) {
	defer qt.Recovering("QBrush::setColor")

	if ptr.Pointer() != nil {
		C.QBrush_SetColor2(ptr.Pointer(), C.int(color))
	}
}

func (ptr *QBrush) SetStyle(style core.Qt__BrushStyle) {
	defer qt.Recovering("QBrush::setStyle")

	if ptr.Pointer() != nil {
		C.QBrush_SetStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QBrush) SetTexture(pixmap QPixmap_ITF) {
	defer qt.Recovering("QBrush::setTexture")

	if ptr.Pointer() != nil {
		C.QBrush_SetTexture(ptr.Pointer(), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QBrush) SetTextureImage(image QImage_ITF) {
	defer qt.Recovering("QBrush::setTextureImage")

	if ptr.Pointer() != nil {
		C.QBrush_SetTextureImage(ptr.Pointer(), PointerFromQImage(image))
	}
}

func (ptr *QBrush) SetTransform(matrix QTransform_ITF) {
	defer qt.Recovering("QBrush::setTransform")

	if ptr.Pointer() != nil {
		C.QBrush_SetTransform(ptr.Pointer(), PointerFromQTransform(matrix))
	}
}

func (ptr *QBrush) Style() core.Qt__BrushStyle {
	defer qt.Recovering("QBrush::style")

	if ptr.Pointer() != nil {
		return core.Qt__BrushStyle(C.QBrush_Style(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBrush) Swap(other QBrush_ITF) {
	defer qt.Recovering("QBrush::swap")

	if ptr.Pointer() != nil {
		C.QBrush_Swap(ptr.Pointer(), PointerFromQBrush(other))
	}
}

func (ptr *QBrush) DestroyQBrush() {
	defer qt.Recovering("QBrush::~QBrush")

	if ptr.Pointer() != nil {
		C.QBrush_DestroyQBrush(ptr.Pointer())
	}
}
