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

type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func PointerFromQPixmap(ptr QPixmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmap_PTR().Pointer()
	}
	return nil
}

func NewQPixmapFromPointer(ptr unsafe.Pointer) *QPixmap {
	var n = new(QPixmap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap {
	return ptr
}

func (ptr *QPixmap) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) IsQBitmap() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_IsQBitmap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPixmap_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) ConvertFromImage(image QImage_ITF, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_ConvertFromImage(ptr.Pointer(), PointerFromQImage(image), C.int(flags)) != 0
	}
	return false
}

func QPixmap_DefaultDepth() int {
	return int(C.QPixmap_QPixmap_DefaultDepth())
}

func (ptr *QPixmap) Detach() {
	if ptr.Pointer() != nil {
		C.QPixmap_Detach(ptr.Pointer())
	}
}

func (ptr *QPixmap) DevicePixelRatio() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPixmap_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) Fill(color QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Fill(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QPixmap) HasAlpha() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlpha(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) HasAlphaChannel() bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) Load(fileName string, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Load(ptr.Pointer(), C.CString(fileName), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) LoadFromData2(data core.QByteArray_ITF, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_LoadFromData2(ptr.Pointer(), core.PointerFromQByteArray(data), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) Save2(device core.QIODevice_ITF, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Save(fileName string, format string, quality int) bool {
	if ptr.Pointer() != nil {
		return C.QPixmap_Save(ptr.Pointer(), C.CString(fileName), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Scroll2(dx int, dy int, rect core.QRect_ITF, exposed QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll2(ptr.Pointer(), C.int(dx), C.int(dy), core.PointerFromQRect(rect), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Scroll(ptr.Pointer(), C.int(dx), C.int(dy), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	if ptr.Pointer() != nil {
		C.QPixmap_SetDevicePixelRatio(ptr.Pointer(), C.double(scaleFactor))
	}
}

func (ptr *QPixmap) SetMask(mask QBitmap_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_SetMask(ptr.Pointer(), PointerFromQBitmap(mask))
	}
}

func (ptr *QPixmap) Swap(other QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QPixmap_Swap(ptr.Pointer(), PointerFromQPixmap(other))
	}
}

func (ptr *QPixmap) DestroyQPixmap() {
	if ptr.Pointer() != nil {
		C.QPixmap_DestroyQPixmap(ptr.Pointer())
	}
}
