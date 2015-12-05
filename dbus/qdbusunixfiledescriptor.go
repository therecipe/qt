package dbus

//#include "dbus.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDBusUnixFileDescriptor struct {
	ptr unsafe.Pointer
}

type QDBusUnixFileDescriptor_ITF interface {
	QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor
}

func (p *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusUnixFileDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) *QDBusUnixFileDescriptor {
	var n = new(QDBusUnixFileDescriptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusUnixFileDescriptor) QDBusUnixFileDescriptor_PTR() *QDBusUnixFileDescriptor {
	return ptr
}

func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")
		}
	}()

	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")
		}
	}()

	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::QDBusUnixFileDescriptor")
		}
	}()

	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(fileDescriptor)))
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::fileDescriptor")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer()))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::isSupported")
		}
	}()

	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::setFileDescriptor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(fileDescriptor))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusUnixFileDescriptor::~QDBusUnixFileDescriptor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(ptr.Pointer())
	}
}
