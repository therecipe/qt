package core

//#include "qtemporaryfile.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTemporaryFile struct {
	QFile
}

type QTemporaryFile_ITF interface {
	QFile_ITF
	QTemporaryFile_PTR() *QTemporaryFile
}

func PointerFromQTemporaryFile(ptr QTemporaryFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTemporaryFile_PTR().Pointer()
	}
	return nil
}

func NewQTemporaryFileFromPointer(ptr unsafe.Pointer) *QTemporaryFile {
	var n = new(QTemporaryFile)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTemporaryFile_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTemporaryFile) QTemporaryFile_PTR() *QTemporaryFile {
	return ptr
}

func NewQTemporaryFile() *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile())
}

func NewQTemporaryFile3(parent QObject_ITF) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile3(PointerFromQObject(parent)))
}

func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile2(C.CString(templateName)))
}

func NewQTemporaryFile4(templateName string, parent QObject_ITF) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile4(C.CString(templateName), PointerFromQObject(parent)))
}

func (ptr *QTemporaryFile) AutoRemove() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryFile_AutoRemove(ptr.Pointer()) != 0
	}
	return false
}

func QTemporaryFile_CreateNativeFile(file QFile_ITF) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile(PointerFromQFile(file)))
}

func QTemporaryFile_CreateNativeFile2(fileName string) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile2(C.CString(fileName)))
}

func (ptr *QTemporaryFile) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTemporaryFile) FileTemplate() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileTemplate(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTemporaryFile) Open() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryFile_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTemporaryFile) SetAutoRemove(b bool) {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetAutoRemove(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTemporaryFile) SetFileTemplate(name string) {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetFileTemplate(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTemporaryFile) DestroyQTemporaryFile() {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_DestroyQTemporaryFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
