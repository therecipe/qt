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

type QFileInfo_ITF interface {
	QFileInfo_PTR() *QFileInfo
}

func (p *QFileInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFileInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFileInfo(ptr QFileInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileInfo_PTR().Pointer()
	}
	return nil
}

func NewQFileInfoFromPointer(ptr unsafe.Pointer) *QFileInfo {
	var n = new(QFileInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileInfo) QFileInfo_PTR() *QFileInfo {
	return ptr
}

func NewQFileInfo() *QFileInfo {
	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo())
}

func NewQFileInfo5(dir QDir_ITF, file string) *QFileInfo {
	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo5(PointerFromQDir(dir), C.CString(file)))
}

func NewQFileInfo4(file QFile_ITF) *QFileInfo {
	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo4(PointerFromQFile(file)))
}

func NewQFileInfo6(fileinfo QFileInfo_ITF) *QFileInfo {
	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo6(PointerFromQFileInfo(fileinfo)))
}

func NewQFileInfo3(file string) *QFileInfo {
	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo3(C.CString(file)))
}

func (ptr *QFileInfo) AbsoluteDir() *QDir {
	if ptr.Pointer() != nil {
		return NewQDirFromPointer(C.QFileInfo_AbsoluteDir(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) AbsoluteFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsoluteFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) AbsolutePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsolutePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) BaseName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) BundleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BundleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Caching() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_Caching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) CanonicalFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CanonicalPath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CompleteBaseName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteBaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CompleteSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Created() *QDateTime {
	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_Created(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) Dir() *QDir {
	if ptr.Pointer() != nil {
		return NewQDirFromPointer(C.QFileInfo_Dir(ptr.Pointer()))
	}
	return nil
}

func QFileInfo_Exists2(file string) bool {
	return C.QFileInfo_QFileInfo_Exists2(C.CString(file)) != 0
}

func (ptr *QFileInfo) Exists() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_Exists(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) FilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Group() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Group(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) IsAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsBundle() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsBundle(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsDir() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsDir(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsExecutable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsExecutable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsFile() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsFile(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsNativePath() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsNativePath(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRelative() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRelative(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRoot() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRoot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsSymLink() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsSymLink(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) LastModified() *QDateTime {
	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_LastModified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) LastRead() *QDateTime {
	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_LastRead(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) MakeAbsolute() bool {
	if ptr.Pointer() != nil {
		return C.QFileInfo_MakeAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) Owner() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Owner(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Refresh() {
	if ptr.Pointer() != nil {
		C.QFileInfo_Refresh(ptr.Pointer())
	}
}

func (ptr *QFileInfo) SetCaching(enable bool) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetCaching(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileInfo) SetFile3(dir QDir_ITF, file string) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile3(ptr.Pointer(), PointerFromQDir(dir), C.CString(file))
	}
}

func (ptr *QFileInfo) SetFile2(file QFile_ITF) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile2(ptr.Pointer(), PointerFromQFile(file))
	}
}

func (ptr *QFileInfo) SetFile(file string) {
	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QFileInfo) Suffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Swap(other QFileInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QFileInfo_Swap(ptr.Pointer(), PointerFromQFileInfo(other))
	}
}

func (ptr *QFileInfo) SymLinkTarget() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_SymLinkTarget(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) DestroyQFileInfo() {
	if ptr.Pointer() != nil {
		C.QFileInfo_DestroyQFileInfo(ptr.Pointer())
	}
}
