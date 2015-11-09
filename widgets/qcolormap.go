package widgets

//#include "qcolormap.h"
import "C"
import (
	"unsafe"
)

type QColormap struct {
	ptr unsafe.Pointer
}

type QColormap_ITF interface {
	QColormap_PTR() *QColormap
}

func (p *QColormap) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QColormap) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQColormap(ptr QColormap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColormap_PTR().Pointer()
	}
	return nil
}

func NewQColormapFromPointer(ptr unsafe.Pointer) *QColormap {
	var n = new(QColormap)
	n.SetPointer(ptr)
	return n
}

func (ptr *QColormap) QColormap_PTR() *QColormap {
	return ptr
}

//QColormap::Mode
type QColormap__Mode int64

const (
	QColormap__Direct  = QColormap__Mode(0)
	QColormap__Indexed = QColormap__Mode(1)
	QColormap__Gray    = QColormap__Mode(2)
)

func NewQColormap(colormap QColormap_ITF) *QColormap {
	return NewQColormapFromPointer(C.QColormap_NewQColormap(PointerFromQColormap(colormap)))
}

func (ptr *QColormap) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QColormap_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) Mode() QColormap__Mode {
	if ptr.Pointer() != nil {
		return QColormap__Mode(C.QColormap_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QColormap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) DestroyQColormap() {
	if ptr.Pointer() != nil {
		C.QColormap_DestroyQColormap(ptr.Pointer())
	}
}
