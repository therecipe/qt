package core

//#include "qfileinfo.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFileInfo struct {
	ptr unsafe.Pointer
}

type QFileInfoITF interface {
	QFileInfoPTR() *QFileInfo
}

func (p *QFileInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFileInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFileInfo(ptr QFileInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileInfoPTR().Pointer()
	}
	return nil
}

func QFileInfoFromPointer(ptr unsafe.Pointer) *QFileInfo {
	var n = new(QFileInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileInfo) QFileInfoPTR() *QFileInfo {
	return ptr
}

func NewQFileInfo() *QFileInfo {
	return QFileInfoFromPointer(unsafe.Pointer(C.QFileInfo_NewQFileInfo()))
}

func NewQFileInfo5(dir QDirITF, file string) *QFileInfo {
	return QFileInfoFromPointer(unsafe.Pointer(C.QFileInfo_NewQFileInfo5(C.QtObjectPtr(PointerFromQDir(dir)), C.CString(file))))
}

func NewQFileInfo4(file QFileITF) *QFileInfo {
	return QFileInfoFromPointer(unsafe.Pointer(C.QFileInfo_NewQFileInfo4(C.QtObjectPtr(PointerFromQFile(file)))))
}

func NewQFileInfo6(fileinfo QFileInfoITF) *QFileInfo {
	return QFileInfoFromPointer(unsafe.Pointer(C.QFileInfo_NewQFileInfo6(C.QtObjectPtr(PointerFromQFileInfo(fileinfo)))))
}

func NewQFileInfo3(file string) *QFileInfo {
	return QFileInfoFromPointer(unsafe.Pointer(C.QFileInfo_NewQFileInfo3(C.CString(file))))
}

func (ptr *QFileInfo) AbsoluteFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsoluteFilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) AbsolutePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsolutePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) BaseName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BaseName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) BundleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BundleName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) Caching() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_Caching(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) CanonicalFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalFilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) CanonicalPath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalPath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) CompleteBaseName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteBaseName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) CompleteSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteSuffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QFileInfo_Exists2(file string) bool {
	return C.QFileInfo_QFileInfo_Exists2(C.CString(file)) != 0
}

func (ptr *QFileInfo) Exists() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_Exists(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) FilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) Group() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Group(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) IsAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsAbsolute(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsBundle() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsBundle(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsDir() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsDir(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsExecutable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsExecutable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsFile() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsFile(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsNativePath() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsNativePath(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsReadable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRelative() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRelative(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRoot() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRoot(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsSymLink() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsSymLink(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) MakeAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_MakeAbsolute(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileInfo) Owner() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Owner(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) Refresh() {
	if ptr.Pointer() != nil {
		C.QFileInfo_Refresh(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QFileInfo) SetCaching(enable bool) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetCaching(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileInfo) SetFile3(dir QDirITF, file string) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDir(dir)), C.CString(file))
	}
}

func (ptr *QFileInfo) SetFile2(file QFileITF) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFile(file)))
	}
}

func (ptr *QFileInfo) SetFile(file string) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile(C.QtObjectPtr(ptr.Pointer()), C.CString(file))
	}
}

func (ptr *QFileInfo) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Suffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) Swap(other QFileInfoITF) {
	if ptr.Pointer() != nil {
		C.QFileInfo_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFileInfo(other)))
	}
}

func (ptr *QFileInfo) SymLinkTarget() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_SymLinkTarget(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileInfo) DestroyQFileInfo() {
	if ptr.Pointer() != nil {
		C.QFileInfo_DestroyQFileInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
