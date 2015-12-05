package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPicture struct {
	QPaintDevice
}

type QPicture_ITF interface {
	QPaintDevice_ITF
	QPicture_PTR() *QPicture
}

func PointerFromQPicture(ptr QPicture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPicture_PTR().Pointer()
	}
	return nil
}

func NewQPictureFromPointer(ptr unsafe.Pointer) *QPicture {
	var n = new(QPicture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPicture) QPicture_PTR() *QPicture {
	return ptr
}

func (ptr *QPicture) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPicture) Load2(dev core.QIODevice_ITF, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_Load2(ptr.Pointer(), core.PointerFromQIODevice(dev), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Load(fileName string, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::load")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_Load(ptr.Pointer(), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Play(painter QPainter_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::play")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_Play(ptr.Pointer(), PointerFromQPainter(painter)) != 0
	}
	return false
}

func (ptr *QPicture) Save2(dev core.QIODevice_ITF, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::save")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_Save2(ptr.Pointer(), core.PointerFromQIODevice(dev), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) Save(fileName string, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::save")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPicture_Save(ptr.Pointer(), C.CString(fileName), C.CString(format)) != 0
	}
	return false
}

func (ptr *QPicture) SetBoundingRect(r core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::setBoundingRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPicture_SetBoundingRect(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QPicture) Swap(other QPicture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPicture_Swap(ptr.Pointer(), PointerFromQPicture(other))
	}
}

func (ptr *QPicture) DestroyQPicture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPicture::~QPicture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPicture_DestroyQPicture(ptr.Pointer())
	}
}
