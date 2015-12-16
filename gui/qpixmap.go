package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QPixmap::depth")

	if ptr.Pointer() != nil {
		return int(C.QPixmap_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) Height() int {
	defer qt.Recovering("QPixmap::height")

	if ptr.Pointer() != nil {
		return int(C.QPixmap_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) IsNull() bool {
	defer qt.Recovering("QPixmap::isNull")

	if ptr.Pointer() != nil {
		return C.QPixmap_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) IsQBitmap() bool {
	defer qt.Recovering("QPixmap::isQBitmap")

	if ptr.Pointer() != nil {
		return C.QPixmap_IsQBitmap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) Rect() *core.QRect {
	defer qt.Recovering("QPixmap::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QPixmap_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPixmap) Size() *core.QSize {
	defer qt.Recovering("QPixmap::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QPixmap_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPixmap) Width() int {
	defer qt.Recovering("QPixmap::width")

	if ptr.Pointer() != nil {
		return int(C.QPixmap_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) CacheKey() int64 {
	defer qt.Recovering("QPixmap::cacheKey")

	if ptr.Pointer() != nil {
		return int64(C.QPixmap_CacheKey(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) ConvertFromImage(image QImage_ITF, flags core.Qt__ImageConversionFlag) bool {
	defer qt.Recovering("QPixmap::convertFromImage")

	if ptr.Pointer() != nil {
		return C.QPixmap_ConvertFromImage(ptr.Pointer(), PointerFromQImage(image), C.int(flags)) != 0
	}
	return false
}

func QPixmap_DefaultDepth() int {
	defer qt.Recovering("QPixmap::defaultDepth")

	return int(C.QPixmap_QPixmap_DefaultDepth())
}

func (ptr *QPixmap) Detach() {
	defer qt.Recovering("QPixmap::detach")

	if ptr.Pointer() != nil {
		C.QPixmap_Detach(ptr.Pointer())
	}
}

func (ptr *QPixmap) DevicePixelRatio() float64 {
	defer qt.Recovering("QPixmap::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QPixmap_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPixmap) Fill(color QColor_ITF) {
	defer qt.Recovering("QPixmap::fill")

	if ptr.Pointer() != nil {
		C.QPixmap_Fill(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QPixmap) HasAlpha() bool {
	defer qt.Recovering("QPixmap::hasAlpha")

	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlpha(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) HasAlphaChannel() bool {
	defer qt.Recovering("QPixmap::hasAlphaChannel")

	if ptr.Pointer() != nil {
		return C.QPixmap_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPixmap) Load(fileName string, format string, flags core.Qt__ImageConversionFlag) bool {
	defer qt.Recovering("QPixmap::load")

	if ptr.Pointer() != nil {
		return C.QPixmap_Load(ptr.Pointer(), C.CString(fileName), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) LoadFromData2(data core.QByteArray_ITF, format string, flags core.Qt__ImageConversionFlag) bool {
	defer qt.Recovering("QPixmap::loadFromData")

	if ptr.Pointer() != nil {
		return C.QPixmap_LoadFromData2(ptr.Pointer(), core.PointerFromQByteArray(data), C.CString(format), C.int(flags)) != 0
	}
	return false
}

func (ptr *QPixmap) Save2(device core.QIODevice_ITF, format string, quality int) bool {
	defer qt.Recovering("QPixmap::save")

	if ptr.Pointer() != nil {
		return C.QPixmap_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Save(fileName string, format string, quality int) bool {
	defer qt.Recovering("QPixmap::save")

	if ptr.Pointer() != nil {
		return C.QPixmap_Save(ptr.Pointer(), C.CString(fileName), C.CString(format), C.int(quality)) != 0
	}
	return false
}

func (ptr *QPixmap) Scroll2(dx int, dy int, rect core.QRect_ITF, exposed QRegion_ITF) {
	defer qt.Recovering("QPixmap::scroll")

	if ptr.Pointer() != nil {
		C.QPixmap_Scroll2(ptr.Pointer(), C.int(dx), C.int(dy), core.PointerFromQRect(rect), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int, exposed QRegion_ITF) {
	defer qt.Recovering("QPixmap::scroll")

	if ptr.Pointer() != nil {
		C.QPixmap_Scroll(ptr.Pointer(), C.int(dx), C.int(dy), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQRegion(exposed))
	}
}

func (ptr *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	defer qt.Recovering("QPixmap::setDevicePixelRatio")

	if ptr.Pointer() != nil {
		C.QPixmap_SetDevicePixelRatio(ptr.Pointer(), C.double(scaleFactor))
	}
}

func (ptr *QPixmap) SetMask(mask QBitmap_ITF) {
	defer qt.Recovering("QPixmap::setMask")

	if ptr.Pointer() != nil {
		C.QPixmap_SetMask(ptr.Pointer(), PointerFromQBitmap(mask))
	}
}

func (ptr *QPixmap) Swap(other QPixmap_ITF) {
	defer qt.Recovering("QPixmap::swap")

	if ptr.Pointer() != nil {
		C.QPixmap_Swap(ptr.Pointer(), PointerFromQPixmap(other))
	}
}

func (ptr *QPixmap) DestroyQPixmap() {
	defer qt.Recovering("QPixmap::~QPixmap")

	if ptr.Pointer() != nil {
		C.QPixmap_DestroyQPixmap(ptr.Pointer())
	}
}
