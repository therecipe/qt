package gui

//#include "qcursor.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCursor struct {
	ptr unsafe.Pointer
}

type QCursor_ITF interface {
	QCursor_PTR() *QCursor
}

func (p *QCursor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCursor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCursor(ptr QCursor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCursor_PTR().Pointer()
	}
	return nil
}

func NewQCursorFromPointer(ptr unsafe.Pointer) *QCursor {
	var n = new(QCursor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCursor) QCursor_PTR() *QCursor {
	return ptr
}

func QCursor_SetPos2(screen QScreen_ITF, x int, y int) {
	C.QCursor_QCursor_SetPos2(PointerFromQScreen(screen), C.int(x), C.int(y))
}

func QCursor_SetPos(x int, y int) {
	C.QCursor_QCursor_SetPos(C.int(x), C.int(y))
}

func NewQCursor() *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor())
}

func NewQCursor6(other QCursor_ITF) *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor6(PointerFromQCursor(other)))
}

func NewQCursor2(shape core.Qt__CursorShape) *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor2(C.int(shape)))
}

func NewQCursor3(bitmap QBitmap_ITF, mask QBitmap_ITF, hotX int, hotY int) *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor3(PointerFromQBitmap(bitmap), PointerFromQBitmap(mask), C.int(hotX), C.int(hotY)))
}

func NewQCursor5(c QCursor_ITF) *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor5(PointerFromQCursor(c)))
}

func NewQCursor4(pixmap QPixmap_ITF, hotX int, hotY int) *QCursor {
	return NewQCursorFromPointer(C.QCursor_NewQCursor4(PointerFromQPixmap(pixmap), C.int(hotX), C.int(hotY)))
}

func (ptr *QCursor) Bitmap() *QBitmap {
	if ptr.Pointer() != nil {
		return NewQBitmapFromPointer(C.QCursor_Bitmap(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCursor) Mask() *QBitmap {
	if ptr.Pointer() != nil {
		return NewQBitmapFromPointer(C.QCursor_Mask(ptr.Pointer()))
	}
	return nil
}

func QCursor_SetPos4(screen QScreen_ITF, p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos4(PointerFromQScreen(screen), core.PointerFromQPoint(p))
}

func QCursor_SetPos3(p core.QPoint_ITF) {
	C.QCursor_QCursor_SetPos3(core.PointerFromQPoint(p))
}

func (ptr *QCursor) SetShape(shape core.Qt__CursorShape) {
	if ptr.Pointer() != nil {
		C.QCursor_SetShape(ptr.Pointer(), C.int(shape))
	}
}

func (ptr *QCursor) Shape() core.Qt__CursorShape {
	if ptr.Pointer() != nil {
		return core.Qt__CursorShape(C.QCursor_Shape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCursor) DestroyQCursor() {
	if ptr.Pointer() != nil {
		C.QCursor_DestroyQCursor(ptr.Pointer())
	}
}
