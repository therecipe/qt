package core

//#include "qfile.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFile struct {
	QFileDevice
}

type QFileITF interface {
	QFileDeviceITF
	QFilePTR() *QFile
}

func PointerFromQFile(ptr QFileITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFilePTR().Pointer()
	}
	return nil
}

func QFileFromPointer(ptr unsafe.Pointer) *QFile {
	var n = new(QFile)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFile_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFile) QFilePTR() *QFile {
	return ptr
}

func NewQFile3(parent QObjectITF) *QFile {
	return QFileFromPointer(unsafe.Pointer(C.QFile_NewQFile3(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQFile(name string) *QFile {
	return QFileFromPointer(unsafe.Pointer(C.QFile_NewQFile(C.CString(name))))
}

func NewQFile4(name string, parent QObjectITF) *QFile {
	return QFileFromPointer(unsafe.Pointer(C.QFile_NewQFile4(C.CString(name), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func QFile_Copy2(fileName string, newName string) bool {
	return C.QFile_QFile_Copy2(C.CString(fileName), C.CString(newName)) != 0
}

func (ptr *QFile) Copy(newName string) bool {
	if ptr.Pointer() != nil {
		return C.QFile_Copy(C.QtObjectPtr(ptr.Pointer()), C.CString(newName)) != 0
	}
	return false
}

func QFile_DecodeName(localFileName QByteArrayITF) string {
	return C.GoString(C.QFile_QFile_DecodeName(C.QtObjectPtr(PointerFromQByteArray(localFileName))))
}

func QFile_DecodeName2(localFileName string) string {
	return C.GoString(C.QFile_QFile_DecodeName2(C.CString(localFileName)))
}

func QFile_Exists(fileName string) bool {
	return C.QFile_QFile_Exists(C.CString(fileName)) != 0
}

func (ptr *QFile) Exists2() bool {
	if ptr.Pointer() != nil {
		return C.QFile_Exists2(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFile) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFile_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QFile_Link2(fileName string, linkName string) bool {
	return C.QFile_QFile_Link2(C.CString(fileName), C.CString(linkName)) != 0
}

func (ptr *QFile) Link(linkName string) bool {
	if ptr.Pointer() != nil {
		return C.QFile_Link(C.QtObjectPtr(ptr.Pointer()), C.CString(linkName)) != 0
	}
	return false
}

func (ptr *QFile) Open(mode QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QFile_Open(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QFile) Open3(fd int, mode QIODevice__OpenModeFlag, handleFlags QFileDevice__FileHandleFlag) bool {
	if ptr.Pointer() != nil {
		return C.QFile_Open3(C.QtObjectPtr(ptr.Pointer()), C.int(fd), C.int(mode), C.int(handleFlags)) != 0
	}
	return false
}

func QFile_Permissions2(fileName string) QFileDevice__Permission {
	return QFileDevice__Permission(C.QFile_QFile_Permissions2(C.CString(fileName)))
}

func (ptr *QFile) Permissions() QFileDevice__Permission {
	if ptr.Pointer() != nil {
		return QFileDevice__Permission(C.QFile_Permissions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFile) Rename(newName string) bool {
	if ptr.Pointer() != nil {
		return C.QFile_Rename(C.QtObjectPtr(ptr.Pointer()), C.CString(newName)) != 0
	}
	return false
}

func QFile_Rename2(oldName string, newName string) bool {
	return C.QFile_QFile_Rename2(C.CString(oldName), C.CString(newName)) != 0
}

func (ptr *QFile) SetFileName(name string) {
	if ptr.Pointer() != nil {
		C.QFile_SetFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QFile) SetPermissions(permissions QFileDevice__Permission) bool {
	if ptr.Pointer() != nil {
		return C.QFile_SetPermissions(C.QtObjectPtr(ptr.Pointer()), C.int(permissions)) != 0
	}
	return false
}

func QFile_SetPermissions2(fileName string, permissions QFileDevice__Permission) bool {
	return C.QFile_QFile_SetPermissions2(C.CString(fileName), C.int(permissions)) != 0
}

func QFile_SymLinkTarget(fileName string) string {
	return C.GoString(C.QFile_QFile_SymLinkTarget(C.CString(fileName)))
}

func (ptr *QFile) SymLinkTarget2() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFile_SymLinkTarget2(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFile) DestroyQFile() {
	if ptr.Pointer() != nil {
		C.QFile_DestroyQFile(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
