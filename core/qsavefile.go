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

type QSaveFile_ITF interface {
	QFileDevice_ITF
	QSaveFile_PTR() *QSaveFile
}

func PointerFromQSaveFile(ptr QSaveFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSaveFile_PTR().Pointer()
	}
	return nil
}

func NewQSaveFileFromPointer(ptr unsafe.Pointer) *QSaveFile {
	var n = new(QSaveFile)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QSaveFile_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSaveFile) QSaveFile_PTR() *QSaveFile {
	return ptr
}

func NewQSaveFile2(parent QObject_ITF) *QSaveFile {
	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile2(PointerFromQObject(parent)))
}

func NewQSaveFile(name string) *QSaveFile {
	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile(C.CString(name)))
}

func NewQSaveFile3(name string, parent QObject_ITF) *QSaveFile {
	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile3(C.CString(name), PointerFromQObject(parent)))
}

func (ptr *QSaveFile) CancelWriting() {
	if ptr.Pointer() != nil {
		C.QSaveFile_CancelWriting(ptr.Pointer())
	}
}

func (ptr *QSaveFile) Commit() bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) DirectWriteFallback() bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_DirectWriteFallback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSaveFile_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSaveFile) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QSaveFile_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSaveFile) SetDirectWriteFallback(enabled bool) {
	if ptr.Pointer() != nil {
		C.QSaveFile_SetDirectWriteFallback(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSaveFile) SetFileName(name string) {
	if ptr.Pointer() != nil {
		C.QSaveFile_SetFileName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSaveFile) DestroyQSaveFile() {
	if ptr.Pointer() != nil {
		C.QSaveFile_DestroyQSaveFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
