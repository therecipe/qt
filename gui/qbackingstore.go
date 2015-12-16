package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	defer qt.Recovering("QBackingStore::paintDevice")

	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QBackingStore_PaintDevice(ptr.Pointer()))
	}
	return nil
}

func NewQBackingStore(window QWindow_ITF) *QBackingStore {
	defer qt.Recovering("QBackingStore::QBackingStore")

	return NewQBackingStoreFromPointer(C.QBackingStore_NewQBackingStore(PointerFromQWindow(window)))
}

func (ptr *QBackingStore) BeginPaint(region QRegion_ITF) {
	defer qt.Recovering("QBackingStore::beginPaint")

	if ptr.Pointer() != nil {
		C.QBackingStore_BeginPaint(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QBackingStore) EndPaint() {
	defer qt.Recovering("QBackingStore::endPaint")

	if ptr.Pointer() != nil {
		C.QBackingStore_EndPaint(ptr.Pointer())
	}
}

func (ptr *QBackingStore) Flush(region QRegion_ITF, win QWindow_ITF, offset core.QPoint_ITF) {
	defer qt.Recovering("QBackingStore::flush")

	if ptr.Pointer() != nil {
		C.QBackingStore_Flush(ptr.Pointer(), PointerFromQRegion(region), PointerFromQWindow(win), core.PointerFromQPoint(offset))
	}
}

func (ptr *QBackingStore) HasStaticContents() bool {
	defer qt.Recovering("QBackingStore::hasStaticContents")

	if ptr.Pointer() != nil {
		return C.QBackingStore_HasStaticContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBackingStore) Resize(size core.QSize_ITF) {
	defer qt.Recovering("QBackingStore::resize")

	if ptr.Pointer() != nil {
		C.QBackingStore_Resize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QBackingStore) Scroll(area QRegion_ITF, dx int, dy int) bool {
	defer qt.Recovering("QBackingStore::scroll")

	if ptr.Pointer() != nil {
		return C.QBackingStore_Scroll(ptr.Pointer(), PointerFromQRegion(area), C.int(dx), C.int(dy)) != 0
	}
	return false
}

func (ptr *QBackingStore) SetStaticContents(region QRegion_ITF) {
	defer qt.Recovering("QBackingStore::setStaticContents")

	if ptr.Pointer() != nil {
		C.QBackingStore_SetStaticContents(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QBackingStore) Size() *core.QSize {
	defer qt.Recovering("QBackingStore::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QBackingStore_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBackingStore) StaticContents() *QRegion {
	defer qt.Recovering("QBackingStore::staticContents")

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QBackingStore_StaticContents(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBackingStore) Window() *QWindow {
	defer qt.Recovering("QBackingStore::window")

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QBackingStore_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBackingStore) DestroyQBackingStore() {
	defer qt.Recovering("QBackingStore::~QBackingStore")

	if ptr.Pointer() != nil {
		C.QBackingStore_DestroyQBackingStore(ptr.Pointer())
	}
}
