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

type QCursorITF interface {
	QCursorPTR() *QCursor
}

func (p *QCursor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCursor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCursor(ptr QCursorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCursorPTR().Pointer()
	}
	return nil
}

func QCursorFromPointer(ptr unsafe.Pointer) *QCursor {
	var n = new(QCursor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCursor) QCursorPTR() *QCursor {
	return ptr
}

func QCursor_SetPos2(screen QScreenITF, x int, y int) {
	C.QCursor_QCursor_SetPos2(C.QtObjectPtr(PointerFromQScreen(screen)), C.int(x), C.int(y))
}

func QCursor_SetPos(x int, y int) {
	C.QCursor_QCursor_SetPos(C.int(x), C.int(y))
}

func NewQCursor() *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor()))
}

func NewQCursor6(other QCursorITF) *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor6(C.QtObjectPtr(PointerFromQCursor(other)))))
}

func NewQCursor2(shape core.Qt__CursorShape) *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor2(C.int(shape))))
}

func NewQCursor3(bitmap QBitmapITF, mask QBitmapITF, hotX int, hotY int) *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor3(C.QtObjectPtr(PointerFromQBitmap(bitmap)), C.QtObjectPtr(PointerFromQBitmap(mask)), C.int(hotX), C.int(hotY))))
}

func NewQCursor5(c QCursorITF) *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor5(C.QtObjectPtr(PointerFromQCursor(c)))))
}

func NewQCursor4(pixmap QPixmapITF, hotX int, hotY int) *QCursor {
	return QCursorFromPointer(unsafe.Pointer(C.QCursor_NewQCursor4(C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(hotX), C.int(hotY))))
}

func (ptr *QCursor) Bitmap() *QBitmap {
	if ptr.Pointer() != nil {
		return QBitmapFromPointer(unsafe.Pointer(C.QCursor_Bitmap(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QCursor) Mask() *QBitmap {
	if ptr.Pointer() != nil {
		return QBitmapFromPointer(unsafe.Pointer(C.QCursor_Mask(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QCursor_SetPos4(screen QScreenITF, p core.QPointITF) {
	C.QCursor_QCursor_SetPos4(C.QtObjectPtr(PointerFromQScreen(screen)), C.QtObjectPtr(core.PointerFromQPoint(p)))
}

func QCursor_SetPos3(p core.QPointITF) {
	C.QCursor_QCursor_SetPos3(C.QtObjectPtr(core.PointerFromQPoint(p)))
}

func (ptr *QCursor) SetShape(shape core.Qt__CursorShape) {
	if ptr.Pointer() != nil {
		C.QCursor_SetShape(C.QtObjectPtr(ptr.Pointer()), C.int(shape))
	}
}

func (ptr *QCursor) Shape() core.Qt__CursorShape {
	if ptr.Pointer() != nil {
		return core.Qt__CursorShape(C.QCursor_Shape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCursor) DestroyQCursor() {
	if ptr.Pointer() != nil {
		C.QCursor_DestroyQCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}
