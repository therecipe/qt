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

type QTemporaryFileITF interface {
	QFileITF
	QTemporaryFilePTR() *QTemporaryFile
}

func PointerFromQTemporaryFile(ptr QTemporaryFileITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTemporaryFilePTR().Pointer()
	}
	return nil
}

func QTemporaryFileFromPointer(ptr unsafe.Pointer) *QTemporaryFile {
	var n = new(QTemporaryFile)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTemporaryFile_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTemporaryFile) QTemporaryFilePTR() *QTemporaryFile {
	return ptr
}

func NewQTemporaryFile() *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_NewQTemporaryFile()))
}

func NewQTemporaryFile3(parent QObjectITF) *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_NewQTemporaryFile3(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_NewQTemporaryFile2(C.CString(templateName))))
}

func NewQTemporaryFile4(templateName string, parent QObjectITF) *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_NewQTemporaryFile4(C.CString(templateName), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QTemporaryFile) AutoRemove() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryFile_AutoRemove(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QTemporaryFile_CreateNativeFile(file QFileITF) *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile(C.QtObjectPtr(PointerFromQFile(file)))))
}

func QTemporaryFile_CreateNativeFile2(fileName string) *QTemporaryFile {
	return QTemporaryFileFromPointer(unsafe.Pointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile2(C.CString(fileName))))
}

func (ptr *QTemporaryFile) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTemporaryFile) FileTemplate() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileTemplate(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTemporaryFile) Open() bool {
	if ptr.Pointer() != nil {
		return C.QTemporaryFile_Open(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTemporaryFile) SetAutoRemove(b bool) {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetAutoRemove(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTemporaryFile) SetFileTemplate(name string) {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetFileTemplate(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QTemporaryFile) DestroyQTemporaryFile() {
	if ptr.Pointer() != nil {
		C.QTemporaryFile_DestroyQTemporaryFile(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
