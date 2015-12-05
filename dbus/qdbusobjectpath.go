package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDBusObjectPath struct {
	ptr unsafe.Pointer
}

type QDBusObjectPath_ITF interface {
	QDBusObjectPath_PTR() *QDBusObjectPath
}

func (p *QDBusObjectPath) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusObjectPath) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusObjectPath(ptr QDBusObjectPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusObjectPath_PTR().Pointer()
	}
	return nil
}

func NewQDBusObjectPathFromPointer(ptr unsafe.Pointer) *QDBusObjectPath {
	var n = new(QDBusObjectPath)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusObjectPath) QDBusObjectPath_PTR() *QDBusObjectPath {
	return ptr
}

func NewQDBusObjectPath() *QDBusObjectPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::QDBusObjectPath")
		}
	}()

	return NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath())
}

func NewQDBusObjectPath3(path core.QLatin1String_ITF) *QDBusObjectPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::QDBusObjectPath")
		}
	}()

	return NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath3(core.PointerFromQLatin1String(path)))
}

func NewQDBusObjectPath4(path string) *QDBusObjectPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::QDBusObjectPath")
		}
	}()

	return NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath4(C.CString(path)))
}

func NewQDBusObjectPath2(path string) *QDBusObjectPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::QDBusObjectPath")
		}
	}()

	return NewQDBusObjectPathFromPointer(C.QDBusObjectPath_NewQDBusObjectPath2(C.CString(path)))
}

func (ptr *QDBusObjectPath) Path() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::path")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusObjectPath_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusObjectPath) SetPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusObjectPath::setPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusObjectPath_SetPath(ptr.Pointer(), C.CString(path))
	}
}
