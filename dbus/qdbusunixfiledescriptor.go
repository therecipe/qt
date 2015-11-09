package dbus

//#include "qdbusunixfiledescriptor.h"
import "C"
import (
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
	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor())
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptor_ITF) *QDBusUnixFileDescriptor {
	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(PointerFromQDBusUnixFileDescriptor(other)))
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	return NewQDBusUnixFileDescriptorFromPointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(fileDescriptor)))
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	if ptr.Pointer() != nil {
		return int(C.QDBusUnixFileDescriptor_FileDescriptor(ptr.Pointer()))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(ptr.Pointer(), C.int(fileDescriptor))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(ptr.Pointer(), PointerFromQDBusUnixFileDescriptor(other))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(ptr.Pointer())
	}
}
