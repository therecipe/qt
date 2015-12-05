package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTextItem struct {
	ptr unsafe.Pointer
}

type QTextItem_ITF interface {
	QTextItem_PTR() *QTextItem
}

func (p *QTextItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextItem(ptr QTextItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextItem_PTR().Pointer()
	}
	return nil
}

func NewQTextItemFromPointer(ptr unsafe.Pointer) *QTextItem {
	var n = new(QTextItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextItem) QTextItem_PTR() *QTextItem {
	return ptr
}

//QTextItem::RenderFlag
type QTextItem__RenderFlag int64

const (
	QTextItem__RightToLeft = QTextItem__RenderFlag(0x1)
	QTextItem__Overline    = QTextItem__RenderFlag(0x10)
	QTextItem__Underline   = QTextItem__RenderFlag(0x20)
	QTextItem__StrikeOut   = QTextItem__RenderFlag(0x40)
	QTextItem__Dummy       = QTextItem__RenderFlag(0xffffffff)
)

func (ptr *QTextItem) Ascent() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextItem::ascent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextItem_Ascent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextItem) Descent() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextItem::descent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextItem_Descent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextItem) RenderFlags() QTextItem__RenderFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextItem::renderFlags")
		}
	}()

	if ptr.Pointer() != nil {
		return QTextItem__RenderFlag(C.QTextItem_RenderFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextItem) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextItem_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTextItem) Width() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextItem::width")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextItem_Width(ptr.Pointer()))
	}
	return 0
}
