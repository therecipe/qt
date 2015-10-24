package gui

//#include "qopenglpixeltransferoptions.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLPixelTransferOptions struct {
	ptr unsafe.Pointer
}

type QOpenGLPixelTransferOptionsITF interface {
	QOpenGLPixelTransferOptionsPTR() *QOpenGLPixelTransferOptions
}

func (p *QOpenGLPixelTransferOptions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLPixelTransferOptions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLPixelTransferOptions(ptr QOpenGLPixelTransferOptionsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLPixelTransferOptionsPTR().Pointer()
	}
	return nil
}

func QOpenGLPixelTransferOptionsFromPointer(ptr unsafe.Pointer) *QOpenGLPixelTransferOptions {
	var n = new(QOpenGLPixelTransferOptions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLPixelTransferOptions) QOpenGLPixelTransferOptionsPTR() *QOpenGLPixelTransferOptions {
	return ptr
}

func NewQOpenGLPixelTransferOptions() *QOpenGLPixelTransferOptions {
	return QOpenGLPixelTransferOptionsFromPointer(unsafe.Pointer(C.QOpenGLPixelTransferOptions_NewQOpenGLPixelTransferOptions()))
}

func (ptr *QOpenGLPixelTransferOptions) Alignment() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) ImageHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_ImageHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) IsLeastSignificantBitFirst() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLPixelTransferOptions_IsLeastSignificantBitFirst(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLPixelTransferOptions) IsSwapBytesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLPixelTransferOptions_IsSwapBytesEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLPixelTransferOptions) SetAlignment(alignment int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetImageHeight(imageHeight int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetImageHeight(C.QtObjectPtr(ptr.Pointer()), C.int(imageHeight))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetSkipImages(skipImages int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetSkipImages(C.QtObjectPtr(ptr.Pointer()), C.int(skipImages))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetSkipPixels(skipPixels int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetSkipPixels(C.QtObjectPtr(ptr.Pointer()), C.int(skipPixels))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetSkipRows(skipRows int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetSkipRows(C.QtObjectPtr(ptr.Pointer()), C.int(skipRows))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SkipImages() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_SkipImages(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) SkipPixels() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_SkipPixels(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) SkipRows() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_SkipRows(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) RowLength() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLPixelTransferOptions_RowLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLPixelTransferOptions) SetLeastSignificantByteFirst(lsbFirst bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetLeastSignificantByteFirst(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(lsbFirst)))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetRowLength(rowLength int) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetRowLength(C.QtObjectPtr(ptr.Pointer()), C.int(rowLength))
	}
}

func (ptr *QOpenGLPixelTransferOptions) SetSwapBytesEnabled(swapBytes bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_SetSwapBytesEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(swapBytes)))
	}
}

func (ptr *QOpenGLPixelTransferOptions) DestroyQOpenGLPixelTransferOptions() {
	if ptr.Pointer() != nil {
		C.QOpenGLPixelTransferOptions_DestroyQOpenGLPixelTransferOptions(C.QtObjectPtr(ptr.Pointer()))
	}
}
