package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QColormap::QColormap")

	return NewQColormapFromPointer(C.QColormap_NewQColormap(PointerFromQColormap(colormap)))
}

func (ptr *QColormap) Depth() int {
	defer qt.Recovering("QColormap::depth")

	if ptr.Pointer() != nil {
		return int(C.QColormap_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) Mode() QColormap__Mode {
	defer qt.Recovering("QColormap::mode")

	if ptr.Pointer() != nil {
		return QColormap__Mode(C.QColormap_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) Size() int {
	defer qt.Recovering("QColormap::size")

	if ptr.Pointer() != nil {
		return int(C.QColormap_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColormap) DestroyQColormap() {
	defer qt.Recovering("QColormap::~QColormap")

	if ptr.Pointer() != nil {
		C.QColormap_DestroyQColormap(ptr.Pointer())
	}
}
