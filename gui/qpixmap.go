package gui

//#include "qpixmap.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPixmap struct {
	QPaintDevice
}

type QPixmapITF interface {
	QPaintDeviceITF
	QPixmapPTR() *QPixmap
}

func PointerFromQPixmap(ptr QPixmapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmapPTR().Pointer()
	}
	return nil
}

func QPixmapFromPointer(ptr unsafe.Pointer) *QPixmap {
	var n = new(QPixmap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixmap) QPixmapPTR() *QPixmap {
	return ptr
}

func (ptr *QPixmap) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPixmap) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPixmap) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPixmap) IsQBitmap() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_IsQBitmap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPixmap) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPixmap) ConvertFromImage(image QImageITF, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_ConvertFromImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image)), C.int(flags)) != 0
	}
	return false
}

func QPixmap_DefaultDepth() int {
	return int(C.QPixmap_QPixmap_DefaultDepth())
}

func (ptr *QPixmap) Detach() {
	if ptr.Pointer() != nil {
		C.QPixmap_Detach(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPixmap) Fill(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Fill(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPixmap) HasAlpha() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlpha(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPixmap) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlphaChannel(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPixmap) Load(fileName string, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) LoadFromData2(data core.QByteArrayITF, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_LoadFromData2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(data)), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) Save2(device core.QIODeviceITF, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Save2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Save(fileName string, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Save(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Scroll2(dx int, dy int, rect core.QRectITF, exposed QRegionITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll2(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy), C.QtObjectPtr(core.PointerFromQRect(rect)), C.QtObjectPtr(PointerFromQRegion(exposed)))
	}
}

func (ptr *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegionITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy), C.int(x), C.int(y), C.int(width), C.int(height), C.QtObjectPtr(PointerFromQRegion(exposed)))
	}
}

func (ptr *QPixmap) SetMask(mask QBitmapITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_SetMask(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBitmap(mask)))
	}
}

func (ptr *QPixmap) Swap(other QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPixmap(other)))
	}
}

func (ptr *QPixmap) DestroyQPixmap() {
	if ptr.Pointer() != nil {
		C.QPixmap_DestroyQPixmap(C.QtObjectPtr(ptr.Pointer()))
	}
}
