package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBackingStore struct {
	ptr unsafe.Pointer
}

type QBackingStore_ITF interface {
	QBackingStore_PTR() *QBackingStore
}

func (p *QBackingStore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBackingStore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBackingStore(ptr QBackingStore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBackingStore_PTR().Pointer()
	}
	return nil
}

func NewQBackingStoreFromPointer(ptr unsafe.Pointer) *QBackingStore {
	var n = new(QBackingStore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBackingStore) QBackingStore_PTR() *QBackingStore {
	return ptr
}

func (ptr *QBackingStore) PaintDevice() *QPaintDevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::paintDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QBackingStore_PaintDevice(ptr.Pointer()))
	}
	return nil
}

func NewQBackingStore(window QWindow_ITF) *QBackingStore {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::QBackingStore")
		}
	}()

	return NewQBackingStoreFromPointer(C.QBackingStore_NewQBackingStore(PointerFromQWindow(window)))
}

func (ptr *QBackingStore) BeginPaint(region QRegion_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::beginPaint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_BeginPaint(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QBackingStore) EndPaint() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::endPaint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_EndPaint(ptr.Pointer())
	}
}

func (ptr *QBackingStore) Flush(region QRegion_ITF, win QWindow_ITF, offset core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::flush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_Flush(ptr.Pointer(), PointerFromQRegion(region), PointerFromQWindow(win), core.PointerFromQPoint(offset))
	}
}

func (ptr *QBackingStore) HasStaticContents() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::hasStaticContents")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBackingStore_HasStaticContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBackingStore) Resize(size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::resize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_Resize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QBackingStore) Scroll(area QRegion_ITF, dx int, dy int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::scroll")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBackingStore_Scroll(ptr.Pointer(), PointerFromQRegion(area), C.int(dx), C.int(dy)) != 0
	}
	return false
}

func (ptr *QBackingStore) SetStaticContents(region QRegion_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::setStaticContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_SetStaticContents(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QBackingStore) StaticContents() *QRegion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::staticContents")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QBackingStore_StaticContents(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBackingStore) Window() *QWindow {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::window")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QBackingStore_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBackingStore) DestroyQBackingStore() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBackingStore::~QBackingStore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBackingStore_DestroyQBackingStore(ptr.Pointer())
	}
}
