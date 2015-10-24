package dbus

//#include "qdbusobjectpath.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusObjectPath struct {
	ptr unsafe.Pointer
}

type QDBusObjectPathITF interface {
	QDBusObjectPathPTR() *QDBusObjectPath
}

func (p *QDBusObjectPath) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusObjectPath) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPathITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPathPTR().Pointer()
	}
	return nil
}

func QDBusObjectPathFromPointer(ptr unsafe.Pointer) *QDBusObjectPath {
	var n = new(QDBusObjectPath)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusObjectPath) QDBusObjectPathPTR() *QDBusObjectPath {
	return ptr
}

func NewQDBusObjectPath() *QDBusObjectPath {
	return QDBusObjectPathFromPointer(unsafe.Pointer(C.QDBusObjectPath_NewQDBusObjectPath()))
}

func NewQDBusObjectPath3(path core.QLatin1StringITF) *QDBusObjectPath {
	return QDBusObjectPathFromPointer(unsafe.Pointer(C.QDBusObjectPath_NewQDBusObjectPath3(C.QtObjectPtr(core.PointerFromQLatin1String(path)))))
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	return QDBusObjectPathFromPointer(unsafe.Pointer(C.QDBusObjectPath_NewQDBusObjectPath4(C.CString(path))))
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	return QDBusObjectPathFromPointer(unsafe.Pointer(C.QDBusObjectPath_NewQDBusObjectPath2(C.CString(path))))
}

func (ptr *QDBusObjectPath) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusObjectPath_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	if ptr.Pointer() != nil {
		C.QDBusObjectPath_SetPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))
	}
}
