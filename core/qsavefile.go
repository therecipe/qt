package core

//#include "qsavefile.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSaveFile struct {
	QFileDevice
}

type QSaveFileITF interface {
	QFileDeviceITF
	QSaveFilePTR() *QSaveFile
}

func PointerFromQSaveFile(ptr QSaveFileITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSaveFilePTR().Pointer()
	}
	return nil
}

func QSaveFileFromPointer(ptr unsafe.Pointer) *QSaveFile {
	var n = new(QSaveFile)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSaveFile_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSaveFile) QSaveFilePTR() *QSaveFile {
	return ptr
}

func NewQSaveFile2(parent QObjectITF) *QSaveFile {
	return QSaveFileFromPointer(unsafe.Pointer(C.QSaveFile_NewQSaveFile2(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSaveFile(name string) *QSaveFile {
	return QSaveFileFromPointer(unsafe.Pointer(C.QSaveFile_NewQSaveFile(C.CString(name))))
}

func NewQSaveFile3(name string, parent QObjectITF) *QSaveFile {
	return QSaveFileFromPointer(unsafe.Pointer(C.QSaveFile_NewQSaveFile3(C.CString(name), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QSaveFile) CancelWriting() {
	if ptr.Pointer() != nil {
		C.QSaveFile_CancelWriting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSaveFile) Commit() bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_Commit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSaveFile) DirectWriteFallback() bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_DirectWriteFallback(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSaveFile) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSaveFile_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSaveFile) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_Open(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSaveFile) SetDirectWriteFallback(enabled bool) {
	if ptr.Pointer() != nil {
		C.QSaveFile_SetDirectWriteFallback(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSaveFile) SetFileName(name string) {
	if ptr.Pointer() != nil {
		C.QSaveFile_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QSaveFile) DestroyQSaveFile() {
	if ptr.Pointer() != nil {
		C.QSaveFile_DestroyQSaveFile(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
