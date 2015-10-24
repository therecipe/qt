package gui

//#include "qbackingstore.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBackingStore struct {
	ptr unsafe.Pointer
}

type QBackingStoreITF interface {
	QBackingStorePTR() *QBackingStore
}

func (p *QBackingStore) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBackingStore) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBackingStore(ptr QBackingStoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBackingStorePTR().Pointer()
	}
	return nil
}

func QBackingStoreFromPointer(ptr unsafe.Pointer) *QBackingStore {
	var n = new(QBackingStore)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBackingStore) QBackingStorePTR() *QBackingStore {
	return ptr
}

func (ptr *QBackingStore) PaintDevice() *QPaintDevice {
	if ptr.Pointer() != nil {
		return QPaintDeviceFromPointer(unsafe.Pointer(C.QBackingStore_PaintDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQBackingStore(window QWindowITF) *QBackingStore {
	return QBackingStoreFromPointer(unsafe.Pointer(C.QBackingStore_NewQBackingStore(C.QtObjectPtr(PointerFromQWindow(window)))))
}

func (ptr *QBackingStore) BeginPaint(region QRegionITF) {
	if ptr.Pointer() != nil {
		C.QBackingStore_BeginPaint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)))
	}
}

func (ptr *QBackingStore) EndPaint() {
	if ptr.Pointer() != nil {
		C.QBackingStore_EndPaint(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBackingStore) Flush(region QRegionITF, win QWindowITF, offset core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QBackingStore_Flush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)), C.QtObjectPtr(PointerFromQWindow(win)), C.QtObjectPtr(core.PointerFromQPoint(offset)))
	}
}

func (ptr *QBackingStore) HasStaticContents() bool {
	if ptr.Pointer() != nil {
		return C.QBackingStore_HasStaticContents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBackingStore) Resize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QBackingStore_Resize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QBackingStore) Scroll(area QRegionITF, dx int, dy int) bool {
	if ptr.Pointer() != nil {
		return C.QBackingStore_Scroll(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(area)), C.int(dx), C.int(dy)) != 0
	}
	return false
}

func (ptr *QBackingStore) SetStaticContents(region QRegionITF) {
	if ptr.Pointer() != nil {
		C.QBackingStore_SetStaticContents(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)))
	}
}

func (ptr *QBackingStore) Window() *QWindow {
	if ptr.Pointer() != nil {
		return QWindowFromPointer(unsafe.Pointer(C.QBackingStore_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QBackingStore) DestroyQBackingStore() {
	if ptr.Pointer() != nil {
		C.QBackingStore_DestroyQBackingStore(C.QtObjectPtr(ptr.Pointer()))
	}
}
