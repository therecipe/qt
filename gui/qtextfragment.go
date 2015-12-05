package gui

//#include "gui.h"
import "C"
import (
	"log"
	"unsafe"
)

type QTextFragment struct {
	ptr unsafe.Pointer
}

type QTextFragment_ITF interface {
	QTextFragment_PTR() *QTextFragment
}

func (p *QTextFragment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextFragment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextFragment(ptr QTextFragment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFragment_PTR().Pointer()
	}
	return nil
}

func NewQTextFragmentFromPointer(ptr unsafe.Pointer) *QTextFragment {
	var n = new(QTextFragment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextFragment) QTextFragment_PTR() *QTextFragment {
	return ptr
}

func NewQTextFragment() *QTextFragment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::QTextFragment")
		}
	}()

	return NewQTextFragmentFromPointer(C.QTextFragment_NewQTextFragment())
}

func NewQTextFragment3(other QTextFragment_ITF) *QTextFragment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::QTextFragment")
		}
	}()

	return NewQTextFragmentFromPointer(C.QTextFragment_NewQTextFragment3(PointerFromQTextFragment(other)))
}

func (ptr *QTextFragment) CharFormatIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::charFormatIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFragment_CharFormatIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFragment) Contains(position int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFragment_Contains(ptr.Pointer(), C.int(position)) != 0
	}
	return false
}

func (ptr *QTextFragment) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFragment_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFragment) Length() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::length")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFragment_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFragment) Position() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::position")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFragment_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFragment) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFragment::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextFragment_Text(ptr.Pointer()))
	}
	return ""
}
