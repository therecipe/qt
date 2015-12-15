package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFile struct {
	QFileDevice
}

type QFile_ITF interface {
	QFileDevice_ITF
	QFile_PTR() *QFile
}

func PointerFromQFile(ptr QFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFile_PTR().Pointer()
	}
	return nil
}

func NewQFileFromPointer(ptr unsafe.Pointer) *QFile {
	var n = new(QFile)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFile_") {
		n.SetObjectName("QFile_" + qt.Identifier())
	}
	return n
}

func (ptr *QFile) QFile_PTR() *QFile {
	return ptr
}

func NewQFile3(parent QObject_ITF) *QFile {
	defer qt.Recovering("QFile::QFile")

	return NewQFileFromPointer(C.QFile_NewQFile3(PointerFromQObject(parent)))
}

func NewQFile(name string) *QFile {
	defer qt.Recovering("QFile::QFile")

	return NewQFileFromPointer(C.QFile_NewQFile(C.CString(name)))
}

func NewQFile4(name string, parent QObject_ITF) *QFile {
	defer qt.Recovering("QFile::QFile")

	return NewQFileFromPointer(C.QFile_NewQFile4(C.CString(name), PointerFromQObject(parent)))
}

func QFile_Copy2(fileName string, newName string) bool {
	defer qt.Recovering("QFile::copy")

	return C.QFile_QFile_Copy2(C.CString(fileName), C.CString(newName)) != 0
}

func (ptr *QFile) Copy(newName string) bool {
	defer qt.Recovering("QFile::copy")

	if ptr.Pointer() != nil {
		return C.QFile_Copy(ptr.Pointer(), C.CString(newName)) != 0
	}
	return false
}

func QFile_DecodeName(localFileName QByteArray_ITF) string {
	defer qt.Recovering("QFile::decodeName")

	return C.GoString(C.QFile_QFile_DecodeName(PointerFromQByteArray(localFileName)))
}

func QFile_DecodeName2(localFileName string) string {
	defer qt.Recovering("QFile::decodeName")

	return C.GoString(C.QFile_QFile_DecodeName2(C.CString(localFileName)))
}

func QFile_EncodeName(fileName string) *QByteArray {
	defer qt.Recovering("QFile::encodeName")

	return NewQByteArrayFromPointer(C.QFile_QFile_EncodeName(C.CString(fileName)))
}

func QFile_Exists(fileName string) bool {
	defer qt.Recovering("QFile::exists")

	return C.QFile_QFile_Exists(C.CString(fileName)) != 0
}

func (ptr *QFile) Exists2() bool {
	defer qt.Recovering("QFile::exists")

	if ptr.Pointer() != nil {
		return C.QFile_Exists2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFile) FileName() string {
	defer qt.Recovering("QFile::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFile_FileName(ptr.Pointer()))
	}
	return ""
}

func QFile_Link2(fileName string, linkName string) bool {
	defer qt.Recovering("QFile::link")

	return C.QFile_QFile_Link2(C.CString(fileName), C.CString(linkName)) != 0
}

func (ptr *QFile) Link(linkName string) bool {
	defer qt.Recovering("QFile::link")

	if ptr.Pointer() != nil {
		return C.QFile_Link(ptr.Pointer(), C.CString(linkName)) != 0
	}
	return false
}

func (ptr *QFile) Open(mode QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QFile::open")

	if ptr.Pointer() != nil {
		return C.QFile_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QFile) Open3(fd int, mode QIODevice__OpenModeFlag, handleFlags QFileDevice__FileHandleFlag) bool {
	defer qt.Recovering("QFile::open")

	if ptr.Pointer() != nil {
		return C.QFile_Open3(ptr.Pointer(), C.int(fd), C.int(mode), C.int(handleFlags)) != 0
	}
	return false
}

func QFile_Permissions2(fileName string) QFileDevice__Permission {
	defer qt.Recovering("QFile::permissions")

	return QFileDevice__Permission(C.QFile_QFile_Permissions2(C.CString(fileName)))
}

func (ptr *QFile) Permissions() QFileDevice__Permission {
	defer qt.Recovering("QFile::permissions")

	if ptr.Pointer() != nil {
		return QFileDevice__Permission(C.QFile_Permissions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFile) Rename(newName string) bool {
	defer qt.Recovering("QFile::rename")

	if ptr.Pointer() != nil {
		return C.QFile_Rename(ptr.Pointer(), C.CString(newName)) != 0
	}
	return false
}

func QFile_Rename2(oldName string, newName string) bool {
	defer qt.Recovering("QFile::rename")

	return C.QFile_QFile_Rename2(C.CString(oldName), C.CString(newName)) != 0
}

func (ptr *QFile) SetFileName(name string) {
	defer qt.Recovering("QFile::setFileName")

	if ptr.Pointer() != nil {
		C.QFile_SetFileName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QFile) SetPermissions(permissions QFileDevice__Permission) bool {
	defer qt.Recovering("QFile::setPermissions")

	if ptr.Pointer() != nil {
		return C.QFile_SetPermissions(ptr.Pointer(), C.int(permissions)) != 0
	}
	return false
}

func QFile_SetPermissions2(fileName string, permissions QFileDevice__Permission) bool {
	defer qt.Recovering("QFile::setPermissions")

	return C.QFile_QFile_SetPermissions2(C.CString(fileName), C.int(permissions)) != 0
}

func QFile_SymLinkTarget(fileName string) string {
	defer qt.Recovering("QFile::symLinkTarget")

	return C.GoString(C.QFile_QFile_SymLinkTarget(C.CString(fileName)))
}

func (ptr *QFile) SymLinkTarget2() string {
	defer qt.Recovering("QFile::symLinkTarget")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFile_SymLinkTarget2(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFile) DestroyQFile() {
	defer qt.Recovering("QFile::~QFile")

	if ptr.Pointer() != nil {
		C.QFile_DestroyQFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
