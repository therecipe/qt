package dbus

//#include "qdbuscontext.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDBusContext struct {
	ptr unsafe.Pointer
}

type QDBusContextITF interface {
	QDBusContextPTR() *QDBusContext
}

func (p *QDBusContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusContext(ptr QDBusContextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContextPTR().Pointer()
	}
	return nil
}

func QDBusContextFromPointer(ptr unsafe.Pointer) *QDBusContext {
	var n = new(QDBusContext)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusContext) QDBusContextPTR() *QDBusContext {
	return ptr
}

func NewQDBusContext() *QDBusContext {
	return QDBusContextFromPointer(unsafe.Pointer(C.QDBusContext_NewQDBusContext()))
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	if ptr.Pointer() != nil {
		return C.QDBusContext_CalledFromDBus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	if ptr.Pointer() != nil {
		return C.QDBusContext_IsDelayedReply(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply2(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(msg))
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(msg))
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusContext) DestroyQDBusContext() {
	if ptr.Pointer() != nil {
		C.QDBusContext_DestroyQDBusContext(C.QtObjectPtr(ptr.Pointer()))
	}
}
