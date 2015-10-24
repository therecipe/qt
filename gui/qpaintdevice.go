package gui

//#include "qpaintdevice.h"
import "C"
import (
	"unsafe"
)

type QPaintDevice struct {
	ptr unsafe.Pointer
}

type QPaintDeviceITF interface {
	QPaintDevicePTR() *QPaintDevice
}

func (p *QPaintDevice) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintDevice) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintDevice(ptr QPaintDeviceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevicePTR().Pointer()
	}
	return nil
}

func QPaintDeviceFromPointer(ptr unsafe.Pointer) *QPaintDevice {
	var n = new(QPaintDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintDevice) QPaintDevicePTR() *QPaintDevice {
	return ptr
}

//QPaintDevice::PaintDeviceMetric
type QPaintDevice__PaintDeviceMetric int

var (
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
		C.QPaintDevice_DestroyQPaintDevice(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPaintDevice) ColorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_ColorCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) DevicePixelRatio() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_DevicePixelRatio(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) HeightMM() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_HeightMM(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) LogicalDpiX() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_LogicalDpiX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) LogicalDpiY() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_LogicalDpiY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return QPaintEngineFromPointer(unsafe.Pointer(C.QPaintDevice_PaintEngine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPaintDevice) PaintingActive() bool {
	if ptr.Pointer() != nil {
		return C.QPaintDevice_PaintingActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintDevice) PhysicalDpiX() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_PhysicalDpiX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) PhysicalDpiY() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_PhysicalDpiY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintDevice) WidthMM() int {
	if ptr.Pointer() != nil {
		return int(C.QPaintDevice_WidthMM(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
