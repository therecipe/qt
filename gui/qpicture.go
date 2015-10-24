package gui

//#include "qpicture.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPicture struct {
	QPaintDevice
}

type QPictureITF interface {
	QPaintDeviceITF
	QPicturePTR() *QPicture
}

func PointerFromQPicture(ptr QPictureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPicturePTR().Pointer()
	}
	return nil
}

func QPictureFromPointer(ptr unsafe.Pointer) *QPicture {
	var n = new(QPicture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPicture) QPicturePTR() *QPicture {
	return ptr
}

func (ptr *QPicture) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QPicture_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPicture) Load2(dev core.QIODeviceITF, format string) bool {
	if ptr.Pointer() != nil {
		return C.QPicture_Load2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(dev)), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Load(fileName string, format string) bool {
	if ptr.Pointer() != nil {
		return C.QPicture_Load(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Play(painter QPainterITF) bool {
	if ptr.Pointer() != nil {
		return C.QPicture_Play(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter))) != 0
	}
	return false
}

func (ptr *QPicture) Save2(dev core.QIODeviceITF, format string) bool {
	if ptr.Pointer() != nil {
		return C.QPicture_Save2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(dev)), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Save(fileName string, format string) bool {
	if ptr.Pointer() != nil {
		return C.QPicture_Save(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) SetBoundingRect(r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPicture_SetBoundingRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QPicture) Swap(other QPictureITF) {
	if ptr.Pointer() != nil {
		C.QPicture_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPicture(other)))
	}
}

func (ptr *QPicture) DestroyQPicture() {
	if ptr.Pointer() != nil {
		C.QPicture_DestroyQPicture(C.QtObjectPtr(ptr.Pointer()))
	}
}
