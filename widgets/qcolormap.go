package widgets

//#include "qcolormap.h"
import "C"
import (
	"unsafe"
)

type QColormap struct {
	ptr unsafe.Pointer
}

type QColormapITF interface {
	QColormapPTR() *QColormap
}

func (p *QColormap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QColormap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQColormap(ptr QColormapITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColormapPTR().Pointer()
	}
	return nil
}

func QColormapFromPointer(ptr unsafe.Pointer) *QColormap {
	var n = new(QColormap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QColormap) QColormapPTR() *QColormap {
	return ptr
}

//QColormap::Mode
type QColormap__Mode int

var (
	QColormap__Direct  = QColormap__Mode(0)
	QColormap__Indexed = QColormap__Mode(1)
	QColormap__Gray    = QColormap__Mode(2)
)

func NewQColormap(colormap QColormapITF) *QColormap {
	return QColormapFromPointer(unsafe.Pointer(C.QColormap_NewQColormap(C.QtObjectPtr(PointerFromQColormap(colormap)))))
}

func (ptr *QColormap) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QColormap_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColormap) Mode() QColormap__Mode {
	if ptr.Pointer() != nil {
		return QColormap__Mode(C.QColormap_Mode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColormap) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QColormap_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QColormap) DestroyQColormap() {
	if ptr.Pointer() != nil {
		C.QColormap_DestroyQColormap(C.QtObjectPtr(ptr.Pointer()))
	}
}
