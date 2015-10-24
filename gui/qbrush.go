package gui

//#include "qbrush.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBrush struct {
	ptr unsafe.Pointer
}

type QBrushITF interface {
	QBrushPTR() *QBrush
}

func (p *QBrush) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBrush) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBrush(ptr QBrushITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBrushPTR().Pointer()
	}
	return nil
}

func QBrushFromPointer(ptr unsafe.Pointer) *QBrush {
	var n = new(QBrush)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBrush) QBrushPTR() *QBrush {
	return ptr
}

func NewQBrush4(color core.Qt__GlobalColor, style core.Qt__BrushStyle) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush4(C.int(color), C.int(style))))
}

func (ptr *QBrush) SetColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QBrush_SetColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func NewQBrush() *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush()))
}

func NewQBrush2(style core.Qt__BrushStyle) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush2(C.int(style))))
}

func NewQBrush6(color core.Qt__GlobalColor, pixmap QPixmapITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush6(C.int(color), C.QtObjectPtr(PointerFromQPixmap(pixmap)))))
}

func NewQBrush9(other QBrushITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush9(C.QtObjectPtr(PointerFromQBrush(other)))))
}

func NewQBrush3(color QColorITF, style core.Qt__BrushStyle) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush3(C.QtObjectPtr(PointerFromQColor(color)), C.int(style))))
}

func NewQBrush5(color QColorITF, pixmap QPixmapITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush5(C.QtObjectPtr(PointerFromQColor(color)), C.QtObjectPtr(PointerFromQPixmap(pixmap)))))
}

func NewQBrush10(gradient QGradientITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush10(C.QtObjectPtr(PointerFromQGradient(gradient)))))
}

func NewQBrush8(image QImageITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush8(C.QtObjectPtr(PointerFromQImage(image)))))
}

func NewQBrush7(pixmap QPixmapITF) *QBrush {
	return QBrushFromPointer(unsafe.Pointer(C.QBrush_NewQBrush7(C.QtObjectPtr(PointerFromQPixmap(pixmap)))))
}

func (ptr *QBrush) Gradient() *QGradient {
	if ptr.Pointer() != nil {
		return QGradientFromPointer(unsafe.Pointer(C.QBrush_Gradient(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QBrush) IsOpaque() bool {
	if ptr.Pointer() != nil {
		return C.QBrush_IsOpaque(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBrush) SetColor2(color core.Qt__GlobalColor) {
	if ptr.Pointer() != nil {
		C.QBrush_SetColor2(C.QtObjectPtr(ptr.Pointer()), C.int(color))
	}
}

func (ptr *QBrush) SetStyle(style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QBrush_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QBrush) SetTexture(pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QBrush_SetTexture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QBrush) SetTextureImage(image QImageITF) {
	if ptr.Pointer() != nil {
		C.QBrush_SetTextureImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image)))
	}
}

func (ptr *QBrush) SetTransform(matrix QTransformITF) {
	if ptr.Pointer() != nil {
		C.QBrush_SetTransform(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTransform(matrix)))
	}
}

func (ptr *QBrush) Style() core.Qt__BrushStyle {
	if ptr.Pointer() != nil {
		return core.Qt__BrushStyle(C.QBrush_Style(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBrush) Swap(other QBrushITF) {
	if ptr.Pointer() != nil {
		C.QBrush_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(other)))
	}
}

func (ptr *QBrush) DestroyQBrush() {
	if ptr.Pointer() != nil {
		C.QBrush_DestroyQBrush(C.QtObjectPtr(ptr.Pointer()))
	}
}
