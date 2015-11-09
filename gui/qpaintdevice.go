package gui

//#include "qpaintdevice.h"
import "C"
import (
	"unsafe"
)

type QPaintDevice struct {
	ptr unsafe.Pointer
}

type QPaintDevice_ITF interface {
	QPaintDevice_PTR() *QPaintDevice
}

func (p *QPaintDevice) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintDevice) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintDevice(ptr QPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPaintDeviceFromPointer(ptr unsafe.Pointer) *QPaintDevice {
	var n = new(QPaintDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintDevice) QPaintDevice_PTR() *QPaintDevice {
	return ptr
}

//QPaintDevice::PaintDeviceMetric
type QPaintDevice__PaintDeviceMetric int64

const (
	QPaintDevice__PdmWidth            = QPaintDevice__PaintDeviceMetric(1)
	QPaintDevice__PdmHeight           = QPaintDevice__PaintDeviceMetric(2)
	QPaintDevice__PdmWidthMM          = QPaintDevice__PaintDeviceMetric(3)
	QPaintDevice__PdmHeightMM         = QPaintDevice__PaintDeviceMetric(4)
	QPaintDevice__PdmNumColors        = QPaintDevice__PaintDeviceMetric(5)
	QPaintDevice__PdmDepth            = QPaintDevice__PaintDeviceMetric(6)
	QPaintDevice__PdmDpiX             = QPaintDevice__PaintDeviceMetric(7)
	QPaintDevice__PdmDpiY             = QPaintDevice__PaintDeviceMetric(8)
	QPaintDevice__PdmPhysicalDpiX     = QPaintDevice__PaintDeviceMetric(9)
	QPaintDevice__PdmPhysicalDpiY     = QPaintDevice__PaintDeviceMetric(10)
	QPaintDevice__PdmDevicePixelRatio = QPaintDevice__PaintDeviceMetric(11)
)

func (ptr *QPaintDevice) DestroyQPaintDevice() {
	if ptr.Pointer() != nil {
		C.QPaintDevice_DestroyQPaintDevice(ptr.Pointer())
	}
}

func (ptr *QPaintDevice) ColorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_ColorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) DevicePixelRatio() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) HeightMM() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_HeightMM(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) LogicalDpiX() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_LogicalDpiX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) LogicalDpiY() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_LogicalDpiY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPaintDevice_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintDevice) PaintingActive() bool {
	if ptr.Pointer() != nil {
		return C.QPaintDevice_PaintingActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintDevice) PhysicalDpiX() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_PhysicalDpiX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) PhysicalDpiY() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_PhysicalDpiY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintDevice) WidthMM() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_WidthMM(ptr.Pointer()))
	}
	return 0
}
