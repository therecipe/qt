package core

//#include "core.h"
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
	for len(n.ObjectName()) < len("QSaveFile_") {
		n.SetObjectName("QSaveFile_" + qt.Identifier())
	}
	return n
}

func (ptr *QSaveFile) QSaveFile_PTR() *QSaveFile {
	return ptr
}

func NewQSaveFile2(parent QObject_ITF) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile2(PointerFromQObject(parent)))
}

func NewQSaveFile(name string) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile(C.CString(name)))
}

func NewQSaveFile3(name string, parent QObject_ITF) *QSaveFile {
	defer qt.Recovering("QSaveFile::QSaveFile")

	return NewQSaveFileFromPointer(C.QSaveFile_NewQSaveFile3(C.CString(name), PointerFromQObject(parent)))
}

func (ptr *QSaveFile) CancelWriting() {
	defer qt.Recovering("QSaveFile::cancelWriting")

	if ptr.Pointer() != nil {
		C.QSaveFile_CancelWriting(ptr.Pointer())
	}
}

func (ptr *QSaveFile) Commit() bool {
	defer qt.Recovering("QSaveFile::commit")

	if ptr.Pointer() != nil {
		return C.QSaveFile_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) DirectWriteFallback() bool {
	defer qt.Recovering("QSaveFile::directWriteFallback")

	if ptr.Pointer() != nil {
		return C.QSaveFile_DirectWriteFallback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSaveFile) FileName() string {
	defer qt.Recovering("QSaveFile::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSaveFile_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSaveFile) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QSaveFile::open")

	if ptr.Pointer() != nil {
		return C.QSaveFile_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSaveFile) SetDirectWriteFallback(enabled bool) {
	defer qt.Recovering("QSaveFile::setDirectWriteFallback")

	if ptr.Pointer() != nil {
		C.QSaveFile_SetDirectWriteFallback(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSaveFile) SetFileName(name string) {
	defer qt.Recovering("QSaveFile::setFileName")

	if ptr.Pointer() != nil {
		C.QSaveFile_SetFileName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSaveFile) DestroyQSaveFile() {
	defer qt.Recovering("QSaveFile::~QSaveFile")

	if ptr.Pointer() != nil {
		C.QSaveFile_DestroyQSaveFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
