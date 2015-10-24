package dbus

//#include "qdbusunixfiledescriptor.h"
import "C"
import (
	"unsafe"
)

type QDBusUnixFileDescriptor struct {
	ptr unsafe.Pointer
}

type QDBusUnixFileDescriptorITF interface {
	QDBusUnixFileDescriptorPTR() *QDBusUnixFileDescriptor
}

func (p *QDBusUnixFileDescriptor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusUnixFileDescriptor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusUnixFileDescriptor(ptr QDBusUnixFileDescriptorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusUnixFileDescriptorPTR().Pointer()
	}
	return nil
}

func QDBusUnixFileDescriptorFromPointer(ptr unsafe.Pointer) *QDBusUnixFileDescriptor {
	var n = new(QDBusUnixFileDescriptor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusUnixFileDescriptor) QDBusUnixFileDescriptorPTR() *QDBusUnixFileDescriptor {
	return ptr
}

func NewQDBusUnixFileDescriptor() *QDBusUnixFileDescriptor {
	return QDBusUnixFileDescriptorFromPointer(unsafe.Pointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor()))
}

func NewQDBusUnixFileDescriptor3(other QDBusUnixFileDescriptorITF) *QDBusUnixFileDescriptor {
	return QDBusUnixFileDescriptorFromPointer(unsafe.Pointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(C.QtObjectPtr(PointerFromQDBusUnixFileDescriptor(other)))))
}

func NewQDBusUnixFileDescriptor2(fileDescriptor int) *QDBusUnixFileDescriptor {
	return QDBusUnixFileDescriptorFromPointer(unsafe.Pointer(C.QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(C.int(fileDescriptor))))
}

func (ptr *QDBusUnixFileDescriptor) FileDescriptor() int {
	if ptr.Pointer() != nil {
		return int(C.QDBusUnixFileDescriptor_FileDescriptor(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QDBusUnixFileDescriptor_IsSupported() bool {
	return C.QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported() != 0
}

func (ptr *QDBusUnixFileDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusUnixFileDescriptor_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusUnixFileDescriptor) SetFileDescriptor(fileDescriptor int) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_SetFileDescriptor(C.QtObjectPtr(ptr.Pointer()), C.int(fileDescriptor))
	}
}

func (ptr *QDBusUnixFileDescriptor) Swap(other QDBusUnixFileDescriptorITF) {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusUnixFileDescriptor(other)))
	}
}

func (ptr *QDBusUnixFileDescriptor) DestroyQDBusUnixFileDescriptor() {
	if ptr.Pointer() != nil {
		C.QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(C.QtObjectPtr(ptr.Pointer()))
	}
}
