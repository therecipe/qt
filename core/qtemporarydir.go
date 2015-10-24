package core

//#include "qtemporarydir.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTemporaryDir struct {
	ptr unsafe.Pointer
}

type QTemporaryDirITF interface {
	QTemporaryDirPTR() *QTemporaryDir
}

func (p *QTemporaryDir) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTemporaryDir) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTemporaryDir(ptr QTemporaryDirITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTemporaryDirPTR().Pointer()
	}
	return nil
}

func QTemporaryDirFromPointer(ptr unsafe.Pointer) *QTemporaryDir {
	var n = new(QTemporaryDir)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTemporaryDir) QTemporaryDirPTR() *QTemporaryDir {
	return ptr
}

func NewQTemporaryDir() *QTemporaryDir {
	return QTemporaryDirFromPointer(unsafe.Pointer(C.QTemporaryDir_NewQTemporaryDir()))
}

func NewQTemporaryDir2(templatePath string) *QTemporaryDir {
	return QTemporaryDirFromPointer(unsafe.Pointer(C.QTemporaryDir_NewQTemporaryDir2(C.CString(templatePath))))
}

func (ptr *QTemporaryDir) AutoRemove() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryDir_AutoRemove(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTemporaryDir) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryDir_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTemporaryDir) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryDir_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTemporaryDir) SetAutoRemove(b bool) {
	if ptr.Pointer() != nil {
		C.QTemporaryDir_SetAutoRemove(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTemporaryDir) DestroyQTemporaryDir() {
	if ptr.Pointer() != nil {
		C.QTemporaryDir_DestroyQTemporaryDir(C.QtObjectPtr(ptr.Pointer()))
	}
}
